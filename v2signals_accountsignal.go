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
	V2SignalsAccountSignalTypeFraudulentMerchant  V2SignalsAccountSignalType = "fraudulent_merchant"
	V2SignalsAccountSignalTypeMerchantDelinquency V2SignalsAccountSignalType = "merchant_delinquency"
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

// An automatically evaluated signal on a v2 account.
type V2SignalsAccountSignal struct {
	APIResource
	// The account or customer this signal is associated with.
	AccountDetails *V2SignalsAccountSignalAccountDetails `json:"account_details,omitempty"`
	// Timestamp at which the signal was created.
	Created time.Time `json:"created"`
	// Data for the fraudulent merchant signal. Present only when type is fraudulent_merchant.
	FraudulentMerchant *V2SignalsAccountSignalFraudulentMerchant `json:"fraudulent_merchant,omitempty"`
	// Unique identifier for the account signal.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Data for the merchant delinquency signal. Present only when type is merchant_delinquency.
	MerchantDelinquency *V2SignalsAccountSignalMerchantDelinquency `json:"merchant_delinquency,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The type of signal.
	Type V2SignalsAccountSignalType `json:"type"`
}
