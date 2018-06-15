package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestBitcoinReceiver_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v BitcoinReceiver
		err := json.Unmarshal([]byte(`"btcrcv_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "btcrcv_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := BitcoinReceiver{ID: "btcrcv_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "btcrcv_123", v.ID)
	}
}
