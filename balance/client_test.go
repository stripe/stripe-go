package balance

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestBalanceGet(t *testing.T) {
	balance, err := Get(nil)
	assert.Nil(t, err)
	assert.NotNil(t, balance)
}

func TestBalanceList(t *testing.T) {
	i := List(&stripe.TxListParams{})

	// Verify that we can get at least one transaction
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Transaction())
}

func TestBalanceTxGet(t *testing.T) {
	transaction, err := GetTx("txn_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
}
