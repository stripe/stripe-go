package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestRefund_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Refund
		err := json.Unmarshal([]byte(`"re_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "re_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Refund{ID: "re_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "re_123", v.ID)
	}
}
