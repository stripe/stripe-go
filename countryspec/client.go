// Package countryspec provides the /country_specs APIs
package countryspec

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /country_specs and countryspec-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns a CountrySpec for a given country code
// For more details see https://stripe.com/docs/api/ruby#retrieve_country_spec
func Get(country string) (*stripe.CountrySpec, error) {
	return getC().Get(country)
}

func (c Client) Get(country string) (*stripe.CountrySpec, error) {
	path := stripe.FormatURLPath("/country_specs/%s", country)
	countrySpec := &stripe.CountrySpec{}
	err := c.B.Call("GET", path, c.Key, nil, countrySpec)
	return countrySpec, err
}

// List lists available CountrySpecs.
func List(params *stripe.CountrySpecListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.CountrySpecListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.CountrySpecList{}
		err := c.B.CallRaw("GET", "/country_specs", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of CountrySpecs.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// CountrySpec returns the most recent CountrySpec
// visited by a call to Next.
func (i *Iter) CountrySpec() *stripe.CountrySpec {
	return i.Current().(*stripe.CountrySpec)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
