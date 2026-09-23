//
//
// File generated from our OpenAPI spec
//
//

// Package consent provides the /v1/financial_connections/consents APIs
package consent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke /v1/financial_connections/consents APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a Financial Connections Consent object for an account holder.
func New(params *stripe.FinancialConnectionsConsentParams) (*stripe.FinancialConnectionsConsent, error) {
	return getC().New(params)
}

// Creates a Financial Connections Consent object for an account holder.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.FinancialConnectionsConsentParams) (*stripe.FinancialConnectionsConsent, error) {
	consent := &stripe.FinancialConnectionsConsent{}
	err := c.B.Call(
		http.MethodPost, "/v1/financial_connections/consents", c.Key, params, consent)
	return consent, err
}

// Retrieves the details of a Financial Connections Consent.
func Get(id string, params *stripe.FinancialConnectionsConsentParams) (*stripe.FinancialConnectionsConsent, error) {
	return getC().Get(id, params)
}

// Retrieves the details of a Financial Connections Consent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.FinancialConnectionsConsentParams) (*stripe.FinancialConnectionsConsent, error) {
	path := stripe.FormatURLPath("/v1/financial_connections/consents/%s", id)
	consent := &stripe.FinancialConnectionsConsent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, consent)
	return consent, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
