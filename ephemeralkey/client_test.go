package ephemeralkey

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v55"
	_ "github.com/stripe/stripe-go/v55/testing"
)

func TestEphemeralKeyDel(t *testing.T) {
	key, err := Del("ephkey_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, key)
}

func TestEphemeralKeyNew(t *testing.T) {
	key, err := New(&stripe.EphemeralKeyParams{
		Customer:      stripe.String("cus_123"),
		StripeVersion: stripe.String("2018-02-06"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, key)
}
