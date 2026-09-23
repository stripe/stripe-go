//
//
// File generated from our OpenAPI spec
//
//

// Package providerconnectionrequest provides the providerconnectionrequest related APIs
package providerconnectionrequest

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke providerconnectionrequest related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new provider connection.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2ProvisioningProviderConnectionRequestParams) (*stripe.V2ProvisioningProviderConnectionRequest, error) {
	providerconnectionrequest := &stripe.V2ProvisioningProviderConnectionRequest{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/provider_connection_requests", c.Key, params, providerconnectionrequest)
	return providerconnectionrequest, err
}

// Retrieves a provider connection.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2ProvisioningProviderConnectionRequestParams) (*stripe.V2ProvisioningProviderConnectionRequest, error) {
	path := stripe.FormatURLPath(
		"/v2/provisioning/provider_connection_requests/%s", id)
	providerconnectionrequest := &stripe.V2ProvisioningProviderConnectionRequest{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, providerconnectionrequest)
	return providerconnectionrequest, err
}

// Submits additional information requested by the provider for a provider connection.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SubmitInformation(id string, params *stripe.V2ProvisioningProviderConnectionRequestSubmitInformationParams) (*stripe.V2ProvisioningProviderConnectionRequest, error) {
	path := stripe.FormatURLPath(
		"/v2/provisioning/provider_connection_requests/%s/submit_information", id)
	providerconnectionrequest := &stripe.V2ProvisioningProviderConnectionRequest{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, providerconnectionrequest)
	return providerconnectionrequest, err
}
