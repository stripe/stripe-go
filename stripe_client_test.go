package stripe_test

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v82"
	. "github.com/stripe/stripe-go/v82/testing"
)

func TestNewClient(t *testing.T) {
	client := stripe.NewClient("foo")
	assert.NotNil(t, client)
}

func TestNewClientWithNilBackend(t *testing.T) {
	client := stripe.NewClient("foo", nil)
	assert.NotNil(t, client)
}

func TestParseEventNotification(t *testing.T) {
	eventNotification := &stripe.V1BillingMeterErrorReportTriggeredEventNotification{
		EventNotification: stripe.EventNotification{
			ID:       "evt_123",
			Object:   "event",
			Type:     "v1.billing.meter.error_report_triggered",
			Livemode: true,
			Created:  time.Now(),
			Context:  nil,
		},
		RelatedObject: stripe.RelatedObject{
			ID:   "bm_123",
			Type: "billing.meter",
			URL:  "/v1/billing/meters/bm_123",
		},
	}
	eventNotificationJSON, err := json.Marshal(eventNotification)
	assert.NoError(t, err)

	p := stripe.SignedPayload{
		UnsignedPayload: stripe.UnsignedPayload{
			Payload:   eventNotificationJSON,
			Timestamp: time.Now(),
			Secret:    "whsec_test_secret",
			Scheme:    "v1",
		},
	}
	p.Signature = stripe.ComputeSignature(p.Timestamp, p.Payload, p.Secret)
	p.Header = generateHeader(p)
	sc := stripe.NewClient("sk_test_secret")
	eventNotificationContainer, err := sc.ParseEventNotification(p.Payload, p.Header, p.Secret)
	assert.NoError(t, err)

	// Type assert the container itself, not the result of GetEventNotification()
	specificEventNotification, ok := eventNotificationContainer.(*stripe.V1BillingMeterErrorReportTriggeredEventNotification)
	assert.True(t, ok)

	assert.Equal(t, "evt_123", specificEventNotification.ID)
	assert.Equal(t, "event", specificEventNotification.Object)
	assert.Equal(t, "v1.billing.meter.error_report_triggered", specificEventNotification.Type)
	assert.Equal(t, true, specificEventNotification.Livemode)
	assert.Equal(t, "bm_123", specificEventNotification.RelatedObject.ID)
	assert.Equal(t, "billing.meter", specificEventNotification.RelatedObject.Type)
	assert.Equal(t, "/v1/billing/meters/bm_123", specificEventNotification.RelatedObject.URL)
}

func TestParseEventNotificationNoTolerance(t *testing.T) {
	unknownEvent := &stripe.UnknownEventNotification{
		EventNotification: stripe.EventNotification{
			ID:       "evt_123",
			Object:   "event",
			Type:     "charge.succeeded",
			Livemode: true,
			Created:  time.Now(),
			Context:  nil,
		},
		RelatedObject: &stripe.RelatedObject{
			ID:   "ch_123",
			Type: "charge",
			URL:  "/v1/charges/ch_123",
		},
	}
	eventJSON, err := json.Marshal(unknownEvent)
	assert.NoError(t, err)
	p := stripe.SignedPayload{
		UnsignedPayload: stripe.UnsignedPayload{
			Payload:   eventJSON,
			Timestamp: time.Now(),
			Secret:    "whsec_test_secret",
			Scheme:    "v1",
		},
	}
	p.Signature = stripe.ComputeSignature(p.Timestamp, p.Payload, p.Secret)
	p.Header = generateHeader(p)
	sc := stripe.NewClient("sk_test_secret")
	_, err = sc.ParseEventNotification(p.Payload, p.Header, p.Secret, stripe.WithTolerance(0))
	assert.Equal(t, stripe.ErrWebhookTooOld, err)
}

func generateHeader(p stripe.SignedPayload) string {
	return fmt.Sprintf("t=%d,%s=%s", p.Timestamp.Unix(), p.Scheme, hex.EncodeToString(p.Signature))
}

func TestParseEventNotificationPing(t *testing.T) {
	// Simulate a meter error event notification
	payload := []byte(`{
		"id": "evt_test_webhook",
		"type": "v2.core.event_destination.ping"
		}`)

	evt, err := stripe.EventNotificationFromJson(payload, *stripe.NewClient("sk_test_secret"))
	if err != nil {
		t.Errorf("Failed to parse event notification: %v", err)
	}
	assert.IsType(t, &stripe.V2CoreEventDestinationPingEvent{}, evt)
}

func TestParseUnknownEventNotification(t *testing.T) {
	payload := []byte(`{
		"id": "evt_test_webhook",
		"type": "imaginary"
		}`)

	evt, err := stripe.EventNotificationFromJson(payload, *stripe.NewClient("sk_test_secret"))
	if err != nil {
		t.Errorf("Failed to parse event notification: %v", err)
	}
	notif := evt.(*stripe.UnknownEventNotification)
	assert.IsType(t, &stripe.UnknownEventNotification{}, evt)

	// just returns nill
	_, err = notif.FetchRelatedObject(context.TODO())
	assert.Nil(t, err)
}

func TestFetchEventHTTPCall(t *testing.T) {
	// Create mock response for the V2 event retrieve call

	mockEventResponse := `{
		"id": "evt_123",
		"object": "event",
		"type": "v1.billing.meter.error_report_triggered",
		"data": {
			"developer_message_summary": "very cool"
		},
		"related_object": {
			"id": "bm_123",
			"type": "billing.meter",
			"url": "/v1/billing/meters/bm_123"
		}
	}`

	// Create mock server that expects a GET to /v2/core/events/evt_123
	server := MockServer(t, http.MethodGet, "/v2/core/events/evt_123", nil, mockEventResponse)
	defer server.Close()

	// Create client with custom backend pointing to mock server
	backend := stripe.GetBackendWithConfig(stripe.APIBackend, &stripe.BackendConfig{
		URL: stripe.String(server.URL),
	})
	client := stripe.NewClient(TestAPIKey, stripe.WithBackends(&stripe.Backends{
		API:         backend,
		Connect:     backend,
		Uploads:     backend,
		MeterEvents: backend,
	}))

	// Create a simple event payload for an unknown event type
	payload := []byte(`{
		"id": "evt_123",
		"object": "event",
		"type": "v1.billing.meter.error_report_triggered"
	}`)

	// Create signed payload for ParseEventNotification
	p := stripe.SignedPayload{
		UnsignedPayload: stripe.UnsignedPayload{
			Payload:   payload,
			Timestamp: time.Now(),
			Secret:    "whsec_test_secret",
			Scheme:    "v1",
		},
	}
	p.Signature = stripe.ComputeSignature(p.Timestamp, p.Payload, p.Secret)
	p.Header = generateHeader(p)

	// Parse the event notification - this will set up the client properly
	eventNotificationContainer, err := client.ParseEventNotification(p.Payload, p.Header, p.Secret)
	assert.NoError(t, err)

	eventNotification, ok := eventNotificationContainer.(*stripe.V1BillingMeterErrorReportTriggeredEventNotification)
	assert.True(t, ok)
	assert.NotNil(t, eventNotification)

	event, err := eventNotification.FetchEvent(context.TODO())
	assert.NoError(t, err)

	assert.IsType(t, &stripe.V1BillingMeterErrorReportTriggeredEvent{}, event)
}
func TestFetchRelatedObject(t *testing.T) {
	// Create mock response for the V2 event retrieve call

	mockRelatedObjectResponse := `{
		"id": "bm_123",
		"object": "billing.meter",
		"display_name": "my_cool_meter"
	}`

	// Create mock server that expects a GET to /v1/billing/meters/bm_123
	server := MockServer(t, http.MethodGet, "/v1/billing/meters/bm_123", nil, mockRelatedObjectResponse)
	defer server.Close()

	// Create client with custom backend pointing to mock server
	backend := stripe.GetBackendWithConfig(stripe.APIBackend, &stripe.BackendConfig{
		URL: stripe.String(server.URL),
	})
	client := stripe.NewClient(TestAPIKey, stripe.WithBackends(&stripe.Backends{
		API:         backend,
		Connect:     backend,
		Uploads:     backend,
		MeterEvents: backend,
	}))

	// Create a simple event payload for an unknown event type
	payload := []byte(`{
		"id": "evt_123",
		"object": "event",
		"type": "v1.billing.meter.error_report_triggered",
		"related_object": {
			"id": "bm_123",
			"type": "billing.meter",
			"url": "/v1/billing/meters/bm_123"
		}
	}`)

	// Create signed payload for ParseEventNotification
	p := stripe.SignedPayload{
		UnsignedPayload: stripe.UnsignedPayload{
			Payload:   payload,
			Timestamp: time.Now(),
			Secret:    "whsec_test_secret",
			Scheme:    "v1",
		},
	}
	p.Signature = stripe.ComputeSignature(p.Timestamp, p.Payload, p.Secret)
	p.Header = generateHeader(p)

	// Parse the event notification - this will set up the client properly
	eventNotificationContainer, err := client.ParseEventNotification(p.Payload, p.Header, p.Secret)
	assert.NoError(t, err)

	eventNotification, ok := eventNotificationContainer.(*stripe.V1BillingMeterErrorReportTriggeredEventNotification)
	assert.True(t, ok)
	assert.NotNil(t, eventNotification)
	assert.NotNil(t, eventNotification.RelatedObject)

	event, err := eventNotification.FetchRelatedObject(context.TODO())
	assert.NoError(t, err)

	assert.IsType(t, &stripe.BillingMeter{}, event)
}

func TestFetchRelatedObjectUnknownEvent(t *testing.T) {
	// Create mock response for the V2 event retrieve call

	mockRelatedObjectResponse := `{
		"id": "bm_123",
		"object": "billing.meter",
		"display_name": "my_cool_meter"
	}`

	// Create mock server that expects a GET to /v1/billing/meters/bm_123
	server := MockServer(t, http.MethodGet, "/v1/billing/meters/bm_123", nil, mockRelatedObjectResponse)
	defer server.Close()

	// Create client with custom backend pointing to mock server
	backend := stripe.GetBackendWithConfig(stripe.APIBackend, &stripe.BackendConfig{
		URL: stripe.String(server.URL),
	})
	client := stripe.NewClient(TestAPIKey, stripe.WithBackends(&stripe.Backends{
		API:         backend,
		Connect:     backend,
		Uploads:     backend,
		MeterEvents: backend,
	}))

	// Create a simple event payload for an unknown event type
	payload := []byte(`{
		"id": "evt_123",
		"object": "event",
		"type": "imaginary_with_related_object",
		"related_object": {
			"id": "bm_123",
			"type": "billing.meter",
			"url": "/v1/billing/meters/bm_123"
		}
	}`)

	// Create signed payload for ParseEventNotification
	p := stripe.SignedPayload{
		UnsignedPayload: stripe.UnsignedPayload{
			Payload:   payload,
			Timestamp: time.Now(),
			Secret:    "whsec_test_secret",
			Scheme:    "v1",
		},
	}
	p.Signature = stripe.ComputeSignature(p.Timestamp, p.Payload, p.Secret)
	p.Header = generateHeader(p)

	// Parse the event notification - this will set up the client properly
	eventNotificationContainer, err := client.ParseEventNotification(p.Payload, p.Header, p.Secret)
	assert.NoError(t, err)

	eventNotification, ok := eventNotificationContainer.(*stripe.UnknownEventNotification)
	assert.True(t, ok)
	assert.NotNil(t, eventNotification)
	assert.NotNil(t, eventNotification.RelatedObject)

	related_obj, err := eventNotification.FetchRelatedObject(context.TODO())
	assert.NoError(t, err)
	assert.IsType(t, &stripe.APIResource{}, related_obj)

	type UnknownResponse struct {
		Object string `json:"object"`
	}

	var rawResp UnknownResponse
	err = json.Unmarshal(related_obj.LastResponse.RawJSON, &rawResp)
	assert.NoError(t, err)
	assert.Equal(t, "billing.meter", rawResp.Object)
}
