//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Amount object.
type V2MoneyManagementCurrencyConversionFromAmount struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// The from block containing what was debited.
type V2MoneyManagementCurrencyConversionFrom struct {
	// Amount object.
	Amount *V2MoneyManagementCurrencyConversionFromAmount `json:"amount"`
}

// Amount object.
type V2MoneyManagementCurrencyConversionToAmount struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// The to block containing what was credited.
type V2MoneyManagementCurrencyConversionTo struct {
	// Amount object.
	Amount *V2MoneyManagementCurrencyConversionToAmount `json:"amount"`
}

// The CurrencyConversion object. Contains details such as the amount debited and credited and the FinancialAccount the CurrencyConversion was performed on.
type V2MoneyManagementCurrencyConversion struct {
	APIResource
	// The time the CurrencyConversion was created at.
	Created time.Time `json:"created"`
	// The exchange rate used when processing the CurrencyConversion.
	ExchangeRate string `json:"exchange_rate"`
	// The FinancialAccount the CurrencyConversion was performed on.
	FinancialAccount string `json:"financial_account"`
	// The from block containing what was debited.
	From *V2MoneyManagementCurrencyConversionFrom `json:"from"`
	// The id of the CurrencyConversion.
	ID string `json:"id"`
	// If the CurrencyConversion was performed in livemode or not.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The to block containing what was credited.
	To *V2MoneyManagementCurrencyConversionTo `json:"to"`
}
