package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestSource_UnmarshalJSON(t *testing.T) {
	{
		// We build the JSON object manually here because it's key that the
		// `object` field is included so that the source knows what type to
		// decode
		data := []byte(`{"id":"src_123", "type":"ach_debit", "ach_debit":{"bank_name":"bar"}}`)

		var v Source
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "src_123", v.ID)
		assert.Equal(t, "ach_debit", v.Type)

		assert.Equal(t, "bar", v.ACHDebit.BankName)
	}

	// Test a degenerate case without a type key (this shouldn't happen, but
	// also shouldn't crash)
	{
		data := []byte(`{"id":"src_123", "type":"ach_debit"}`)

		var v Source
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "src_123", v.ID)
		assert.Equal(t, "ach_debit", v.Type)
	}
}
