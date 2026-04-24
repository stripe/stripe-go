//
//
// File generated from our OpenAPI spec
//
//

// Package workflow provides the workflow related APIs
package workflow

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke workflow related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Workflow by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2ExtendWorkflowParams) (*stripe.V2ExtendWorkflow, error) {
	path := stripe.FormatURLPath("/v2/extend/workflows/%s", id)
	workflow := &stripe.V2ExtendWorkflow{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, workflow)
	return workflow, err
}

// Invokes an on-demand Workflow with parameters, to launch a new Workflow Run.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Invoke(id string, params *stripe.V2ExtendWorkflowInvokeParams) (*stripe.V2ExtendWorkflowRun, error) {
	path := stripe.FormatURLPath("/v2/extend/workflows/%s/invoke", id)
	workflowrun := &stripe.V2ExtendWorkflowRun{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, workflowrun)
	return workflowrun, err
}

// List all Workflows.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2ExtendWorkflowListParams) stripe.Seq2[*stripe.V2ExtendWorkflow, error] {
	if listParams == nil {
		listParams = &stripe.V2ExtendWorkflowListParams{}
	}
	return stripe.NewV2List("/v2/extend/workflows", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2ExtendWorkflow], error) {
		page := &stripe.V2Page[*stripe.V2ExtendWorkflow]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
