package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// InvoiceItemParams is the set of parameters that can be used when creating or updating an invoice item.
// For more details see https://stripe.com/docs/api#create_invoiceitem and https://stripe.com/docs/api#update_invoiceitem.
type InvoiceItemParams struct {
	Params
	Customer           string
	Amount             int64
	Currency           Currency
	Invoice, Desc, Sub string
}

// InvoiceItemListparams is the set of parameters that can be used when listing invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
type InvoiceItemListParams struct {
	ListParams
	Created  int64
	Customer string
}

// InvoiceItem is the resource represneting a Stripe invoice item.
// For more details see https://stripe.com/docs/api#invoiceitems.
type InvoiceItem struct {
	Id        string            `json:"id"`
	Live      bool              `json:"livemode"`
	Amount    int64             `json:"amount"`
	Currency  Currency          `json:"currency"`
	Customer  *Customer         `json:"customer"`
	Date      int64             `json:"date"`
	Proration bool              `json:"proration"`
	Desc      string            `json:"description"`
	Invoice   *Invoice          `json:"invoice"`
	Meta      map[string]string `json:"metadata"`
	Sub       string            `json:"subscription"`
}

// InvoiceItemList represents a list object for invoice items.
type InvoiceItemList struct {
	ListResponse
	Values []*InvoiceItem `json:"data"`
}

// InvoiceItemClient is the client used to invoke /invoiceitems APIs.
type InvoiceItemClient struct {
	api   Api
	token string
}

// Create POSTs new invoice items.
// For more details see https://stripe.com/docs/api#create_invoiceitem.
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

// Get returns the details of an invoice item.
// For more details see https://stripe.com/docs/api#retrieve_invoiceitem.
func (c *InvoiceItemClient) Get(id string) (*InvoiceItem, error) {
	invoiceItem := &InvoiceItem{}
	err := c.api.Call("GET", "/invoiceitems/"+id, c.token, nil, invoiceItem)

	return invoiceItem, err
}

// Update updates an invoice item's properties.
// For more details see https://stripe.com/docs/api#update_invoiceitem.
func (c *InvoiceItemClient) Update(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Amount != 0 {
			body.Add("amount", strconv.FormatInt(params.Amount, 10))
		}

		if len(params.Desc) > 0 {
			body.Add("description", params.Desc)
		}

		for k, v := range params.Meta {
			body.Add(fmt.Sprintf("metadata[%v]", k), v)
		}
	}

	invoiceItem := &InvoiceItem{}
	err := c.api.Call("POST", "/invoiceitems/"+id, c.token, body, invoiceItem)

	return invoiceItem, err
}

// Delete removes an invoice item.
// For more details see https://stripe.com/docs/api#delete_invoiceitem.
func (c *InvoiceItemClient) Delete(id string) error {
	return c.api.Call("DELETE", "/invoiceitems/"+id, c.token, nil, nil)
}

// List returns a list of invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
func (c *InvoiceItemClient) List(params *InvoiceItemListParams) (*InvoiceItemList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Filters.f) > 0 {
			params.Filters.appendTo(body)
		}

		if len(params.Customer) > 0 {
			body.Add("customer", params.Customer)
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

	list := &InvoiceItemList{}
	err := c.api.Call("GET", "/invoiceitems", c.token, body, list)

	return list, err
}

func (i *InvoiceItem) UnmarshalJSON(data []byte) error {
	type invoiceitem InvoiceItem
	var ii invoiceitem
	err := json.Unmarshal(data, &ii)
	if err == nil {
		*i = InvoiceItem(ii)
	} else {
		// the id is surrounded by escaped \, so ignore those
		i.Id = string(data[1 : len(data)-1])
	}

	return nil
}
