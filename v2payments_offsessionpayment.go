//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The frequency of the underlying payment that this OSP represents.
type V2PaymentsOffSessionPaymentCadence string

// List of values that V2PaymentsOffSessionPaymentCadence can take
const (
	V2PaymentsOffSessionPaymentCadenceRecurring   V2PaymentsOffSessionPaymentCadence = "recurring"
	V2PaymentsOffSessionPaymentCadenceUnscheduled V2PaymentsOffSessionPaymentCadence = "unscheduled"
)

// Reason why the OSP failed.
type V2PaymentsOffSessionPaymentFailureReason string

// List of values that V2PaymentsOffSessionPaymentFailureReason can take
const (
	V2PaymentsOffSessionPaymentFailureReasonRejectedByPartner V2PaymentsOffSessionPaymentFailureReason = "rejected_by_partner"
	V2PaymentsOffSessionPaymentFailureReasonRetriesExhausted  V2PaymentsOffSessionPaymentFailureReason = "retries_exhausted"
)

// How you want Stripe to retry the payment.
type V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy string

// List of values that V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy can take
const (
	V2PaymentsOffSessionPaymentRetryDetailsRetryStrategyNone  V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy = "none"
	V2PaymentsOffSessionPaymentRetryDetailsRetryStrategySmart V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy = "smart"
)

// Status of the OSP.
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

// Details about the OSP retries.
type V2PaymentsOffSessionPaymentRetryDetails struct {
	// Number of authorization attempts so far.
	Attempts int64 `json:"attempts"`
	// How you want Stripe to retry the payment.
	RetryStrategy V2PaymentsOffSessionPaymentRetryDetailsRetryStrategy `json:"retry_strategy"`
}

// Instructions for the transfer to be made with this OSP after successful money movement.
type V2PaymentsOffSessionPaymentTransferData struct {
	// Amount in minor units that you want to transfer.
	Amount int64 `json:"amount"`
	// ID of the connected account where you want money to go.
	Destination string `json:"destination"`
}

// Off-session payment resource.
type V2PaymentsOffSessionPayment struct {
	APIResource
	// The amount you requested to be collected on the OSP upon creation.
	AmountRequested Amount `json:"amount_requested"`
	// Number of authorization attempts.
	Attempts int64 `json:"attempts"`
	// The frequency of the underlying payment that this OSP represents.
	Cadence V2PaymentsOffSessionPaymentCadence `json:"cadence"`
	// ID of owning compartment.
	CompartmentID string `json:"compartment_id"`
	// Timestamp of creation.
	Created time.Time `json:"created"`
	// Customer owning the supplied payment method.
	Customer string `json:"customer"`
	// Reason why the OSP failed.
	FailureReason V2PaymentsOffSessionPaymentFailureReason `json:"failure_reason"`
	// ID of the OSP.
	ID string `json:"id"`
	// Last error returned by the financial partner for a failed authorization.
	LastAuthorizationAttemptError string `json:"last_authorization_attempt_error"`
	// Payment attempt record for the latest attempt, if one exists.
	LatestPaymentAttemptRecord string `json:"latest_payment_attempt_record"`
	// True if the txn is livemode, false otherwise.
	Livemode bool `json:"livemode"`
	// Metadata you provided.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// OBO, same as on the PI.
	OnBehalfOf string `json:"on_behalf_of"`
	// ID of payment method.
	PaymentMethod string `json:"payment_method"`
	// Payment record associated with the OSP. consistent across attempts.
	PaymentRecord string `json:"payment_record"`
	// Details about the OSP retries.
	RetryDetails *V2PaymentsOffSessionPaymentRetryDetails `json:"retry_details"`
	// Statement descriptor you provided.
	StatementDescriptor string `json:"statement_descriptor"`
	// Statement descriptor suffix you provided, similar to that on the PI.
	StatementDescriptorSuffix string `json:"statement_descriptor_suffix"`
	// Status of the OSP.
	Status V2PaymentsOffSessionPaymentStatus `json:"status"`
	// Test clock to be used to advance the retry attempts.
	TestClock string `json:"test_clock"`
	// Instructions for the transfer to be made with this OSP after successful money movement.
	TransferData *V2PaymentsOffSessionPaymentTransferData `json:"transfer_data"`
}
