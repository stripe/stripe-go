package feerefund

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestApplicationFeeRefundGet(t *testing.T) {
	refund, err := Get("fr_123", &stripe.ApplicationFeeRefundParams{
		ApplicationFee: "fee_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}

func TestApplicationFeeRefundList(t *testing.T) {
	i := List(&stripe.ApplicationFeeRefundListParams{
		ApplicationFee: "fee_123",
	})

	// Verify that we can get at least one refund
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ApplicationFeeRefund())
}

func TestApplicationFeeRefundNew(t *testing.T) {
	refund, err := New(&stripe.ApplicationFeeRefundParams{
		ApplicationFee: "fee_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}

func TestApplicationFeeRefundUpdate(t *testing.T) {
	refund, err := Update("fr_123", &stripe.ApplicationFeeRefundParams{
		ApplicationFee: "fee_123",
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}
