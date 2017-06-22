package coupon

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestCouponDel(t *testing.T) {
	coupon, err := Del("25OFF", nil)
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
}

func TestCouponGet(t *testing.T) {
	coupon, err := Get("25OFF", nil)
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
}

func TestCouponList(t *testing.T) {
	i := List(&stripe.CouponListParams{})

	// Verify that we can get at least one coupon
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Coupon())
}

func TestCouponNew(t *testing.T) {
	coupon, err := New(&stripe.CouponParams{
		Duration:       "repeating",
		DurationPeriod: 3,
		ID:             "25OFF",
		Percent:        25,
	})
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
}

func TestCouponUpdate(t *testing.T) {
	coupon, err := Update("25OFF", &stripe.CouponParams{
		Redemptions: 5,
	})
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
}
