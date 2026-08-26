//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The feedback option's status.
type BillingFeedbackOptionStatus string

// List of values that BillingFeedbackOptionStatus can take
const (
	BillingFeedbackOptionStatusActive   BillingFeedbackOptionStatus = "active"
	BillingFeedbackOptionStatusInactive BillingFeedbackOptionStatus = "inactive"
)

// An API method for listing the feedback options model
type BillingFeedbackOptionListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Filter results to only include feedback options with the given status.
	Status *string `form:"status" json:"status,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingFeedbackOptionListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a new feedback option.
type BillingFeedbackOptionParams struct {
	Params      `form:"*"`
	Description *string `form:"description" json:"description,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingFeedbackOptionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Deactivates a feedback option. Deactivated feedback options cannot be used in portal configurations.
type BillingFeedbackOptionDeactivateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingFeedbackOptionDeactivateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a new feedback option.
type BillingFeedbackOptionCreateParams struct {
	Params      `form:"*"`
	Description *string `form:"description" json:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingFeedbackOptionCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a feedback options object given an ID.
type BillingFeedbackOptionRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingFeedbackOptionRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates the description of an existing feedback option.
type BillingFeedbackOptionUpdateParams struct {
	Params      `form:"*"`
	Description *string `form:"description" json:"description,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingFeedbackOptionUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type BillingFeedbackOptionStatusTransitions struct {
	// The time the feedback option was deactivated, if any. Measured in seconds since Unix epoch.
	DeactivatedAt int64 `json:"deactivated_at"`
}

// A resource for the feedback options model (for custom cancellation reasons)
type BillingFeedbackOption struct {
	APIResource
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The feedback option's status.
	Status            BillingFeedbackOptionStatus             `json:"status"`
	StatusTransitions *BillingFeedbackOptionStatusTransitions `json:"status_transitions"`
}

// BillingFeedbackOptionList is a list of FeedbackOptions as retrieved from a list endpoint.
type BillingFeedbackOptionList struct {
	APIResource
	ListMeta
	Data []*BillingFeedbackOption `json:"data"`
}

// UnmarshalJSON handles deserialization of a BillingFeedbackOption.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BillingFeedbackOption) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type billingFeedbackOption BillingFeedbackOption
	var v billingFeedbackOption
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BillingFeedbackOption(v)
	return nil
}
