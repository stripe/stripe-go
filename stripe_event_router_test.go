package stripe

import (
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
func TestEventRouter_RoutesEventToRegisteredHandler(t *testing.T) {
	// Setup
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	backend.(*BackendImplementation).StripeContext = String("original_context_123")

	client := &Client{
		backend: backend,
		key:     "sk_test_1234",
	}

	var unhandledCalled bool
	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		unhandledCalled = true
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	var handlerCalled bool
	var receivedEvent *V1BillingMeterErrorReportTriggeredEventNotification
	var receivedClient *Client

	handler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		handlerCalled = true
		receivedEvent = event
		receivedClient = client
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler)
	assert.NoError(t, err)

	// Execute
	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = router.Handle([]byte(payload), sigHeader)

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
func TestEventRouter_RoutesDifferentEventsToCorrectHandlers(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	var unhandledCalled bool
	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		unhandledCalled = true
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	var billingHandlerCalled, noMeterHandlerCalled bool
	var billingEvent *V1BillingMeterErrorReportTriggeredEventNotification
	var noMeterEvent *V1BillingMeterNoMeterFoundEventNotification

	billingHandler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		billingHandlerCalled = true
		billingEvent = event
		return nil
	}

	noMeterHandler := func(event *V1BillingMeterNoMeterFoundEventNotification, client *Client) error {
		noMeterHandlerCalled = true
		noMeterEvent = event
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(billingHandler)
	assert.NoError(t, err)
	err = router.OnV1BillingMeterNoMeterFound(noMeterHandler)
	assert.NoError(t, err)

	// Execute first event
	payload1 := v1BillingMeterPayload()
	sigHeader1 := generateTestSignature(payload1)
	err = router.Handle([]byte(payload1), sigHeader1)
	assert.NoError(t, err)

	// Execute second event
	payload2 := v1BillingMeterNoMeterFoundPayload()
	sigHeader2 := generateTestSignature(payload2)
	err = router.Handle([]byte(payload2), sigHeader2)
	assert.NoError(t, err)

	// Assert
	assert.True(t, billingHandlerCalled, "Billing handler should have been called")
	assert.True(t, noMeterHandlerCalled, "NoMeter handler should have been called")
	assert.Equal(t, "v1.billing.meter.error_report_triggered", billingEvent.Type)
	assert.Equal(t, "v1.billing.meter.no_meter_found", noMeterEvent.Type)
	assert.False(t, unhandledCalled, "Unhandled handler should not be called")
}

// Test: Handler receives correct runtime type
func TestEventRouter_HandlerReceivesCorrectRuntimeType(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	var receivedEvent *V1BillingMeterErrorReportTriggeredEventNotification
	var receivedClient *Client

	handler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		receivedEvent = event
		receivedClient = client
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler)
	assert.NoError(t, err)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = router.Handle([]byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.NotNil(t, receivedEvent)
	assert.Equal(t, "v1.billing.meter.error_report_triggered", receivedEvent.Type)
	assert.Equal(t, "evt_123", receivedEvent.ID)
	assert.Equal(t, "mtr_123", receivedEvent.RelatedObject.ID)
	assert.NotNil(t, receivedClient)
}

// Test: Cannot register handler after handling
func TestEventRouter_CannotRegisterHandlerAfterHandling(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	handler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler)
	assert.NoError(t, err)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = router.Handle([]byte(payload), sigHeader)
	assert.NoError(t, err)

	// Try to register another handler after handling
	handler2 := func(event *V1BillingMeterNoMeterFoundEventNotification, client *Client) error {
		return nil
	}

	err = router.OnV1BillingMeterNoMeterFound(handler2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot register new event handlers after handling an event")
}

// Test: Cannot register duplicate handler
func TestEventRouter_CannotRegisterDuplicateHandler(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	handler1 := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		return nil
	}

	handler2 := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler1)
	assert.NoError(t, err)

	err = router.OnV1BillingMeterErrorReportTriggered(handler2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "handler for event type v1.billing.meter.error_report_triggered is already registered")
}

// Test: Handler uses event stripe context
func TestEventRouter_HandlerUsesEventStripeContext(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	backend.(*BackendImplementation).StripeContext = String("original_context_123")

	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	var receivedContext *string

	handler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		backendImpl := client.backend.(*BackendImplementation)
		receivedContext = backendImpl.StripeContext
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler)
	assert.NoError(t, err)

	// Verify original context
	originalContext := backend.(*BackendImplementation).StripeContext
	assert.Equal(t, "original_context_123", *originalContext)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = router.Handle([]byte(payload), sigHeader)
	assert.NoError(t, err)

	// Assert handler received event context
	assert.NotNil(t, receivedContext)
	assert.Equal(t, "event_context_456", *receivedContext)
}

// Test: Stripe context restored after handler success
func TestEventRouter_StripeContextRestoredAfterHandlerSuccess(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	backend.(*BackendImplementation).StripeContext = String("original_context_123")

	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	handler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		// Verify context is event context during handler
		backendImpl := client.backend.(*BackendImplementation)
		assert.Equal(t, "event_context_456", *backendImpl.StripeContext)
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler)
	assert.NoError(t, err)

	// Verify original context before handle
	assert.Equal(t, "original_context_123", *backend.(*BackendImplementation).StripeContext)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = router.Handle([]byte(payload), sigHeader)
	assert.NoError(t, err)

	// Verify context is restored after handle
	assert.Equal(t, "original_context_123", *backend.(*BackendImplementation).StripeContext)
}

// Test: Stripe context restored after handler error
func TestEventRouter_StripeContextRestoredAfterHandlerError(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	backend.(*BackendImplementation).StripeContext = String("original_context_123")

	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	handler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		// Verify context is event context during handler
		backendImpl := client.backend.(*BackendImplementation)
		assert.Equal(t, "event_context_456", *backendImpl.StripeContext)
		// Return an error
		return fmt.Errorf("handler error")
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler)
	assert.NoError(t, err)

	// Verify original context before handle
	assert.Equal(t, "original_context_123", *backend.(*BackendImplementation).StripeContext)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = router.Handle([]byte(payload), sigHeader)
	assert.Error(t, err) // Handler returned an error

	// Verify context is still restored after handler error
	assert.Equal(t, "original_context_123", *backend.(*BackendImplementation).StripeContext)
}

// Test: Stripe context set to nil when event has no context
func TestEventRouter_StripeContextSetToNilWhenEventHasNoContext(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	backend.(*BackendImplementation).StripeContext = String("original_context_123")

	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	var receivedContext *string

	handler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		backendImpl := client.backend.(*BackendImplementation)
		receivedContext = backendImpl.StripeContext
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler)
	assert.NoError(t, err)

	// Verify original context
	assert.Equal(t, "original_context_123", *backend.(*BackendImplementation).StripeContext)

	payload := v1BillingMeterPayloadWithNilContext()
	sigHeader := generateTestSignature(payload)
	err = router.Handle([]byte(payload), sigHeader)
	assert.NoError(t, err)

	// Assert handler received nil context
	assert.Nil(t, receivedContext)

	// Verify context is restored
	assert.Equal(t, "original_context_123", *backend.(*BackendImplementation).StripeContext)
}

// Test: Unknown event Oout
func TestEventRouter_UnknownEventRoutesToOnUnhandled(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	var unhandledCalled bool
	var unhandledEvent EventNotificationContainer
	var unhandledClient *Client
	var unhandledDetails UnhandledNotificationDetails

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		unhandledCalled = true
		unhandledEvent = notif
		unhandledClient = client
		unhandledDetails = details
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	payload := unknownEventPayload()
	sigHeader := generateTestSignature(payload)
	err := router.Handle([]byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.True(t, unhandledCalled, "Unhandled handler should have been called")
	assert.NotNil(t, unhandledEvent)

	// Check if it's an UnknownEventNotification
	_, isUnknown := unhandledEvent.(*UnknownEventNotification)
	assert.True(t, isUnknown, "Event should be UnknownEventNotification")

	assert.Equal(t, "llama.created", unhandledEvent.GetEventNotification().Type)
	assert.NotNil(t, unhandledClient)
	assert.False(t, unhandledDetails.IsKnownType, "Unknown event should have IsKnownType=false")
}

// Test: Known unregistered event Oout
func TestEventRouter_KnownUnregisteredEventRoutesToOnUnhandled(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	var unhandledCalled bool
	var unhandledEvent EventNotificationContainer
	var unhandledClient *Client
	var unhandledDetails UnhandledNotificationDetails

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		unhandledCalled = true
		unhandledEvent = notif
		unhandledClient = client
		unhandledDetails = details
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	// Don't register a handler for this known event type
	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err := router.Handle([]byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.True(t, unhandledCalled, "Unhandled handler should have been called")
	assert.NotNil(t, unhandledEvent)

	// Check if it's a known event type
	_, isKnown := unhandledEvent.(*V1BillingMeterErrorReportTriggeredEventNotification)
	assert.True(t, isKnown, "Event should be V1BillingMeterErrorReportTriggeredEventNotification")

	assert.Equal(t, "v1.billing.meter.error_report_triggered", unhandledEvent.GetEventNotification().Type)
	assert.NotNil(t, unhandledClient)
	assert.True(t, unhandledDetails.IsKnownType, "Known event should have IsKnownType=true")
}

// Test: Registered event does Oot
func TestEventRouter_RegisteredEventDoesNotCallOnUnhandled(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	var unhandledCalled bool
	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		unhandledCalled = true
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	var handlerCalled bool
	handler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		handlerCalled = true
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler)
	assert.NoError(t, err)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = router.Handle([]byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.True(t, handlerCalled, "Handler should have been called")
	assert.False(t, unhandledCalled, "Unhandled handler should not be called")
}

// Test: Handler client retains configuration
func TestEventRouter_HandlerClientRetainsConfiguration(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	backend.(*BackendImplementation).StripeContext = String("original_context_xyz")

	apiKey := "sk_test_custom_key"
	client := &Client{backend: backend, key: apiKey}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	var receivedAPIKey string
	var receivedContext *string

	handler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		receivedAPIKey = client.key
		backendImpl := client.backend.(*BackendImplementation)
		receivedContext = backendImpl.StripeContext
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler)
	assert.NoError(t, err)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err = router.Handle([]byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.Equal(t, apiKey, receivedAPIKey, "Handler should receive client with original API key")
	assert.Equal(t, "event_context_456", *receivedContext, "Handler should receive event context")
	assert.Equal(t, "original_context_xyz", *backend.(*BackendImplementation).StripeContext, "Original context should be restored")
}

// Test: on_unhandled receives correct info for unknown
func TestEventRouter_OnUnhandledReceivesCorrectInfoForUnknown(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	var details UnhandledNotificationDetails
	onUnhandled := func(notif EventNotificationContainer, client *Client, d UnhandledNotificationDetails) error {
		details = d
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	payload := unknownEventPayload()
	sigHeader := generateTestSignature(payload)
	err := router.Handle([]byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.False(t, details.IsKnownType)
}

// Test: on_unhandled receives correct info for known unregistered
func TestEventRouter_OnUnhandledReceivesCorrectInfoForKnownUnregistered(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	var details UnhandledNotificationDetails
	onUnhandled := func(notif EventNotificationContainer, client *Client, d UnhandledNotificationDetails) error {
		details = d
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	payload := v1BillingMeterPayload()
	sigHeader := generateTestSignature(payload)
	err := router.Handle([]byte(payload), sigHeader)

	assert.NoError(t, err)
	assert.True(t, details.IsKnownType)
}

// Test: Validates webhook signature
func TestEventRouter_ValidatesWebhookSignature(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	payload := v1BillingMeterPayload()
	err := router.Handle([]byte(payload), "invalid_signature")

	assert.Error(t, err)
	// Check if error is related to signature verification
	assert.Contains(t, err.Error(), "Stripe-Signature")
}

// Test: RegisteredEventTypes returns empty list
func TestEventRouter_RegisteredEventTypesEmpty(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	types := router.RegisteredEventTypes()
	assert.Empty(t, types)
}

// Test: RegisteredEventTypes returns single event type
func TestEventRouter_RegisteredEventTypesSingle(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	handler := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		return nil
	}

	err := router.OnV1BillingMeterErrorReportTriggered(handler)
	assert.NoError(t, err)

	types := router.RegisteredEventTypes()
	assert.Equal(t, []string{"v1.billing.meter.error_report_triggered"}, types)
}

// Test: RegisteredEventTypes returns multiple event types in alphabetical order
func TestEventRouter_RegisteredEventTypesMultipleAlphabetized(t *testing.T) {
	backend := GetBackendWithConfig(APIBackend, &BackendConfig{})
	client := &Client{backend: backend, key: "sk_test_1234"}

	onUnhandled := func(notif EventNotificationContainer, client *Client, details UnhandledNotificationDetails) error {
		return nil
	}

	router := NewEventRouter(client, testWebhookSecret, onUnhandled)

	handler1 := func(event *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error {
		return nil
	}
	handler2 := func(event *V1BillingMeterNoMeterFoundEventNotification, client *Client) error {
		return nil
	}
	handler3 := func(event *V2CoreAccountCreatedEventNotification, client *Client) error {
		return nil
	}

	// Register in non-alphabetical order
	err := router.OnV2CoreAccountCreated(handler3)
	assert.NoError(t, err)
	err = router.OnV1BillingMeterErrorReportTriggered(handler1)
	assert.NoError(t, err)
	err = router.OnV1BillingMeterNoMeterFound(handler2)
	assert.NoError(t, err)

	types := router.RegisteredEventTypes()
	expected := []string{
		"v1.billing.meter.error_report_triggered",
		"v1.billing.meter.no_meter_found",
		"v2.core.account.created",
	}

	assert.Equal(t, expected, types)
}
