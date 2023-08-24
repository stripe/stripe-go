//
//
// File generated from our OpenAPI spec
//
//

// Package capability provides the /accounts/{account}/capabilities APIs
package capability

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/form"
)

// Client is used to invoke /accounts/{account}/capabilities APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a capability.
func Get(id string, params *stripe.CapabilityParams) (*stripe.Capability, error) {
	return getC().Get(id, params)
}

// Get returns the details of a capability.
func (c Client) Get(id string, params *stripe.CapabilityParams) (*stripe.Capability, error) {
	if params == nil {
		return nil, fmt.Errorf(
			"params cannot be nil, and params.Account must be set",
		)
	}
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/capabilities/%s",
		stripe.StringValue(params.Account),
		id,
	)
	capability := &stripe.Capability{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, capability)
	return capability, err
}

// Update updates a capability's properties.
func Update(id string, params *stripe.CapabilityParams) (*stripe.Capability, error) {
	return getC().Update(id, params)
}

// Update updates a capability's properties.
func (c Client) Update(id string, params *stripe.CapabilityParams) (*stripe.Capability, error) {
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/capabilities/%s",
		stripe.StringValue(params.Account),
		id,
	)
	capability := &stripe.Capability{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, capability)
	return capability, err
}

// List returns a list of capabilities.
func List(params *stripe.CapabilityListParams) *Iter {
	return getC().List(params)
}

// List returns a list of capabilities.
func (c Client) List(listParams *stripe.CapabilityListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/capabilities",
		stripe.StringValue(listParams.Account),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CapabilityList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for capabilities.
type Iter struct {
	*stripe.Iter
}

// Capability returns the capability which the iterator is currently pointing to.
func (i *Iter) Capability() *stripe.Capability {
	return i.Current().(*stripe.Capability)
}

// CapabilityList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CapabilityList() *stripe.CapabilityList {
	return i.List().(*stripe.CapabilityList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
