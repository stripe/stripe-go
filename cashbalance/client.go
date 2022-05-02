//
//
// File generated from our OpenAPI spec
//
//

// Package cashbalance provides the /customers/{customer}/cash_balance APIs
package cashbalance

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /customers/{customer}/cash_balance APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a cash balance.
func Get(params *stripe.CashBalanceParams) (*stripe.CashBalance, error) {
	return getC().Get(params)
}

// Get returns the details of a cash balance.
func (c Client) Get(params *stripe.CashBalanceParams) (*stripe.CashBalance, error) {
	if params == nil || params.Customer == nil {
		return nil, fmt.Errorf(
			"params cannot be nil, and params.Customer must be set",
		)
	}
	path := stripe.FormatURLPath(
		"/v1/customers/%s/cash_balance",
		stripe.StringValue(params.Customer),
	)
	cashbalance := &stripe.CashBalance{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cashbalance)
	return cashbalance, err
}

// Update updates a cash balance's properties.
func Update(params *stripe.CashBalanceParams) (*stripe.CashBalance, error) {
	return getC().Update(params)
}

// Update updates a cash balance's properties.
func (c Client) Update(params *stripe.CashBalanceParams) (*stripe.CashBalance, error) {
	if params == nil || params.Customer == nil {
		return nil, fmt.Errorf(
			"params cannot be nil, and params.Customer must be set",
		)
	}
	path := stripe.FormatURLPath(
		"/v1/customers/%s/cash_balance",
		stripe.StringValue(params.Customer),
	)
	cashbalance := &stripe.CashBalance{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cashbalance)
	return cashbalance, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
