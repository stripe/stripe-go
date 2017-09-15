// Package token provides the /tokens APIs
package token

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	Card stripe.TokenType = "card"
	Bank stripe.TokenType = "bank_account"
	PII  stripe.TokenType = "pii"
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
	body := &form.Values{}
	form.AppendTo(body, params)

	tok := &stripe.Token{}
	err := c.B.Call("POST", "/tokens", c.Key, body, &params.Params, tok)

	return tok, err
}

// Get returns the details of a token.
// For more details see https://stripe.com/docs/api#retrieve_token.
func Get(id string, params *stripe.TokenParams) (*stripe.Token, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.TokenParams) (*stripe.Token, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	token := &stripe.Token{}
	err := c.B.Call("GET", "/tokens/"+id, c.Key, body, commonParams, token)

	return token, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
