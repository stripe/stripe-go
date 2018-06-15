package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestSKU_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v SKU
		err := json.Unmarshal([]byte(`"sku_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "sku_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := SKU{ID: "sku_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "sku_123", v.ID)
	}
}
