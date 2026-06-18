//
//
// File generated from our OpenAPI spec
//
//

// Package financialaddress provides the financialaddress related APIs
package financialaddress

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke financialaddress related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Simulate debiting a FinancialAddress in a Sandbox environment. This can be used to remove virtual funds and decrease your balance for testing.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Debit(id string, params *stripe.V2MoneyManagementTestHelpersFinancialAddressDebitParams) (*stripe.V2MoneyManagementFinancialAddressDebitSimulation, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/test_helpers/financial_addresses/%s/debit", id)
	financialaddressdebitsimulation := &stripe.V2MoneyManagementFinancialAddressDebitSimulation{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaddressdebitsimulation)
	return financialaddressdebitsimulation, err
}
