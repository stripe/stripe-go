package promotioncode

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestPromotionCodeGet(t *testing.T) {
	pc, err := Get("promo_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, pc)
}

func TestPromotionCodeList(t *testing.T) {
	params := &stripe.PromotionCodeListParams{
		Code:     stripe.String("MYCODE"),
		Coupon:   stripe.String("co_123"),
		Customer: stripe.String("cus_123"),
	}
	i := List(params)

	// Verify that we can get at least one pc
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.PromotionCode())
	assert.NotNil(t, i.PromotionCodeList())
}

func TestPromotionCodeNew(t *testing.T) {
	params := &stripe.PromotionCodeParams{
		Code:     stripe.String("MYCODE"),
		Coupon:   stripe.String("co_123"),
		Customer: stripe.String("cus_123"),
		Restrictions: &stripe.PromotionCodeRestrictionsParams{
			FirstTimeTransaction:  stripe.Bool(true),
			MinimumAmount:         stripe.Int64(1234),
			MinimumAmountCurrency: stripe.String(string(stripe.CurrencyUSD)),
		},
	}
	pc, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, pc)
}

func TestPromotionCodeUpdate(t *testing.T) {
	params := &stripe.PromotionCodeParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	}
	pc, err := Update("promo_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, pc)
}
