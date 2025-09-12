package stripe_test

import (
	"context"
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v82"
	. "github.com/stripe/stripe-go/v82/testing"
)

func TestCustomerList_HasLastResponse(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	i := sc.V1Customers.List(context.TODO(), &stripe.CustomerListParams{
		Email: stripe.String("test@example.com"),
	})
	i(func(customer *stripe.Customer, err error) bool {
		assert.Nil(t, err)
		assert.NotNil(t, customer)
		assert.NotNil(t, customer.LastResponse)
		rawJSON := customer.LastResponse.RawJSON
		lastResponseCustomer := &stripe.Customer{}
		if err := json.Unmarshal(rawJSON, &lastResponseCustomer); err != nil {
			assert.NoError(t, err)
		}
		assert.Equal(t, customer.ID, lastResponseCustomer.ID)
		return true
	})
}

func TestCustomerSearch_HasLastResponse(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	params := &stripe.CustomerSearchParams{}
	params.Query = "email:'test@example.com'"
	i := sc.V1Customers.Search(context.TODO(), params)
	i(func(customer *stripe.Customer, err error) bool {
		assert.Nil(t, err)
		assert.NotNil(t, customer)
		assert.NotNil(t, customer.LastResponse)
		rawJSON := customer.LastResponse.RawJSON
		lastResponseCustomer := &stripe.Customer{}
		if err := json.Unmarshal(rawJSON, &lastResponseCustomer); err != nil {
			assert.NoError(t, err)
		}
		assert.Equal(t, customer.ID, lastResponseCustomer.ID)
		return true
	})
}
