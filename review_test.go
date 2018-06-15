package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestReview_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Review
		err := json.Unmarshal([]byte(`"prv_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "prv_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Review{ID: "prv_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "prv_123", v.ID)
	}
}
