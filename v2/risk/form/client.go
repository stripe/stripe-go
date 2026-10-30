//
//
// File generated from our OpenAPI spec
//
//

// Package form provides the form related APIs
package form

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke form related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a risk Form by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2RiskFormParams) (*stripe.V2RiskForm, error) {
	path := stripe.FormatURLPath("/v2/risk/forms/%s", id)
	form := &stripe.V2RiskForm{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, form)
	return form, err
}

// Submits a response to a risk Form.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2RiskFormParams) (*stripe.V2RiskForm, error) {
	path := stripe.FormatURLPath("/v2/risk/forms/%s", id)
	form := &stripe.V2RiskForm{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, form)
	return form, err
}
