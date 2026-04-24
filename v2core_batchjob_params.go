//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The endpoint configuration for the batch job.
type V2CoreBatchJobEndpointParams struct {
	// The HTTP method to use when calling the endpoint.
	HTTPMethod *string `form:"http_method" json:"http_method"`
	// The path of the endpoint to run this batch job against.
	// In the form used in the documentation. For instance, for
	// subscription migration this would be `/v1/subscriptions/:id/migrate`.
	Path *V2CoreBatchJobsPath `form:"path" json:"path"`
}

// Notification suppression settings for the batch job.
type V2CoreBatchJobNotificationSuppressionParams struct {
	// The scope of notification suppression.
	Scope *string `form:"scope" json:"scope"`
}

// Creates a new batch job.
type V2CoreBatchJobParams struct {
	Params `form:"*"`
	// The endpoint configuration for the batch job.
	Endpoint *V2CoreBatchJobEndpointParams `form:"endpoint" json:"endpoint,omitempty"`
	// Optional field that allows the user to control how fast they want this batch job to run.
	// Gives them a control over the number of webhooks they receive.
	MaximumRps *int64 `form:"maximum_rps" json:"maximum_rps,omitempty"`
	// The metadata of the `batch_job`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Notification suppression settings for the batch job.
	NotificationSuppression *V2CoreBatchJobNotificationSuppressionParams `form:"notification_suppression" json:"notification_suppression,omitempty"`
	// Allows the user to skip validation.
	SkipValidation *bool `form:"skip_validation" json:"skip_validation,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreBatchJobParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancels an existing batch job.
type V2CoreBatchJobCancelParams struct {
	Params `form:"*"`
}

// The endpoint configuration for the batch job.
type V2CoreBatchJobCreateEndpointParams struct {
	// The HTTP method to use when calling the endpoint.
	HTTPMethod *string `form:"http_method" json:"http_method"`
	// The path of the endpoint to run this batch job against.
	// In the form used in the documentation. For instance, for
	// subscription migration this would be `/v1/subscriptions/:id/migrate`.
	Path *V2CoreBatchJobsPath `form:"path" json:"path"`
}

// Notification suppression settings for the batch job.
type V2CoreBatchJobCreateNotificationSuppressionParams struct {
	// The scope of notification suppression.
	Scope *string `form:"scope" json:"scope"`
}

// Creates a new batch job.
type V2CoreBatchJobCreateParams struct {
	Params `form:"*"`
	// The endpoint configuration for the batch job.
	Endpoint *V2CoreBatchJobCreateEndpointParams `form:"endpoint" json:"endpoint"`
	// Optional field that allows the user to control how fast they want this batch job to run.
	// Gives them a control over the number of webhooks they receive.
	MaximumRps *int64 `form:"maximum_rps" json:"maximum_rps,omitempty"`
	// The metadata of the `batch_job`.
	Metadata map[string]string `form:"metadata" json:"metadata"`
	// Notification suppression settings for the batch job.
	NotificationSuppression *V2CoreBatchJobCreateNotificationSuppressionParams `form:"notification_suppression" json:"notification_suppression,omitempty"`
	// Allows the user to skip validation.
	SkipValidation *bool `form:"skip_validation" json:"skip_validation"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreBatchJobCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves an existing batch job.
type V2CoreBatchJobRetrieveParams struct {
	Params `form:"*"`
}

// V2CoreBatchJobsPath represents a batchable API endpoint for batch jobs.
type V2CoreBatchJobsPath string

// List of values that V2CoreBatchJobsPath can take.
const (
	V2CoreBatchJobsPathAccountUpdate              V2CoreBatchJobsPath = "/v1/accounts/:account"
	V2CoreBatchJobsPathCreditNoteCreate           V2CoreBatchJobsPath = "/v1/credit_notes"
	V2CoreBatchJobsPathCustomerUpdate             V2CoreBatchJobsPath = "/v1/customers/:customer"
	V2CoreBatchJobsPathInvoiceUpdate              V2CoreBatchJobsPath = "/v1/invoices/:invoice"
	V2CoreBatchJobsPathInvoicePay                 V2CoreBatchJobsPath = "/v1/invoices/:invoice/pay"
	V2CoreBatchJobsPathPromotionCodeCreate        V2CoreBatchJobsPath = "/v1/promotion_codes"
	V2CoreBatchJobsPathPromotionCodeUpdate        V2CoreBatchJobsPath = "/v1/promotion_codes/:promotion_code"
	V2CoreBatchJobsPathSubscriptionUpdate         V2CoreBatchJobsPath = "/v1/subscriptions/:subscription_exposed_id"
	V2CoreBatchJobsPathSubscriptionMigrate        V2CoreBatchJobsPath = "/v1/subscriptions/:subscription/migrate"
	V2CoreBatchJobsPathSubscriptionScheduleCreate V2CoreBatchJobsPath = "/v1/subscription_schedules"
	V2CoreBatchJobsPathSubscriptionScheduleUpdate V2CoreBatchJobsPath = "/v1/subscription_schedules/:schedule"
	V2CoreBatchJobsPathSubscriptionScheduleCancel V2CoreBatchJobsPath = "/v1/subscription_schedules/:schedule/cancel"
)
