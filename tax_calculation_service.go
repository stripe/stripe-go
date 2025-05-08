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

// v1TaxCalculationService is used to invoke /v1/tax/calculations APIs.
type v1TaxCalculationService struct {
	B   Backend
	Key string
}

// Calculates tax based on the input and returns a Tax Calculation object.
func (c v1TaxCalculationService) Create(ctx context.Context, params *TaxCalculationCreateParams) (*TaxCalculation, error) {
	if params == nil {
		params = &TaxCalculationCreateParams{}
	}
	params.Context = ctx
	calculation := &TaxCalculation{}
	err := c.B.Call(
		http.MethodPost, "/v1/tax/calculations", c.Key, params, calculation)
	return calculation, err
}

// Retrieves a Tax Calculation object, if the calculation hasn't expired.
func (c v1TaxCalculationService) Retrieve(ctx context.Context, id string, params *TaxCalculationRetrieveParams) (*TaxCalculation, error) {
	if params == nil {
		params = &TaxCalculationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax/calculations/%s", id)
	calculation := &TaxCalculation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, calculation)
	return calculation, err
}

// Retrieves the line items of a tax calculation as a collection, if the calculation hasn't expired.
func (c v1TaxCalculationService) ListLineItems(ctx context.Context, listParams *TaxCalculationListLineItemsParams) Seq2[*TaxCalculationLineItem, error] {
	if listParams == nil {
		listParams = &TaxCalculationListLineItemsParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/tax/calculations/%s/line_items", StringValue(listParams.Calculation))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TaxCalculationLineItem, ListContainer, error) {
		list := &TaxCalculationLineItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
