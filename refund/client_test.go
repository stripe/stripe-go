package refund

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestRefundGet(t *testing.T) {
	refund, err := Get("gold", nil)
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}

func TestRefundList(t *testing.T) {
	i := List(&stripe.RefundListParams{})

	// Verify that we can get at least one refund
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Refund())
	assert.NotNil(t, i.RefundList())
}

func TestRefundNew(t *testing.T) {
	refund, err := New(&stripe.RefundParams{
		Charge:        stripe.String("ch_123"),
		PaymentIntent: stripe.String("pi_123"),
		Reason:        stripe.String(string(stripe.RefundReasonDuplicate)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}

func TestRefundUpdate(t *testing.T) {
	refund, err := Update("gold", &stripe.RefundParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}
