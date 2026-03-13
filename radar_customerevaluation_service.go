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

// v1RadarCustomerEvaluationService is used to invoke /v1/radar/customer_evaluations APIs.
type v1RadarCustomerEvaluationService struct {
	B   Backend
	Key string
}

// Creates a new CustomerEvaluation object.
func (c v1RadarCustomerEvaluationService) Create(ctx context.Context, params *RadarCustomerEvaluationCreateParams) (*RadarCustomerEvaluation, error) {
	if params == nil {
		params = &RadarCustomerEvaluationCreateParams{}
	}
	params.Context = ctx
	customerevaluation := &RadarCustomerEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/customer_evaluations", c.Key, params, customerevaluation)
	return customerevaluation, err
}

// Reports an event on a CustomerEvaluation object.
func (c v1RadarCustomerEvaluationService) Update(ctx context.Context, id string, params *RadarCustomerEvaluationUpdateParams) (*RadarCustomerEvaluation, error) {
	if params == nil {
		params = &RadarCustomerEvaluationUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/radar/customer_evaluations/%s/report", id)
	customerevaluation := &RadarCustomerEvaluation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, customerevaluation)
	return customerevaluation, err
}
