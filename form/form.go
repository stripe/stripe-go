package form

import (
	"bytes"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const tagName = "form"

type formOptions struct {
	// IndexedArray indicates that contrary to standard "Rack-style" form
	// encoding, array items should be index like `arr[0]=...&arr[1]=...`
	// (normally it'd be `arr[]=...`).
	IndexedArray bool

	// Empty indicates that a field's value should be emptied in that its value
	// should be an empty string. It's used to workaround the fact that an
	// empty string is a string's zero value and wouldn't normally be encoded.
	Empty bool

	// Invert indicates that a boolean field's value should be inverted. False
	// is the zero value for a boolean so it's convention in the library to
	// specify a `No*` field to allow a false to be passed to the API. These
	// fields should be annotated with `invert`.
	Invert bool

	// Zero indicates a field that's specifically defined to workaround the
	// fact because 0 is the "zero value" of all int/float types, we can't
	// properly encode an explicit 0. It indicates that an explicit zero should
	// be sent.
	Zero bool
}

// Appender is the interface implemented by types that can append themselves to
// a collection of form values.
//
// This is usually something that shouldn't be used, but is needed in a few
// places where authors deviated from norms while implementing various
// parameters.
type Appender interface {
	// AppendTo is invoked by the form package on any types found to implement
	// Appender so that they have a chance to encode themselves. Note that
	// AppendTo is called in addition to normal encoding, so other form tags on
	// the struct are still fair game.
	AppendTo(values *Values, keyParts []string)
}

func AppendTo(values *Values, i interface{}) {
	reflectValue(values, reflect.ValueOf(i), nil, nil)
}

// FormatKey takes a series of key parts that may be parameter keyParts, map keys,
// or array indices and unifies them into a single key suitable for Stripe's
// style of form encoding.
func FormatKey(parts []string) string {
	if len(parts) < 1 {
		panic("Not allowed 0-length parts slice")
	}

	key := parts[0]
	for i := 1; i < len(parts); i++ {
		key += "[" + parts[i] + "]"
	}
	return key
}

func reflectValue(values *Values, val reflect.Value, keyParts []string, options *formOptions) {
	t := val.Type()

	// Also do nothing if this is the type's zero value
	if t.Comparable() && val.Interface() == reflect.Zero(t).Interface() {
		return
	}

	// Special case for when a type has implemented special logic to append
	// itself to a form.
	if t.Implements(reflect.TypeOf((*Appender)(nil)).Elem()) {
		val.Interface().(Appender).AppendTo(values, keyParts)
	}

	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		// FormatKey automatically adds square brackets, so just pass an empty
		// string into the breadcrumb trail
		arrNames := append(keyParts, "")

		for i := 0; i < val.Len(); i++ {
			// The one exception to the above is when options have requested
			// that this array/slice be indexed. In that case we produce a hash
			// keyed with integers which the Stripe API knows how to interpret.
			if options != nil && options.IndexedArray {
				arrNames = append(keyParts, strconv.Itoa(i))
			}

			reflectValue(values, val.Index(i), arrNames, nil)
		}

	case reflect.Bool:
		if options != nil {
			if val.Bool() {
				switch {
				case options.Empty:
					values.Add(FormatKey(keyParts), "")
				case options.Invert:
					values.Add(FormatKey(keyParts), strconv.FormatBool(false))
				case options.Zero:
					values.Add(FormatKey(keyParts), "0")
				}
			}
		} else {
			values.Add(FormatKey(keyParts), strconv.FormatBool(val.Bool()))
		}

	case reflect.Float32:
		values.Add(FormatKey(keyParts), strconv.FormatFloat(val.Float(), 'f', 4, 32))

	case reflect.Float64:
		values.Add(FormatKey(keyParts), strconv.FormatFloat(val.Float(), 'f', 4, 64))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		values.Add(FormatKey(keyParts), strconv.FormatInt(val.Int(), 10))

	case reflect.Interface:
		if val.IsNil() {
			return
		}
		reflectValue(values, val.Elem(), keyParts, nil)

	case reflect.Map:
		for _, keyVal := range val.MapKeys() {
			if keyVal.Kind() != reflect.String {
				panic("Don't support serializing maps with non-string keys")
			}

			reflectValue(values, val.MapIndex(keyVal), append(keyParts, keyVal.String()), nil)
		}

	case reflect.Ptr:
		if val.IsNil() {
			return
		}
		reflectValue(values, val.Elem(), keyParts, options)

	case reflect.String:
		values.Add(FormatKey(keyParts), val.String())

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get(tagName)
			if tag == "" {
				continue
			}
			formName, options := parseTag(tag)

			if options != nil &&
				(options.Empty || options.Invert || options.Zero) &&
				val.Field(i).Type().Kind() != reflect.Bool {

				panic("Cannot specify `zero` on non-integer field")
			}

			// The wildcard on a form tag is a "special" value: it indicates a
			// struct field that we should recurse into, but for which no part
			// should be added to the key parts, meaning that its own subfields
			// will be named at the same level as with the fields of the
			// current structure.
			if formName == "*" {
				reflectValue(values, val.Field(i), keyParts, options)
			} else {
				reflectValue(values, val.Field(i), append(keyParts, formName), options)
			}
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		values.Add(FormatKey(keyParts), strconv.FormatUint(val.Uint(), 10))
	}
}

func parseTag(tag string) (string, *formOptions) {
	var options *formOptions
	parts := strings.Split(tag, ",")
	name := parts[0]

	for i := 1; i < len(parts); i++ {
		switch parts[i] {
		case "empty":
			if options == nil {
				options = &formOptions{}
			}
			options.Empty = true

		case "indexed":
			if options == nil {
				options = &formOptions{}
			}
			options.IndexedArray = true

		case "invert":
			if options == nil {
				options = &formOptions{}
			}
			options.Invert = true

		case "zero":
			if options == nil {
				options = &formOptions{}
			}
			options.Zero = true
		}
	}

	return name, options
}

// ---

// Values is a collection of values that can be submitted along with a
// request that specifically allows for duplicate keys and encodes its entries
// in the same order that they were added.
type Values struct {
	values []formValue
}

// Add adds a key/value tuple to the form.
func (f *Values) Add(key, val string) {
	f.values = append(f.values, formValue{key, val})
}

// Encode encodes the values into “URL encoded” form ("bar=baz&foo=quux").
func (f *Values) Encode() string {
	var buf bytes.Buffer
	for _, v := range f.values {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(v.Key))
		buf.WriteString("=")
		buf.WriteString(url.QueryEscape(v.Value))
	}
	return buf.String()
}

// Empty returns true if no parameters have been set.
func (f *Values) Empty() bool {
	return len(f.values) == 0
}

// Set sets the first instance of a parameter for the given key to the given
// value. If no parameters exist with the key, a new one is added.
//
// Note that Set is O(n) and may be quite slow for a very large parameter list.
func (f *Values) Set(key, val string) {
	for i, v := range f.values {
		if v.Key == key {
			f.values[i].Value = val
			return
		}
	}

	f.Add(key, val)
}

// Get retrieves the list of values for the given key.  If no values exist
// for the key, nil will be returned.
//
// Note that Get is O(n) and may be quite slow for a very large parameter list.
func (f *Values) Get(key string) []string {
	var results []string
	for i, v := range f.values {
		if v.Key == key {
			results = append(results, f.values[i].Value)
		}
	}
	return results
}

// ToValues converts an instance of Values into an instance of
// url.Values. This can be useful in cases where it's useful to make an
// unordered comparison of two sets of request values.
//
// Note that url.Values is incapable of representing certain Rack form types in
// a cohesive way. For example, an array of maps in Rack is encoded with a
// string like:
//
//     arr[][foo]=foo0&arr[][bar]=bar0&arr[][foo]=foo1&arr[][bar]=bar1
//
// Because url.Values is a map, values will be handled in a way that's grouped
// by their key instead of in the order they were added. Therefore the above
// may by encoded to something like (maps are unordered so the actual result is
// somewhat non-deterministic):
//
//     arr[][foo]=foo0&arr[][foo]=foo1&arr[][bar]=bar0&arr[][bar]=bar1
//
// And thus result in an incorrect request to Stripe.
func (f *Values) ToValues() url.Values {
	values := url.Values{}
	for _, v := range f.values {
		values.Add(v.Key, v.Value)
	}
	return values
}

// A key/value tuple for use in the Values type.
type formValue struct {
	Key   string
	Value string
}
