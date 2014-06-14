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

type TransferListParams struct {
	Created, Date         int64
	Filters               Filters
	Recipient, Start, End string
	Status                TransferStatus
	Limit                 uint64
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
	Tx        string            `json:"balance_transaction"`
	Desc      string            `json:"description"`
	Meta      map[string]string `json:"metadata"`
	Bank      *BankAccount      `json:"bank_account"`
	Card      *Card             `json:"card"`
	Recipient string            `json:"recipient"`
	Statement string            `json:"statement_description"`
}

type TransferList struct {
	Count  uint16      `json:"total_count"`
	More   bool        `json:"has_more"`
	Url    string      `json:"url"`
	Values []*Transfer `json:"data"`
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
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Desc) > 0 {
			body.Add("description", params.Desc)
		}

		for k, v := range params.Meta {
			body.Add(fmt.Sprintf("metadata[%v]", k), v)
		}
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

func (c *TransferClient) List(params *TransferListParams) (*TransferList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if params.Date > 0 {
			body.Add("date", strconv.FormatInt(params.Date, 10))
		}

		if len(params.Filters.f) > 0 {
			params.Filters.appendTo(body)
		}

		if len(params.Recipient) > 0 {
			body.Add("recipient", params.Recipient)
		}

		if len(params.Status) > 0 {
			body.Add("status", string(params.Status))
		}

		if len(params.Start) > 0 {
			body.Add("starting_after", params.Start)
		}

		if len(params.End) > 0 {
			body.Add("ending_before", params.End)
		}

		if params.Limit > 0 {
			if params.Limit > 100 {
				params.Limit = 100
			}

			body.Add("limit", strconv.FormatUint(params.Limit, 10))
		}
	}

	list := &TransferList{}
	err := c.api.Call("GET", "/transfers", c.token, body, list)

	return list, err
}
