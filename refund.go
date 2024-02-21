//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The type of refund. This can be `refund`, `reversal`, or `pending`.
type RefundDestinationDetailsCardType string

// List of values that RefundDestinationDetailsCardType can take
const (
	RefundDestinationDetailsCardTypePending  RefundDestinationDetailsCardType = "pending"
	RefundDestinationDetailsCardTypeRefund   RefundDestinationDetailsCardType = "refund"
	RefundDestinationDetailsCardTypeReversal RefundDestinationDetailsCardType = "reversal"
)

// Provides the reason for the refund failure. Possible values are: `lost_or_stolen_card`, `expired_or_canceled_card`, `charge_for_pending_refund_disputed`, `insufficient_funds`, `declined`, `merchant_request`, or `unknown`.
type RefundFailureReason string

// List of values that RefundFailureReason can take
const (
	RefundFailureReasonExpiredOrCanceledCard RefundFailureReason = "expired_or_canceled_card"
	RefundFailureReasonLostOrStolenCard      RefundFailureReason = "lost_or_stolen_card"
	RefundFailureReasonUnknown               RefundFailureReason = "unknown"
)

// Reason for the refund, which is either user-provided (`duplicate`, `fraudulent`, or `requested_by_customer`) or generated by Stripe internally (`expired_uncaptured_charge`).
type RefundReason string

// List of values that RefundReason can take
const (
	RefundReasonDuplicate               RefundReason = "duplicate"
	RefundReasonExpiredUncapturedCharge RefundReason = "expired_uncaptured_charge"
	RefundReasonFraudulent              RefundReason = "fraudulent"
	RefundReasonRequestedByCustomer     RefundReason = "requested_by_customer"
)

// Status of the refund. This can be `pending`, `requires_action`, `succeeded`, `failed`, or `canceled`. Learn more about [failed refunds](https://stripe.com/docs/refunds#failed-refunds).
type RefundStatus string

// List of values that RefundStatus can take
const (
	RefundStatusCanceled       RefundStatus = "canceled"
	RefundStatusFailed         RefundStatus = "failed"
	RefundStatusPending        RefundStatus = "pending"
	RefundStatusSucceeded      RefundStatus = "succeeded"
	RefundStatusRequiresAction RefundStatus = "requires_action"
)

// Returns a list of all refunds you created. We return the refunds in sorted order, with the most recent refunds appearing first The 10 most recent refunds are always available by default on the Charge object.
type RefundListParams struct {
	ListParams `form:"*"`
	// Only return refunds for the charge specified by this charge ID.
	Charge *string `form:"charge"`
	// Only return refunds that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return refunds that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return refunds for the PaymentIntent specified by this ID.
	PaymentIntent *string `form:"payment_intent"`
}

// AddExpand appends a new field to expand.
func (p *RefundListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// When you create a new refund, you must specify a Charge or a PaymentIntent object on which to create it.
//
// Creating a new refund will refund a charge that has previously been created but not yet refunded.
// Funds will be refunded to the credit or debit card that was originally charged.
//
// You can optionally refund only part of a charge.
// You can do so multiple times, until the entire charge has been refunded.
//
// Once entirely refunded, a charge can't be refunded again.
// This method will raise an error when called on an already-refunded charge,
// or when trying to refund more money than is left on a charge.
type RefundParams struct {
	Params `form:"*"`
	Amount *int64 `form:"amount"`
	// The identifier of the charge to refund.
	Charge *string `form:"charge"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Customer whose customer balance to refund from.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// For payment methods without native refund support (e.g., Konbini, PromptPay), use this email from the customer to receive refund instructions.
	InstructionsEmail *string `form:"instructions_email"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Origin of the refund
	Origin *string `form:"origin"`
	// The identifier of the PaymentIntent to refund.
	PaymentIntent *string `form:"payment_intent"`
	// String indicating the reason for the refund. If set, possible values are `duplicate`, `fraudulent`, and `requested_by_customer`. If you believe the charge to be fraudulent, specifying `fraudulent` as the reason will add the associated card and email to your [block lists](https://stripe.com/docs/radar/lists), and will also help us improve our fraud detection algorithms.
	Reason *string `form:"reason"`
	// Boolean indicating whether the application fee should be refunded when refunding this charge. If a full charge refund is given, the full application fee will be refunded. Otherwise, the application fee will be refunded in an amount proportional to the amount of the charge refunded. An application fee can be refunded only by the application that created the charge.
	RefundApplicationFee *bool `form:"refund_application_fee"`
	// Boolean indicating whether the transfer should be reversed when refunding this charge. The transfer will be reversed proportionally to the amount being refunded (either the entire or partial amount).
	//
	// A transfer can be reversed only by the application that created the charge.
	ReverseTransfer *bool `form:"reverse_transfer"`
}

// AddExpand appends a new field to expand.
func (p *RefundParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *RefundParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancels a refund with a status of requires_action.
//
// You can't cancel refunds in other states. Only refunds for payment methods that require customer action can enter the requires_action state.
type RefundCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *RefundCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type RefundDestinationDetailsAffirm struct{}
type RefundDestinationDetailsAfterpayClearpay struct{}
type RefundDestinationDetailsAlipay struct{}
type RefundDestinationDetailsAuBankTransfer struct{}
type RefundDestinationDetailsBLIK struct {
	// The reference assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
}
type RefundDestinationDetailsBrBankTransfer struct {
	// The reference assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
}
type RefundDestinationDetailsCard struct {
	// Value of the reference number assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference number on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
	// Type of the reference number assigned to the refund.
	ReferenceType string `json:"reference_type"`
	// The type of refund. This can be `refund`, `reversal`, or `pending`.
	Type RefundDestinationDetailsCardType `json:"type"`
}
type RefundDestinationDetailsCashApp struct{}
type RefundDestinationDetailsCustomerCashBalance struct{}
type RefundDestinationDetailsEPS struct{}
type RefundDestinationDetailsEUBankTransfer struct {
	// The reference assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
}
type RefundDestinationDetailsGBBankTransfer struct {
	// The reference assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
}
type RefundDestinationDetailsGiropay struct{}
type RefundDestinationDetailsGrabpay struct{}
type RefundDestinationDetailsJPBankTransfer struct {
	// The reference assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
}
type RefundDestinationDetailsKlarna struct{}
type RefundDestinationDetailsMXBankTransfer struct {
	// The reference assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
}
type RefundDestinationDetailsP24 struct {
	// The reference assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
}
type RefundDestinationDetailsPayNow struct{}
type RefundDestinationDetailsPaypal struct{}
type RefundDestinationDetailsPix struct{}
type RefundDestinationDetailsRevolut struct{}
type RefundDestinationDetailsSofort struct{}
type RefundDestinationDetailsSwish struct {
	// The reference assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
}
type RefundDestinationDetailsTHBankTransfer struct {
	// The reference assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
}
type RefundDestinationDetailsUSBankTransfer struct {
	// The reference assigned to the refund.
	Reference string `json:"reference"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus string `json:"reference_status"`
}
type RefundDestinationDetailsWeChatPay struct{}
type RefundDestinationDetailsZip struct{}
type RefundDestinationDetails struct {
	Affirm              *RefundDestinationDetailsAffirm              `json:"affirm"`
	AfterpayClearpay    *RefundDestinationDetailsAfterpayClearpay    `json:"afterpay_clearpay"`
	Alipay              *RefundDestinationDetailsAlipay              `json:"alipay"`
	AuBankTransfer      *RefundDestinationDetailsAuBankTransfer      `json:"au_bank_transfer"`
	BLIK                *RefundDestinationDetailsBLIK                `json:"blik"`
	BrBankTransfer      *RefundDestinationDetailsBrBankTransfer      `json:"br_bank_transfer"`
	Card                *RefundDestinationDetailsCard                `json:"card"`
	CashApp             *RefundDestinationDetailsCashApp             `json:"cashapp"`
	CustomerCashBalance *RefundDestinationDetailsCustomerCashBalance `json:"customer_cash_balance"`
	EPS                 *RefundDestinationDetailsEPS                 `json:"eps"`
	EUBankTransfer      *RefundDestinationDetailsEUBankTransfer      `json:"eu_bank_transfer"`
	GBBankTransfer      *RefundDestinationDetailsGBBankTransfer      `json:"gb_bank_transfer"`
	Giropay             *RefundDestinationDetailsGiropay             `json:"giropay"`
	Grabpay             *RefundDestinationDetailsGrabpay             `json:"grabpay"`
	JPBankTransfer      *RefundDestinationDetailsJPBankTransfer      `json:"jp_bank_transfer"`
	Klarna              *RefundDestinationDetailsKlarna              `json:"klarna"`
	MXBankTransfer      *RefundDestinationDetailsMXBankTransfer      `json:"mx_bank_transfer"`
	P24                 *RefundDestinationDetailsP24                 `json:"p24"`
	PayNow              *RefundDestinationDetailsPayNow              `json:"paynow"`
	Paypal              *RefundDestinationDetailsPaypal              `json:"paypal"`
	Pix                 *RefundDestinationDetailsPix                 `json:"pix"`
	Revolut             *RefundDestinationDetailsRevolut             `json:"revolut"`
	Sofort              *RefundDestinationDetailsSofort              `json:"sofort"`
	Swish               *RefundDestinationDetailsSwish               `json:"swish"`
	THBankTransfer      *RefundDestinationDetailsTHBankTransfer      `json:"th_bank_transfer"`
	// The type of transaction-specific details of the payment method used in the refund (e.g., `card`). An additional hash is included on `destination_details` with a name matching this value. It contains information specific to the refund transaction.
	Type           string                                  `json:"type"`
	USBankTransfer *RefundDestinationDetailsUSBankTransfer `json:"us_bank_transfer"`
	WeChatPay      *RefundDestinationDetailsWeChatPay      `json:"wechat_pay"`
	Zip            *RefundDestinationDetailsZip            `json:"zip"`
}
type RefundNextActionDisplayDetailsEmailSent struct {
	// The timestamp when the email was sent.
	EmailSentAt int64 `json:"email_sent_at"`
	// The recipient's email address.
	EmailSentTo string `json:"email_sent_to"`
}

// Contains the refund details.
type RefundNextActionDisplayDetails struct {
	EmailSent *RefundNextActionDisplayDetailsEmailSent `json:"email_sent"`
	// The expiry timestamp.
	ExpiresAt int64 `json:"expires_at"`
}
type RefundNextAction struct {
	// Contains the refund details.
	DisplayDetails *RefundNextActionDisplayDetails `json:"display_details"`
	// Type of the next action to perform.
	Type string `json:"type"`
}

// Refund objects allow you to refund a previously created charge that isn't
// refunded yet. Funds are refunded to the credit or debit card that's
// initially charged.
//
// Related guide: [Refunds](https://stripe.com/docs/refunds)
type Refund struct {
	APIResource
	// Amount, in cents (or local equivalent).
	Amount int64 `json:"amount"`
	// Balance transaction that describes the impact on your account balance.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	// ID of the charge that's refunded.
	Charge *Charge `json:"charge"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// An arbitrary string attached to the object. You can use this for displaying to users (available on non-card refunds only).
	Description        string                    `json:"description"`
	DestinationDetails *RefundDestinationDetails `json:"destination_details"`
	// After the refund fails, this balance transaction describes the adjustment made on your account balance that reverses the initial balance transaction.
	FailureBalanceTransaction *BalanceTransaction `json:"failure_balance_transaction"`
	// Provides the reason for the refund failure. Possible values are: `lost_or_stolen_card`, `expired_or_canceled_card`, `charge_for_pending_refund_disputed`, `insufficient_funds`, `declined`, `merchant_request`, or `unknown`.
	FailureReason RefundFailureReason `json:"failure_reason"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// For payment methods without native refund support (for example, Konbini, PromptPay), provide an email address for the customer to receive refund instructions.
	InstructionsEmail string `json:"instructions_email"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata   map[string]string `json:"metadata"`
	NextAction *RefundNextAction `json:"next_action"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// ID of the PaymentIntent that's refunded.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
	// Reason for the refund, which is either user-provided (`duplicate`, `fraudulent`, or `requested_by_customer`) or generated by Stripe internally (`expired_uncaptured_charge`).
	Reason RefundReason `json:"reason"`
	// This is the transaction number that appears on email receipts sent for this refund.
	ReceiptNumber string `json:"receipt_number"`
	// The transfer reversal that's associated with the refund. Only present if the charge came from another Stripe account.
	SourceTransferReversal *TransferReversal `json:"source_transfer_reversal"`
	// Status of the refund. This can be `pending`, `requires_action`, `succeeded`, `failed`, or `canceled`. Learn more about [failed refunds](https://stripe.com/docs/refunds#failed-refunds).
	Status RefundStatus `json:"status"`
	// This refers to the transfer reversal object if the accompanying transfer reverses. This is only applicable if the charge was created using the destination parameter.
	TransferReversal *TransferReversal `json:"transfer_reversal"`
}

// RefundList is a list of Refunds as retrieved from a list endpoint.
type RefundList struct {
	APIResource
	ListMeta
	Data []*Refund `json:"data"`
}

// UnmarshalJSON handles deserialization of a Refund.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *Refund) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		r.ID = id
		return nil
	}

	type refund Refund
	var v refund
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*r = Refund(v)
	return nil
}
