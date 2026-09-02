//
//
// File generated from our OpenAPI spec
//
//

// Package integrationconfiguration provides the integrationconfiguration related APIs
package integrationconfiguration

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke integrationconfiguration related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve the tax integration configuration for this account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(params *stripe.V2TaxIntegrationConfigurationParams) (*stripe.V2TaxIntegrationConfiguration, error) {
	integrationconfiguration := &stripe.V2TaxIntegrationConfiguration{}
	err := c.B.Call(
		http.MethodGet, "/v2/tax/integration_configurations", c.Key, params, integrationconfiguration)
	return integrationconfiguration, err
}

// Update the tax integration configuration for this account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(params *stripe.V2TaxIntegrationConfigurationParams) (*stripe.V2TaxIntegrationConfiguration, error) {
	integrationconfiguration := &stripe.V2TaxIntegrationConfiguration{}
	err := c.B.Call(
		http.MethodPost, "/v2/tax/integration_configurations", c.Key, params, integrationconfiguration)
	return integrationconfiguration, err
}
