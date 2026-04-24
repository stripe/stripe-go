//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The action that was requested.
type V2CoreApprovalRequestAction string

// List of values that V2CoreApprovalRequestAction can take
const (
	V2CoreApprovalRequestActionChargeCreate                           V2CoreApprovalRequestAction = "charge.create"
	V2CoreApprovalRequestActionDisputeClose                           V2CoreApprovalRequestAction = "dispute.close"
	V2CoreApprovalRequestActionInboundTransfersMoneyManagementCreate  V2CoreApprovalRequestAction = "inbound_transfers.money_management.create"
	V2CoreApprovalRequestActionInvoiceCreate                          V2CoreApprovalRequestAction = "invoice.create"
	V2CoreApprovalRequestActionOutboundPaymentsMoneyManagementCreate  V2CoreApprovalRequestAction = "outbound_payments.money_management.create"
	V2CoreApprovalRequestActionOutboundTransfersMoneyManagementCreate V2CoreApprovalRequestAction = "outbound_transfers.money_management.create"
	V2CoreApprovalRequestActionPaymentIntentCreate                    V2CoreApprovalRequestAction = "payment_intent.create"
	V2CoreApprovalRequestActionPaymentIntentUpdate                    V2CoreApprovalRequestAction = "payment_intent.update"
	V2CoreApprovalRequestActionPayoutCreate                           V2CoreApprovalRequestAction = "payout.create"
	V2CoreApprovalRequestActionPriceUpdate                            V2CoreApprovalRequestAction = "price.update"
	V2CoreApprovalRequestActionRefundCreate                           V2CoreApprovalRequestAction = "refund.create"
	V2CoreApprovalRequestActionSetupIntentCreate                      V2CoreApprovalRequestAction = "setup_intent.create"
	V2CoreApprovalRequestActionSubscriptionCreate                     V2CoreApprovalRequestAction = "subscription.create"
	V2CoreApprovalRequestActionSubscriptionUpdate                     V2CoreApprovalRequestAction = "subscription.update"
	V2CoreApprovalRequestActionTopupCreate                            V2CoreApprovalRequestAction = "topup.create"
	V2CoreApprovalRequestActionTransferCreate                         V2CoreApprovalRequestAction = "transfer.create"
)

// The result of the review.
type V2CoreApprovalRequestReviewResult string

// List of values that V2CoreApprovalRequestReviewResult can take
const (
	V2CoreApprovalRequestReviewResultApproved V2CoreApprovalRequestReviewResult = "approved"
	V2CoreApprovalRequestReviewResultRejected V2CoreApprovalRequestReviewResult = "rejected"
)

// The status of this ApprovalRequest.
type V2CoreApprovalRequestStatus string

// List of values that V2CoreApprovalRequestStatus can take
const (
	V2CoreApprovalRequestStatusApproved           V2CoreApprovalRequestStatus = "approved"
	V2CoreApprovalRequestStatusCanceled           V2CoreApprovalRequestStatus = "canceled"
	V2CoreApprovalRequestStatusExecutionFailed    V2CoreApprovalRequestStatus = "execution_failed"
	V2CoreApprovalRequestStatusExecutionStarted   V2CoreApprovalRequestStatus = "execution_started"
	V2CoreApprovalRequestStatusExecutionSucceeded V2CoreApprovalRequestStatus = "execution_succeeded"
	V2CoreApprovalRequestStatusExpired            V2CoreApprovalRequestStatus = "expired"
	V2CoreApprovalRequestStatusFailed             V2CoreApprovalRequestStatus = "failed"
	V2CoreApprovalRequestStatusPending            V2CoreApprovalRequestStatus = "pending"
	V2CoreApprovalRequestStatusRejected           V2CoreApprovalRequestStatus = "rejected"
	V2CoreApprovalRequestStatusRequiresExecution  V2CoreApprovalRequestStatus = "requires_execution"
	V2CoreApprovalRequestStatusRequiresReview     V2CoreApprovalRequestStatus = "requires_review"
	V2CoreApprovalRequestStatusSucceeded          V2CoreApprovalRequestStatus = "succeeded"
)

// The requester of this ApprovalRequest.
type V2CoreApprovalRequestRequestedBy struct {
	// Stripe-defined identifier for the requester (e.g. a restricted API key token).
	ID string `json:"id"`
	// Merchant-defined name for the requester.
	Name string `json:"name,omitempty"`
}

// The reviewer who performed the review.
type V2CoreApprovalRequestReviewReviewedBy struct {
	// Stripe-defined identifier for the reviewer (e.g. a restricted API key token).
	ID string `json:"id"`
	// Merchant-defined name for the reviewer.
	Name string `json:"name"`
}

// The review of this ApprovalRequest if it has been reviewed.
type V2CoreApprovalRequestReview struct {
	// The reason provided by the reviewer.
	Reason string `json:"reason,omitempty"`
	// The result of the review.
	Result V2CoreApprovalRequestReviewResult `json:"result"`
	// Timestamp when the review was performed.
	ReviewedAt time.Time `json:"reviewed_at"`
	// The reviewer who performed the review.
	ReviewedBy *V2CoreApprovalRequestReviewReviewedBy `json:"reviewed_by"`
}

// The rule associated with this ApprovalRequest.
type V2CoreApprovalRequestRule struct {
	// The name of the rule.
	Name string `json:"name"`
}

// Deprecated: use requires_execution status instead.
type V2CoreApprovalRequestStatusDetailsApproved struct {
	// The reason provided when approving the request.
	Reason string `json:"reason,omitempty"`
}

// Deprecated: use canceled status instead.
type V2CoreApprovalRequestStatusDetailsCanceled struct{}

// Deprecated: use failed status instead.
type V2CoreApprovalRequestStatusDetailsExecutionFailed struct {
	// The error code for the failed execution.
	Code string `json:"code"`
	// The error message for the failed execution.
	Message string `json:"message"`
	// The error type for the failed execution.
	Type string `json:"type"`
}

// Deprecated: use requires_execution status instead.
type V2CoreApprovalRequestStatusDetailsExecutionStarted struct{}

// The result of the successful execution.
type V2CoreApprovalRequestStatusDetailsExecutionSucceededResult struct {
	// The unique identifier of the executed object.
	ID string `json:"id"`
	// The object type of the executed resource.
	Object string `json:"object"`
}

// Deprecated: use succeeded status instead.
type V2CoreApprovalRequestStatusDetailsExecutionSucceeded struct {
	// The result of the successful execution.
	Result *V2CoreApprovalRequestStatusDetailsExecutionSucceededResult `json:"result"`
}

// Deprecated: use expired status instead.
type V2CoreApprovalRequestStatusDetailsExpired struct{}

// Details when the approval request failed.
type V2CoreApprovalRequestStatusDetailsFailed struct {
	// The error code for the failed execution.
	ErrorCode string `json:"error_code"`
	// The error message for the failed execution.
	ErrorMessage string `json:"error_message"`
	// The error type for the failed execution.
	ErrorType string `json:"error_type"`
}

// Deprecated: use requires_review status instead.
type V2CoreApprovalRequestStatusDetailsPending struct{}

// Deprecated: use rejected status instead.
type V2CoreApprovalRequestStatusDetailsRejected struct {
	// The reason provided when rejecting the request.
	Reason string `json:"reason,omitempty"`
}

// The result of the successful execution.
type V2CoreApprovalRequestStatusDetailsSucceededResult struct {
	// The unique identifier of the executed object.
	ID string `json:"id"`
	// The object type of the executed resource.
	Object string `json:"object"`
}

// Details when the approval request succeeded.
type V2CoreApprovalRequestStatusDetailsSucceeded struct {
	// The result of the successful execution.
	Result *V2CoreApprovalRequestStatusDetailsSucceededResult `json:"result"`
}

// The details of the status of this ApprovalRequest.
type V2CoreApprovalRequestStatusDetails struct {
	// Deprecated: use requires_execution status instead.
	Approved *V2CoreApprovalRequestStatusDetailsApproved `json:"approved,omitempty"`
	// Deprecated: use canceled status instead.
	Canceled *V2CoreApprovalRequestStatusDetailsCanceled `json:"canceled,omitempty"`
	// Deprecated: use failed status instead.
	ExecutionFailed *V2CoreApprovalRequestStatusDetailsExecutionFailed `json:"execution_failed,omitempty"`
	// Deprecated: use requires_execution status instead.
	ExecutionStarted *V2CoreApprovalRequestStatusDetailsExecutionStarted `json:"execution_started,omitempty"`
	// Deprecated: use succeeded status instead.
	ExecutionSucceeded *V2CoreApprovalRequestStatusDetailsExecutionSucceeded `json:"execution_succeeded,omitempty"`
	// Deprecated: use expired status instead.
	Expired *V2CoreApprovalRequestStatusDetailsExpired `json:"expired,omitempty"`
	// Details when the approval request failed.
	Failed *V2CoreApprovalRequestStatusDetailsFailed `json:"failed,omitempty"`
	// Deprecated: use requires_review status instead.
	Pending *V2CoreApprovalRequestStatusDetailsPending `json:"pending,omitempty"`
	// Deprecated: use rejected status instead.
	Rejected *V2CoreApprovalRequestStatusDetailsRejected `json:"rejected,omitempty"`
	// Details when the approval request succeeded.
	Succeeded *V2CoreApprovalRequestStatusDetailsSucceeded `json:"succeeded,omitempty"`
}

// The transitions of the status of this ApprovalRequest.
type V2CoreApprovalRequestStatusTransitions struct {
	// Timestamp when the approval request was canceled.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Timestamp when the approval request expired.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	// Timestamp when the approval request failed.
	FailedAt time.Time `json:"failed_at,omitempty"`
	// Timestamp when the approval request was rejected.
	RejectedAt time.Time `json:"rejected_at,omitempty"`
	// Timestamp when the approval request moved to requires_execution status.
	RequiresExecutionAt time.Time `json:"requires_execution_at,omitempty"`
	// Timestamp when the approval request succeeded.
	SucceededAt time.Time `json:"succeeded_at,omitempty"`
}

// An approval request represents a pending action that requires review before execution.
type V2CoreApprovalRequest struct {
	APIResource
	// The action that was requested.
	Action V2CoreApprovalRequestAction `json:"action"`
	// Time this ApprovalRequest was created.
	Created time.Time `json:"created"`
	// The URL to the dashboard for this ApprovalRequest.
	DashboardURL string `json:"dashboard_url,omitempty"`
	// A description of the approval request.
	Description string `json:"description,omitempty"`
	// The timestamp at which this ApprovalRequest will expire.
	ExpiresAt time.Time `json:"expires_at"`
	// The unique identifier for this ApprovalRequest.
	ID string `json:"id"`
	// Whether this ApprovalRequest is livemode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The requester of this ApprovalRequest.
	RequestedBy *V2CoreApprovalRequestRequestedBy `json:"requested_by"`
	// The review of this ApprovalRequest if it has been reviewed.
	Review *V2CoreApprovalRequestReview `json:"review,omitempty"`
	// The rule associated with this ApprovalRequest.
	Rule *V2CoreApprovalRequestRule `json:"rule,omitempty"`
	// The status of this ApprovalRequest.
	Status V2CoreApprovalRequestStatus `json:"status"`
	// The details of the status of this ApprovalRequest.
	StatusDetails *V2CoreApprovalRequestStatusDetails `json:"status_details,omitempty"`
	// The transitions of the status of this ApprovalRequest.
	StatusTransitions *V2CoreApprovalRequestStatusTransitions `json:"status_transitions,omitempty"`
}
