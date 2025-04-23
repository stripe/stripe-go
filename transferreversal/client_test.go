package transferreversal

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestTransferReversalGet(t *testing.T) {
	reversal, err := Get("trr_123", &stripe.TransferReversalParams{
		ID: stripe.String("tr_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}

func TestTransferReversalList(t *testing.T) {
	i := List(&stripe.TransferReversalListParams{
		ID: stripe.String("tr_123"),
	})

	// Verify that we can get at least one reversal
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.TransferReversal())
	assert.NotNil(t, i.TransferReversalList())
}

func TestTransferReversalNew(t *testing.T) {
	reversal, err := New(&stripe.TransferReversalParams{
		Amount: stripe.Int64(123),
		ID:     stripe.String("tr_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}

func TestTransferReversalUpdate(t *testing.T) {
	reversal, err := Update("trr_123", &stripe.TransferReversalParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
		ID: stripe.String("tr_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}
