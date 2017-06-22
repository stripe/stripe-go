package order

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestOrderGet(t *testing.T) {
	order, err := Get("or_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, order)
}

func TestOrderList(t *testing.T) {
	i := List(&stripe.OrderListParams{})

	// Verify that we can get at least one order
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Order())
}

func TestOrderNew(t *testing.T) {
	order, err := New(&stripe.OrderParams{
		Currency: "usd",
	})
	assert.Nil(t, err)
	assert.NotNil(t, order)
}

func TestOrderPay(t *testing.T) {
	order, err := Pay("or_123", &stripe.OrderPayParams{
		Customer: "cus_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, order)
}

func TestOrderReturn(t *testing.T) {
	order, err := Return("or_123", &stripe.OrderReturnParams{})
	assert.Nil(t, err)
	assert.NotNil(t, order)
}

func TestOrderUpdate(t *testing.T) {
	order, err := Update("or_123", &stripe.OrderUpdateParams{
		Status: "fulfilled",
	})
	assert.Nil(t, err)
	assert.NotNil(t, order)
}
