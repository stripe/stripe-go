package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// RecipientType is the list of allowed values for the recipient's type.
// Allowed values are "individual", "corporation".
type RecipientType string

const (
	Individual RecipientType = "individual"
	Corp       RecipientType = "corporation"
)

// RecipientParams is the set of parameters that can be used when creating or updating recipients.
// For more details see https://stripe.com/docs/api#create_recipient and https://stripe.com/docs/api#update_recipient.
type RecipientParams struct {
	Params
	Name                      string
	Type                      RecipientType
	TaxId, Token, Email, Desc string
	Bank                      *BankAccountParams
	Card                      *CardParams
	DefaultCard               string
}

// RecipientListParams is the set of parameters that can be used when listing recipients.
// For more details see https://stripe.com/docs/api#list_recipients.
type RecipientListParams struct {
	ListParams
	Verified bool
}

// BankAccountParams is the set of parameters that can be used when creating or updating a bank account.
type BankAccountParams struct {
	Country, Routing, Account string
}

// Recipient is the resource representing a Stripe recipient.
// For more details see https://stripe.com/docs/api#recipients.
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
	DefaultCard *Card             `json:"default_card"`
}

// BankAccount represents a Stripe bank account.
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

// RecipientList is a list object for recipients.
type RecipientList struct {
	ListResponse
	Values []*Recipient `json:"data"`
}

// RecipientClient is the client used to invoke /recipients APIs.
type RecipientClient struct {
	api   Api
	token string
}

// Create POSTs a new recipient.
// For more details see https://stripe.com/docs/api#create_recipient.
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

// Get returns the details of a recipient.
// For more details see https://stripe.com/docs/api#retrieve_recipient.
func (c *RecipientClient) Get(id string) (*Recipient, error) {
	recipient := &Recipient{}
	err := c.api.Call("GET", "/recipients/"+id, c.token, nil, recipient)

	return recipient, err
}

// Update updates a recipient's properties.
// For more details see https://stripe.com/docs/api#update_recipient.
func (c *RecipientClient) Update(id string, params *RecipientParams) (*Recipient, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

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
	}

	recipient := &Recipient{}
	err := c.api.Call("POST", "/recipients/"+id, c.token, body, recipient)

	return recipient, err
}

// Delete removes a recipient.
// For more details see https://stripe.com/docs/api#delete_recipient.
func (c *RecipientClient) Delete(id string) error {
	return c.api.Call("DELETE", "/recipients/"+id, c.token, nil, nil)
}

// List returns a list of recipients.
// For more details see https://stripe.com/docs/api#list_recipients.
func (c *RecipientClient) List(params *RecipientListParams) (*RecipientList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

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

// appendTo adds the bank account's details to the query string values.
func (b *BankAccountParams) appendTo(values *url.Values) {
	values.Add("bank_account[country]", b.Country)
	values.Add("bank_account[routing_number]", b.Routing)
	values.Add("bank_account[account_number]", b.Account)
}

func (r *Recipient) UnmarshalJSON(data []byte) error {
	type recipient Recipient
	var rr recipient
	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Recipient(rr)
	} else {
		// the id is surrounded by escaped \, so ignore those
		r.Id = string(data[1 : len(data)-1])
	}

	return nil
}
