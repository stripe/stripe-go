//
//
// File generated from our OpenAPI spec
//
//

// Package configuration provides the /v1/terminal/configurations APIs
package configuration

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/terminal/configurations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new Configuration object.
func New(params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	return getC().New(params)
}

// Creates a new Configuration object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	configuration := &stripe.TerminalConfiguration{}
	err := c.B.Call(
		http.MethodPost, "/v1/terminal/configurations", c.Key, params, configuration)
	return configuration, err
}

// Retrieves a Configuration object.
func Get(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	return getC().Get(id, params)
}

// Retrieves a Configuration object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	path := stripe.FormatURLPath("/v1/terminal/configurations/%s", id)
	configuration := &stripe.TerminalConfiguration{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, configuration)
	return configuration, err
}

// Updates a new Configuration object.
func Update(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	return getC().Update(id, params)
}

// Updates a new Configuration object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	path := stripe.FormatURLPath("/v1/terminal/configurations/%s", id)
	configuration := &stripe.TerminalConfiguration{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, configuration)
	return configuration, err
}

// Deletes a Configuration object.
func Del(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	return getC().Del(id, params)
}

// Deletes a Configuration object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	path := stripe.FormatURLPath("/v1/terminal/configurations/%s", id)
	configuration := &stripe.TerminalConfiguration{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, configuration)
	return configuration, err
}

// Returns a list of Configuration objects.
func List(params *stripe.TerminalConfigurationListParams) *Iter {
	return getC().List(params)
}

// Returns a list of Configuration objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TerminalConfigurationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TerminalConfigurationList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/terminal/configurations", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for terminal configurations.
type Iter struct {
	*stripe.Iter
}

// TerminalConfiguration returns the terminal configuration which the iterator is currently pointing to.
func (i *Iter) TerminalConfiguration() *stripe.TerminalConfiguration {
	return i.Current().(*stripe.TerminalConfiguration)
}

// TerminalConfigurationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TerminalConfigurationList() *stripe.TerminalConfigurationList {
	return i.List().(*stripe.TerminalConfigurationList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
