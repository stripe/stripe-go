//
//
// File generated from our OpenAPI spec
//
//

// Package configuration provides the /terminal/configurations APIs
package configuration

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /terminal/configurations APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new terminal configuration.
func New(params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	return getC().New(params)
}

// New creates a new terminal configuration.
func (c Client) New(params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	configuration := &stripe.TerminalConfiguration{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/terminal/configurations",
		c.Key,
		params,
		configuration,
	)
	return configuration, err
}

// Get returns the details of a terminal configuration.
func Get(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	return getC().Get(id, params)
}

// Get returns the details of a terminal configuration.
func (c Client) Get(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	path := stripe.FormatURLPath("/v1/terminal/configurations/%s", id)
	configuration := &stripe.TerminalConfiguration{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, configuration)
	return configuration, err
}

// Update updates a terminal configuration's properties.
func Update(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	return getC().Update(id, params)
}

// Update updates a terminal configuration's properties.
func (c Client) Update(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	path := stripe.FormatURLPath("/v1/terminal/configurations/%s", id)
	configuration := &stripe.TerminalConfiguration{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, configuration)
	return configuration, err
}

// Del removes a terminal configuration.
func Del(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	return getC().Del(id, params)
}

// Del removes a terminal configuration.
func (c Client) Del(id string, params *stripe.TerminalConfigurationParams) (*stripe.TerminalConfiguration, error) {
	path := stripe.FormatURLPath("/v1/terminal/configurations/%s", id)
	configuration := &stripe.TerminalConfiguration{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, configuration)
	return configuration, err
}

// List returns a list of terminal configurations.
func List(params *stripe.TerminalConfigurationListParams) *Iter {
	return getC().List(params)
}

// List returns a list of terminal configurations.
func (c Client) List(listParams *stripe.TerminalConfigurationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TerminalConfigurationList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/terminal/configurations", c.Key, b, p, list)

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
