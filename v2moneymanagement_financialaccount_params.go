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

// Parameters specific to creating `storage` type FinancialAccounts.
type V2MoneyManagementFinancialAccountStorageParams struct {
	// The currencies that the storage FinancialAccount can hold a balance in. Three-letter ISO currency code, in lowercase.
	// Must be a subset of the currencies requested in `configuration.storer.features.holds_currencies` of the v2/accounts API.
	HoldsCurrencies []*string `form:"holds_currencies" json:"holds_currencies"`
}

// Creates a new FinancialAccount.
type V2MoneyManagementFinancialAccountParams struct {
	Params `form:"*"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name will be used in the Stripe Dashboard and embedded components.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to the FinancialAccount. Individual keys can be unset by posting
	// null to them.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// Parameters specific to creating `storage` type FinancialAccounts.
	Storage *V2MoneyManagementFinancialAccountStorageParams `form:"storage" json:"storage,omitempty"`
	// The type of FinancialAccount to create.
	Type *string `form:"type" json:"type,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementFinancialAccountParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Settings that specify where to send positive balances and how to cover negative balances.
// You must provide these settings if the FinancialAccount has a non-zero balance.
type V2MoneyManagementFinancialAccountCloseForwardingSettingsParams struct {
	// ID of the FinancialAccount that supplies funds to cover negative balances
	// during or after this FinancialAccount's closure.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// ID of the FinancialAccount that receives funds forwarded from this FinancialAccount when it closes.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
}

// Closes a FinancialAccount. You must provide forwarding settings if it has a non-zero balance.
type V2MoneyManagementFinancialAccountCloseParams struct {
	Params `form:"*"`
	// Settings that specify where to send positive balances and how to cover negative balances.
	// You must provide these settings if the FinancialAccount has a non-zero balance.
	ForwardingSettings *V2MoneyManagementFinancialAccountCloseForwardingSettingsParams `form:"forwarding_settings" json:"forwarding_settings,omitempty"`
}

// Parameters specific to creating `storage` type FinancialAccounts.
type V2MoneyManagementFinancialAccountCreateStorageParams struct {
	// The currencies that the storage FinancialAccount can hold a balance in. Three-letter ISO currency code, in lowercase.
	// Must be a subset of the currencies requested in `configuration.storer.features.holds_currencies` of the v2/accounts API.
	HoldsCurrencies []*string `form:"holds_currencies" json:"holds_currencies"`
}

// Creates a new FinancialAccount.
type V2MoneyManagementFinancialAccountCreateParams struct {
	Params `form:"*"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name will be used in the Stripe Dashboard and embedded components.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Metadata associated with the FinancialAccount.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Parameters specific to creating `storage` type FinancialAccounts.
	Storage *V2MoneyManagementFinancialAccountCreateStorageParams `form:"storage" json:"storage,omitempty"`
	// The type of FinancialAccount to create.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementFinancialAccountCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
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
