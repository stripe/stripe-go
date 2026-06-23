package stripe

// This file contains minimal type definitions for types that are referenced by
// handwritten Go SDK code but normally defined in generated files. These stubs
// provide just enough structure to satisfy the compiler.

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// PaymentIntent is a stub for the generated PaymentIntent type.
// Only the fields accessed by handwritten code (error.go) are included.
type PaymentIntent struct {
	APIResource
	ClientSecret string `json:"client_secret"`
}

// PaymentMethod is a stub for the generated PaymentMethod type.
type PaymentMethod struct {
	APIResource
}

// PaymentMethodType is the list of allowed values for a payment method's type.
type PaymentMethodType string

// SetupIntent is a stub for the generated SetupIntent type.
// Only the fields accessed by handwritten code (error.go) are included.
type SetupIntent struct {
	APIResource
	ClientSecret string `json:"client_secret"`
}

// PaymentSource is a stub for the generated PaymentSource type.
type PaymentSource struct {
	APIResource
	ID string `json:"id"`
}

// V2CoreEventReasonType represents the type of an event reason.
type V2CoreEventReasonType string

// List of values that V2CoreEventReasonType can take.
const (
	V2CoreEventReasonTypeRequest V2CoreEventReasonType = "request"
)

// V2CoreEventReasonRequest contains information on the API request that
// instigated the event.
type V2CoreEventReasonRequest struct {
	// ID of the API request that caused the event.
	ID string `json:"id"`
	// The idempotency key transmitted during the request.
	IdempotencyKey string `json:"idempotency_key"`
}

// V2CoreEventReason is the reason for the event.
type V2CoreEventReason struct {
	// Information on the API request that instigated the event.
	Request *V2CoreEventReasonRequest `json:"request,omitempty"`
	// Event reason type.
	Type V2CoreEventReasonType `json:"type"`
}

// V2BaseEvent is the base type for all V2 events. It is embedded in every
// specific event type and in V2CoreRawEvent.
type V2BaseEvent struct {
	APIResource
	// Before and after changes for the primary related object.
	Changes map[string]any `json:"changes,omitempty"`
	// Authentication context needed to fetch the event or related object.
	Context string `json:"context,omitempty"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the event.
	ID string `json:"id"`
	// Has the value true if the object exists in live mode or the value false
	// if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type.
	Object string `json:"object"`
	// Reason for the event.
	Reason *V2CoreEventReason `json:"reason,omitempty"`
	// The type of the event.
	Type string `json:"type"`
}

func (e *V2BaseEvent) getBaseEvent() *V2BaseEvent {
	return e
}

// Event is a stub for the generated V1 Event type.
// Referenced by webhooks.go (ConstructEvent) and stripe_client.go (ConstructEvent method).
type Event struct {
	APIResource
	APIVersion string `json:"api_version"`
	Data       *EventData `json:"data"`
	ID         string `json:"id"`
	Type       string `json:"type"`
}

// EventData is a stub for the generated EventData type.
type EventData struct {
	Raw json.RawMessage `json:"object"`
}

// ErrorCodeLockTimeout is a stub for the generated error code constant.
// Referenced by error.go for retry logic on 429 responses.
const ErrorCodeLockTimeout ErrorCode = "lock_timeout"

// V2CoreEventRetrieveParams is a stub for the generated params type.
// Referenced by event_notification.go when fetching full event details.
type V2CoreEventRetrieveParams struct {
	Params `form:"*"`
}

// v2CoreEventService is a stub for the generated V2 core events service.
// Referenced by event_notification.go via client.V2CoreEvents.Retrieve().
type v2CoreEventService struct {
	B   Backend
	Key string
}

// Retrieve fetches a V2 core event by ID.
func (c v2CoreEventService) Retrieve(ctx context.Context, id string, params *V2CoreEventRetrieveParams) (V2CoreEvent, error) {
	if params == nil {
		params = &V2CoreEventRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/events/%s", id)
	event := &V2CoreRawEvent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, event)
	return event, err
}
