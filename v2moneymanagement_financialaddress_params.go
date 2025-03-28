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
