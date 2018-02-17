package order

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
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
		Currency: stripe.String(string(currency.USD)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, order)
}

func TestOrderPay(t *testing.T) {
	order, err := Pay("or_123", &stripe.OrderPayParams{
		Customer: stripe.String("cus_123"),
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
		Status: stripe.String(string(stripe.StatusFulfilled)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, order)
}
