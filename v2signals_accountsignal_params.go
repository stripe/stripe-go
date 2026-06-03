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

// Lists AccountSignals for a given account or customer, filtered by signal type.
type V2SignalsAccountSignalListParams struct {
	Params `form:"*"`
	// The account or customer to list signals for. Exactly one of account_details.account or
	// account_details.customer must be provided.
	AccountDetails *V2SignalsAccountSignalListAccountDetailsParams `form:"account_details" json:"account_details,omitempty"`
	// Maximum number of results to return per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Signal types to filter by.
	Type []*string `form:"type" json:"type"`
}

// Retrieves an AccountSignal by its ID.
type V2SignalsAccountSignalParams struct {
	Params `form:"*"`
}

// Retrieves an AccountSignal by its ID.
type V2SignalsAccountSignalRetrieveParams struct {
	Params `form:"*"`
}
