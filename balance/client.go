//
//
// File generated from our OpenAPI spec
//
//

// Package balance provides the /balance APIs
package balance

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
)

// Client is used to invoke /balance APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a balance.
func Get(params *stripe.BalanceParams) (*stripe.Balance, error) {
	return getC().Get(params)
}

// Get returns the details of a balance.
func (c Client) Get(params *stripe.BalanceParams) (*stripe.Balance, error) {
	balance := &stripe.Balance{}
	err := c.B.Call(http.MethodGet, "/v1/balance", c.Key, params, balance)
	return balance, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
