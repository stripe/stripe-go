package stripe

import (
	"fmt"
	"slices"
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

// RegisteredEventTypes returns a sorted list of all event types with registered handlers
func (r *EventRouter) RegisteredEventTypes() []string {
	types := make([]string, 0, len(r.eventHandlers))
	for eventType := range r.eventHandlers {
		types = append(types, eventType)
	}
	slices.Sort(types)
	return types
}

// event-router-methods: The beginning of the section generated from our OpenAPI spec
// On_V1BillingMeterErrorReportTriggeredEventNotification registers a handler for the "v1.billing.meter.error_report_triggered" event.
func (r *EventRouter) On_V1BillingMeterErrorReportTriggeredEventNotification(handler func(notif *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V1BillingMeterErrorReportTriggeredEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V1BillingMeterErrorReportTriggeredEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v1.billing.meter.error_report_triggered", wrapper)
}

// On_V1BillingMeterNoMeterFoundEventNotification registers a handler for the "v1.billing.meter.no_meter_found" event.
func (r *EventRouter) On_V1BillingMeterNoMeterFoundEventNotification(handler func(notif *V1BillingMeterNoMeterFoundEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V1BillingMeterNoMeterFoundEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V1BillingMeterNoMeterFoundEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v1.billing.meter.no_meter_found", wrapper)
}

// On_V2CoreAccountClosedEventNotification registers a handler for the "v2.core.account.closed" event.
func (r *EventRouter) On_V2CoreAccountClosedEventNotification(handler func(notif *V2CoreAccountClosedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountClosedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountClosedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account.closed", wrapper)
}

// On_V2CoreAccountCreatedEventNotification registers a handler for the "v2.core.account.created" event.
func (r *EventRouter) On_V2CoreAccountCreatedEventNotification(handler func(notif *V2CoreAccountCreatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountCreatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountCreatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account.created", wrapper)
}

// On_V2CoreAccountUpdatedEventNotification registers a handler for the "v2.core.account.updated" event.
func (r *EventRouter) On_V2CoreAccountUpdatedEventNotification(handler func(notif *V2CoreAccountUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account.updated", wrapper)
}

// On_V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification registers a handler for the "v2.core.account[configuration.customer].capability_status_updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.core.account[configuration.customer].capability_status_updated", wrapper)
}

// On_V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification registers a handler for the "v2.core.account[configuration.customer].updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account[configuration.customer].updated", wrapper)
}

// On_V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification registers a handler for the "v2.core.account[configuration.merchant].capability_status_updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.core.account[configuration.merchant].capability_status_updated", wrapper)
}

// On_V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification registers a handler for the "v2.core.account[configuration.merchant].updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account[configuration.merchant].updated", wrapper)
}

// On_V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification registers a handler for the "v2.core.account[configuration.recipient].capability_status_updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.core.account[configuration.recipient].capability_status_updated", wrapper)
}

// On_V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification registers a handler for the "v2.core.account[configuration.recipient].updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account[configuration.recipient].updated", wrapper)
}

// On_V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification registers a handler for the "v2.core.account[configuration.storer].capability_status_updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.core.account[configuration.storer].capability_status_updated", wrapper)
}

// On_V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification registers a handler for the "v2.core.account[configuration.storer].updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account[configuration.storer].updated", wrapper)
}

// On_V2CoreAccountIncludingDefaultsUpdatedEventNotification registers a handler for the "v2.core.account[defaults].updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingDefaultsUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingDefaultsUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingDefaultsUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingDefaultsUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account[defaults].updated", wrapper)
}

// On_V2CoreAccountIncludingIdentityUpdatedEventNotification registers a handler for the "v2.core.account[identity].updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingIdentityUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingIdentityUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingIdentityUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingIdentityUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account[identity].updated", wrapper)
}

// On_V2CoreAccountIncludingRequirementsUpdatedEventNotification registers a handler for the "v2.core.account[requirements].updated" event.
func (r *EventRouter) On_V2CoreAccountIncludingRequirementsUpdatedEventNotification(handler func(notif *V2CoreAccountIncludingRequirementsUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountIncludingRequirementsUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountIncludingRequirementsUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account[requirements].updated", wrapper)
}

// On_V2CoreAccountLinkReturnedEventNotification registers a handler for the "v2.core.account_link.returned" event.
func (r *EventRouter) On_V2CoreAccountLinkReturnedEventNotification(handler func(notif *V2CoreAccountLinkReturnedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountLinkReturnedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountLinkReturnedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account_link.returned", wrapper)
}

// On_V2CoreAccountPersonCreatedEventNotification registers a handler for the "v2.core.account_person.created" event.
func (r *EventRouter) On_V2CoreAccountPersonCreatedEventNotification(handler func(notif *V2CoreAccountPersonCreatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountPersonCreatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountPersonCreatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account_person.created", wrapper)
}

// On_V2CoreAccountPersonDeletedEventNotification registers a handler for the "v2.core.account_person.deleted" event.
func (r *EventRouter) On_V2CoreAccountPersonDeletedEventNotification(handler func(notif *V2CoreAccountPersonDeletedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountPersonDeletedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountPersonDeletedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account_person.deleted", wrapper)
}

// On_V2CoreAccountPersonUpdatedEventNotification registers a handler for the "v2.core.account_person.updated" event.
func (r *EventRouter) On_V2CoreAccountPersonUpdatedEventNotification(handler func(notif *V2CoreAccountPersonUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreAccountPersonUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreAccountPersonUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.account_person.updated", wrapper)
}

// On_V2CoreEventDestinationPingEventNotification registers a handler for the "v2.core.event_destination.ping" event.
func (r *EventRouter) On_V2CoreEventDestinationPingEventNotification(handler func(notif *V2CoreEventDestinationPingEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2CoreEventDestinationPingEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2CoreEventDestinationPingEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.core.event_destination.ping", wrapper)
}

// On_V2MoneyManagementAdjustmentCreatedEventNotification registers a handler for the "v2.money_management.adjustment.created" event.
func (r *EventRouter) On_V2MoneyManagementAdjustmentCreatedEventNotification(handler func(notif *V2MoneyManagementAdjustmentCreatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementAdjustmentCreatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementAdjustmentCreatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.adjustment.created", wrapper)
}

// On_V2MoneyManagementFinancialAccountCreatedEventNotification registers a handler for the "v2.money_management.financial_account.created" event.
func (r *EventRouter) On_V2MoneyManagementFinancialAccountCreatedEventNotification(handler func(notif *V2MoneyManagementFinancialAccountCreatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementFinancialAccountCreatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementFinancialAccountCreatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.financial_account.created", wrapper)
}

// On_V2MoneyManagementFinancialAccountUpdatedEventNotification registers a handler for the "v2.money_management.financial_account.updated" event.
func (r *EventRouter) On_V2MoneyManagementFinancialAccountUpdatedEventNotification(handler func(notif *V2MoneyManagementFinancialAccountUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementFinancialAccountUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementFinancialAccountUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.financial_account.updated", wrapper)
}

// On_V2MoneyManagementFinancialAddressActivatedEventNotification registers a handler for the "v2.money_management.financial_address.activated" event.
func (r *EventRouter) On_V2MoneyManagementFinancialAddressActivatedEventNotification(handler func(notif *V2MoneyManagementFinancialAddressActivatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementFinancialAddressActivatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementFinancialAddressActivatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.financial_address.activated", wrapper)
}

// On_V2MoneyManagementFinancialAddressFailedEventNotification registers a handler for the "v2.money_management.financial_address.failed" event.
func (r *EventRouter) On_V2MoneyManagementFinancialAddressFailedEventNotification(handler func(notif *V2MoneyManagementFinancialAddressFailedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementFinancialAddressFailedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementFinancialAddressFailedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.financial_address.failed", wrapper)
}

// On_V2MoneyManagementInboundTransferAvailableEventNotification registers a handler for the "v2.money_management.inbound_transfer.available" event.
func (r *EventRouter) On_V2MoneyManagementInboundTransferAvailableEventNotification(handler func(notif *V2MoneyManagementInboundTransferAvailableEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementInboundTransferAvailableEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementInboundTransferAvailableEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.inbound_transfer.available", wrapper)
}

// On_V2MoneyManagementInboundTransferBankDebitFailedEventNotification registers a handler for the "v2.money_management.inbound_transfer.bank_debit_failed" event.
func (r *EventRouter) On_V2MoneyManagementInboundTransferBankDebitFailedEventNotification(handler func(notif *V2MoneyManagementInboundTransferBankDebitFailedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementInboundTransferBankDebitFailedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementInboundTransferBankDebitFailedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.money_management.inbound_transfer.bank_debit_failed", wrapper)
}

// On_V2MoneyManagementInboundTransferBankDebitProcessingEventNotification registers a handler for the "v2.money_management.inbound_transfer.bank_debit_processing" event.
func (r *EventRouter) On_V2MoneyManagementInboundTransferBankDebitProcessingEventNotification(handler func(notif *V2MoneyManagementInboundTransferBankDebitProcessingEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementInboundTransferBankDebitProcessingEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementInboundTransferBankDebitProcessingEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.money_management.inbound_transfer.bank_debit_processing", wrapper)
}

// On_V2MoneyManagementInboundTransferBankDebitQueuedEventNotification registers a handler for the "v2.money_management.inbound_transfer.bank_debit_queued" event.
func (r *EventRouter) On_V2MoneyManagementInboundTransferBankDebitQueuedEventNotification(handler func(notif *V2MoneyManagementInboundTransferBankDebitQueuedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementInboundTransferBankDebitQueuedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementInboundTransferBankDebitQueuedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.money_management.inbound_transfer.bank_debit_queued", wrapper)
}

// On_V2MoneyManagementInboundTransferBankDebitReturnedEventNotification registers a handler for the "v2.money_management.inbound_transfer.bank_debit_returned" event.
func (r *EventRouter) On_V2MoneyManagementInboundTransferBankDebitReturnedEventNotification(handler func(notif *V2MoneyManagementInboundTransferBankDebitReturnedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementInboundTransferBankDebitReturnedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementInboundTransferBankDebitReturnedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.money_management.inbound_transfer.bank_debit_returned", wrapper)
}

// On_V2MoneyManagementInboundTransferBankDebitSucceededEventNotification registers a handler for the "v2.money_management.inbound_transfer.bank_debit_succeeded" event.
func (r *EventRouter) On_V2MoneyManagementInboundTransferBankDebitSucceededEventNotification(handler func(notif *V2MoneyManagementInboundTransferBankDebitSucceededEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementInboundTransferBankDebitSucceededEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementInboundTransferBankDebitSucceededEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.money_management.inbound_transfer.bank_debit_succeeded", wrapper)
}

// On_V2MoneyManagementOutboundPaymentCanceledEventNotification registers a handler for the "v2.money_management.outbound_payment.canceled" event.
func (r *EventRouter) On_V2MoneyManagementOutboundPaymentCanceledEventNotification(handler func(notif *V2MoneyManagementOutboundPaymentCanceledEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundPaymentCanceledEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundPaymentCanceledEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_payment.canceled", wrapper)
}

// On_V2MoneyManagementOutboundPaymentCreatedEventNotification registers a handler for the "v2.money_management.outbound_payment.created" event.
func (r *EventRouter) On_V2MoneyManagementOutboundPaymentCreatedEventNotification(handler func(notif *V2MoneyManagementOutboundPaymentCreatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundPaymentCreatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundPaymentCreatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_payment.created", wrapper)
}

// On_V2MoneyManagementOutboundPaymentFailedEventNotification registers a handler for the "v2.money_management.outbound_payment.failed" event.
func (r *EventRouter) On_V2MoneyManagementOutboundPaymentFailedEventNotification(handler func(notif *V2MoneyManagementOutboundPaymentFailedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundPaymentFailedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundPaymentFailedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_payment.failed", wrapper)
}

// On_V2MoneyManagementOutboundPaymentPostedEventNotification registers a handler for the "v2.money_management.outbound_payment.posted" event.
func (r *EventRouter) On_V2MoneyManagementOutboundPaymentPostedEventNotification(handler func(notif *V2MoneyManagementOutboundPaymentPostedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundPaymentPostedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundPaymentPostedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_payment.posted", wrapper)
}

// On_V2MoneyManagementOutboundPaymentReturnedEventNotification registers a handler for the "v2.money_management.outbound_payment.returned" event.
func (r *EventRouter) On_V2MoneyManagementOutboundPaymentReturnedEventNotification(handler func(notif *V2MoneyManagementOutboundPaymentReturnedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundPaymentReturnedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundPaymentReturnedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_payment.returned", wrapper)
}

// On_V2MoneyManagementOutboundPaymentUpdatedEventNotification registers a handler for the "v2.money_management.outbound_payment.updated" event.
func (r *EventRouter) On_V2MoneyManagementOutboundPaymentUpdatedEventNotification(handler func(notif *V2MoneyManagementOutboundPaymentUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundPaymentUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundPaymentUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_payment.updated", wrapper)
}

// On_V2MoneyManagementOutboundTransferCanceledEventNotification registers a handler for the "v2.money_management.outbound_transfer.canceled" event.
func (r *EventRouter) On_V2MoneyManagementOutboundTransferCanceledEventNotification(handler func(notif *V2MoneyManagementOutboundTransferCanceledEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundTransferCanceledEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundTransferCanceledEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_transfer.canceled", wrapper)
}

// On_V2MoneyManagementOutboundTransferCreatedEventNotification registers a handler for the "v2.money_management.outbound_transfer.created" event.
func (r *EventRouter) On_V2MoneyManagementOutboundTransferCreatedEventNotification(handler func(notif *V2MoneyManagementOutboundTransferCreatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundTransferCreatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundTransferCreatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_transfer.created", wrapper)
}

// On_V2MoneyManagementOutboundTransferFailedEventNotification registers a handler for the "v2.money_management.outbound_transfer.failed" event.
func (r *EventRouter) On_V2MoneyManagementOutboundTransferFailedEventNotification(handler func(notif *V2MoneyManagementOutboundTransferFailedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundTransferFailedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundTransferFailedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_transfer.failed", wrapper)
}

// On_V2MoneyManagementOutboundTransferPostedEventNotification registers a handler for the "v2.money_management.outbound_transfer.posted" event.
func (r *EventRouter) On_V2MoneyManagementOutboundTransferPostedEventNotification(handler func(notif *V2MoneyManagementOutboundTransferPostedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundTransferPostedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundTransferPostedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_transfer.posted", wrapper)
}

// On_V2MoneyManagementOutboundTransferReturnedEventNotification registers a handler for the "v2.money_management.outbound_transfer.returned" event.
func (r *EventRouter) On_V2MoneyManagementOutboundTransferReturnedEventNotification(handler func(notif *V2MoneyManagementOutboundTransferReturnedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundTransferReturnedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundTransferReturnedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_transfer.returned", wrapper)
}

// On_V2MoneyManagementOutboundTransferUpdatedEventNotification registers a handler for the "v2.money_management.outbound_transfer.updated" event.
func (r *EventRouter) On_V2MoneyManagementOutboundTransferUpdatedEventNotification(handler func(notif *V2MoneyManagementOutboundTransferUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementOutboundTransferUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementOutboundTransferUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.outbound_transfer.updated", wrapper)
}

// On_V2MoneyManagementPayoutMethodUpdatedEventNotification registers a handler for the "v2.money_management.payout_method.updated" event.
func (r *EventRouter) On_V2MoneyManagementPayoutMethodUpdatedEventNotification(handler func(notif *V2MoneyManagementPayoutMethodUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementPayoutMethodUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementPayoutMethodUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.payout_method.updated", wrapper)
}

// On_V2MoneyManagementReceivedCreditAvailableEventNotification registers a handler for the "v2.money_management.received_credit.available" event.
func (r *EventRouter) On_V2MoneyManagementReceivedCreditAvailableEventNotification(handler func(notif *V2MoneyManagementReceivedCreditAvailableEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementReceivedCreditAvailableEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementReceivedCreditAvailableEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.received_credit.available", wrapper)
}

// On_V2MoneyManagementReceivedCreditFailedEventNotification registers a handler for the "v2.money_management.received_credit.failed" event.
func (r *EventRouter) On_V2MoneyManagementReceivedCreditFailedEventNotification(handler func(notif *V2MoneyManagementReceivedCreditFailedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementReceivedCreditFailedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementReceivedCreditFailedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.received_credit.failed", wrapper)
}

// On_V2MoneyManagementReceivedCreditReturnedEventNotification registers a handler for the "v2.money_management.received_credit.returned" event.
func (r *EventRouter) On_V2MoneyManagementReceivedCreditReturnedEventNotification(handler func(notif *V2MoneyManagementReceivedCreditReturnedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementReceivedCreditReturnedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementReceivedCreditReturnedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.received_credit.returned", wrapper)
}

// On_V2MoneyManagementReceivedCreditSucceededEventNotification registers a handler for the "v2.money_management.received_credit.succeeded" event.
func (r *EventRouter) On_V2MoneyManagementReceivedCreditSucceededEventNotification(handler func(notif *V2MoneyManagementReceivedCreditSucceededEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementReceivedCreditSucceededEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementReceivedCreditSucceededEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.received_credit.succeeded", wrapper)
}

// On_V2MoneyManagementReceivedDebitCanceledEventNotification registers a handler for the "v2.money_management.received_debit.canceled" event.
func (r *EventRouter) On_V2MoneyManagementReceivedDebitCanceledEventNotification(handler func(notif *V2MoneyManagementReceivedDebitCanceledEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementReceivedDebitCanceledEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementReceivedDebitCanceledEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.received_debit.canceled", wrapper)
}

// On_V2MoneyManagementReceivedDebitFailedEventNotification registers a handler for the "v2.money_management.received_debit.failed" event.
func (r *EventRouter) On_V2MoneyManagementReceivedDebitFailedEventNotification(handler func(notif *V2MoneyManagementReceivedDebitFailedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementReceivedDebitFailedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementReceivedDebitFailedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.received_debit.failed", wrapper)
}

// On_V2MoneyManagementReceivedDebitPendingEventNotification registers a handler for the "v2.money_management.received_debit.pending" event.
func (r *EventRouter) On_V2MoneyManagementReceivedDebitPendingEventNotification(handler func(notif *V2MoneyManagementReceivedDebitPendingEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementReceivedDebitPendingEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementReceivedDebitPendingEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.received_debit.pending", wrapper)
}

// On_V2MoneyManagementReceivedDebitSucceededEventNotification registers a handler for the "v2.money_management.received_debit.succeeded" event.
func (r *EventRouter) On_V2MoneyManagementReceivedDebitSucceededEventNotification(handler func(notif *V2MoneyManagementReceivedDebitSucceededEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementReceivedDebitSucceededEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementReceivedDebitSucceededEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.received_debit.succeeded", wrapper)
}

// On_V2MoneyManagementReceivedDebitUpdatedEventNotification registers a handler for the "v2.money_management.received_debit.updated" event.
func (r *EventRouter) On_V2MoneyManagementReceivedDebitUpdatedEventNotification(handler func(notif *V2MoneyManagementReceivedDebitUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementReceivedDebitUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementReceivedDebitUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.received_debit.updated", wrapper)
}

// On_V2MoneyManagementTransactionCreatedEventNotification registers a handler for the "v2.money_management.transaction.created" event.
func (r *EventRouter) On_V2MoneyManagementTransactionCreatedEventNotification(handler func(notif *V2MoneyManagementTransactionCreatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementTransactionCreatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementTransactionCreatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.transaction.created", wrapper)
}

// On_V2MoneyManagementTransactionUpdatedEventNotification registers a handler for the "v2.money_management.transaction.updated" event.
func (r *EventRouter) On_V2MoneyManagementTransactionUpdatedEventNotification(handler func(notif *V2MoneyManagementTransactionUpdatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2MoneyManagementTransactionUpdatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2MoneyManagementTransactionUpdatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.money_management.transaction.updated", wrapper)
}

// On_V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEventNotification registers a handler for the "v2.payments.off_session_payment.authorization_attempt_failed" event.
func (r *EventRouter) On_V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEventNotification(handler func(notif *V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.payments.off_session_payment.authorization_attempt_failed", wrapper)
}

// On_V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEventNotification registers a handler for the "v2.payments.off_session_payment.authorization_attempt_started" event.
func (r *EventRouter) On_V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEventNotification(handler func(notif *V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register(
		"v2.payments.off_session_payment.authorization_attempt_started", wrapper)
}

// On_V2PaymentsOffSessionPaymentCanceledEventNotification registers a handler for the "v2.payments.off_session_payment.canceled" event.
func (r *EventRouter) On_V2PaymentsOffSessionPaymentCanceledEventNotification(handler func(notif *V2PaymentsOffSessionPaymentCanceledEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2PaymentsOffSessionPaymentCanceledEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2PaymentsOffSessionPaymentCanceledEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.payments.off_session_payment.canceled", wrapper)
}

// On_V2PaymentsOffSessionPaymentCreatedEventNotification registers a handler for the "v2.payments.off_session_payment.created" event.
func (r *EventRouter) On_V2PaymentsOffSessionPaymentCreatedEventNotification(handler func(notif *V2PaymentsOffSessionPaymentCreatedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2PaymentsOffSessionPaymentCreatedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2PaymentsOffSessionPaymentCreatedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.payments.off_session_payment.created", wrapper)
}

// On_V2PaymentsOffSessionPaymentFailedEventNotification registers a handler for the "v2.payments.off_session_payment.failed" event.
func (r *EventRouter) On_V2PaymentsOffSessionPaymentFailedEventNotification(handler func(notif *V2PaymentsOffSessionPaymentFailedEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2PaymentsOffSessionPaymentFailedEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2PaymentsOffSessionPaymentFailedEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.payments.off_session_payment.failed", wrapper)
}

// On_V2PaymentsOffSessionPaymentRequiresCaptureEventNotification registers a handler for the "v2.payments.off_session_payment.requires_capture" event.
func (r *EventRouter) On_V2PaymentsOffSessionPaymentRequiresCaptureEventNotification(handler func(notif *V2PaymentsOffSessionPaymentRequiresCaptureEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2PaymentsOffSessionPaymentRequiresCaptureEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2PaymentsOffSessionPaymentRequiresCaptureEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.payments.off_session_payment.requires_capture", wrapper)
}

// On_V2PaymentsOffSessionPaymentSucceededEventNotification registers a handler for the "v2.payments.off_session_payment.succeeded" event.
func (r *EventRouter) On_V2PaymentsOffSessionPaymentSucceededEventNotification(handler func(notif *V2PaymentsOffSessionPaymentSucceededEventNotification, client *Client) error) error {
	wrapper := func(notif EventNotificationContainer, client *Client) error {
		typedNotif, ok := notif.(*V2PaymentsOffSessionPaymentSucceededEventNotification)
		if !ok {
			return fmt.Errorf(
				"failed to cast notification to V2PaymentsOffSessionPaymentSucceededEventNotification")
		}
		return handler(typedNotif, client)
	}
	return r.register("v2.payments.off_session_payment.succeeded", wrapper)
}

// event-router-methods: The end of the section generated from our OpenAPI spec

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

	// todo: Bind event context to client; reset after
	if exists {
		return callback(notif, r.client)
	}
	_, isUnknownEventType := notif.(*UnknownEventNotification)
	details := UnhandledNotificationDetails{
		IsKnownType: !isUnknownEventType,
	}
	return r.onUnhandledHandler(notif, r.client, details)

}
