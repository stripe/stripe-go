package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestProduct_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Product
		err := json.Unmarshal([]byte(`"prod_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "prod_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Product{ID: "prod_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "prod_123", v.ID)
	}
}
