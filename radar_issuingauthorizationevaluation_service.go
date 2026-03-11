//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1RadarIssuingAuthorizationEvaluationService is used to invoke /v1/radar/issuing_authorization_evaluations APIs.
type v1RadarIssuingAuthorizationEvaluationService struct {
	B   Backend
	Key string
}

// Request a fraud risk assessment from Stripe for an Issuing card authorization.
func (c v1RadarIssuingAuthorizationEvaluationService) Create(ctx context.Context, params *RadarIssuingAuthorizationEvaluationCreateParams) (*RadarIssuingAuthorizationEvaluation, error) {
	if params == nil {
		params = &RadarIssuingAuthorizationEvaluationCreateParams{}
	}
	params.Context = ctx
	issuingauthorizationevaluation := &RadarIssuingAuthorizationEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/issuing_authorization_evaluations", c.Key, params, issuingauthorizationevaluation)
	return issuingauthorizationevaluation, err
}
