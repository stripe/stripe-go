//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The frequency of the underlying payment.
type V2PaymentsOffSessionPaymentCadence string

// List of values that V2PaymentsOffSessionPaymentCadence can take
const (
	V2PaymentsOffSessionPaymentCadenceRecurring   V2PaymentsOffSessionPaymentCadence = "recurring"
	V2PaymentsOffSessionPaymentCadenceUnscheduled V2PaymentsOffSessionPaymentCadence = "unscheduled"
)

// The reason why the OffSessionPayment failed.
type V2PaymentsOffSessionPaymentFailureReason string

// List of values that V2PaymentsOffSessionPaymentFailureReason can take
const (
	V2PaymentsOffSessionPaymentFailureReasonAuthorizationExpired V2PaymentsOffSessionPaymentFailureReason = "authorization_expired"
	V2PaymentsOffSessionPaymentFailureReasonRejectedByPartner    V2PaymentsOffSessionPaymentFailureReason = "rejected_by_partner"
	V2PaymentsOffSessionPaymentFailureReasonRetriesExhausted     V2PaymentsOffSessionPaymentFailureReason = "retries_exhausted"
)

// Indicates the strategy for how you want Stripe to retry the payment.
type V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy string

// List of values that V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy can take
const (
	V2PaymentsOffSessionPaymentRetryDetailsRetryStrategyHeuristic V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy = "heuristic"
	V2PaymentsOffSessionPaymentRetryDetailsRetryStrategyNone      V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy = "none"
	V2PaymentsOffSessionPaymentRetryDetailsRetryStrategyScheduled V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy = "scheduled"
	V2PaymentsOffSessionPaymentRetryDetailsRetryStrategySmart     V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy = "smart"
)

// Status of this OffSessionPayment, one of `pending`, `pending_retry`, `processing`,
// `failed`, `canceled`, `requires_capture`, or `succeeded`.
type V2PaymentsOffSessionPaymentStatus string

// List of values that V2PaymentsOffSessionPaymentStatus can take
const (
	V2PaymentsOffSessionPaymentStatusCanceled        V2PaymentsOffSessionPaymentStatus = "canceled"
	V2PaymentsOffSessionPaymentStatusFailed          V2PaymentsOffSessionPaymentStatus = "failed"
	V2PaymentsOffSessionPaymentStatusPending         V2PaymentsOffSessionPaymentStatus = "pending"
	V2PaymentsOffSessionPaymentStatusPendingRetry    V2PaymentsOffSessionPaymentStatus = "pending_retry"
	V2PaymentsOffSessionPaymentStatusProcessing      V2PaymentsOffSessionPaymentStatus = "processing"
	V2PaymentsOffSessionPaymentStatusRequiresCapture V2PaymentsOffSessionPaymentStatus = "requires_capture"
	V2PaymentsOffSessionPaymentStatusSucceeded       V2PaymentsOffSessionPaymentStatus = "succeeded"
)

// Contains information about the tax on the item.
type V2PaymentsOffSessionPaymentAmountDetailsLineItemTax struct {
	// Total portion of the amount that is for tax.
	TotalTaxAmount int64 `json:"total_tax_amount,omitempty"`
}

// A list of line items, each containing information about a product in the PaymentIntent. There is a maximum of 100 line items.
type V2PaymentsOffSessionPaymentAmountDetailsLineItem struct {
	// The amount an item was discounted for. Positive integer.
	DiscountAmount int64 `json:"discount_amount,omitempty"`
	// Unique identifier of the product. At most 12 characters long.
	ProductCode string `json:"product_code,omitempty"`
	// Name of the product. At most 100 characters long.
	ProductName string `json:"product_name"`
	// Number of items of the product. Positive integer.
	Quantity int64 `json:"quantity"`
	// Contains information about the tax on the item.
	Tax *V2PaymentsOffSessionPaymentAmountDetailsLineItemTax `json:"tax,omitempty"`
	// Cost of the product. Non-negative integer.
	UnitCost int64 `json:"unit_cost"`
}

// Contains information about the shipping portion of the amount.
type V2PaymentsOffSessionPaymentAmountDetailsShipping struct {
	// Portion of the amount that is for shipping.
	Amount int64 `json:"amount,omitempty"`
	// The postal code that represents the shipping source.
	FromPostalCode string `json:"from_postal_code,omitempty"`
	// The postal code that represents the shipping destination.
	ToPostalCode string `json:"to_postal_code,omitempty"`
}

// Contains information about the tax portion of the amount.
type V2PaymentsOffSessionPaymentAmountDetailsTax struct {
	// Total portion of the amount that is for tax.
	TotalTaxAmount int64 `json:"total_tax_amount,omitempty"`
}

// Provides industry-specific information about the amount.
type V2PaymentsOffSessionPaymentAmountDetails struct {
	// The amount the total transaction was discounted for.
	DiscountAmount int64 `json:"discount_amount,omitempty"`
	// A list of line items, each containing information about a product in the PaymentIntent. There is a maximum of 100 line items.
	LineItems []*V2PaymentsOffSessionPaymentAmountDetailsLineItem `json:"line_items"`
	// Contains information about the shipping portion of the amount.
	Shipping *V2PaymentsOffSessionPaymentAmountDetailsShipping `json:"shipping,omitempty"`
	// Contains information about the tax portion of the amount.
	Tax *V2PaymentsOffSessionPaymentAmountDetailsTax `json:"tax,omitempty"`
}

// Details about the payments orchestration configuration.
type V2PaymentsOffSessionPaymentPaymentsOrchestration struct {
	// True when you want to enable payments orchestration for this off-session payment. False otherwise.
	Enabled bool `json:"enabled"`
}

// Details about the OffSessionPayment retries.
type V2PaymentsOffSessionPaymentRetryDetails struct {
	// Number of authorization attempts so far.
	Attempts int64 `json:"attempts"`
	// The pre-configured retry policy to use for the payment.
	RetryPolicy string `json:"retry_policy,omitempty"`
	// Indicates the strategy for how you want Stripe to retry the payment.
	RetryStrategy V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy `json:"retry_strategy"`
}

// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.corp.stripe.com/payments/connected-accounts).
type V2PaymentsOffSessionPaymentTransferData struct {
	// The amount transferred to the destination account. This transfer will occur
	// automatically after the payment succeeds. If no amount is specified, by default
	// the entire payment amount is transferred to the destination account. The amount
	// must be less than or equal to the
	// [amount_requested](https://docs.corp.stripe.com/api/v2/off-session-payments/object?api-version=2025-05-28.preview#v2_off_session_payment_object-amount_requested),
	// and must be a positive integer representing how much to transfer in the smallest
	// currency unit (e.g., 100 cents to charge $1.00).
	Amount int64 `json:"amount,omitempty"`
	// The account (if any) that the payment is attributed to for tax reporting, and
	// where funds from the payment are transferred to after payment success.
	Destination string `json:"destination"`
}

// OffSessionPayment resource.
type V2PaymentsOffSessionPayment struct {
	APIResource
	// Provides industry-specific information about the amount.
	AmountDetails *V2PaymentsOffSessionPaymentAmountDetails `json:"amount_details,omitempty"`
	// The “presentment amount” to be collected from the customer.
	AmountRequested Amount `json:"amount_requested"`
	// The frequency of the underlying payment.
	Cadence V2PaymentsOffSessionPaymentCadence `json:"cadence"`
	// ID of the owning compartment.
	CompartmentID string `json:"compartment_id"`
	// Creation time of the OffSessionPayment. Represented as a RFC 3339 date & time UTC
	// value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// ID of the Customer to which this OffSessionPayment belongs.
	Customer string `json:"customer"`
	// The reason why the OffSessionPayment failed.
	FailureReason V2PaymentsOffSessionPaymentFailureReason `json:"failure_reason,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The payment error encountered in the previous attempt to authorize the payment.
	LastAuthorizationAttemptError string `json:"last_authorization_attempt_error,omitempty"`
	// Payment attempt record for the latest attempt, if one exists.
	LatestPaymentAttemptRecord string `json:"latest_payment_attempt_record,omitempty"`
	// Has the value true if the object exists in live mode or the value false if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.corp.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.corp.stripe.com/payments/payment-intents#storing-information-in-metadata).
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The account (if any) for which the funds of the OffSessionPayment are intended.
	OnBehalfOf string `json:"on_behalf_of,omitempty"`
	// ID of the payment method used in this OffSessionPayment.
	PaymentMethod string `json:"payment_method"`
	// Payment record associated with the OffSessionPayment.
	PaymentRecord string `json:"payment_record,omitempty"`
	// Details about the payments orchestration configuration.
	PaymentsOrchestration *V2PaymentsOffSessionPaymentPaymentsOrchestration `json:"payments_orchestration"`
	// Details about the OffSessionPayment retries.
	RetryDetails *V2PaymentsOffSessionPaymentRetryDetails `json:"retry_details"`
	// Text that appears on the customer's statement as the statement descriptor for a
	// non-card charge. This value overrides the account's default statement descriptor.
	// For information about requirements, including the 22-character limit, see the
	// [Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// Provides information about a card charge. Concatenated to the account's
	// [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static)
	// to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix string `json:"statement_descriptor_suffix,omitempty"`
	// Status of this OffSessionPayment, one of `pending`, `pending_retry`, `processing`,
	// `failed`, `canceled`, `requires_capture`, or `succeeded`.
	Status V2PaymentsOffSessionPaymentStatus `json:"status"`
	// Test clock that can be used to advance the retry attempts in a sandbox.
	TestClock string `json:"test_clock,omitempty"`
	// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.corp.stripe.com/payments/connected-accounts).
	TransferData *V2PaymentsOffSessionPaymentTransferData `json:"transfer_data,omitempty"`
}
