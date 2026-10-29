//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of Transactions that match the provided filters.
type V2MoneyManagementTransactionListParams struct {
	Params `form:"*"`
	// Filter for Transactions created within a specified time range.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// Filter for Transactions belonging to a specific FinancialAccount. This filter only supports filtering by the FinancialAccount's ID.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Filter for Transactions created by a specific flow object. This filter only supports filtering by
	// the flow object's ID, such as `ibt_123` for an InboundTransfer. A value that doesn't match a flow ID
	// returns an empty list.
	Flow *string `form:"flow" json:"flow,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves the details of a Transaction by ID.
type V2MoneyManagementTransactionParams struct {
	Params `form:"*"`
	// A description of the Transaction, up to 100 characters. Pass `null` to clear the description,
	// or omit this parameter to leave it unchanged.
	Description *string `form:"description" json:"description,omitempty"`
	// Set of key-value pairs that you can attach to the Transaction. Individual keys can be unset by posting
	// null to them.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementTransactionParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Creates a fresh hosted URL for a Transaction's regulatory receipt.
type V2MoneyManagementTransactionRefreshRegulatoryReceiptParams struct {
	Params `form:"*"`
}

// Retrieves the details of a Transaction by ID.
type V2MoneyManagementTransactionRetrieveParams struct {
	Params `form:"*"`
}

// Updates an existing Transaction. You can update its description and metadata.
type V2MoneyManagementTransactionUpdateParams struct {
	Params `form:"*"`
	// A description of the Transaction, up to 100 characters. Pass `null` to clear the description,
	// or omit this parameter to leave it unchanged.
	Description *string `form:"description" json:"description,omitempty"`
	// Set of key-value pairs that you can attach to the Transaction. Individual keys can be unset by posting
	// null to them.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementTransactionUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
