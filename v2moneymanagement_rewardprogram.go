//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Determines who performs reward money movement after rule evaluation. Immutable after creation.
type V2MoneyManagementRewardProgramFulfilledBy string

// List of values that V2MoneyManagementRewardProgramFulfilledBy can take
const (
	V2MoneyManagementRewardProgramFulfilledByApplication V2MoneyManagementRewardProgramFulfilledBy = "application"
)

// The status of the RewardProgram. Inactive by default on creation; use activate/deactivate to change.
type V2MoneyManagementRewardProgramStatus string

// List of values that V2MoneyManagementRewardProgramStatus can take
const (
	V2MoneyManagementRewardProgramStatusActive   V2MoneyManagementRewardProgramStatus = "active"
	V2MoneyManagementRewardProgramStatusInactive V2MoneyManagementRewardProgramStatus = "inactive"
)

// A RewardProgram represents shared card-spend cashback reward configuration for Issuing Card-Spend
// Rewards. It is not scoped to a single Issuer Program or connected account.
type V2MoneyManagementRewardProgram struct {
	APIResource
	// Time at which the RewardProgram was created.
	Created time.Time `json:"created"`
	// Determines who performs reward money movement after rule evaluation. Immutable after creation.
	FulfilledBy V2MoneyManagementRewardProgramFulfilledBy `json:"fulfilled_by"`
	// Unique identifier for the RewardProgram.
	ID string `json:"id"`
	// Whether this object exists in livemode or testmode.
	Livemode bool `json:"livemode"`
	// Display name for the RewardProgram.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The status of the RewardProgram. Inactive by default on creation; use activate/deactivate to change.
	Status V2MoneyManagementRewardProgramStatus `json:"status"`
}
