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

// v1TaxRateService is used to invoke /v1/tax_rates APIs.
type v1TaxRateService struct {
	B   Backend
	Key string
}

// Creates a new tax rate.
func (c v1TaxRateService) Create(ctx context.Context, params *TaxRateCreateParams) (*TaxRate, error) {
	if params == nil {
		params = &TaxRateCreateParams{}
	}
	params.Context = ctx
	taxrate := &TaxRate{}
	err := c.B.Call(http.MethodPost, "/v1/tax_rates", c.Key, params, taxrate)
	return taxrate, err
}

// Retrieves a tax rate with the given ID
func (c v1TaxRateService) Retrieve(ctx context.Context, id string, params *TaxRateRetrieveParams) (*TaxRate, error) {
	if params == nil {
		params = &TaxRateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax_rates/%s", id)
	taxrate := &TaxRate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, taxrate)
	return taxrate, err
}

// Updates an existing tax rate.
func (c v1TaxRateService) Update(ctx context.Context, id string, params *TaxRateUpdateParams) (*TaxRate, error) {
	if params == nil {
		params = &TaxRateUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax_rates/%s", id)
	taxrate := &TaxRate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, taxrate)
	return taxrate, err
}

// Returns a list of your tax rates. Tax rates are returned sorted by creation date, with the most recently created tax rates appearing first.
func (c v1TaxRateService) List(ctx context.Context, listParams *TaxRateListParams) Seq2[*TaxRate, error] {
	if listParams == nil {
		listParams = &TaxRateListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TaxRate, ListContainer, error) {
		list := &TaxRateList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/tax_rates", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
