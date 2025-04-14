//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Type of payment object associated with this invoice payment.
type InvoicePaymentPaymentType string

// List of values that InvoicePaymentPaymentType can take
const (
	InvoicePaymentPaymentTypeCharge        InvoicePaymentPaymentType = "charge"
	InvoicePaymentPaymentTypePaymentIntent InvoicePaymentPaymentType = "payment_intent"
)

// The payment details of the invoice payments to return.
type InvoicePaymentListPaymentParams struct {
	// Only return invoice payments associated by this payment intent ID.
	PaymentIntent *string `form:"payment_intent"`
	// Only return invoice payments associated by this payment type.
	Type *string `form:"type"`
}

// When retrieving an invoice, there is an includable payments property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of payments.
type InvoicePaymentListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The identifier of the invoice whose payments to return.
	Invoice *string `form:"invoice"`
	// The payment details of the invoice payments to return.
	Payment *InvoicePaymentListPaymentParams `form:"payment"`
	// The status of the invoice payments to return.
	Status *string `form:"status"`
}

// AddExpand appends a new field to expand.
func (p *InvoicePaymentListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the invoice payment with the given ID.
type InvoicePaymentParams struct {
	Params `form:"*"`
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
	// Type of payment object associated with this invoice payment.
	Type InvoicePaymentPaymentType `json:"type"`
}
type InvoicePaymentStatusTransitions struct {
	// The time that the payment was canceled.
	CanceledAt int64 `json:"canceled_at"`
	// The time that the payment succeeded.
	PaidAt int64 `json:"paid_at"`
}

// The invoice payment object
type InvoicePayment struct {
	APIResource
	// Amount that was actually paid for this invoice, in cents (or local equivalent). This field is null until the payment is `paid`. This amount can be less than the `amount_requested` if the PaymentIntent's `amount_received` is not sufficient to pay all of the invoices that it is attached to.
	AmountPaid int64 `json:"amount_paid"`
	// Amount intended to be paid toward this invoice, in cents (or local equivalent)
	AmountRequested int64 `json:"amount_requested"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
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
