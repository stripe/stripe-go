//
//
// File generated from our OpenAPI spec
//
//

// Package group provides the group related APIs
package group

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke group related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a BillingGroup.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingGroupParams) (*stripe.V2BillingGroup, error) {
	group := &stripe.V2BillingGroup{}
	err := c.B.Call(http.MethodPost, "/v2/billing/groups", c.Key, params, group)
	return group, err
}

// Retrieve a BillingGroup.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingGroupParams) (*stripe.V2BillingGroup, error) {
	path := stripe.FormatURLPath("/v2/billing/groups/%s", id)
	group := &stripe.V2BillingGroup{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, group)
	return group, err
}

// Update a BillingGroup.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingGroupParams) (*stripe.V2BillingGroup, error) {
	path := stripe.FormatURLPath("/v2/billing/groups/%s", id)
	group := &stripe.V2BillingGroup{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, group)
	return group, err
}

// Cancel a BillingGroup.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2BillingGroupCancelParams) (*stripe.V2BillingGroup, error) {
	path := stripe.FormatURLPath("/v2/billing/groups/%s/cancel", id)
	group := &stripe.V2BillingGroup{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, group)
	return group, err
}

// List Billing Groups.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingGroupListParams) stripe.Seq2[*stripe.V2BillingGroup, error] {
	if listParams == nil {
		listParams = &stripe.V2BillingGroupListParams{}
	}
	return stripe.NewV2List("/v2/billing/groups", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingGroup], error) {
		page := &stripe.V2Page[*stripe.V2BillingGroup]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
