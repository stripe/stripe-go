package stripe

import (
	"strings"
	"testing"
)

const testEventBridgePayload = `{
	"version": "0",
	"id": "17e8dff5-d6cd-3770-ace9-aeac02b6ac3f",
	"detail-type": "customer.created",
	"source": "aws.partner/stripe.com/ed_123",
	"account": "506417113029",
	"time": "2024-03-07T18:27:56Z",
	"region": "us-west-2",
	"resources": [],
	"detail": {
		"id": "evt_test_123",
		"object": "event",
		"api_version": "2023-10-16",
		"created": 1709836076,
		"data": {"object": {"id": "cus_123", "object": "customer"}},
		"livemode": true,
		"pending_webhooks": 0,
		"request": {"id": "req_123", "idempotency_key": null},
		"type": "customer.created"
	}
}`

const testEventGridPayload = `{
	"specversion": "1.0",
	"type": "customer.created",
	"source": "/providers/stripe/ed_test_123",
	"id": "9aeb0fdf-c01e-0131-0922-9eb54906e209",
	"time": "2025-07-11T14:30:00Z",
	"subject": null,
	"dataContentType": "application/cloudevents+json",
	"data": {
		"id": "evt_test_456",
		"object": "event",
		"api_version": "2023-10-16",
		"created": 1709836076,
		"data": {"object": {"id": "cus_456", "object": "customer"}},
		"livemode": false,
		"pending_webhooks": 0,
		"request": {"id": "req_456", "idempotency_key": null},
		"type": "customer.created"
	}
}`

var eventBridgeNotificationPayload = []byte(`{
	"version": "0",
	"id": "17e8dff5-d6cd-3770-ace9-aeac02b6ac3f",
	"detail-type": "v2.core.event_destination.ping",
	"source": "aws.partner/stripe.com/ed_123",
	"detail": {
		"id": "evt_test_789",
		"object": "v2.core.event",
		"type": "v2.core.event_destination.ping",
		"created": "2024-03-07T18:27:56.000Z",
		"livemode": true
	}
}`)

var eventGridNotificationPayload = []byte(`{
	"specversion": "1.0",
	"type": "v2.core.event_destination.ping",
	"source": "/providers/stripe/ed_test_123",
	"id": "9aeb0fdf-c01e-0131-0922-9eb54906e209",
	"time": "2025-07-11T14:30:00Z",
	"subject": null,
	"dataContentType": "application/cloudevents+json",
	"data": {
		"id": "evt_test_789",
		"object": "v2.core.event",
		"type": "v2.core.event_destination.ping",
		"created": "2024-03-07T18:27:56.000Z",
		"livemode": true
	}
}`)

func TestConstructEventWithoutVerification_EventBridge(t *testing.T) {
	client := &Client{}
	event, err := client.ConstructEventWithoutVerification([]byte(testEventBridgePayload))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if event.ID != "evt_test_123" {
		t.Errorf("expected event ID 'evt_test_123', got '%s'", event.ID)
	}
	if event.Type != "customer.created" {
		t.Errorf("expected event type 'customer.created', got '%s'", event.Type)
	}
}

func TestConstructEventWithoutVerification_EventGrid(t *testing.T) {
	client := &Client{}
	event, err := client.ConstructEventWithoutVerification([]byte(testEventGridPayload))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if event.ID != "evt_test_456" {
		t.Errorf("expected event ID 'evt_test_456', got '%s'", event.ID)
	}
	if event.Type != "customer.created" {
		t.Errorf("expected event type 'customer.created', got '%s'", event.Type)
	}
}

func TestConstructEventWithoutVerification_RawEvent(t *testing.T) {
	rawEvent := `{"id":"evt_test_123","object":"event","type":"customer.created","api_version":"2023-10-16","created":1709836076,"data":{"object":{}},"livemode":false,"pending_webhooks":0,"request":null}`
	event, err := ConstructEventWithoutVerification([]byte(rawEvent))
	if err != nil {
		t.Fatalf("unexpected error for raw Stripe Event: %v", err)
	}
	if event.ID != "evt_test_123" {
		t.Errorf("expected event ID 'evt_test_123', got '%s'", event.ID)
	}
	if event.Type != "customer.created" {
		t.Errorf("expected event type 'customer.created', got '%s'", event.Type)
	}
}

func TestConstructEventWithoutVerification_RejectsV2ThinEvent(t *testing.T) {
	client := &Client{}
	_, err := client.ConstructEventWithoutVerification(eventBridgeNotificationPayload)
	if err == nil {
		t.Fatal("expected error when EventBridge envelope contains a v2 thin event")
	}
	if !strings.Contains(err.Error(), "EventNotification") {
		t.Errorf("expected error to mention EventNotification, got: %s", err.Error())
	}
}

func TestConstructEventWithoutVerification_InvalidJSON(t *testing.T) {
	client := &Client{}
	_, err := client.ConstructEventWithoutVerification([]byte("not valid json"))
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestConstructEventWithoutVerification_UnrecognizedFormat(t *testing.T) {
	client := &Client{}
	_, err := client.ConstructEventWithoutVerification([]byte(`{"foo":"bar"}`))
	if err == nil {
		t.Fatal("expected error for unrecognized format")
	}
	if !strings.Contains(err.Error(), "unrecognized event format") {
		t.Errorf("expected error about unrecognized format, got: %s", err.Error())
	}
}

func TestParseEventNotificationWithoutVerification_EventBridge(t *testing.T) {
	client := &Client{}
	notification, err := client.ParseEventNotificationWithoutVerification(eventBridgeNotificationPayload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	inner := notification.GetEventNotification()
	if inner.ID != "evt_test_789" {
		t.Errorf("expected notification ID 'evt_test_789', got '%s'", inner.ID)
	}
	if inner.Type != "v2.core.event_destination.ping" {
		t.Errorf("expected notification type 'v2.core.event_destination.ping', got '%s'", inner.Type)
	}
}

func TestParseEventNotificationWithoutVerification_EventGrid(t *testing.T) {
	client := &Client{}
	notification, err := client.ParseEventNotificationWithoutVerification(eventGridNotificationPayload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	inner := notification.GetEventNotification()
	if inner.ID != "evt_test_789" {
		t.Errorf("expected notification ID 'evt_test_789', got '%s'", inner.ID)
	}
	if inner.Type != "v2.core.event_destination.ping" {
		t.Errorf("expected notification type 'v2.core.event_destination.ping', got '%s'", inner.Type)
	}
}

func TestParseEventNotificationWithoutVerification_V1EventError(t *testing.T) {
	client := &Client{}
	_, err := client.ParseEventNotificationWithoutVerification([]byte(testEventBridgePayload))
	if err == nil {
		t.Fatal("expected error when cloud envelope contains a v1 Event")
	}
	if !strings.Contains(err.Error(), "ConstructEvent") {
		t.Errorf("expected error to mention ConstructEvent, got: %s", err.Error())
	}
}

func TestParseEventNotificationWithoutVerification_InvalidJSON(t *testing.T) {
	client := &Client{}
	_, err := client.ParseEventNotificationWithoutVerification([]byte("not valid json"))
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestParseEventNotificationWithoutVerification_UnrecognizedFormat(t *testing.T) {
	client := &Client{}
	_, err := client.ParseEventNotificationWithoutVerification([]byte(`{"foo": "bar"}`))
	if err == nil {
		t.Fatal("expected error for unrecognized format")
	}
	if !strings.Contains(err.Error(), "unrecognized event format") {
		t.Errorf("expected error about unrecognized format, got: %s", err.Error())
	}
}

func TestConstructEventWithoutVerification_EventGridMissingData(t *testing.T) {
	client := &Client{}
	payload := []byte(`{"specversion":"1.0","type":"customer.created","source":"/providers/stripe/ed_test_123","id":"test-missing-data"}`)
	_, err := client.ConstructEventWithoutVerification(payload)
	if err == nil {
		t.Fatal("expected error for Azure envelope missing data field")
	}
	if !strings.Contains(err.Error(), "unrecognized event format") {
		t.Errorf("expected unrecognized format error, got: %s", err.Error())
	}
}

func TestParseEventNotificationWithoutVerification_UnexpectedObjectType(t *testing.T) {
	client := &Client{}
	// Wrap in EventBridge envelope so the inner payload reaches EventNotificationFromJSON
	payload := []byte(`{"version":"0","id":"test","detail-type":"customer.created","source":"aws.partner/stripe.com/ed_123","detail":{"object":"customer","type":"customer.created","id":"cus_123"}}`)
	_, err := client.ParseEventNotificationWithoutVerification(payload)
	if err == nil {
		t.Fatal("expected error for unexpected object type")
	}
	if !strings.Contains(err.Error(), "unexpected object type") {
		t.Errorf("expected error about unexpected object type, got: %s", err.Error())
	}
}

func TestParseEventNotificationWithoutVerification_EventGridMissingData(t *testing.T) {
	client := &Client{}
	payload := []byte(`{"specversion":"1.0","type":"v2.core.event_destination.ping","source":"/providers/stripe/ed_test_123","id":"test-missing-data"}`)
	_, err := client.ParseEventNotificationWithoutVerification(payload)
	if err == nil {
		t.Fatal("expected error for Azure envelope missing data field")
	}
	if !strings.Contains(err.Error(), "unrecognized event format") {
		t.Errorf("expected unrecognized format error, got: %s", err.Error())
	}
}

func TestParseEventNotificationWithoutVerification_RawNotificationPassthrough(t *testing.T) {
	rawNotification := []byte(`{"id": "evt_234", "object": "v2.core.event", "type": "v2.core.event_destination.ping", "created": "2024-03-07T18:27:56.000Z", "livemode": true}`)
	client := &Client{}
	notification, err := client.ParseEventNotificationWithoutVerification(rawNotification)
	if err != nil {
		t.Fatalf("unexpected error for raw v2 notification: %v", err)
	}
	inner := notification.GetEventNotification()
	if inner.ID != "evt_234" {
		t.Errorf("expected notification ID 'evt_234', got '%s'", inner.ID)
	}
}
