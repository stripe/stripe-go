//
//
// File generated from our OpenAPI spec
//
//

// Package paymentrecord provides the /v1/payment_records APIs
package paymentrecord

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/payment_records APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Payment Record with the given ID
func Get(id string, params *stripe.PaymentRecordParams) (*stripe.PaymentRecord, error) {
	return getC().Get(id, params)
}

// Retrieves a Payment Record with the given ID
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.PaymentRecordParams) (*stripe.PaymentRecord, error) {
	path := stripe.FormatURLPath("/v1/payment_records/%s", id)
	paymentrecord := &stripe.PaymentRecord{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report a new Payment Record. You may report a Payment Record as it is
//
//	initialized and later report updates through the other report_* methods, or report Payment
//	Records in a terminal state directly, through this method.
func ReportPayment(params *stripe.PaymentRecordReportPaymentParams) (*stripe.PaymentRecord, error) {
	return getC().ReportPayment(params)
}

// Report a new Payment Record. You may report a Payment Record as it is
//
//	initialized and later report updates through the other report_* methods, or report Payment
//	Records in a terminal state directly, through this method.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPayment(params *stripe.PaymentRecordReportPaymentParams) (*stripe.PaymentRecord, error) {
	paymentrecord := &stripe.PaymentRecord{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_records/report_payment", c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report a new payment attempt on the specified Payment Record. A new payment
//
//	attempt can only be specified if all other payment attempts are canceled or failed.
func ReportPaymentAttempt(id string, params *stripe.PaymentRecordReportPaymentAttemptParams) (*stripe.PaymentRecord, error) {
	return getC().ReportPaymentAttempt(id, params)
}

// Report a new payment attempt on the specified Payment Record. A new payment
//
//	attempt can only be specified if all other payment attempts are canceled or failed.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPaymentAttempt(id string, params *stripe.PaymentRecordReportPaymentAttemptParams) (*stripe.PaymentRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_records/%s/report_payment_attempt", id)
	paymentrecord := &stripe.PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was canceled.
func ReportPaymentAttemptCanceled(id string, params *stripe.PaymentRecordReportPaymentAttemptCanceledParams) (*stripe.PaymentRecord, error) {
	return getC().ReportPaymentAttemptCanceled(id, params)
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was canceled.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPaymentAttemptCanceled(id string, params *stripe.PaymentRecordReportPaymentAttemptCanceledParams) (*stripe.PaymentRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_records/%s/report_payment_attempt_canceled", id)
	paymentrecord := &stripe.PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	failed or errored.
func ReportPaymentAttemptFailed(id string, params *stripe.PaymentRecordReportPaymentAttemptFailedParams) (*stripe.PaymentRecord, error) {
	return getC().ReportPaymentAttemptFailed(id, params)
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	failed or errored.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPaymentAttemptFailed(id string, params *stripe.PaymentRecordReportPaymentAttemptFailedParams) (*stripe.PaymentRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_records/%s/report_payment_attempt_failed", id)
	paymentrecord := &stripe.PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was guaranteed.
func ReportPaymentAttemptGuaranteed(id string, params *stripe.PaymentRecordReportPaymentAttemptGuaranteedParams) (*stripe.PaymentRecord, error) {
	return getC().ReportPaymentAttemptGuaranteed(id, params)
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was guaranteed.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPaymentAttemptGuaranteed(id string, params *stripe.PaymentRecordReportPaymentAttemptGuaranteedParams) (*stripe.PaymentRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_records/%s/report_payment_attempt_guaranteed", id)
	paymentrecord := &stripe.PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report informational updates on the specified Payment Record.
func ReportPaymentAttemptInformational(id string, params *stripe.PaymentRecordReportPaymentAttemptInformationalParams) (*stripe.PaymentRecord, error) {
	return getC().ReportPaymentAttemptInformational(id, params)
}

// Report informational updates on the specified Payment Record.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportPaymentAttemptInformational(id string, params *stripe.PaymentRecordReportPaymentAttemptInformationalParams) (*stripe.PaymentRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/payment_records/%s/report_payment_attempt_informational", id)
	paymentrecord := &stripe.PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was refunded.
func ReportRefund(id string, params *stripe.PaymentRecordReportRefundParams) (*stripe.PaymentRecord, error) {
	return getC().ReportRefund(id, params)
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was refunded.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportRefund(id string, params *stripe.PaymentRecordReportRefundParams) (*stripe.PaymentRecord, error) {
	path := stripe.FormatURLPath("/v1/payment_records/%s/report_refund", id)
	paymentrecord := &stripe.PaymentRecord{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentrecord)
	return paymentrecord, err
}

// Search for PaymentRecords you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func Search(params *stripe.PaymentRecordSearchParams) *SearchIter {
	return getC().Search(params)
}

// Search for PaymentRecords you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Search(params *stripe.PaymentRecordSearchParams) *SearchIter {
	return &SearchIter{
		SearchIter: stripe.GetSearchIter(params, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.SearchContainer, error) {
			list := &stripe.PaymentRecordSearchResult{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_records/search", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// SearchIter is an iterator for payment records.
type SearchIter struct {
	*stripe.SearchIter
}

// PaymentRecord returns the payment record which the iterator is currently pointing to.
func (i *SearchIter) PaymentRecord() *stripe.PaymentRecord {
	return i.Current().(*stripe.PaymentRecord)
}

// PaymentRecordSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) PaymentRecordSearchResult() *stripe.PaymentRecordSearchResult {
	return i.SearchResult().(*stripe.PaymentRecordSearchResult)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
