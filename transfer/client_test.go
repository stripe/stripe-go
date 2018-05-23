package transfer

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
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
		Amount:            stripe.Int64(123),
		Currency:          stripe.String(string(currency.USD)),
		Destination:       stripe.String("acct_123"),
		SourceTransaction: stripe.String("ch_123"),
		SourceType:        stripe.String(string(SourceCard)),
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
