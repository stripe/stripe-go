package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestDispute_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Dispute
		err := json.Unmarshal([]byte(`"dp_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "dp_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Dispute{ID: "dp_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "dp_123", v.ID)
	}
}
