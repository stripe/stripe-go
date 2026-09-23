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

// v1FinancialConnectionsConsentService is used to invoke /v1/financial_connections/consents APIs.
type v1FinancialConnectionsConsentService struct {
	B   Backend
	Key string
}

// Creates a Financial Connections Consent object for an account holder.
func (c v1FinancialConnectionsConsentService) Create(ctx context.Context, params *FinancialConnectionsConsentCreateParams) (*FinancialConnectionsConsent, error) {
	if params == nil {
		params = &FinancialConnectionsConsentCreateParams{}
	}
	params.Context = ctx
	consent := &FinancialConnectionsConsent{}
	err := c.B.Call(
		http.MethodPost, "/v1/financial_connections/consents", c.Key, params, consent)
	return consent, err
}

// Retrieves the details of a Financial Connections Consent.
func (c v1FinancialConnectionsConsentService) Retrieve(ctx context.Context, id string, params *FinancialConnectionsConsentRetrieveParams) (*FinancialConnectionsConsent, error) {
	if params == nil {
		params = &FinancialConnectionsConsentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/financial_connections/consents/%s", id)
	consent := &FinancialConnectionsConsent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, consent)
	return consent, err
}
