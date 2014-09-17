// package account provides the /account APIs
package account

import (
	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /account APIs.
type Client struct {
	B   Backend
	Key string
}

// Get returns the details of your account.
// For more details see https://stripe.com/docs/api/#retrieve_account.
func Get() (*Account, error) {
	return getC().Get()
}

func (c Client) Get() (*Account, error) {
	account := &Account{}
	err := c.B.Call("GET", "/account", c.Key, nil, account)

	return account, err
}

func getC() Client {
	return Client{GetBackend(), Key}
}
