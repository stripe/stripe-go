//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists FinancialAccounts in this compartment.
type V2MoneyManagementFinancialAccountListParams struct {
	Params `form:"*"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The status of the FinancialAccount to filter by. By default, closed FinancialAccounts are not returned.
	Status *string `form:"status" json:"status,omitempty"`
}

// Parameters specific to creating `storage` type FinancialAccounts.
type V2MoneyManagementFinancialAccountStorageParams struct {
	// The currencies that this FinancialAccount can hold.
	HoldsCurrencies []*string `form:"holds_currencies" json:"holds_currencies"`
}

// Creates a new FinancialAccount.
type V2MoneyManagementFinancialAccountParams struct {
	Params `form:"*"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name will be used in the Stripe Dashboard and embedded components.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Metadata associated with the FinancialAccount.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// Parameters specific to creating `storage` type FinancialAccounts.
	Storage *V2MoneyManagementFinancialAccountStorageParams `form:"storage" json:"storage,omitempty"`
	// The type of FinancialAccount to create.
	Type *string `form:"type" json:"type,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementFinancialAccountParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = &value
}

// The addresses to forward any incoming transactions to.
type V2MoneyManagementFinancialAccountCloseForwardingSettingsParams struct {
	// The address to send forwarded payments to.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// The address to send forwarded payouts to.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
}

// Closes a FinancialAccount with or without forwarding settings.
type V2MoneyManagementFinancialAccountCloseParams struct {
	Params `form:"*"`
	// The addresses to forward any incoming transactions to.
	ForwardingSettings *V2MoneyManagementFinancialAccountCloseForwardingSettingsParams `form:"forwarding_settings" json:"forwarding_settings,omitempty"`
}

// Parameters specific to creating `storage` type FinancialAccounts.
type V2MoneyManagementFinancialAccountCreateStorageParams struct {
	// The currencies that this FinancialAccount can hold.
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
}

// Updates an existing FinancialAccount.
type V2MoneyManagementFinancialAccountUpdateParams struct {
	Params `form:"*"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name will be used in the Stripe Dashboard and embedded components.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Metadata associated with the FinancialAccount.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementFinancialAccountUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = &value
}
