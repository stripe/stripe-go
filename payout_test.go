package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPayout_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Payout
		err := json.Unmarshal([]byte(`"po_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "po_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Payout{ID: "po_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "po_123", v.ID)
	}
}

func TestPayoutDestination_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v PayoutDestination
		err := json.Unmarshal([]byte(`"ba_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "ba_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		// We build the JSON object manually here because it's key that the
		// `object` field is included so that the source knows what type to
		// decode
		data := []byte(`{"id":"ba_123", "object":"bank_account"}`)

		var v PayoutDestination
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, PayoutDestinationTypeBankAccount, v.Type)

		// The destination has a field for each possible type, so the bank
		// account is located one level down
		assert.Equal(t, "ba_123", v.BankAccount.ID)
	}
}
