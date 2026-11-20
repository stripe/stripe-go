//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The FinancialAccount holder's reason for disputing the money movement.
type V2MoneyManagementDisputeReason string

// List of values that V2MoneyManagementDisputeReason can take
const (
	V2MoneyManagementDisputeReasonAuthorizationRevoked V2MoneyManagementDisputeReason = "authorization_revoked"
	V2MoneyManagementDisputeReasonDuplicate            V2MoneyManagementDisputeReason = "duplicate"
	V2MoneyManagementDisputeReasonIncorrectAmount      V2MoneyManagementDisputeReason = "incorrect_amount"
	V2MoneyManagementDisputeReasonIncorrectDate        V2MoneyManagementDisputeReason = "incorrect_date"
	V2MoneyManagementDisputeReasonOther                V2MoneyManagementDisputeReason = "other"
	V2MoneyManagementDisputeReasonUnauthorized         V2MoneyManagementDisputeReason = "unauthorized"
)

// The current status of the dispute.
type V2MoneyManagementDisputeStatus string

// List of values that V2MoneyManagementDisputeStatus can take
const (
	V2MoneyManagementDisputeStatusLost        V2MoneyManagementDisputeStatus = "lost"
	V2MoneyManagementDisputeStatusUnderReview V2MoneyManagementDisputeStatus = "under_review"
	V2MoneyManagementDisputeStatusWon         V2MoneyManagementDisputeStatus = "won"
)

// The type of money movement being disputed.
type V2MoneyManagementDisputeTargetType string

// List of values that V2MoneyManagementDisputeTargetType can take
const (
	V2MoneyManagementDisputeTargetTypeReceivedDebit V2MoneyManagementDisputeTargetType = "received_debit"
)

// Timestamps describing when the dispute changed status.
type V2MoneyManagementDisputeStatusTransitions struct {
	// The time at which the Dispute became lost. Only present if status == lost.
	LostAt time.Time `json:"lost_at,omitempty"`
	// The time at which the Dispute became won. Only present if status == won.
	WonAt time.Time `json:"won_at,omitempty"`
}

// The money movement being disputed.
type V2MoneyManagementDisputeTarget struct {
	// If applicable, the ID of the ReceivedDebit being disputed.
	ReceivedDebit string `json:"received_debit,omitempty"`
	// The type of money movement being disputed.
	Type V2MoneyManagementDisputeTargetType `json:"type"`
}

// A Dispute tracks a claim challenging a money movement on behalf of a FinancialAccount holder.
type V2MoneyManagementDispute struct {
	APIResource
	// The amount and currency of the disputed money movement.
	Amount Amount `json:"amount"`
	// The time at which the dispute was created.
	Created time.Time `json:"created"`
	// Additional explanation for the dispute, if supplied.
	Explanation string `json:"explanation,omitempty"`
	// The FinancialAccount associated with the disputed money movement.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the Dispute.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The FinancialAccount holder's reason for disputing the money movement.
	Reason V2MoneyManagementDisputeReason `json:"reason"`
	// The current status of the dispute.
	Status V2MoneyManagementDisputeStatus `json:"status"`
	// Timestamps describing when the dispute changed status.
	StatusTransitions *V2MoneyManagementDisputeStatusTransitions `json:"status_transitions,omitempty"`
	// The money movement being disputed.
	Target *V2MoneyManagementDisputeTarget `json:"target"`
}
