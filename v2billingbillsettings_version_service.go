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

// v2BillingBillSettingsVersionService is used to invoke version related APIs.
type v2BillingBillSettingsVersionService struct {
	B   Backend
	Key string
}

// Retrieve a BillSettingVersion by ID.
func (c v2BillingBillSettingsVersionService) Retrieve(ctx context.Context, id string, params *V2BillingBillSettingsVersionRetrieveParams) (*V2BillingBillSettingVersion, error) {
	if params == nil {
		params = &V2BillingBillSettingsVersionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/bill_settings/%s/versions/%s", StringValue(
			params.BillSettingID), id)
	billsettingversion := &V2BillingBillSettingVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, billsettingversion)
	return billsettingversion, err
}

// List all BillSettingVersions by BillSetting ID.
func (c v2BillingBillSettingsVersionService) List(ctx context.Context, listParams *V2BillingBillSettingsVersionListParams) Seq2[*V2BillingBillSettingVersion, error] {
	if listParams == nil {
		listParams = &V2BillingBillSettingsVersionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/bill_settings/%s/versions", StringValue(
			listParams.BillSettingID))
	return NewV2List(path, listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingBillSettingVersion], error) {
		page := &V2Page[*V2BillingBillSettingVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
