package stripe

import (
	"fmt"
	"net/url"
	"strconv"
)

type RecipientType string

const (
	Individual RecipientType = "individual"
	Corp       RecipientType = "corporation"
)

type RecipientParams struct {
	Name                      string
	Type                      RecipientType
	TaxId, Token, Email, Desc string
	Bank                      *BankAccountParams
	Card                      *CardParams
	Meta                      map[string]string
	DefaultCard               string
}

type RecipientListParams struct {
	Filters    Filters
	Start, End string
	Limit      uint64
	Verified   bool
}

type BankAccountParams struct {
	Country, Routing, Account string
}

type Recipient struct {
	Id          string            `json:"id"`
	Live        bool              `json:"livemode"`
	Created     int64             `json:"created"`
	Type        RecipientType     `json:"type"`
	Bank        *BankAccount      `json:"active_account"`
	Desc        string            `json:"description"`
	Email       string            `json:"email"`
	Meta        map[string]string `json:"metadata"`
	Name        string            `json:"name"`
	Cards       *CardList         `json:"cards"`
	DefaultCard string            `json:"default_card"`
}

type BankAccount struct {
	Id          string   `json:"id"`
	Name        string   `json:"bank_name"`
	Country     string   `json:"country"`
	Currency    Currency `json:"currency"`
	LastFour    string   `json:"last4"`
	Disabled    bool     `json:"disabled"`
	Fingerprint string   `json:"fingerprint"`
	Valid       bool     `json:"validated"`
}

type RecipientList struct {
	Count  uint16       `json:"total_count"`
	More   bool         `json:"has_more"`
	Url    string       `json:"url"`
	Values []*Recipient `json:"data"`
}

type RecipientClient struct {
	api   Api
	token string
}

func (c *RecipientClient) Create(params *RecipientParams) (*Recipient, error) {
	body := &url.Values{
		"name": {params.Name},
		"type": {string(params.Type)},
	}

	if params.Bank != nil {
		params.Bank.appendTo(body)
	}

	if len(params.Token) > 0 {
		body.Add("card", params.Token)
	} else if params.Card != nil {
		params.Card.appendTo(body, true)
	}

	if len(params.TaxId) > 0 {
		body.Add("tax_id", params.TaxId)
	}

	if len(params.Email) > 0 {
		body.Add("email", params.Email)
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	recipient := &Recipient{}
	err := c.api.Call("POST", "/recipients", c.token, body, recipient)

	return recipient, err
}

func (c *RecipientClient) Get(id string) (*Recipient, error) {
	recipient := &Recipient{}
	err := c.api.Call("GET", "/recipients/"+id, c.token, nil, recipient)

	return recipient, err
}

func (c *RecipientClient) Update(id string, params *RecipientParams) (*Recipient, error) {
	body := &url.Values{}

	if len(params.Name) > 0 {
		body.Add("name", params.Name)
	}

	if params.Bank != nil {
		params.Bank.appendTo(body)
	}

	if len(params.Token) > 0 {
		body.Add("card", params.Token)
	} else if params.Card != nil {
		params.Card.appendTo(body, true)
	}

	if len(params.TaxId) > 0 {
		body.Add("tax_id", params.TaxId)
	}

	if len(params.DefaultCard) > 0 {
		body.Add("default_card", params.DefaultCard)
	}

	if len(params.Email) > 0 {
		body.Add("email", params.Email)
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	recipient := &Recipient{}
	err := c.api.Call("POST", "/recipients/"+id, c.token, body, recipient)

	return recipient, err
}

func (c *RecipientClient) Delete(id string) error {
	return c.api.Call("DELETE", "/recipients/"+id, c.token, nil, nil)
}

func (c *RecipientClient) List(params *RecipientListParams) (*RecipientList, error) {
	body := &url.Values{}

	if params != nil {
		if params.Verified {
			body.Add("verified", strconv.FormatBool(true))
		}

		if len(params.Filters.f) > 0 {
			params.Filters.appendTo(body)
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

	list := &RecipientList{}
	err := c.api.Call("GET", "/recipients", c.token, body, list)

	return list, err
}
func (b *BankAccountParams) appendTo(values *url.Values) {
	values.Add("bank_account[country]", b.Country)
	values.Add("bank_account[routing_number]", b.Routing)
	values.Add("bank_account[account_number]", b.Account)
}
