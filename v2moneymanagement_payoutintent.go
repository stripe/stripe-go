//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of payout.
type V2MoneyManagementPayoutIntentLatestPayoutType string

// List of values that V2MoneyManagementPayoutIntentLatestPayoutType can take
const (
	V2MoneyManagementPayoutIntentLatestPayoutTypeOutboundPayment  V2MoneyManagementPayoutIntentLatestPayoutType = "outbound_payment"
	V2MoneyManagementPayoutIntentLatestPayoutTypeOutboundTransfer V2MoneyManagementPayoutIntentLatestPayoutType = "outbound_transfer"
)

// Open Enum. The reason for the failure.
type V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason string

// List of values that V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason can take
const (
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonAccountNotConfiguredAsRecipient                V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "account_not_configured_as_recipient"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonCurrencyNotSupportedForFinancialAccountBalance V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "currency_not_supported_for_financial_account_balance"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonCurrencyRequired                               V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "currency_required"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonFeatureNotActiveForRecipient                   V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "feature_not_active_for_recipient"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonFxRateDriftExceededAfterReview                 V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "fx_rate_drift_exceeded_after_review"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonInsufficientFunds                              V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "insufficient_funds"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodAccountTypeIncorrect               V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_account_type_incorrect"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodAmountLimitExceeded                V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_amount_limit_exceeded"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodCanceledByCustomer                 V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_canceled_by_customer"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodClosed                             V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_closed"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodCurrencyUnsupported                V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_currency_unsupported"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodDeclined                           V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_declined"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodDoesNotExist                       V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_does_not_exist"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodExpired                            V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_expired"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodHolderAddressIncorrect             V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_holder_address_incorrect"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodHolderDetailsIncorrect             V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_holder_details_incorrect"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodHolderNameIncorrect                V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_holder_name_incorrect"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodInvalidAccountNumber               V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_invalid_account_number"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodRestricted                         V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_restricted"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodUnsupported                        V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_unsupported"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodUsageFrequencyLimitExceeded        V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_usage_frequency_limit_exceeded"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonRecalled                                       V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "recalled"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonReviewRejected                                 V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "review_rejected"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonToDestinationInvalid                           V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "to_destination_invalid"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonUnknownFailure                                 V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "unknown_failure"
)

// Open Enum. The type of next action required.
type V2MoneyManagementPayoutIntentNextActionType string

// List of values that V2MoneyManagementPayoutIntentNextActionType can take
const (
	V2MoneyManagementPayoutIntentNextActionTypeHandleFailure V2MoneyManagementPayoutIntentNextActionType = "handle_failure"
)

// Closed Enum. Configuration option to enable or disable notifications to recipients.
// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
type V2MoneyManagementPayoutIntentRecipientNotificationSetting string

// List of values that V2MoneyManagementPayoutIntentRecipientNotificationSetting can take
const (
	V2MoneyManagementPayoutIntentRecipientNotificationSettingConfigured V2MoneyManagementPayoutIntentRecipientNotificationSetting = "configured"
	V2MoneyManagementPayoutIntentRecipientNotificationSettingNone       V2MoneyManagementPayoutIntentRecipientNotificationSetting = "none"
)

// Open Enum. Current status of the PayoutIntent: `pending`, `processing`, `posted`, `canceled`, `requires_action`.
type V2MoneyManagementPayoutIntentStatus string

// List of values that V2MoneyManagementPayoutIntentStatus can take
const (
	V2MoneyManagementPayoutIntentStatusCanceled       V2MoneyManagementPayoutIntentStatus = "canceled"
	V2MoneyManagementPayoutIntentStatusPending        V2MoneyManagementPayoutIntentStatus = "pending"
	V2MoneyManagementPayoutIntentStatusPosted         V2MoneyManagementPayoutIntentStatus = "posted"
	V2MoneyManagementPayoutIntentStatusProcessing     V2MoneyManagementPayoutIntentStatus = "processing"
	V2MoneyManagementPayoutIntentStatusRequiresAction V2MoneyManagementPayoutIntentStatus = "requires_action"
)

// Open Enum. ACH submission timing.
type V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission string

// List of values that V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission can take
const (
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmissionNextDay V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission = "next_day"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmissionSameDay V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission = "same_day"
)

// The transaction purpose for this ACH payment.
type V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose string

// List of values that V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose can take
const (
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurposePayroll V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose = "payroll"
)

// The preferred networks to use for this PayoutIntent.
type V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork string

// List of values that V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork can take
const (
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkACH         V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "ach"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkBECS        V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "becs"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkEft         V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "eft"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkFedwire     V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "fedwire"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkFPS         V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "fps"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkNpp         V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "npp"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkRTP         V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "rtp"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkSEPACredit  V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "sepa_credit"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkSEPAInstant V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "sepa_instant"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkSwift       V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "swift"
)

// The FinancialAccount that funds are pulled from.
type V2MoneyManagementPayoutIntentFrom struct {
	// The currency of the financial account.
	Currency Currency `json:"currency"`
	// The FinancialAccount that funds are pulled from.
	FinancialAccount string `json:"financial_account"`
}

// Details about the latest payout associated with this PayoutIntent.
type V2MoneyManagementPayoutIntentLatestPayout struct {
	// The ID of the OutboundPayment, if applicable.
	OutboundPayment string `json:"outbound_payment,omitempty"`
	// The ID of the OutboundTransfer, if applicable.
	OutboundTransfer string `json:"outbound_transfer,omitempty"`
	// The type of payout.
	Type V2MoneyManagementPayoutIntentLatestPayoutType `json:"type"`
}

// Details about a failure that requires user action. Populated when type is handle_failure.
type V2MoneyManagementPayoutIntentNextActionHandleFailure struct {
	// Open Enum. The reason for the failure.
	FailureReason V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason `json:"failure_reason"`
}

// Next action required for a PayoutIntent in the requires_action state. Populated when status is requires_action.
type V2MoneyManagementPayoutIntentNextAction struct {
	// Details about a failure that requires user action. Populated when type is handle_failure.
	HandleFailure *V2MoneyManagementPayoutIntentNextActionHandleFailure `json:"handle_failure,omitempty"`
	// Open Enum. The type of next action required.
	Type V2MoneyManagementPayoutIntentNextActionType `json:"type"`
}

// Details about the OutboundPayment notification settings for recipient. Only applicable to OutboundPayment.
type V2MoneyManagementPayoutIntentRecipientNotification struct {
	// Closed Enum. Configuration option to enable or disable notifications to recipients.
	// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
	Setting V2MoneyManagementPayoutIntentRecipientNotificationSetting `json:"setting"`
}

// Scheduling options for the payout. If this is nil, we assume immediate execution.
type V2MoneyManagementPayoutIntentScheduleOptions struct {
	// The date when the payout should be executed, in YYYY-MM-DD format.
	ExecuteOn string `json:"execute_on,omitempty"`
}

// Hash containing timestamps of when transitioned to a particular status.
type V2MoneyManagementPayoutIntentStatusTransitions struct {
	// Timestamp describing when a PayoutIntent changed status to `canceled`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Timestamp describing when a PayoutIntent changed status to `posted`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	PostedAt time.Time `json:"posted_at,omitempty"`
	// Timestamp describing when a PayoutIntent changed status to `processing`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ProcessingAt time.Time `json:"processing_at,omitempty"`
	// Timestamp describing when a PayoutIntent changed status to `requires_action`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	RequiresActionAt time.Time `json:"requires_action_at,omitempty"`
}

// ACH-specific network options.
type V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACH struct {
	// Open Enum. ACH submission timing.
	Submission V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission `json:"submission,omitempty"`
	// The transaction purpose for this ACH payment.
	TransactionPurpose V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose `json:"transaction_purpose,omitempty"`
}

// Per-network configuration options.
type V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptions struct {
	// ACH-specific network options.
	ACH *V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACH `json:"ach,omitempty"`
}

// Options for bank account payout methods.
type V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccount struct {
	// Per-network configuration options.
	PreferredNetworkOptions *V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptions `json:"preferred_network_options,omitempty"`
	// The preferred networks to use for this PayoutIntent.
	PreferredNetworks []V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork `json:"preferred_networks"`
}

// Payout method options for the PayoutIntent.
type V2MoneyManagementPayoutIntentToPayoutMethodOptions struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccount `json:"bank_account,omitempty"`
}

// To which payout method the payout is sent.
type V2MoneyManagementPayoutIntentTo struct {
	// The currency to send to the recipient.
	Currency Currency `json:"currency,omitempty"`
	// The payout method ID. Optional for OutboundPayment if recipient has default payment method. Required for OutboundTransfer.
	PayoutMethod string `json:"payout_method,omitempty"`
	// Payout method options for the PayoutIntent.
	PayoutMethodOptions *V2MoneyManagementPayoutIntentToPayoutMethodOptions `json:"payout_method_options,omitempty"`
	// The recipient ID. Only relevant for OutboundPayment.
	Recipient string `json:"recipient,omitempty"`
}

// PayoutIntent represents an intent to send funds from a Financial Account to a payout method.
type V2MoneyManagementPayoutIntent struct {
	APIResource
	// The monetary amount to be sent.
	Amount Amount `json:"amount"`
	// Time at which the PayoutIntent was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// An arbitrary string attached to the PayoutIntent. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// The FinancialAccount that funds are pulled from.
	From *V2MoneyManagementPayoutIntentFrom `json:"from"`
	// Unique identifier for the PayoutIntent.
	ID string `json:"id"`
	// Details about the latest payout associated with this PayoutIntent.
	LatestPayout *V2MoneyManagementPayoutIntentLatestPayout `json:"latest_payout"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Next action required for a PayoutIntent in the requires_action state. Populated when status is requires_action.
	NextAction *V2MoneyManagementPayoutIntentNextAction `json:"next_action,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Details about the OutboundPayment notification settings for recipient. Only applicable to OutboundPayment.
	RecipientNotification *V2MoneyManagementPayoutIntentRecipientNotification `json:"recipient_notification,omitempty"`
	// Scheduling options for the payout. If this is nil, we assume immediate execution.
	ScheduleOptions *V2MoneyManagementPayoutIntentScheduleOptions `json:"schedule_options,omitempty"`
	// The description that appears on the receiving end for the payout (for example, on a bank statement).
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// Open Enum. Current status of the PayoutIntent: `pending`, `processing`, `posted`, `canceled`, `requires_action`.
	Status V2MoneyManagementPayoutIntentStatus `json:"status"`
	// Hash containing timestamps of when transitioned to a particular status.
	StatusTransitions *V2MoneyManagementPayoutIntentStatusTransitions `json:"status_transitions,omitempty"`
	// To which payout method the payout is sent.
	To *V2MoneyManagementPayoutIntentTo `json:"to"`
}
