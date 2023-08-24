//
//
// File generated from our OpenAPI spec
//
//

// Package quote provides the /quotes APIs
package quote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/form"
)

// Client is used to invoke /quotes APIs.
type Client struct {
	B        stripe.Backend
	BUploads stripe.Backend
	Key      string
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

// DraftQuote is the method for the `POST /v1/quotes/{quote}/mark_draft` API.
func DraftQuote(id string, params *stripe.QuoteDraftQuoteParams) (*stripe.Quote, error) {
	return getC().DraftQuote(id, params)
}

// DraftQuote is the method for the `POST /v1/quotes/{quote}/mark_draft` API.
func (c Client) DraftQuote(id string, params *stripe.QuoteDraftQuoteParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/mark_draft", id)
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

// MarkStaleQuote is the method for the `POST /v1/quotes/{quote}/mark_stale` API.
func MarkStaleQuote(id string, params *stripe.QuoteMarkStaleQuoteParams) (*stripe.Quote, error) {
	return getC().MarkStaleQuote(id, params)
}

// MarkStaleQuote is the method for the `POST /v1/quotes/{quote}/mark_stale` API.
func (c Client) MarkStaleQuote(id string, params *stripe.QuoteMarkStaleQuoteParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/mark_stale", id)
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
	err := c.BUploads.CallStreaming(http.MethodGet, path, c.Key, params, stream)
	return stream, err
}

// Reestimate is the method for the `POST /v1/quotes/{quote}/reestimate` API.
func Reestimate(id string, params *stripe.QuoteReestimateParams) (*stripe.Quote, error) {
	return getC().Reestimate(id, params)
}

// Reestimate is the method for the `POST /v1/quotes/{quote}/reestimate` API.
func (c Client) Reestimate(id string, params *stripe.QuoteReestimateParams) (*stripe.Quote, error) {
	path := stripe.FormatURLPath("/v1/quotes/%s/reestimate", id)
	quote := &stripe.Quote{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, quote)
	return quote, err
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

// ListLines is the method for the `GET /v1/quotes/{quote}/lines` API.
func ListLines(params *stripe.QuoteListLinesParams) *LineIter {
	return getC().ListLines(params)
}

// ListLines is the method for the `GET /v1/quotes/{quote}/lines` API.
func (c Client) ListLines(listParams *stripe.QuoteListLinesParams) *LineIter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/lines",
		stripe.StringValue(listParams.Quote),
	)
	return &LineIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.QuoteLineList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

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

// ListPreviewInvoiceLines is the method for the `GET /v1/quotes/{quote}/preview_invoices/{preview_invoice}/lines` API.
func ListPreviewInvoiceLines(params *stripe.QuoteListPreviewInvoiceLinesParams) *InvoiceLineItemIter {
	return getC().ListPreviewInvoiceLines(params)
}

// ListPreviewInvoiceLines is the method for the `GET /v1/quotes/{quote}/preview_invoices/{preview_invoice}/lines` API.
func (c Client) ListPreviewInvoiceLines(listParams *stripe.QuoteListPreviewInvoiceLinesParams) *InvoiceLineItemIter {
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

// ListPreviewInvoices is the method for the `GET /v1/quotes/{quote}/preview_invoices` API.
func ListPreviewInvoices(params *stripe.QuoteListPreviewInvoicesParams) *PreviewInvoiceIter {
	return getC().ListPreviewInvoices(params)
}

// ListPreviewInvoices is the method for the `GET /v1/quotes/{quote}/preview_invoices` API.
func (c Client) ListPreviewInvoices(listParams *stripe.QuoteListPreviewInvoicesParams) *PreviewInvoiceIter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/preview_invoices",
		stripe.StringValue(listParams.Quote),
	)
	return &PreviewInvoiceIter{
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

// PreviewInvoiceIter is an iterator for quote preview invoices.
type PreviewInvoiceIter struct {
	*stripe.Iter
}

// QuotePreviewInvoice returns the quote preview invoice which the iterator is currently pointing to.
func (i *PreviewInvoiceIter) QuotePreviewInvoice() *stripe.QuotePreviewInvoice {
	return i.Current().(*stripe.QuotePreviewInvoice)
}

// QuotePreviewInvoiceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *PreviewInvoiceIter) QuotePreviewInvoiceList() *stripe.QuotePreviewInvoiceList {
	return i.List().(*stripe.QuotePreviewInvoiceList)
}

// ListPreviewSubscriptionSchedules is the method for the `GET /v1/quotes/{quote}/preview_subscription_schedules` API.
func ListPreviewSubscriptionSchedules(params *stripe.QuoteListPreviewSubscriptionSchedulesParams) *PreviewScheduleIter {
	return getC().ListPreviewSubscriptionSchedules(params)
}

// ListPreviewSubscriptionSchedules is the method for the `GET /v1/quotes/{quote}/preview_subscription_schedules` API.
func (c Client) ListPreviewSubscriptionSchedules(listParams *stripe.QuoteListPreviewSubscriptionSchedulesParams) *PreviewScheduleIter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/preview_subscription_schedules",
		stripe.StringValue(listParams.Quote),
	)
	return &PreviewScheduleIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.QuotePreviewScheduleList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// PreviewScheduleIter is an iterator for quote preview schedules.
type PreviewScheduleIter struct {
	*stripe.Iter
}

// QuotePreviewSchedule returns the quote preview schedule which the iterator is currently pointing to.
func (i *PreviewScheduleIter) QuotePreviewSchedule() *stripe.QuotePreviewSchedule {
	return i.Current().(*stripe.QuotePreviewSchedule)
}

// QuotePreviewScheduleList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *PreviewScheduleIter) QuotePreviewScheduleList() *stripe.QuotePreviewScheduleList {
	return i.List().(*stripe.QuotePreviewScheduleList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.GetBackend(stripe.UploadsBackend), stripe.Key}
}
