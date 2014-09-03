package stripe

import (
	"errors"
	"net/url"
)

// TokenType is the list of allowed values for a token's type.
// Allowed values are "card", "bank_account".
type TokenType string

const (
	CardToken TokenType = "card"
	BankToken TokenType = "bank_account"
)

// TokenParams is the set of parameters that can be used when creating a token.
// For more details see https://stripe.com/docs/api#create_card_token and https://stripe.com/docs/api#create_bank_account_token.
type TokenParams struct {
	Params
	Card     *CardParams
	Bank     *BankAccountParams
	Customer string
}

// Token is the resource representing a Stripe token.
// For more details see https://stripe.com/docs/api#tokens.
type Token struct {
	Id      string       `json:"id"`
	Live    bool         `json:"livemode"`
	Created int64        `json:"created"`
	Type    TokenType    `json:"type"`
	Used    bool         `json:"used"`
	Bank    *BankAccount `json:"bank_account"`
	Card    *Card        `json:"card"`
}

// TokenClient is the client used to invoke /tokens APIs.
type TokenClient struct {
	api   Api
	token string
}

// Create POSTs a new card or bank account.
// For more details see https://stripe.com/docs/api#create_card_token and https://stripe.com/docs/api#create_bank_account_token.
func (c *TokenClient) Create(params *TokenParams) (*Token, error) {
	body := &url.Values{}
	token := c.token

	if len(params.Customer) > 0 {
		if len(params.AccessToken) == 0 {
			err := errors.New("Invalid token params: an access token is required for customer")
			return nil, err
		}

		body.Add("customer", params.Customer)
		token = params.AccessToken
	}

	if params.Card != nil {
		params.Card.appendTo(body, true)
	} else if params.Bank != nil {
		params.Bank.appendTo(body)
	} else if len(params.Customer) == 0 {
		err := errors.New("Invalid token params: either Card or Bank need to be set")
		return nil, err
	}

	tok := &Token{}
	err := c.api.Call("POST", "/tokens", token, body, tok)

	return tok, err
}

// Get returns the details of a token.
// For more details see https://stripe.com/docs/api#retrieve_token.
func (c *TokenClient) Get(id string) (*Token, error) {
	token := &Token{}
	err := c.api.Call("GET", "/tokens/"+id, c.token, nil, token)

	return token, err
}
