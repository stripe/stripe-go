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

// v2ExtendWorkflowRunService is used to invoke workflowrun related APIs.
type v2ExtendWorkflowRunService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Workflow Run by ID.
func (c v2ExtendWorkflowRunService) Retrieve(ctx context.Context, id string, params *V2ExtendWorkflowRunRetrieveParams) (*V2ExtendWorkflowRun, error) {
	if params == nil {
		params = &V2ExtendWorkflowRunRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/extend/workflow_runs/%s", id)
	workflowrun := &V2ExtendWorkflowRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, workflowrun)
	return workflowrun, err
}

// List all Workflow Runs.
func (c v2ExtendWorkflowRunService) List(ctx context.Context, listParams *V2ExtendWorkflowRunListParams) *V2List[*V2ExtendWorkflowRun] {
	if listParams == nil {
		listParams = &V2ExtendWorkflowRunListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/extend/workflow_runs", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2ExtendWorkflowRun], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2ExtendWorkflowRun]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
