package form

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
)

type testStruct struct {
	String    string  `form:"string"`
	StringPtr *string `form:"string_ptr"`
}

func TestAppendTo(t *testing.T) {
	stringVal := "123"

	testCases := []struct {
		field string
		data  *testStruct
		want  interface{}
	}{
		{"string", &testStruct{String: stringVal}, stringVal},
		{"string_ptr", &testStruct{StringPtr: &stringVal}, stringVal},
	}
	for _, tc := range testCases {
		t.Run(tc.field, func(t *testing.T) {
			form := &RequestValues{}
			AppendTo(form, tc.data)
			assert.Equal(t, tc.want, form.ToValues().Get(tc.field))
		})
	}
}
