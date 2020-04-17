package bitcointransaction

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v70"
	_ "github.com/stripe/stripe-go/v70/testing"
)

func TestBitcoinTransactionList(t *testing.T) {
	i := List(&stripe.BitcoinTransactionListParams{
		Receiver: stripe.String("btcrcv_123"),
	})

	// Verify that we can get at least one transaction
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BitcoinTransaction())
}
