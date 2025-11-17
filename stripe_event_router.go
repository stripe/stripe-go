package stripe

import (
	"fmt"
)

type HandlerFunc = func(EventNotificationContainer, *Client) error
type UnhandledHandlerFunc = func(EventNotificationContainer, *Client, UnhandledNotificationDetails) error

type UnhandledNotificationDetails struct {
	IsKnownType bool
}

type EventRouter struct {
	client             *Client
	webhookSecret      string
	eventHandlers      map[string]HandlerFunc
	hasHandledEvent    bool
	onUnhandledHandler UnhandledHandlerFunc
}

func NewEventHandler(client *Client, webhook_secret string, onUnhandledHandler UnhandledHandlerFunc) *EventRouter {
	return &EventRouter{
		client:             client,
		webhookSecret:      webhook_secret,
		eventHandlers:      make(map[string]HandlerFunc),
		hasHandledEvent:    false,
		onUnhandledHandler: onUnhandledHandler,
	}
}

func (r *EventRouter) register(eventType string, callback HandlerFunc) error {
	if r.hasHandledEvent {
		return fmt.Errorf("cannot register new event handlers after handling an event. This is indicative of a bug.")
	}

	if r.eventHandlers[eventType] != nil {
		return fmt.Errorf("handler for event type %s is already registered", eventType)
	}

	r.eventHandlers[eventType] = callback
	return nil
}

func (r *EventRouter) On_V1BillingMeterErrorCompleted(handler func(notification V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error) error {
	return r.register("V1BillingMeterErrorCompleted", handler)
}

// // pre-setup
// client := stripe.NewClient("sk_test_...")
// handler := stripe.NewEventHandler(client, "whsec_...")
// handler.Register[stripe.V1](func (notif, client){})

// // per-webhook
// handler.Handle(webhook_body, sig_header)

func (r *EventRouter) Handle(webhook_body []byte, sig_header string) error {
	r.hasHandledEvent = true

	notif, err := r.client.ParseEventNotification(webhook_body, sig_header, r.webhookSecret)
	if err != nil {
		return err
	}

	n := notif.GetEventNotification()
	eventType := n.Type

	callback, exists := r.eventHandlers[eventType]

	if exists {
		return callback(notif, r.client)
	} else {
		isKnownType := false
		details := UnhandledNotificationDetails{
			IsKnownType: isKnownType,
		}
		return r.onUnhandledHandler(notif, r.client, details)
	}

	// todo: Bind event context to client
}
