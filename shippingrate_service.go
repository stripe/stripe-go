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

// v1ShippingRateService is used to invoke /v1/shipping_rates APIs.
type v1ShippingRateService struct {
	B   Backend
	Key string
}

// Creates a new shipping rate object.
func (c v1ShippingRateService) Create(ctx context.Context, params *ShippingRateCreateParams) (*ShippingRate, error) {
	if params == nil {
		params = &ShippingRateCreateParams{}
	}
	params.Context = ctx
	shippingrate := &ShippingRate{}
	err := c.B.Call(
		http.MethodPost, "/v1/shipping_rates", c.Key, params, shippingrate)
	return shippingrate, err
}

// Returns the shipping rate object with the given ID.
func (c v1ShippingRateService) Retrieve(ctx context.Context, id string, params *ShippingRateRetrieveParams) (*ShippingRate, error) {
	if params == nil {
		params = &ShippingRateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/shipping_rates/%s", id)
	shippingrate := &ShippingRate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, shippingrate)
	return shippingrate, err
}

// Updates an existing shipping rate object.
func (c v1ShippingRateService) Update(ctx context.Context, id string, params *ShippingRateUpdateParams) (*ShippingRate, error) {
	if params == nil {
		params = &ShippingRateUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/shipping_rates/%s", id)
	shippingrate := &ShippingRate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, shippingrate)
	return shippingrate, err
}

// Returns a list of your shipping rates.
func (c v1ShippingRateService) List(ctx context.Context, listParams *ShippingRateListParams) Seq2[*ShippingRate, error] {
	if listParams == nil {
		listParams = &ShippingRateListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*ShippingRate, ListContainer, error) {
		list := &ShippingRateList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/shipping_rates", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
