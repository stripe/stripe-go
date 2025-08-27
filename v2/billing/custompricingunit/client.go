//
//
// File generated from our OpenAPI spec
//
//

// Package custompricingunit provides the custompricingunit related APIs
package custompricingunit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke custompricingunit related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Custom Pricing Unit object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingCustomPricingUnitParams) (*stripe.V2BillingCustomPricingUnit, error) {
	custompricingunit := &stripe.V2BillingCustomPricingUnit{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/custom_pricing_units", c.Key, params, custompricingunit)
	return custompricingunit, err
}

// Retrieve a Custom Pricing Unit object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingCustomPricingUnitParams) (*stripe.V2BillingCustomPricingUnit, error) {
	path := stripe.FormatURLPath("/v2/billing/custom_pricing_units/%s", id)
	custompricingunit := &stripe.V2BillingCustomPricingUnit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, custompricingunit)
	return custompricingunit, err
}

// Update a Custom Pricing Unit object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingCustomPricingUnitParams) (*stripe.V2BillingCustomPricingUnit, error) {
	path := stripe.FormatURLPath("/v2/billing/custom_pricing_units/%s", id)
	custompricingunit := &stripe.V2BillingCustomPricingUnit{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, custompricingunit)
	return custompricingunit, err
}

// List all Custom Pricing Unit objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingCustomPricingUnitListParams) stripe.Seq2[*stripe.V2BillingCustomPricingUnit, error] {
	return stripe.NewV2List("/v2/billing/custom_pricing_units", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingCustomPricingUnit], error) {
		page := &stripe.V2Page[*stripe.V2BillingCustomPricingUnit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
