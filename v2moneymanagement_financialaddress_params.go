//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all FinancialAddresses for a FinancialAccount (V2 shape).
type V2MoneyManagementFinancialAddressListParams struct {
	Params `form:"*"`
	// The ID of the Account that owns the FinancialAddresses.
	Account *string `form:"account" json:"account,omitempty"`
	// The ID of the FinancialAccount for which FinancialAddresses are to be returned.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Properties for creating a bank account FinancialAddress.
type V2MoneyManagementFinancialAddressBankAccountParams struct {
	// The country where Stripe provisions the bank account. Each currency supports the country listed in its description. If omitted, Stripe selects that country.
	Country *string `form:"country" json:"country,omitempty"`
	// The currency of the bank account to provision. Some currencies require preview or rollout access, as configured for your account.
	Currency *string `form:"currency" json:"currency"`
}

// Properties for creating a crypto wallet FinancialAddress. Requires `settlement_currency`.
type V2MoneyManagementFinancialAddressCryptoWalletParams struct {
	// The blockchain network of the crypto wallet.
	Network *string `form:"network" json:"network"`
}

// Create a new FinancialAddress for a FinancialAccount (V2 shape).
type V2MoneyManagementFinancialAddressParams struct {
	Params `form:"*"`
	// The ID of the Account that owns this FinancialAddress.
	Account *string `form:"account" json:"account,omitempty"`
	// Properties for creating a bank account FinancialAddress.
	BankAccount *V2MoneyManagementFinancialAddressBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
	// Properties for creating a crypto wallet FinancialAddress. Requires `settlement_currency`.
	CryptoWallet *V2MoneyManagementFinancialAddressCryptoWalletParams `form:"crypto_wallet" json:"crypto_wallet,omitempty"`
	// The ID of the FinancialAccount the new FinancialAddress should be associated with.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Open Enum. The currency the FinancialAddress settles into the FinancialAccount. Optional for `bank_account`, where it defaults to `bank_account.currency`; required for all other types.
	SettlementCurrency *string `form:"settlement_currency" json:"settlement_currency,omitempty"`
	// The type of FinancialAddress to create. Must agree with which branch of financial_address_type_properties is set.
	Type *string `form:"type" json:"type,omitempty"`
}

// Properties for creating a bank account FinancialAddress.
type V2MoneyManagementFinancialAddressCreateBankAccountParams struct {
	// The country where Stripe provisions the bank account. Each currency supports the country listed in its description. If omitted, Stripe selects that country.
	Country *string `form:"country" json:"country,omitempty"`
	// The currency of the bank account to provision. Some currencies require preview or rollout access, as configured for your account.
	Currency *string `form:"currency" json:"currency"`
}

// Properties for creating a crypto wallet FinancialAddress. Requires `settlement_currency`.
type V2MoneyManagementFinancialAddressCreateCryptoWalletParams struct {
	// The blockchain network of the crypto wallet.
	Network *string `form:"network" json:"network"`
}

// Create a new FinancialAddress for a FinancialAccount (V2 shape).
type V2MoneyManagementFinancialAddressCreateParams struct {
	Params `form:"*"`
	// The ID of the Account that owns this FinancialAddress.
	Account *string `form:"account" json:"account,omitempty"`
	// Properties for creating a bank account FinancialAddress.
	BankAccount *V2MoneyManagementFinancialAddressCreateBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
	// Properties for creating a crypto wallet FinancialAddress. Requires `settlement_currency`.
	CryptoWallet *V2MoneyManagementFinancialAddressCreateCryptoWalletParams `form:"crypto_wallet" json:"crypto_wallet,omitempty"`
	// The ID of the FinancialAccount the new FinancialAddress should be associated with.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// Open Enum. The currency the FinancialAddress settles into the FinancialAccount. Optional for `bank_account`, where it defaults to `bank_account.currency`; required for all other types.
	SettlementCurrency *string `form:"settlement_currency" json:"settlement_currency,omitempty"`
	// The type of FinancialAddress to create. Must agree with which branch of financial_address_type_properties is set.
	Type *string `form:"type" json:"type"`
}

// Retrieve a FinancialAddress (V2 shape).
type V2MoneyManagementFinancialAddressRetrieveParams struct {
	Params `form:"*"`
}
