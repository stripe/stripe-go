//
//
// File generated from our OpenAPI spec
//
//

// Package customer provides the /customers APIs
package customer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v75"
)

// Client is used to invoke /customers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// FundCashBalance is the method for the `POST /v1/test_helpers/customers/{customer}/fund_cash_balance` API.
func FundCashBalance(id string, params *stripe.TestHelpersCustomerFundCashBalanceParams) (*stripe.CustomerCashBalanceTransaction, error) {
	return getC().FundCashBalance(id, params)
}

// FundCashBalance is the method for the `POST /v1/test_helpers/customers/{customer}/fund_cash_balance` API.
func (c Client) FundCashBalance(id string, params *stripe.TestHelpersCustomerFundCashBalanceParams) (*stripe.CustomerCashBalanceTransaction, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/customers/%s/fund_cash_balance",
		id,
	)
	customercashbalancetransaction := &stripe.CustomerCashBalanceTransaction{}
	err := c.B.Call(
		http.MethodPost,
		path,
		c.Key,
		params,
		customercashbalancetransaction,
	)
	return customercashbalancetransaction, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
