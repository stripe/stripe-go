// Package account provides the /account APIs
package account

import (
	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /account APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of your account.
// For more details see https://stripe.com/docs/api/#retrieve_account.
func Get() (*stripe.Account, error) {
	return getC().Get()
}

func (c Client) Get() (*stripe.Account, error) {
	account := &stripe.Account{}
	err := c.B.Call("GET", "/account", c.Key, nil, nil, account)

	return account, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
