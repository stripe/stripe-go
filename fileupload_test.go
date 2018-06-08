package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestFileUpload_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v FileUpload
		err := json.Unmarshal([]byte(`"file_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "file_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := FileUpload{ID: "file_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "file_123", v.ID)
	}
}
