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

// v2BillingGroupsSpendModifierRuleService is used to invoke spendmodifierrule related APIs.
type v2BillingGroupsSpendModifierRuleService struct {
	B   Backend
	Key string
}

// Retrieve a spend modifier associated with a billing group.
func (c v2BillingGroupsSpendModifierRuleService) Retrieve(ctx context.Context, id string, params *V2BillingGroupsSpendModifierRuleRetrieveParams) (*V2BillingGroupSpendModifier, error) {
	if params == nil {
		params = &V2BillingGroupsSpendModifierRuleRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/groups/%s/spend_modifier_rules/%s", StringValue(
			params.GroupID), id)
	groupspendmodifier := &V2BillingGroupSpendModifier{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, groupspendmodifier)
	return groupspendmodifier, err
}

// List spend modifiers associated with a billing group.
func (c v2BillingGroupsSpendModifierRuleService) List(ctx context.Context, listParams *V2BillingGroupsSpendModifierRuleListParams) *V2List[*V2BillingGroupSpendModifier] {
	if listParams == nil {
		listParams = &V2BillingGroupsSpendModifierRuleListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/groups/%s/spend_modifier_rules", StringValue(
			listParams.GroupID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingGroupSpendModifier], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingGroupSpendModifier]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
