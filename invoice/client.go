// Package invoice provides the /invoices APIs
package invoice

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is the client used to invoke /invoices APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new invoices.
// For more details see https://stripe.com/docs/api#create_invoice.
func New(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	invoice := &stripe.Invoice{}
	err := c.B.Call2("POST", "/invoices", c.Key, params, invoice)
	return invoice, err
}

// Get returns the details of an invoice.
// For more details see https://stripe.com/docs/api#retrieve_invoice.
func Get(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/invoices/%s", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call2("GET", path, c.Key, params, invoice)
	return invoice, err
}

// Pay pays an invoice.
// For more details see https://stripe.com/docs/api#pay_invoice.
func Pay(id string, params *stripe.InvoicePayParams) (*stripe.Invoice, error) {
	return getC().Pay(id, params)
}

func (c Client) Pay(id string, params *stripe.InvoicePayParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/invoices/%s/pay", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call2("POST", path, c.Key, params, invoice)
	return invoice, err
}

// Update updates an invoice's properties.
// For more details see https://stripe.com/docs/api#update_invoice.
func Update(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/invoices/%s", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call2("POST", path, c.Key, params, invoice)
	return invoice, err
}

// GetNext returns the upcoming invoice's properties.
// For more details see https://stripe.com/docs/api#retrieve_customer_invoice.
func GetNext(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().GetNext(params)
}

func (c Client) GetNext(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	invoice := &stripe.Invoice{}
	err := c.B.Call2("GET", "/invoices/upcoming", c.Key, params, invoice)
	return invoice, err
}

// List returns a list of invoices.
// For more details see https://stripe.com/docs/api#list_customer_invoices.
func List(params *stripe.InvoiceListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.InvoiceListParams) *Iter {
	return &Iter{stripe.GetIter2(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.InvoiceList{}
		err := c.B.CallRaw("GET", "/invoices", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// ListLines returns a list of line items.
// For more details see https://stripe.com/docs/api#invoice_lines.
func ListLines(params *stripe.InvoiceLineListParams) *LineIter {
	return getC().ListLines(params)
}

func (c Client) ListLines(listParams *stripe.InvoiceLineListParams) *LineIter {
	path := stripe.FormatURLPath("/invoices/%s/lines", stripe.StringValue(listParams.ID))
	return &LineIter{stripe.GetIter2(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.InvoiceLineList{}
		err := c.B.CallRaw("GET", path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Invoices.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Invoice returns the most recent Invoice
// visited by a call to Next.
func (i *Iter) Invoice() *stripe.Invoice {
	return i.Current().(*stripe.Invoice)
}

// LineIter is an iterator for lists of InvoiceLines.
// The embedded Iter carries methods with it;
// see its documentation for details.
type LineIter struct {
	*stripe.Iter
}

// InvoiceLine returns the most recent InvoiceLine
// visited by a call to Next.
func (i *LineIter) InvoiceLine() *stripe.InvoiceLine {
	return i.Current().(*stripe.InvoiceLine)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
