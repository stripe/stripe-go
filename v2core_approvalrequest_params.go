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
	// Maximum number of results to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// GET /v2/core/approval_requests/:id
// Retrieves an approval request by ID.
type V2CoreApprovalRequestParams struct {
	Params `form:"*"`
}

// POST /v2/core/approval_requests/:id/cancel
// Cancels a pending approval request.
type V2CoreApprovalRequestCancelParams struct {
	Params `form:"*"`
}

// POST /v2/core/approval_requests/:id/execute
// Executes an approved approval request.
type V2CoreApprovalRequestExecuteParams struct {
	Params `form:"*"`
}

// POST /v2/core/approval_requests/:id/submit
// Moves a pending approval request into the reviewer queue for auto-execution upon approval.
type V2CoreApprovalRequestSubmitParams struct {
	Params `form:"*"`
	// The reason for submitting the approval request.
	Reason *string `form:"reason" json:"reason,omitempty"`
}

// GET /v2/core/approval_requests/:id
// Retrieves an approval request by ID.
type V2CoreApprovalRequestRetrieveParams struct {
	Params `form:"*"`
}
