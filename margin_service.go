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

// v1MarginService is used to invoke /v1/billing/margins APIs.
type v1MarginService struct {
	B   Backend
	Key string
}

// Create a margin object to be used with invoices, invoice items, and invoice line items for a customer to represent a partner discount. A margin has a percent_off which is the percent that will be taken off the subtotal after all items and other discounts and promotions) of any invoices for a customer. Calculation of prorations do not include any partner margins applied on the original invoice item.
func (c v1MarginService) Create(ctx context.Context, params *MarginCreateParams) (*Margin, error) {
	if params == nil {
		params = &MarginCreateParams{}
	}
	params.Context = ctx
	margin := &Margin{}
	err := c.B.Call(http.MethodPost, "/v1/billing/margins", c.Key, params, margin)
	return margin, err
}

// Retrieve a margin object with the given ID.
func (c v1MarginService) Retrieve(ctx context.Context, id string, params *MarginRetrieveParams) (*Margin, error) {
	if params == nil {
		params = &MarginRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/margins/%s", id)
	margin := &Margin{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, margin)
	return margin, err
}

// Update the specified margin object. Certain fields of the margin object are not editable.
func (c v1MarginService) Update(ctx context.Context, id string, params *MarginUpdateParams) (*Margin, error) {
	if params == nil {
		params = &MarginUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/margins/%s", id)
	margin := &Margin{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, margin)
	return margin, err
}

// Retrieve a list of your margins.
func (c v1MarginService) List(ctx context.Context, listParams *MarginListParams) Seq2[*Margin, error] {
	if listParams == nil {
		listParams = &MarginListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*Margin], error) {
		list := &v1Page[*Margin]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/billing/margins", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
