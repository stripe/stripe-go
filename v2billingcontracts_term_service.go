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

// v2BillingContractsTermService is used to invoke term related APIs.
type v2BillingContractsTermService struct {
	B   Backend
	Key string
}

// Retrieve a term on a contract.
func (c v2BillingContractsTermService) RetrieveContractTerm(ctx context.Context, id string, params *V2BillingContractsTermRetrieveContractTermParams) (*V2BillingContractTerm, error) {
	if params == nil {
		params = &V2BillingContractsTermRetrieveContractTermParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/contracts/%s/terms/%s", StringValue(params.ContractID), id)
	contractterm := &V2BillingContractTerm{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, contractterm)
	return contractterm, err
}

// List the terms on a contract, optionally filtered by type.
func (c v2BillingContractsTermService) ListContractTerms(ctx context.Context, listParams *V2BillingContractsTermListContractTermsParams) *V2List[*V2BillingContractTerm] {
	if listParams == nil {
		listParams = &V2BillingContractsTermListContractTermsParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/contracts/%s/terms", StringValue(listParams.ContractID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingContractTerm], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingContractTerm]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
