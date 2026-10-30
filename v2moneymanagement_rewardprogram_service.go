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

// v2MoneyManagementRewardProgramService is used to invoke rewardprogram related APIs.
type v2MoneyManagementRewardProgramService struct {
	B   Backend
	Key string
}

// Creates a new RewardProgram. RewardPrograms are inactive by default on creation; use
// ActivateRewardProgram to make them eligible for reward evaluation.
func (c v2MoneyManagementRewardProgramService) Create(ctx context.Context, params *V2MoneyManagementRewardProgramCreateParams) (*V2MoneyManagementRewardProgram, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramCreateParams{}
	}
	params.Context = ctx
	rewardprogram := &V2MoneyManagementRewardProgram{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/reward_programs", c.Key, params, rewardprogram)
	return rewardprogram, err
}

// Retrieves the details of an existing RewardProgram.
func (c v2MoneyManagementRewardProgramService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementRewardProgramRetrieveParams) (*V2MoneyManagementRewardProgram, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/reward_programs/%s", id)
	rewardprogram := &V2MoneyManagementRewardProgram{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, rewardprogram)
	return rewardprogram, err
}

// Updates the mutable fields of an existing RewardProgram. Use ActivateRewardProgram /
// DeactivateRewardProgram to change the RewardProgram's status.
func (c v2MoneyManagementRewardProgramService) Update(ctx context.Context, id string, params *V2MoneyManagementRewardProgramUpdateParams) (*V2MoneyManagementRewardProgram, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/reward_programs/%s", id)
	rewardprogram := &V2MoneyManagementRewardProgram{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogram)
	return rewardprogram, err
}

// Activates a RewardProgram, making it eligible for reward evaluation. Idempotent: calling this on a
// RewardProgram that is already active returns 200 with the current object unchanged.
func (c v2MoneyManagementRewardProgramService) Activate(ctx context.Context, id string, params *V2MoneyManagementRewardProgramActivateParams) (*V2MoneyManagementRewardProgram, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramActivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/reward_programs/%s/activate", id)
	rewardprogram := &V2MoneyManagementRewardProgram{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogram)
	return rewardprogram, err
}

// Deactivates a RewardProgram. Idempotent: calling this on a RewardProgram that is already inactive
// returns 200 with the current object unchanged.
func (c v2MoneyManagementRewardProgramService) Deactivate(ctx context.Context, id string, params *V2MoneyManagementRewardProgramDeactivateParams) (*V2MoneyManagementRewardProgram, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramDeactivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/reward_programs/%s/deactivate", id)
	rewardprogram := &V2MoneyManagementRewardProgram{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogram)
	return rewardprogram, err
}

// Returns a list of RewardPrograms.
func (c v2MoneyManagementRewardProgramService) List(ctx context.Context, listParams *V2MoneyManagementRewardProgramListParams) *V2List[*V2MoneyManagementRewardProgram] {
	if listParams == nil {
		listParams = &V2MoneyManagementRewardProgramListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/reward_programs", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementRewardProgram], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementRewardProgram]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
