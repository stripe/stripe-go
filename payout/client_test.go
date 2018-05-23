package payout

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
	_ "github.com/stripe/stripe-go/testing"
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
	i := List(&stripe.PayoutListParams{})

	// Verify that we can get at least one payout
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Payout())
}

func TestPayoutNew(t *testing.T) {
	payout, err := New(&stripe.PayoutParams{
		Amount:   stripe.Int64(123),
		Currency: stripe.String(string(currency.USD)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, payout)
}

func TestPayoutUpdate(t *testing.T) {
	payout, err := Update("tr_123", &stripe.PayoutParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, payout)
}
