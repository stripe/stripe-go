//
//
// File generated from our OpenAPI spec
//
//

// Package source provides the /sources APIs
package source

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /sources APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new source.
func New(params *stripe.SourceObjectParams) (*stripe.Source, error) {
	return getC().New(params)
}

// New creates a new source.
func (c Client) New(params *stripe.SourceObjectParams) (*stripe.Source, error) {
	source := &stripe.Source{}
	err := c.B.Call(http.MethodPost, "/v1/sources", c.Key, params, source)
	return source, err
}

// Get returns the details of a source.
func Get(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	return getC().Get(id, params)
}

// Get returns the details of a source.
func (c Client) Get(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	path := stripe.FormatURLPath("/v1/sources/%s", id)
	source := &stripe.Source{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, source)
	return source, err
}

// Update updates a source's properties.
func Update(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	return getC().Update(id, params)
}

// Update updates a source's properties.
func (c Client) Update(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	path := stripe.FormatURLPath("/v1/sources/%s", id)
	source := &stripe.Source{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, source)
	return source, err
}

// Detach is the method for the `DELETE /v1/customers/{customer}/sources/{id}` API.
func Detach(id string, params *stripe.SourceObjectDetachParams) (*stripe.Source, error) {
	return getC().Detach(id, params)
}

// Detach is the method for the `DELETE /v1/customers/{customer}/sources/{id}` API.
func (c Client) Detach(id string, params *stripe.SourceObjectDetachParams) (*stripe.Source, error) {
	if params.Customer == nil {
		return nil, fmt.Errorf(
			"Invalid source detach params: Customer needs to be set",
		)
	}
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	source := &stripe.Source{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, source)
	return source, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
