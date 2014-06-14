package stripe

import (
	"net/url"
)

type TokenType string

const (
	CardToken TokenType = "card"
	BankToken TokenType = "bank_account"
)

type TokenParams struct {
	Card     *CardParams
	Customer string
}

type Token struct {
	Id      string       `json:"id"`
	Live    bool         `json:"livemode"`
	Created int64        `json:"created"`
	Type    TokenType    `json:"type"`
	Used    bool         `json:"used"`
	Bank    *BankAccount `json:"bank_account"`
	Card    *Card        `json:"card"`
}

type TokenClient struct {
	api   Api
	token string
}

func (c *TokenClient) Create(params *TokenParams) (*Token, error) {
	body := &url.Values{}

	params.Card.appendTo(body, true)

	if len(params.Customer) > 0 {
		body.Add("customer", params.Customer)
	}

	token := &Token{}
	err := c.api.Call("POST", "/tokens", c.token, body, token)

	return token, err
}

func (c *TokenClient) Get(id string) (*Token, error) {
	token := &Token{}
	err := c.api.Call("GET", "/tokens/"+id, c.token, nil, token)

	return token, err
}
