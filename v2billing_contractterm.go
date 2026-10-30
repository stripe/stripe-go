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

// The type of billing schedule.
type V2BillingContractTermAdHocProductBillScheduleType string

// List of values that V2BillingContractTermAdHocProductBillScheduleType can take
const (
	V2BillingContractTermAdHocProductBillScheduleTypeCustom V2BillingContractTermAdHocProductBillScheduleType = "custom"
)

// Whether to include or exclude the selected recurring or ad_hoc_product terms.
type V2BillingContractTermPricingOverrideCriterionType string

// List of values that V2BillingContractTermPricingOverrideCriterionType can take
const (
	V2BillingContractTermPricingOverrideCriterionTypeExclude V2BillingContractTermPricingOverrideCriterionType = "exclude"
	V2BillingContractTermPricingOverrideCriterionTypeInclude V2BillingContractTermPricingOverrideCriterionType = "include"
)

// The pricing operation.
type V2BillingContractTermPricingOverrideType string

// List of values that V2BillingContractTermPricingOverrideType can take
const (
	V2BillingContractTermPricingOverrideTypeMultiply  V2BillingContractTermPricingOverrideType = "multiply"
	V2BillingContractTermPricingOverrideTypeOverwrite V2BillingContractTermPricingOverrideType = "overwrite"
)

// The type of this term.
type V2BillingContractTermType string

// List of values that V2BillingContractTermType can take
const (
	V2BillingContractTermTypeAdHocProduct    V2BillingContractTermType = "ad_hoc_product"
	V2BillingContractTermTypePricingOverride V2BillingContractTermType = "pricing_override"
	V2BillingContractTermTypeRecurringPrice  V2BillingContractTermType = "recurring_price"
)

// The resolved bill date.
type V2BillingContractTermAdHocProductBillScheduleCustomBillAt struct {
	// The timestamp at which the fee will be billed.
	Timestamp time.Time `json:"timestamp"`
}

// Scheduled installments and billed history.
type V2BillingContractTermAdHocProductBillScheduleCustom struct {
	// Amount in the contract's currency.
	Amount Amount `json:"amount"`
	// The resolved bill date.
	BillAt *V2BillingContractTermAdHocProductBillScheduleCustomBillAt `json:"bill_at"`
}

// Explicit bill dates and amounts, including billed history.
type V2BillingContractTermAdHocProductBillSchedule struct {
	// Scheduled installments and billed history.
	Custom []*V2BillingContractTermAdHocProductBillScheduleCustom `json:"custom"`
	// The type of billing schedule.
	Type V2BillingContractTermAdHocProductBillScheduleType `json:"type"`
}

// The exclusive end of the service period.
type V2BillingContractTermAdHocProductServicePeriodEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// The inclusive start of the service period.
type V2BillingContractTermAdHocProductServicePeriodStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// Informational invoice and revenue-recognition period with timestamp objects; does not affect billing or proration.
type V2BillingContractTermAdHocProductServicePeriod struct {
	// The exclusive end of the service period.
	EndsAt *V2BillingContractTermAdHocProductServicePeriodEndsAt `json:"ends_at"`
	// The inclusive start of the service period.
	StartsAt *V2BillingContractTermAdHocProductServicePeriodStartsAt `json:"starts_at"`
}

// Present when type is ad_hoc_product.
type V2BillingContractTermAdHocProduct struct {
	// Explicit bill dates and amounts, including billed history.
	BillSchedule *V2BillingContractTermAdHocProductBillSchedule `json:"bill_schedule"`
	// The V1 Product being billed.
	Product string `json:"product"`
	// Informational invoice and revenue-recognition period with timestamp objects; does not affect billing or proration.
	ServicePeriod *V2BillingContractTermAdHocProductServicePeriod `json:"service_period,omitempty"`
}

// Recurring or ad_hoc_product terms selected by this override. An override without criteria applies to all recurring and ad_hoc_product terms.
type V2BillingContractTermPricingOverrideCriterion struct {
	// Resolved IDs of recurring or ad_hoc_product terms. Override terms are not valid targets.
	Terms []string `json:"terms,omitempty"`
	// Whether to include or exclude the selected recurring or ad_hoc_product terms.
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

// Present when type is overwrite.
type V2BillingContractTermPricingOverrideOverwrite struct {
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
	// Recurring or ad_hoc_product terms selected by this override. An override without criteria applies to all recurring and ad_hoc_product terms.
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

// A commercial term. Display order is independent of billing evaluation order.
type V2BillingContractTerm struct {
	APIResource
	// Present when type is ad_hoc_product.
	AdHocProduct *V2BillingContractTermAdHocProduct `json:"ad_hoc_product,omitempty"`
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
	// The type of this term.
	Type V2BillingContractTermType `json:"type"`
}
