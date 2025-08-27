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

// v2BillingLicenseFeeService is used to invoke licensefee related APIs.
type v2BillingLicenseFeeService struct {
	B   Backend
	Key string
}

// Create a License Fee object.
func (c v2BillingLicenseFeeService) Create(ctx context.Context, params *V2BillingLicenseFeeCreateParams) (*V2BillingLicenseFee, error) {
	if params == nil {
		params = &V2BillingLicenseFeeCreateParams{}
	}
	params.Context = ctx
	licensefee := &V2BillingLicenseFee{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/license_fees", c.Key, params, licensefee)
	return licensefee, err
}

// Retrieve a License Fee object.
func (c v2BillingLicenseFeeService) Retrieve(ctx context.Context, id string, params *V2BillingLicenseFeeRetrieveParams) (*V2BillingLicenseFee, error) {
	if params == nil {
		params = &V2BillingLicenseFeeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/license_fees/%s", id)
	licensefee := &V2BillingLicenseFee{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, licensefee)
	return licensefee, err
}

// Update a License Fee object.
func (c v2BillingLicenseFeeService) Update(ctx context.Context, id string, params *V2BillingLicenseFeeUpdateParams) (*V2BillingLicenseFee, error) {
	if params == nil {
		params = &V2BillingLicenseFeeUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/license_fees/%s", id)
	licensefee := &V2BillingLicenseFee{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, licensefee)
	return licensefee, err
}

// List all License Fee objects.
func (c v2BillingLicenseFeeService) List(ctx context.Context, listParams *V2BillingLicenseFeeListParams) Seq2[*V2BillingLicenseFee, error] {
	if listParams == nil {
		listParams = &V2BillingLicenseFeeListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/license_fees", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingLicenseFee], error) {
		page := &V2Page[*V2BillingLicenseFee]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
