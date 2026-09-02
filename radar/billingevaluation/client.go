//
//
// File generated from our OpenAPI spec
//
//

// Package billingevaluation provides the /v1/radar/billing_evaluations APIs
package billingevaluation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke /v1/radar/billing_evaluations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Request Stripe Radar's assessment of the non-payment abuse risk of an upcoming charge, before the payment is attempted.
func New(params *stripe.RadarBillingEvaluationParams) (*stripe.RadarBillingEvaluation, error) {
	return getC().New(params)
}

// Request Stripe Radar's assessment of the non-payment abuse risk of an upcoming charge, before the payment is attempted.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.RadarBillingEvaluationParams) (*stripe.RadarBillingEvaluation, error) {
	billingevaluation := &stripe.RadarBillingEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/billing_evaluations", c.Key, params, billingevaluation)
	return billingevaluation, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
