//
//
// File generated from our OpenAPI spec
//
//

// Package service provides the service related APIs
package service

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke service related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Lists services available in the catalog.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2ProvisioningCatalogServiceListParams) stripe.Seq2[*stripe.V2ProvisioningProviderServiceDetail, error] {
	if listParams == nil {
		listParams = &stripe.V2ProvisioningCatalogServiceListParams{}
	}
	return stripe.NewV2List("/v2/provisioning/catalog/services", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2ProvisioningProviderServiceDetail], error) {
		page := &stripe.V2Page[*stripe.V2ProvisioningProviderServiceDetail]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
