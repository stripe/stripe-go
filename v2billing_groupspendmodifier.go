//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of the spend modifier amount.
type V2BillingGroupSpendModifierMaxBillingPeriodSpendAmountType string

// List of values that V2BillingGroupSpendModifierMaxBillingPeriodSpendAmountType can take
const (
	V2BillingGroupSpendModifierMaxBillingPeriodSpendAmountTypeCustomPricingUnit V2BillingGroupSpendModifierMaxBillingPeriodSpendAmountType = "custom_pricing_unit"
)

// The type of the spend modifier.
type V2BillingGroupSpendModifierType string

// List of values that V2BillingGroupSpendModifierType can take
const (
	V2BillingGroupSpendModifierTypeMaxBillingPeriodSpend V2BillingGroupSpendModifierType = "max_billing_period_spend"
)

// The custom pricing unit amount. Present if type is `custom_pricing_unit`.
type V2BillingGroupSpendModifierMaxBillingPeriodSpendAmountCustomPricingUnit struct {
	// The ID of the custom pricing unit.
	ID string `json:"id"`
	// The decimal value of custom pricing units, represented as a string.
	Value string `json:"value"`
}

// The maximum invoice spend, measured in custom pricing units.
type V2BillingGroupSpendModifierMaxBillingPeriodSpendAmount struct {
	// The custom pricing unit amount. Present if type is `custom_pricing_unit`.
	CustomPricingUnit *V2BillingGroupSpendModifierMaxBillingPeriodSpendAmountCustomPricingUnit `json:"custom_pricing_unit,omitempty"`
	// The type of the spend modifier amount.
	Type V2BillingGroupSpendModifierMaxBillingPeriodSpendAmountType `json:"type"`
}

// Max invoice spend details. Present if type is `max_billing_period_spend`.
type V2BillingGroupSpendModifierMaxBillingPeriodSpend struct {
	// The billing alert associated with the maximum spend threshold.
	Alert string `json:"alert"`
	// The maximum invoice spend, measured in custom pricing units.
	Amount *V2BillingGroupSpendModifierMaxBillingPeriodSpendAmount `json:"amount"`
	// The timestamp from which this spend modifier is effective.
	EffectiveFrom time.Time `json:"effective_from"`
	// The timestamp until which this spend modifier is effective. If not set, the modifier is effective indefinitely.
	EffectiveUntil time.Time `json:"effective_until,omitempty"`
}

// A spend modifier changes how spend is computed for a billing group.
type V2BillingGroupSpendModifier struct {
	APIResource
	// The ID of the billing group this spend modifier is associated with.
	BillingGroup string `json:"billing_group"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Max invoice spend details. Present if type is `max_billing_period_spend`.
	MaxBillingPeriodSpend *V2BillingGroupSpendModifierMaxBillingPeriodSpend `json:"max_billing_period_spend,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The type of the spend modifier.
	Type V2BillingGroupSpendModifierType `json:"type"`
}
