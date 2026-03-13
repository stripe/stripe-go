//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1PaymentAttemptRecordService is used to invoke /v1/payment_attempt_records APIs.
type v1PaymentAttemptRecordService struct {
	B   Backend
	Key string
}

// Retrieves a Payment Attempt Record with the given ID
func (c v1PaymentAttemptRecordService) Retrieve(ctx context.Context, id string, params *PaymentAttemptRecordRetrieveParams) (*PaymentAttemptRecord, error) {
	if params == nil {
		params = &PaymentAttemptRecordRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_attempt_records/%s", id)
	paymentattemptrecord := &PaymentAttemptRecord{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record was authenticated.
func (c v1PaymentAttemptRecordService) ReportAuthenticated(ctx context.Context, id string, params *PaymentAttemptRecordReportAuthenticatedParams) (*PaymentAttemptRecord, error) {
	if params == nil {
		params = &PaymentAttemptRecordReportAuthenticatedParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/payment_attempt_records/%s/report_authenticated", id)
	paymentattemptrecord := &PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record was canceled.
func (c v1PaymentAttemptRecordService) ReportCanceled(ctx context.Context, id string, params *PaymentAttemptRecordReportCanceledParams) (*PaymentAttemptRecord, error) {
	if params == nil {
		params = &PaymentAttemptRecordReportCanceledParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_attempt_records/%s/report_canceled", id)
	paymentattemptrecord := &PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record failed.
func (c v1PaymentAttemptRecordService) ReportFailed(ctx context.Context, id string, params *PaymentAttemptRecordReportFailedParams) (*PaymentAttemptRecord, error) {
	if params == nil {
		params = &PaymentAttemptRecordReportFailedParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_attempt_records/%s/report_failed", id)
	paymentattemptrecord := &PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record was guaranteed.
func (c v1PaymentAttemptRecordService) ReportGuaranteed(ctx context.Context, id string, params *PaymentAttemptRecordReportGuaranteedParams) (*PaymentAttemptRecord, error) {
	if params == nil {
		params = &PaymentAttemptRecordReportGuaranteedParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_attempt_records/%s/report_guaranteed", id)
	paymentattemptrecord := &PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report informational updates on the specified Payment Attempt Record.
func (c v1PaymentAttemptRecordService) ReportInformational(ctx context.Context, id string, params *PaymentAttemptRecordReportInformationalParams) (*PaymentAttemptRecord, error) {
	if params == nil {
		params = &PaymentAttemptRecordReportInformationalParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/payment_attempt_records/%s/report_informational", id)
	paymentattemptrecord := &PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record was refunded.
func (c v1PaymentAttemptRecordService) ReportRefund(ctx context.Context, id string, params *PaymentAttemptRecordReportRefundParams) (*PaymentAttemptRecord, error) {
	if params == nil {
		params = &PaymentAttemptRecordReportRefundParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_attempt_records/%s/report_refund", id)
	paymentattemptrecord := &PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// List all the Payment Attempt Records attached to the specified Payment Record.
func (c v1PaymentAttemptRecordService) List(ctx context.Context, listParams *PaymentAttemptRecordListParams) Seq2[*PaymentAttemptRecord, error] {
	if listParams == nil {
		listParams = &PaymentAttemptRecordListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*PaymentAttemptRecord], error) {
		list := &v1Page[*PaymentAttemptRecord]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_attempt_records", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
