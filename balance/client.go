//
//
// File generated from our OpenAPI spec
//
//

// Package balance provides the /balance APIs
package balance

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
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
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, "/v1/balance", c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, balance)
	return balance, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
