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

// v2TaxOperationService is used to invoke operation related APIs.
type v2TaxOperationService struct {
	B   Backend
	Key string
}

// Resolves an address to its tax precision level.
func (c v2TaxOperationService) ResolveAddress(ctx context.Context, params *V2TaxOperationResolveAddressParams) (*V2TaxOperationsResolveAddressResult, error) {
	if params == nil {
		params = &V2TaxOperationResolveAddressParams{}
	}
	params.Context = ctx
	operationsresolveaddressresult := &V2TaxOperationsResolveAddressResult{}
	err := c.B.Call(
		http.MethodPost, "/v2/tax/operations/resolve_address", c.Key, params, operationsresolveaddressresult)
	return operationsresolveaddressresult, err
}
