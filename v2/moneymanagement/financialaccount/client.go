//
//
// File generated from our OpenAPI spec
//
//

// Package financialaccount provides the financialaccount related APIs
package financialaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke financialaccount related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an existing FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementFinancialAccountParams) (*stripe.V2MoneyManagementFinancialAccount, error) {
	path := stripe.FormatURLPath("/v2/money_management/financial_accounts/%s", id)
	financialaccount := &stripe.V2MoneyManagementFinancialAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Lists FinancialAccounts in this compartment.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementFinancialAccountListParams) stripe.Seq2[*stripe.V2MoneyManagementFinancialAccount, error] {
	return stripe.NewV2List("/v2/money_management/financial_accounts", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementFinancialAccount], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementFinancialAccount]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
