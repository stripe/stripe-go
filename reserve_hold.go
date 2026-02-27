//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Indicates which party created this ReserveHold.
type ReserveHoldCreatedBy string

// List of values that ReserveHoldCreatedBy can take
const (
	ReserveHoldCreatedByApplication ReserveHoldCreatedBy = "application"
	ReserveHoldCreatedByStripe      ReserveHoldCreatedBy = "stripe"
)

// The reason for the ReserveHold.
type ReserveHoldReason string

// List of values that ReserveHoldReason can take
const (
	ReserveHoldReasonCharge     ReserveHoldReason = "charge"
	ReserveHoldReasonStandalone ReserveHoldReason = "standalone"
)

// Which source balance type this ReserveHold reserves funds from. One of `bank_account`, `card`, or `fpx`.
type ReserveHoldSourceType string

// List of values that ReserveHoldSourceType can take
const (
	ReserveHoldSourceTypeBankAccount ReserveHoldSourceType = "bank_account"
	ReserveHoldSourceTypeCard        ReserveHoldSourceType = "card"
	ReserveHoldSourceTypeFPX         ReserveHoldSourceType = "fpx"
)

type ReserveHoldReleaseSchedule struct {
	// The time after which the ReserveHold is requested to be released.
	ReleaseAfter int64 `json:"release_after"`
	// The time at which the ReserveHold is scheduled to be released, automatically set to midnight UTC of the day after `release_after`.
	ScheduledRelease int64 `json:"scheduled_release"`
}

// ReserveHolds are used to place a temporary ReserveHold on a merchant's funds.
type ReserveHold struct {
	// Amount reserved. A positive integer representing how much is reserved in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount int64 `json:"amount"`
	// Amount in cents that can be released from this ReserveHold
	AmountReleasable int64 `json:"amount_releasable"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Indicates which party created this ReserveHold.
	CreatedBy ReserveHoldCreatedBy `json:"created_by"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Whether there are any funds available to release on this ReserveHold. Note that if the ReserveHold is in the process of being released, this could be false, even though the funds haven't been fully released yet.
	IsReleasable bool `json:"is_releasable"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The reason for the ReserveHold.
	Reason          ReserveHoldReason           `json:"reason"`
	ReleaseSchedule *ReserveHoldReleaseSchedule `json:"release_schedule"`
	// The ReservePlan which produced this ReserveHold (i.e., resplan_123)
	ReservePlan *ReservePlan `json:"reserve_plan"`
	// The Charge which funded this ReserveHold (e.g., ch_123)
	SourceCharge *Charge `json:"source_charge"`
	// Which source balance type this ReserveHold reserves funds from. One of `bank_account`, `card`, or `fpx`.
	SourceType ReserveHoldSourceType `json:"source_type"`
}

// UnmarshalJSON handles deserialization of a ReserveHold.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *ReserveHold) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		r.ID = id
		return nil
	}

	type reserveHold ReserveHold
	var v reserveHold
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*r = ReserveHold(v)
	return nil
}
