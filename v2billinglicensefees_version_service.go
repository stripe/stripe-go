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

// v2BillingLicenseFeesVersionService is used to invoke version related APIs.
type v2BillingLicenseFeesVersionService struct {
	B   Backend
	Key string
}

// Retrieve a License Fee Version object.
func (c v2BillingLicenseFeesVersionService) Retrieve(ctx context.Context, id string, params *V2BillingLicenseFeesVersionRetrieveParams) (*V2BillingLicenseFeeVersion, error) {
	if params == nil {
		params = &V2BillingLicenseFeesVersionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/license_fees/%s/versions/%s", StringValue(
			params.LicenseFeeID), id)
	licensefeeversion := &V2BillingLicenseFeeVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, licensefeeversion)
	return licensefeeversion, err
}

// List all versions of a License Fee object.
func (c v2BillingLicenseFeesVersionService) List(ctx context.Context, listParams *V2BillingLicenseFeesVersionListParams) *V2List[*V2BillingLicenseFeeVersion] {
	if listParams == nil {
		listParams = &V2BillingLicenseFeesVersionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/license_fees/%s/versions", StringValue(
			listParams.LicenseFeeID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingLicenseFeeVersion], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingLicenseFeeVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
