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
	// Filter by the received debit mandate ID. Deprecated: use `spending_mandate` instead.
	ReceivedDebitMandate *string `form:"received_debit_mandate" json:"received_debit_mandate,omitempty"`
	// Filter by the SpendingMandate ID.
	SpendingMandate *string `form:"spending_mandate" json:"spending_mandate,omitempty"`
}

// Retrieves a single ReceivedDebit by ID.
type V2MoneyManagementReceivedDebitParams struct {
	Params `form:"*"`
}

// Retrieves a single ReceivedDebit by ID.
type V2MoneyManagementReceivedDebitRetrieveParams struct {
	Params `form:"*"`
}
