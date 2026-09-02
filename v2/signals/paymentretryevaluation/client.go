//
//
// File generated from our OpenAPI spec
//
//

// Package paymentretryevaluation provides the paymentretryevaluation related APIs
package paymentretryevaluation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke paymentretryevaluation related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new payment retry evaluation for a failed payment.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2SignalsPaymentRetryEvaluationParams) (*stripe.V2SignalsPaymentRetryEvaluation, error) {
	paymentretryevaluation := &stripe.V2SignalsPaymentRetryEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v2/signals/payment_retry_evaluations", c.Key, params, paymentretryevaluation)
	return paymentretryevaluation, err
}

// Retrieves a payment retry evaluation by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2SignalsPaymentRetryEvaluationParams) (*stripe.V2SignalsPaymentRetryEvaluation, error) {
	path := stripe.FormatURLPath("/v2/signals/payment_retry_evaluations/%s", id)
	paymentretryevaluation := &stripe.V2SignalsPaymentRetryEvaluation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentretryevaluation)
	return paymentretryevaluation, err
}

// Updates an active payment retry evaluation with a replacement payment identifier.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2SignalsPaymentRetryEvaluationParams) (*stripe.V2SignalsPaymentRetryEvaluation, error) {
	path := stripe.FormatURLPath("/v2/signals/payment_retry_evaluations/%s", id)
	paymentretryevaluation := &stripe.V2SignalsPaymentRetryEvaluation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentretryevaluation)
	return paymentretryevaluation, err
}

// Cancels an active payment retry evaluation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2SignalsPaymentRetryEvaluationCancelParams) (*stripe.V2SignalsPaymentRetryEvaluation, error) {
	path := stripe.FormatURLPath(
		"/v2/signals/payment_retry_evaluations/%s/cancel", id)
	paymentretryevaluation := &stripe.V2SignalsPaymentRetryEvaluation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentretryevaluation)
	return paymentretryevaluation, err
}
