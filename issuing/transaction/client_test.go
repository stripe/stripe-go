package transaction

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestIssuingTransactionGet(t *testing.T) {
	transaction, err := Get("iauth_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, "issuing.transaction", transaction.Object)
}

func TestIssuingTransactionList(t *testing.T) {
	i := List(&stripe.IssuingTransactionListParams{})

	// Verify that we can get at least one transaction
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IssuingTransaction())
	assert.Equal(t, "issuing.transaction", i.IssuingTransaction().Object)
	assert.NotNil(t, i.IssuingTransactionList())
}

func TestIssuingTransactionUpdate(t *testing.T) {
	params := &stripe.IssuingTransactionParams{}
	params.AddMetadata("key", "value")
	transaction, err := Update("iauth_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, "issuing.transaction", transaction.Object)
}
