//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves a list of ReceivedCredits.
type V2MoneyManagementReceivedCreditListParams struct {
	Params `form:"*"`
	// Hash of options for filtering on creation time.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a ReceivedCredit by ID.
type V2MoneyManagementReceivedCreditParams struct {
	Params `form:"*"`
}

// Retrieve a ReceivedCredit by ID.
type V2MoneyManagementReceivedCreditRetrieveParams struct {
	Params `form:"*"`
}
