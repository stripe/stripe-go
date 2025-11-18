//
//
// File generated from our OpenAPI spec
//
//

// Package financialaddress provides the financialaddress related APIs
package financialaddress

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke financialaddress related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a new FinancialAddress for a FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementFinancialAddressParams) (*stripe.V2MoneyManagementFinancialAddress, error) {
	financialaddress := &stripe.V2MoneyManagementFinancialAddress{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/financial_addresses", c.Key, params, financialaddress)
	return financialaddress, err
}

// Retrieve a FinancialAddress. By default, the FinancialAddress will be returned in its unexpanded state, revealing only the last 4 digits of the account number.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementFinancialAddressParams) (*stripe.V2MoneyManagementFinancialAddress, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/financial_addresses/%s", id)
	financialaddress := &stripe.V2MoneyManagementFinancialAddress{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaddress)
	return financialaddress, err
}

// List all FinancialAddresses for a FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementFinancialAddressListParams) stripe.Seq2[*stripe.V2MoneyManagementFinancialAddress, error] {
	return stripe.NewV2List("/v2/money_management/financial_addresses", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementFinancialAddress], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementFinancialAddress]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
