package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestInvoice_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Invoice
		err := json.Unmarshal([]byte(`"in_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "in_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Invoice{ID: "in_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "in_123", v.ID)
	}
}
