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

// v1RadarBillingEvaluationService is used to invoke /v1/radar/billing_evaluations APIs.
type v1RadarBillingEvaluationService struct {
	B   Backend
	Key string
}

// Request Stripe Radar's assessment of the non-payment abuse risk of an upcoming charge, before the payment is attempted.
func (c v1RadarBillingEvaluationService) Create(ctx context.Context, params *RadarBillingEvaluationCreateParams) (*RadarBillingEvaluation, error) {
	if params == nil {
		params = &RadarBillingEvaluationCreateParams{}
	}
	params.Context = ctx
	billingevaluation := &RadarBillingEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/billing_evaluations", c.Key, params, billingevaluation)
	return billingevaluation, err
}
