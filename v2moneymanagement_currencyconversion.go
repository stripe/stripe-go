//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The amount that was debited from the source currency.
type V2MoneyManagementCurrencyConversionFrom struct {
	// A specific amount.
	Amount Amount `json:"amount"`
}

// The amount that was credited to the destination currency.
type V2MoneyManagementCurrencyConversionTo struct {
	// A specific amount.
	Amount Amount `json:"amount"`
}

// A Currency Conversion represents a standalone conversion between currencies in a [Financial Account](https://docs.stripe.com/api/v2/money-management/financial-accounts) or the payments balance.
type V2MoneyManagementCurrencyConversion struct {
	APIResource
	// The time the Currency Conversion was created at.
	Created time.Time `json:"created"`
	// The exchange rate used, excluding applicable fees.
	ExchangeRate string `json:"exchange_rate"`
	// The ID of the Financial Account the Currency Conversion was created in. If `null`, it was created in the payments balance.
	FinancialAccount string `json:"financial_account,omitempty"`
	// The amount that was debited from the source currency.
	From *V2MoneyManagementCurrencyConversionFrom `json:"from"`
	// The ID of the Currency Conversion.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The amount that was credited to the destination currency.
	To *V2MoneyManagementCurrencyConversionTo `json:"to"`
}
