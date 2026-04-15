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

// v2CoreWorkflowRunService is used to invoke workflowrun related APIs.
type v2CoreWorkflowRunService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Workflow Run by ID.
func (c v2CoreWorkflowRunService) Retrieve(ctx context.Context, id string, params *V2CoreWorkflowRunRetrieveParams) (*V2CoreWorkflowRun, error) {
	if params == nil {
		params = &V2CoreWorkflowRunRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/workflow_runs/%s", id)
	workflowrun := &V2CoreWorkflowRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, workflowrun)
	return workflowrun, err
}

// List all Workflow Runs.
func (c v2CoreWorkflowRunService) List(ctx context.Context, listParams *V2CoreWorkflowRunListParams) *V2List[*V2CoreWorkflowRun] {
	if listParams == nil {
		listParams = &V2CoreWorkflowRunListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/core/workflow_runs", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2CoreWorkflowRun], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2CoreWorkflowRun]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
