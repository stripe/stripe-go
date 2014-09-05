package account

import (
	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /account APIs.
type Client struct {
	B     Backend
	Token string
}

var c *Client

// Get returns the details of your account.
// For more details see https://stripe.com/docs/api/#retrieve_account.
func Get() (*Account, error) {
	refresh()
	return c.Get()
}

func (c *Client) Get() (*Account, error) {
	account := &Account{}
	err := c.B.Call("GET", "/account", c.Token, nil, account)

	return account, err
}

func refresh() {
	if c == nil {
		c = &Client{B: GetBackend()}
	}

	c.Token = Key
}
