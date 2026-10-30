//
//
// File generated from our OpenAPI spec
//
//

// Package rule provides the rule related APIs
package rule

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke rule related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new RewardProgramRule under the given RewardProgram. RewardProgramRules are inactive by
// default on creation; use ActivateRewardProgramRule to make them eligible for reward evaluation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementRewardProgramsRuleParams) (*stripe.V2MoneyManagementRewardProgramRule, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules", stripe.StringValue(
			params.RewardProgramID))
	rewardprogramrule := &stripe.V2MoneyManagementRewardProgramRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramrule)
	return rewardprogramrule, err
}

// Retrieves the details of an existing RewardProgramRule.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementRewardProgramsRuleParams) (*stripe.V2MoneyManagementRewardProgramRule, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules/%s", stripe.StringValue(
			params.RewardProgramID), id)
	rewardprogramrule := &stripe.V2MoneyManagementRewardProgramRule{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, rewardprogramrule)
	return rewardprogramrule, err
}

// Updates the mutable fields of an existing RewardProgramRule. Use ActivateRewardProgramRule /
// DeactivateRewardProgramRule to change the RewardProgramRule's status.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2MoneyManagementRewardProgramsRuleParams) (*stripe.V2MoneyManagementRewardProgramRule, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules/%s", stripe.StringValue(
			params.RewardProgramID), id)
	rewardprogramrule := &stripe.V2MoneyManagementRewardProgramRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramrule)
	return rewardprogramrule, err
}

// Activates a RewardProgramRule. At most one RewardProgramRule may be active per RewardProgram at a
// time. Idempotent: calling this on a RewardProgramRule that is already active returns 200 with the
// current object unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Activate(id string, params *stripe.V2MoneyManagementRewardProgramsRuleActivateParams) (*stripe.V2MoneyManagementRewardProgramRule, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules/%s/activate", stripe.StringValue(
			params.RewardProgramID), id)
	rewardprogramrule := &stripe.V2MoneyManagementRewardProgramRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramrule)
	return rewardprogramrule, err
}

// Deactivates a RewardProgramRule. Idempotent: calling this on a RewardProgramRule that is already
// inactive returns 200 with the current object unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.V2MoneyManagementRewardProgramsRuleDeactivateParams) (*stripe.V2MoneyManagementRewardProgramRule, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules/%s/deactivate", stripe.StringValue(
			params.RewardProgramID), id)
	rewardprogramrule := &stripe.V2MoneyManagementRewardProgramRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramrule)
	return rewardprogramrule, err
}

// Returns a list of RewardProgramRules for a RewardProgram.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementRewardProgramsRuleListParams) stripe.Seq2[*stripe.V2MoneyManagementRewardProgramRule, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementRewardProgramsRuleListParams{}
	}
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_programs/%s/rules", stripe.StringValue(
			listParams.RewardProgramID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementRewardProgramRule], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementRewardProgramRule]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
