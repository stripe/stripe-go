//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of account liable for tax.
type V2PaymentsOffSessionPaymentAutomaticTaxLiabilityType string

// List of values that V2PaymentsOffSessionPaymentAutomaticTaxLiabilityType can take
const (
	V2PaymentsOffSessionPaymentAutomaticTaxLiabilityTypeSelf V2PaymentsOffSessionPaymentAutomaticTaxLiabilityType = "self"
)

// The reason why the OffSessionPayment failed.
type V2PaymentsOffSessionPaymentFailureReason string

// List of values that V2PaymentsOffSessionPaymentFailureReason can take
const (
	V2PaymentsOffSessionPaymentFailureReasonCustomerInvalid      V2PaymentsOffSessionPaymentFailureReason = "customer_invalid"
	V2PaymentsOffSessionPaymentFailureReasonExceededRetryWindow  V2PaymentsOffSessionPaymentFailureReason = "exceeded_retry_window"
	V2PaymentsOffSessionPaymentFailureReasonNoValidPaymentMethod V2PaymentsOffSessionPaymentFailureReason = "no_valid_payment_method"
	V2PaymentsOffSessionPaymentFailureReasonRejectedByPartner    V2PaymentsOffSessionPaymentFailureReason = "rejected_by_partner"
	V2PaymentsOffSessionPaymentFailureReasonRetriesExhausted     V2PaymentsOffSessionPaymentFailureReason = "retries_exhausted"
)

// The type of item. Currently, only `one_time_price` is supported.
type V2PaymentsOffSessionPaymentItemType string

// List of values that V2PaymentsOffSessionPaymentItemType can take
const (
	V2PaymentsOffSessionPaymentItemTypeOneTimePrice V2PaymentsOffSessionPaymentItemType = "one_time_price"
)

// Status of this OffSessionPayment, one of `pending_initial_attempt`, `pending_retry`,
// `processing`, `failed`, `canceled`, `succeeded`, or `paused`.
type V2PaymentsOffSessionPaymentStatus string

// List of values that V2PaymentsOffSessionPaymentStatus can take
const (
	V2PaymentsOffSessionPaymentStatusCanceled              V2PaymentsOffSessionPaymentStatus = "canceled"
	V2PaymentsOffSessionPaymentStatusFailed                V2PaymentsOffSessionPaymentStatus = "failed"
	V2PaymentsOffSessionPaymentStatusPaused                V2PaymentsOffSessionPaymentStatus = "paused"
	V2PaymentsOffSessionPaymentStatusPendingInitialAttempt V2PaymentsOffSessionPaymentStatus = "pending_initial_attempt"
	V2PaymentsOffSessionPaymentStatusPendingRetry          V2PaymentsOffSessionPaymentStatus = "pending_retry"
	V2PaymentsOffSessionPaymentStatusProcessing            V2PaymentsOffSessionPaymentStatus = "processing"
	V2PaymentsOffSessionPaymentStatusSucceeded             V2PaymentsOffSessionPaymentStatus = "succeeded"
)

// The account liable for tax. Null when automatic tax is disabled; otherwise,
// an object with type set to `self`, indicating the account making the API request.
type V2PaymentsOffSessionPaymentAutomaticTaxLiability struct {
	// The type of account liable for tax.
	Type V2PaymentsOffSessionPaymentAutomaticTaxLiabilityType `json:"type"`
}

// Details about automatic tax calculation.
type V2PaymentsOffSessionPaymentAutomaticTax struct {
	// Whether automatic tax calculation is enabled for this OffSessionPayment.
	Enabled bool `json:"enabled"`
	// The account liable for tax. Null when automatic tax is disabled; otherwise,
	// an object with type set to `self`, indicating the account making the API request.
	Liability *V2PaymentsOffSessionPaymentAutomaticTaxLiability `json:"liability,omitempty"`
}

// The resolved price IDs and quantities in their original order.
type V2PaymentsOffSessionPaymentItemOneTimePriceItem struct {
	// The ID of the resolved one-time Price.
	Price string `json:"price"`
	// Quantity of this price in the item group.
	Quantity int64 `json:"quantity"`
}

// The persisted one-time prices in this item group.
type V2PaymentsOffSessionPaymentItemOneTimePrice struct {
	// A description of this item group.
	Description string `json:"description,omitempty"`
	// The resolved price IDs and quantities in their original order.
	Items []*V2PaymentsOffSessionPaymentItemOneTimePriceItem `json:"items"`
	// Additional information about this item group.
	Metadata map[string]string `json:"metadata,omitempty"`
}

// The persisted item groups for this OffSessionPayment, in their original order. Null for raw-amount payments.
type V2PaymentsOffSessionPaymentItem struct {
	// The persisted one-time prices in this item group.
	OneTimePrice *V2PaymentsOffSessionPaymentItemOneTimePrice `json:"one_time_price"`
	// The type of item. Currently, only `one_time_price` is supported.
	Type V2PaymentsOffSessionPaymentItemType `json:"type"`
}

// Retry configuration and attempts for this OffSessionPayment.
type V2PaymentsOffSessionPaymentRetries struct {
	// Number of authorization attempts so far.
	Attempts int64 `json:"attempts"`
	// Whether retries are enabled for this OffSessionPayment.
	Enabled bool `json:"enabled"`
}

// Subtotal and tax for an itemized OffSessionPayment. Null for raw-amount payments.
type V2PaymentsOffSessionPaymentTotalDetails struct {
	// The subtotal of the items, including any inclusive tax and excluding exclusive tax.
	AmountSubtotal Amount `json:"amount_subtotal"`
	// The calculated tax amount, including zero. Null when automatic tax was disabled.
	AmountTax Amount `json:"amount_tax,omitempty"`
}

// An OffSessionPayment allows you to create your own payment flow for recurring and unscheduled payments. In a single API request, you can
// initiate a payment and leverage Stripe's AI-powered [Smart Retries](https://docs.stripe.com/payments/off-session-payments/smart-retries)
// to choose the best times to retry failures to increase the chance of a successful payment.
// To properly interface with OffSessionPayments, we recommend setting up an event destination to listen to webhook notifications, as the
// OffSessionPayment will transition through its lifecycle asynchronously.
// Related guide: [Off-Session Payments API](https://docs.stripe.com/payments/off-session-payments).
type V2PaymentsOffSessionPayment struct {
	APIResource
	// The total amount to collect, including tax when applicable.
	AmountTotal Amount `json:"amount_total"`
	// Details about automatic tax calculation.
	AutomaticTax *V2PaymentsOffSessionPaymentAutomaticTax `json:"automatic_tax,omitempty"`
	// Creation time of the OffSessionPayment. Represented as a RFC 3339 date & time UTC
	// value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Three-letter ISO currency code, in lowercase.
	Currency Currency `json:"currency"`
	// ID of the Customer to which this OffSessionPayment belongs, if one exists.
	Customer string `json:"customer,omitempty"`
	// ID of the Account representing the customer to which this OffSessionPayment belongs, if one exists.
	CustomerAccount string `json:"customer_account,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// The reason why the OffSessionPayment failed.
	FailureReason V2PaymentsOffSessionPaymentFailureReason `json:"failure_reason,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The persisted item groups for this OffSessionPayment, in their original order. Null for raw-amount payments.
	Items []*V2PaymentsOffSessionPaymentItem `json:"items,omitempty"`
	// Payment attempt record for the latest attempt, if one exists.
	LatestPaymentAttemptRecord string `json:"latest_payment_attempt_record,omitempty"`
	// Has the value true if the object exists in live mode or the value false if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.stripe.com/payments/payment-intents#storing-information-in-metadata).
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// ID of the payment method used in this OffSessionPayment.
	PaymentMethod string `json:"payment_method"`
	// Payment record associated with the OffSessionPayment.
	PaymentRecord string `json:"payment_record,omitempty"`
	// Retry configuration and attempts for this OffSessionPayment.
	Retries *V2PaymentsOffSessionPaymentRetries `json:"retries"`
	// Text that appears on the customer's statement as the statement descriptor for a
	// non-card charge. This value overrides the account's default statement descriptor.
	// For information about requirements, including the 22-character limit, see the
	// [Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's
	// [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static)
	// to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix string `json:"statement_descriptor_suffix,omitempty"`
	// Status of this OffSessionPayment, one of `pending_initial_attempt`, `pending_retry`,
	// `processing`, `failed`, `canceled`, `succeeded`, or `paused`.
	Status V2PaymentsOffSessionPaymentStatus `json:"status"`
	// Provides a way to schedule deferred payments for cards and debits, in YYYY-MM-DD format.
	TargetDate string `json:"target_date,omitempty"`
	// Test clock that can be used to advance the retry attempts in a sandbox.
	TestClock string `json:"test_clock,omitempty"`
	// Subtotal and tax for an itemized OffSessionPayment. Null for raw-amount payments.
	TotalDetails *V2PaymentsOffSessionPaymentTotalDetails `json:"total_details,omitempty"`
}
