//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2MoneyManagementFinancialAccountService is used to invoke financialaccount related APIs.
type v2MoneyManagementFinancialAccountService struct {
	B   Backend
	Key string
}

// Creates a new FinancialAccount.
func (c v2MoneyManagementFinancialAccountService) Create(ctx context.Context, params *V2MoneyManagementFinancialAccountCreateParams) (*V2MoneyManagementFinancialAccount, error) {
	if params == nil {
		params = &V2MoneyManagementFinancialAccountCreateParams{}
	}
	params.Context = ctx
	financialaccount := &V2MoneyManagementFinancialAccount{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/financial_accounts", c.Key, params, financialaccount)
	return financialaccount, err
}

// Retrieves the details of an existing FinancialAccount.
func (c v2MoneyManagementFinancialAccountService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementFinancialAccountRetrieveParams) (*V2MoneyManagementFinancialAccount, error) {
	if params == nil {
		params = &V2MoneyManagementFinancialAccountRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/financial_accounts/%s", id)
	financialaccount := &V2MoneyManagementFinancialAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Closes a FinancialAccount with or without forwarding settings.
func (c v2MoneyManagementFinancialAccountService) Close(ctx context.Context, id string, params *V2MoneyManagementFinancialAccountCloseParams) (*V2MoneyManagementFinancialAccount, error) {
	if params == nil {
		params = &V2MoneyManagementFinancialAccountCloseParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/financial_accounts/%s/close", id)
	financialaccount := &V2MoneyManagementFinancialAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Lists FinancialAccounts in this compartment.
func (c v2MoneyManagementFinancialAccountService) List(ctx context.Context, listParams *V2MoneyManagementFinancialAccountListParams) Seq2[*V2MoneyManagementFinancialAccount, error] {
	if listParams == nil {
		listParams = &V2MoneyManagementFinancialAccountListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/money_management/financial_accounts", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementFinancialAccount], error) {
		page := &V2Page[*V2MoneyManagementFinancialAccount]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
