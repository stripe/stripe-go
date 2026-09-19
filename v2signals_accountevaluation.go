//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Categorical assessment of the fraudulent website risk.
type V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevel string

// List of values that V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevel can take
const (
	V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevelElevated V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevel = "elevated"
	V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevelHighest  V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevel = "highest"
	V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevelLow      V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevel = "low"
	V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevelNormal   V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevel = "normal"
	V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevelUnknown  V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevel = "unknown"
)

// Categorical assessment of the account-sharing risk.
type V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevel string

// List of values that V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevel can take
const (
	V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevelElevated V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevel = "elevated"
	V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevelHighest  V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevel = "highest"
	V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevelLow      V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevel = "low"
	V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevelNormal   V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevel = "normal"
	V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevelUnknown  V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevel = "unknown"
)

// Categorical assessment of the multi-accounting risk.
type V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevel string

// List of values that V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevel can take
const (
	V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevelElevated V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevel = "elevated"
	V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevelHighest  V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevel = "highest"
	V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevelLow      V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevel = "low"
	V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevelNormal   V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevel = "normal"
	V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevelUnknown  V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevel = "unknown"
)

// List of signals still pending evaluation.
type V2SignalsAccountEvaluationPendingSignal string

// List of values that V2SignalsAccountEvaluationPendingSignal can take
const (
	V2SignalsAccountEvaluationPendingSignalFraudulentWebsite   V2SignalsAccountEvaluationPendingSignal = "fraudulent_website"
	V2SignalsAccountEvaluationPendingSignalUserAccountSharing  V2SignalsAccountEvaluationPendingSignal = "user_account_sharing"
	V2SignalsAccountEvaluationPendingSignalUserMultiAccounting V2SignalsAccountEvaluationPendingSignal = "user_multi_accounting"
)

// List of signals requested for evaluation.
type V2SignalsAccountEvaluationRequestedSignal string

// List of values that V2SignalsAccountEvaluationRequestedSignal can take
const (
	V2SignalsAccountEvaluationRequestedSignalFraudulentWebsite   V2SignalsAccountEvaluationRequestedSignal = "fraudulent_website"
	V2SignalsAccountEvaluationRequestedSignalUserAccountSharing  V2SignalsAccountEvaluationRequestedSignal = "user_account_sharing"
	V2SignalsAccountEvaluationRequestedSignalUserMultiAccounting V2SignalsAccountEvaluationRequestedSignal = "user_multi_accounting"
)

// Account activity recorded alongside this evaluation, when applicable.
type V2SignalsAccountEvaluationAccountActivityDetails struct {
	// The ID of the account activity created or associated with the evaluation.
	AccountActivity string `json:"account_activity,omitempty"`
}

// Account profile data.
type V2SignalsAccountEvaluationAccountDetailsDataDefaultsProfile struct {
	// The business URL.
	BusinessURL string `json:"business_url"`
	// Doing business as (DBA) name.
	DoingBusinessAs string `json:"doing_business_as,omitempty"`
	// Description of the account's product or service.
	ProductDescription string `json:"product_description,omitempty"`
}

// Default account settings.
type V2SignalsAccountEvaluationAccountDetailsDataDefaults struct {
	// Account profile data.
	Profile *V2SignalsAccountEvaluationAccountDetailsDataDefaultsProfile `json:"profile"`
}

// Business details for identity data.
type V2SignalsAccountEvaluationAccountDetailsDataIdentityBusinessDetails struct {
	// Registered business name.
	RegisteredName string `json:"registered_name,omitempty"`
}

// Identity data.
type V2SignalsAccountEvaluationAccountDetailsDataIdentity struct {
	// Business details for identity data.
	BusinessDetails *V2SignalsAccountEvaluationAccountDetailsDataIdentityBusinessDetails `json:"business_details"`
}

// Inline account data to evaluate without creating a v2 account.
type V2SignalsAccountEvaluationAccountDetailsData struct {
	// Default account settings.
	Defaults *V2SignalsAccountEvaluationAccountDetailsDataDefaults `json:"defaults,omitempty"`
	// Identity data.
	Identity *V2SignalsAccountEvaluationAccountDetailsDataIdentity `json:"identity,omitempty"`
}

// The account, customer, or inline account data being evaluated.
type V2SignalsAccountEvaluationAccountDetails struct {
	// The v2 account ID of the account.
	Account string `json:"account,omitempty"`
	// The v1 customer ID of the account, for users not yet migrated to v2/accounts.
	Customer string `json:"customer,omitempty"`
	// Inline account data to evaluate without creating a v2 account.
	Data *V2SignalsAccountEvaluationAccountDetailsData `json:"data,omitempty"`
}

// Fraudulent website result for the evaluation, when available.
type V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsite struct {
	// Human-readable details about the fraudulent website evaluation, when available.
	Details string `json:"details,omitempty"`
	// Timestamp at which the signal was evaluated.
	EvaluatedAt time.Time `json:"evaluated_at,omitempty"`
	// Categorical assessment of the fraudulent website risk.
	RiskLevel V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsiteRiskLevel `json:"risk_level"`
	// The account signal ID containing the full fraudulent website signal result.
	Signal string `json:"signal,omitempty"`
}

// User account-sharing result for the evaluation, when available.
type V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharing struct {
	// Timestamp at which the signal was evaluated.
	EvaluatedAt time.Time `json:"evaluated_at,omitempty"`
	// Categorical assessment of the account-sharing risk.
	RiskLevel V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharingRiskLevel `json:"risk_level"`
	// The specific risk score for the account, between 0.00 and 100.00, when available.
	Score float64 `json:"score,string,omitempty"`
	// The account signal ID containing the full user account-sharing signal result.
	Signal string `json:"signal,omitempty"`
}

// User multi-accounting result for the evaluation, when available.
type V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccounting struct {
	// Timestamp at which the signal was evaluated.
	EvaluatedAt time.Time `json:"evaluated_at,omitempty"`
	// Categorical assessment of the multi-accounting risk.
	RiskLevel V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccountingRiskLevel `json:"risk_level"`
	// The specific risk score for the account, between 0.00 and 100.00, when available.
	Score float64 `json:"score,string,omitempty"`
	// The account signal ID containing the full user multi-accounting signal result.
	Signal string `json:"signal,omitempty"`
}

// Signal results that are available for the evaluation.
type V2SignalsAccountEvaluationEvaluatedSignals struct {
	// Fraudulent website result for the evaluation, when available.
	FraudulentWebsite *V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsite `json:"fraudulent_website,omitempty"`
	// User account-sharing result for the evaluation, when available.
	UserAccountSharing *V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharing `json:"user_account_sharing,omitempty"`
	// User multi-accounting result for the evaluation, when available.
	UserMultiAccounting *V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccounting `json:"user_multi_accounting,omitempty"`
}

// Account Evaluation resource for the Signals API.
type V2SignalsAccountEvaluation struct {
	APIResource
	// Account activity recorded alongside this evaluation, when applicable.
	AccountActivityDetails *V2SignalsAccountEvaluationAccountActivityDetails `json:"account_activity_details,omitempty"`
	// The account, customer, or inline account data being evaluated.
	AccountDetails *V2SignalsAccountEvaluationAccountDetails `json:"account_details"`
	// Timestamp at which the evaluation was created.
	Created time.Time `json:"created"`
	// Signal results that are available for the evaluation.
	EvaluatedSignals *V2SignalsAccountEvaluationEvaluatedSignals `json:"evaluated_signals,omitempty"`
	// Unique identifier for the account evaluation.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// List of signals still pending evaluation.
	PendingSignals []V2SignalsAccountEvaluationPendingSignal `json:"pending_signals"`
	// List of signals requested for evaluation.
	RequestedSignals []V2SignalsAccountEvaluationRequestedSignal `json:"requested_signals"`
}
