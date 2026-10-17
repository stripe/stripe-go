//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v87/form"
)

// v1RiskFormService is used to invoke /v1/risk/forms APIs.
type v1RiskFormService struct {
	B   Backend
	Key string
}

// Retrieves a risk Form by ID.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
func (c v1RiskFormService) Retrieve(ctx context.Context, id string, params *RiskFormRetrieveParams) (*RiskForm, error) {
	if params == nil {
		params = &RiskFormRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/risk/forms/%s", id)
	form := &RiskForm{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, form)
	return form, err
}

// Updates a risk Form with the connected account's response.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
// | form_closed | 400 | The Form is closed and cannot accept updates. |
// | invalid_field_for_form_type | 400 | A submitted field is not supported by this Form type. |
// | form_field_already_submitted | 400 | A submitted field has already been submitted. |
// | form_insufficient_fields | 400 | Required fields are missing for this Form type. |
// | resource_missing | 404 | A submitted file does not exist or is not accessible to the requesting account. |
func (c v1RiskFormService) Update(ctx context.Context, id string, params *RiskFormUpdateParams) (*RiskForm, error) {
	if params == nil {
		params = &RiskFormUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/risk/forms/%s", id)
	form := &RiskForm{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, form)
	return form, err
}
