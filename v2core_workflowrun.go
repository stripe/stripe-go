//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The current Workflow Run execution status.
type V2CoreWorkflowRunStatus string

// List of values that V2CoreWorkflowRunStatus can take
const (
	V2CoreWorkflowRunStatusFailed    V2CoreWorkflowRunStatus = "failed"
	V2CoreWorkflowRunStatusStarted   V2CoreWorkflowRunStatus = "started"
	V2CoreWorkflowRunStatusSucceeded V2CoreWorkflowRunStatus = "succeeded"
)

// Which type of trigger this is.
type V2CoreWorkflowRunTriggerType string

// List of values that V2CoreWorkflowRunTriggerType can take
const (
	V2CoreWorkflowRunTriggerTypeEventTrigger V2CoreWorkflowRunTriggerType = "event_trigger"
	V2CoreWorkflowRunTriggerTypeManual       V2CoreWorkflowRunTriggerType = "manual"
)

// Details about the Workflow Run's transition into the FAILED state.
type V2CoreWorkflowRunStatusDetailsFailed struct {
	// If this Workflow Run failed during the processing of an action step, the step identifier.
	Action string `json:"action,omitempty"`
	// Category of the failure result.
	ErrorCode string `json:"error_code"`
	// Optional details about the failure result.
	ErrorMessage string `json:"error_message,omitempty"`
}

// Details about the Workflow Run's transition in to the STARTED state.
type V2CoreWorkflowRunStatusDetailsStarted struct{}

// Details about the Workflow Run's transition into the SUCCEEDED state.
type V2CoreWorkflowRunStatusDetailsSucceeded struct {
	// Category of the success result.
	StatusCode string `json:"status_code"`
	// Optional details about the success result.
	StatusMessage string `json:"status_message,omitempty"`
}

// Details about the Workflow Run's status transitions.
type V2CoreWorkflowRunStatusDetails struct {
	// Details about the Workflow Run's transition into the FAILED state.
	Failed *V2CoreWorkflowRunStatusDetailsFailed `json:"failed,omitempty"`
	// Details about the Workflow Run's transition in to the STARTED state.
	Started *V2CoreWorkflowRunStatusDetailsStarted `json:"started,omitempty"`
	// Details about the Workflow Run's transition into the SUCCEEDED state.
	Succeeded *V2CoreWorkflowRunStatusDetailsSucceeded `json:"succeeded,omitempty"`
}

// Summary information about the Workflow Run's status transitions.
type V2CoreWorkflowRunStatusTransitions struct {
	// When the Workflow Run failed.
	FailedAt time.Time `json:"failed_at,omitempty"`
	// When the Workflow Run was started.
	StartedAt time.Time `json:"started_at,omitempty"`
	// When the Workflow Run succeeded.
	SucceededAt time.Time `json:"succeeded_at,omitempty"`
}

// The Workflow Run was launched when Stripe emitted a certain event.
type V2CoreWorkflowRunTriggerEventTrigger struct {
	// The Stripe event that triggered this Run.
	ID string `json:"id"`
	// The Stripe event type triggered this Run.
	Type string `json:"type"`
}

// The Workflow Run was launched through a direct call, using either the Dashboard or the Stripe API.
type V2CoreWorkflowRunTriggerManual struct {
	// The input parameters used when launching the Run.
	InputParameters map[string]any `json:"input_parameters"`
}

// A record of the trigger that launched this Workflow Run.
type V2CoreWorkflowRunTrigger struct {
	// The Workflow Run was launched when Stripe emitted a certain event.
	EventTrigger *V2CoreWorkflowRunTriggerEventTrigger `json:"event_trigger,omitempty"`
	// The Workflow Run was launched through a direct call, using either the Dashboard or the Stripe API.
	Manual *V2CoreWorkflowRunTriggerManual `json:"manual,omitempty"`
	// Which type of trigger this is.
	Type V2CoreWorkflowRunTriggerType `json:"type"`
}

// An execution of a Workflow in response to a triggering event.
type V2CoreWorkflowRun struct {
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
	Status V2CoreWorkflowRunStatus `json:"status"`
	// Details about the Workflow Run's status transitions.
	StatusDetails *V2CoreWorkflowRunStatusDetails `json:"status_details,omitempty"`
	// Summary information about the Workflow Run's status transitions.
	StatusTransitions *V2CoreWorkflowRunStatusTransitions `json:"status_transitions"`
	// A record of the trigger that launched this Workflow Run.
	Trigger *V2CoreWorkflowRunTrigger `json:"trigger"`
	// The Workflow this Run belongs to.
	Workflow string `json:"workflow"`
}
