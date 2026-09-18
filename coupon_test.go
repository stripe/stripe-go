package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v86/form"
)

func TestCouponScriptConfiguration(t *testing.T) {
	var coupon Coupon
	err := json.Unmarshal([]byte(`{"script":{"configuration":{"enabled":true,"nested":{"key":"value"}}}}`), &coupon)
	assert.NoError(t, err)
	assert.Equal(t, true, coupon.Script.Configuration["enabled"])
	assert.Equal(t, "value", coupon.Script.Configuration["nested"].(map[string]interface{})["key"])

	for _, params := range []interface{}{
		&CouponParams{Script: &CouponScriptParams{Configuration: map[string]any{"enabled": true, "nested": map[string]any{"key": "value"}}}},
		&CouponCreateParams{Script: &CouponCreateScriptParams{Configuration: map[string]any{"enabled": true, "nested": map[string]any{"key": "value"}}}},
	} {
		values := &form.Values{}
		form.AppendTo(values, params)
		assert.Equal(t, []string{"true"}, values.Get("script[configuration][enabled]"))
		assert.Equal(t, []string{"value"}, values.Get("script[configuration][nested][key]"))
	}
}

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
