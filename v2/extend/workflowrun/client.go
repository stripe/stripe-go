//
//
// File generated from our OpenAPI spec
//
//

// Package workflowrun provides the workflowrun related APIs
package workflowrun

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke workflowrun related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Workflow Run by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2ExtendWorkflowRunParams) (*stripe.V2ExtendWorkflowRun, error) {
	path := stripe.FormatURLPath("/v2/extend/workflow_runs/%s", id)
	workflowrun := &stripe.V2ExtendWorkflowRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, workflowrun)
	return workflowrun, err
}

// List all Workflow Runs.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2ExtendWorkflowRunListParams) stripe.Seq2[*stripe.V2ExtendWorkflowRun, error] {
	if listParams == nil {
		listParams = &stripe.V2ExtendWorkflowRunListParams{}
	}
	return stripe.NewV2List("/v2/extend/workflow_runs", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2ExtendWorkflowRun], error) {
		page := &stripe.V2Page[*stripe.V2ExtendWorkflowRun]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
