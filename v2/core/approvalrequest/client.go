//
//
// File generated from our OpenAPI spec
//
//

// Package approvalrequest provides the approvalrequest related APIs
package approvalrequest

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke approvalrequest related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// GET /v2/core/approval_requests/:id
// Retrieves an approval request by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreApprovalRequestParams) (*stripe.V2CoreApprovalRequest, error) {
	path := stripe.FormatURLPath("/v2/core/approval_requests/%s", id)
	approvalrequest := &stripe.V2CoreApprovalRequest{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, approvalrequest)
	return approvalrequest, err
}

// POST /v2/core/approval_requests/:id/cancel
// Cancels a pending approval request.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2CoreApprovalRequestCancelParams) (*stripe.V2CoreApprovalRequest, error) {
	path := stripe.FormatURLPath("/v2/core/approval_requests/%s/cancel", id)
	approvalrequest := &stripe.V2CoreApprovalRequest{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, approvalrequest)
	return approvalrequest, err
}

// POST /v2/core/approval_requests/:id/execute
// Executes an approved approval request.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Execute(id string, params *stripe.V2CoreApprovalRequestExecuteParams) (*stripe.V2CoreApprovalRequest, error) {
	path := stripe.FormatURLPath("/v2/core/approval_requests/%s/execute", id)
	approvalrequest := &stripe.V2CoreApprovalRequest{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, approvalrequest)
	return approvalrequest, err
}

// POST /v2/core/approval_requests/:id/submit
// Moves a pending approval request into the reviewer queue for auto-execution upon approval.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Submit(id string, params *stripe.V2CoreApprovalRequestSubmitParams) (*stripe.V2CoreApprovalRequest, error) {
	path := stripe.FormatURLPath("/v2/core/approval_requests/%s/submit", id)
	approvalrequest := &stripe.V2CoreApprovalRequest{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, approvalrequest)
	return approvalrequest, err
}

// GET /v2/core/approval_requests
// Lists approval requests with optional filtering.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreApprovalRequestListParams) stripe.Seq2[*stripe.V2CoreApprovalRequest, error] {
	if listParams == nil {
		listParams = &stripe.V2CoreApprovalRequestListParams{}
	}
	return stripe.NewV2List("/v2/core/approval_requests", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2CoreApprovalRequest], error) {
		page := &stripe.V2Page[*stripe.V2CoreApprovalRequest]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
