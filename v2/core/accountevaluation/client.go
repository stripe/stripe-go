//
//
// File generated from our OpenAPI spec
//
//

// Package accountevaluation provides the accountevaluation related APIs
package accountevaluation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke accountevaluation related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new account evaluation to trigger signal evaluations on an account or account data.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreAccountEvaluationParams) (*stripe.V2CoreAccountEvaluation, error) {
	accountevaluation := &stripe.V2CoreAccountEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/account_evaluations", c.Key, params, accountevaluation)
	return accountevaluation, err
}
