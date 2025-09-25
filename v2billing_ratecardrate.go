//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Defines whether the tiering price should be graduated or volume-based. In volume-based tiering, the maximum
// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
// grows into new tiers. Can only be set if `tiers` is set.
type V2BillingRateCardRateTieringMode string

// List of values that V2BillingRateCardRateTieringMode can take
const (
	V2BillingRateCardRateTieringModeGraduated V2BillingRateCardRateTieringMode = "graduated"
	V2BillingRateCardRateTieringModeVolume    V2BillingRateCardRateTieringMode = "volume"
)

// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
type V2BillingRateCardRateTierUpToInf string

// List of values that V2BillingRateCardRateTierUpToInf can take
const (
	V2BillingRateCardRateTierUpToInfInf V2BillingRateCardRateTierUpToInf = "inf"
)

// After division, round the result up or down.
type V2BillingRateCardRateTransformQuantityRound string

// List of values that V2BillingRateCardRateTransformQuantityRound can take
const (
	V2BillingRateCardRateTransformQuantityRoundDown V2BillingRateCardRateTransformQuantityRound = "down"
	V2BillingRateCardRateTransformQuantityRoundUp   V2BillingRateCardRateTransformQuantityRound = "up"
)

// The custom pricing unit that this rate binds to.
type V2BillingRateCardRateCustomPricingUnitAmount struct {
	// The Custom Pricing Unit object.
	CustomPricingUnitDetails *V2BillingCustomPricingUnit `json:"custom_pricing_unit_details,omitempty"`
	// The id of the custom pricing unit.
	ID string `json:"id"`
	// The unit value for the custom pricing unit, as a string.
	Value string `json:"value"`
}

// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
type V2BillingRateCardRateTier struct {
	// Price for the entire tier, represented as a decimal string in minor currency units with at most 12 decimal places.
	FlatAmount string `json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units with at
	// most 12 decimal places.
	UnitAmount string `json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier. Only one of `up_to_decimal` and `up_to_inf` may
	// be set.
	UpToDecimal string `json:"up_to_decimal,omitempty"`
	// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
	UpToInf V2BillingRateCardRateTierUpToInf `json:"up_to_inf,omitempty"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed.
type V2BillingRateCardRateTransformQuantity struct {
	// Divide usage by this number.
	DivideBy int64 `json:"divide_by"`
	// After division, round the result up or down.
	Round V2BillingRateCardRateTransformQuantityRound `json:"round"`
}
type V2BillingRateCardRate struct {
	APIResource
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// The custom pricing unit that this rate binds to.
	CustomPricingUnitAmount *V2BillingRateCardRateCustomPricingUnitAmount `json:"custom_pricing_unit_amount,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// A Metered Item represents a billable item whose pricing is based on usage, measured by a meter. You can use rate cards
	// to specify the pricing and create subscriptions to these items.
	MeteredItem *V2BillingMeteredItem `json:"metered_item"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the Rate Card it belongs to.
	RateCard string `json:"rate_card"`
	// The ID of the Rate Card Version it was created on.
	RateCardVersion string `json:"rate_card_version"`
	// Defines whether the tiering price should be graduated or volume-based. In volume-based tiering, the maximum
	// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
	// grows into new tiers. Can only be set if `tiers` is set.
	TieringMode V2BillingRateCardRateTieringMode `json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
	Tiers []*V2BillingRateCardRateTier `json:"tiers"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed.
	TransformQuantity *V2BillingRateCardRateTransformQuantity `json:"transform_quantity,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units with at most 12 decimal
	// places. Cannot be set if `tiers` is provided.
	UnitAmount string `json:"unit_amount,omitempty"`
}
