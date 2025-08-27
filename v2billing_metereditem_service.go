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

// v2BillingMeteredItemService is used to invoke metereditem related APIs.
type v2BillingMeteredItemService struct {
	B   Backend
	Key string
}

// Create a Metered Item object.
func (c v2BillingMeteredItemService) Create(ctx context.Context, params *V2BillingMeteredItemCreateParams) (*V2BillingMeteredItem, error) {
	if params == nil {
		params = &V2BillingMeteredItemCreateParams{}
	}
	params.Context = ctx
	metereditem := &V2BillingMeteredItem{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/metered_items", c.Key, params, metereditem)
	return metereditem, err
}

// Retrieve a Metered Item object.
func (c v2BillingMeteredItemService) Retrieve(ctx context.Context, id string, params *V2BillingMeteredItemRetrieveParams) (*V2BillingMeteredItem, error) {
	if params == nil {
		params = &V2BillingMeteredItemRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/metered_items/%s", id)
	metereditem := &V2BillingMeteredItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, metereditem)
	return metereditem, err
}

// Update a Metered Item object. At least one of the fields is required.
func (c v2BillingMeteredItemService) Update(ctx context.Context, id string, params *V2BillingMeteredItemUpdateParams) (*V2BillingMeteredItem, error) {
	if params == nil {
		params = &V2BillingMeteredItemUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/metered_items/%s", id)
	metereditem := &V2BillingMeteredItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, metereditem)
	return metereditem, err
}

// List all Metered Item objects in reverse chronological order of creation.
func (c v2BillingMeteredItemService) List(ctx context.Context, listParams *V2BillingMeteredItemListParams) Seq2[*V2BillingMeteredItem, error] {
	if listParams == nil {
		listParams = &V2BillingMeteredItemListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/metered_items", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingMeteredItem], error) {
		page := &V2Page[*V2BillingMeteredItem]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
