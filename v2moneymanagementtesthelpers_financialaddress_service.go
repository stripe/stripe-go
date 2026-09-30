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

// Simulate crediting a FinancialAddress in a Sandbox environment. This can be used to add virtual funds and increase your balance for testing.
func (c v2MoneyManagementTestHelpersFinancialAddressService) Credit(ctx context.Context, id string, params *V2MoneyManagementTestHelpersFinancialAddressCreditParams) (*V2MoneyManagementFinancialAddressCreditSimulation, error) {
	if params == nil {
		params = &V2MoneyManagementTestHelpersFinancialAddressCreditParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/test_helpers/financial_addresses/%s/credit", id)
	financialaddresscreditsimulation := &V2MoneyManagementFinancialAddressCreditSimulation{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaddresscreditsimulation)
	return financialaddresscreditsimulation, err
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

// Generates microdeposits for a FinancialAddress in a Sandbox environment.
func (c v2MoneyManagementTestHelpersFinancialAddressService) GenerateMicrodeposits(ctx context.Context, id string, params *V2MoneyManagementTestHelpersFinancialAddressGenerateMicrodepositsParams) (*V2MoneyManagementFinancialAddressGeneratedMicrodeposits, error) {
	if params == nil {
		params = &V2MoneyManagementTestHelpersFinancialAddressGenerateMicrodepositsParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/test_helpers/financial_addresses/%s/generate_microdeposits", id)
	financialaddressgeneratedmicrodeposits := &V2MoneyManagementFinancialAddressGeneratedMicrodeposits{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaddressgeneratedmicrodeposits)
	return financialaddressgeneratedmicrodeposits, err
}
