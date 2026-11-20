//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"github.com/shopspring/decimal"
	"time"
)

// The financial objective the recommended controls are intended to improve.
type V2RiskEvaluationOutcomeDelinquencyBusinessObjectiveType string

// List of values that V2RiskEvaluationOutcomeDelinquencyBusinessObjectiveType can take
const (
	V2RiskEvaluationOutcomeDelinquencyBusinessObjectiveTypeProfitability V2RiskEvaluationOutcomeDelinquencyBusinessObjectiveType = "profitability"
)

// Whether to update, maintain, or clear reserves.
type V2RiskEvaluationOutcomeDelinquencyControlRiskReserveAction string

// List of values that V2RiskEvaluationOutcomeDelinquencyControlRiskReserveAction can take
const (
	V2RiskEvaluationOutcomeDelinquencyControlRiskReserveActionClearReserve         V2RiskEvaluationOutcomeDelinquencyControlRiskReserveAction = "clear_reserve"
	V2RiskEvaluationOutcomeDelinquencyControlRiskReserveActionMaintainCurrentState V2RiskEvaluationOutcomeDelinquencyControlRiskReserveAction = "maintain_current_state"
	V2RiskEvaluationOutcomeDelinquencyControlRiskReserveActionUpdateReserve        V2RiskEvaluationOutcomeDelinquencyControlRiskReserveAction = "update_reserve"
)

// How reserved funds are released.
type V2RiskEvaluationOutcomeDelinquencyControlRiskReserveTargetReleaseType string

// List of values that V2RiskEvaluationOutcomeDelinquencyControlRiskReserveTargetReleaseType can take
const (
	V2RiskEvaluationOutcomeDelinquencyControlRiskReserveTargetReleaseTypeRolling V2RiskEvaluationOutcomeDelinquencyControlRiskReserveTargetReleaseType = "rolling"
)

// The kind of control recommended to manage risk.
type V2RiskEvaluationOutcomeDelinquencyControlType string

// List of values that V2RiskEvaluationOutcomeDelinquencyControlType can take
const (
	V2RiskEvaluationOutcomeDelinquencyControlTypeRiskReserve V2RiskEvaluationOutcomeDelinquencyControlType = "risk_reserve"
)

// How the recommendation prioritizes protection from losses relative to revenue.
type V2RiskEvaluationOutcomeDelinquencyRiskAppetiteMode string

// List of values that V2RiskEvaluationOutcomeDelinquencyRiskAppetiteMode can take
const (
	V2RiskEvaluationOutcomeDelinquencyRiskAppetiteModeBalanced           V2RiskEvaluationOutcomeDelinquencyRiskAppetiteMode = "balanced"
	V2RiskEvaluationOutcomeDelinquencyRiskAppetiteModeMaximizeProtection V2RiskEvaluationOutcomeDelinquencyRiskAppetiteMode = "maximize_protection"
	V2RiskEvaluationOutcomeDelinquencyRiskAppetiteModeMaximizeRevenue    V2RiskEvaluationOutcomeDelinquencyRiskAppetiteMode = "maximize_revenue"
)

// Assessed risk level.
type V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevel string

// List of values that V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevel can take
const (
	V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevelElevated V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevel = "elevated"
	V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevelHighest  V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevel = "highest"
	V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevelLow      V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevel = "low"
	V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevelNormal   V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevel = "normal"
	V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevelUnknown  V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevel = "unknown"
)

// How the outcome was generated.
type V2RiskEvaluationOutcomeOrigin string

// List of values that V2RiskEvaluationOutcomeOrigin can take
const (
	V2RiskEvaluationOutcomeOriginRuleOutcome          V2RiskEvaluationOutcomeOrigin = "rule_outcome"
	V2RiskEvaluationOutcomeOriginStripeRecommendation V2RiskEvaluationOutcomeOrigin = "stripe_recommendation"
)

// Scenario evaluated; determines which type-specific payload is provided.
type V2RiskEvaluationOutcomeType string

// List of values that V2RiskEvaluationOutcomeType can take
const (
	V2RiskEvaluationOutcomeTypeDelinquency V2RiskEvaluationOutcomeType = "delinquency"
)

// Inline subject data for an evaluation without a Stripe entity.
type V2RiskEvaluationOutcomeAccountDetailsData struct {
	// Contact email supplied for the evaluation, or null when unavailable.
	ContactEmail string `json:"contact_email,omitempty"`
}

// Subject of the evaluation. Exactly one of account, customer, or data is provided.
type V2RiskEvaluationOutcomeAccountDetails struct {
	// Connected account evaluated.
	Account string `json:"account,omitempty"`
	// Customer evaluated.
	Customer string `json:"customer,omitempty"`
	// Inline subject data for an evaluation without a Stripe entity.
	Data *V2RiskEvaluationOutcomeAccountDetailsData `json:"data,omitempty"`
}

// Expected profitability, revenue, and losses if the current reserve configuration remains unchanged.
// Provides a comparison for projected_outcome. Null when not reliably quantifiable.
type V2RiskEvaluationOutcomeDelinquencyBusinessObjectiveBaselineOutcome struct {
	// Expected financial loss from delinquency under this reserve configuration, expressed as a negative amount.
	Loss Amount `json:"loss"`
	// Expected financial result after accounting for revenue and losses. Can be negative.
	Profitability Amount `json:"profitability"`
	// Expected revenue under this reserve configuration.
	Revenue Amount `json:"revenue"`
}

// Expected profitability, revenue, and losses if the recommended controls are applied. Equals
// baseline_outcome when the recommendation maintains the current state. Null when not reliably quantifiable.
type V2RiskEvaluationOutcomeDelinquencyBusinessObjectiveProjectedOutcome struct {
	// Expected financial loss from delinquency under this reserve configuration, expressed as a negative amount.
	Loss Amount `json:"loss"`
	// Expected financial result after accounting for revenue and losses. Can be negative.
	Profitability Amount `json:"profitability"`
	// Expected revenue under this reserve configuration.
	Revenue Amount `json:"revenue"`
}

// The objective of the recommendation, with expected financial outcomes for the recommended
// controls and the current configuration when those outcomes can be quantified.
type V2RiskEvaluationOutcomeDelinquencyBusinessObjective struct {
	// Expected profitability, revenue, and losses if the current reserve configuration remains unchanged.
	// Provides a comparison for projected_outcome. Null when not reliably quantifiable.
	BaselineOutcome *V2RiskEvaluationOutcomeDelinquencyBusinessObjectiveBaselineOutcome `json:"baseline_outcome,omitempty"`
	// Expected profitability, revenue, and losses if the recommended controls are applied. Equals
	// baseline_outcome when the recommendation maintains the current state. Null when not reliably quantifiable.
	ProjectedOutcome *V2RiskEvaluationOutcomeDelinquencyBusinessObjectiveProjectedOutcome `json:"projected_outcome,omitempty"`
	// The financial objective the recommended controls are intended to improve.
	Type V2RiskEvaluationOutcomeDelinquencyBusinessObjectiveType `json:"type"`
}

// Recommended reserve configuration, specifying the percentage of funds to hold, when to release
// them, and whether to include available funds. Populated for update_reserve; null for
// maintain_current_state and clear_reserve.
type V2RiskEvaluationOutcomeDelinquencyControlRiskReserveTarget struct {
	// Integer percentage of funds to reserve.
	Percent int64 `json:"percent"`
	// Integer number of days after which reserved funds are released.
	ReleaseAfterDays int64 `json:"release_after_days"`
	// How reserved funds are released.
	ReleaseType V2RiskEvaluationOutcomeDelinquencyControlRiskReserveTargetReleaseType `json:"release_type"`
	// Whether to reserve already available funds as well as future funds.
	SweepAvailableFunds bool `json:"sweep_available_funds"`
}

// The recommended action for funds held in reserve, including a target configuration when
// updating reserves. Populated when type is risk_reserve; null otherwise.
type V2RiskEvaluationOutcomeDelinquencyControlRiskReserve struct {
	// Whether to update, maintain, or clear reserves.
	Action V2RiskEvaluationOutcomeDelinquencyControlRiskReserveAction `json:"action"`
	// Recommended reserve configuration, specifying the percentage of funds to hold, when to release
	// them, and whether to include available funds. Populated for update_reserve; null for
	// maintain_current_state and clear_reserve.
	Target *V2RiskEvaluationOutcomeDelinquencyControlRiskReserveTarget `json:"target,omitempty"`
}

// Risk controls recommended to manage the account's delinquency risk.
type V2RiskEvaluationOutcomeDelinquencyControl struct {
	// The recommended action for funds held in reserve, including a target configuration when
	// updating reserves. Populated when type is risk_reserve; null otherwise.
	RiskReserve *V2RiskEvaluationOutcomeDelinquencyControlRiskReserve `json:"risk_reserve,omitempty"`
	// The kind of control recommended to manage risk.
	Type V2RiskEvaluationOutcomeDelinquencyControlType `json:"type"`
}

// The balance between loss protection and revenue used to generate this recommendation.
type V2RiskEvaluationOutcomeDelinquencyRiskAppetite struct {
	// How the recommendation prioritizes protection from losses relative to revenue.
	Mode V2RiskEvaluationOutcomeDelinquencyRiskAppetiteMode `json:"mode"`
}

// The account's assessed probability and level of delinquency risk. Null when unavailable.
type V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquency struct {
	// Time the signal was evaluated.
	EvaluatedAt time.Time `json:"evaluated_at"`
	// Probability of delinquency, represented as a decimal string. Null when the risk level is
	// unknown or the account signal does not include a numeric score.
	Probability decimal.Decimal `json:"probability,omitempty"`
	// Assessed risk level.
	RiskLevel V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquencyRiskLevel `json:"risk_level"`
	// ID of the supporting account signal.
	Signal string `json:"signal"`
}

// The account's assessed financial exposure if it becomes delinquent. Null when unavailable.
type V2RiskEvaluationOutcomeDelinquencySupportingSignalsPaymentDelinquencyExposure struct {
	// Time the signal was evaluated.
	EvaluatedAt time.Time `json:"evaluated_at"`
	// Exposure if the account becomes delinquent.
	ExposureAmount Amount `json:"exposure_amount"`
	// ID of the supporting account signal.
	Signal string `json:"signal"`
}

// Signals that informed this evaluation, with their values at evaluation time. Null when unavailable.
type V2RiskEvaluationOutcomeDelinquencySupportingSignals struct {
	// The account's assessed probability and level of delinquency risk. Null when unavailable.
	MerchantDelinquency *V2RiskEvaluationOutcomeDelinquencySupportingSignalsMerchantDelinquency `json:"merchant_delinquency,omitempty"`
	// The account's assessed financial exposure if it becomes delinquent. Null when unavailable.
	PaymentDelinquencyExposure *V2RiskEvaluationOutcomeDelinquencySupportingSignalsPaymentDelinquencyExposure `json:"payment_delinquency_exposure,omitempty"`
}

// Recommended controls, expected financial impact, and supporting signals for a delinquency
// evaluation. Populated when type is delinquency; null otherwise.
type V2RiskEvaluationOutcomeDelinquency struct {
	// The objective of the recommendation, with expected financial outcomes for the recommended
	// controls and the current configuration when those outcomes can be quantified.
	BusinessObjective *V2RiskEvaluationOutcomeDelinquencyBusinessObjective `json:"business_objective"`
	// Risk controls recommended to manage the account's delinquency risk.
	Controls []*V2RiskEvaluationOutcomeDelinquencyControl `json:"controls"`
	// The balance between loss protection and revenue used to generate this recommendation.
	RiskAppetite *V2RiskEvaluationOutcomeDelinquencyRiskAppetite `json:"risk_appetite"`
	// Signals that informed this evaluation, with their values at evaluation time. Null when unavailable.
	SupportingSignals *V2RiskEvaluationOutcomeDelinquencySupportingSignals `json:"supporting_signals,omitempty"`
}

// An immutable, point-in-time risk evaluation result that captures recommended controls and
// their expected impact. The type identifies the scenario evaluated and its corresponding details.
// Unavailable optional values are null.
type V2RiskEvaluationOutcome struct {
	APIResource
	// Subject of the evaluation. Exactly one of account, customer, or data is provided.
	AccountDetails *V2RiskEvaluationOutcomeAccountDetails `json:"account_details"`
	// Timestamp when the evaluation was created.
	Created time.Time `json:"created"`
	// Recommended controls, expected financial impact, and supporting signals for a delinquency
	// evaluation. Populated when type is delinquency; null otherwise.
	Delinquency *V2RiskEvaluationOutcomeDelinquency `json:"delinquency,omitempty"`
	// Unique identifier for the evaluation outcome, prefixed with evalout_.
	ID string `json:"id"`
	// Whether the evaluation outcome is in live mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// How the outcome was generated.
	Origin V2RiskEvaluationOutcomeOrigin `json:"origin"`
	// Scenario evaluated; determines which type-specific payload is provided.
	Type V2RiskEvaluationOutcomeType `json:"type"`
}
