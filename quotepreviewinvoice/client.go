//
//
// File generated from our OpenAPI spec
//
//

// Package quotepreviewinvoice provides the /quotes/{quote}/preview_invoices APIs
package quotepreviewinvoice

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/form"
)

// Client is used to invoke /quotes/{quote}/preview_invoices APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of quote preview invoices.
func List(params *stripe.QuotePreviewInvoiceListParams) *Iter {
	return getC().List(params)
}

// List returns a list of quote preview invoices.
func (c Client) List(listParams *stripe.QuotePreviewInvoiceListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/preview_invoices",
		stripe.StringValue(listParams.Quote),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.QuotePreviewInvoiceList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for quote preview invoices.
type Iter struct {
	*stripe.Iter
}

// QuotePreviewInvoice returns the quote preview invoice which the iterator is currently pointing to.
func (i *Iter) QuotePreviewInvoice() *stripe.QuotePreviewInvoice {
	return i.Current().(*stripe.QuotePreviewInvoice)
}

// QuotePreviewInvoiceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) QuotePreviewInvoiceList() *stripe.QuotePreviewInvoiceList {
	return i.List().(*stripe.QuotePreviewInvoiceList)
}

// ListLines is the method for the `GET /v1/quotes/{quote}/preview_invoices/{preview_invoice}/lines` API.
func ListLines(params *stripe.QuotePreviewInvoiceListLinesParams) *InvoiceLineItemIter {
	return getC().ListLines(params)
}

// ListLines is the method for the `GET /v1/quotes/{quote}/preview_invoices/{preview_invoice}/lines` API.
func (c Client) ListLines(listParams *stripe.QuotePreviewInvoiceListLinesParams) *InvoiceLineItemIter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/preview_invoices/%s/lines",
		stripe.StringValue(listParams.Quote),
		stripe.StringValue(listParams.PreviewInvoice),
	)
	return &InvoiceLineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.InvoiceLineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// InvoiceLineItemIter is an iterator for invoice line items.
type InvoiceLineItemIter struct {
	*stripe.Iter
}

// InvoiceLineItem returns the invoice line item which the iterator is currently pointing to.
func (i *InvoiceLineItemIter) InvoiceLineItem() *stripe.InvoiceLineItem {
	return i.Current().(*stripe.InvoiceLineItem)
}

// InvoiceLineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *InvoiceLineItemIter) InvoiceLineItemList() *stripe.InvoiceLineItemList {
	return i.List().(*stripe.InvoiceLineItemList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
