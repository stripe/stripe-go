package threedsecure

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestThreeDSecureGet(t *testing.T) {
	threeDSecure, err := Get("tdsrc_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, threeDSecure)
}

func TestThreeDSecureNew(t *testing.T) {
	threeDSecure, err := New(&stripe.ThreeDSecureParams{
		Amount:    stripe.Int64(1000),
		Currency:  stripe.String(string(stripe.CurrencyUSD)),
		Customer:  stripe.String("cus_123"),
		Card:      stripe.String("card_123"),
		ReturnURL: stripe.String("https://test.com"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, threeDSecure)
}
