//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The account or customer to list signals for. Exactly one of account_details.account or
// account_details.customer must be provided.
type V2SignalsAccountSignalListAccountDetailsParams struct {
	// The v2 account ID of the account.
	Account *string `form:"account" json:"account,omitempty"`
	// The v1 customer ID of the account, for users not yet migrated to v2/accounts.
	Customer *string `form:"customer" json:"customer,omitempty"`
}

// Lists AccountSignals for a given account or customer, newest first. Signals more than 180 days old are omitted.
// By default, returns only the latest matching AccountSignal for each requested signal type.
// Set latest_per_type to false to include matching historical signals. Filtering and latest-per-type selection occur before pagination.
type V2SignalsAccountSignalListParams struct {
	Params `form:"*"`
	// The account or customer to list signals for. Exactly one of account_details.account or
	// account_details.customer must be provided.
	AccountDetails *V2SignalsAccountSignalListAccountDetailsParams `form:"account_details" json:"account_details,omitempty"`
	// Filter signals by their creation timestamp before latest-per-type selection and pagination.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// Whether to return only the latest matching signal for each requested type. Defaults to true.
	// Set to false to include matching historical signals.
	LatestPerType *bool `form:"latest_per_type" json:"latest_per_type,omitempty"`
	// Maximum number of results to return per page. Defaults to 10.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Signal types to filter by.
	Type []*string `form:"type" json:"type"`
}

// Retrieves an AccountSignal by its ID for up to 180 days after creation. Signals more than 180 days old are inaccessible.
type V2SignalsAccountSignalParams struct {
	Params `form:"*"`
}

// Retrieves an AccountSignal by its ID for up to 180 days after creation. Signals more than 180 days old are inaccessible.
type V2SignalsAccountSignalRetrieveParams struct {
	Params `form:"*"`
}
