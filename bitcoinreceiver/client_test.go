package bitcoinreceiver

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestBitcoinReceiverGet(t *testing.T) {
	receiver, err := Get("btcrcv_123")
	assert.Nil(t, err)
	assert.NotNil(t, receiver)
}

func TestBitcoinReceiverList(t *testing.T) {
	i := List(&stripe.BitcoinReceiverListParams{})

	// Verify that we can get at least one receiver
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BitcoinReceiver())
	assert.NotNil(t, i.BitcoinReceiverList())
}

// New and Update endpoints for Bitcoin receivers no longer exist in the API or
// stripe-mock. We've left the functions in the client so as not to break
// anyone that may still have been referencing these, but this entire package
// is very much deprecated.
