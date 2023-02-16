package form

import (
	"fmt"
	"net/url"
	"sync"
	"testing"

	assert "github.com/stretchr/testify/require"
)

type benchStruct struct {
	Bool         bool              `form:"bool"`
	Ignored      string            `form:"-"`
	Int          int64             `form:"int64"`
	String       string            `form:"string"`
	SubSubStruct *testSubSubStruct `form:"subsubstruct"`
}

type testStruct struct {
	// Note that only a pointer can implement the Appender interface, so only
	// the pointer of testAppender is checked.
	Appender *testAppender `form:"appender"`

	Array    [3]string  `form:"array"`
	ArrayPtr *[3]string `form:"array_ptr"`

	Bool    bool  `form:"bool"`
	BoolPtr *bool `form:"bool_ptr"`

	Emptied bool `form:"emptied,empty"`

	Float32    float32  `form:"float32"`
	Float32Ptr *float32 `form:"float32_ptr"`

	Float32Precise    float32  `form:"float32_precise,high_precision"`
	Float32PrecisePtr *float32 `form:"float32_precise_ptr,high_precision"`

	Float64    float64  `form:"float64"`
	Float64Ptr *float64 `form:"float64_ptr"`

	Float64Precise    float64  `form:"float64_precise,high_precision"`
	Float64PrecisePtr *float64 `form:"float64_precise_ptr,high_precision"`

	Ignored string `form:"-"`

	Int      int    `form:"int"`
	IntPtr   *int   `form:"int_ptr"`
	Int8     int8   `form:"int8"`
	Int8Ptr  *int8  `form:"int8_ptr"`
	Int16    int16  `form:"int16"`
	Int16Ptr *int16 `form:"int16_ptr"`
	Int32    int32  `form:"int32"`
	Int32Ptr *int32 `form:"int32_ptr"`
	Int64    int64  `form:"int64"`
	Int64Ptr *int64 `form:"int64_ptr"`

	Map map[string]interface{} `form:"map"`

	Slice    []string  `form:"slice"`
	SlicePtr *[]string `form:"slice_ptr"`

	String    string  `form:"string"`
	StringPtr *string `form:"string_ptr"`

	SubStruct    testSubStruct  `form:"substruct"`
	SubStructPtr *testSubStruct `form:"substruct_ptr"`

	SubStructFlat    testSubStruct  `form:"*"`
	SubStructFlatPtr *testSubStruct `form:"*"`

	Uuint      uint    `form:"uint"`
	UuintPtr   *uint   `form:"uint_ptr"`
	Uuint8     uint8   `form:"uint8"`
	Uuint8Ptr  *uint8  `form:"uint8_ptr"`
	Uuint16    uint16  `form:"uint16"`
	Uuint16Ptr *uint16 `form:"uint16_ptr"`
	Uuint32    uint32  `form:"uint32"`
	Uuint32Ptr *uint32 `form:"uint32_ptr"`
	Uuint64    uint64  `form:"uint64"`
	Uuint64Ptr *uint64 `form:"uint64_ptr"`
}

type testAppender struct {
	String string `form:"-"` // Value added manually
}

func (a *testAppender) AppendTo(values *Values, keyParts []string) {
	values.Add(FormatKey(keyParts), a.String)
}

type testSubStruct struct {
	SubSubStruct testSubSubStruct `form:"subsubstruct"`
}

type testSubSubStruct struct {
	String string `form:"string"`
}

func init() {
	Strict = true
}

func BenchmarkAppendTo(b *testing.B) {
	// Disable strict mode for the duration of the benchmark (most real
	// installations should not have it turned on)
	Strict = false
	defer func() {
		Strict = true
	}()

	data := &benchStruct{
		Bool:         true,
		Ignored:      "123",
		Int:          123,
		String:       "123",
		SubSubStruct: &testSubSubStruct{String: "123"},
	}

	for i := 0; i < b.N; i++ {
		form := &Values{}
		AppendTo(form, data)
	}
}

func TestAppendTo(t *testing.T) {
	var arrayVal = [3]string{"1", "2", "3"}
	var arrayVal0 = [3]string{}

	var boolValT = true
	var boolValF = false

	var float32Val float32 = 1.2345
	var float32Val0 float32

	var float32PreciseVal float32 = 0.123456789012

	var float64Val = 1.2345
	var float64Val0 = 0.0

	var float64PreciseVal = 0.123456789012345678901234

	var intVal = 123
	var intVal0 = 0
	var int8Val int8 = 123
	var int8Val0 int8
	var int16Val int16 = 123
	var int16Val0 int16
	var int32Val int32 = 123
	var int32Val0 int32
	var int64Val int64 = 123
	var int64Val0 int64
	var sliceVal = []string{"1", "2", "3"}
	var sliceVal0 = []string{}

	var stringVal = "123"
	var stringVal0 = ""

	var subStructVal = testSubStruct{
		SubSubStruct: testSubSubStruct{
			String: "123",
		},
	}

	var uintVal uint = 123
	var uintVal0 uint
	var uint8Val uint8 = 123
	var uint8Val0 uint8
	var uint16Val uint16 = 123
	var uint16Val0 uint16
	var uint32Val uint32 = 123
	var uint32Val0 uint32
	var uint64Val uint64 = 123
	var uint64Val0 uint64

	testCases := []struct {
		field string
		data  *testStruct

		// Set to a pointer to the desired value or nil if the value shouldn't
		// be present.
		want *string
	}{
		{"appender", &testStruct{Appender: &testAppender{String: "123"}}, stringPtr("123")},

		{"array[2]", &testStruct{Array: arrayVal}, stringPtr("3")},
		{"array", &testStruct{Array: arrayVal0}, nil},

		{"array_ptr[2]", &testStruct{ArrayPtr: &arrayVal}, stringPtr("3")},
		{"array_ptr", &testStruct{ArrayPtr: &arrayVal0}, nil},

		{"bool", &testStruct{Bool: boolValT}, stringPtr("true")},
		{"bool_ptr", &testStruct{}, nil},
		{"bool_ptr", &testStruct{BoolPtr: &boolValT}, stringPtr("true")},
		{"bool_ptr", &testStruct{BoolPtr: &boolValF}, stringPtr("false")},

		{"emptied", &testStruct{Emptied: true}, stringPtr("")},

		{"float32", &testStruct{Float32: float32Val}, stringPtr("1.2345")},
		{"float32_ptr", &testStruct{Float32Ptr: &float32Val}, stringPtr("1.2345")},
		{"float32_ptr", &testStruct{Float32Ptr: &float32Val0}, stringPtr("0.0000")},
		{"float32_ptr", &testStruct{}, nil},

		// Tests float32 with high precision
		{"float32_precise", &testStruct{Float32Precise: float32PreciseVal}, stringPtr("0.12345679")},
		{"float32_precise_ptr", &testStruct{Float32PrecisePtr: &float32PreciseVal}, stringPtr("0.12345679")},
		{"float32_precise_ptr", &testStruct{Float32PrecisePtr: &float32Val0}, stringPtr("0")},
		{"float32_precise_ptr", &testStruct{}, nil},

		// The 32-bit test value we're using it already beyond the length of
		// 32-bit precision. The 64-bit value (which starts with the same
		// decimals) is well beyond it, and therefore encodes to the same
		// value.
		{"float32_precise", &testStruct{Float32Precise: float32(float64PreciseVal)}, stringPtr("0.12345679")},

		{"float64", &testStruct{Float64: float64Val}, stringPtr("1.2345")},
		{"float64_ptr", &testStruct{Float64Ptr: &float64Val}, stringPtr("1.2345")},
		{"float64_ptr", &testStruct{Float64Ptr: &float64Val0}, stringPtr("0.0000")},
		{"float64_ptr", &testStruct{}, nil},

		// Tests float64 with high precision
		{"float64_precise", &testStruct{Float64Precise: float64PreciseVal}, stringPtr("0.12345678901234568")},
		{"float64_precise_ptr", &testStruct{Float64PrecisePtr: &float64PreciseVal}, stringPtr("0.12345678901234568")},
		{"float64_precise_ptr", &testStruct{Float64PrecisePtr: &float64Val0}, stringPtr("0")},
		{"float64_precise_ptr", &testStruct{}, nil},

		{"int", &testStruct{Int: intVal}, stringPtr("123")},
		{"int_ptr", &testStruct{IntPtr: &intVal}, stringPtr("123")},
		{"int_ptr", &testStruct{IntPtr: &intVal0}, stringPtr("0")},
		{"int_ptr", &testStruct{}, nil},
		{"int8", &testStruct{Int8: int8Val}, stringPtr("123")},
		{"int8_ptr", &testStruct{Int8Ptr: &int8Val}, stringPtr("123")},
		{"int8_ptr", &testStruct{Int8Ptr: &int8Val0}, stringPtr("0")},
		{"int8_ptr", &testStruct{}, nil},
		{"int16", &testStruct{Int16: int16Val}, stringPtr("123")},
		{"int16_ptr", &testStruct{Int16Ptr: &int16Val}, stringPtr("123")},
		{"int16_ptr", &testStruct{Int16Ptr: &int16Val0}, stringPtr("0")},
		{"int16_ptr", &testStruct{}, nil},
		{"int32", &testStruct{Int32: int32Val}, stringPtr("123")},
		{"int32_ptr", &testStruct{Int32Ptr: &int32Val}, stringPtr("123")},
		{"int32_ptr", &testStruct{Int32Ptr: &int32Val0}, stringPtr("0")},
		{"int32_ptr", &testStruct{}, nil},
		{"int64", &testStruct{Int64: int64Val}, stringPtr("123")},
		{"int64_ptr", &testStruct{Int64Ptr: &int64Val}, stringPtr("123")},
		{"int64_ptr", &testStruct{Int64Ptr: &int64Val0}, stringPtr("0")},
		{"int64_ptr", &testStruct{}, nil},

		// Tests map
		{
			"map[foo]",
			&testStruct{Map: map[string]interface{}{
				"foo": "bar",
			}},
			stringPtr("bar"),
		},

		// Tests map with an empty value
		{
			"map[empty]",
			&testStruct{Map: map[string]interface{}{
				// Note that we use an empty integer instead of an empty string
				// here because `Value`''s `Get` implementation will return an
				// empty string for an unset value which means that we can't
				// differentiate between a missing and empty value. The empty
				// value for int64 is 0, so we can.
				"empty": int64(0),
			}},
			stringPtr("0"),
		},

		// Tests map nested inside of another map
		{
			"map[foo][bar]",
			&testStruct{Map: map[string]interface{}{
				"foo": map[string]interface{}{"bar": "baz"},
			}},
			stringPtr("baz"),
		},

		{"slice[2]", &testStruct{Slice: sliceVal}, stringPtr("3")},
		{"slice", &testStruct{Slice: []string{}}, stringPtr("")},
		{"slice", &testStruct{Slice: nil}, nil},

		{"slice_ptr[2]", &testStruct{SlicePtr: &sliceVal}, stringPtr("3")},

		// A slice pointer given an explicit but empty slice should encode to
		// an empty string (this tells the Stripe API to "zero" the array).
		{"slice_ptr", &testStruct{SlicePtr: &sliceVal0}, stringPtr("")},

		{"slice_ptr", &testStruct{SlicePtr: nil}, nil},

		{"string", &testStruct{String: stringVal}, &stringVal},
		{"string_ptr", &testStruct{StringPtr: &stringVal}, &stringVal},
		{"string_ptr", &testStruct{StringPtr: &stringVal0}, &stringVal0},
		{"string_ptr", &testStruct{}, nil},

		{"substruct[subsubstruct][string]", &testStruct{SubStruct: subStructVal}, stringPtr("123")},
		{"substruct_ptr[subsubstruct][string]", &testStruct{SubStructPtr: &subStructVal}, stringPtr("123")},

		{"subsubstruct[string]", &testStruct{SubStructFlat: subStructVal}, stringPtr("123")},
		{"subsubstruct[string]", &testStruct{SubStructFlatPtr: &subStructVal}, stringPtr("123")},

		{"uint", &testStruct{Uuint: uintVal}, stringPtr("123")},
		{"uint_ptr", &testStruct{UuintPtr: &uintVal}, stringPtr("123")},
		{"uint_ptr", &testStruct{UuintPtr: &uintVal0}, stringPtr("0")},
		{"uint_ptr", &testStruct{}, nil},
		{"uint8", &testStruct{Uuint8: uint8Val}, stringPtr("123")},
		{"uint8_ptr", &testStruct{Uuint8Ptr: &uint8Val}, stringPtr("123")},
		{"uint8_ptr", &testStruct{Uuint8Ptr: &uint8Val0}, stringPtr("0")},
		{"uint8_ptr", &testStruct{}, nil},
		{"uint16", &testStruct{Uuint16: uint16Val}, stringPtr("123")},
		{"uint16_ptr", &testStruct{Uuint16Ptr: &uint16Val}, stringPtr("123")},
		{"uint16_ptr", &testStruct{Uuint16Ptr: &uint16Val0}, stringPtr("0")},
		{"uint16_ptr", &testStruct{}, nil},
		{"uint32", &testStruct{Uuint32: uint32Val}, stringPtr("123")},
		{"uint32_ptr", &testStruct{Uuint32Ptr: &uint32Val}, stringPtr("123")},
		{"uint32_ptr", &testStruct{Uuint32Ptr: &uint32Val0}, stringPtr("0")},
		{"uint32_ptr", &testStruct{}, nil},
		{"uint64", &testStruct{Uuint64: uint64Val}, stringPtr("123")},
		{"uint64_ptr", &testStruct{Uuint64Ptr: &uint64Val}, stringPtr("123")},
		{"uint64_ptr", &testStruct{Uuint64Ptr: &uint64Val0}, stringPtr("0")},
		{"uint64_ptr", &testStruct{}, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.field, func(t *testing.T) {
			form := &Values{}
			AppendTo(form, tc.data)
			values := form.ToValues()
			t.Logf("values: %+v encoded: \"%+v\"", values, form.Encode())

			actuals, ok := values[tc.field]
			if ok {
				if tc.want == nil {
					assert.Fail(t, fmt.Sprintf(
						"Value not expected for %s, but got '%+v'",
						tc.field, actuals))
				}
			} else {
				if tc.want == nil {
					// Got no value and expected no value
					return
				}

				assert.Fail(t, fmt.Sprintf(
					"No value present for %s, but expected '%+v'",
					tc.field, *tc.want))
			}

			if len(actuals) > 1 {
				assert.Fail(t, fmt.Sprintf(
					"Got more than one value for %s: %+v",
					tc.field, actuals))
			}

			actual := actuals[0]
			assert.Equal(t, *tc.want, actual)
		})
	}
}

func TestAppendTo_IgnoredFields(t *testing.T) {
	form := &Values{}
	data := &testStruct{Ignored: "value"}
	AppendTo(form, data)
	assert.Equal(t, &Values{}, form)
}

func TestAppendTo_ZeroValues(t *testing.T) {
	form := &Values{}
	data := &testStruct{}
	AppendTo(form, data)
	assert.Equal(t, &Values{}, form)
}

func TestAppendToPrefixed(t *testing.T) {
	form := &Values{}
	data := &testStruct{String: "foo"}
	AppendToPrefixed(form, data, []string{"prefix"})
	assert.Equal(t, []string{"foo"}, form.Get("prefix[string]"))
}

func TestEncode(t *testing.T) {
	form := &Values{}
	form.Add("foo", "bar")
	form.Add("foo[bar]", "baz")
	assert.Equal(t, "foo=bar&foo[bar]=baz", form.Encode())
}

func TestEncodeDeterministically(t *testing.T) {
	data := map[string]string{
		"first":  "foo",
		"second": "bar",
	}
	initialForm := &Values{}
	AppendTo(initialForm, data)
	encoded := initialForm.Encode()

	for i := 0; i < 100; i++ {
		form := &Values{}
		AppendTo(form, data)
		assert.Equal(t, encoded, form.Encode(), "iteration %d", i)
	}
}

func TestEncodeMapNonStringKey(t *testing.T) {
	// Disable strict mode for this test, so the non-string key is ignored.
	Strict = false
	defer func() {
		Strict = true
	}()

	form := &Values{}
	assert.NotPanics(t, func() {
		AppendTo(form, map[int]string{1: "foo"})
	})
	assert.Len(t, form.values, 0)
}

func TestFormatKey(t *testing.T) {
	assert.Equal(t, "param", FormatKey([]string{"param"}))
	assert.Equal(t, "param[key]", FormatKey([]string{"param", "key"}))
	assert.Equal(t, "param[key][]", FormatKey([]string{"param", "key", ""}))
	assert.Equal(t, "param[key][0]", FormatKey([]string{"param", "key", "0"}))
}

// The encoder uses a type cache for speed. This test is designed to help
// verify that concurrent access to it is safe.
//
// I had good success in reproducing concurrent access errors on my computer
// using this test, but given that we're just throwing lots of Goroutines out
// there and hoping for an error, your mileage may vary depending on your
// system. It may be necessary to increase the value of `n` or introduce other
// constructs (although hopefully this package will stay concurrency-safe).
func TestCacheConcurrency(t *testing.T) {
	// Clear out anything in the existing cache
	encoderCache.m = nil
	structCache.m = nil

	var wg sync.WaitGroup
	n := 10
	val := &testStruct{String: "123"}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			form := &Values{}
			AppendTo(form, val)
			wg.Done()
		}()
	}

	wg.Wait()
}

func TestParseTag(t *testing.T) {
	// Disable strict mode for the duration of this test so that we can test
	// some malformed tags
	Strict = false
	defer func() {
		Strict = true
	}()

	testCases := []struct {
		tag         string
		wantName    string
		wantOptions *formOptions
	}{
		{"id", "id", nil},
		{"id,empty", "id", &formOptions{Empty: true}},

		// invalid invocations
		{"id,", "id", nil},
		{"id,,", "id", nil},
		{"id,foo", "id", nil},
		{"id,foo=bar", "id", nil},
	}
	for _, tc := range testCases {
		t.Run(tc.tag, func(t *testing.T) {
			name, options := parseTag(tc.tag)
			assert.Equal(t, tc.wantName, name)
			assert.Equal(t, tc.wantOptions, options)
		})
	}
}

func TestValues(t *testing.T) {
	values := &Values{}

	assert.Equal(t, "", values.Encode())
	assert.True(t, values.Empty())

	values = &Values{}
	values.Add("foo", "bar")

	assert.Equal(t, "foo=bar", values.Encode())
	assert.False(t, values.Empty())
	assert.Equal(t, []string{"bar"}, values.Get("foo"))

	values = &Values{}
	values.Add("foo", "bar")
	values.Add("foo", "bar")
	values.Add("baz", "bar")

	assert.Equal(t, "foo=bar&foo=bar&baz=bar", values.Encode())
	assert.Equal(t, []string{"bar", "bar"}, values.Get("foo"))
	assert.Equal(t, []string{"bar"}, values.Get("baz"))

	values.Set("foo", "firstbar")

	assert.Equal(t, "foo=firstbar&foo=bar&baz=bar", values.Encode())
	assert.Equal(t, []string{"firstbar", "bar"}, values.Get("foo"))
	assert.Equal(t, []string{"bar"}, values.Get("baz"))

	values.Set("new", "appended")

	assert.Equal(t, "foo=firstbar&foo=bar&baz=bar&new=appended", values.Encode())

	assert.Equal(t, url.Values{
		"baz": {"bar"},
		"foo": {"firstbar", "bar"},
		"new": {"appended"},
	}, values.ToValues())
	assert.Equal(t, []string{"appended"}, values.Get("new"))

	assert.Nil(t, values.Get("boguskey"))
}

//
// Private functions
//

// Trivial helper to get is a string pointer from a string (because you're not
// allowed to take the address of a literal directly).
func stringPtr(s string) *string {
	return &s
}
