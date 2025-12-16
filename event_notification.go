package stripe

import (
	"context"
	"net/http"
	"time"
)

// V2CoreEventNotification represents the json that's delivered from an Event Destination.
// Use it to check basic information about an event.
// If you want more details, use the `FetchEvent()` instance method
// to fetch the full Event object.
type V2CoreEventNotification struct {
	// Unique identifier for the event
	ID string `json:"id"`
	// The string "event"
	Object string `json:"object"`
	// The type of the event
	Type string `json:"type"`
	// Livemode indicates if the event is from a production(true) or test(false) account
	Livemode bool `json:"livemode"`
	// Time at which the event was created
	Created time.Time `json:"created"`
	// [Optional] Authentication context needed to fetch the event or related object
	Context *Context `json:"context"`
	// [Optional] Reason for the event
	Reason *V2CoreEventReason `json:"reason"`

	client Client
}

func (n *V2CoreEventNotification) GetEventNotification() *V2CoreEventNotification {
	return n
}

func (n *V2CoreEventNotification) fetchEvent(ctx context.Context) (V2CoreEvent, error) {
	// TODO: usage?
	params := &V2CoreEventRetrieveParams{}
	params.SetStripeContextFrom(n.Context)

	return n.client.V2CoreEvents.Retrieve(ctx, n.ID, params)
}

// interface to return from ParseEventNotification
// this is how go's unions basically work?
type EventNotificationContainer interface {
	GetEventNotification() *V2CoreEventNotification
}

type V2CoreEventRelatedObject struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type UnknownEventNotification struct {
	V2CoreEventNotification

	// [Optional] Object containing the reference to API resource relevant to the event
	RelatedObject *V2CoreEventRelatedObject `json:"related_object"`
}

func (n *UnknownEventNotification) FetchEvent(ctx context.Context) (V2CoreEvent, error) {
	return n.fetchEvent(ctx)
}

// FetchRelatedObject tries to fetch the related object, if one exists. Returns nil if the struct doesn't have a RelatedObject property
func (n *UnknownEventNotification) FetchRelatedObject(ctx context.Context) (*APIResource, error) {
	if n.RelatedObject == nil {
		return nil, nil
	}

	// TODO: usage?
	obj := &APIResource{}
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)

	err := n.client.backends.API.Call(http.MethodGet, n.RelatedObject.URL, n.client.key, params, obj)
	return obj, err
}
