package stripe

import (
	"context"
	"fmt"
	"sort"
)

// CallbackFunc is run when an event of a registered type is received.
type CallbackFunc = func(context.Context, EventNotificationContainer, *Client) error

// FallbackCallbackFunc is run when an event is received that does not match any registered type. It contains additional details about the unhandled event (as compared to a CallbackFunc).
type FallbackCallbackFunc = func(context.Context, EventNotificationContainer, *Client, UnhandledNotificationDetails) error

// PreHandleFunc runs after Handle() parses the payload but before any user callback fires. Returning false stops handling entirely: no callback runs. Returning a non-nil error aborts handling and is returned from Handle.
type PreHandleFunc = func(context.Context, EventNotificationContainer, *Client) (bool, error)

type UnhandledNotificationDetails struct {
	// IsKnownEventType indicates whether the unhandled event is of a known type (i.e., it has a defined struct in the SDK) or is completely unknown.
	IsKnownEventType bool
}

// eventNotificationHandlerBase holds the shared registration and dispatch machinery
// shared by the two user-facing event handlers.
type eventNotificationHandlerBase struct {
	client           *Client
	eventHandlers    map[string]CallbackFunc
	hasHandledEvent  bool
	fallbackCallback FallbackCallbackFunc
	preHandleFunc    PreHandleFunc
}

// newEventNotificationHandlerBase builds the shared handler state used by both the
// verifying and non-verifying constructors.
func newEventNotificationHandlerBase(client *Client, fallbackCallback FallbackCallbackFunc) *eventNotificationHandlerBase {
	return &eventNotificationHandlerBase{
		client:           client,
		eventHandlers:    make(map[string]CallbackFunc),
		hasHandledEvent:  false,
		fallbackCallback: fallbackCallback,
	}
}

// EventNotificationHandler routes incoming Stripe event notifications to registered handlers based on event type.
type EventNotificationHandler struct {
	*eventNotificationHandlerBase
	webhookSecret string
}

func NewEventNotificationHandler(client *Client, webhookSecret string, fallbackCallback FallbackCallbackFunc) *EventNotificationHandler {
	if webhookSecret == "" {
		panic("webhookSecret must be a non-empty string")
	}
	return &EventNotificationHandler{
		eventNotificationHandlerBase: newEventNotificationHandlerBase(client, fallbackCallback),
		webhookSecret:                webhookSecret,
	}
}

// assertCanRegister reports an error if callbacks can no longer be registered. Callbacks are
// expected to be registered once on startup, so registering anything after handling has begun
// indicates a bug.
//
// intentionally not worried about concurrency because we expect all registrations to happen
// synchronously on startup, so it'll only be read after it's done being written.
func (h *eventNotificationHandlerBase) assertCanRegister() error {
	if h.hasHandledEvent {
		return fmt.Errorf("cannot register new callbacks after an event has been handled. This is indicative of a bug.")
	}

	return nil
}

func (h *eventNotificationHandlerBase) register(eventType string, callback CallbackFunc) error {
	if err := h.assertCanRegister(); err != nil {
		return err
	}

	if h.eventHandlers[eventType] != nil {
		return fmt.Errorf("callback for event type %q is already registered", eventType)
	}

	h.eventHandlers[eventType] = callback
	return nil
}

// PreHandle registers a hook that runs after Handle() parses the payload but before any other callback fires. Returning `false` from the hook stops handling: no callback runs. Returning an error aborts handling and is returned from the original Handle() call.
func (h *eventNotificationHandlerBase) PreHandle(callback PreHandleFunc) error {
	if err := h.assertCanRegister(); err != nil {
		return err
	}

	if h.preHandleFunc != nil {
		return fmt.Errorf("a PreHandle callback is already registered")
	}

	h.preHandleFunc = callback
	return nil
}

// RegisteredEventTypes returns a sorted list of all event types with registered handlers
func (h *eventNotificationHandlerBase) RegisteredEventTypes() []string {
	types := make([]string, 0, len(h.eventHandlers))
	for eventType := range h.eventHandlers {
		types = append(types, eventType)
	}
	sort.Strings(types)
	return types
}

func registerTypedHandler[T EventNotificationContainer](
	r *eventNotificationHandlerBase,
	eventType string,
	handler func(context.Context, T, *Client) error,
) error {
	wrapper := func(ctx context.Context, notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(T)
		if !ok {
			// Use a zero value to get the type name for the error message
			var zero T
			return fmt.Errorf("failed to cast notification to %T", zero)
		}
		return handler(ctx, typedNotif, client)
	}
	return r.register(eventType, wrapper)
}

// event-handler-methods: The beginning of the section generated from our OpenAPI spec

// OnV1BillingMeterErrorReportTriggered registers a callback to handle notifications about the "v1.billing.meter.error_report_triggered" event.
func (h *eventNotificationHandlerBase) OnV1BillingMeterErrorReportTriggered(callback func(ctx context.Context, notif *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.billing.meter.error_report_triggered", callback)
}

// OnV1BillingMeterNoMeterFound registers a callback to handle notifications about the "v1.billing.meter.no_meter_found" event.
func (h *eventNotificationHandlerBase) OnV1BillingMeterNoMeterFound(callback func(ctx context.Context, notif *V1BillingMeterNoMeterFoundEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.meter.no_meter_found", callback)
}

// OnV2CommerceProductCatalogImportsFailed registers a callback to handle notifications about the "v2.commerce.product_catalog.imports.failed" event.
func (h *eventNotificationHandlerBase) OnV2CommerceProductCatalogImportsFailed(callback func(ctx context.Context, notif *V2CommerceProductCatalogImportsFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.commerce.product_catalog.imports.failed", callback)
}

// OnV2CommerceProductCatalogImportsProcessing registers a callback to handle notifications about the "v2.commerce.product_catalog.imports.processing" event.
func (h *eventNotificationHandlerBase) OnV2CommerceProductCatalogImportsProcessing(callback func(ctx context.Context, notif *V2CommerceProductCatalogImportsProcessingEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.commerce.product_catalog.imports.processing", callback)
}

// OnV2CommerceProductCatalogImportsSucceeded registers a callback to handle notifications about the "v2.commerce.product_catalog.imports.succeeded" event.
func (h *eventNotificationHandlerBase) OnV2CommerceProductCatalogImportsSucceeded(callback func(ctx context.Context, notif *V2CommerceProductCatalogImportsSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.commerce.product_catalog.imports.succeeded", callback)
}

// OnV2CommerceProductCatalogImportsSucceededWithErrors registers a callback to handle notifications about the "v2.commerce.product_catalog.imports.succeeded_with_errors" event.
func (h *eventNotificationHandlerBase) OnV2CommerceProductCatalogImportsSucceededWithErrors(callback func(ctx context.Context, notif *V2CommerceProductCatalogImportsSucceededWithErrorsEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.commerce.product_catalog.imports.succeeded_with_errors", callback)
}

// OnV2CoreAccountClosed registers a callback to handle notifications about the "v2.core.account.closed" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountClosed(callback func(ctx context.Context, notif *V2CoreAccountClosedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account.closed", callback)
}

// OnV2CoreAccountCreated registers a callback to handle notifications about the "v2.core.account.created" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountCreated(callback func(ctx context.Context, notif *V2CoreAccountCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account.created", callback)
}

// OnV2CoreAccountUpdated registers a callback to handle notifications about the "v2.core.account.updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountUpdated(callback func(ctx context.Context, notif *V2CoreAccountUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account.updated", callback)
}

// OnV2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdated registers a callback to handle notifications about the "v2.core.account[configuration.customer].capability_status_updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.customer].capability_status_updated", callback)
}

// OnV2CoreAccountIncludingConfigurationCustomerUpdated registers a callback to handle notifications about the "v2.core.account[configuration.customer].updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountIncludingConfigurationCustomerUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.customer].updated", callback)
}

// OnV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdated registers a callback to handle notifications about the "v2.core.account[configuration.merchant].capability_status_updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.merchant].capability_status_updated", callback)
}

// OnV2CoreAccountIncludingConfigurationMerchantUpdated registers a callback to handle notifications about the "v2.core.account[configuration.merchant].updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountIncludingConfigurationMerchantUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.merchant].updated", callback)
}

// OnV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdated registers a callback to handle notifications about the "v2.core.account[configuration.recipient].capability_status_updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.recipient].capability_status_updated", callback)
}

// OnV2CoreAccountIncludingConfigurationRecipientUpdated registers a callback to handle notifications about the "v2.core.account[configuration.recipient].updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountIncludingConfigurationRecipientUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.recipient].updated", callback)
}

// OnV2CoreAccountIncludingDefaultsUpdated registers a callback to handle notifications about the "v2.core.account[defaults].updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountIncludingDefaultsUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingDefaultsUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account[defaults].updated", callback)
}

// OnV2CoreAccountIncludingFutureRequirementsUpdated registers a callback to handle notifications about the "v2.core.account[future_requirements].updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountIncludingFutureRequirementsUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingFutureRequirementsUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[future_requirements].updated", callback)
}

// OnV2CoreAccountIncludingIdentityUpdated registers a callback to handle notifications about the "v2.core.account[identity].updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountIncludingIdentityUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingIdentityUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account[identity].updated", callback)
}

// OnV2CoreAccountIncludingRequirementsUpdated registers a callback to handle notifications about the "v2.core.account[requirements].updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountIncludingRequirementsUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingRequirementsUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[requirements].updated", callback)
}

// OnV2CoreAccountLinkReturned registers a callback to handle notifications about the "v2.core.account_link.returned" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountLinkReturned(callback func(ctx context.Context, notif *V2CoreAccountLinkReturnedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account_link.returned", callback)
}

// OnV2CoreAccountPersonCreated registers a callback to handle notifications about the "v2.core.account_person.created" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountPersonCreated(callback func(ctx context.Context, notif *V2CoreAccountPersonCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account_person.created", callback)
}

// OnV2CoreAccountPersonDeleted registers a callback to handle notifications about the "v2.core.account_person.deleted" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountPersonDeleted(callback func(ctx context.Context, notif *V2CoreAccountPersonDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account_person.deleted", callback)
}

// OnV2CoreAccountPersonUpdated registers a callback to handle notifications about the "v2.core.account_person.updated" event.
func (h *eventNotificationHandlerBase) OnV2CoreAccountPersonUpdated(callback func(ctx context.Context, notif *V2CoreAccountPersonUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account_person.updated", callback)
}

// OnV2CoreEventDestinationPing registers a callback to handle notifications about the "v2.core.event_destination.ping" event.
func (h *eventNotificationHandlerBase) OnV2CoreEventDestinationPing(callback func(ctx context.Context, notif *V2CoreEventDestinationPingEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.event_destination.ping", callback)
}

// event-handler-methods: The end of the section generated from our OpenAPI spec

// createClientWithContext creates a new Client with a custom stripe_context.
// It reuses the HTTPClient and other expensive resources from the base backend
// to avoid re-establishing TLS connections (Flyweight pattern).
func (h *eventNotificationHandlerBase) createClientWithContext(stripeContext *string) (*Client, error) {
	baseConfig := h.client.backends.config
	if baseConfig == nil {
		return nil, fmt.Errorf("EventNotificationHandler requires a Backend created with NewBackendsWithConfig. If you're seeing this error, please file an issue at https://github.com/stripe/stripe-go/issues")
	}
	newConfig := *baseConfig
	newConfig.StripeContext = stripeContext
	return NewClient(h.client.key, WithBackends(NewBackendsWithConfig(&newConfig))), nil
}

// Handle processes an incoming webhook payload by routing it to the appropriate CallbackFunc (or the FallbackCallbackFunc if none is available).
func (h *EventNotificationHandler) Handle(ctx context.Context, webhookBody []byte, sigHeader string) error {
	// intentionally not worried about concurrency because we expect all registrations to happen
	// synchronously on startup, so it'll only be read after it's done being written.
	h.hasHandledEvent = true

	notif, err := h.client.ParseEventNotification(webhookBody, sigHeader, h.webhookSecret)
	if err != nil {
		return err
	}

	return h.dispatch(ctx, notif)
}

func (h *eventNotificationHandlerBase) dispatch(ctx context.Context, notif EventNotificationContainer) error {
	n := notif.GetEventNotification()
	eventType := n.Type

	// Create a new client with the event's context instead of modifying the shared backend
	// This makes the code thread-safe for parallel webhook processing
	clientWithContext, err := h.createClientWithContext(n.Context.StringPtr())
	if err != nil {
		return err
	}

	if h.preHandleFunc != nil {
		shouldContinue, err := h.preHandleFunc(ctx, notif, clientWithContext)
		if err != nil {
			return err
		}
		if !shouldContinue {
			return nil
		}
	}

	callback, ok := h.eventHandlers[eventType]
	if !ok {
		_, isUnknownEventType := notif.(*UnknownEventNotification)
		details := UnhandledNotificationDetails{
			IsKnownEventType: !isUnknownEventType,
		}
		return h.fallbackCallback(ctx, notif, clientWithContext, details)
	}

	return callback(ctx, notif, clientWithContext)
}

// EventNotificationHandlerWithoutVerification routes incoming Stripe event
// notifications to registered handlers without verifying webhook signatures.
// Intended for pre-authenticated channels like AWS EventBridge, Azure Event Grid,
// or your own pre-authenticated event queue.
//
// Use NewEventNotificationHandlerWithoutVerification to create an instance.
type EventNotificationHandlerWithoutVerification struct {
	*eventNotificationHandlerBase
}

// NewEventNotificationHandlerWithoutVerification creates a handler that processes
// events without webhook signature verification.
func NewEventNotificationHandlerWithoutVerification(client *Client, fallbackCallback FallbackCallbackFunc) *EventNotificationHandlerWithoutVerification {
	return &EventNotificationHandlerWithoutVerification{
		eventNotificationHandlerBase: newEventNotificationHandlerBase(client, fallbackCallback),
	}
}

// Handle processes an incoming webhook payload without signature verification,
// routing it to the appropriate CallbackFunc (or the FallbackCallbackFunc if none is available).
func (h *EventNotificationHandlerWithoutVerification) Handle(ctx context.Context, webhookBody []byte) error {
	h.hasHandledEvent = true

	notif, err := h.client.ParseEventNotificationWithoutVerification(webhookBody)
	if err != nil {
		return err
	}

	return h.dispatch(ctx, notif)
}
