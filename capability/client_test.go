package capability

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestCapabilityGet(t *testing.T) {
	capability, err := Get("acap_123", &stripe.CapabilityParams{
		Account: stripe.String("acct_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, capability)
}

func TestCapabilityList(t *testing.T) {
	i := List(&stripe.CapabilityListParams{
		Account: stripe.String("acct_123"),
	})

	// Verify that we can get at least one capability
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Capability())
	assert.NotNil(t, i.CapabilityList())
}

func TestCapabilityUpdate(t *testing.T) {
	capability, err := Update("acap_123", &stripe.CapabilityParams{
		Account:   stripe.String("acct_123"),
		Requested: stripe.Bool(true),
	})
	assert.Nil(t, err)
	assert.NotNil(t, capability)
}
