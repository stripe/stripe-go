package stripe

import (
	"context"
	"net/http"
	"time"
)

// V2EventNotification represents the json that's delivered from an Event Destination.
// Use it to check basic information about a delivered event.
// If you want more details, use `sc.V2Events.Get(thinEvent.ID)`
// to fetch the full event object.
type V2EventNotification struct {
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
	Context *string `json:"context"`
	// [Optional] Reason for the event
	Reason *V2EventReason `json:"reason"`

	client Client
}

func (en *V2EventNotification) fetchEvent(ctx context.Context) (V2Event, error) {
	// TODO: usage?
	return en.client.V2CoreEvents.Retrieve(ctx, en.ID, &V2CoreEventRetrieveParams{
		Params{
			StripeContext: en.Context,
		},
	})
}

// interface to return from ParseEventNotification
// this is how go's unions basically work?
type EventNotificationContainer interface {
	GetEventNotification() *V2EventNotification
}

type RelatedObject struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type V2UnknownEventNotification struct {
	V2EventNotification

	// [Optional] Object containing the reference to API resource relevant to the event
	RelatedObject *RelatedObject `json:"related_object"`
}

func (en *V2UnknownEventNotification) GetEventNotification() *V2EventNotification {
	return &en.V2EventNotification
}

func (en *V2UnknownEventNotification) FetchEvent(ctx context.Context) (V2Event, error) {
	return en.fetchEvent(ctx)
}

// FetchRelatedObject tries to fetch the related object, if one exists. Returns nil if the struct doesn't have a RelatedObject property
func (en *V2UnknownEventNotification) FetchRelatedObject(ctx context.Context) (*APIResource, error) {
	if en.RelatedObject == nil {
		return nil, nil
	}

	// TODO: usage?
	obj := &APIResource{}

	err := en.client.backend.Call(http.MethodGet, en.RelatedObject.URL, en.client.key, &eventNotificationParams{Params: Params{Context: ctx, StripeContext: en.Context}}, obj)
	return obj, err
}
