//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List quantity changes for a licensed recurring-price term.
type V2BillingContractsTermsQuantityChangeListContractTermQuantityChangesParams struct {
	Params `form:"*"`
	// The ID of a licensed recurring-price term.
	Term *string `form:"-" json:"-"` // Included in URL
	// The contract whose term to query.
	ContractID *string `form:"-" json:"-"` // Included in URL
	// Maximum results per page.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}
