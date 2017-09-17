package form

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const tagName = "form"

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

type encoderFunc func(values *Values, v reflect.Value, keyParts []string, options *formOptions)

// field represents a single field found in a struct. It caches information
// about that field so that we can make encoding faster.
type field struct {
	formName   string
	index      int
	isAppender bool
	isPtr      bool
	options    *formOptions
}

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

type structEncoder struct {
	fields    []*field
	fieldEncs []encoderFunc
}

func (se *structEncoder) encode(values *Values, v reflect.Value, keyParts []string, _ *formOptions) {
	for i, f := range se.fields {
		var fieldKeyParts []string
		fieldV := v.Field(f.index)

		// The wildcard on a form tag is a "special" value: it indicates a
		// struct field that we should recurse into, but for which no part
		// should be added to the key parts, meaning that its own subfields
		// will be named at the same level as with the fields of the
		// current structure.
		if f.formName == "*" {
			fieldKeyParts = keyParts
		} else {
			fieldKeyParts = append(keyParts, f.formName)
		}

		se.fieldEncs[i](values, fieldV, fieldKeyParts, f.options)
		if f.isAppender && (!f.isPtr || !fieldV.IsNil()) {
			fieldV.Interface().(Appender).AppendTo(values, fieldKeyParts)
		}
	}
}

// ---

// Strict enables strict mode wherein the package will panic on an AppendTo
// function if it finds that a tag string was malformed.
var Strict = false

var encoderCache struct {
	mu    sync.Mutex   // used only by writers
	value atomic.Value // map[reflect.Type]encoderFunc
}

var structCache struct {
	mu    sync.Mutex   // used only by writers
	value atomic.Value // map[reflect.Type]*structEncoder
}

// AppendTo uses reflection to form encode into the given values collection
// based off the form tags that it defines.
func AppendTo(values *Values, i interface{}) {
	reflectValue(values, reflect.ValueOf(i), nil)
}

// AppendToPrefixed is the same as AppendTo, but it allows a slice of key parts
// to be specified to prefix the form values.
//
// I was hoping not to have to expose this function, but I ended up needing it
// for recipients. Recipients is going away, and when it does, we can probably
// remove it again.
func AppendToPrefixed(values *Values, i interface{}, keyParts []string) {
	reflectValue(values, reflect.ValueOf(i), keyParts)
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

// ---

func boolEncoder(values *Values, v reflect.Value, keyParts []string, options *formOptions) {
	if !v.Bool() {
		return
	}

	if options != nil {
		if v.Bool() {
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
		values.Add(FormatKey(keyParts), strconv.FormatBool(v.Bool()))
	}
}

func buildArrayOrSliceEncoder(t reflect.Type) encoderFunc {
	// Gets an encoder for the type that the array or slice will hold
	elemF := getCachedOrBuildTypeEncoder(t.Elem())

	return func(values *Values, v reflect.Value, keyParts []string, options *formOptions) {
		// FormatKey automatically adds square brackets, so just pass an empty
		// string into the breadcrumb trail
		arrNames := append(keyParts, "")

		for i := 0; i < v.Len(); i++ {
			// The one exception to the above is when options have requested
			// that this array/slice be indexed. In that case we produce a hash
			// keyed with integers which the Stripe API knows how to interpret.
			if options != nil && options.IndexedArray {
				arrNames = append(keyParts, strconv.Itoa(i))
			}

			indexV := v.Index(i)
			elemF(values, indexV, arrNames, nil)

			if isAppender(indexV.Type()) && !indexV.IsNil() {
				v.Interface().(Appender).AppendTo(values, arrNames)
			}
		}
	}
}

func buildPtrEncoder(t reflect.Type) encoderFunc {
	// Gets an encoder for the type that the pointer wraps
	elemF := getCachedOrBuildTypeEncoder(t.Elem())

	return func(values *Values, v reflect.Value, keyParts []string, options *formOptions) {
		if v.IsNil() {
			return
		}
		elemF(values, v.Elem(), keyParts, options)
	}
}

func buildStructEncoder(t reflect.Type) encoderFunc {
	se := getCachedOrBuildStructEncoder(t)
	return se.encode
}

func float32Encoder(values *Values, v reflect.Value, keyParts []string, _ *formOptions) {
	if v.Float() == 0.0 {
		return
	}
	values.Add(FormatKey(keyParts), strconv.FormatFloat(v.Float(), 'f', 4, 32))
}

func float64Encoder(values *Values, v reflect.Value, keyParts []string, _ *formOptions) {
	if v.Float() == 0.0 {
		return
	}
	values.Add(FormatKey(keyParts), strconv.FormatFloat(v.Float(), 'f', 4, 64))
}

func getCachedOrBuildStructEncoder(t reflect.Type) *structEncoder {
	m, _ := structCache.value.Load().(map[reflect.Type]*structEncoder)
	f := m[t]
	if f != nil {
		return f
	}

	// We do the work to get the encoder without holding a lock. This could
	// result in duplicate work, but it will help us avoid a deadlock. Encoders
	// may be built and stored recursively in the cases of something like an
	// array or slice, so we need to make sure that this function is properly
	// re-entrant.
	f = makeStructEncoder(t)

	structCache.mu.Lock()
	defer structCache.mu.Unlock()

	m, _ = structCache.value.Load().(map[reflect.Type]*structEncoder)
	newM := m
	if newM == nil {
		newM = make(map[reflect.Type]*structEncoder)
		structCache.value.Store(newM)
	}
	newM[t] = f

	return f
}

// getCachedOrBuildTypeEncoder tries to get an encoderFunc for the type from
// the cache, and falls back to building one if there wasn't a cached one
// available. If an encoder is built, it's stored back to the cache.
func getCachedOrBuildTypeEncoder(t reflect.Type) encoderFunc {
	m, _ := encoderCache.value.Load().(map[reflect.Type]encoderFunc)
	f := m[t]
	if f != nil {
		return f
	}

	// We do the work to get the encoder without holding a lock. This could
	// result in duplicate work, but it will help us avoid a deadlock. Encoders
	// may be built and stored recursively in the cases of something like an
	// array or slice, so we need to make sure that this function is properly
	// re-entrant.
	f = makeTypeEncoder(t)

	encoderCache.mu.Lock()
	defer encoderCache.mu.Unlock()

	m, _ = encoderCache.value.Load().(map[reflect.Type]encoderFunc)
	newM := m
	if newM == nil {
		newM = make(map[reflect.Type]encoderFunc)
		encoderCache.value.Store(newM)
	}
	newM[t] = f

	return f
}

func intEncoder(values *Values, v reflect.Value, keyParts []string, _ *formOptions) {
	if v.Int() == 0 {
		return
	}
	values.Add(FormatKey(keyParts), strconv.FormatInt(v.Int(), 10))
}

func interfaceEncoder(values *Values, v reflect.Value, keyParts []string, _ *formOptions) {
	if v.IsNil() {
		return
	}
	reflectValue(values, v.Elem(), keyParts)
}

func isAppender(t reflect.Type) bool {
	return t.Implements(reflect.TypeOf((*Appender)(nil)).Elem())
}

func mapEncoder(values *Values, v reflect.Value, keyParts []string, _ *formOptions) {
	for _, keyVal := range v.MapKeys() {
		if Strict && keyVal.Kind() != reflect.String {
			panic("Don't support serializing maps with non-string keys")
		}

		reflectValue(values, v.MapIndex(keyVal), append(keyParts, keyVal.String()))
	}
}

func stringEncoder(values *Values, v reflect.Value, keyParts []string, _ *formOptions) {
	if v.String() == "" {
		return
	}
	values.Add(FormatKey(keyParts), v.String())
}

func uintEncoder(values *Values, v reflect.Value, keyParts []string, _ *formOptions) {
	if v.Uint() == 0 {
		return
	}
	values.Add(FormatKey(keyParts), strconv.FormatUint(v.Uint(), 10))
}

func reflectValue(values *Values, v reflect.Value, keyParts []string) {
	t := v.Type()

	f := getCachedOrBuildTypeEncoder(t)
	if f != nil {
		f(values, v, keyParts, nil)
	}

	if isAppender(t) {
		v.Interface().(Appender).AppendTo(values, keyParts)
	}
}

func makeStructEncoder(t reflect.Type) *structEncoder {
	// Don't specify capacity because we don't know how many fields are tagged with
	// `form`
	se := &structEncoder{}

	for i := 0; i < t.NumField(); i++ {
		reflectField := t.Field(i)
		tag := reflectField.Tag.Get(tagName)
		if Strict && tag == "" {
			panic(fmt.Sprintf(
				"All fields in structs to be form-encoded must have `form` tag; on: %s/%s "+
					"(hint: use an explicit `form:\"-\"` if the field should not be encoded",
				t.Name(), reflectField.Name,
			))
		}

		formName, options := parseTag(tag)

		// Like with encoding/json, a hyphen is an explicit way of saying
		// that this field should not be encoded
		if formName == "-" {
			continue
		}

		if Strict && options != nil &&
			(options.Empty || options.Invert || options.Zero) &&
			reflectField.Type.Kind() != reflect.Bool {

			panic(fmt.Sprintf(
				"Cannot specify `empty`, `invert`, or `zero` for non-boolean field; on: %s/%s",
				t.Name(), reflectField.Name,
			))
		}

		var isPtr bool
		if reflectField.Type.Kind() == reflect.Ptr {
			isPtr = true
		}

		se.fields = append(se.fields, &field{
			formName:   formName,
			index:      i,
			isAppender: isAppender(reflectField.Type),
			isPtr:      isPtr,
			options:    options,
		})
		se.fieldEncs = append(se.fieldEncs,
			getCachedOrBuildTypeEncoder(reflectField.Type))
	}

	return se
}

func makeTypeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		return buildArrayOrSliceEncoder(t)

	case reflect.Bool:
		return boolEncoder

	case reflect.Float32:
		return float32Encoder

	case reflect.Float64:
		return float64Encoder

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intEncoder

	case reflect.Interface:
		return interfaceEncoder

	case reflect.Map:
		return mapEncoder

	case reflect.Ptr:
		return buildPtrEncoder(t)

	case reflect.String:
		return stringEncoder

	case reflect.Struct:
		return buildStructEncoder(t)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintEncoder
	}

	return nil
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

		default:
			if Strict {
				part := parts[i]
				if part == "" {
					part = "(empty)"
				}
				panic(fmt.Sprintf("Don't know how to handle form tag part: %s (tag: %s)",
					part, tag))
			}
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
