//
//
// File generated from our OpenAPI spec
//
//

// Package authorization provides the /v1/financial_connections/authorizations APIs
package authorization

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/financial_connections/authorizations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an Financial Connections Authorization.
func Get(id string, params *stripe.FinancialConnectionsAuthorizationParams) (*stripe.FinancialConnectionsAuthorization, error) {
	return getC().Get(id, params)
}

// Retrieves the details of an Financial Connections Authorization.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.FinancialConnectionsAuthorizationParams) (*stripe.FinancialConnectionsAuthorization, error) {
	path := stripe.FormatURLPath(
		"/v1/financial_connections/authorizations/%s", id)
	authorization := &stripe.FinancialConnectionsAuthorization{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, authorization)
	return authorization, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
