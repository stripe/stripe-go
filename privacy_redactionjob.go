//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List redaction jobs method...
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

// The objects at the root level that are subject to redaction.
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

// Create redaction job method
type PrivacyRedactionJobParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The objects at the root level that are subject to redaction.
	Objects *PrivacyRedactionJobObjectsParams `form:"objects"`
	// Default is "error". If "error", we will make sure all objects in the graph are
	// redactable in the 1st traversal, otherwise error. If "fix", where possible, we will
	// auto-fix any validation errors (e.g. by auto-transitioning objects to a terminal
	// state, etc.) in the 2nd traversal before redacting
	ValidationBehavior *string `form:"validation_behavior"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Cancel redaction job method
type PrivacyRedactionJobCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Run redaction job method
type PrivacyRedactionJobRunParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobRunParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Validate redaction job method
type PrivacyRedactionJobValidateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobValidateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The objects at the root level that are subject to redaction.
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

// Create redaction job method
type PrivacyRedactionJobCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The objects at the root level that are subject to redaction.
	Objects *PrivacyRedactionJobCreateObjectsParams `form:"objects"`
	// Default is "error". If "error", we will make sure all objects in the graph are
	// redactable in the 1st traversal, otherwise error. If "fix", where possible, we will
	// auto-fix any validation errors (e.g. by auto-transitioning objects to a terminal
	// state, etc.) in the 2nd traversal before redacting
	ValidationBehavior *string `form:"validation_behavior"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieve redaction job method
type PrivacyRedactionJobRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Update redaction job method
type PrivacyRedactionJobUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand             []*string `form:"expand"`
	ValidationBehavior *string   `form:"validation_behavior"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Redaction Jobs store the status of a redaction request. They are created
// when a redaction request is made and track the redaction validation and execution.
type PrivacyRedactionJob struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The objects at the root level that are subject to redaction.
	Objects *PrivacyRedactionJobRootObjects `json:"objects"`
	// The status field represents the current state of the redaction job. It can take on any of the following values: VALIDATING, READY, REDACTING, SUCCEEDED, CANCELED, FAILED.
	Status string `json:"status"`
	// Default is "error". If "error", we will make sure all objects in the graph are redactable in the 1st traversal, otherwise error. If "fix", where possible, we will auto-fix any validation errors (e.g. by auto-transitioning objects to a terminal state, etc.) in the 2nd traversal before redacting
	ValidationBehavior string `json:"validation_behavior"`
}

// PrivacyRedactionJobList is a list of RedactionJobs as retrieved from a list endpoint.
type PrivacyRedactionJobList struct {
	APIResource
	ListMeta
	Data []*PrivacyRedactionJob `json:"data"`
}
