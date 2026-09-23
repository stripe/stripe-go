//
//
// File generated from our OpenAPI spec
//
//

// Package providerconnection provides the providerconnection related APIs
package providerconnection

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke providerconnection related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Unlinks a provider connection so it can no longer be used to create resources.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Unlink(id string, params *stripe.V2ProvisioningProviderConnectionUnlinkParams) (*stripe.V2ProvisioningProviderConnection, error) {
	path := stripe.FormatURLPath(
		"/v2/provisioning/provider_connections/%s/unlink", id)
	providerconnection := &stripe.V2ProvisioningProviderConnection{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, providerconnection)
	return providerconnection, err
}

// Lists the provider connections for the account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2ProvisioningProviderConnectionListParams) stripe.Seq2[*stripe.V2ProvisioningProviderConnection, error] {
	if listParams == nil {
		listParams = &stripe.V2ProvisioningProviderConnectionListParams{}
	}
	return stripe.NewV2List("/v2/provisioning/provider_connections", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2ProvisioningProviderConnection], error) {
		page := &stripe.V2Page[*stripe.V2ProvisioningProviderConnection]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
