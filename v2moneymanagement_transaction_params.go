//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Returns a list of Transactions that match the provided filters.
type V2MoneyManagementTransactionListParams struct {
	Params `form:"*"`
	// Filter for Transactions created at an exact time.
	Created *time.Time `form:"created" json:"created,omitempty"`
	// Filter for Transactions created after the specified timestamp.
	CreatedGt *time.Time `form:"created_gt" json:"created_gt,omitempty"`
	// Filter for Transactions created at or after the specified timestamp.
	CreatedGTE *time.Time `form:"created_gte" json:"created_gte,omitempty"`
	// Filter for Transactions created before the specified timestamp.
	CreatedLT *time.Time `form:"created_lt" json:"created_lt,omitempty"`
	// Filter for Transactions created at or before the specified timestamp.
	CreatedLte *time.Time `form:"created_lte" json:"created_lte,omitempty"`
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
