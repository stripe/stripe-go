//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. Method for bank account.
type V2MoneyManagementOutboundTransferDeliveryOptionsBankAccount string

// List of values that V2MoneyManagementOutboundTransferDeliveryOptionsBankAccount can take
const (
	V2MoneyManagementOutboundTransferDeliveryOptionsBankAccountAutomatic V2MoneyManagementOutboundTransferDeliveryOptionsBankAccount = "automatic"
	V2MoneyManagementOutboundTransferDeliveryOptionsBankAccountLocal     V2MoneyManagementOutboundTransferDeliveryOptionsBankAccount = "local"
	V2MoneyManagementOutboundTransferDeliveryOptionsBankAccountWire      V2MoneyManagementOutboundTransferDeliveryOptionsBankAccount = "wire"
)

// Closed Enum. Current status of the OutboundTransfer: `processing`, `failed`, `posted`, `returned`, `canceled`.
// An OutboundTransfer is `processing` if it has been created and is processing.
// The status changes to `posted` once the OutboundTransfer has been "confirmed" and funds have left the account, or to `failed` or `canceled`.
// If an OutboundTransfer fails to arrive at its payout method, its status will change to `returned`.
type V2MoneyManagementOutboundTransferStatus string

// List of values that V2MoneyManagementOutboundTransferStatus can take
const (
	V2MoneyManagementOutboundTransferStatusCanceled   V2MoneyManagementOutboundTransferStatus = "canceled"
	V2MoneyManagementOutboundTransferStatusFailed     V2MoneyManagementOutboundTransferStatus = "failed"
	V2MoneyManagementOutboundTransferStatusPosted     V2MoneyManagementOutboundTransferStatus = "posted"
	V2MoneyManagementOutboundTransferStatusProcessing V2MoneyManagementOutboundTransferStatus = "processing"
	V2MoneyManagementOutboundTransferStatusReturned   V2MoneyManagementOutboundTransferStatus = "returned"
)

// Open Enum. The `failed` status reason.
type V2MoneyManagementOutboundTransferStatusDetailsFailedReason string

// List of values that V2MoneyManagementOutboundTransferStatusDetailsFailedReason can take
const (
	V2MoneyManagementOutboundTransferStatusDetailsFailedReasonPayoutMethodDeclined                    V2MoneyManagementOutboundTransferStatusDetailsFailedReason = "payout_method_declined"
	V2MoneyManagementOutboundTransferStatusDetailsFailedReasonPayoutMethodDoesNotExist                V2MoneyManagementOutboundTransferStatusDetailsFailedReason = "payout_method_does_not_exist"
	V2MoneyManagementOutboundTransferStatusDetailsFailedReasonPayoutMethodExpired                     V2MoneyManagementOutboundTransferStatusDetailsFailedReason = "payout_method_expired"
	V2MoneyManagementOutboundTransferStatusDetailsFailedReasonPayoutMethodUnsupported                 V2MoneyManagementOutboundTransferStatusDetailsFailedReason = "payout_method_unsupported"
	V2MoneyManagementOutboundTransferStatusDetailsFailedReasonPayoutMethodUsageFrequencyLimitExceeded V2MoneyManagementOutboundTransferStatusDetailsFailedReason = "payout_method_usage_frequency_limit_exceeded"
	V2MoneyManagementOutboundTransferStatusDetailsFailedReasonUnknownFailure                          V2MoneyManagementOutboundTransferStatusDetailsFailedReason = "unknown_failure"
)

// Open Enum. The `returned` status reason.
type V2MoneyManagementOutboundTransferStatusDetailsReturnedReason string

// List of values that V2MoneyManagementOutboundTransferStatusDetailsReturnedReason can take
const (
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonPayoutMethodCanceledByCustomer     V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "payout_method_canceled_by_customer"
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonPayoutMethodClosed                 V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "payout_method_closed"
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonPayoutMethodCurrencyUnsupported    V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "payout_method_currency_unsupported"
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonPayoutMethodDoesNotExist           V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "payout_method_does_not_exist"
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonPayoutMethodHolderAddressIncorrect V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "payout_method_holder_address_incorrect"
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonPayoutMethodHolderDetailsIncorrect V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "payout_method_holder_details_incorrect"
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonPayoutMethodHolderNameIncorrect    V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "payout_method_holder_name_incorrect"
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonPayoutMethodInvalidAccountNumber   V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "payout_method_invalid_account_number"
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonPayoutMethodRestricted             V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "payout_method_restricted"
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonRecalled                           V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "recalled"
	V2MoneyManagementOutboundTransferStatusDetailsReturnedReasonUnknownFailure                     V2MoneyManagementOutboundTransferStatusDetailsReturnedReason = "unknown_failure"
)

// Possible values are `pending`, `supported`, and `unsupported`. Initially set to `pending`, it changes to
// `supported` when the recipient bank provides a trace ID, or `unsupported` if the recipient bank doesn't support it.
// Note that this status may not align with the OutboundPayment or OutboundTransfer status and can remain `pending`
// even after the payment or transfer is posted.
type V2MoneyManagementOutboundTransferTraceIDStatus string

// List of values that V2MoneyManagementOutboundTransferTraceIDStatus can take
const (
	V2MoneyManagementOutboundTransferTraceIDStatusPending     V2MoneyManagementOutboundTransferTraceIDStatus = "pending"
	V2MoneyManagementOutboundTransferTraceIDStatusSupported   V2MoneyManagementOutboundTransferTraceIDStatus = "supported"
	V2MoneyManagementOutboundTransferTraceIDStatusUnsupported V2MoneyManagementOutboundTransferTraceIDStatus = "unsupported"
)

// Delivery options to be used to send the OutboundTransfer.
type V2MoneyManagementOutboundTransferDeliveryOptions struct {
	// Open Enum. Method for bank account.
	BankAccount V2MoneyManagementOutboundTransferDeliveryOptionsBankAccount `json:"bank_account"`
}

// The FinancialAccount that funds were pulled from.
type V2MoneyManagementOutboundTransferFrom struct {
	// The monetary amount debited from the sender, only set on responses.
	Debited Amount `json:"debited"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount string `json:"financial_account"`
}

// The `failed` status reason.
type V2MoneyManagementOutboundTransferStatusDetailsFailed struct {
	// Open Enum. The `failed` status reason.
	Reason V2MoneyManagementOutboundTransferStatusDetailsFailedReason `json:"reason"`
}

// The `returned` status reason.
type V2MoneyManagementOutboundTransferStatusDetailsReturned struct {
	// Open Enum. The `returned` status reason.
	Reason V2MoneyManagementOutboundTransferStatusDetailsReturnedReason `json:"reason"`
}

// Status details for an OutboundTransfer in a `failed` or `returned` state.
type V2MoneyManagementOutboundTransferStatusDetails struct {
	// The `failed` status reason.
	Failed *V2MoneyManagementOutboundTransferStatusDetailsFailed `json:"failed"`
	// The `returned` status reason.
	Returned *V2MoneyManagementOutboundTransferStatusDetailsReturned `json:"returned"`
}

// Hash containing timestamps of when the object transitioned to a particular status.
type V2MoneyManagementOutboundTransferStatusTransitions struct {
	// Timestamp describing when an OutboundTransfer changed status to `canceled`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	CanceledAt time.Time `json:"canceled_at"`
	// Timestamp describing when an OutboundTransfer changed status to `failed`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	FailedAt time.Time `json:"failed_at"`
	// Timestamp describing when an OutboundTransfer changed status to `posted`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	PostedAt time.Time `json:"posted_at"`
	// Timestamp describing when an OutboundTransfer changed status to `returned`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ReturnedAt time.Time `json:"returned_at"`
}

// To which payout method the OutboundTransfer was sent.
type V2MoneyManagementOutboundTransferTo struct {
	// The monetary amount being credited to the destination.
	Credited Amount `json:"credited"`
	// The payout method which the OutboundTransfer uses to send payout.
	PayoutMethod string `json:"payout_method"`
}

// A unique identifier that can be used to track this OutboundTransfer with recipient bank. Banks might call this a “reference number” or something similar.
type V2MoneyManagementOutboundTransferTraceID struct {
	// Possible values are `pending`, `supported`, and `unsupported`. Initially set to `pending`, it changes to
	// `supported` when the recipient bank provides a trace ID, or `unsupported` if the recipient bank doesn't support it.
	// Note that this status may not align with the OutboundPayment or OutboundTransfer status and can remain `pending`
	// even after the payment or transfer is posted.
	Status V2MoneyManagementOutboundTransferTraceIDStatus `json:"status"`
	// The trace ID value if `trace_id.status` is `supported`, otherwise empty.
	Value string `json:"value"`
}

// OutboundTransfer represents a single money movement from one FinancialAccount you own to a payout method you also own.
type V2MoneyManagementOutboundTransfer struct {
	APIResource
	// The "presentment amount" for the OutboundTransfer.
	Amount Amount `json:"amount"`
	// Returns true if the OutboundTransfer can be canceled, and false otherwise.
	Cancelable bool `json:"cancelable"`
	// Time at which the OutboundTransfer was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Delivery options to be used to send the OutboundTransfer.
	DeliveryOptions *V2MoneyManagementOutboundTransferDeliveryOptions `json:"delivery_options"`
	// An arbitrary string attached to the OutboundTransfer. Often useful for displaying to users.
	Description string `json:"description"`
	// The date when funds are expected to arrive in the payout method. This field is not set if the payout method is in a `failed`, `canceled`, or `returned` state.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ExpectedArrivalDate time.Time `json:"expected_arrival_date"`
	// The FinancialAccount that funds were pulled from.
	From *V2MoneyManagementOutboundTransferFrom `json:"from"`
	// Unique identifier for the OutboundTransfer.
	ID string `json:"id"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// A hosted transaction receipt URL that is provided when money movement is considered regulated under Stripe's money transmission licenses.
	ReceiptURL string `json:"receipt_url"`
	// The description that appears on the receiving end for an OutboundTransfer (for example, bank statement for external bank transfer).
	StatementDescriptor string `json:"statement_descriptor"`
	// Closed Enum. Current status of the OutboundTransfer: `processing`, `failed`, `posted`, `returned`, `canceled`.
	// An OutboundTransfer is `processing` if it has been created and is processing.
	// The status changes to `posted` once the OutboundTransfer has been "confirmed" and funds have left the account, or to `failed` or `canceled`.
	// If an OutboundTransfer fails to arrive at its payout method, its status will change to `returned`.
	Status V2MoneyManagementOutboundTransferStatus `json:"status"`
	// Status details for an OutboundTransfer in a `failed` or `returned` state.
	StatusDetails *V2MoneyManagementOutboundTransferStatusDetails `json:"status_details"`
	// Hash containing timestamps of when the object transitioned to a particular status.
	StatusTransitions *V2MoneyManagementOutboundTransferStatusTransitions `json:"status_transitions"`
	// To which payout method the OutboundTransfer was sent.
	To *V2MoneyManagementOutboundTransferTo `json:"to"`
	// A unique identifier that can be used to track this OutboundTransfer with recipient bank. Banks might call this a “reference number” or something similar.
	TraceID *V2MoneyManagementOutboundTransferTraceID `json:"trace_id"`
}
