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

// v2TestHelpersFinancialAddressService is used to invoke financialaddress related APIs.
type v2TestHelpersFinancialAddressService struct {
	B   Backend
	Key string
}

// Simulate crediting a FinancialAddress in a Sandbox environment. This can be used to add virtual funds and increase your balance for testing.
func (c v2TestHelpersFinancialAddressService) Credit(ctx context.Context, id string, params *V2TestHelpersFinancialAddressCreditParams) (*V2FinancialAddressCreditSimulation, error) {
	path := FormatURLPath("/v2/test_helpers/financial_addresses/%s/credit", id)
	financialaddresscreditsimulation := &V2FinancialAddressCreditSimulation{}
	if params == nil {
		params = &V2TestHelpersFinancialAddressCreditParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaddresscreditsimulation)
	return financialaddresscreditsimulation, err
}

// Generates microdeposits for a FinancialAddress in a Sandbox environment.
func (c v2TestHelpersFinancialAddressService) GenerateMicrodeposits(ctx context.Context, id string, params *V2TestHelpersFinancialAddressGenerateMicrodepositsParams) (*V2FinancialAddressGeneratedMicrodeposits, error) {
	path := FormatURLPath(
		"/v2/test_helpers/financial_addresses/%s/generate_microdeposits", id)
	financialaddressgeneratedmicrodeposits := &V2FinancialAddressGeneratedMicrodeposits{}
	if params == nil {
		params = &V2TestHelpersFinancialAddressGenerateMicrodepositsParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaddressgeneratedmicrodeposits)
	return financialaddressgeneratedmicrodeposits, err
}
