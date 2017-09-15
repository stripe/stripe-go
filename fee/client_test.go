package fee

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestFeeGet(t *testing.T) {
	fee, err := Get("fee_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, fee)
}

func TestFeeList(t *testing.T) {
	i := List(&stripe.FeeListParams{})

	// Verify that we can get at least one fee
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Fee())
}
