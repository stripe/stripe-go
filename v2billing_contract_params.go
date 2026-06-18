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

// Timestamp to indicate when the contract line ends.
type V2BillingContractContractLineEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp"`
}

// Timestamp to indicate when the override ends.
type V2BillingContractContractLineOverrideEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp"`
}

// The amount of the credit grant.
type V2BillingContractContractLineOverrideServiceActionAddCreditGrantAmountParams struct {
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary *Amount `form:"monetary" json:"monetary,omitempty"`
	// The type of the credit grant amount.
	Type *string `form:"type" json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingContractContractLineOverrideServiceActionAddCreditGrantApplicabilityConfigScopeParams struct {
	// The billable items to apply the credit grant to.
	BillableItems []*string `form:"billable_items" json:"billable_items,omitempty"`
	// The price type that credit grants can apply to.
	PriceType *string `form:"price_type" json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingContractContractLineOverrideServiceActionAddCreditGrantApplicabilityConfigParams struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingContractContractLineOverrideServiceActionAddCreditGrantApplicabilityConfigScopeParams `form:"scope" json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingContractContractLineOverrideServiceActionAddCreditGrantExpiryConfigParams struct {
	// The type of the expiry configuration.
	Type *string `form:"type" json:"type"`
}

// Details for the credit grant. Required if `type` is `credit_grant`.
type V2BillingContractContractLineOverrideServiceActionAddCreditGrantParams struct {
	// The amount of the credit grant.
	Amount *V2BillingContractContractLineOverrideServiceActionAddCreditGrantAmountParams `form:"amount" json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingContractContractLineOverrideServiceActionAddCreditGrantApplicabilityConfigParams `form:"applicability_config" json:"applicability_config"`
	// The category of the credit grant.
	Category *string `form:"category" json:"category,omitempty"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingContractContractLineOverrideServiceActionAddCreditGrantExpiryConfigParams `form:"expiry_config" json:"expiry_config"`
	// A descriptive name.
	Name *string `form:"name" json:"name"`
	// The desired priority for applying this credit grant. The highest priority is 0 and the lowest is 100.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
}

// Parameters for adding a new service action.
type V2BillingContractContractLineOverrideServiceActionAddParams struct {
	// Details for the credit grant. Required if `type` is `credit_grant`.
	CreditGrant *V2BillingContractContractLineOverrideServiceActionAddCreditGrantParams `form:"credit_grant" json:"credit_grant,omitempty"`
	// The interval for assessing service.
	ServiceInterval *string `form:"service_interval" json:"service_interval"`
	// The length of the interval for assessing service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count"`
	// The type of the service action.
	Type *string `form:"type" json:"type"`
}

// The amount of the credit grant.
type V2BillingContractContractLineOverrideServiceActionReplaceCreditGrantAmountParams struct {
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary *Amount `form:"monetary" json:"monetary,omitempty"`
	// The type of the credit grant amount.
	Type *string `form:"type" json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingContractContractLineOverrideServiceActionReplaceCreditGrantApplicabilityConfigScopeParams struct {
	// The billable items to apply the credit grant to.
	BillableItems []*string `form:"billable_items" json:"billable_items,omitempty"`
	// The price type that credit grants can apply to.
	PriceType *string `form:"price_type" json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingContractContractLineOverrideServiceActionReplaceCreditGrantApplicabilityConfigParams struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingContractContractLineOverrideServiceActionReplaceCreditGrantApplicabilityConfigScopeParams `form:"scope" json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingContractContractLineOverrideServiceActionReplaceCreditGrantExpiryConfigParams struct {
	// The type of the expiry configuration.
	Type *string `form:"type" json:"type"`
}

// Details for the credit grant. Required if `type` is `credit_grant`.
type V2BillingContractContractLineOverrideServiceActionReplaceCreditGrantParams struct {
	// The amount of the credit grant.
	Amount *V2BillingContractContractLineOverrideServiceActionReplaceCreditGrantAmountParams `form:"amount" json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingContractContractLineOverrideServiceActionReplaceCreditGrantApplicabilityConfigParams `form:"applicability_config" json:"applicability_config"`
	// The category of the credit grant.
	Category *string `form:"category" json:"category,omitempty"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingContractContractLineOverrideServiceActionReplaceCreditGrantExpiryConfigParams `form:"expiry_config" json:"expiry_config"`
	// A descriptive name.
	Name *string `form:"name" json:"name"`
	// The desired priority for applying this credit grant. The highest priority is 0 and the lowest is 100.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
}

// Parameters for replacing an existing service action.
type V2BillingContractContractLineOverrideServiceActionReplaceParams struct {
	// Details for the credit grant. Required if `type` is `credit_grant`.
	CreditGrant *V2BillingContractContractLineOverrideServiceActionReplaceCreditGrantParams `form:"credit_grant" json:"credit_grant,omitempty"`
	// The ID of the service action to replace.
	ID *string `form:"id" json:"id,omitempty"`
	// The lookup key for the service action to replace.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// The interval for assessing service.
	ServiceInterval *string `form:"service_interval" json:"service_interval"`
	// The length of the interval for assessing service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count"`
	// The type of the service action.
	Type *string `form:"type" json:"type"`
}

// Service action override parameters. Required if `type` is `service_action`.
type V2BillingContractContractLineOverrideServiceActionParams struct {
	// Parameters for adding a new service action.
	Add *V2BillingContractContractLineOverrideServiceActionAddParams `form:"add" json:"add,omitempty"`
	// Parameters for replacing an existing service action.
	Replace *V2BillingContractContractLineOverrideServiceActionReplaceParams `form:"replace" json:"replace,omitempty"`
	// The type of service action override.
	Type *string `form:"type" json:"type"`
}

// Timestamp to indicate when the override starts.
type V2BillingContractContractLineOverrideStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp"`
}

// List of overrides. Later overrides in the list override earlier ones.
type V2BillingContractContractLineOverrideParams struct {
	// Timestamp to indicate when the override ends.
	EndsAt *V2BillingContractContractLineOverrideEndsAtParams `form:"ends_at" json:"ends_at"`
	// Service action override parameters. Required if `type` is `service_action`.
	ServiceAction *V2BillingContractContractLineOverrideServiceActionParams `form:"service_action" json:"service_action,omitempty"`
	// Timestamp to indicate when the override starts.
	StartsAt *V2BillingContractContractLineOverrideStartsAtParams `form:"starts_at" json:"starts_at"`
	// The type of the override.
	Type *string `form:"type" json:"type"`
}

// The pricing configuration for the contract line.
type V2BillingContractContractLinePricingParams struct{}

// Timestamp to indicate when the contract line starts.
type V2BillingContractContractLineStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp"`
}

// A list of contract lines to create with the contract.
type V2BillingContractContractLineParams struct {
	// Timestamp to indicate when the contract line ends.
	EndsAt *V2BillingContractContractLineEndsAtParams `form:"ends_at" json:"ends_at"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// List of overrides. Later overrides in the list override earlier ones.
	Overrides []*V2BillingContractContractLineOverrideParams `form:"overrides" json:"overrides"`
	// The pricing configuration for the contract line.
	Pricing *V2BillingContractContractLinePricingParams `form:"pricing" json:"pricing"`
	// Timestamp to indicate when the contract line starts.
	StartsAt *V2BillingContractContractLineStartsAtParams `form:"starts_at" json:"starts_at"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractContractLineParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The effective at for the license quantity action.
type V2BillingContractLicenseQuantityActionEffectiveAtParams struct {
	// The timestamp for the effective at.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the effective at.
	Type *string `form:"type" json:"type"`
}

// The set quantity for the license quantity action.
type V2BillingContractLicenseQuantityActionSetParams struct {
	// The quantity to set.
	Quantity *int64 `form:"quantity" json:"quantity"`
}

// A list of license quantity actions to create with the contract.
type V2BillingContractLicenseQuantityActionParams struct {
	// The effective at for the license quantity action.
	EffectiveAt *V2BillingContractLicenseQuantityActionEffectiveAtParams `form:"effective_at" json:"effective_at"`
	// The ID of the license pricing.
	LicensePricingID *string `form:"license_pricing_id" json:"license_pricing_id,omitempty"`
	// The lookup key for the license pricing.
	LicensePricingLookupKey *string `form:"license_pricing_lookup_key" json:"license_pricing_lookup_key,omitempty"`
	// The type of the license pricing.
	LicensePricingType *string `form:"license_pricing_type" json:"license_pricing_type"`
	// The pricing line ID for the license quantity action.
	PricingLine *string `form:"pricing_line" json:"pricing_line,omitempty"`
	// The pricing line lookup key for the license quantity action.
	PricingLineLookupKey *string `form:"pricing_line_lookup_key" json:"pricing_line_lookup_key,omitempty"`
	// The set quantity for the license quantity action.
	Set *V2BillingContractLicenseQuantityActionSetParams `form:"set" json:"set,omitempty"`
	// The type of the license quantity action.
	Type *string `form:"type" json:"type"`
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

// V1 price details. Required if `type` is `price`.
type V2BillingContractPricingLinePricingPriceDetailsParams struct {
	// The ID of the V1 price.
	Price *string `form:"price" json:"price"`
	// The quantity for the price. Only applicable for licensed prices.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
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

// Each element represents a pricing tier.
type V2BillingContractPricingOverrideOverwritePriceTierParams struct {
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
type V2BillingContractPricingOverrideOverwritePriceParams struct {
	// The ID of the V1 price to overwrite.
	Price *string `form:"price" json:"price"`
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractPricingOverrideOverwritePriceTierParams `form:"tiers" json:"tiers,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
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
	// Parameters for an overwrite_price override. Required if `type` is `overwrite_price`.
	OverwritePrice *V2BillingContractPricingOverrideOverwritePriceParams `form:"overwrite_price" json:"overwrite_price,omitempty"`
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
	// The billing settings for the contract.
	BillingSettings *V2BillingContractBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// A list of contract lines to create with the contract.
	ContractLines []*V2BillingContractContractLineParams `form:"contract_lines" json:"contract_lines,omitempty"`
	// A unique user-provided contract number e.g. C-2026-0001.
	ContractNumber *string `form:"contract_number" json:"contract_number,omitempty"`
	// Currency of the contract.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Schema-only: License quantity actions (implementation to follow).
	LicenseQuantityActions []*V2BillingContractLicenseQuantityActionParams `form:"license_quantity_actions" json:"license_quantity_actions,omitempty"`
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

// V1 price details. Required if `type` is `price`.
type V2BillingContractPricingLineActionAddPricingPriceDetailsParams struct {
	// The ID of the V1 price.
	Price *string `form:"price" json:"price"`
	// The quantity for the price. Only applicable for licensed prices.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
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
	// The ID of the V1 price to overwrite.
	Price *string `form:"price" json:"price"`
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

// Cancel a Contract object by ID.
type V2BillingContractCancelParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
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

// Timestamp to indicate when the contract line ends.
type V2BillingContractCreateContractLineEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp"`
}

// Timestamp to indicate when the override ends.
type V2BillingContractCreateContractLineOverrideEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp"`
}

// The amount of the credit grant.
type V2BillingContractCreateContractLineOverrideServiceActionAddCreditGrantAmountParams struct {
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary *Amount `form:"monetary" json:"monetary,omitempty"`
	// The type of the credit grant amount.
	Type *string `form:"type" json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingContractCreateContractLineOverrideServiceActionAddCreditGrantApplicabilityConfigScopeParams struct {
	// The billable items to apply the credit grant to.
	BillableItems []*string `form:"billable_items" json:"billable_items,omitempty"`
	// The price type that credit grants can apply to.
	PriceType *string `form:"price_type" json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingContractCreateContractLineOverrideServiceActionAddCreditGrantApplicabilityConfigParams struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingContractCreateContractLineOverrideServiceActionAddCreditGrantApplicabilityConfigScopeParams `form:"scope" json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingContractCreateContractLineOverrideServiceActionAddCreditGrantExpiryConfigParams struct {
	// The type of the expiry configuration.
	Type *string `form:"type" json:"type"`
}

// Details for the credit grant. Required if `type` is `credit_grant`.
type V2BillingContractCreateContractLineOverrideServiceActionAddCreditGrantParams struct {
	// The amount of the credit grant.
	Amount *V2BillingContractCreateContractLineOverrideServiceActionAddCreditGrantAmountParams `form:"amount" json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingContractCreateContractLineOverrideServiceActionAddCreditGrantApplicabilityConfigParams `form:"applicability_config" json:"applicability_config"`
	// The category of the credit grant.
	Category *string `form:"category" json:"category,omitempty"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingContractCreateContractLineOverrideServiceActionAddCreditGrantExpiryConfigParams `form:"expiry_config" json:"expiry_config"`
	// A descriptive name.
	Name *string `form:"name" json:"name"`
	// The desired priority for applying this credit grant. The highest priority is 0 and the lowest is 100.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
}

// Parameters for adding a new service action.
type V2BillingContractCreateContractLineOverrideServiceActionAddParams struct {
	// Details for the credit grant. Required if `type` is `credit_grant`.
	CreditGrant *V2BillingContractCreateContractLineOverrideServiceActionAddCreditGrantParams `form:"credit_grant" json:"credit_grant,omitempty"`
	// The interval for assessing service.
	ServiceInterval *string `form:"service_interval" json:"service_interval"`
	// The length of the interval for assessing service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count"`
	// The type of the service action.
	Type *string `form:"type" json:"type"`
}

// The amount of the credit grant.
type V2BillingContractCreateContractLineOverrideServiceActionReplaceCreditGrantAmountParams struct {
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary *Amount `form:"monetary" json:"monetary,omitempty"`
	// The type of the credit grant amount.
	Type *string `form:"type" json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingContractCreateContractLineOverrideServiceActionReplaceCreditGrantApplicabilityConfigScopeParams struct {
	// The billable items to apply the credit grant to.
	BillableItems []*string `form:"billable_items" json:"billable_items,omitempty"`
	// The price type that credit grants can apply to.
	PriceType *string `form:"price_type" json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingContractCreateContractLineOverrideServiceActionReplaceCreditGrantApplicabilityConfigParams struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingContractCreateContractLineOverrideServiceActionReplaceCreditGrantApplicabilityConfigScopeParams `form:"scope" json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingContractCreateContractLineOverrideServiceActionReplaceCreditGrantExpiryConfigParams struct {
	// The type of the expiry configuration.
	Type *string `form:"type" json:"type"`
}

// Details for the credit grant. Required if `type` is `credit_grant`.
type V2BillingContractCreateContractLineOverrideServiceActionReplaceCreditGrantParams struct {
	// The amount of the credit grant.
	Amount *V2BillingContractCreateContractLineOverrideServiceActionReplaceCreditGrantAmountParams `form:"amount" json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingContractCreateContractLineOverrideServiceActionReplaceCreditGrantApplicabilityConfigParams `form:"applicability_config" json:"applicability_config"`
	// The category of the credit grant.
	Category *string `form:"category" json:"category,omitempty"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingContractCreateContractLineOverrideServiceActionReplaceCreditGrantExpiryConfigParams `form:"expiry_config" json:"expiry_config"`
	// A descriptive name.
	Name *string `form:"name" json:"name"`
	// The desired priority for applying this credit grant. The highest priority is 0 and the lowest is 100.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
}

// Parameters for replacing an existing service action.
type V2BillingContractCreateContractLineOverrideServiceActionReplaceParams struct {
	// Details for the credit grant. Required if `type` is `credit_grant`.
	CreditGrant *V2BillingContractCreateContractLineOverrideServiceActionReplaceCreditGrantParams `form:"credit_grant" json:"credit_grant,omitempty"`
	// The ID of the service action to replace.
	ID *string `form:"id" json:"id,omitempty"`
	// The lookup key for the service action to replace.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// The interval for assessing service.
	ServiceInterval *string `form:"service_interval" json:"service_interval"`
	// The length of the interval for assessing service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count"`
	// The type of the service action.
	Type *string `form:"type" json:"type"`
}

// Service action override parameters. Required if `type` is `service_action`.
type V2BillingContractCreateContractLineOverrideServiceActionParams struct {
	// Parameters for adding a new service action.
	Add *V2BillingContractCreateContractLineOverrideServiceActionAddParams `form:"add" json:"add,omitempty"`
	// Parameters for replacing an existing service action.
	Replace *V2BillingContractCreateContractLineOverrideServiceActionReplaceParams `form:"replace" json:"replace,omitempty"`
	// The type of service action override.
	Type *string `form:"type" json:"type"`
}

// Timestamp to indicate when the override starts.
type V2BillingContractCreateContractLineOverrideStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp"`
}

// List of overrides. Later overrides in the list override earlier ones.
type V2BillingContractCreateContractLineOverrideParams struct {
	// Timestamp to indicate when the override ends.
	EndsAt *V2BillingContractCreateContractLineOverrideEndsAtParams `form:"ends_at" json:"ends_at"`
	// Service action override parameters. Required if `type` is `service_action`.
	ServiceAction *V2BillingContractCreateContractLineOverrideServiceActionParams `form:"service_action" json:"service_action,omitempty"`
	// Timestamp to indicate when the override starts.
	StartsAt *V2BillingContractCreateContractLineOverrideStartsAtParams `form:"starts_at" json:"starts_at"`
	// The type of the override.
	Type *string `form:"type" json:"type"`
}

// The pricing configuration for the contract line.
type V2BillingContractCreateContractLinePricingParams struct{}

// Timestamp to indicate when the contract line starts.
type V2BillingContractCreateContractLineStartsAtParams struct {
	// The timestamp when the item starts.
	Timestamp *time.Time `form:"timestamp" json:"timestamp"`
}

// A list of contract lines to create with the contract.
type V2BillingContractCreateContractLineParams struct {
	// Timestamp to indicate when the contract line ends.
	EndsAt *V2BillingContractCreateContractLineEndsAtParams `form:"ends_at" json:"ends_at"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// List of overrides. Later overrides in the list override earlier ones.
	Overrides []*V2BillingContractCreateContractLineOverrideParams `form:"overrides" json:"overrides"`
	// The pricing configuration for the contract line.
	Pricing *V2BillingContractCreateContractLinePricingParams `form:"pricing" json:"pricing"`
	// Timestamp to indicate when the contract line starts.
	StartsAt *V2BillingContractCreateContractLineStartsAtParams `form:"starts_at" json:"starts_at"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractCreateContractLineParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The effective at for the license quantity action.
type V2BillingContractCreateLicenseQuantityActionEffectiveAtParams struct {
	// The timestamp for the effective at.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the effective at.
	Type *string `form:"type" json:"type"`
}

// The set quantity for the license quantity action.
type V2BillingContractCreateLicenseQuantityActionSetParams struct {
	// The quantity to set.
	Quantity *int64 `form:"quantity" json:"quantity"`
}

// A list of license quantity actions to create with the contract.
type V2BillingContractCreateLicenseQuantityActionParams struct {
	// The effective at for the license quantity action.
	EffectiveAt *V2BillingContractCreateLicenseQuantityActionEffectiveAtParams `form:"effective_at" json:"effective_at"`
	// The ID of the license pricing.
	LicensePricingID *string `form:"license_pricing_id" json:"license_pricing_id,omitempty"`
	// The lookup key for the license pricing.
	LicensePricingLookupKey *string `form:"license_pricing_lookup_key" json:"license_pricing_lookup_key,omitempty"`
	// The type of the license pricing.
	LicensePricingType *string `form:"license_pricing_type" json:"license_pricing_type"`
	// The pricing line for the license quantity action.
	PricingLine *string `form:"pricing_line" json:"pricing_line,omitempty"`
	// The set quantity for the license quantity action.
	Set *V2BillingContractCreateLicenseQuantityActionSetParams `form:"set" json:"set,omitempty"`
	// The type of the license quantity action.
	Type *string `form:"type" json:"type"`
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

// V1 price details. Required if `type` is `price`.
type V2BillingContractCreatePricingLinePricingPriceDetailsParams struct {
	// The ID of the V1 price.
	Price *string `form:"price" json:"price"`
	// The quantity for the price. Only applicable for licensed prices.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
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

// Each element represents a pricing tier.
type V2BillingContractCreatePricingOverrideOverwritePriceTierParams struct {
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
type V2BillingContractCreatePricingOverrideOverwritePriceParams struct {
	// The ID of the V1 price to overwrite.
	Price *string `form:"price" json:"price"`
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractCreatePricingOverrideOverwritePriceTierParams `form:"tiers" json:"tiers,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
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
	// Parameters for an overwrite_price override. Required if `type` is `overwrite_price`.
	OverwritePrice *V2BillingContractCreatePricingOverrideOverwritePriceParams `form:"overwrite_price" json:"overwrite_price,omitempty"`
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
	// The billing settings for the contract.
	BillingSettings *V2BillingContractCreateBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// A list of contract lines to create with the contract.
	ContractLines []*V2BillingContractCreateContractLineParams `form:"contract_lines" json:"contract_lines"`
	// A unique user-provided contract number e.g. C-2026-0001.
	ContractNumber *string `form:"contract_number" json:"contract_number"`
	// Currency of the contract.
	Currency *string `form:"currency" json:"currency"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// A list of license quantity actions to create with the contract.
	LicenseQuantityActions []*V2BillingContractCreateLicenseQuantityActionParams `form:"license_quantity_actions" json:"license_quantity_actions"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// A list of one-time fees to create with the contract. Each fee is billed as individual invoice items per its bill_schedule.
	OneTimeFees []*V2BillingContractCreateOneTimeFeeParams `form:"one_time_fees" json:"one_time_fees,omitempty"`
	// A list of pricing lines to create with the contract.
	PricingLines []*V2BillingContractCreatePricingLineParams `form:"pricing_lines" json:"pricing_lines"`
	// A list of pricing overrides to create with the contract.
	PricingOverrides []*V2BillingContractCreatePricingOverrideParams `form:"pricing_overrides" json:"pricing_overrides"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingContractCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a Contract object by ID.
type V2BillingContractRetrieveParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// The effective at for the license quantity action.
type V2BillingContractUpdateLicenseQuantityActionEffectiveAtParams struct {
	// The timestamp for the effective at.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of the effective at.
	Type *string `form:"type" json:"type"`
}

// The set quantity for the license quantity action.
type V2BillingContractUpdateLicenseQuantityActionSetParams struct {
	// The quantity to set.
	Quantity *int64 `form:"quantity" json:"quantity"`
}

// Schema-only: License quantity actions (implementation to follow).
type V2BillingContractUpdateLicenseQuantityActionParams struct {
	// The effective at for the license quantity action.
	EffectiveAt *V2BillingContractUpdateLicenseQuantityActionEffectiveAtParams `form:"effective_at" json:"effective_at"`
	// The ID of the license pricing.
	LicensePricingID *string `form:"license_pricing_id" json:"license_pricing_id,omitempty"`
	// The lookup key for the license pricing.
	LicensePricingLookupKey *string `form:"license_pricing_lookup_key" json:"license_pricing_lookup_key,omitempty"`
	// The type of the license pricing.
	LicensePricingType *string `form:"license_pricing_type" json:"license_pricing_type"`
	// The pricing line ID for the license quantity action.
	PricingLine *string `form:"pricing_line" json:"pricing_line,omitempty"`
	// The pricing line lookup key for the license quantity action.
	PricingLineLookupKey *string `form:"pricing_line_lookup_key" json:"pricing_line_lookup_key,omitempty"`
	// The set quantity for the license quantity action.
	Set *V2BillingContractUpdateLicenseQuantityActionSetParams `form:"set" json:"set,omitempty"`
	// The type of the license quantity action.
	Type *string `form:"type" json:"type"`
}

// The end time for the pricing line.
type V2BillingContractUpdatePricingLineActionAddEndsAtParams struct {
	// The timestamp when the item ends.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// The type of end time to apply.
	Type *string `form:"type" json:"type"`
}

// V1 price details. Required if `type` is `price`.
type V2BillingContractUpdatePricingLineActionAddPricingPriceDetailsParams struct {
	// The ID of the V1 price.
	Price *string `form:"price" json:"price"`
	// The quantity for the price. Only applicable for licensed prices.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
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
	// The ID of the V1 price to overwrite.
	Price *string `form:"price" json:"price"`
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
	// Schema-only: License quantity actions (implementation to follow).
	LicenseQuantityActions []*V2BillingContractUpdateLicenseQuantityActionParams `form:"license_quantity_actions" json:"license_quantity_actions,omitempty"`
	// Pricing line actions to apply.
	PricingLineActions []*V2BillingContractUpdatePricingLineActionParams `form:"pricing_line_actions" json:"pricing_line_actions,omitempty"`
	// Pricing override actions to apply.
	PricingOverrideActions []*V2BillingContractUpdatePricingOverrideActionParams `form:"pricing_override_actions" json:"pricing_override_actions,omitempty"`
}
