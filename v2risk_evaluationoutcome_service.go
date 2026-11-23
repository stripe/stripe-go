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

// v2RiskEvaluationOutcomeService is used to invoke evaluationoutcome related APIs.
type v2RiskEvaluationOutcomeService struct {
	B   Backend
	Key string
}

// Retrieves an evaluation outcome by its ID.
func (c v2RiskEvaluationOutcomeService) Retrieve(ctx context.Context, id string, params *V2RiskEvaluationOutcomeRetrieveParams) (*V2RiskEvaluationOutcome, error) {
	if params == nil {
		params = &V2RiskEvaluationOutcomeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/risk/evaluation_outcomes/%s", id)
	evaluationoutcome := &V2RiskEvaluationOutcome{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, evaluationoutcome)
	return evaluationoutcome, err
}

// Lists all matching historical evaluation outcomes, newest first.
func (c v2RiskEvaluationOutcomeService) List(ctx context.Context, listParams *V2RiskEvaluationOutcomeListParams) *V2List[*V2RiskEvaluationOutcome] {
	if listParams == nil {
		listParams = &V2RiskEvaluationOutcomeListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/risk/evaluation_outcomes", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2RiskEvaluationOutcome], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2RiskEvaluationOutcome]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
