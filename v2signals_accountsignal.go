//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of signal.
type V2SignalsAccountSignalType string

// List of values that V2SignalsAccountSignalType can take
const (
	V2SignalsAccountSignalTypeUserAccountSharing  V2SignalsAccountSignalType = "user_account_sharing"
	V2SignalsAccountSignalTypeUserMultiAccounting V2SignalsAccountSignalType = "user_multi_accounting"
)

// Categorical assessment of the account-sharing risk.
type V2SignalsAccountSignalUserAccountSharingRiskLevel string

// List of values that V2SignalsAccountSignalUserAccountSharingRiskLevel can take
const (
	V2SignalsAccountSignalUserAccountSharingRiskLevelElevated V2SignalsAccountSignalUserAccountSharingRiskLevel = "elevated"
	V2SignalsAccountSignalUserAccountSharingRiskLevelHighest  V2SignalsAccountSignalUserAccountSharingRiskLevel = "highest"
	V2SignalsAccountSignalUserAccountSharingRiskLevelLow      V2SignalsAccountSignalUserAccountSharingRiskLevel = "low"
	V2SignalsAccountSignalUserAccountSharingRiskLevelNormal   V2SignalsAccountSignalUserAccountSharingRiskLevel = "normal"
	V2SignalsAccountSignalUserAccountSharingRiskLevelUnknown  V2SignalsAccountSignalUserAccountSharingRiskLevel = "unknown"
)

// Categorical assessment of the multi-accounting risk.
type V2SignalsAccountSignalUserMultiAccountingRiskLevel string

// List of values that V2SignalsAccountSignalUserMultiAccountingRiskLevel can take
const (
	V2SignalsAccountSignalUserMultiAccountingRiskLevelElevated V2SignalsAccountSignalUserMultiAccountingRiskLevel = "elevated"
	V2SignalsAccountSignalUserMultiAccountingRiskLevelHighest  V2SignalsAccountSignalUserMultiAccountingRiskLevel = "highest"
	V2SignalsAccountSignalUserMultiAccountingRiskLevelLow      V2SignalsAccountSignalUserMultiAccountingRiskLevel = "low"
	V2SignalsAccountSignalUserMultiAccountingRiskLevelNormal   V2SignalsAccountSignalUserMultiAccountingRiskLevel = "normal"
	V2SignalsAccountSignalUserMultiAccountingRiskLevelUnknown  V2SignalsAccountSignalUserMultiAccountingRiskLevel = "unknown"
)

// The account or customer this signal is associated with.
type V2SignalsAccountSignalAccountDetails struct {
	// The v2 account ID of the account.
	Account string `json:"account,omitempty"`
	// The v1 customer ID of the account, for users not yet migrated to v2/accounts.
	Customer string `json:"customer,omitempty"`
}

// Data for the user account-sharing signal. Present only when type is user_account_sharing.
type V2SignalsAccountSignalUserAccountSharing struct {
	// Categorical assessment of the account-sharing risk.
	RiskLevel V2SignalsAccountSignalUserAccountSharingRiskLevel `json:"risk_level"`
	// The specific risk score for the account, between 0.00 and 100.00. Absent when risk level is
	// not_assessed or unknown, or when the user is not on a product tier that includes numeric scores.
	Score float64 `json:"score,string,omitempty"`
}

// Data for the user multi-accounting signal. Present only when type is user_multi_accounting.
type V2SignalsAccountSignalUserMultiAccounting struct {
	// Categorical assessment of the multi-accounting risk.
	RiskLevel V2SignalsAccountSignalUserMultiAccountingRiskLevel `json:"risk_level"`
	// The specific risk score for the account, between 0.00 and 100.00. Absent when risk level is
	// not_assessed or unknown, or when the user is not on a product tier that includes numeric scores.
	Score float64 `json:"score,string,omitempty"`
}

// An automatically evaluated signal on an account. Each Account Signal object corresponds to
// exactly one signal type, indicated by type. Only the type-specific field is populated; other
// type-specific payload fields are null. If an account has multiple signals, Stripe creates
// separate account signal objects.
type V2SignalsAccountSignal struct {
	APIResource
	// The account or customer this signal is associated with.
	AccountDetails *V2SignalsAccountSignalAccountDetails `json:"account_details,omitempty"`
	// The account evaluation that produced this signal, if applicable.
	AccountEvaluation string `json:"account_evaluation,omitempty"`
	// Timestamp at which the signal was created.
	Created time.Time `json:"created"`
	// Unique identifier for the account signal.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The type of signal.
	Type V2SignalsAccountSignalType `json:"type"`
	// Data for the user account-sharing signal. Present only when type is user_account_sharing.
	UserAccountSharing *V2SignalsAccountSignalUserAccountSharing `json:"user_account_sharing,omitempty"`
	// Data for the user multi-accounting signal. Present only when type is user_multi_accounting.
	UserMultiAccounting *V2SignalsAccountSignalUserMultiAccounting `json:"user_multi_accounting,omitempty"`
}
