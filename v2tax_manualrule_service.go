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

// v2TaxManualRuleService is used to invoke manualrule related APIs.
type v2TaxManualRuleService struct {
	B   Backend
	Key string
}

// Creates a ManualRule object.
func (c v2TaxManualRuleService) Create(ctx context.Context, params *V2TaxManualRuleCreateParams) (*V2TaxManualRule, error) {
	if params == nil {
		params = &V2TaxManualRuleCreateParams{}
	}
	params.Context = ctx
	manualrule := &V2TaxManualRule{}
	err := c.B.Call(
		http.MethodPost, "/v2/tax/manual_rules", c.Key, params, manualrule)
	return manualrule, err
}

// Retrieves a ManualRule object by ID.
func (c v2TaxManualRuleService) Retrieve(ctx context.Context, id string, params *V2TaxManualRuleRetrieveParams) (*V2TaxManualRule, error) {
	if params == nil {
		params = &V2TaxManualRuleRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/tax/manual_rules/%s", id)
	manualrule := &V2TaxManualRule{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, manualrule)
	return manualrule, err
}

// Updates the Tax configuration for a ManualRule object.
func (c v2TaxManualRuleService) Update(ctx context.Context, id string, params *V2TaxManualRuleUpdateParams) (*V2TaxManualRule, error) {
	if params == nil {
		params = &V2TaxManualRuleUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/tax/manual_rules/%s", id)
	manualrule := &V2TaxManualRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, manualrule)
	return manualrule, err
}

// Deactivates a ManualRule object.
func (c v2TaxManualRuleService) Deactivate(ctx context.Context, id string, params *V2TaxManualRuleDeactivateParams) (*V2TaxManualRule, error) {
	if params == nil {
		params = &V2TaxManualRuleDeactivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/tax/manual_rules/%s/deactivate", id)
	manualrule := &V2TaxManualRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, manualrule)
	return manualrule, err
}

// List all ManualRule objects.
func (c v2TaxManualRuleService) List(ctx context.Context, listParams *V2TaxManualRuleListParams) Seq2[*V2TaxManualRule, error] {
	if listParams == nil {
		listParams = &V2TaxManualRuleListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/tax/manual_rules", listParams, func(path string, p ParamsContainer) (*V2Page[*V2TaxManualRule], error) {
		page := &V2Page[*V2TaxManualRule]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
