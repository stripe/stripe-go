//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Indicates which party created this ReservePlan.
type ReservePlanCreatedBy string

// List of values that ReservePlanCreatedBy can take
const (
	ReservePlanCreatedByApplication ReservePlanCreatedBy = "application"
	ReservePlanCreatedByStripe      ReservePlanCreatedBy = "stripe"
)

// The current status of the ReservePlan. The ReservePlan only affects charges if it is `active`.
type ReservePlanStatus string

// List of values that ReservePlanStatus can take
const (
	ReservePlanStatusActive   ReservePlanStatus = "active"
	ReservePlanStatusDisabled ReservePlanStatus = "disabled"
	ReservePlanStatusExpired  ReservePlanStatus = "expired"
)

// The type of the ReservePlan.
type ReservePlanType string

// List of values that ReservePlanType can take
const (
	ReservePlanTypeFixedRelease   ReservePlanType = "fixed_release"
	ReservePlanTypeRollingRelease ReservePlanType = "rolling_release"
)

type ReservePlanFixedRelease struct {
	// The time after which all reserved funds are requested for release.
	ReleaseAfter int64 `json:"release_after"`
	// The time at which reserved funds are scheduled for release, automatically set to midnight UTC of the day after `release_after`.
	ScheduledRelease int64 `json:"scheduled_release"`
}
type ReservePlanRollingRelease struct {
	// The number of days to reserve funds before releasing.
	DaysAfterCharge int64 `json:"days_after_charge"`
	// The time at which the ReservePlan expires.
	ExpiresOn int64 `json:"expires_on"`
}

// ReservePlans are used to automatically place holds on a merchant's funds until the plan expires. It takes a portion of each incoming Charge (including those resulting from a Transfer from a platform account).
type ReservePlan struct {
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Indicates which party created this ReservePlan.
	CreatedBy ReservePlanCreatedBy `json:"created_by"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies). An unset currency indicates that the plan applies to all currencies.
	Currency Currency `json:"currency"`
	// Time at which the ReservePlan was disabled.
	DisabledAt   int64                    `json:"disabled_at"`
	FixedRelease *ReservePlanFixedRelease `json:"fixed_release"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The percent of each Charge to reserve.
	Percent        int64                      `json:"percent"`
	RollingRelease *ReservePlanRollingRelease `json:"rolling_release"`
	// The current status of the ReservePlan. The ReservePlan only affects charges if it is `active`.
	Status ReservePlanStatus `json:"status"`
	// The type of the ReservePlan.
	Type ReservePlanType `json:"type"`
}

// UnmarshalJSON handles deserialization of a ReservePlan.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *ReservePlan) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		r.ID = id
		return nil
	}

	type reservePlan ReservePlan
	var v reservePlan
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*r = ReservePlan(v)
	return nil
}
