//
//
// File generated from our OpenAPI spec
//
//

// Package customerevaluation provides the /v1/radar/customer_evaluations APIs
package customerevaluation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/radar/customer_evaluations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new CustomerEvaluation object.
func New(params *stripe.RadarCustomerEvaluationParams) (*stripe.RadarCustomerEvaluation, error) {
	return getC().New(params)
}

// Creates a new CustomerEvaluation object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.RadarCustomerEvaluationParams) (*stripe.RadarCustomerEvaluation, error) {
	customerevaluation := &stripe.RadarCustomerEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/customer_evaluations", c.Key, params, customerevaluation)
	return customerevaluation, err
}

// Reports an event on a CustomerEvaluation object.
func Update(id string, params *stripe.RadarCustomerEvaluationParams) (*stripe.RadarCustomerEvaluation, error) {
	return getC().Update(id, params)
}

// Reports an event on a CustomerEvaluation object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.RadarCustomerEvaluationParams) (*stripe.RadarCustomerEvaluation, error) {
	path := stripe.FormatURLPath("/v1/radar/customer_evaluations/%s/report", id)
	customerevaluation := &stripe.RadarCustomerEvaluation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, customerevaluation)
	return customerevaluation, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
