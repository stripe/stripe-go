package fee

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestApplicationFeeGet(t *testing.T) {
	fee, err := Get("fee_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, fee)
}

func TestApplicationFeeList(t *testing.T) {
	i := List(&stripe.ApplicationFeeListParams{})

	// Verify that we can get at least one fee
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ApplicationFee())
	assert.NotNil(t, i.ApplicationFeeList())
}
