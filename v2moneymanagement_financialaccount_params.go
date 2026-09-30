//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists FinancialAccounts in this compartment.
type V2MoneyManagementFinancialAccountListParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter for FinancialAccount `status`. By default, closed FinancialAccounts are not returned.
	Statuses []*string `form:"statuses" json:"statuses,omitempty"`
	// Filter for FinancialAccount `type`. By default, FinancialAccounts of any `type` are returned.
	Types []*string `form:"types" json:"types,omitempty"`
}

// Parameters specific to creating `savings` type FinancialAccounts.
type V2MoneyManagementFinancialAccountSavingsParams struct {
	// The currencies that this savings FinancialAccount can hold. Three-letter ISO currency code, in lowercase.
	HoldsCurrencies []*string `form:"holds_currencies" json:"holds_currencies"`
}

// Crypto-specific storage configuration. Only populated when `storage.crypto` is passed in the `include` parameter and the FinancialAccount stores crypto assets. Fiat currencies remain configured only through `holds_currencies`.
type V2MoneyManagementFinancialAccountStorageCryptoParams struct {
	// The blockchain network configured for each crypto currency. Keys are lowercase currency codes and must identify crypto currencies also present in `holds_currencies`.
	CurrencyNetworks map[string]string `form:"currency_networks" json:"currency_networks"`
	// Describes who controls the private keys for the crypto storage.
	CustodyModel *string `form:"custody_model" json:"custody_model"`
}

// Array of eligibility objects, segmented by bank name and deposit insurance scheme.
type V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityParams struct {
	// The bank where funds are stored.
	BankName *string `form:"bank_name" json:"bank_name"`
	// Currencies eligible for deposit insurance at this bank under this scheme.
	Currencies []*string `form:"currencies" json:"currencies"`
	// The deposit insurance scheme.
	Type *string `form:"type" json:"type"`
}

// Parameters specific to creating `storage` type FinancialAccounts.
type V2MoneyManagementFinancialAccountStorageParams struct {
	// Crypto-specific storage configuration. Only populated when `storage.crypto` is passed in the `include` parameter and the FinancialAccount stores crypto assets. Fiat currencies remain configured only through `holds_currencies`.
	Crypto *V2MoneyManagementFinancialAccountStorageCryptoParams `form:"crypto" json:"crypto,omitempty"`
	// Array of eligibility objects, segmented by bank name and deposit insurance scheme.
	DepositInsuranceEligibility []*V2MoneyManagementFinancialAccountStorageDepositInsuranceEligibilityParams `form:"deposit_insurance_eligibility" json:"deposit_insurance_eligibility,omitempty"`
	// The usage type for funds in this FinancialAccount. Can be used to specify that the funds are for Consumer activity.
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
	// Metadata associated with the FinancialAccount.
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
	// The address to send forwarded payments to.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// The address to send forwarded payouts to.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Whether to skip forwarding exportable self-custodied wallet balances. Defaults to false. This does not skip non-exportable or fiat balances, inbound-pending checks, or negative-balance requirements.
	SkipExportableBalances *bool `form:"skip_exportable_balances" json:"skip_exportable_balances,omitempty"`
}

// The addresses to forward any incoming transactions to.
type V2MoneyManagementFinancialAccountCloseForwardingSettingsParams struct {
	// The address to send forwarded payments to.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// The address to send forwarded payouts to.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Whether to skip forwarding exportable self-custodied wallet balances. Defaults to false. This does not skip non-exportable or fiat balances, inbound-pending checks, or negative-balance requirements.
	SkipExportableBalances *bool `form:"skip_exportable_balances" json:"skip_exportable_balances,omitempty"`
}

// Closes a FinancialAccount with or without forwarding settings.
type V2MoneyManagementFinancialAccountCloseParams struct {
	Params `form:"*"`
	// The addresses to forward any incoming transactions to.
	ForwardingSettings *V2MoneyManagementFinancialAccountCloseForwardingSettingsParams `form:"forwarding_settings" json:"forwarding_settings,omitempty"`
}

// Parameters specific to creating `savings` type FinancialAccounts.
type V2MoneyManagementFinancialAccountCreateSavingsParams struct {
	// The currencies that this savings FinancialAccount can hold. Three-letter ISO currency code, in lowercase.
	HoldsCurrencies []*string `form:"holds_currencies" json:"holds_currencies"`
}

// Crypto-specific storage configuration. Only populated when `storage.crypto` is passed in the `include` parameter and the FinancialAccount stores crypto assets. Fiat currencies remain configured only through `holds_currencies`.
type V2MoneyManagementFinancialAccountCreateStorageCryptoParams struct {
	// The blockchain network configured for each crypto currency. Keys are lowercase currency codes and must identify crypto currencies also present in `holds_currencies`.
	CurrencyNetworks map[string]string `form:"currency_networks" json:"currency_networks"`
	// Describes who controls the private keys for the crypto storage.
	CustodyModel *string `form:"custody_model" json:"custody_model"`
}

// Array of eligibility objects, segmented by bank name and deposit insurance scheme.
type V2MoneyManagementFinancialAccountCreateStorageDepositInsuranceEligibilityParams struct {
	// The bank where funds are stored.
	BankName *string `form:"bank_name" json:"bank_name"`
	// Currencies eligible for deposit insurance at this bank under this scheme.
	Currencies []*string `form:"currencies" json:"currencies"`
	// The deposit insurance scheme.
	Type *string `form:"type" json:"type"`
}

// Parameters specific to creating `storage` type FinancialAccounts.
type V2MoneyManagementFinancialAccountCreateStorageParams struct {
	// Crypto-specific storage configuration. Only populated when `storage.crypto` is passed in the `include` parameter and the FinancialAccount stores crypto assets. Fiat currencies remain configured only through `holds_currencies`.
	Crypto *V2MoneyManagementFinancialAccountCreateStorageCryptoParams `form:"crypto" json:"crypto,omitempty"`
	// Array of eligibility objects, segmented by bank name and deposit insurance scheme.
	DepositInsuranceEligibility []*V2MoneyManagementFinancialAccountCreateStorageDepositInsuranceEligibilityParams `form:"deposit_insurance_eligibility" json:"deposit_insurance_eligibility,omitempty"`
	// The usage type for funds in this FinancialAccount. Can be used to specify that the funds are for Consumer activity.
	FundsUsageType *string `form:"funds_usage_type" json:"funds_usage_type,omitempty"`
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
	// The address to send forwarded payments to.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// The address to send forwarded payouts to.
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

// Updates an existing FinancialAccount.
type V2MoneyManagementFinancialAccountUpdateParams struct {
	Params `form:"*"`
	// A descriptive name for the FinancialAccount, up to 50 characters long. This name will be used in the Stripe Dashboard and embedded components.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Forwarding settings for a closed FinancialAccount. Post-close forwarding updates are not yet implemented.
	ForwardingSettings *V2MoneyManagementFinancialAccountUpdateForwardingSettingsParams `form:"forwarding_settings" json:"forwarding_settings,omitempty"`
	// Metadata associated with the FinancialAccount.
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
