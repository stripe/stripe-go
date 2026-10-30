//
//
// File generated from our OpenAPI spec
//
//

// Package rewardprogram provides the rewardprogram related APIs
package rewardprogram

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke rewardprogram related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new RewardProgram. RewardPrograms are inactive by default on creation; use
// ActivateRewardProgram to make them eligible for reward evaluation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementRewardProgramParams) (*stripe.V2MoneyManagementRewardProgram, error) {
	rewardprogram := &stripe.V2MoneyManagementRewardProgram{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/reward_programs", c.Key, params, rewardprogram)
	return rewardprogram, err
}

// Retrieves the details of an existing RewardProgram.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementRewardProgramParams) (*stripe.V2MoneyManagementRewardProgram, error) {
	path := stripe.FormatURLPath("/v2/money_management/reward_programs/%s", id)
	rewardprogram := &stripe.V2MoneyManagementRewardProgram{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, rewardprogram)
	return rewardprogram, err
}

// Updates the mutable fields of an existing RewardProgram. Use ActivateRewardProgram /
// DeactivateRewardProgram to change the RewardProgram's status.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2MoneyManagementRewardProgramParams) (*stripe.V2MoneyManagementRewardProgram, error) {
	path := stripe.FormatURLPath("/v2/money_management/reward_programs/%s", id)
	rewardprogram := &stripe.V2MoneyManagementRewardProgram{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogram)
	return rewardprogram, err
}

// Activates a RewardProgram, making it eligible for reward evaluation. Idempotent: calling this on a
// RewardProgram that is already active returns 200 with the current object unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Activate(id string, params *stripe.V2MoneyManagementRewardProgramActivateParams) (*stripe.V2MoneyManagementRewardProgram, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_programs/%s/activate", id)
	rewardprogram := &stripe.V2MoneyManagementRewardProgram{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogram)
	return rewardprogram, err
}

// Deactivates a RewardProgram. Idempotent: calling this on a RewardProgram that is already inactive
// returns 200 with the current object unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.V2MoneyManagementRewardProgramDeactivateParams) (*stripe.V2MoneyManagementRewardProgram, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_programs/%s/deactivate", id)
	rewardprogram := &stripe.V2MoneyManagementRewardProgram{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogram)
	return rewardprogram, err
}

// Returns a list of RewardPrograms.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementRewardProgramListParams) stripe.Seq2[*stripe.V2MoneyManagementRewardProgram, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementRewardProgramListParams{}
	}
	return stripe.NewV2List("/v2/money_management/reward_programs", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementRewardProgram], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementRewardProgram]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
