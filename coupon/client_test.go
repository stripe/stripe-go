package coupon_test

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/client"
	. "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestCouponDel(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	coupon, err := sc.Coupons.Del("25OFF", nil)
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
}

func TestCouponGet(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	coupon, err := sc.Coupons.Get("25OFF", nil)
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
}

func TestCouponList(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	i := sc.Coupons.List(&stripe.CouponListParams{})

	// Verify that we can get at least one coupon
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Coupon())
	assert.NotNil(t, i.CouponList())
}

func TestCouponNew(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	coupon, err := sc.Coupons.New(&stripe.CouponParams{
		AppliesTo: &stripe.CouponAppliesToParams{
			Products: stripe.StringSlice([]string{
				"prod_123",
				"prod_abc",
			}),
		},
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Duration:   stripe.String(string(stripe.CouponDurationOnce)),
		ID:         stripe.String("25OFF"),
		PercentOff: stripe.Float64(12.5),
	})
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
}

func TestCouponUpdate(t *testing.T) {
	sc := client.New(TestAPIKey, nil)
	coupon, err := sc.Coupons.Update("25OFF", &stripe.CouponParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
}
