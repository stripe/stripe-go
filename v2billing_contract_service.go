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

// v2BillingContractService is used to invoke contract related APIs.
type v2BillingContractService struct {
	B   Backend
	Key string
}

// Create a Contract object.
func (c v2BillingContractService) Create(ctx context.Context, params *V2BillingContractCreateParams) (*V2BillingContract, error) {
	if params == nil {
		params = &V2BillingContractCreateParams{}
	}
	params.Context = ctx
	contract := &V2BillingContract{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/contracts", c.Key, params, contract)
	return contract, err
}

// Retrieve a Contract object by ID.
func (c v2BillingContractService) Retrieve(ctx context.Context, id string, params *V2BillingContractRetrieveParams) (*V2BillingContract, error) {
	if params == nil {
		params = &V2BillingContractRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/contracts/%s", id)
	contract := &V2BillingContract{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, contract)
	return contract, err
}

// Update a Contract object by ID.
func (c v2BillingContractService) Update(ctx context.Context, id string, params *V2BillingContractUpdateParams) (*V2BillingContract, error) {
	if params == nil {
		params = &V2BillingContractUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/contracts/%s", id)
	contract := &V2BillingContract{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, contract)
	return contract, err
}

// Delete a draft Contract object by ID.
func (c v2BillingContractService) Delete(ctx context.Context, id string, params *V2BillingContractDeleteParams) (*V2DeletedObject, error) {
	if params == nil {
		params = &V2BillingContractDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/contracts/%s", id)
	deletedObj := &V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}

// Activate a Draft Contract object by ID.
func (c v2BillingContractService) Activate(ctx context.Context, id string, params *V2BillingContractActivateParams) (*V2BillingContract, error) {
	if params == nil {
		params = &V2BillingContractActivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/contracts/%s/activate", id)
	contract := &V2BillingContract{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, contract)
	return contract, err
}

// Cancel a Contract object by ID.
func (c v2BillingContractService) Cancel(ctx context.Context, id string, params *V2BillingContractCancelParams) (*V2BillingContract, error) {
	if params == nil {
		params = &V2BillingContractCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/contracts/%s/cancel", id)
	contract := &V2BillingContract{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, contract)
	return contract, err
}

// List Contract objects with pagination.
func (c v2BillingContractService) List(ctx context.Context, listParams *V2BillingContractListParams) *V2List[*V2BillingContract] {
	if listParams == nil {
		listParams = &V2BillingContractListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/billing/contracts", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingContract], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingContract]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
