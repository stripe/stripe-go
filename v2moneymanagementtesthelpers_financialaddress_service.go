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

// v2MoneyManagementTestHelpersFinancialAddressService is used to invoke financialaddress related APIs.
type v2MoneyManagementTestHelpersFinancialAddressService struct {
	B   Backend
	Key string
}

// Simulate debiting a FinancialAddress in a Sandbox environment. This can be used to remove virtual funds and decrease your balance for testing.
func (c v2MoneyManagementTestHelpersFinancialAddressService) Debit(ctx context.Context, id string, params *V2MoneyManagementTestHelpersFinancialAddressDebitParams) (*V2MoneyManagementFinancialAddressDebitSimulation, error) {
	if params == nil {
		params = &V2MoneyManagementTestHelpersFinancialAddressDebitParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/test_helpers/financial_addresses/%s/debit", id)
	financialaddressdebitsimulation := &V2MoneyManagementFinancialAddressDebitSimulation{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaddressdebitsimulation)
	return financialaddressdebitsimulation, err
}
