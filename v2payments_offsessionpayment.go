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

// The method to use to capture the payment.
type V2PaymentsOffSessionPaymentCaptureCaptureMethod string

// List of values that V2PaymentsOffSessionPaymentCaptureCaptureMethod can take
const (
	V2PaymentsOffSessionPaymentCaptureCaptureMethodAutomatic V2PaymentsOffSessionPaymentCaptureCaptureMethod = "automatic"
	V2PaymentsOffSessionPaymentCaptureCaptureMethodManual    V2PaymentsOffSessionPaymentCaptureCaptureMethod = "manual"
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

// The amount available to be captured.
type V2PaymentsOffSessionPaymentAmountCapturable struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// The “presentment amount” to be collected from the customer.
type V2PaymentsOffSessionPaymentAmountRequested struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// Details about the capture configuration for the OffSessionPayment.
type V2PaymentsOffSessionPaymentCapture struct {
	// The timestamp when this payment is no longer eligible to be captured.
	CaptureBefore time.Time `json:"capture_before,omitempty"`
	// The method to use to capture the payment.
	CaptureMethod V2PaymentsOffSessionPaymentCaptureCaptureMethod `json:"capture_method"`
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

// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.stripe.com/payments/connected-accounts).
type V2PaymentsOffSessionPaymentTransferData struct {
	// The amount transferred to the destination account. This transfer will occur
	// automatically after the payment succeeds. If no amount is specified, by default
	// the entire payment amount is transferred to the destination account. The amount
	// must be less than or equal to the
	// [amount_requested](https://docs.stripe.com/api/v2/off-session-payments/object?api-version=2025-05-28.preview#v2_off_session_payment_object-amount_requested),
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
	// The amount available to be captured.
	AmountCapturable *V2PaymentsOffSessionPaymentAmountCapturable `json:"amount_capturable,omitempty"`
	// The “presentment amount” to be collected from the customer.
	AmountRequested *V2PaymentsOffSessionPaymentAmountRequested `json:"amount_requested"`
	// The frequency of the underlying payment.
	Cadence V2PaymentsOffSessionPaymentCadence `json:"cadence"`
	// Details about the capture configuration for the OffSessionPayment.
	Capture *V2PaymentsOffSessionPaymentCapture `json:"capture,omitempty"`
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
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can
	// attach to an object. This can be useful for storing additional information about
	// the object in a structured format. Learn more about
	// [storing information in metadata](https://docs.stripe.com/payments/payment-intents#storing-information-in-metadata).
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
	// The data that automatically creates a Transfer after the payment finalizes. Learn more about the use case for [connected accounts](https://docs.stripe.com/payments/connected-accounts).
	TransferData *V2PaymentsOffSessionPaymentTransferData `json:"transfer_data,omitempty"`
}
