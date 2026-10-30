//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The status of the RewardProgramEnrollment. Active by default on creation; use activate/deactivate to
// change.
type V2MoneyManagementRewardProgramEnrollmentStatus string

// List of values that V2MoneyManagementRewardProgramEnrollmentStatus can take
const (
	V2MoneyManagementRewardProgramEnrollmentStatusActive   V2MoneyManagementRewardProgramEnrollmentStatus = "active"
	V2MoneyManagementRewardProgramEnrollmentStatusInactive V2MoneyManagementRewardProgramEnrollmentStatus = "inactive"
)

// Issuing-specific per-enrollment configuration. Immutable after creation.
type V2MoneyManagementRewardProgramEnrollmentCardSpend struct {
	// The Issuer Program for this connected account. Immutable after creation.
	IssuerProgram string `json:"issuer_program"`
}

// A RewardProgramEnrollment represents the binding of a connected account / Issuer Program to a shared
// RewardProgram. RewardProgramEnrollments are platform-owned canonical objects and are active by default
// on creation. At most one RewardProgramEnrollment may be active per Issuer Program globally, across all
// RewardPrograms.
type V2MoneyManagementRewardProgramEnrollment struct {
	APIResource
	// The connected account enrolled in this RewardProgram. Immutable after creation.
	Account string `json:"account"`
	// Issuing-specific per-enrollment configuration. Immutable after creation.
	CardSpend *V2MoneyManagementRewardProgramEnrollmentCardSpend `json:"card_spend"`
	// Time at which the RewardProgramEnrollment was created.
	Created time.Time `json:"created"`
	// Unique identifier for the RewardProgramEnrollment.
	ID string `json:"id"`
	// Whether this object exists in livemode or testmode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the RewardProgram this enrollment belongs to. Immutable after creation.
	RewardProgram string `json:"reward_program"`
	// The status of the RewardProgramEnrollment. Active by default on creation; use activate/deactivate to
	// change.
	Status V2MoneyManagementRewardProgramEnrollmentStatus `json:"status"`
}
