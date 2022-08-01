package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestTransferReversal_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v TransferReversal
		err := json.Unmarshal([]byte(`"trr_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "trr_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := TransferReversal{ID: "trr_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "trr_123", v.ID)
	}
}
