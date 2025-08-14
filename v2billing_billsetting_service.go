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

// v2BillingBillSettingService is used to invoke billsetting related APIs.
type v2BillingBillSettingService struct {
	B   Backend
	Key string
}

// Create a BillSetting object.
func (c v2BillingBillSettingService) Create(ctx context.Context, params *V2BillingBillSettingCreateParams) (*V2BillingBillSetting, error) {
	if params == nil {
		params = &V2BillingBillSettingCreateParams{}
	}
	params.Context = ctx
	billsetting := &V2BillingBillSetting{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/bill_settings", c.Key, params, billsetting)
	return billsetting, err
}

// Retrieve a BillSetting object by ID.
func (c v2BillingBillSettingService) Retrieve(ctx context.Context, id string, params *V2BillingBillSettingRetrieveParams) (*V2BillingBillSetting, error) {
	if params == nil {
		params = &V2BillingBillSettingRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/bill_settings/%s", id)
	billsetting := &V2BillingBillSetting{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, billsetting)
	return billsetting, err
}

// Update fields on an existing BillSetting object.
func (c v2BillingBillSettingService) Update(ctx context.Context, id string, params *V2BillingBillSettingUpdateParams) (*V2BillingBillSetting, error) {
	if params == nil {
		params = &V2BillingBillSettingUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/bill_settings/%s", id)
	billsetting := &V2BillingBillSetting{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, billsetting)
	return billsetting, err
}

// List all BillSetting objects.
func (c v2BillingBillSettingService) List(ctx context.Context, listParams *V2BillingBillSettingListParams) Seq2[*V2BillingBillSetting, error] {
	if listParams == nil {
		listParams = &V2BillingBillSettingListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/bill_settings", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingBillSetting], error) {
		page := &V2Page[*V2BillingBillSetting]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
