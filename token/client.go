// package token provides the /tokens APIs
package token

import (
	"errors"
	"net/url"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /tokens APIs.
type Client struct {
	B   Backend
	Key string
}

// New POSTs a new card or bank account.
// For more details see https://stripe.com/docs/api#create_card_Token and https://stripe.com/docs/api#create_bank_account_token.
func New(params *TokenParams) (*Token, error) {
	return getC().New(params)
}

func (c Client) New(params *TokenParams) (*Token, error) {
	body := &url.Values{}
	token := c.Key

	if len(params.Customer) > 0 {
		body.Add("customer", params.Customer)
	}

	if params.Card != nil {
		params.Card.AppendDetails(body, true)
	} else if params.Bank != nil {
		params.Bank.AppendDetails(body)
	} else if len(params.Customer) == 0 {
		err := errors.New("Invalid Token params: either Card or Bank need to be set")
		return nil, err
	}

	params.AppendTo(body)

	tok := &Token{}
	err := c.B.Call("POST", "/tokens", token, body, tok)

	return tok, err
}

// Get returns the details of a token.
// For more details see https://stripe.com/docs/api#retrieve_token.
func Get(id string, params *TokenParams) (*Token, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *TokenParams) (*Token, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	token := &Token{}
	err := c.B.Call("GET", "/tokens/"+id, c.Key, body, token)

	return token, err
}

func getC() Client {
	return Client{GetBackend(), Key}
}
