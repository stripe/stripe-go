package stripe

// most of this file is codegen, but not quite all of it.
// V2Events: The beginning of the section generated from our OpenAPI spec
import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// Open Enum.
type V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode string

// List of values that V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode can take
const (
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeArchivedMeter                   V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "archived_meter"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeMeterEventCustomerNotFound      V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "meter_event_customer_not_found"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeMeterEventDimensionCountTooHigh V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "meter_event_dimension_count_too_high"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeMeterEventInvalidValue          V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "meter_event_invalid_value"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeMeterEventNoCustomerDefined     V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "meter_event_no_customer_defined"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeMissingDimensionPayloadKeys     V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "missing_dimension_payload_keys"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeNoMeter                         V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "no_meter"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeTimestampInFuture               V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "timestamp_in_future"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeTimestampTooFarInPast           V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "timestamp_too_far_in_past"
)

// Open Enum.
type V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode string

// List of values that V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode can take
const (
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeArchivedMeter                   V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "archived_meter"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeMeterEventCustomerNotFound      V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "meter_event_customer_not_found"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeMeterEventDimensionCountTooHigh V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "meter_event_dimension_count_too_high"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeMeterEventInvalidValue          V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "meter_event_invalid_value"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeMeterEventNoCustomerDefined     V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "meter_event_no_customer_defined"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeMissingDimensionPayloadKeys     V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "missing_dimension_payload_keys"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeNoMeter                         V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "no_meter"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeTimestampInFuture               V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "timestamp_in_future"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeTimestampTooFarInPast           V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "timestamp_too_far_in_past"
)

// Open Enum. The capability which had its status updated.
type V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapability string

// List of values that V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapability can take
const (
	V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapabilityAutomaticIndirectTax V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapability = "automatic_indirect_tax"
)

// Open Enum. The capability which had its status updated.
type V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability string

// List of values that V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability can take
const (
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityACHDebitPayments         V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "ach_debit_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityACSSDebitPayments        V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "acss_debit_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityAffirmPayments           V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "affirm_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityAfterpayClearpayPayments V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "afterpay_clearpay_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityAlmaPayments             V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "alma_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityAmazonPayPayments        V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "amazon_pay_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityAUBECSDebitPayments      V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "au_becs_debit_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityBACSDebitPayments        V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "bacs_debit_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityBancontactPayments       V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "bancontact_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityBLIKPayments             V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "blik_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityBoletoPayments           V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "boleto_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityCardPayments             V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "card_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityCartesBancairesPayments  V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "cartes_bancaires_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityCashAppPayments          V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "cashapp_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityEPSPayments              V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "eps_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityFPXPayments              V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "fpx_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityGBBankTransferPayments   V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "gb_bank_transfer_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityGrabpayPayments          V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "grabpay_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityIDEALPayments            V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "ideal_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityJCBPayments              V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "jcb_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityJPBankTransferPayments   V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "jp_bank_transfer_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityKakaoPayPayments         V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "kakao_pay_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityKlarnaPayments           V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "klarna_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityKonbiniPayments          V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "konbini_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityKrCardPayments           V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "kr_card_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityLinkPayments             V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "link_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityMobilepayPayments        V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "mobilepay_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityMultibancoPayments       V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "multibanco_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityMXBankTransferPayments   V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "mx_bank_transfer_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityNaverPayPayments         V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "naver_pay_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityOXXOPayments             V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "oxxo_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityP24Payments              V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "p24_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityPaycoPayments            V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "payco_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityPayNowPayments           V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "paynow_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityStripeBalancePayouts     V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "stripe_balance.payouts"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityPayByBankPayments        V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "pay_by_bank_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityPromptPayPayments        V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "promptpay_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityRevolutPayPayments       V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "revolut_pay_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilitySamsungPayPayments       V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "samsung_pay_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilitySEPABankTransferPayments V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "sepa_bank_transfer_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilitySEPADebitPayments        V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "sepa_debit_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilitySwishPayments            V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "swish_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityTWINTPayments            V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "twint_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityUSBankTransferPayments   V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "us_bank_transfer_payments"
	V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityZipPayments              V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "zip_payments"
)

// Open Enum. The capability which had its status updated.
type V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability string

// List of values that V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability can take
const (
	V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityBankAccountsLocal            V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "bank_accounts.local"
	V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityBankAccountsWire             V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "bank_accounts.wire"
	V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityCards                        V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "cards"
	V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityStripeBalancePayouts         V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "stripe_balance.payouts"
	V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityStripeBalanceStripeTransfers V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "stripe_balance.stripe_transfers"
	V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityStripeTransfers              V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "stripe.transfers"
)

// Open Enum. The use case type of the account link that has been completed.
type V2CoreAccountLinkReturnedEventDataUseCase string

// List of values that V2CoreAccountLinkReturnedEventDataUseCase can take
const (
	V2CoreAccountLinkReturnedEventDataUseCaseAccountOnboarding V2CoreAccountLinkReturnedEventDataUseCase = "account_onboarding"
	V2CoreAccountLinkReturnedEventDataUseCaseAccountUpdate     V2CoreAccountLinkReturnedEventDataUseCase = "account_update"
)

// Configurations on the Account that was onboarded via the account link.
type V2CoreAccountLinkReturnedEventDataConfiguration string

// List of values that V2CoreAccountLinkReturnedEventDataConfiguration can take
const (
	V2CoreAccountLinkReturnedEventDataConfigurationCustomer  V2CoreAccountLinkReturnedEventDataConfiguration = "customer"
	V2CoreAccountLinkReturnedEventDataConfigurationMerchant  V2CoreAccountLinkReturnedEventDataConfiguration = "merchant"
	V2CoreAccountLinkReturnedEventDataConfigurationRecipient V2CoreAccountLinkReturnedEventDataConfiguration = "recipient"
)

// V2CoreEvent is the interface implemented by V2 Events. To get the underlying Event,
// use a type switch or type assertion to one of the concrete event types.
type V2CoreEvent interface {
	getBaseEvent() *V2BaseEvent
}

// V2CoreRawEvent is the raw event type for V2 events. It is used to unmarshal the
// event data into a generic structure, and can also be used a default event
// type when the event type is not known.
type V2CoreRawEvent struct {
	V2BaseEvent
	Data          *json.RawMessage          `json:"data"`
	RelatedObject *V2CoreEventRelatedObject `json:"related_object"`
}

// Used for everything internal to the EventNotifications
type eventNotificationParams struct {
	Params `form:"*"`
}

// V1BillingMeterErrorReportTriggeredEvent is the Go struct for the "v1.billing.meter.error_report_triggered" event.
// Occurs when a Meter has invalid async usage events.
type V1BillingMeterErrorReportTriggeredEvent struct {
	V2BaseEvent
	Data               V1BillingMeterErrorReportTriggeredEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject                    `json:"related_object"`
	fetchRelatedObject func() (*BillingMeter, error)
}

// FetchRelatedObject fetches the BillingMeter related to the event.
func (e *V1BillingMeterErrorReportTriggeredEvent) FetchRelatedObject(ctx context.Context) (*BillingMeter, error) {
	return e.fetchRelatedObject()
}

// V1BillingMeterErrorReportTriggeredEventNotification is the webhook payload you'll get when handling an event with type "v1.billing.meter.error_report_triggered"
// Occurs when a Meter has invalid async usage events.
type V1BillingMeterErrorReportTriggeredEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V1BillingMeterErrorReportTriggeredEvent that created this Notification
func (n *V1BillingMeterErrorReportTriggeredEventNotification) FetchEvent(ctx context.Context) (*V1BillingMeterErrorReportTriggeredEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V1BillingMeterErrorReportTriggeredEvent), nil
}

// FetchRelatedObject fetches the BillingMeter related to the event.
func (n *V1BillingMeterErrorReportTriggeredEventNotification) FetchRelatedObject(ctx context.Context) (*BillingMeter, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &BillingMeter{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V1BillingMeterNoMeterFoundEvent is the Go struct for the "v1.billing.meter.no_meter_found" event.
// Occurs when a Meter's id is missing or invalid in async usage events.
type V1BillingMeterNoMeterFoundEvent struct {
	V2BaseEvent
	Data V1BillingMeterNoMeterFoundEventData `json:"data"`
}

// V1BillingMeterNoMeterFoundEventNotification is the webhook payload you'll get when handling an event with type "v1.billing.meter.no_meter_found"
// Occurs when a Meter's id is missing or invalid in async usage events.
type V1BillingMeterNoMeterFoundEventNotification struct {
	V2CoreEventNotification
}

// FetchEvent retrieves the V1BillingMeterNoMeterFoundEvent that created this Notification
func (n *V1BillingMeterNoMeterFoundEventNotification) FetchEvent(ctx context.Context) (*V1BillingMeterNoMeterFoundEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V1BillingMeterNoMeterFoundEvent), nil
}

// V2CoreAccountClosedEvent is the Go struct for the "v2.core.account.closed" event.
// This event occurs when an account is closed.
type V2CoreAccountClosedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountClosedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountClosedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account.closed"
// This event occurs when an account is closed.
type V2CoreAccountClosedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountClosedEvent that created this Notification
func (n *V2CoreAccountClosedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountClosedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountClosedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountClosedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountCreatedEvent is the Go struct for the "v2.core.account.created" event.
// Occurs when an Account is created.
type V2CoreAccountCreatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountCreatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountCreatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account.created"
// Occurs when an Account is created.
type V2CoreAccountCreatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountCreatedEvent that created this Notification
func (n *V2CoreAccountCreatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountCreatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountCreatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountCreatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountUpdatedEvent is the Go struct for the "v2.core.account.updated" event.
// Occurs when an Account is updated.
type V2CoreAccountUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account.updated"
// Occurs when an Account is updated.
type V2CoreAccountUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountUpdatedEvent that created this Notification
func (n *V2CoreAccountUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.customer].capability_status_updated" event.
// Occurs when the status of an Account's customer configuration capability is updated.
type V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject                                                    `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[configuration.customer].capability_status_updated"
// Occurs when the status of an Account's customer configuration capability is updated.
type V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingConfigurationCustomerUpdatedEvent is the Go struct for the "v2.core.account[configuration.customer].updated" event.
// Occurs when an Account's customer configuration is updated.
type V2CoreAccountIncludingConfigurationCustomerUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingConfigurationCustomerUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[configuration.customer].updated"
// Occurs when an Account's customer configuration is updated.
type V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingConfigurationCustomerUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingConfigurationCustomerUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingConfigurationCustomerUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.merchant].capability_status_updated" event.
// Occurs when the status of an Account's merchant configuration capability is updated.
type V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject                                                    `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[configuration.merchant].capability_status_updated"
// Occurs when the status of an Account's merchant configuration capability is updated.
type V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingConfigurationMerchantUpdatedEvent is the Go struct for the "v2.core.account[configuration.merchant].updated" event.
// Occurs when an Account's merchant configuration is updated.
type V2CoreAccountIncludingConfigurationMerchantUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingConfigurationMerchantUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[configuration.merchant].updated"
// Occurs when an Account's merchant configuration is updated.
type V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingConfigurationMerchantUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingConfigurationMerchantUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingConfigurationMerchantUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.recipient].capability_status_updated" event.
// Occurs when the status of an Account's recipient configuration capability is updated.
type V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject                                                     `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[configuration.recipient].capability_status_updated"
// Occurs when the status of an Account's recipient configuration capability is updated.
type V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingConfigurationRecipientUpdatedEvent is the Go struct for the "v2.core.account[configuration.recipient].updated" event.
// Occurs when a Recipient's configuration is updated.
type V2CoreAccountIncludingConfigurationRecipientUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingConfigurationRecipientUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[configuration.recipient].updated"
// Occurs when a Recipient's configuration is updated.
type V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingConfigurationRecipientUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingConfigurationRecipientUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingConfigurationRecipientUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingDefaultsUpdatedEvent is the Go struct for the "v2.core.account[defaults].updated" event.
// This event occurs when account defaults are created or updated.
type V2CoreAccountIncludingDefaultsUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingDefaultsUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingDefaultsUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[defaults].updated"
// This event occurs when account defaults are created or updated.
type V2CoreAccountIncludingDefaultsUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingDefaultsUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingDefaultsUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingDefaultsUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingDefaultsUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingDefaultsUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingFutureRequirementsUpdatedEvent is the Go struct for the "v2.core.account[future_requirements].updated" event.
// Occurs when an Account's future requirements are updated.
type V2CoreAccountIncludingFutureRequirementsUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingFutureRequirementsUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingFutureRequirementsUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[future_requirements].updated"
// Occurs when an Account's future requirements are updated.
type V2CoreAccountIncludingFutureRequirementsUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingFutureRequirementsUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingFutureRequirementsUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingFutureRequirementsUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingFutureRequirementsUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingFutureRequirementsUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingIdentityUpdatedEvent is the Go struct for the "v2.core.account[identity].updated" event.
// Occurs when an Identity is updated.
type V2CoreAccountIncludingIdentityUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingIdentityUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingIdentityUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[identity].updated"
// Occurs when an Identity is updated.
type V2CoreAccountIncludingIdentityUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingIdentityUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingIdentityUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingIdentityUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingIdentityUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingIdentityUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingRequirementsUpdatedEvent is the Go struct for the "v2.core.account[requirements].updated" event.
// Occurs when an Account's requirements are updated.
type V2CoreAccountIncludingRequirementsUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingRequirementsUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingRequirementsUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[requirements].updated"
// Occurs when an Account's requirements are updated.
type V2CoreAccountIncludingRequirementsUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingRequirementsUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingRequirementsUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingRequirementsUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingRequirementsUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingRequirementsUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountLinkReturnedEvent is the Go struct for the "v2.core.account_link.returned" event.
// Occurs when the generated AccountLink is completed.
type V2CoreAccountLinkReturnedEvent struct {
	V2BaseEvent
	Data V2CoreAccountLinkReturnedEventData `json:"data"`
}

// V2CoreAccountLinkReturnedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account_link.returned"
// Occurs when the generated AccountLink is completed.
type V2CoreAccountLinkReturnedEventNotification struct {
	V2CoreEventNotification
}

// FetchEvent retrieves the V2CoreAccountLinkReturnedEvent that created this Notification
func (n *V2CoreAccountLinkReturnedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountLinkReturnedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountLinkReturnedEvent), nil
}

// V2CoreAccountPersonCreatedEvent is the Go struct for the "v2.core.account_person.created" event.
// Occurs when a Person is created.
type V2CoreAccountPersonCreatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountPersonCreatedEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject            `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccountPerson, error)
}

// FetchRelatedObject fetches the V2CoreAccountPerson related to the event.
func (e *V2CoreAccountPersonCreatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccountPerson, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountPersonCreatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account_person.created"
// Occurs when a Person is created.
type V2CoreAccountPersonCreatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountPersonCreatedEvent that created this Notification
func (n *V2CoreAccountPersonCreatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountPersonCreatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountPersonCreatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccountPerson related to the event.
func (n *V2CoreAccountPersonCreatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccountPerson, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccountPerson{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountPersonDeletedEvent is the Go struct for the "v2.core.account_person.deleted" event.
// Occurs when a Person is deleted.
type V2CoreAccountPersonDeletedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountPersonDeletedEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject            `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccountPerson, error)
}

// FetchRelatedObject fetches the V2CoreAccountPerson related to the event.
func (e *V2CoreAccountPersonDeletedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccountPerson, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountPersonDeletedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account_person.deleted"
// Occurs when a Person is deleted.
type V2CoreAccountPersonDeletedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountPersonDeletedEvent that created this Notification
func (n *V2CoreAccountPersonDeletedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountPersonDeletedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountPersonDeletedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccountPerson related to the event.
func (n *V2CoreAccountPersonDeletedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccountPerson, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccountPerson{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountPersonUpdatedEvent is the Go struct for the "v2.core.account_person.updated" event.
// Occurs when a Person is updated.
type V2CoreAccountPersonUpdatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountPersonUpdatedEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject            `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccountPerson, error)
}

// FetchRelatedObject fetches the V2CoreAccountPerson related to the event.
func (e *V2CoreAccountPersonUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccountPerson, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountPersonUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account_person.updated"
// Occurs when a Person is updated.
type V2CoreAccountPersonUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountPersonUpdatedEvent that created this Notification
func (n *V2CoreAccountPersonUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountPersonUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountPersonUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccountPerson related to the event.
func (n *V2CoreAccountPersonUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccountPerson, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccountPerson{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreEventDestinationPingEvent is the Go struct for the "v2.core.event_destination.ping" event.
// A ping event used to test the connection to an EventDestination.
type V2CoreEventDestinationPingEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreEventDestination, error)
}

// FetchRelatedObject fetches the V2CoreEventDestination related to the event.
func (e *V2CoreEventDestinationPingEvent) FetchRelatedObject(ctx context.Context) (*V2CoreEventDestination, error) {
	return e.fetchRelatedObject()
}

// V2CoreEventDestinationPingEventNotification is the webhook payload you'll get when handling an event with type "v2.core.event_destination.ping"
// A ping event used to test the connection to an EventDestination.
type V2CoreEventDestinationPingEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreEventDestinationPingEvent that created this Notification
func (n *V2CoreEventDestinationPingEventNotification) FetchEvent(ctx context.Context) (*V2CoreEventDestinationPingEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreEventDestinationPingEvent), nil
}

// FetchRelatedObject fetches the V2CoreEventDestination related to the event.
func (n *V2CoreEventDestinationPingEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreEventDestination, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreEventDestination{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// The request causes the error.
type V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeSampleErrorRequest struct {
	// The request idempotency key.
	Identifier string `json:"identifier"`
}

// A list of sample errors of this type.
type V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeSampleError struct {
	// The error message.
	ErrorMessage string `json:"error_message"`
	// The request causes the error.
	Request *V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeSampleErrorRequest `json:"request"`
}

// The error details.
type V1BillingMeterErrorReportTriggeredEventDataReasonErrorType struct {
	// Open Enum.
	Code V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode `json:"code"`
	// The number of errors of this type.
	ErrorCount int64 `json:"error_count"`
	// A list of sample errors of this type.
	SampleErrors []*V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeSampleError `json:"sample_errors"`
}

// This contains information about why meter error happens.
type V1BillingMeterErrorReportTriggeredEventDataReason struct {
	// The total error count within this window.
	ErrorCount int64 `json:"error_count"`
	// The error details.
	ErrorTypes []*V1BillingMeterErrorReportTriggeredEventDataReasonErrorType `json:"error_types"`
}

// Occurs when a Meter has invalid async usage events.
type V1BillingMeterErrorReportTriggeredEventData struct {
	// Extra field included in the event's `data` when fetched from /v2/events.
	DeveloperMessageSummary string `json:"developer_message_summary"`
	// This contains information about why meter error happens.
	Reason *V1BillingMeterErrorReportTriggeredEventDataReason `json:"reason"`
	// The end of the window that is encapsulated by this summary.
	ValidationEnd time.Time `json:"validation_end"`
	// The start of the window that is encapsulated by this summary.
	ValidationStart time.Time `json:"validation_start"`
}

// The request causes the error.
type V1BillingMeterNoMeterFoundEventDataReasonErrorTypeSampleErrorRequest struct {
	// The request idempotency key.
	Identifier string `json:"identifier"`
}

// A list of sample errors of this type.
type V1BillingMeterNoMeterFoundEventDataReasonErrorTypeSampleError struct {
	// The error message.
	ErrorMessage string `json:"error_message"`
	// The request causes the error.
	Request *V1BillingMeterNoMeterFoundEventDataReasonErrorTypeSampleErrorRequest `json:"request"`
}

// The error details.
type V1BillingMeterNoMeterFoundEventDataReasonErrorType struct {
	// Open Enum.
	Code V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode `json:"code"`
	// The number of errors of this type.
	ErrorCount int64 `json:"error_count"`
	// A list of sample errors of this type.
	SampleErrors []*V1BillingMeterNoMeterFoundEventDataReasonErrorTypeSampleError `json:"sample_errors"`
}

// This contains information about why meter error happens.
type V1BillingMeterNoMeterFoundEventDataReason struct {
	// The total error count within this window.
	ErrorCount int64 `json:"error_count"`
	// The error details.
	ErrorTypes []*V1BillingMeterNoMeterFoundEventDataReasonErrorType `json:"error_types"`
}

// Occurs when a Meter's id is missing or invalid in async usage events.
type V1BillingMeterNoMeterFoundEventData struct {
	// Extra field included in the event's `data` when fetched from /v2/events.
	DeveloperMessageSummary string `json:"developer_message_summary"`
	// This contains information about why meter error happens.
	Reason *V1BillingMeterNoMeterFoundEventDataReason `json:"reason"`
	// The end of the window that is encapsulated by this summary.
	ValidationEnd time.Time `json:"validation_end"`
	// The start of the window that is encapsulated by this summary.
	ValidationStart time.Time `json:"validation_start"`
}

// Occurs when the status of an Account's customer configuration capability is updated.
type V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventData struct {
	// Open Enum. The capability which had its status updated.
	UpdatedCapability V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapability `json:"updated_capability"`
}

// Occurs when the status of an Account's merchant configuration capability is updated.
type V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventData struct {
	// Open Enum. The capability which had its status updated.
	UpdatedCapability V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability `json:"updated_capability"`
}

// Occurs when the status of an Account's recipient configuration capability is updated.
type V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventData struct {
	// Open Enum. The capability which had its status updated.
	UpdatedCapability V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability `json:"updated_capability"`
}

// Occurs when the generated AccountLink is completed.
type V2CoreAccountLinkReturnedEventData struct {
	// The ID of the v2 account.
	AccountID string `json:"account_id"`
	// Configurations on the Account that was onboarded via the account link.
	Configurations []V2CoreAccountLinkReturnedEventDataConfiguration `json:"configurations"`
	// Open Enum. The use case type of the account link that has been completed.
	UseCase V2CoreAccountLinkReturnedEventDataUseCase `json:"use_case"`
}

// Occurs when a Person is created.
type V2CoreAccountPersonCreatedEventData struct {
	// The ID of the v2 account.
	AccountID string `json:"account_id"`
}

// Occurs when a Person is deleted.
type V2CoreAccountPersonDeletedEventData struct {
	// The ID of the v2 account.
	AccountID string `json:"account_id"`
}

// Occurs when a Person is updated.
type V2CoreAccountPersonUpdatedEventData struct {
	// The ID of the v2 account.
	AccountID string `json:"account_id"`
}

// ConvertRawEvent converts a raw event to a concrete event type.
// If the event type is not known, it returns the raw event.
func ConvertRawEvent(event *V2CoreRawEvent, backend Backend, key string) (V2CoreEvent, error) {
	switch event.Type {
	case "v1.billing.meter.error_report_triggered":
		result := &V1BillingMeterErrorReportTriggeredEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*BillingMeter, error) {
			v := &BillingMeter{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v1.billing.meter.no_meter_found":
		result := &V1BillingMeterNoMeterFoundEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account.closed":
		result := &V2CoreAccountClosedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.core.account.created":
		result := &V2CoreAccountCreatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.core.account.updated":
		result := &V2CoreAccountUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.core.account[configuration.customer].capability_status_updated":
		result := &V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account[configuration.customer].updated":
		result := &V2CoreAccountIncludingConfigurationCustomerUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.core.account[configuration.merchant].capability_status_updated":
		result := &V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account[configuration.merchant].updated":
		result := &V2CoreAccountIncludingConfigurationMerchantUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.core.account[configuration.recipient].capability_status_updated":
		result := &V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account[configuration.recipient].updated":
		result := &V2CoreAccountIncludingConfigurationRecipientUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.core.account[defaults].updated":
		result := &V2CoreAccountIncludingDefaultsUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.core.account[future_requirements].updated":
		result := &V2CoreAccountIncludingFutureRequirementsUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.core.account[identity].updated":
		result := &V2CoreAccountIncludingIdentityUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.core.account[requirements].updated":
		result := &V2CoreAccountIncludingRequirementsUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccount, error) {
			v := &V2CoreAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.core.account_link.returned":
		result := &V2CoreAccountLinkReturnedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account_person.created":
		result := &V2CoreAccountPersonCreatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccountPerson, error) {
			v := &V2CoreAccountPerson{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account_person.deleted":
		result := &V2CoreAccountPersonDeletedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccountPerson, error) {
			v := &V2CoreAccountPerson{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account_person.updated":
		result := &V2CoreAccountPersonUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccountPerson, error) {
			v := &V2CoreAccountPerson{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.event_destination.ping":
		result := &V2CoreEventDestinationPingEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreEventDestination, error) {
			v := &V2CoreEventDestination{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	default:
		return event, nil
	}
}

// V2Events: The end of the section generated from our OpenAPI spec

// EventNotificationFromJSON is a helper for constructing an Event Notification. Doesn't perform signature validation,
// so you should use [Client.ParseEventNotification] instead for initial handling.
// This is useful in unit tests and working with EventNotifications that you've already validated the authenticity of.
func EventNotificationFromJSON(payload []byte, client Client) (EventNotificationContainer, error) {
	// we can pull the type out to base our matching on
	var result = &struct {
		Type string `json:"type"`
	}{}
	if err := json.Unmarshal(payload, result); err != nil {
		return nil, err
	}

	// V2EventNotificationTypes: The beginning of the section generated from our OpenAPI spec
	switch result.Type {
	case "v1.billing.meter.error_report_triggered":
		evt := V1BillingMeterErrorReportTriggeredEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v1.billing.meter.no_meter_found":
		evt := V1BillingMeterNoMeterFoundEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account.closed":
		evt := V2CoreAccountClosedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account.created":
		evt := V2CoreAccountCreatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account.updated":
		evt := V2CoreAccountUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[configuration.customer].capability_status_updated":
		evt := V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[configuration.customer].updated":
		evt := V2CoreAccountIncludingConfigurationCustomerUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[configuration.merchant].capability_status_updated":
		evt := V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[configuration.merchant].updated":
		evt := V2CoreAccountIncludingConfigurationMerchantUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[configuration.recipient].capability_status_updated":
		evt := V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[configuration.recipient].updated":
		evt := V2CoreAccountIncludingConfigurationRecipientUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[defaults].updated":
		evt := V2CoreAccountIncludingDefaultsUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[future_requirements].updated":
		evt := V2CoreAccountIncludingFutureRequirementsUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[identity].updated":
		evt := V2CoreAccountIncludingIdentityUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[requirements].updated":
		evt := V2CoreAccountIncludingRequirementsUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account_link.returned":
		evt := V2CoreAccountLinkReturnedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account_person.created":
		evt := V2CoreAccountPersonCreatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account_person.deleted":
		evt := V2CoreAccountPersonDeletedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account_person.updated":
		evt := V2CoreAccountPersonUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.event_destination.ping":
		evt := V2CoreEventDestinationPingEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	default:
		evt := UnknownEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	}
	// V2EventNotificationTypes: The end of the section generated from our OpenAPI spec
}
