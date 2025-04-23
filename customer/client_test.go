package customer

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
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
	assert.NotNil(t, i.CustomerList())
}

func TestCustomerNew(t *testing.T) {
	customer, err := New(&stripe.CustomerParams{
		Email: stripe.String("foo@example.com"),
		Shipping: &stripe.CustomerShippingParams{
			Address: &stripe.AddressParams{
				Line1: stripe.String("line1"),
				City:  stripe.String("city"),
			},
			Name: stripe.String("name"),
		},
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
		Email: stripe.String("foo@example.com"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestCustomerListPaymentMethods(t *testing.T) {
	i := ListPaymentMethods(&stripe.CustomerListPaymentMethodsParams{
		Customer: stripe.String("cus_123"),
		Type:     stripe.String("card"),
	})
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.PaymentMethod())
	assert.NotNil(t, i.PaymentMethodList())
}
