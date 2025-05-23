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

// v1FinancialConnectionsSessionService is used to invoke /v1/financial_connections/sessions APIs.
type v1FinancialConnectionsSessionService struct {
	B   Backend
	Key string
}

// To launch the Financial Connections authorization flow, create a Session. The session's client_secret can be used to launch the flow using Stripe.js.
func (c v1FinancialConnectionsSessionService) Create(ctx context.Context, params *FinancialConnectionsSessionCreateParams) (*FinancialConnectionsSession, error) {
	if params == nil {
		params = &FinancialConnectionsSessionCreateParams{}
	}
	params.Context = ctx
	session := &FinancialConnectionsSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/financial_connections/sessions", c.Key, params, session)
	return session, err
}

// Retrieves the details of a Financial Connections Session
func (c v1FinancialConnectionsSessionService) Retrieve(ctx context.Context, id string, params *FinancialConnectionsSessionRetrieveParams) (*FinancialConnectionsSession, error) {
	if params == nil {
		params = &FinancialConnectionsSessionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/financial_connections/sessions/%s", id)
	session := &FinancialConnectionsSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, session)
	return session, err
}
