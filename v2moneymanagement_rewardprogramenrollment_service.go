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

// v2MoneyManagementRewardProgramEnrollmentService is used to invoke rewardprogramenrollment related APIs.
type v2MoneyManagementRewardProgramEnrollmentService struct {
	B   Backend
	Key string
}

// Creates a new RewardProgramEnrollment, binding a connected account / Issuer Program to a
// RewardProgram. RewardProgramEnrollments are active by default on creation.
func (c v2MoneyManagementRewardProgramEnrollmentService) Create(ctx context.Context, params *V2MoneyManagementRewardProgramEnrollmentCreateParams) (*V2MoneyManagementRewardProgramEnrollment, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramEnrollmentCreateParams{}
	}
	params.Context = ctx
	rewardprogramenrollment := &V2MoneyManagementRewardProgramEnrollment{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/reward_program_enrollments", c.Key, params, rewardprogramenrollment)
	return rewardprogramenrollment, err
}

// Retrieves the details of an existing RewardProgramEnrollment.
func (c v2MoneyManagementRewardProgramEnrollmentService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementRewardProgramEnrollmentRetrieveParams) (*V2MoneyManagementRewardProgramEnrollment, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramEnrollmentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/reward_program_enrollments/%s", id)
	rewardprogramenrollment := &V2MoneyManagementRewardProgramEnrollment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, rewardprogramenrollment)
	return rewardprogramenrollment, err
}

// Activates a RewardProgramEnrollment. At most one RewardProgramEnrollment may be active per Issuer
// Program globally, across all RewardPrograms. Idempotent: calling this on a RewardProgramEnrollment
// that is already active returns 200 with the current object unchanged.
func (c v2MoneyManagementRewardProgramEnrollmentService) Activate(ctx context.Context, id string, params *V2MoneyManagementRewardProgramEnrollmentActivateParams) (*V2MoneyManagementRewardProgramEnrollment, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramEnrollmentActivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/reward_program_enrollments/%s/activate", id)
	rewardprogramenrollment := &V2MoneyManagementRewardProgramEnrollment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramenrollment)
	return rewardprogramenrollment, err
}

// Deactivates a RewardProgramEnrollment. Idempotent: calling this on a RewardProgramEnrollment that is
// already inactive returns 200 with the current object unchanged.
func (c v2MoneyManagementRewardProgramEnrollmentService) Deactivate(ctx context.Context, id string, params *V2MoneyManagementRewardProgramEnrollmentDeactivateParams) (*V2MoneyManagementRewardProgramEnrollment, error) {
	if params == nil {
		params = &V2MoneyManagementRewardProgramEnrollmentDeactivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/reward_program_enrollments/%s/deactivate", id)
	rewardprogramenrollment := &V2MoneyManagementRewardProgramEnrollment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramenrollment)
	return rewardprogramenrollment, err
}

// Returns a list of RewardProgramEnrollments. Omitting `reward_program` returns platform-wide
// enrollments across all RewardPrograms.
func (c v2MoneyManagementRewardProgramEnrollmentService) List(ctx context.Context, listParams *V2MoneyManagementRewardProgramEnrollmentListParams) *V2List[*V2MoneyManagementRewardProgramEnrollment] {
	if listParams == nil {
		listParams = &V2MoneyManagementRewardProgramEnrollmentListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/reward_program_enrollments", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementRewardProgramEnrollment], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementRewardProgramEnrollment]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
