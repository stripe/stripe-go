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

// v2MoneyManagementFinancialAccountsStatementService is used to invoke statement related APIs.
type v2MoneyManagementFinancialAccountsStatementService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Financial Account Statement.
func (c v2MoneyManagementFinancialAccountsStatementService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementFinancialAccountsStatementRetrieveParams) (*V2MoneyManagementFinancialAccountStatement, error) {
	if params == nil {
		params = &V2MoneyManagementFinancialAccountsStatementRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/financial_accounts/%s/statements/%s", StringValue(
			params.FinancialAccountID), id)
	financialaccountstatement := &V2MoneyManagementFinancialAccountStatement{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, financialaccountstatement)
	return financialaccountstatement, err
}

// Returns a list of statements for a Financial Account.
func (c v2MoneyManagementFinancialAccountsStatementService) List(ctx context.Context, listParams *V2MoneyManagementFinancialAccountsStatementListParams) *V2List[*V2MoneyManagementFinancialAccountStatement] {
	if listParams == nil {
		listParams = &V2MoneyManagementFinancialAccountsStatementListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/financial_accounts/%s/statements", StringValue(
			listParams.FinancialAccountID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementFinancialAccountStatement], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementFinancialAccountStatement]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
