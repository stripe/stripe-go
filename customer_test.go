package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCustomer_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Customer
		err := json.Unmarshal([]byte(`"cus_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "cus_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Customer{ID: "cus_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "cus_123", v.ID)
	}
}
