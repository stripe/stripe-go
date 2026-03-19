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

// v2BillingCadencesSpendModifierRuleService is used to invoke spendmodifierrule related APIs.
type v2BillingCadencesSpendModifierRuleService struct {
	B   Backend
	Key string
}

// Retrieve a Spend Modifier associated with a Billing Cadence.
func (c v2BillingCadencesSpendModifierRuleService) Retrieve(ctx context.Context, id string, params *V2BillingCadencesSpendModifierRuleRetrieveParams) (*V2BillingCadenceSpendModifier, error) {
	if params == nil {
		params = &V2BillingCadencesSpendModifierRuleRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/cadences/%s/spend_modifier_rules/%s", StringValue(
			params.CadenceID), id)
	cadencespendmodifier := &V2BillingCadenceSpendModifier{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cadencespendmodifier)
	return cadencespendmodifier, err
}

// List all Spend Modifiers associated with a Billing Cadence.
func (c v2BillingCadencesSpendModifierRuleService) List(ctx context.Context, listParams *V2BillingCadencesSpendModifierRuleListParams) *V2List[*V2BillingCadenceSpendModifier] {
	if listParams == nil {
		listParams = &V2BillingCadencesSpendModifierRuleListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/cadences/%s/spend_modifier_rules", StringValue(
			listParams.CadenceID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingCadenceSpendModifier], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingCadenceSpendModifier]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
