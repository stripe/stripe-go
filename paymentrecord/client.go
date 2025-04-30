//
//
// File generated from our OpenAPI spec
//
//

// Package paymentrecord provides the /v1/payment_records APIs
package paymentrecord

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
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

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
