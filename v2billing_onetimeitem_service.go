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

// v2BillingOneTimeItemService is used to invoke onetimeitem related APIs.
type v2BillingOneTimeItemService struct {
	B   Backend
	Key string
}

// Create a One-Time Item object.
func (c v2BillingOneTimeItemService) Create(ctx context.Context, params *V2BillingOneTimeItemCreateParams) (*V2BillingOneTimeItem, error) {
	if params == nil {
		params = &V2BillingOneTimeItemCreateParams{}
	}
	params.Context = ctx
	onetimeitem := &V2BillingOneTimeItem{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/one_time_items", c.Key, params, onetimeitem)
	return onetimeitem, err
}

// Retrieve a One-Time Item object.
func (c v2BillingOneTimeItemService) Retrieve(ctx context.Context, id string, params *V2BillingOneTimeItemRetrieveParams) (*V2BillingOneTimeItem, error) {
	if params == nil {
		params = &V2BillingOneTimeItemRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/one_time_items/%s", id)
	onetimeitem := &V2BillingOneTimeItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, onetimeitem)
	return onetimeitem, err
}

// Update a One-Time Item object.
func (c v2BillingOneTimeItemService) Update(ctx context.Context, id string, params *V2BillingOneTimeItemUpdateParams) (*V2BillingOneTimeItem, error) {
	if params == nil {
		params = &V2BillingOneTimeItemUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/one_time_items/%s", id)
	onetimeitem := &V2BillingOneTimeItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, onetimeitem)
	return onetimeitem, err
}

// List all One-Time Item objects in reverse chronological order of creation.
func (c v2BillingOneTimeItemService) List(ctx context.Context, listParams *V2BillingOneTimeItemListParams) *V2List[*V2BillingOneTimeItem] {
	if listParams == nil {
		listParams = &V2BillingOneTimeItemListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/billing/one_time_items", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingOneTimeItem], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingOneTimeItem]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
