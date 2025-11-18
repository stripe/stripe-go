//
//
// File generated from our OpenAPI spec
//
//

// Package customersession provides the /v1/customer_sessions APIs
package customersession

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/customer_sessions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a Customer Session object that includes a single-use client secret that you can use on your front-end to grant client-side API access for certain customer resources.
func New(params *stripe.CustomerSessionParams) (*stripe.CustomerSession, error) {
	return getC().New(params)
}

// Creates a Customer Session object that includes a single-use client secret that you can use on your front-end to grant client-side API access for certain customer resources.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.CustomerSessionParams) (*stripe.CustomerSession, error) {
	customersession := &stripe.CustomerSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/customer_sessions", c.Key, params, customersession)
	return customersession, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
