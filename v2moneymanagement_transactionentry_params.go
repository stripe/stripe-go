//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of TransactionEntries that match the provided filters.
type V2MoneyManagementTransactionEntryListParams struct {
	Params `form:"*"`
	// Set of filters to query TransactionEntries within a range of `created` timestamps.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter for TransactionEntries belonging to a Transaction.
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
