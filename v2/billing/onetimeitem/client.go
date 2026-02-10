//
//
// File generated from our OpenAPI spec
//
//

// Package onetimeitem provides the onetimeitem related APIs
package onetimeitem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke onetimeitem related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a One-Time Item object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingOneTimeItemParams) (*stripe.V2BillingOneTimeItem, error) {
	onetimeitem := &stripe.V2BillingOneTimeItem{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/one_time_items", c.Key, params, onetimeitem)
	return onetimeitem, err
}

// Retrieve a One-Time Item object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingOneTimeItemParams) (*stripe.V2BillingOneTimeItem, error) {
	path := stripe.FormatURLPath("/v2/billing/one_time_items/%s", id)
	onetimeitem := &stripe.V2BillingOneTimeItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, onetimeitem)
	return onetimeitem, err
}

// Update a One-Time Item object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingOneTimeItemParams) (*stripe.V2BillingOneTimeItem, error) {
	path := stripe.FormatURLPath("/v2/billing/one_time_items/%s", id)
	onetimeitem := &stripe.V2BillingOneTimeItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, onetimeitem)
	return onetimeitem, err
}

// List all One-Time Item objects in reverse chronological order of creation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingOneTimeItemListParams) stripe.Seq2[*stripe.V2BillingOneTimeItem, error] {
	return stripe.NewV2List("/v2/billing/one_time_items", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingOneTimeItem], error) {
		page := &stripe.V2Page[*stripe.V2BillingOneTimeItem]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
