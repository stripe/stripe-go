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

// v2BillingContractsPricingLinesQuantityChangeService is used to invoke quantitychange related APIs.
type v2BillingContractsPricingLinesQuantityChangeService struct {
	B   Backend
	Key string
}

// List quantity changes for a pricing line on a contract.
func (c v2BillingContractsPricingLinesQuantityChangeService) ListContractPricingLineQuantityChanges(ctx context.Context, listParams *V2BillingContractsPricingLinesQuantityChangeListContractPricingLineQuantityChangesParams) *V2List[*V2BillingContractPricingLineQuantityChange] {
	if listParams == nil {
		listParams = &V2BillingContractsPricingLinesQuantityChangeListContractPricingLineQuantityChangesParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/contracts/%s/pricing_lines/%s/quantity_changes", StringValue(
			listParams.ContractID), StringValue(listParams.PricingLineID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingContractPricingLineQuantityChange], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingContractPricingLineQuantityChange]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
