package form

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

type testStruct struct {
	Array    [3]string  `form:"array"`
	ArrayPtr *[3]string `form:"array_ptr"`

	Bool    bool  `form:"bool"`
	BoolPtr *bool `form:"bool_ptr"`

	Float32    float32  `form:"float32"`
	Float32Ptr *float32 `form:"float32_ptr"`

	Float64    float64  `form:"float64"`
	Float64Ptr *float64 `form:"float64_ptr"`

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

	Slice    []string  `form:"slice"`
	SlicePtr *[]string `form:"slice_ptr"`

	String    string  `form:"string"`
	StringPtr *string `form:"string_ptr"`

	SubStruct    testSubStruct  `form:"substruct"`
	SubStructPtr *testSubStruct `form:"substruct_ptr"`

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

type testSubStruct struct {
	SubSubStruct testSubSubStruct `form:"subsubstruct"`
}

type testSubSubStruct struct {
	String string `form:"string"`
}

func TestAppendTo(t *testing.T) {
	var boolVal bool = true

	var float32Val float32 = 1.2345

	var float64Val float64 = 1.2345

	var intVal int = 123
	var int8Val int8 = 123
	var int16Val int16 = 123
	var int32Val int32 = 123
	var int64Val int64 = 123

	var stringVal string = "123"

	var subStructVal testSubStruct = testSubStruct{
		SubSubStruct: testSubSubStruct{
			String: "123",
		},
	}

	var uintVal uint = 123
	var uint8Val uint8 = 123
	var uint16Val uint16 = 123
	var uint32Val uint32 = 123
	var uint64Val uint64 = 123

	testCases := []struct {
		field string
		data  *testStruct
		want  interface{}
	}{
		{"bool", &testStruct{Bool: boolVal}, "true"},
		{"bool_ptr", &testStruct{BoolPtr: &boolVal}, "true"},

		{"float32", &testStruct{Float32: float32Val}, "1.2345"},
		{"float32_ptr", &testStruct{Float32Ptr: &float32Val}, "1.2345"},

		{"float64", &testStruct{Float64: float64Val}, "1.2345"},
		{"float64_ptr", &testStruct{Float64Ptr: &float64Val}, "1.2345"},

		{"int", &testStruct{Int: intVal}, "123"},
		{"int_ptr", &testStruct{IntPtr: &intVal}, "123"},
		{"int8", &testStruct{Int8: int8Val}, "123"},
		{"int8_ptr", &testStruct{Int8Ptr: &int8Val}, "123"},
		{"int16", &testStruct{Int16: int16Val}, "123"},
		{"int16_ptr", &testStruct{Int16Ptr: &int16Val}, "123"},
		{"int32", &testStruct{Int32: int32Val}, "123"},
		{"int32_ptr", &testStruct{Int32Ptr: &int32Val}, "123"},
		{"int64", &testStruct{Int64: int64Val}, "123"},
		{"int64_ptr", &testStruct{Int64Ptr: &int64Val}, "123"},

		{"string", &testStruct{String: stringVal}, stringVal},
		{"string_ptr", &testStruct{StringPtr: &stringVal}, stringVal},

		{"substruct[subsubstruct][string]", &testStruct{SubStruct: subStructVal}, "123"},
		{"substruct_ptr[subsubstruct][string]", &testStruct{SubStructPtr: &subStructVal}, "123"},

		{"uint", &testStruct{Uuint: uintVal}, "123"},
		{"uint_ptr", &testStruct{UuintPtr: &uintVal}, "123"},
		{"uint8", &testStruct{Uuint8: uint8Val}, "123"},
		{"uint8_ptr", &testStruct{Uuint8Ptr: &uint8Val}, "123"},
		{"uint16", &testStruct{Uuint16: uint16Val}, "123"},
		{"uint16_ptr", &testStruct{Uuint16Ptr: &uint16Val}, "123"},
		{"uint32", &testStruct{Uuint32: uint32Val}, "123"},
		{"uint32_ptr", &testStruct{Uuint32Ptr: &uint32Val}, "123"},
		{"uint64", &testStruct{Uuint64: uint64Val}, "123"},
		{"uint64_ptr", &testStruct{Uuint64Ptr: &uint64Val}, "123"},
	}
	for _, tc := range testCases {
		t.Run(tc.field, func(t *testing.T) {
			form := &RequestValues{}
			AppendTo(form, tc.data)
			values := form.ToValues()
			//fmt.Printf("values = %+v", values)
			assert.Equal(t, tc.want, values.Get(tc.field))
		})
	}
}

func TestAppendTo_DuplicatedNames(t *testing.T) {
	arrayVal := [3]string{"1", "2", "3"}
	sliceVal := []string{"1", "2", "3"}

	testCases := []struct {
		field string
		data  *testStruct
		want  interface{}
	}{
		{"array[]", &testStruct{Array: arrayVal}, sliceVal},
		{"array_ptr[]", &testStruct{ArrayPtr: &arrayVal}, sliceVal},
		{"slice[]", &testStruct{Slice: sliceVal}, sliceVal},
		{"slice_ptr[]", &testStruct{SlicePtr: &sliceVal}, sliceVal},
	}
	for _, tc := range testCases {
		t.Run(tc.field, func(t *testing.T) {
			form := &RequestValues{}
			AppendTo(form, tc.data)
			values := form.ToValues()
			//fmt.Printf("values = %+v", values)

			// This is the only difference between this test case and the one
			// above: we used square brackets to grab a []string instead of
			// just a single value.
			assert.Equal(t, tc.want, values[tc.field])
		})
	}
}

func TestAppendTo_ZeroValues(t *testing.T) {
	form := &RequestValues{}
	data := &testStruct{}
	AppendTo(form, data)
	assert.Equal(t, &RequestValues{}, form)
}
