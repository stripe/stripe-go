package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestSourceTransaction_UnmarshalJSON(t *testing.T) {
	{
		// We build the JSON object manually here because it's key that the
		// `object` field is included so that the source knows what type to
		// decode
		data := []byte(`{"type":"ach_credit_transfer", "ach_credit_transfer":{"routing_number":"bar"}}`)

		var v SourceTransaction
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "ach_credit_transfer", v.Type)

		assert.Equal(t, "bar", v.ACHCreditTransfer.RoutingNumber)
	}

	// Test a degenerate case without a type key (this shouldn't happen, but
	// also shouldn't crash)
	{
		data := []byte(`{"type":"ach"}`)

		var v SourceTransaction
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "ach", v.Type)
	}
}
