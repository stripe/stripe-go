package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestFileLink_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v FileLink
		err := json.Unmarshal([]byte(`"link_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "link_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := FileLink{ID: "link_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "link_123", v.ID)
	}
}
