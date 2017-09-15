package ephemeralkey

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestEphemeralKeyDel(t *testing.T) {
	key, err := Del("ephkey_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, key)
}

func TestEphemeralKeyNew(t *testing.T) {
	key, err := New(&stripe.EphemeralKeyParams{
		Customer:      "cus_123",
		StripeVersion: "2017-05-25",
	})
	assert.Nil(t, err)
	assert.NotNil(t, key)
}
