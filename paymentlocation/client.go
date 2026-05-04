//
//
// File generated from our OpenAPI spec
//
//

// Package paymentlocation provides the /v1/payment_locations APIs
package paymentlocation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/payment_locations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Payment Location.
func New(params *stripe.PaymentLocationParams) (*stripe.PaymentLocation, error) {
	return getC().New(params)
}

// Create a Payment Location.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.PaymentLocationParams) (*stripe.PaymentLocation, error) {
	paymentlocation := &stripe.PaymentLocation{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_locations", c.Key, params, paymentlocation)
	return paymentlocation, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
