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
	// Filter for FinancialAccount `type`. By default, FinancialAccounts of any `type` are returned.
	Types []*string `form:"types" json:"types,omitempty"`
}

// Parameters specific to creating `savings` type FinancialAccounts.
type V2MoneyManagementFinancialAccountSavingsParams struct {
	// The currencies that this savings FinancialAccount can hold. Three-letter ISO currency code, in lowercase.
	HoldsCurrencies []*string `form:"holds_currencies" json:"holds_currencies"`
}

// Crypto-specific storage configuration. Fiat currencies remain configured only through `holds_currencies`.
type V2MoneyManagementFinancialAccountStorageCryptoParams struct {
	// The blockchain network configured for each crypto currency. Keys are lowercase currency codes and must identify crypto currencies also present in `holds_currencies`.
	CurrencyNetworks map[string]string `form:"currency_networks" json:"currency_networks"`
	// Describes who controls the private keys for the crypto storage.
	CustodyModel *string `form:"custody_model" json:"custody_model"`
}

// Parameters specific to creating `storage` type FinancialAccounts.
type V2MoneyManagementFinancialAccountStorageParams struct {
	// Crypto-specific storage configuration used when adding crypto to a fiat-only FinancialAccount.
	// `custody_model` is required for the initial crypto update and cannot be changed afterward.
	Crypto *V2MoneyManagementFinancialAccountStorageCryptoParams `form:"crypto" json:"crypto,omitempty"`
	// The usage type for funds in this FinancialAccount. Defaults to BUSINESS if not specified.
	FundsUsageType *string `form:"funds_usage_type" json:"funds_usage_type,omitempty"`
	// The currencies that this storage FinancialAccount can hold a balance in. Three-letter ISO currency code, in lowercase.
	// Adding currencies requires the corresponding holds_currencies storer capabilities to be enabled.
	// Removing currencies is not supported as of March 2026.
	HoldsCurrencies []*string `form:"holds_currencies" json:"holds_currencies,omitempty"`
}

// Creates a new FinancialAccount.
type V2MoneyManagementFinancialAccountParams struct {
	Params `form:"*"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name will be used in the Stripe Dashboard and embedded components.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Forwarding settings for a closed FinancialAccount. Post-close forwarding updates are not yet implemented.
	ForwardingSettings *V2MoneyManagementFinancialAccountForwardingSettingsParams `form:"forwarding_settings" json:"forwarding_settings,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to the FinancialAccount. Individual keys can be unset by posting
	// null to them.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// Parameters specific to creating `savings` type FinancialAccounts.
	Savings *V2MoneyManagementFinancialAccountSavingsParams `form:"savings" json:"savings,omitempty"`
	// Parameters for updating storage-specific fields on the FinancialAccount.
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

// Forwarding settings for a closed FinancialAccount. Post-close forwarding updates are not yet implemented.
type V2MoneyManagementFinancialAccountForwardingSettingsParams struct {
	// ID of the FinancialAccount that supplies funds to cover negative balances
	// during or after this FinancialAccount's closure.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// ID of the FinancialAccount that receives funds forwarded from this FinancialAccount when it closes.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Whether to skip forwarding exportable self-custodied wallet balances. Defaults to false. This does not skip non-exportable or fiat balances, inbound-pending checks, or negative-balance requirements.
	SkipExportableBalances *bool `form:"skip_exportable_balances" json:"skip_exportable_balances,omitempty"`
}

// Settings that specify where to send positive balances and how to cover negative balances.
// You must provide these settings if the FinancialAccount has a non-zero balance.
type V2MoneyManagementFinancialAccountCloseForwardingSettingsParams struct {
	// ID of the FinancialAccount that supplies funds to cover negative balances
	// during or after this FinancialAccount's closure.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// ID of the FinancialAccount that receives funds forwarded from this FinancialAccount when it closes.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Whether to skip forwarding exportable self-custodied wallet balances. Defaults to false. This does not skip non-exportable or fiat balances, inbound-pending checks, or negative-balance requirements.
	SkipExportableBalances *bool `form:"skip_exportable_balances" json:"skip_exportable_balances,omitempty"`
}

// Closes a FinancialAccount. You must provide forwarding settings if it has a non-zero balance.
type V2MoneyManagementFinancialAccountCloseParams struct {
	Params `form:"*"`
	// Settings that specify where to send positive balances and how to cover negative balances.
	// You must provide these settings if the FinancialAccount has a non-zero balance.
	ForwardingSettings *V2MoneyManagementFinancialAccountCloseForwardingSettingsParams `form:"forwarding_settings" json:"forwarding_settings,omitempty"`
}

// Parameters specific to creating `savings` type FinancialAccounts.
type V2MoneyManagementFinancialAccountCreateSavingsParams struct {
	// The currencies that this savings FinancialAccount can hold. Three-letter ISO currency code, in lowercase.
	HoldsCurrencies []*string `form:"holds_currencies" json:"holds_currencies"`
}

// Crypto-specific storage configuration. Fiat currencies remain configured only through `holds_currencies`.
type V2MoneyManagementFinancialAccountCreateStorageCryptoParams struct {
	// The blockchain network configured for each crypto currency. Keys are lowercase currency codes and must identify crypto currencies also present in `holds_currencies`.
	CurrencyNetworks map[string]string `form:"currency_networks" json:"currency_networks"`
	// Describes who controls the private keys for the crypto storage.
	CustodyModel *string `form:"custody_model" json:"custody_model"`
}

// Parameters specific to creating `storage` type FinancialAccounts.
type V2MoneyManagementFinancialAccountCreateStorageParams struct {
	// Crypto-specific storage configuration. Fiat currencies remain configured only through `holds_currencies`.
	Crypto *V2MoneyManagementFinancialAccountCreateStorageCryptoParams `form:"crypto" json:"crypto,omitempty"`
	// The usage type for funds in this FinancialAccount. Defaults to BUSINESS if not specified.
	FundsUsageType *string `form:"funds_usage_type" json:"funds_usage_type,omitempty"`
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
	// Parameters specific to creating `savings` type FinancialAccounts.
	Savings *V2MoneyManagementFinancialAccountCreateSavingsParams `form:"savings" json:"savings,omitempty"`
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

// Forwarding settings for a closed FinancialAccount. Post-close forwarding updates are not yet implemented.
type V2MoneyManagementFinancialAccountUpdateForwardingSettingsParams struct {
	// ID of the FinancialAccount that supplies funds to cover negative balances
	// during or after this FinancialAccount's closure.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// ID of the FinancialAccount that receives funds forwarded from this FinancialAccount when it closes.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Whether to skip forwarding exportable self-custodied wallet balances. Defaults to false. This does not skip non-exportable or fiat balances, inbound-pending checks, or negative-balance requirements.
	SkipExportableBalances *bool `form:"skip_exportable_balances" json:"skip_exportable_balances,omitempty"`
}

// Crypto-specific storage configuration used when adding crypto to a fiat-only FinancialAccount.
// `custody_model` is required for the initial crypto update and cannot be changed afterward.
type V2MoneyManagementFinancialAccountUpdateStorageCryptoParams struct {
	// The blockchain network configured for each crypto currency. Keys are lowercase currency codes and must identify crypto currencies also present in `holds_currencies`.
	CurrencyNetworks map[string]string `form:"currency_networks" json:"currency_networks"`
	// Describes who controls the private keys for the crypto storage.
	CustodyModel *string `form:"custody_model" json:"custody_model"`
}

// Parameters for updating storage-specific fields on the FinancialAccount.
type V2MoneyManagementFinancialAccountUpdateStorageParams struct {
	// Crypto-specific storage configuration used when adding crypto to a fiat-only FinancialAccount.
	// `custody_model` is required for the initial crypto update and cannot be changed afterward.
	Crypto *V2MoneyManagementFinancialAccountUpdateStorageCryptoParams `form:"crypto" json:"crypto,omitempty"`
	// The currencies that this storage FinancialAccount can hold a balance in. Three-letter ISO currency code, in lowercase.
	// Adding currencies requires the corresponding holds_currencies storer capabilities to be enabled.
	// Removing currencies is not supported as of March 2026.
	HoldsCurrencies []*string `form:"holds_currencies" json:"holds_currencies,omitempty"`
}

// Updates an existing FinancialAccount using the provided parameters.
type V2MoneyManagementFinancialAccountUpdateParams struct {
	Params `form:"*"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name appears in the Stripe Dashboard and embedded components.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Forwarding settings for a closed FinancialAccount. Post-close forwarding updates are not yet implemented.
	ForwardingSettings *V2MoneyManagementFinancialAccountUpdateForwardingSettingsParams `form:"forwarding_settings" json:"forwarding_settings,omitempty"`
	// Set of key-value pairs that you can attach to the FinancialAccount. Individual keys can be unset by posting
	// null to them.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// Parameters for updating storage-specific fields on the FinancialAccount.
	Storage *V2MoneyManagementFinancialAccountUpdateStorageParams `form:"storage" json:"storage,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementFinancialAccountUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
