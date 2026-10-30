//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of TransactionEntries that match the provided filters.
type V2MoneyManagementTransactionEntryListParams struct {
	Params `form:"*"`
	// Filter for TransactionEntries created within a specified time range.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter for TransactionEntries belonging to a specific Transaction. This filter only supports
	// filtering by a Transaction ID, which appears in each TransactionEntry's `transaction` field.
	Transaction *string `form:"transaction" json:"transaction,omitempty"`
}

// Retrieves the details of a TransactionEntry by ID.
type V2MoneyManagementTransactionEntryParams struct {
	Params `form:"*"`
}

// Retrieves the details of a TransactionEntry by ID.
type V2MoneyManagementTransactionEntryRetrieveParams struct {
	Params `form:"*"`
}
