//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Whether this Workflow is active, inactive, or in some other state. Only active Workflows may be invoked.
type V2ExtendWorkflowStatus string

// List of values that V2ExtendWorkflowStatus can take
const (
	V2ExtendWorkflowStatusActive   V2ExtendWorkflowStatus = "active"
	V2ExtendWorkflowStatusArchived V2ExtendWorkflowStatus = "archived"
	V2ExtendWorkflowStatusDraft    V2ExtendWorkflowStatus = "draft"
	V2ExtendWorkflowStatusInactive V2ExtendWorkflowStatus = "inactive"
)

// Which type of trigger this is.
type V2ExtendWorkflowTriggerType string

// List of values that V2ExtendWorkflowTriggerType can take
const (
	V2ExtendWorkflowTriggerTypeEventTrigger V2ExtendWorkflowTriggerType = "event_trigger"
	V2ExtendWorkflowTriggerTypeManual       V2ExtendWorkflowTriggerType = "manual"
)

// The Workflow can be launched when Stripe emits a certain event.
type V2ExtendWorkflowTriggerEventTrigger struct {
	// Specifies which accounts' events will trigger this Workflow.
	EventsFrom []string `json:"events_from"`
	// The Stripe event type that will trigger this Workflow.
	Type string `json:"type"`
}

// The Workflow can be launched through a direct call, using either the Dashboard or the Stripe API.
type V2ExtendWorkflowTriggerManual struct{}

// Under what conditions will this Workflow launch.
type V2ExtendWorkflowTrigger struct {
	// The Workflow can be launched when Stripe emits a certain event.
	EventTrigger *V2ExtendWorkflowTriggerEventTrigger `json:"event_trigger,omitempty"`
	// The Workflow can be launched through a direct call, using either the Dashboard or the Stripe API.
	Manual *V2ExtendWorkflowTriggerManual `json:"manual,omitempty"`
	// Which type of trigger this is.
	Type V2ExtendWorkflowTriggerType `json:"type"`
}

// A Stripe Workflow is a sequence of actions, like Stripe API calls, that are taken in response to an initiating trigger.
// A trigger can be a Stripe API event, or a manual invocation.
type V2ExtendWorkflow struct {
	APIResource
	// When the Workflow was created.
	Created time.Time `json:"created"`
	// The unique ID of the Workflow.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Whether this Workflow is active, inactive, or in some other state. Only active Workflows may be invoked.
	Status V2ExtendWorkflowStatus `json:"status"`
	// Workflow title.
	Title string `json:"title"`
	// Under what conditions will this Workflow launch.
	Triggers []*V2ExtendWorkflowTrigger `json:"triggers"`
}
