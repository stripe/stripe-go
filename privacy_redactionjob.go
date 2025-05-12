//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The status of the job.
type PrivacyRedactionJobStatus string

// List of values that PrivacyRedactionJobStatus can take
const (
	PrivacyRedactionJobStatusCanceled   PrivacyRedactionJobStatus = "canceled"
	PrivacyRedactionJobStatusCanceling  PrivacyRedactionJobStatus = "canceling"
	PrivacyRedactionJobStatusCreated    PrivacyRedactionJobStatus = "created"
	PrivacyRedactionJobStatusFailed     PrivacyRedactionJobStatus = "failed"
	PrivacyRedactionJobStatusReady      PrivacyRedactionJobStatus = "ready"
	PrivacyRedactionJobStatusRedacting  PrivacyRedactionJobStatus = "redacting"
	PrivacyRedactionJobStatusSucceeded  PrivacyRedactionJobStatus = "succeeded"
	PrivacyRedactionJobStatusValidating PrivacyRedactionJobStatus = "validating"
)

// Validation behavior determines how a job validates objects for redaction eligibility. Default is `error`.
type PrivacyRedactionJobValidationBehavior string

// List of values that PrivacyRedactionJobValidationBehavior can take
const (
	PrivacyRedactionJobValidationBehaviorError PrivacyRedactionJobValidationBehavior = "error"
	PrivacyRedactionJobValidationBehaviorFix   PrivacyRedactionJobValidationBehavior = "fix"
)

// Returns a list of redaction jobs.
type PrivacyRedactionJobListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	Status *string   `form:"status"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The objects to redact. These root objects and their related ones will be validated for redaction.
type PrivacyRedactionJobObjectsParams struct {
	Charges                      []*string `form:"charges"`
	CheckoutSessions             []*string `form:"checkout_sessions"`
	Customers                    []*string `form:"customers"`
	IdentityVerificationSessions []*string `form:"identity_verification_sessions"`
	Invoices                     []*string `form:"invoices"`
	IssuingCardholders           []*string `form:"issuing_cardholders"`
	IssuingCards                 []*string `form:"issuing_cards"`
	PaymentIntents               []*string `form:"payment_intents"`
	RadarValueListItems          []*string `form:"radar_value_list_items"`
	SetupIntents                 []*string `form:"setup_intents"`
}

// Creates a redaction job. When a job is created, it will start to validate.
type PrivacyRedactionJobParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The objects to redact. These root objects and their related ones will be validated for redaction.
	Objects *PrivacyRedactionJobObjectsParams `form:"objects"`
	// Determines the validation behavior of the job. Default is `error`.
	ValidationBehavior *string `form:"validation_behavior"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// You can cancel a redaction job when it's in one of these statuses: ready, failed.
//
// Canceling the redaction job will abandon its attempt to redact the configured objects. A canceled job cannot be used again.
type PrivacyRedactionJobCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Run a redaction job in a ready status.
//
// When you run a job, the configured objects will be redacted asynchronously. This action is irreversible and cannot be canceled once started.
//
// The status of the job will move to redacting. Once all of the objects are redacted, the status will become succeeded. If the job's validation_behavior is set to fix, the automatic fixes will be applied to objects at this step.
type PrivacyRedactionJobRunParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobRunParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Validate a redaction job when it is in a failed status.
//
// When a job is created, it automatically begins to validate on the configured objects' eligibility for redaction. Use this to validate the job again after its validation errors are resolved or the job's validation_behavior is changed.
//
// The status of the job will move to validating. Once all of the objects are validated, the status of the job will become ready. If there are any validation errors preventing the job from running, the status will become failed.
type PrivacyRedactionJobValidateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobValidateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The objects to redact. These root objects and their related ones will be validated for redaction.
type PrivacyRedactionJobCreateObjectsParams struct {
	Charges                      []*string `form:"charges"`
	CheckoutSessions             []*string `form:"checkout_sessions"`
	Customers                    []*string `form:"customers"`
	IdentityVerificationSessions []*string `form:"identity_verification_sessions"`
	Invoices                     []*string `form:"invoices"`
	IssuingCardholders           []*string `form:"issuing_cardholders"`
	IssuingCards                 []*string `form:"issuing_cards"`
	PaymentIntents               []*string `form:"payment_intents"`
	RadarValueListItems          []*string `form:"radar_value_list_items"`
	SetupIntents                 []*string `form:"setup_intents"`
}

// Creates a redaction job. When a job is created, it will start to validate.
type PrivacyRedactionJobCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The objects to redact. These root objects and their related ones will be validated for redaction.
	Objects *PrivacyRedactionJobCreateObjectsParams `form:"objects"`
	// Determines the validation behavior of the job. Default is `error`.
	ValidationBehavior *string `form:"validation_behavior"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a previously created redaction job.
type PrivacyRedactionJobRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates the properties of a redaction job without running or canceling the job.
//
// If the job to update is in a failed status, it will not automatically start to validate. Once you applied all of the changes, use the validate API to start validation again.
type PrivacyRedactionJobUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Determines the validation behavior of the job. Default is `error`.
	ValidationBehavior *string `form:"validation_behavior"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The objects to redact in this job.
type PrivacyRedactionJobObjects struct {
	// Charge object identifiers usually starting with `ch_`
	Charges []string `json:"charges"`
	// CheckoutSession object identifiers starting with `cs_`
	CheckoutSessions []string `json:"checkout_sessions"`
	// Customer object identifiers starting with `cus_`
	Customers []string `json:"customers"`
	// Identity VerificationSessions object identifiers starting with `vs_`
	IdentityVerificationSessions []string `json:"identity_verification_sessions"`
	// Invoice object identifiers starting with `in_`
	Invoices []string `json:"invoices"`
	// Issuing Cardholder object identifiers starting with `ich_`
	IssuingCardholders []string `json:"issuing_cardholders"`
	// PaymentIntent object identifiers starting with `pi_`
	PaymentIntents []string `json:"payment_intents"`
	// Fraud ValueListItem object identifiers starting with `rsli_`
	RadarValueListItems []string `json:"radar_value_list_items"`
	// SetupIntent object identifiers starting with `seti_`
	SetupIntents []string `json:"setup_intents"`
}

// The Redaction Job object is used to redact Stripe objects. It is used
// to coordinate the removal of personal information from selected
// objects, making them permanently inaccessible in the Stripe Dashboard
// and API.
type PrivacyRedactionJob struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The objects to redact in this job.
	Objects *PrivacyRedactionJobObjects `json:"objects"`
	// The status of the job.
	Status PrivacyRedactionJobStatus `json:"status"`
	// Validation behavior determines how a job validates objects for redaction eligibility. Default is `error`.
	ValidationBehavior PrivacyRedactionJobValidationBehavior `json:"validation_behavior"`
}

// PrivacyRedactionJobList is a list of RedactionJobs as retrieved from a list endpoint.
type PrivacyRedactionJobList struct {
	APIResource
	ListMeta
	Data []*PrivacyRedactionJob `json:"data"`
}
