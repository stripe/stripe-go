//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List the terms on a contract, optionally filtered by type.
type V2BillingContractsTermListContractTermsParams struct {
	Params `form:"*"`
	// The contract whose terms to list.
	ContractID *string `form:"-" json:"-"` // Included in URL
	// Maximum results per page.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter terms by type.
	Type *string `form:"type" json:"type,omitempty"`
}

// Retrieve a term on a contract.
type V2BillingContractsTermRetrieveContractTermParams struct {
	Params `form:"*"`
	// The contract containing the term.
	ContractID *string `form:"-" json:"-"` // Included in URL
}
