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
		Amount:    1000,
		Currency:  "usd",
		Customer:  "cus_123",
		Card:      "card_123",
		ReturnURL: "https://test.com",
	})
	assert.Nil(t, err)
	assert.NotNil(t, threeDSecure)
}
