//
//
// File copied from our code generator; changes here will be overwritten.
//
//

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

// OnV1AccountApplicationAuthorized registers a callback to handle notifications about the "v1.account.application.authorized" event.
func (h *EventNotificationHandler) OnV1AccountApplicationAuthorized(callback func(ctx context.Context, notif *V1AccountApplicationAuthorizedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.account.application.authorized", callback)
}

// OnV1AccountApplicationDeauthorized registers a callback to handle notifications about the "v1.account.application.deauthorized" event.
func (h *EventNotificationHandler) OnV1AccountApplicationDeauthorized(callback func(ctx context.Context, notif *V1AccountApplicationDeauthorizedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.account.application.deauthorized", callback)
}

// OnV1AccountExternalAccountCreated registers a callback to handle notifications about the "v1.account.external_account.created" event.
func (h *EventNotificationHandler) OnV1AccountExternalAccountCreated(callback func(ctx context.Context, notif *V1AccountExternalAccountCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.account.external_account.created", callback)
}

// OnV1AccountExternalAccountDeleted registers a callback to handle notifications about the "v1.account.external_account.deleted" event.
func (h *EventNotificationHandler) OnV1AccountExternalAccountDeleted(callback func(ctx context.Context, notif *V1AccountExternalAccountDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.account.external_account.deleted", callback)
}

// OnV1AccountExternalAccountUpdated registers a callback to handle notifications about the "v1.account.external_account.updated" event.
func (h *EventNotificationHandler) OnV1AccountExternalAccountUpdated(callback func(ctx context.Context, notif *V1AccountExternalAccountUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.account.external_account.updated", callback)
}

// OnV1AccountUpdated registers a callback to handle notifications about the "v1.account.updated" event.
func (h *EventNotificationHandler) OnV1AccountUpdated(callback func(ctx context.Context, notif *V1AccountUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.account.updated", callback)
}

// OnV1AccountSignalsIncludingDelinquencyCreated registers a callback to handle notifications about the "v1.account_signals[delinquency].created" event.
func (h *EventNotificationHandler) OnV1AccountSignalsIncludingDelinquencyCreated(callback func(ctx context.Context, notif *V1AccountSignalsIncludingDelinquencyCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.account_signals[delinquency].created", callback)
}

// OnV1ApplicationFeeCreated registers a callback to handle notifications about the "v1.application_fee.created" event.
func (h *EventNotificationHandler) OnV1ApplicationFeeCreated(callback func(ctx context.Context, notif *V1ApplicationFeeCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.application_fee.created", callback)
}

// OnV1ApplicationFeeRefundUpdated registers a callback to handle notifications about the "v1.application_fee.refund.updated" event.
func (h *EventNotificationHandler) OnV1ApplicationFeeRefundUpdated(callback func(ctx context.Context, notif *V1ApplicationFeeRefundUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.application_fee.refund.updated", callback)
}

// OnV1ApplicationFeeRefunded registers a callback to handle notifications about the "v1.application_fee.refunded" event.
func (h *EventNotificationHandler) OnV1ApplicationFeeRefunded(callback func(ctx context.Context, notif *V1ApplicationFeeRefundedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.application_fee.refunded", callback)
}

// OnV1BalanceAvailable registers a callback to handle notifications about the "v1.balance.available" event.
func (h *EventNotificationHandler) OnV1BalanceAvailable(callback func(ctx context.Context, notif *V1BalanceAvailableEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.balance.available", callback)
}

// OnV1BillingAlertTriggered registers a callback to handle notifications about the "v1.billing.alert.triggered" event.
func (h *EventNotificationHandler) OnV1BillingAlertTriggered(callback func(ctx context.Context, notif *V1BillingAlertTriggeredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.alert.triggered", callback)
}

// OnV1BillingMeterErrorReportTriggered registers a callback to handle notifications about the "v1.billing.meter.error_report_triggered" event.
func (h *EventNotificationHandler) OnV1BillingMeterErrorReportTriggered(callback func(ctx context.Context, notif *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.billing.meter.error_report_triggered", callback)
}

// OnV1BillingMeterNoMeterFound registers a callback to handle notifications about the "v1.billing.meter.no_meter_found" event.
func (h *EventNotificationHandler) OnV1BillingMeterNoMeterFound(callback func(ctx context.Context, notif *V1BillingMeterNoMeterFoundEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing.meter.no_meter_found", callback)
}

// OnV1BillingPortalConfigurationCreated registers a callback to handle notifications about the "v1.billing_portal.configuration.created" event.
func (h *EventNotificationHandler) OnV1BillingPortalConfigurationCreated(callback func(ctx context.Context, notif *V1BillingPortalConfigurationCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.billing_portal.configuration.created", callback)
}

// OnV1BillingPortalConfigurationUpdated registers a callback to handle notifications about the "v1.billing_portal.configuration.updated" event.
func (h *EventNotificationHandler) OnV1BillingPortalConfigurationUpdated(callback func(ctx context.Context, notif *V1BillingPortalConfigurationUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.billing_portal.configuration.updated", callback)
}

// OnV1BillingPortalSessionCreated registers a callback to handle notifications about the "v1.billing_portal.session.created" event.
func (h *EventNotificationHandler) OnV1BillingPortalSessionCreated(callback func(ctx context.Context, notif *V1BillingPortalSessionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.billing_portal.session.created", callback)
}

// OnV1CapabilityUpdated registers a callback to handle notifications about the "v1.capability.updated" event.
func (h *EventNotificationHandler) OnV1CapabilityUpdated(callback func(ctx context.Context, notif *V1CapabilityUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.capability.updated", callback)
}

// OnV1CashBalanceFundsAvailable registers a callback to handle notifications about the "v1.cash_balance.funds_available" event.
func (h *EventNotificationHandler) OnV1CashBalanceFundsAvailable(callback func(ctx context.Context, notif *V1CashBalanceFundsAvailableEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.cash_balance.funds_available", callback)
}

// OnV1ChargeCaptured registers a callback to handle notifications about the "v1.charge.captured" event.
func (h *EventNotificationHandler) OnV1ChargeCaptured(callback func(ctx context.Context, notif *V1ChargeCapturedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.captured", callback)
}

// OnV1ChargeDisputeClosed registers a callback to handle notifications about the "v1.charge.dispute.closed" event.
func (h *EventNotificationHandler) OnV1ChargeDisputeClosed(callback func(ctx context.Context, notif *V1ChargeDisputeClosedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.dispute.closed", callback)
}

// OnV1ChargeDisputeCreated registers a callback to handle notifications about the "v1.charge.dispute.created" event.
func (h *EventNotificationHandler) OnV1ChargeDisputeCreated(callback func(ctx context.Context, notif *V1ChargeDisputeCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.dispute.created", callback)
}

// OnV1ChargeDisputeFundsReinstated registers a callback to handle notifications about the "v1.charge.dispute.funds_reinstated" event.
func (h *EventNotificationHandler) OnV1ChargeDisputeFundsReinstated(callback func(ctx context.Context, notif *V1ChargeDisputeFundsReinstatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.dispute.funds_reinstated", callback)
}

// OnV1ChargeDisputeFundsWithdrawn registers a callback to handle notifications about the "v1.charge.dispute.funds_withdrawn" event.
func (h *EventNotificationHandler) OnV1ChargeDisputeFundsWithdrawn(callback func(ctx context.Context, notif *V1ChargeDisputeFundsWithdrawnEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.dispute.funds_withdrawn", callback)
}

// OnV1ChargeDisputeUpdated registers a callback to handle notifications about the "v1.charge.dispute.updated" event.
func (h *EventNotificationHandler) OnV1ChargeDisputeUpdated(callback func(ctx context.Context, notif *V1ChargeDisputeUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.dispute.updated", callback)
}

// OnV1ChargeExpired registers a callback to handle notifications about the "v1.charge.expired" event.
func (h *EventNotificationHandler) OnV1ChargeExpired(callback func(ctx context.Context, notif *V1ChargeExpiredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.expired", callback)
}

// OnV1ChargeFailed registers a callback to handle notifications about the "v1.charge.failed" event.
func (h *EventNotificationHandler) OnV1ChargeFailed(callback func(ctx context.Context, notif *V1ChargeFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.failed", callback)
}

// OnV1ChargePending registers a callback to handle notifications about the "v1.charge.pending" event.
func (h *EventNotificationHandler) OnV1ChargePending(callback func(ctx context.Context, notif *V1ChargePendingEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.pending", callback)
}

// OnV1ChargeRefundUpdated registers a callback to handle notifications about the "v1.charge.refund.updated" event.
func (h *EventNotificationHandler) OnV1ChargeRefundUpdated(callback func(ctx context.Context, notif *V1ChargeRefundUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.refund.updated", callback)
}

// OnV1ChargeRefunded registers a callback to handle notifications about the "v1.charge.refunded" event.
func (h *EventNotificationHandler) OnV1ChargeRefunded(callback func(ctx context.Context, notif *V1ChargeRefundedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.refunded", callback)
}

// OnV1ChargeSucceeded registers a callback to handle notifications about the "v1.charge.succeeded" event.
func (h *EventNotificationHandler) OnV1ChargeSucceeded(callback func(ctx context.Context, notif *V1ChargeSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.succeeded", callback)
}

// OnV1ChargeUpdated registers a callback to handle notifications about the "v1.charge.updated" event.
func (h *EventNotificationHandler) OnV1ChargeUpdated(callback func(ctx context.Context, notif *V1ChargeUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.charge.updated", callback)
}

// OnV1CheckoutSessionAsyncPaymentFailed registers a callback to handle notifications about the "v1.checkout.session.async_payment_failed" event.
func (h *EventNotificationHandler) OnV1CheckoutSessionAsyncPaymentFailed(callback func(ctx context.Context, notif *V1CheckoutSessionAsyncPaymentFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.checkout.session.async_payment_failed", callback)
}

// OnV1CheckoutSessionAsyncPaymentSucceeded registers a callback to handle notifications about the "v1.checkout.session.async_payment_succeeded" event.
func (h *EventNotificationHandler) OnV1CheckoutSessionAsyncPaymentSucceeded(callback func(ctx context.Context, notif *V1CheckoutSessionAsyncPaymentSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.checkout.session.async_payment_succeeded", callback)
}

// OnV1CheckoutSessionCompleted registers a callback to handle notifications about the "v1.checkout.session.completed" event.
func (h *EventNotificationHandler) OnV1CheckoutSessionCompleted(callback func(ctx context.Context, notif *V1CheckoutSessionCompletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.checkout.session.completed", callback)
}

// OnV1CheckoutSessionExpired registers a callback to handle notifications about the "v1.checkout.session.expired" event.
func (h *EventNotificationHandler) OnV1CheckoutSessionExpired(callback func(ctx context.Context, notif *V1CheckoutSessionExpiredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.checkout.session.expired", callback)
}

// OnV1ClimateOrderCanceled registers a callback to handle notifications about the "v1.climate.order.canceled" event.
func (h *EventNotificationHandler) OnV1ClimateOrderCanceled(callback func(ctx context.Context, notif *V1ClimateOrderCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.order.canceled", callback)
}

// OnV1ClimateOrderCreated registers a callback to handle notifications about the "v1.climate.order.created" event.
func (h *EventNotificationHandler) OnV1ClimateOrderCreated(callback func(ctx context.Context, notif *V1ClimateOrderCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.order.created", callback)
}

// OnV1ClimateOrderDelayed registers a callback to handle notifications about the "v1.climate.order.delayed" event.
func (h *EventNotificationHandler) OnV1ClimateOrderDelayed(callback func(ctx context.Context, notif *V1ClimateOrderDelayedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.order.delayed", callback)
}

// OnV1ClimateOrderDelivered registers a callback to handle notifications about the "v1.climate.order.delivered" event.
func (h *EventNotificationHandler) OnV1ClimateOrderDelivered(callback func(ctx context.Context, notif *V1ClimateOrderDeliveredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.order.delivered", callback)
}

// OnV1ClimateOrderProductSubstituted registers a callback to handle notifications about the "v1.climate.order.product_substituted" event.
func (h *EventNotificationHandler) OnV1ClimateOrderProductSubstituted(callback func(ctx context.Context, notif *V1ClimateOrderProductSubstitutedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.climate.order.product_substituted", callback)
}

// OnV1ClimateProductCreated registers a callback to handle notifications about the "v1.climate.product.created" event.
func (h *EventNotificationHandler) OnV1ClimateProductCreated(callback func(ctx context.Context, notif *V1ClimateProductCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.product.created", callback)
}

// OnV1ClimateProductPricingUpdated registers a callback to handle notifications about the "v1.climate.product.pricing_updated" event.
func (h *EventNotificationHandler) OnV1ClimateProductPricingUpdated(callback func(ctx context.Context, notif *V1ClimateProductPricingUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.climate.product.pricing_updated", callback)
}

// OnV1CouponCreated registers a callback to handle notifications about the "v1.coupon.created" event.
func (h *EventNotificationHandler) OnV1CouponCreated(callback func(ctx context.Context, notif *V1CouponCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.coupon.created", callback)
}

// OnV1CouponDeleted registers a callback to handle notifications about the "v1.coupon.deleted" event.
func (h *EventNotificationHandler) OnV1CouponDeleted(callback func(ctx context.Context, notif *V1CouponDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.coupon.deleted", callback)
}

// OnV1CouponUpdated registers a callback to handle notifications about the "v1.coupon.updated" event.
func (h *EventNotificationHandler) OnV1CouponUpdated(callback func(ctx context.Context, notif *V1CouponUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.coupon.updated", callback)
}

// OnV1CreditNoteCreated registers a callback to handle notifications about the "v1.credit_note.created" event.
func (h *EventNotificationHandler) OnV1CreditNoteCreated(callback func(ctx context.Context, notif *V1CreditNoteCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.credit_note.created", callback)
}

// OnV1CreditNoteUpdated registers a callback to handle notifications about the "v1.credit_note.updated" event.
func (h *EventNotificationHandler) OnV1CreditNoteUpdated(callback func(ctx context.Context, notif *V1CreditNoteUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.credit_note.updated", callback)
}

// OnV1CreditNoteVoided registers a callback to handle notifications about the "v1.credit_note.voided" event.
func (h *EventNotificationHandler) OnV1CreditNoteVoided(callback func(ctx context.Context, notif *V1CreditNoteVoidedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.credit_note.voided", callback)
}

// OnV1CustomerCreated registers a callback to handle notifications about the "v1.customer.created" event.
func (h *EventNotificationHandler) OnV1CustomerCreated(callback func(ctx context.Context, notif *V1CustomerCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.created", callback)
}

// OnV1CustomerDeleted registers a callback to handle notifications about the "v1.customer.deleted" event.
func (h *EventNotificationHandler) OnV1CustomerDeleted(callback func(ctx context.Context, notif *V1CustomerDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.deleted", callback)
}

// OnV1CustomerSubscriptionCreated registers a callback to handle notifications about the "v1.customer.subscription.created" event.
func (h *EventNotificationHandler) OnV1CustomerSubscriptionCreated(callback func(ctx context.Context, notif *V1CustomerSubscriptionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.subscription.created", callback)
}

// OnV1CustomerSubscriptionDeleted registers a callback to handle notifications about the "v1.customer.subscription.deleted" event.
func (h *EventNotificationHandler) OnV1CustomerSubscriptionDeleted(callback func(ctx context.Context, notif *V1CustomerSubscriptionDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.subscription.deleted", callback)
}

// OnV1CustomerSubscriptionPaused registers a callback to handle notifications about the "v1.customer.subscription.paused" event.
func (h *EventNotificationHandler) OnV1CustomerSubscriptionPaused(callback func(ctx context.Context, notif *V1CustomerSubscriptionPausedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.subscription.paused", callback)
}

// OnV1CustomerSubscriptionPendingUpdateApplied registers a callback to handle notifications about the "v1.customer.subscription.pending_update_applied" event.
func (h *EventNotificationHandler) OnV1CustomerSubscriptionPendingUpdateApplied(callback func(ctx context.Context, notif *V1CustomerSubscriptionPendingUpdateAppliedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.customer.subscription.pending_update_applied", callback)
}

// OnV1CustomerSubscriptionPendingUpdateExpired registers a callback to handle notifications about the "v1.customer.subscription.pending_update_expired" event.
func (h *EventNotificationHandler) OnV1CustomerSubscriptionPendingUpdateExpired(callback func(ctx context.Context, notif *V1CustomerSubscriptionPendingUpdateExpiredEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.customer.subscription.pending_update_expired", callback)
}

// OnV1CustomerSubscriptionResumed registers a callback to handle notifications about the "v1.customer.subscription.resumed" event.
func (h *EventNotificationHandler) OnV1CustomerSubscriptionResumed(callback func(ctx context.Context, notif *V1CustomerSubscriptionResumedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.subscription.resumed", callback)
}

// OnV1CustomerSubscriptionTrialWillEnd registers a callback to handle notifications about the "v1.customer.subscription.trial_will_end" event.
func (h *EventNotificationHandler) OnV1CustomerSubscriptionTrialWillEnd(callback func(ctx context.Context, notif *V1CustomerSubscriptionTrialWillEndEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.customer.subscription.trial_will_end", callback)
}

// OnV1CustomerSubscriptionUpdated registers a callback to handle notifications about the "v1.customer.subscription.updated" event.
func (h *EventNotificationHandler) OnV1CustomerSubscriptionUpdated(callback func(ctx context.Context, notif *V1CustomerSubscriptionUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.subscription.updated", callback)
}

// OnV1CustomerTaxIDCreated registers a callback to handle notifications about the "v1.customer.tax_id.created" event.
func (h *EventNotificationHandler) OnV1CustomerTaxIDCreated(callback func(ctx context.Context, notif *V1CustomerTaxIDCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.tax_id.created", callback)
}

// OnV1CustomerTaxIDDeleted registers a callback to handle notifications about the "v1.customer.tax_id.deleted" event.
func (h *EventNotificationHandler) OnV1CustomerTaxIDDeleted(callback func(ctx context.Context, notif *V1CustomerTaxIDDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.tax_id.deleted", callback)
}

// OnV1CustomerTaxIDUpdated registers a callback to handle notifications about the "v1.customer.tax_id.updated" event.
func (h *EventNotificationHandler) OnV1CustomerTaxIDUpdated(callback func(ctx context.Context, notif *V1CustomerTaxIDUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.tax_id.updated", callback)
}

// OnV1CustomerUpdated registers a callback to handle notifications about the "v1.customer.updated" event.
func (h *EventNotificationHandler) OnV1CustomerUpdated(callback func(ctx context.Context, notif *V1CustomerUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.customer.updated", callback)
}

// OnV1CustomerCashBalanceTransactionCreated registers a callback to handle notifications about the "v1.customer_cash_balance_transaction.created" event.
func (h *EventNotificationHandler) OnV1CustomerCashBalanceTransactionCreated(callback func(ctx context.Context, notif *V1CustomerCashBalanceTransactionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.customer_cash_balance_transaction.created", callback)
}

// OnV1EntitlementsActiveEntitlementSummaryUpdated registers a callback to handle notifications about the "v1.entitlements.active_entitlement_summary.updated" event.
func (h *EventNotificationHandler) OnV1EntitlementsActiveEntitlementSummaryUpdated(callback func(ctx context.Context, notif *V1EntitlementsActiveEntitlementSummaryUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.entitlements.active_entitlement_summary.updated", callback)
}

// OnV1FileCreated registers a callback to handle notifications about the "v1.file.created" event.
func (h *EventNotificationHandler) OnV1FileCreated(callback func(ctx context.Context, notif *V1FileCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.file.created", callback)
}

// OnV1FinancialConnectionsAccountCreated registers a callback to handle notifications about the "v1.financial_connections.account.created" event.
func (h *EventNotificationHandler) OnV1FinancialConnectionsAccountCreated(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.created", callback)
}

// OnV1FinancialConnectionsAccountDeactivated registers a callback to handle notifications about the "v1.financial_connections.account.deactivated" event.
func (h *EventNotificationHandler) OnV1FinancialConnectionsAccountDeactivated(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountDeactivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.deactivated", callback)
}

// OnV1FinancialConnectionsAccountDisconnected registers a callback to handle notifications about the "v1.financial_connections.account.disconnected" event.
func (h *EventNotificationHandler) OnV1FinancialConnectionsAccountDisconnected(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountDisconnectedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.disconnected", callback)
}

// OnV1FinancialConnectionsAccountReactivated registers a callback to handle notifications about the "v1.financial_connections.account.reactivated" event.
func (h *EventNotificationHandler) OnV1FinancialConnectionsAccountReactivated(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountReactivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.reactivated", callback)
}

// OnV1FinancialConnectionsAccountRefreshedBalance registers a callback to handle notifications about the "v1.financial_connections.account.refreshed_balance" event.
func (h *EventNotificationHandler) OnV1FinancialConnectionsAccountRefreshedBalance(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountRefreshedBalanceEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.refreshed_balance", callback)
}

// OnV1FinancialConnectionsAccountRefreshedOwnership registers a callback to handle notifications about the "v1.financial_connections.account.refreshed_ownership" event.
func (h *EventNotificationHandler) OnV1FinancialConnectionsAccountRefreshedOwnership(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountRefreshedOwnershipEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.refreshed_ownership", callback)
}

// OnV1FinancialConnectionsAccountRefreshedTransactions registers a callback to handle notifications about the "v1.financial_connections.account.refreshed_transactions" event.
func (h *EventNotificationHandler) OnV1FinancialConnectionsAccountRefreshedTransactions(callback func(ctx context.Context, notif *V1FinancialConnectionsAccountRefreshedTransactionsEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.financial_connections.account.refreshed_transactions", callback)
}

// OnV1IdentityVerificationSessionCanceled registers a callback to handle notifications about the "v1.identity.verification_session.canceled" event.
func (h *EventNotificationHandler) OnV1IdentityVerificationSessionCanceled(callback func(ctx context.Context, notif *V1IdentityVerificationSessionCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.canceled", callback)
}

// OnV1IdentityVerificationSessionCreated registers a callback to handle notifications about the "v1.identity.verification_session.created" event.
func (h *EventNotificationHandler) OnV1IdentityVerificationSessionCreated(callback func(ctx context.Context, notif *V1IdentityVerificationSessionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.created", callback)
}

// OnV1IdentityVerificationSessionProcessing registers a callback to handle notifications about the "v1.identity.verification_session.processing" event.
func (h *EventNotificationHandler) OnV1IdentityVerificationSessionProcessing(callback func(ctx context.Context, notif *V1IdentityVerificationSessionProcessingEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.processing", callback)
}

// OnV1IdentityVerificationSessionRedacted registers a callback to handle notifications about the "v1.identity.verification_session.redacted" event.
func (h *EventNotificationHandler) OnV1IdentityVerificationSessionRedacted(callback func(ctx context.Context, notif *V1IdentityVerificationSessionRedactedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.redacted", callback)
}

// OnV1IdentityVerificationSessionRequiresInput registers a callback to handle notifications about the "v1.identity.verification_session.requires_input" event.
func (h *EventNotificationHandler) OnV1IdentityVerificationSessionRequiresInput(callback func(ctx context.Context, notif *V1IdentityVerificationSessionRequiresInputEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.requires_input", callback)
}

// OnV1IdentityVerificationSessionVerified registers a callback to handle notifications about the "v1.identity.verification_session.verified" event.
func (h *EventNotificationHandler) OnV1IdentityVerificationSessionVerified(callback func(ctx context.Context, notif *V1IdentityVerificationSessionVerifiedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.identity.verification_session.verified", callback)
}

// OnV1InvoiceCreated registers a callback to handle notifications about the "v1.invoice.created" event.
func (h *EventNotificationHandler) OnV1InvoiceCreated(callback func(ctx context.Context, notif *V1InvoiceCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.created", callback)
}

// OnV1InvoiceDeleted registers a callback to handle notifications about the "v1.invoice.deleted" event.
func (h *EventNotificationHandler) OnV1InvoiceDeleted(callback func(ctx context.Context, notif *V1InvoiceDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.deleted", callback)
}

// OnV1InvoiceFinalizationFailed registers a callback to handle notifications about the "v1.invoice.finalization_failed" event.
func (h *EventNotificationHandler) OnV1InvoiceFinalizationFailed(callback func(ctx context.Context, notif *V1InvoiceFinalizationFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.finalization_failed", callback)
}

// OnV1InvoiceFinalized registers a callback to handle notifications about the "v1.invoice.finalized" event.
func (h *EventNotificationHandler) OnV1InvoiceFinalized(callback func(ctx context.Context, notif *V1InvoiceFinalizedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.finalized", callback)
}

// OnV1InvoiceMarkedUncollectible registers a callback to handle notifications about the "v1.invoice.marked_uncollectible" event.
func (h *EventNotificationHandler) OnV1InvoiceMarkedUncollectible(callback func(ctx context.Context, notif *V1InvoiceMarkedUncollectibleEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.marked_uncollectible", callback)
}

// OnV1InvoiceOverdue registers a callback to handle notifications about the "v1.invoice.overdue" event.
func (h *EventNotificationHandler) OnV1InvoiceOverdue(callback func(ctx context.Context, notif *V1InvoiceOverdueEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.overdue", callback)
}

// OnV1InvoiceOverpaid registers a callback to handle notifications about the "v1.invoice.overpaid" event.
func (h *EventNotificationHandler) OnV1InvoiceOverpaid(callback func(ctx context.Context, notif *V1InvoiceOverpaidEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.overpaid", callback)
}

// OnV1InvoicePaid registers a callback to handle notifications about the "v1.invoice.paid" event.
func (h *EventNotificationHandler) OnV1InvoicePaid(callback func(ctx context.Context, notif *V1InvoicePaidEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.paid", callback)
}

// OnV1InvoicePaymentActionRequired registers a callback to handle notifications about the "v1.invoice.payment_action_required" event.
func (h *EventNotificationHandler) OnV1InvoicePaymentActionRequired(callback func(ctx context.Context, notif *V1InvoicePaymentActionRequiredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.payment_action_required", callback)
}

// OnV1InvoicePaymentFailed registers a callback to handle notifications about the "v1.invoice.payment_failed" event.
func (h *EventNotificationHandler) OnV1InvoicePaymentFailed(callback func(ctx context.Context, notif *V1InvoicePaymentFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.payment_failed", callback)
}

// OnV1InvoicePaymentSucceeded registers a callback to handle notifications about the "v1.invoice.payment_succeeded" event.
func (h *EventNotificationHandler) OnV1InvoicePaymentSucceeded(callback func(ctx context.Context, notif *V1InvoicePaymentSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.payment_succeeded", callback)
}

// OnV1InvoiceSent registers a callback to handle notifications about the "v1.invoice.sent" event.
func (h *EventNotificationHandler) OnV1InvoiceSent(callback func(ctx context.Context, notif *V1InvoiceSentEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.sent", callback)
}

// OnV1InvoiceUpcoming registers a callback to handle notifications about the "v1.invoice.upcoming" event.
func (h *EventNotificationHandler) OnV1InvoiceUpcoming(callback func(ctx context.Context, notif *V1InvoiceUpcomingEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.upcoming", callback)
}

// OnV1InvoiceUpdated registers a callback to handle notifications about the "v1.invoice.updated" event.
func (h *EventNotificationHandler) OnV1InvoiceUpdated(callback func(ctx context.Context, notif *V1InvoiceUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.updated", callback)
}

// OnV1InvoiceVoided registers a callback to handle notifications about the "v1.invoice.voided" event.
func (h *EventNotificationHandler) OnV1InvoiceVoided(callback func(ctx context.Context, notif *V1InvoiceVoidedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.voided", callback)
}

// OnV1InvoiceWillBeDue registers a callback to handle notifications about the "v1.invoice.will_be_due" event.
func (h *EventNotificationHandler) OnV1InvoiceWillBeDue(callback func(ctx context.Context, notif *V1InvoiceWillBeDueEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice.will_be_due", callback)
}

// OnV1InvoicePaymentPaid registers a callback to handle notifications about the "v1.invoice_payment.paid" event.
func (h *EventNotificationHandler) OnV1InvoicePaymentPaid(callback func(ctx context.Context, notif *V1InvoicePaymentPaidEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoice_payment.paid", callback)
}

// OnV1InvoiceitemCreated registers a callback to handle notifications about the "v1.invoiceitem.created" event.
func (h *EventNotificationHandler) OnV1InvoiceitemCreated(callback func(ctx context.Context, notif *V1InvoiceitemCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoiceitem.created", callback)
}

// OnV1InvoiceitemDeleted registers a callback to handle notifications about the "v1.invoiceitem.deleted" event.
func (h *EventNotificationHandler) OnV1InvoiceitemDeleted(callback func(ctx context.Context, notif *V1InvoiceitemDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.invoiceitem.deleted", callback)
}

// OnV1IssuingAuthorizationCreated registers a callback to handle notifications about the "v1.issuing_authorization.created" event.
func (h *EventNotificationHandler) OnV1IssuingAuthorizationCreated(callback func(ctx context.Context, notif *V1IssuingAuthorizationCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_authorization.created", callback)
}

// OnV1IssuingAuthorizationRequest registers a callback to handle notifications about the "v1.issuing_authorization.request" event.
func (h *EventNotificationHandler) OnV1IssuingAuthorizationRequest(callback func(ctx context.Context, notif *V1IssuingAuthorizationRequestEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_authorization.request", callback)
}

// OnV1IssuingAuthorizationUpdated registers a callback to handle notifications about the "v1.issuing_authorization.updated" event.
func (h *EventNotificationHandler) OnV1IssuingAuthorizationUpdated(callback func(ctx context.Context, notif *V1IssuingAuthorizationUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_authorization.updated", callback)
}

// OnV1IssuingCardCreated registers a callback to handle notifications about the "v1.issuing_card.created" event.
func (h *EventNotificationHandler) OnV1IssuingCardCreated(callback func(ctx context.Context, notif *V1IssuingCardCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_card.created", callback)
}

// OnV1IssuingCardUpdated registers a callback to handle notifications about the "v1.issuing_card.updated" event.
func (h *EventNotificationHandler) OnV1IssuingCardUpdated(callback func(ctx context.Context, notif *V1IssuingCardUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_card.updated", callback)
}

// OnV1IssuingCardholderCreated registers a callback to handle notifications about the "v1.issuing_cardholder.created" event.
func (h *EventNotificationHandler) OnV1IssuingCardholderCreated(callback func(ctx context.Context, notif *V1IssuingCardholderCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_cardholder.created", callback)
}

// OnV1IssuingCardholderUpdated registers a callback to handle notifications about the "v1.issuing_cardholder.updated" event.
func (h *EventNotificationHandler) OnV1IssuingCardholderUpdated(callback func(ctx context.Context, notif *V1IssuingCardholderUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_cardholder.updated", callback)
}

// OnV1IssuingDisputeClosed registers a callback to handle notifications about the "v1.issuing_dispute.closed" event.
func (h *EventNotificationHandler) OnV1IssuingDisputeClosed(callback func(ctx context.Context, notif *V1IssuingDisputeClosedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_dispute.closed", callback)
}

// OnV1IssuingDisputeCreated registers a callback to handle notifications about the "v1.issuing_dispute.created" event.
func (h *EventNotificationHandler) OnV1IssuingDisputeCreated(callback func(ctx context.Context, notif *V1IssuingDisputeCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_dispute.created", callback)
}

// OnV1IssuingDisputeFundsReinstated registers a callback to handle notifications about the "v1.issuing_dispute.funds_reinstated" event.
func (h *EventNotificationHandler) OnV1IssuingDisputeFundsReinstated(callback func(ctx context.Context, notif *V1IssuingDisputeFundsReinstatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_dispute.funds_reinstated", callback)
}

// OnV1IssuingDisputeFundsRescinded registers a callback to handle notifications about the "v1.issuing_dispute.funds_rescinded" event.
func (h *EventNotificationHandler) OnV1IssuingDisputeFundsRescinded(callback func(ctx context.Context, notif *V1IssuingDisputeFundsRescindedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_dispute.funds_rescinded", callback)
}

// OnV1IssuingDisputeSubmitted registers a callback to handle notifications about the "v1.issuing_dispute.submitted" event.
func (h *EventNotificationHandler) OnV1IssuingDisputeSubmitted(callback func(ctx context.Context, notif *V1IssuingDisputeSubmittedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_dispute.submitted", callback)
}

// OnV1IssuingDisputeUpdated registers a callback to handle notifications about the "v1.issuing_dispute.updated" event.
func (h *EventNotificationHandler) OnV1IssuingDisputeUpdated(callback func(ctx context.Context, notif *V1IssuingDisputeUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_dispute.updated", callback)
}

// OnV1IssuingPersonalizationDesignActivated registers a callback to handle notifications about the "v1.issuing_personalization_design.activated" event.
func (h *EventNotificationHandler) OnV1IssuingPersonalizationDesignActivated(callback func(ctx context.Context, notif *V1IssuingPersonalizationDesignActivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_personalization_design.activated", callback)
}

// OnV1IssuingPersonalizationDesignDeactivated registers a callback to handle notifications about the "v1.issuing_personalization_design.deactivated" event.
func (h *EventNotificationHandler) OnV1IssuingPersonalizationDesignDeactivated(callback func(ctx context.Context, notif *V1IssuingPersonalizationDesignDeactivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_personalization_design.deactivated", callback)
}

// OnV1IssuingPersonalizationDesignRejected registers a callback to handle notifications about the "v1.issuing_personalization_design.rejected" event.
func (h *EventNotificationHandler) OnV1IssuingPersonalizationDesignRejected(callback func(ctx context.Context, notif *V1IssuingPersonalizationDesignRejectedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_personalization_design.rejected", callback)
}

// OnV1IssuingPersonalizationDesignUpdated registers a callback to handle notifications about the "v1.issuing_personalization_design.updated" event.
func (h *EventNotificationHandler) OnV1IssuingPersonalizationDesignUpdated(callback func(ctx context.Context, notif *V1IssuingPersonalizationDesignUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_personalization_design.updated", callback)
}

// OnV1IssuingTokenCreated registers a callback to handle notifications about the "v1.issuing_token.created" event.
func (h *EventNotificationHandler) OnV1IssuingTokenCreated(callback func(ctx context.Context, notif *V1IssuingTokenCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_token.created", callback)
}

// OnV1IssuingTokenUpdated registers a callback to handle notifications about the "v1.issuing_token.updated" event.
func (h *EventNotificationHandler) OnV1IssuingTokenUpdated(callback func(ctx context.Context, notif *V1IssuingTokenUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_token.updated", callback)
}

// OnV1IssuingTransactionCreated registers a callback to handle notifications about the "v1.issuing_transaction.created" event.
func (h *EventNotificationHandler) OnV1IssuingTransactionCreated(callback func(ctx context.Context, notif *V1IssuingTransactionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_transaction.created", callback)
}

// OnV1IssuingTransactionPurchaseDetailsReceiptUpdated registers a callback to handle notifications about the "v1.issuing_transaction.purchase_details_receipt_updated" event.
func (h *EventNotificationHandler) OnV1IssuingTransactionPurchaseDetailsReceiptUpdated(callback func(ctx context.Context, notif *V1IssuingTransactionPurchaseDetailsReceiptUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.issuing_transaction.purchase_details_receipt_updated", callback)
}

// OnV1IssuingTransactionUpdated registers a callback to handle notifications about the "v1.issuing_transaction.updated" event.
func (h *EventNotificationHandler) OnV1IssuingTransactionUpdated(callback func(ctx context.Context, notif *V1IssuingTransactionUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.issuing_transaction.updated", callback)
}

// OnV1MandateUpdated registers a callback to handle notifications about the "v1.mandate.updated" event.
func (h *EventNotificationHandler) OnV1MandateUpdated(callback func(ctx context.Context, notif *V1MandateUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.mandate.updated", callback)
}

// OnV1PaymentIntentAmountCapturableUpdated registers a callback to handle notifications about the "v1.payment_intent.amount_capturable_updated" event.
func (h *EventNotificationHandler) OnV1PaymentIntentAmountCapturableUpdated(callback func(ctx context.Context, notif *V1PaymentIntentAmountCapturableUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.payment_intent.amount_capturable_updated", callback)
}

// OnV1PaymentIntentCanceled registers a callback to handle notifications about the "v1.payment_intent.canceled" event.
func (h *EventNotificationHandler) OnV1PaymentIntentCanceled(callback func(ctx context.Context, notif *V1PaymentIntentCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.canceled", callback)
}

// OnV1PaymentIntentCreated registers a callback to handle notifications about the "v1.payment_intent.created" event.
func (h *EventNotificationHandler) OnV1PaymentIntentCreated(callback func(ctx context.Context, notif *V1PaymentIntentCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.created", callback)
}

// OnV1PaymentIntentPartiallyFunded registers a callback to handle notifications about the "v1.payment_intent.partially_funded" event.
func (h *EventNotificationHandler) OnV1PaymentIntentPartiallyFunded(callback func(ctx context.Context, notif *V1PaymentIntentPartiallyFundedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.partially_funded", callback)
}

// OnV1PaymentIntentPaymentFailed registers a callback to handle notifications about the "v1.payment_intent.payment_failed" event.
func (h *EventNotificationHandler) OnV1PaymentIntentPaymentFailed(callback func(ctx context.Context, notif *V1PaymentIntentPaymentFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.payment_failed", callback)
}

// OnV1PaymentIntentProcessing registers a callback to handle notifications about the "v1.payment_intent.processing" event.
func (h *EventNotificationHandler) OnV1PaymentIntentProcessing(callback func(ctx context.Context, notif *V1PaymentIntentProcessingEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.processing", callback)
}

// OnV1PaymentIntentRequiresAction registers a callback to handle notifications about the "v1.payment_intent.requires_action" event.
func (h *EventNotificationHandler) OnV1PaymentIntentRequiresAction(callback func(ctx context.Context, notif *V1PaymentIntentRequiresActionEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.requires_action", callback)
}

// OnV1PaymentIntentSucceeded registers a callback to handle notifications about the "v1.payment_intent.succeeded" event.
func (h *EventNotificationHandler) OnV1PaymentIntentSucceeded(callback func(ctx context.Context, notif *V1PaymentIntentSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_intent.succeeded", callback)
}

// OnV1PaymentLinkCreated registers a callback to handle notifications about the "v1.payment_link.created" event.
func (h *EventNotificationHandler) OnV1PaymentLinkCreated(callback func(ctx context.Context, notif *V1PaymentLinkCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_link.created", callback)
}

// OnV1PaymentLinkUpdated registers a callback to handle notifications about the "v1.payment_link.updated" event.
func (h *EventNotificationHandler) OnV1PaymentLinkUpdated(callback func(ctx context.Context, notif *V1PaymentLinkUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_link.updated", callback)
}

// OnV1PaymentMethodAttached registers a callback to handle notifications about the "v1.payment_method.attached" event.
func (h *EventNotificationHandler) OnV1PaymentMethodAttached(callback func(ctx context.Context, notif *V1PaymentMethodAttachedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_method.attached", callback)
}

// OnV1PaymentMethodAutomaticallyUpdated registers a callback to handle notifications about the "v1.payment_method.automatically_updated" event.
func (h *EventNotificationHandler) OnV1PaymentMethodAutomaticallyUpdated(callback func(ctx context.Context, notif *V1PaymentMethodAutomaticallyUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.payment_method.automatically_updated", callback)
}

// OnV1PaymentMethodDetached registers a callback to handle notifications about the "v1.payment_method.detached" event.
func (h *EventNotificationHandler) OnV1PaymentMethodDetached(callback func(ctx context.Context, notif *V1PaymentMethodDetachedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_method.detached", callback)
}

// OnV1PaymentMethodUpdated registers a callback to handle notifications about the "v1.payment_method.updated" event.
func (h *EventNotificationHandler) OnV1PaymentMethodUpdated(callback func(ctx context.Context, notif *V1PaymentMethodUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payment_method.updated", callback)
}

// OnV1PayoutCanceled registers a callback to handle notifications about the "v1.payout.canceled" event.
func (h *EventNotificationHandler) OnV1PayoutCanceled(callback func(ctx context.Context, notif *V1PayoutCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.canceled", callback)
}

// OnV1PayoutCreated registers a callback to handle notifications about the "v1.payout.created" event.
func (h *EventNotificationHandler) OnV1PayoutCreated(callback func(ctx context.Context, notif *V1PayoutCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.created", callback)
}

// OnV1PayoutFailed registers a callback to handle notifications about the "v1.payout.failed" event.
func (h *EventNotificationHandler) OnV1PayoutFailed(callback func(ctx context.Context, notif *V1PayoutFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.failed", callback)
}

// OnV1PayoutPaid registers a callback to handle notifications about the "v1.payout.paid" event.
func (h *EventNotificationHandler) OnV1PayoutPaid(callback func(ctx context.Context, notif *V1PayoutPaidEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.paid", callback)
}

// OnV1PayoutReconciliationCompleted registers a callback to handle notifications about the "v1.payout.reconciliation_completed" event.
func (h *EventNotificationHandler) OnV1PayoutReconciliationCompleted(callback func(ctx context.Context, notif *V1PayoutReconciliationCompletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.reconciliation_completed", callback)
}

// OnV1PayoutUpdated registers a callback to handle notifications about the "v1.payout.updated" event.
func (h *EventNotificationHandler) OnV1PayoutUpdated(callback func(ctx context.Context, notif *V1PayoutUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.payout.updated", callback)
}

// OnV1PersonCreated registers a callback to handle notifications about the "v1.person.created" event.
func (h *EventNotificationHandler) OnV1PersonCreated(callback func(ctx context.Context, notif *V1PersonCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.person.created", callback)
}

// OnV1PersonDeleted registers a callback to handle notifications about the "v1.person.deleted" event.
func (h *EventNotificationHandler) OnV1PersonDeleted(callback func(ctx context.Context, notif *V1PersonDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.person.deleted", callback)
}

// OnV1PersonUpdated registers a callback to handle notifications about the "v1.person.updated" event.
func (h *EventNotificationHandler) OnV1PersonUpdated(callback func(ctx context.Context, notif *V1PersonUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.person.updated", callback)
}

// OnV1PlanCreated registers a callback to handle notifications about the "v1.plan.created" event.
func (h *EventNotificationHandler) OnV1PlanCreated(callback func(ctx context.Context, notif *V1PlanCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.plan.created", callback)
}

// OnV1PlanDeleted registers a callback to handle notifications about the "v1.plan.deleted" event.
func (h *EventNotificationHandler) OnV1PlanDeleted(callback func(ctx context.Context, notif *V1PlanDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.plan.deleted", callback)
}

// OnV1PlanUpdated registers a callback to handle notifications about the "v1.plan.updated" event.
func (h *EventNotificationHandler) OnV1PlanUpdated(callback func(ctx context.Context, notif *V1PlanUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.plan.updated", callback)
}

// OnV1PriceCreated registers a callback to handle notifications about the "v1.price.created" event.
func (h *EventNotificationHandler) OnV1PriceCreated(callback func(ctx context.Context, notif *V1PriceCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.price.created", callback)
}

// OnV1PriceDeleted registers a callback to handle notifications about the "v1.price.deleted" event.
func (h *EventNotificationHandler) OnV1PriceDeleted(callback func(ctx context.Context, notif *V1PriceDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.price.deleted", callback)
}

// OnV1PriceUpdated registers a callback to handle notifications about the "v1.price.updated" event.
func (h *EventNotificationHandler) OnV1PriceUpdated(callback func(ctx context.Context, notif *V1PriceUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.price.updated", callback)
}

// OnV1ProductCreated registers a callback to handle notifications about the "v1.product.created" event.
func (h *EventNotificationHandler) OnV1ProductCreated(callback func(ctx context.Context, notif *V1ProductCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.product.created", callback)
}

// OnV1ProductDeleted registers a callback to handle notifications about the "v1.product.deleted" event.
func (h *EventNotificationHandler) OnV1ProductDeleted(callback func(ctx context.Context, notif *V1ProductDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.product.deleted", callback)
}

// OnV1ProductUpdated registers a callback to handle notifications about the "v1.product.updated" event.
func (h *EventNotificationHandler) OnV1ProductUpdated(callback func(ctx context.Context, notif *V1ProductUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.product.updated", callback)
}

// OnV1PromotionCodeCreated registers a callback to handle notifications about the "v1.promotion_code.created" event.
func (h *EventNotificationHandler) OnV1PromotionCodeCreated(callback func(ctx context.Context, notif *V1PromotionCodeCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.promotion_code.created", callback)
}

// OnV1PromotionCodeUpdated registers a callback to handle notifications about the "v1.promotion_code.updated" event.
func (h *EventNotificationHandler) OnV1PromotionCodeUpdated(callback func(ctx context.Context, notif *V1PromotionCodeUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.promotion_code.updated", callback)
}

// OnV1QuoteAccepted registers a callback to handle notifications about the "v1.quote.accepted" event.
func (h *EventNotificationHandler) OnV1QuoteAccepted(callback func(ctx context.Context, notif *V1QuoteAcceptedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.quote.accepted", callback)
}

// OnV1QuoteCanceled registers a callback to handle notifications about the "v1.quote.canceled" event.
func (h *EventNotificationHandler) OnV1QuoteCanceled(callback func(ctx context.Context, notif *V1QuoteCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.quote.canceled", callback)
}

// OnV1QuoteCreated registers a callback to handle notifications about the "v1.quote.created" event.
func (h *EventNotificationHandler) OnV1QuoteCreated(callback func(ctx context.Context, notif *V1QuoteCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.quote.created", callback)
}

// OnV1QuoteFinalized registers a callback to handle notifications about the "v1.quote.finalized" event.
func (h *EventNotificationHandler) OnV1QuoteFinalized(callback func(ctx context.Context, notif *V1QuoteFinalizedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.quote.finalized", callback)
}

// OnV1RadarEarlyFraudWarningCreated registers a callback to handle notifications about the "v1.radar.early_fraud_warning.created" event.
func (h *EventNotificationHandler) OnV1RadarEarlyFraudWarningCreated(callback func(ctx context.Context, notif *V1RadarEarlyFraudWarningCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.radar.early_fraud_warning.created", callback)
}

// OnV1RadarEarlyFraudWarningUpdated registers a callback to handle notifications about the "v1.radar.early_fraud_warning.updated" event.
func (h *EventNotificationHandler) OnV1RadarEarlyFraudWarningUpdated(callback func(ctx context.Context, notif *V1RadarEarlyFraudWarningUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.radar.early_fraud_warning.updated", callback)
}

// OnV1RefundCreated registers a callback to handle notifications about the "v1.refund.created" event.
func (h *EventNotificationHandler) OnV1RefundCreated(callback func(ctx context.Context, notif *V1RefundCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.refund.created", callback)
}

// OnV1RefundFailed registers a callback to handle notifications about the "v1.refund.failed" event.
func (h *EventNotificationHandler) OnV1RefundFailed(callback func(ctx context.Context, notif *V1RefundFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.refund.failed", callback)
}

// OnV1RefundUpdated registers a callback to handle notifications about the "v1.refund.updated" event.
func (h *EventNotificationHandler) OnV1RefundUpdated(callback func(ctx context.Context, notif *V1RefundUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.refund.updated", callback)
}

// OnV1ReviewClosed registers a callback to handle notifications about the "v1.review.closed" event.
func (h *EventNotificationHandler) OnV1ReviewClosed(callback func(ctx context.Context, notif *V1ReviewClosedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.review.closed", callback)
}

// OnV1ReviewOpened registers a callback to handle notifications about the "v1.review.opened" event.
func (h *EventNotificationHandler) OnV1ReviewOpened(callback func(ctx context.Context, notif *V1ReviewOpenedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.review.opened", callback)
}

// OnV1SetupIntentCanceled registers a callback to handle notifications about the "v1.setup_intent.canceled" event.
func (h *EventNotificationHandler) OnV1SetupIntentCanceled(callback func(ctx context.Context, notif *V1SetupIntentCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.setup_intent.canceled", callback)
}

// OnV1SetupIntentCreated registers a callback to handle notifications about the "v1.setup_intent.created" event.
func (h *EventNotificationHandler) OnV1SetupIntentCreated(callback func(ctx context.Context, notif *V1SetupIntentCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.setup_intent.created", callback)
}

// OnV1SetupIntentRequiresAction registers a callback to handle notifications about the "v1.setup_intent.requires_action" event.
func (h *EventNotificationHandler) OnV1SetupIntentRequiresAction(callback func(ctx context.Context, notif *V1SetupIntentRequiresActionEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.setup_intent.requires_action", callback)
}

// OnV1SetupIntentSetupFailed registers a callback to handle notifications about the "v1.setup_intent.setup_failed" event.
func (h *EventNotificationHandler) OnV1SetupIntentSetupFailed(callback func(ctx context.Context, notif *V1SetupIntentSetupFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.setup_intent.setup_failed", callback)
}

// OnV1SetupIntentSucceeded registers a callback to handle notifications about the "v1.setup_intent.succeeded" event.
func (h *EventNotificationHandler) OnV1SetupIntentSucceeded(callback func(ctx context.Context, notif *V1SetupIntentSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.setup_intent.succeeded", callback)
}

// OnV1SigmaScheduledQueryRunCreated registers a callback to handle notifications about the "v1.sigma.scheduled_query_run.created" event.
func (h *EventNotificationHandler) OnV1SigmaScheduledQueryRunCreated(callback func(ctx context.Context, notif *V1SigmaScheduledQueryRunCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.sigma.scheduled_query_run.created", callback)
}

// OnV1SourceCanceled registers a callback to handle notifications about the "v1.source.canceled" event.
func (h *EventNotificationHandler) OnV1SourceCanceled(callback func(ctx context.Context, notif *V1SourceCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.source.canceled", callback)
}

// OnV1SourceChargeable registers a callback to handle notifications about the "v1.source.chargeable" event.
func (h *EventNotificationHandler) OnV1SourceChargeable(callback func(ctx context.Context, notif *V1SourceChargeableEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.source.chargeable", callback)
}

// OnV1SourceFailed registers a callback to handle notifications about the "v1.source.failed" event.
func (h *EventNotificationHandler) OnV1SourceFailed(callback func(ctx context.Context, notif *V1SourceFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.source.failed", callback)
}

// OnV1SourceRefundAttributesRequired registers a callback to handle notifications about the "v1.source.refund_attributes_required" event.
func (h *EventNotificationHandler) OnV1SourceRefundAttributesRequired(callback func(ctx context.Context, notif *V1SourceRefundAttributesRequiredEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.source.refund_attributes_required", callback)
}

// OnV1SubscriptionScheduleAborted registers a callback to handle notifications about the "v1.subscription_schedule.aborted" event.
func (h *EventNotificationHandler) OnV1SubscriptionScheduleAborted(callback func(ctx context.Context, notif *V1SubscriptionScheduleAbortedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.aborted", callback)
}

// OnV1SubscriptionScheduleCanceled registers a callback to handle notifications about the "v1.subscription_schedule.canceled" event.
func (h *EventNotificationHandler) OnV1SubscriptionScheduleCanceled(callback func(ctx context.Context, notif *V1SubscriptionScheduleCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.canceled", callback)
}

// OnV1SubscriptionScheduleCompleted registers a callback to handle notifications about the "v1.subscription_schedule.completed" event.
func (h *EventNotificationHandler) OnV1SubscriptionScheduleCompleted(callback func(ctx context.Context, notif *V1SubscriptionScheduleCompletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.completed", callback)
}

// OnV1SubscriptionScheduleCreated registers a callback to handle notifications about the "v1.subscription_schedule.created" event.
func (h *EventNotificationHandler) OnV1SubscriptionScheduleCreated(callback func(ctx context.Context, notif *V1SubscriptionScheduleCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.created", callback)
}

// OnV1SubscriptionScheduleExpiring registers a callback to handle notifications about the "v1.subscription_schedule.expiring" event.
func (h *EventNotificationHandler) OnV1SubscriptionScheduleExpiring(callback func(ctx context.Context, notif *V1SubscriptionScheduleExpiringEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.expiring", callback)
}

// OnV1SubscriptionScheduleReleased registers a callback to handle notifications about the "v1.subscription_schedule.released" event.
func (h *EventNotificationHandler) OnV1SubscriptionScheduleReleased(callback func(ctx context.Context, notif *V1SubscriptionScheduleReleasedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.released", callback)
}

// OnV1SubscriptionScheduleUpdated registers a callback to handle notifications about the "v1.subscription_schedule.updated" event.
func (h *EventNotificationHandler) OnV1SubscriptionScheduleUpdated(callback func(ctx context.Context, notif *V1SubscriptionScheduleUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.subscription_schedule.updated", callback)
}

// OnV1TaxSettingsUpdated registers a callback to handle notifications about the "v1.tax.settings.updated" event.
func (h *EventNotificationHandler) OnV1TaxSettingsUpdated(callback func(ctx context.Context, notif *V1TaxSettingsUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.tax.settings.updated", callback)
}

// OnV1TaxRateCreated registers a callback to handle notifications about the "v1.tax_rate.created" event.
func (h *EventNotificationHandler) OnV1TaxRateCreated(callback func(ctx context.Context, notif *V1TaxRateCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.tax_rate.created", callback)
}

// OnV1TaxRateUpdated registers a callback to handle notifications about the "v1.tax_rate.updated" event.
func (h *EventNotificationHandler) OnV1TaxRateUpdated(callback func(ctx context.Context, notif *V1TaxRateUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.tax_rate.updated", callback)
}

// OnV1TerminalReaderActionFailed registers a callback to handle notifications about the "v1.terminal.reader.action_failed" event.
func (h *EventNotificationHandler) OnV1TerminalReaderActionFailed(callback func(ctx context.Context, notif *V1TerminalReaderActionFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.terminal.reader.action_failed", callback)
}

// OnV1TerminalReaderActionSucceeded registers a callback to handle notifications about the "v1.terminal.reader.action_succeeded" event.
func (h *EventNotificationHandler) OnV1TerminalReaderActionSucceeded(callback func(ctx context.Context, notif *V1TerminalReaderActionSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.terminal.reader.action_succeeded", callback)
}

// OnV1TerminalReaderActionUpdated registers a callback to handle notifications about the "v1.terminal.reader.action_updated" event.
func (h *EventNotificationHandler) OnV1TerminalReaderActionUpdated(callback func(ctx context.Context, notif *V1TerminalReaderActionUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.terminal.reader.action_updated", callback)
}

// OnV1TestHelpersTestClockAdvancing registers a callback to handle notifications about the "v1.test_helpers.test_clock.advancing" event.
func (h *EventNotificationHandler) OnV1TestHelpersTestClockAdvancing(callback func(ctx context.Context, notif *V1TestHelpersTestClockAdvancingEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.test_helpers.test_clock.advancing", callback)
}

// OnV1TestHelpersTestClockCreated registers a callback to handle notifications about the "v1.test_helpers.test_clock.created" event.
func (h *EventNotificationHandler) OnV1TestHelpersTestClockCreated(callback func(ctx context.Context, notif *V1TestHelpersTestClockCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.test_helpers.test_clock.created", callback)
}

// OnV1TestHelpersTestClockDeleted registers a callback to handle notifications about the "v1.test_helpers.test_clock.deleted" event.
func (h *EventNotificationHandler) OnV1TestHelpersTestClockDeleted(callback func(ctx context.Context, notif *V1TestHelpersTestClockDeletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.test_helpers.test_clock.deleted", callback)
}

// OnV1TestHelpersTestClockInternalFailure registers a callback to handle notifications about the "v1.test_helpers.test_clock.internal_failure" event.
func (h *EventNotificationHandler) OnV1TestHelpersTestClockInternalFailure(callback func(ctx context.Context, notif *V1TestHelpersTestClockInternalFailureEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v1.test_helpers.test_clock.internal_failure", callback)
}

// OnV1TestHelpersTestClockReady registers a callback to handle notifications about the "v1.test_helpers.test_clock.ready" event.
func (h *EventNotificationHandler) OnV1TestHelpersTestClockReady(callback func(ctx context.Context, notif *V1TestHelpersTestClockReadyEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.test_helpers.test_clock.ready", callback)
}

// OnV1TopupCanceled registers a callback to handle notifications about the "v1.topup.canceled" event.
func (h *EventNotificationHandler) OnV1TopupCanceled(callback func(ctx context.Context, notif *V1TopupCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.topup.canceled", callback)
}

// OnV1TopupCreated registers a callback to handle notifications about the "v1.topup.created" event.
func (h *EventNotificationHandler) OnV1TopupCreated(callback func(ctx context.Context, notif *V1TopupCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.topup.created", callback)
}

// OnV1TopupFailed registers a callback to handle notifications about the "v1.topup.failed" event.
func (h *EventNotificationHandler) OnV1TopupFailed(callback func(ctx context.Context, notif *V1TopupFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.topup.failed", callback)
}

// OnV1TopupReversed registers a callback to handle notifications about the "v1.topup.reversed" event.
func (h *EventNotificationHandler) OnV1TopupReversed(callback func(ctx context.Context, notif *V1TopupReversedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.topup.reversed", callback)
}

// OnV1TopupSucceeded registers a callback to handle notifications about the "v1.topup.succeeded" event.
func (h *EventNotificationHandler) OnV1TopupSucceeded(callback func(ctx context.Context, notif *V1TopupSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.topup.succeeded", callback)
}

// OnV1TransferCreated registers a callback to handle notifications about the "v1.transfer.created" event.
func (h *EventNotificationHandler) OnV1TransferCreated(callback func(ctx context.Context, notif *V1TransferCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.transfer.created", callback)
}

// OnV1TransferReversed registers a callback to handle notifications about the "v1.transfer.reversed" event.
func (h *EventNotificationHandler) OnV1TransferReversed(callback func(ctx context.Context, notif *V1TransferReversedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.transfer.reversed", callback)
}

// OnV1TransferUpdated registers a callback to handle notifications about the "v1.transfer.updated" event.
func (h *EventNotificationHandler) OnV1TransferUpdated(callback func(ctx context.Context, notif *V1TransferUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v1.transfer.updated", callback)
}

// OnV2BillingCadenceBilled registers a callback to handle notifications about the "v2.billing.cadence.billed" event.
func (h *EventNotificationHandler) OnV2BillingCadenceBilled(callback func(ctx context.Context, notif *V2BillingCadenceBilledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.cadence.billed", callback)
}

// OnV2BillingCadenceCanceled registers a callback to handle notifications about the "v2.billing.cadence.canceled" event.
func (h *EventNotificationHandler) OnV2BillingCadenceCanceled(callback func(ctx context.Context, notif *V2BillingCadenceCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.cadence.canceled", callback)
}

// OnV2BillingCadenceCreated registers a callback to handle notifications about the "v2.billing.cadence.created" event.
func (h *EventNotificationHandler) OnV2BillingCadenceCreated(callback func(ctx context.Context, notif *V2BillingCadenceCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.cadence.created", callback)
}

// OnV2BillingLicenseFeeCreated registers a callback to handle notifications about the "v2.billing.license_fee.created" event.
func (h *EventNotificationHandler) OnV2BillingLicenseFeeCreated(callback func(ctx context.Context, notif *V2BillingLicenseFeeCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.license_fee.created", callback)
}

// OnV2BillingLicenseFeeUpdated registers a callback to handle notifications about the "v2.billing.license_fee.updated" event.
func (h *EventNotificationHandler) OnV2BillingLicenseFeeUpdated(callback func(ctx context.Context, notif *V2BillingLicenseFeeUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.license_fee.updated", callback)
}

// OnV2BillingLicenseFeeVersionCreated registers a callback to handle notifications about the "v2.billing.license_fee_version.created" event.
func (h *EventNotificationHandler) OnV2BillingLicenseFeeVersionCreated(callback func(ctx context.Context, notif *V2BillingLicenseFeeVersionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.license_fee_version.created", callback)
}

// OnV2BillingLicensedItemCreated registers a callback to handle notifications about the "v2.billing.licensed_item.created" event.
func (h *EventNotificationHandler) OnV2BillingLicensedItemCreated(callback func(ctx context.Context, notif *V2BillingLicensedItemCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.licensed_item.created", callback)
}

// OnV2BillingLicensedItemUpdated registers a callback to handle notifications about the "v2.billing.licensed_item.updated" event.
func (h *EventNotificationHandler) OnV2BillingLicensedItemUpdated(callback func(ctx context.Context, notif *V2BillingLicensedItemUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.licensed_item.updated", callback)
}

// OnV2BillingMeteredItemCreated registers a callback to handle notifications about the "v2.billing.metered_item.created" event.
func (h *EventNotificationHandler) OnV2BillingMeteredItemCreated(callback func(ctx context.Context, notif *V2BillingMeteredItemCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.metered_item.created", callback)
}

// OnV2BillingMeteredItemUpdated registers a callback to handle notifications about the "v2.billing.metered_item.updated" event.
func (h *EventNotificationHandler) OnV2BillingMeteredItemUpdated(callback func(ctx context.Context, notif *V2BillingMeteredItemUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.metered_item.updated", callback)
}

// OnV2BillingPricingPlanCreated registers a callback to handle notifications about the "v2.billing.pricing_plan.created" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanCreated(callback func(ctx context.Context, notif *V2BillingPricingPlanCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.pricing_plan.created", callback)
}

// OnV2BillingPricingPlanUpdated registers a callback to handle notifications about the "v2.billing.pricing_plan.updated" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanUpdated(callback func(ctx context.Context, notif *V2BillingPricingPlanUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.pricing_plan.updated", callback)
}

// OnV2BillingPricingPlanComponentCreated registers a callback to handle notifications about the "v2.billing.pricing_plan_component.created" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanComponentCreated(callback func(ctx context.Context, notif *V2BillingPricingPlanComponentCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_component.created", callback)
}

// OnV2BillingPricingPlanComponentUpdated registers a callback to handle notifications about the "v2.billing.pricing_plan_component.updated" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanComponentUpdated(callback func(ctx context.Context, notif *V2BillingPricingPlanComponentUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_component.updated", callback)
}

// OnV2BillingPricingPlanSubscriptionCollectionAwaitingCustomerAction registers a callback to handle notifications about the "v2.billing.pricing_plan_subscription.collection_awaiting_customer_action" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanSubscriptionCollectionAwaitingCustomerAction(callback func(ctx context.Context, notif *V2BillingPricingPlanSubscriptionCollectionAwaitingCustomerActionEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_subscription.collection_awaiting_customer_action", callback)
}

// OnV2BillingPricingPlanSubscriptionCollectionCurrent registers a callback to handle notifications about the "v2.billing.pricing_plan_subscription.collection_current" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanSubscriptionCollectionCurrent(callback func(ctx context.Context, notif *V2BillingPricingPlanSubscriptionCollectionCurrentEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_subscription.collection_current", callback)
}

// OnV2BillingPricingPlanSubscriptionCollectionPastDue registers a callback to handle notifications about the "v2.billing.pricing_plan_subscription.collection_past_due" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanSubscriptionCollectionPastDue(callback func(ctx context.Context, notif *V2BillingPricingPlanSubscriptionCollectionPastDueEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_subscription.collection_past_due", callback)
}

// OnV2BillingPricingPlanSubscriptionCollectionPaused registers a callback to handle notifications about the "v2.billing.pricing_plan_subscription.collection_paused" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanSubscriptionCollectionPaused(callback func(ctx context.Context, notif *V2BillingPricingPlanSubscriptionCollectionPausedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_subscription.collection_paused", callback)
}

// OnV2BillingPricingPlanSubscriptionCollectionUnpaid registers a callback to handle notifications about the "v2.billing.pricing_plan_subscription.collection_unpaid" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanSubscriptionCollectionUnpaid(callback func(ctx context.Context, notif *V2BillingPricingPlanSubscriptionCollectionUnpaidEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_subscription.collection_unpaid", callback)
}

// OnV2BillingPricingPlanSubscriptionServicingActivated registers a callback to handle notifications about the "v2.billing.pricing_plan_subscription.servicing_activated" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanSubscriptionServicingActivated(callback func(ctx context.Context, notif *V2BillingPricingPlanSubscriptionServicingActivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_subscription.servicing_activated", callback)
}

// OnV2BillingPricingPlanSubscriptionServicingCanceled registers a callback to handle notifications about the "v2.billing.pricing_plan_subscription.servicing_canceled" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanSubscriptionServicingCanceled(callback func(ctx context.Context, notif *V2BillingPricingPlanSubscriptionServicingCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_subscription.servicing_canceled", callback)
}

// OnV2BillingPricingPlanSubscriptionServicingPaused registers a callback to handle notifications about the "v2.billing.pricing_plan_subscription.servicing_paused" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanSubscriptionServicingPaused(callback func(ctx context.Context, notif *V2BillingPricingPlanSubscriptionServicingPausedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_subscription.servicing_paused", callback)
}

// OnV2BillingPricingPlanVersionCreated registers a callback to handle notifications about the "v2.billing.pricing_plan_version.created" event.
func (h *EventNotificationHandler) OnV2BillingPricingPlanVersionCreated(callback func(ctx context.Context, notif *V2BillingPricingPlanVersionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.pricing_plan_version.created", callback)
}

// OnV2BillingRateCardCreated registers a callback to handle notifications about the "v2.billing.rate_card.created" event.
func (h *EventNotificationHandler) OnV2BillingRateCardCreated(callback func(ctx context.Context, notif *V2BillingRateCardCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.rate_card.created", callback)
}

// OnV2BillingRateCardUpdated registers a callback to handle notifications about the "v2.billing.rate_card.updated" event.
func (h *EventNotificationHandler) OnV2BillingRateCardUpdated(callback func(ctx context.Context, notif *V2BillingRateCardUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.rate_card.updated", callback)
}

// OnV2BillingRateCardCustomPricingUnitOverageRateCreated registers a callback to handle notifications about the "v2.billing.rate_card_custom_pricing_unit_overage_rate.created" event.
func (h *EventNotificationHandler) OnV2BillingRateCardCustomPricingUnitOverageRateCreated(callback func(ctx context.Context, notif *V2BillingRateCardCustomPricingUnitOverageRateCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_custom_pricing_unit_overage_rate.created", callback)
}

// OnV2BillingRateCardRateCreated registers a callback to handle notifications about the "v2.billing.rate_card_rate.created" event.
func (h *EventNotificationHandler) OnV2BillingRateCardRateCreated(callback func(ctx context.Context, notif *V2BillingRateCardRateCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.billing.rate_card_rate.created", callback)
}

// OnV2BillingRateCardSubscriptionActivated registers a callback to handle notifications about the "v2.billing.rate_card_subscription.activated" event.
func (h *EventNotificationHandler) OnV2BillingRateCardSubscriptionActivated(callback func(ctx context.Context, notif *V2BillingRateCardSubscriptionActivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_subscription.activated", callback)
}

// OnV2BillingRateCardSubscriptionCanceled registers a callback to handle notifications about the "v2.billing.rate_card_subscription.canceled" event.
func (h *EventNotificationHandler) OnV2BillingRateCardSubscriptionCanceled(callback func(ctx context.Context, notif *V2BillingRateCardSubscriptionCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_subscription.canceled", callback)
}

// OnV2BillingRateCardSubscriptionCollectionAwaitingCustomerAction registers a callback to handle notifications about the "v2.billing.rate_card_subscription.collection_awaiting_customer_action" event.
func (h *EventNotificationHandler) OnV2BillingRateCardSubscriptionCollectionAwaitingCustomerAction(callback func(ctx context.Context, notif *V2BillingRateCardSubscriptionCollectionAwaitingCustomerActionEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_subscription.collection_awaiting_customer_action", callback)
}

// OnV2BillingRateCardSubscriptionCollectionCurrent registers a callback to handle notifications about the "v2.billing.rate_card_subscription.collection_current" event.
func (h *EventNotificationHandler) OnV2BillingRateCardSubscriptionCollectionCurrent(callback func(ctx context.Context, notif *V2BillingRateCardSubscriptionCollectionCurrentEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_subscription.collection_current", callback)
}

// OnV2BillingRateCardSubscriptionCollectionPastDue registers a callback to handle notifications about the "v2.billing.rate_card_subscription.collection_past_due" event.
func (h *EventNotificationHandler) OnV2BillingRateCardSubscriptionCollectionPastDue(callback func(ctx context.Context, notif *V2BillingRateCardSubscriptionCollectionPastDueEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_subscription.collection_past_due", callback)
}

// OnV2BillingRateCardSubscriptionCollectionPaused registers a callback to handle notifications about the "v2.billing.rate_card_subscription.collection_paused" event.
func (h *EventNotificationHandler) OnV2BillingRateCardSubscriptionCollectionPaused(callback func(ctx context.Context, notif *V2BillingRateCardSubscriptionCollectionPausedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_subscription.collection_paused", callback)
}

// OnV2BillingRateCardSubscriptionCollectionUnpaid registers a callback to handle notifications about the "v2.billing.rate_card_subscription.collection_unpaid" event.
func (h *EventNotificationHandler) OnV2BillingRateCardSubscriptionCollectionUnpaid(callback func(ctx context.Context, notif *V2BillingRateCardSubscriptionCollectionUnpaidEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_subscription.collection_unpaid", callback)
}

// OnV2BillingRateCardSubscriptionServicingActivated registers a callback to handle notifications about the "v2.billing.rate_card_subscription.servicing_activated" event.
func (h *EventNotificationHandler) OnV2BillingRateCardSubscriptionServicingActivated(callback func(ctx context.Context, notif *V2BillingRateCardSubscriptionServicingActivatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_subscription.servicing_activated", callback)
}

// OnV2BillingRateCardSubscriptionServicingCanceled registers a callback to handle notifications about the "v2.billing.rate_card_subscription.servicing_canceled" event.
func (h *EventNotificationHandler) OnV2BillingRateCardSubscriptionServicingCanceled(callback func(ctx context.Context, notif *V2BillingRateCardSubscriptionServicingCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_subscription.servicing_canceled", callback)
}

// OnV2BillingRateCardSubscriptionServicingPaused registers a callback to handle notifications about the "v2.billing.rate_card_subscription.servicing_paused" event.
func (h *EventNotificationHandler) OnV2BillingRateCardSubscriptionServicingPaused(callback func(ctx context.Context, notif *V2BillingRateCardSubscriptionServicingPausedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_subscription.servicing_paused", callback)
}

// OnV2BillingRateCardVersionCreated registers a callback to handle notifications about the "v2.billing.rate_card_version.created" event.
func (h *EventNotificationHandler) OnV2BillingRateCardVersionCreated(callback func(ctx context.Context, notif *V2BillingRateCardVersionCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.billing.rate_card_version.created", callback)
}

// OnV2CommerceProductCatalogImportsFailed registers a callback to handle notifications about the "v2.commerce.product_catalog.imports.failed" event.
func (h *EventNotificationHandler) OnV2CommerceProductCatalogImportsFailed(callback func(ctx context.Context, notif *V2CommerceProductCatalogImportsFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.commerce.product_catalog.imports.failed", callback)
}

// OnV2CommerceProductCatalogImportsProcessing registers a callback to handle notifications about the "v2.commerce.product_catalog.imports.processing" event.
func (h *EventNotificationHandler) OnV2CommerceProductCatalogImportsProcessing(callback func(ctx context.Context, notif *V2CommerceProductCatalogImportsProcessingEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.commerce.product_catalog.imports.processing", callback)
}

// OnV2CommerceProductCatalogImportsSucceeded registers a callback to handle notifications about the "v2.commerce.product_catalog.imports.succeeded" event.
func (h *EventNotificationHandler) OnV2CommerceProductCatalogImportsSucceeded(callback func(ctx context.Context, notif *V2CommerceProductCatalogImportsSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.commerce.product_catalog.imports.succeeded", callback)
}

// OnV2CommerceProductCatalogImportsSucceededWithErrors registers a callback to handle notifications about the "v2.commerce.product_catalog.imports.succeeded_with_errors" event.
func (h *EventNotificationHandler) OnV2CommerceProductCatalogImportsSucceededWithErrors(callback func(ctx context.Context, notif *V2CommerceProductCatalogImportsSucceededWithErrorsEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.commerce.product_catalog.imports.succeeded_with_errors", callback)
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

// OnV2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdated registers a callback to handle notifications about the "v2.core.account[configuration.card_creator].capability_status_updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.card_creator].capability_status_updated", callback)
}

// OnV2CoreAccountIncludingConfigurationCardCreatorUpdated registers a callback to handle notifications about the "v2.core.account[configuration.card_creator].updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingConfigurationCardCreatorUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingConfigurationCardCreatorUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[configuration.card_creator].updated", callback)
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

// OnV2CoreAccountIncludingFutureRequirementsUpdated registers a callback to handle notifications about the "v2.core.account[future_requirements].updated" event.
func (h *EventNotificationHandler) OnV2CoreAccountIncludingFutureRequirementsUpdated(callback func(ctx context.Context, notif *V2CoreAccountIncludingFutureRequirementsUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account[future_requirements].updated", callback)
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

// OnV2CoreAccountSignalsFraudulentWebsiteReady registers a callback to handle notifications about the "v2.core.account_signals.fraudulent_website_ready" event.
func (h *EventNotificationHandler) OnV2CoreAccountSignalsFraudulentWebsiteReady(callback func(ctx context.Context, notif *V2CoreAccountSignalsFraudulentWebsiteReadyEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.account_signals.fraudulent_website_ready", callback)
}

// OnV2CoreApprovalRequestApproved registers a callback to handle notifications about the "v2.core.approval_request.approved" event.
func (h *EventNotificationHandler) OnV2CoreApprovalRequestApproved(callback func(ctx context.Context, notif *V2CoreApprovalRequestApprovedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.approval_request.approved", callback)
}

// OnV2CoreApprovalRequestCanceled registers a callback to handle notifications about the "v2.core.approval_request.canceled" event.
func (h *EventNotificationHandler) OnV2CoreApprovalRequestCanceled(callback func(ctx context.Context, notif *V2CoreApprovalRequestCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.approval_request.canceled", callback)
}

// OnV2CoreApprovalRequestCreated registers a callback to handle notifications about the "v2.core.approval_request.created" event.
func (h *EventNotificationHandler) OnV2CoreApprovalRequestCreated(callback func(ctx context.Context, notif *V2CoreApprovalRequestCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.approval_request.created", callback)
}

// OnV2CoreApprovalRequestExpired registers a callback to handle notifications about the "v2.core.approval_request.expired" event.
func (h *EventNotificationHandler) OnV2CoreApprovalRequestExpired(callback func(ctx context.Context, notif *V2CoreApprovalRequestExpiredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.approval_request.expired", callback)
}

// OnV2CoreApprovalRequestFailed registers a callback to handle notifications about the "v2.core.approval_request.failed" event.
func (h *EventNotificationHandler) OnV2CoreApprovalRequestFailed(callback func(ctx context.Context, notif *V2CoreApprovalRequestFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.approval_request.failed", callback)
}

// OnV2CoreApprovalRequestRejected registers a callback to handle notifications about the "v2.core.approval_request.rejected" event.
func (h *EventNotificationHandler) OnV2CoreApprovalRequestRejected(callback func(ctx context.Context, notif *V2CoreApprovalRequestRejectedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.approval_request.rejected", callback)
}

// OnV2CoreApprovalRequestSucceeded registers a callback to handle notifications about the "v2.core.approval_request.succeeded" event.
func (h *EventNotificationHandler) OnV2CoreApprovalRequestSucceeded(callback func(ctx context.Context, notif *V2CoreApprovalRequestSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.approval_request.succeeded", callback)
}

// OnV2CoreBatchJobBatchFailed registers a callback to handle notifications about the "v2.core.batch_job.batch_failed" event.
func (h *EventNotificationHandler) OnV2CoreBatchJobBatchFailed(callback func(ctx context.Context, notif *V2CoreBatchJobBatchFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.batch_job.batch_failed", callback)
}

// OnV2CoreBatchJobCanceled registers a callback to handle notifications about the "v2.core.batch_job.canceled" event.
func (h *EventNotificationHandler) OnV2CoreBatchJobCanceled(callback func(ctx context.Context, notif *V2CoreBatchJobCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.batch_job.canceled", callback)
}

// OnV2CoreBatchJobCompleted registers a callback to handle notifications about the "v2.core.batch_job.completed" event.
func (h *EventNotificationHandler) OnV2CoreBatchJobCompleted(callback func(ctx context.Context, notif *V2CoreBatchJobCompletedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.batch_job.completed", callback)
}

// OnV2CoreBatchJobCreated registers a callback to handle notifications about the "v2.core.batch_job.created" event.
func (h *EventNotificationHandler) OnV2CoreBatchJobCreated(callback func(ctx context.Context, notif *V2CoreBatchJobCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.batch_job.created", callback)
}

// OnV2CoreBatchJobReadyForUpload registers a callback to handle notifications about the "v2.core.batch_job.ready_for_upload" event.
func (h *EventNotificationHandler) OnV2CoreBatchJobReadyForUpload(callback func(ctx context.Context, notif *V2CoreBatchJobReadyForUploadEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.batch_job.ready_for_upload", callback)
}

// OnV2CoreBatchJobTimeout registers a callback to handle notifications about the "v2.core.batch_job.timeout" event.
func (h *EventNotificationHandler) OnV2CoreBatchJobTimeout(callback func(ctx context.Context, notif *V2CoreBatchJobTimeoutEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.batch_job.timeout", callback)
}

// OnV2CoreBatchJobUpdated registers a callback to handle notifications about the "v2.core.batch_job.updated" event.
func (h *EventNotificationHandler) OnV2CoreBatchJobUpdated(callback func(ctx context.Context, notif *V2CoreBatchJobUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.batch_job.updated", callback)
}

// OnV2CoreBatchJobUploadTimeout registers a callback to handle notifications about the "v2.core.batch_job.upload_timeout" event.
func (h *EventNotificationHandler) OnV2CoreBatchJobUploadTimeout(callback func(ctx context.Context, notif *V2CoreBatchJobUploadTimeoutEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.batch_job.upload_timeout", callback)
}

// OnV2CoreBatchJobValidating registers a callback to handle notifications about the "v2.core.batch_job.validating" event.
func (h *EventNotificationHandler) OnV2CoreBatchJobValidating(callback func(ctx context.Context, notif *V2CoreBatchJobValidatingEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.batch_job.validating", callback)
}

// OnV2CoreBatchJobValidationFailed registers a callback to handle notifications about the "v2.core.batch_job.validation_failed" event.
func (h *EventNotificationHandler) OnV2CoreBatchJobValidationFailed(callback func(ctx context.Context, notif *V2CoreBatchJobValidationFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.batch_job.validation_failed", callback)
}

// OnV2CoreClaimableSandboxClaimed registers a callback to handle notifications about the "v2.core.claimable_sandbox.claimed" event.
func (h *EventNotificationHandler) OnV2CoreClaimableSandboxClaimed(callback func(ctx context.Context, notif *V2CoreClaimableSandboxClaimedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.claimable_sandbox.claimed", callback)
}

// OnV2CoreClaimableSandboxCreated registers a callback to handle notifications about the "v2.core.claimable_sandbox.created" event.
func (h *EventNotificationHandler) OnV2CoreClaimableSandboxCreated(callback func(ctx context.Context, notif *V2CoreClaimableSandboxCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.claimable_sandbox.created", callback)
}

// OnV2CoreClaimableSandboxExpired registers a callback to handle notifications about the "v2.core.claimable_sandbox.expired" event.
func (h *EventNotificationHandler) OnV2CoreClaimableSandboxExpired(callback func(ctx context.Context, notif *V2CoreClaimableSandboxExpiredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.claimable_sandbox.expired", callback)
}

// OnV2CoreClaimableSandboxExpiring registers a callback to handle notifications about the "v2.core.claimable_sandbox.expiring" event.
func (h *EventNotificationHandler) OnV2CoreClaimableSandboxExpiring(callback func(ctx context.Context, notif *V2CoreClaimableSandboxExpiringEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.claimable_sandbox.expiring", callback)
}

// OnV2CoreClaimableSandboxUpdated registers a callback to handle notifications about the "v2.core.claimable_sandbox.updated" event.
func (h *EventNotificationHandler) OnV2CoreClaimableSandboxUpdated(callback func(ctx context.Context, notif *V2CoreClaimableSandboxUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.claimable_sandbox.updated", callback)
}

// OnV2CoreEventDestinationPing registers a callback to handle notifications about the "v2.core.event_destination.ping" event.
func (h *EventNotificationHandler) OnV2CoreEventDestinationPing(callback func(ctx context.Context, notif *V2CoreEventDestinationPingEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.event_destination.ping", callback)
}

// OnV2CoreHealthAPIErrorFiring registers a callback to handle notifications about the "v2.core.health.api_error.firing" event.
func (h *EventNotificationHandler) OnV2CoreHealthAPIErrorFiring(callback func(ctx context.Context, notif *V2CoreHealthAPIErrorFiringEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.health.api_error.firing", callback)
}

// OnV2CoreHealthAPIErrorResolved registers a callback to handle notifications about the "v2.core.health.api_error.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthAPIErrorResolved(callback func(ctx context.Context, notif *V2CoreHealthAPIErrorResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.health.api_error.resolved", callback)
}

// OnV2CoreHealthAPILatencyFiring registers a callback to handle notifications about the "v2.core.health.api_latency.firing" event.
func (h *EventNotificationHandler) OnV2CoreHealthAPILatencyFiring(callback func(ctx context.Context, notif *V2CoreHealthAPILatencyFiringEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.core.health.api_latency.firing", callback)
}

// OnV2CoreHealthAPILatencyResolved registers a callback to handle notifications about the "v2.core.health.api_latency.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthAPILatencyResolved(callback func(ctx context.Context, notif *V2CoreHealthAPILatencyResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.api_latency.resolved", callback)
}

// OnV2CoreHealthAuthorizationRateDropFiring registers a callback to handle notifications about the "v2.core.health.authorization_rate_drop.firing" event.
func (h *EventNotificationHandler) OnV2CoreHealthAuthorizationRateDropFiring(callback func(ctx context.Context, notif *V2CoreHealthAuthorizationRateDropFiringEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.authorization_rate_drop.firing", callback)
}

// OnV2CoreHealthAuthorizationRateDropResolved registers a callback to handle notifications about the "v2.core.health.authorization_rate_drop.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthAuthorizationRateDropResolved(callback func(ctx context.Context, notif *V2CoreHealthAuthorizationRateDropResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.authorization_rate_drop.resolved", callback)
}

// OnV2CoreHealthEventGenerationFailureResolved registers a callback to handle notifications about the "v2.core.health.event_generation_failure.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthEventGenerationFailureResolved(callback func(ctx context.Context, notif *V2CoreHealthEventGenerationFailureResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.event_generation_failure.resolved", callback)
}

// OnV2CoreHealthFraudRateIncreased registers a callback to handle notifications about the "v2.core.health.fraud_rate.increased" event.
func (h *EventNotificationHandler) OnV2CoreHealthFraudRateIncreased(callback func(ctx context.Context, notif *V2CoreHealthFraudRateIncreasedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.fraud_rate.increased", callback)
}

// OnV2CoreHealthIssuingAuthorizationRequestErrorsFiring registers a callback to handle notifications about the "v2.core.health.issuing_authorization_request_errors.firing" event.
func (h *EventNotificationHandler) OnV2CoreHealthIssuingAuthorizationRequestErrorsFiring(callback func(ctx context.Context, notif *V2CoreHealthIssuingAuthorizationRequestErrorsFiringEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.issuing_authorization_request_errors.firing", callback)
}

// OnV2CoreHealthIssuingAuthorizationRequestErrorsResolved registers a callback to handle notifications about the "v2.core.health.issuing_authorization_request_errors.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthIssuingAuthorizationRequestErrorsResolved(callback func(ctx context.Context, notif *V2CoreHealthIssuingAuthorizationRequestErrorsResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.issuing_authorization_request_errors.resolved", callback)
}

// OnV2CoreHealthIssuingAuthorizationRequestTimeoutFiring registers a callback to handle notifications about the "v2.core.health.issuing_authorization_request_timeout.firing" event.
func (h *EventNotificationHandler) OnV2CoreHealthIssuingAuthorizationRequestTimeoutFiring(callback func(ctx context.Context, notif *V2CoreHealthIssuingAuthorizationRequestTimeoutFiringEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.issuing_authorization_request_timeout.firing", callback)
}

// OnV2CoreHealthIssuingAuthorizationRequestTimeoutResolved registers a callback to handle notifications about the "v2.core.health.issuing_authorization_request_timeout.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthIssuingAuthorizationRequestTimeoutResolved(callback func(ctx context.Context, notif *V2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.issuing_authorization_request_timeout.resolved", callback)
}

// OnV2CoreHealthMeterEventSummariesDelayedFiring registers a callback to handle notifications about the "v2.core.health.meter_event_summaries_delayed.firing" event.
func (h *EventNotificationHandler) OnV2CoreHealthMeterEventSummariesDelayedFiring(callback func(ctx context.Context, notif *V2CoreHealthMeterEventSummariesDelayedFiringEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.meter_event_summaries_delayed.firing", callback)
}

// OnV2CoreHealthMeterEventSummariesDelayedResolved registers a callback to handle notifications about the "v2.core.health.meter_event_summaries_delayed.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthMeterEventSummariesDelayedResolved(callback func(ctx context.Context, notif *V2CoreHealthMeterEventSummariesDelayedResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.meter_event_summaries_delayed.resolved", callback)
}

// OnV2CoreHealthPaymentMethodErrorFiring registers a callback to handle notifications about the "v2.core.health.payment_method_error.firing" event.
func (h *EventNotificationHandler) OnV2CoreHealthPaymentMethodErrorFiring(callback func(ctx context.Context, notif *V2CoreHealthPaymentMethodErrorFiringEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.payment_method_error.firing", callback)
}

// OnV2CoreHealthPaymentMethodErrorResolved registers a callback to handle notifications about the "v2.core.health.payment_method_error.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthPaymentMethodErrorResolved(callback func(ctx context.Context, notif *V2CoreHealthPaymentMethodErrorResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.payment_method_error.resolved", callback)
}

// OnV2CoreHealthSEPADebitDelayedFiring registers a callback to handle notifications about the "v2.core.health.sepa_debit_delayed.firing" event.
func (h *EventNotificationHandler) OnV2CoreHealthSEPADebitDelayedFiring(callback func(ctx context.Context, notif *V2CoreHealthSEPADebitDelayedFiringEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.sepa_debit_delayed.firing", callback)
}

// OnV2CoreHealthSEPADebitDelayedResolved registers a callback to handle notifications about the "v2.core.health.sepa_debit_delayed.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthSEPADebitDelayedResolved(callback func(ctx context.Context, notif *V2CoreHealthSEPADebitDelayedResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.sepa_debit_delayed.resolved", callback)
}

// OnV2CoreHealthTrafficVolumeDropFiring registers a callback to handle notifications about the "v2.core.health.traffic_volume_drop.firing" event.
func (h *EventNotificationHandler) OnV2CoreHealthTrafficVolumeDropFiring(callback func(ctx context.Context, notif *V2CoreHealthTrafficVolumeDropFiringEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.traffic_volume_drop.firing", callback)
}

// OnV2CoreHealthTrafficVolumeDropResolved registers a callback to handle notifications about the "v2.core.health.traffic_volume_drop.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthTrafficVolumeDropResolved(callback func(ctx context.Context, notif *V2CoreHealthTrafficVolumeDropResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.traffic_volume_drop.resolved", callback)
}

// OnV2CoreHealthWebhookLatencyFiring registers a callback to handle notifications about the "v2.core.health.webhook_latency.firing" event.
func (h *EventNotificationHandler) OnV2CoreHealthWebhookLatencyFiring(callback func(ctx context.Context, notif *V2CoreHealthWebhookLatencyFiringEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.webhook_latency.firing", callback)
}

// OnV2CoreHealthWebhookLatencyResolved registers a callback to handle notifications about the "v2.core.health.webhook_latency.resolved" event.
func (h *EventNotificationHandler) OnV2CoreHealthWebhookLatencyResolved(callback func(ctx context.Context, notif *V2CoreHealthWebhookLatencyResolvedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.core.health.webhook_latency.resolved", callback)
}

// OnV2DataReportingQueryRunCreated registers a callback to handle notifications about the "v2.data.reporting.query_run.created" event.
func (h *EventNotificationHandler) OnV2DataReportingQueryRunCreated(callback func(ctx context.Context, notif *V2DataReportingQueryRunCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.data.reporting.query_run.created", callback)
}

// OnV2DataReportingQueryRunFailed registers a callback to handle notifications about the "v2.data.reporting.query_run.failed" event.
func (h *EventNotificationHandler) OnV2DataReportingQueryRunFailed(callback func(ctx context.Context, notif *V2DataReportingQueryRunFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.data.reporting.query_run.failed", callback)
}

// OnV2DataReportingQueryRunSucceeded registers a callback to handle notifications about the "v2.data.reporting.query_run.succeeded" event.
func (h *EventNotificationHandler) OnV2DataReportingQueryRunSucceeded(callback func(ctx context.Context, notif *V2DataReportingQueryRunSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.data.reporting.query_run.succeeded", callback)
}

// OnV2DataReportingQueryRunUpdated registers a callback to handle notifications about the "v2.data.reporting.query_run.updated" event.
func (h *EventNotificationHandler) OnV2DataReportingQueryRunUpdated(callback func(ctx context.Context, notif *V2DataReportingQueryRunUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.data.reporting.query_run.updated", callback)
}

// OnV2ExtendExtensionRunFailed registers a callback to handle notifications about the "v2.extend.extension_run.failed" event.
func (h *EventNotificationHandler) OnV2ExtendExtensionRunFailed(callback func(ctx context.Context, notif *V2ExtendExtensionRunFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.extend.extension_run.failed", callback)
}

// OnV2ExtendWorkflowRunFailed registers a callback to handle notifications about the "v2.extend.workflow_run.failed" event.
func (h *EventNotificationHandler) OnV2ExtendWorkflowRunFailed(callback func(ctx context.Context, notif *V2ExtendWorkflowRunFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.extend.workflow_run.failed", callback)
}

// OnV2ExtendWorkflowRunStarted registers a callback to handle notifications about the "v2.extend.workflow_run.started" event.
func (h *EventNotificationHandler) OnV2ExtendWorkflowRunStarted(callback func(ctx context.Context, notif *V2ExtendWorkflowRunStartedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.extend.workflow_run.started", callback)
}

// OnV2ExtendWorkflowRunSucceeded registers a callback to handle notifications about the "v2.extend.workflow_run.succeeded" event.
func (h *EventNotificationHandler) OnV2ExtendWorkflowRunSucceeded(callback func(ctx context.Context, notif *V2ExtendWorkflowRunSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.extend.workflow_run.succeeded", callback)
}

// OnV2IamAPIKeyCreated registers a callback to handle notifications about the "v2.iam.api_key.created" event.
func (h *EventNotificationHandler) OnV2IamAPIKeyCreated(callback func(ctx context.Context, notif *V2IamAPIKeyCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.iam.api_key.created", callback)
}

// OnV2IamAPIKeyDefaultSecretRevealed registers a callback to handle notifications about the "v2.iam.api_key.default_secret_revealed" event.
func (h *EventNotificationHandler) OnV2IamAPIKeyDefaultSecretRevealed(callback func(ctx context.Context, notif *V2IamAPIKeyDefaultSecretRevealedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.iam.api_key.default_secret_revealed", callback)
}

// OnV2IamAPIKeyExpired registers a callback to handle notifications about the "v2.iam.api_key.expired" event.
func (h *EventNotificationHandler) OnV2IamAPIKeyExpired(callback func(ctx context.Context, notif *V2IamAPIKeyExpiredEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.iam.api_key.expired", callback)
}

// OnV2IamAPIKeyPermissionsUpdated registers a callback to handle notifications about the "v2.iam.api_key.permissions_updated" event.
func (h *EventNotificationHandler) OnV2IamAPIKeyPermissionsUpdated(callback func(ctx context.Context, notif *V2IamAPIKeyPermissionsUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.iam.api_key.permissions_updated", callback)
}

// OnV2IamAPIKeyRotated registers a callback to handle notifications about the "v2.iam.api_key.rotated" event.
func (h *EventNotificationHandler) OnV2IamAPIKeyRotated(callback func(ctx context.Context, notif *V2IamAPIKeyRotatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.iam.api_key.rotated", callback)
}

// OnV2IamAPIKeyUpdated registers a callback to handle notifications about the "v2.iam.api_key.updated" event.
func (h *EventNotificationHandler) OnV2IamAPIKeyUpdated(callback func(ctx context.Context, notif *V2IamAPIKeyUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.iam.api_key.updated", callback)
}

// OnV2IamStripeAccessGrantApproved registers a callback to handle notifications about the "v2.iam.stripe_access_grant.approved" event.
func (h *EventNotificationHandler) OnV2IamStripeAccessGrantApproved(callback func(ctx context.Context, notif *V2IamStripeAccessGrantApprovedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.iam.stripe_access_grant.approved", callback)
}

// OnV2IamStripeAccessGrantCanceled registers a callback to handle notifications about the "v2.iam.stripe_access_grant.canceled" event.
func (h *EventNotificationHandler) OnV2IamStripeAccessGrantCanceled(callback func(ctx context.Context, notif *V2IamStripeAccessGrantCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.iam.stripe_access_grant.canceled", callback)
}

// OnV2IamStripeAccessGrantDenied registers a callback to handle notifications about the "v2.iam.stripe_access_grant.denied" event.
func (h *EventNotificationHandler) OnV2IamStripeAccessGrantDenied(callback func(ctx context.Context, notif *V2IamStripeAccessGrantDeniedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.iam.stripe_access_grant.denied", callback)
}

// OnV2IamStripeAccessGrantRemoved registers a callback to handle notifications about the "v2.iam.stripe_access_grant.removed" event.
func (h *EventNotificationHandler) OnV2IamStripeAccessGrantRemoved(callback func(ctx context.Context, notif *V2IamStripeAccessGrantRemovedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.iam.stripe_access_grant.removed", callback)
}

// OnV2IamStripeAccessGrantRequested registers a callback to handle notifications about the "v2.iam.stripe_access_grant.requested" event.
func (h *EventNotificationHandler) OnV2IamStripeAccessGrantRequested(callback func(ctx context.Context, notif *V2IamStripeAccessGrantRequestedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.iam.stripe_access_grant.requested", callback)
}

// OnV2IamStripeAccessGrantUpdated registers a callback to handle notifications about the "v2.iam.stripe_access_grant.updated" event.
func (h *EventNotificationHandler) OnV2IamStripeAccessGrantUpdated(callback func(ctx context.Context, notif *V2IamStripeAccessGrantUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.iam.stripe_access_grant.updated", callback)
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

// OnV2MoneyManagementRecipientVerificationCreated registers a callback to handle notifications about the "v2.money_management.recipient_verification.created" event.
func (h *EventNotificationHandler) OnV2MoneyManagementRecipientVerificationCreated(callback func(ctx context.Context, notif *V2MoneyManagementRecipientVerificationCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.recipient_verification.created", callback)
}

// OnV2MoneyManagementRecipientVerificationUpdated registers a callback to handle notifications about the "v2.money_management.recipient_verification.updated" event.
func (h *EventNotificationHandler) OnV2MoneyManagementRecipientVerificationUpdated(callback func(ctx context.Context, notif *V2MoneyManagementRecipientVerificationUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.money_management.recipient_verification.updated", callback)
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

// OnV2OrchestratedCommerceAgreementConfirmed registers a callback to handle notifications about the "v2.orchestrated_commerce.agreement.confirmed" event.
func (h *EventNotificationHandler) OnV2OrchestratedCommerceAgreementConfirmed(callback func(ctx context.Context, notif *V2OrchestratedCommerceAgreementConfirmedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.orchestrated_commerce.agreement.confirmed", callback)
}

// OnV2OrchestratedCommerceAgreementCreated registers a callback to handle notifications about the "v2.orchestrated_commerce.agreement.created" event.
func (h *EventNotificationHandler) OnV2OrchestratedCommerceAgreementCreated(callback func(ctx context.Context, notif *V2OrchestratedCommerceAgreementCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.orchestrated_commerce.agreement.created", callback)
}

// OnV2OrchestratedCommerceAgreementPartiallyConfirmed registers a callback to handle notifications about the "v2.orchestrated_commerce.agreement.partially_confirmed" event.
func (h *EventNotificationHandler) OnV2OrchestratedCommerceAgreementPartiallyConfirmed(callback func(ctx context.Context, notif *V2OrchestratedCommerceAgreementPartiallyConfirmedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.orchestrated_commerce.agreement.partially_confirmed", callback)
}

// OnV2OrchestratedCommerceAgreementTerminated registers a callback to handle notifications about the "v2.orchestrated_commerce.agreement.terminated" event.
func (h *EventNotificationHandler) OnV2OrchestratedCommerceAgreementTerminated(callback func(ctx context.Context, notif *V2OrchestratedCommerceAgreementTerminatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.orchestrated_commerce.agreement.terminated", callback)
}

// OnV2PaymentsOffSessionPaymentAttemptFailed registers a callback to handle notifications about the "v2.payments.off_session_payment.attempt_failed" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentAttemptFailed(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentAttemptFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.attempt_failed", callback)
}

// OnV2PaymentsOffSessionPaymentAttemptStarted registers a callback to handle notifications about the "v2.payments.off_session_payment.attempt_started" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentAttemptStarted(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentAttemptStartedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.attempt_started", callback)
}

// OnV2PaymentsOffSessionPaymentAuthorizationAttemptFailed registers a callback to handle notifications about the "v2.payments.off_session_payment.authorization_attempt_failed" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentAuthorizationAttemptFailed(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.authorization_attempt_failed", callback)
}

// OnV2PaymentsOffSessionPaymentAuthorizationAttemptStarted registers a callback to handle notifications about the "v2.payments.off_session_payment.authorization_attempt_started" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentAuthorizationAttemptStarted(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.authorization_attempt_started", callback)
}

// OnV2PaymentsOffSessionPaymentCanceled registers a callback to handle notifications about the "v2.payments.off_session_payment.canceled" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentCanceled(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.canceled", callback)
}

// OnV2PaymentsOffSessionPaymentCreated registers a callback to handle notifications about the "v2.payments.off_session_payment.created" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentCreated(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.created", callback)
}

// OnV2PaymentsOffSessionPaymentFailed registers a callback to handle notifications about the "v2.payments.off_session_payment.failed" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentFailed(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.failed", callback)
}

// OnV2PaymentsOffSessionPaymentPaused registers a callback to handle notifications about the "v2.payments.off_session_payment.paused" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentPaused(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentPausedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.paused", callback)
}

// OnV2PaymentsOffSessionPaymentRequiresCapture registers a callback to handle notifications about the "v2.payments.off_session_payment.requires_capture" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentRequiresCapture(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentRequiresCaptureEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.requires_capture", callback)
}

// OnV2PaymentsOffSessionPaymentResumed registers a callback to handle notifications about the "v2.payments.off_session_payment.resumed" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentResumed(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentResumedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.resumed", callback)
}

// OnV2PaymentsOffSessionPaymentSucceeded registers a callback to handle notifications about the "v2.payments.off_session_payment.succeeded" event.
func (h *EventNotificationHandler) OnV2PaymentsOffSessionPaymentSucceeded(callback func(ctx context.Context, notif *V2PaymentsOffSessionPaymentSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.off_session_payment.succeeded", callback)
}

// OnV2PaymentsSettlementAllocationIntentCanceled registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent.canceled" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentCanceled(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent.canceled", callback)
}

// OnV2PaymentsSettlementAllocationIntentCreated registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent.created" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentCreated(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent.created", callback)
}

// OnV2PaymentsSettlementAllocationIntentErrored registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent.errored" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentErrored(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentErroredEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent.errored", callback)
}

// OnV2PaymentsSettlementAllocationIntentFundsNotReceived registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent.funds_not_received" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentFundsNotReceived(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentFundsNotReceivedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent.funds_not_received", callback)
}

// OnV2PaymentsSettlementAllocationIntentMatched registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent.matched" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentMatched(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentMatchedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent.matched", callback)
}

// OnV2PaymentsSettlementAllocationIntentNotFound registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent.not_found" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentNotFound(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentNotFoundEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent.not_found", callback)
}

// OnV2PaymentsSettlementAllocationIntentSettled registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent.settled" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentSettled(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentSettledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent.settled", callback)
}

// OnV2PaymentsSettlementAllocationIntentSubmitted registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent.submitted" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentSubmitted(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentSubmittedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent.submitted", callback)
}

// OnV2PaymentsSettlementAllocationIntentSplitCanceled registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent_split.canceled" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentSplitCanceled(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentSplitCanceledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent_split.canceled", callback)
}

// OnV2PaymentsSettlementAllocationIntentSplitCreated registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent_split.created" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentSplitCreated(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentSplitCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent_split.created", callback)
}

// OnV2PaymentsSettlementAllocationIntentSplitSettled registers a callback to handle notifications about the "v2.payments.settlement_allocation_intent_split.settled" event.
func (h *EventNotificationHandler) OnV2PaymentsSettlementAllocationIntentSplitSettled(callback func(ctx context.Context, notif *V2PaymentsSettlementAllocationIntentSplitSettledEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.payments.settlement_allocation_intent_split.settled", callback)
}

// OnV2ReportingReportRunCreated registers a callback to handle notifications about the "v2.reporting.report_run.created" event.
func (h *EventNotificationHandler) OnV2ReportingReportRunCreated(callback func(ctx context.Context, notif *V2ReportingReportRunCreatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.reporting.report_run.created", callback)
}

// OnV2ReportingReportRunFailed registers a callback to handle notifications about the "v2.reporting.report_run.failed" event.
func (h *EventNotificationHandler) OnV2ReportingReportRunFailed(callback func(ctx context.Context, notif *V2ReportingReportRunFailedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.reporting.report_run.failed", callback)
}

// OnV2ReportingReportRunSucceeded registers a callback to handle notifications about the "v2.reporting.report_run.succeeded" event.
func (h *EventNotificationHandler) OnV2ReportingReportRunSucceeded(callback func(ctx context.Context, notif *V2ReportingReportRunSucceededEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.reporting.report_run.succeeded", callback)
}

// OnV2ReportingReportRunUpdated registers a callback to handle notifications about the "v2.reporting.report_run.updated" event.
func (h *EventNotificationHandler) OnV2ReportingReportRunUpdated(callback func(ctx context.Context, notif *V2ReportingReportRunUpdatedEventNotification, client *Client) error) error {
	return registerTypedHandler(h, "v2.reporting.report_run.updated", callback)
}

// OnV2SignalsAccountSignalFraudulentMerchantReady registers a callback to handle notifications about the "v2.signals.account_signal.fraudulent_merchant_ready" event.
func (h *EventNotificationHandler) OnV2SignalsAccountSignalFraudulentMerchantReady(callback func(ctx context.Context, notif *V2SignalsAccountSignalFraudulentMerchantReadyEventNotification, client *Client) error) error {
	return registerTypedHandler(
		h, "v2.signals.account_signal.fraudulent_merchant_ready", callback)
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
