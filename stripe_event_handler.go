package stripe

import (
	"fmt"
	"reflect"
)

type Callback = func(EventNotificationContainer, *Client) error

type EventHandler struct {
	client            *Client
	webhook_secret    string
	event_handlers    map[reflect.Type]Callback
	has_handled_event bool
}

func NewEventHandler(client *Client, webhook_secret string) *EventHandler {
	return &EventHandler{
		client:            client,
		webhook_secret:    webhook_secret,
		event_handlers:    make(map[reflect.Type]Callback),
		has_handled_event: false,
	}
}

func Register[T EventNotificationContainer](handler *EventHandler, callback func(T, *Client) error) error {
	if handler.has_handled_event {
		return fmt.Errorf("cannot register new event handlers after handling an event. This is indicative of a bug.")
	}

	// get the concrete type of T
	eventType := reflect.TypeOf((*T)(nil)).Elem()
	fmt.Printf("storing %s\n", eventType)

	if handler.event_handlers[eventType] != nil {
		return fmt.Errorf("handler for event type %s is already registered", eventType)
	}

	// Wrap the typed callback in a type-erased callback that does the type assertion
	handler.event_handlers[eventType] = func(notif EventNotificationContainer, client *Client) error {
		typedNotif := notif.(T)
		return callback(typedNotif, client)
	}

	return nil
}

// // pre-setup
// client := stripe.NewClient("sk_test_...")
// handler := stripe.NewEventHandler(client, "whsec_...")
// handler.Register[stripe.V1](func (notif, client){})

// // per-webhook
// handler.Handle(webhook_body, sig_header)

func (h *EventHandler) Handle(webhook_body []byte, sig_header string) error {
	h.has_handled_event = true

	notif, err := h.client.ParseEventNotification(webhook_body, sig_header, h.webhook_secret)
	if err != nil {
		return err
	}

	eventType := reflect.TypeOf(notif)
	fmt.Printf("looking up handler for %s", eventType)

	callback, exists := h.event_handlers[eventType]
	if !exists {
		// maybe route to unhandled?
		return fmt.Errorf("no handler registered for event type %s", eventType)
	}

	// todo: Bind event context to client
	return callback(notif, h.client)
}
