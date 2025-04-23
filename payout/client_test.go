package payout

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestPayoutCancel(t *testing.T) {
	payout, err := Cancel("po_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, payout)
}

func TestPayoutGet(t *testing.T) {
	payout, err := Get("po_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, payout)
}

func TestPayoutList(t *testing.T) {
	params := &stripe.PayoutListParams{}
	i := List(params)

	// Verify that we can get at least one payout
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Payout())
	assert.NotNil(t, i.PayoutList())
}

func TestPayoutNew(t *testing.T) {
	payout, err := New(&stripe.PayoutParams{
		Amount:   stripe.Int64(123),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, payout)
}

func TestPayoutReverse(t *testing.T) {
	params := &stripe.PayoutReverseParams{}
	payout, err := Reverse("po_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, payout)
}

func TestPayoutUpdate(t *testing.T) {
	params := &stripe.PayoutParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	}
	payout, err := Update("po_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, payout)
}
