//
//
// File generated from our OpenAPI spec
//
//

// Package quotephase provides the /quote_phases APIs
package quotephase

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/form"
)

// Client is used to invoke /quote_phases APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the quote phase with the given ID.
func Get(id string, params *stripe.QuotePhaseParams) (*stripe.QuotePhase, error) {
	return getC().Get(id, params)
}

// Retrieves the quote phase with the given ID.
func (c Client) Get(id string, params *stripe.QuotePhaseParams) (*stripe.QuotePhase, error) {
	path := stripe.FormatURLPath("/v1/quote_phases/%s", id)
	quotephase := &stripe.QuotePhase{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, quotephase)
	return quotephase, err
}

// Returns a list of quote phases.
func List(params *stripe.QuotePhaseListParams) *Iter {
	return getC().List(params)
}

// Returns a list of quote phases.
func (c Client) List(listParams *stripe.QuotePhaseListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.QuotePhaseList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/quote_phases", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for quote phases.
type Iter struct {
	*stripe.Iter
}

// QuotePhase returns the quote phase which the iterator is currently pointing to.
func (i *Iter) QuotePhase() *stripe.QuotePhase {
	return i.Current().(*stripe.QuotePhase)
}

// QuotePhaseList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) QuotePhaseList() *stripe.QuotePhaseList {
	return i.List().(*stripe.QuotePhaseList)
}

// When retrieving a quote phase, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func ListLineItems(params *stripe.QuotePhaseListLineItemsParams) *LineItemIter {
	return getC().ListLineItems(params)
}

// When retrieving a quote phase, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func (c Client) ListLineItems(listParams *stripe.QuotePhaseListLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/quote_phases/%s/line_items",
		stripe.StringValue(listParams.QuotePhase),
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
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
