//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of the spend modifier amount.
type V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmountType string

// List of values that V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmountType can take
const (
	V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmountTypeCustomPricingUnit V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmountType = "custom_pricing_unit"
)

// The type of the spend modifier.
type V2BillingCadenceSpendModifierType string

// List of values that V2BillingCadenceSpendModifierType can take
const (
	V2BillingCadenceSpendModifierTypeMaxBillingPeriodSpend V2BillingCadenceSpendModifierType = "max_billing_period_spend"
)

// The custom pricing unit amount. Present if type is `custom_pricing_unit`.
type V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmountCustomPricingUnit struct {
	// The id of the custom pricing unit.
	ID string `json:"id,omitempty"`
	// The decimal value of custom pricing units, represented as a string.
	Value string `json:"value"`
}

// The maximum amount of invoice spend.
type V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmount struct {
	// The custom pricing unit amount. Present if type is `custom_pricing_unit`.
	CustomPricingUnit *V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmountCustomPricingUnit `json:"custom_pricing_unit,omitempty"`
	// The type of the spend modifier amount.
	Type V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmountType `json:"type"`
}

// The configuration for the overage rate for the custom pricing unit.
type V2BillingCadenceSpendModifierMaxBillingPeriodSpendCustomPricingUnitOverageRate struct {
	// ID of the custom pricing unit overage rate.
	ID string `json:"id"`
}

// Max invoice spend details. Present if type is `max_billing_period_spend`.
type V2BillingCadenceSpendModifierMaxBillingPeriodSpend struct {
	// The billing alert associated with the maximum spend threshold.
	Alert string `json:"alert"`
	// The maximum amount of invoice spend.
	Amount *V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmount `json:"amount"`
	// The configuration for the overage rate for the custom pricing unit.
	CustomPricingUnitOverageRate *V2BillingCadenceSpendModifierMaxBillingPeriodSpendCustomPricingUnitOverageRate `json:"custom_pricing_unit_overage_rate"`
	// The timestamp from which this spend modifier is effective.
	EffectiveFrom time.Time `json:"effective_from"`
	// The timestamp until which this spend modifier is effective. If not set, the modifier is effective indefinitely.
	EffectiveUntil time.Time `json:"effective_until,omitempty"`
}

// A Spend Modifier changes how spend is computed when billing for specific cadence.
type V2BillingCadenceSpendModifier struct {
	APIResource
	// The ID of the Billing Cadence this spend modifier is associated with.
	BillingCadence string `json:"billing_cadence"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Max invoice spend details. Present if type is `max_billing_period_spend`.
	MaxBillingPeriodSpend *V2BillingCadenceSpendModifierMaxBillingPeriodSpend `json:"max_billing_period_spend,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The type of the spend modifier.
	Type V2BillingCadenceSpendModifierType `json:"type"`
}
