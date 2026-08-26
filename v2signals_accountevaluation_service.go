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

// v2SignalsAccountEvaluationService is used to invoke accountevaluation related APIs.
type v2SignalsAccountEvaluationService struct {
	B   Backend
	Key string
}

// Creates a new account evaluation to request signal evaluations on an account, customer, or inline account data.
func (c v2SignalsAccountEvaluationService) Create(ctx context.Context, params *V2SignalsAccountEvaluationCreateParams) (*V2SignalsAccountEvaluation, error) {
	if params == nil {
		params = &V2SignalsAccountEvaluationCreateParams{}
	}
	params.Context = ctx
	accountevaluation := &V2SignalsAccountEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v2/signals/account_evaluations", c.Key, params, accountevaluation)
	return accountevaluation, err
}

// Retrieves an AccountEvaluation by its ID.
func (c v2SignalsAccountEvaluationService) Retrieve(ctx context.Context, id string, params *V2SignalsAccountEvaluationRetrieveParams) (*V2SignalsAccountEvaluation, error) {
	if params == nil {
		params = &V2SignalsAccountEvaluationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/signals/account_evaluations/%s", id)
	accountevaluation := &V2SignalsAccountEvaluation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountevaluation)
	return accountevaluation, err
}
