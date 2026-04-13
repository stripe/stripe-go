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

// v2CoreWorkflowService is used to invoke workflow related APIs.
type v2CoreWorkflowService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Workflow by ID.
func (c v2CoreWorkflowService) Retrieve(ctx context.Context, id string, params *V2CoreWorkflowRetrieveParams) (*V2CoreWorkflow, error) {
	if params == nil {
		params = &V2CoreWorkflowRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/workflows/%s", id)
	workflow := &V2CoreWorkflow{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, workflow)
	return workflow, err
}

// Invokes an on-demand Workflow with parameters, to launch a new Workflow Run.
func (c v2CoreWorkflowService) Invoke(ctx context.Context, id string, params *V2CoreWorkflowInvokeParams) (*V2CoreWorkflowRun, error) {
	if params == nil {
		params = &V2CoreWorkflowInvokeParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/workflows/%s/invoke", id)
	workflowrun := &V2CoreWorkflowRun{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, workflowrun)
	return workflowrun, err
}

// List all Workflows.
func (c v2CoreWorkflowService) List(ctx context.Context, listParams *V2CoreWorkflowListParams) *V2List[*V2CoreWorkflow] {
	if listParams == nil {
		listParams = &V2CoreWorkflowListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/core/workflows", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2CoreWorkflow], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2CoreWorkflow]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
