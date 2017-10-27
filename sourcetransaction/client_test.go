package sourcetransaction

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestSourceTransactionList(t *testing.T) {
	// TODO: unskip the test once stripe-mock supports the /v1/sources/src_.../source_transactions endpoint
	t.Skip("not yet supported by stripe-mock")

	i := List(&stripe.SourceTransactionListParams{
		Source: "src_123",
	})

	// Verify that we can get at least one transaction
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SourceTransaction())
}
