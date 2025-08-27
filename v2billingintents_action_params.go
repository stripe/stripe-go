//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List Billing Intent Actions.
type V2BillingIntentsActionListParams struct {
	Params `form:"*"`
	// ID of the Billing Intent to list Billing Intent Actions for.
	IntentID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 10.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a Billing Intent Action.
type V2BillingIntentsActionParams struct {
	Params `form:"*"`
	// The ID of the Billing Intent the Billing Intent Action belongs to.
	IntentID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a Billing Intent Action.
type V2BillingIntentsActionRetrieveParams struct {
	Params `form:"*"`
	// The ID of the Billing Intent the Billing Intent Action belongs to.
	IntentID *string `form:"-" json:"-"` // Included in URL
}
