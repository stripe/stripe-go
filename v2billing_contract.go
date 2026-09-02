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
type V2BillingContractPricingOverridesDataMultiplyPricingCriterionType string

// List of values that V2BillingContractPricingOverridesDataMultiplyPricingCriterionType can take
const (
	V2BillingContractPricingOverridesDataMultiplyPricingCriterionTypeExclude V2BillingContractPricingOverridesDataMultiplyPricingCriterionType = "exclude"
	V2BillingContractPricingOverridesDataMultiplyPricingCriterionTypeInclude V2BillingContractPricingOverridesDataMultiplyPricingCriterionType = "include"
)

// The type of pricing override.
type V2BillingContractPricingOverridesDataType string

// List of values that V2BillingContractPricingOverridesDataType can take
const (
	V2BillingContractPricingOverridesDataTypeMultiplyPricing V2BillingContractPricingOverridesDataType = "multiply_pricing"
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

// The billing cycle anchor.
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

// The billing settings.
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
	// The id of the one-time fee.
	ID string `json:"id"`
	// The user-provided lookup key.
	LookupKey string `json:"lookup_key,omitempty"`
	// The id of the product for this fee.
	Product string `json:"product"`
}

// The one-time fees. Only populated when `one_time_fees` is passed in the `include` parameter.
type V2BillingContractOneTimeFees struct {
	// The one-time fees for this page.
	Data []*V2BillingContractOneTimeFeesData `json:"data"`
}

// Timestamp when the pricing line ends.
type V2BillingContractPricingLinesDataEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// Timestamp when this override ends.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// Details for an overwrite_price override.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePrice struct {
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount string `json:"unit_amount,omitempty"`
}

// Timestamp when this override starts.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// The pricing line overrides.
type V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesData struct {
	// Timestamp when this override ends.
	EndsAt *V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataEndsAt `json:"ends_at"`
	// The ID of the pricing override.
	ID string `json:"id"`
	// The user-provided lookup key for this override.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of key-value pairs.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Details for an overwrite_price override.
	OverwritePrice *V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePrice `json:"overwrite_price,omitempty"`
	// The priority of this override relative to others. Lower number = higher priority.
	Priority int64 `json:"priority"`
	// Timestamp when this override starts.
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

// Timestamp when the pricing line starts.
type V2BillingContractPricingLinesDataStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// The pricing lines for this page.
type V2BillingContractPricingLinesData struct {
	// Timestamp when the pricing line ends.
	EndsAt *V2BillingContractPricingLinesDataEndsAt `json:"ends_at"`
	// The id of the pricing line.
	ID string `json:"id"`
	// The user-provided lookup key for the pricing line.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of key-value pairs.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The pricing configuration for the pricing line.
	Pricing *V2BillingContractPricingLinesDataPricing `json:"pricing"`
	// Timestamp when the pricing line starts.
	StartsAt *V2BillingContractPricingLinesDataStartsAt `json:"starts_at"`
}

// The pricing lines. Only populated when `pricing_lines` is passed in the `include` parameter.
type V2BillingContractPricingLines struct {
	// The pricing lines for this page.
	Data []*V2BillingContractPricingLinesData `json:"data"`
}

// Resolved timestamp when the pricing override ends.
type V2BillingContractPricingOverridesDataEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// Criteria determining which rates the multiply_pricing override applies to.
type V2BillingContractPricingOverridesDataMultiplyPricingCriterion struct {
	// Filter by pricing line IDs.
	PricingLineIDs []string `json:"pricing_line_ids,omitempty"`
	// Filter by pricing line lookup keys.
	PricingLineLookupKeys []string `json:"pricing_line_lookup_keys,omitempty"`
	// Whether to include or exclude items matching these criteria.
	Type V2BillingContractPricingOverridesDataMultiplyPricingCriterionType `json:"type"`
}

// Details for a multiply_pricing override.
type V2BillingContractPricingOverridesDataMultiplyPricing struct {
	// Criteria determining which rates the multiply_pricing override applies to.
	Criteria []*V2BillingContractPricingOverridesDataMultiplyPricingCriterion `json:"criteria"`
	// The multiply_pricing factor, represented as a decimal string. e.g. "0.8" for a 20% reduction.
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
	// Set of key-value pairs.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Details for a multiply_pricing override.
	MultiplyPricing *V2BillingContractPricingOverridesDataMultiplyPricing `json:"multiply_pricing,omitempty"`
	// The priority of this override relative to others. Lower number = higher priority.
	Priority int64 `json:"priority"`
	// Resolved timestamp when the pricing override starts.
	StartsAt *V2BillingContractPricingOverridesDataStartsAt `json:"starts_at"`
	// The type of pricing override.
	Type V2BillingContractPricingOverridesDataType `json:"type"`
}

// The pricing overrides. Only populated when `pricing_overrides` is passed in the `include` parameter.
type V2BillingContractPricingOverrides struct {
	// The pricing overrides for this page.
	Data []*V2BillingContractPricingOverridesData `json:"data"`
}

// Historical timestamps of when the contract transitioned into each status.
type V2BillingContractStatusTransitions struct {
	// The timestamp when the contract was activated.
	ActivatedAt time.Time `json:"activated_at,omitempty"`
	// The timestamp when the contract was canceled.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// The timestamp when the contract ended.
	EndedAt time.Time `json:"ended_at,omitempty"`
}

// Contract resource representing a comprehensive sales agreement
type V2BillingContract struct {
	APIResource
	// The billing cycle anchor.
	BillingCycleAnchor *V2BillingContractBillingCycleAnchor `json:"billing_cycle_anchor,omitempty"`
	// The billing settings.
	BillingSettings *V2BillingContractBillingSettings `json:"billing_settings,omitempty"`
	// A unique user-provided contract number e.g. C-2026-0001.
	ContractNumber string `json:"contract_number"`
	// Timestamp of when the contract was created.
	Created time.Time `json:"created"`
	// The currency.
	Currency Currency `json:"currency"`
	// The customer id.
	Customer string `json:"customer"`
	// The contract id.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The one-time fees. Only populated when `one_time_fees` is passed in the `include` parameter.
	OneTimeFees *V2BillingContractOneTimeFees `json:"one_time_fees,omitempty"`
	// The pricing lines. Only populated when `pricing_lines` is passed in the `include` parameter.
	PricingLines *V2BillingContractPricingLines `json:"pricing_lines,omitempty"`
	// The pricing overrides. Only populated when `pricing_overrides` is passed in the `include` parameter.
	PricingOverrides *V2BillingContractPricingOverrides `json:"pricing_overrides,omitempty"`
	// The current status of the contract.
	Status V2BillingContractStatus `json:"status"`
	// Historical timestamps of when the contract transitioned into each status.
	StatusTransitions *V2BillingContractStatusTransitions `json:"status_transitions,omitempty"`
}
