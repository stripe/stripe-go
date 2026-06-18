//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List Contract objects with pagination.
type V2BillingContractListParams struct {
	Params `form:"*"`
	// Filter by customer ID.
	Customer *string `form:"customer" json:"customer,omitempty"`
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
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxParams struct {
	// The type of tax calculation.
	Type *string `form:"type" json:"type"`
}

// Calculation settings.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationParams struct {
	// Tax calculation settings.
	Tax *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxParams `form:"tax" json:"tax,omitempty"`
}

// The number of time units before the invoice is past due.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueParams struct {
	// The interval unit.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Invoice settings.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceParams struct {
	// The number of time units before the invoice is past due.
	TimeUntilDue *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueParams `form:"time_until_due" json:"time_until_due,omitempty"`
}

// The bill settings details.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsParams struct {
	// Calculation settings.
	Calculation *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationParams `form:"calculation" json:"calculation,omitempty"`
	// Invoice settings.
	Invoice *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// The billing profile details.
type V2BillingContractBillingSettingsContractBillingDetailsBillingProfileDetailsParams struct {
	// The customer who pays for the contract invoice.
	Customer *string `form:"customer" json:"customer"`
	// The default payment method for the contract.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
}

// The collection settings details.
type V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsParams struct {
	// The collection method.
	CollectionMethod *string `form:"collection_method" json:"collection_method"`
	// The payment method configuration.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
}

// Billing settings details for the contract.
type V2BillingContractBillingSettingsContractBillingDetailsParams struct {
	// The billing profile details.
	BillingProfileDetails *V2BillingContractBillingSettingsContractBillingDetailsBillingProfileDetailsParams `form:"billing_profile_details" json:"billing_profile_details"`
	// The bill settings details.
	BillSettingsDetails *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsParams `form:"bill_settings_details" json:"bill_settings_details,omitempty"`
	// The collection settings details.
	CollectionSettingsDetails *V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsParams `form:"collection_settings_details" json:"collection_settings_details"`
}

// The billing settings for the contract.
type V2BillingContractBillingSettingsParams struct {
	// Billing settings details for the contract.
	ContractBillingDetails *V2BillingContractBillingSettingsContractBillingDetailsParams `form:"contract_billing_details" json:"contract_billing_details,omitempty"`
}

// When this entry should be billed.
type V2BillingContractOneTimeFeeBillScheduleBillAtParams struct {
	// The datetime at which the entry should be billed. Required if `type` is `datetime`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the bill_at.
	Type *string `form:"type" json:"type"`
}

// The bill schedule for the fee. Each entry produces an individual invoice item billed at `bill_at`.
type V2BillingContractOneTimeFeeBillScheduleParams struct {
	// When this entry should be billed.
	BillAt *V2BillingContractOneTimeFeeBillScheduleBillAtParams `form:"bill_at" json:"bill_at"`
	// The amount to bill for this entry, in the smallest currency unit.
	Value *int64 `form:"value" json:"value,string"`
}

// Details for a product billable target. Required when `billable_item_type` is `product`.
type V2BillingContractOneTimeFeeProductDetailsParams struct {
	// The ID of the v1 Product.
	Product *string `form:"product" json:"product"`
}

// A list of one-time fees to create with the contract. Each fee is billed as individual invoice items per its bill_schedule.
type V2BillingContractOneTimeFeeParams struct {
	// The type of billable item that this fee references.
	BillableItemType *string `form:"billable_item_type" json:"billable_item_type"`
	// The bill schedule for the fee. Each entry produces an individual invoice item billed at `bill_at`.
	BillSchedule []*V2BillingContractOneTimeFeeBillScheduleParams `form:"bill_schedule" json:"bill_schedule"`
	// Details for a product billable target. Required when `billable_item_type` is `product`.
	ProductDetails *V2BillingContractOneTimeFeeProductDetailsParams `form:"product_details" json:"product_details,omitempty"`
}

// When the pricing line ends.
type V2BillingContractPricingLineEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// When the override ends. Defaults to the pricing line's end if not specified.
type V2BillingContractPricingLinePricingPriceDetailsPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Each element represents a pricing tier.
type V2BillingContractPricingLinePricingPriceDetailsPricingOverrideOverwritePriceTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Parameters for the overwrite_price override. Required if `type` is `overwrite_price`.
type V2BillingContractPricingLinePricingPriceDetailsPricingOverrideOverwritePriceParams struct {
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractPricingLinePricingPriceDetailsPricingOverrideOverwritePriceTierParams `form:"tiers" json:"tiers,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// When the override starts. Defaults to the pricing line's start if not specified.
type V2BillingContractPricingLinePricingPriceDetailsPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Pricing overrides embedded directly on this pricing line.
type V2BillingContractPricingLinePricingPriceDetailsPricingOverrideParams struct {
	// When the override ends. Defaults to the pricing line's end if not specified.
	EndsAt *V2BillingContractPricingLinePricingPriceDetailsPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// A user-provided lookup key to reference this override.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Parameters for the overwrite_price override. Required if `type` is `overwrite_price`.
	OverwritePrice *V2BillingContractPricingLinePricingPriceDetailsPricingOverrideOverwritePriceParams `form:"overwrite_price" json:"overwrite_price,omitempty"`
	// The priority of this override relative to others. 0 is highest, 100 is lowest. Defaults to 50.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
	// When the override starts. Defaults to the pricing line's start if not specified.
	StartsAt *V2BillingContractPricingLinePricingPriceDetailsPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
	// The type of override. Currently only `overwrite_price` is supported.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractPricingLinePricingPriceDetailsPricingOverrideParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// When this quantity change takes effect.
type V2BillingContractPricingLinePricingPriceDetailsQuantityChangeEffectiveAtParams struct {
	// The timestamp for the effective at.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the effective at.
	Type *string `form:"type" json:"type"`
}

// Quantity changes for the pricing line. For now, at most one entry is allowed.
// A quantity change clears all future quantity changes on this pricing line.
type V2BillingContractPricingLinePricingPriceDetailsQuantityChangeParams struct {
	// When this quantity change takes effect.
	EffectiveAt *V2BillingContractPricingLinePricingPriceDetailsQuantityChangeEffectiveAtParams `form:"effective_at" json:"effective_at"`
	// The quantity to set.
	Set *float64 `form:"set,high_precision" json:"set,string"`
}

// V1 price details. Required if `type` is `price`.
type V2BillingContractPricingLinePricingPriceDetailsParams struct {
	// The ID of the V1 price.
	Price *string `form:"price" json:"price"`
	// Pricing overrides embedded directly on this pricing line.
	PricingOverrides []*V2BillingContractPricingLinePricingPriceDetailsPricingOverrideParams `form:"pricing_overrides" json:"pricing_overrides,omitempty"`
	// Quantity changes for the pricing line. For now, at most one entry is allowed.
	// A quantity change clears all future quantity changes on this pricing line.
	QuantityChanges []*V2BillingContractPricingLinePricingPriceDetailsQuantityChangeParams `form:"quantity_changes" json:"quantity_changes,omitempty"`
}

// The pricing configuration for the pricing line.
type V2BillingContractPricingLinePricingParams struct {
	// V1 price details. Required if `type` is `price`.
	PriceDetails *V2BillingContractPricingLinePricingPriceDetailsParams `form:"price_details" json:"price_details,omitempty"`
	// The type of pricing.
	Type *string `form:"type" json:"type"`
}

// When the pricing line starts.
type V2BillingContractPricingLineStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// A list of pricing lines to create with the contract.
type V2BillingContractPricingLineParams struct {
	// When the pricing line ends.
	EndsAt *V2BillingContractPricingLineEndsAtParams `form:"ends_at" json:"ends_at"`
	// A user-provided lookup key to reference this pricing line.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The pricing configuration for the pricing line.
	Pricing *V2BillingContractPricingLinePricingParams `form:"pricing" json:"pricing"`
	// When the pricing line starts.
	StartsAt *V2BillingContractPricingLineStartsAtParams `form:"starts_at" json:"starts_at"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractPricingLineParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// When the pricing override ends.
type V2BillingContractPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// All of these key-value conditions must match.
type V2BillingContractPricingOverrideMultiplierCriterionMetadataConditionAllOfParams struct {
	// The metadata key.
	Key *string `form:"key" json:"key"`
	// The metadata value.
	Value *string `form:"value" json:"value"`
}

// Filter by metadata conditions.
type V2BillingContractPricingOverrideMultiplierCriterionMetadataConditionParams struct {
	// All of these key-value conditions must match.
	AllOf []*V2BillingContractPricingOverrideMultiplierCriterionMetadataConditionAllOfParams `form:"all_of" json:"all_of"`
}

// Criteria determining which rates the multiplier applies to.
type V2BillingContractPricingOverrideMultiplierCriterionParams struct {
	// Filter by billable item IDs.
	BillableItemIDs []*string `form:"billable_item_ids" json:"billable_item_ids"`
	// Filter by billable item lookup keys.
	BillableItemLookupKeys []*string `form:"billable_item_lookup_keys" json:"billable_item_lookup_keys"`
	// Filter by billable item type.
	BillableItemTypes []*string `form:"billable_item_types" json:"billable_item_types"`
	// Filter by metadata conditions.
	MetadataConditions []*V2BillingContractPricingOverrideMultiplierCriterionMetadataConditionParams `form:"metadata_conditions" json:"metadata_conditions"`
	// Filter by rate card IDs. Only applicable for `multiplier` overrides.
	RateCardIDs []*string `form:"rate_card_ids" json:"rate_card_ids"`
	// Whether to include or exclude items matching these criteria.
	Type *string `form:"type" json:"type"`
}

// Parameters for a multiplier override. Required if `type` is `multiplier`.
type V2BillingContractPricingOverrideMultiplierParams struct {
	// Criteria determining which rates the multiplier applies to.
	Criteria []*V2BillingContractPricingOverrideMultiplierCriterionParams `form:"criteria" json:"criteria,omitempty"`
	// The multiplier factor, represented as a decimal string. e.g. "0.8" for a 20% reduction.
	Factor *string `form:"factor" json:"factor"`
}

// When the pricing override starts.
type V2BillingContractPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// A list of pricing overrides to create with the contract.
type V2BillingContractPricingOverrideParams struct {
	// When the pricing override ends.
	EndsAt *V2BillingContractPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at"`
	// A user-provided lookup key to reference this pricing override.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Parameters for a multiplier override. Required if `type` is `multiplier`.
	Multiplier *V2BillingContractPricingOverrideMultiplierParams `form:"multiplier" json:"multiplier,omitempty"`
	// The priority of this override relative to others. The highest priority is 0 and the lowest is 100.
	Priority *int64 `form:"priority" json:"priority"`
	// When the pricing override starts.
	StartsAt *V2BillingContractPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at"`
	// The type of pricing override.
	Type *string `form:"type" json:"type"`
}

// Create a Contract object.
type V2BillingContractParams struct {
	Params `form:"*"`
	// The billing cycle anchor for the contract. If not provided, defaults to the pricing line start time.
	// It is only at the top-level of the contract with no option to override at the pricing line level.
	BillingCycleAnchor *V2BillingContractBillingCycleAnchorParams `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// The billing settings for the contract.
	BillingSettings *V2BillingContractBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// A unique user-provided contract number e.g. C-2026-0001.
	ContractNumber *string `form:"contract_number" json:"contract_number,omitempty"`
	// Currency of the contract.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// A list of one-time fees to create with the contract. Each fee is billed as individual invoice items per its bill_schedule.
	OneTimeFees []*V2BillingContractOneTimeFeeParams `form:"one_time_fees" json:"one_time_fees,omitempty"`
	// Pricing line actions to apply.
	PricingLineActions []*V2BillingContractPricingLineActionParams `form:"pricing_line_actions" json:"pricing_line_actions,omitempty"`
	// A list of pricing lines to create with the contract.
	PricingLines []*V2BillingContractPricingLineParams `form:"pricing_lines" json:"pricing_lines,omitempty"`
	// Pricing override actions to apply.
	PricingOverrideActions []*V2BillingContractPricingOverrideActionParams `form:"pricing_override_actions" json:"pricing_override_actions,omitempty"`
	// A list of pricing overrides to create with the contract.
	PricingOverrides []*V2BillingContractPricingOverrideParams `form:"pricing_overrides" json:"pricing_overrides,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The end time for the pricing line.
type V2BillingContractPricingLineActionAddEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// When the override ends. Defaults to the pricing line's end if not specified.
type V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Each element represents a pricing tier.
type V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideOverwritePriceTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Parameters for the overwrite_price override. Required if `type` is `overwrite_price`.
type V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideOverwritePriceParams struct {
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideOverwritePriceTierParams `form:"tiers" json:"tiers,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// When the override starts. Defaults to the pricing line's start if not specified.
type V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Pricing overrides embedded directly on this pricing line.
type V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideParams struct {
	// When the override ends. Defaults to the pricing line's end if not specified.
	EndsAt *V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// A user-provided lookup key to reference this override.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Parameters for the overwrite_price override. Required if `type` is `overwrite_price`.
	OverwritePrice *V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideOverwritePriceParams `form:"overwrite_price" json:"overwrite_price,omitempty"`
	// The priority of this override relative to others. 0 is highest, 100 is lowest. Defaults to 50.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
	// When the override starts. Defaults to the pricing line's start if not specified.
	StartsAt *V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
	// The type of override. Currently only `overwrite_price` is supported.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// When this quantity change takes effect.
type V2BillingContractPricingLineActionAddPricingPriceDetailsQuantityChangeEffectiveAtParams struct {
	// The timestamp for the effective at.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the effective at.
	Type *string `form:"type" json:"type"`
}

// Quantity changes for the pricing line. For now, at most one entry is allowed.
// A quantity change clears all future quantity changes on this pricing line.
type V2BillingContractPricingLineActionAddPricingPriceDetailsQuantityChangeParams struct {
	// When this quantity change takes effect.
	EffectiveAt *V2BillingContractPricingLineActionAddPricingPriceDetailsQuantityChangeEffectiveAtParams `form:"effective_at" json:"effective_at"`
	// The quantity to set.
	Set *float64 `form:"set,high_precision" json:"set,string"`
}

// V1 price details. Required if `type` is `price`.
type V2BillingContractPricingLineActionAddPricingPriceDetailsParams struct {
	// The ID of the V1 price.
	Price *string `form:"price" json:"price"`
	// Pricing overrides embedded directly on this pricing line.
	PricingOverrides []*V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideParams `form:"pricing_overrides" json:"pricing_overrides,omitempty"`
	// Quantity changes for the pricing line. For now, at most one entry is allowed.
	// A quantity change clears all future quantity changes on this pricing line.
	QuantityChanges []*V2BillingContractPricingLineActionAddPricingPriceDetailsQuantityChangeParams `form:"quantity_changes" json:"quantity_changes,omitempty"`
}

// The pricing configuration for the pricing line.
type V2BillingContractPricingLineActionAddPricingParams struct {
	// V1 price details. Required if `type` is `price`.
	PriceDetails *V2BillingContractPricingLineActionAddPricingPriceDetailsParams `form:"price_details" json:"price_details,omitempty"`
	// The type of pricing.
	Type *string `form:"type" json:"type"`
}

// The start time for the pricing line.
type V2BillingContractPricingLineActionAddStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for adding a pricing line.
type V2BillingContractPricingLineActionAddParams struct {
	// The end time for the pricing line.
	EndsAt *V2BillingContractPricingLineActionAddEndsAtParams `form:"ends_at" json:"ends_at"`
	// A lookup key for the pricing line.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Metadata for the pricing line.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The pricing configuration for the pricing line.
	Pricing *V2BillingContractPricingLineActionAddPricingParams `form:"pricing" json:"pricing"`
	// The start time for the pricing line.
	StartsAt *V2BillingContractPricingLineActionAddStartsAtParams `form:"starts_at" json:"starts_at"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractPricingLineActionAddParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Parameters for removing a pricing line.
type V2BillingContractPricingLineActionRemoveParams struct {
	// The ID of the pricing line to remove.
	ID *string `form:"id" json:"id"`
}

// The updated end time for the pricing line.
type V2BillingContractPricingLineActionUpdateEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// The end time for the override.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// Each element represents a pricing tier.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddOverwritePriceTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Parameters for an overwrite_price override. Required if `type` is `overwrite_price`.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddOverwritePriceParams struct {
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddOverwritePriceTierParams `form:"tiers" json:"tiers,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// The start time for the override.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for adding a pricing line override.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddParams struct {
	// The end time for the override.
	EndsAt *V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddEndsAtParams `form:"ends_at" json:"ends_at"`
	// A lookup key for the override.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Parameters for an overwrite_price override. Required if `type` is `overwrite_price`.
	OverwritePrice *V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddOverwritePriceParams `form:"overwrite_price" json:"overwrite_price,omitempty"`
	// The priority of this override relative to others. 0 is highest, 100 is lowest. Defaults to 50.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
	// The start time for the override.
	StartsAt *V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddStartsAtParams `form:"starts_at" json:"starts_at"`
	// The type of override to add.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Parameters for removing a pricing line override.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionRemoveParams struct {
	// The ID of the pricing line override to remove.
	ID *string `form:"id" json:"id,omitempty"`
	// A lookup key for the override to remove.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
}

// The updated end time for the override.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// The updated start time for the override.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for updating a pricing line override.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateParams struct {
	// The updated end time for the override.
	EndsAt *V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// The ID of the pricing line override to update.
	ID *string `form:"id" json:"id,omitempty"`
	// A lookup key for the override to update.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The updated start time for the override.
	StartsAt *V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Pricing override actions to apply to the overrides embedded on this pricing line.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionParams struct {
	// Parameters for adding a pricing line override.
	Add *V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddParams `form:"add" json:"add,omitempty"`
	// Parameters for removing a pricing line override.
	Remove *V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionRemoveParams `form:"remove" json:"remove,omitempty"`
	// The type of pricing line override action.
	Type *string `form:"type" json:"type"`
	// Parameters for updating a pricing line override.
	Update *V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateParams `form:"update" json:"update,omitempty"`
}

// When this quantity change takes effect.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsQuantityChangeEffectiveAtParams struct {
	// The timestamp for the effective at.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the effective at.
	Type *string `form:"type" json:"type"`
}

// Quantity changes for the pricing line. Setting this clears all future quantity changes.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsQuantityChangeParams struct {
	// When this quantity change takes effect.
	EffectiveAt *V2BillingContractPricingLineActionUpdatePricingPriceDetailsQuantityChangeEffectiveAtParams `form:"effective_at" json:"effective_at"`
	// The quantity to set.
	Set *float64 `form:"set,high_precision" json:"set,string"`
}

// V1 price details. Present when the pricing line type is `price`.
type V2BillingContractPricingLineActionUpdatePricingPriceDetailsParams struct {
	// Pricing override actions to apply to the overrides embedded on this pricing line.
	PricingOverrideActions []*V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionParams `form:"pricing_override_actions" json:"pricing_override_actions,omitempty"`
	// Quantity changes for the pricing line. Setting this clears all future quantity changes.
	QuantityChanges []*V2BillingContractPricingLineActionUpdatePricingPriceDetailsQuantityChangeParams `form:"quantity_changes" json:"quantity_changes,omitempty"`
}

// Pricing updates for the pricing line (quantity changes and pricing override actions).
type V2BillingContractPricingLineActionUpdatePricingParams struct {
	// V1 price details. Present when the pricing line type is `price`.
	PriceDetails *V2BillingContractPricingLineActionUpdatePricingPriceDetailsParams `form:"price_details" json:"price_details,omitempty"`
}

// The updated start time for the pricing line.
type V2BillingContractPricingLineActionUpdateStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for updating a pricing line.
type V2BillingContractPricingLineActionUpdateParams struct {
	// The updated end time for the pricing line.
	EndsAt *V2BillingContractPricingLineActionUpdateEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// The ID of the pricing line.
	ID *string `form:"id" json:"id"`
	// Pricing updates for the pricing line (quantity changes and pricing override actions).
	Pricing *V2BillingContractPricingLineActionUpdatePricingParams `form:"pricing" json:"pricing,omitempty"`
	// The updated start time for the pricing line.
	StartsAt *V2BillingContractPricingLineActionUpdateStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
}

// Pricing line actions to apply.
type V2BillingContractPricingLineActionParams struct {
	// Parameters for adding a pricing line.
	Add *V2BillingContractPricingLineActionAddParams `form:"add" json:"add,omitempty"`
	// Parameters for removing a pricing line.
	Remove *V2BillingContractPricingLineActionRemoveParams `form:"remove" json:"remove,omitempty"`
	// The type of pricing line action.
	Type *string `form:"type" json:"type"`
	// Parameters for updating a pricing line.
	Update *V2BillingContractPricingLineActionUpdateParams `form:"update" json:"update,omitempty"`
}

// The end time for the pricing override.
type V2BillingContractPricingOverrideActionAddEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// All of these key-value conditions must match.
type V2BillingContractPricingOverrideActionAddMultiplierCriterionMetadataConditionAllOfParams struct {
	// The metadata key.
	Key *string `form:"key" json:"key"`
	// The metadata value.
	Value *string `form:"value" json:"value"`
}

// Filter by metadata conditions.
type V2BillingContractPricingOverrideActionAddMultiplierCriterionMetadataConditionParams struct {
	// All of these key-value conditions must match.
	AllOf []*V2BillingContractPricingOverrideActionAddMultiplierCriterionMetadataConditionAllOfParams `form:"all_of" json:"all_of"`
}

// Criteria determining which rates the multiplier applies to.
type V2BillingContractPricingOverrideActionAddMultiplierCriterionParams struct {
	// Filter by billable item IDs.
	BillableItemIDs []*string `form:"billable_item_ids" json:"billable_item_ids"`
	// Filter by billable item lookup keys.
	BillableItemLookupKeys []*string `form:"billable_item_lookup_keys" json:"billable_item_lookup_keys"`
	// Filter by billable item type.
	BillableItemTypes []*string `form:"billable_item_types" json:"billable_item_types"`
	// Filter by metadata conditions.
	MetadataConditions []*V2BillingContractPricingOverrideActionAddMultiplierCriterionMetadataConditionParams `form:"metadata_conditions" json:"metadata_conditions"`
	// Filter by rate card IDs. Only applicable for `multiplier` overrides.
	RateCardIDs []*string `form:"rate_card_ids" json:"rate_card_ids"`
	// Whether to include or exclude items matching these criteria.
	Type *string `form:"type" json:"type"`
}

// A multiplier override to add.
type V2BillingContractPricingOverrideActionAddMultiplierParams struct {
	// Criteria determining which rates the multiplier applies to.
	Criteria []*V2BillingContractPricingOverrideActionAddMultiplierCriterionParams `form:"criteria" json:"criteria"`
	// The multiplier factor, represented as a decimal string. e.g. "0.8" for a 20% reduction.
	Factor *string `form:"factor" json:"factor"`
}

// Each element represents a pricing tier.
type V2BillingContractPricingOverrideActionAddOverwritePriceTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// An overwrite price override to add.
type V2BillingContractPricingOverrideActionAddOverwritePriceParams struct {
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractPricingOverrideActionAddOverwritePriceTierParams `form:"tiers" json:"tiers"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// The start time for the pricing override.
type V2BillingContractPricingOverrideActionAddStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for adding a pricing override.
type V2BillingContractPricingOverrideActionAddParams struct {
	// The end time for the pricing override.
	EndsAt *V2BillingContractPricingOverrideActionAddEndsAtParams `form:"ends_at" json:"ends_at"`
	// A lookup key for the pricing override.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// A multiplier override to add.
	Multiplier *V2BillingContractPricingOverrideActionAddMultiplierParams `form:"multiplier" json:"multiplier,omitempty"`
	// An overwrite price override to add.
	OverwritePrice *V2BillingContractPricingOverrideActionAddOverwritePriceParams `form:"overwrite_price" json:"overwrite_price,omitempty"`
	// The priority for the pricing override. The highest priority is 0 and the lowest is 100.
	Priority *int64 `form:"priority" json:"priority"`
	// The start time for the pricing override.
	StartsAt *V2BillingContractPricingOverrideActionAddStartsAtParams `form:"starts_at" json:"starts_at"`
	// The type of pricing override to add.
	Type *string `form:"type" json:"type"`
}

// Parameters for removing a pricing override.
type V2BillingContractPricingOverrideActionRemoveParams struct {
	// The ID of the pricing override to remove.
	ID *string `form:"id" json:"id"`
}

// The updated end time for the pricing override.
type V2BillingContractPricingOverrideActionUpdateEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// The updated start time for the pricing override.
type V2BillingContractPricingOverrideActionUpdateStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for updating a pricing override.
type V2BillingContractPricingOverrideActionUpdateParams struct {
	// The updated end time for the pricing override.
	EndsAt *V2BillingContractPricingOverrideActionUpdateEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// The ID of the pricing override.
	ID *string `form:"id" json:"id"`
	// The updated start time for the pricing override.
	StartsAt *V2BillingContractPricingOverrideActionUpdateStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
}

// Pricing override actions to apply.
type V2BillingContractPricingOverrideActionParams struct {
	// Parameters for adding a pricing override.
	Add *V2BillingContractPricingOverrideActionAddParams `form:"add" json:"add,omitempty"`
	// Parameters for removing a pricing override.
	Remove *V2BillingContractPricingOverrideActionRemoveParams `form:"remove" json:"remove,omitempty"`
	// The type of pricing override action.
	Type *string `form:"type" json:"type"`
	// Parameters for updating a pricing override.
	Update *V2BillingContractPricingOverrideActionUpdateParams `form:"update" json:"update,omitempty"`
}

// Activate a Draft Contract object by ID.
type V2BillingContractActivateParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// Per-pricing-line proration behavior overrides. Falls back to `proration_behavior` if
// not specified for a given line.
type V2BillingContractCancelCancelPricingLineParams struct {
	// The ID of the pricing line.
	ID *string `form:"id" json:"id,omitempty"`
	// The lookup key of the pricing line.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Proration behavior scoped to this pricing line. If not provided, falls back to the
	// top-level `proration_behavior` on the cancel request. Defaults to `prorated`.
	ProrationBehavior *string `form:"proration_behavior" json:"proration_behavior,omitempty"`
}

// Cancel a Contract object by ID.
type V2BillingContractCancelParams struct {
	Params `form:"*"`
	// Per-pricing-line proration behavior overrides. Falls back to `proration_behavior` if
	// not specified for a given line.
	CancelPricingLines []*V2BillingContractCancelCancelPricingLineParams `form:"cancel_pricing_lines" json:"cancel_pricing_lines,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Top-level proration behavior for the cancellation. Defaults to `prorated` if not set.
	ProrationBehavior *string `form:"proration_behavior" json:"proration_behavior,omitempty"`
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
type V2BillingContractCreateBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxParams struct {
	// The type of tax calculation.
	Type *string `form:"type" json:"type"`
}

// Calculation settings.
type V2BillingContractCreateBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationParams struct {
	// Tax calculation settings.
	Tax *V2BillingContractCreateBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxParams `form:"tax" json:"tax,omitempty"`
}

// The number of time units before the invoice is past due.
type V2BillingContractCreateBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueParams struct {
	// The interval unit.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Invoice settings.
type V2BillingContractCreateBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceParams struct {
	// The number of time units before the invoice is past due.
	TimeUntilDue *V2BillingContractCreateBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueParams `form:"time_until_due" json:"time_until_due,omitempty"`
}

// The bill settings details.
type V2BillingContractCreateBillingSettingsContractBillingDetailsBillSettingsDetailsParams struct {
	// Calculation settings.
	Calculation *V2BillingContractCreateBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationParams `form:"calculation" json:"calculation,omitempty"`
	// Invoice settings.
	Invoice *V2BillingContractCreateBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// The billing profile details.
type V2BillingContractCreateBillingSettingsContractBillingDetailsBillingProfileDetailsParams struct {
	// The customer who pays for the contract invoice.
	Customer *string `form:"customer" json:"customer"`
	// The default payment method for the contract.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
}

// The collection settings details.
type V2BillingContractCreateBillingSettingsContractBillingDetailsCollectionSettingsDetailsParams struct {
	// The collection method.
	CollectionMethod *string `form:"collection_method" json:"collection_method"`
	// The payment method configuration.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
}

// Billing settings details for the contract.
type V2BillingContractCreateBillingSettingsContractBillingDetailsParams struct {
	// The billing profile details.
	BillingProfileDetails *V2BillingContractCreateBillingSettingsContractBillingDetailsBillingProfileDetailsParams `form:"billing_profile_details" json:"billing_profile_details"`
	// The bill settings details.
	BillSettingsDetails *V2BillingContractCreateBillingSettingsContractBillingDetailsBillSettingsDetailsParams `form:"bill_settings_details" json:"bill_settings_details,omitempty"`
	// The collection settings details.
	CollectionSettingsDetails *V2BillingContractCreateBillingSettingsContractBillingDetailsCollectionSettingsDetailsParams `form:"collection_settings_details" json:"collection_settings_details"`
}

// The billing settings for the contract.
type V2BillingContractCreateBillingSettingsParams struct {
	// Billing settings details for the contract.
	ContractBillingDetails *V2BillingContractCreateBillingSettingsContractBillingDetailsParams `form:"contract_billing_details" json:"contract_billing_details,omitempty"`
}

// When this entry should be billed.
type V2BillingContractCreateOneTimeFeeBillScheduleBillAtParams struct {
	// The datetime at which the entry should be billed. Required if `type` is `datetime`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the bill_at.
	Type *string `form:"type" json:"type"`
}

// The bill schedule for the fee. Each entry produces an individual invoice item billed at `bill_at`.
type V2BillingContractCreateOneTimeFeeBillScheduleParams struct {
	// When this entry should be billed.
	BillAt *V2BillingContractCreateOneTimeFeeBillScheduleBillAtParams `form:"bill_at" json:"bill_at"`
	// The amount to bill for this entry, in the smallest currency unit.
	Value *int64 `form:"value" json:"value,string"`
}

// Details for a product billable target. Required when `billable_item_type` is `product`.
type V2BillingContractCreateOneTimeFeeProductDetailsParams struct {
	// The ID of the v1 Product.
	Product *string `form:"product" json:"product"`
}

// A list of one-time fees to create with the contract. Each fee is billed as individual invoice items per its bill_schedule.
type V2BillingContractCreateOneTimeFeeParams struct {
	// The type of billable item that this fee references.
	BillableItemType *string `form:"billable_item_type" json:"billable_item_type"`
	// The bill schedule for the fee. Each entry produces an individual invoice item billed at `bill_at`.
	BillSchedule []*V2BillingContractCreateOneTimeFeeBillScheduleParams `form:"bill_schedule" json:"bill_schedule"`
	// Details for a product billable target. Required when `billable_item_type` is `product`.
	ProductDetails *V2BillingContractCreateOneTimeFeeProductDetailsParams `form:"product_details" json:"product_details,omitempty"`
}

// When the pricing line ends.
type V2BillingContractCreatePricingLineEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// When the override ends. Defaults to the pricing line's end if not specified.
type V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Each element represents a pricing tier.
type V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideOverwritePriceTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Parameters for the overwrite_price override. Required if `type` is `overwrite_price`.
type V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideOverwritePriceParams struct {
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideOverwritePriceTierParams `form:"tiers" json:"tiers,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// When the override starts. Defaults to the pricing line's start if not specified.
type V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Pricing overrides embedded directly on this pricing line.
type V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideParams struct {
	// When the override ends. Defaults to the pricing line's end if not specified.
	EndsAt *V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// A user-provided lookup key to reference this override.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Parameters for the overwrite_price override. Required if `type` is `overwrite_price`.
	OverwritePrice *V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideOverwritePriceParams `form:"overwrite_price" json:"overwrite_price,omitempty"`
	// The priority of this override relative to others. 0 is highest, 100 is lowest. Defaults to 50.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
	// When the override starts. Defaults to the pricing line's start if not specified.
	StartsAt *V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
	// The type of override. Currently only `overwrite_price` is supported.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// When this quantity change takes effect.
type V2BillingContractCreatePricingLinePricingPriceDetailsQuantityChangeEffectiveAtParams struct {
	// The timestamp for the effective at.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the effective at.
	Type *string `form:"type" json:"type"`
}

// Quantity changes for the pricing line. For now, at most one entry is allowed.
// A quantity change clears all future quantity changes on this pricing line.
type V2BillingContractCreatePricingLinePricingPriceDetailsQuantityChangeParams struct {
	// When this quantity change takes effect.
	EffectiveAt *V2BillingContractCreatePricingLinePricingPriceDetailsQuantityChangeEffectiveAtParams `form:"effective_at" json:"effective_at"`
	// The quantity to set.
	Set *float64 `form:"set,high_precision" json:"set,string"`
}

// V1 price details. Required if `type` is `price`.
type V2BillingContractCreatePricingLinePricingPriceDetailsParams struct {
	// The ID of the V1 price.
	Price *string `form:"price" json:"price"`
	// Pricing overrides embedded directly on this pricing line.
	PricingOverrides []*V2BillingContractCreatePricingLinePricingPriceDetailsPricingOverrideParams `form:"pricing_overrides" json:"pricing_overrides,omitempty"`
	// Quantity changes for the pricing line. For now, at most one entry is allowed.
	// A quantity change clears all future quantity changes on this pricing line.
	QuantityChanges []*V2BillingContractCreatePricingLinePricingPriceDetailsQuantityChangeParams `form:"quantity_changes" json:"quantity_changes,omitempty"`
}

// The pricing configuration for the pricing line.
type V2BillingContractCreatePricingLinePricingParams struct {
	// V1 price details. Required if `type` is `price`.
	PriceDetails *V2BillingContractCreatePricingLinePricingPriceDetailsParams `form:"price_details" json:"price_details,omitempty"`
	// The type of pricing.
	Type *string `form:"type" json:"type"`
}

// When the pricing line starts.
type V2BillingContractCreatePricingLineStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// A list of pricing lines to create with the contract.
type V2BillingContractCreatePricingLineParams struct {
	// When the pricing line ends.
	EndsAt *V2BillingContractCreatePricingLineEndsAtParams `form:"ends_at" json:"ends_at"`
	// A user-provided lookup key to reference this pricing line.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The pricing configuration for the pricing line.
	Pricing *V2BillingContractCreatePricingLinePricingParams `form:"pricing" json:"pricing"`
	// When the pricing line starts.
	StartsAt *V2BillingContractCreatePricingLineStartsAtParams `form:"starts_at" json:"starts_at"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractCreatePricingLineParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// When the pricing override ends.
type V2BillingContractCreatePricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// All of these key-value conditions must match.
type V2BillingContractCreatePricingOverrideMultiplierCriterionMetadataConditionAllOfParams struct {
	// The metadata key.
	Key *string `form:"key" json:"key"`
	// The metadata value.
	Value *string `form:"value" json:"value"`
}

// Filter by metadata conditions.
type V2BillingContractCreatePricingOverrideMultiplierCriterionMetadataConditionParams struct {
	// All of these key-value conditions must match.
	AllOf []*V2BillingContractCreatePricingOverrideMultiplierCriterionMetadataConditionAllOfParams `form:"all_of" json:"all_of"`
}

// Criteria determining which rates the multiplier applies to.
type V2BillingContractCreatePricingOverrideMultiplierCriterionParams struct {
	// Filter by billable item IDs.
	BillableItemIDs []*string `form:"billable_item_ids" json:"billable_item_ids"`
	// Filter by billable item lookup keys.
	BillableItemLookupKeys []*string `form:"billable_item_lookup_keys" json:"billable_item_lookup_keys"`
	// Filter by billable item type.
	BillableItemTypes []*string `form:"billable_item_types" json:"billable_item_types"`
	// Filter by metadata conditions.
	MetadataConditions []*V2BillingContractCreatePricingOverrideMultiplierCriterionMetadataConditionParams `form:"metadata_conditions" json:"metadata_conditions"`
	// Filter by rate card IDs. Only applicable for `multiplier` overrides.
	RateCardIDs []*string `form:"rate_card_ids" json:"rate_card_ids"`
	// Whether to include or exclude items matching these criteria.
	Type *string `form:"type" json:"type"`
}

// Parameters for a multiplier override. Required if `type` is `multiplier`.
type V2BillingContractCreatePricingOverrideMultiplierParams struct {
	// Criteria determining which rates the multiplier applies to.
	Criteria []*V2BillingContractCreatePricingOverrideMultiplierCriterionParams `form:"criteria" json:"criteria,omitempty"`
	// The multiplier factor, represented as a decimal string. e.g. "0.8" for a 20% reduction.
	Factor *string `form:"factor" json:"factor"`
}

// When the pricing override starts.
type V2BillingContractCreatePricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// A list of pricing overrides to create with the contract.
type V2BillingContractCreatePricingOverrideParams struct {
	// When the pricing override ends.
	EndsAt *V2BillingContractCreatePricingOverrideEndsAtParams `form:"ends_at" json:"ends_at"`
	// A user-provided lookup key to reference this pricing override.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Parameters for a multiplier override. Required if `type` is `multiplier`.
	Multiplier *V2BillingContractCreatePricingOverrideMultiplierParams `form:"multiplier" json:"multiplier,omitempty"`
	// The priority of this override relative to others. The highest priority is 0 and the lowest is 100.
	Priority *int64 `form:"priority" json:"priority"`
	// When the pricing override starts.
	StartsAt *V2BillingContractCreatePricingOverrideStartsAtParams `form:"starts_at" json:"starts_at"`
	// The type of pricing override.
	Type *string `form:"type" json:"type"`
}

// Create a Contract object.
type V2BillingContractCreateParams struct {
	Params `form:"*"`
	// The billing cycle anchor for the contract. If not provided, defaults to the pricing line start time.
	// It is only at the top-level of the contract with no option to override at the pricing line level.
	BillingCycleAnchor *V2BillingContractCreateBillingCycleAnchorParams `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// The billing settings for the contract.
	BillingSettings *V2BillingContractCreateBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// A unique user-provided contract number e.g. C-2026-0001.
	ContractNumber *string `form:"contract_number" json:"contract_number"`
	// Currency of the contract.
	Currency *string `form:"currency" json:"currency"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// A list of one-time fees to create with the contract. Each fee is billed as individual invoice items per its bill_schedule.
	OneTimeFees []*V2BillingContractCreateOneTimeFeeParams `form:"one_time_fees" json:"one_time_fees,omitempty"`
	// A list of pricing lines to create with the contract.
	PricingLines []*V2BillingContractCreatePricingLineParams `form:"pricing_lines" json:"pricing_lines"`
	// A list of pricing overrides to create with the contract.
	PricingOverrides []*V2BillingContractCreatePricingOverrideParams `form:"pricing_overrides" json:"pricing_overrides,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Delete a draft Contract object by ID.
type V2BillingContractDeleteParams struct {
	Params `form:"*"`
}

// Retrieve a Contract object by ID.
type V2BillingContractRetrieveParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// The end time for the pricing line.
type V2BillingContractUpdatePricingLineActionAddEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// When the override ends. Defaults to the pricing line's end if not specified.
type V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideEndsAtParams struct {
	// The timestamp when the item ends. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the ends_at.
	Type *string `form:"type" json:"type"`
}

// Each element represents a pricing tier.
type V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideOverwritePriceTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Parameters for the overwrite_price override. Required if `type` is `overwrite_price`.
type V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideOverwritePriceParams struct {
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideOverwritePriceTierParams `form:"tiers" json:"tiers,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// When the override starts. Defaults to the pricing line's start if not specified.
type V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideStartsAtParams struct {
	// The timestamp when the item starts. Required if `type` is `timestamp`.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the starts_at.
	Type *string `form:"type" json:"type"`
}

// Pricing overrides embedded directly on this pricing line.
type V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideParams struct {
	// When the override ends. Defaults to the pricing line's end if not specified.
	EndsAt *V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// A user-provided lookup key to reference this override.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Parameters for the overwrite_price override. Required if `type` is `overwrite_price`.
	OverwritePrice *V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideOverwritePriceParams `form:"overwrite_price" json:"overwrite_price,omitempty"`
	// The priority of this override relative to others. 0 is highest, 100 is lowest. Defaults to 50.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
	// When the override starts. Defaults to the pricing line's start if not specified.
	StartsAt *V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
	// The type of override. Currently only `overwrite_price` is supported.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// When this quantity change takes effect.
type V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsQuantityChangeEffectiveAtParams struct {
	// The timestamp for the effective at.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the effective at.
	Type *string `form:"type" json:"type"`
}

// Quantity changes for the pricing line. For now, at most one entry is allowed.
// A quantity change clears all future quantity changes on this pricing line.
type V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsQuantityChangeParams struct {
	// When this quantity change takes effect.
	EffectiveAt *V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsQuantityChangeEffectiveAtParams `form:"effective_at" json:"effective_at"`
	// The quantity to set.
	Set *float64 `form:"set,high_precision" json:"set,string"`
}

// V1 price details. Required if `type` is `price`.
type V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsParams struct {
	// The ID of the V1 price.
	Price *string `form:"price" json:"price"`
	// Pricing overrides embedded directly on this pricing line.
	PricingOverrides []*V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsPricingOverrideParams `form:"pricing_overrides" json:"pricing_overrides,omitempty"`
	// Quantity changes for the pricing line. For now, at most one entry is allowed.
	// A quantity change clears all future quantity changes on this pricing line.
	QuantityChanges []*V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsQuantityChangeParams `form:"quantity_changes" json:"quantity_changes,omitempty"`
}

// The pricing configuration for the pricing line.
type V2BillingContractUpdatePricingLineActionAddPricingParams struct {
	// V1 price details. Required if `type` is `price`.
	PriceDetails *V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsParams `form:"price_details" json:"price_details,omitempty"`
	// The type of pricing.
	Type *string `form:"type" json:"type"`
}

// The start time for the pricing line.
type V2BillingContractUpdatePricingLineActionAddStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for adding a pricing line.
type V2BillingContractUpdatePricingLineActionAddParams struct {
	// The end time for the pricing line.
	EndsAt *V2BillingContractUpdatePricingLineActionAddEndsAtParams `form:"ends_at" json:"ends_at"`
	// A lookup key for the pricing line.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Metadata for the pricing line.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The pricing configuration for the pricing line.
	Pricing *V2BillingContractUpdatePricingLineActionAddPricingParams `form:"pricing" json:"pricing"`
	// The start time for the pricing line.
	StartsAt *V2BillingContractUpdatePricingLineActionAddStartsAtParams `form:"starts_at" json:"starts_at"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractUpdatePricingLineActionAddParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Parameters for removing a pricing line.
type V2BillingContractUpdatePricingLineActionRemoveParams struct {
	// The ID of the pricing line to remove.
	ID *string `form:"id" json:"id"`
}

// The updated end time for the pricing line.
type V2BillingContractUpdatePricingLineActionUpdateEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// The end time for the override.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// Each element represents a pricing tier.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddOverwritePriceTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Parameters for an overwrite_price override. Required if `type` is `overwrite_price`.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddOverwritePriceParams struct {
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddOverwritePriceTierParams `form:"tiers" json:"tiers,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// The start time for the override.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for adding a pricing line override.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddParams struct {
	// The end time for the override.
	EndsAt *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddEndsAtParams `form:"ends_at" json:"ends_at"`
	// A lookup key for the override.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Parameters for an overwrite_price override. Required if `type` is `overwrite_price`.
	OverwritePrice *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddOverwritePriceParams `form:"overwrite_price" json:"overwrite_price,omitempty"`
	// The priority of this override relative to others. 0 is highest, 100 is lowest. Defaults to 50.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
	// The start time for the override.
	StartsAt *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddStartsAtParams `form:"starts_at" json:"starts_at"`
	// The type of override to add.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Parameters for removing a pricing line override.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionRemoveParams struct {
	// The ID of the pricing line override to remove.
	ID *string `form:"id" json:"id,omitempty"`
	// A lookup key for the override to remove.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
}

// The updated end time for the override.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// The updated start time for the override.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for updating a pricing line override.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateParams struct {
	// The updated end time for the override.
	EndsAt *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// The ID of the pricing line override to update.
	ID *string `form:"id" json:"id,omitempty"`
	// A lookup key for the override to update.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The updated start time for the override.
	StartsAt *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Pricing override actions to apply to the overrides embedded on this pricing line.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionParams struct {
	// Parameters for adding a pricing line override.
	Add *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddParams `form:"add" json:"add,omitempty"`
	// Parameters for removing a pricing line override.
	Remove *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionRemoveParams `form:"remove" json:"remove,omitempty"`
	// The type of pricing line override action.
	Type *string `form:"type" json:"type"`
	// Parameters for updating a pricing line override.
	Update *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateParams `form:"update" json:"update,omitempty"`
}

// When this quantity change takes effect.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsQuantityChangeEffectiveAtParams struct {
	// The timestamp for the effective at.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the effective at.
	Type *string `form:"type" json:"type"`
}

// Quantity changes for the pricing line. Setting this clears all future quantity changes.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsQuantityChangeParams struct {
	// When this quantity change takes effect.
	EffectiveAt *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsQuantityChangeEffectiveAtParams `form:"effective_at" json:"effective_at"`
	// The quantity to set.
	Set *float64 `form:"set,high_precision" json:"set,string"`
}

// V1 price details. Present when the pricing line type is `price`.
type V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsParams struct {
	// Pricing override actions to apply to the overrides embedded on this pricing line.
	PricingOverrideActions []*V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsPricingOverrideActionParams `form:"pricing_override_actions" json:"pricing_override_actions,omitempty"`
	// Quantity changes for the pricing line. Setting this clears all future quantity changes.
	QuantityChanges []*V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsQuantityChangeParams `form:"quantity_changes" json:"quantity_changes,omitempty"`
}

// Pricing updates for the pricing line (quantity changes and pricing override actions).
type V2BillingContractUpdatePricingLineActionUpdatePricingParams struct {
	// V1 price details. Present when the pricing line type is `price`.
	PriceDetails *V2BillingContractUpdatePricingLineActionUpdatePricingPriceDetailsParams `form:"price_details" json:"price_details,omitempty"`
}

// The updated start time for the pricing line.
type V2BillingContractUpdatePricingLineActionUpdateStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for updating a pricing line.
type V2BillingContractUpdatePricingLineActionUpdateParams struct {
	// The updated end time for the pricing line.
	EndsAt *V2BillingContractUpdatePricingLineActionUpdateEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// The ID of the pricing line.
	ID *string `form:"id" json:"id"`
	// Pricing updates for the pricing line (quantity changes and pricing override actions).
	Pricing *V2BillingContractUpdatePricingLineActionUpdatePricingParams `form:"pricing" json:"pricing,omitempty"`
	// The updated start time for the pricing line.
	StartsAt *V2BillingContractUpdatePricingLineActionUpdateStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
}

// Pricing line actions to apply.
type V2BillingContractUpdatePricingLineActionParams struct {
	// Parameters for adding a pricing line.
	Add *V2BillingContractUpdatePricingLineActionAddParams `form:"add" json:"add,omitempty"`
	// Parameters for removing a pricing line.
	Remove *V2BillingContractUpdatePricingLineActionRemoveParams `form:"remove" json:"remove,omitempty"`
	// The type of pricing line action.
	Type *string `form:"type" json:"type"`
	// Parameters for updating a pricing line.
	Update *V2BillingContractUpdatePricingLineActionUpdateParams `form:"update" json:"update,omitempty"`
}

// The end time for the pricing override.
type V2BillingContractUpdatePricingOverrideActionAddEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// All of these key-value conditions must match.
type V2BillingContractUpdatePricingOverrideActionAddMultiplierCriterionMetadataConditionAllOfParams struct {
	// The metadata key.
	Key *string `form:"key" json:"key"`
	// The metadata value.
	Value *string `form:"value" json:"value"`
}

// Filter by metadata conditions.
type V2BillingContractUpdatePricingOverrideActionAddMultiplierCriterionMetadataConditionParams struct {
	// All of these key-value conditions must match.
	AllOf []*V2BillingContractUpdatePricingOverrideActionAddMultiplierCriterionMetadataConditionAllOfParams `form:"all_of" json:"all_of"`
}

// Criteria determining which rates the multiplier applies to.
type V2BillingContractUpdatePricingOverrideActionAddMultiplierCriterionParams struct {
	// Filter by billable item IDs.
	BillableItemIDs []*string `form:"billable_item_ids" json:"billable_item_ids"`
	// Filter by billable item lookup keys.
	BillableItemLookupKeys []*string `form:"billable_item_lookup_keys" json:"billable_item_lookup_keys"`
	// Filter by billable item type.
	BillableItemTypes []*string `form:"billable_item_types" json:"billable_item_types"`
	// Filter by metadata conditions.
	MetadataConditions []*V2BillingContractUpdatePricingOverrideActionAddMultiplierCriterionMetadataConditionParams `form:"metadata_conditions" json:"metadata_conditions"`
	// Filter by rate card IDs. Only applicable for `multiplier` overrides.
	RateCardIDs []*string `form:"rate_card_ids" json:"rate_card_ids"`
	// Whether to include or exclude items matching these criteria.
	Type *string `form:"type" json:"type"`
}

// A multiplier override to add.
type V2BillingContractUpdatePricingOverrideActionAddMultiplierParams struct {
	// Criteria determining which rates the multiplier applies to.
	Criteria []*V2BillingContractUpdatePricingOverrideActionAddMultiplierCriterionParams `form:"criteria" json:"criteria"`
	// The multiplier factor, represented as a decimal string. e.g. "0.8" for a 20% reduction.
	Factor *string `form:"factor" json:"factor"`
}

// Each element represents a pricing tier.
type V2BillingContractUpdatePricingOverrideActionAddOverwritePriceTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// An overwrite price override to add.
type V2BillingContractUpdatePricingOverrideActionAddOverwritePriceParams struct {
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractUpdatePricingOverrideActionAddOverwritePriceTierParams `form:"tiers" json:"tiers"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// The start time for the pricing override.
type V2BillingContractUpdatePricingOverrideActionAddStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for adding a pricing override.
type V2BillingContractUpdatePricingOverrideActionAddParams struct {
	// The end time for the pricing override.
	EndsAt *V2BillingContractUpdatePricingOverrideActionAddEndsAtParams `form:"ends_at" json:"ends_at"`
	// A lookup key for the pricing override.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// A multiplier override to add.
	Multiplier *V2BillingContractUpdatePricingOverrideActionAddMultiplierParams `form:"multiplier" json:"multiplier,omitempty"`
	// An overwrite price override to add.
	OverwritePrice *V2BillingContractUpdatePricingOverrideActionAddOverwritePriceParams `form:"overwrite_price" json:"overwrite_price,omitempty"`
	// The priority for the pricing override. The highest priority is 0 and the lowest is 100.
	Priority *int64 `form:"priority" json:"priority"`
	// The start time for the pricing override.
	StartsAt *V2BillingContractUpdatePricingOverrideActionAddStartsAtParams `form:"starts_at" json:"starts_at"`
	// The type of pricing override to add.
	Type *string `form:"type" json:"type"`
}

// Parameters for removing a pricing override.
type V2BillingContractUpdatePricingOverrideActionRemoveParams struct {
	// The ID of the pricing override to remove.
	ID *string `form:"id" json:"id"`
}

// The updated end time for the pricing override.
type V2BillingContractUpdatePricingOverrideActionUpdateEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// The updated start time for the pricing override.
type V2BillingContractUpdatePricingOverrideActionUpdateStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of start time to apply.
	Type *string `form:"type" json:"type"`
}

// Parameters for updating a pricing override.
type V2BillingContractUpdatePricingOverrideActionUpdateParams struct {
	// The updated end time for the pricing override.
	EndsAt *V2BillingContractUpdatePricingOverrideActionUpdateEndsAtParams `form:"ends_at" json:"ends_at,omitempty"`
	// The ID of the pricing override.
	ID *string `form:"id" json:"id"`
	// The updated start time for the pricing override.
	StartsAt *V2BillingContractUpdatePricingOverrideActionUpdateStartsAtParams `form:"starts_at" json:"starts_at,omitempty"`
}

// Pricing override actions to apply.
type V2BillingContractUpdatePricingOverrideActionParams struct {
	// Parameters for adding a pricing override.
	Add *V2BillingContractUpdatePricingOverrideActionAddParams `form:"add" json:"add,omitempty"`
	// Parameters for removing a pricing override.
	Remove *V2BillingContractUpdatePricingOverrideActionRemoveParams `form:"remove" json:"remove,omitempty"`
	// The type of pricing override action.
	Type *string `form:"type" json:"type"`
	// Parameters for updating a pricing override.
	Update *V2BillingContractUpdatePricingOverrideActionUpdateParams `form:"update" json:"update,omitempty"`
}

// Update a Contract object by ID.
type V2BillingContractUpdateParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Pricing line actions to apply.
	PricingLineActions []*V2BillingContractUpdatePricingLineActionParams `form:"pricing_line_actions" json:"pricing_line_actions,omitempty"`
	// Pricing override actions to apply.
	PricingOverrideActions []*V2BillingContractUpdatePricingOverrideActionParams `form:"pricing_override_actions" json:"pricing_override_actions,omitempty"`
}
