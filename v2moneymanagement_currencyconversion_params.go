//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists all Currency Conversions that match the provided filters.
type V2MoneyManagementCurrencyConversionListParams struct {
	Params `form:"*"`
	// The ID of the Financial Account for which conversions should be listed.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The maximum number of Currency Conversions to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// The source currency to debit, optionally including an amount. Either `from` or `to` must contain a specific amount (but not both).
type V2MoneyManagementCurrencyConversionFromParams struct {
	// An amount specified for conversion.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
}

// The destination currency to credit, optionally including an amount.
type V2MoneyManagementCurrencyConversionToParams struct {
	// An amount specified for conversion.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
}

// Create a new Currency Conversion in a Financial Account or the payments balance.
type V2MoneyManagementCurrencyConversionParams struct {
	Params `form:"*"`
	// The ID of the Financial Account in which the Currency Conversion should be created. If omitted, it will be created in the payments balance.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The source currency to debit, optionally including an amount. Either `from` or `to` must contain a specific amount (but not both).
	From *V2MoneyManagementCurrencyConversionFromParams `form:"from" json:"from,omitempty"`
	// The ID of the [FX Quote](https://docs.stripe.com/api/fx_quotes) to use.
	FxQuote *string `form:"fx_quote" json:"fx_quote,omitempty"`
	// The destination currency to credit, optionally including an amount.
	To *V2MoneyManagementCurrencyConversionToParams `form:"to" json:"to,omitempty"`
}

// The source currency to debit, optionally including an amount. Either `from` or `to` must contain a specific amount (but not both).
type V2MoneyManagementCurrencyConversionCreateFromParams struct {
	// An amount specified for conversion.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
}

// The destination currency to credit, optionally including an amount.
type V2MoneyManagementCurrencyConversionCreateToParams struct {
	// An amount specified for conversion.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
}

// Create a new Currency Conversion in a Financial Account or the payments balance.
type V2MoneyManagementCurrencyConversionCreateParams struct {
	Params `form:"*"`
	// The ID of the Financial Account in which the Currency Conversion should be created. If omitted, it will be created in the payments balance.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The source currency to debit, optionally including an amount. Either `from` or `to` must contain a specific amount (but not both).
	From *V2MoneyManagementCurrencyConversionCreateFromParams `form:"from" json:"from"`
	// The ID of the [FX Quote](https://docs.stripe.com/api/fx_quotes) to use.
	FxQuote *string `form:"fx_quote" json:"fx_quote,omitempty"`
	// The destination currency to credit, optionally including an amount.
	To *V2MoneyManagementCurrencyConversionCreateToParams `form:"to" json:"to"`
}

// Retrieves the details of a Currency Conversion by its ID.
type V2MoneyManagementCurrencyConversionRetrieveParams struct {
	Params `form:"*"`
}
