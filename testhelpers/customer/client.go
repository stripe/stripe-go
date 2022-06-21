//
//
// File generated from our OpenAPI spec
//
//

// Package customer provides the /customers APIs
package customer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /customers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// FundCashBalance is the method for the `POST /v1/test_helpers/customers/{customer}/fund_cash_balance` API.
func FundCashBalance(id string, params *stripe.TestHelpersCustomerFundCashBalanceParams) (*stripe.Customer, error) {
	return getC().FundCashBalance(id, params)
}

// FundCashBalance is the method for the `POST /v1/test_helpers/customers/{customer}/fund_cash_balance` API.
func (c Client) FundCashBalance(id string, params *stripe.TestHelpersCustomerFundCashBalanceParams) (*stripe.Customer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/customers/%s/fund_cash_balance",
		id,
	)
	customer := &stripe.Customer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, customer)
	return customer, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
