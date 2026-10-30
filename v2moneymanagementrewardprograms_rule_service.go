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

// v2MoneyManagementRewardProgramsRuleService is used to invoke rule related APIs.
type v2MoneyManagementRewardProgramsRuleService struct {
	B   Backend
	Key string
}

// Creates a new RewardProgramRule under the given RewardProgram. RewardProgramRules are inactive by
// default on creation; use ActivateRewardProgramRule to make them eligible for reward evaluation.
func (c v2MoneyManagementRewardProgramsRuleService) Create(ctx context.Context, params *V2MoneyManagementRewardProgramsRuleCreateParams) (*V2MoneyManagementRewardProgramRule, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramsRuleCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules", StringValue(
			params.RewardProgramID))
	rewardprogramrule := &V2MoneyManagementRewardProgramRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramrule)
	return rewardprogramrule, err
}

// Retrieves the details of an existing RewardProgramRule.
func (c v2MoneyManagementRewardProgramsRuleService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementRewardProgramsRuleRetrieveParams) (*V2MoneyManagementRewardProgramRule, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramsRuleRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules/%s", StringValue(
			params.RewardProgramID), id)
	rewardprogramrule := &V2MoneyManagementRewardProgramRule{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, rewardprogramrule)
	return rewardprogramrule, err
}

// Updates the mutable fields of an existing RewardProgramRule. Use ActivateRewardProgramRule /
// DeactivateRewardProgramRule to change the RewardProgramRule's status.
func (c v2MoneyManagementRewardProgramsRuleService) Update(ctx context.Context, id string, params *V2MoneyManagementRewardProgramsRuleUpdateParams) (*V2MoneyManagementRewardProgramRule, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramsRuleUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules/%s", StringValue(
			params.RewardProgramID), id)
	rewardprogramrule := &V2MoneyManagementRewardProgramRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramrule)
	return rewardprogramrule, err
}

// Activates a RewardProgramRule. At most one RewardProgramRule may be active per RewardProgram at a
// time. Idempotent: calling this on a RewardProgramRule that is already active returns 200 with the
// current object unchanged.
func (c v2MoneyManagementRewardProgramsRuleService) Activate(ctx context.Context, id string, params *V2MoneyManagementRewardProgramsRuleActivateParams) (*V2MoneyManagementRewardProgramRule, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramsRuleActivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules/%s/activate", StringValue(
			params.RewardProgramID), id)
	rewardprogramrule := &V2MoneyManagementRewardProgramRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramrule)
	return rewardprogramrule, err
}

// Deactivates a RewardProgramRule. Idempotent: calling this on a RewardProgramRule that is already
// inactive returns 200 with the current object unchanged.
func (c v2MoneyManagementRewardProgramsRuleService) Deactivate(ctx context.Context, id string, params *V2MoneyManagementRewardProgramsRuleDeactivateParams) (*V2MoneyManagementRewardProgramRule, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramsRuleDeactivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules/%s/deactivate", StringValue(
			params.RewardProgramID), id)
	rewardprogramrule := &V2MoneyManagementRewardProgramRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramrule)
	return rewardprogramrule, err
}

// Returns a list of RewardProgramRules for a RewardProgram.
func (c v2MoneyManagementRewardProgramsRuleService) List(ctx context.Context, listParams *V2MoneyManagementRewardProgramsRuleListParams) *V2List[*V2MoneyManagementRewardProgramRule] {
	if listParams == nil {
		listParams = &V2MoneyManagementRewardProgramsRuleListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules", StringValue(
			listParams.RewardProgramID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementRewardProgramRule], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementRewardProgramRule]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
