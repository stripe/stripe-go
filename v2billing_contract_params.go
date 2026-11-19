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

// List contracts.
type V2BillingContractListParams struct {
	Params `form:"*"`
	// Filter by customer id.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// The limit for the number of results per page.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Configuration for determining the billing cycle anchor by calendar fields.
type V2BillingContractBillingCycleAnchorConfigParams struct {
	// Day of month (1-31).
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// Hour of day in UTC (0-23).
	Hour *int64 `form:"hour" json:"hour,omitempty"`
	// Minute of hour (0-59).
	Minute *int64 `form:"minute" json:"minute,omitempty"`
	// Month of year (1-12).
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// Second of minute (0-59).
	Second *int64 `form:"second" json:"second,omitempty"`
}

// The billing cycle anchor for the contract. If not provided, defaults to the pricing line start time.
// It is only at the top-level of the contract with no option to override at the pricing line level.
type V2BillingContractBillingCycleAnchorParams struct {
	// Configuration for determining the billing cycle anchor by calendar fields.
	Config *V2BillingContractBillingCycleAnchorConfigParams `form:"config" json:"config,omitempty"`
	// A specific timestamp to use as the billing cycle anchor.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of billing cycle anchor.
	Type *string `form:"type" json:"type"`
}

// Tax calculation settings.
type V2BillingContractBillingSettingsBillSettingsDetailsCalculationTaxParams struct {
	// The type of tax calculation.
	Type *string `form:"type" json:"type"`
}

// Calculation settings.
type V2BillingContractBillingSettingsBillSettingsDetailsCalculationParams struct {
	// Tax calculation settings.
	Tax *V2BillingContractBillingSettingsBillSettingsDetailsCalculationTaxParams `form:"tax" json:"tax,omitempty"`
}

// How long the customer has to pay the invoice before it's past due.
type V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueParams struct {
	// The interval unit.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Invoice settings.
type V2BillingContractBillingSettingsBillSettingsDetailsInvoiceParams struct {
	// How long the customer has to pay the invoice before it's past due.
	TimeUntilDue *V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueParams `form:"time_until_due" json:"time_until_due,omitempty"`
}

// The bill settings details configures invoice and tax settings for the contract.
type V2BillingContractBillingSettingsBillSettingsDetailsParams struct {
	// The tax calculation settings to update.
	Calculation *V2BillingContractBillingSettingsBillSettingsDetailsCalculationParams `form:"calculation" json:"calculation,omitempty"`
	// The invoice settings to update.
	Invoice *V2BillingContractBillingSettingsBillSettingsDetailsInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// The billing profile details configures who is charged for the contract.
type V2BillingContractBillingSettingsBillingProfileDetailsParams struct {
	// The customer who pays for the contract invoice.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The default payment method to charge for the contract.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
}

// The collection settings details configures how payments are collected on the contract.
type V2BillingContractBillingSettingsCollectionSettingsDetailsParams struct {
	// How payment is collected for the contract. An omitted value leaves the
	// collection method unchanged.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// The payment method configuration.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
}

// The billing settings for the contract.
type V2BillingContractBillingSettingsParams struct {
	// The billing profile details configures who is charged for the contract.
	BillingProfileDetails *V2BillingContractBillingSettingsBillingProfileDetailsParams `form:"billing_profile_details" json:"billing_profile_details,omitempty"`
	// The bill settings to update (tax calculation type and/or invoice time until due).
	BillSettingsDetails *V2BillingContractBillingSettingsBillSettingsDetailsParams `form:"bill_settings_details" json:"bill_settings_details,omitempty"`
	// The collection settings details configures how payments are collected on the contract.
	CollectionSettingsDetails *V2BillingContractBillingSettingsCollectionSettingsDetailsParams `form:"collection_settings_details" json:"collection_settings_details,omitempty"`
}

// Term selectors. Each ID or lookup key resolves to one fixed term ID.
// Omitted or empty selects all charge-bearing terms. Override terms cannot be targets.
type V2BillingContractTermPricingOverrideCriterionTermParams struct {
	// Select this term by its stable ID.
	ID *string `form:"id" json:"id,omitempty"`
	// Select this term by its unique lookup key.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Whether the selector uses a term ID or lookup key.
	Type *string `form:"type" json:"type"`
}

// Optional criteria selecting recurring or scheduled_fee terms. If omitted, the override applies to all recurring and scheduled_fee terms.
type V2BillingContractTermPricingOverrideCriterionParams struct {
	// Term selectors. Each ID or lookup key resolves to one fixed term ID.
	// Omitted or empty selects all charge-bearing terms. Override terms cannot be targets.
	Terms []*V2BillingContractTermPricingOverrideCriterionTermParams `form:"terms" json:"terms,omitempty"`
	// Whether to include or exclude the selected recurring or scheduled_fee terms.
	Type *string `form:"type" json:"type"`
}

// The exclusive end. Omitted means open-ended.
type V2BillingContractTermPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is multiply.
type V2BillingContractTermPricingOverrideMultiplyParams struct {
	// The multiplier, for example 0.85 for a 15% discount.
	Factor *decimal.Decimal `form:"factor" json:"factor"`
}

// Required when type is overwrite.
type V2BillingContractTermPricingOverrideOverwriteParams struct {
	// The replacement per-unit amount in minor currency units.
	UnitAmountDecimal *decimal.Decimal `form:"unit_amount_decimal" json:"unit_amount_decimal,omitempty"`
}

// The inclusive start of the validity window.
type V2BillingContractTermPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is pricing_override.
type V2BillingContractTermPricingOverrideParams struct {
	// Optional criteria selecting recurring or scheduled_fee terms. If omitted, the override applies to all recurring and scheduled_fee terms.
	Criteria []*V2BillingContractTermPricingOverrideCriterionParams `form:"criteria" json:"criteria,omitempty"`
	// The exclusive end. Omitted means open-ended.
	EndsAt *V2BillingContractTermPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// Required when type is multiply.
	Multiply *V2BillingContractTermPricingOverrideMultiplyParams `form:"multiply" json:"multiply,omitempty"`
	// Required when type is overwrite.
	Overwrite *V2BillingContractTermPricingOverrideOverwriteParams `form:"overwrite" json:"overwrite,omitempty"`
	// Lower numbers win; defaults to 50.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
	// The inclusive start of the validity window.
	StartsAt *V2BillingContractTermPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at"`
	// The pricing operation. Exactly one matching pricing payload is required.
	Type *string `form:"type" json:"type"`
}

// The exclusive end. Omitted means open-ended.
type V2BillingContractTermRecurringPriceEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// The inclusive start of the term.
type V2BillingContractTermRecurringPriceStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is recurring_price.
type V2BillingContractTermRecurringPriceParams struct {
	// The exclusive end. Omitted means open-ended.
	EndsAt *V2BillingContractTermRecurringPriceEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// The V1 recurring Price.
	Price *string `form:"price" json:"price"`
	// Initial licensed quantity. Defaults to 1; must be omitted for metered Prices.
	Quantity *decimal.Decimal `form:"quantity" json:"quantity,omitempty"`
	// The inclusive start of the term.
	StartsAt *V2BillingContractTermRecurringPriceStartsAtParams `form:"starts_at" json:"starts_at"`
}

// A finite bill date. Resolved once per request, including on idempotent retries.
type V2BillingContractTermScheduledFeeBillScheduleCustomBillAtParams struct {
	// The timestamp at which the entry should be billed. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the bill_at.
	Type *string `form:"type" json:"type"`
}

// Between one and ten installments. Empty schedules are allowed only on update.
type V2BillingContractTermScheduledFeeBillScheduleCustomParams struct {
	// Non-negative amount with an explicit currency that must match the contract's currency.
	Amount *Amount `form:"amount" json:"amount"`
	// A finite bill date. Resolved once per request, including on idempotent retries.
	BillAt *V2BillingContractTermScheduledFeeBillScheduleCustomBillAtParams `form:"bill_at" json:"bill_at"`
}

// Explicit bill dates and amounts.
type V2BillingContractTermScheduledFeeBillScheduleParams struct {
	// Between one and ten installments. Empty schedules are allowed only on update.
	Custom []*V2BillingContractTermScheduledFeeBillScheduleCustomParams `form:"custom" json:"custom,omitempty"`
	// The type of schedule.
	Type *string `form:"type" json:"type"`
}

// The finite exclusive end; must be after starts_at. Never is not supported.
type V2BillingContractTermScheduledFeeServicePeriodEndsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// The finite inclusive start of the service period.
type V2BillingContractTermScheduledFeeServicePeriodStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Informational service period, independent of bill dates.
type V2BillingContractTermScheduledFeeServicePeriodParams struct {
	// The finite exclusive end; must be after starts_at. Never is not supported.
	EndsAt *V2BillingContractTermScheduledFeeServicePeriodEndsAtParams `form:"ends_at" json:"ends_at"`
	// The finite inclusive start of the service period.
	StartsAt *V2BillingContractTermScheduledFeeServicePeriodStartsAtParams `form:"starts_at" json:"starts_at"`
}

// Required when type is scheduled_fee.
type V2BillingContractTermScheduledFeeParams struct {
	// Explicit bill dates and amounts.
	BillSchedule *V2BillingContractTermScheduledFeeBillScheduleParams `form:"bill_schedule" json:"bill_schedule"`
	// The V1 Product being billed.
	Product *string `form:"product" json:"product"`
	// Informational service period, independent of bill dates.
	ServicePeriod *V2BillingContractTermScheduledFeeServicePeriodParams `form:"service_period" json:"service_period,omitempty"`
}

// Terms to create with the contract.
type V2BillingContractTermParams struct {
	// Lookup key unique across all term types on the contract.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs attached to the term.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Required when type is pricing_override.
	PricingOverride *V2BillingContractTermPricingOverrideParams `form:"pricing_override" json:"pricing_override,omitempty"`
	// Required when type is recurring_price.
	RecurringPrice *V2BillingContractTermRecurringPriceParams `form:"recurring_price" json:"recurring_price,omitempty"`
	// Required when type is scheduled_fee.
	ScheduledFee *V2BillingContractTermScheduledFeeParams `form:"scheduled_fee" json:"scheduled_fee,omitempty"`
	// The term type. Production types require exactly one matching payload.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractTermParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Create a draft contract.
type V2BillingContractParams struct {
	Params `form:"*"`
	// Apply all term actions atomically.
	Actions []*V2BillingContractActionParams `form:"actions" json:"actions,omitempty"`
	// The billing cycle anchor for the contract. If not provided, defaults to the pricing line start time.
	// It is only at the top-level of the contract with no option to override at the pricing line level.
	BillingCycleAnchor *V2BillingContractBillingCycleAnchorParams `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// The billing settings to update on the contract.
	BillingSettings *V2BillingContractBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// The unique user-provided contract number. This can only be changed while the contract is in draft.
	ContractNumber *string `form:"contract_number" json:"contract_number,omitempty"`
	// Currency of the contract.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// Terms to create with the contract.
	Terms []*V2BillingContractTermParams `form:"terms" json:"terms,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Term selectors. Each ID or lookup key resolves to one fixed term ID.
// Omitted or empty selects all charge-bearing terms. Override terms cannot be targets.
type V2BillingContractActionAddTermPricingOverrideCriterionTermParams struct {
	// Select this term by its stable ID.
	ID *string `form:"id" json:"id,omitempty"`
	// Select this term by its unique lookup key.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Whether the selector uses a term ID or lookup key.
	Type *string `form:"type" json:"type"`
}

// Optional criteria selecting recurring or scheduled_fee terms. If omitted, the override applies to all recurring and scheduled_fee terms.
type V2BillingContractActionAddTermPricingOverrideCriterionParams struct {
	// Term selectors. Each ID or lookup key resolves to one fixed term ID.
	// Omitted or empty selects all charge-bearing terms. Override terms cannot be targets.
	Terms []*V2BillingContractActionAddTermPricingOverrideCriterionTermParams `form:"terms" json:"terms,omitempty"`
	// Whether to include or exclude the selected recurring or scheduled_fee terms.
	Type *string `form:"type" json:"type"`
}

// The exclusive end. Omitted means open-ended.
type V2BillingContractActionAddTermPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is multiply.
type V2BillingContractActionAddTermPricingOverrideMultiplyParams struct {
	// The multiplier, for example 0.85 for a 15% discount.
	Factor *decimal.Decimal `form:"factor" json:"factor"`
}

// Required when type is overwrite.
type V2BillingContractActionAddTermPricingOverrideOverwriteParams struct {
	// The replacement per-unit amount in minor currency units.
	UnitAmountDecimal *decimal.Decimal `form:"unit_amount_decimal" json:"unit_amount_decimal,omitempty"`
}

// The inclusive start of the validity window.
type V2BillingContractActionAddTermPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is pricing_override.
type V2BillingContractActionAddTermPricingOverrideParams struct {
	// Optional criteria selecting recurring or scheduled_fee terms. If omitted, the override applies to all recurring and scheduled_fee terms.
	Criteria []*V2BillingContractActionAddTermPricingOverrideCriterionParams `form:"criteria" json:"criteria,omitempty"`
	// The exclusive end. Omitted means open-ended.
	EndsAt *V2BillingContractActionAddTermPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// Required when type is multiply.
	Multiply *V2BillingContractActionAddTermPricingOverrideMultiplyParams `form:"multiply" json:"multiply,omitempty"`
	// Required when type is overwrite.
	Overwrite *V2BillingContractActionAddTermPricingOverrideOverwriteParams `form:"overwrite" json:"overwrite,omitempty"`
	// Lower numbers win; defaults to 50.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
	// The inclusive start of the validity window.
	StartsAt *V2BillingContractActionAddTermPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at"`
	// The pricing operation. Exactly one matching pricing payload is required.
	Type *string `form:"type" json:"type"`
}

// The exclusive end. Omitted means open-ended.
type V2BillingContractActionAddTermRecurringPriceEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// The inclusive start of the term.
type V2BillingContractActionAddTermRecurringPriceStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is recurring_price.
type V2BillingContractActionAddTermRecurringPriceParams struct {
	// The exclusive end. Omitted means open-ended.
	EndsAt *V2BillingContractActionAddTermRecurringPriceEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// The V1 recurring Price.
	Price *string `form:"price" json:"price"`
	// Initial licensed quantity. Defaults to 1; must be omitted for metered Prices.
	Quantity *decimal.Decimal `form:"quantity" json:"quantity,omitempty"`
	// The inclusive start of the term.
	StartsAt *V2BillingContractActionAddTermRecurringPriceStartsAtParams `form:"starts_at" json:"starts_at"`
}

// A finite bill date. Resolved once per request, including on idempotent retries.
type V2BillingContractActionAddTermScheduledFeeBillScheduleCustomBillAtParams struct {
	// The timestamp at which the entry should be billed. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the bill_at.
	Type *string `form:"type" json:"type"`
}

// Between one and ten installments. Empty schedules are allowed only on update.
type V2BillingContractActionAddTermScheduledFeeBillScheduleCustomParams struct {
	// Non-negative amount with an explicit currency that must match the contract's currency.
	Amount *Amount `form:"amount" json:"amount"`
	// A finite bill date. Resolved once per request, including on idempotent retries.
	BillAt *V2BillingContractActionAddTermScheduledFeeBillScheduleCustomBillAtParams `form:"bill_at" json:"bill_at"`
}

// Explicit bill dates and amounts.
type V2BillingContractActionAddTermScheduledFeeBillScheduleParams struct {
	// Between one and ten installments. Empty schedules are allowed only on update.
	Custom []*V2BillingContractActionAddTermScheduledFeeBillScheduleCustomParams `form:"custom" json:"custom,omitempty"`
	// The type of schedule.
	Type *string `form:"type" json:"type"`
}

// The finite exclusive end; must be after starts_at. Never is not supported.
type V2BillingContractActionAddTermScheduledFeeServicePeriodEndsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// The finite inclusive start of the service period.
type V2BillingContractActionAddTermScheduledFeeServicePeriodStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Informational service period, independent of bill dates.
type V2BillingContractActionAddTermScheduledFeeServicePeriodParams struct {
	// The finite exclusive end; must be after starts_at. Never is not supported.
	EndsAt *V2BillingContractActionAddTermScheduledFeeServicePeriodEndsAtParams `form:"ends_at" json:"ends_at"`
	// The finite inclusive start of the service period.
	StartsAt *V2BillingContractActionAddTermScheduledFeeServicePeriodStartsAtParams `form:"starts_at" json:"starts_at"`
}

// Required when type is scheduled_fee.
type V2BillingContractActionAddTermScheduledFeeParams struct {
	// Explicit bill dates and amounts.
	BillSchedule *V2BillingContractActionAddTermScheduledFeeBillScheduleParams `form:"bill_schedule" json:"bill_schedule"`
	// The V1 Product being billed.
	Product *string `form:"product" json:"product"`
	// Informational service period, independent of bill dates.
	ServicePeriod *V2BillingContractActionAddTermScheduledFeeServicePeriodParams `form:"service_period" json:"service_period,omitempty"`
}

// Prepend a term, preserving relative order among additions in this request.
type V2BillingContractActionAddTermParams struct {
	// Lookup key unique across all term types on the contract.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs attached to the term.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Required when type is pricing_override.
	PricingOverride *V2BillingContractActionAddTermPricingOverrideParams `form:"pricing_override" json:"pricing_override,omitempty"`
	// Required when type is recurring_price.
	RecurringPrice *V2BillingContractActionAddTermRecurringPriceParams `form:"recurring_price" json:"recurring_price,omitempty"`
	// Required when type is scheduled_fee.
	ScheduledFee *V2BillingContractActionAddTermScheduledFeeParams `form:"scheduled_fee" json:"scheduled_fee,omitempty"`
	// The term type. Production types require exactly one matching payload.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractActionAddTermParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Select a term by ID or its unique lookup key.
type V2BillingContractActionRemoveTermTermParams struct {
	// Select this term by its stable ID.
	ID *string `form:"id" json:"id,omitempty"`
	// Select this term by its unique lookup key.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Whether the selector uses a term ID or lookup key.
	Type *string `form:"type" json:"type"`
}

// Remove a term that has not started.
type V2BillingContractActionRemoveTermParams struct {
	// Select a term by ID or its unique lookup key.
	Term *V2BillingContractActionRemoveTermTermParams `form:"term" json:"term"`
}

// Omission leaves the end unchanged; never requests an open end.
type V2BillingContractActionUpdateTermPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Can change only before the override starts. Accepts timestamp or now.
type V2BillingContractActionUpdateTermPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Only validity dates can change; pricing, criteria, and priority are immutable.
type V2BillingContractActionUpdateTermPricingOverrideParams struct {
	// Omission leaves the end unchanged; never requests an open end.
	EndsAt *V2BillingContractActionUpdateTermPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// Can change only before the override starts. Accepts timestamp or now.
	StartsAt *V2BillingContractActionUpdateTermPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
}

// Omission leaves the end unchanged; never requests an open end.
type V2BillingContractActionUpdateTermRecurringPriceEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Reschedule an unstarted term in place.
type V2BillingContractActionUpdateTermRecurringPriceStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Must match the existing term's type.
type V2BillingContractActionUpdateTermRecurringPriceParams struct {
	// Omission leaves the end unchanged; never requests an open end.
	EndsAt *V2BillingContractActionUpdateTermRecurringPriceEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// Update an active term's licensed quantity at the contract's current clock time,
	// or an unstarted term's initial quantity, preserving its ID. Ended terms cannot change quantity.
	Quantity *decimal.Decimal `form:"quantity" json:"quantity,omitempty"`
	// Reschedule an unstarted term in place.
	StartsAt *V2BillingContractActionUpdateTermRecurringPriceStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
}

// A finite bill date. Resolved once per request, including on idempotent retries.
type V2BillingContractActionUpdateTermScheduledFeeBillScheduleCustomBillAtParams struct {
	// The timestamp at which the entry should be billed. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the bill_at.
	Type *string `form:"type" json:"type"`
}

// Full replacement of up to ten unbilled installments. An explicit empty array cancels all of them.
type V2BillingContractActionUpdateTermScheduledFeeBillScheduleCustomParams struct {
	// Non-negative amount with an explicit currency that must match the contract's currency.
	Amount *Amount `form:"amount" json:"amount"`
	// A finite bill date. Resolved once per request, including on idempotent retries.
	BillAt *V2BillingContractActionUpdateTermScheduledFeeBillScheduleCustomBillAtParams `form:"bill_at" json:"bill_at"`
}

// Replace remaining unbilled installments. Billed entries remain immutable.
type V2BillingContractActionUpdateTermScheduledFeeBillScheduleParams struct {
	// Full replacement of up to ten unbilled installments. An explicit empty array cancels all of them.
	Custom []*V2BillingContractActionUpdateTermScheduledFeeBillScheduleCustomParams `form:"custom" json:"custom,omitempty"`
	// The type of schedule.
	Type *string `form:"type" json:"type"`
}

// The finite exclusive end; must be after starts_at. Never is not supported.
type V2BillingContractActionUpdateTermScheduledFeeServicePeriodEndsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// The finite inclusive start of the service period.
type V2BillingContractActionUpdateTermScheduledFeeServicePeriodStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Informational service period; cannot change historical billing.
type V2BillingContractActionUpdateTermScheduledFeeServicePeriodParams struct {
	// The finite exclusive end; must be after starts_at. Never is not supported.
	EndsAt *V2BillingContractActionUpdateTermScheduledFeeServicePeriodEndsAtParams `form:"ends_at" json:"ends_at"`
	// The finite inclusive start of the service period.
	StartsAt *V2BillingContractActionUpdateTermScheduledFeeServicePeriodStartsAtParams `form:"starts_at" json:"starts_at"`
}

// Must match the existing term's type.
type V2BillingContractActionUpdateTermScheduledFeeParams struct {
	// Replace remaining unbilled installments. Billed entries remain immutable.
	BillSchedule *V2BillingContractActionUpdateTermScheduledFeeBillScheduleParams `form:"bill_schedule" json:"bill_schedule,omitempty"`
	// Informational service period; cannot change historical billing.
	ServicePeriod *V2BillingContractActionUpdateTermScheduledFeeServicePeriodParams `form:"service_period" json:"service_period,omitempty"`
}

// Select a term by ID or its unique lookup key.
type V2BillingContractActionUpdateTermTermParams struct {
	// Select this term by its stable ID.
	ID *string `form:"id" json:"id,omitempty"`
	// Select this term by its unique lookup key.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Whether the selector uses a term ID or lookup key.
	Type *string `form:"type" json:"type"`
}

// Edit a term without changing its display position.
type V2BillingContractActionUpdateTermParams struct {
	// Metadata changes to apply immediately.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// Only validity dates can change; pricing, criteria, and priority are immutable.
	PricingOverride *V2BillingContractActionUpdateTermPricingOverrideParams `form:"pricing_override" json:"pricing_override,omitempty"`
	// Must match the existing term's type.
	RecurringPrice *V2BillingContractActionUpdateTermRecurringPriceParams `form:"recurring_price" json:"recurring_price,omitempty"`
	// Must match the existing term's type.
	ScheduledFee *V2BillingContractActionUpdateTermScheduledFeeParams `form:"scheduled_fee" json:"scheduled_fee,omitempty"`
	// Select a term by ID or its unique lookup key.
	Term *V2BillingContractActionUpdateTermTermParams `form:"term" json:"term"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractActionUpdateTermParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Apply all term actions atomically.
type V2BillingContractActionParams struct {
	// Prepend a term, preserving relative order among additions in this request.
	AddTerm *V2BillingContractActionAddTermParams `form:"add_term" json:"add_term,omitempty"`
	// Remove a term that has not started.
	RemoveTerm *V2BillingContractActionRemoveTermParams `form:"remove_term" json:"remove_term,omitempty"`
	// The action type. Exactly one matching payload is required.
	Type *string `form:"type" json:"type"`
	// Edit a term without changing its display position.
	UpdateTerm *V2BillingContractActionUpdateTermParams `form:"update_term" json:"update_term,omitempty"`
}

// Activate a draft contract.
type V2BillingContractActivateParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// Select a term by ID or its unique lookup key.
type V2BillingContractCancelTermTermParams struct {
	// Select this term by its stable ID.
	ID *string `form:"id" json:"id,omitempty"`
	// Select this term by its unique lookup key.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Whether the selector uses a term ID or lookup key.
	Type *string `form:"type" json:"type"`
}

// Per-term proration overrides.
type V2BillingContractCancelTermParams struct {
	// Falls back to the contract cancellation's top-level proration behavior.
	ProrationBehavior *string `form:"proration_behavior" json:"proration_behavior,omitempty"`
	// Select a term by ID or its unique lookup key.
	Term *V2BillingContractCancelTermTermParams `form:"term" json:"term"`
}

// Cancel an active contract.
type V2BillingContractCancelParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Top-level proration behavior for the cancellation. Defaults to `prorated` if not set.
	ProrationBehavior *string `form:"proration_behavior" json:"proration_behavior,omitempty"`
	// Per-term proration overrides.
	Terms []*V2BillingContractCancelTermParams `form:"terms" json:"terms,omitempty"`
}

// Configuration for determining the billing cycle anchor by calendar fields.
type V2BillingContractCreateBillingCycleAnchorConfigParams struct {
	// Day of month (1-31).
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// Hour of day in UTC (0-23).
	Hour *int64 `form:"hour" json:"hour,omitempty"`
	// Minute of hour (0-59).
	Minute *int64 `form:"minute" json:"minute,omitempty"`
	// Month of year (1-12).
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// Second of minute (0-59).
	Second *int64 `form:"second" json:"second,omitempty"`
}

// The billing cycle anchor for the contract. If not provided, defaults to the pricing line start time.
// It is only at the top-level of the contract with no option to override at the pricing line level.
type V2BillingContractCreateBillingCycleAnchorParams struct {
	// Configuration for determining the billing cycle anchor by calendar fields.
	Config *V2BillingContractCreateBillingCycleAnchorConfigParams `form:"config" json:"config,omitempty"`
	// A specific timestamp to use as the billing cycle anchor.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of billing cycle anchor.
	Type *string `form:"type" json:"type"`
}

// Tax calculation settings.
type V2BillingContractCreateBillingSettingsBillSettingsDetailsCalculationTaxParams struct {
	// The type of tax calculation.
	Type *string `form:"type" json:"type"`
}

// Calculation settings.
type V2BillingContractCreateBillingSettingsBillSettingsDetailsCalculationParams struct {
	// Tax calculation settings.
	Tax *V2BillingContractCreateBillingSettingsBillSettingsDetailsCalculationTaxParams `form:"tax" json:"tax,omitempty"`
}

// How long the customer has to pay the invoice before it's past due.
type V2BillingContractCreateBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueParams struct {
	// The interval unit.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Invoice settings.
type V2BillingContractCreateBillingSettingsBillSettingsDetailsInvoiceParams struct {
	// How long the customer has to pay the invoice before it's past due.
	TimeUntilDue *V2BillingContractCreateBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueParams `form:"time_until_due" json:"time_until_due,omitempty"`
}

// The bill settings details configures invoice and tax settings for the contract.
type V2BillingContractCreateBillingSettingsBillSettingsDetailsParams struct {
	// Calculation settings.
	Calculation *V2BillingContractCreateBillingSettingsBillSettingsDetailsCalculationParams `form:"calculation" json:"calculation,omitempty"`
	// Invoice settings.
	Invoice *V2BillingContractCreateBillingSettingsBillSettingsDetailsInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// The billing profile details configures who is charged for the contract.
type V2BillingContractCreateBillingSettingsBillingProfileDetailsParams struct {
	// The customer who pays for the contract invoice.
	Customer *string `form:"customer" json:"customer"`
	// The default payment method for the contract.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
}

// The collection settings details configures how payments are collected on the contract.
type V2BillingContractCreateBillingSettingsCollectionSettingsDetailsParams struct {
	// How payment is collected for the contract.
	CollectionMethod *string `form:"collection_method" json:"collection_method"`
	// The payment method configuration.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
}

// The billing settings for the contract.
type V2BillingContractCreateBillingSettingsParams struct {
	// The billing profile details configures who is charged for the contract.
	BillingProfileDetails *V2BillingContractCreateBillingSettingsBillingProfileDetailsParams `form:"billing_profile_details" json:"billing_profile_details"`
	// The bill settings details configures invoice and tax settings for the contract.
	BillSettingsDetails *V2BillingContractCreateBillingSettingsBillSettingsDetailsParams `form:"bill_settings_details" json:"bill_settings_details,omitempty"`
	// The collection settings details configures how payments are collected on the contract.
	CollectionSettingsDetails *V2BillingContractCreateBillingSettingsCollectionSettingsDetailsParams `form:"collection_settings_details" json:"collection_settings_details"`
}

// Term selectors. Each ID or lookup key resolves to one fixed term ID.
// Omitted or empty selects all charge-bearing terms. Override terms cannot be targets.
type V2BillingContractCreateTermPricingOverrideCriterionTermParams struct {
	// Select this term by its stable ID.
	ID *string `form:"id" json:"id,omitempty"`
	// Select this term by its unique lookup key.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Whether the selector uses a term ID or lookup key.
	Type *string `form:"type" json:"type"`
}

// Optional criteria selecting recurring or scheduled_fee terms. If omitted, the override applies to all recurring and scheduled_fee terms.
type V2BillingContractCreateTermPricingOverrideCriterionParams struct {
	// Term selectors. Each ID or lookup key resolves to one fixed term ID.
	// Omitted or empty selects all charge-bearing terms. Override terms cannot be targets.
	Terms []*V2BillingContractCreateTermPricingOverrideCriterionTermParams `form:"terms" json:"terms,omitempty"`
	// Whether to include or exclude the selected recurring or scheduled_fee terms.
	Type *string `form:"type" json:"type"`
}

// The exclusive end. Omitted means open-ended.
type V2BillingContractCreateTermPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is multiply.
type V2BillingContractCreateTermPricingOverrideMultiplyParams struct {
	// The multiplier, for example 0.85 for a 15% discount.
	Factor *decimal.Decimal `form:"factor" json:"factor"`
}

// Required when type is overwrite.
type V2BillingContractCreateTermPricingOverrideOverwriteParams struct {
	// The replacement per-unit amount in minor currency units.
	UnitAmountDecimal *decimal.Decimal `form:"unit_amount_decimal" json:"unit_amount_decimal,omitempty"`
}

// The inclusive start of the validity window.
type V2BillingContractCreateTermPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is pricing_override.
type V2BillingContractCreateTermPricingOverrideParams struct {
	// Optional criteria selecting recurring or scheduled_fee terms. If omitted, the override applies to all recurring and scheduled_fee terms.
	Criteria []*V2BillingContractCreateTermPricingOverrideCriterionParams `form:"criteria" json:"criteria,omitempty"`
	// The exclusive end. Omitted means open-ended.
	EndsAt *V2BillingContractCreateTermPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// Required when type is multiply.
	Multiply *V2BillingContractCreateTermPricingOverrideMultiplyParams `form:"multiply" json:"multiply,omitempty"`
	// Required when type is overwrite.
	Overwrite *V2BillingContractCreateTermPricingOverrideOverwriteParams `form:"overwrite" json:"overwrite,omitempty"`
	// Lower numbers win; defaults to 50.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
	// The inclusive start of the validity window.
	StartsAt *V2BillingContractCreateTermPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at"`
	// The pricing operation. Exactly one matching pricing payload is required.
	Type *string `form:"type" json:"type"`
}

// The exclusive end. Omitted means open-ended.
type V2BillingContractCreateTermRecurringPriceEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// The inclusive start of the term.
type V2BillingContractCreateTermRecurringPriceStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is recurring_price.
type V2BillingContractCreateTermRecurringPriceParams struct {
	// The exclusive end. Omitted means open-ended.
	EndsAt *V2BillingContractCreateTermRecurringPriceEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// The V1 recurring Price.
	Price *string `form:"price" json:"price"`
	// Initial licensed quantity. Defaults to 1; must be omitted for metered Prices.
	Quantity *decimal.Decimal `form:"quantity" json:"quantity,omitempty"`
	// The inclusive start of the term.
	StartsAt *V2BillingContractCreateTermRecurringPriceStartsAtParams `form:"starts_at" json:"starts_at"`
}

// A finite bill date. Resolved once per request, including on idempotent retries.
type V2BillingContractCreateTermScheduledFeeBillScheduleCustomBillAtParams struct {
	// The timestamp at which the entry should be billed. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the bill_at.
	Type *string `form:"type" json:"type"`
}

// Between one and ten installments. Empty schedules are allowed only on update.
type V2BillingContractCreateTermScheduledFeeBillScheduleCustomParams struct {
	// Non-negative amount with an explicit currency that must match the contract's currency.
	Amount *Amount `form:"amount" json:"amount"`
	// A finite bill date. Resolved once per request, including on idempotent retries.
	BillAt *V2BillingContractCreateTermScheduledFeeBillScheduleCustomBillAtParams `form:"bill_at" json:"bill_at"`
}

// Explicit bill dates and amounts.
type V2BillingContractCreateTermScheduledFeeBillScheduleParams struct {
	// Between one and ten installments. Empty schedules are allowed only on update.
	Custom []*V2BillingContractCreateTermScheduledFeeBillScheduleCustomParams `form:"custom" json:"custom,omitempty"`
	// The type of schedule.
	Type *string `form:"type" json:"type"`
}

// The finite exclusive end; must be after starts_at. Never is not supported.
type V2BillingContractCreateTermScheduledFeeServicePeriodEndsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// The finite inclusive start of the service period.
type V2BillingContractCreateTermScheduledFeeServicePeriodStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Informational service period, independent of bill dates.
type V2BillingContractCreateTermScheduledFeeServicePeriodParams struct {
	// The finite exclusive end; must be after starts_at. Never is not supported.
	EndsAt *V2BillingContractCreateTermScheduledFeeServicePeriodEndsAtParams `form:"ends_at" json:"ends_at"`
	// The finite inclusive start of the service period.
	StartsAt *V2BillingContractCreateTermScheduledFeeServicePeriodStartsAtParams `form:"starts_at" json:"starts_at"`
}

// Required when type is scheduled_fee.
type V2BillingContractCreateTermScheduledFeeParams struct {
	// Explicit bill dates and amounts.
	BillSchedule *V2BillingContractCreateTermScheduledFeeBillScheduleParams `form:"bill_schedule" json:"bill_schedule"`
	// The V1 Product being billed.
	Product *string `form:"product" json:"product"`
	// Informational service period, independent of bill dates.
	ServicePeriod *V2BillingContractCreateTermScheduledFeeServicePeriodParams `form:"service_period" json:"service_period,omitempty"`
}

// Terms to create with the contract.
type V2BillingContractCreateTermParams struct {
	// Lookup key unique across all term types on the contract.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs attached to the term.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Required when type is pricing_override.
	PricingOverride *V2BillingContractCreateTermPricingOverrideParams `form:"pricing_override" json:"pricing_override,omitempty"`
	// Required when type is recurring_price.
	RecurringPrice *V2BillingContractCreateTermRecurringPriceParams `form:"recurring_price" json:"recurring_price,omitempty"`
	// Required when type is scheduled_fee.
	ScheduledFee *V2BillingContractCreateTermScheduledFeeParams `form:"scheduled_fee" json:"scheduled_fee,omitempty"`
	// The term type. Production types require exactly one matching payload.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractCreateTermParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Create a draft contract.
type V2BillingContractCreateParams struct {
	Params `form:"*"`
	// The billing cycle anchor for the contract. If not provided, defaults to the pricing line start time.
	// It is only at the top-level of the contract with no option to override at the pricing line level.
	BillingCycleAnchor *V2BillingContractCreateBillingCycleAnchorParams `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// The billing settings for the contract.
	BillingSettings *V2BillingContractCreateBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// A unique user-provided contract number e.g. C-2026-0001.
	// Maximum length of 200 characters.
	ContractNumber *string `form:"contract_number" json:"contract_number"`
	// Currency of the contract.
	Currency *string `form:"currency" json:"currency"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Terms to create with the contract.
	Terms []*V2BillingContractCreateTermParams `form:"terms" json:"terms,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Delete a draft contract.
type V2BillingContractDeleteParams struct {
	Params `form:"*"`
}

// Retrieve a contract.
type V2BillingContractRetrieveParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// Term selectors. Each ID or lookup key resolves to one fixed term ID.
// Omitted or empty selects all charge-bearing terms. Override terms cannot be targets.
type V2BillingContractUpdateActionAddTermPricingOverrideCriterionTermParams struct {
	// Select this term by its stable ID.
	ID *string `form:"id" json:"id,omitempty"`
	// Select this term by its unique lookup key.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Whether the selector uses a term ID or lookup key.
	Type *string `form:"type" json:"type"`
}

// Optional criteria selecting recurring or scheduled_fee terms. If omitted, the override applies to all recurring and scheduled_fee terms.
type V2BillingContractUpdateActionAddTermPricingOverrideCriterionParams struct {
	// Term selectors. Each ID or lookup key resolves to one fixed term ID.
	// Omitted or empty selects all charge-bearing terms. Override terms cannot be targets.
	Terms []*V2BillingContractUpdateActionAddTermPricingOverrideCriterionTermParams `form:"terms" json:"terms,omitempty"`
	// Whether to include or exclude the selected recurring or scheduled_fee terms.
	Type *string `form:"type" json:"type"`
}

// The exclusive end. Omitted means open-ended.
type V2BillingContractUpdateActionAddTermPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is multiply.
type V2BillingContractUpdateActionAddTermPricingOverrideMultiplyParams struct {
	// The multiplier, for example 0.85 for a 15% discount.
	Factor *decimal.Decimal `form:"factor" json:"factor"`
}

// Required when type is overwrite.
type V2BillingContractUpdateActionAddTermPricingOverrideOverwriteParams struct {
	// The replacement per-unit amount in minor currency units.
	UnitAmountDecimal *decimal.Decimal `form:"unit_amount_decimal" json:"unit_amount_decimal,omitempty"`
}

// The inclusive start of the validity window.
type V2BillingContractUpdateActionAddTermPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is pricing_override.
type V2BillingContractUpdateActionAddTermPricingOverrideParams struct {
	// Optional criteria selecting recurring or scheduled_fee terms. If omitted, the override applies to all recurring and scheduled_fee terms.
	Criteria []*V2BillingContractUpdateActionAddTermPricingOverrideCriterionParams `form:"criteria" json:"criteria,omitempty"`
	// The exclusive end. Omitted means open-ended.
	EndsAt *V2BillingContractUpdateActionAddTermPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// Required when type is multiply.
	Multiply *V2BillingContractUpdateActionAddTermPricingOverrideMultiplyParams `form:"multiply" json:"multiply,omitempty"`
	// Required when type is overwrite.
	Overwrite *V2BillingContractUpdateActionAddTermPricingOverrideOverwriteParams `form:"overwrite" json:"overwrite,omitempty"`
	// Lower numbers win; defaults to 50.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
	// The inclusive start of the validity window.
	StartsAt *V2BillingContractUpdateActionAddTermPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at"`
	// The pricing operation. Exactly one matching pricing payload is required.
	Type *string `form:"type" json:"type"`
}

// The exclusive end. Omitted means open-ended.
type V2BillingContractUpdateActionAddTermRecurringPriceEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// The inclusive start of the term.
type V2BillingContractUpdateActionAddTermRecurringPriceStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Required when type is recurring_price.
type V2BillingContractUpdateActionAddTermRecurringPriceParams struct {
	// The exclusive end. Omitted means open-ended.
	EndsAt *V2BillingContractUpdateActionAddTermRecurringPriceEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// The V1 recurring Price.
	Price *string `form:"price" json:"price"`
	// Initial licensed quantity. Defaults to 1; must be omitted for metered Prices.
	Quantity *decimal.Decimal `form:"quantity" json:"quantity,omitempty"`
	// The inclusive start of the term.
	StartsAt *V2BillingContractUpdateActionAddTermRecurringPriceStartsAtParams `form:"starts_at" json:"starts_at"`
}

// A finite bill date. Resolved once per request, including on idempotent retries.
type V2BillingContractUpdateActionAddTermScheduledFeeBillScheduleCustomBillAtParams struct {
	// The timestamp at which the entry should be billed. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the bill_at.
	Type *string `form:"type" json:"type"`
}

// Between one and ten installments. Empty schedules are allowed only on update.
type V2BillingContractUpdateActionAddTermScheduledFeeBillScheduleCustomParams struct {
	// Non-negative amount with an explicit currency that must match the contract's currency.
	Amount *Amount `form:"amount" json:"amount"`
	// A finite bill date. Resolved once per request, including on idempotent retries.
	BillAt *V2BillingContractUpdateActionAddTermScheduledFeeBillScheduleCustomBillAtParams `form:"bill_at" json:"bill_at"`
}

// Explicit bill dates and amounts.
type V2BillingContractUpdateActionAddTermScheduledFeeBillScheduleParams struct {
	// Between one and ten installments. Empty schedules are allowed only on update.
	Custom []*V2BillingContractUpdateActionAddTermScheduledFeeBillScheduleCustomParams `form:"custom" json:"custom,omitempty"`
	// The type of schedule.
	Type *string `form:"type" json:"type"`
}

// The finite exclusive end; must be after starts_at. Never is not supported.
type V2BillingContractUpdateActionAddTermScheduledFeeServicePeriodEndsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// The finite inclusive start of the service period.
type V2BillingContractUpdateActionAddTermScheduledFeeServicePeriodStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Informational service period, independent of bill dates.
type V2BillingContractUpdateActionAddTermScheduledFeeServicePeriodParams struct {
	// The finite exclusive end; must be after starts_at. Never is not supported.
	EndsAt *V2BillingContractUpdateActionAddTermScheduledFeeServicePeriodEndsAtParams `form:"ends_at" json:"ends_at"`
	// The finite inclusive start of the service period.
	StartsAt *V2BillingContractUpdateActionAddTermScheduledFeeServicePeriodStartsAtParams `form:"starts_at" json:"starts_at"`
}

// Required when type is scheduled_fee.
type V2BillingContractUpdateActionAddTermScheduledFeeParams struct {
	// Explicit bill dates and amounts.
	BillSchedule *V2BillingContractUpdateActionAddTermScheduledFeeBillScheduleParams `form:"bill_schedule" json:"bill_schedule"`
	// The V1 Product being billed.
	Product *string `form:"product" json:"product"`
	// Informational service period, independent of bill dates.
	ServicePeriod *V2BillingContractUpdateActionAddTermScheduledFeeServicePeriodParams `form:"service_period" json:"service_period,omitempty"`
}

// Prepend a term, preserving relative order among additions in this request.
type V2BillingContractUpdateActionAddTermParams struct {
	// Lookup key unique across all term types on the contract.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs attached to the term.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Required when type is pricing_override.
	PricingOverride *V2BillingContractUpdateActionAddTermPricingOverrideParams `form:"pricing_override" json:"pricing_override,omitempty"`
	// Required when type is recurring_price.
	RecurringPrice *V2BillingContractUpdateActionAddTermRecurringPriceParams `form:"recurring_price" json:"recurring_price,omitempty"`
	// Required when type is scheduled_fee.
	ScheduledFee *V2BillingContractUpdateActionAddTermScheduledFeeParams `form:"scheduled_fee" json:"scheduled_fee,omitempty"`
	// The term type. Production types require exactly one matching payload.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractUpdateActionAddTermParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Select a term by ID or its unique lookup key.
type V2BillingContractUpdateActionRemoveTermTermParams struct {
	// Select this term by its stable ID.
	ID *string `form:"id" json:"id,omitempty"`
	// Select this term by its unique lookup key.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Whether the selector uses a term ID or lookup key.
	Type *string `form:"type" json:"type"`
}

// Remove a term that has not started.
type V2BillingContractUpdateActionRemoveTermParams struct {
	// Select a term by ID or its unique lookup key.
	Term *V2BillingContractUpdateActionRemoveTermTermParams `form:"term" json:"term"`
}

// Omission leaves the end unchanged; never requests an open end.
type V2BillingContractUpdateActionUpdateTermPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Can change only before the override starts. Accepts timestamp or now.
type V2BillingContractUpdateActionUpdateTermPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Only validity dates can change; pricing, criteria, and priority are immutable.
type V2BillingContractUpdateActionUpdateTermPricingOverrideParams struct {
	// Omission leaves the end unchanged; never requests an open end.
	EndsAt *V2BillingContractUpdateActionUpdateTermPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// Can change only before the override starts. Accepts timestamp or now.
	StartsAt *V2BillingContractUpdateActionUpdateTermPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
}

// Omission leaves the end unchanged; never requests an open end.
type V2BillingContractUpdateActionUpdateTermRecurringPriceEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Reschedule an unstarted term in place.
type V2BillingContractUpdateActionUpdateTermRecurringPriceStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Must match the existing term's type.
type V2BillingContractUpdateActionUpdateTermRecurringPriceParams struct {
	// Omission leaves the end unchanged; never requests an open end.
	EndsAt *V2BillingContractUpdateActionUpdateTermRecurringPriceEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// Update an active term's licensed quantity at the contract's current clock time,
	// or an unstarted term's initial quantity, preserving its ID. Ended terms cannot change quantity.
	Quantity *decimal.Decimal `form:"quantity" json:"quantity,omitempty"`
	// Reschedule an unstarted term in place.
	StartsAt *V2BillingContractUpdateActionUpdateTermRecurringPriceStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
}

// A finite bill date. Resolved once per request, including on idempotent retries.
type V2BillingContractUpdateActionUpdateTermScheduledFeeBillScheduleCustomBillAtParams struct {
	// The timestamp at which the entry should be billed. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the bill_at.
	Type *string `form:"type" json:"type"`
}

// Full replacement of up to ten unbilled installments. An explicit empty array cancels all of them.
type V2BillingContractUpdateActionUpdateTermScheduledFeeBillScheduleCustomParams struct {
	// Non-negative amount with an explicit currency that must match the contract's currency.
	Amount *Amount `form:"amount" json:"amount"`
	// A finite bill date. Resolved once per request, including on idempotent retries.
	BillAt *V2BillingContractUpdateActionUpdateTermScheduledFeeBillScheduleCustomBillAtParams `form:"bill_at" json:"bill_at"`
}

// Replace remaining unbilled installments. Billed entries remain immutable.
type V2BillingContractUpdateActionUpdateTermScheduledFeeBillScheduleParams struct {
	// Full replacement of up to ten unbilled installments. An explicit empty array cancels all of them.
	Custom []*V2BillingContractUpdateActionUpdateTermScheduledFeeBillScheduleCustomParams `form:"custom" json:"custom,omitempty"`
	// The type of schedule.
	Type *string `form:"type" json:"type"`
}

// The finite exclusive end; must be after starts_at. Never is not supported.
type V2BillingContractUpdateActionUpdateTermScheduledFeeServicePeriodEndsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// The finite inclusive start of the service period.
type V2BillingContractUpdateActionUpdateTermScheduledFeeServicePeriodStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Informational service period; cannot change historical billing.
type V2BillingContractUpdateActionUpdateTermScheduledFeeServicePeriodParams struct {
	// The finite exclusive end; must be after starts_at. Never is not supported.
	EndsAt *V2BillingContractUpdateActionUpdateTermScheduledFeeServicePeriodEndsAtParams `form:"ends_at" json:"ends_at"`
	// The finite inclusive start of the service period.
	StartsAt *V2BillingContractUpdateActionUpdateTermScheduledFeeServicePeriodStartsAtParams `form:"starts_at" json:"starts_at"`
}

// Must match the existing term's type.
type V2BillingContractUpdateActionUpdateTermScheduledFeeParams struct {
	// Replace remaining unbilled installments. Billed entries remain immutable.
	BillSchedule *V2BillingContractUpdateActionUpdateTermScheduledFeeBillScheduleParams `form:"bill_schedule" json:"bill_schedule,omitempty"`
	// Informational service period; cannot change historical billing.
	ServicePeriod *V2BillingContractUpdateActionUpdateTermScheduledFeeServicePeriodParams `form:"service_period" json:"service_period,omitempty"`
}

// Select a term by ID or its unique lookup key.
type V2BillingContractUpdateActionUpdateTermTermParams struct {
	// Select this term by its stable ID.
	ID *string `form:"id" json:"id,omitempty"`
	// Select this term by its unique lookup key.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Whether the selector uses a term ID or lookup key.
	Type *string `form:"type" json:"type"`
}

// Edit a term without changing its display position.
type V2BillingContractUpdateActionUpdateTermParams struct {
	// Metadata changes to apply immediately.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// Only validity dates can change; pricing, criteria, and priority are immutable.
	PricingOverride *V2BillingContractUpdateActionUpdateTermPricingOverrideParams `form:"pricing_override" json:"pricing_override,omitempty"`
	// Must match the existing term's type.
	RecurringPrice *V2BillingContractUpdateActionUpdateTermRecurringPriceParams `form:"recurring_price" json:"recurring_price,omitempty"`
	// Must match the existing term's type.
	ScheduledFee *V2BillingContractUpdateActionUpdateTermScheduledFeeParams `form:"scheduled_fee" json:"scheduled_fee,omitempty"`
	// Select a term by ID or its unique lookup key.
	Term *V2BillingContractUpdateActionUpdateTermTermParams `form:"term" json:"term"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractUpdateActionUpdateTermParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Apply all term actions atomically.
type V2BillingContractUpdateActionParams struct {
	// Prepend a term, preserving relative order among additions in this request.
	AddTerm *V2BillingContractUpdateActionAddTermParams `form:"add_term" json:"add_term,omitempty"`
	// Remove a term that has not started.
	RemoveTerm *V2BillingContractUpdateActionRemoveTermParams `form:"remove_term" json:"remove_term,omitempty"`
	// The action type. Exactly one matching payload is required.
	Type *string `form:"type" json:"type"`
	// Edit a term without changing its display position.
	UpdateTerm *V2BillingContractUpdateActionUpdateTermParams `form:"update_term" json:"update_term,omitempty"`
}

// Tax calculation settings.
type V2BillingContractUpdateBillingSettingsBillSettingsDetailsCalculationTaxParams struct {
	// The type of tax calculation.
	Type *string `form:"type" json:"type"`
}

// The tax calculation settings to update.
type V2BillingContractUpdateBillingSettingsBillSettingsDetailsCalculationParams struct {
	// Tax calculation settings.
	Tax *V2BillingContractUpdateBillingSettingsBillSettingsDetailsCalculationTaxParams `form:"tax" json:"tax,omitempty"`
}

// How long the customer has to pay the invoice before it's past due.
type V2BillingContractUpdateBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueParams struct {
	// The interval unit.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// The invoice settings to update.
type V2BillingContractUpdateBillingSettingsBillSettingsDetailsInvoiceParams struct {
	// How long the customer has to pay the invoice before it's past due.
	TimeUntilDue *V2BillingContractUpdateBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueParams `form:"time_until_due" json:"time_until_due,omitempty"`
}

// The bill settings to update (tax calculation type and/or invoice time until due).
type V2BillingContractUpdateBillingSettingsBillSettingsDetailsParams struct {
	// The tax calculation settings to update.
	Calculation *V2BillingContractUpdateBillingSettingsBillSettingsDetailsCalculationParams `form:"calculation" json:"calculation,omitempty"`
	// The invoice settings to update.
	Invoice *V2BillingContractUpdateBillingSettingsBillSettingsDetailsInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// The billing profile details to update.
type V2BillingContractUpdateBillingSettingsBillingProfileDetailsParams struct {
	// The default payment method to charge for the contract.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
}

// The collection settings details to update on the contract.
type V2BillingContractUpdateBillingSettingsCollectionSettingsDetailsParams struct {
	// How payment is collected for the contract. An omitted value leaves the
	// collection method unchanged.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// The payment method configuration.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
}

// The billing settings to update on the contract.
type V2BillingContractUpdateBillingSettingsParams struct {
	// The billing profile details to update.
	BillingProfileDetails *V2BillingContractUpdateBillingSettingsBillingProfileDetailsParams `form:"billing_profile_details" json:"billing_profile_details,omitempty"`
	// The bill settings to update (tax calculation type and/or invoice time until due).
	BillSettingsDetails *V2BillingContractUpdateBillingSettingsBillSettingsDetailsParams `form:"bill_settings_details" json:"bill_settings_details,omitempty"`
	// The collection settings details to update on the contract.
	CollectionSettingsDetails *V2BillingContractUpdateBillingSettingsCollectionSettingsDetailsParams `form:"collection_settings_details" json:"collection_settings_details,omitempty"`
}

// Update a draft or active contract.
type V2BillingContractUpdateParams struct {
	Params `form:"*"`
	// Apply all term actions atomically.
	Actions []*V2BillingContractUpdateActionParams `form:"actions" json:"actions,omitempty"`
	// The billing settings to update on the contract.
	BillingSettings *V2BillingContractUpdateBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// The unique user-provided contract number. This can only be changed while the contract is in draft.
	ContractNumber *string `form:"contract_number" json:"contract_number,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
