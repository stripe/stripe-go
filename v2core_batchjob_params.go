//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a new batch job.
type V2CoreBatchJobParams struct {
	Params `form:"*"`
	// The API endpoint to batch (e.g., /v1/customers/:id for batch customer updates).
	Endpoint *string `form:"endpoint" json:"endpoint"`
}

// Creates a new batch job.
type V2CoreBatchJobCreateParams struct {
	Params `form:"*"`
	// The API endpoint to batch (e.g., /v1/customers/:id for batch customer updates).
	Endpoint *string `form:"endpoint" json:"endpoint"`
}

// BatchItemParams represents a single item in a batch job JSONL file.
type BatchItemParams struct {
	// Unique identifier for this batch item.
	ID string `json:"id"`
	// Optional Stripe-Context header value.
	Context string `json:"context,omitempty"`
	// Optional Stripe-Version header value.
	StripeVersion string `json:"stripe_version,omitempty"`
	// Map of path parameter names to values.
	PathParams map[string]string `json:"path_params,omitempty"`
	// The params object for the API call.
	Params interface{} `json:"params"`
}
