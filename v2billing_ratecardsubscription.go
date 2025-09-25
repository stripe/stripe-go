//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The payment status of a Rate Card Subscription.
type V2BillingRateCardSubscriptionCollectionStatus string

// List of values that V2BillingRateCardSubscriptionCollectionStatus can take
const (
	V2BillingRateCardSubscriptionCollectionStatusAwaitingCustomerAction V2BillingRateCardSubscriptionCollectionStatus = "awaiting_customer_action"
	V2BillingRateCardSubscriptionCollectionStatusCurrent                V2BillingRateCardSubscriptionCollectionStatus = "current"
	V2BillingRateCardSubscriptionCollectionStatusPastDue                V2BillingRateCardSubscriptionCollectionStatus = "past_due"
	V2BillingRateCardSubscriptionCollectionStatusPaused                 V2BillingRateCardSubscriptionCollectionStatus = "paused"
	V2BillingRateCardSubscriptionCollectionStatusUnpaid                 V2BillingRateCardSubscriptionCollectionStatus = "unpaid"
)

// The servicing status of a Rate Card Subscription.
type V2BillingRateCardSubscriptionServicingStatus string

// List of values that V2BillingRateCardSubscriptionServicingStatus can take
const (
	V2BillingRateCardSubscriptionServicingStatusActive   V2BillingRateCardSubscriptionServicingStatus = "active"
	V2BillingRateCardSubscriptionServicingStatusCanceled V2BillingRateCardSubscriptionServicingStatus = "canceled"
	V2BillingRateCardSubscriptionServicingStatusPaused   V2BillingRateCardSubscriptionServicingStatus = "paused"
	V2BillingRateCardSubscriptionServicingStatusPending  V2BillingRateCardSubscriptionServicingStatus = "pending"
)

// The collection status transitions of the Rate Card Subscription.
type V2BillingRateCardSubscriptionCollectionStatusTransitions struct {
	// When the collection status transitioned to awaiting customer action.
	AwaitingCustomerActionAt string `json:"awaiting_customer_action_at,omitempty"`
	// When the collection status transitioned to current.
	CurrentAt string `json:"current_at,omitempty"`
	// When the collection status transitioned to past due.
	PastDueAt string `json:"past_due_at,omitempty"`
	// When the collection status transitioned to paused.
	PausedAt string `json:"paused_at,omitempty"`
	// When the collection status transitioned to unpaid.
	UnpaidAt string `json:"unpaid_at,omitempty"`
}

// The servicing status transitions of the Rate Card Subscription.
type V2BillingRateCardSubscriptionServicingStatusTransitions struct {
	// When the servicing status transitioned to activated.
	ActivatedAt string `json:"activated_at,omitempty"`
	// When the servicing status transitioned to canceled.
	CanceledAt string `json:"canceled_at,omitempty"`
	// When the servicing status transitioned to paused.
	PausedAt string `json:"paused_at,omitempty"`
	// When the servicing is scheduled to transition to activate.
	WillActivateAt string `json:"will_activate_at,omitempty"`
	// When the servicing is scheduled to cancel.
	WillCancelAt string `json:"will_cancel_at,omitempty"`
}
type V2BillingRateCardSubscription struct {
	APIResource
	// The ID of the Billing Cadence.
	BillingCadence string `json:"billing_cadence"`
	// The payment status of a Rate Card Subscription.
	CollectionStatus V2BillingRateCardSubscriptionCollectionStatus `json:"collection_status,omitempty"`
	// The collection status transitions of the Rate Card Subscription.
	CollectionStatusTransitions *V2BillingRateCardSubscriptionCollectionStatusTransitions `json:"collection_status_transitions,omitempty"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the Rate Card.
	RateCard string `json:"rate_card"`
	// The ID of the Rate Card Version.
	RateCardVersion string `json:"rate_card_version"`
	// The servicing status of a Rate Card Subscription.
	ServicingStatus V2BillingRateCardSubscriptionServicingStatus `json:"servicing_status,omitempty"`
	// The servicing status transitions of the Rate Card Subscription.
	ServicingStatusTransitions *V2BillingRateCardSubscriptionServicingStatusTransitions `json:"servicing_status_transitions,omitempty"`
	// The ID of the Test Clock, if any.
	TestClock string `json:"test_clock,omitempty"`
}
