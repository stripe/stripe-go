package coupon

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
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
		Currency:         stripe.String(string(currency.USD)),
		Duration:         stripe.String(string(Repeating)),
		DurationInMonths: stripe.Int64(3),
		ID:               stripe.String("25OFF"),
		PercentOff:       stripe.Int64(25),
	})
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
}

func TestCouponUpdate(t *testing.T) {
	coupon, err := Update("25OFF", &stripe.CouponParams{
		MaxRedemptions: stripe.Int64(5),
	})
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
}
