//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The effect this indicator had on the overall risk level.
type AccountSignalsDelinquencyIndicatorImpact string

// List of values that AccountSignalsDelinquencyIndicatorImpact can take
const (
	AccountSignalsDelinquencyIndicatorImpactDecrease       AccountSignalsDelinquencyIndicatorImpact = "decrease"
	AccountSignalsDelinquencyIndicatorImpactNeutral        AccountSignalsDelinquencyIndicatorImpact = "neutral"
	AccountSignalsDelinquencyIndicatorImpactSlightIncrease AccountSignalsDelinquencyIndicatorImpact = "slight_increase"
	AccountSignalsDelinquencyIndicatorImpactStrongIncrease AccountSignalsDelinquencyIndicatorImpact = "strong_increase"
)

// The name of the specific indicator used in the risk assessment.
type AccountSignalsDelinquencyIndicatorIndicator string

// List of values that AccountSignalsDelinquencyIndicatorIndicator can take
const (
	AccountSignalsDelinquencyIndicatorIndicatorAccountBalance      AccountSignalsDelinquencyIndicatorIndicator = "account_balance"
	AccountSignalsDelinquencyIndicatorIndicatorAov                 AccountSignalsDelinquencyIndicatorIndicator = "aov"
	AccountSignalsDelinquencyIndicatorIndicatorChargeConcentration AccountSignalsDelinquencyIndicatorIndicator = "charge_concentration"
	AccountSignalsDelinquencyIndicatorIndicatorDisputeWindow       AccountSignalsDelinquencyIndicatorIndicator = "dispute_window"
	AccountSignalsDelinquencyIndicatorIndicatorDisputes            AccountSignalsDelinquencyIndicatorIndicator = "disputes"
	AccountSignalsDelinquencyIndicatorIndicatorDuplicates          AccountSignalsDelinquencyIndicatorIndicator = "duplicates"
	AccountSignalsDelinquencyIndicatorIndicatorExposure            AccountSignalsDelinquencyIndicatorIndicator = "exposure"
	AccountSignalsDelinquencyIndicatorIndicatorFirmographic        AccountSignalsDelinquencyIndicatorIndicator = "firmographic"
	AccountSignalsDelinquencyIndicatorIndicatorLifetimeMetrics     AccountSignalsDelinquencyIndicatorIndicator = "lifetime_metrics"
	AccountSignalsDelinquencyIndicatorIndicatorPaymentProcessing   AccountSignalsDelinquencyIndicatorIndicator = "payment_processing"
	AccountSignalsDelinquencyIndicatorIndicatorPaymentVolume       AccountSignalsDelinquencyIndicatorIndicator = "payment_volume"
	AccountSignalsDelinquencyIndicatorIndicatorPayouts             AccountSignalsDelinquencyIndicatorIndicator = "payouts"
	AccountSignalsDelinquencyIndicatorIndicatorRefunds             AccountSignalsDelinquencyIndicatorIndicator = "refunds"
	AccountSignalsDelinquencyIndicatorIndicatorRelatedAccounts     AccountSignalsDelinquencyIndicatorIndicator = "related_accounts"
	AccountSignalsDelinquencyIndicatorIndicatorTenure              AccountSignalsDelinquencyIndicatorIndicator = "tenure"
	AccountSignalsDelinquencyIndicatorIndicatorTransfers           AccountSignalsDelinquencyIndicatorIndicator = "transfers"
)

// Categorical assessment of the delinquency risk based on probability.
type AccountSignalsDelinquencyRiskLevel string

// List of values that AccountSignalsDelinquencyRiskLevel can take
const (
	AccountSignalsDelinquencyRiskLevelElevated    AccountSignalsDelinquencyRiskLevel = "elevated"
	AccountSignalsDelinquencyRiskLevelHighest     AccountSignalsDelinquencyRiskLevel = "highest"
	AccountSignalsDelinquencyRiskLevelLow         AccountSignalsDelinquencyRiskLevel = "low"
	AccountSignalsDelinquencyRiskLevelNormal      AccountSignalsDelinquencyRiskLevel = "normal"
	AccountSignalsDelinquencyRiskLevelNotAssessed AccountSignalsDelinquencyRiskLevel = "not_assessed"
	AccountSignalsDelinquencyRiskLevelUnknown     AccountSignalsDelinquencyRiskLevel = "unknown"
)

// The effect this indicator had on the overall risk level.
type AccountSignalsFraudIntentIndicatorImpact string

// List of values that AccountSignalsFraudIntentIndicatorImpact can take
const (
	AccountSignalsFraudIntentIndicatorImpactDecrease       AccountSignalsFraudIntentIndicatorImpact = "decrease"
	AccountSignalsFraudIntentIndicatorImpactNeutral        AccountSignalsFraudIntentIndicatorImpact = "neutral"
	AccountSignalsFraudIntentIndicatorImpactSlightIncrease AccountSignalsFraudIntentIndicatorImpact = "slight_increase"
	AccountSignalsFraudIntentIndicatorImpactStrongIncrease AccountSignalsFraudIntentIndicatorImpact = "strong_increase"
)

// The name of the specific indicator used in the risk assessment.
type AccountSignalsFraudIntentIndicatorIndicator string

// List of values that AccountSignalsFraudIntentIndicatorIndicator can take
const (
	AccountSignalsFraudIntentIndicatorIndicatorBankAccount                           AccountSignalsFraudIntentIndicatorIndicator = "bank_account"
	AccountSignalsFraudIntentIndicatorIndicatorBusinessInformationAndAccountActivity AccountSignalsFraudIntentIndicatorIndicator = "business_information_and_account_activity"
	AccountSignalsFraudIntentIndicatorIndicatorDisputes                              AccountSignalsFraudIntentIndicatorIndicator = "disputes"
	AccountSignalsFraudIntentIndicatorIndicatorFailures                              AccountSignalsFraudIntentIndicatorIndicator = "failures"
	AccountSignalsFraudIntentIndicatorIndicatorGeoLocation                           AccountSignalsFraudIntentIndicatorIndicator = "geo_location"
	AccountSignalsFraudIntentIndicatorIndicatorOther                                 AccountSignalsFraudIntentIndicatorIndicator = "other"
	AccountSignalsFraudIntentIndicatorIndicatorOtherRelatedAccounts                  AccountSignalsFraudIntentIndicatorIndicator = "other_related_accounts"
	AccountSignalsFraudIntentIndicatorIndicatorOtherTransactionActivity              AccountSignalsFraudIntentIndicatorIndicator = "other_transaction_activity"
	AccountSignalsFraudIntentIndicatorIndicatorOwnerEmail                            AccountSignalsFraudIntentIndicatorIndicator = "owner_email"
	AccountSignalsFraudIntentIndicatorIndicatorWebPresence                           AccountSignalsFraudIntentIndicatorIndicator = "web_presence"
)

// Categorical assessment of the fraud intent risk based on probability.
type AccountSignalsFraudIntentRiskLevel string

// List of values that AccountSignalsFraudIntentRiskLevel can take
const (
	AccountSignalsFraudIntentRiskLevelElevated    AccountSignalsFraudIntentRiskLevel = "elevated"
	AccountSignalsFraudIntentRiskLevelHighest     AccountSignalsFraudIntentRiskLevel = "highest"
	AccountSignalsFraudIntentRiskLevelLow         AccountSignalsFraudIntentRiskLevel = "low"
	AccountSignalsFraudIntentRiskLevelNormal      AccountSignalsFraudIntentRiskLevel = "normal"
	AccountSignalsFraudIntentRiskLevelNotAssessed AccountSignalsFraudIntentRiskLevel = "not_assessed"
	AccountSignalsFraudIntentRiskLevelUnknown     AccountSignalsFraudIntentRiskLevel = "unknown"
)

// Retrieves the account's Signal objects
type AccountSignalsParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *AccountSignalsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the account's Signal objects
type AccountSignalsRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *AccountSignalsRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Array of objects representing individual factors that contributed to the calculated probability of delinquency.
type AccountSignalsDelinquencyIndicator struct {
	// A brief explanation of how this indicator contributed to the delinquency probability.
	Description string `json:"description"`
	// The effect this indicator had on the overall risk level.
	Impact AccountSignalsDelinquencyIndicatorImpact `json:"impact"`
	// The name of the specific indicator used in the risk assessment.
	Indicator AccountSignalsDelinquencyIndicatorIndicator `json:"indicator"`
}

// The delinquency signal of the account.
type AccountSignalsDelinquency struct {
	// Time at which the signal was evaluated, measured in seconds since the Unix epoch.
	EvaluatedAt int64 `json:"evaluated_at"`
	// Array of objects representing individual factors that contributed to the calculated probability of delinquency.
	Indicators []*AccountSignalsDelinquencyIndicator `json:"indicators"`
	// The probability of delinquency. Can be between 0.00 and 100.00
	Probability float64 `json:"probability"`
	// Categorical assessment of the delinquency risk based on probability.
	RiskLevel AccountSignalsDelinquencyRiskLevel `json:"risk_level"`
	// Unique identifier for the delinquency signal.
	SignalID string `json:"signal_id"`
}

// Array of objects representing individual factors that contributed to the calculated probability of fraud intent.
type AccountSignalsFraudIntentIndicator struct {
	// A brief explanation of how this indicator contributed to the delinquency probability.
	Description string `json:"description"`
	// The effect this indicator had on the overall risk level.
	Impact AccountSignalsFraudIntentIndicatorImpact `json:"impact"`
	// The name of the specific indicator used in the risk assessment.
	Indicator AccountSignalsFraudIntentIndicatorIndicator `json:"indicator"`
}

// The fraud intent signal of the account.
type AccountSignalsFraudIntent struct {
	// Time at which the signal was evaluated, measured in seconds since the Unix epoch.
	EvaluatedAt int64 `json:"evaluated_at"`
	// Array of objects representing individual factors that contributed to the calculated probability of fraud intent.
	Indicators []*AccountSignalsFraudIntentIndicator `json:"indicators"`
	// The probability of fraud intent. Can be between 0.00 and 100.00
	Probability float64 `json:"probability"`
	// Categorical assessment of the fraud intent risk based on probability.
	RiskLevel AccountSignalsFraudIntentRiskLevel `json:"risk_level"`
	// Unique identifier for the fraud intent signal.
	SignalID string `json:"signal_id"`
}

// The Account Signals API provides risk related signals that can be used to better manage risks.
type AccountSignals struct {
	APIResource
	// The account for which the signals belong to.
	Account string `json:"account"`
	// The delinquency signal of the account.
	Delinquency *AccountSignalsDelinquency `json:"delinquency,omitempty"`
	// The fraud intent signal of the account.
	FraudIntent *AccountSignalsFraudIntent `json:"fraud_intent,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
