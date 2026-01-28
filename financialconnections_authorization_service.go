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

// v1FinancialConnectionsAuthorizationService is used to invoke /v1/financial_connections/authorizations APIs.
type v1FinancialConnectionsAuthorizationService struct {
	B   Backend
	Key string
}

// Retrieves the details of an Financial Connections Authorization.
func (c v1FinancialConnectionsAuthorizationService) Retrieve(ctx context.Context, id string, params *FinancialConnectionsAuthorizationRetrieveParams) (*FinancialConnectionsAuthorization, error) {
	if params == nil {
		params = &FinancialConnectionsAuthorizationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/financial_connections/authorizations/%s", id)
	authorization := &FinancialConnectionsAuthorization{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, authorization)
	return authorization, err
}
