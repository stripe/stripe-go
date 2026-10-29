//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Closed Enum. Configuration option to enable or disable notifications to recipients.
// Do not send notifications when `setting` is `none`. Use the account setting when `setting` is `configured` or omitted.
type V2MoneyManagementOutboundPaymentRecipientNotificationSetting string

// List of values that V2MoneyManagementOutboundPaymentRecipientNotificationSetting can take
const (
	V2MoneyManagementOutboundPaymentRecipientNotificationSettingConfigured V2MoneyManagementOutboundPaymentRecipientNotificationSetting = "configured"
	V2MoneyManagementOutboundPaymentRecipientNotificationSettingNone       V2MoneyManagementOutboundPaymentRecipientNotificationSetting = "none"
)

// Current status of the OutboundPayment.
// An OutboundPayment is `processing` after creation. It becomes `posted` when funds leave the account,
// `failed` if it cannot be sent, `canceled` if canceled before posting, or `returned` if the funds are returned.
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
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonFxRateDriftExceededAfterReview          V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "fx_rate_drift_exceeded_after_review"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodAmountLimitExceeded         V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_amount_limit_exceeded"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodDeclined                    V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_declined"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodDoesNotExist                V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_does_not_exist"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodExpired                     V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_expired"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodUnsupported                 V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_unsupported"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodUsageFrequencyLimitExceeded V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_usage_frequency_limit_exceeded"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonReviewRejected                          V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "review_rejected"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonUnknownFailure                          V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "unknown_failure"
)

// Open Enum. The `processing` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsProcessingReason string

// List of values that V2MoneyManagementOutboundPaymentStatusDetailsProcessingReason can take
const (
	V2MoneyManagementOutboundPaymentStatusDetailsProcessingReasonUnderReview V2MoneyManagementOutboundPaymentStatusDetailsProcessingReason = "under_review"
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

// Open Enum. ACH submission timing.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission string

// List of values that V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission can take
const (
	V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmissionNextDay V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission = "next_day"
	V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmissionSameDay V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission = "same_day"
)

// The transaction purpose for this ACH payment.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose string

// List of values that V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose can take
const (
	V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurposePayroll V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose = "payroll"
)

// The preferred networks to use for this OutboundPayment.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetwork string

// List of values that V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetwork can take
const (
	V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkACH     V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetwork = "ach"
	V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkFedwire V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetwork = "fedwire"
	V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkFPS     V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetwork = "fps"
	V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkRTP     V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetwork = "rtp"
	V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkSEPA    V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetwork = "sepa"
	V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkSwift   V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetwork = "swift"
)

// The status of the tracking details for this OutboundPayment.
type V2MoneyManagementOutboundPaymentTrackingDetailsStatus string

// List of values that V2MoneyManagementOutboundPaymentTrackingDetailsStatus can take
const (
	V2MoneyManagementOutboundPaymentTrackingDetailsStatusPending     V2MoneyManagementOutboundPaymentTrackingDetailsStatus = "pending"
	V2MoneyManagementOutboundPaymentTrackingDetailsStatusSupported   V2MoneyManagementOutboundPaymentTrackingDetailsStatus = "supported"
	V2MoneyManagementOutboundPaymentTrackingDetailsStatusUnsupported V2MoneyManagementOutboundPaymentTrackingDetailsStatus = "unsupported"
)

// The type of tracking details for this OutboundPayment.
type V2MoneyManagementOutboundPaymentTrackingDetailsType string

// List of values that V2MoneyManagementOutboundPaymentTrackingDetailsType can take
const (
	V2MoneyManagementOutboundPaymentTrackingDetailsTypeBankAccount V2MoneyManagementOutboundPaymentTrackingDetailsType = "bank_account"
	V2MoneyManagementOutboundPaymentTrackingDetailsTypeCard        V2MoneyManagementOutboundPaymentTrackingDetailsType = "card"
)

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
	// Do not send notifications when `setting` is `none`. Use the account setting when `setting` is `configured` or omitted.
	Setting V2MoneyManagementOutboundPaymentRecipientNotificationSetting `json:"setting"`
}

// The `failed` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsFailed struct {
	// Open Enum. The `failed` status reason.
	Reason V2MoneyManagementOutboundPaymentStatusDetailsFailedReason `json:"reason"`
}

// The `processing` status details.
type V2MoneyManagementOutboundPaymentStatusDetailsProcessing struct {
	// Open Enum. The `processing` status reason.
	Reason V2MoneyManagementOutboundPaymentStatusDetailsProcessingReason `json:"reason"`
}

// The `returned` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsReturned struct {
	// Open Enum. The `returned` status reason.
	Reason V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason `json:"reason"`
}

// Status details for an OutboundPayment in a `processing`, `failed`, or `returned` state.
type V2MoneyManagementOutboundPaymentStatusDetails struct {
	// The `failed` status reason.
	Failed *V2MoneyManagementOutboundPaymentStatusDetailsFailed `json:"failed,omitempty"`
	// The `processing` status details.
	Processing *V2MoneyManagementOutboundPaymentStatusDetailsProcessing `json:"processing,omitempty"`
	// The `returned` status reason.
	Returned *V2MoneyManagementOutboundPaymentStatusDetailsReturned `json:"returned,omitempty"`
}

// Hash containing timestamps of when the object transitioned to a particular status.
type V2MoneyManagementOutboundPaymentStatusTransitions struct {
	// Timestamp describing when an OutboundPayment changed status to `canceled`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Timestamp describing when an OutboundPayment changed status to `failed`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	FailedAt time.Time `json:"failed_at,omitempty"`
	// Timestamp describing when an OutboundPayment changed status to `posted`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	PostedAt time.Time `json:"posted_at,omitempty"`
	// Timestamp describing when an OutboundPayment changed status to `returned`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ReturnedAt time.Time `json:"returned_at,omitempty"`
}

// ACH-specific network options.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACH struct {
	// Open Enum. ACH submission timing.
	Submission V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission `json:"submission,omitempty"`
	// The transaction purpose for this ACH payment.
	TransactionPurpose V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose `json:"transaction_purpose,omitempty"`
}

// Per-network configuration options.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptions struct {
	// ACH-specific network options.
	ACH *V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACH `json:"ach,omitempty"`
}

// Options for bank account payout methods.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccount struct {
	// Per-network configuration options.
	PreferredNetworkOptions *V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptions `json:"preferred_network_options,omitempty"`
	// The preferred networks to use for this OutboundPayment.
	PreferredNetworks []V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetwork `json:"preferred_networks,omitempty"`
}

// Payout method options for the OutboundPayment.
type V2MoneyManagementOutboundPaymentToPayoutMethodOptions struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccount `json:"bank_account,omitempty"`
}

// To which payout method the OutboundPayment was sent.
type V2MoneyManagementOutboundPaymentTo struct {
	// The monetary amount being credited to the destination.
	Credited Amount `json:"credited"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod string `json:"payout_method,omitempty"`
	// Payout method options for the OutboundPayment.
	PayoutMethodOptions *V2MoneyManagementOutboundPaymentToPayoutMethodOptions `json:"payout_method_options,omitempty"`
	// To which account the OutboundPayment is sent.
	Recipient string `json:"recipient"`
}

// Bank account tracking details.
type V2MoneyManagementOutboundPaymentTrackingDetailsBankAccount struct {
	// A unique identifier that can be used to track this OutboundPayment with the recipient bank. Banks might call this a “reference number” or something similar.
	TraceID string `json:"trace_id,omitempty"`
}

// Card tracking details.
type V2MoneyManagementOutboundPaymentTrackingDetailsCard struct {
	// A unique identifier that can be used to track this OutboundPayment with the recipient card.
	RetrievalReference string `json:"retrieval_reference,omitempty"`
}

// Information to track this OutboundPayment with the recipient bank.
type V2MoneyManagementOutboundPaymentTrackingDetails struct {
	// Bank account tracking details.
	BankAccount *V2MoneyManagementOutboundPaymentTrackingDetailsBankAccount `json:"bank_account,omitempty"`
	// Card tracking details.
	Card *V2MoneyManagementOutboundPaymentTrackingDetailsCard `json:"card,omitempty"`
	// The status of the tracking details for this OutboundPayment.
	Status V2MoneyManagementOutboundPaymentTrackingDetailsStatus `json:"status"`
	// The type of tracking details for this OutboundPayment.
	Type V2MoneyManagementOutboundPaymentTrackingDetailsType `json:"type"`
}

// OutboundPayment represents a single money movement from one FinancialAccount you own to a payout method someone else owns.
type V2MoneyManagementOutboundPayment struct {
	APIResource
	// The amount and currency requested for this OutboundPayment.
	Amount Amount `json:"amount"`
	// Returns true if the OutboundPayment can be canceled, and false otherwise.
	Cancelable bool `json:"cancelable"`
	// Time at which the OutboundPayment was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// An arbitrary string attached to the OutboundPayment. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// The date when funds are expected to arrive in the payout method. This field is not set if the payout method is in a `failed`, `canceled`, or `returned` state.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ExpectedArrivalDate time.Time `json:"expected_arrival_date,omitempty"`
	// The FinancialAccount that funds were pulled from.
	From *V2MoneyManagementOutboundPaymentFrom `json:"from"`
	// Unique identifier for the OutboundPayment.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The PayoutIntent ID that triggered this OutboundPayment.
	PayoutIntent string `json:"payout_intent,omitempty"`
	// Details about the OutboundPayment notification settings for recipient.
	RecipientNotification *V2MoneyManagementOutboundPaymentRecipientNotification `json:"recipient_notification"`
	// The ID of the RecipientVerification for this OutboundPayment. Required where recipient-name verification is mandated before creating an OutboundPayment.
	RecipientVerification string `json:"recipient_verification,omitempty"`
	// The description that appears on the receiving end of an OutboundPayment (for example, on a bank statement). Defaults to `STRIPE` if not configured in the account settings.
	StatementDescriptor string `json:"statement_descriptor"`
	// Current status of the OutboundPayment.
	// An OutboundPayment is `processing` after creation. It becomes `posted` when funds leave the account,
	// `failed` if it cannot be sent, `canceled` if canceled before posting, or `returned` if the funds are returned.
	Status V2MoneyManagementOutboundPaymentStatus `json:"status"`
	// Status details for an OutboundPayment in a `processing`, `failed`, or `returned` state.
	StatusDetails *V2MoneyManagementOutboundPaymentStatusDetails `json:"status_details,omitempty"`
	// Hash containing timestamps of when the object transitioned to a particular status.
	StatusTransitions *V2MoneyManagementOutboundPaymentStatusTransitions `json:"status_transitions,omitempty"`
	// To which payout method the OutboundPayment was sent.
	To *V2MoneyManagementOutboundPaymentTo `json:"to"`
	// Information to track this OutboundPayment with the recipient bank.
	TrackingDetails *V2MoneyManagementOutboundPaymentTrackingDetails `json:"tracking_details,omitempty"`
}
