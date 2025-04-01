//
//
// File generated from our OpenAPI spec
//
//

// Package financialaddress provides the financialaddress related APIs
package financialaddress

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke financialaddress related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Simulate crediting a FinancialAddress in a Sandbox environment. This can be used to add virtual funds and increase your balance for testing.
func (c Client) Credit(id string, params *stripe.V2TestHelpersFinancialAddressCreditParams) (*stripe.V2FinancialAddressCreditSimulation, error) {
	path := stripe.FormatURLPath(
		"/v2/test_helpers/financial_addresses/%s/credit", id)
	financialaddresscreditsimulation := &stripe.V2FinancialAddressCreditSimulation{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaddresscreditsimulation)
	return financialaddresscreditsimulation, err
}

// Generates microdeposits for a FinancialAddress in a Sandbox environment.
func (c Client) GenerateMicrodeposits(id string, params *stripe.V2TestHelpersFinancialAddressGenerateMicrodepositsParams) (*stripe.V2FinancialAddressGeneratedMicrodeposits, error) {
	path := stripe.FormatURLPath(
		"/v2/test_helpers/financial_addresses/%s/generate_microdeposits", id)
	financialaddressgeneratedmicrodeposits := &stripe.V2FinancialAddressGeneratedMicrodeposits{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaddressgeneratedmicrodeposits)
	return financialaddressgeneratedmicrodeposits, err
}
