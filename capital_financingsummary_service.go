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

// v1CapitalFinancingSummaryService is used to invoke /v1/capital/financing_summary APIs.
type v1CapitalFinancingSummaryService struct {
	B   Backend
	Key string
}

// Retrieve the financing state for the account that was authenticated in the request.
func (c v1CapitalFinancingSummaryService) Retrieve(ctx context.Context, params *CapitalFinancingSummaryRetrieveParams) (*CapitalFinancingSummary, error) {
	financingsummary := &CapitalFinancingSummary{}
	if params == nil {
		params = &CapitalFinancingSummaryRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodGet, "/v1/capital/financing_summary", c.Key, params, financingsummary)
	return financingsummary, err
}
