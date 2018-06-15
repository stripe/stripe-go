package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestApplication_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Application
		err := json.Unmarshal([]byte(`"ca_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "ca_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Application{ID: "ca_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "ca_123", v.ID)
	}
}
