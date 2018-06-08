package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestRecipientTransfer_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v RecipientTransfer
		err := json.Unmarshal([]byte(`"rtr_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "rtr_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := RecipientTransfer{ID: "rtr_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "rtr_123", v.ID)
	}
}

func TestRecipientTransferDestination_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v RecipientTransferDestination
		err := json.Unmarshal([]byte(`"card_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "card_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		// We build the JSON object manually here because it's key that the
		// `object` field is included so that the source knows what type to
		// decode
		data := []byte(`{"id":"card_123", "object":"card"}`)

		var v RecipientTransferDestination
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, RecipientTransferDestinationCard, v.Type)

		// The source has a field for each possible type, so the card is
		// located one level down
		assert.Equal(t, "card_123", v.Card.ID)
	}
}
