//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of FinancialAccounts in your account that match the provided filters.
type V2MoneyManagementFinancialAccountListParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Return FinancialAccounts whose `status` matches any of the provided values.
	// If omitted, closed FinancialAccounts are not returned.
	Statuses []*string `form:"statuses" json:"statuses,omitempty"`
}

// Retrieves the details of an existing FinancialAccount.
type V2MoneyManagementFinancialAccountParams struct {
	Params `form:"*"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name appears in the Stripe Dashboard and embedded components.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to the FinancialAccount. Individual keys can be unset by posting
	// null to them.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementFinancialAccountParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of an existing FinancialAccount.
type V2MoneyManagementFinancialAccountRetrieveParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// Updates an existing FinancialAccount using the provided parameters.
type V2MoneyManagementFinancialAccountUpdateParams struct {
	Params `form:"*"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name appears in the Stripe Dashboard and embedded components.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Set of key-value pairs that you can attach to the FinancialAccount. Individual keys can be unset by posting
	// null to them.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementFinancialAccountUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
