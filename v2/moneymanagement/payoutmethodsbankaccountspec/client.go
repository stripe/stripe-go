//
//
// File generated from our OpenAPI spec
//
//

// Package payoutmethodsbankaccountspec provides the payoutmethodsbankaccountspec related APIs
package payoutmethodsbankaccountspec

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke payoutmethodsbankaccountspec related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Fetch the specifications for a set of countries to know which
// credential fields are required, the validations for each fields, and how to translate these
// country-specific fields to the generic fields in the PayoutMethodBankAccount type.
func (c Client) Get(params *stripe.V2MoneyManagementPayoutMethodsBankAccountSpecParams) (*stripe.V2MoneyManagementPayoutMethodsBankAccountSpec, error) {
	payoutmethodsbankaccountspec := &stripe.V2MoneyManagementPayoutMethodsBankAccountSpec{}
	err := c.B.Call(
		http.MethodGet, "/v2/money_management/payout_methods_bank_account_spec", c.Key, params, payoutmethodsbankaccountspec)
	return payoutmethodsbankaccountspec, err
}
