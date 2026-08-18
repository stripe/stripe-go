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

type BillingFeedbackOptionStatusTransitions struct {
	// The time the feedback option was deactivated, if any. Measured in seconds since Unix epoch.
	DeactivatedAt int64 `json:"deactivated_at"`
}

// A resource for the feedback options model (for custom cancellation reasons)
type BillingFeedbackOption struct {
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
