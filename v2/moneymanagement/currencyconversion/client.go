//
//
// File generated from our OpenAPI spec
//
//

// Package currencyconversion provides the currencyconversion related APIs
package currencyconversion

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke currencyconversion related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a CurrencyConversion.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementCurrencyConversionParams) (*stripe.V2MoneyManagementCurrencyConversion, error) {
	currencyconversion := &stripe.V2MoneyManagementCurrencyConversion{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/currency_conversions", c.Key, params, currencyconversion)
	return currencyconversion, err
}

// Retrieve details of a CurrencyConversion by id.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementCurrencyConversionParams) (*stripe.V2MoneyManagementCurrencyConversion, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/currency_conversions/%s", id)
	currencyconversion := &stripe.V2MoneyManagementCurrencyConversion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, currencyconversion)
	return currencyconversion, err
}

// List all CurrencyConversion on an account with the option to filter by FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementCurrencyConversionListParams) stripe.Seq2[*stripe.V2MoneyManagementCurrencyConversion, error] {
	return stripe.NewV2List("/v2/money_management/currency_conversions", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementCurrencyConversion], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementCurrencyConversion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
