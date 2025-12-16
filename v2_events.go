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

// Open Enum. The capability which had its status updated.
type V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability string

// List of values that V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability can take
const (
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityFinancialAddresssesBankAccounts    V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "financial_addressses.bank_accounts"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityFinancialAddresssesCryptoWallets   V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "financial_addressses.crypto_wallets"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityHoldsCurrenciesEUR                 V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "holds_currencies.eur"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityHoldsCurrenciesGBP                 V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "holds_currencies.gbp"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityHoldsCurrenciesUSD                 V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "holds_currencies.usd"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityHoldsCurrenciesUsdc                V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "holds_currencies.usdc"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityInboundTransfersBankAccounts       V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "inbound_transfers.bank_accounts"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundPaymentsBankAccounts       V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_payments.bank_accounts"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundPaymentsCards              V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_payments.cards"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundPaymentsCryptoWallets      V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_payments.crypto_wallets"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundPaymentsFinancialAccounts  V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_payments.financial_accounts"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundTransfersBankAccounts      V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_transfers.bank_accounts"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundTransfersCryptoWallets     V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_transfers.crypto_wallets"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundTransfersFinancialAccounts V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_transfers.financial_accounts"
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
	V2CoreAccountLinkReturnedEventDataConfigurationStorer    V2CoreAccountLinkReturnedEventDataConfiguration = "storer"
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.storer].capability_status_updated" event.
// Occurs when the status of an Account's storer configuration capability is updated.
type V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject                                                  `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[configuration.storer].capability_status_updated"
// Occurs when the status of an Account's storer configuration capability is updated.
type V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreAccountIncludingConfigurationStorerUpdatedEvent is the Go struct for the "v2.core.account[configuration.storer].updated" event.
// Occurs when a Storer's configuration is updated.
type V2CoreAccountIncludingConfigurationStorerUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (e *V2CoreAccountIncludingConfigurationStorerUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.account[configuration.storer].updated"
// Occurs when a Storer's configuration is updated.
type V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2CoreAccountIncludingConfigurationStorerUpdatedEvent that created this Notification
func (n *V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2CoreAccountIncludingConfigurationStorerUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreAccountIncludingConfigurationStorerUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2CoreAccount related to the event.
func (n *V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2CoreAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2CoreAccount{}
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
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
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2CoreHealthEventGenerationFailureResolvedEvent is the Go struct for the "v2.core.health.event_generation_failure.resolved" event.
// Occurs when an event generation failure alert is resolved.
type V2CoreHealthEventGenerationFailureResolvedEvent struct {
	V2BaseEvent
	Data V2CoreHealthEventGenerationFailureResolvedEventData `json:"data"`
}

// V2CoreHealthEventGenerationFailureResolvedEventNotification is the webhook payload you'll get when handling an event with type "v2.core.health.event_generation_failure.resolved"
// Occurs when an event generation failure alert is resolved.
type V2CoreHealthEventGenerationFailureResolvedEventNotification struct {
	V2CoreEventNotification
}

// FetchEvent retrieves the V2CoreHealthEventGenerationFailureResolvedEvent that created this Notification
func (n *V2CoreHealthEventGenerationFailureResolvedEventNotification) FetchEvent(ctx context.Context) (*V2CoreHealthEventGenerationFailureResolvedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreHealthEventGenerationFailureResolvedEvent), nil
}

// V2MoneyManagementAdjustmentCreatedEvent is the Go struct for the "v2.money_management.adjustment.created" event.
// Occurs when an Adjustment is created.
type V2MoneyManagementAdjustmentCreatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementAdjustment, error)
}

// FetchRelatedObject fetches the V2MoneyManagementAdjustment related to the event.
func (e *V2MoneyManagementAdjustmentCreatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementAdjustment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementAdjustmentCreatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.adjustment.created"
// Occurs when an Adjustment is created.
type V2MoneyManagementAdjustmentCreatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementAdjustmentCreatedEvent that created this Notification
func (n *V2MoneyManagementAdjustmentCreatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementAdjustmentCreatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementAdjustmentCreatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementAdjustment related to the event.
func (n *V2MoneyManagementAdjustmentCreatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementAdjustment, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementAdjustment{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementFinancialAccountCreatedEvent is the Go struct for the "v2.money_management.financial_account.created" event.
// Occurs when a FinancialAccount is created.
type V2MoneyManagementFinancialAccountCreatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementFinancialAccount, error)
}

// FetchRelatedObject fetches the V2MoneyManagementFinancialAccount related to the event.
func (e *V2MoneyManagementFinancialAccountCreatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementFinancialAccount, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementFinancialAccountCreatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.financial_account.created"
// Occurs when a FinancialAccount is created.
type V2MoneyManagementFinancialAccountCreatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementFinancialAccountCreatedEvent that created this Notification
func (n *V2MoneyManagementFinancialAccountCreatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementFinancialAccountCreatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementFinancialAccountCreatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementFinancialAccount related to the event.
func (n *V2MoneyManagementFinancialAccountCreatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementFinancialAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementFinancialAccount{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementFinancialAccountUpdatedEvent is the Go struct for the "v2.money_management.financial_account.updated" event.
// Occurs when a FinancialAccount is updated.
type V2MoneyManagementFinancialAccountUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementFinancialAccount, error)
}

// FetchRelatedObject fetches the V2MoneyManagementFinancialAccount related to the event.
func (e *V2MoneyManagementFinancialAccountUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementFinancialAccount, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementFinancialAccountUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.financial_account.updated"
// Occurs when a FinancialAccount is updated.
type V2MoneyManagementFinancialAccountUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementFinancialAccountUpdatedEvent that created this Notification
func (n *V2MoneyManagementFinancialAccountUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementFinancialAccountUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementFinancialAccountUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementFinancialAccount related to the event.
func (n *V2MoneyManagementFinancialAccountUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementFinancialAccount, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementFinancialAccount{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementFinancialAddressActivatedEvent is the Go struct for the "v2.money_management.financial_address.activated" event.
// Occurs when a FinancialAddress is activated and is ready to receive funds.
type V2MoneyManagementFinancialAddressActivatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementFinancialAddress, error)
}

// FetchRelatedObject fetches the V2MoneyManagementFinancialAddress related to the event.
func (e *V2MoneyManagementFinancialAddressActivatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementFinancialAddress, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementFinancialAddressActivatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.financial_address.activated"
// Occurs when a FinancialAddress is activated and is ready to receive funds.
type V2MoneyManagementFinancialAddressActivatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementFinancialAddressActivatedEvent that created this Notification
func (n *V2MoneyManagementFinancialAddressActivatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementFinancialAddressActivatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementFinancialAddressActivatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementFinancialAddress related to the event.
func (n *V2MoneyManagementFinancialAddressActivatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementFinancialAddress, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementFinancialAddress{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementFinancialAddressFailedEvent is the Go struct for the "v2.money_management.financial_address.failed" event.
// Occurs when a FinancialAddress fails to activate and can not receive funds.
type V2MoneyManagementFinancialAddressFailedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementFinancialAddress, error)
}

// FetchRelatedObject fetches the V2MoneyManagementFinancialAddress related to the event.
func (e *V2MoneyManagementFinancialAddressFailedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementFinancialAddress, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementFinancialAddressFailedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.financial_address.failed"
// Occurs when a FinancialAddress fails to activate and can not receive funds.
type V2MoneyManagementFinancialAddressFailedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementFinancialAddressFailedEvent that created this Notification
func (n *V2MoneyManagementFinancialAddressFailedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementFinancialAddressFailedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementFinancialAddressFailedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementFinancialAddress related to the event.
func (n *V2MoneyManagementFinancialAddressFailedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementFinancialAddress, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementFinancialAddress{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementInboundTransferAvailableEvent is the Go struct for the "v2.money_management.inbound_transfer.available" event.
// Occurs when an InboundTransfer's funds are made available.
type V2MoneyManagementInboundTransferAvailableEvent struct {
	V2BaseEvent
	Data               V2MoneyManagementInboundTransferAvailableEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject                           `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (e *V2MoneyManagementInboundTransferAvailableEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferAvailableEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.inbound_transfer.available"
// Occurs when an InboundTransfer's funds are made available.
type V2MoneyManagementInboundTransferAvailableEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementInboundTransferAvailableEvent that created this Notification
func (n *V2MoneyManagementInboundTransferAvailableEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementInboundTransferAvailableEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementInboundTransferAvailableEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (n *V2MoneyManagementInboundTransferAvailableEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementInboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementInboundTransferBankDebitFailedEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_failed" event.
// Occurs when an InboundTransfer fails.
type V2MoneyManagementInboundTransferBankDebitFailedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (e *V2MoneyManagementInboundTransferBankDebitFailedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitFailedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.inbound_transfer.bank_debit_failed"
// Occurs when an InboundTransfer fails.
type V2MoneyManagementInboundTransferBankDebitFailedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementInboundTransferBankDebitFailedEvent that created this Notification
func (n *V2MoneyManagementInboundTransferBankDebitFailedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementInboundTransferBankDebitFailedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementInboundTransferBankDebitFailedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (n *V2MoneyManagementInboundTransferBankDebitFailedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementInboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementInboundTransferBankDebitProcessingEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_processing" event.
// Occurs when an InboundTransfer starts processing.
type V2MoneyManagementInboundTransferBankDebitProcessingEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (e *V2MoneyManagementInboundTransferBankDebitProcessingEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitProcessingEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.inbound_transfer.bank_debit_processing"
// Occurs when an InboundTransfer starts processing.
type V2MoneyManagementInboundTransferBankDebitProcessingEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementInboundTransferBankDebitProcessingEvent that created this Notification
func (n *V2MoneyManagementInboundTransferBankDebitProcessingEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementInboundTransferBankDebitProcessingEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementInboundTransferBankDebitProcessingEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (n *V2MoneyManagementInboundTransferBankDebitProcessingEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementInboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementInboundTransferBankDebitQueuedEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_queued" event.
// Occurs when an InboundTransfer is queued.
type V2MoneyManagementInboundTransferBankDebitQueuedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (e *V2MoneyManagementInboundTransferBankDebitQueuedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitQueuedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.inbound_transfer.bank_debit_queued"
// Occurs when an InboundTransfer is queued.
type V2MoneyManagementInboundTransferBankDebitQueuedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementInboundTransferBankDebitQueuedEvent that created this Notification
func (n *V2MoneyManagementInboundTransferBankDebitQueuedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementInboundTransferBankDebitQueuedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementInboundTransferBankDebitQueuedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (n *V2MoneyManagementInboundTransferBankDebitQueuedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementInboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementInboundTransferBankDebitReturnedEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_returned" event.
// Occurs when an InboundTransfer is returned.
type V2MoneyManagementInboundTransferBankDebitReturnedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (e *V2MoneyManagementInboundTransferBankDebitReturnedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitReturnedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.inbound_transfer.bank_debit_returned"
// Occurs when an InboundTransfer is returned.
type V2MoneyManagementInboundTransferBankDebitReturnedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementInboundTransferBankDebitReturnedEvent that created this Notification
func (n *V2MoneyManagementInboundTransferBankDebitReturnedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementInboundTransferBankDebitReturnedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementInboundTransferBankDebitReturnedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (n *V2MoneyManagementInboundTransferBankDebitReturnedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementInboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementInboundTransferBankDebitSucceededEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_succeeded" event.
// Occurs when an InboundTransfer succeeds.
type V2MoneyManagementInboundTransferBankDebitSucceededEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (e *V2MoneyManagementInboundTransferBankDebitSucceededEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitSucceededEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.inbound_transfer.bank_debit_succeeded"
// Occurs when an InboundTransfer succeeds.
type V2MoneyManagementInboundTransferBankDebitSucceededEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementInboundTransferBankDebitSucceededEvent that created this Notification
func (n *V2MoneyManagementInboundTransferBankDebitSucceededEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementInboundTransferBankDebitSucceededEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementInboundTransferBankDebitSucceededEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementInboundTransfer related to the event.
func (n *V2MoneyManagementInboundTransferBankDebitSucceededEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementInboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementInboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundPaymentCanceledEvent is the Go struct for the "v2.money_management.outbound_payment.canceled" event.
// Occurs when an OutboundPayment transitions into the canceled state.
type V2MoneyManagementOutboundPaymentCanceledEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (e *V2MoneyManagementOutboundPaymentCanceledEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentCanceledEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_payment.canceled"
// Occurs when an OutboundPayment transitions into the canceled state.
type V2MoneyManagementOutboundPaymentCanceledEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundPaymentCanceledEvent that created this Notification
func (n *V2MoneyManagementOutboundPaymentCanceledEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundPaymentCanceledEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundPaymentCanceledEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (n *V2MoneyManagementOutboundPaymentCanceledEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundPayment{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundPaymentCreatedEvent is the Go struct for the "v2.money_management.outbound_payment.created" event.
// Occurs when an OutboundPayment is created.
type V2MoneyManagementOutboundPaymentCreatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (e *V2MoneyManagementOutboundPaymentCreatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentCreatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_payment.created"
// Occurs when an OutboundPayment is created.
type V2MoneyManagementOutboundPaymentCreatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundPaymentCreatedEvent that created this Notification
func (n *V2MoneyManagementOutboundPaymentCreatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundPaymentCreatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundPaymentCreatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (n *V2MoneyManagementOutboundPaymentCreatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundPayment{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundPaymentFailedEvent is the Go struct for the "v2.money_management.outbound_payment.failed" event.
// Occurs when an OutboundPayment transitions into the failed state.
type V2MoneyManagementOutboundPaymentFailedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (e *V2MoneyManagementOutboundPaymentFailedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentFailedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_payment.failed"
// Occurs when an OutboundPayment transitions into the failed state.
type V2MoneyManagementOutboundPaymentFailedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundPaymentFailedEvent that created this Notification
func (n *V2MoneyManagementOutboundPaymentFailedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundPaymentFailedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundPaymentFailedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (n *V2MoneyManagementOutboundPaymentFailedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundPayment{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundPaymentPostedEvent is the Go struct for the "v2.money_management.outbound_payment.posted" event.
// Occurs when an OutboundPayment transitions into the posted state.
type V2MoneyManagementOutboundPaymentPostedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (e *V2MoneyManagementOutboundPaymentPostedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentPostedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_payment.posted"
// Occurs when an OutboundPayment transitions into the posted state.
type V2MoneyManagementOutboundPaymentPostedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundPaymentPostedEvent that created this Notification
func (n *V2MoneyManagementOutboundPaymentPostedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundPaymentPostedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundPaymentPostedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (n *V2MoneyManagementOutboundPaymentPostedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundPayment{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundPaymentReturnedEvent is the Go struct for the "v2.money_management.outbound_payment.returned" event.
// Occurs when an OutboundPayment transitions into the returned state.
type V2MoneyManagementOutboundPaymentReturnedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (e *V2MoneyManagementOutboundPaymentReturnedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentReturnedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_payment.returned"
// Occurs when an OutboundPayment transitions into the returned state.
type V2MoneyManagementOutboundPaymentReturnedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundPaymentReturnedEvent that created this Notification
func (n *V2MoneyManagementOutboundPaymentReturnedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundPaymentReturnedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundPaymentReturnedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (n *V2MoneyManagementOutboundPaymentReturnedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundPayment{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundPaymentUpdatedEvent is the Go struct for the "v2.money_management.outbound_payment.updated" event.
// Occurs when an OutboundPayment is updated.
type V2MoneyManagementOutboundPaymentUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (e *V2MoneyManagementOutboundPaymentUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_payment.updated"
// Occurs when an OutboundPayment is updated.
type V2MoneyManagementOutboundPaymentUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundPaymentUpdatedEvent that created this Notification
func (n *V2MoneyManagementOutboundPaymentUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundPaymentUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundPaymentUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundPayment related to the event.
func (n *V2MoneyManagementOutboundPaymentUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundPayment, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundPayment{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundTransferCanceledEvent is the Go struct for the "v2.money_management.outbound_transfer.canceled" event.
// Occurs when an OutboundTransfer transitions into the canceled state.
type V2MoneyManagementOutboundTransferCanceledEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (e *V2MoneyManagementOutboundTransferCanceledEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferCanceledEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_transfer.canceled"
// Occurs when an OutboundTransfer transitions into the canceled state.
type V2MoneyManagementOutboundTransferCanceledEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundTransferCanceledEvent that created this Notification
func (n *V2MoneyManagementOutboundTransferCanceledEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundTransferCanceledEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundTransferCanceledEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (n *V2MoneyManagementOutboundTransferCanceledEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundTransferCreatedEvent is the Go struct for the "v2.money_management.outbound_transfer.created" event.
// Occurs when an OutboundTransfer is created.
type V2MoneyManagementOutboundTransferCreatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (e *V2MoneyManagementOutboundTransferCreatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferCreatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_transfer.created"
// Occurs when an OutboundTransfer is created.
type V2MoneyManagementOutboundTransferCreatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundTransferCreatedEvent that created this Notification
func (n *V2MoneyManagementOutboundTransferCreatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundTransferCreatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundTransferCreatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (n *V2MoneyManagementOutboundTransferCreatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundTransferFailedEvent is the Go struct for the "v2.money_management.outbound_transfer.failed" event.
// Occurs when an OutboundTransfer transitions into the failed state.
type V2MoneyManagementOutboundTransferFailedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (e *V2MoneyManagementOutboundTransferFailedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferFailedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_transfer.failed"
// Occurs when an OutboundTransfer transitions into the failed state.
type V2MoneyManagementOutboundTransferFailedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundTransferFailedEvent that created this Notification
func (n *V2MoneyManagementOutboundTransferFailedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundTransferFailedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundTransferFailedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (n *V2MoneyManagementOutboundTransferFailedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundTransferPostedEvent is the Go struct for the "v2.money_management.outbound_transfer.posted" event.
// Occurs when an OutboundTransfer transitions into the posted state.
type V2MoneyManagementOutboundTransferPostedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (e *V2MoneyManagementOutboundTransferPostedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferPostedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_transfer.posted"
// Occurs when an OutboundTransfer transitions into the posted state.
type V2MoneyManagementOutboundTransferPostedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundTransferPostedEvent that created this Notification
func (n *V2MoneyManagementOutboundTransferPostedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundTransferPostedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundTransferPostedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (n *V2MoneyManagementOutboundTransferPostedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundTransferReturnedEvent is the Go struct for the "v2.money_management.outbound_transfer.returned" event.
// Occurs when an OutboundTransfer transitions into the returned state.
type V2MoneyManagementOutboundTransferReturnedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (e *V2MoneyManagementOutboundTransferReturnedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferReturnedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_transfer.returned"
// Occurs when an OutboundTransfer transitions into the returned state.
type V2MoneyManagementOutboundTransferReturnedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundTransferReturnedEvent that created this Notification
func (n *V2MoneyManagementOutboundTransferReturnedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundTransferReturnedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundTransferReturnedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (n *V2MoneyManagementOutboundTransferReturnedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementOutboundTransferUpdatedEvent is the Go struct for the "v2.money_management.outbound_transfer.updated" event.
// Event that is emitted every time an Outbound Transfer is updated.
type V2MoneyManagementOutboundTransferUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (e *V2MoneyManagementOutboundTransferUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.outbound_transfer.updated"
// Event that is emitted every time an Outbound Transfer is updated.
type V2MoneyManagementOutboundTransferUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementOutboundTransferUpdatedEvent that created this Notification
func (n *V2MoneyManagementOutboundTransferUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementOutboundTransferUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementOutboundTransferUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementOutboundTransfer related to the event.
func (n *V2MoneyManagementOutboundTransferUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementOutboundTransfer, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementOutboundTransfer{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementPayoutMethodCreatedEvent is the Go struct for the "v2.money_management.payout_method.created" event.
// Occurs when a PayoutMethod is created.
type V2MoneyManagementPayoutMethodCreatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementPayoutMethod, error)
}

// FetchRelatedObject fetches the V2MoneyManagementPayoutMethod related to the event.
func (e *V2MoneyManagementPayoutMethodCreatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementPayoutMethod, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementPayoutMethodCreatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.payout_method.created"
// Occurs when a PayoutMethod is created.
type V2MoneyManagementPayoutMethodCreatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementPayoutMethodCreatedEvent that created this Notification
func (n *V2MoneyManagementPayoutMethodCreatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementPayoutMethodCreatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementPayoutMethodCreatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementPayoutMethod related to the event.
func (n *V2MoneyManagementPayoutMethodCreatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementPayoutMethod, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementPayoutMethod{}
	err := n.client.backend.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementPayoutMethodUpdatedEvent is the Go struct for the "v2.money_management.payout_method.updated" event.
// Occurs when a PayoutMethod is updated.
type V2MoneyManagementPayoutMethodUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementPayoutMethod, error)
}

// FetchRelatedObject fetches the V2MoneyManagementPayoutMethod related to the event.
func (e *V2MoneyManagementPayoutMethodUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementPayoutMethod, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementPayoutMethodUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.payout_method.updated"
// Occurs when a PayoutMethod is updated.
type V2MoneyManagementPayoutMethodUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementPayoutMethodUpdatedEvent that created this Notification
func (n *V2MoneyManagementPayoutMethodUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementPayoutMethodUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementPayoutMethodUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementPayoutMethod related to the event.
func (n *V2MoneyManagementPayoutMethodUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementPayoutMethod, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementPayoutMethod{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementReceivedCreditAvailableEvent is the Go struct for the "v2.money_management.received_credit.available" event.
// Occurs when a ReceivedCredit's funds are received and are available in your balance.
type V2MoneyManagementReceivedCreditAvailableEvent struct {
	V2BaseEvent
	Data               V2MoneyManagementReceivedCreditAvailableEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject                          `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedCredit related to the event.
func (e *V2MoneyManagementReceivedCreditAvailableEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditAvailableEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.received_credit.available"
// Occurs when a ReceivedCredit's funds are received and are available in your balance.
type V2MoneyManagementReceivedCreditAvailableEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementReceivedCreditAvailableEvent that created this Notification
func (n *V2MoneyManagementReceivedCreditAvailableEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementReceivedCreditAvailableEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementReceivedCreditAvailableEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedCredit related to the event.
func (n *V2MoneyManagementReceivedCreditAvailableEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedCredit, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementReceivedCredit{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementReceivedCreditFailedEvent is the Go struct for the "v2.money_management.received_credit.failed" event.
// Occurs when a ReceivedCredit is attempted to your balance and fails. See the status_details for more information.
type V2MoneyManagementReceivedCreditFailedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedCredit related to the event.
func (e *V2MoneyManagementReceivedCreditFailedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditFailedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.received_credit.failed"
// Occurs when a ReceivedCredit is attempted to your balance and fails. See the status_details for more information.
type V2MoneyManagementReceivedCreditFailedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementReceivedCreditFailedEvent that created this Notification
func (n *V2MoneyManagementReceivedCreditFailedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementReceivedCreditFailedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementReceivedCreditFailedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedCredit related to the event.
func (n *V2MoneyManagementReceivedCreditFailedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedCredit, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementReceivedCredit{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementReceivedCreditReturnedEvent is the Go struct for the "v2.money_management.received_credit.returned" event.
// Occurs when a ReceivedCredit is reversed, returned to the originator, and deducted from your balance.
type V2MoneyManagementReceivedCreditReturnedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedCredit related to the event.
func (e *V2MoneyManagementReceivedCreditReturnedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditReturnedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.received_credit.returned"
// Occurs when a ReceivedCredit is reversed, returned to the originator, and deducted from your balance.
type V2MoneyManagementReceivedCreditReturnedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementReceivedCreditReturnedEvent that created this Notification
func (n *V2MoneyManagementReceivedCreditReturnedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementReceivedCreditReturnedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementReceivedCreditReturnedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedCredit related to the event.
func (n *V2MoneyManagementReceivedCreditReturnedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedCredit, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementReceivedCredit{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementReceivedCreditSucceededEvent is the Go struct for the "v2.money_management.received_credit.succeeded" event.
// Occurs when a ReceivedCredit succeeds.
type V2MoneyManagementReceivedCreditSucceededEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedCredit related to the event.
func (e *V2MoneyManagementReceivedCreditSucceededEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditSucceededEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.received_credit.succeeded"
// Occurs when a ReceivedCredit succeeds.
type V2MoneyManagementReceivedCreditSucceededEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementReceivedCreditSucceededEvent that created this Notification
func (n *V2MoneyManagementReceivedCreditSucceededEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementReceivedCreditSucceededEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementReceivedCreditSucceededEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedCredit related to the event.
func (n *V2MoneyManagementReceivedCreditSucceededEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedCredit, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementReceivedCredit{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementReceivedDebitCanceledEvent is the Go struct for the "v2.money_management.received_debit.canceled" event.
// Occurs when a ReceivedDebit is canceled.
type V2MoneyManagementReceivedDebitCanceledEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedDebit related to the event.
func (e *V2MoneyManagementReceivedDebitCanceledEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitCanceledEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.received_debit.canceled"
// Occurs when a ReceivedDebit is canceled.
type V2MoneyManagementReceivedDebitCanceledEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementReceivedDebitCanceledEvent that created this Notification
func (n *V2MoneyManagementReceivedDebitCanceledEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementReceivedDebitCanceledEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementReceivedDebitCanceledEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedDebit related to the event.
func (n *V2MoneyManagementReceivedDebitCanceledEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedDebit, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementReceivedDebit{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementReceivedDebitFailedEvent is the Go struct for the "v2.money_management.received_debit.failed" event.
// Occurs when a ReceivedDebit fails.
type V2MoneyManagementReceivedDebitFailedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedDebit related to the event.
func (e *V2MoneyManagementReceivedDebitFailedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitFailedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.received_debit.failed"
// Occurs when a ReceivedDebit fails.
type V2MoneyManagementReceivedDebitFailedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementReceivedDebitFailedEvent that created this Notification
func (n *V2MoneyManagementReceivedDebitFailedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementReceivedDebitFailedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementReceivedDebitFailedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedDebit related to the event.
func (n *V2MoneyManagementReceivedDebitFailedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedDebit, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementReceivedDebit{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementReceivedDebitPendingEvent is the Go struct for the "v2.money_management.received_debit.pending" event.
// Occurs when a ReceivedDebit is set to pending.
type V2MoneyManagementReceivedDebitPendingEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedDebit related to the event.
func (e *V2MoneyManagementReceivedDebitPendingEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitPendingEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.received_debit.pending"
// Occurs when a ReceivedDebit is set to pending.
type V2MoneyManagementReceivedDebitPendingEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementReceivedDebitPendingEvent that created this Notification
func (n *V2MoneyManagementReceivedDebitPendingEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementReceivedDebitPendingEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementReceivedDebitPendingEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedDebit related to the event.
func (n *V2MoneyManagementReceivedDebitPendingEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedDebit, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementReceivedDebit{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementReceivedDebitSucceededEvent is the Go struct for the "v2.money_management.received_debit.succeeded" event.
// Occurs when a ReceivedDebit succeeds.
type V2MoneyManagementReceivedDebitSucceededEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedDebit related to the event.
func (e *V2MoneyManagementReceivedDebitSucceededEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitSucceededEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.received_debit.succeeded"
// Occurs when a ReceivedDebit succeeds.
type V2MoneyManagementReceivedDebitSucceededEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementReceivedDebitSucceededEvent that created this Notification
func (n *V2MoneyManagementReceivedDebitSucceededEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementReceivedDebitSucceededEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementReceivedDebitSucceededEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedDebit related to the event.
func (n *V2MoneyManagementReceivedDebitSucceededEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedDebit, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementReceivedDebit{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementReceivedDebitUpdatedEvent is the Go struct for the "v2.money_management.received_debit.updated" event.
// Occurs when a ReceivedDebit is updated.
type V2MoneyManagementReceivedDebitUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedDebit related to the event.
func (e *V2MoneyManagementReceivedDebitUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.received_debit.updated"
// Occurs when a ReceivedDebit is updated.
type V2MoneyManagementReceivedDebitUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementReceivedDebitUpdatedEvent that created this Notification
func (n *V2MoneyManagementReceivedDebitUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementReceivedDebitUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementReceivedDebitUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementReceivedDebit related to the event.
func (n *V2MoneyManagementReceivedDebitUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementReceivedDebit, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementReceivedDebit{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementTransactionCreatedEvent is the Go struct for the "v2.money_management.transaction.created" event.
// Occurs when a Transaction is created.
type V2MoneyManagementTransactionCreatedEvent struct {
	V2BaseEvent
	Data               V2MoneyManagementTransactionCreatedEventData `json:"data"`
	RelatedObject      V2CoreEventRelatedObject                     `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementTransaction, error)
}

// FetchRelatedObject fetches the V2MoneyManagementTransaction related to the event.
func (e *V2MoneyManagementTransactionCreatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementTransaction, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementTransactionCreatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.transaction.created"
// Occurs when a Transaction is created.
type V2MoneyManagementTransactionCreatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementTransactionCreatedEvent that created this Notification
func (n *V2MoneyManagementTransactionCreatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementTransactionCreatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementTransactionCreatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementTransaction related to the event.
func (n *V2MoneyManagementTransactionCreatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementTransaction, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementTransaction{}
	err := n.client.backends.API.Call(
		http.MethodGet, n.RelatedObject.URL, n.client.key, params, relatedObj)
	return relatedObj, err
}

// V2MoneyManagementTransactionUpdatedEvent is the Go struct for the "v2.money_management.transaction.updated" event.
// Occurs when a Transaction is updated.
type V2MoneyManagementTransactionUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      V2CoreEventRelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementTransaction, error)
}

// FetchRelatedObject fetches the V2MoneyManagementTransaction related to the event.
func (e *V2MoneyManagementTransactionUpdatedEvent) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementTransaction, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementTransactionUpdatedEventNotification is the webhook payload you'll get when handling an event with type "v2.money_management.transaction.updated"
// Occurs when a Transaction is updated.
type V2MoneyManagementTransactionUpdatedEventNotification struct {
	V2CoreEventNotification
	RelatedObject V2CoreEventRelatedObject `json:"related_object"`
}

// FetchEvent retrieves the V2MoneyManagementTransactionUpdatedEvent that created this Notification
func (n *V2MoneyManagementTransactionUpdatedEventNotification) FetchEvent(ctx context.Context) (*V2MoneyManagementTransactionUpdatedEvent, error) {
	evt, err := n.V2CoreEventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2MoneyManagementTransactionUpdatedEvent), nil
}

// FetchRelatedObject fetches the V2MoneyManagementTransaction related to the event.
func (n *V2MoneyManagementTransactionUpdatedEventNotification) FetchRelatedObject(ctx context.Context) (*V2MoneyManagementTransaction, error) {
	params := &eventNotificationParams{Params: Params{Context: ctx}}
	params.SetStripeContextFrom(n.Context)
	relatedObj := &V2MoneyManagementTransaction{}
	err := n.client.backends.API.Call(
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

// Occurs when the status of an Account's storer configuration capability is updated.
type V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventData struct {
	// Open Enum. The capability which had its status updated.
	UpdatedCapability V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability `json:"updated_capability"`
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

// The related object details.
type V2CoreHealthEventGenerationFailureResolvedEventDataImpactRelatedObject struct {
	// The ID of the related object (e.g., "pi_...").
	ID string `json:"id"`
	// The type of the related object (e.g., "payment_intent").
	Type string `json:"type"`
	// The API URL for the related object (e.g., "/v1/payment_intents/pi_...").
	URL string `json:"url"`
}

// The user impact.
type V2CoreHealthEventGenerationFailureResolvedEventDataImpact struct {
	// The context the event should have been generated for. Only present when the account is a connected account.
	Context string `json:"context,omitempty"`
	// The type of event that Stripe failed to generate.
	EventType string `json:"event_type"`
	// The related object details.
	RelatedObject *V2CoreHealthEventGenerationFailureResolvedEventDataImpactRelatedObject `json:"related_object"`
}

// Occurs when an event generation failure alert is resolved.
type V2CoreHealthEventGenerationFailureResolvedEventData struct {
	// The alert ID.
	AlertID string `json:"alert_id"`
	// The grouping key for the alert.
	GroupingKey string `json:"grouping_key"`
	// The user impact.
	Impact *V2CoreHealthEventGenerationFailureResolvedEventDataImpact `json:"impact"`
	// The time when the user experience has returned to expected levels.
	ResolvedAt time.Time `json:"resolved_at"`
	// A short description of the alert.
	Summary string `json:"summary"`
}

// Occurs when an InboundTransfer's funds are made available.
type V2MoneyManagementInboundTransferAvailableEventData struct {
	// The transaction ID of the received credit.
	TransactionID string `json:"transaction_id"`
}

// Occurs when a ReceivedCredit's funds are received and are available in your balance.
type V2MoneyManagementReceivedCreditAvailableEventData struct {
	// The transaction ID of the received credit.
	TransactionID string `json:"transaction_id"`
}

// Occurs when a Transaction is created.
type V2MoneyManagementTransactionCreatedEventData struct {
	// Id of the v1 Transaction corresponding to this Transaction.
	V1ID string `json:"v1_id,omitempty"`
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
	case "v2.core.account[configuration.storer].capability_status_updated":
		result := &V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent{}
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
	case "v2.core.account[configuration.storer].updated":
		result := &V2CoreAccountIncludingConfigurationStorerUpdatedEvent{}
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
	case "v2.core.health.event_generation_failure.resolved":
		result := &V2CoreHealthEventGenerationFailureResolvedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.money_management.adjustment.created":
		result := &V2MoneyManagementAdjustmentCreatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementAdjustment, error) {
			v := &V2MoneyManagementAdjustment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.financial_account.created":
		result := &V2MoneyManagementFinancialAccountCreatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementFinancialAccount, error) {
			v := &V2MoneyManagementFinancialAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.financial_account.updated":
		result := &V2MoneyManagementFinancialAccountUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementFinancialAccount, error) {
			v := &V2MoneyManagementFinancialAccount{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.financial_address.activated":
		result := &V2MoneyManagementFinancialAddressActivatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementFinancialAddress, error) {
			v := &V2MoneyManagementFinancialAddress{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.financial_address.failed":
		result := &V2MoneyManagementFinancialAddressFailedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementFinancialAddress, error) {
			v := &V2MoneyManagementFinancialAddress{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.inbound_transfer.available":
		result := &V2MoneyManagementInboundTransferAvailableEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementInboundTransfer, error) {
			v := &V2MoneyManagementInboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.money_management.inbound_transfer.bank_debit_failed":
		result := &V2MoneyManagementInboundTransferBankDebitFailedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementInboundTransfer, error) {
			v := &V2MoneyManagementInboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.inbound_transfer.bank_debit_processing":
		result := &V2MoneyManagementInboundTransferBankDebitProcessingEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementInboundTransfer, error) {
			v := &V2MoneyManagementInboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.inbound_transfer.bank_debit_queued":
		result := &V2MoneyManagementInboundTransferBankDebitQueuedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementInboundTransfer, error) {
			v := &V2MoneyManagementInboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.inbound_transfer.bank_debit_returned":
		result := &V2MoneyManagementInboundTransferBankDebitReturnedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementInboundTransfer, error) {
			v := &V2MoneyManagementInboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.inbound_transfer.bank_debit_succeeded":
		result := &V2MoneyManagementInboundTransferBankDebitSucceededEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementInboundTransfer, error) {
			v := &V2MoneyManagementInboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_payment.canceled":
		result := &V2MoneyManagementOutboundPaymentCanceledEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundPayment, error) {
			v := &V2MoneyManagementOutboundPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_payment.created":
		result := &V2MoneyManagementOutboundPaymentCreatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundPayment, error) {
			v := &V2MoneyManagementOutboundPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_payment.failed":
		result := &V2MoneyManagementOutboundPaymentFailedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundPayment, error) {
			v := &V2MoneyManagementOutboundPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_payment.posted":
		result := &V2MoneyManagementOutboundPaymentPostedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundPayment, error) {
			v := &V2MoneyManagementOutboundPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_payment.returned":
		result := &V2MoneyManagementOutboundPaymentReturnedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundPayment, error) {
			v := &V2MoneyManagementOutboundPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_payment.updated":
		result := &V2MoneyManagementOutboundPaymentUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundPayment, error) {
			v := &V2MoneyManagementOutboundPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_transfer.canceled":
		result := &V2MoneyManagementOutboundTransferCanceledEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundTransfer, error) {
			v := &V2MoneyManagementOutboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_transfer.created":
		result := &V2MoneyManagementOutboundTransferCreatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundTransfer, error) {
			v := &V2MoneyManagementOutboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_transfer.failed":
		result := &V2MoneyManagementOutboundTransferFailedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundTransfer, error) {
			v := &V2MoneyManagementOutboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_transfer.posted":
		result := &V2MoneyManagementOutboundTransferPostedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundTransfer, error) {
			v := &V2MoneyManagementOutboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_transfer.returned":
		result := &V2MoneyManagementOutboundTransferReturnedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundTransfer, error) {
			v := &V2MoneyManagementOutboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.outbound_transfer.updated":
		result := &V2MoneyManagementOutboundTransferUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementOutboundTransfer, error) {
			v := &V2MoneyManagementOutboundTransfer{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.payout_method.created":
		result := &V2MoneyManagementPayoutMethodCreatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementPayoutMethod, error) {
			v := &V2MoneyManagementPayoutMethod{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.payout_method.updated":
		result := &V2MoneyManagementPayoutMethodUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementPayoutMethod, error) {
			v := &V2MoneyManagementPayoutMethod{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.received_credit.available":
		result := &V2MoneyManagementReceivedCreditAvailableEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementReceivedCredit, error) {
			v := &V2MoneyManagementReceivedCredit{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.money_management.received_credit.failed":
		result := &V2MoneyManagementReceivedCreditFailedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementReceivedCredit, error) {
			v := &V2MoneyManagementReceivedCredit{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.received_credit.returned":
		result := &V2MoneyManagementReceivedCreditReturnedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementReceivedCredit, error) {
			v := &V2MoneyManagementReceivedCredit{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.received_credit.succeeded":
		result := &V2MoneyManagementReceivedCreditSucceededEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementReceivedCredit, error) {
			v := &V2MoneyManagementReceivedCredit{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.received_debit.canceled":
		result := &V2MoneyManagementReceivedDebitCanceledEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementReceivedDebit, error) {
			v := &V2MoneyManagementReceivedDebit{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.received_debit.failed":
		result := &V2MoneyManagementReceivedDebitFailedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementReceivedDebit, error) {
			v := &V2MoneyManagementReceivedDebit{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.received_debit.pending":
		result := &V2MoneyManagementReceivedDebitPendingEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementReceivedDebit, error) {
			v := &V2MoneyManagementReceivedDebit{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.received_debit.succeeded":
		result := &V2MoneyManagementReceivedDebitSucceededEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementReceivedDebit, error) {
			v := &V2MoneyManagementReceivedDebit{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.received_debit.updated":
		result := &V2MoneyManagementReceivedDebitUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementReceivedDebit, error) {
			v := &V2MoneyManagementReceivedDebit{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.money_management.transaction.created":
		result := &V2MoneyManagementTransactionCreatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementTransaction, error) {
			v := &V2MoneyManagementTransaction{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.money_management.transaction.updated":
		result := &V2MoneyManagementTransactionUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2MoneyManagementTransaction, error) {
			v := &V2MoneyManagementTransaction{}
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
	case "v2.core.account[configuration.storer].capability_status_updated":
		evt := V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.account[configuration.storer].updated":
		evt := V2CoreAccountIncludingConfigurationStorerUpdatedEventNotification{}
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
	case "v2.core.health.event_generation_failure.resolved":
		evt := V2CoreHealthEventGenerationFailureResolvedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.adjustment.created":
		evt := V2MoneyManagementAdjustmentCreatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.financial_account.created":
		evt := V2MoneyManagementFinancialAccountCreatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.financial_account.updated":
		evt := V2MoneyManagementFinancialAccountUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.financial_address.activated":
		evt := V2MoneyManagementFinancialAddressActivatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.financial_address.failed":
		evt := V2MoneyManagementFinancialAddressFailedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.inbound_transfer.available":
		evt := V2MoneyManagementInboundTransferAvailableEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.inbound_transfer.bank_debit_failed":
		evt := V2MoneyManagementInboundTransferBankDebitFailedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.inbound_transfer.bank_debit_processing":
		evt := V2MoneyManagementInboundTransferBankDebitProcessingEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.inbound_transfer.bank_debit_queued":
		evt := V2MoneyManagementInboundTransferBankDebitQueuedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.inbound_transfer.bank_debit_returned":
		evt := V2MoneyManagementInboundTransferBankDebitReturnedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.inbound_transfer.bank_debit_succeeded":
		evt := V2MoneyManagementInboundTransferBankDebitSucceededEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_payment.canceled":
		evt := V2MoneyManagementOutboundPaymentCanceledEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_payment.created":
		evt := V2MoneyManagementOutboundPaymentCreatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_payment.failed":
		evt := V2MoneyManagementOutboundPaymentFailedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_payment.posted":
		evt := V2MoneyManagementOutboundPaymentPostedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_payment.returned":
		evt := V2MoneyManagementOutboundPaymentReturnedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_payment.updated":
		evt := V2MoneyManagementOutboundPaymentUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_transfer.canceled":
		evt := V2MoneyManagementOutboundTransferCanceledEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_transfer.created":
		evt := V2MoneyManagementOutboundTransferCreatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_transfer.failed":
		evt := V2MoneyManagementOutboundTransferFailedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_transfer.posted":
		evt := V2MoneyManagementOutboundTransferPostedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_transfer.returned":
		evt := V2MoneyManagementOutboundTransferReturnedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.outbound_transfer.updated":
		evt := V2MoneyManagementOutboundTransferUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.payout_method.created":
		evt := V2MoneyManagementPayoutMethodCreatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.payout_method.updated":
		evt := V2MoneyManagementPayoutMethodUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.received_credit.available":
		evt := V2MoneyManagementReceivedCreditAvailableEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.received_credit.failed":
		evt := V2MoneyManagementReceivedCreditFailedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.received_credit.returned":
		evt := V2MoneyManagementReceivedCreditReturnedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.received_credit.succeeded":
		evt := V2MoneyManagementReceivedCreditSucceededEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.received_debit.canceled":
		evt := V2MoneyManagementReceivedDebitCanceledEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.received_debit.failed":
		evt := V2MoneyManagementReceivedDebitFailedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.received_debit.pending":
		evt := V2MoneyManagementReceivedDebitPendingEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.received_debit.succeeded":
		evt := V2MoneyManagementReceivedDebitSucceededEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.received_debit.updated":
		evt := V2MoneyManagementReceivedDebitUpdatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.transaction.created":
		evt := V2MoneyManagementTransactionCreatedEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.money_management.transaction.updated":
		evt := V2MoneyManagementTransactionUpdatedEventNotification{}
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
