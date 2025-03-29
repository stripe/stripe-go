//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves a list of ReceivedDebits, given the selected filters.
type V2MoneyManagementReceivedDebitListParams struct {
	Params `form:"*"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves a single ReceivedDebit by ID.
type V2MoneyManagementReceivedDebitParams struct {
	Params `form:"*"`
}
