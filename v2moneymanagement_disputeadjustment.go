//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of Dispute associated with this DisputeAdjustment.
type V2MoneyManagementDisputeAdjustmentRelatedDisputeType string

// List of values that V2MoneyManagementDisputeAdjustmentRelatedDisputeType can take
const (
	V2MoneyManagementDisputeAdjustmentRelatedDisputeTypeDispute V2MoneyManagementDisputeAdjustmentRelatedDisputeType = "dispute"
)

// The Dispute associated with this DisputeAdjustment.
type V2MoneyManagementDisputeAdjustmentRelatedDispute struct {
	// The ID of the Dispute associated with this DisputeAdjustment. Present when `type` is `dispute`.
	Dispute string `json:"dispute,omitempty"`
	// The type of Dispute associated with this DisputeAdjustment.
	Type V2MoneyManagementDisputeAdjustmentRelatedDisputeType `json:"type"`
}

// A DisputeAdjustment records a money movement associated with a Dispute.
type V2MoneyManagementDisputeAdjustment struct {
	APIResource
	// The amount and currency of the DisputeAdjustment.
	Amount Amount `json:"amount"`
	// The time at which the DisputeAdjustment was created.
	Created time.Time `json:"created"`
	// The ID of the FinancialAccount associated with the DisputeAdjustment.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the DisputeAdjustment.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The Dispute associated with this DisputeAdjustment.
	RelatedDispute *V2MoneyManagementDisputeAdjustmentRelatedDispute `json:"related_dispute"`
}
