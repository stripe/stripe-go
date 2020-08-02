package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPromotionCode_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v PromotionCode
		err := json.Unmarshal([]byte(`"promo_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "promo_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := PromotionCode{ID: "promo_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "promo_123", v.ID)
	}
}
