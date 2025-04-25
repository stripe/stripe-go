//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1QuoteService is used to invoke /v1/quotes APIs.
type v1QuoteService struct {
	B        Backend
	BUploads Backend
	Key      string
}

// A quote models prices and services for a customer. Default options for header, description, footer, and expires_at can be set in the dashboard via the [quote template](https://dashboard.stripe.com/settings/billing/quote).
func (c v1QuoteService) Create(ctx context.Context, params *QuoteCreateParams) (*Quote, error) {
	quote := &Quote{}
	if params == nil {
		params = &QuoteCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, "/v1/quotes", c.Key, params, quote)
	return quote, err
}

// Retrieves the quote with the given ID.
func (c v1QuoteService) Retrieve(ctx context.Context, id string, params *QuoteRetrieveParams) (*Quote, error) {
	path := FormatURLPath("/v1/quotes/%s", id)
	quote := &Quote{}
	if params == nil {
		params = &QuoteRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, quote)
	return quote, err
}

// A quote models prices and services for a customer.
func (c v1QuoteService) Update(ctx context.Context, id string, params *QuoteUpdateParams) (*Quote, error) {
	path := FormatURLPath("/v1/quotes/%s", id)
	quote := &Quote{}
	if params == nil {
		params = &QuoteUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Accepts the specified quote.
func (c v1QuoteService) Accept(ctx context.Context, id string, params *QuoteAcceptParams) (*Quote, error) {
	path := FormatURLPath("/v1/quotes/%s/accept", id)
	quote := &Quote{}
	if params == nil {
		params = &QuoteAcceptParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Cancels the quote.
func (c v1QuoteService) Cancel(ctx context.Context, id string, params *QuoteCancelParams) (*Quote, error) {
	path := FormatURLPath("/v1/quotes/%s/cancel", id)
	quote := &Quote{}
	if params == nil {
		params = &QuoteCancelParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Finalizes the quote.
func (c v1QuoteService) FinalizeQuote(ctx context.Context, id string, params *QuoteFinalizeQuoteParams) (*Quote, error) {
	path := FormatURLPath("/v1/quotes/%s/finalize", id)
	quote := &Quote{}
	if params == nil {
		params = &QuoteFinalizeQuoteParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Download the PDF for a finalized quote. Explanation for special handling can be found [here](https://docs.stripe.com/quotes/overview#quote_pdf)
func (c v1QuoteService) PDF(ctx context.Context, id string, params *QuotePDFParams) (*APIStream, error) {
	path := FormatURLPath("/v1/quotes/%s/pdf", id)
	stream := &APIStream{}
	if params == nil {
		params = &QuotePDFParams{}
	}
	params.Context = ctx
	err := c.BUploads.CallStreaming(http.MethodGet, path, c.Key, params, stream)
	return stream, err
}

// Returns a list of your quotes.
func (c v1QuoteService) List(ctx context.Context, listParams *QuoteListParams) Seq2[*Quote, error] {
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Quote, ListContainer, error) {
		list := &QuoteList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/quotes", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// When retrieving a quote, there is an includable [computed.upfront.line_items](https://stripe.com/docs/api/quotes/object#quote_object-computed-upfront-line_items) property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of upfront line items.
func (c v1QuoteService) ListComputedUpfrontLineItems(ctx context.Context, listParams *QuoteListComputedUpfrontLineItemsParams) Seq2[*LineItem, error] {
	path := FormatURLPath(
		"/v1/quotes/%s/computed_upfront_line_items", StringValue(listParams.Quote))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*LineItem, ListContainer, error) {
		list := &LineItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// When retrieving a quote, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func (c v1QuoteService) ListLineItems(ctx context.Context, listParams *QuoteListLineItemsParams) Seq2[*LineItem, error] {
	path := FormatURLPath(
		"/v1/quotes/%s/line_items", StringValue(listParams.Quote))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*LineItem, ListContainer, error) {
		list := &LineItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
