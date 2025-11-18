//
//
// File generated from our OpenAPI spec
//
//

// Package metereditem provides the metereditem related APIs
package metereditem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke metereditem related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Metered Item object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingMeteredItemParams) (*stripe.V2BillingMeteredItem, error) {
	metereditem := &stripe.V2BillingMeteredItem{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/metered_items", c.Key, params, metereditem)
	return metereditem, err
}

// Retrieve a Metered Item object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingMeteredItemParams) (*stripe.V2BillingMeteredItem, error) {
	path := stripe.FormatURLPath("/v2/billing/metered_items/%s", id)
	metereditem := &stripe.V2BillingMeteredItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, metereditem)
	return metereditem, err
}

// Update a Metered Item object. At least one of the fields is required.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingMeteredItemParams) (*stripe.V2BillingMeteredItem, error) {
	path := stripe.FormatURLPath("/v2/billing/metered_items/%s", id)
	metereditem := &stripe.V2BillingMeteredItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, metereditem)
	return metereditem, err
}

// List all Metered Item objects in reverse chronological order of creation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingMeteredItemListParams) stripe.Seq2[*stripe.V2BillingMeteredItem, error] {
	return stripe.NewV2List("/v2/billing/metered_items", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingMeteredItem], error) {
		page := &stripe.V2Page[*stripe.V2BillingMeteredItem]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
