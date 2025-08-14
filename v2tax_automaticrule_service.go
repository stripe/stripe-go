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

// v2TaxAutomaticRuleService is used to invoke automaticrule related APIs.
type v2TaxAutomaticRuleService struct {
	B   Backend
	Key string
}

// Creates an AutomaticRule object.
func (c v2TaxAutomaticRuleService) Create(ctx context.Context, params *V2TaxAutomaticRuleCreateParams) (*V2TaxAutomaticRule, error) {
	if params == nil {
		params = &V2TaxAutomaticRuleCreateParams{}
	}
	params.Context = ctx
	automaticrule := &V2TaxAutomaticRule{}
	err := c.B.Call(
		http.MethodPost, "/v2/tax/automatic_rules", c.Key, params, automaticrule)
	return automaticrule, err
}

// Retrieves an AutomaticRule object by ID.
func (c v2TaxAutomaticRuleService) Retrieve(ctx context.Context, id string, params *V2TaxAutomaticRuleRetrieveParams) (*V2TaxAutomaticRule, error) {
	if params == nil {
		params = &V2TaxAutomaticRuleRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/tax/automatic_rules/%s", id)
	automaticrule := &V2TaxAutomaticRule{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, automaticrule)
	return automaticrule, err
}

// Updates the automatic Tax configuration for an AutomaticRule object.
func (c v2TaxAutomaticRuleService) Update(ctx context.Context, id string, params *V2TaxAutomaticRuleUpdateParams) (*V2TaxAutomaticRule, error) {
	if params == nil {
		params = &V2TaxAutomaticRuleUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/tax/automatic_rules/%s", id)
	automaticrule := &V2TaxAutomaticRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, automaticrule)
	return automaticrule, err
}

// Deactivates an AutomaticRule object. Deactivated AutomaticRule objects are ignored in future tax calculations.
func (c v2TaxAutomaticRuleService) Deactivate(ctx context.Context, id string, params *V2TaxAutomaticRuleDeactivateParams) (*V2TaxAutomaticRule, error) {
	if params == nil {
		params = &V2TaxAutomaticRuleDeactivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/tax/automatic_rules/%s/deactivate", id)
	automaticrule := &V2TaxAutomaticRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, automaticrule)
	return automaticrule, err
}

// Finds an AutomaticRule object by BillableItem ID.
func (c v2TaxAutomaticRuleService) Find(ctx context.Context, params *V2TaxAutomaticRuleFindParams) (*V2TaxAutomaticRule, error) {
	if params == nil {
		params = &V2TaxAutomaticRuleFindParams{}
	}
	params.Context = ctx
	automaticrule := &V2TaxAutomaticRule{}
	err := c.B.Call(
		http.MethodGet, "/v2/tax/automatic_rules/find", c.Key, params, automaticrule)
	return automaticrule, err
}
