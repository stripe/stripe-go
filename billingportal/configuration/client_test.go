package configuration

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestBillingPortalConfigurationGet(t *testing.T) {
	configuration, err := Get("bpc_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, configuration)
}

func TestBillingPortalConfigurationList(t *testing.T) {
	i := List(&stripe.BillingPortalConfigurationListParams{})

	// Verify that we can get at least one configuration
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BillingPortalConfiguration())
	assert.NotNil(t, i.BillingPortalConfigurationList())
}

func TestBillingPortalConfigurationNew(t *testing.T) {
	configuration, err := New(&stripe.BillingPortalConfigurationParams{
		BusinessProfile: &stripe.BillingPortalConfigurationBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/tos"),
		},
		Features: &stripe.BillingPortalConfigurationFeaturesParams{
			CustomerUpdate: &stripe.BillingPortalConfigurationFeaturesCustomerUpdateParams{
				AllowedUpdates: []*string{stripe.String("address")},
				Enabled:        stripe.Bool(true),
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, configuration)
}

func TestBillingPortalConfigurationUpdate(t *testing.T) {
	configuration, err := Update("bpc_123", &stripe.BillingPortalConfigurationParams{
		Active: stripe.Bool(false),
	})
	assert.Nil(t, err)
	assert.NotNil(t, configuration)
}
