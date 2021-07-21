//
//
// File generated from our OpenAPI spec
//
//

// Package quote provides the /quotes APIs
package quote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /quotes APIs.
type Client struct {
	B          stripe.Backend
	PDFBackend stripe.Backend
	Key        string
}

// New creates a new quote.
func New(params *stripe.QuoteParams) (*stripe.Quote, error) {
	return getC().New(params)
}

// New creates a new quote.
func (c Client) New(params *stripe.QuoteParams) (*stripe.Quote, error) {
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, "/v1/quotes", c.Key, params, quote)
	return quote, err
}

// Get returns the details of a quote.
func Get(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	return getC().Get(id, params)
}

// Get returns the details of a quote.
func (c Client) Get(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, quote)
	return quote, err
}

// Update updates a quote's properties.
func Update(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	return getC().Update(id, params)
}

// Update updates a quote's properties.
func (c Client) Update(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Accept is the method for the `POST /v1/quotes/{quote}/accept` API.
func Accept(id string, params *stripe.QuoteAcceptParams) (*stripe.Quote, error) {
	return getC().Accept(id, params)
}

// Accept is the method for the `POST /v1/quotes/{quote}/accept` API.
func (c Client) Accept(id string, params *stripe.QuoteAcceptParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/accept", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Cancel is the method for the `POST /v1/quotes/{quote}/cancel` API.
func Cancel(id string, params *stripe.QuoteCancelParams) (*stripe.Quote, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/quotes/{quote}/cancel` API.
func (c Client) Cancel(id string, params *stripe.QuoteCancelParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/cancel", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// FinalizeQuote is the method for the `POST /v1/quotes/{quote}/finalize` API.
func FinalizeQuote(id string, params *stripe.QuoteFinalizeQuoteParams) (*stripe.Quote, error) {
	return getC().FinalizeQuote(id, params)
}

// FinalizeQuote is the method for the `POST /v1/quotes/{quote}/finalize` API.
func (c Client) FinalizeQuote(id string, params *stripe.QuoteFinalizeQuoteParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/finalize", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// PDF is the method for the `GET /v1/quotes/{quote}/pdf` API.
func PDF(id string, params *stripe.QuotePDFParams) (*stripe.APIStream, error) {
	return getC().PDF(id, params)
}

// PDF is the method for the `GET /v1/quotes/{quote}/pdf` API.
func (c Client) PDF(id string, params *stripe.QuotePDFParams) (*stripe.APIStream, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/pdf", id)
	stream := &stripe.APIStream{}
	err := c.PDFBackend.CallStreaming(http.MethodGet, path, c.Key, params, stream)
	return stream, err
}

// List returns a list of quotes.
func List(params *stripe.QuoteListParams) *Iter {
	return getC().List(params)
}

// List returns a list of quotes.
func (c Client) List(listParams *stripe.QuoteListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.QuoteList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/quotes", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for quotes.
type Iter struct {
	*stripe.Iter
}

// Quote returns the quote which the iterator is currently pointing to.
func (i *Iter) Quote() *stripe.Quote {
	return i.Current().(*stripe.Quote)
}

// QuoteList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) QuoteList() *stripe.QuoteList {
	return i.List().(*stripe.QuoteList)
}

// ListComputedUpfrontLineItems is the method for the `GET /v1/quotes/{quote}/computed_upfront_line_items` API.
func ListComputedUpfrontLineItems(params *stripe.QuoteListComputedUpfrontLineItemsParams) *LineItemIter {
	return getC().ListComputedUpfrontLineItems(params)
}

// ListComputedUpfrontLineItems is the method for the `GET /v1/quotes/{quote}/computed_upfront_line_items` API.
func (c Client) ListComputedUpfrontLineItems(listParams *stripe.QuoteListComputedUpfrontLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/computed_upfront_line_items",
		stripe.StringValue(listParams.Quote),
	)
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.LineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// ListLineItems is the method for the `GET /v1/quotes/{quote}/line_items` API.
func ListLineItems(params *stripe.QuoteListLineItemsParams) *LineItemIter {
	return getC().ListLineItems(params)
}

// ListLineItems is the method for the `GET /v1/quotes/{quote}/line_items` API.
func (c Client) ListLineItems(listParams *stripe.QuoteListLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/line_items",
		stripe.StringValue(listParams.Quote),
	)
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.LineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// LineItemIter is an iterator for line items.
type LineItemIter struct {
	*stripe.Iter
}

// LineItem returns the line item which the iterator is currently pointing to.
func (i *LineItemIter) LineItem() *stripe.LineItem {
	return i.Current().(*stripe.LineItem)
}

// LineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) LineItemList() *stripe.LineItemList {
	return i.List().(*stripe.LineItemList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.GetBackend(stripe.UploadsBackend), stripe.Key}
}
