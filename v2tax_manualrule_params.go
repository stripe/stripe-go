//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List all ManualRule objects.
type V2TaxManualRuleListParams struct {
	Params `form:"*"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Location where the rule applies.
type V2TaxManualRuleLocationParams struct {
	// Country ISO-3166.
	Country *string `form:"country" json:"country"`
	// State ISO-3166.
	State *string `form:"state" json:"state,omitempty"`
}

// Products associated with the rule.
type V2TaxManualRuleProductParams struct {
	// The licensed item identifier.
	LicensedItem *string `form:"licensed_item" json:"licensed_item,omitempty"`
	// The metered item identifier.
	MeteredItem *string `form:"metered_item" json:"metered_item,omitempty"`
	// The tax code for the product.
	TaxCode *string `form:"tax_code" json:"tax_code,omitempty"`
	// The type of the product.
	Type *string `form:"type" json:"type"`
}

// The tax rates to be applied.
type V2TaxManualRuleScheduledTaxRateRateParams struct {
	// Country of the tax rate.
	Country *string `form:"country" json:"country,omitempty"`
	// Description of the tax rate.
	Description *string `form:"description" json:"description,omitempty"`
	// Display name of the tax rate as it will be shown on the invoice.
	DisplayName *string `form:"display_name" json:"display_name"`
	// Jurisdiction of the tax rate should apply as it will be shown on the invoice.
	Jurisdiction *string `form:"jurisdiction" json:"jurisdiction,omitempty"`
	// Percentage of the tax rate. Must be positive and maximum of 4 decimal points.
	Percentage *string `form:"percentage" json:"percentage"`
	// State of the tax rate.
	State *string `form:"state" json:"state,omitempty"`
}

// Tax rates to be applied.
type V2TaxManualRuleScheduledTaxRateParams struct {
	// The tax rates to be applied.
	Rates []*V2TaxManualRuleScheduledTaxRateRateParams `form:"rates" json:"rates"`
	// The start time for the tax rate.
	StartsAt *time.Time `form:"starts_at" json:"starts_at,omitempty"`
}

// Creates a ManualRule object.
type V2TaxManualRuleParams struct {
	Params `form:"*"`
	// Location where the rule applies.
	Location *V2TaxManualRuleLocationParams `form:"location" json:"location,omitempty"`
	// Products associated with the rule.
	Products []*V2TaxManualRuleProductParams `form:"products" json:"products,omitempty"`
	// Tax rates to be applied.
	ScheduledTaxRates []*V2TaxManualRuleScheduledTaxRateParams `form:"scheduled_tax_rates" json:"scheduled_tax_rates,omitempty"`
}

// Deactivates a ManualRule object.
type V2TaxManualRuleDeactivateParams struct {
	Params `form:"*"`
}

// Location where the rule applies.
type V2TaxManualRuleCreateLocationParams struct {
	// Country ISO-3166.
	Country *string `form:"country" json:"country"`
	// State ISO-3166.
	State *string `form:"state" json:"state,omitempty"`
}

// Products associated with the rule.
type V2TaxManualRuleCreateProductParams struct {
	// The licensed item identifier.
	LicensedItem *string `form:"licensed_item" json:"licensed_item,omitempty"`
	// The metered item identifier.
	MeteredItem *string `form:"metered_item" json:"metered_item,omitempty"`
	// The tax code for the product.
	TaxCode *string `form:"tax_code" json:"tax_code,omitempty"`
	// The type of the product.
	Type *string `form:"type" json:"type"`
}

// The tax rates to be applied.
type V2TaxManualRuleCreateScheduledTaxRateRateParams struct {
	// Country of the tax rate.
	Country *string `form:"country" json:"country,omitempty"`
	// Description of the tax rate.
	Description *string `form:"description" json:"description,omitempty"`
	// Display name of the tax rate as it will be shown on the invoice.
	DisplayName *string `form:"display_name" json:"display_name"`
	// Jurisdiction of the tax rate should apply as it will be shown on the invoice.
	Jurisdiction *string `form:"jurisdiction" json:"jurisdiction,omitempty"`
	// Percentage of the tax rate. Must be positive and maximum of 4 decimal points.
	Percentage *string `form:"percentage" json:"percentage"`
	// State of the tax rate.
	State *string `form:"state" json:"state,omitempty"`
}

// Tax rates to be applied.
type V2TaxManualRuleCreateScheduledTaxRateParams struct {
	// The tax rates to be applied.
	Rates []*V2TaxManualRuleCreateScheduledTaxRateRateParams `form:"rates" json:"rates"`
	// The start time for the tax rate.
	StartsAt *time.Time `form:"starts_at" json:"starts_at,omitempty"`
}

// Creates a ManualRule object.
type V2TaxManualRuleCreateParams struct {
	Params `form:"*"`
	// Location where the rule applies.
	Location *V2TaxManualRuleCreateLocationParams `form:"location" json:"location,omitempty"`
	// Products associated with the rule.
	Products []*V2TaxManualRuleCreateProductParams `form:"products" json:"products,omitempty"`
	// Tax rates to be applied.
	ScheduledTaxRates []*V2TaxManualRuleCreateScheduledTaxRateParams `form:"scheduled_tax_rates" json:"scheduled_tax_rates"`
}

// Retrieves a ManualRule object by ID.
type V2TaxManualRuleRetrieveParams struct {
	Params `form:"*"`
}

// Location where the rule applies.
type V2TaxManualRuleUpdateLocationParams struct {
	// Country ISO-3166.
	Country *string `form:"country" json:"country"`
	// State ISO-3166.
	State *string `form:"state" json:"state,omitempty"`
}

// Products associated with the rule.
type V2TaxManualRuleUpdateProductParams struct {
	// The licensed item identifier.
	LicensedItem *string `form:"licensed_item" json:"licensed_item,omitempty"`
	// The metered item identifier.
	MeteredItem *string `form:"metered_item" json:"metered_item,omitempty"`
	// The tax code for the product.
	TaxCode *string `form:"tax_code" json:"tax_code,omitempty"`
	// The type of the product.
	Type *string `form:"type" json:"type"`
}

// The tax rates to be applied.
type V2TaxManualRuleUpdateScheduledTaxRateRateParams struct {
	// Country of the tax rate.
	Country *string `form:"country" json:"country,omitempty"`
	// Description of the tax rate.
	Description *string `form:"description" json:"description,omitempty"`
	// Display name of the tax rate as it will be shown on the invoice.
	DisplayName *string `form:"display_name" json:"display_name"`
	// Jurisdiction of the tax rate should apply as it will be shown on the invoice.
	Jurisdiction *string `form:"jurisdiction" json:"jurisdiction,omitempty"`
	// Percentage of the tax rate. Must be positive and maximum of 4 decimal points.
	Percentage *string `form:"percentage" json:"percentage"`
	// State of the tax rate.
	State *string `form:"state" json:"state,omitempty"`
}

// Tax rates to be applied.
type V2TaxManualRuleUpdateScheduledTaxRateParams struct {
	// The tax rates to be applied.
	Rates []*V2TaxManualRuleUpdateScheduledTaxRateRateParams `form:"rates" json:"rates"`
	// The start time for the tax rate.
	StartsAt *time.Time `form:"starts_at" json:"starts_at,omitempty"`
}

// Updates the Tax configuration for a ManualRule object.
type V2TaxManualRuleUpdateParams struct {
	Params `form:"*"`
	// Location where the rule applies.
	Location *V2TaxManualRuleUpdateLocationParams `form:"location" json:"location,omitempty"`
	// Products associated with the rule.
	Products []*V2TaxManualRuleUpdateProductParams `form:"products" json:"products,omitempty"`
	// Tax rates to be applied.
	ScheduledTaxRates []*V2TaxManualRuleUpdateScheduledTaxRateParams `form:"scheduled_tax_rates" json:"scheduled_tax_rates"`
}
