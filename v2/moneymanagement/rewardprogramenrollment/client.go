//
//
// File generated from our OpenAPI spec
//
//

// Package rewardprogramenrollment provides the rewardprogramenrollment related APIs
package rewardprogramenrollment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke rewardprogramenrollment related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new RewardProgramEnrollment, binding a connected account / Issuer Program to a
// RewardProgram. RewardProgramEnrollments are active by default on creation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementRewardProgramEnrollmentParams) (*stripe.V2MoneyManagementRewardProgramEnrollment, error) {
	rewardprogramenrollment := &stripe.V2MoneyManagementRewardProgramEnrollment{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/reward_program_enrollments", c.Key, params, rewardprogramenrollment)
	return rewardprogramenrollment, err
}

// Retrieves the details of an existing RewardProgramEnrollment.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementRewardProgramEnrollmentParams) (*stripe.V2MoneyManagementRewardProgramEnrollment, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_program_enrollments/%s", id)
	rewardprogramenrollment := &stripe.V2MoneyManagementRewardProgramEnrollment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, rewardprogramenrollment)
	return rewardprogramenrollment, err
}

// Activates a RewardProgramEnrollment. At most one RewardProgramEnrollment may be active per Issuer
// Program globally, across all RewardPrograms. Idempotent: calling this on a RewardProgramEnrollment
// that is already active returns 200 with the current object unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Activate(id string, params *stripe.V2MoneyManagementRewardProgramEnrollmentActivateParams) (*stripe.V2MoneyManagementRewardProgramEnrollment, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_program_enrollments/%s/activate", id)
	rewardprogramenrollment := &stripe.V2MoneyManagementRewardProgramEnrollment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramenrollment)
	return rewardprogramenrollment, err
}

// Deactivates a RewardProgramEnrollment. Idempotent: calling this on a RewardProgramEnrollment that is
// already inactive returns 200 with the current object unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.V2MoneyManagementRewardProgramEnrollmentDeactivateParams) (*stripe.V2MoneyManagementRewardProgramEnrollment, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/reward_program_enrollments/%s/deactivate", id)
	rewardprogramenrollment := &stripe.V2MoneyManagementRewardProgramEnrollment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, rewardprogramenrollment)
	return rewardprogramenrollment, err
}

// Returns a list of RewardProgramEnrollments. Omitting `reward_program` returns platform-wide
// enrollments across all RewardPrograms.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementRewardProgramEnrollmentListParams) stripe.Seq2[*stripe.V2MoneyManagementRewardProgramEnrollment, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementRewardProgramEnrollmentListParams{}
	}
	return stripe.NewV2List("/v2/money_management/reward_program_enrollments", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementRewardProgramEnrollment], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementRewardProgramEnrollment]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
