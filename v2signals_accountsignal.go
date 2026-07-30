//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The effect this indicator had on the overall risk level.
type V2SignalsAccountSignalFraudulentMerchantIndicatorImpact string

// List of values that V2SignalsAccountSignalFraudulentMerchantIndicatorImpact can take
const (
	V2SignalsAccountSignalFraudulentMerchantIndicatorImpactDecrease       V2SignalsAccountSignalFraudulentMerchantIndicatorImpact = "decrease"
	V2SignalsAccountSignalFraudulentMerchantIndicatorImpactNeutral        V2SignalsAccountSignalFraudulentMerchantIndicatorImpact = "neutral"
	V2SignalsAccountSignalFraudulentMerchantIndicatorImpactSlightIncrease V2SignalsAccountSignalFraudulentMerchantIndicatorImpact = "slight_increase"
	V2SignalsAccountSignalFraudulentMerchantIndicatorImpactStrongIncrease V2SignalsAccountSignalFraudulentMerchantIndicatorImpact = "strong_increase"
)

// The name of the specific indicator used in the risk assessment.
type V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator string

// List of values that V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator can take
const (
	V2SignalsAccountSignalFraudulentMerchantIndicatorIndicatorBankAccount                           V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator = "bank_account"
	V2SignalsAccountSignalFraudulentMerchantIndicatorIndicatorBusinessInformationAndAccountActivity V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator = "business_information_and_account_activity"
	V2SignalsAccountSignalFraudulentMerchantIndicatorIndicatorDisputes                              V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator = "disputes"
	V2SignalsAccountSignalFraudulentMerchantIndicatorIndicatorFailures                              V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator = "failures"
	V2SignalsAccountSignalFraudulentMerchantIndicatorIndicatorGeolocation                           V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator = "geolocation"
	V2SignalsAccountSignalFraudulentMerchantIndicatorIndicatorOther                                 V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator = "other"
	V2SignalsAccountSignalFraudulentMerchantIndicatorIndicatorOtherRelatedAccounts                  V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator = "other_related_accounts"
	V2SignalsAccountSignalFraudulentMerchantIndicatorIndicatorOtherTransactionActivity              V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator = "other_transaction_activity"
	V2SignalsAccountSignalFraudulentMerchantIndicatorIndicatorOwnerEmail                            V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator = "owner_email"
)

// Categorical assessment of the fraudulent merchant risk based on probability.
type V2SignalsAccountSignalFraudulentMerchantRiskLevel string

// List of values that V2SignalsAccountSignalFraudulentMerchantRiskLevel can take
const (
	V2SignalsAccountSignalFraudulentMerchantRiskLevelElevated    V2SignalsAccountSignalFraudulentMerchantRiskLevel = "elevated"
	V2SignalsAccountSignalFraudulentMerchantRiskLevelHighest     V2SignalsAccountSignalFraudulentMerchantRiskLevel = "highest"
	V2SignalsAccountSignalFraudulentMerchantRiskLevelLow         V2SignalsAccountSignalFraudulentMerchantRiskLevel = "low"
	V2SignalsAccountSignalFraudulentMerchantRiskLevelNormal      V2SignalsAccountSignalFraudulentMerchantRiskLevel = "normal"
	V2SignalsAccountSignalFraudulentMerchantRiskLevelNotAssessed V2SignalsAccountSignalFraudulentMerchantRiskLevel = "not_assessed"
	V2SignalsAccountSignalFraudulentMerchantRiskLevelUnknown     V2SignalsAccountSignalFraudulentMerchantRiskLevel = "unknown"
)

// Categorical assessment of the fraudulent website risk.
type V2SignalsAccountSignalFraudulentWebsiteRiskLevel string

// List of values that V2SignalsAccountSignalFraudulentWebsiteRiskLevel can take
const (
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelElevated    V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "elevated"
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelHighest     V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "highest"
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelLow         V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "low"
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelNormal      V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "normal"
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelNotAssessed V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "not_assessed"
	V2SignalsAccountSignalFraudulentWebsiteRiskLevelUnknown     V2SignalsAccountSignalFraudulentWebsiteRiskLevel = "unknown"
)

// The effect this indicator had on the overall risk level.
type V2SignalsAccountSignalMerchantDelinquencyIndicatorImpact string

// List of values that V2SignalsAccountSignalMerchantDelinquencyIndicatorImpact can take
const (
	V2SignalsAccountSignalMerchantDelinquencyIndicatorImpactDecrease       V2SignalsAccountSignalMerchantDelinquencyIndicatorImpact = "decrease"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorImpactNeutral        V2SignalsAccountSignalMerchantDelinquencyIndicatorImpact = "neutral"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorImpactSlightIncrease V2SignalsAccountSignalMerchantDelinquencyIndicatorImpact = "slight_increase"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorImpactStrongIncrease V2SignalsAccountSignalMerchantDelinquencyIndicatorImpact = "strong_increase"
)

// The name of the specific indicator used in the risk assessment.
type V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator string

// List of values that V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator can take
const (
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorAccountBalance      V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "account_balance"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorAov                 V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "aov"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorChargeConcentration V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "charge_concentration"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorDisputes            V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "disputes"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorDisputeWindow       V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "dispute_window"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorExposure            V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "exposure"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorFirmographic        V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "firmographic"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorLifetimeMetrics     V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "lifetime_metrics"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorOther               V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "other"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorPaymentProcessing   V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "payment_processing"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorPaymentVolume       V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "payment_volume"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorPayouts             V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "payouts"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorRefunds             V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "refunds"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorRelatedAccounts     V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "related_accounts"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorTenure              V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "tenure"
	V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicatorTransfers           V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator = "transfers"
)

// Categorical assessment of the delinquency risk based on probability.
type V2SignalsAccountSignalMerchantDelinquencyRiskLevel string

// List of values that V2SignalsAccountSignalMerchantDelinquencyRiskLevel can take
const (
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelElevated    V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "elevated"
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelHighest     V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "highest"
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelLow         V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "low"
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelNormal      V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "normal"
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelNotAssessed V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "not_assessed"
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelUnknown     V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "unknown"
)

// The type of signal.
type V2SignalsAccountSignalType string

// List of values that V2SignalsAccountSignalType can take
const (
	V2SignalsAccountSignalTypeFraudulentMerchant         V2SignalsAccountSignalType = "fraudulent_merchant"
	V2SignalsAccountSignalTypeFraudulentWebsite          V2SignalsAccountSignalType = "fraudulent_website"
	V2SignalsAccountSignalTypeMerchantDelinquency        V2SignalsAccountSignalType = "merchant_delinquency"
	V2SignalsAccountSignalTypePaymentDelinquencyExposure V2SignalsAccountSignalType = "payment_delinquency_exposure"
	V2SignalsAccountSignalTypeUserAccountSharing         V2SignalsAccountSignalType = "user_account_sharing"
	V2SignalsAccountSignalTypeUserMultiAccounting        V2SignalsAccountSignalType = "user_multi_accounting"
)

// Categorical assessment of the account-sharing risk.
type V2SignalsAccountSignalUserAccountSharingRiskLevel string

// List of values that V2SignalsAccountSignalUserAccountSharingRiskLevel can take
const (
	V2SignalsAccountSignalUserAccountSharingRiskLevelElevated    V2SignalsAccountSignalUserAccountSharingRiskLevel = "elevated"
	V2SignalsAccountSignalUserAccountSharingRiskLevelHighest     V2SignalsAccountSignalUserAccountSharingRiskLevel = "highest"
	V2SignalsAccountSignalUserAccountSharingRiskLevelLow         V2SignalsAccountSignalUserAccountSharingRiskLevel = "low"
	V2SignalsAccountSignalUserAccountSharingRiskLevelNormal      V2SignalsAccountSignalUserAccountSharingRiskLevel = "normal"
	V2SignalsAccountSignalUserAccountSharingRiskLevelNotAssessed V2SignalsAccountSignalUserAccountSharingRiskLevel = "not_assessed"
	V2SignalsAccountSignalUserAccountSharingRiskLevelUnknown     V2SignalsAccountSignalUserAccountSharingRiskLevel = "unknown"
)

// Categorical assessment of the multi-accounting risk.
type V2SignalsAccountSignalUserMultiAccountingRiskLevel string

// List of values that V2SignalsAccountSignalUserMultiAccountingRiskLevel can take
const (
	V2SignalsAccountSignalUserMultiAccountingRiskLevelElevated    V2SignalsAccountSignalUserMultiAccountingRiskLevel = "elevated"
	V2SignalsAccountSignalUserMultiAccountingRiskLevelHighest     V2SignalsAccountSignalUserMultiAccountingRiskLevel = "highest"
	V2SignalsAccountSignalUserMultiAccountingRiskLevelLow         V2SignalsAccountSignalUserMultiAccountingRiskLevel = "low"
	V2SignalsAccountSignalUserMultiAccountingRiskLevelNormal      V2SignalsAccountSignalUserMultiAccountingRiskLevel = "normal"
	V2SignalsAccountSignalUserMultiAccountingRiskLevelNotAssessed V2SignalsAccountSignalUserMultiAccountingRiskLevel = "not_assessed"
	V2SignalsAccountSignalUserMultiAccountingRiskLevelUnknown     V2SignalsAccountSignalUserMultiAccountingRiskLevel = "unknown"
)

// The account or customer this signal is associated with.
type V2SignalsAccountSignalAccountDetails struct {
	// The v2 account ID of the account.
	Account string `json:"account,omitempty"`
	// The v1 customer ID of the account, for users not yet migrated to v2/accounts.
	Customer string `json:"customer,omitempty"`
}

// Array of objects representing individual factors that contributed to the calculated probability. Absent when risk level is not_assessed or unknown,
// or when the user is not on a product tier that includes indicators.
type V2SignalsAccountSignalFraudulentMerchantIndicator struct {
	// A brief explanation of how this indicator contributed to the fraudulent merchant probability.
	Explanation string `json:"explanation"`
	// The effect this indicator had on the overall risk level.
	Impact V2SignalsAccountSignalFraudulentMerchantIndicatorImpact `json:"impact"`
	// The name of the specific indicator used in the risk assessment.
	Indicator V2SignalsAccountSignalFraudulentMerchantIndicatorIndicator `json:"indicator"`
}

// Data for the fraudulent merchant signal. Present only when type is fraudulent_merchant.
type V2SignalsAccountSignalFraudulentMerchant struct {
	// Array of objects representing individual factors that contributed to the calculated probability. Absent when risk level is not_assessed or unknown,
	// or when the user is not on a product tier that includes indicators.
	Indicators []*V2SignalsAccountSignalFraudulentMerchantIndicator `json:"indicators"`
	// The probability of the merchant being fraudulent. Can be between 0.00 and 100.00. Absent when risk level is not_assessed or unknown,
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

// Array of objects representing individual factors that contributed to the calculated probability of delinquency. Absent when risk level is not_assessed or unknown,
// or when the user is not on a product tier that includes indicators.
type V2SignalsAccountSignalMerchantDelinquencyIndicator struct {
	// A brief explanation of how this indicator contributed to the delinquency probability.
	Explanation string `json:"explanation"`
	// The effect this indicator had on the overall risk level.
	Impact V2SignalsAccountSignalMerchantDelinquencyIndicatorImpact `json:"impact"`
	// The name of the specific indicator used in the risk assessment.
	Indicator V2SignalsAccountSignalMerchantDelinquencyIndicatorIndicator `json:"indicator"`
}

// Data for the merchant delinquency signal. Present only when type is merchant_delinquency.
type V2SignalsAccountSignalMerchantDelinquency struct {
	// Array of objects representing individual factors that contributed to the calculated probability of delinquency. Absent when risk level is not_assessed or unknown,
	// or when the user is not on a product tier that includes indicators.
	Indicators []*V2SignalsAccountSignalMerchantDelinquencyIndicator `json:"indicators"`
	// The probability of delinquency. Can be between 0.00 and 100.00. Absent when risk level is not_assessed or unknown,
	// or when the user is not on a product tier that includes numeric scores.
	Probability float64 `json:"probability,string,omitempty"`
	// Categorical assessment of the delinquency risk based on probability.
	RiskLevel V2SignalsAccountSignalMerchantDelinquencyRiskLevel `json:"risk_level"`
}

// Total payments still exposed to dispute or refund risk in the event of delinquency.
type V2SignalsAccountSignalPaymentDelinquencyExposureAdditionalDetailsGrossExposureAmount struct {
	// ISO 4217 currency code.
	Currency Currency `json:"currency"`
	// Amount in minor units for the given currency.
	Value int64 `json:"value,string"`
}

// Additional details about the exposure assessment.
type V2SignalsAccountSignalPaymentDelinquencyExposureAdditionalDetails struct {
	// Total payments still exposed to dispute or refund risk in the event of delinquency.
	GrossExposureAmount *V2SignalsAccountSignalPaymentDelinquencyExposureAdditionalDetailsGrossExposureAmount `json:"gross_exposure_amount,omitempty"`
	// Percentage of Gross Exposure expected to be disputed or refunded and materialize as a loss in the event of delinquency.
	LossGivenDefaultInPercentages int64 `json:"loss_given_default_in_percentages,omitempty"`
	// Predicted window size in days until dispute is raised.
	PredictedDisputeWindowInDays int64 `json:"predicted_dispute_window_in_days,omitempty"`
}

// The exposure amount if this account becomes delinquent.
type V2SignalsAccountSignalPaymentDelinquencyExposureExposureAmount struct {
	// ISO 4217 currency code.
	Currency Currency `json:"currency"`
	// Amount in minor units for the given currency.
	Value int64 `json:"value,string"`
}

// Data for the payment delinquency exposure signal. Present only when type is payment_delinquency_exposure.
type V2SignalsAccountSignalPaymentDelinquencyExposure struct {
	// Additional details about the exposure assessment.
	AdditionalDetails *V2SignalsAccountSignalPaymentDelinquencyExposureAdditionalDetails `json:"additional_details"`
	// The exposure amount if this account becomes delinquent.
	ExposureAmount *V2SignalsAccountSignalPaymentDelinquencyExposureExposureAmount `json:"exposure_amount"`
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
	// Data for the merchant delinquency signal. Present only when type is merchant_delinquency.
	MerchantDelinquency *V2SignalsAccountSignalMerchantDelinquency `json:"merchant_delinquency,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Data for the payment delinquency exposure signal. Present only when type is payment_delinquency_exposure.
	PaymentDelinquencyExposure *V2SignalsAccountSignalPaymentDelinquencyExposure `json:"payment_delinquency_exposure,omitempty"`
	// The type of signal.
	Type V2SignalsAccountSignalType `json:"type"`
	// Data for the user account-sharing signal. Present only when type is user_account_sharing.
	UserAccountSharing *V2SignalsAccountSignalUserAccountSharing `json:"user_account_sharing,omitempty"`
	// Data for the user multi-accounting signal. Present only when type is user_multi_accounting.
	UserMultiAccounting *V2SignalsAccountSignalUserMultiAccounting `json:"user_multi_accounting,omitempty"`
}
