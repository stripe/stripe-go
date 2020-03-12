package sourcetransaction

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/channelmeter/stripe-go"
	_ "github.com/channelmeter/stripe-go/testing"
)

func TestSourceTransactionList(t *testing.T) {
	i := List(&stripe.SourceTransactionListParams{
		Source: stripe.String("src_123"),
	})

	// Verify that we can get at least one transaction
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SourceTransaction())
}
