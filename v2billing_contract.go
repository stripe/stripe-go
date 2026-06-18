//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of tax calculation.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxType string

// List of values that V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxType can take
const (
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxTypeAutomatic V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxType = "automatic"
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxTypeManual    V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxType = "manual"
)

// The interval unit.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval string

// List of values that V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval can take
const (
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueIntervalDay   V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval = "day"
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueIntervalMonth V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval = "month"
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueIntervalWeek  V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval = "week"
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueIntervalYear  V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval = "year"
)

// The collection method.
type V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethod string

// List of values that V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethod can take
const (
	V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethodChargeAutomatically V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethod = "charge_automatically"
	V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethodSendInvoice         V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethod = "send_invoice"
)

// The type of the bill_at.
type V2BillingContractOneTimeFeesDataBillScheduleBillAtType string

// List of values that V2BillingContractOneTimeFeesDataBillScheduleBillAtType can take
const (
	V2BillingContractOneTimeFeesDataBillScheduleBillAtTypeContractStart V2BillingContractOneTimeFeesDataBillScheduleBillAtType = "contract_start"
	V2BillingContractOneTimeFeesDataBillScheduleBillAtTypeDatetime      V2BillingContractOneTimeFeesDataBillScheduleBillAtType = "datetime"
)

// The type of billable item that this fee references.
type V2BillingContractOneTimeFeesDataBillableItemType string

// List of values that V2BillingContractOneTimeFeesDataBillableItemType can take
const (
	V2BillingContractOneTimeFeesDataBillableItemTypeProduct V2BillingContractOneTimeFeesDataBillableItemType = "product"
)

// Defines whether the tiered price should be graduated or volume-based.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTieringMode string

// List of values that V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTieringMode can take
const (
	V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTieringModeGraduated V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTieringMode = "graduated"
	V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTieringModeVolume    V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTieringMode = "volume"
)

// No upper bound to this tier.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTierUpToInf string

// List of values that V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTierUpToInf can take
const (
	V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTierUpToInfInf V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTierUpToInf = "inf"
)

// The type of override.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataType string

// List of values that V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataType can take
const (
	V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataTypeOverwritePrice V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataType = "overwrite_price"
)

// The type of pricing.
type V2BillingContractPricingLinesDataPricingType string

// List of values that V2BillingContractPricingLinesDataPricingType can take
const (
	V2BillingContractPricingLinesDataPricingTypePrice V2BillingContractPricingLinesDataPricingType = "price"
)

// Filter by billable item type.
type V2BillingContractPricingOverridesDataMultiplierCriterionBillableItemType string

// List of values that V2BillingContractPricingOverridesDataMultiplierCriterionBillableItemType can take
const (
	V2BillingContractPricingOverridesDataMultiplierCriterionBillableItemTypeLicensed V2BillingContractPricingOverridesDataMultiplierCriterionBillableItemType = "licensed"
	V2BillingContractPricingOverridesDataMultiplierCriterionBillableItemTypeMetered  V2BillingContractPricingOverridesDataMultiplierCriterionBillableItemType = "metered"
)

// Whether to include or exclude items matching these criteria.
type V2BillingContractPricingOverridesDataMultiplierCriterionType string

// List of values that V2BillingContractPricingOverridesDataMultiplierCriterionType can take
const (
	V2BillingContractPricingOverridesDataMultiplierCriterionTypeExclude V2BillingContractPricingOverridesDataMultiplierCriterionType = "exclude"
	V2BillingContractPricingOverridesDataMultiplierCriterionTypeInclude V2BillingContractPricingOverridesDataMultiplierCriterionType = "include"
)

// The type of pricing override.
type V2BillingContractPricingOverridesDataType string

// List of values that V2BillingContractPricingOverridesDataType can take
const (
	V2BillingContractPricingOverridesDataTypeMultiplier V2BillingContractPricingOverridesDataType = "multiplier"
)

// The current status of the contract.
type V2BillingContractStatus string

// List of values that V2BillingContractStatus can take
const (
	V2BillingContractStatusActive   V2BillingContractStatus = "active"
	V2BillingContractStatusCanceled V2BillingContractStatus = "canceled"
	V2BillingContractStatusDraft    V2BillingContractStatus = "draft"
	V2BillingContractStatusEnded    V2BillingContractStatus = "ended"
)

// The billing cycle anchor for the contract.
type V2BillingContractBillingCycleAnchor struct {
	// The billing cycle anchor as a UTC timestamp.
	Timestamp time.Time `json:"timestamp"`
}

// Tax calculation settings.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTax struct {
	// The type of tax calculation.
	Type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxType `json:"type"`
}

// Calculation settings.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculation struct {
	// Tax calculation settings.
	Tax *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTax `json:"tax,omitempty"`
}

// The number of time units before the invoice is past due.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDue struct {
	// The interval unit.
	Interval V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval `json:"interval"`
	// The number of intervals.
	IntervalCount int64 `json:"interval_count"`
}

// Invoice settings.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoice struct {
	// The number of time units before the invoice is past due.
	TimeUntilDue *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDue `json:"time_until_due,omitempty"`
}

// The bill settings details.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetails struct {
	// Calculation settings.
	Calculation *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculation `json:"calculation,omitempty"`
	// Invoice settings.
	Invoice *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoice `json:"invoice,omitempty"`
}

// The billing profile details.
type V2BillingContractBillingSettingsContractBillingDetailsBillingProfileDetails struct {
	// The customer who pays for the contract invoice.
	Customer string `json:"customer"`
	// The default payment method for the contract.
	DefaultPaymentMethod string `json:"default_payment_method,omitempty"`
}

// The collection settings details.
type V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetails struct {
	// The collection method.
	CollectionMethod V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethod `json:"collection_method"`
	// The payment method configuration.
	PaymentMethodConfiguration string `json:"payment_method_configuration,omitempty"`
}

// Billing settings details for the contract.
type V2BillingContractBillingSettingsContractBillingDetails struct {
	// The billing profile details.
	BillingProfileDetails *V2BillingContractBillingSettingsContractBillingDetailsBillingProfileDetails `json:"billing_profile_details"`
	// The bill settings details.
	BillSettingsDetails *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetails `json:"bill_settings_details,omitempty"`
	// The collection settings details.
	CollectionSettingsDetails *V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetails `json:"collection_settings_details"`
}

// The billing settings for the contract.
type V2BillingContractBillingSettings struct {
	// Billing settings details for the contract.
	ContractBillingDetails *V2BillingContractBillingSettingsContractBillingDetails `json:"contract_billing_details,omitempty"`
}

// When this entry will be billed.
type V2BillingContractOneTimeFeesDataBillScheduleBillAt struct {
	// The datetime at which the entry will be billed. Set when `type` is `datetime`.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// The type of the bill_at.
	Type V2BillingContractOneTimeFeesDataBillScheduleBillAtType `json:"type"`
}

// The resolved bill schedule for the fee.
type V2BillingContractOneTimeFeesDataBillSchedule struct {
	// When this entry will be billed.
	BillAt *V2BillingContractOneTimeFeesDataBillScheduleBillAt `json:"bill_at"`
	// The amount to bill for this entry, in the smallest currency unit.
	Value int64 `json:"value,string"`
}

// Details for a product billable target. Set when `billable_item_type` is `product`.
type V2BillingContractOneTimeFeesDataProductDetails struct {
	// The ID of the v1 Product.
	Product string `json:"product"`
}

// The one-time fees for this page.
type V2BillingContractOneTimeFeesData struct {
	// The type of billable item that this fee references.
	BillableItemType V2BillingContractOneTimeFeesDataBillableItemType `json:"billable_item_type"`
	// The resolved bill schedule for the fee.
	BillSchedule []*V2BillingContractOneTimeFeesDataBillSchedule `json:"bill_schedule"`
	// The ID of the one-time fee.
	ID string `json:"id"`
	// The user-provided lookup key.
	LookupKey string `json:"lookup_key,omitempty"`
	// Details for a product billable target. Set when `billable_item_type` is `product`.
	ProductDetails *V2BillingContractOneTimeFeesDataProductDetails `json:"product_details,omitempty"`
}

// The one-time fees of the contract. Only populated when `one_time_fees` is passed in the `include` parameter.
type V2BillingContractOneTimeFees struct {
	// The one-time fees for this page.
	Data []*V2BillingContractOneTimeFeesData `json:"data"`
}

// Resolved timestamp when the pricing line ends.
type V2BillingContractPricingLinesDataEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// Resolved timestamp when this override ends.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// Each element represents a pricing tier.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTier struct {
	// Price for the entire tier, represented as a decimal string in minor currency units.
	FlatAmount string `json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units.
	UnitAmount string `json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier.
	UpToDecimal float64 `json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier.
	UpToInf V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTierUpToInf `json:"up_to_inf,omitempty"`
}

// Details for an overwrite_price override.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePrice struct {
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTieringMode `json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePriceTier `json:"tiers"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount string `json:"unit_amount,omitempty"`
}

// Resolved timestamp when this override starts.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// The pricing line overrides.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesData struct {
	// Resolved timestamp when this override ends.
	EndsAt *V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataEndsAt `json:"ends_at"`
	// The user-provided lookup key for this override.
	LookupKey string `json:"lookup_key,omitempty"`
	// Details for an overwrite_price override.
	OverwritePrice *V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePrice `json:"overwrite_price,omitempty"`
	// The ID of the pricing line override.
	PricingOverride string `json:"pricing_override"`
	// Resolved timestamp when this override starts.
	StartsAt *V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataStartsAt `json:"starts_at"`
	// The type of override.
	Type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataType `json:"type"`
}

// The overwrite_price overrides embedded directly on this pricing line.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverrides struct {
	// The pricing line overrides.
	Data []*V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesData `json:"data"`
}

// V1 price details. Present when `type` is `price`.
type V2BillingContractPricingLinesDataPricingPriceDetails struct {
	// The current quantity on this pricing line.
	CurrentQuantity float64 `json:"current_quantity,string"`
	// The ID of the V1 price.
	Price string `json:"price"`
	// The overwrite_price overrides embedded directly on this pricing line.
	PricingOverrides *V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverrides `json:"pricing_overrides,omitempty"`
}

// The pricing configuration for the pricing line.
type V2BillingContractPricingLinesDataPricing struct {
	// V1 price details. Present when `type` is `price`.
	PriceDetails *V2BillingContractPricingLinesDataPricingPriceDetails `json:"price_details,omitempty"`
	// The type of pricing.
	Type V2BillingContractPricingLinesDataPricingType `json:"type"`
}

// Resolved timestamp when the pricing line starts.
type V2BillingContractPricingLinesDataStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// The pricing lines for this page.
type V2BillingContractPricingLinesData struct {
	// Resolved timestamp when the pricing line ends.
	EndsAt *V2BillingContractPricingLinesDataEndsAt `json:"ends_at"`
	// The ID of the pricing line.
	ID string `json:"id"`
	// The user-provided lookup key for the pricing line.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The pricing configuration for the pricing line.
	Pricing *V2BillingContractPricingLinesDataPricing `json:"pricing"`
	// Resolved timestamp when the pricing line starts.
	StartsAt *V2BillingContractPricingLinesDataStartsAt `json:"starts_at"`
}

// The pricing lines of the contract. Only populated when `pricing_lines` is passed in the `include` parameter.
type V2BillingContractPricingLines struct {
	// The pricing lines for this page.
	Data []*V2BillingContractPricingLinesData `json:"data"`
}

// Resolved timestamp when the pricing override ends.
type V2BillingContractPricingOverridesDataEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// All of these key-value conditions must match.
type V2BillingContractPricingOverridesDataMultiplierCriterionMetadataConditionAllOf struct {
	// The metadata key.
	Key string `json:"key"`
	// The metadata value.
	Value string `json:"value"`
}

// Filter by metadata conditions.
type V2BillingContractPricingOverridesDataMultiplierCriterionMetadataCondition struct {
	// All of these key-value conditions must match.
	AllOf []*V2BillingContractPricingOverridesDataMultiplierCriterionMetadataConditionAllOf `json:"all_of"`
}

// Criteria determining which rates the multiplier applies to.
type V2BillingContractPricingOverridesDataMultiplierCriterion struct {
	// Filter by billable item IDs.
	BillableItemIDs []string `json:"billable_item_ids"`
	// Filter by billable item lookup keys.
	BillableItemLookupKeys []string `json:"billable_item_lookup_keys"`
	// Filter by billable item type.
	BillableItemTypes []V2BillingContractPricingOverridesDataMultiplierCriterionBillableItemType `json:"billable_item_types"`
	// Filter by metadata conditions.
	MetadataConditions []*V2BillingContractPricingOverridesDataMultiplierCriterionMetadataCondition `json:"metadata_conditions"`
	// Filter by rate card IDs. Only applicable for `multiplier` overrides.
	RateCardIDs []string `json:"rate_card_ids"`
	// Whether to include or exclude items matching these criteria.
	Type V2BillingContractPricingOverridesDataMultiplierCriterionType `json:"type"`
}

// Details for a multiplier override.
type V2BillingContractPricingOverridesDataMultiplier struct {
	// Criteria determining which rates the multiplier applies to.
	Criteria []*V2BillingContractPricingOverridesDataMultiplierCriterion `json:"criteria"`
	// The multiplier factor, represented as a decimal string. e.g. "0.8" for a 20% reduction.
	Factor string `json:"factor"`
}

// Resolved timestamp when the pricing override starts.
type V2BillingContractPricingOverridesDataStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// The pricing overrides for this page.
type V2BillingContractPricingOverridesData struct {
	// Resolved timestamp when the pricing override ends.
	EndsAt *V2BillingContractPricingOverridesDataEndsAt `json:"ends_at"`
	// The ID of the pricing override.
	ID string `json:"id"`
	// The user-provided lookup key for the pricing override.
	LookupKey string `json:"lookup_key,omitempty"`
	// Details for a multiplier override.
	Multiplier *V2BillingContractPricingOverridesDataMultiplier `json:"multiplier,omitempty"`
	// The priority of this override relative to others. Lower number = higher priority.
	Priority int64 `json:"priority"`
	// Resolved timestamp when the pricing override starts.
	StartsAt *V2BillingContractPricingOverridesDataStartsAt `json:"starts_at"`
	// The type of pricing override.
	Type V2BillingContractPricingOverridesDataType `json:"type"`
}

// The pricing overrides of the contract. Only populated when `pricing_overrides` is passed in the `include` parameter.
type V2BillingContractPricingOverrides struct {
	// The pricing overrides for this page.
	Data []*V2BillingContractPricingOverridesData `json:"data"`
}

// Details of the active contract status.
type V2BillingContractStatusDetailsActive struct {
	// The timestamp when the contract was activated.
	ActivatedAt time.Time `json:"activated_at"`
}

// Details of the canceled contract status.
type V2BillingContractStatusDetailsCanceled struct {
	// The timestamp when the contract was canceled.
	CanceledAt time.Time `json:"canceled_at"`
}

// Information about the contract status transitions.
type V2BillingContractStatusDetails struct {
	// Details of the active contract status.
	Active *V2BillingContractStatusDetailsActive `json:"active,omitempty"`
	// Details of the canceled contract status.
	Canceled *V2BillingContractStatusDetailsCanceled `json:"canceled,omitempty"`
}

// Main Contract resource representing a comprehensive billing agreement
type V2BillingContract struct {
	APIResource
	// The billing cycle anchor for the contract.
	BillingCycleAnchor *V2BillingContractBillingCycleAnchor `json:"billing_cycle_anchor,omitempty"`
	// The billing settings for the contract.
	BillingSettings *V2BillingContractBillingSettings `json:"billing_settings,omitempty"`
	// A unique user-provided contract number e.g. C-2026-0001.
	ContractNumber string `json:"contract_number"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// The currency of the contract.
	Currency Currency `json:"currency"`
	// The ID of the customer associated with the contract.
	Customer string `json:"customer"`
	// The ID of the contract object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The one-time fees of the contract. Only populated when `one_time_fees` is passed in the `include` parameter.
	OneTimeFees *V2BillingContractOneTimeFees `json:"one_time_fees,omitempty"`
	// The pricing lines of the contract. Only populated when `pricing_lines` is passed in the `include` parameter.
	PricingLines *V2BillingContractPricingLines `json:"pricing_lines,omitempty"`
	// The pricing overrides of the contract. Only populated when `pricing_overrides` is passed in the `include` parameter.
	PricingOverrides *V2BillingContractPricingOverrides `json:"pricing_overrides,omitempty"`
	// The current status of the contract.
	Status V2BillingContractStatus `json:"status"`
	// Information about the contract status transitions.
	StatusDetails *V2BillingContractStatusDetails `json:"status_details"`
}
