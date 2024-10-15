//
//
// File generated from our OpenAPI spec
//
//

// Package customercashbalancetransaction provides the /customers/{customer}/cash_balance_transactions APIs
package customercashbalancetransaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v80"
)

// Client is used to invoke /customers/{customer}/cash_balance_transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Simulate various customer cash balance side-effects by creating synthetic cash balance transactions in testmode.
func New(params *stripe.TestHelpersCustomerCashBalanceTransactionParams) (*stripe.CustomerCashBalanceTransaction, error) {
	return getC().New(params)
}

// Simulate various customer cash balance side-effects by creating synthetic cash balance transactions in testmode.
func (c Client) New(params *stripe.TestHelpersCustomerCashBalanceTransactionParams) (*stripe.CustomerCashBalanceTransaction, error) {
	customercashbalancetransaction := &stripe.CustomerCashBalanceTransaction{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/test_helpers/customer_cash_balance_transactions",
		c.Key,
		params,
		customercashbalancetransaction,
	)
	return customercashbalancetransaction, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
