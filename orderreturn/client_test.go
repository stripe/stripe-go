package orderreturn

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestOrderReturnList(t *testing.T) {
	i := List(&stripe.OrderReturnListParams{})

	// Verify that we can get at least one order return
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.OrderReturn())
}
