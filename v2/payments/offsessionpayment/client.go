//
//
// File generated from our OpenAPI spec
//
//

// Package offsessionpayment provides the offsessionpayment related APIs
package offsessionpayment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke offsessionpayment related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an OffSessionPayment object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2PaymentsOffSessionPaymentParams) (*stripe.V2PaymentsOffSessionPayment, error) {
	offsessionpayment := &stripe.V2PaymentsOffSessionPayment{}
	err := c.B.Call(
		http.MethodPost, "/v2/payments/off_session_payments", c.Key, params, offsessionpayment)
	return offsessionpayment, err
}

// Retrieves the details of an OffSessionPayment that has previously been created.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2PaymentsOffSessionPaymentParams) (*stripe.V2PaymentsOffSessionPayment, error) {
	path := stripe.FormatURLPath("/v2/payments/off_session_payments/%s", id)
	offsessionpayment := &stripe.V2PaymentsOffSessionPayment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, offsessionpayment)
	return offsessionpayment, err
}

// Cancel an OffSessionPayment that has previously been created.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2PaymentsOffSessionPaymentCancelParams) (*stripe.V2PaymentsOffSessionPayment, error) {
	path := stripe.FormatURLPath(
		"/v2/payments/off_session_payments/%s/cancel", id)
	offsessionpayment := &stripe.V2PaymentsOffSessionPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, offsessionpayment)
	return offsessionpayment, err
}

// Captures an OffSessionPayment that has previously been created.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Capture(id string, params *stripe.V2PaymentsOffSessionPaymentCaptureParams) (*stripe.V2PaymentsOffSessionPayment, error) {
	path := stripe.FormatURLPath(
		"/v2/payments/off_session_payments/%s/capture", id)
	offsessionpayment := &stripe.V2PaymentsOffSessionPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, offsessionpayment)
	return offsessionpayment, err
}

// Returns a list of OffSessionPayments matching a filter.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2PaymentsOffSessionPaymentListParams) stripe.Seq2[*stripe.V2PaymentsOffSessionPayment, error] {
	return stripe.NewV2List("/v2/payments/off_session_payments", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2PaymentsOffSessionPayment], error) {
		page := &stripe.V2Page[*stripe.V2PaymentsOffSessionPayment]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
