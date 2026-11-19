//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List spend modifiers associated with a billing group.
type V2BillingGroupsSpendModifierRuleListParams struct {
	Params `form:"*"`
	// The ID of the billing group to list spend modifiers for.
	GroupID *string `form:"-" json:"-"` // Included in URL
	// Return only spend modifiers that are effective at the specified timestamp.
	EffectiveAt *time.Time `form:"effective_at" json:"effective_at,omitempty"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a spend modifier associated with a billing group.
type V2BillingGroupsSpendModifierRuleParams struct {
	Params `form:"*"`
	// The ID of the billing group.
	GroupID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a spend modifier associated with a billing group.
type V2BillingGroupsSpendModifierRuleRetrieveParams struct {
	Params `form:"*"`
	// The ID of the billing group.
	GroupID *string `form:"-" json:"-"` // Included in URL
}
