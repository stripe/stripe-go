//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"net/http"
	"time"
)

// Configurations on the Account that was onboarded via the account link.
type V2CoreAccountLinkCompletedEventDataConfiguration string

// List of values that V2CoreAccountLinkCompletedEventDataConfiguration can take
const (
	V2CoreAccountLinkCompletedEventDataConfigurationRecipient V2CoreAccountLinkCompletedEventDataConfiguration = "recipient"
)

// Open Enum. The use case type of the account link that has been completed.
type V2CoreAccountLinkCompletedEventDataUseCase string

// List of values that V2CoreAccountLinkCompletedEventDataUseCase can take
const (
	V2CoreAccountLinkCompletedEventDataUseCaseAccountOnboarding V2CoreAccountLinkCompletedEventDataUseCase = "account_onboarding"
	V2CoreAccountLinkCompletedEventDataUseCaseAccountUpdate     V2CoreAccountLinkCompletedEventDataUseCase = "account_update"
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
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityHoldsCurrenciesEUR                 V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "holds_currencies.eur"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityHoldsCurrenciesGBP                 V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "holds_currencies.gbp"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityHoldsCurrenciesUSD                 V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "holds_currencies.usd"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityInboundTransfersBankAccounts       V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "inbound_transfers.bank_accounts"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundPaymentsBankAccounts       V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_payments.bank_accounts"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundPaymentsCards              V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_payments.cards"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundPaymentsFinancialAccounts  V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_payments.financial_accounts"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundTransfersBankAccounts      V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_transfers.bank_accounts"
	V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapabilityOutboundTransfersFinancialAccounts V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventDataUpdatedCapability = "outbound_transfers.financial_accounts"
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

// V2Event is the interface implemented by V2 Events. To get the underlying Event,
// use a type switch or type assertion to one of the concrete event types.
type V2Event interface {
	getBaseEvent() *V2BaseEvent
}

// V2RawEvent is the raw event type for V2 events. It is used to unmarshal the
// event data into a generic structure, and can also be used a default event
// type when the event type is not known.
type V2RawEvent struct {
	V2BaseEvent
	Data          *json.RawMessage `json:"data"`
	RelatedObject *RelatedObject   `json:"related_object"`
}

// V2CoreAccountIncludingRequirementsUpdatedEvent is the Go struct for the "v2.core.account[requirements].updated" event.
// Occurs when an Account's requirements are updated.
type V2CoreAccountIncludingRequirementsUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingRequirementsUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountLinkCompletedEvent is the Go struct for the "v2.core.account_link.completed" event.
// Occurs when the generated AccountLink is completed.
type V2CoreAccountLinkCompletedEvent struct {
	V2BaseEvent
	Data V2CoreAccountLinkCompletedEventData `json:"data"`
}

// V2CoreAccountClosedEvent is the Go struct for the "v2.core.account.closed" event.
// This event occurs when an account is closed.
type V2CoreAccountClosedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountClosedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountCreatedEvent is the Go struct for the "v2.core.account.created" event.
// Occurs when an Account is created.
type V2CoreAccountCreatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountCreatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountUpdatedEvent is the Go struct for the "v2.core.account.updated" event.
// Occurs when an Account is updated.
type V2CoreAccountUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingDefaultsUpdatedEvent is the Go struct for the "v2.core.account[defaults].updated" event.
// This event occurs when account defaults are created or updated.
type V2CoreAccountIncludingDefaultsUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingDefaultsUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.customer].capability_status_updated" event.
// Occurs when the status of an Account's customer configuration capability is updated.
type V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventData `json:"data"`
	RelatedObject      RelatedObject                                                               `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationCustomerUpdatedEvent is the Go struct for the "v2.core.account[configuration.customer].updated" event.
// Occurs when an Account's customer configuration is updated.
type V2CoreAccountIncludingConfigurationCustomerUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingConfigurationCustomerUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingIdentityUpdatedEvent is the Go struct for the "v2.core.account[identity].updated" event.
// Occurs when an Identity is updated.
type V2CoreAccountIncludingIdentityUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingIdentityUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountPersonCreatedEvent is the Go struct for the "v2.core.account_person.created" event.
// Occurs when a Person is created.
type V2CoreAccountPersonCreatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountPersonCreatedEventData `json:"data"`
	RelatedObject      RelatedObject                       `json:"related_object"`
	fetchRelatedObject func() (*V2CorePerson, error)
}

// FetchRelatedObject fetches the related V2CorePerson object for the event.
func (e V2CoreAccountPersonCreatedEvent) FetchRelatedObject() (*V2CorePerson, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountPersonDeletedEvent is the Go struct for the "v2.core.account_person.deleted" event.
// Occurs when a Person is deleted.
type V2CoreAccountPersonDeletedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountPersonDeletedEventData `json:"data"`
	RelatedObject      RelatedObject                       `json:"related_object"`
	fetchRelatedObject func() (*V2CorePerson, error)
}

// FetchRelatedObject fetches the related V2CorePerson object for the event.
func (e V2CoreAccountPersonDeletedEvent) FetchRelatedObject() (*V2CorePerson, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountPersonUpdatedEvent is the Go struct for the "v2.core.account_person.updated" event.
// Occurs when a Person is updated.
type V2CoreAccountPersonUpdatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountPersonUpdatedEventData `json:"data"`
	RelatedObject      RelatedObject                       `json:"related_object"`
	fetchRelatedObject func() (*V2CorePerson, error)
}

// FetchRelatedObject fetches the related V2CorePerson object for the event.
func (e V2CoreAccountPersonUpdatedEvent) FetchRelatedObject() (*V2CorePerson, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.merchant].capability_status_updated" event.
// Occurs when the status of an Account's merchant configuration capability is updated.
type V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEventData `json:"data"`
	RelatedObject      RelatedObject                                                               `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationMerchantUpdatedEvent is the Go struct for the "v2.core.account[configuration.merchant].updated" event.
// Occurs when an Account's merchant configuration is updated.
type V2CoreAccountIncludingConfigurationMerchantUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingConfigurationMerchantUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.recipient].capability_status_updated" event.
// Occurs when the status of an Account's recipient configuration capability is updated.
type V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEventData `json:"data"`
	RelatedObject      RelatedObject                                                                `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationRecipientUpdatedEvent is the Go struct for the "v2.core.account[configuration.recipient].updated" event.
// Occurs when a Recipient's configuration is updated.
type V2CoreAccountIncludingConfigurationRecipientUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingConfigurationRecipientUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.storer].capability_status_updated" event.
// Occurs when the status of an Account's storer configuration capability is updated.
type V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent struct {
	V2BaseEvent
	Data               V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEventData `json:"data"`
	RelatedObject      RelatedObject                                                             `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountIncludingConfigurationStorerUpdatedEvent is the Go struct for the "v2.core.account[configuration.storer].updated" event.
// Occurs when a Storer's configuration is updated.
type V2CoreAccountIncludingConfigurationStorerUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2CoreAccount, error)
}

// FetchRelatedObject fetches the related V2CoreAccount object for the event.
func (e V2CoreAccountIncludingConfigurationStorerUpdatedEvent) FetchRelatedObject() (*V2CoreAccount, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementAdjustmentCreatedEvent is the Go struct for the "v2.money_management.adjustment.created" event.
// Occurs when an Adjustment is created.
type V2MoneyManagementAdjustmentCreatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementAdjustment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementAdjustment object for the event.
func (e V2MoneyManagementAdjustmentCreatedEvent) FetchRelatedObject() (*V2MoneyManagementAdjustment, error) {
	return e.fetchRelatedObject()
}

// V1BillingMeterErrorReportTriggeredEvent is the Go struct for the "v1.billing.meter.error_report_triggered" event.
// Occurs when a Meter has invalid async usage events.
type V1BillingMeterErrorReportTriggeredEvent struct {
	V2BaseEvent
	Data               V1BillingMeterErrorReportTriggeredEventData `json:"data"`
	RelatedObject      RelatedObject                               `json:"related_object"`
	fetchRelatedObject func() (*BillingMeter, error)
}

// FetchRelatedObject fetches the related BillingMeter object for the event.
func (e V1BillingMeterErrorReportTriggeredEvent) FetchRelatedObject() (*BillingMeter, error) {
	return e.fetchRelatedObject()
}

// V1BillingMeterNoMeterFoundEvent is the Go struct for the "v1.billing.meter.no_meter_found" event.
// Occurs when a Meter's id is missing or invalid in async usage events.
type V1BillingMeterNoMeterFoundEvent struct {
	V2BaseEvent
	Data V1BillingMeterNoMeterFoundEventData `json:"data"`
}

// V2MoneyManagementFinancialAccountCreatedEvent is the Go struct for the "v2.money_management.financial_account.created" event.
// Occurs when a FinancialAccount is created.
type V2MoneyManagementFinancialAccountCreatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementFinancialAccount, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementFinancialAccount object for the event.
func (e V2MoneyManagementFinancialAccountCreatedEvent) FetchRelatedObject() (*V2MoneyManagementFinancialAccount, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementFinancialAccountUpdatedEvent is the Go struct for the "v2.money_management.financial_account.updated" event.
// Occurs when a FinancialAccount is updated.
type V2MoneyManagementFinancialAccountUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementFinancialAccount, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementFinancialAccount object for the event.
func (e V2MoneyManagementFinancialAccountUpdatedEvent) FetchRelatedObject() (*V2MoneyManagementFinancialAccount, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementFinancialAddressActivatedEvent is the Go struct for the "v2.money_management.financial_address.activated" event.
// Occurs when a FinancialAddress is activated and is ready to receive funds.
type V2MoneyManagementFinancialAddressActivatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementFinancialAddress, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementFinancialAddress object for the event.
func (e V2MoneyManagementFinancialAddressActivatedEvent) FetchRelatedObject() (*V2MoneyManagementFinancialAddress, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementFinancialAddressFailedEvent is the Go struct for the "v2.money_management.financial_address.failed" event.
// Occurs when a FinancialAddress fails to activate and can not receive funds.
type V2MoneyManagementFinancialAddressFailedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementFinancialAddress, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementFinancialAddress object for the event.
func (e V2MoneyManagementFinancialAddressFailedEvent) FetchRelatedObject() (*V2MoneyManagementFinancialAddress, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferAvailableEvent is the Go struct for the "v2.money_management.inbound_transfer.available" event.
// Occurs when an InboundTransfer's funds are made available.
type V2MoneyManagementInboundTransferAvailableEvent struct {
	V2BaseEvent
	Data               V2MoneyManagementInboundTransferAvailableEventData `json:"data"`
	RelatedObject      RelatedObject                                      `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferAvailableEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitFailedEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_failed" event.
// Occurs when an InboundTransfer fails.
type V2MoneyManagementInboundTransferBankDebitFailedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferBankDebitFailedEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitProcessingEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_processing" event.
// Occurs when an InboundTransfer starts processing.
type V2MoneyManagementInboundTransferBankDebitProcessingEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferBankDebitProcessingEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitQueuedEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_queued" event.
// Occurs when an InboundTransfer is queued.
type V2MoneyManagementInboundTransferBankDebitQueuedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferBankDebitQueuedEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitReturnedEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_returned" event.
// Occurs when an InboundTransfer is returned.
type V2MoneyManagementInboundTransferBankDebitReturnedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferBankDebitReturnedEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitSucceededEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_succeeded" event.
// Occurs when an InboundTransfer succeeds.
type V2MoneyManagementInboundTransferBankDebitSucceededEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferBankDebitSucceededEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2CoreEventDestinationPingEvent is the Go struct for the "v2.core.event_destination.ping" event.
// A ping event used to test the connection to an EventDestination.
type V2CoreEventDestinationPingEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2EventDestination, error)
}

// FetchRelatedObject fetches the related V2EventDestination object for the event.
func (e V2CoreEventDestinationPingEvent) FetchRelatedObject() (*V2EventDestination, error) {
	return e.fetchRelatedObject()
}

// V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEvent is the Go struct for the "v2.payments.off_session_payment.authorization_attempt_failed" event.
// Sent after a failed authorization if there are still retries available on the OffSessionPayment.
type V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2PaymentsOffSessionPayment, error)
}

// FetchRelatedObject fetches the related V2PaymentsOffSessionPayment object for the event.
func (e V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEvent) FetchRelatedObject() (*V2PaymentsOffSessionPayment, error) {
	return e.fetchRelatedObject()
}

// V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEvent is the Go struct for the "v2.payments.off_session_payment.authorization_attempt_started" event.
// Sent when our internal scheduling system kicks off an attempt at authorization, whether it's a
// retry or an initial authorization.
type V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2PaymentsOffSessionPayment, error)
}

// FetchRelatedObject fetches the related V2PaymentsOffSessionPayment object for the event.
func (e V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEvent) FetchRelatedObject() (*V2PaymentsOffSessionPayment, error) {
	return e.fetchRelatedObject()
}

// V2PaymentsOffSessionPaymentCanceledEvent is the Go struct for the "v2.payments.off_session_payment.canceled" event.
// Sent immediately following a user's call to the Off-Session Payments cancel endpoint.
type V2PaymentsOffSessionPaymentCanceledEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2PaymentsOffSessionPayment, error)
}

// FetchRelatedObject fetches the related V2PaymentsOffSessionPayment object for the event.
func (e V2PaymentsOffSessionPaymentCanceledEvent) FetchRelatedObject() (*V2PaymentsOffSessionPayment, error) {
	return e.fetchRelatedObject()
}

// V2PaymentsOffSessionPaymentCreatedEvent is the Go struct for the "v2.payments.off_session_payment.created" event.
// Sent immediately following a user's call to the Off-Session Payments create endpoint.
type V2PaymentsOffSessionPaymentCreatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2PaymentsOffSessionPayment, error)
}

// FetchRelatedObject fetches the related V2PaymentsOffSessionPayment object for the event.
func (e V2PaymentsOffSessionPaymentCreatedEvent) FetchRelatedObject() (*V2PaymentsOffSessionPayment, error) {
	return e.fetchRelatedObject()
}

// V2PaymentsOffSessionPaymentFailedEvent is the Go struct for the "v2.payments.off_session_payment.failed" event.
// Sent after a failed authorization if there are no retries remaining, or if the failure is unretryable.
type V2PaymentsOffSessionPaymentFailedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2PaymentsOffSessionPayment, error)
}

// FetchRelatedObject fetches the related V2PaymentsOffSessionPayment object for the event.
func (e V2PaymentsOffSessionPaymentFailedEvent) FetchRelatedObject() (*V2PaymentsOffSessionPayment, error) {
	return e.fetchRelatedObject()
}

// V2PaymentsOffSessionPaymentSucceededEvent is the Go struct for the "v2.payments.off_session_payment.succeeded" event.
// Sent immediately after a successful authorization.
type V2PaymentsOffSessionPaymentSucceededEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2PaymentsOffSessionPayment, error)
}

// FetchRelatedObject fetches the related V2PaymentsOffSessionPayment object for the event.
func (e V2PaymentsOffSessionPaymentSucceededEvent) FetchRelatedObject() (*V2PaymentsOffSessionPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentCanceledEvent is the Go struct for the "v2.money_management.outbound_payment.canceled" event.
// Occurs when an OutboundPayment transitions into the canceled state.
type V2MoneyManagementOutboundPaymentCanceledEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentCanceledEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentCreatedEvent is the Go struct for the "v2.money_management.outbound_payment.created" event.
// Occurs when an OutboundPayment is created.
type V2MoneyManagementOutboundPaymentCreatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentCreatedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentFailedEvent is the Go struct for the "v2.money_management.outbound_payment.failed" event.
// Occurs when an OutboundPayment transitions into the failed state.
type V2MoneyManagementOutboundPaymentFailedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentFailedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentPostedEvent is the Go struct for the "v2.money_management.outbound_payment.posted" event.
// Occurs when an OutboundPayment transitions into the posted state.
type V2MoneyManagementOutboundPaymentPostedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentPostedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentReturnedEvent is the Go struct for the "v2.money_management.outbound_payment.returned" event.
// Occurs when an OutboundPayment transitions into the returned state.
type V2MoneyManagementOutboundPaymentReturnedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentReturnedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentUpdatedEvent is the Go struct for the "v2.money_management.outbound_payment.updated" event.
// Occurs when an OutboundPayment is updated.
type V2MoneyManagementOutboundPaymentUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentUpdatedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferCanceledEvent is the Go struct for the "v2.money_management.outbound_transfer.canceled" event.
// Occurs when an OutboundTransfer transitions into the canceled state.
type V2MoneyManagementOutboundTransferCanceledEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferCanceledEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferCreatedEvent is the Go struct for the "v2.money_management.outbound_transfer.created" event.
// Occurs when an OutboundTransfer is created.
type V2MoneyManagementOutboundTransferCreatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferCreatedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferFailedEvent is the Go struct for the "v2.money_management.outbound_transfer.failed" event.
// Occurs when an OutboundTransfer transitions into the failed state.
type V2MoneyManagementOutboundTransferFailedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferFailedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferPostedEvent is the Go struct for the "v2.money_management.outbound_transfer.posted" event.
// Occurs when an OutboundTransfer transitions into the posted state.
type V2MoneyManagementOutboundTransferPostedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferPostedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferReturnedEvent is the Go struct for the "v2.money_management.outbound_transfer.returned" event.
// Occurs when an OutboundTransfer transitions into the returned state.
type V2MoneyManagementOutboundTransferReturnedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferReturnedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferUpdatedEvent is the Go struct for the "v2.money_management.outbound_transfer.updated" event.
// Event that is emitted every time an Outbound Transfer is updated.
type V2MoneyManagementOutboundTransferUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferUpdatedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementPayoutMethodUpdatedEvent is the Go struct for the "v2.money_management.payout_method.updated" event.
// Occurs when a PayoutMethod is updated.
type V2MoneyManagementPayoutMethodUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementPayoutMethod, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementPayoutMethod object for the event.
func (e V2MoneyManagementPayoutMethodUpdatedEvent) FetchRelatedObject() (*V2MoneyManagementPayoutMethod, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditAvailableEvent is the Go struct for the "v2.money_management.received_credit.available" event.
// Occurs when a ReceivedCredit's funds are received and are available in your balance.
type V2MoneyManagementReceivedCreditAvailableEvent struct {
	V2BaseEvent
	Data               V2MoneyManagementReceivedCreditAvailableEventData `json:"data"`
	RelatedObject      RelatedObject                                     `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedCredit object for the event.
func (e V2MoneyManagementReceivedCreditAvailableEvent) FetchRelatedObject() (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditFailedEvent is the Go struct for the "v2.money_management.received_credit.failed" event.
// Occurs when a ReceivedCredit is attempted to your balance and fails. See the status_details for more information.
type V2MoneyManagementReceivedCreditFailedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedCredit object for the event.
func (e V2MoneyManagementReceivedCreditFailedEvent) FetchRelatedObject() (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditReturnedEvent is the Go struct for the "v2.money_management.received_credit.returned" event.
// Occurs when a ReceivedCredit is reversed, returned to the originator, and deducted from your balance.
type V2MoneyManagementReceivedCreditReturnedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedCredit object for the event.
func (e V2MoneyManagementReceivedCreditReturnedEvent) FetchRelatedObject() (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditSucceededEvent is the Go struct for the "v2.money_management.received_credit.succeeded" event.
// Occurs when a ReceivedCredit succeeds.
type V2MoneyManagementReceivedCreditSucceededEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedCredit object for the event.
func (e V2MoneyManagementReceivedCreditSucceededEvent) FetchRelatedObject() (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitCanceledEvent is the Go struct for the "v2.money_management.received_debit.canceled" event.
// Occurs when a ReceivedDebit is canceled.
type V2MoneyManagementReceivedDebitCanceledEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedDebit object for the event.
func (e V2MoneyManagementReceivedDebitCanceledEvent) FetchRelatedObject() (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitFailedEvent is the Go struct for the "v2.money_management.received_debit.failed" event.
// Occurs when a ReceivedDebit fails.
type V2MoneyManagementReceivedDebitFailedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedDebit object for the event.
func (e V2MoneyManagementReceivedDebitFailedEvent) FetchRelatedObject() (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitPendingEvent is the Go struct for the "v2.money_management.received_debit.pending" event.
// Occurs when a ReceivedDebit is set to pending.
type V2MoneyManagementReceivedDebitPendingEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedDebit object for the event.
func (e V2MoneyManagementReceivedDebitPendingEvent) FetchRelatedObject() (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitSucceededEvent is the Go struct for the "v2.money_management.received_debit.succeeded" event.
// Occurs when a ReceivedDebit succeeds.
type V2MoneyManagementReceivedDebitSucceededEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedDebit object for the event.
func (e V2MoneyManagementReceivedDebitSucceededEvent) FetchRelatedObject() (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitUpdatedEvent is the Go struct for the "v2.money_management.received_debit.updated" event.
// Occurs when a ReceivedDebit is updated.
type V2MoneyManagementReceivedDebitUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedDebit object for the event.
func (e V2MoneyManagementReceivedDebitUpdatedEvent) FetchRelatedObject() (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementTransactionCreatedEvent is the Go struct for the "v2.money_management.transaction.created" event.
// Occurs when a Transaction is created.
type V2MoneyManagementTransactionCreatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementTransaction, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementTransaction object for the event.
func (e V2MoneyManagementTransactionCreatedEvent) FetchRelatedObject() (*V2MoneyManagementTransaction, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementTransactionUpdatedEvent is the Go struct for the "v2.money_management.transaction.updated" event.
// Occurs when a Transaction is updated.
type V2MoneyManagementTransactionUpdatedEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2MoneyManagementTransaction, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementTransaction object for the event.
func (e V2MoneyManagementTransactionUpdatedEvent) FetchRelatedObject() (*V2MoneyManagementTransaction, error) {
	return e.fetchRelatedObject()
}

// Occurs when the generated AccountLink is completed.
type V2CoreAccountLinkCompletedEventData struct {
	// The ID of the v2 account.
	AccountID string `json:"account_id"`
	// Configurations on the Account that was onboarded via the account link.
	Configurations []V2CoreAccountLinkCompletedEventDataConfiguration `json:"configurations"`
	// Open Enum. The use case type of the account link that has been completed.
	UseCase V2CoreAccountLinkCompletedEventDataUseCase `json:"use_case"`
}

// Occurs when the status of an Account's customer configuration capability is updated.
type V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventData struct {
	// Open Enum. The capability which had its status updated.
	UpdatedCapability V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapability `json:"updated_capability"`
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

// ConvertRawEvent converts a raw event to a concrete event type.
// If the event type is not known, it returns the raw event.
func ConvertRawEvent(event *V2RawEvent, backend Backend, key string) (V2Event, error) {
	switch event.Type {
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
	case "v2.core.account_link.completed":
		result := &V2CoreAccountLinkCompletedEvent{}
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
	case "v2.core.account_person.created":
		result := &V2CoreAccountPersonCreatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CorePerson, error) {
			v := &V2CorePerson{}
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
		result.fetchRelatedObject = func() (*V2CorePerson, error) {
			v := &V2CorePerson{}
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
		result.fetchRelatedObject = func() (*V2CorePerson, error) {
			v := &V2CorePerson{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
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
	case "v2.core.event_destination.ping":
		result := &V2CoreEventDestinationPingEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2EventDestination, error) {
			v := &V2EventDestination{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.payments.off_session_payment.authorization_attempt_failed":
		result := &V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2PaymentsOffSessionPayment, error) {
			v := &V2PaymentsOffSessionPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.payments.off_session_payment.authorization_attempt_started":
		result := &V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2PaymentsOffSessionPayment, error) {
			v := &V2PaymentsOffSessionPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.payments.off_session_payment.canceled":
		result := &V2PaymentsOffSessionPaymentCanceledEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2PaymentsOffSessionPayment, error) {
			v := &V2PaymentsOffSessionPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.payments.off_session_payment.created":
		result := &V2PaymentsOffSessionPaymentCreatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2PaymentsOffSessionPayment, error) {
			v := &V2PaymentsOffSessionPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.payments.off_session_payment.failed":
		result := &V2PaymentsOffSessionPaymentFailedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2PaymentsOffSessionPayment, error) {
			v := &V2PaymentsOffSessionPayment{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	case "v2.payments.off_session_payment.succeeded":
		result := &V2PaymentsOffSessionPaymentSucceededEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2PaymentsOffSessionPayment, error) {
			v := &V2PaymentsOffSessionPayment{}
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
