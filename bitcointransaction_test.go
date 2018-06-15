package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestBitcoinTransaction_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v BitcoinTransaction
		err := json.Unmarshal([]byte(`"btctxn_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "btctxn_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := BitcoinTransaction{ID: "btctxn_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "btctxn_123", v.ID)
	}
}
