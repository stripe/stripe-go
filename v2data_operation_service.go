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

// v2DataOperationService is used to invoke operation related APIs.
type v2DataOperationService struct {
	B   Backend
	Key string
}

// Resolve a natural-language query into ranked public analytical resources
// (schemas, reports, and metrics) with execution context.
func (c v2DataOperationService) Search(ctx context.Context, params *V2DataOperationSearchParams) (*V2DataSearchResultList, error) {
	if params == nil {
		params = &V2DataOperationSearchParams{}
	}
	params.Context = ctx
	searchresultlist := &V2DataSearchResultList{}
	err := c.B.Call(
		http.MethodPost, "/v2/data/operations/search", c.Key, params, searchresultlist)
	return searchresultlist, err
}
