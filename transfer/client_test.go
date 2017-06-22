package transfer

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestTransferGet(t *testing.T) {
	transfer, err := Get("tr_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, transfer)
}

func TestTransferList(t *testing.T) {
	i := List(&stripe.TransferListParams{})

	// Verify that we can get at least one transfer
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Transfer())
}

func TestTransferNew(t *testing.T) {
	transfer, err := New(&stripe.TransferParams{
		Amount:   123,
		Currency: "usd",
		Dest:     "acct_123",
		SourceTx: "ch_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, transfer)
}

func TestTransferUpdate(t *testing.T) {
	transfer, err := Update("tr_123", &stripe.TransferParams{
		Params: stripe.Params{
			Meta: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, transfer)
}
