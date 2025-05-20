//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"fmt"
	"net/http"
)

// v1SourceService is used to invoke /v1/sources APIs.
type v1SourceService struct {
	B   Backend
	Key string
}

// Creates a new source object.
func (c v1SourceService) Create(ctx context.Context, params *SourceCreateParams) (*Source, error) {
	if params == nil {
		params = &SourceCreateParams{}
	}
	params.Context = ctx
	source := &Source{}
	err := c.B.Call(http.MethodPost, "/v1/sources", c.Key, params, source)
	return source, err
}

// Retrieves an existing source object. Supply the unique source ID from a source creation request and Stripe will return the corresponding up-to-date source object information.
func (c v1SourceService) Retrieve(ctx context.Context, id string, params *SourceRetrieveParams) (*Source, error) {
	if params == nil {
		params = &SourceRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/sources/%s", id)
	source := &Source{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, source)
	return source, err
}

// Updates the specified source by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request accepts the metadata and owner as arguments. It is also possible to update type specific information for selected payment methods. Please refer to our [payment method guides](https://docs.stripe.com/docs/sources) for more detail.
func (c v1SourceService) Update(ctx context.Context, id string, params *SourceUpdateParams) (*Source, error) {
	if params == nil {
		params = &SourceUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/sources/%s", id)
	source := &Source{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, source)
	return source, err
}

// Delete a specified source for a given customer.
func (c v1SourceService) Detach(ctx context.Context, id string, params *SourceDetachParams) (*Source, error) {
	if params.Customer == nil {
		return nil, fmt.Errorf(
			"Invalid source detach params: Customer needs to be set")
	}
	if params == nil {
		params = &SourceDetachParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	source := &Source{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, source)
	return source, err
}
