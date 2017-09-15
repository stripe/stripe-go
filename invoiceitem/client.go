// Package invoiceitem provides the /invoiceitems APIs
package invoiceitem

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
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
	body := &form.Values{}
	form.AppendTo(body, params)

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
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
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
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	invoiceItem := &stripe.InvoiceItem{}
	err := c.B.Call("POST", "/invoiceitems/"+id, c.Key, body, commonParams, invoiceItem)

	return invoiceItem, err
}

// Del removes an invoice item.
// For more details see https://stripe.com/docs/api#delete_invoiceitem.
func Del(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.InvoiceItemParams) (*stripe.InvoiceItem, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	ii := &stripe.InvoiceItem{}
	err := c.B.Call("DELETE", "/invoiceitems/"+id, c.Key, body, commonParams, ii)

	return ii, err
}

// List returns a list of invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
func List(params *stripe.InvoiceItemListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.InvoiceItemListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.InvoiceItemList{}
		err := c.B.Call("GET", "/invoiceitems", c.Key, b, p, list)

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
