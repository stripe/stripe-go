//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Current status of the Billing Intent.
type V2BillingIntentStatus string

// List of values that V2BillingIntentStatus can take
const (
	V2BillingIntentStatusCanceled  V2BillingIntentStatus = "canceled"
	V2BillingIntentStatusCommitted V2BillingIntentStatus = "committed"
	V2BillingIntentStatusDraft     V2BillingIntentStatus = "draft"
	V2BillingIntentStatusReserved  V2BillingIntentStatus = "reserved"
)

// Breakdown of the amount for this Billing Intent.
type V2BillingIntentAmountDetails struct {
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency Currency `json:"currency"`
	// Amount of discount applied.
	Discount string `json:"discount"`
	// Amount of shipping charges.
	Shipping string `json:"shipping"`
	// Subtotal amount before tax and discounts.
	Subtotal string `json:"subtotal"`
	// Amount of tax.
	Tax string `json:"tax"`
	// Total amount for the Billing Intent.
	Total string `json:"total"`
}

// Timestamps for status transitions of the Billing Intent.
type V2BillingIntentStatusTransitions struct {
	// Time at which the Billing Intent was canceled.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Time at which the Billing Intent was committed.
	CommittedAt time.Time `json:"committed_at,omitempty"`
	// Time at which the Billing Intent was drafted.
	DraftedAt time.Time `json:"drafted_at,omitempty"`
	// Time at which the Billing Intent was reserved.
	ReservedAt time.Time `json:"reserved_at,omitempty"`
}
type V2BillingIntent struct {
	APIResource
	// Breakdown of the amount for this Billing Intent.
	AmountDetails *V2BillingIntentAmountDetails `json:"amount_details"`
	// ID of an existing Cadence to use.
	Cadence string `json:"cadence,omitempty"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Current status of the Billing Intent.
	Status V2BillingIntentStatus `json:"status"`
	// Timestamps for status transitions of the Billing Intent.
	StatusTransitions *V2BillingIntentStatusTransitions `json:"status_transitions"`
}
