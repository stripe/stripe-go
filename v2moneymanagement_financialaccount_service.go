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

// Updates an existing FinancialAccount using the provided parameters.
func (c v2MoneyManagementFinancialAccountService) Update(ctx context.Context, id string, params *V2MoneyManagementFinancialAccountUpdateParams) (*V2MoneyManagementFinancialAccount, error) {
	if params == nil {
		params = &V2MoneyManagementFinancialAccountUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/financial_accounts/%s", id)
	financialaccount := &V2MoneyManagementFinancialAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Returns a list of FinancialAccounts in your account that match the provided filters.
func (c v2MoneyManagementFinancialAccountService) List(ctx context.Context, listParams *V2MoneyManagementFinancialAccountListParams) *V2List[*V2MoneyManagementFinancialAccount] {
	if listParams == nil {
		listParams = &V2MoneyManagementFinancialAccountListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/financial_accounts", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementFinancialAccount], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementFinancialAccount]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
