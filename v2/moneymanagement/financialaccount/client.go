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
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an existing FinancialAccount.
func (c Client) Get(id string, params *stripe.V2MoneyManagementFinancialAccountParams) (*stripe.V2MoneyManagementFinancialAccount, error) {
	path := stripe.FormatURLPath("/v2/money_management/financial_accounts/%s", id)
	financialaccount := &stripe.V2MoneyManagementFinancialAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Lists FinancialAccounts in this compartment.
func (c Client) All(listParams *stripe.V2MoneyManagementFinancialAccountListParams) stripe.Seq2[*stripe.V2MoneyManagementFinancialAccount, error] {
	return stripe.NewV2List("/v2/money_management/financial_accounts", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementFinancialAccount], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementFinancialAccount]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
