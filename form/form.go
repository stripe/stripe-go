package form

import (
	"bytes"
	"net/url"
	"reflect"
	"strconv"
)

const tagName = "form"

func AppendTo(values *RequestValues, i interface{}) {
	//v := reflect.ValueOf(i).Elem()
	reflectValue(values, reflect.ValueOf(i), nil)
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

func reflectValue(values *RequestValues, val reflect.Value, names []string) {
	// Dereference a pointer if necessary
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Do nothing if this is a nil pointer
	if !val.IsValid() {
		return
	}

	t := val.Type()

	// Also do nothing if this is the type's zero value
	if val.Interface() == reflect.Zero(t).Interface() {
		return
	}

	switch val.Kind() {
	case reflect.Array:
		panic("Arrays not handled by encoder")

	case reflect.Bool:
		values.Add(formatName(names), strconv.FormatBool(val.Bool()))

	case reflect.Float32:
		values.Add(formatName(names), strconv.FormatFloat(val.Float(), 'f', 4, 32))

	case reflect.Float64:
		values.Add(formatName(names), strconv.FormatFloat(val.Float(), 'f', 4, 64))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		values.Add(formatName(names), strconv.FormatInt(val.Int(), 10))

	case reflect.Interface:
		panic("Interfaces not handled by encoder")

	case reflect.String:
		values.Add(formatName(names), val.String())

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get(tagName)
			if tag == "" {
				continue
			}
			formName := tag
			reflectValue(values, val.Field(i), append(names, formName))
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		values.Add(formatName(names), strconv.FormatUint(val.Uint(), 10))
	}
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
