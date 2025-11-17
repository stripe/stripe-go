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

// v1RadarAccountEvaluationService is used to invoke /v1/radar/account_evaluations APIs.
type v1RadarAccountEvaluationService struct {
	B   Backend
	Key string
}

// Creates a new AccountEvaluation object.
func (c v1RadarAccountEvaluationService) Create(ctx context.Context, params *RadarAccountEvaluationCreateParams) (*RadarAccountEvaluation, error) {
	if params == nil {
		params = &RadarAccountEvaluationCreateParams{}
	}
	params.Context = ctx
	accountevaluation := &RadarAccountEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/account_evaluations", c.Key, params, accountevaluation)
	return accountevaluation, err
}

// Retrieves an AccountEvaluation object.
func (c v1RadarAccountEvaluationService) Retrieve(ctx context.Context, id string, params *RadarAccountEvaluationRetrieveParams) (*RadarAccountEvaluation, error) {
	if params == nil {
		params = &RadarAccountEvaluationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/radar/account_evaluations/%s", id)
	accountevaluation := &RadarAccountEvaluation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountevaluation)
	return accountevaluation, err
}

// Reports an event on an AccountEvaluation object.
func (c v1RadarAccountEvaluationService) Update(ctx context.Context, id string, params *RadarAccountEvaluationUpdateParams) (*RadarAccountEvaluation, error) {
	if params == nil {
		params = &RadarAccountEvaluationUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/radar/account_evaluations/%s/report_event", id)
	accountevaluation := &RadarAccountEvaluation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountevaluation)
	return accountevaluation, err
}
