package stripe

import (
	"context"
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
)

const testWebhookSecret = "whsec_test_secret"

// Test helper to generate a valid webhook signature
func generateTestSignature(payload string) string {
	signedPayload := GenerateTestSignedPayload(&UnsignedPayload{
		Payload: []byte(payload),
		Secret:  testWebhookSecret,
	})
	return signedPayload.Header
}

// Test payloads
func v1BillingMeterPayload() string {
	return `{
		"id": "evt_123",
		"object": "v2.core.event",
		"type": "v1.billing.meter.error_report_triggered",
		"livemode": false,
		"created": "2022-02-15T00:27:45.330Z",
		"context": "event_context_456",
		"related_object": {
			"id": "mtr_123",
			"type": "billing.meter",
			"url": "/v1/billing/meters/mtr_123"
		}
	}`
}

func v1BillingMeterNoMeterFoundPayload() string {
	return `{
		"id": "evt_456",
		"object": "v2.core.event",
		"type": "v1.billing.meter.no_meter_found",
		"livemode": false,
		"created": "2022-02-15T00:27:45.330Z",
		"context": "event_context_789"
	}`
}

func v1BillingMeterPayloadWithNilContext() string {
	return `{
		"id": "evt_789",
		"object": "v2.core.event",
		"type": "v1.billing.meter.error_report_triggered",
		"livemode": false,
		"created": "2022-02-15T00:27:45.330Z",
		"context": null,
		"related_object": {
			"id": "mtr_456",
			"type": "billing.meter",
			"url": "/v1/billing/meters/mtr_456"
		}
	}`
}

func unknownEventPayload() string {
	return `{
		"id": "evt_unknown",
		"object": "v2.core.event",
		"type": "llama.created",
		"livemode": false,
		"created": "2022-02-15T00:27:45.330Z",
		"context": "event_context_unknown",
		"related_object": {
			"id": "llama_123",
			"type": "llama",
			"url": "/v1/llamas/llama_123"
		}
	}`
}

// Test: Event is routed to registered handler
func TestEventHandler_RoutesEventToRegisteredHandler(t *testing.T) {
	// Setup
	config := &BackendConfig{StripeContext: String("original_context_123")}
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(config)))

	var unhandledCalled bool
	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		unhandledCalled = true
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	var handlerCalled bool
	var receivedEvent *V1BillingMeterErrorReportTriggeredEventNotification
	var receivedClient *Client

	callback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		handlerCalled = true
		receivedEvent = event
		receivedClient = client
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback)
	assert.NoError(t, err)

	// Execute
	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = handler.Handle(context.TODO(), []byte(payload), sigHeader)

	// Assert
	assert.NoError(t, err)
	assert.True(t, handlerCalled, "Handler should have been called")
	assert.NotNil(t, receivedEvent, "Handler should have received event")
	assert.Equal(t, "v1.billing.meter.error_report_triggered", receivedEvent.Type)
	assert.Equal(t, "evt_123", receivedEvent.ID)
	assert.Equal(t, "mtr_123", receivedEvent.RelatedObject.ID)
	assert.NotNil(t, receivedClient, "Handler should have received client")
	assert.False(t, unhandledCalled, "Unhandled handler should not be called")
}

// Test: Different events route to correct handlers
func TestEventHandler_RoutesDifferentEventsToCorrectHandlers(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	var unhandledCalled bool
	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		unhandledCalled = true
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	var billingCallbackCalled, noMeterCallbackCalled bool
	var billingEvent *V1BillingMeterErrorReportTriggeredEventNotification
	var noMeterEvent *V1BillingMeterNoMeterFoundEventNotification

	billingCallback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		billingCallbackCalled = true
		billingEvent = event
		return nil
	}

	noMeterCallback := func(event *V1BillingMeterNoMeterFoundEventNotification, client *Client) error {
		noMeterCallbackCalled = true
		noMeterEvent = event
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(billingCallback)
	assert.NoError(t, err)
	err = handler.OnV1BillingMeterNoMeterFound(noMeterCallback)
	assert.NoError(t, err)

	// Execute first event
	payload1 := v1BillingMeterPayload()
	sigHeader1 := generateTestSignature(payload1)
	err = handler.Handle(context.TODO(), []byte(payload1), sigHeader1)
	assert.NoError(t, err)

	// Execute second event
	payload2 := v1BillingMeterNoMeterFoundPayload()
	sigHeader2 := generateTestSignature(payload2)
	err = handler.Handle(context.TODO(), []byte(payload2), sigHeader2)
	assert.NoError(t, err)

	// Assert
	assert.True(t, billingCallbackCalled, "Billing handler should have been called")
	assert.True(t, noMeterCallbackCalled, "NoMeter handler should have been called")
	assert.Equal(t, "v1.billing.meter.error_report_triggered", billingEvent.Type)
	assert.Equal(t, "v1.billing.meter.no_meter_found", noMeterEvent.Type)
	assert.False(t, unhandledCalled, "Unhandled handler should not be called")
}

// Test: Handler receives correct runtime type
func TestEventHandler_HandlerReceivesCorrectRuntimeType(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	var receivedEvent *V1BillingMeterErrorReportTriggeredEventNotification
	var receivedClient *Client

	callback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		receivedEvent = event
		receivedClient = client
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback)
	assert.NoError(t, err)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = handler.Handle(context.TODO(), []byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.NotNil(t, receivedEvent)
	assert.Equal(t, "v1.billing.meter.error_report_triggered", receivedEvent.Type)
	assert.Equal(t, "evt_123", receivedEvent.ID)
	assert.Equal(t, "mtr_123", receivedEvent.RelatedObject.ID)
	assert.NotNil(t, receivedClient)
}

// Test: Cannot register handler after handling
func TestEventHandler_CannotRegisterHandlerAfterHandling(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	callback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback)
	assert.NoError(t, err)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = handler.Handle(context.TODO(), []byte(payload), sigHeader)
	assert.NoError(t, err)

	// Try to register another handler after handling
	callback2 := func(event *V1BillingMeterNoMeterFoundEventNotification, client *Client) error {
		return nil
	}

	err = handler.OnV1BillingMeterNoMeterFound(callback2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot register new event handlers after handling an event")
}

// Test: Cannot register duplicate handler
func TestEventHandler_CannotRegisterDuplicateHandler(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	callback1 := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		return nil
	}

	callback2 := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback1)
	assert.NoError(t, err)

	err = handler.OnV1BillingMeterErrorReportTriggered(callback2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "handler for event type v1.billing.meter.error_report_triggered is already registered")
}

// Test: Handler uses event stripe context
func TestEventHandler_HandlerUsesEventStripeContext(t *testing.T) {
	config := &BackendConfig{StripeContext: String("original_context_123")}
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(config)))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	var receivedContext *string

	callback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		receivedContext = client.backends.config.StripeContext
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback)
	assert.NoError(t, err)

	// Verify original context
	assert.Equal(t, "original_context_123", *client.backends.config.StripeContext)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = handler.Handle(context.TODO(), []byte(payload), sigHeader)
	assert.NoError(t, err)

	// Assert handler received event context
	assert.NotNil(t, receivedContext)
	assert.Equal(t, "event_context_456", *receivedContext)
}

// Test: Stripe context restored after handler success
func TestEventHandler_StripeContextRestoredAfterHandlerSuccess(t *testing.T) {
	config := &BackendConfig{StripeContext: String("original_context_123")}
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(config)))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	callback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		// Verify context is event context during handler
		assert.Equal(t, "event_context_456", *client.backends.config.StripeContext)
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback)
	assert.NoError(t, err)

	// Verify original context before handle
	assert.Equal(t, "original_context_123", *client.backends.config.StripeContext)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = handler.Handle(context.TODO(), []byte(payload), sigHeader)
	assert.NoError(t, err)

	// Verify original client context is unchanged after handle
	assert.Equal(t, "original_context_123", *client.backends.config.StripeContext)
}

// Test: Stripe context restored after handler error
func TestEventHandler_StripeContextRestoredAfterHandlerError(t *testing.T) {
	config := &BackendConfig{StripeContext: String("original_context_123")}
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(config)))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	callback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		// Verify context is event context during handler
		assert.Equal(t, "event_context_456", *client.backends.config.StripeContext)
		// Return an error
		return fmt.Errorf("handler error")
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback)
	assert.NoError(t, err)

	// Verify original context before handle
	assert.Equal(t, "original_context_123", *client.backends.config.StripeContext)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = handler.Handle(context.TODO(), []byte(payload), sigHeader)
	assert.Error(t, err) // Handler returned an error

	// Verify original client context is unchanged after handler error
	assert.Equal(t, "original_context_123", *client.backends.config.StripeContext)
}

// Test: Stripe context set to nil when event has no context
func TestEventHandler_StripeContextSetToNilWhenEventHasNoContext(t *testing.T) {
	config := &BackendConfig{StripeContext: String("original_context_123")}
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(config)))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	var receivedContext *string
	var receivedContextWasChecked bool

	callback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		receivedContext = client.backends.config.StripeContext
		receivedContextWasChecked = true
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback)
	assert.NoError(t, err)

	// Verify original context
	assert.Equal(t, "original_context_123", *client.backends.config.StripeContext)

	payload := v1BillingMeterPayloadWithNilContext()
	sigHeader := generateTestSignature(payload)
	err = handler.Handle(context.TODO(), []byte(payload), sigHeader)
	assert.NoError(t, err)

	// Assert handler received nil context
	assert.True(t, receivedContextWasChecked)
	assert.Nil(t, receivedContext)

	// Verify original client context is unchanged
	assert.Equal(t, "original_context_123", *client.backends.config.StripeContext)
}

// Test: Unknown event Oout
func TestEventHandler_UnknownEventRoutesToOnUnhandled(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	var unhandledCalled bool
	var unhandledEvent EventNotificationContainer
	var unhandledClient *Client
	var unhandledDetails UnhandledNotificationDetails

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		unhandledCalled = true
		unhandledEvent = notif
		unhandledClient = client
		unhandledDetails = details
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	payload := unknownEventPayload()
	sigHeader := generateTestSignature(payload)
	err := handler.Handle(context.TODO(), []byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.True(t, unhandledCalled, "Unhandled handler should have been called")
	assert.NotNil(t, unhandledEvent)

	// Check if it's an UnknownEventNotification
	_, isUnknown := unhandledEvent.(*UnknownEventNotification)
	assert.True(t, isUnknown, "Event should be UnknownEventNotification")

	assert.Equal(t, "llama.created", unhandledEvent.GetEventNotification().Type)
	assert.NotNil(t, unhandledClient)
	assert.False(t, unhandledDetails.IsKnownEventType, "Unknown event should have IsKnownEventType=false")
}

// Test: Known unregistered event Oout
func TestEventHandler_KnownUnregisteredEventRoutesToOnUnhandled(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	var unhandledCalled bool
	var unhandledEvent EventNotificationContainer
	var unhandledClient *Client
	var unhandledDetails UnhandledNotificationDetails

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		unhandledCalled = true
		unhandledEvent = notif
		unhandledClient = client
		unhandledDetails = details
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	// Don't register a handler for this known event type
	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err := handler.Handle(context.TODO(), []byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.True(t, unhandledCalled, "Unhandled handler should have been called")
	assert.NotNil(t, unhandledEvent)

	// Check if it's a known event type
	_, isKnown := unhandledEvent.(*V1BillingMeterErrorReportTriggeredEventNotification)
	assert.True(t, isKnown, "Event should be V1BillingMeterErrorReportTriggeredEventNotification")

	assert.Equal(t, "v1.billing.meter.error_report_triggered", unhandledEvent.GetEventNotification().Type)
	assert.NotNil(t, unhandledClient)
	assert.True(t, unhandledDetails.IsKnownEventType, "Known event should have IsKnownEventType=true")
}

// Test: Registered event does Oot
func TestEventHandler_RegisteredEventDoesNotCallOnUnhandled(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	var unhandledCalled bool
	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		unhandledCalled = true
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	var callbackCalled bool
	callback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		callbackCalled = true
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback)
	assert.NoError(t, err)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = handler.Handle(context.TODO(), []byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.True(t, callbackCalled, "Handler should have been called")
	assert.False(t, unhandledCalled, "Unhandled handler should not be called")
}

// Test: Handler client retains configuration
func TestEventHandler_HandlerClientRetainsConfiguration(t *testing.T) {
	apiKey := "sk_test_custom_key"
	config := &BackendConfig{StripeContext: String("original_context_xyz")}
	client := NewClient(apiKey, WithBackends(NewBackendsWithConfig(config)))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	var receivedAPIKey string
	var receivedContext *string

	callback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		receivedAPIKey = client.key
		receivedContext = client.backends.config.StripeContext
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback)
	assert.NoError(t, err)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = handler.Handle(context.TODO(), []byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.Equal(t, apiKey, receivedAPIKey, "Handler should receive client with original API key")
	assert.Equal(t, "event_context_456", *receivedContext, "Handler should receive event context")
	assert.Equal(t, "original_context_xyz", *client.backends.config.StripeContext, "Original client context should be unchanged")
}

// Test: on_unhandled receives correct info for unknown
func TestEventHandler_OnUnhandledReceivesCorrectInfoForUnknown(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	var details UnhandledNotificationDetails
	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, d UnhandledNotificationDetails) error {
		details = d
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	payload := unknownEventPayload()
	sigHeader := generateTestSignature(payload)
	err := handler.Handle(context.TODO(), []byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.False(t, details.IsKnownEventType)
}

// Test: on_unhandled receives correct info for known unregistered
func TestEventHandler_OnUnhandledReceivesCorrectInfoForKnownUnregistered(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	var details UnhandledNotificationDetails
	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, d UnhandledNotificationDetails) error {
		details = d
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err := handler.Handle(context.TODO(), []byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.True(t, details.IsKnownEventType)
}

// Test: Validates webhook signature
func TestEventHandler_ValidatesWebhookSignature(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	payload := v1BillingMeterPayload()
	err := handler.Handle(context.TODO(), []byte(payload), "invalid_signature")

	assert.Error(t, err)
	// Check if error is related to signature verification
	assert.Contains(t, err.Error(), "Stripe-Signature")
}

// Test: RegisteredEventTypes returns empty list
func TestEventHandler_RegisteredEventTypesEmpty(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	types := handler.RegisteredEventTypes()
	assert.Empty(t, types)
}

// Test: RegisteredEventTypes returns single event type
func TestEventHandler_RegisteredEventTypesSingle(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	handler := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	callback := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		return nil
	}

	err := handler.OnV1BillingMeterErrorReportTriggered(callback)
	assert.NoError(t, err)

	types := handler.RegisteredEventTypes()
	assert.Equal(t, []string{"v1.billing.meter.error_report_triggered"}, types)
}

// Test: RegisteredEventTypes returns multiple event types in alphabetical order
func TestEventHandler_RegisteredEventTypesMultipleAlphabetized(t *testing.T) {
	client := NewClient("sk_test_1234", WithBackends(NewBackendsWithConfig(&BackendConfig{})))

	onUnhandled := func(ctx context.Context, notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	callback := NewEventNotificationHandler(client, testWebhookSecret, onUnhandled)

	callback1 := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		return nil
	}
	callback2 := func(event *V1BillingMeterNoMeterFoundEventNotification, client *Client) error {
		return nil
	}
	callback3 := func(event *V2CoreAccountCreatedEventNotification, client *Client) error {
		return nil
	}

	// Register in non-alphabetical order
	err := callback.OnV2CoreAccountCreated(callback3)
	assert.NoError(t, err)
	err = callback.OnV1BillingMeterErrorReportTriggered(callback1)
	assert.NoError(t, err)
	err = callback.OnV1BillingMeterNoMeterFound(callback2)
	assert.NoError(t, err)

	types := callback.RegisteredEventTypes()
	expected := []string{
		"v1.billing.meter.error_report_triggered",
		"v1.billing.meter.no_meter_found",
		"v2.core.account.created",
	}

	assert.Equal(t, expected, types)
}
