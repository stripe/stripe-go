//
//
// File generated from our OpenAPI spec
//
//

// Package location provides the /v1/tax/locations APIs
package location

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/form"
)

// Client is used to invoke /v1/tax/locations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a tax location to use in calculating taxes for a service, ticket, or other type of product. The resulting object contains the id, address, name, description, and current operational status of the tax location.
func New(params *stripe.TaxLocationParams) (*stripe.TaxLocation, error) {
	return getC().New(params)
}

// Create a tax location to use in calculating taxes for a service, ticket, or other type of product. The resulting object contains the id, address, name, description, and current operational status of the tax location.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TaxLocationParams) (*stripe.TaxLocation, error) {
	location := &stripe.TaxLocation{}
	err := c.B.Call(http.MethodPost, "/v1/tax/locations", c.Key, params, location)
	return location, err
}

// Fetch the details of a specific tax location using its unique identifier. Use a tax location to calculate taxes based on the location of the end product, such as a performance, instead of the customer address. For more details, check the [integration guide](https://docs.stripe.com/tax/tax-for-tickets/integration-guide).
func Get(id string, params *stripe.TaxLocationParams) (*stripe.TaxLocation, error) {
	return getC().Get(id, params)
}

// Fetch the details of a specific tax location using its unique identifier. Use a tax location to calculate taxes based on the location of the end product, such as a performance, instead of the customer address. For more details, check the [integration guide](https://docs.stripe.com/tax/tax-for-tickets/integration-guide).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TaxLocationParams) (*stripe.TaxLocation, error) {
	path := stripe.FormatURLPath("/v1/tax/locations/%s", id)
	location := &stripe.TaxLocation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, location)
	return location, err
}

// Retrieve a list of all tax locations. Tax locations can represent the venues for services, tickets, or other product types.
//
// The response includes detailed information for each tax location, such as its address, name, description, and current operational status.
//
// You can paginate through the list by using the limit parameter to control the number of results returned in each request.
func List(params *stripe.TaxLocationListParams) *Iter {
	return getC().List(params)
}

// Retrieve a list of all tax locations. Tax locations can represent the venues for services, tickets, or other product types.
//
// The response includes detailed information for each tax location, such as its address, name, description, and current operational status.
//
// You can paginate through the list by using the limit parameter to control the number of results returned in each request.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TaxLocationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxLocationList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/tax/locations", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for tax locations.
type Iter struct {
	*stripe.Iter
}

// TaxLocation returns the tax location which the iterator is currently pointing to.
func (i *Iter) TaxLocation() *stripe.TaxLocation {
	return i.Current().(*stripe.TaxLocation)
}

// TaxLocationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxLocationList() *stripe.TaxLocationList {
	return i.List().(*stripe.TaxLocationList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
