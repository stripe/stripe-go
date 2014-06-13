package stripe

import (
	"fmt"
	"net/url"
	"strconv"
)

type TransferStatus string
type TransferType string

const (
	Paid             TransferStatus = "paid"
	Pending          TransferStatus = "pending"
	Failed           TransferStatus = "failed"
	TransferCanceled TransferStatus = "canceled"

	CardTransfer TransferType = "card"
	BankTransfer TransferType = "bank_account"
)

type TransferParams struct {
	Amount                      int64
	Currency                    Currency
	Recipient                   string
	Desc, Statement, Bank, Card string
	Meta                        map[string]string
}

type Transfer struct {
	Id        string            `json:"id"`
	Live      bool              `json:"livemode"`
	Amount    int64             `json:"amount"`
	Currency  Currency          `json:"currency"`
	Created   int64             `json:"created"`
	Date      int64             `json:"date"`
	Status    TransferStatus    `json:"status"`
	Type      TransferType      `json:"type"`
	Balance   string            `json:"balance_transaction"`
	Desc      string            `json:"description"`
	Meta      map[string]string `json:"metadata"`
	Bank      *BankAccount      `json:"bank_account"`
	Card      *Card             `json:"card"`
	Recipient string            `json:"recipient"`
	Statement string            `json:"statement_description"`
}

type TransferClient struct {
	api   Api
	token string
}

func (c *TransferClient) Create(params *TransferParams) (*Transfer, error) {
	body := &url.Values{
		"amount":    {strconv.FormatInt(params.Amount, 10)},
		"currency":  {string(params.Currency)},
		"recipient": {params.Recipient},
	}

	if len(params.Bank) > 0 {
		body.Add("bank_account", params.Bank)
	} else if len(params.Card) > 0 {
		body.Add("card", params.Card)
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	if len(params.Statement) > 0 {
		body.Add("statement_description", params.Statement)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	transfer := &Transfer{}
	err := c.api.Call("POST", "/transfers", c.token, body, transfer)

	return transfer, err
}

func (c *TransferClient) Get(id string) (*Transfer, error) {
	transfer := &Transfer{}
	err := c.api.Call("GET", "/transfers/"+id, c.token, nil, transfer)

	return transfer, err
}

func (c *TransferClient) Update(id string, params *TransferParams) (*Transfer, error) {
	body := &url.Values{}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	transfer := &Transfer{}
	err := c.api.Call("POST", "/transfers/"+id, c.token, body, transfer)

	return transfer, err
}

func (c *TransferClient) Cancel(id string) (*Transfer, error) {
	transfer := &Transfer{}
	err := c.api.Call("POST", fmt.Sprintf("/transfers/%v/cancel", id), c.token, nil, transfer)

	return transfer, err
}
