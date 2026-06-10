//
//
// File generated from our OpenAPI spec
//
//

// Package statement provides the statement related APIs
package statement

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke statement related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Financial Account Statement.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementFinancialAccountsStatementParams) (*stripe.V2MoneyManagementFinancialAccountStatement, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/financial_accounts/%s/statements/%s", stripe.StringValue(
			params.FinancialAccountID), id)
	financialaccountstatement := &stripe.V2MoneyManagementFinancialAccountStatement{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, financialaccountstatement)
	return financialaccountstatement, err
}

// Returns a list of statements for a Financial Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementFinancialAccountsStatementListParams) stripe.Seq2[*stripe.V2MoneyManagementFinancialAccountStatement, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementFinancialAccountsStatementListParams{}
	}
	path := stripe.FormatURLPath(
		"/v2/money_management/financial_accounts/%s/statements", stripe.StringValue(
			listParams.FinancialAccountID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementFinancialAccountStatement], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementFinancialAccountStatement]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
