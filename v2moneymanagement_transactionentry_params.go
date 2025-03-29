//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Returns a list of TransactionEntries that match the provided filters.
type V2MoneyManagementTransactionEntryListParams struct {
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
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter for TransactionEntries belonging to a Transaction.
	Transaction *string `form:"transaction" json:"transaction,omitempty"`
}

// Retrieves the details of a TransactionEntry by ID.
type V2MoneyManagementTransactionEntryParams struct {
	Params `form:"*"`
}
