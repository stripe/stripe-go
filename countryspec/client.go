//
//
// File generated from our OpenAPI spec
//
//

// Package countryspec provides the /country_specs APIs
package countryspec

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/form"
)

// Client is used to invoke /country_specs APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a country spec.
func Get(id string, params *stripe.CountrySpecParams) (*stripe.CountrySpec, error) {
	return getC().Get(id, params)
}

// Get returns the details of a country spec.
func (c Client) Get(id string, params *stripe.CountrySpecParams) (*stripe.CountrySpec, error) {
	path := stripe.FormatURLPath("/v1/country_specs/%s", id)
	countryspec := &stripe.CountrySpec{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, countryspec)
	return countryspec, err
}

// List returns a list of country specs.
func List(params *stripe.CountrySpecListParams) *Iter {
	return getC().List(params)
}

// List returns a list of country specs.
func (c Client) List(listParams *stripe.CountrySpecListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CountrySpecList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/country_specs", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for country specs.
type Iter struct {
	*stripe.Iter
}

// CountrySpec returns the country spec which the iterator is currently pointing to.
func (i *Iter) CountrySpec() *stripe.CountrySpec {
	return i.Current().(*stripe.CountrySpec)
}

// CountrySpecList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CountrySpecList() *stripe.CountrySpecList {
	return i.List().(*stripe.CountrySpecList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
