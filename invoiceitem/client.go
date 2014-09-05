package invoiceitem

import (
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /invoiceitems APIs.
type Client struct {
	B     Backend
	Token string
}

var c *Client

// Create POSTs new invoice items.
// For more details see https://stripe.com/docs/api#create_invoiceitem.
func Create(params *InvoiceItemParams) (*InvoiceItem, error) {
	refresh()
	return c.Create(params)
}

func (c *Client) Create(params *InvoiceItemParams) (*InvoiceItem, error) {
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

	params.AppendTo(body)

	invoiceItem := &InvoiceItem{}
	err := c.B.Call("POST", "/invoiceitems", c.Token, body, invoiceItem)

	return invoiceItem, err
}

// Get returns the details of an invoice item.
// For more details see https://stripe.com/docs/api#retrieve_invoiceitem.
func Get(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	refresh()
	return c.Get(id, params)
}

func (c *Client) Get(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	invoiceItem := &InvoiceItem{}
	err := c.B.Call("GET", "/invoiceitems/"+id, c.Token, body, invoiceItem)

	return invoiceItem, err
}

// Update updates an invoice item's properties.
// For more details see https://stripe.com/docs/api#update_invoiceitem.
func Update(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	refresh()
	return c.Update(id, params)
}

func (c *Client) Update(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Amount != 0 {
			body.Add("amount", strconv.FormatInt(params.Amount, 10))
		}

		if len(params.Desc) > 0 {
			body.Add("description", params.Desc)
		}

		params.AppendTo(body)
	}

	invoiceItem := &InvoiceItem{}
	err := c.B.Call("POST", "/invoiceitems/"+id, c.Token, body, invoiceItem)

	return invoiceItem, err
}

// Delete removes an invoice item.
// For more details see https://stripe.com/docs/api#delete_invoiceitem.
func Delete(id string) error {
	refresh()
	return c.Delete(id)
}

func (c *Client) Delete(id string) error {
	return c.B.Call("DELETE", "/invoiceitems/"+id, c.Token, nil, nil)
}

// List returns a list of invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
func List(params *InvoiceItemListParams) (*InvoiceItemList, error) {
	refresh()
	return c.List(params)
}

func (c *Client) List(params *InvoiceItemListParams) (*InvoiceItemList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Customer) > 0 {
			body.Add("customer", params.Customer)
		}

		params.AppendTo(body)
	}

	list := &InvoiceItemList{}
	err := c.B.Call("GET", "/invoiceitems", c.Token, body, list)

	return list, err
}

func refresh() {
	if c == nil {
		c = &Client{B: GetBackend()}
	}

	c.Token = Key
}
