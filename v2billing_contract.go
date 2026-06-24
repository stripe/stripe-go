//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of tax calculation.
type V2BillingContractBillingSettingsBillSettingsDetailsCalculationTaxType string

// List of values that V2BillingContractBillingSettingsBillSettingsDetailsCalculationTaxType can take
const (
	V2BillingContractBillingSettingsBillSettingsDetailsCalculationTaxTypeAutomatic V2BillingContractBillingSettingsBillSettingsDetailsCalculationTaxType = "automatic"
	V2BillingContractBillingSettingsBillSettingsDetailsCalculationTaxTypeManual    V2BillingContractBillingSettingsBillSettingsDetailsCalculationTaxType = "manual"
)

// The interval unit.
type V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueInterval string

// List of values that V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueInterval can take
const (
	V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueIntervalDay   V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueInterval = "day"
	V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueIntervalMonth V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueInterval = "month"
	V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueIntervalWeek  V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueInterval = "week"
	V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueIntervalYear  V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueInterval = "year"
)

// The collection method.
type V2BillingContractBillingSettingsCollectionSettingsDetailsCollectionMethod string

// List of values that V2BillingContractBillingSettingsCollectionSettingsDetailsCollectionMethod can take
const (
	V2BillingContractBillingSettingsCollectionSettingsDetailsCollectionMethodChargeAutomatically V2BillingContractBillingSettingsCollectionSettingsDetailsCollectionMethod = "charge_automatically"
	V2BillingContractBillingSettingsCollectionSettingsDetailsCollectionMethodSendInvoice         V2BillingContractBillingSettingsCollectionSettingsDetailsCollectionMethod = "send_invoice"
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
type V2BillingContractBillingSettingsBillSettingsDetailsCalculationTax struct {
	// The type of tax calculation.
	Type V2BillingContractBillingSettingsBillSettingsDetailsCalculationTaxType `json:"type"`
}

// Calculation settings.
type V2BillingContractBillingSettingsBillSettingsDetailsCalculation struct {
	// Tax calculation settings.
	Tax *V2BillingContractBillingSettingsBillSettingsDetailsCalculationTax `json:"tax,omitempty"`
}

// The number of time units before the invoice is past due.
type V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDue struct {
	// The interval unit.
	Interval V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDueInterval `json:"interval"`
	// The number of intervals.
	IntervalCount int64 `json:"interval_count"`
}

// Invoice settings.
type V2BillingContractBillingSettingsBillSettingsDetailsInvoice struct {
	// The number of time units before the invoice is past due.
	TimeUntilDue *V2BillingContractBillingSettingsBillSettingsDetailsInvoiceTimeUntilDue `json:"time_until_due,omitempty"`
}

// The bill settings details configures invoice and tax settings for the contract.
type V2BillingContractBillingSettingsBillSettingsDetails struct {
	// Calculation settings.
	Calculation *V2BillingContractBillingSettingsBillSettingsDetailsCalculation `json:"calculation,omitempty"`
	// Invoice settings.
	Invoice *V2BillingContractBillingSettingsBillSettingsDetailsInvoice `json:"invoice,omitempty"`
}

// The billing profile details configures who is charged for the contract.
type V2BillingContractBillingSettingsBillingProfileDetails struct {
	// The customer who pays for the contract invoice.
	Customer string `json:"customer"`
	// The default payment method for the contract.
	DefaultPaymentMethod string `json:"default_payment_method,omitempty"`
}

// The collection settings details configures how payments are collected on the contract.
type V2BillingContractBillingSettingsCollectionSettingsDetails struct {
	// The collection method.
	CollectionMethod V2BillingContractBillingSettingsCollectionSettingsDetailsCollectionMethod `json:"collection_method"`
	// The payment method configuration.
	PaymentMethodConfiguration string `json:"payment_method_configuration,omitempty"`
}

// The billing settings for the contract.
type V2BillingContractBillingSettings struct {
	// The billing profile details configures who is charged for the contract.
	BillingProfileDetails *V2BillingContractBillingSettingsBillingProfileDetails `json:"billing_profile_details"`
	// The bill settings details configures invoice and tax settings for the contract.
	BillSettingsDetails *V2BillingContractBillingSettingsBillSettingsDetails `json:"bill_settings_details,omitempty"`
	// The collection settings details configures how payments are collected on the contract.
	CollectionSettingsDetails *V2BillingContractBillingSettingsCollectionSettingsDetails `json:"collection_settings_details"`
}

// When this fee will be billed. Always contains a concrete timestamp.
type V2BillingContractOneTimeFeesDataBillAt struct {
	// The timestamp at which the fee will be billed.
	Timestamp time.Time `json:"timestamp"`
}

// The one-time fees for this page.
type V2BillingContractOneTimeFeesData struct {
	// The amount billed for this fee.
	Amount Amount `json:"amount"`
	// When this fee will be billed. Always contains a concrete timestamp.
	BillAt *V2BillingContractOneTimeFeesDataBillAt `json:"bill_at"`
	// The ID of the one-time fee.
	ID string `json:"id"`
	// The user-provided lookup key.
	LookupKey string `json:"lookup_key,omitempty"`
	// The ID of the v1 Product for this fee.
	Product string `json:"product"`
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

// Criteria determining which rates the multiplier applies to.
type V2BillingContractPricingOverridesDataMultiplierCriterion struct {
	// Filter by pricing line IDs.
	PricingLineIDs []string `json:"pricing_line_ids,omitempty"`
	// Filter by pricing line lookup keys.
	PricingLineLookupKeys []string `json:"pricing_line_lookup_keys,omitempty"`
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
