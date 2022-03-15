//
//
// File generated from our OpenAPI spec
//
//

// Package balance provides the /balance APIs
package balance

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
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

// Iter is an iterator for balance transactions.
type Iter struct {
	*stripe.Iter
}

// BalanceTransaction returns the balance transaction which the iterator is currently pointing to.
func (i *Iter) BalanceTransaction() *stripe.BalanceTransaction {
	return i.Current().(*stripe.BalanceTransaction)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
