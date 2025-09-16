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

// v1TerminalConfigurationService is used to invoke /v1/terminal/configurations APIs.
type v1TerminalConfigurationService struct {
	B   Backend
	Key string
}

// Creates a new Configuration object.
func (c v1TerminalConfigurationService) Create(ctx context.Context, params *TerminalConfigurationCreateParams) (*TerminalConfiguration, error) {
	if params == nil {
		params = &TerminalConfigurationCreateParams{}
	}
	params.Context = ctx
	configuration := &TerminalConfiguration{}
	err := c.B.Call(
		http.MethodPost, "/v1/terminal/configurations", c.Key, params, configuration)
	return configuration, err
}

// Retrieves a Configuration object.
func (c v1TerminalConfigurationService) Retrieve(ctx context.Context, id string, params *TerminalConfigurationRetrieveParams) (*TerminalConfiguration, error) {
	if params == nil {
		params = &TerminalConfigurationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/configurations/%s", id)
	configuration := &TerminalConfiguration{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, configuration)
	return configuration, err
}

// Updates a new Configuration object.
func (c v1TerminalConfigurationService) Update(ctx context.Context, id string, params *TerminalConfigurationUpdateParams) (*TerminalConfiguration, error) {
	if params == nil {
		params = &TerminalConfigurationUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/configurations/%s", id)
	configuration := &TerminalConfiguration{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, configuration)
	return configuration, err
}

// Deletes a Configuration object.
func (c v1TerminalConfigurationService) Delete(ctx context.Context, id string, params *TerminalConfigurationDeleteParams) (*TerminalConfiguration, error) {
	if params == nil {
		params = &TerminalConfigurationDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/terminal/configurations/%s", id)
	configuration := &TerminalConfiguration{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, configuration)
	return configuration, err
}

// Returns a list of Configuration objects.
func (c v1TerminalConfigurationService) List(ctx context.Context, listParams *TerminalConfigurationListParams) Seq2[*TerminalConfiguration, error] {
	if listParams == nil {
		listParams = &TerminalConfigurationListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TerminalConfiguration, ListContainer, error) {
		list := &TerminalConfigurationList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/terminal/configurations", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
