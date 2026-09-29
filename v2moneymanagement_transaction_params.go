//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of Transactions that match the provided filters.
type V2MoneyManagementTransactionListParams struct {
	Params `form:"*"`
	// Set of filters to query Transactions within a range of `created` timestamps.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// Filter for Transactions belonging to a FinancialAccount.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Filter for Transactions corresponding to a Flow.
	Flow *string `form:"flow" json:"flow,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves the details of a Transaction by ID.
type V2MoneyManagementTransactionParams struct {
	Params `form:"*"`
}

// Retrieves the details of a Transaction by ID.
type V2MoneyManagementTransactionRetrieveParams struct {
	Params `form:"*"`
}
