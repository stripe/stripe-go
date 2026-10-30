//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2RiskFormService is used to invoke form related APIs.
type v2RiskFormService struct {
	B   Backend
	Key string
}

// Retrieves a risk Form by ID.
func (c v2RiskFormService) Retrieve(ctx context.Context, id string, params *V2RiskFormRetrieveParams) (*V2RiskForm, error) {
	if params == nil {
		params = &V2RiskFormRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/risk/forms/%s", id)
	form := &V2RiskForm{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, form)
	return form, err
}

// Submits a response to a risk Form.
func (c v2RiskFormService) Update(ctx context.Context, id string, params *V2RiskFormUpdateParams) (*V2RiskForm, error) {
	if params == nil {
		params = &V2RiskFormUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/risk/forms/%s", id)
	form := &V2RiskForm{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, form)
	return form, err
}
