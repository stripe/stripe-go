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

type UnhandledNotificationDetails struct {
	// IsKnownEventType indicates whether the unhandled event is of a known type (i.e., it has a defined struct in the SDK) or is completely unknown.
	IsKnownEventType bool
}

// EventNotificationHandler routes incoming Stripe event notifications to registered handlers based on event type.
type EventNotificationHandler struct {
	client           *Client
	webhookSecret    string
	eventHandlers    map[string]CallbackFunc
	hasHandledEvent  bool
	fallbackCallback FallbackCallbackFunc
}

func NewEventNotificationHandler(client *Client, webhookSecret string, fallbackCallback FallbackCallbackFunc) *EventNotificationHandler {
	return &EventNotificationHandler{
		client:           client,
		webhookSecret:    webhookSecret,
		eventHandlers:    make(map[string]CallbackFunc),
		hasHandledEvent:  false,
		fallbackCallback: fallbackCallback,
	}
}

func (h *EventNotificationHandler) register(eventType string, callback CallbackFunc) error {
	// intentionally not worried about concurrency because we expect all registrations to happen
	// synchronously on startup, so it'll only be read after it's done being written.
	if h.hasHandledEvent {
		return fmt.Errorf("cannot register new event handlers after handling an event. This is indicative of a bug.")
	}

	if h.eventHandlers[eventType] != nil {
		return fmt.Errorf("handler for event type %s is already registered", eventType)
	}

	h.eventHandlers[eventType] = callback
	return nil
}

// RegisteredEventTypes returns a sorted list of all event types with registered handlers
func (h *EventNotificationHandler) RegisteredEventTypes() []string {
	types := make([]string, 0, len(h.eventHandlers))
	for eventType := range h.eventHandlers {
		types = append(types, eventType)
	}
	sort.Strings(types)
	return types
}

func registerTypedHandler[T EventNotificationContainer](
	r *EventNotificationHandler,
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
func (h *EventNotificationHandler) OnV1BillingMeterErrorReportTriggered(callback func(ctx context.Context, notif *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.billing.meter.error_report_triggered", callback)
}

// OnV1BillingMeterNoMeterFound registers a callback to handle notifications about the "v1.billing.meter.no_meter_found" event.
func (h *EventNotificationHandler) OnV1BillingMeterNoMeterFound(callback func(ctx context.Context, notif *V1BillingMeterNoMeterFoundEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.meter.no_meter_found", callback)
}

// OnV2CoreAccountClosed registers a callback to handle notifications about the "v2.core.account.closed" event.
func (h *EventNotificationHandler) OnV2CoreAccountClosed(callback func(ctx context.Context, notif *V2CoreAccountClosedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account.closed", callback)
}

// OnV2CoreAccountCreated registers a callback to handle notifications about the "v2.core.account.created" event.
func (h *EventNotificationHandler) OnV2CoreAccountCreated(callback func(ctx context.Context, notif *V2CoreAccountCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account.created", callback)
}

// OnV2CoreAccountUpdated registers a callback to handle notifications about the "v2.core.account.updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountUpdated(callback func(ctx context.Context, notif *V2CoreAccountUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account.updated", callback)
}

// OnV2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdated registers a callback to handle notifications about the "v2.core.account[configuration.customer].capability_status_updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.customer].capability_status_updated", callback)
}

// OnV2CoreAccountIncludingConfigurationCustomerUpdated registers a callback to handle notifications about the "v2.core.account[configuration.customer].updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingConfigurationCustomerUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.customer].updated", callback)
}

// OnV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdated registers a callback to handle notifications about the "v2.core.account[configuration.merchant].capability_status_updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.merchant].capability_status_updated", callback)
}

// OnV2CoreAccountIncludingConfigurationMerchantUpdated registers a callback to handle notifications about the "v2.core.account[configuration.merchant].updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingConfigurationMerchantUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.merchant].updated", callback)
}

// OnV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdated registers a callback to handle notifications about the "v2.core.account[configuration.recipient].capability_status_updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.recipient].capability_status_updated", callback)
}

// OnV2CoreAccountIncludingConfigurationRecipientUpdated registers a callback to handle notifications about the "v2.core.account[configuration.recipient].updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingConfigurationRecipientUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.recipient].updated", callback)
}

// OnV2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdated registers a callback to handle notifications about the "v2.core.account[configuration.storer].capability_status_updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.storer].capability_status_updated", callback)
}

// OnV2CoreAccountIncludingConfigurationStorerUpdated registers a callback to handle notifications about the "v2.core.account[configuration.storer].updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingConfigurationStorerUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.storer].updated", callback)
}

// OnV2CoreAccountIncludingDefaultsUpdated registers a callback to handle notifications about the "v2.core.account[defaults].updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingDefaultsUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingDefaultsUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account[defaults].updated", callback)
}

// OnV2CoreAccountIncludingIdentityUpdated registers a callback to handle notifications about the "v2.core.account[identity].updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingIdentityUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingIdentityUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account[identity].updated", callback)
}

// OnV2CoreAccountIncludingRequirementsUpdated registers a callback to handle notifications about the "v2.core.account[requirements].updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingRequirementsUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingRequirementsUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[requirements].updated", callback)
}

// OnV2CoreAccountLinkReturned registers a callback to handle notifications about the "v2.core.account_link.returned" event.
func (h *EventNotificationHandler) OnV2CoreAccountLinkReturned(callback func(ctx context.Context, notif *V2CoreAccountLinkReturnedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account_link.returned", callback)
}

// OnV2CoreAccountPersonCreated registers a callback to handle notifications about the "v2.core.account_person.created" event.
func (h *EventNotificationHandler) OnV2CoreAccountPersonCreated(callback func(ctx context.Context, notif *V2CoreAccountPersonCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account_person.created", callback)
}

// OnV2CoreAccountPersonDeleted registers a callback to handle notifications about the "v2.core.account_person.deleted" event.
func (h *EventNotificationHandler) OnV2CoreAccountPersonDeleted(callback func(ctx context.Context, notif *V2CoreAccountPersonDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account_person.deleted", callback)
}

// OnV2CoreAccountPersonUpdated registers a callback to handle notifications about the "v2.core.account_person.updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountPersonUpdated(callback func(ctx context.Context, notif *V2CoreAccountPersonUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.account_person.updated", callback)
}

// OnV2CoreEventDestinationPing registers a callback to handle notifications about the "v2.core.event_destination.ping" event.
func (h *EventNotificationHandler) OnV2CoreEventDestinationPing(callback func(ctx context.Context, notif *V2CoreEventDestinationPingEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.event_destination.ping", callback)
}

// OnV2CoreHealthEventGenerationFailureResolved registers a callback to handle notifications about the "v2.core.health.event_generation_failure.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthEventGenerationFailureResolved(callback func(ctx context.Context, notif *V2CoreHealthEventGenerationFailureResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.event_generation_failure.resolved", callback)
}

// OnV2MoneyManagementAdjustmentCreated registers a callback to handle notifications about the "v2.money_management.adjustment.created" event.
func (h *EventNotificationHandler) OnV2MoneyManagementAdjustmentCreated(callback func(ctx context.Context, notif *V2MoneyManagementAdjustmentCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.adjustment.created", callback)
}

// OnV2MoneyManagementFinancialAccountCreated registers a callback to handle notifications about the "v2.money_management.financial_account.created" event.
func (h *EventNotificationHandler) OnV2MoneyManagementFinancialAccountCreated(callback func(ctx context.Context, notif *V2MoneyManagementFinancialAccountCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.financial_account.created", callback)
}

// OnV2MoneyManagementFinancialAccountUpdated registers a callback to handle notifications about the "v2.money_management.financial_account.updated" event.
func (h *EventNotificationHandler) OnV2MoneyManagementFinancialAccountUpdated(callback func(ctx context.Context, notif *V2MoneyManagementFinancialAccountUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.financial_account.updated", callback)
}

// OnV2MoneyManagementFinancialAddressActivated registers a callback to handle notifications about the "v2.money_management.financial_address.activated" event.
func (h *EventNotificationHandler) OnV2MoneyManagementFinancialAddressActivated(callback func(ctx context.Context, notif *V2MoneyManagementFinancialAddressActivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.financial_address.activated", callback)
}

// OnV2MoneyManagementFinancialAddressFailed registers a callback to handle notifications about the "v2.money_management.financial_address.failed" event.
func (h *EventNotificationHandler) OnV2MoneyManagementFinancialAddressFailed(callback func(ctx context.Context, notif *V2MoneyManagementFinancialAddressFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.financial_address.failed", callback)
}

// OnV2MoneyManagementInboundTransferAvailable registers a callback to handle notifications about the "v2.money_management.inbound_transfer.available" event.
func (h *EventNotificationHandler) OnV2MoneyManagementInboundTransferAvailable(callback func(ctx context.Context, notif *V2MoneyManagementInboundTransferAvailableEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.inbound_transfer.available", callback)
}

// OnV2MoneyManagementInboundTransferBankDebitFailed registers a callback to handle notifications about the "v2.money_management.inbound_transfer.bank_debit_failed" event.
func (h *EventNotificationHandler) OnV2MoneyManagementInboundTransferBankDebitFailed(callback func(ctx context.Context, notif *V2MoneyManagementInboundTransferBankDebitFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.inbound_transfer.bank_debit_failed", callback)
}

// OnV2MoneyManagementInboundTransferBankDebitProcessing registers a callback to handle notifications about the "v2.money_management.inbound_transfer.bank_debit_processing" event.
func (h *EventNotificationHandler) OnV2MoneyManagementInboundTransferBankDebitProcessing(callback func(ctx context.Context, notif *V2MoneyManagementInboundTransferBankDebitProcessingEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.inbound_transfer.bank_debit_processing", callback)
}

// OnV2MoneyManagementInboundTransferBankDebitQueued registers a callback to handle notifications about the "v2.money_management.inbound_transfer.bank_debit_queued" event.
func (h *EventNotificationHandler) OnV2MoneyManagementInboundTransferBankDebitQueued(callback func(ctx context.Context, notif *V2MoneyManagementInboundTransferBankDebitQueuedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.inbound_transfer.bank_debit_queued", callback)
}

// OnV2MoneyManagementInboundTransferBankDebitReturned registers a callback to handle notifications about the "v2.money_management.inbound_transfer.bank_debit_returned" event.
func (h *EventNotificationHandler) OnV2MoneyManagementInboundTransferBankDebitReturned(callback func(ctx context.Context, notif *V2MoneyManagementInboundTransferBankDebitReturnedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.inbound_transfer.bank_debit_returned", callback)
}

// OnV2MoneyManagementInboundTransferBankDebitSucceeded registers a callback to handle notifications about the "v2.money_management.inbound_transfer.bank_debit_succeeded" event.
func (h *EventNotificationHandler) OnV2MoneyManagementInboundTransferBankDebitSucceeded(callback func(ctx context.Context, notif *V2MoneyManagementInboundTransferBankDebitSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.inbound_transfer.bank_debit_succeeded", callback)
}

// OnV2MoneyManagementOutboundPaymentCanceled registers a callback to handle notifications about the "v2.money_management.outbound_payment.canceled" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundPaymentCanceled(callback func(ctx context.Context, notif *V2MoneyManagementOutboundPaymentCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_payment.canceled", callback)
}

// OnV2MoneyManagementOutboundPaymentCreated registers a callback to handle notifications about the "v2.money_management.outbound_payment.created" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundPaymentCreated(callback func(ctx context.Context, notif *V2MoneyManagementOutboundPaymentCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_payment.created", callback)
}

// OnV2MoneyManagementOutboundPaymentFailed registers a callback to handle notifications about the "v2.money_management.outbound_payment.failed" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundPaymentFailed(callback func(ctx context.Context, notif *V2MoneyManagementOutboundPaymentFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_payment.failed", callback)
}

// OnV2MoneyManagementOutboundPaymentPosted registers a callback to handle notifications about the "v2.money_management.outbound_payment.posted" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundPaymentPosted(callback func(ctx context.Context, notif *V2MoneyManagementOutboundPaymentPostedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_payment.posted", callback)
}

// OnV2MoneyManagementOutboundPaymentReturned registers a callback to handle notifications about the "v2.money_management.outbound_payment.returned" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundPaymentReturned(callback func(ctx context.Context, notif *V2MoneyManagementOutboundPaymentReturnedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_payment.returned", callback)
}

// OnV2MoneyManagementOutboundPaymentUpdated registers a callback to handle notifications about the "v2.money_management.outbound_payment.updated" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundPaymentUpdated(callback func(ctx context.Context, notif *V2MoneyManagementOutboundPaymentUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_payment.updated", callback)
}

// OnV2MoneyManagementOutboundTransferCanceled registers a callback to handle notifications about the "v2.money_management.outbound_transfer.canceled" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundTransferCanceled(callback func(ctx context.Context, notif *V2MoneyManagementOutboundTransferCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_transfer.canceled", callback)
}

// OnV2MoneyManagementOutboundTransferCreated registers a callback to handle notifications about the "v2.money_management.outbound_transfer.created" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundTransferCreated(callback func(ctx context.Context, notif *V2MoneyManagementOutboundTransferCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_transfer.created", callback)
}

// OnV2MoneyManagementOutboundTransferFailed registers a callback to handle notifications about the "v2.money_management.outbound_transfer.failed" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundTransferFailed(callback func(ctx context.Context, notif *V2MoneyManagementOutboundTransferFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_transfer.failed", callback)
}

// OnV2MoneyManagementOutboundTransferPosted registers a callback to handle notifications about the "v2.money_management.outbound_transfer.posted" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundTransferPosted(callback func(ctx context.Context, notif *V2MoneyManagementOutboundTransferPostedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_transfer.posted", callback)
}

// OnV2MoneyManagementOutboundTransferReturned registers a callback to handle notifications about the "v2.money_management.outbound_transfer.returned" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundTransferReturned(callback func(ctx context.Context, notif *V2MoneyManagementOutboundTransferReturnedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_transfer.returned", callback)
}

// OnV2MoneyManagementOutboundTransferUpdated registers a callback to handle notifications about the "v2.money_management.outbound_transfer.updated" event.
func (h *EventNotificationHandler) OnV2MoneyManagementOutboundTransferUpdated(callback func(ctx context.Context, notif *V2MoneyManagementOutboundTransferUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.outbound_transfer.updated", callback)
}

// OnV2MoneyManagementPayoutMethodCreated registers a callback to handle notifications about the "v2.money_management.payout_method.created" event.
func (h *EventNotificationHandler) OnV2MoneyManagementPayoutMethodCreated(callback func(ctx context.Context, notif *V2MoneyManagementPayoutMethodCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.payout_method.created", callback)
}

// OnV2MoneyManagementPayoutMethodUpdated registers a callback to handle notifications about the "v2.money_management.payout_method.updated" event.
func (h *EventNotificationHandler) OnV2MoneyManagementPayoutMethodUpdated(callback func(ctx context.Context, notif *V2MoneyManagementPayoutMethodUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.payout_method.updated", callback)
}

// OnV2MoneyManagementReceivedCreditAvailable registers a callback to handle notifications about the "v2.money_management.received_credit.available" event.
func (h *EventNotificationHandler) OnV2MoneyManagementReceivedCreditAvailable(callback func(ctx context.Context, notif *V2MoneyManagementReceivedCreditAvailableEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.received_credit.available", callback)
}

// OnV2MoneyManagementReceivedCreditFailed registers a callback to handle notifications about the "v2.money_management.received_credit.failed" event.
func (h *EventNotificationHandler) OnV2MoneyManagementReceivedCreditFailed(callback func(ctx context.Context, notif *V2MoneyManagementReceivedCreditFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.received_credit.failed", callback)
}

// OnV2MoneyManagementReceivedCreditReturned registers a callback to handle notifications about the "v2.money_management.received_credit.returned" event.
func (h *EventNotificationHandler) OnV2MoneyManagementReceivedCreditReturned(callback func(ctx context.Context, notif *V2MoneyManagementReceivedCreditReturnedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.received_credit.returned", callback)
}

// OnV2MoneyManagementReceivedCreditSucceeded registers a callback to handle notifications about the "v2.money_management.received_credit.succeeded" event.
func (h *EventNotificationHandler) OnV2MoneyManagementReceivedCreditSucceeded(callback func(ctx context.Context, notif *V2MoneyManagementReceivedCreditSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.received_credit.succeeded", callback)
}

// OnV2MoneyManagementReceivedDebitCanceled registers a callback to handle notifications about the "v2.money_management.received_debit.canceled" event.
func (h *EventNotificationHandler) OnV2MoneyManagementReceivedDebitCanceled(callback func(ctx context.Context, notif *V2MoneyManagementReceivedDebitCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.received_debit.canceled", callback)
}

// OnV2MoneyManagementReceivedDebitFailed registers a callback to handle notifications about the "v2.money_management.received_debit.failed" event.
func (h *EventNotificationHandler) OnV2MoneyManagementReceivedDebitFailed(callback func(ctx context.Context, notif *V2MoneyManagementReceivedDebitFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.received_debit.failed", callback)
}

// OnV2MoneyManagementReceivedDebitPending registers a callback to handle notifications about the "v2.money_management.received_debit.pending" event.
func (h *EventNotificationHandler) OnV2MoneyManagementReceivedDebitPending(callback func(ctx context.Context, notif *V2MoneyManagementReceivedDebitPendingEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.received_debit.pending", callback)
}

// OnV2MoneyManagementReceivedDebitSucceeded registers a callback to handle notifications about the "v2.money_management.received_debit.succeeded" event.
func (h *EventNotificationHandler) OnV2MoneyManagementReceivedDebitSucceeded(callback func(ctx context.Context, notif *V2MoneyManagementReceivedDebitSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.received_debit.succeeded", callback)
}

// OnV2MoneyManagementReceivedDebitUpdated registers a callback to handle notifications about the "v2.money_management.received_debit.updated" event.
func (h *EventNotificationHandler) OnV2MoneyManagementReceivedDebitUpdated(callback func(ctx context.Context, notif *V2MoneyManagementReceivedDebitUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.received_debit.updated", callback)
}

// OnV2MoneyManagementTransactionCreated registers a callback to handle notifications about the "v2.money_management.transaction.created" event.
func (h *EventNotificationHandler) OnV2MoneyManagementTransactionCreated(callback func(ctx context.Context, notif *V2MoneyManagementTransactionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.transaction.created", callback)
}

// OnV2MoneyManagementTransactionUpdated registers a callback to handle notifications about the "v2.money_management.transaction.updated" event.
func (h *EventNotificationHandler) OnV2MoneyManagementTransactionUpdated(callback func(ctx context.Context, notif *V2MoneyManagementTransactionUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.transaction.updated", callback)
}

// event-handler-methods: The end of the section generated from our OpenAPI spec

// createClientWithContext creates a new Client with a custom stripe_context.
// It reuses the HTTPClient and other expensive resources from the base backend
// to avoid re-establishing TLS connections (Flyweight pattern).
func (h *EventNotificationHandler) createClientWithContext(stripeContext *string) (*Client, error) {
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

	n := notif.GetEventNotification()
	eventType := n.Type

	// Create a new client with the event's context instead of modifying the shared backend
	// This makes the code thread-safe for parallel webhook processing
	clientWithContext, err := h.createClientWithContext(n.Context.StringPtr())
	if err != nil {
		return err
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
