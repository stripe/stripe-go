//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Current collection status of this subscription.
type V2BillingPricingPlanSubscriptionCollectionStatus string

// List of values that V2BillingPricingPlanSubscriptionCollectionStatus can take
const (
	V2BillingPricingPlanSubscriptionCollectionStatusAwaitingCustomerAction V2BillingPricingPlanSubscriptionCollectionStatus = "awaiting_customer_action"
	V2BillingPricingPlanSubscriptionCollectionStatusCurrent                V2BillingPricingPlanSubscriptionCollectionStatus = "current"
	V2BillingPricingPlanSubscriptionCollectionStatusPastDue                V2BillingPricingPlanSubscriptionCollectionStatus = "past_due"
	V2BillingPricingPlanSubscriptionCollectionStatusPaused                 V2BillingPricingPlanSubscriptionCollectionStatus = "paused"
	V2BillingPricingPlanSubscriptionCollectionStatusUnpaid                 V2BillingPricingPlanSubscriptionCollectionStatus = "unpaid"
)

// Current servicing status of this subscription.
type V2BillingPricingPlanSubscriptionServicingStatus string

// List of values that V2BillingPricingPlanSubscriptionServicingStatus can take
const (
	V2BillingPricingPlanSubscriptionServicingStatusActive   V2BillingPricingPlanSubscriptionServicingStatus = "active"
	V2BillingPricingPlanSubscriptionServicingStatusCanceled V2BillingPricingPlanSubscriptionServicingStatus = "canceled"
	V2BillingPricingPlanSubscriptionServicingStatusPaused   V2BillingPricingPlanSubscriptionServicingStatus = "paused"
	V2BillingPricingPlanSubscriptionServicingStatusPending  V2BillingPricingPlanSubscriptionServicingStatus = "pending"
)

// Timestamps for collection status transitions.
type V2BillingPricingPlanSubscriptionCollectionStatusTransitions struct {
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

// Timestamps for servicing status transitions.
type V2BillingPricingPlanSubscriptionServicingStatusTransitions struct {
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
type V2BillingPricingPlanSubscription struct {
	APIResource
	// The ID of the Billing Cadence this subscription is billed on.
	BillingCadence string `json:"billing_cadence"`
	// Current collection status of this subscription.
	CollectionStatus V2BillingPricingPlanSubscriptionCollectionStatus `json:"collection_status"`
	// Timestamps for collection status transitions.
	CollectionStatusTransitions *V2BillingPricingPlanSubscriptionCollectionStatusTransitions `json:"collection_status_transitions"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the Pricing Plan for this subscription.
	PricingPlan string `json:"pricing_plan"`
	// The ID of the Pricing Plan Version for this subscription.
	PricingPlanVersion string `json:"pricing_plan_version"`
	// Current servicing status of this subscription.
	ServicingStatus V2BillingPricingPlanSubscriptionServicingStatus `json:"servicing_status"`
	// Timestamps for servicing status transitions.
	ServicingStatusTransitions *V2BillingPricingPlanSubscriptionServicingStatusTransitions `json:"servicing_status_transitions"`
	// The ID of the Test Clock of the associated Billing Cadence, if any.
	TestClock string `json:"test_clock,omitempty"`
}
