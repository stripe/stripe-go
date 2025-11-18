//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1FxQuoteService is used to invoke /v1/fx_quotes APIs.
type v1FxQuoteService struct {
	B   Backend
	Key string
}

// Creates an FX Quote object
func (c v1FxQuoteService) Create(ctx context.Context, params *FxQuoteCreateParams) (*FxQuote, error) {
	if params == nil {
		params = &FxQuoteCreateParams{}
	}
	params.Context = ctx
	fxquote := &FxQuote{}
	err := c.B.Call(http.MethodPost, "/v1/fx_quotes", c.Key, params, fxquote)
	return fxquote, err
}

// Retrieve an FX Quote object
func (c v1FxQuoteService) Retrieve(ctx context.Context, id string, params *FxQuoteRetrieveParams) (*FxQuote, error) {
	if params == nil {
		params = &FxQuoteRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/fx_quotes/%s", id)
	fxquote := &FxQuote{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, fxquote)
	return fxquote, err
}

// Returns a list of FX quotes that have been issued. The FX quotes are returned in sorted order, with the most recent FX quotes appearing first.
func (c v1FxQuoteService) List(ctx context.Context, listParams *FxQuoteListParams) Seq2[*FxQuote, error] {
	if listParams == nil {
		listParams = &FxQuoteListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*FxQuote], error) {
		list := &v1Page[*FxQuote]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/fx_quotes", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
