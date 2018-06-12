package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCoupon_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Coupon
		err := json.Unmarshal([]byte(`"25OFF"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "25OFF", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Coupon{ID: "25OFF"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "25OFF", v.ID)
	}
}
