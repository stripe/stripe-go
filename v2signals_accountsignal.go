//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The effect this indicator had on the overall risk level.
type V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpact string

// List of values that V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpact can take
const (
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpactDecrease       V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpact = "decrease"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpactNeutral        V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpact = "neutral"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpactSlightIncrease V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpact = "slight_increase"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpactStrongIncrease V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpact = "strong_increase"
)

// The name of the specific indicator used in the risk assessment.
type V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator string

// List of values that V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator can take
const (
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicatorBankAccount                           V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator = "bank_account"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicatorBusinessInformationAndAccountActivity V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator = "business_information_and_account_activity"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicatorDisputes                              V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator = "disputes"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicatorFailures                              V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator = "failures"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicatorGeolocation                           V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator = "geolocation"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicatorOther                                 V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator = "other"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicatorOtherRelatedAccounts                  V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator = "other_related_accounts"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicatorOtherTransactionActivity              V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator = "other_transaction_activity"
	V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicatorOwnerEmail                            V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator = "owner_email"
)

// Categorical assessment of the fraudulent merchant risk based on probability.
type V2SignalsAccountSignalFraudulentMerchantRiskLevel string

// List of values that V2SignalsAccountSignalFraudulentMerchantRiskLevel can take
const (
	V2SignalsAccountSignalFraudulentMerchantRiskLevelElevated V2SignalsAccountSignalFraudulentMerchantRiskLevel = "elevated"
	V2SignalsAccountSignalFraudulentMerchantRiskLevelHighest  V2SignalsAccountSignalFraudulentMerchantRiskLevel = "highest"
	V2SignalsAccountSignalFraudulentMerchantRiskLevelLow      V2SignalsAccountSignalFraudulentMerchantRiskLevel = "low"
	V2SignalsAccountSignalFraudulentMerchantRiskLevelNormal   V2SignalsAccountSignalFraudulentMerchantRiskLevel = "normal"
	V2SignalsAccountSignalFraudulentMerchantRiskLevelUnknown  V2SignalsAccountSignalFraudulentMerchantRiskLevel = "unknown"
)

// Categorical assessment of the fraudulent website risk.
type V2SignalsAccountSignalFraudulentWebsiteRiskLevel string

// List of values that V2SignalsAccountSignalFraudulentWebsiteRiskLevel can take
const (
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelElevated V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "elevated"
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelHighest  V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "highest"
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelLow      V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "low"
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelNormal   V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "normal"
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelUnknown  V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "unknown"
)

// The type of signal.
type V2SignalsAccountSignalType string

// List of values that V2SignalsAccountSignalType can take
const (
	V2SignalsAccountSignalTypeFraudulentMerchant  V2SignalsAccountSignalType = "fraudulent_merchant"
	V2SignalsAccountSignalTypeFraudulentWebsite   V2SignalsAccountSignalType = "fraudulent_website"
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

// Array of objects representing individual factors that contributed to the calculated probability. Absent when risk level is unknown,
// or when the user is not on a product tier that includes indicators.
type V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicator struct {
	// A brief explanation of how this indicator contributed to the fraudulent merchant probability.
	Explanation string `json:"explanation"`
	// The effect this indicator had on the overall risk level.
	Impact V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorImpact `json:"impact"`
	// The name of the specific indicator used in the risk assessment.
	Indicator V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicatorIndicator `json:"indicator"`
}

// Supplementary contextual data for the signal, including indicators.
type V2SignalsAccountSignalFraudulentMerchantAdditionalDetails struct {
	// Array of objects representing individual factors that contributed to the calculated probability. Absent when risk level is unknown,
	// or when the user is not on a product tier that includes indicators.
	Indicators []*V2SignalsAccountSignalFraudulentMerchantAdditionalDetailsIndicator `json:"indicators"`
}

// Data for the fraudulent merchant signal. Present only when type is fraudulent_merchant.
type V2SignalsAccountSignalFraudulentMerchant struct {
	// Supplementary contextual data for the signal, including indicators.
	AdditionalDetails *V2SignalsAccountSignalFraudulentMerchantAdditionalDetails `json:"additional_details,omitempty"`
	// The probability of the merchant being fraudulent. Can be between 0.00 and 100.00. Absent when risk level is unknown,
	// or when the user is not on a product tier that includes numeric scores.
	Probability float64 `json:"probability,string,omitempty"`
	// Categorical assessment of the fraudulent merchant risk based on probability.
	RiskLevel V2SignalsAccountSignalFraudulentMerchantRiskLevel `json:"risk_level"`
}

// Data for the fraudulent website signal. Present only when type is fraudulent_website.
type V2SignalsAccountSignalFraudulentWebsite struct {
	// Human-readable details about the fraudulent website evaluation.
	Details string `json:"details,omitempty"`
	// Categorical assessment of the fraudulent website risk.
	RiskLevel V2SignalsAccountSignalFraudulentWebsiteRiskLevel `json:"risk_level"`
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
	// Data for the fraudulent merchant signal. Present only when type is fraudulent_merchant.
	FraudulentMerchant *V2SignalsAccountSignalFraudulentMerchant `json:"fraudulent_merchant,omitempty"`
	// Data for the fraudulent website signal. Present only when type is fraudulent_website.
	FraudulentWebsite *V2SignalsAccountSignalFraudulentWebsite `json:"fraudulent_website,omitempty"`
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
