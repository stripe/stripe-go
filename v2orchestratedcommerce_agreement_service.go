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

// v2OrchestratedCommerceAgreementService is used to invoke agreement related APIs.
type v2OrchestratedCommerceAgreementService struct {
	B   Backend
	Key string
}

// Create a new Agreement.
func (c v2OrchestratedCommerceAgreementService) Create(ctx context.Context, params *V2OrchestratedCommerceAgreementCreateParams) (*V2OrchestratedCommerceAgreement, error) {
	if params == nil {
		params = &V2OrchestratedCommerceAgreementCreateParams{}
	}
	params.Context = ctx
	agreement := &V2OrchestratedCommerceAgreement{}
	err := c.B.Call(
		http.MethodPost, "/v2/orchestrated_commerce/agreements", c.Key, params, agreement)
	return agreement, err
}

// Retrieve an Agreement by ID.
func (c v2OrchestratedCommerceAgreementService) Retrieve(ctx context.Context, id string, params *V2OrchestratedCommerceAgreementRetrieveParams) (*V2OrchestratedCommerceAgreement, error) {
	if params == nil {
		params = &V2OrchestratedCommerceAgreementRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/orchestrated_commerce/agreements/%s", id)
	agreement := &V2OrchestratedCommerceAgreement{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, agreement)
	return agreement, err
}

// Confirm an Agreement.
func (c v2OrchestratedCommerceAgreementService) Confirm(ctx context.Context, id string, params *V2OrchestratedCommerceAgreementConfirmParams) (*V2OrchestratedCommerceAgreement, error) {
	if params == nil {
		params = &V2OrchestratedCommerceAgreementConfirmParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/orchestrated_commerce/agreements/%s/confirm", id)
	agreement := &V2OrchestratedCommerceAgreement{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, agreement)
	return agreement, err
}

// Terminate an Agreement.
func (c v2OrchestratedCommerceAgreementService) Terminate(ctx context.Context, id string, params *V2OrchestratedCommerceAgreementTerminateParams) (*V2OrchestratedCommerceAgreement, error) {
	if params == nil {
		params = &V2OrchestratedCommerceAgreementTerminateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/orchestrated_commerce/agreements/%s/terminate", id)
	agreement := &V2OrchestratedCommerceAgreement{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, agreement)
	return agreement, err
}

// List Agreements for the profile associated with the authenticated merchant.
func (c v2OrchestratedCommerceAgreementService) List(ctx context.Context, listParams *V2OrchestratedCommerceAgreementListParams) *V2List[*V2OrchestratedCommerceAgreement] {
	if listParams == nil {
		listParams = &V2OrchestratedCommerceAgreementListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/orchestrated_commerce/agreements", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2OrchestratedCommerceAgreement], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2OrchestratedCommerceAgreement]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
