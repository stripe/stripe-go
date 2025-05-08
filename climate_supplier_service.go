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

// v1ClimateSupplierService is used to invoke /v1/climate/suppliers APIs.
type v1ClimateSupplierService struct {
	B   Backend
	Key string
}

// Retrieves a Climate supplier object.
func (c v1ClimateSupplierService) Retrieve(ctx context.Context, id string, params *ClimateSupplierRetrieveParams) (*ClimateSupplier, error) {
	if params == nil {
		params = &ClimateSupplierRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/climate/suppliers/%s", id)
	supplier := &ClimateSupplier{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, supplier)
	return supplier, err
}

// Lists all available Climate supplier objects.
func (c v1ClimateSupplierService) List(ctx context.Context, listParams *ClimateSupplierListParams) Seq2[*ClimateSupplier, error] {
	if listParams == nil {
		listParams = &ClimateSupplierListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*ClimateSupplier, ListContainer, error) {
		list := &ClimateSupplierList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/climate/suppliers", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
