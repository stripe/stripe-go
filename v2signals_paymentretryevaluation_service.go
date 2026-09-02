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

// v2SignalsPaymentRetryEvaluationService is used to invoke paymentretryevaluation related APIs.
type v2SignalsPaymentRetryEvaluationService struct {
	B   Backend
	Key string
}

// Creates a new payment retry evaluation for a failed payment.
func (c v2SignalsPaymentRetryEvaluationService) Create(ctx context.Context, params *V2SignalsPaymentRetryEvaluationCreateParams) (*V2SignalsPaymentRetryEvaluation, error) {
	if params == nil {
		params = &V2SignalsPaymentRetryEvaluationCreateParams{}
	}
	params.Context = ctx
	paymentretryevaluation := &V2SignalsPaymentRetryEvaluation{}
	err := c.B.Call(
		http.MethodPost, "/v2/signals/payment_retry_evaluations", c.Key, params, paymentretryevaluation)
	return paymentretryevaluation, err
}

// Retrieves a payment retry evaluation by ID.
func (c v2SignalsPaymentRetryEvaluationService) Retrieve(ctx context.Context, id string, params *V2SignalsPaymentRetryEvaluationRetrieveParams) (*V2SignalsPaymentRetryEvaluation, error) {
	if params == nil {
		params = &V2SignalsPaymentRetryEvaluationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/signals/payment_retry_evaluations/%s", id)
	paymentretryevaluation := &V2SignalsPaymentRetryEvaluation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentretryevaluation)
	return paymentretryevaluation, err
}

// Updates an active payment retry evaluation with a replacement payment identifier.
func (c v2SignalsPaymentRetryEvaluationService) Update(ctx context.Context, id string, params *V2SignalsPaymentRetryEvaluationUpdateParams) (*V2SignalsPaymentRetryEvaluation, error) {
	if params == nil {
		params = &V2SignalsPaymentRetryEvaluationUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/signals/payment_retry_evaluations/%s", id)
	paymentretryevaluation := &V2SignalsPaymentRetryEvaluation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentretryevaluation)
	return paymentretryevaluation, err
}

// Cancels an active payment retry evaluation.
func (c v2SignalsPaymentRetryEvaluationService) Cancel(ctx context.Context, id string, params *V2SignalsPaymentRetryEvaluationCancelParams) (*V2SignalsPaymentRetryEvaluation, error) {
	if params == nil {
		params = &V2SignalsPaymentRetryEvaluationCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/signals/payment_retry_evaluations/%s/cancel", id)
	paymentretryevaluation := &V2SignalsPaymentRetryEvaluation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentretryevaluation)
	return paymentretryevaluation, err
}
