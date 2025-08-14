//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates an AutomaticRule object.
type V2TaxAutomaticRuleParams struct {
	Params `form:"*"`
	// The BillableItem ID to set automatic Tax configuration for.
	BillableItem *string `form:"billable_item" json:"billable_item,omitempty"`
	// The TaxCode object to be used for automatic tax calculations.
	TaxCode *string `form:"tax_code" json:"tax_code,omitempty"`
}

// Finds an AutomaticRule object by BillableItem ID.
type V2TaxAutomaticRuleFindParams struct {
	Params `form:"*"`
	// The BillableItem ID to search by.
	BillableItem *string `form:"billable_item" json:"billable_item"`
}

// Deactivates an AutomaticRule object. Deactivated AutomaticRule objects are ignored in future tax calculations.
type V2TaxAutomaticRuleDeactivateParams struct {
	Params `form:"*"`
}

// Creates an AutomaticRule object.
type V2TaxAutomaticRuleCreateParams struct {
	Params `form:"*"`
	// The BillableItem ID to set automatic Tax configuration for.
	BillableItem *string `form:"billable_item" json:"billable_item"`
	// The TaxCode object to be used for automatic tax calculations.
	TaxCode *string `form:"tax_code" json:"tax_code"`
}

// Retrieves an AutomaticRule object by ID.
type V2TaxAutomaticRuleRetrieveParams struct {
	Params `form:"*"`
}

// Updates the automatic Tax configuration for an AutomaticRule object.
type V2TaxAutomaticRuleUpdateParams struct {
	Params `form:"*"`
	// The TaxCode object to be used for automatic tax calculations.
	TaxCode *string `form:"tax_code" json:"tax_code"`
}
