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

// v2BillingGroupService is used to invoke group related APIs.
type v2BillingGroupService struct {
	B   Backend
	Key string
}

// Create a BillingGroup.
func (c v2BillingGroupService) Create(ctx context.Context, params *V2BillingGroupCreateParams) (*V2BillingGroup, error) {
	if params == nil {
		params = &V2BillingGroupCreateParams{}
	}
	params.Context = ctx
	group := &V2BillingGroup{}
	err := c.B.Call(http.MethodPost, "/v2/billing/groups", c.Key, params, group)
	return group, err
}

// Retrieve a BillingGroup.
func (c v2BillingGroupService) Retrieve(ctx context.Context, id string, params *V2BillingGroupRetrieveParams) (*V2BillingGroup, error) {
	if params == nil {
		params = &V2BillingGroupRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/groups/%s", id)
	group := &V2BillingGroup{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, group)
	return group, err
}

// Update a BillingGroup.
func (c v2BillingGroupService) Update(ctx context.Context, id string, params *V2BillingGroupUpdateParams) (*V2BillingGroup, error) {
	if params == nil {
		params = &V2BillingGroupUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/groups/%s", id)
	group := &V2BillingGroup{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, group)
	return group, err
}

// Cancel a BillingGroup.
func (c v2BillingGroupService) Cancel(ctx context.Context, id string, params *V2BillingGroupCancelParams) (*V2BillingGroup, error) {
	if params == nil {
		params = &V2BillingGroupCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/groups/%s/cancel", id)
	group := &V2BillingGroup{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, group)
	return group, err
}

// List Billing Groups.
func (c v2BillingGroupService) List(ctx context.Context, listParams *V2BillingGroupListParams) *V2List[*V2BillingGroup] {
	if listParams == nil {
		listParams = &V2BillingGroupListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/billing/groups", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingGroup], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingGroup]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
