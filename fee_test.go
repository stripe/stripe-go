package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestApplicationFee_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v ApplicationFee
		err := json.Unmarshal([]byte(`"fee_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "fee_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := ApplicationFee{ID: "fee_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "fee_123", v.ID)
	}
}
