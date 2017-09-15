package refund

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
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
}

func TestRefundNew(t *testing.T) {
	refund, err := New(&stripe.RefundParams{
		Charge: "ch_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}

func TestRefundUpdate(t *testing.T) {
	refund, err := Update("gold", &stripe.RefundParams{
		Params: stripe.Params{
			Meta: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}
