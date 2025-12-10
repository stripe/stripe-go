//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of the product.
type V2TaxManualRuleProductType string

// List of values that V2TaxManualRuleProductType can take
const (
	V2TaxManualRuleProductTypeLicensedItem V2TaxManualRuleProductType = "licensed_item"
	V2TaxManualRuleProductTypeMeteredItem  V2TaxManualRuleProductType = "metered_item"
	V2TaxManualRuleProductTypeTaxCode      V2TaxManualRuleProductType = "tax_code"
)

// The status of the ManualRule object.
type V2TaxManualRuleStatus string

// List of values that V2TaxManualRuleStatus can take
const (
	V2TaxManualRuleStatusActive   V2TaxManualRuleStatus = "active"
	V2TaxManualRuleStatusInactive V2TaxManualRuleStatus = "inactive"
)

// Location where the rule applies.
type V2TaxManualRuleLocation struct {
	// Country ISO-3166.
	Country string `json:"country"`
	// State ISO-3166.
	State string `json:"state,omitempty"`
}

// Products associated with the rule.
type V2TaxManualRuleProduct struct {
	// The licensed item identifier.
	LicensedItem string `json:"licensed_item,omitempty"`
	// The metered item identifier.
	MeteredItem string `json:"metered_item,omitempty"`
	// The tax code for the product.
	TaxCode string `json:"tax_code,omitempty"`
	// The type of the product.
	Type V2TaxManualRuleProductType `json:"type"`
}

// The tax rates to be applied.
type V2TaxManualRuleScheduledTaxRateRate struct {
	// Country of the tax rate.
	Country string `json:"country,omitempty"`
	// Description of the tax rate.
	Description string `json:"description,omitempty"`
	// Display name of the tax rate as it will be shown on the invoice.
	DisplayName string `json:"display_name"`
	// Jurisdiction of the tax rate should apply as it will be shown on the invoice.
	Jurisdiction string `json:"jurisdiction,omitempty"`
	// Percentage of the tax rate. Must be positive and maximum of 4 decimal points.
	Percentage string `json:"percentage"`
	// State of the tax rate.
	State string `json:"state,omitempty"`
}

// Tax rates to be applied.
type V2TaxManualRuleScheduledTaxRate struct {
	// The tax rates to be applied.
	Rates []*V2TaxManualRuleScheduledTaxRateRate `json:"rates"`
	// The start time for the tax rate.
	StartsAt time.Time `json:"starts_at,omitempty"`
}

// A ManualRule holds tax rates for a BillableItem. It can hold at most 5 distinct tax rates.
type V2TaxManualRule struct {
	APIResource
	// The time at which the ManualRule object was created.
	Created time.Time `json:"created"`
	// The ID of the ManualRule object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Location where the rule applies.
	Location *V2TaxManualRuleLocation `json:"location,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Products associated with the rule.
	Products []*V2TaxManualRuleProduct `json:"products"`
	// Tax rates to be applied.
	ScheduledTaxRates []*V2TaxManualRuleScheduledTaxRate `json:"scheduled_tax_rates"`
	// The status of the ManualRule object.
	Status V2TaxManualRuleStatus `json:"status"`
}
