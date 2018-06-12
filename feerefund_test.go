package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestFeeRefund_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v FeeRefund
		err := json.Unmarshal([]byte(`"fr_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "fr_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := FeeRefund{ID: "fr_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "fr_123", v.ID)
	}
}
