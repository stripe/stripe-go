//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// When the BillingIntent will take effect.
type V2BillingIntentEffectiveAt string

// List of values that V2BillingIntentEffectiveAt can take
const (
	V2BillingIntentEffectiveAtCurrentBillingPeriodStart V2BillingIntentEffectiveAt = "current_billing_period_start"
	V2BillingIntentEffectiveAtOnCommit                  V2BillingIntentEffectiveAt = "on_commit"
	V2BillingIntentEffectiveAtOnReserve                 V2BillingIntentEffectiveAt = "on_reserve"
)

// Current status of the BillingIntent.
type V2BillingIntentStatus string

// List of values that V2BillingIntentStatus can take
const (
	V2BillingIntentStatusCanceled  V2BillingIntentStatus = "canceled"
	V2BillingIntentStatusCommitted V2BillingIntentStatus = "committed"
	V2BillingIntentStatusDraft     V2BillingIntentStatus = "draft"
	V2BillingIntentStatusReserved  V2BillingIntentStatus = "reserved"
)

// Breakdown of the amount for this BillingIntent.
type V2BillingIntentAmountDetails struct {
	// Three-letter ISO currency code, in lowercase.
	Currency Currency `json:"currency"`
	// Amount of discount applied.
	Discount string `json:"discount"`
	// Amount of shipping charges.
	Shipping string `json:"shipping"`
	// Subtotal amount before tax and discounts.
	Subtotal string `json:"subtotal"`
	// Amount of tax.
	Tax string `json:"tax"`
	// Total amount for the BillingIntent.
	Total string `json:"total"`
}

// Timestamps for status transitions of the BillingIntent.
type V2BillingIntentStatusTransitions struct {
	// Time at which the BillingIntent was canceled.
	CanceledAt time.Time `json:"canceled_at"`
	// Time at which the BillingIntent was committed.
	CommittedAt time.Time `json:"committed_at"`
	// Time at which the BillingIntent was drafted.
	DraftedAt time.Time `json:"drafted_at"`
	// Time at which the BillingIntent was reserved.
	ReservedAt time.Time `json:"reserved_at"`
}
type V2BillingIntent struct {
	APIResource
	// Breakdown of the amount for this BillingIntent.
	AmountDetails *V2BillingIntentAmountDetails `json:"amount_details"`
	// ID of an existing Cadence to use.
	Cadence string `json:"cadence"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Three-letter ISO currency code, in lowercase.
	Currency Currency `json:"currency"`
	// When the BillingIntent will take effect.
	EffectiveAt V2BillingIntentEffectiveAt `json:"effective_at"`
	// Unique identifier for the BillingIntent.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Current status of the BillingIntent.
	Status V2BillingIntentStatus `json:"status"`
	// Timestamps for status transitions of the BillingIntent.
	StatusTransitions *V2BillingIntentStatusTransitions `json:"status_transitions"`
}
