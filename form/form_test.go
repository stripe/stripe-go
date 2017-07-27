package form

import (
	"fmt"
	"testing"
)

type testStruct struct {
	StringField    string  `form:"string_field"`
	StringPtrField *string `form:"string_ptr_field"`
}

func TestAppendTo(t *testing.T) {
	stringVal := "123"

	values := &RequestValues{}
	s := &testStruct{
		StringField:    stringVal,
		StringPtrField: &stringVal,
	}
	AppendTo(values, s)
	fmt.Printf("values = %+v", values)
}
