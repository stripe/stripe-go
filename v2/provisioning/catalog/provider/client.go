//
//
// File generated from our OpenAPI spec
//
//

// Package provider provides the provider related APIs
package provider

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke provider related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Lists providers available in the catalog.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2ProvisioningCatalogProviderListParams) stripe.Seq2[*stripe.V2ProvisioningProvider, error] {
	if listParams == nil {
		listParams = &stripe.V2ProvisioningCatalogProviderListParams{}
	}
	return stripe.NewV2List("/v2/provisioning/catalog/providers", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2ProvisioningProvider], error) {
		page := &stripe.V2Page[*stripe.V2ProvisioningProvider]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
