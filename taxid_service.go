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

// v1TaxIDService is used to invoke /v1/tax_ids APIs.
type v1TaxIDService struct {
	B   Backend
	Key string
}

// Creates a new tax_id object for a customer.
func (c v1TaxIDService) Create(ctx context.Context, params *TaxIDCreateParams) (*TaxID, error) {
	path := "/v1/tax_ids"
	if params == nil {
		params = &TaxIDCreateParams{}
	}
	params.Context = ctx
	if params.Customer != nil {
		path = FormatURLPath(
			"/v1/customers/%s/tax_ids", StringValue(params.Customer))
	}
	taxid := &TaxID{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, taxid)
	return taxid, err
}

// Retrieves the tax_id object with the given identifier.
func (c v1TaxIDService) Retrieve(ctx context.Context, id string, params *TaxIDRetrieveParams) (*TaxID, error) {
	path := FormatURLPath("/v1/tax_ids/%s", id)
	if params == nil {
		params = &TaxIDRetrieveParams{}
	}
	params.Context = ctx
	if params.Customer != nil {
		path = FormatURLPath(
			"/v1/customers/%s/tax_ids/%s", StringValue(params.Customer), id)
	}
	taxid := &TaxID{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, taxid)
	return taxid, err
}

// Deletes an existing tax_id object.
func (c v1TaxIDService) Delete(ctx context.Context, id string, params *TaxIDDeleteParams) (*TaxID, error) {
	path := FormatURLPath("/v1/tax_ids/%s", id)
	if params == nil {
		params = &TaxIDDeleteParams{}
	}
	params.Context = ctx
	if params.Customer != nil {
		path = FormatURLPath(
			"/v1/customers/%s/tax_ids/%s", StringValue(params.Customer), id)
	}
	taxid := &TaxID{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, taxid)
	return taxid, err
}

// Returns a list of tax IDs for a customer.
func (c v1TaxIDService) List(ctx context.Context, listParams *TaxIDListParams) Seq2[*TaxID, error] {
	path := "/v1/tax_ids"
	if listParams != nil && listParams.Customer != nil {
		path = FormatURLPath(
			"/v1/customers/%s/tax_ids", StringValue(listParams.Customer))
	}
	if listParams == nil {
		listParams = &TaxIDListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TaxID, ListContainer, error) {
		list := &TaxIDList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
