//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

// Type of payment object associated with this invoice payment.
type InvoicePaymentPaymentType string

// List of values that InvoicePaymentPaymentType can take
const (
	InvoicePaymentPaymentTypeCharge           InvoicePaymentPaymentType = "charge"
	InvoicePaymentPaymentTypeOutOfBandPayment InvoicePaymentPaymentType = "out_of_band_payment"
	InvoicePaymentPaymentTypePaymentIntent    InvoicePaymentPaymentType = "payment_intent"
	InvoicePaymentPaymentTypePaymentRecord    InvoicePaymentPaymentType = "payment_record"
)

// When retrieving an invoice, there is an includable payments property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of payments.
type InvoicePaymentListParams struct {
	ListParams `form:"*"`
	Invoice    *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *InvoicePaymentListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the invoice payment with the given ID.
type InvoicePaymentParams struct {
	Params  `form:"*"`
	Invoice *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *InvoicePaymentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type InvoicePaymentPayment struct {
	// ID of the successful charge for this payment when `type` is `charge`.
	Charge *Charge `json:"charge"`
	// ID of the PaymentIntent associated with this payment when `type` is `payment_intent`. Note: This property is only populated for invoices finalized on or after March 15th, 2019.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
	// ID of the PaymentRecord associated with this payment when `type` is `payment_record`.
	PaymentRecord *PaymentRecord `json:"payment_record"`
	// Type of payment object associated with this invoice payment.
	Type InvoicePaymentPaymentType `json:"type"`
}
type InvoicePaymentStatusTransitions struct {
	// The time that the payment was canceled.
	CanceledAt time.Time `json:"canceled_at"`
	// The time that the payment succeeded.
	PaidAt time.Time `json:"paid_at"`
}

// The invoice payment object
type InvoicePayment struct {
	APIResource
	// Excess payment that was received for this invoice and credited to the customer's `invoice_credit_balance`. This field is null until the payment is `paid`. Overpayment can happen when you attach more than one PaymentIntent to the invoice, and each of them succeeds. To avoid overpayment, cancel any PaymentIntents that you do not need before attaching more.
	AmountOverpaid int64 `json:"amount_overpaid"`
	// Amount that was actually paid for this invoice, in cents (or local equivalent). This field is null until the payment is `paid`. This amount can be less than the `amount_requested` if the PaymentIntent's `amount_received` is not sufficient to pay all of the invoices that it is attached to.
	AmountPaid int64 `json:"amount_paid"`
	// Amount intended to be paid toward this invoice, in cents (or local equivalent)
	AmountRequested int64 `json:"amount_requested"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created time.Time `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The invoice that was paid.
	Invoice *Invoice `json:"invoice"`
	// Stripe automatically creates a default InvoicePayment when the invoice is finalized, and keeps it synchronized with the invoice's `amount_remaining`. The PaymentIntent associated with the default payment can't be edited or canceled directly.
	IsDefault bool `json:"is_default"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object  string                 `json:"object"`
	Payment *InvoicePaymentPayment `json:"payment"`
	// The status of the payment, one of `open`, `paid`, or `canceled`.
	Status            string                           `json:"status"`
	StatusTransitions *InvoicePaymentStatusTransitions `json:"status_transitions"`
}

// InvoicePaymentList is a list of InvoicePayments as retrieved from a list endpoint.
type InvoicePaymentList struct {
	APIResource
	ListMeta
	Data []*InvoicePayment `json:"data"`
}

// UnmarshalJSON handles deserialization of an InvoicePaymentStatusTransitions.
// This custom unmarshaling is needed to handle the time fields correctly.
func (i *InvoicePaymentStatusTransitions) UnmarshalJSON(data []byte) error {
	type invoicePaymentStatusTransitions InvoicePaymentStatusTransitions
	v := struct {
		CanceledAt int64 `json:"canceled_at"`
		PaidAt     int64 `json:"paid_at"`
		*invoicePaymentStatusTransitions
	}{
		invoicePaymentStatusTransitions: (*invoicePaymentStatusTransitions)(i),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	i.CanceledAt = time.Unix(v.CanceledAt, 0)
	i.PaidAt = time.Unix(v.PaidAt, 0)
	return nil
}

// UnmarshalJSON handles deserialization of an InvoicePayment.
// This custom unmarshaling is needed to handle the time fields correctly.
func (i *InvoicePayment) UnmarshalJSON(data []byte) error {
	type invoicePayment InvoicePayment
	v := struct {
		Created int64 `json:"created"`
		*invoicePayment
	}{
		invoicePayment: (*invoicePayment)(i),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	i.Created = time.Unix(v.Created, 0)
	return nil
}

// MarshalJSON handles serialization of an InvoicePaymentStatusTransitions.
// This custom marshaling is needed to handle the time fields correctly.
func (i InvoicePaymentStatusTransitions) MarshalJSON() ([]byte, error) {
	type invoicePaymentStatusTransitions InvoicePaymentStatusTransitions
	v := struct {
		CanceledAt int64 `json:"canceled_at"`
		PaidAt     int64 `json:"paid_at"`
		invoicePaymentStatusTransitions
	}{
		invoicePaymentStatusTransitions: (invoicePaymentStatusTransitions)(i),
		CanceledAt:                      i.CanceledAt.Unix(),
		PaidAt:                          i.PaidAt.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}

// MarshalJSON handles serialization of an InvoicePayment.
// This custom marshaling is needed to handle the time fields correctly.
func (i InvoicePayment) MarshalJSON() ([]byte, error) {
	type invoicePayment InvoicePayment
	v := struct {
		Created int64 `json:"created"`
		invoicePayment
	}{
		invoicePayment: (invoicePayment)(i),
		Created:        i.Created.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
