//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all FinancialAddresses for a FinancialAccount.
type V2MoneyManagementFinancialAddressListParams struct {
	Params `form:"*"`
	// The ID of the FinancialAccount for which FinancialAddresses are to be returned.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Open Enum. A list of fields to reveal in the FinancialAddresses returned.
	Include []*string `form:"include" json:"include,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Properties needed to create a FinancialAddress for an FA with USDC currency.
type V2MoneyManagementFinancialAddressCryptoPropertiesParams struct {
	// The blockchain network of the crypto wallet.
	Network *string `form:"network" json:"network"`
}

// Optional SEPA Bank account options, used to configure the type of SEPA Bank account to create, such as the originating country.
type V2MoneyManagementFinancialAddressSEPABankAccountParams struct {
	// The originating country of the SEPA Bank account.
	Country *string `form:"country" json:"country"`
}

// Create a new FinancialAddress for a FinancialAccount.
type V2MoneyManagementFinancialAddressParams struct {
	Params `form:"*"`
	// Properties needed to create a FinancialAddress for an FA with USDC currency.
	CryptoProperties *V2MoneyManagementFinancialAddressCryptoPropertiesParams `form:"crypto_properties" json:"crypto_properties,omitempty"`
	// The ID of the FinancialAccount the new FinancialAddress should be associated with.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Open Enum. A list of fields to reveal in the FinancialAddresses returned.
	Include []*string `form:"include" json:"include,omitempty"`
	// Optional SEPA Bank account options, used to configure the type of SEPA Bank account to create, such as the originating country.
	SEPABankAccount *V2MoneyManagementFinancialAddressSEPABankAccountParams `form:"sepa_bank_account" json:"sepa_bank_account,omitempty"`
	// Open Enum. The currency the FinancialAddress settles into the FinancialAccount. Currently, only the `usd`, `gbp` and `usdc` values are supported.
	SettlementCurrency *string `form:"settlement_currency" json:"settlement_currency,omitempty"`
	// The type of FinancialAddress details to provision.
	Type *string `form:"type" json:"type,omitempty"`
}

// Properties needed to create a FinancialAddress for an FA with USDC currency.
type V2MoneyManagementFinancialAddressCreateCryptoPropertiesParams struct {
	// The blockchain network of the crypto wallet.
	Network *string `form:"network" json:"network"`
}

// Optional SEPA Bank account options, used to configure the type of SEPA Bank account to create, such as the originating country.
type V2MoneyManagementFinancialAddressCreateSEPABankAccountParams struct {
	// The originating country of the SEPA Bank account.
	Country *string `form:"country" json:"country"`
}

// Create a new FinancialAddress for a FinancialAccount.
type V2MoneyManagementFinancialAddressCreateParams struct {
	Params `form:"*"`
	// Properties needed to create a FinancialAddress for an FA with USDC currency.
	CryptoProperties *V2MoneyManagementFinancialAddressCreateCryptoPropertiesParams `form:"crypto_properties" json:"crypto_properties,omitempty"`
	// The ID of the FinancialAccount the new FinancialAddress should be associated with.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// Optional SEPA Bank account options, used to configure the type of SEPA Bank account to create, such as the originating country.
	SEPABankAccount *V2MoneyManagementFinancialAddressCreateSEPABankAccountParams `form:"sepa_bank_account" json:"sepa_bank_account,omitempty"`
	// Open Enum. The currency the FinancialAddress settles into the FinancialAccount. Currently, only the `usd`, `gbp` and `usdc` values are supported.
	SettlementCurrency *string `form:"settlement_currency" json:"settlement_currency,omitempty"`
	// The type of FinancialAddress details to provision.
	Type *string `form:"type" json:"type"`
}

// Retrieve a FinancialAddress. By default, the FinancialAddress will be returned in its unexpanded state, revealing only the last 4 digits of the account number.
type V2MoneyManagementFinancialAddressRetrieveParams struct {
	Params `form:"*"`
	// Open Enum. A list of fields to reveal in the FinancialAddresses returned.
	Include []*string `form:"include" json:"include,omitempty"`
}
