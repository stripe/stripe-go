//
//
// File generated from our OpenAPI spec
//
//

package stripe

// GET /v2/core/approval_requests
// Lists approval requests with optional filtering.
type V2CoreApprovalRequestListParams struct {
	Params `form:"*"`
	// Filter by action type (e.g. "refund.create", "payment_intent.create", "payout.create").
	Action *string `form:"action" json:"action,omitempty"`
	// Filter by creation time.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// Maximum number of results to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by approval request status (e.g. "requires_review", "approved", "succeeded", "failed", "rejected", "canceled", "expired").
	Status *string `form:"status" json:"status,omitempty"`
}

// GET /v2/core/approval_requests/:id
// Retrieves an approval request by ID.
type V2CoreApprovalRequestParams struct {
	Params `form:"*"`
	// The updated reason for the approval request.
	Reason *string `form:"reason" json:"reason,omitempty"`
}

// POST /v2/core/approval_requests/:id/cancel
// Cancels a pending approval request.
type V2CoreApprovalRequestCancelParams struct {
	Params `form:"*"`
}

// GET /v2/core/approval_requests/:id
// Retrieves an approval request by ID.
type V2CoreApprovalRequestRetrieveParams struct {
	Params `form:"*"`
}

// POST /v2/core/approval_requests/:id
// Updates a pending approval request's mutable fields.
type V2CoreApprovalRequestUpdateParams struct {
	Params `form:"*"`
	// The updated reason for the approval request.
	Reason *string `form:"reason" json:"reason,omitempty"`
}
