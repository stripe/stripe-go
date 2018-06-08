package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCard_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Card
		err := json.Unmarshal([]byte(`"card_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "card_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Card{ID: "card_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "card_123", v.ID)
	}
}
