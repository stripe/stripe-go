//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all CurrencyConversion on an account with the option to filter by FinancialAccount.
type V2MoneyManagementCurrencyConversionListParams struct {
	Params `form:"*"`
	// The ID of the FinancialAccount to filter by.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Amount object.
type V2MoneyManagementCurrencyConversionFromAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// From amount object indicating the from currency or optional amount.
type V2MoneyManagementCurrencyConversionFromParams struct {
	// Amount object.
	Amount *V2MoneyManagementCurrencyConversionFromAmountParams `form:"amount" json:"amount,omitempty"`
	// A lowercase alpha3 currency code like "usd".
	Currency *string `form:"currency" json:"currency,omitempty"`
}

// Amount object.
type V2MoneyManagementCurrencyConversionToAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// To amount object indicating the to currency or optional amount.
type V2MoneyManagementCurrencyConversionToParams struct {
	// Amount object.
	Amount *V2MoneyManagementCurrencyConversionToAmountParams `form:"amount" json:"amount,omitempty"`
	// A lowercase alpha3 currency code like "usd".
	Currency *string `form:"currency" json:"currency,omitempty"`
}

// Create a CurrencyConversion.
type V2MoneyManagementCurrencyConversionParams struct {
	Params `form:"*"`
	// The FinancialAccount id to create the CurrencyConversion on.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// From amount object indicating the from currency or optional amount.
	From *V2MoneyManagementCurrencyConversionFromParams `form:"from" json:"from,omitempty"`
	// To amount object indicating the to currency or optional amount.
	To *V2MoneyManagementCurrencyConversionToParams `form:"to" json:"to,omitempty"`
}

// Amount object.
type V2MoneyManagementCurrencyConversionCreateFromAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// From amount object indicating the from currency or optional amount.
type V2MoneyManagementCurrencyConversionCreateFromParams struct {
	// Amount object.
	Amount *V2MoneyManagementCurrencyConversionCreateFromAmountParams `form:"amount" json:"amount,omitempty"`
	// A lowercase alpha3 currency code like "usd".
	Currency *string `form:"currency" json:"currency,omitempty"`
}

// Amount object.
type V2MoneyManagementCurrencyConversionCreateToAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// To amount object indicating the to currency or optional amount.
type V2MoneyManagementCurrencyConversionCreateToParams struct {
	// Amount object.
	Amount *V2MoneyManagementCurrencyConversionCreateToAmountParams `form:"amount" json:"amount,omitempty"`
	// A lowercase alpha3 currency code like "usd".
	Currency *string `form:"currency" json:"currency,omitempty"`
}

// Create a CurrencyConversion.
type V2MoneyManagementCurrencyConversionCreateParams struct {
	Params `form:"*"`
	// The FinancialAccount id to create the CurrencyConversion on.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// From amount object indicating the from currency or optional amount.
	From *V2MoneyManagementCurrencyConversionCreateFromParams `form:"from" json:"from"`
	// To amount object indicating the to currency or optional amount.
	To *V2MoneyManagementCurrencyConversionCreateToParams `form:"to" json:"to"`
}

// Retrieve details of a CurrencyConversion by id.
type V2MoneyManagementCurrencyConversionRetrieveParams struct {
	Params `form:"*"`
}
