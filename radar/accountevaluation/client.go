//
//
// File generated from our OpenAPI spec
//
//

// Package accountevaluation provides the /v1/radar/account_evaluations APIs
package accountevaluation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke /v1/radar/account_evaluations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new AccountEvaluation object.
func New(params *stripe.RadarAccountEvaluationParams) (*stripe.RadarAccountEvaluation, error) {
	return getC().New(params)
}

// Creates a new AccountEvaluation object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.RadarAccountEvaluationParams) (*stripe.RadarAccountEvaluation, error) {
	accountevaluation := &stripe.RadarAccountEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/account_evaluations", c.Key, params, accountevaluation)
	return accountevaluation, err
}

// Retrieves an AccountEvaluation object.
func Get(id string, params *stripe.RadarAccountEvaluationParams) (*stripe.RadarAccountEvaluation, error) {
	return getC().Get(id, params)
}

// Retrieves an AccountEvaluation object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.RadarAccountEvaluationParams) (*stripe.RadarAccountEvaluation, error) {
	path := stripe.FormatURLPath("/v1/radar/account_evaluations/%s", id)
	accountevaluation := &stripe.RadarAccountEvaluation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountevaluation)
	return accountevaluation, err
}

// Reports an event on an AccountEvaluation object.
func Update(id string, params *stripe.RadarAccountEvaluationParams) (*stripe.RadarAccountEvaluation, error) {
	return getC().Update(id, params)
}

// Reports an event on an AccountEvaluation object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.RadarAccountEvaluationParams) (*stripe.RadarAccountEvaluation, error) {
	path := stripe.FormatURLPath(
		"/v1/radar/account_evaluations/%s/report_event", id)
	accountevaluation := &stripe.RadarAccountEvaluation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountevaluation)
	return accountevaluation, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
