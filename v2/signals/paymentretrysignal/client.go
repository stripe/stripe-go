//
//
// File generated from our OpenAPI spec
//
//

// Package paymentretrysignal provides the paymentretrysignal related APIs
package paymentretrysignal

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke paymentretrysignal related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a payment retry signal by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2SignalsPaymentRetrySignalParams) (*stripe.V2SignalsPaymentRetrySignal, error) {
	path := stripe.FormatURLPath("/v2/signals/payment_retry_signals/%s", id)
	paymentretrysignal := &stripe.V2SignalsPaymentRetrySignal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentretrysignal)
	return paymentretrysignal, err
}
