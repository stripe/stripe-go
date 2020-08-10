package feerefund

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestFeeRefundGet(t *testing.T) {
	refund, err := Get("fr_123", &stripe.FeeRefundParams{
		ApplicationFee: stripe.String("fee_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}

func TestFeeRefundList(t *testing.T) {
	i := List(&stripe.FeeRefundListParams{
		ApplicationFee: stripe.String("fee_123"),
	})

	// Verify that we can get at least one refund
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.FeeRefund())
	assert.NotNil(t, i.FeeRefundList())
}

func TestFeeRefundNew(t *testing.T) {
	refund, err := New(&stripe.FeeRefundParams{
		ApplicationFee: stripe.String("fee_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}

func TestFeeRefundUpdate(t *testing.T) {
	refund, err := Update("fr_123", &stripe.FeeRefundParams{
		ApplicationFee: stripe.String("fee_123"),
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, refund)
}
