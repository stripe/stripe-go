//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List all Spend Modifiers associated with a Billing Cadence.
type V2BillingCadencesSpendModifierRuleListParams struct {
	Params `form:"*"`
	// The ID of the Billing Cadence to list spend modifiers for.
	CadenceID *string `form:"-" json:"-"` // Included in URL
	// Return only spend modifiers that are effective at the specified timestamp.
	EffectiveAt *time.Time `form:"effective_at" json:"effective_at,omitempty"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a Spend Modifier associated with a Billing Cadence.
type V2BillingCadencesSpendModifierRuleParams struct {
	Params `form:"*"`
	// The ID of the Billing Cadence.
	CadenceID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a Spend Modifier associated with a Billing Cadence.
type V2BillingCadencesSpendModifierRuleRetrieveParams struct {
	Params `form:"*"`
	// The ID of the Billing Cadence.
	CadenceID *string `form:"-" json:"-"` // Included in URL
}
