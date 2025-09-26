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

// v2PaymentsOffSessionPaymentService is used to invoke offsessionpayment related APIs.
type v2PaymentsOffSessionPaymentService struct {
	B   Backend
	Key string
}

// Creates an OffSessionPayment object.
func (c v2PaymentsOffSessionPaymentService) Create(ctx context.Context, params *V2PaymentsOffSessionPaymentCreateParams) (*V2PaymentsOffSessionPayment, error) {
	if params == nil {
		params = &V2PaymentsOffSessionPaymentCreateParams{}
	}
	params.Context = ctx
	offsessionpayment := &V2PaymentsOffSessionPayment{}
	err := c.B.Call(
		http.MethodPost, "/v2/payments/off_session_payments", c.Key, params, offsessionpayment)
	return offsessionpayment, err
}

// Retrieves the details of an OffSessionPayment that has previously been created.
func (c v2PaymentsOffSessionPaymentService) Retrieve(ctx context.Context, id string, params *V2PaymentsOffSessionPaymentRetrieveParams) (*V2PaymentsOffSessionPayment, error) {
	if params == nil {
		params = &V2PaymentsOffSessionPaymentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/payments/off_session_payments/%s", id)
	offsessionpayment := &V2PaymentsOffSessionPayment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, offsessionpayment)
	return offsessionpayment, err
}

// Cancel an OffSessionPayment that has previously been created.
func (c v2PaymentsOffSessionPaymentService) Cancel(ctx context.Context, id string, params *V2PaymentsOffSessionPaymentCancelParams) (*V2PaymentsOffSessionPayment, error) {
	if params == nil {
		params = &V2PaymentsOffSessionPaymentCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/payments/off_session_payments/%s/cancel", id)
	offsessionpayment := &V2PaymentsOffSessionPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, offsessionpayment)
	return offsessionpayment, err
}

// Captures an OffSessionPayment that has previously been created.
func (c v2PaymentsOffSessionPaymentService) Capture(ctx context.Context, id string, params *V2PaymentsOffSessionPaymentCaptureParams) (*V2PaymentsOffSessionPayment, error) {
	if params == nil {
		params = &V2PaymentsOffSessionPaymentCaptureParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/payments/off_session_payments/%s/capture", id)
	offsessionpayment := &V2PaymentsOffSessionPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, offsessionpayment)
	return offsessionpayment, err
}

// Returns a list of OffSessionPayments matching a filter.
func (c v2PaymentsOffSessionPaymentService) List(ctx context.Context, listParams *V2PaymentsOffSessionPaymentListParams) Seq2[*V2PaymentsOffSessionPayment, error] {
	if listParams == nil {
		listParams = &V2PaymentsOffSessionPaymentListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/payments/off_session_payments", listParams, func(path string, p ParamsContainer) (*V2Page[*V2PaymentsOffSessionPayment], error) {
		page := &V2Page[*V2PaymentsOffSessionPayment]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
