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

// v2CoreApprovalRequestService is used to invoke approvalrequest related APIs.
type v2CoreApprovalRequestService struct {
	B   Backend
	Key string
}

// GET /v2/core/approval_requests/:id
// Retrieves an approval request by ID.
func (c v2CoreApprovalRequestService) Retrieve(ctx context.Context, id string, params *V2CoreApprovalRequestRetrieveParams) (*V2CoreApprovalRequest, error) {
	if params == nil {
		params = &V2CoreApprovalRequestRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/approval_requests/%s", id)
	approvalrequest := &V2CoreApprovalRequest{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, approvalrequest)
	return approvalrequest, err
}

// POST /v2/core/approval_requests/:id/cancel
// Cancels a pending approval request.
func (c v2CoreApprovalRequestService) Cancel(ctx context.Context, id string, params *V2CoreApprovalRequestCancelParams) (*V2CoreApprovalRequest, error) {
	if params == nil {
		params = &V2CoreApprovalRequestCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/approval_requests/%s/cancel", id)
	approvalrequest := &V2CoreApprovalRequest{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, approvalrequest)
	return approvalrequest, err
}

// POST /v2/core/approval_requests/:id/execute
// Executes an approved approval request.
func (c v2CoreApprovalRequestService) Execute(ctx context.Context, id string, params *V2CoreApprovalRequestExecuteParams) (*V2CoreApprovalRequest, error) {
	if params == nil {
		params = &V2CoreApprovalRequestExecuteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/approval_requests/%s/execute", id)
	approvalrequest := &V2CoreApprovalRequest{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, approvalrequest)
	return approvalrequest, err
}

// POST /v2/core/approval_requests/:id/submit
// Moves a pending approval request into the reviewer queue for auto-execution upon approval.
func (c v2CoreApprovalRequestService) Submit(ctx context.Context, id string, params *V2CoreApprovalRequestSubmitParams) (*V2CoreApprovalRequest, error) {
	if params == nil {
		params = &V2CoreApprovalRequestSubmitParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/approval_requests/%s/submit", id)
	approvalrequest := &V2CoreApprovalRequest{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, approvalrequest)
	return approvalrequest, err
}

// GET /v2/core/approval_requests
// Lists approval requests with optional filtering.
func (c v2CoreApprovalRequestService) List(ctx context.Context, listParams *V2CoreApprovalRequestListParams) *V2List[*V2CoreApprovalRequest] {
	if listParams == nil {
		listParams = &V2CoreApprovalRequestListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/core/approval_requests", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2CoreApprovalRequest], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2CoreApprovalRequest]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
