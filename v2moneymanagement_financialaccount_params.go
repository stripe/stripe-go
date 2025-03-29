//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists FinancialAccounts in this compartment.
type V2MoneyManagementFinancialAccountListParams struct {
	Params `form:"*"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves the details of an existing FinancialAccount.
type V2MoneyManagementFinancialAccountParams struct {
	Params `form:"*"`
}
