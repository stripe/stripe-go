//
//
// File generated from our OpenAPI spec
//
//

// Package accountevaluation provides the accountevaluation related APIs
package accountevaluation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke accountevaluation related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new account evaluation to request signal evaluations on an account, customer, or inline account data.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2SignalsAccountEvaluationParams) (*stripe.V2SignalsAccountEvaluation, error) {
	accountevaluation := &stripe.V2SignalsAccountEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v2/signals/account_evaluations", c.Key, params, accountevaluation)
	return accountevaluation, err
}

// Retrieves an AccountEvaluation by its ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2SignalsAccountEvaluationParams) (*stripe.V2SignalsAccountEvaluation, error) {
	path := stripe.FormatURLPath("/v2/signals/account_evaluations/%s", id)
	accountevaluation := &stripe.V2SignalsAccountEvaluation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountevaluation)
	return accountevaluation, err
}
