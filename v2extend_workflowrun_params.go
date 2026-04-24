//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Workflow Runs.
type V2ExtendWorkflowRunListParams struct {
	Params `form:"*"`
	// Restrict page size to no more than this number.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// When retrieving Workflow Runs, include only those with the specified status values. If not specified, all Runs are returned.
	Status []*string `form:"status" json:"status,omitempty"`
	// When retrieving Workflow Runs, include only those associated with the Workflows specified. If not specified, all Runs are returned.
	Workflow []*string `form:"workflow" json:"workflow,omitempty"`
}

// Retrieves the details of a Workflow Run by ID.
type V2ExtendWorkflowRunParams struct {
	Params `form:"*"`
}

// Retrieves the details of a Workflow Run by ID.
type V2ExtendWorkflowRunRetrieveParams struct {
	Params `form:"*"`
}
