package customerbalancetransaction

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestCustomerBalanceTransactionGet(t *testing.T) {
	transaction, err := Get("cbtxn_123", &stripe.CustomerBalanceTransactionParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
}

func TestCustomerBalanceTransactionList(t *testing.T) {
	i := List(&stripe.CustomerBalanceTransactionListParams{
		Customer: stripe.String("cus_123"),
	})

	// Verify that we can get at least one transaction
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.CustomerBalanceTransaction())
	assert.NotNil(t, i.CustomerBalanceTransactionList())
}

func TestCustomerBalanceTransactionNew(t *testing.T) {
	transaction, err := New(&stripe.CustomerBalanceTransactionParams{
		Amount:   stripe.Int64(1234),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
}

func TestCustomerBalanceTransactionUpdate(t *testing.T) {
	transaction, err := Update("cbtxn_123", &stripe.CustomerBalanceTransactionParams{
		Customer:    stripe.String("cus_123"),
		Description: stripe.String("description"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
}
