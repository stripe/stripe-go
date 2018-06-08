package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestInvoiceItem_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v InvoiceItem
		err := json.Unmarshal([]byte(`"ii_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "ii_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := InvoiceItem{ID: "ii_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "ii_123", v.ID)
	}
}
