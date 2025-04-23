package balancetransaction

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestBalanceTransactionGet(t *testing.T) {
	transaction, err := Get("txn_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
}

func TestBalanceTransactionList(t *testing.T) {
	i := List(&stripe.BalanceTransactionListParams{})

	// Verify that we can get at least one balance transaction
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BalanceTransaction())
	assert.NotNil(t, i.BalanceTransactionList())
}
