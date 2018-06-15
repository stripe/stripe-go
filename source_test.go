package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestSourceObjectParams_AppendTo(t *testing.T) {
	// Test to make sure that TypeData makes it to the root object level of
	// encoding
	{
		params := &SourceObjectParams{
			Type: String("source_type"),
			TypeData: map[string]string{
				"foo": "bar",
			},
		}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"bar"}, body.Get("source_type[foo]"))
	}
}

func TestSource_UnmarshalJSON(t *testing.T) {
	{
		// We build the JSON object manually here because it's key that the
		// `object` field is included so that the source knows what type to
		// decode
		data := []byte(`{"id":"src_123", "type":"ach", "ach":{"foo":"bar"}}`)

		var v Source
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "src_123", v.ID)
		assert.Equal(t, "ach", v.Type)

		// The source data is extracted to the special TypeData field
		assert.Equal(t, "bar", v.TypeData["foo"])
	}

	// Test a degenerate case without a type key (this shouldn't happen, but
	// also shouldn't crash)
	{
		data := []byte(`{"id":"src_123", "type":"ach"}`)

		var v Source
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "src_123", v.ID)
		assert.Equal(t, "ach", v.Type)
	}
}
