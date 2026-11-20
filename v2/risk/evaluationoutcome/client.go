//
//
// File generated from our OpenAPI spec
//
//

// Package evaluationoutcome provides the evaluationoutcome related APIs
package evaluationoutcome

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke evaluationoutcome related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an evaluation outcome by its ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2RiskEvaluationOutcomeParams) (*stripe.V2RiskEvaluationOutcome, error) {
	path := stripe.FormatURLPath("/v2/risk/evaluation_outcomes/%s", id)
	evaluationoutcome := &stripe.V2RiskEvaluationOutcome{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, evaluationoutcome)
	return evaluationoutcome, err
}

// Lists all matching historical evaluation outcomes, newest first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2RiskEvaluationOutcomeListParams) stripe.Seq2[*stripe.V2RiskEvaluationOutcome, error] {
	if listParams == nil {
		listParams = &stripe.V2RiskEvaluationOutcomeListParams{}
	}
	return stripe.NewV2List("/v2/risk/evaluation_outcomes", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2RiskEvaluationOutcome], error) {
		page := &stripe.V2Page[*stripe.V2RiskEvaluationOutcome]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
