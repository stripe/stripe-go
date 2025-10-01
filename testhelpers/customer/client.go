//
//
// File generated from our OpenAPI spec
//
//

// Package customer provides the /v1/customers APIs
package customer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke /v1/customers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an incoming testmode bank transfer
func FundCashBalance(id string, params *stripe.TestHelpersCustomerFundCashBalanceParams) (*stripe.CustomerCashBalanceTransaction, error) {
	return getC().FundCashBalance(id, params)
}

// Create an incoming testmode bank transfer
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) FundCashBalance(id string, params *stripe.TestHelpersCustomerFundCashBalanceParams) (*stripe.CustomerCashBalanceTransaction, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/customers/%s/fund_cash_balance", id)
	customercashbalancetransaction := &stripe.CustomerCashBalanceTransaction{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, customercashbalancetransaction)
	return customercashbalancetransaction, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
