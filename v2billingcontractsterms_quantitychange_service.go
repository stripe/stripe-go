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

// v2BillingContractsTermsQuantityChangeService is used to invoke quantitychange related APIs.
type v2BillingContractsTermsQuantityChangeService struct {
	B   Backend
	Key string
}

// List quantity changes for a licensed recurring-price term.
func (c v2BillingContractsTermsQuantityChangeService) ListContractTermQuantityChanges(ctx context.Context, listParams *V2BillingContractsTermsQuantityChangeListContractTermQuantityChangesParams) *V2List[*V2BillingContractTermQuantityChange] {
	if listParams == nil {
		listParams = &V2BillingContractsTermsQuantityChangeListContractTermQuantityChangesParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/contracts/%s/terms/%s/quantity_changes", StringValue(
			listParams.ContractID), StringValue(listParams.Term))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingContractTermQuantityChange], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingContractTermQuantityChange]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
