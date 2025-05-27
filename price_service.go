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

// v1PriceService is used to invoke /v1/prices APIs.
type v1PriceService struct {
	B   Backend
	Key string
}

// Creates a new [Price for an existing <a href="https://docs.stripe.com/api/products">Product](https://docs.stripe.com/api/prices). The Price can be recurring or one-time.
func (c v1PriceService) Create(ctx context.Context, params *PriceCreateParams) (*Price, error) {
	if params == nil {
		params = &PriceCreateParams{}
	}
	params.Context = ctx
	price := &Price{}
	err := c.B.Call(http.MethodPost, "/v1/prices", c.Key, params, price)
	return price, err
}

// Retrieves the price with the given ID.
func (c v1PriceService) Retrieve(ctx context.Context, id string, params *PriceRetrieveParams) (*Price, error) {
	if params == nil {
		params = &PriceRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/prices/%s", id)
	price := &Price{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, price)
	return price, err
}

// Updates the specified price by setting the values of the parameters passed. Any parameters not provided are left unchanged.
func (c v1PriceService) Update(ctx context.Context, id string, params *PriceUpdateParams) (*Price, error) {
	if params == nil {
		params = &PriceUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/prices/%s", id)
	price := &Price{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, price)
	return price, err
}

// Returns a list of your active prices, excluding [inline prices](https://docs.stripe.com/docs/products-prices/pricing-models#inline-pricing). For the list of inactive prices, set active to false.
func (c v1PriceService) List(ctx context.Context, listParams *PriceListParams) Seq2[*Price, error] {
	if listParams == nil {
		listParams = &PriceListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Price, ListContainer, error) {
		list := &PriceList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/prices", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// Search for prices you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func (c v1PriceService) Search(ctx context.Context, params *PriceSearchParams) Seq2[*Price, error] {
	if params == nil {
		params = &PriceSearchParams{}
	}
	params.Context = ctx
	return newV1SearchList(params, func(p *Params, b *form.Values) ([]*Price, SearchContainer, error) {
		list := &PriceSearchResult{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/prices/search", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
