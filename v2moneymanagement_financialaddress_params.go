//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Create a new FinancialAddress for a FinancialAccount.
type V2MoneyManagementFinancialAddressParams struct {
	Params `form:"*"`
	// Open Enum. The currency the FinancialAddress should support. Currently, only the `usd` and `gbp` values are supported.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The ID of the FinancialAccount the new FinancialAddress should be associated with.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Open Enum. A list of fields to reveal in the FinancialAddresses returned.
	Include []*string `form:"include" json:"include,omitempty"`
}

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

// Create a new FinancialAddress for a FinancialAccount.
type V2MoneyManagementFinancialAddressCreateParams struct {
	Params `form:"*"`
	// Open Enum. The currency the FinancialAddress should support. Currently, only the `usd` and `gbp` values are supported.
	Currency *string `form:"currency" json:"currency"`
	// The ID of the FinancialAccount the new FinancialAddress should be associated with.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Retrieve a FinancialAddress. By default, the FinancialAddress will be returned in it's unexpanded state, revealing only the last 4 digits of the account number.
type V2MoneyManagementFinancialAddressRetrieveParams struct {
	Params `form:"*"`
	// Open Enum. A list of fields to reveal in the FinancialAddresses returned.
	Include []*string `form:"include" json:"include,omitempty"`
}
