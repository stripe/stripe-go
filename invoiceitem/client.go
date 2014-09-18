// package invoiceitem provides the /invoiceitems APIs
package invoiceitem

import (
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /invoiceitems APIs.
type Client struct {
	B   Backend
	Key string
}

// New POSTs new invoice items.
// For more details see https://stripe.com/docs/api#create_invoiceitem.
func New(params *InvoiceItemParams) (*InvoiceItem, error) {
	return getC().New(params)
}

func (c Client) New(params *InvoiceItemParams) (*InvoiceItem, error) {
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
	err := c.B.Call("POST", "/invoiceitems", c.Key, body, invoiceItem)

	return invoiceItem, err
}

// Get returns the details of an invoice item.
// For more details see https://stripe.com/docs/api#retrieve_invoiceitem.
func Get(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	invoiceItem := &InvoiceItem{}
	err := c.B.Call("GET", "/invoiceitems/"+id, c.Key, body, invoiceItem)

	return invoiceItem, err
}

// Update updates an invoice item's properties.
// For more details see https://stripe.com/docs/api#update_invoiceitem.
func Update(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
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
	err := c.B.Call("POST", "/invoiceitems/"+id, c.Key, body, invoiceItem)

	return invoiceItem, err
}

// Del removes an invoice item.
// For more details see https://stripe.com/docs/api#delete_invoiceitem.
func Del(id string) error {
	return getC().Del(id)
}

func (c Client) Del(id string) error {
	return c.B.Call("DELETE", "/invoiceitems/"+id, c.Key, nil, nil)
}

// List returns a list of invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
func List(params *InvoiceItemListParams) *InvoiceItemIter {
	return getC().List(params)
}

func (c Client) List(params *InvoiceItemListParams) *InvoiceItemIter {
	type invoiceItemList struct {
		ListMeta
		Values []*InvoiceItem `json:"data"`
	}

	var body *url.Values
	var lp *ListParams

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

	return &InvoiceItemIter{GetIter(lp, body, func(b url.Values) ([]interface{}, ListMeta, error) {
		list := &invoiceItemList{}
		err := c.B.Call("GET", "/invoiceitems", c.Key, &b, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func getC() Client {
	return Client{GetBackend(), Key}
}
