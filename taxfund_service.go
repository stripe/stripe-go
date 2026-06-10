//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v86/form"
)

// v1TaxFundService is used to invoke /v1/tax_funds APIs.
type v1TaxFundService struct {
	B   Backend
	Key string
}

// Retrieves a tax fund object by its ID.
func (c v1TaxFundService) Retrieve(ctx context.Context, id string, params *TaxFundRetrieveParams) (*TaxFund, error) {
	if params == nil {
		params = &TaxFundRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax_funds/%s", id)
	taxfund := &TaxFund{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, taxfund)
	return taxfund, err
}

// Returns a list of tax funds in reverse chronological order.
func (c v1TaxFundService) List(ctx context.Context, listParams *TaxFundListParams) *V1List[*TaxFund] {
	if listParams == nil {
		listParams = &TaxFundListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*TaxFund], error) {
		list := &v1Page[*TaxFund]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/tax_funds", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
