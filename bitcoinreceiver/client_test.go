package bitcoinreceiver

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
	_ "github.com/stripe/stripe-go/testing"
)

func TestBitcoinReceiverGet(t *testing.T) {
	receiver, err := Get("btcrcv_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, receiver)
}

func TestBitcoinReceiverList(t *testing.T) {
	i := List(&stripe.BitcoinReceiverListParams{})

	// Verify that we can get at least one receiver
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BitcoinReceiver())
}

func TestBitcoinReceiverNew(t *testing.T) {
	receiver, err := New(&stripe.BitcoinReceiverParams{
		Amount:   1000,
		Currency: currency.USD,
		Email:    "a@b.com",
		Desc:     "some details",
	})
	assert.Nil(t, err)
	assert.NotNil(t, receiver)
}

func TestBitcoinReceiverUpdate(t *testing.T) {
	receiver, err := Update("btcrcv_123", &stripe.BitcoinReceiverUpdateParams{
		Desc: "some other details",
	})
	assert.Nil(t, err)
	assert.NotNil(t, receiver)
}
