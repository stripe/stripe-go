package order

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
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
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Items: []*stripe.OrderItemParams{
			{Amount: stripe.Int64(1), Description: stripe.String("Item 1")},
			{Amount: stripe.Int64(1), Description: stripe.String("Item 2")},
		},
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

func TestOrderReturn_RequestParams(t *testing.T) {
	// Create an ephemeral test server so that we can inspect request attributes.
	var lastRequest *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		fmt.Fprintln(w, bytes.NewBuffer([]byte("{}")))
	}))
	defer ts.Close()

	// Configure the stripe client to use the ephemeral backend.
	backend := stripe.GetBackendWithConfig(stripe.APIBackend, &stripe.BackendConfig{
		URL: ts.URL,
	})
	orderClient := Client{B: backend, Key: stripe.Key}

	p := &stripe.OrderReturnParams{}
	p.SetStripeAccount("acct_123")

	order, err := orderClient.Return("or_123", p)
	assert.Nil(t, err)
	assert.NotNil(t, order)

	assert.Equal(t, lastRequest.Header.Get("Stripe-Account"), "acct_123")
}

func TestOrderUpdate(t *testing.T) {
	order, err := Update("or_123", &stripe.OrderUpdateParams{
		Status: stripe.String(string(stripe.OrderStatusFulfilled)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, order)
}
