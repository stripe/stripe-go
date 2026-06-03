//
//
// File generated from our OpenAPI spec
//
//

// Package order provides the /v1/delegated_checkout/orders APIs
package order

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/delegated_checkout/orders APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a delegated checkout order.
func Get(id string, params *stripe.DelegatedCheckoutOrderParams) (*stripe.DelegatedCheckoutOrder, error) {
	return getC().Get(id, params)
}

// Retrieves a delegated checkout order.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.DelegatedCheckoutOrderParams) (*stripe.DelegatedCheckoutOrder, error) {
	path := stripe.FormatURLPath("/v1/delegated_checkout/orders/%s", id)
	order := &stripe.DelegatedCheckoutOrder{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, order)
	return order, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
