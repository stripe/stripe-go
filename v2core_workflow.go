//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Whether this Workflow is active, inactive, or in some other state. Only active Workflows may be invoked.
type V2CoreWorkflowStatus string

// List of values that V2CoreWorkflowStatus can take
const (
	V2CoreWorkflowStatusActive   V2CoreWorkflowStatus = "active"
	V2CoreWorkflowStatusArchived V2CoreWorkflowStatus = "archived"
	V2CoreWorkflowStatusDraft    V2CoreWorkflowStatus = "draft"
	V2CoreWorkflowStatusInactive V2CoreWorkflowStatus = "inactive"
)

// Which type of trigger this is.
type V2CoreWorkflowTriggerType string

// List of values that V2CoreWorkflowTriggerType can take
const (
	V2CoreWorkflowTriggerTypeEventTrigger V2CoreWorkflowTriggerType = "event_trigger"
	V2CoreWorkflowTriggerTypeManual       V2CoreWorkflowTriggerType = "manual"
)

// The Workflow can be launched when Stripe emits a certain event.
type V2CoreWorkflowTriggerEventTrigger struct {
	// The Stripe event type that will trigger this Workflow.
	Type string `json:"type"`
}

// The Workflow can be launched through a direct call, using either the Dashboard or the Stripe API.
type V2CoreWorkflowTriggerManual struct {
	// An unprocessed version of the trigger's input parameter schema.
	RawSchema string `json:"raw_schema"`
}

// Under what conditions will this Workflow launch.
type V2CoreWorkflowTrigger struct {
	// The Workflow can be launched when Stripe emits a certain event.
	EventTrigger *V2CoreWorkflowTriggerEventTrigger `json:"event_trigger,omitempty"`
	// The Workflow can be launched through a direct call, using either the Dashboard or the Stripe API.
	Manual *V2CoreWorkflowTriggerManual `json:"manual,omitempty"`
	// Which type of trigger this is.
	Type V2CoreWorkflowTriggerType `json:"type"`
}

// A Stripe Workflow is a sequence of actions, like Stripe API calls, that are taken in response to an initiating trigger.
// A trigger can be a Stripe API event, or a manual invocation.
type V2CoreWorkflow struct {
	APIResource
	// When the Workflow was created.
	Created time.Time `json:"created"`
	// Workflow description.
	Description string `json:"description"`
	// The unique ID of the Workflow.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Whether this Workflow is active, inactive, or in some other state. Only active Workflows may be invoked.
	Status V2CoreWorkflowStatus `json:"status"`
	// Under what conditions will this Workflow launch.
	Triggers []*V2CoreWorkflowTrigger `json:"triggers"`
}
