package transfer

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
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
	assert.NotNil(t, i.TransferList())
}

func TestTransferNew(t *testing.T) {
	transfer, err := New(&stripe.TransferParams{
		Amount:            stripe.Int64(123),
		Currency:          stripe.String(string(stripe.CurrencyUSD)),
		Destination:       stripe.String("acct_123"),
		SourceTransaction: stripe.String("ch_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, transfer)
}

func TestTransferUpdate(t *testing.T) {
	transfer, err := Update("tr_123", &stripe.TransferParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, transfer)
}
