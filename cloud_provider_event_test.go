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

func TestConstructEventFromCloudProvider_EventBridge(t *testing.T) {
	client := &Client{}
	event, err := client.ConstructEventFromCloudProvider([]byte(testEventBridgePayload))
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

func TestConstructEventFromCloudProvider_EventGrid(t *testing.T) {
	client := &Client{}
	event, err := client.ConstructEventFromCloudProvider([]byte(testEventGridPayload))
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

func TestConstructEventFromCloudProvider_InvalidJSON(t *testing.T) {
	client := &Client{}
	_, err := client.ConstructEventFromCloudProvider([]byte("not valid json"))
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestConstructEventFromCloudProvider_RawEventFromClient(t *testing.T) {
	rawEvent := `{"id":"evt_test_123","object":"event","type":"customer.created"}`
	client := &Client{}
	_, err := client.ConstructEventFromCloudProvider([]byte(rawEvent))
	if err == nil {
		t.Fatal("expected error for raw Stripe Event")
	}
	if !strings.Contains(err.Error(), "ConstructEvent") {
		t.Errorf("expected error to mention ConstructEvent, got: %s", err.Error())
	}
}

func TestConstructEventFromCloudProvider_UnrecognizedFormat(t *testing.T) {
	client := &Client{}
	_, err := client.ConstructEventFromCloudProvider([]byte(`{"foo":"bar"}`))
	if err == nil {
		t.Fatal("expected error for unrecognized format")
	}
	if !strings.Contains(err.Error(), "unrecognized cloud event format") {
		t.Errorf("expected error about unrecognized format, got: %s", err.Error())
	}
}

func TestParseEventNotificationFromCloudProvider_EventBridge(t *testing.T) {
	client := &Client{}
	notification, err := client.ParseEventNotificationFromCloudProvider(eventBridgeNotificationPayload)
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

func TestParseEventNotificationFromCloudProvider_EventGrid(t *testing.T) {
	client := &Client{}
	notification, err := client.ParseEventNotificationFromCloudProvider(eventGridNotificationPayload)
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

func TestParseEventNotificationFromCloudProvider_V1EventError(t *testing.T) {
	client := &Client{}
	_, err := client.ParseEventNotificationFromCloudProvider([]byte(testEventBridgePayload))
	if err == nil {
		t.Fatal("expected error when cloud envelope contains a v1 Event")
	}
	if !strings.Contains(err.Error(), "ConstructEvent") {
		t.Errorf("expected error to mention ConstructEventFromCloudProvider, got: %s", err.Error())
	}
}
