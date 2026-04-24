//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The interval for assessing service.
type V2BillingLicenseFeeServiceCycleInterval string

// List of values that V2BillingLicenseFeeServiceCycleInterval can take
const (
	V2BillingLicenseFeeServiceCycleIntervalDay   V2BillingLicenseFeeServiceCycleInterval = "day"
	V2BillingLicenseFeeServiceCycleIntervalMonth V2BillingLicenseFeeServiceCycleInterval = "month"
	V2BillingLicenseFeeServiceCycleIntervalWeek  V2BillingLicenseFeeServiceCycleInterval = "week"
	V2BillingLicenseFeeServiceCycleIntervalYear  V2BillingLicenseFeeServiceCycleInterval = "year"
)

// The interval for assessing service.
type V2BillingLicenseFeeServiceInterval string

// List of values that V2BillingLicenseFeeServiceInterval can take
const (
	V2BillingLicenseFeeServiceIntervalDay   V2BillingLicenseFeeServiceInterval = "day"
	V2BillingLicenseFeeServiceIntervalMonth V2BillingLicenseFeeServiceInterval = "month"
	V2BillingLicenseFeeServiceIntervalWeek  V2BillingLicenseFeeServiceInterval = "week"
	V2BillingLicenseFeeServiceIntervalYear  V2BillingLicenseFeeServiceInterval = "year"
)

// The tax behavior for Stripe Tax — whether the license fee price includes or excludes tax.
type V2BillingLicenseFeeTaxBehavior string

// List of values that V2BillingLicenseFeeTaxBehavior can take
const (
	V2BillingLicenseFeeTaxBehaviorExclusive V2BillingLicenseFeeTaxBehavior = "exclusive"
	V2BillingLicenseFeeTaxBehaviorInclusive V2BillingLicenseFeeTaxBehavior = "inclusive"
)

// Defines whether the tiering price is graduated or volume-based. In volume-based tiering, the maximum
// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
// grows into new tiers. Can only be set if `tiers` is set.
type V2BillingLicenseFeeTieringMode string

// List of values that V2BillingLicenseFeeTieringMode can take
const (
	V2BillingLicenseFeeTieringModeGraduated V2BillingLicenseFeeTieringMode = "graduated"
	V2BillingLicenseFeeTieringModeVolume    V2BillingLicenseFeeTieringMode = "volume"
)

// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
type V2BillingLicenseFeeTierUpToInf string

// List of values that V2BillingLicenseFeeTierUpToInf can take
const (
	V2BillingLicenseFeeTierUpToInfInf V2BillingLicenseFeeTierUpToInf = "inf"
)

// After division, round the result up or down.
type V2BillingLicenseFeeTransformQuantityRound string

// List of values that V2BillingLicenseFeeTransformQuantityRound can take
const (
	V2BillingLicenseFeeTransformQuantityRoundDown V2BillingLicenseFeeTransformQuantityRound = "down"
	V2BillingLicenseFeeTransformQuantityRoundUp   V2BillingLicenseFeeTransformQuantityRound = "up"
)

// The service cycle configuration for this License Fee.
type V2BillingLicenseFeeServiceCycle struct {
	// The interval for assessing service.
	Interval V2BillingLicenseFeeServiceCycleInterval `json:"interval"`
	// The length of the interval for assessing service. For example, set this to 3 and `interval` to `"month"`
	// to specify quarterly service.
	IntervalCount int64 `json:"interval_count"`
}

// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
type V2BillingLicenseFeeTier struct {
	// Price for the entire tier, represented as a decimal string in minor currency units with at most 12 decimal places.
	FlatAmount string `json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units with at
	// most 12 decimal places.
	UnitAmount string `json:"unit_amount,omitempty"`
	// Up to and including this quantity is contained in the tier. Only one of `up_to_decimal` and `up_to_inf` may
	// be set.
	UpToDecimal float64 `json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
	UpToInf V2BillingLicenseFeeTierUpToInf `json:"up_to_inf,omitempty"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed.
type V2BillingLicenseFeeTransformQuantity struct {
	// Divide usage by this number.
	DivideBy int64 `json:"divide_by,string"`
	// After division, round the result up or down.
	Round V2BillingLicenseFeeTransformQuantityRound `json:"round"`
}

// A License Fee represents a versioned recurring charge for a Licensed Item, typically used for seat-based or quantity-based
// pricing. Each License Fee defines the pricing structure (flat unit amount or tiered pricing) and service interval. After
// creating a License Fee, you can subscribe customers to it by creating a License Fee Subscription.
type V2BillingLicenseFee struct {
	APIResource
	// Whether this License Fee is active. Inactive License Fees cannot be used in new activations or be modified.
	Active bool `json:"active"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency Currency `json:"currency"`
	// A customer-facing name for the license fee.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The ID of the license fee's most recently created version.
	LatestVersion string `json:"latest_version"`
	// A Licensed Item represents a billable item whose pricing is based on license fees. You can use license fees
	// to specify the pricing and create subscriptions to these items.
	LicensedItem *V2BillingLicensedItem `json:"licensed_item"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The ID of the License Fee Version used by all subscriptions when no specific version is specified.
	LiveVersion string `json:"live_version"`
	// An internal key you can use to search for a particular License Fee. Maximum length of 200 characters.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The service cycle configuration for this License Fee.
	ServiceCycle *V2BillingLicenseFeeServiceCycle `json:"service_cycle"`
	// The interval for assessing service.
	ServiceInterval V2BillingLicenseFeeServiceInterval `json:"service_interval"`
	// The length of the interval for assessing service. For example, set this to 3 and `service_interval` to `"month"`
	// to specify quarterly service.
	ServiceIntervalCount int64 `json:"service_interval_count"`
	// The tax behavior for Stripe Tax — whether the license fee price includes or excludes tax.
	TaxBehavior V2BillingLicenseFeeTaxBehavior `json:"tax_behavior"`
	// Defines whether the tiering price is graduated or volume-based. In volume-based tiering, the maximum
	// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
	// grows into new tiers. Can only be set if `tiers` is set.
	TieringMode V2BillingLicenseFeeTieringMode `json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
	Tiers []*V2BillingLicenseFeeTier `json:"tiers"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed.
	TransformQuantity *V2BillingLicenseFeeTransformQuantity `json:"transform_quantity,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units with at most 12 decimal
	// places. Cannot be set if `tiers` is provided.
	UnitAmount string `json:"unit_amount,omitempty"`
}
