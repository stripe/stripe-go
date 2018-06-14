// Package token provides the /tokens APIs
package token

import (
	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /tokens APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new card or bank account.
// For more details see https://stripe.com/docs/api#create_card_Token and https://stripe.com/docs/api#create_bank_account_token.
func New(params *stripe.TokenParams) (*stripe.Token, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.TokenParams) (*stripe.Token, error) {
	tok := &stripe.Token{}
	err := c.B.Call("POST", "/tokens", c.Key, params, tok)
	return tok, err
}

// Get returns the details of a token.
// For more details see https://stripe.com/docs/api#retrieve_token.
func Get(id string, params *stripe.TokenParams) (*stripe.Token, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.TokenParams) (*stripe.Token, error) {
	path := stripe.FormatURLPath("/tokens/%s", id)
	token := &stripe.Token{}
	err := c.B.Call("GET", path, c.Key, params, token)

	return token, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
