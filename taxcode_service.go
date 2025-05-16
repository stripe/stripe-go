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

// v1TaxCodeService is used to invoke /v1/tax_codes APIs.
type v1TaxCodeService struct {
	B   Backend
	Key string
}

// Retrieves the details of an existing tax code. Supply the unique tax code ID and Stripe will return the corresponding tax code information.
func (c v1TaxCodeService) Retrieve(ctx context.Context, id string, params *TaxCodeRetrieveParams) (*TaxCode, error) {
	if params == nil {
		params = &TaxCodeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax_codes/%s", id)
	taxcode := &TaxCode{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, taxcode)
	return taxcode, err
}

// A list of [all tax codes available](https://stripe.com/docs/tax/tax-categories) to add to Products in order to allow specific tax calculations.
func (c v1TaxCodeService) List(ctx context.Context, listParams *TaxCodeListParams) Seq2[*TaxCode, error] {
	if listParams == nil {
		listParams = &TaxCodeListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TaxCode, ListContainer, error) {
		list := &TaxCodeList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/tax_codes", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
