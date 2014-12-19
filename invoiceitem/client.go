// Package invoiceitem provides the /invoiceitems APIs
package invoiceitem

import (
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /invoiceitems APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new invoice items.
// For more details see https://stripe.com/docs/api#create_invoiceitem.
func New(params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
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

	invoiceItem := &stripe.InvoiceItem{}
	err := c.B.Call("POST", "/invoiceitems", c.Key, body, &params.Params, invoiceItem)

	return invoiceItem, err
}

// Get returns the details of an invoice item.
// For more details see https://stripe.com/docs/api#retrieve_invoiceitem.
func Get(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}
		params.AppendTo(body)
	}

	invoiceItem := &stripe.InvoiceItem{}
	err := c.B.Call("GET", "/invoiceitems/"+id, c.Key, body, commonParams, invoiceItem)

	return invoiceItem, err
}

// Update updates an invoice item's properties.
// For more details see https://stripe.com/docs/api#update_invoiceitem.
func Update(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}

		if params.Amount != 0 {
			body.Add("amount", strconv.FormatInt(params.Amount, 10))
		}

		if len(params.Desc) > 0 {
			body.Add("description", params.Desc)
		}

		params.AppendTo(body)
	}

	invoiceItem := &stripe.InvoiceItem{}
	err := c.B.Call("POST", "/invoiceitems/"+id, c.Key, body, commonParams, invoiceItem)

	return invoiceItem, err
}

// Del removes an invoice item.
// For more details see https://stripe.com/docs/api#delete_invoiceitem.
func Del(id string) error {
	return getC().Del(id)
}

func (c Client) Del(id string) error {
	return c.B.Call("DELETE", "/invoiceitems/"+id, c.Key, nil, nil, nil)
}

// List returns a list of invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
func List(params *stripe.InvoiceItemListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.InvoiceItemListParams) *Iter {
	type invoiceItemList struct {
		stripe.ListMeta
		Values []*stripe.InvoiceItem `json:"data"`
	}

	var body *url.Values
	var lp *stripe.ListParams

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Customer) > 0 {
			body.Add("customer", params.Customer)
		}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &invoiceItemList{}
		err := c.B.Call("GET", "/invoiceitems", c.Key, &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of InvoiceItems.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// InvoiceItem returns the most recent InvoiceItem
// visited by a call to Next.
func (i *Iter) InvoiceItem() *stripe.InvoiceItem {
	return i.Current().(*stripe.InvoiceItem)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
