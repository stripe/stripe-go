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

// PreHandleFunc is the function type registered via PreHandle.
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

// PreHandle registers a function that will be run before any event-specific callbacks. A useful
// place to store event-agnostic logic, such as logging or checking for duplicate event deliveries
// (https://docs.stripe.com/webhooks#handle-duplicate-events).
//
// Returning true causes handling to continue as normal; returning false returns from Handle()
// immediately, so neither the registered callback nor the fallback callback are called. A non-nil
// error aborts handling and is returned from Handle().
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

// OnV1AccountApplicationAuthorized registers a callback to handle notifications about the "v1.account.application.authorized" event.
func (h *eventNotificationHandlerBase) OnV1AccountApplicationAuthorized(callback func(ctx context.Context, notif *V1AccountApplicationAuthorizedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.account.application.authorized", callback)
}

// OnV1AccountApplicationDeauthorized registers a callback to handle notifications about the "v1.account.application.deauthorized" event.
func (h *eventNotificationHandlerBase) OnV1AccountApplicationDeauthorized(callback func(ctx context.Context, notif *V1AccountApplicationDeauthorizedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.account.application.deauthorized", callback)
}

// OnV1AccountExternalAccountCreated registers a callback to handle notifications about the "v1.account.external_account.created" event.
func (h *eventNotificationHandlerBase) OnV1AccountExternalAccountCreated(callback func(ctx context.Context, notif *V1AccountExternalAccountCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.account.external_account.created", callback)
}

// OnV1AccountExternalAccountDeleted registers a callback to handle notifications about the "v1.account.external_account.deleted" event.
func (h *eventNotificationHandlerBase) OnV1AccountExternalAccountDeleted(callback func(ctx context.Context, notif *V1AccountExternalAccountDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.account.external_account.deleted", callback)
}

// OnV1AccountExternalAccountUpdated registers a callback to handle notifications about the "v1.account.external_account.updated" event.
func (h *eventNotificationHandlerBase) OnV1AccountExternalAccountUpdated(callback func(ctx context.Context, notif *V1AccountExternalAccountUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.account.external_account.updated", callback)
}

// OnV1AccountUpdated registers a callback to handle notifications about the "v1.account.updated" event.
func (h *eventNotificationHandlerBase) OnV1AccountUpdated(callback func(ctx context.Context, notif *V1AccountUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.account.updated", callback)
}

// OnV1ApplicationFeeCreated registers a callback to handle notifications about the "v1.application_fee.created" event.
func (h *eventNotificationHandlerBase) OnV1ApplicationFeeCreated(callback func(ctx context.Context, notif *V1ApplicationFeeCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.application_fee.created", callback)
}

// OnV1ApplicationFeeRefundUpdated registers a callback to handle notifications about the "v1.application_fee.refund.updated" event.
func (h *eventNotificationHandlerBase) OnV1ApplicationFeeRefundUpdated(callback func(ctx context.Context, notif *V1ApplicationFeeRefundUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.application_fee.refund.updated", callback)
}

// OnV1ApplicationFeeRefunded registers a callback to handle notifications about the "v1.application_fee.refunded" event.
func (h *eventNotificationHandlerBase) OnV1ApplicationFeeRefunded(callback func(ctx context.Context, notif *V1ApplicationFeeRefundedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.application_fee.refunded", callback)
}

// OnV1BalanceAvailable registers a callback to handle notifications about the "v1.balance.available" event.
func (h *eventNotificationHandlerBase) OnV1BalanceAvailable(callback func(ctx context.Context, notif *V1BalanceAvailableEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.balance.available", callback)
}

// OnV1BalanceSettingsUpdated registers a callback to handle notifications about the "v1.balance_settings.updated" event.
func (h *eventNotificationHandlerBase) OnV1BalanceSettingsUpdated(callback func(ctx context.Context, notif *V1BalanceSettingsUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.balance_settings.updated", callback)
}

// OnV1BillingAlertTriggered registers a callback to handle notifications about the "v1.billing.alert.triggered" event.
func (h *eventNotificationHandlerBase) OnV1BillingAlertTriggered(callback func(ctx context.Context, notif *V1BillingAlertTriggeredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.alert.triggered", callback)
}

// OnV1BillingCreditBalanceTransactionCreated registers a callback to handle notifications about the "v1.billing.credit_balance_transaction.created" event.
func (h *eventNotificationHandlerBase) OnV1BillingCreditBalanceTransactionCreated(callback func(ctx context.Context, notif *V1BillingCreditBalanceTransactionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.billing.credit_balance_transaction.created", callback)
}

// OnV1BillingCreditGrantCreated registers a callback to handle notifications about the "v1.billing.credit_grant.created" event.
func (h *eventNotificationHandlerBase) OnV1BillingCreditGrantCreated(callback func(ctx context.Context, notif *V1BillingCreditGrantCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.credit_grant.created", callback)
}

// OnV1BillingCreditGrantUpdated registers a callback to handle notifications about the "v1.billing.credit_grant.updated" event.
func (h *eventNotificationHandlerBase) OnV1BillingCreditGrantUpdated(callback func(ctx context.Context, notif *V1BillingCreditGrantUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.credit_grant.updated", callback)
}

// OnV1BillingMeterCreated registers a callback to handle notifications about the "v1.billing.meter.created" event.
func (h *eventNotificationHandlerBase) OnV1BillingMeterCreated(callback func(ctx context.Context, notif *V1BillingMeterCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.meter.created", callback)
}

// OnV1BillingMeterDeactivated registers a callback to handle notifications about the "v1.billing.meter.deactivated" event.
func (h *eventNotificationHandlerBase) OnV1BillingMeterDeactivated(callback func(ctx context.Context, notif *V1BillingMeterDeactivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.meter.deactivated", callback)
}

// OnV1BillingMeterErrorReportTriggered registers a callback to handle notifications about the "v1.billing.meter.error_report_triggered" event.
func (h *eventNotificationHandlerBase) OnV1BillingMeterErrorReportTriggered(callback func(ctx context.Context, notif *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.billing.meter.error_report_triggered", callback)
}

// OnV1BillingMeterNoMeterFound registers a callback to handle notifications about the "v1.billing.meter.no_meter_found" event.
func (h *eventNotificationHandlerBase) OnV1BillingMeterNoMeterFound(callback func(ctx context.Context, notif *V1BillingMeterNoMeterFoundEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.meter.no_meter_found", callback)
}

// OnV1BillingMeterReactivated registers a callback to handle notifications about the "v1.billing.meter.reactivated" event.
func (h *eventNotificationHandlerBase) OnV1BillingMeterReactivated(callback func(ctx context.Context, notif *V1BillingMeterReactivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.meter.reactivated", callback)
}

// OnV1BillingMeterUpdated registers a callback to handle notifications about the "v1.billing.meter.updated" event.
func (h *eventNotificationHandlerBase) OnV1BillingMeterUpdated(callback func(ctx context.Context, notif *V1BillingMeterUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.meter.updated", callback)
}

// OnV1BillingPortalConfigurationCreated registers a callback to handle notifications about the "v1.billing_portal.configuration.created" event.
func (h *eventNotificationHandlerBase) OnV1BillingPortalConfigurationCreated(callback func(ctx context.Context, notif *V1BillingPortalConfigurationCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.billing_portal.configuration.created", callback)
}

// OnV1BillingPortalConfigurationUpdated registers a callback to handle notifications about the "v1.billing_portal.configuration.updated" event.
func (h *eventNotificationHandlerBase) OnV1BillingPortalConfigurationUpdated(callback func(ctx context.Context, notif *V1BillingPortalConfigurationUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.billing_portal.configuration.updated", callback)
}

// OnV1BillingPortalSessionCreated registers a callback to handle notifications about the "v1.billing_portal.session.created" event.
func (h *eventNotificationHandlerBase) OnV1BillingPortalSessionCreated(callback func(ctx context.Context, notif *V1BillingPortalSessionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing_portal.session.created", callback)
}

// OnV1CapabilityUpdated registers a callback to handle notifications about the "v1.capability.updated" event.
func (h *eventNotificationHandlerBase) OnV1CapabilityUpdated(callback func(ctx context.Context, notif *V1CapabilityUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.capability.updated", callback)
}

// OnV1CashBalanceFundsAvailable registers a callback to handle notifications about the "v1.cash_balance.funds_available" event.
func (h *eventNotificationHandlerBase) OnV1CashBalanceFundsAvailable(callback func(ctx context.Context, notif *V1CashBalanceFundsAvailableEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.cash_balance.funds_available", callback)
}

// OnV1ChargeCaptured registers a callback to handle notifications about the "v1.charge.captured" event.
func (h *eventNotificationHandlerBase) OnV1ChargeCaptured(callback func(ctx context.Context, notif *V1ChargeCapturedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.captured", callback)
}

// OnV1ChargeDisputeClosed registers a callback to handle notifications about the "v1.charge.dispute.closed" event.
func (h *eventNotificationHandlerBase) OnV1ChargeDisputeClosed(callback func(ctx context.Context, notif *V1ChargeDisputeClosedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.dispute.closed", callback)
}

// OnV1ChargeDisputeCreated registers a callback to handle notifications about the "v1.charge.dispute.created" event.
func (h *eventNotificationHandlerBase) OnV1ChargeDisputeCreated(callback func(ctx context.Context, notif *V1ChargeDisputeCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.dispute.created", callback)
}

// OnV1ChargeDisputeFundsReinstated registers a callback to handle notifications about the "v1.charge.dispute.funds_reinstated" event.
func (h *eventNotificationHandlerBase) OnV1ChargeDisputeFundsReinstated(callback func(ctx context.Context, notif *V1ChargeDisputeFundsReinstatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.dispute.funds_reinstated", callback)
}

// OnV1ChargeDisputeFundsWithdrawn registers a callback to handle notifications about the "v1.charge.dispute.funds_withdrawn" event.
func (h *eventNotificationHandlerBase) OnV1ChargeDisputeFundsWithdrawn(callback func(ctx context.Context, notif *V1ChargeDisputeFundsWithdrawnEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.dispute.funds_withdrawn", callback)
}

// OnV1ChargeDisputeUpdated registers a callback to handle notifications about the "v1.charge.dispute.updated" event.
func (h *eventNotificationHandlerBase) OnV1ChargeDisputeUpdated(callback func(ctx context.Context, notif *V1ChargeDisputeUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.dispute.updated", callback)
}

// OnV1ChargeExpired registers a callback to handle notifications about the "v1.charge.expired" event.
func (h *eventNotificationHandlerBase) OnV1ChargeExpired(callback func(ctx context.Context, notif *V1ChargeExpiredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.expired", callback)
}

// OnV1ChargeFailed registers a callback to handle notifications about the "v1.charge.failed" event.
func (h *eventNotificationHandlerBase) OnV1ChargeFailed(callback func(ctx context.Context, notif *V1ChargeFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.failed", callback)
}

// OnV1ChargePending registers a callback to handle notifications about the "v1.charge.pending" event.
func (h *eventNotificationHandlerBase) OnV1ChargePending(callback func(ctx context.Context, notif *V1ChargePendingEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.pending", callback)
}

// OnV1ChargeRefundUpdated registers a callback to handle notifications about the "v1.charge.refund.updated" event.
func (h *eventNotificationHandlerBase) OnV1ChargeRefundUpdated(callback func(ctx context.Context, notif *V1ChargeRefundUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.refund.updated", callback)
}

// OnV1ChargeRefunded registers a callback to handle notifications about the "v1.charge.refunded" event.
func (h *eventNotificationHandlerBase) OnV1ChargeRefunded(callback func(ctx context.Context, notif *V1ChargeRefundedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.refunded", callback)
}

// OnV1ChargeSucceeded registers a callback to handle notifications about the "v1.charge.succeeded" event.
func (h *eventNotificationHandlerBase) OnV1ChargeSucceeded(callback func(ctx context.Context, notif *V1ChargeSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.succeeded", callback)
}

// OnV1ChargeUpdated registers a callback to handle notifications about the "v1.charge.updated" event.
func (h *eventNotificationHandlerBase) OnV1ChargeUpdated(callback func(ctx context.Context, notif *V1ChargeUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.updated", callback)
}

// OnV1CheckoutSessionAsyncPaymentFailed registers a callback to handle notifications about the "v1.checkout.session.async_payment_failed" event.
func (h *eventNotificationHandlerBase) OnV1CheckoutSessionAsyncPaymentFailed(callback func(ctx context.Context, notif *V1CheckoutSessionAsyncPaymentFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.checkout.session.async_payment_failed", callback)
}

// OnV1CheckoutSessionAsyncPaymentSucceeded registers a callback to handle notifications about the "v1.checkout.session.async_payment_succeeded" event.
func (h *eventNotificationHandlerBase) OnV1CheckoutSessionAsyncPaymentSucceeded(callback func(ctx context.Context, notif *V1CheckoutSessionAsyncPaymentSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.checkout.session.async_payment_succeeded", callback)
}

// OnV1CheckoutSessionCompleted registers a callback to handle notifications about the "v1.checkout.session.completed" event.
func (h *eventNotificationHandlerBase) OnV1CheckoutSessionCompleted(callback func(ctx context.Context, notif *V1CheckoutSessionCompletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.checkout.session.completed", callback)
}

// OnV1CheckoutSessionExpired registers a callback to handle notifications about the "v1.checkout.session.expired" event.
func (h *eventNotificationHandlerBase) OnV1CheckoutSessionExpired(callback func(ctx context.Context, notif *V1CheckoutSessionExpiredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.checkout.session.expired", callback)
}

// OnV1ClimateOrderCanceled registers a callback to handle notifications about the "v1.climate.order.canceled" event.
func (h *eventNotificationHandlerBase) OnV1ClimateOrderCanceled(callback func(ctx context.Context, notif *V1ClimateOrderCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.order.canceled", callback)
}

// OnV1ClimateOrderCreated registers a callback to handle notifications about the "v1.climate.order.created" event.
func (h *eventNotificationHandlerBase) OnV1ClimateOrderCreated(callback func(ctx context.Context, notif *V1ClimateOrderCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.order.created", callback)
}

// OnV1ClimateOrderDelayed registers a callback to handle notifications about the "v1.climate.order.delayed" event.
func (h *eventNotificationHandlerBase) OnV1ClimateOrderDelayed(callback func(ctx context.Context, notif *V1ClimateOrderDelayedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.order.delayed", callback)
}

// OnV1ClimateOrderDelivered registers a callback to handle notifications about the "v1.climate.order.delivered" event.
func (h *eventNotificationHandlerBase) OnV1ClimateOrderDelivered(callback func(ctx context.Context, notif *V1ClimateOrderDeliveredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.order.delivered", callback)
}

// OnV1ClimateOrderProductSubstituted registers a callback to handle notifications about the "v1.climate.order.product_substituted" event.
func (h *eventNotificationHandlerBase) OnV1ClimateOrderProductSubstituted(callback func(ctx context.Context, notif *V1ClimateOrderProductSubstitutedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.climate.order.product_substituted", callback)
}

// OnV1ClimateProductCreated registers a callback to handle notifications about the "v1.climate.product.created" event.
func (h *eventNotificationHandlerBase) OnV1ClimateProductCreated(callback func(ctx context.Context, notif *V1ClimateProductCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.product.created", callback)
}

// OnV1ClimateProductPricingUpdated registers a callback to handle notifications about the "v1.climate.product.pricing_updated" event.
func (h *eventNotificationHandlerBase) OnV1ClimateProductPricingUpdated(callback func(ctx context.Context, notif *V1ClimateProductPricingUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.product.pricing_updated", callback)
}

// OnV1CouponCreated registers a callback to handle notifications about the "v1.coupon.created" event.
func (h *eventNotificationHandlerBase) OnV1CouponCreated(callback func(ctx context.Context, notif *V1CouponCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.coupon.created", callback)
}

// OnV1CouponDeleted registers a callback to handle notifications about the "v1.coupon.deleted" event.
func (h *eventNotificationHandlerBase) OnV1CouponDeleted(callback func(ctx context.Context, notif *V1CouponDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.coupon.deleted", callback)
}

// OnV1CouponUpdated registers a callback to handle notifications about the "v1.coupon.updated" event.
func (h *eventNotificationHandlerBase) OnV1CouponUpdated(callback func(ctx context.Context, notif *V1CouponUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.coupon.updated", callback)
}

// OnV1CreditNoteCreated registers a callback to handle notifications about the "v1.credit_note.created" event.
func (h *eventNotificationHandlerBase) OnV1CreditNoteCreated(callback func(ctx context.Context, notif *V1CreditNoteCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.credit_note.created", callback)
}

// OnV1CreditNoteUpdated registers a callback to handle notifications about the "v1.credit_note.updated" event.
func (h *eventNotificationHandlerBase) OnV1CreditNoteUpdated(callback func(ctx context.Context, notif *V1CreditNoteUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.credit_note.updated", callback)
}

// OnV1CreditNoteVoided registers a callback to handle notifications about the "v1.credit_note.voided" event.
func (h *eventNotificationHandlerBase) OnV1CreditNoteVoided(callback func(ctx context.Context, notif *V1CreditNoteVoidedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.credit_note.voided", callback)
}

// OnV1CustomerCreated registers a callback to handle notifications about the "v1.customer.created" event.
func (h *eventNotificationHandlerBase) OnV1CustomerCreated(callback func(ctx context.Context, notif *V1CustomerCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.created", callback)
}

// OnV1CustomerDeleted registers a callback to handle notifications about the "v1.customer.deleted" event.
func (h *eventNotificationHandlerBase) OnV1CustomerDeleted(callback func(ctx context.Context, notif *V1CustomerDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.deleted", callback)
}

// OnV1CustomerDiscountCreated registers a callback to handle notifications about the "v1.customer.discount.created" event.
func (h *eventNotificationHandlerBase) OnV1CustomerDiscountCreated(callback func(ctx context.Context, notif *V1CustomerDiscountCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.discount.created", callback)
}

// OnV1CustomerDiscountDeleted registers a callback to handle notifications about the "v1.customer.discount.deleted" event.
func (h *eventNotificationHandlerBase) OnV1CustomerDiscountDeleted(callback func(ctx context.Context, notif *V1CustomerDiscountDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.discount.deleted", callback)
}

// OnV1CustomerDiscountUpdated registers a callback to handle notifications about the "v1.customer.discount.updated" event.
func (h *eventNotificationHandlerBase) OnV1CustomerDiscountUpdated(callback func(ctx context.Context, notif *V1CustomerDiscountUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.discount.updated", callback)
}

// OnV1CustomerSubscriptionCreated registers a callback to handle notifications about the "v1.customer.subscription.created" event.
func (h *eventNotificationHandlerBase) OnV1CustomerSubscriptionCreated(callback func(ctx context.Context, notif *V1CustomerSubscriptionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.subscription.created", callback)
}

// OnV1CustomerSubscriptionDeleted registers a callback to handle notifications about the "v1.customer.subscription.deleted" event.
func (h *eventNotificationHandlerBase) OnV1CustomerSubscriptionDeleted(callback func(ctx context.Context, notif *V1CustomerSubscriptionDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.subscription.deleted", callback)
}

// OnV1CustomerSubscriptionPaused registers a callback to handle notifications about the "v1.customer.subscription.paused" event.
func (h *eventNotificationHandlerBase) OnV1CustomerSubscriptionPaused(callback func(ctx context.Context, notif *V1CustomerSubscriptionPausedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.subscription.paused", callback)
}

// OnV1CustomerSubscriptionPendingUpdateApplied registers a callback to handle notifications about the "v1.customer.subscription.pending_update_applied" event.
func (h *eventNotificationHandlerBase) OnV1CustomerSubscriptionPendingUpdateApplied(callback func(ctx context.Context, notif *V1CustomerSubscriptionPendingUpdateAppliedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.customer.subscription.pending_update_applied", callback)
}

// OnV1CustomerSubscriptionPendingUpdateExpired registers a callback to handle notifications about the "v1.customer.subscription.pending_update_expired" event.
func (h *eventNotificationHandlerBase) OnV1CustomerSubscriptionPendingUpdateExpired(callback func(ctx context.Context, notif *V1CustomerSubscriptionPendingUpdateExpiredEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.customer.subscription.pending_update_expired", callback)
}

// OnV1CustomerSubscriptionResumed registers a callback to handle notifications about the "v1.customer.subscription.resumed" event.
func (h *eventNotificationHandlerBase) OnV1CustomerSubscriptionResumed(callback func(ctx context.Context, notif *V1CustomerSubscriptionResumedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.subscription.resumed", callback)
}

// OnV1CustomerSubscriptionTrialWillEnd registers a callback to handle notifications about the "v1.customer.subscription.trial_will_end" event.
func (h *eventNotificationHandlerBase) OnV1CustomerSubscriptionTrialWillEnd(callback func(ctx context.Context, notif *V1CustomerSubscriptionTrialWillEndEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.customer.subscription.trial_will_end", callback)
}

// OnV1CustomerSubscriptionUpdated registers a callback to handle notifications about the "v1.customer.subscription.updated" event.
func (h *eventNotificationHandlerBase) OnV1CustomerSubscriptionUpdated(callback func(ctx context.Context, notif *V1CustomerSubscriptionUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.subscription.updated", callback)
}

// OnV1CustomerTaxIDCreated registers a callback to handle notifications about the "v1.customer.tax_id.created" event.
func (h *eventNotificationHandlerBase) OnV1CustomerTaxIDCreated(callback func(ctx context.Context, notif *V1CustomerTaxIDCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.tax_id.created", callback)
}

// OnV1CustomerTaxIDDeleted registers a callback to handle notifications about the "v1.customer.tax_id.deleted" event.
func (h *eventNotificationHandlerBase) OnV1CustomerTaxIDDeleted(callback func(ctx context.Context, notif *V1CustomerTaxIDDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.tax_id.deleted", callback)
}

// OnV1CustomerTaxIDUpdated registers a callback to handle notifications about the "v1.customer.tax_id.updated" event.
func (h *eventNotificationHandlerBase) OnV1CustomerTaxIDUpdated(callback func(ctx context.Context, notif *V1CustomerTaxIDUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.tax_id.updated", callback)
}

// OnV1CustomerUpdated registers a callback to handle notifications about the "v1.customer.updated" event.
func (h *eventNotificationHandlerBase) OnV1CustomerUpdated(callback func(ctx context.Context, notif *V1CustomerUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.updated", callback)
}

// OnV1CustomerCashBalanceTransactionCreated registers a callback to handle notifications about the "v1.customer_cash_balance_transaction.created" event.
func (h *eventNotificationHandlerBase) OnV1CustomerCashBalanceTransactionCreated(callback func(ctx context.Context, notif *V1CustomerCashBalanceTransactionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.customer_cash_balance_transaction.created", callback)
}

// OnV1EntitlementsActiveEntitlementSummaryUpdated registers a callback to handle notifications about the "v1.entitlements.active_entitlement_summary.updated" event.
func (h *eventNotificationHandlerBase) OnV1EntitlementsActiveEntitlementSummaryUpdated(callback func(ctx context.Context, notif *V1EntitlementsActiveEntitlementSummaryUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.entitlements.active_entitlement_summary.updated", callback)
}

// OnV1FileCreated registers a callback to handle notifications about the "v1.file.created" event.
func (h *eventNotificationHandlerBase) OnV1FileCreated(callback func(ctx context.Context, notif *V1FileCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.file.created", callback)
}

// OnV1FinancialConnectionsAccountAccountNumbersUpdated registers a callback to handle notifications about the "v1.financial_connections.account.account_numbers_updated" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountAccountNumbersUpdated(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountAccountNumbersUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.account_numbers_updated", callback)
}

// OnV1FinancialConnectionsAccountCreated registers a callback to handle notifications about the "v1.financial_connections.account.created" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountCreated(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.created", callback)
}

// OnV1FinancialConnectionsAccountDeactivated registers a callback to handle notifications about the "v1.financial_connections.account.deactivated" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountDeactivated(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountDeactivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.deactivated", callback)
}

// OnV1FinancialConnectionsAccountDisconnected registers a callback to handle notifications about the "v1.financial_connections.account.disconnected" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountDisconnected(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountDisconnectedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.disconnected", callback)
}

// OnV1FinancialConnectionsAccountExpectedDeactivationDateUpdated registers a callback to handle notifications about the "v1.financial_connections.account.expected_deactivation_date_updated" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountExpectedDeactivationDateUpdated(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountExpectedDeactivationDateUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.expected_deactivation_date_updated", callback)
}

// OnV1FinancialConnectionsAccountReactivated registers a callback to handle notifications about the "v1.financial_connections.account.reactivated" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountReactivated(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountReactivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.reactivated", callback)
}

// OnV1FinancialConnectionsAccountRefreshedBalance registers a callback to handle notifications about the "v1.financial_connections.account.refreshed_balance" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountRefreshedBalance(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountRefreshedBalanceEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.refreshed_balance", callback)
}

// OnV1FinancialConnectionsAccountRefreshedOwnership registers a callback to handle notifications about the "v1.financial_connections.account.refreshed_ownership" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountRefreshedOwnership(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountRefreshedOwnershipEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.refreshed_ownership", callback)
}

// OnV1FinancialConnectionsAccountRefreshedTransactions registers a callback to handle notifications about the "v1.financial_connections.account.refreshed_transactions" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountRefreshedTransactions(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountRefreshedTransactionsEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.refreshed_transactions", callback)
}

// OnV1FinancialConnectionsAccountSupportedPaymentMethodTypesUpdated registers a callback to handle notifications about the "v1.financial_connections.account.supported_payment_method_types_updated" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountSupportedPaymentMethodTypesUpdated(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountSupportedPaymentMethodTypesUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.supported_payment_method_types_updated", callback)
}

// OnV1FinancialConnectionsAccountUpcomingAccountNumberExpiry registers a callback to handle notifications about the "v1.financial_connections.account.upcoming_account_number_expiry" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountUpcomingAccountNumberExpiry(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountUpcomingAccountNumberExpiryEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.upcoming_account_number_expiry", callback)
}

// OnV1FinancialConnectionsAccountUpcomingDeactivation registers a callback to handle notifications about the "v1.financial_connections.account.upcoming_deactivation" event.
func (h *eventNotificationHandlerBase) OnV1FinancialConnectionsAccountUpcomingDeactivation(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountUpcomingDeactivationEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.upcoming_deactivation", callback)
}

// OnV1IdentityVerificationSessionCanceled registers a callback to handle notifications about the "v1.identity.verification_session.canceled" event.
func (h *eventNotificationHandlerBase) OnV1IdentityVerificationSessionCanceled(callback func(ctx context.Context, notif *V1IdentityVerificationSessionCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.canceled", callback)
}

// OnV1IdentityVerificationSessionCreated registers a callback to handle notifications about the "v1.identity.verification_session.created" event.
func (h *eventNotificationHandlerBase) OnV1IdentityVerificationSessionCreated(callback func(ctx context.Context, notif *V1IdentityVerificationSessionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.created", callback)
}

// OnV1IdentityVerificationSessionProcessing registers a callback to handle notifications about the "v1.identity.verification_session.processing" event.
func (h *eventNotificationHandlerBase) OnV1IdentityVerificationSessionProcessing(callback func(ctx context.Context, notif *V1IdentityVerificationSessionProcessingEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.processing", callback)
}

// OnV1IdentityVerificationSessionRedacted registers a callback to handle notifications about the "v1.identity.verification_session.redacted" event.
func (h *eventNotificationHandlerBase) OnV1IdentityVerificationSessionRedacted(callback func(ctx context.Context, notif *V1IdentityVerificationSessionRedactedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.redacted", callback)
}

// OnV1IdentityVerificationSessionRequiresInput registers a callback to handle notifications about the "v1.identity.verification_session.requires_input" event.
func (h *eventNotificationHandlerBase) OnV1IdentityVerificationSessionRequiresInput(callback func(ctx context.Context, notif *V1IdentityVerificationSessionRequiresInputEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.requires_input", callback)
}

// OnV1IdentityVerificationSessionVerified registers a callback to handle notifications about the "v1.identity.verification_session.verified" event.
func (h *eventNotificationHandlerBase) OnV1IdentityVerificationSessionVerified(callback func(ctx context.Context, notif *V1IdentityVerificationSessionVerifiedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.verified", callback)
}

// OnV1InvoiceCreated registers a callback to handle notifications about the "v1.invoice.created" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceCreated(callback func(ctx context.Context, notif *V1InvoiceCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.created", callback)
}

// OnV1InvoiceDeleted registers a callback to handle notifications about the "v1.invoice.deleted" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceDeleted(callback func(ctx context.Context, notif *V1InvoiceDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.deleted", callback)
}

// OnV1InvoiceFinalizationFailed registers a callback to handle notifications about the "v1.invoice.finalization_failed" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceFinalizationFailed(callback func(ctx context.Context, notif *V1InvoiceFinalizationFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.finalization_failed", callback)
}

// OnV1InvoiceFinalized registers a callback to handle notifications about the "v1.invoice.finalized" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceFinalized(callback func(ctx context.Context, notif *V1InvoiceFinalizedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.finalized", callback)
}

// OnV1InvoiceMarkedUncollectible registers a callback to handle notifications about the "v1.invoice.marked_uncollectible" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceMarkedUncollectible(callback func(ctx context.Context, notif *V1InvoiceMarkedUncollectibleEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.marked_uncollectible", callback)
}

// OnV1InvoiceOverdue registers a callback to handle notifications about the "v1.invoice.overdue" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceOverdue(callback func(ctx context.Context, notif *V1InvoiceOverdueEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.overdue", callback)
}

// OnV1InvoiceOverpaid registers a callback to handle notifications about the "v1.invoice.overpaid" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceOverpaid(callback func(ctx context.Context, notif *V1InvoiceOverpaidEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.overpaid", callback)
}

// OnV1InvoicePaid registers a callback to handle notifications about the "v1.invoice.paid" event.
func (h *eventNotificationHandlerBase) OnV1InvoicePaid(callback func(ctx context.Context, notif *V1InvoicePaidEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.paid", callback)
}

// OnV1InvoicePaymentActionRequired registers a callback to handle notifications about the "v1.invoice.payment_action_required" event.
func (h *eventNotificationHandlerBase) OnV1InvoicePaymentActionRequired(callback func(ctx context.Context, notif *V1InvoicePaymentActionRequiredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.payment_action_required", callback)
}

// OnV1InvoicePaymentAttemptRequired registers a callback to handle notifications about the "v1.invoice.payment_attempt_required" event.
func (h *eventNotificationHandlerBase) OnV1InvoicePaymentAttemptRequired(callback func(ctx context.Context, notif *V1InvoicePaymentAttemptRequiredEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.invoice.payment_attempt_required", callback)
}

// OnV1InvoicePaymentFailed registers a callback to handle notifications about the "v1.invoice.payment_failed" event.
func (h *eventNotificationHandlerBase) OnV1InvoicePaymentFailed(callback func(ctx context.Context, notif *V1InvoicePaymentFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.payment_failed", callback)
}

// OnV1InvoicePaymentSucceeded registers a callback to handle notifications about the "v1.invoice.payment_succeeded" event.
func (h *eventNotificationHandlerBase) OnV1InvoicePaymentSucceeded(callback func(ctx context.Context, notif *V1InvoicePaymentSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.payment_succeeded", callback)
}

// OnV1InvoiceSent registers a callback to handle notifications about the "v1.invoice.sent" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceSent(callback func(ctx context.Context, notif *V1InvoiceSentEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.sent", callback)
}

// OnV1InvoiceUpcoming registers a callback to handle notifications about the "v1.invoice.upcoming" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceUpcoming(callback func(ctx context.Context, notif *V1InvoiceUpcomingEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.upcoming", callback)
}

// OnV1InvoiceUpdated registers a callback to handle notifications about the "v1.invoice.updated" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceUpdated(callback func(ctx context.Context, notif *V1InvoiceUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.updated", callback)
}

// OnV1InvoiceVoided registers a callback to handle notifications about the "v1.invoice.voided" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceVoided(callback func(ctx context.Context, notif *V1InvoiceVoidedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.voided", callback)
}

// OnV1InvoiceWillBeDue registers a callback to handle notifications about the "v1.invoice.will_be_due" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceWillBeDue(callback func(ctx context.Context, notif *V1InvoiceWillBeDueEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.will_be_due", callback)
}

// OnV1InvoicePaymentPaid registers a callback to handle notifications about the "v1.invoice_payment.paid" event.
func (h *eventNotificationHandlerBase) OnV1InvoicePaymentPaid(callback func(ctx context.Context, notif *V1InvoicePaymentPaidEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice_payment.paid", callback)
}

// OnV1InvoiceitemCreated registers a callback to handle notifications about the "v1.invoiceitem.created" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceitemCreated(callback func(ctx context.Context, notif *V1InvoiceitemCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoiceitem.created", callback)
}

// OnV1InvoiceitemDeleted registers a callback to handle notifications about the "v1.invoiceitem.deleted" event.
func (h *eventNotificationHandlerBase) OnV1InvoiceitemDeleted(callback func(ctx context.Context, notif *V1InvoiceitemDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoiceitem.deleted", callback)
}

// OnV1IssuingAuthorizationCreated registers a callback to handle notifications about the "v1.issuing_authorization.created" event.
func (h *eventNotificationHandlerBase) OnV1IssuingAuthorizationCreated(callback func(ctx context.Context, notif *V1IssuingAuthorizationCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_authorization.created", callback)
}

// OnV1IssuingAuthorizationRequest registers a callback to handle notifications about the "v1.issuing_authorization.request" event.
func (h *eventNotificationHandlerBase) OnV1IssuingAuthorizationRequest(callback func(ctx context.Context, notif *V1IssuingAuthorizationRequestEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_authorization.request", callback)
}

// OnV1IssuingAuthorizationUpdated registers a callback to handle notifications about the "v1.issuing_authorization.updated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingAuthorizationUpdated(callback func(ctx context.Context, notif *V1IssuingAuthorizationUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_authorization.updated", callback)
}

// OnV1IssuingCardCreated registers a callback to handle notifications about the "v1.issuing_card.created" event.
func (h *eventNotificationHandlerBase) OnV1IssuingCardCreated(callback func(ctx context.Context, notif *V1IssuingCardCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_card.created", callback)
}

// OnV1IssuingCardUpdated registers a callback to handle notifications about the "v1.issuing_card.updated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingCardUpdated(callback func(ctx context.Context, notif *V1IssuingCardUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_card.updated", callback)
}

// OnV1IssuingCardholderCreated registers a callback to handle notifications about the "v1.issuing_cardholder.created" event.
func (h *eventNotificationHandlerBase) OnV1IssuingCardholderCreated(callback func(ctx context.Context, notif *V1IssuingCardholderCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_cardholder.created", callback)
}

// OnV1IssuingCardholderUpdated registers a callback to handle notifications about the "v1.issuing_cardholder.updated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingCardholderUpdated(callback func(ctx context.Context, notif *V1IssuingCardholderUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_cardholder.updated", callback)
}

// OnV1IssuingDisputeClosed registers a callback to handle notifications about the "v1.issuing_dispute.closed" event.
func (h *eventNotificationHandlerBase) OnV1IssuingDisputeClosed(callback func(ctx context.Context, notif *V1IssuingDisputeClosedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_dispute.closed", callback)
}

// OnV1IssuingDisputeCreated registers a callback to handle notifications about the "v1.issuing_dispute.created" event.
func (h *eventNotificationHandlerBase) OnV1IssuingDisputeCreated(callback func(ctx context.Context, notif *V1IssuingDisputeCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_dispute.created", callback)
}

// OnV1IssuingDisputeFundsReinstated registers a callback to handle notifications about the "v1.issuing_dispute.funds_reinstated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingDisputeFundsReinstated(callback func(ctx context.Context, notif *V1IssuingDisputeFundsReinstatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_dispute.funds_reinstated", callback)
}

// OnV1IssuingDisputeFundsRescinded registers a callback to handle notifications about the "v1.issuing_dispute.funds_rescinded" event.
func (h *eventNotificationHandlerBase) OnV1IssuingDisputeFundsRescinded(callback func(ctx context.Context, notif *V1IssuingDisputeFundsRescindedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_dispute.funds_rescinded", callback)
}

// OnV1IssuingDisputeSubmitted registers a callback to handle notifications about the "v1.issuing_dispute.submitted" event.
func (h *eventNotificationHandlerBase) OnV1IssuingDisputeSubmitted(callback func(ctx context.Context, notif *V1IssuingDisputeSubmittedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_dispute.submitted", callback)
}

// OnV1IssuingDisputeUpdated registers a callback to handle notifications about the "v1.issuing_dispute.updated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingDisputeUpdated(callback func(ctx context.Context, notif *V1IssuingDisputeUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_dispute.updated", callback)
}

// OnV1IssuingPersonalizationDesignActivated registers a callback to handle notifications about the "v1.issuing_personalization_design.activated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingPersonalizationDesignActivated(callback func(ctx context.Context, notif *V1IssuingPersonalizationDesignActivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_personalization_design.activated", callback)
}

// OnV1IssuingPersonalizationDesignDeactivated registers a callback to handle notifications about the "v1.issuing_personalization_design.deactivated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingPersonalizationDesignDeactivated(callback func(ctx context.Context, notif *V1IssuingPersonalizationDesignDeactivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_personalization_design.deactivated", callback)
}

// OnV1IssuingPersonalizationDesignRejected registers a callback to handle notifications about the "v1.issuing_personalization_design.rejected" event.
func (h *eventNotificationHandlerBase) OnV1IssuingPersonalizationDesignRejected(callback func(ctx context.Context, notif *V1IssuingPersonalizationDesignRejectedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_personalization_design.rejected", callback)
}

// OnV1IssuingPersonalizationDesignUpdated registers a callback to handle notifications about the "v1.issuing_personalization_design.updated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingPersonalizationDesignUpdated(callback func(ctx context.Context, notif *V1IssuingPersonalizationDesignUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_personalization_design.updated", callback)
}

// OnV1IssuingTokenCreated registers a callback to handle notifications about the "v1.issuing_token.created" event.
func (h *eventNotificationHandlerBase) OnV1IssuingTokenCreated(callback func(ctx context.Context, notif *V1IssuingTokenCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_token.created", callback)
}

// OnV1IssuingTokenUpdated registers a callback to handle notifications about the "v1.issuing_token.updated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingTokenUpdated(callback func(ctx context.Context, notif *V1IssuingTokenUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_token.updated", callback)
}

// OnV1IssuingTransactionCreated registers a callback to handle notifications about the "v1.issuing_transaction.created" event.
func (h *eventNotificationHandlerBase) OnV1IssuingTransactionCreated(callback func(ctx context.Context, notif *V1IssuingTransactionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_transaction.created", callback)
}

// OnV1IssuingTransactionPurchaseDetailsReceiptUpdated registers a callback to handle notifications about the "v1.issuing_transaction.purchase_details_receipt_updated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingTransactionPurchaseDetailsReceiptUpdated(callback func(ctx context.Context, notif *V1IssuingTransactionPurchaseDetailsReceiptUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_transaction.purchase_details_receipt_updated", callback)
}

// OnV1IssuingTransactionUpdated registers a callback to handle notifications about the "v1.issuing_transaction.updated" event.
func (h *eventNotificationHandlerBase) OnV1IssuingTransactionUpdated(callback func(ctx context.Context, notif *V1IssuingTransactionUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_transaction.updated", callback)
}

// OnV1MandateUpdated registers a callback to handle notifications about the "v1.mandate.updated" event.
func (h *eventNotificationHandlerBase) OnV1MandateUpdated(callback func(ctx context.Context, notif *V1MandateUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.mandate.updated", callback)
}

// OnV1PaymentIntentAmountCapturableUpdated registers a callback to handle notifications about the "v1.payment_intent.amount_capturable_updated" event.
func (h *eventNotificationHandlerBase) OnV1PaymentIntentAmountCapturableUpdated(callback func(ctx context.Context, notif *V1PaymentIntentAmountCapturableUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.payment_intent.amount_capturable_updated", callback)
}

// OnV1PaymentIntentCanceled registers a callback to handle notifications about the "v1.payment_intent.canceled" event.
func (h *eventNotificationHandlerBase) OnV1PaymentIntentCanceled(callback func(ctx context.Context, notif *V1PaymentIntentCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.canceled", callback)
}

// OnV1PaymentIntentCreated registers a callback to handle notifications about the "v1.payment_intent.created" event.
func (h *eventNotificationHandlerBase) OnV1PaymentIntentCreated(callback func(ctx context.Context, notif *V1PaymentIntentCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.created", callback)
}

// OnV1PaymentIntentPartiallyFunded registers a callback to handle notifications about the "v1.payment_intent.partially_funded" event.
func (h *eventNotificationHandlerBase) OnV1PaymentIntentPartiallyFunded(callback func(ctx context.Context, notif *V1PaymentIntentPartiallyFundedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.partially_funded", callback)
}

// OnV1PaymentIntentPaymentFailed registers a callback to handle notifications about the "v1.payment_intent.payment_failed" event.
func (h *eventNotificationHandlerBase) OnV1PaymentIntentPaymentFailed(callback func(ctx context.Context, notif *V1PaymentIntentPaymentFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.payment_failed", callback)
}

// OnV1PaymentIntentProcessing registers a callback to handle notifications about the "v1.payment_intent.processing" event.
func (h *eventNotificationHandlerBase) OnV1PaymentIntentProcessing(callback func(ctx context.Context, notif *V1PaymentIntentProcessingEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.processing", callback)
}

// OnV1PaymentIntentRequiresAction registers a callback to handle notifications about the "v1.payment_intent.requires_action" event.
func (h *eventNotificationHandlerBase) OnV1PaymentIntentRequiresAction(callback func(ctx context.Context, notif *V1PaymentIntentRequiresActionEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.requires_action", callback)
}

// OnV1PaymentIntentSucceeded registers a callback to handle notifications about the "v1.payment_intent.succeeded" event.
func (h *eventNotificationHandlerBase) OnV1PaymentIntentSucceeded(callback func(ctx context.Context, notif *V1PaymentIntentSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.succeeded", callback)
}

// OnV1PaymentLinkCreated registers a callback to handle notifications about the "v1.payment_link.created" event.
func (h *eventNotificationHandlerBase) OnV1PaymentLinkCreated(callback func(ctx context.Context, notif *V1PaymentLinkCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_link.created", callback)
}

// OnV1PaymentLinkUpdated registers a callback to handle notifications about the "v1.payment_link.updated" event.
func (h *eventNotificationHandlerBase) OnV1PaymentLinkUpdated(callback func(ctx context.Context, notif *V1PaymentLinkUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_link.updated", callback)
}

// OnV1PaymentMethodAttached registers a callback to handle notifications about the "v1.payment_method.attached" event.
func (h *eventNotificationHandlerBase) OnV1PaymentMethodAttached(callback func(ctx context.Context, notif *V1PaymentMethodAttachedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_method.attached", callback)
}

// OnV1PaymentMethodAutomaticallyUpdated registers a callback to handle notifications about the "v1.payment_method.automatically_updated" event.
func (h *eventNotificationHandlerBase) OnV1PaymentMethodAutomaticallyUpdated(callback func(ctx context.Context, notif *V1PaymentMethodAutomaticallyUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.payment_method.automatically_updated", callback)
}

// OnV1PaymentMethodDetached registers a callback to handle notifications about the "v1.payment_method.detached" event.
func (h *eventNotificationHandlerBase) OnV1PaymentMethodDetached(callback func(ctx context.Context, notif *V1PaymentMethodDetachedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_method.detached", callback)
}

// OnV1PaymentMethodUpdated registers a callback to handle notifications about the "v1.payment_method.updated" event.
func (h *eventNotificationHandlerBase) OnV1PaymentMethodUpdated(callback func(ctx context.Context, notif *V1PaymentMethodUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_method.updated", callback)
}

// OnV1PayoutCanceled registers a callback to handle notifications about the "v1.payout.canceled" event.
func (h *eventNotificationHandlerBase) OnV1PayoutCanceled(callback func(ctx context.Context, notif *V1PayoutCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.canceled", callback)
}

// OnV1PayoutCreated registers a callback to handle notifications about the "v1.payout.created" event.
func (h *eventNotificationHandlerBase) OnV1PayoutCreated(callback func(ctx context.Context, notif *V1PayoutCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.created", callback)
}

// OnV1PayoutFailed registers a callback to handle notifications about the "v1.payout.failed" event.
func (h *eventNotificationHandlerBase) OnV1PayoutFailed(callback func(ctx context.Context, notif *V1PayoutFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.failed", callback)
}

// OnV1PayoutPaid registers a callback to handle notifications about the "v1.payout.paid" event.
func (h *eventNotificationHandlerBase) OnV1PayoutPaid(callback func(ctx context.Context, notif *V1PayoutPaidEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.paid", callback)
}

// OnV1PayoutReconciliationCompleted registers a callback to handle notifications about the "v1.payout.reconciliation_completed" event.
func (h *eventNotificationHandlerBase) OnV1PayoutReconciliationCompleted(callback func(ctx context.Context, notif *V1PayoutReconciliationCompletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.reconciliation_completed", callback)
}

// OnV1PayoutUpdated registers a callback to handle notifications about the "v1.payout.updated" event.
func (h *eventNotificationHandlerBase) OnV1PayoutUpdated(callback func(ctx context.Context, notif *V1PayoutUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.updated", callback)
}

// OnV1PersonCreated registers a callback to handle notifications about the "v1.person.created" event.
func (h *eventNotificationHandlerBase) OnV1PersonCreated(callback func(ctx context.Context, notif *V1PersonCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.person.created", callback)
}

// OnV1PersonDeleted registers a callback to handle notifications about the "v1.person.deleted" event.
func (h *eventNotificationHandlerBase) OnV1PersonDeleted(callback func(ctx context.Context, notif *V1PersonDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.person.deleted", callback)
}

// OnV1PersonUpdated registers a callback to handle notifications about the "v1.person.updated" event.
func (h *eventNotificationHandlerBase) OnV1PersonUpdated(callback func(ctx context.Context, notif *V1PersonUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.person.updated", callback)
}

// OnV1PlanCreated registers a callback to handle notifications about the "v1.plan.created" event.
func (h *eventNotificationHandlerBase) OnV1PlanCreated(callback func(ctx context.Context, notif *V1PlanCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.plan.created", callback)
}

// OnV1PlanDeleted registers a callback to handle notifications about the "v1.plan.deleted" event.
func (h *eventNotificationHandlerBase) OnV1PlanDeleted(callback func(ctx context.Context, notif *V1PlanDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.plan.deleted", callback)
}

// OnV1PlanUpdated registers a callback to handle notifications about the "v1.plan.updated" event.
func (h *eventNotificationHandlerBase) OnV1PlanUpdated(callback func(ctx context.Context, notif *V1PlanUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.plan.updated", callback)
}

// OnV1PriceCreated registers a callback to handle notifications about the "v1.price.created" event.
func (h *eventNotificationHandlerBase) OnV1PriceCreated(callback func(ctx context.Context, notif *V1PriceCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.price.created", callback)
}

// OnV1PriceDeleted registers a callback to handle notifications about the "v1.price.deleted" event.
func (h *eventNotificationHandlerBase) OnV1PriceDeleted(callback func(ctx context.Context, notif *V1PriceDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.price.deleted", callback)
}

// OnV1PriceUpdated registers a callback to handle notifications about the "v1.price.updated" event.
func (h *eventNotificationHandlerBase) OnV1PriceUpdated(callback func(ctx context.Context, notif *V1PriceUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.price.updated", callback)
}

// OnV1ProductCreated registers a callback to handle notifications about the "v1.product.created" event.
func (h *eventNotificationHandlerBase) OnV1ProductCreated(callback func(ctx context.Context, notif *V1ProductCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.product.created", callback)
}

// OnV1ProductDeleted registers a callback to handle notifications about the "v1.product.deleted" event.
func (h *eventNotificationHandlerBase) OnV1ProductDeleted(callback func(ctx context.Context, notif *V1ProductDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.product.deleted", callback)
}

// OnV1ProductUpdated registers a callback to handle notifications about the "v1.product.updated" event.
func (h *eventNotificationHandlerBase) OnV1ProductUpdated(callback func(ctx context.Context, notif *V1ProductUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.product.updated", callback)
}

// OnV1PromotionCodeCreated registers a callback to handle notifications about the "v1.promotion_code.created" event.
func (h *eventNotificationHandlerBase) OnV1PromotionCodeCreated(callback func(ctx context.Context, notif *V1PromotionCodeCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.promotion_code.created", callback)
}

// OnV1PromotionCodeUpdated registers a callback to handle notifications about the "v1.promotion_code.updated" event.
func (h *eventNotificationHandlerBase) OnV1PromotionCodeUpdated(callback func(ctx context.Context, notif *V1PromotionCodeUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.promotion_code.updated", callback)
}

// OnV1QuoteAccepted registers a callback to handle notifications about the "v1.quote.accepted" event.
func (h *eventNotificationHandlerBase) OnV1QuoteAccepted(callback func(ctx context.Context, notif *V1QuoteAcceptedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.quote.accepted", callback)
}

// OnV1QuoteCanceled registers a callback to handle notifications about the "v1.quote.canceled" event.
func (h *eventNotificationHandlerBase) OnV1QuoteCanceled(callback func(ctx context.Context, notif *V1QuoteCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.quote.canceled", callback)
}

// OnV1QuoteCreated registers a callback to handle notifications about the "v1.quote.created" event.
func (h *eventNotificationHandlerBase) OnV1QuoteCreated(callback func(ctx context.Context, notif *V1QuoteCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.quote.created", callback)
}

// OnV1QuoteFinalized registers a callback to handle notifications about the "v1.quote.finalized" event.
func (h *eventNotificationHandlerBase) OnV1QuoteFinalized(callback func(ctx context.Context, notif *V1QuoteFinalizedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.quote.finalized", callback)
}

// OnV1RadarEarlyFraudWarningCreated registers a callback to handle notifications about the "v1.radar.early_fraud_warning.created" event.
func (h *eventNotificationHandlerBase) OnV1RadarEarlyFraudWarningCreated(callback func(ctx context.Context, notif *V1RadarEarlyFraudWarningCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.radar.early_fraud_warning.created", callback)
}

// OnV1RadarEarlyFraudWarningUpdated registers a callback to handle notifications about the "v1.radar.early_fraud_warning.updated" event.
func (h *eventNotificationHandlerBase) OnV1RadarEarlyFraudWarningUpdated(callback func(ctx context.Context, notif *V1RadarEarlyFraudWarningUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.radar.early_fraud_warning.updated", callback)
}

// OnV1RefundCreated registers a callback to handle notifications about the "v1.refund.created" event.
func (h *eventNotificationHandlerBase) OnV1RefundCreated(callback func(ctx context.Context, notif *V1RefundCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.refund.created", callback)
}

// OnV1RefundFailed registers a callback to handle notifications about the "v1.refund.failed" event.
func (h *eventNotificationHandlerBase) OnV1RefundFailed(callback func(ctx context.Context, notif *V1RefundFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.refund.failed", callback)
}

// OnV1RefundUpdated registers a callback to handle notifications about the "v1.refund.updated" event.
func (h *eventNotificationHandlerBase) OnV1RefundUpdated(callback func(ctx context.Context, notif *V1RefundUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.refund.updated", callback)
}

// OnV1ReviewClosed registers a callback to handle notifications about the "v1.review.closed" event.
func (h *eventNotificationHandlerBase) OnV1ReviewClosed(callback func(ctx context.Context, notif *V1ReviewClosedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.review.closed", callback)
}

// OnV1ReviewOpened registers a callback to handle notifications about the "v1.review.opened" event.
func (h *eventNotificationHandlerBase) OnV1ReviewOpened(callback func(ctx context.Context, notif *V1ReviewOpenedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.review.opened", callback)
}

// OnV1SetupIntentCanceled registers a callback to handle notifications about the "v1.setup_intent.canceled" event.
func (h *eventNotificationHandlerBase) OnV1SetupIntentCanceled(callback func(ctx context.Context, notif *V1SetupIntentCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.setup_intent.canceled", callback)
}

// OnV1SetupIntentCreated registers a callback to handle notifications about the "v1.setup_intent.created" event.
func (h *eventNotificationHandlerBase) OnV1SetupIntentCreated(callback func(ctx context.Context, notif *V1SetupIntentCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.setup_intent.created", callback)
}

// OnV1SetupIntentRequiresAction registers a callback to handle notifications about the "v1.setup_intent.requires_action" event.
func (h *eventNotificationHandlerBase) OnV1SetupIntentRequiresAction(callback func(ctx context.Context, notif *V1SetupIntentRequiresActionEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.setup_intent.requires_action", callback)
}

// OnV1SetupIntentSetupFailed registers a callback to handle notifications about the "v1.setup_intent.setup_failed" event.
func (h *eventNotificationHandlerBase) OnV1SetupIntentSetupFailed(callback func(ctx context.Context, notif *V1SetupIntentSetupFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.setup_intent.setup_failed", callback)
}

// OnV1SetupIntentSucceeded registers a callback to handle notifications about the "v1.setup_intent.succeeded" event.
func (h *eventNotificationHandlerBase) OnV1SetupIntentSucceeded(callback func(ctx context.Context, notif *V1SetupIntentSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.setup_intent.succeeded", callback)
}

// OnV1SigmaScheduledQueryRunCreated registers a callback to handle notifications about the "v1.sigma.scheduled_query_run.created" event.
func (h *eventNotificationHandlerBase) OnV1SigmaScheduledQueryRunCreated(callback func(ctx context.Context, notif *V1SigmaScheduledQueryRunCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.sigma.scheduled_query_run.created", callback)
}

// OnV1SourceCanceled registers a callback to handle notifications about the "v1.source.canceled" event.
func (h *eventNotificationHandlerBase) OnV1SourceCanceled(callback func(ctx context.Context, notif *V1SourceCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.source.canceled", callback)
}

// OnV1SourceChargeable registers a callback to handle notifications about the "v1.source.chargeable" event.
func (h *eventNotificationHandlerBase) OnV1SourceChargeable(callback func(ctx context.Context, notif *V1SourceChargeableEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.source.chargeable", callback)
}

// OnV1SourceFailed registers a callback to handle notifications about the "v1.source.failed" event.
func (h *eventNotificationHandlerBase) OnV1SourceFailed(callback func(ctx context.Context, notif *V1SourceFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.source.failed", callback)
}

// OnV1SourceRefundAttributesRequired registers a callback to handle notifications about the "v1.source.refund_attributes_required" event.
func (h *eventNotificationHandlerBase) OnV1SourceRefundAttributesRequired(callback func(ctx context.Context, notif *V1SourceRefundAttributesRequiredEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.source.refund_attributes_required", callback)
}

// OnV1SubscriptionScheduleAborted registers a callback to handle notifications about the "v1.subscription_schedule.aborted" event.
func (h *eventNotificationHandlerBase) OnV1SubscriptionScheduleAborted(callback func(ctx context.Context, notif *V1SubscriptionScheduleAbortedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.aborted", callback)
}

// OnV1SubscriptionScheduleCanceled registers a callback to handle notifications about the "v1.subscription_schedule.canceled" event.
func (h *eventNotificationHandlerBase) OnV1SubscriptionScheduleCanceled(callback func(ctx context.Context, notif *V1SubscriptionScheduleCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.canceled", callback)
}

// OnV1SubscriptionScheduleCompleted registers a callback to handle notifications about the "v1.subscription_schedule.completed" event.
func (h *eventNotificationHandlerBase) OnV1SubscriptionScheduleCompleted(callback func(ctx context.Context, notif *V1SubscriptionScheduleCompletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.completed", callback)
}

// OnV1SubscriptionScheduleCreated registers a callback to handle notifications about the "v1.subscription_schedule.created" event.
func (h *eventNotificationHandlerBase) OnV1SubscriptionScheduleCreated(callback func(ctx context.Context, notif *V1SubscriptionScheduleCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.created", callback)
}

// OnV1SubscriptionScheduleExpiring registers a callback to handle notifications about the "v1.subscription_schedule.expiring" event.
func (h *eventNotificationHandlerBase) OnV1SubscriptionScheduleExpiring(callback func(ctx context.Context, notif *V1SubscriptionScheduleExpiringEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.expiring", callback)
}

// OnV1SubscriptionScheduleReleased registers a callback to handle notifications about the "v1.subscription_schedule.released" event.
func (h *eventNotificationHandlerBase) OnV1SubscriptionScheduleReleased(callback func(ctx context.Context, notif *V1SubscriptionScheduleReleasedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.released", callback)
}

// OnV1SubscriptionScheduleUpdated registers a callback to handle notifications about the "v1.subscription_schedule.updated" event.
func (h *eventNotificationHandlerBase) OnV1SubscriptionScheduleUpdated(callback func(ctx context.Context, notif *V1SubscriptionScheduleUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.updated", callback)
}

// OnV1TaxSettingsUpdated registers a callback to handle notifications about the "v1.tax.settings.updated" event.
func (h *eventNotificationHandlerBase) OnV1TaxSettingsUpdated(callback func(ctx context.Context, notif *V1TaxSettingsUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.tax.settings.updated", callback)
}

// OnV1TaxRateCreated registers a callback to handle notifications about the "v1.tax_rate.created" event.
func (h *eventNotificationHandlerBase) OnV1TaxRateCreated(callback func(ctx context.Context, notif *V1TaxRateCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.tax_rate.created", callback)
}

// OnV1TaxRateUpdated registers a callback to handle notifications about the "v1.tax_rate.updated" event.
func (h *eventNotificationHandlerBase) OnV1TaxRateUpdated(callback func(ctx context.Context, notif *V1TaxRateUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.tax_rate.updated", callback)
}

// OnV1TerminalReaderActionFailed registers a callback to handle notifications about the "v1.terminal.reader.action_failed" event.
func (h *eventNotificationHandlerBase) OnV1TerminalReaderActionFailed(callback func(ctx context.Context, notif *V1TerminalReaderActionFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.terminal.reader.action_failed", callback)
}

// OnV1TerminalReaderActionSucceeded registers a callback to handle notifications about the "v1.terminal.reader.action_succeeded" event.
func (h *eventNotificationHandlerBase) OnV1TerminalReaderActionSucceeded(callback func(ctx context.Context, notif *V1TerminalReaderActionSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.terminal.reader.action_succeeded", callback)
}

// OnV1TerminalReaderActionUpdated registers a callback to handle notifications about the "v1.terminal.reader.action_updated" event.
func (h *eventNotificationHandlerBase) OnV1TerminalReaderActionUpdated(callback func(ctx context.Context, notif *V1TerminalReaderActionUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.terminal.reader.action_updated", callback)
}

// OnV1TestHelpersTestClockAdvancing registers a callback to handle notifications about the "v1.test_helpers.test_clock.advancing" event.
func (h *eventNotificationHandlerBase) OnV1TestHelpersTestClockAdvancing(callback func(ctx context.Context, notif *V1TestHelpersTestClockAdvancingEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.test_helpers.test_clock.advancing", callback)
}

// OnV1TestHelpersTestClockCreated registers a callback to handle notifications about the "v1.test_helpers.test_clock.created" event.
func (h *eventNotificationHandlerBase) OnV1TestHelpersTestClockCreated(callback func(ctx context.Context, notif *V1TestHelpersTestClockCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.test_helpers.test_clock.created", callback)
}

// OnV1TestHelpersTestClockDeleted registers a callback to handle notifications about the "v1.test_helpers.test_clock.deleted" event.
func (h *eventNotificationHandlerBase) OnV1TestHelpersTestClockDeleted(callback func(ctx context.Context, notif *V1TestHelpersTestClockDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.test_helpers.test_clock.deleted", callback)
}

// OnV1TestHelpersTestClockInternalFailure registers a callback to handle notifications about the "v1.test_helpers.test_clock.internal_failure" event.
func (h *eventNotificationHandlerBase) OnV1TestHelpersTestClockInternalFailure(callback func(ctx context.Context, notif *V1TestHelpersTestClockInternalFailureEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.test_helpers.test_clock.internal_failure", callback)
}

// OnV1TestHelpersTestClockReady registers a callback to handle notifications about the "v1.test_helpers.test_clock.ready" event.
func (h *eventNotificationHandlerBase) OnV1TestHelpersTestClockReady(callback func(ctx context.Context, notif *V1TestHelpersTestClockReadyEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.test_helpers.test_clock.ready", callback)
}

// OnV1TopupCanceled registers a callback to handle notifications about the "v1.topup.canceled" event.
func (h *eventNotificationHandlerBase) OnV1TopupCanceled(callback func(ctx context.Context, notif *V1TopupCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.topup.canceled", callback)
}

// OnV1TopupCreated registers a callback to handle notifications about the "v1.topup.created" event.
func (h *eventNotificationHandlerBase) OnV1TopupCreated(callback func(ctx context.Context, notif *V1TopupCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.topup.created", callback)
}

// OnV1TopupFailed registers a callback to handle notifications about the "v1.topup.failed" event.
func (h *eventNotificationHandlerBase) OnV1TopupFailed(callback func(ctx context.Context, notif *V1TopupFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.topup.failed", callback)
}

// OnV1TopupReversed registers a callback to handle notifications about the "v1.topup.reversed" event.
func (h *eventNotificationHandlerBase) OnV1TopupReversed(callback func(ctx context.Context, notif *V1TopupReversedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.topup.reversed", callback)
}

// OnV1TopupSucceeded registers a callback to handle notifications about the "v1.topup.succeeded" event.
func (h *eventNotificationHandlerBase) OnV1TopupSucceeded(callback func(ctx context.Context, notif *V1TopupSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.topup.succeeded", callback)
}

// OnV1TransferCreated registers a callback to handle notifications about the "v1.transfer.created" event.
func (h *eventNotificationHandlerBase) OnV1TransferCreated(callback func(ctx context.Context, notif *V1TransferCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.transfer.created", callback)
}

// OnV1TransferReversed registers a callback to handle notifications about the "v1.transfer.reversed" event.
func (h *eventNotificationHandlerBase) OnV1TransferReversed(callback func(ctx context.Context, notif *V1TransferReversedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.transfer.reversed", callback)
}

// OnV1TransferUpdated registers a callback to handle notifications about the "v1.transfer.updated" event.
func (h *eventNotificationHandlerBase) OnV1TransferUpdated(callback func(ctx context.Context, notif *V1TransferUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.transfer.updated", callback)
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
