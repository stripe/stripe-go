//
//
// File generated from our OpenAPI spec
//
//

// Package issuingauthorizationevaluation provides the /v1/radar/issuing_authorization_evaluations APIs
package issuingauthorizationevaluation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/radar/issuing_authorization_evaluations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Request a fraud risk assessment from Stripe for an Issuing card authorization.
func New(params *stripe.RadarIssuingAuthorizationEvaluationParams) (*stripe.RadarIssuingAuthorizationEvaluation, error) {
	return getC().New(params)
}

// Request a fraud risk assessment from Stripe for an Issuing card authorization.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.RadarIssuingAuthorizationEvaluationParams) (*stripe.RadarIssuingAuthorizationEvaluation, error) {
	issuingauthorizationevaluation := &stripe.RadarIssuingAuthorizationEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/issuing_authorization_evaluations", c.Key, params, issuingauthorizationevaluation)
	return issuingauthorizationevaluation, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
