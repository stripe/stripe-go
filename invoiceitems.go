package stripe

import (
	"fmt"
	"net/url"
	"strconv"
)

type InvoiceItemParams struct {
	Customer           string
	Amount             int64
	Currency           Currency
	Invoice, Desc, Sub string
	Meta               map[string]string
}

type InvoiceItem struct {
	Id        string            `json:"id"`
	Live      bool              `json:"livemode"`
	Amount    int64             `json:"amount"`
	Currency  Currency          `json:"currency"`
	Customer  string            `json:"customer"`
	Date      int64             `json:"date"`
	Proration bool              `json:"proration"`
	Desc      string            `json:"description"`
	Invoice   string            `json:invoice"`
	Sub       string            `json:"subscription"`
	Meta      map[string]string `json:"metadata"`
}

type InvoiceItemClient struct {
	api   Api
	token string
}

func (c *InvoiceItemClient) Create(params *InvoiceItemParams) (*InvoiceItem, error) {
	body := &url.Values{
		"customer": {params.Customer},
		"amount":   {strconv.FormatInt(params.Amount, 10)},
		"currency": {string(params.Currency)},
	}

	if len(params.Invoice) > 0 {
		body.Add("invoice", params.Invoice)
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	if len(params.Sub) > 0 {
		body.Add("subscription", params.Sub)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	invoiceItem := &InvoiceItem{}
	err := c.api.Call("POST", "/invoiceitems", c.token, body, invoiceItem)

	return invoiceItem, err
}

func (c *InvoiceItemClient) Get(id string) (*InvoiceItem, error) {
	invoiceItem := &InvoiceItem{}
	err := c.api.Call("GET", "/invoiceitems/"+id, c.token, nil, invoiceItem)

	return invoiceItem, err
}

func (c *InvoiceItemClient) Update(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	body := &url.Values{}

	if params.Amount != 0 {
		body.Add("amount", strconv.FormatInt(params.Amount, 10))
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	invoiceItem := &InvoiceItem{}
	err := c.api.Call("POST", "/invoiceitems/"+id, c.token, body, invoiceItem)

	return invoiceItem, err
}

func (c *InvoiceItemClient) Delete(id string) error {
	return c.api.Call("DELETE", "/invoiceitems/"+id, c.token, nil, nil)
}
