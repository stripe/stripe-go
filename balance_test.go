package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestBalanceTransaction_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v BalanceTransaction
		err := json.Unmarshal([]byte(`"bt_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "bt_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := BalanceTransaction{ID: "bt_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "bt_123", v.ID)
	}
}

func TestBalanceTransactionSource_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v BalanceTransactionSource
		err := json.Unmarshal([]byte(`"ch_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "ch_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		// We build the JSON object manually here because it's key that the
		// `object` field is included so that the source knows what type to
		// decode
		data := []byte(`{"id":"ch_123", "object":"charge"}`)

		var v BalanceTransactionSource
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, BalanceTransactionSourceTypeCharge, v.Type)

		// The source has a field for each possible type, so the charge is
		// located one level down
		assert.Equal(t, "ch_123", v.Charge.ID)
	}
}
