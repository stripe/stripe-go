//
//
// File generated from our OpenAPI spec
//
//

// Package quotepreviewinvoice provides the /quotes/{quote}/preview_invoices APIs
package quotepreviewinvoice

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
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

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
