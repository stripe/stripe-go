//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1TerminalLocationService is used to invoke /v1/terminal/locations APIs.
type v1TerminalLocationService struct {
	B   Backend
	Key string
}

// Creates a new Location object.
// For further details, including which address fields are required in each country, see the [Manage locations](https://docs.stripe.com/docs/terminal/fleet/locations) guide.
func (c v1TerminalLocationService) Create(ctx context.Context, params *TerminalLocationCreateParams) (*TerminalLocation, error) {
	if params == nil {
		params = &TerminalLocationCreateParams{}
	}
	params.Context = ctx
	location := &TerminalLocation{}
	err := c.B.Call(
		http.MethodPost, "/v1/terminal/locations", c.Key, params, location)
	return location, err
}

// Retrieves a Location object.
func (c v1TerminalLocationService) Retrieve(ctx context.Context, id string, params *TerminalLocationRetrieveParams) (*TerminalLocation, error) {
	if params == nil {
		params = &TerminalLocationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/locations/%s", id)
	location := &TerminalLocation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, location)
	return location, err
}

// Updates a Location object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1TerminalLocationService) Update(ctx context.Context, id string, params *TerminalLocationUpdateParams) (*TerminalLocation, error) {
	if params == nil {
		params = &TerminalLocationUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/locations/%s", id)
	location := &TerminalLocation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, location)
	return location, err
}

// Deletes a Location object.
func (c v1TerminalLocationService) Delete(ctx context.Context, id string, params *TerminalLocationDeleteParams) (*TerminalLocation, error) {
	if params == nil {
		params = &TerminalLocationDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/locations/%s", id)
	location := &TerminalLocation{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, location)
	return location, err
}

// Returns a list of Location objects.
func (c v1TerminalLocationService) List(ctx context.Context, listParams *TerminalLocationListParams) Seq2[*TerminalLocation, error] {
	if listParams == nil {
		listParams = &TerminalLocationListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TerminalLocation, ListContainer, error) {
		list := &TerminalLocationList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/terminal/locations", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
