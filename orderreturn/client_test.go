package orderreturn

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestOrderReturnNew(t *testing.T) {
	orret, err := New(&stripe.OrderReturnParams{
		Order: stripe.String("or_123"),
	})

	// Verify that we can create an order return
	assert.Nil(t, err)
	assert.NotNil(t, orret)
}

func TestOrderReturnGet(t *testing.T) {
	orret, err := Get("orret_123", &stripe.OrderReturnParams{})

	// Verify that we can get an order return
	assert.Nil(t, err)
	assert.NotNil(t, orret)
}

func TestOrderReturnList(t *testing.T) {
	i := List(&stripe.OrderReturnListParams{})

	// Verify that we can get at least one order return
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.OrderReturn())
}
