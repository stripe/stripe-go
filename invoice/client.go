//
//
// File generated from our OpenAPI spec
//
//

// Package invoice provides the /invoices APIs
package invoice

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /invoices APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new invoice.
func New(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().New(params)
}

// New creates a new invoice.
func (c Client) New(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, "/v1/invoices", c.Key, params, invoice)
	return invoice, err
}

// Get returns the details of an invoice.
func Get(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().Get(id, params)
}

// Get returns the details of an invoice.
func (c Client) Get(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, invoice)
	return invoice, err
}

// Update updates an invoice's properties.
func Update(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().Update(id, params)
}

// Update updates an invoice's properties.
func (c Client) Update(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Del removes an invoice.
func Del(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().Del(id, params)
}

// Del removes an invoice.
func (c Client) Del(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, invoice)
	return invoice, err
}

// FinalizeInvoice is the method for the `POST /v1/invoices/{invoice}/finalize` API.
func FinalizeInvoice(id string, params *stripe.InvoiceFinalizeParams) (*stripe.Invoice, error) {
	return getC().FinalizeInvoice(id, params)
}

// FinalizeInvoice is the method for the `POST /v1/invoices/{invoice}/finalize` API.
func (c Client) FinalizeInvoice(id string, params *stripe.InvoiceFinalizeParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/finalize", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// GetNext is the method for the `GET /v1/invoices/upcoming` API.
func GetNext(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	return getC().GetNext(params)
}

// GetNext is the method for the `GET /v1/invoices/upcoming` API.
func (c Client) GetNext(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	invoice := &stripe.Invoice{}
	err := c.B.Call(
		http.MethodGet,
		"/v1/invoices/upcoming",
		c.Key,
		params,
		invoice,
	)
	return invoice, err
}

// MarkUncollectible is the method for the `POST /v1/invoices/{invoice}/mark_uncollectible` API.
func MarkUncollectible(id string, params *stripe.InvoiceMarkUncollectibleParams) (*stripe.Invoice, error) {
	return getC().MarkUncollectible(id, params)
}

// MarkUncollectible is the method for the `POST /v1/invoices/{invoice}/mark_uncollectible` API.
func (c Client) MarkUncollectible(id string, params *stripe.InvoiceMarkUncollectibleParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/mark_uncollectible", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// Pay is the method for the `POST /v1/invoices/{invoice}/pay` API.
func Pay(id string, params *stripe.InvoicePayParams) (*stripe.Invoice, error) {
	return getC().Pay(id, params)
}

// Pay is the method for the `POST /v1/invoices/{invoice}/pay` API.
func (c Client) Pay(id string, params *stripe.InvoicePayParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/pay", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// SendInvoice is the method for the `POST /v1/invoices/{invoice}/send` API.
func SendInvoice(id string, params *stripe.InvoiceSendParams) (*stripe.Invoice, error) {
	return getC().SendInvoice(id, params)
}

// SendInvoice is the method for the `POST /v1/invoices/{invoice}/send` API.
func (c Client) SendInvoice(id string, params *stripe.InvoiceSendParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/send", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// VoidInvoice is the method for the `POST /v1/invoices/{invoice}/void` API.
func VoidInvoice(id string, params *stripe.InvoiceVoidParams) (*stripe.Invoice, error) {
	return getC().VoidInvoice(id, params)
}

// VoidInvoice is the method for the `POST /v1/invoices/{invoice}/void` API.
func (c Client) VoidInvoice(id string, params *stripe.InvoiceVoidParams) (*stripe.Invoice, error) {
	path := stripe.FormatURLPath("/v1/invoices/%s/void", id)
	invoice := &stripe.Invoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoice)
	return invoice, err
}

// List returns a list of invoices.
func List(params *stripe.InvoiceListParams) *Iter {
	return getC().List(params)
}

// List returns a list of invoices.
func (c Client) List(listParams *stripe.InvoiceListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.InvoiceList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/invoices", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for invoices.
type Iter struct {
	*stripe.Iter
}

// Invoice returns the invoice which the iterator is currently pointing to.
func (i *Iter) Invoice() *stripe.Invoice {
	return i.Current().(*stripe.Invoice)
}

// InvoiceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) InvoiceList() *stripe.InvoiceList {
	return i.List().(*stripe.InvoiceList)
}

// ListLines is the method for the `GET /v1/invoices/{invoice}/lines` API.
func ListLines(params *stripe.InvoiceLineListParams) *LineIter {
	return getC().ListLines(params)
}

// ListLines is the method for the `GET /v1/invoices/{invoice}/lines` API.
func (c Client) ListLines(listParams *stripe.InvoiceLineListParams) *LineIter {
	path := stripe.FormatURLPath(
		"/v1/invoices/%s/lines",
		stripe.StringValue(listParams.ID),
	)
	return &LineIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.InvoiceLineList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// LineIter is an iterator for invoice line items.
type LineIter struct {
	*stripe.Iter
}

// InvoiceLine returns the invoice line item which the iterator is currently pointing to.
func (i *LineIter) InvoiceLine() *stripe.InvoiceLine {
	return i.Current().(*stripe.InvoiceLine)
}

// InvoiceLineList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineIter) InvoiceLineList() *stripe.InvoiceLineList {
	return i.List().(*stripe.InvoiceLineList)
}

// Search returns a search result containing invoices.
func Search(params *stripe.InvoiceSearchParams) *SearchIter {
	return getC().Search(params)
}

// Search returns a search result containing invoices.
func (c Client) Search(params *stripe.InvoiceSearchParams) *SearchIter {
	return &SearchIter{
		SearchIter: stripe.GetSearchIter(params, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.SearchContainer, error) {
			list := &stripe.InvoiceSearchResult{}
			err := c.B.CallRaw(http.MethodGet, "/v1/invoices/search", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// SearchIter is an iterator for invoices.
type SearchIter struct {
	*stripe.SearchIter
}

// Invoice returns the invoice which the iterator is currently pointing to.
func (i *SearchIter) Invoice() *stripe.Invoice {
	return i.Current().(*stripe.Invoice)
}

// InvoiceSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) InvoiceSearchResult() *stripe.InvoiceSearchResult {
	return i.SearchResult().(*stripe.InvoiceSearchResult)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
