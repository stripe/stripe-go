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
type V2CoreAccountConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapability string

// List of values that V2CoreAccountConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapability can take
const (
	V2CoreAccountConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapabilityAutomaticIndirectTax V2CoreAccountConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapability = "automatic_indirect_tax"
)

// Open Enum. The capability which had its status updated.
type V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability string

// List of values that V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability can take
const (
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityACHDebitPayments         V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "ach_debit_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityACSSDebitPayments        V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "acss_debit_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityAffirmPayments           V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "affirm_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityAfterpayClearpayPayments V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "afterpay_clearpay_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityAlmaPayments             V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "alma_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityAmazonPayPayments        V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "amazon_pay_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityAUBECSDebitPayments      V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "au_becs_debit_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityBACSDebitPayments        V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "bacs_debit_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityBancontactPayments       V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "bancontact_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityBLIKPayments             V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "blik_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityBoletoPayments           V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "boleto_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityCardPayments             V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "card_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityCartesBancairesPayments  V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "cartes_bancaires_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityCashAppPayments          V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "cashapp_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityEPSPayments              V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "eps_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityFPXPayments              V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "fpx_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityGBBankTransferPayments   V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "gb_bank_transfer_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityGrabpayPayments          V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "grabpay_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityIDEALPayments            V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "ideal_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityJCBPayments              V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "jcb_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityJPBankTransferPayments   V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "jp_bank_transfer_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityKakaoPayPayments         V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "kakao_pay_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityKlarnaPayments           V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "klarna_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityKonbiniPayments          V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "konbini_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityKrCardPayments           V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "kr_card_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityLinkPayments             V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "link_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityMobilepayPayments        V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "mobilepay_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityMultibancoPayments       V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "multibanco_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityMXBankTransferPayments   V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "mx_bank_transfer_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityNaverPayPayments         V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "naver_pay_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityOXXOPayments             V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "oxxo_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityP24Payments              V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "p24_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityPaycoPayments            V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "payco_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityPayNowPayments           V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "paynow_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityPayByBankPayments        V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "pay_by_bank_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityPromptPayPayments        V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "promptpay_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityRevolutPayPayments       V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "revolut_pay_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilitySamsungPayPayments       V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "samsung_pay_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilitySEPABankTransferPayments V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "sepa_bank_transfer_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilitySEPADebitPayments        V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "sepa_debit_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilitySwishPayments            V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "swish_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityTWINTPayments            V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "twint_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityUSBankTransferPayments   V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "us_bank_transfer_payments"
	V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapabilityZipPayments              V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability = "zip_payments"
)

// Open Enum. The capability which had its status updated.
type V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability string

// List of values that V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability can take
const (
	V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityBankAccountsLocal            V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "bank_accounts.local"
	V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityBankAccountsLocalUk          V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "bank_accounts.local_uk"
	V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityBankAccountsWire             V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "bank_accounts.wire"
	V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityBankAccountsWireUk           V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "bank_accounts.wire_uk"
	V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityCards                        V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "cards"
	V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityCardsUk                      V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "cards_uk"
	V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityCryptoWalletsV2              V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "crypto_wallets_v2"
	V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityStripeBalanceStripeTransfers V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "stripe_balance.stripe_transfers"
	V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapabilityStripeTransfers              V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability = "stripe.transfers"
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
	GetBaseEvent() *V2BaseEvent
}

// V2RawEvent is the raw event type for V2 events. It is used to unmarshal the
// event data into a generic structure, and can also be used a default event
// type when the event type is not known.
type V2RawEvent struct {
	V2BaseEvent
	Data          *json.RawMessage `json:"data"`
	RelatedObject *RelatedObject   `json:"related_object"`
}

// V2CoreAccountRequirementsUpdatedEvent is the Go struct for the "v2.core.account[requirements].updated" event.
// This event occurs when the account's requirements are updated.
type V2CoreAccountRequirementsUpdatedEvent struct{ V2RawEvent }

// V2CoreAccountLinkCompletedEvent is the Go struct for the "v2.core.account_link.completed" event.
// The generated account link has been completed.
type V2CoreAccountLinkCompletedEvent struct {
	V2RawEvent
	Data               V2CoreAccountLinkCompletedEventData
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2CoreAccountLink, error)
}

// FetchRelatedObject fetches the related V2CoreAccountLink object for the event.
func (e V2CoreAccountLinkCompletedEvent) FetchRelatedObject() (*V2CoreAccountLink, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountConfigurationCustomerCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.customer].capability_status_updated" event.
// The status of a customer config capability was updated.
type V2CoreAccountConfigurationCustomerCapabilityStatusUpdatedEvent struct {
	V2RawEvent
	Data V2CoreAccountConfigurationCustomerCapabilityStatusUpdatedEventData
}

// V2CoreAccountConfigurationCustomerUpdatedEvent is the Go struct for the "v2.core.account[configuration.customer].updated" event.
// A customer config was updated.
type V2CoreAccountConfigurationCustomerUpdatedEvent struct{ V2RawEvent }

// V2CoreAccountIdentityUpdatedEvent is the Go struct for the "v2.core.account[identity].updated" event.
// This event occurs when identity is updated.
type V2CoreAccountIdentityUpdatedEvent struct{ V2RawEvent }

// V2CoreAccountPersonCreatedEvent is the Go struct for the "v2.core.account_person.created" event.
// This event occurs when a person is created.
type V2CoreAccountPersonCreatedEvent struct {
	V2RawEvent
	Data               V2CoreAccountPersonCreatedEventData
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2CorePerson, error)
}

// FetchRelatedObject fetches the related V2CorePerson object for the event.
func (e V2CoreAccountPersonCreatedEvent) FetchRelatedObject() (*V2CorePerson, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountPersonDeletedEvent is the Go struct for the "v2.core.account_person.deleted" event.
// This event occurs when a person is deleted.
type V2CoreAccountPersonDeletedEvent struct {
	V2RawEvent
	Data               V2CoreAccountPersonDeletedEventData
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2CorePerson, error)
}

// FetchRelatedObject fetches the related V2CorePerson object for the event.
func (e V2CoreAccountPersonDeletedEvent) FetchRelatedObject() (*V2CorePerson, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountPersonUpdatedEvent is the Go struct for the "v2.core.account_person.updated" event.
// This event occurs when a person is updated.
type V2CoreAccountPersonUpdatedEvent struct {
	V2RawEvent
	Data               V2CoreAccountPersonUpdatedEventData
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2CorePerson, error)
}

// FetchRelatedObject fetches the related V2CorePerson object for the event.
func (e V2CoreAccountPersonUpdatedEvent) FetchRelatedObject() (*V2CorePerson, error) {
	return e.fetchRelatedObject()
}

// V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.merchant].capability_status_updated" event.
// The status of a merchant config capability was updated.
type V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEvent struct {
	V2RawEvent
	Data V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventData
}

// V2CoreAccountConfigurationMerchantUpdatedEvent is the Go struct for the "v2.core.account[configuration.merchant].updated" event.
// A merchant config was updated.
type V2CoreAccountConfigurationMerchantUpdatedEvent struct{ V2RawEvent }

// V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEvent is the Go struct for the "v2.core.account[configuration.recipient].capability_status_updated" event.
// The status of a recipient config capability was updated.
type V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEvent struct {
	V2RawEvent
	Data V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventData
}

// V2CoreAccountConfigurationRecipientUpdatedEvent is the Go struct for the "v2.core.account[configuration.recipient].updated" event.
// A recipient config was updated.
type V2CoreAccountConfigurationRecipientUpdatedEvent struct{ V2RawEvent }

// V1BillingMeterErrorReportTriggeredEvent is the Go struct for the "v1.billing.meter.error_report_triggered" event.
// This event occurs when there are invalid async usage events for a given meter.
type V1BillingMeterErrorReportTriggeredEvent struct {
	V2RawEvent
	Data               V1BillingMeterErrorReportTriggeredEventData
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*BillingMeter, error)
}

// FetchRelatedObject fetches the related BillingMeter object for the event.
func (e V1BillingMeterErrorReportTriggeredEvent) FetchRelatedObject() (*BillingMeter, error) {
	return e.fetchRelatedObject()
}

// V1BillingMeterNoMeterFoundEvent is the Go struct for the "v1.billing.meter.no_meter_found" event.
// This event occurs when async usage events have missing or invalid meter ids.
type V1BillingMeterNoMeterFoundEvent struct {
	V2RawEvent
	Data V1BillingMeterNoMeterFoundEventData
}

// V2MoneyManagementFinancialAccountCreatedEvent is the Go struct for the "v2.money_management.financial_account.created" event.
// Occurs when a financial account is created.
type V2MoneyManagementFinancialAccountCreatedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementFinancialAccount, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementFinancialAccount object for the event.
func (e V2MoneyManagementFinancialAccountCreatedEvent) FetchRelatedObject() (*V2MoneyManagementFinancialAccount, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementFinancialAddressActivatedEvent is the Go struct for the "v2.money_management.financial_address.activated" event.
// The FinancialAddress is now active and ready to receive funds using the credentials provided.
type V2MoneyManagementFinancialAddressActivatedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementFinancialAddress, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementFinancialAddress object for the event.
func (e V2MoneyManagementFinancialAddressActivatedEvent) FetchRelatedObject() (*V2MoneyManagementFinancialAddress, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementFinancialAddressFailedEvent is the Go struct for the "v2.money_management.financial_address.failed" event.
// The FinancialAddress could not be activated and can not receive funds.
type V2MoneyManagementFinancialAddressFailedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementFinancialAddress, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementFinancialAddress object for the event.
func (e V2MoneyManagementFinancialAddressFailedEvent) FetchRelatedObject() (*V2MoneyManagementFinancialAddress, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferAvailableEvent is the Go struct for the "v2.money_management.inbound_transfer.available" event.
// Funds from an InboundTransfer were just made available.
type V2MoneyManagementInboundTransferAvailableEvent struct {
	V2RawEvent
	Data               V2MoneyManagementInboundTransferAvailableEventData
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferAvailableEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitFailedEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_failed" event.
// An InboundTransfer failed.
type V2MoneyManagementInboundTransferBankDebitFailedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferBankDebitFailedEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitProcessingEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_processing" event.
// An InboundTransfer is now processing.
type V2MoneyManagementInboundTransferBankDebitProcessingEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferBankDebitProcessingEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitQueuedEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_queued" event.
// An InboundTransfer has been queued.
type V2MoneyManagementInboundTransferBankDebitQueuedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferBankDebitQueuedEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitReturnedEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_returned" event.
// An InboundTransfer was returned.
type V2MoneyManagementInboundTransferBankDebitReturnedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferBankDebitReturnedEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementInboundTransferBankDebitSucceededEvent is the Go struct for the "v2.money_management.inbound_transfer.bank_debit_succeeded" event.
// An InboundTransfer succeeded.
type V2MoneyManagementInboundTransferBankDebitSucceededEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementInboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementInboundTransfer object for the event.
func (e V2MoneyManagementInboundTransferBankDebitSucceededEvent) FetchRelatedObject() (*V2MoneyManagementInboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentCanceledEvent is the Go struct for the "v2.money_management.outbound_payment.canceled" event.
// An OutboundPayment has transitioned into the canceled state.
type V2MoneyManagementOutboundPaymentCanceledEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentCanceledEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentCreatedEvent is the Go struct for the "v2.money_management.outbound_payment.created" event.
// A new OutboundPayment has been created.
type V2MoneyManagementOutboundPaymentCreatedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentCreatedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentFailedEvent is the Go struct for the "v2.money_management.outbound_payment.failed" event.
// An OutboundPayment has transitioned into the failed state.
type V2MoneyManagementOutboundPaymentFailedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentFailedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentPostedEvent is the Go struct for the "v2.money_management.outbound_payment.posted" event.
// An OutboundPayment has transitioned into the posted state.
type V2MoneyManagementOutboundPaymentPostedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentPostedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundPaymentReturnedEvent is the Go struct for the "v2.money_management.outbound_payment.returned" event.
// An OutboundPayment has transitioned into the returned state.
type V2MoneyManagementOutboundPaymentReturnedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementOutboundPayment, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundPayment object for the event.
func (e V2MoneyManagementOutboundPaymentReturnedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundPayment, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferCanceledEvent is the Go struct for the "v2.money_management.outbound_transfer.canceled" event.
// An OutboundTransfer has transitioned into the canceled state.
type V2MoneyManagementOutboundTransferCanceledEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferCanceledEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferCreatedEvent is the Go struct for the "v2.money_management.outbound_transfer.created" event.
// A new OutboundTransfer has been created.
type V2MoneyManagementOutboundTransferCreatedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferCreatedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferFailedEvent is the Go struct for the "v2.money_management.outbound_transfer.failed" event.
// An OutboundTransfer has transitioned into the failed state.
type V2MoneyManagementOutboundTransferFailedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferFailedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferPostedEvent is the Go struct for the "v2.money_management.outbound_transfer.posted" event.
// An OutboundTransfer has transitioned into the posted state.
type V2MoneyManagementOutboundTransferPostedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferPostedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementOutboundTransferReturnedEvent is the Go struct for the "v2.money_management.outbound_transfer.returned" event.
// An OutboundTransfer has transitioned into the returned state.
type V2MoneyManagementOutboundTransferReturnedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementOutboundTransfer, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementOutboundTransfer object for the event.
func (e V2MoneyManagementOutboundTransferReturnedEvent) FetchRelatedObject() (*V2MoneyManagementOutboundTransfer, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditAvailableEvent is the Go struct for the "v2.money_management.received_credit.available" event.
// The funds related to the received credit are available in your balance.
type V2MoneyManagementReceivedCreditAvailableEvent struct {
	V2RawEvent
	Data               V2MoneyManagementReceivedCreditAvailableEventData
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedCredit object for the event.
func (e V2MoneyManagementReceivedCreditAvailableEvent) FetchRelatedObject() (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditFailedEvent is the Go struct for the "v2.money_management.received_credit.failed" event.
// A credit was attempted to your balance and was not successful. See the status_details for more information.
type V2MoneyManagementReceivedCreditFailedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedCredit object for the event.
func (e V2MoneyManagementReceivedCreditFailedEvent) FetchRelatedObject() (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditReturnedEvent is the Go struct for the "v2.money_management.received_credit.returned" event.
// The previously received credit has been reversed, returned to the originator and deducted from your balance.
type V2MoneyManagementReceivedCreditReturnedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedCredit object for the event.
func (e V2MoneyManagementReceivedCreditReturnedEvent) FetchRelatedObject() (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedCreditSucceededEvent is the Go struct for the "v2.money_management.received_credit.succeeded" event.
// A credit was received successfully and processed by Stripe.
type V2MoneyManagementReceivedCreditSucceededEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementReceivedCredit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedCredit object for the event.
func (e V2MoneyManagementReceivedCreditSucceededEvent) FetchRelatedObject() (*V2MoneyManagementReceivedCredit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitCanceledEvent is the Go struct for the "v2.money_management.received_debit.canceled" event.
// This event is sent when a received debit is canceled.
type V2MoneyManagementReceivedDebitCanceledEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedDebit object for the event.
func (e V2MoneyManagementReceivedDebitCanceledEvent) FetchRelatedObject() (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitFailedEvent is the Go struct for the "v2.money_management.received_debit.failed" event.
// This event is sent when a received debit fails.
type V2MoneyManagementReceivedDebitFailedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedDebit object for the event.
func (e V2MoneyManagementReceivedDebitFailedEvent) FetchRelatedObject() (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitPendingEvent is the Go struct for the "v2.money_management.received_debit.pending" event.
// This event is sent when a ReceivedDebit is set to pending.
type V2MoneyManagementReceivedDebitPendingEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedDebit object for the event.
func (e V2MoneyManagementReceivedDebitPendingEvent) FetchRelatedObject() (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitSucceededEvent is the Go struct for the "v2.money_management.received_debit.succeeded" event.
// This event is sent when a ReceivedDebit succeeds.
type V2MoneyManagementReceivedDebitSucceededEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedDebit object for the event.
func (e V2MoneyManagementReceivedDebitSucceededEvent) FetchRelatedObject() (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// V2MoneyManagementReceivedDebitUpdatedEvent is the Go struct for the "v2.money_management.received_debit.updated" event.
// This event is sent when a ReceivedDebit is updated.
type V2MoneyManagementReceivedDebitUpdatedEvent struct {
	V2RawEvent
	RelatedObject      RelatedObject
	fetchRelatedObject func() (*V2MoneyManagementReceivedDebit, error)
}

// FetchRelatedObject fetches the related V2MoneyManagementReceivedDebit object for the event.
func (e V2MoneyManagementReceivedDebitUpdatedEvent) FetchRelatedObject() (*V2MoneyManagementReceivedDebit, error) {
	return e.fetchRelatedObject()
}

// The generated account link has been completed.
type V2CoreAccountLinkCompletedEventData struct {
	// The ID of the v2 account.
	AccountID string `json:"account_id"`
	// Configurations on the Account that was onboarded via the account link.
	Configurations []V2CoreAccountLinkCompletedEventDataConfiguration `json:"configurations"`
	// Open Enum. The use case type of the account link that has been completed.
	UseCase V2CoreAccountLinkCompletedEventDataUseCase `json:"use_case"`
}

// The status of a customer config capability was updated.
type V2CoreAccountConfigurationCustomerCapabilityStatusUpdatedEventData struct {
	// Open Enum. The capability which had its status updated.
	UpdatedCapability V2CoreAccountConfigurationCustomerCapabilityStatusUpdatedEventDataUpdatedCapability `json:"updated_capability"`
}

// This event occurs when a person is created.
type V2CoreAccountPersonCreatedEventData struct {
	// The ID of the v2 account.
	AccountID string `json:"account_id"`
}

// This event occurs when a person is deleted.
type V2CoreAccountPersonDeletedEventData struct {
	// The ID of the v2 account.
	AccountID string `json:"account_id"`
}

// This event occurs when a person is updated.
type V2CoreAccountPersonUpdatedEventData struct {
	// The ID of the v2 account.
	AccountID string `json:"account_id"`
}

// The status of a merchant config capability was updated.
type V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventData struct {
	// Open Enum. The capability which had its status updated.
	UpdatedCapability V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEventDataUpdatedCapability `json:"updated_capability"`
}

// The status of a recipient config capability was updated.
type V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventData struct {
	// Open Enum. The capability which had its status updated.
	UpdatedCapability V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEventDataUpdatedCapability `json:"updated_capability"`
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

// This event occurs when there are invalid async usage events for a given meter.
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

// This event occurs when async usage events have missing or invalid meter ids.
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

// Funds from an InboundTransfer were just made available.
type V2MoneyManagementInboundTransferAvailableEventData struct {
	// The transaction ID of the received credit.
	TransactionID string `json:"transaction_id"`
}

// The funds related to the received credit are available in your balance.
type V2MoneyManagementReceivedCreditAvailableEventData struct {
	// The transaction ID of the received credit.
	TransactionID string `json:"transaction_id"`
}

// ConvertRawEvent converts a raw event to a concrete event type.
// If the event type is not known, it returns the raw event.
func ConvertRawEvent(event *V2RawEvent, backend Backend, key string) (V2Event, error) {
	switch event.Type {
	case "v2.core.account[requirements].updated":
		result := &V2CoreAccountRequirementsUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account_link.completed":
		result := &V2CoreAccountLinkCompletedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2CoreAccountLink, error) {
			v := &V2CoreAccountLink{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account[configuration.customer].capability_status_updated":
		result := &V2CoreAccountConfigurationCustomerCapabilityStatusUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account[configuration.customer].updated":
		result := &V2CoreAccountConfigurationCustomerUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account[identity].updated":
		result := &V2CoreAccountIdentityUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account[configuration.merchant].capability_status_updated":
		result := &V2CoreAccountConfigurationMerchantCapabilityStatusUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account[configuration.merchant].updated":
		result := &V2CoreAccountConfigurationMerchantUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account[configuration.recipient].capability_status_updated":
		result := &V2CoreAccountConfigurationRecipientCapabilityStatusUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.account[configuration.recipient].updated":
		result := &V2CoreAccountConfigurationRecipientUpdatedEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
		}
		return result, nil
	case "v1.billing.meter.no_meter_found":
		result := &V1BillingMeterNoMeterFoundEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, result); err != nil {
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
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
		if err := json.Unmarshal(*event.Data, result); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return event, nil
	}
}
