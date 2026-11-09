//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Controls whether the PayoutIntent requires explicit confirmation before transitioning to `pending`.
// Even when `automatic`, regulatory requirements may require confirmation. See [Send money with Global Payouts](https://docs.stripe.com/global-payouts/send-money) for details.
type V2MoneyManagementPayoutIntentConfirmationMethod string

// List of values that V2MoneyManagementPayoutIntentConfirmationMethod can take
const (
	V2MoneyManagementPayoutIntentConfirmationMethodAutomatic V2MoneyManagementPayoutIntentConfirmationMethod = "automatic"
	V2MoneyManagementPayoutIntentConfirmationMethodManual    V2MoneyManagementPayoutIntentConfirmationMethod = "manual"
)

// The network associated with the fee.
type V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetwork string

// List of values that V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetwork can take
const (
	V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkACH     V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetwork = "ach"
	V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkFedwire V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetwork = "fedwire"
	V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkFPS     V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetwork = "fps"
	V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkLocal   V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetwork = "local"
	V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkRTP     V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetwork = "rtp"
	V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkSEPA    V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetwork = "sepa"
	V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkSwift   V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetwork = "swift"
)

// Open Enum. ACH submission timing.
type V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptionsACHSubmission string

// List of values that V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptionsACHSubmission can take
const (
	V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptionsACHSubmissionNextDay V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptionsACHSubmission = "next_day"
	V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptionsACHSubmissionSameDay V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptionsACHSubmission = "same_day"
)

// Open Enum. The type of fee.
type V2MoneyManagementPayoutIntentEstimatedFeeType string

// List of values that V2MoneyManagementPayoutIntentEstimatedFeeType can take
const (
	V2MoneyManagementPayoutIntentEstimatedFeeTypeCrossBorderFee       V2MoneyManagementPayoutIntentEstimatedFeeType = "cross_border_fee"
	V2MoneyManagementPayoutIntentEstimatedFeeTypeForeignExchangeFee   V2MoneyManagementPayoutIntentEstimatedFeeType = "foreign_exchange_fee"
	V2MoneyManagementPayoutIntentEstimatedFeeTypeInstantCardPayoutFee V2MoneyManagementPayoutIntentEstimatedFeeType = "instant_card_payout_fee"
	V2MoneyManagementPayoutIntentEstimatedFeeTypeNetworkFee           V2MoneyManagementPayoutIntentEstimatedFeeType = "network_fee"
	V2MoneyManagementPayoutIntentEstimatedFeeTypeNextDayPayoutFee     V2MoneyManagementPayoutIntentEstimatedFeeType = "next_day_payout_fee"
	V2MoneyManagementPayoutIntentEstimatedFeeTypeRealTimePayoutFee    V2MoneyManagementPayoutIntentEstimatedFeeType = "real_time_payout_fee"
	V2MoneyManagementPayoutIntentEstimatedFeeTypeStandardPayoutFee    V2MoneyManagementPayoutIntentEstimatedFeeType = "standard_payout_fee"
	V2MoneyManagementPayoutIntentEstimatedFeeTypeWirePayoutFee        V2MoneyManagementPayoutIntentEstimatedFeeType = "wire_payout_fee"
)

// Open Enum. Duration of the FX rate lock.
type V2MoneyManagementPayoutIntentFxQuoteLockDuration string

// List of values that V2MoneyManagementPayoutIntentFxQuoteLockDuration can take
const (
	V2MoneyManagementPayoutIntentFxQuoteLockDurationFiveMinutes V2MoneyManagementPayoutIntentFxQuoteLockDuration = "five_minutes"
	V2MoneyManagementPayoutIntentFxQuoteLockDurationNone        V2MoneyManagementPayoutIntentFxQuoteLockDuration = "none"
)

// Open Enum. Lock status of the FX rate.
type V2MoneyManagementPayoutIntentFxQuoteLockStatus string

// List of values that V2MoneyManagementPayoutIntentFxQuoteLockStatus can take
const (
	V2MoneyManagementPayoutIntentFxQuoteLockStatusActive  V2MoneyManagementPayoutIntentFxQuoteLockStatus = "active"
	V2MoneyManagementPayoutIntentFxQuoteLockStatusExpired V2MoneyManagementPayoutIntentFxQuoteLockStatus = "expired"
	V2MoneyManagementPayoutIntentFxQuoteLockStatusNone    V2MoneyManagementPayoutIntentFxQuoteLockStatus = "none"
)

// The type of payout.
type V2MoneyManagementPayoutIntentLatestPayoutType string

// List of values that V2MoneyManagementPayoutIntentLatestPayoutType can take
const (
	V2MoneyManagementPayoutIntentLatestPayoutTypeOutboundPayment  V2MoneyManagementPayoutIntentLatestPayoutType = "outbound_payment"
	V2MoneyManagementPayoutIntentLatestPayoutTypeOutboundTransfer V2MoneyManagementPayoutIntentLatestPayoutType = "outbound_transfer"
)

// The PayoutIntent fields that must be reviewed before confirmation.
type V2MoneyManagementPayoutIntentNextActionConfirmFieldsToReview string

// List of values that V2MoneyManagementPayoutIntentNextActionConfirmFieldsToReview can take
const (
	V2MoneyManagementPayoutIntentNextActionConfirmFieldsToReviewEstimatedFees         V2MoneyManagementPayoutIntentNextActionConfirmFieldsToReview = "estimated_fees"
	V2MoneyManagementPayoutIntentNextActionConfirmFieldsToReviewFxQuote               V2MoneyManagementPayoutIntentNextActionConfirmFieldsToReview = "fx_quote"
	V2MoneyManagementPayoutIntentNextActionConfirmFieldsToReviewRecipientVerification V2MoneyManagementPayoutIntentNextActionConfirmFieldsToReview = "recipient_verification"
)

// Open Enum. The reason the PayoutIntent requires confirmation.
type V2MoneyManagementPayoutIntentNextActionConfirmReason string

// List of values that V2MoneyManagementPayoutIntentNextActionConfirmReason can take
const (
	V2MoneyManagementPayoutIntentNextActionConfirmReasonAutomaticallyRequired V2MoneyManagementPayoutIntentNextActionConfirmReason = "automatically_required"
	V2MoneyManagementPayoutIntentNextActionConfirmReasonManuallyRequested     V2MoneyManagementPayoutIntentNextActionConfirmReason = "manually_requested"
)

// Open Enum. The reason for the failure.
type V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason string

// List of values that V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason can take
const (
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonAccountNotConfiguredAsRecipient         V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "account_not_configured_as_recipient"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonFxRateDriftExceededAfterReview          V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "fx_rate_drift_exceeded_after_review"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonInsufficientFunds                       V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "insufficient_funds"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodAccountTypeIncorrect        V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_account_type_incorrect"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodAmountLimitExceeded         V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_amount_limit_exceeded"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodCanceledByCustomer          V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_canceled_by_customer"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodClosed                      V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_closed"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodCurrencyUnsupported         V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_currency_unsupported"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodDeclined                    V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_declined"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodDoesNotExist                V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_does_not_exist"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodExpired                     V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_expired"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodHolderAddressIncorrect      V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_holder_address_incorrect"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodHolderDetailsIncorrect      V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_holder_details_incorrect"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodHolderNameIncorrect         V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_holder_name_incorrect"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodInvalid                     V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_invalid"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodInvalidAccountNumber        V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_invalid_account_number"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodRestricted                  V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_restricted"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodUnsupported                 V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_unsupported"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonPayoutMethodUsageFrequencyLimitExceeded V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "payout_method_usage_frequency_limit_exceeded"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonRecalled                                V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "recalled"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonRecipientCapabilityNotActive            V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "recipient_capability_not_active"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonReviewRejected                          V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "review_rejected"
	V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReasonUnknownFailure                          V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason = "unknown_failure"
)

// Open Enum. The type of next action required.
type V2MoneyManagementPayoutIntentNextActionType string

// List of values that V2MoneyManagementPayoutIntentNextActionType can take
const (
	V2MoneyManagementPayoutIntentNextActionTypeConfirm       V2MoneyManagementPayoutIntentNextActionType = "confirm"
	V2MoneyManagementPayoutIntentNextActionTypeHandleFailure V2MoneyManagementPayoutIntentNextActionType = "handle_failure"
)

// Closed Enum. Configuration option to enable or disable notifications to recipients.
// Do not send notifications when `setting` is `none`. Use the account setting when `setting` is `configured` or omitted.
type V2MoneyManagementPayoutIntentRecipientNotificationSetting string

// List of values that V2MoneyManagementPayoutIntentRecipientNotificationSetting can take
const (
	V2MoneyManagementPayoutIntentRecipientNotificationSettingConfigured V2MoneyManagementPayoutIntentRecipientNotificationSetting = "configured"
	V2MoneyManagementPayoutIntentRecipientNotificationSettingNone       V2MoneyManagementPayoutIntentRecipientNotificationSetting = "none"
)

// Closed Enum. The recipient-verification provider's match result.
type V2MoneyManagementPayoutIntentRecipientVerificationMatchResult string

// List of values that V2MoneyManagementPayoutIntentRecipientVerificationMatchResult can take
const (
	V2MoneyManagementPayoutIntentRecipientVerificationMatchResultCloseMatch  V2MoneyManagementPayoutIntentRecipientVerificationMatchResult = "close_match"
	V2MoneyManagementPayoutIntentRecipientVerificationMatchResultMatch       V2MoneyManagementPayoutIntentRecipientVerificationMatchResult = "match"
	V2MoneyManagementPayoutIntentRecipientVerificationMatchResultNoMatch     V2MoneyManagementPayoutIntentRecipientVerificationMatchResult = "no_match"
	V2MoneyManagementPayoutIntentRecipientVerificationMatchResultUnavailable V2MoneyManagementPayoutIntentRecipientVerificationMatchResult = "unavailable"
)

// The business type returned by the recipient-verification provider.
type V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsMatchedBusinessType string

// List of values that V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsMatchedBusinessType can take
const (
	V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsMatchedBusinessTypeBusiness V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsMatchedBusinessType = "business"
	V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsMatchedBusinessTypePersonal V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsMatchedBusinessType = "personal"
)

// The business type provided for recipient verification.
type V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsProvidedBusinessType string

// List of values that V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsProvidedBusinessType can take
const (
	V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsProvidedBusinessTypeBusiness V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsProvidedBusinessType = "business"
	V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsProvidedBusinessTypePersonal V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsProvidedBusinessType = "personal"
)

// Current status of the PayoutIntent.
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
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkACH     V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "ach"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkFedwire V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "fedwire"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkFPS     V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "fps"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkRTP     V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "rtp"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkSEPA    V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "sepa"
	V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkSwift   V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork = "swift"
)

// ACH-specific network fee options.
type V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptionsACH struct {
	// Open Enum. ACH submission timing.
	Submission V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptionsACHSubmission `json:"submission,omitempty"`
}

// Per-network options that affect the fee.
type V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptions struct {
	// ACH-specific network fee options.
	ACH *V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptionsACH `json:"ach,omitempty"`
}

// Details about the network and options associated with this fee. Present when `type` is `network_fee`.
type V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetails struct {
	// The network associated with the fee.
	Network V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetwork `json:"network"`
	// Per-network options that affect the fee.
	NetworkOptions *V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetailsNetworkOptions `json:"network_options"`
}

// Tax charged for this fee, if applicable. Value expressed as a decimal string in major units.
type V2MoneyManagementPayoutIntentEstimatedFeeTaxAmount struct {
	// Currency code.
	Currency Currency `json:"currency"`
	// Tax amount value represented as a decimal string in major units.
	ValueDecimal string `json:"value_decimal"`
}

// Estimated fees and taxes.
type V2MoneyManagementPayoutIntentEstimatedFee struct {
	// The fee amount.
	Amount Amount `json:"amount"`
	// Details about the network and options associated with this fee. Present when `type` is `network_fee`.
	NetworkFeeDetails *V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetails `json:"network_fee_details,omitempty"`
	// Tax charged for this fee, if applicable. Value expressed as a decimal string in major units.
	TaxAmount *V2MoneyManagementPayoutIntentEstimatedFeeTaxAmount `json:"tax_amount,omitempty"`
	// Open Enum. The type of fee.
	Type V2MoneyManagementPayoutIntentEstimatedFeeType `json:"type"`
}

// The FinancialAccount that funds are pulled from.
type V2MoneyManagementPayoutIntentFrom struct {
	// The currency of the financial account.
	Currency Currency `json:"currency"`
	// Estimated amount to be debited from the financial account.
	Debited Amount `json:"debited,omitempty"`
	// The FinancialAccount that funds are pulled from.
	FinancialAccount string `json:"financial_account"`
}

// Maps each source currency to its exchange rate into `to_currency`.
type V2MoneyManagementPayoutIntentFxQuoteRates struct {
	// The exchange rate from the source currency to `to_currency`, represented as a decimal string
	// (for example, `1.1520`) to preserve the full precision of the rate.
	ExchangeRate string `json:"exchange_rate"`
}

// FX rate information for fee transparency.
type V2MoneyManagementPayoutIntentFxQuote struct {
	// Open Enum. Duration of the FX rate lock.
	LockDuration V2MoneyManagementPayoutIntentFxQuoteLockDuration `json:"lock_duration"`
	// Timestamp when the rate lock expires. Null when rate locking is not supported.
	LockExpiresAt time.Time `json:"lock_expires_at,omitempty"`
	// Open Enum. Lock status of the FX rate.
	LockStatus V2MoneyManagementPayoutIntentFxQuoteLockStatus `json:"lock_status"`
	// Maps each source currency to its exchange rate into `to_currency`.
	Rates map[string]*V2MoneyManagementPayoutIntentFxQuoteRates `json:"rates"`
	// The destination currency.
	ToCurrency Currency `json:"to_currency"`
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

// Details about a confirmation required. Populated when `type` is `confirm`.
type V2MoneyManagementPayoutIntentNextActionConfirm struct {
	// The PayoutIntent fields that must be reviewed before confirmation.
	FieldsToReview []V2MoneyManagementPayoutIntentNextActionConfirmFieldsToReview `json:"fields_to_review,omitempty"`
	// Open Enum. The reason the PayoutIntent requires confirmation.
	Reason V2MoneyManagementPayoutIntentNextActionConfirmReason `json:"reason"`
}

// Details about a failure that requires user action. Populated when `type` is `handle_failure`.
type V2MoneyManagementPayoutIntentNextActionHandleFailure struct {
	// Open Enum. The reason for the failure.
	FailureReason V2MoneyManagementPayoutIntentNextActionHandleFailureFailureReason `json:"failure_reason"`
}

// Next action required for a PayoutIntent in the `requires_action` state. Populated when `status` is `requires_action`.
type V2MoneyManagementPayoutIntentNextAction struct {
	// Details about a confirmation required. Populated when `type` is `confirm`.
	Confirm *V2MoneyManagementPayoutIntentNextActionConfirm `json:"confirm,omitempty"`
	// Details about a failure that requires user action. Populated when `type` is `handle_failure`.
	HandleFailure *V2MoneyManagementPayoutIntentNextActionHandleFailure `json:"handle_failure,omitempty"`
	// Open Enum. The type of next action required.
	Type V2MoneyManagementPayoutIntentNextActionType `json:"type"`
}

// Recipient notification settings for the PayoutIntent. Only applicable when the PayoutIntent creates an OutboundPayment.
type V2MoneyManagementPayoutIntentRecipientNotification struct {
	// Closed Enum. Configuration option to enable or disable notifications to recipients.
	// Do not send notifications when `setting` is `none`. Use the account setting when `setting` is `configured` or omitted.
	Setting V2MoneyManagementPayoutIntentRecipientNotificationSetting `json:"setting"`
}

// Details associated with the match result.
type V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetails struct {
	// The business type returned by the recipient-verification provider.
	MatchedBusinessType V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsMatchedBusinessType `json:"matched_business_type,omitempty"`
	// The account name returned by the recipient-verification provider.
	MatchedName string `json:"matched_name,omitempty"`
	// A message describing the match result.
	Message string `json:"message"`
	// The business type provided for recipient verification.
	ProvidedBusinessType V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetailsProvidedBusinessType `json:"provided_business_type,omitempty"`
	// The name associated with the provided recipient.
	ProvidedName string `json:"provided_name"`
}

// Recipient verification details for this PayoutIntent.
type V2MoneyManagementPayoutIntentRecipientVerification struct {
	// Closed Enum. The recipient-verification provider's match result.
	MatchResult V2MoneyManagementPayoutIntentRecipientVerificationMatchResult `json:"match_result"`
	// Details associated with the match result.
	MatchResultDetails *V2MoneyManagementPayoutIntentRecipientVerificationMatchResultDetails `json:"match_result_details"`
}

// Scheduling options for the payout. If omitted, the payout is scheduled for immediate execution.
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
	PreferredNetworks []V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetwork `json:"preferred_networks,omitempty"`
}

// Payout method options for the PayoutIntent.
type V2MoneyManagementPayoutIntentToPayoutMethodOptions struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccount `json:"bank_account,omitempty"`
}

// To which payout method the payout is sent.
type V2MoneyManagementPayoutIntentTo struct {
	// Estimated amount to be credited to the recipient in the destination currency.
	Credited Amount `json:"credited,omitempty"`
	// The currency to send to the recipient.
	Currency Currency `json:"currency,omitempty"`
	// The payout method ID for the PayoutIntent. Optional when the PayoutIntent creates an OutboundPayment and the recipient has a default payout method. Required when the PayoutIntent creates an OutboundTransfer.
	PayoutMethod string `json:"payout_method,omitempty"`
	// Payout method options for the PayoutIntent.
	PayoutMethodOptions *V2MoneyManagementPayoutIntentToPayoutMethodOptions `json:"payout_method_options,omitempty"`
	// The recipient ID for the PayoutIntent.
	Recipient string `json:"recipient,omitempty"`
}

// PayoutIntent represents an intent to send funds from a Financial Account to a payout method.
type V2MoneyManagementPayoutIntent struct {
	APIResource
	// The monetary amount to be sent.
	Amount Amount `json:"amount"`
	// Controls whether the PayoutIntent requires explicit confirmation before transitioning to `pending`.
	// Even when `automatic`, regulatory requirements may require confirmation. See [Send money with Global Payouts](https://docs.stripe.com/global-payouts/send-money) for details.
	ConfirmationMethod V2MoneyManagementPayoutIntentConfirmationMethod `json:"confirmation_method"`
	// Time at which the PayoutIntent was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// An arbitrary string attached to the PayoutIntent. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// Estimated fees and taxes.
	EstimatedFees []*V2MoneyManagementPayoutIntentEstimatedFee `json:"estimated_fees,omitempty"`
	// The FinancialAccount that funds are pulled from.
	From *V2MoneyManagementPayoutIntentFrom `json:"from"`
	// FX rate information for fee transparency.
	FxQuote *V2MoneyManagementPayoutIntentFxQuote `json:"fx_quote,omitempty"`
	// Unique identifier for the PayoutIntent.
	ID string `json:"id"`
	// Details about the latest payout associated with this PayoutIntent.
	LatestPayout *V2MoneyManagementPayoutIntentLatestPayout `json:"latest_payout,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Next action required for a PayoutIntent in the `requires_action` state. Populated when `status` is `requires_action`.
	NextAction *V2MoneyManagementPayoutIntentNextAction `json:"next_action,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Recipient notification settings for the PayoutIntent. Only applicable when the PayoutIntent creates an OutboundPayment.
	RecipientNotification *V2MoneyManagementPayoutIntentRecipientNotification `json:"recipient_notification,omitempty"`
	// Recipient verification details for this PayoutIntent.
	RecipientVerification *V2MoneyManagementPayoutIntentRecipientVerification `json:"recipient_verification,omitempty"`
	// Scheduling options for the payout. If omitted, the payout is scheduled for immediate execution.
	ScheduleOptions *V2MoneyManagementPayoutIntentScheduleOptions `json:"schedule_options,omitempty"`
	// The description that appears on the receiving end for the payout (for example, on a bank statement).
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// Current status of the PayoutIntent.
	Status V2MoneyManagementPayoutIntentStatus `json:"status"`
	// Hash containing timestamps of when transitioned to a particular status.
	StatusTransitions *V2MoneyManagementPayoutIntentStatusTransitions `json:"status_transitions,omitempty"`
	// To which payout method the payout is sent.
	To *V2MoneyManagementPayoutIntentTo `json:"to"`
}
