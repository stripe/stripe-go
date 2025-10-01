//
//
// File generated from our OpenAPI spec
//
//

// Package serviceaction provides the serviceaction related APIs
package serviceaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke serviceaction related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Service Action object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingServiceActionParams) (*stripe.V2BillingServiceAction, error) {
	serviceaction := &stripe.V2BillingServiceAction{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/service_actions", c.Key, params, serviceaction)
	return serviceaction, err
}

// Retrieve a Service Action object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingServiceActionParams) (*stripe.V2BillingServiceAction, error) {
	path := stripe.FormatURLPath("/v2/billing/service_actions/%s", id)
	serviceaction := &stripe.V2BillingServiceAction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, serviceaction)
	return serviceaction, err
}

// Update a ServiceAction object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingServiceActionParams) (*stripe.V2BillingServiceAction, error) {
	path := stripe.FormatURLPath("/v2/billing/service_actions/%s", id)
	serviceaction := &stripe.V2BillingServiceAction{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, serviceaction)
	return serviceaction, err
}
