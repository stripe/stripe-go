//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1TaxLocationService is used to invoke /v1/tax/locations APIs.
type v1TaxLocationService struct {
	B   Backend
	Key string
}

// Create a tax location to use in calculating taxes for a service, ticket, or other type of product. The resulting object contains the id, address, name, description, and current operational status of the tax location.
func (c v1TaxLocationService) Create(ctx context.Context, params *TaxLocationCreateParams) (*TaxLocation, error) {
	if params == nil {
		params = &TaxLocationCreateParams{}
	}
	params.Context = ctx
	location := &TaxLocation{}
	err := c.B.Call(http.MethodPost, "/v1/tax/locations", c.Key, params, location)
	return location, err
}

// Fetch the details of a specific tax location using its unique identifier. Use a tax location to calculate taxes based on the location of the end product, such as a performance, instead of the customer address. For more details, check the [integration guide](https://docs.stripe.com/tax/tax-for-tickets/integration-guide).
func (c v1TaxLocationService) Retrieve(ctx context.Context, id string, params *TaxLocationRetrieveParams) (*TaxLocation, error) {
	if params == nil {
		params = &TaxLocationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax/locations/%s", id)
	location := &TaxLocation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, location)
	return location, err
}

// Retrieve a list of all tax locations. Tax locations can represent the venues for services, tickets, or other product types.
//
// The response includes detailed information for each tax location, such as its address, name, description, and current operational status.
//
// You can paginate through the list by using the limit parameter to control the number of results returned in each request.
func (c v1TaxLocationService) List(ctx context.Context, listParams *TaxLocationListParams) Seq2[*TaxLocation, error] {
	if listParams == nil {
		listParams = &TaxLocationListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*TaxLocation], error) {
		list := &v1Page[*TaxLocation]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/tax/locations", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
