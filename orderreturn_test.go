package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestOrderReturn_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v OrderReturn
		err := json.Unmarshal([]byte(`"orret_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "orret_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := OrderReturn{ID: "orret_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "orret_123", v.ID)
	}
}
