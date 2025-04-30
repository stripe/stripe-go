//
//
// File generated from our OpenAPI spec
//
//

// Package quote provides the /v1/quotes APIs
package quote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/quotes APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B        stripe.Backend
	BUploads stripe.Backend
	Key      string
}

// A quote models prices and services for a customer. Default options for header, description, footer, and expires_at can be set in the dashboard via the [quote template](https://dashboard.stripe.com/settings/billing/quote).
func New(params *stripe.QuoteParams) (*stripe.Quote, error) {
	return getC().New(params)
}

// A quote models prices and services for a customer. Default options for header, description, footer, and expires_at can be set in the dashboard via the [quote template](https://dashboard.stripe.com/settings/billing/quote).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.QuoteParams) (*stripe.Quote, error) {
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, "/v1/quotes", c.Key, params, quote)
	return quote, err
}

// Retrieves the quote with the given ID.
func Get(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	return getC().Get(id, params)
}

// Retrieves the quote with the given ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, quote)
	return quote, err
}

// A quote models prices and services for a customer.
func Update(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	return getC().Update(id, params)
}

// A quote models prices and services for a customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Accepts the specified quote.
func Accept(id string, params *stripe.QuoteAcceptParams) (*stripe.Quote, error) {
	return getC().Accept(id, params)
}

// Accepts the specified quote.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Accept(id string, params *stripe.QuoteAcceptParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/accept", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Cancels the quote.
func Cancel(id string, params *stripe.QuoteCancelParams) (*stripe.Quote, error) {
	return getC().Cancel(id, params)
}

// Cancels the quote.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.QuoteCancelParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/cancel", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Finalizes the quote.
func FinalizeQuote(id string, params *stripe.QuoteFinalizeQuoteParams) (*stripe.Quote, error) {
	return getC().FinalizeQuote(id, params)
}

// Finalizes the quote.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) FinalizeQuote(id string, params *stripe.QuoteFinalizeQuoteParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/finalize", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Converts a stale quote to draft.
func MarkDraft(id string, params *stripe.QuoteMarkDraftParams) (*stripe.Quote, error) {
	return getC().MarkDraft(id, params)
}

// Converts a stale quote to draft.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) MarkDraft(id string, params *stripe.QuoteMarkDraftParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/mark_draft", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Converts a draft or open quote to stale.
func MarkStale(id string, params *stripe.QuoteMarkStaleParams) (*stripe.Quote, error) {
	return getC().MarkStale(id, params)
}

// Converts a draft or open quote to stale.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) MarkStale(id string, params *stripe.QuoteMarkStaleParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/mark_stale", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Download the PDF for a finalized quote. Explanation for special handling can be found [here](https://docs.stripe.com/quotes/overview#quote_pdf)
func PDF(id string, params *stripe.QuotePDFParams) (*stripe.APIStream, error) {
	return getC().PDF(id, params)
}

// Download the PDF for a finalized quote. Explanation for special handling can be found [here](https://docs.stripe.com/quotes/overview#quote_pdf)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) PDF(id string, params *stripe.QuotePDFParams) (*stripe.APIStream, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/pdf", id)
	stream := &stripe.APIStream{}
	err := c.BUploads.CallStreaming(http.MethodGet, path, c.Key, params, stream)
	return stream, err
}

// Recompute the upcoming invoice estimate for the quote.
func Reestimate(id string, params *stripe.QuoteReestimateParams) (*stripe.Quote, error) {
	return getC().Reestimate(id, params)
}

// Recompute the upcoming invoice estimate for the quote.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Reestimate(id string, params *stripe.QuoteReestimateParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/reestimate", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
}

// Returns a list of your quotes.
func List(params *stripe.QuoteListParams) *Iter {
	return getC().List(params)
}

// Returns a list of your quotes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.QuoteListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.QuoteList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/quotes", c.Key, []byte(b.Encode()), p, list)

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

// When retrieving a quote, there is an includable [computed.upfront.line_items](https://stripe.com/docs/api/quotes/object#quote_object-computed-upfront-line_items) property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of upfront line items.
func ListComputedUpfrontLineItems(params *stripe.QuoteListComputedUpfrontLineItemsParams) *LineItemIter {
	return getC().ListComputedUpfrontLineItems(params)
}

// When retrieving a quote, there is an includable [computed.upfront.line_items](https://stripe.com/docs/api/quotes/object#quote_object-computed-upfront-line_items) property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of upfront line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListComputedUpfrontLineItems(listParams *stripe.QuoteListComputedUpfrontLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/computed_upfront_line_items", stripe.StringValue(
			listParams.Quote))
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.LineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

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

// When retrieving a quote, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func ListLineItems(params *stripe.QuoteListLineItemsParams) *LineItemIter {
	return getC().ListLineItems(params)
}

// When retrieving a quote, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLineItems(listParams *stripe.QuoteListLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/line_items", stripe.StringValue(listParams.Quote))
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.LineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Retrieves a paginated list of lines for a quote. These lines describe changes that will be used to create new subscription schedules or update existing subscription schedules when the quote is accepted.
func ListLines(params *stripe.QuoteListLinesParams) *LineIter {
	return getC().ListLines(params)
}

// Retrieves a paginated list of lines for a quote. These lines describe changes that will be used to create new subscription schedules or update existing subscription schedules when the quote is accepted.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLines(listParams *stripe.QuoteListLinesParams) *LineIter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/lines", stripe.StringValue(listParams.Quote))
	return &LineIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.QuoteLineList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// LineIter is an iterator for quote lines.
type LineIter struct {
	*stripe.Iter
}

// QuoteLine returns the quote line which the iterator is currently pointing to.
func (i *LineIter) QuoteLine() *stripe.QuoteLine {
	return i.Current().(*stripe.QuoteLine)
}

// QuoteLineList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineIter) QuoteLineList() *stripe.QuoteLineList {
	return i.List().(*stripe.QuoteLineList)
}

// Preview the invoice line items that would be generated by accepting the quote.
func ListPreviewInvoiceLines(params *stripe.QuoteListPreviewInvoiceLinesParams) *InvoiceLineItemIter {
	return getC().ListPreviewInvoiceLines(params)
}

// Preview the invoice line items that would be generated by accepting the quote.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListPreviewInvoiceLines(listParams *stripe.QuoteListPreviewInvoiceLinesParams) *InvoiceLineItemIter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/preview_invoices/%s/lines", stripe.StringValue(
			listParams.Quote), stripe.StringValue(listParams.PreviewInvoice))
	return &InvoiceLineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.InvoiceLineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

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
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.GetBackend(stripe.UploadsBackend), stripe.Key}
}
