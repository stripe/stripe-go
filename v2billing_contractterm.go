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

// Whether to include or exclude the selected recurring or scheduled_fee terms.
type V2BillingContractTermPricingOverrideCriterionType string

// List of values that V2BillingContractTermPricingOverrideCriterionType can take
const (
	V2BillingContractTermPricingOverrideCriterionTypeExclude V2BillingContractTermPricingOverrideCriterionType = "exclude"
	V2BillingContractTermPricingOverrideCriterionTypeInclude V2BillingContractTermPricingOverrideCriterionType = "include"
)

// Whether the replacement pricing is per unit or tiered.
type V2BillingContractTermPricingOverrideOverwriteBillingScheme string

// List of values that V2BillingContractTermPricingOverrideOverwriteBillingScheme can take
const (
	V2BillingContractTermPricingOverrideOverwriteBillingSchemePerUnit V2BillingContractTermPricingOverrideOverwriteBillingScheme = "per_unit"
	V2BillingContractTermPricingOverrideOverwriteBillingSchemeTiered  V2BillingContractTermPricingOverrideOverwriteBillingScheme = "tiered"
)

// How the pricing tiers are applied.
type V2BillingContractTermPricingOverrideOverwriteTiersMode string

// List of values that V2BillingContractTermPricingOverrideOverwriteTiersMode can take
const (
	V2BillingContractTermPricingOverrideOverwriteTiersModeGraduated V2BillingContractTermPricingOverrideOverwriteTiersMode = "graduated"
	V2BillingContractTermPricingOverrideOverwriteTiersModeVolume    V2BillingContractTermPricingOverrideOverwriteTiersMode = "volume"
)

// The pricing operation.
type V2BillingContractTermPricingOverrideType string

// List of values that V2BillingContractTermPricingOverrideType can take
const (
	V2BillingContractTermPricingOverrideTypeMultiply  V2BillingContractTermPricingOverrideType = "multiply"
	V2BillingContractTermPricingOverrideTypeOverwrite V2BillingContractTermPricingOverrideType = "overwrite"
)

// The type of billing schedule.
type V2BillingContractTermScheduledFeeBillScheduleType string

// List of values that V2BillingContractTermScheduledFeeBillScheduleType can take
const (
	V2BillingContractTermScheduledFeeBillScheduleTypeCustom V2BillingContractTermScheduledFeeBillScheduleType = "custom"
)

// The type of this term.
type V2BillingContractTermType string

// List of values that V2BillingContractTermType can take
const (
	V2BillingContractTermTypePricingOverride V2BillingContractTermType = "pricing_override"
	V2BillingContractTermTypeRecurringPrice  V2BillingContractTermType = "recurring_price"
	V2BillingContractTermTypeScheduledFee    V2BillingContractTermType = "scheduled_fee"
)

// Recurring or scheduled_fee terms selected by this override. An override without criteria applies to all recurring and scheduled_fee terms.
type V2BillingContractTermPricingOverrideCriterion struct {
	// Resolved IDs of recurring or scheduled_fee terms. Override terms are not valid targets.
	Terms []string `json:"terms,omitempty"`
	// Whether to include or exclude the selected recurring or scheduled_fee terms.
	Type V2BillingContractTermPricingOverrideCriterionType `json:"type"`
}

// The exclusive end of the validity window. Null for an open end.
type V2BillingContractTermPricingOverrideEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// Present when type is multiply.
type V2BillingContractTermPricingOverrideMultiply struct {
	// The multiplier, for example 0.85 for a 15% discount.
	Factor decimal.Decimal `json:"factor"`
}

// Replacement pricing tiers.
type V2BillingContractTermPricingOverrideOverwriteTier struct {
	// Flat amount in minor currency units.
	FlatAmountDecimal decimal.Decimal `json:"flat_amount_decimal,omitempty"`
	// Per-unit amount in minor currency units.
	UnitAmountDecimal decimal.Decimal `json:"unit_amount_decimal,omitempty"`
	// Inclusive whole-number upper bound. Omitted for an unbounded final tier.
	UpTo int64 `json:"up_to,string,omitempty"`
}

// Present when type is overwrite.
type V2BillingContractTermPricingOverrideOverwrite struct {
	// Whether the replacement pricing is per unit or tiered.
	BillingScheme V2BillingContractTermPricingOverrideOverwriteBillingScheme `json:"billing_scheme,omitempty"`
	// Replacement pricing tiers.
	Tiers []*V2BillingContractTermPricingOverrideOverwriteTier `json:"tiers,omitempty"`
	// How the pricing tiers are applied.
	TiersMode V2BillingContractTermPricingOverrideOverwriteTiersMode `json:"tiers_mode,omitempty"`
	// The replacement per-unit amount in minor currency units.
	UnitAmountDecimal decimal.Decimal `json:"unit_amount_decimal,omitempty"`
}

// The inclusive start of the override's validity window.
type V2BillingContractTermPricingOverrideStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// Present when type is pricing_override.
type V2BillingContractTermPricingOverride struct {
	// Recurring or scheduled_fee terms selected by this override. An override without criteria applies to all recurring and scheduled_fee terms.
	Criteria []*V2BillingContractTermPricingOverrideCriterion `json:"criteria"`
	// The exclusive end of the validity window. Null for an open end.
	EndsAt *V2BillingContractTermPricingOverrideEndsAt `json:"ends_at,omitempty"`
	// Present when type is multiply.
	Multiply *V2BillingContractTermPricingOverrideMultiply `json:"multiply,omitempty"`
	// Present when type is overwrite.
	Overwrite *V2BillingContractTermPricingOverrideOverwrite `json:"overwrite,omitempty"`
	// Lower numbers win. Equal priorities resolve by display order; overrides do not stack.
	Priority int64 `json:"priority"`
	// The inclusive start of the override's validity window.
	StartsAt *V2BillingContractTermPricingOverrideStartsAt `json:"starts_at"`
	// The pricing operation.
	Type V2BillingContractTermPricingOverrideType `json:"type"`
}

// The exclusive end of the term. Null for an open end.
type V2BillingContractTermRecurringPriceEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// The inclusive start of the term.
type V2BillingContractTermRecurringPriceStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// Present when type is recurring_price.
type V2BillingContractTermRecurringPrice struct {
	// The exclusive end of the term. Null for an open end.
	EndsAt *V2BillingContractTermRecurringPriceEndsAt `json:"ends_at,omitempty"`
	// The V1 recurring Price.
	Price string `json:"price"`
	// Current licensed quantity. Null for metered Prices; history is available separately.
	Quantity decimal.Decimal `json:"quantity,omitempty"`
	// The inclusive start of the term.
	StartsAt *V2BillingContractTermRecurringPriceStartsAt `json:"starts_at"`
}

// The resolved bill date.
type V2BillingContractTermScheduledFeeBillScheduleCustomBillAt struct {
	// The timestamp at which the fee will be billed.
	Timestamp time.Time `json:"timestamp"`
}

// Scheduled installments and billed history.
type V2BillingContractTermScheduledFeeBillScheduleCustom struct {
	// The resolved bill date.
	BillAt *V2BillingContractTermScheduledFeeBillScheduleCustomBillAt `json:"bill_at"`
	// Resolved installment quantity. Always present; defaults to 1 when omitted on input.
	Quantity decimal.Decimal `json:"quantity"`
	// Unit price in minor units of the contract's currency, before multiplying by quantity.
	UnitAmountDecimal decimal.Decimal `json:"unit_amount_decimal"`
}

// Explicit bill dates and amounts, including billed history.
type V2BillingContractTermScheduledFeeBillSchedule struct {
	// Scheduled installments and billed history.
	Custom []*V2BillingContractTermScheduledFeeBillScheduleCustom `json:"custom"`
	// The type of billing schedule.
	Type V2BillingContractTermScheduledFeeBillScheduleType `json:"type"`
}

// The exclusive end of the service period.
type V2BillingContractTermScheduledFeeServicePeriodEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// The inclusive start of the service period.
type V2BillingContractTermScheduledFeeServicePeriodStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// Informational invoice and revenue-recognition period with timestamp objects; does not affect billing or proration.
type V2BillingContractTermScheduledFeeServicePeriod struct {
	// The exclusive end of the service period.
	EndsAt *V2BillingContractTermScheduledFeeServicePeriodEndsAt `json:"ends_at"`
	// The inclusive start of the service period.
	StartsAt *V2BillingContractTermScheduledFeeServicePeriodStartsAt `json:"starts_at"`
}

// Present when type is scheduled_fee.
type V2BillingContractTermScheduledFee struct {
	// Explicit bill dates and amounts, including billed history.
	BillSchedule *V2BillingContractTermScheduledFeeBillSchedule `json:"bill_schedule"`
	// The V1 Product being billed.
	Product string `json:"product"`
	// Informational invoice and revenue-recognition period with timestamp objects; does not affect billing or proration.
	ServicePeriod *V2BillingContractTermScheduledFeeServicePeriod `json:"service_period,omitempty"`
}

// A commercial term. Display order is independent of billing evaluation order.
type V2BillingContractTerm struct {
	APIResource
	// The stable ID of the term.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The lookup key, unique across all term types on this contract.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of key-value pairs attached to the term.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Present when type is pricing_override.
	PricingOverride *V2BillingContractTermPricingOverride `json:"pricing_override,omitempty"`
	// Present when type is recurring_price.
	RecurringPrice *V2BillingContractTermRecurringPrice `json:"recurring_price,omitempty"`
	// Present when type is scheduled_fee.
	ScheduledFee *V2BillingContractTermScheduledFee `json:"scheduled_fee,omitempty"`
	// The type of this term.
	Type V2BillingContractTermType `json:"type"`
}
