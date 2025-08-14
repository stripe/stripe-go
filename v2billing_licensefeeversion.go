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
type V2BillingLicenseFeeVersionTieringMode string

// List of values that V2BillingLicenseFeeVersionTieringMode can take
const (
	V2BillingLicenseFeeVersionTieringModeGraduated V2BillingLicenseFeeVersionTieringMode = "graduated"
	V2BillingLicenseFeeVersionTieringModeVolume    V2BillingLicenseFeeVersionTieringMode = "volume"
)

// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
type V2BillingLicenseFeeVersionTierUpToInf string

// List of values that V2BillingLicenseFeeVersionTierUpToInf can take
const (
	V2BillingLicenseFeeVersionTierUpToInfInf V2BillingLicenseFeeVersionTierUpToInf = "inf"
)

// After division, round the result up or down.
type V2BillingLicenseFeeVersionTransformQuantityRound string

// List of values that V2BillingLicenseFeeVersionTransformQuantityRound can take
const (
	V2BillingLicenseFeeVersionTransformQuantityRoundDown V2BillingLicenseFeeVersionTransformQuantityRound = "down"
	V2BillingLicenseFeeVersionTransformQuantityRoundUp   V2BillingLicenseFeeVersionTransformQuantityRound = "up"
)

// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
type V2BillingLicenseFeeVersionTier struct {
	// Price for the entire tier, represented as a decimal string in minor currency units with at most 12 decimal places.
	FlatAmount string `json:"flat_amount"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units with at
	// most 12 decimal places.
	UnitAmount string `json:"unit_amount"`
	// Up to and including this quantity will be contained in the tier. Only one of `up_to_decimal` and `up_to_inf` may
	// be set.
	UpToDecimal string `json:"up_to_decimal"`
	// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
	UpToInf V2BillingLicenseFeeVersionTierUpToInf `json:"up_to_inf"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed.
type V2BillingLicenseFeeVersionTransformQuantity struct {
	// Divide usage by this number.
	DivideBy int64 `json:"divide_by"`
	// After division, round the result up or down.
	Round V2BillingLicenseFeeVersionTransformQuantityRound `json:"round"`
}
type V2BillingLicenseFeeVersion struct {
	APIResource
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// The ID of the LicenseFeeVersion.
	ID string `json:"id"`
	// The ID of the parent LicenseFee.
	LicenseFeeID string `json:"license_fee_id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Defines whether the tiering price should be graduated or volume-based. In volume-based tiering, the maximum
	// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
	// grows into new tiers. Can only be set if `tiers` is set.
	TieringMode V2BillingLicenseFeeVersionTieringMode `json:"tiering_mode"`
	// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
	Tiers []*V2BillingLicenseFeeVersionTier `json:"tiers"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed.
	TransformQuantity *V2BillingLicenseFeeVersionTransformQuantity `json:"transform_quantity"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units with at most 12 decimal
	// places. Cannot be set if `tiers` is provided.
	UnitAmount string `json:"unit_amount"`
}
