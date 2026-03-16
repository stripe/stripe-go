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

// v1RadarPaymentEvaluationService is used to invoke /v1/radar/payment_evaluations APIs.
type v1RadarPaymentEvaluationService struct {
	B   Backend
	Key string
}

// Request a Radar API fraud risk score from Stripe for a payment before sending it for external processor authorization.
func (c v1RadarPaymentEvaluationService) Create(ctx context.Context, params *RadarPaymentEvaluationCreateParams) (*RadarPaymentEvaluation, error) {
	if params == nil {
		params = &RadarPaymentEvaluationCreateParams{}
	}
	params.Context = ctx
	paymentevaluation := &RadarPaymentEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/payment_evaluations", c.Key, params, paymentevaluation)
	return paymentevaluation, err
}
