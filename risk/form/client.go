//
//
// File generated from our OpenAPI spec
//
//

// Package form provides the /v1/risk/forms APIs
package form

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
	"github.com/stripe/stripe-go/v87/form"
)

// Client is used to invoke /v1/risk/forms APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a risk Form by ID.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
func Get(id string, params *stripe.RiskFormParams) (*stripe.RiskForm, error) {
	return getC().Get(id, params)
}

// Retrieves a risk Form by ID.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.RiskFormParams) (*stripe.RiskForm, error) {
	path := stripe.FormatURLPath("/v1/risk/forms/%s", id)
	form := &stripe.RiskForm{}
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
func Update(id string, params *stripe.RiskFormParams) (*stripe.RiskForm, error) {
	return getC().Update(id, params)
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
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.RiskFormParams) (*stripe.RiskForm, error) {
	path := stripe.FormatURLPath("/v1/risk/forms/%s", id)
	form := &stripe.RiskForm{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, form)
	return form, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
