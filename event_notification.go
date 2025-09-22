package stripe

import (
	"context"
	"time"
)

// EventNotification represents the json that's delivered from an Event Destination.
// Use it to check basic information about a delivered event.
// If you want more details, use `sc.V2Events.Get(thinEvent.ID)`
// to fetch the full event object.
type EventNotification struct {
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

func (en *EventNotification) FetchEvent(ctx context.Context) (V2Event, error) {
	// TODO: usage?
	return en.client.V2CoreEvents.Retrieve(ctx, en.ID, &V2CoreEventRetrieveParams{
		Params{
			StripeContext: en.Context,
		},
	})
}

// implement getThinEvent
// interface must have that

// interface to return from parseThinEvent
// this is how go's unions basically work?
type EventNotificationContainer interface {
	GetEventNotification() *EventNotification
}

type RelatedObject struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type UnknownEventNotification struct {
	EventNotification

	// [Optional] Object containing the reference to API resource relevant to the event
	RelatedObject *RelatedObject `json:"related_object"`
}

func (u *UnknownEventNotification) GetEventNotification() *EventNotification {
	return &u.EventNotification
}

func (u *UnknownEventNotification) FetchRelatedObject(ctx context.Context) (*APIResource, error) {
	if u.RelatedObject == nil {
		return nil, nil
	}

	// TODO: usage?
	// this probably doesn't work for actually accessing the thing? Since we have to know the type
	// maybe need a lookup map here?
	obj := &APIResource{}

	err := u.client.backend.Call("get", u.RelatedObject.URL, u.client.key, nil, obj)

	if err != nil {
		return nil, err
	}

	return obj, nil

	// (ctx, en.ID, &V2CoreEventRetrieveParams{
	// 	Params{
	// 		StripeContext: en.Context,
	// 	},
	// })
}

// FIXME: remove
type V1BillingMeterErrorReportTriggeredEventNotification struct {
	EventNotification
	RelatedObject      RelatedObject                                            `json:"related_object"`
	FetchRelatedObject func() (*BillingMeter, error)                            `json:"-"`
	FetchEvent         func() (*V1BillingMeterErrorReportTriggeredEvent, error) `json:"-"`
}

func (v *V1BillingMeterErrorReportTriggeredEventNotification) GetEventNotification() *EventNotification {
	return &v.EventNotification
}
