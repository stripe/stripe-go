//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The current Workflow Run execution status.
type V2ExtendWorkflowRunStatus string

// List of values that V2ExtendWorkflowRunStatus can take
const (
	V2ExtendWorkflowRunStatusFailed    V2ExtendWorkflowRunStatus = "failed"
	V2ExtendWorkflowRunStatusStarted   V2ExtendWorkflowRunStatus = "started"
	V2ExtendWorkflowRunStatusSucceeded V2ExtendWorkflowRunStatus = "succeeded"
)

// Which type of trigger this is.
type V2ExtendWorkflowRunTriggerType string

// List of values that V2ExtendWorkflowRunTriggerType can take
const (
	V2ExtendWorkflowRunTriggerTypeEventTrigger V2ExtendWorkflowRunTriggerType = "event_trigger"
	V2ExtendWorkflowRunTriggerTypeManual       V2ExtendWorkflowRunTriggerType = "manual"
)

// Details about the Workflow Run's transition into the FAILED state.
type V2ExtendWorkflowRunStatusDetailsFailed struct {
	// Optional details about the failure result.
	ErrorMessage string `json:"error_message,omitempty"`
}

// Details about the Workflow Run's transition in to the STARTED state.
type V2ExtendWorkflowRunStatusDetailsStarted struct{}

// Details about the Workflow Run's transition into the SUCCEEDED state.
type V2ExtendWorkflowRunStatusDetailsSucceeded struct{}

// Details about the Workflow Run's status transitions.
type V2ExtendWorkflowRunStatusDetails struct {
	// Details about the Workflow Run's transition into the FAILED state.
	Failed *V2ExtendWorkflowRunStatusDetailsFailed `json:"failed,omitempty"`
	// Details about the Workflow Run's transition in to the STARTED state.
	Started *V2ExtendWorkflowRunStatusDetailsStarted `json:"started,omitempty"`
	// Details about the Workflow Run's transition into the SUCCEEDED state.
	Succeeded *V2ExtendWorkflowRunStatusDetailsSucceeded `json:"succeeded,omitempty"`
}

// Summary information about the Workflow Run's status transitions.
type V2ExtendWorkflowRunStatusTransitions struct {
	// When the Workflow Run failed.
	FailedAt time.Time `json:"failed_at,omitempty"`
	// When the Workflow Run was started.
	StartedAt time.Time `json:"started_at,omitempty"`
	// When the Workflow Run succeeded.
	SucceededAt time.Time `json:"succeeded_at,omitempty"`
}

// The Workflow Run was launched when Stripe emitted a certain event.
type V2ExtendWorkflowRunTriggerEventTrigger struct {
	// The account that generated the triggering event.
	Context string `json:"context"`
	// The Stripe event that triggered this Run.
	ID string `json:"id"`
	// The Stripe event type triggered this Run.
	Type string `json:"type"`
}

// The Workflow Run was launched through a direct call, using either the Dashboard or the Stripe API.
type V2ExtendWorkflowRunTriggerManual struct {
	// The input parameters used when launching the Run.
	InputParameters map[string]any `json:"input_parameters"`
}

// A record of the trigger that launched this Workflow Run.
type V2ExtendWorkflowRunTrigger struct {
	// The Workflow Run was launched when Stripe emitted a certain event.
	EventTrigger *V2ExtendWorkflowRunTriggerEventTrigger `json:"event_trigger,omitempty"`
	// The Workflow Run was launched through a direct call, using either the Dashboard or the Stripe API.
	Manual *V2ExtendWorkflowRunTriggerManual `json:"manual,omitempty"`
	// Which type of trigger this is.
	Type V2ExtendWorkflowRunTriggerType `json:"type"`
}

// An execution of a Workflow in response to a triggering event.
type V2ExtendWorkflowRun struct {
	APIResource
	// When the Workflow Run was created.
	Created time.Time `json:"created"`
	// The unique ID of the Workflow Run.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The current Workflow Run execution status.
	Status V2ExtendWorkflowRunStatus `json:"status"`
	// Details about the Workflow Run's status transitions.
	StatusDetails *V2ExtendWorkflowRunStatusDetails `json:"status_details,omitempty"`
	// Summary information about the Workflow Run's status transitions.
	StatusTransitions *V2ExtendWorkflowRunStatusTransitions `json:"status_transitions"`
	// A record of the trigger that launched this Workflow Run.
	Trigger *V2ExtendWorkflowRunTrigger `json:"trigger"`
	// The Workflow this Run belongs to.
	Workflow string `json:"workflow"`
}
