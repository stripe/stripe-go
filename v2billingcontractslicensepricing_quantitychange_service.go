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

// v2BillingContractsLicensePricingQuantityChangeService is used to invoke quantitychange related APIs.
type v2BillingContractsLicensePricingQuantityChangeService struct {
	B   Backend
	Key string
}

// List license quantity changes for a contract given a license pricing ID.
func (c v2BillingContractsLicensePricingQuantityChangeService) ListQuantityChanges(ctx context.Context, listParams *V2BillingContractsLicensePricingQuantityChangeListQuantityChangesParams) *V2List[*V2BillingContractLicensePricingQuantityChange] {
	if listParams == nil {
		listParams = &V2BillingContractsLicensePricingQuantityChangeListQuantityChangesParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/contracts/%s/license_pricing/%s/quantity_changes", StringValue(
			listParams.ContractID), StringValue(listParams.LicensePricingID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingContractLicensePricingQuantityChange], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingContractLicensePricingQuantityChange]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
