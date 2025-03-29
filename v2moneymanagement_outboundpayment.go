//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. Method for bank account.
type V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount string

// List of values that V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount can take
const (
	V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccountAutomatic V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount = "automatic"
	V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccountLocal     V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount = "local"
	V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccountWire      V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount = "wire"
)

// Closed Enum. Configuration option to enable or disable notifications to recipients.
// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
type V2MoneyManagementOutboundPaymentRecipientNotificationSetting string

// List of values that V2MoneyManagementOutboundPaymentRecipientNotificationSetting can take
const (
	V2MoneyManagementOutboundPaymentRecipientNotificationSettingConfigured V2MoneyManagementOutboundPaymentRecipientNotificationSetting = "configured"
	V2MoneyManagementOutboundPaymentRecipientNotificationSettingNone       V2MoneyManagementOutboundPaymentRecipientNotificationSetting = "none"
)

// Closed Enum. Current status of the OutboundPayment: `processing`, `failed`, `posted`, `returned`, `canceled`.
// An OutboundPayment is `processing` if it has been created and is processing.
// The status changes to `posted` once the OutboundPayment has been "confirmed" and funds have left the account, or to `failed` or `canceled`.
// If an OutboundPayment fails to arrive at its payout method, its status will change to `returned`.
type V2MoneyManagementOutboundPaymentStatus string

// List of values that V2MoneyManagementOutboundPaymentStatus can take
const (
	V2MoneyManagementOutboundPaymentStatusCanceled   V2MoneyManagementOutboundPaymentStatus = "canceled"
	V2MoneyManagementOutboundPaymentStatusFailed     V2MoneyManagementOutboundPaymentStatus = "failed"
	V2MoneyManagementOutboundPaymentStatusPosted     V2MoneyManagementOutboundPaymentStatus = "posted"
	V2MoneyManagementOutboundPaymentStatusProcessing V2MoneyManagementOutboundPaymentStatus = "processing"
	V2MoneyManagementOutboundPaymentStatusReturned   V2MoneyManagementOutboundPaymentStatus = "returned"
)

// Open Enum. The `failed` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsFailedReason string

// List of values that V2MoneyManagementOutboundPaymentStatusDetailsFailedReason can take
const (
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodDeclined                    V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_declined"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodDoesNotExist                V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_does_not_exist"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodExpired                     V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_expired"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodUnsupported                 V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_unsupported"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodUsageFrequencyLimitExceeded V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_usage_frequency_limit_exceeded"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonUnknownFailure                          V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "unknown_failure"
)

// Open Enum. The `returned` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason string

// List of values that V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason can take
const (
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodCanceledByCustomer     V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_canceled_by_customer"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodClosed                 V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_closed"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodCurrencyUnsupported    V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_currency_unsupported"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodDoesNotExist           V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_does_not_exist"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodHolderAddressIncorrect V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_holder_address_incorrect"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodHolderDetailsIncorrect V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_holder_details_incorrect"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodHolderNameIncorrect    V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_holder_name_incorrect"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodInvalidAccountNumber   V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_invalid_account_number"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodRestricted             V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_restricted"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonRecalled                           V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "recalled"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonUnknownFailure                     V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "unknown_failure"
)

// Possible values are `pending`, `supported`, and `unsupported`. Initially set to `pending`, it changes to
// `supported` when the recipient bank provides a trace ID, or `unsupported` if the recipient bank doesn't support it.
// Note that this status may not align with the OutboundPayment or OutboundTransfer status and can remain `pending`
// even after the payment or transfer is posted.
type V2MoneyManagementOutboundPaymentTraceIDStatus string

// List of values that V2MoneyManagementOutboundPaymentTraceIDStatus can take
const (
	V2MoneyManagementOutboundPaymentTraceIDStatusPending     V2MoneyManagementOutboundPaymentTraceIDStatus = "pending"
	V2MoneyManagementOutboundPaymentTraceIDStatusSupported   V2MoneyManagementOutboundPaymentTraceIDStatus = "supported"
	V2MoneyManagementOutboundPaymentTraceIDStatusUnsupported V2MoneyManagementOutboundPaymentTraceIDStatus = "unsupported"
)

// Delivery options to be used to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentDeliveryOptions struct {
	// Open Enum. Method for bank account.
	BankAccount V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount `json:"bank_account"`
}

// The FinancialAccount that funds were pulled from.
type V2MoneyManagementOutboundPaymentFrom struct {
	// The monetary amount debited from the sender, only set on responses.
	Debited Amount `json:"debited"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount string `json:"financial_account"`
}

// Details about the OutboundPayment notification settings for recipient.
type V2MoneyManagementOutboundPaymentRecipientNotification struct {
	// Closed Enum. Configuration option to enable or disable notifications to recipients.
	// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
	Setting V2MoneyManagementOutboundPaymentRecipientNotificationSetting `json:"setting"`
}

// The `failed` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsFailed struct {
	// Open Enum. The `failed` status reason.
	Reason V2MoneyManagementOutboundPaymentStatusDetailsFailedReason `json:"reason"`
}

// The `returned` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsReturned struct {
	// Open Enum. The `returned` status reason.
	Reason V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason `json:"reason"`
}

// Status details for an OutboundPayment in a `failed` or `returned` state.
type V2MoneyManagementOutboundPaymentStatusDetails struct {
	// The `failed` status reason.
	Failed *V2MoneyManagementOutboundPaymentStatusDetailsFailed `json:"failed"`
	// The `returned` status reason.
	Returned *V2MoneyManagementOutboundPaymentStatusDetailsReturned `json:"returned"`
}

// Hash containing timestamps of when the object transitioned to a particular status.
type V2MoneyManagementOutboundPaymentStatusTransitions struct {
	// Timestamp describing when an OutboundPayment changed status to `canceled`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	CanceledAt time.Time `json:"canceled_at"`
	// Timestamp describing when an OutboundPayment changed status to `failed`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	FailedAt time.Time `json:"failed_at"`
	// Timestamp describing when an OutboundPayment changed status to `posted`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	PostedAt time.Time `json:"posted_at"`
	// Timestamp describing when an OutboundPayment changed status to `returned`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ReturnedAt time.Time `json:"returned_at"`
}

// To which payout method the OutboundPayment was sent.
type V2MoneyManagementOutboundPaymentTo struct {
	// The monetary amount being credited to the destination.
	Credited Amount `json:"credited"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod string `json:"payout_method"`
	// To which account the OutboundPayment is sent.
	Recipient string `json:"recipient"`
}

// A unique identifier that can be used to track this OutboundPayment with recipient bank. Banks might call this a “reference number” or something similar.
type V2MoneyManagementOutboundPaymentTraceID struct {
	// Possible values are `pending`, `supported`, and `unsupported`. Initially set to `pending`, it changes to
	// `supported` when the recipient bank provides a trace ID, or `unsupported` if the recipient bank doesn't support it.
	// Note that this status may not align with the OutboundPayment or OutboundTransfer status and can remain `pending`
	// even after the payment or transfer is posted.
	Status V2MoneyManagementOutboundPaymentTraceIDStatus `json:"status"`
	// The trace ID value if `trace_id.status` is `supported`, otherwise empty.
	Value string `json:"value"`
}

// OutboundPayment represents a single money movement from one FinancialAccount you own to a payout method someone else owns.
type V2MoneyManagementOutboundPayment struct {
	APIResource
	// The "presentment amount" for the OutboundPayment.
	Amount Amount `json:"amount"`
	// Returns true if the OutboundPayment can be canceled, and false otherwise.
	Cancelable bool `json:"cancelable"`
	// Time at which the OutboundPayment was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Delivery options to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentDeliveryOptions `json:"delivery_options"`
	// An arbitrary string attached to the OutboundPayment. Often useful for displaying to users.
	Description string `json:"description"`
	// The date when funds are expected to arrive in the payout method. This field is not set if the payout method is in a `failed`, `canceled`, or `returned` state.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ExpectedArrivalDate time.Time `json:"expected_arrival_date"`
	// The FinancialAccount that funds were pulled from.
	From *V2MoneyManagementOutboundPaymentFrom `json:"from"`
	// Unique identifier for the OutboundPayment.
	ID string `json:"id"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The quote for this OutboundPayment. Only required for countries with regulatory mandates to display fee estimates before OutboundPayment creation.
	OutboundPaymentQuote string `json:"outbound_payment_quote"`
	// A hosted transaction receipt URL that is provided when money movement is considered regulated under Stripe's money transmission licenses.
	ReceiptURL string `json:"receipt_url"`
	// Details about the OutboundPayment notification settings for recipient.
	RecipientNotification *V2MoneyManagementOutboundPaymentRecipientNotification `json:"recipient_notification"`
	// The description that appears on the receiving end for an OutboundPayment (for example, bank statement for external bank transfer).
	StatementDescriptor string `json:"statement_descriptor"`
	// Closed Enum. Current status of the OutboundPayment: `processing`, `failed`, `posted`, `returned`, `canceled`.
	// An OutboundPayment is `processing` if it has been created and is processing.
	// The status changes to `posted` once the OutboundPayment has been "confirmed" and funds have left the account, or to `failed` or `canceled`.
	// If an OutboundPayment fails to arrive at its payout method, its status will change to `returned`.
	Status V2MoneyManagementOutboundPaymentStatus `json:"status"`
	// Status details for an OutboundPayment in a `failed` or `returned` state.
	StatusDetails *V2MoneyManagementOutboundPaymentStatusDetails `json:"status_details"`
	// Hash containing timestamps of when the object transitioned to a particular status.
	StatusTransitions *V2MoneyManagementOutboundPaymentStatusTransitions `json:"status_transitions"`
	// To which payout method the OutboundPayment was sent.
	To *V2MoneyManagementOutboundPaymentTo `json:"to"`
	// A unique identifier that can be used to track this OutboundPayment with recipient bank. Banks might call this a “reference number” or something similar.
	TraceID *V2MoneyManagementOutboundPaymentTraceID `json:"trace_id"`
}
