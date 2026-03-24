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

// v2CoreAccountEvaluationService is used to invoke accountevaluation related APIs.
type v2CoreAccountEvaluationService struct {
	B   Backend
	Key string
}

// Creates a new account evaluation to trigger signal evaluations on an account or account data.
func (c v2CoreAccountEvaluationService) Create(ctx context.Context, params *V2CoreAccountEvaluationCreateParams) (*V2CoreAccountEvaluation, error) {
	if params == nil {
		params = &V2CoreAccountEvaluationCreateParams{}
	}
	params.Context = ctx
	accountevaluation := &V2CoreAccountEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/account_evaluations", c.Key, params, accountevaluation)
	return accountevaluation, err
}
