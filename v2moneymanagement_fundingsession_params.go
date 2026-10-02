//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Options for a crypto wallet FinancialAddress. Required if `crypto_wallet` is requested.
type V2MoneyManagementFundingSessionFinancialAddressOptionsCryptoWalletParams struct {
	// Open Enum. The currency the crypto wallet FinancialAddress settles into the FinancialAccount. Required.
	SettlementCurrency *string `form:"settlement_currency" json:"settlement_currency"`
}

// Per-type options used when creating the FinancialAddress. Required.
type V2MoneyManagementFundingSessionFinancialAddressOptionsParams struct {
	// Options for a crypto wallet FinancialAddress. Required if `crypto_wallet` is requested.
	CryptoWallet *V2MoneyManagementFundingSessionFinancialAddressOptionsCryptoWalletParams `form:"crypto_wallet" json:"crypto_wallet,omitempty"`
}

// Create a FundingSession: a hosted funding surface for a customer to fund a FinancialAccount.
type V2MoneyManagementFundingSessionParams struct {
	Params `form:"*"`
	// The ID of the Account that owns the FinancialAccount. Required.
	Account *string `form:"account" json:"account"`
	// The ID of the FinancialAccount to fund. Required.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// Per-type options used when creating the FinancialAddress. Required.
	FinancialAddressOptions *V2MoneyManagementFundingSessionFinancialAddressOptionsParams `form:"financial_address_options" json:"financial_address_options"`
	// Open Enum. The types of FinancialAddress that can be funded in this session. At least one is required.
	FinancialAddressTypes []*string `form:"financial_address_types" json:"financial_address_types"`
	// The URL the customer is redirected to after completing or abandoning the funding session. Required.
	ReturnURL *string `form:"return_url" json:"return_url"`
}

// Options for a crypto wallet FinancialAddress. Required if `crypto_wallet` is requested.
type V2MoneyManagementFundingSessionCreateFinancialAddressOptionsCryptoWalletParams struct {
	// Open Enum. The currency the crypto wallet FinancialAddress settles into the FinancialAccount. Required.
	SettlementCurrency *string `form:"settlement_currency" json:"settlement_currency"`
}

// Per-type options used when creating the FinancialAddress. Required.
type V2MoneyManagementFundingSessionCreateFinancialAddressOptionsParams struct {
	// Options for a crypto wallet FinancialAddress. Required if `crypto_wallet` is requested.
	CryptoWallet *V2MoneyManagementFundingSessionCreateFinancialAddressOptionsCryptoWalletParams `form:"crypto_wallet" json:"crypto_wallet,omitempty"`
}

// Create a FundingSession: a hosted funding surface for a customer to fund a FinancialAccount.
type V2MoneyManagementFundingSessionCreateParams struct {
	Params `form:"*"`
	// The ID of the Account that owns the FinancialAccount. Required.
	Account *string `form:"account" json:"account"`
	// The ID of the FinancialAccount to fund. Required.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// Per-type options used when creating the FinancialAddress. Required.
	FinancialAddressOptions *V2MoneyManagementFundingSessionCreateFinancialAddressOptionsParams `form:"financial_address_options" json:"financial_address_options"`
	// Open Enum. The types of FinancialAddress that can be funded in this session. At least one is required.
	FinancialAddressTypes []*string `form:"financial_address_types" json:"financial_address_types"`
	// The URL the customer is redirected to after completing or abandoning the funding session. Required.
	ReturnURL *string `form:"return_url" json:"return_url"`
}
