package customer

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestCustomerDel(t *testing.T) {
	customer, err := Del("cus_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestCustomerGet(t *testing.T) {
	customer, err := Get("cus_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestCustomerList(t *testing.T) {
	i := List(&stripe.CustomerListParams{})

	// Verify that we can get at least one customer
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Customer())
}

func TestCustomerNew(t *testing.T) {
	customer, err := New(&stripe.CustomerParams{
		Email: "foo@example.com",
	})
	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestCustomerNew_NilParams(t *testing.T) {
	customer, err := New(nil)
	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestCustomerUpdate(t *testing.T) {
	customer, err := Update("cus_123", &stripe.CustomerParams{
		Email: "foo@example.com",
	})
	assert.Nil(t, err)
	assert.NotNil(t, customer)
}
