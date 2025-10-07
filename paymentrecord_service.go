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

// v1PaymentRecordService is used to invoke /v1/payment_records APIs.
type v1PaymentRecordService struct {
	B   Backend
	Key string
}

// Retrieves a Payment Record with the given ID
func (c v1PaymentRecordService) Retrieve(ctx context.Context, id string, params *PaymentRecordRetrieveParams) (*PaymentRecord, error) {
	if params == nil {
		params = &PaymentRecordRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_records/%s", id)
	paymentrecord := &PaymentRecord{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report a new Payment Record. You may report a Payment Record as it is
//
//	initialized and later report updates through the other report_* methods, or report Payment
//	Records in a terminal state directly, through this method.
func (c v1PaymentRecordService) ReportPayment(ctx context.Context, params *PaymentRecordReportPaymentParams) (*PaymentRecord, error) {
	if params == nil {
		params = &PaymentRecordReportPaymentParams{}
	}
	params.Context = ctx
	paymentrecord := &PaymentRecord{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_records/report_payment", c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report a new payment attempt on the specified Payment Record. A new payment
//
//	attempt can only be specified if all other payment attempts are canceled or failed.
func (c v1PaymentRecordService) ReportPaymentAttempt(ctx context.Context, id string, params *PaymentRecordReportPaymentAttemptParams) (*PaymentRecord, error) {
	if params == nil {
		params = &PaymentRecordReportPaymentAttemptParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_records/%s/report_payment_attempt", id)
	paymentrecord := &PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was canceled.
func (c v1PaymentRecordService) ReportPaymentAttemptCanceled(ctx context.Context, id string, params *PaymentRecordReportPaymentAttemptCanceledParams) (*PaymentRecord, error) {
	if params == nil {
		params = &PaymentRecordReportPaymentAttemptCanceledParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/payment_records/%s/report_payment_attempt_canceled", id)
	paymentrecord := &PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	failed or errored.
func (c v1PaymentRecordService) ReportPaymentAttemptFailed(ctx context.Context, id string, params *PaymentRecordReportPaymentAttemptFailedParams) (*PaymentRecord, error) {
	if params == nil {
		params = &PaymentRecordReportPaymentAttemptFailedParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/payment_records/%s/report_payment_attempt_failed", id)
	paymentrecord := &PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was guaranteed.
func (c v1PaymentRecordService) ReportPaymentAttemptGuaranteed(ctx context.Context, id string, params *PaymentRecordReportPaymentAttemptGuaranteedParams) (*PaymentRecord, error) {
	if params == nil {
		params = &PaymentRecordReportPaymentAttemptGuaranteedParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/payment_records/%s/report_payment_attempt_guaranteed", id)
	paymentrecord := &PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report informational updates on the specified Payment Record.
func (c v1PaymentRecordService) ReportPaymentAttemptInformational(ctx context.Context, id string, params *PaymentRecordReportPaymentAttemptInformationalParams) (*PaymentRecord, error) {
	if params == nil {
		params = &PaymentRecordReportPaymentAttemptInformationalParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/payment_records/%s/report_payment_attempt_informational", id)
	paymentrecord := &PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was refunded.
func (c v1PaymentRecordService) ReportRefund(ctx context.Context, id string, params *PaymentRecordReportRefundParams) (*PaymentRecord, error) {
	if params == nil {
		params = &PaymentRecordReportRefundParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_records/%s/report_refund", id)
	paymentrecord := &PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}
