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

// v2ExtendWorkflowService is used to invoke workflow related APIs.
type v2ExtendWorkflowService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Workflow by ID.
func (c v2ExtendWorkflowService) Retrieve(ctx context.Context, id string, params *V2ExtendWorkflowRetrieveParams) (*V2ExtendWorkflow, error) {
	if params == nil {
		params = &V2ExtendWorkflowRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/extend/workflows/%s", id)
	workflow := &V2ExtendWorkflow{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, workflow)
	return workflow, err
}

// Invokes an on-demand Workflow with parameters, to launch a new Workflow Run.
func (c v2ExtendWorkflowService) Invoke(ctx context.Context, id string, params *V2ExtendWorkflowInvokeParams) (*V2ExtendWorkflowRun, error) {
	if params == nil {
		params = &V2ExtendWorkflowInvokeParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/extend/workflows/%s/invoke", id)
	workflowrun := &V2ExtendWorkflowRun{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, workflowrun)
	return workflowrun, err
}

// List all Workflows.
func (c v2ExtendWorkflowService) List(ctx context.Context, listParams *V2ExtendWorkflowListParams) *V2List[*V2ExtendWorkflow] {
	if listParams == nil {
		listParams = &V2ExtendWorkflowListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/extend/workflows", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2ExtendWorkflow], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2ExtendWorkflow]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
