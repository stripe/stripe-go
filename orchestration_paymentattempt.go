//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves orchestration information for the given payment attempt record (e.g. return url).
type OrchestrationPaymentAttemptParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *OrchestrationPaymentAttemptParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves orchestration information for the given payment attempt record (e.g. return url).
type OrchestrationPaymentAttemptRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *OrchestrationPaymentAttemptRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Represents orchestration information for a payment attempt record (e.g. return url).
type OrchestrationPaymentAttempt struct {
	APIResource
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// If present, the return URL for this payment attempt.
	ReturnURL string `json:"return_url"`
}
