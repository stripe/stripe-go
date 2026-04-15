//
//
// File generated from our OpenAPI spec
//
//

// Package paymentattemptrecord provides the /v1/payment_attempt_records APIs
package paymentattemptrecord

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
	"github.com/stripe/stripe-go/v85/form"
)

// Client is used to invoke /v1/payment_attempt_records APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Payment Attempt Record with the given ID
func Get(id string, params *stripe.PaymentAttemptRecordParams) (*stripe.PaymentAttemptRecord, error) {
	return getC().Get(id, params)
}

// Retrieves a Payment Attempt Record with the given ID
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.PaymentAttemptRecordParams) (*stripe.PaymentAttemptRecord, error) {
	path := stripe.FormatURLPath("/v1/payment_attempt_records/%s", id)
	paymentattemptrecord := &stripe.PaymentAttemptRecord{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record was authenticated.
func ReportAuthenticated(id string, params *stripe.PaymentAttemptRecordReportAuthenticatedParams) (*stripe.PaymentAttemptRecord, error) {
	return getC().ReportAuthenticated(id, params)
}

// Report that the specified Payment Attempt Record was authenticated.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportAuthenticated(id string, params *stripe.PaymentAttemptRecordReportAuthenticatedParams) (*stripe.PaymentAttemptRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_attempt_records/%s/report_authenticated", id)
	paymentattemptrecord := &stripe.PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record was authorized.
func ReportAuthorized(id string, params *stripe.PaymentAttemptRecordReportAuthorizedParams) (*stripe.PaymentAttemptRecord, error) {
	return getC().ReportAuthorized(id, params)
}

// Report that the specified Payment Attempt Record was authorized.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportAuthorized(id string, params *stripe.PaymentAttemptRecordReportAuthorizedParams) (*stripe.PaymentAttemptRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_attempt_records/%s/report_authorized", id)
	paymentattemptrecord := &stripe.PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record was canceled.
func ReportCanceled(id string, params *stripe.PaymentAttemptRecordReportCanceledParams) (*stripe.PaymentAttemptRecord, error) {
	return getC().ReportCanceled(id, params)
}

// Report that the specified Payment Attempt Record was canceled.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportCanceled(id string, params *stripe.PaymentAttemptRecordReportCanceledParams) (*stripe.PaymentAttemptRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_attempt_records/%s/report_canceled", id)
	paymentattemptrecord := &stripe.PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record failed.
func ReportFailed(id string, params *stripe.PaymentAttemptRecordReportFailedParams) (*stripe.PaymentAttemptRecord, error) {
	return getC().ReportFailed(id, params)
}

// Report that the specified Payment Attempt Record failed.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportFailed(id string, params *stripe.PaymentAttemptRecordReportFailedParams) (*stripe.PaymentAttemptRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_attempt_records/%s/report_failed", id)
	paymentattemptrecord := &stripe.PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record was guaranteed.
func ReportGuaranteed(id string, params *stripe.PaymentAttemptRecordReportGuaranteedParams) (*stripe.PaymentAttemptRecord, error) {
	return getC().ReportGuaranteed(id, params)
}

// Report that the specified Payment Attempt Record was guaranteed.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportGuaranteed(id string, params *stripe.PaymentAttemptRecordReportGuaranteedParams) (*stripe.PaymentAttemptRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_attempt_records/%s/report_guaranteed", id)
	paymentattemptrecord := &stripe.PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report informational updates on the specified Payment Attempt Record.
func ReportInformational(id string, params *stripe.PaymentAttemptRecordReportInformationalParams) (*stripe.PaymentAttemptRecord, error) {
	return getC().ReportInformational(id, params)
}

// Report informational updates on the specified Payment Attempt Record.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportInformational(id string, params *stripe.PaymentAttemptRecordReportInformationalParams) (*stripe.PaymentAttemptRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_attempt_records/%s/report_informational", id)
	paymentattemptrecord := &stripe.PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// Report that the specified Payment Attempt Record was refunded.
func ReportRefund(id string, params *stripe.PaymentAttemptRecordReportRefundParams) (*stripe.PaymentAttemptRecord, error) {
	return getC().ReportRefund(id, params)
}

// Report that the specified Payment Attempt Record was refunded.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportRefund(id string, params *stripe.PaymentAttemptRecordReportRefundParams) (*stripe.PaymentAttemptRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_attempt_records/%s/report_refund", id)
	paymentattemptrecord := &stripe.PaymentAttemptRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// List all the Payment Attempt Records attached to the specified Payment Record.
func List(params *stripe.PaymentAttemptRecordListParams) *Iter {
	return getC().List(params)
}

// List all the Payment Attempt Records attached to the specified Payment Record.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PaymentAttemptRecordListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentAttemptRecordList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_attempt_records", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment attempt records.
type Iter struct {
	*stripe.Iter
}

// PaymentAttemptRecord returns the payment attempt record which the iterator is currently pointing to.
func (i *Iter) PaymentAttemptRecord() *stripe.PaymentAttemptRecord {
	return i.Current().(*stripe.PaymentAttemptRecord)
}

// PaymentAttemptRecordList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentAttemptRecordList() *stripe.PaymentAttemptRecordList {
	return i.List().(*stripe.PaymentAttemptRecordList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
