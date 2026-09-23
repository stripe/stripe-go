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

// The effect this indicator had on the overall risk level.
type V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpact string

// List of values that V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpact can take
const (
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpactDecrease       V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpact = "decrease"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpactNeutral        V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpact = "neutral"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpactSlightIncrease V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpact = "slight_increase"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpactStrongIncrease V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpact = "strong_increase"
)

// The name of the specific indicator used in the risk assessment.
type V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator string

// List of values that V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator can take
const (
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorAccountBalance      V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "account_balance"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorAov                 V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "aov"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorChargeConcentration V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "charge_concentration"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorDisputes            V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "disputes"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorDisputeWindow       V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "dispute_window"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorExposure            V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "exposure"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorFirmographic        V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "firmographic"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorLifetimeMetrics     V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "lifetime_metrics"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorOther               V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "other"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorPaymentProcessing   V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "payment_processing"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorPaymentVolume       V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "payment_volume"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorPayouts             V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "payouts"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorRefunds             V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "refunds"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorRelatedAccounts     V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "related_accounts"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorTenure              V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "tenure"
	V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicatorTransfers           V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator = "transfers"
)

// Categorical assessment of the delinquency risk based on probability.
type V2SignalsAccountSignalMerchantDelinquencyRiskLevel string

// List of values that V2SignalsAccountSignalMerchantDelinquencyRiskLevel can take
const (
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelElevated V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "elevated"
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelHighest  V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "highest"
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelLow      V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "low"
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelNormal   V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "normal"
	V2SignalsAccountSignalMerchantDelinquencyRiskLevelUnknown  V2SignalsAccountSignalMerchantDelinquencyRiskLevel = "unknown"
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

// Array of objects representing individual factors that contributed to the calculated probability of delinquency. Absent when risk level is unknown,
// or when the user is not on a product tier that includes indicators.
type V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicator struct {
	// A brief explanation of how this indicator contributed to the delinquency probability.
	Explanation string `json:"explanation"`
	// The effect this indicator had on the overall risk level.
	Impact V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorImpact `json:"impact"`
	// The name of the specific indicator used in the risk assessment.
	Indicator V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicatorIndicator `json:"indicator"`
}

// Supplementary contextual data for the signal, including indicators.
type V2SignalsAccountSignalMerchantDelinquencyAdditionalDetails struct {
	// Array of objects representing individual factors that contributed to the calculated probability of delinquency. Absent when risk level is unknown,
	// or when the user is not on a product tier that includes indicators.
	Indicators []*V2SignalsAccountSignalMerchantDelinquencyAdditionalDetailsIndicator `json:"indicators"`
}

// Data for the merchant delinquency signal. Present only when type is merchant_delinquency.
type V2SignalsAccountSignalMerchantDelinquency struct {
	// Supplementary contextual data for the signal, including indicators.
	AdditionalDetails *V2SignalsAccountSignalMerchantDelinquencyAdditionalDetails `json:"additional_details,omitempty"`
	// The probability of delinquency. Can be between 0.00 and 100.00. Absent when risk level is unknown,
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
