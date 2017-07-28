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

func AppendTo(values *RequestValues, i interface{}) {
	reflectValue(values, reflect.ValueOf(i), nil, nil)
}

func formatName(names []string) string {
	if len(names) < 1 {
		panic("Not allowed 0-length names slice")
	}

	fullName := names[0]
	for i := 1; i < len(names); i++ {
		fullName += "[" + names[i] + "]"
	}
	return fullName
}

func reflectValue(values *RequestValues, val reflect.Value, names []string, options *formOptions) {
	t := val.Type()

	// Also do nothing if this is the type's zero value
	if t.Comparable() && val.Interface() == reflect.Zero(t).Interface() {
		return
	}

	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		// formatName automatically adds square brackets, so just pass an empty
		// string into the breadcrumb trail
		arrNames := append(names, "")

		for i := 0; i < val.Len(); i++ {
			// The one exception to the above is when options have requested
			// that this array/slice be indexed. In that case we produce a hash
			// keyed with integers which the Stripe API knows how to interpret.
			if options != nil && options.IndexedArray {
				arrNames = append(names, strconv.Itoa(i))
			}

			reflectValue(values, val.Index(i), arrNames, nil)
		}

	case reflect.Bool:
		if options != nil {
			if val.Bool() {
				if options.Invert {
					values.Add(formatName(names), strconv.FormatBool(false))
				} else if options.Zero {
					values.Add(formatName(names), "0")
				}
			}
		} else {
			values.Add(formatName(names), strconv.FormatBool(val.Bool()))
		}

	case reflect.Float32:
		values.Add(formatName(names), strconv.FormatFloat(val.Float(), 'f', 4, 32))

	case reflect.Float64:
		values.Add(formatName(names), strconv.FormatFloat(val.Float(), 'f', 4, 64))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		values.Add(formatName(names), strconv.FormatInt(val.Int(), 10))

	case reflect.Interface:
		if val.IsNil() {
			return
		}
		reflectValue(values, val.Elem(), names, nil)

	case reflect.Map:
		for _, keyVal := range val.MapKeys() {
			if keyVal.Kind() != reflect.String {
				panic("Don't support serializing maps with non-string keys")
			}

			reflectValue(values, val.MapIndex(keyVal), append(names, keyVal.String()), nil)
		}

	case reflect.Ptr:
		if val.IsNil() {
			return
		}
		reflectValue(values, val.Elem(), names, options)

	case reflect.String:
		values.Add(formatName(names), val.String())

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get(tagName)
			if tag == "" {
				continue
			}
			formName, options := parseTag(tag)

			if options != nil && (options.Invert || options.Zero) &&
				val.Field(i).Type().Kind() != reflect.Bool {
				panic("Cannot specify `zero` on non-integer field")
			}

			reflectValue(values, val.Field(i), append(names, formName), options)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		values.Add(formatName(names), strconv.FormatUint(val.Uint(), 10))
	}
}

func parseTag(tag string) (string, *formOptions) {
	var options *formOptions
	parts := strings.Split(tag, ",")
	name := parts[0]

	for i := 1; i < len(parts); i++ {
		switch parts[i] {
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

// RequestValues is a collection of values that can be submitted along with a
// request that specifically allows for duplicate keys and encodes its entries
// in the same order that they were added.
type RequestValues struct {
	values []formValue
}

// Add adds a key/value tuple to the form.
func (f *RequestValues) Add(key, val string) {
	f.values = append(f.values, formValue{key, val})
}

// Encode encodes the values into “URL encoded” form ("bar=baz&foo=quux").
func (f *RequestValues) Encode() string {
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
func (f *RequestValues) Empty() bool {
	return len(f.values) == 0
}

// Set sets the first instance of a parameter for the given key to the given
// value. If no parameters exist with the key, a new one is added.
//
// Note that Set is O(n) and may be quite slow for a very large parameter list.
func (f *RequestValues) Set(key, val string) {
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
func (f *RequestValues) Get(key string) []string {
	var results []string
	for i, v := range f.values {
		if v.Key == key {
			results = append(results, f.values[i].Value)
		}
	}
	return results
}

// ToValues converts an instance of RequestValues into an instance of
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
func (f *RequestValues) ToValues() url.Values {
	values := url.Values{}
	for _, v := range f.values {
		values.Add(v.Key, v.Value)
	}
	return values
}

// A key/value tuple for use in the RequestValues type.
type formValue struct {
	Key   string
	Value string
}
