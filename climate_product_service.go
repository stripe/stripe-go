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

// v1ClimateProductService is used to invoke /v1/climate/products APIs.
type v1ClimateProductService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Climate product with the given ID.
func (c v1ClimateProductService) Retrieve(ctx context.Context, id string, params *ClimateProductRetrieveParams) (*ClimateProduct, error) {
	if params == nil {
		params = &ClimateProductRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/climate/products/%s", id)
	product := &ClimateProduct{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, product)
	return product, err
}

// Lists all available Climate product objects.
func (c v1ClimateProductService) List(ctx context.Context, listParams *ClimateProductListParams) Seq2[*ClimateProduct, error] {
	if listParams == nil {
		listParams = &ClimateProductListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*ClimateProduct, ListContainer, error) {
		list := &ClimateProductList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/climate/products", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
