//
//
// File generated from our OpenAPI spec
//
//

// Package location provides the /v1/terminal/locations APIs
package location

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/terminal/locations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new Location object.
// For further details, including which address fields are required in each country, see the [Manage locations](https://docs.stripe.com/docs/terminal/fleet/locations) guide.
func New(params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	return getC().New(params)
}

// Creates a new Location object.
// For further details, including which address fields are required in each country, see the [Manage locations](https://docs.stripe.com/docs/terminal/fleet/locations) guide.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	location := &stripe.TerminalLocation{}
	err := c.B.Call(
		http.MethodPost, "/v1/terminal/locations", c.Key, params, location)
	return location, err
}

// Retrieves a Location object.
func Get(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	return getC().Get(id, params)
}

// Retrieves a Location object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	path := stripe.FormatURLPath("/v1/terminal/locations/%s", id)
	location := &stripe.TerminalLocation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, location)
	return location, err
}

// Updates a Location object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	return getC().Update(id, params)
}

// Updates a Location object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	path := stripe.FormatURLPath("/v1/terminal/locations/%s", id)
	location := &stripe.TerminalLocation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, location)
	return location, err
}

// Deletes a Location object.
func Del(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	return getC().Del(id, params)
}

// Deletes a Location object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	path := stripe.FormatURLPath("/v1/terminal/locations/%s", id)
	location := &stripe.TerminalLocation{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, location)
	return location, err
}

// Returns a list of Location objects.
func List(params *stripe.TerminalLocationListParams) *Iter {
	return getC().List(params)
}

// Returns a list of Location objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TerminalLocationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TerminalLocationList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/terminal/locations", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for terminal locations.
type Iter struct {
	*stripe.Iter
}

// TerminalLocation returns the terminal location which the iterator is currently pointing to.
func (i *Iter) TerminalLocation() *stripe.TerminalLocation {
	return i.Current().(*stripe.TerminalLocation)
}

// TerminalLocationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TerminalLocationList() *stripe.TerminalLocationList {
	return i.List().(*stripe.TerminalLocationList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
