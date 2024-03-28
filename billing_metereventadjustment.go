//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The meter event adjustment's status.
type BillingMeterEventAdjustmentStatus string

// List of values that BillingMeterEventAdjustmentStatus can take
const (
	BillingMeterEventAdjustmentStatusComplete BillingMeterEventAdjustmentStatus = "complete"
	BillingMeterEventAdjustmentStatusPending  BillingMeterEventAdjustmentStatus = "pending"
)

// Specifies which event to cancel.
type BillingMeterEventAdjustmentCancelParams struct {
	// Unique identifier for the event.
	Identifier *string `form:"identifier"`
}

// Creates a billing meter event adjustment
type BillingMeterEventAdjustmentParams struct {
	Params `form:"*"`
	// Specifies which event to cancel.
	Cancel *BillingMeterEventAdjustmentCancelParams `form:"cancel"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Specifies whether to cancel a single event or a range of events for a time period.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *BillingMeterEventAdjustmentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A billing meter event adjustment represents the status of a meter event adjustment.
type BillingMeterEventAdjustment struct {
	APIResource
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The meter event adjustment's status.
	Status BillingMeterEventAdjustmentStatus `json:"status"`
}
