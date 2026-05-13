//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of statements for a Financial Account.
type V2MoneyManagementFinancialAccountsStatementListParams struct {
	Params `form:"*"`
	// The ID of the Financial Account to list statements for.
	FinancialAccountID *string `form:"-" json:"-"` // Included in URL
	// The maximum number of results to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The field by which to sort results. Defaults to "created".
	OrderBy *string `form:"order_by" json:"order_by,omitempty"`
}

// Retrieves the details of a Financial Account Statement.
type V2MoneyManagementFinancialAccountsStatementParams struct {
	Params `form:"*"`
	// The ID of the Financial Account.
	FinancialAccountID *string `form:"-" json:"-"` // Included in URL
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// Retrieves the details of a Financial Account Statement.
type V2MoneyManagementFinancialAccountsStatementRetrieveParams struct {
	Params `form:"*"`
	// The ID of the Financial Account.
	FinancialAccountID *string `form:"-" json:"-"` // Included in URL
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}
