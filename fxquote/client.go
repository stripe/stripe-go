//
//
// File generated from our OpenAPI spec
//
//

// Package fxquote provides the /v1/fx_quotes APIs
package fxquote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/fx_quotes APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an FX Quote object
func New(params *stripe.FxQuoteParams) (*stripe.FxQuote, error) {
	return getC().New(params)
}

// Creates an FX Quote object
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.FxQuoteParams) (*stripe.FxQuote, error) {
	fxquote := &stripe.FxQuote{}
	err := c.B.Call(http.MethodPost, "/v1/fx_quotes", c.Key, params, fxquote)
	return fxquote, err
}

// Retrieve an FX Quote object
func Get(id string, params *stripe.FxQuoteParams) (*stripe.FxQuote, error) {
	return getC().Get(id, params)
}

// Retrieve an FX Quote object
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.FxQuoteParams) (*stripe.FxQuote, error) {
	path := stripe.FormatURLPath("/v1/fx_quotes/%s", id)
	fxquote := &stripe.FxQuote{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, fxquote)
	return fxquote, err
}

// Returns a list of FX quotes that have been issued. The FX quotes are returned in sorted order, with the most recent FX quotes appearing first.
func List(params *stripe.FxQuoteListParams) *Iter {
	return getC().List(params)
}

// Returns a list of FX quotes that have been issued. The FX quotes are returned in sorted order, with the most recent FX quotes appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.FxQuoteListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.FxQuoteList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/fx_quotes", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for fx quotes.
type Iter struct {
	*stripe.Iter
}

// FxQuote returns the fx quote which the iterator is currently pointing to.
func (i *Iter) FxQuote() *stripe.FxQuote {
	return i.Current().(*stripe.FxQuote)
}

// FxQuoteList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FxQuoteList() *stripe.FxQuoteList {
	return i.List().(*stripe.FxQuoteList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
