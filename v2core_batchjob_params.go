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
	Endpoint *V2BatchJobsEndpoint `form:"endpoint" json:"endpoint"`
}

// Serializes a Customer update request into a batch job JSONL line.
type V2CoreBatchJobSerializeV1CustomerUpdateParams struct {
	Params `form:"*"`
}

// Serializes a Subscription update request into a batch job JSONL line.
type V2CoreBatchJobSerializeV1SubscriptionUpdateParams struct {
	Params `form:"*"`
}

// Serializes a SubscriptionSchedule create request into a batch job JSONL line.
type V2CoreBatchJobSerializeV1SubscriptionScheduleCreateParams struct {
	Params `form:"*"`
}

// Creates a new batch job.
type V2CoreBatchJobCreateParams struct {
	Params `form:"*"`
	// The API endpoint to batch (e.g., /v1/customers/:id for batch customer updates).
	Endpoint *V2BatchJobsEndpoint `form:"endpoint" json:"endpoint"`
}

// V2BatchJobsEndpoint represents a batchable API endpoint for batch jobs.
type V2BatchJobsEndpoint string

// List of values that V2BatchJobsEndpoint can take.
const (
	V2BatchJobsEndpointCustomerUpdate             V2BatchJobsEndpoint = "/v1/customers/:customer"
	V2BatchJobsEndpointSubscriptionUpdate         V2BatchJobsEndpoint = "/v1/subscriptions/:subscription_exposed_id"
	V2BatchJobsEndpointSubscriptionScheduleCreate V2BatchJobsEndpoint = "/v1/subscription_schedules"
)

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

// batchItemParamsContainer is an internal interface for extracting batch job fields from params.
type batchItemParamsContainer interface {
	GetStripeVersion() *string
	GetStripeContext() *string
}
