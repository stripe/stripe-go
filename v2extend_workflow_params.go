//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Workflows.
type V2ExtendWorkflowListParams struct {
	Params `form:"*"`
	// Restrict page size to no more than this number.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// When retrieving Workflows, include only those with the specified status values.
	// If not specified, all Workflows in active and inactive status are returned.
	Status []*string `form:"status" json:"status,omitempty"`
}

// Retrieves the details of a Workflow by ID.
type V2ExtendWorkflowParams struct {
	Params `form:"*"`
}

// Invokes an on-demand Workflow with parameters, to launch a new Workflow Run.
type V2ExtendWorkflowInvokeParams struct {
	Params `form:"*"`
	// Parameters used to invoke the Workflow Run.
	InputParameters map[string]any `form:"input_parameters" json:"input_parameters"`
}

// Retrieves the details of a Workflow by ID.
type V2ExtendWorkflowRetrieveParams struct {
	Params `form:"*"`
}
