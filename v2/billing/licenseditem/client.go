//
//
// File generated from our OpenAPI spec
//
//

// Package licenseditem provides the licenseditem related APIs
package licenseditem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke licenseditem related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Licensed Item object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingLicensedItemParams) (*stripe.V2BillingLicensedItem, error) {
	licenseditem := &stripe.V2BillingLicensedItem{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/licensed_items", c.Key, params, licenseditem)
	return licenseditem, err
}

// Retrieve a Licensed Item object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingLicensedItemParams) (*stripe.V2BillingLicensedItem, error) {
	path := stripe.FormatURLPath("/v2/billing/licensed_items/%s", id)
	licenseditem := &stripe.V2BillingLicensedItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, licenseditem)
	return licenseditem, err
}

// Update a Licensed Item object. At least one of the fields is required.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingLicensedItemParams) (*stripe.V2BillingLicensedItem, error) {
	path := stripe.FormatURLPath("/v2/billing/licensed_items/%s", id)
	licenseditem := &stripe.V2BillingLicensedItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, licenseditem)
	return licenseditem, err
}

// List all Licensed Item objects in reverse chronological order of creation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingLicensedItemListParams) stripe.Seq2[*stripe.V2BillingLicensedItem, error] {
	return stripe.NewV2List("/v2/billing/licensed_items", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingLicensedItem], error) {
		page := &stripe.V2Page[*stripe.V2BillingLicensedItem]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
