//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The HTTP method.
type V2CoreHealthAlertAPIErrorHTTPMethod string

// List of values that V2CoreHealthAlertAPIErrorHTTPMethod can take
const (
	V2CoreHealthAlertAPIErrorHTTPMethodDELETE V2CoreHealthAlertAPIErrorHTTPMethod = "DELETE"
	V2CoreHealthAlertAPIErrorHTTPMethodGET    V2CoreHealthAlertAPIErrorHTTPMethod = "GET"
	V2CoreHealthAlertAPIErrorHTTPMethodPOST   V2CoreHealthAlertAPIErrorHTTPMethod = "POST"
	V2CoreHealthAlertAPIErrorHTTPMethodPUT    V2CoreHealthAlertAPIErrorHTTPMethod = "PUT"
)

// The HTTP method.
type V2CoreHealthAlertAPILatencyHTTPMethod string

// List of values that V2CoreHealthAlertAPILatencyHTTPMethod can take
const (
	V2CoreHealthAlertAPILatencyHTTPMethodDELETE V2CoreHealthAlertAPILatencyHTTPMethod = "DELETE"
	V2CoreHealthAlertAPILatencyHTTPMethodGET    V2CoreHealthAlertAPILatencyHTTPMethod = "GET"
	V2CoreHealthAlertAPILatencyHTTPMethodPOST   V2CoreHealthAlertAPILatencyHTTPMethod = "POST"
	V2CoreHealthAlertAPILatencyHTTPMethodPUT    V2CoreHealthAlertAPILatencyHTTPMethod = "PUT"
)

// The type of the charge.
type V2CoreHealthAlertAuthorizationRateDropChargeType string

// List of values that V2CoreHealthAlertAuthorizationRateDropChargeType can take
const (
	V2CoreHealthAlertAuthorizationRateDropChargeTypeMoneyMoving V2CoreHealthAlertAuthorizationRateDropChargeType = "money_moving"
	V2CoreHealthAlertAuthorizationRateDropChargeTypeValidation  V2CoreHealthAlertAuthorizationRateDropChargeType = "validation"
)

// The type of the dimension. Determines which field in dimension_details is populated.
type V2CoreHealthAlertAuthorizationRateDropDimensionType string

// List of values that V2CoreHealthAlertAuthorizationRateDropDimensionType can take
const (
	V2CoreHealthAlertAuthorizationRateDropDimensionTypeIssuer V2CoreHealthAlertAuthorizationRateDropDimensionType = "issuer"
)

// The type of the payment method.
type V2CoreHealthAlertAuthorizationRateDropPaymentMethodType string

// List of values that V2CoreHealthAlertAuthorizationRateDropPaymentMethodType can take
const (
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeACSSDebit            V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "acss_debit"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeAffirm               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "affirm"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeAfterpayClearpay     V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "afterpay_clearpay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeAlipay               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "alipay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeAlma                 V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "alma"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeAmazonPay            V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "amazon_pay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeApplePay             V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "apple_pay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeAUBECSDebit          V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "au_becs_debit"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeBACSDebit            V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "bacs_debit"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeBancontact           V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "bancontact"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeBillie               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "billie"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeBLIK                 V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "blik"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeBoleto               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "boleto"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeCard                 V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "card"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeCardPresent          V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "card_present"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeCartesBancaires      V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "cartes_bancaires"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeCashApp              V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "cashapp"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeCrypto               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "crypto"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeDummyPassthroughCard V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "dummy_passthrough_card"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeEPS                  V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "eps"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeFPX                  V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "fpx"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeGiropay              V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "giropay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeGrabpay              V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "grabpay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeIDEAL                V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "ideal"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeInteracPresent       V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "interac_present"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeKakaoPay             V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "kakao_pay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeKlarna               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "klarna"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeKonbini              V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "konbini"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeKriya                V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "kriya"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeKrCard               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "kr_card"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeLink                 V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "link"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeMbWay                V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "mb_way"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeMobilepay            V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "mobilepay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeMondu                V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "mondu"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeMultibanco           V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "multibanco"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeNaverPay             V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "naver_pay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeNgBank               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "ng_bank"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeNgBankTransfer       V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "ng_bank_transfer"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeNgCard               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "ng_card"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeNgMarket             V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "ng_market"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeNgUssd               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "ng_ussd"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeNgWallet             V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "ng_wallet"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeOXXO                 V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "oxxo"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeP24                  V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "p24"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypePaperCheck           V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "paper_check"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypePayco                V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "payco"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypePayNow               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "paynow"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypePaypal               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "paypal"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypePaypay               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "paypay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypePayto                V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "payto"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypePayByBank            V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "pay_by_bank"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypePix                  V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "pix"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypePromptPay            V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "promptpay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeRechnung             V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "rechnung"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeRevolutPay           V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "revolut_pay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeSamsungPay           V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "samsung_pay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeSatispay             V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "satispay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeScalapay             V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "scalapay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeSEPADebit            V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "sepa_debit"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeSequra               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "sequra"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeSofort               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "sofort"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeSunbit               V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "sunbit"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeSwish                V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "swish"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeTWINT                V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "twint"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeUpi                  V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "upi"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeUSBankAccount        V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "us_bank_account"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeVipps                V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "vipps"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeWeChatPay            V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "wechat_pay"
	V2CoreHealthAlertAuthorizationRateDropPaymentMethodTypeZip                  V2CoreHealthAlertAuthorizationRateDropPaymentMethodType = "zip"
)

// The type of the element.
type V2CoreHealthAlertElementsErrorElementType string

// List of values that V2CoreHealthAlertElementsErrorElementType can take
const (
	V2CoreHealthAlertElementsErrorElementTypeExpressCheckout V2CoreHealthAlertElementsErrorElementType = "expressCheckout"
	V2CoreHealthAlertElementsErrorElementTypePayment         V2CoreHealthAlertElementsErrorElementType = "payment"
)

// Fraud attack type.
type V2CoreHealthAlertFraudRateAttackType string

// List of values that V2CoreHealthAlertFraudRateAttackType can take
const (
	V2CoreHealthAlertFraudRateAttackTypeSpike           V2CoreHealthAlertFraudRateAttackType = "spike"
	V2CoreHealthAlertFraudRateAttackTypeSustainedAttack V2CoreHealthAlertFraudRateAttackType = "sustained_attack"
)

// Whether the alert is linked to an incident or is a self-contained problem.
type V2CoreHealthAlertGroupingType string

// List of values that V2CoreHealthAlertGroupingType can take
const (
	V2CoreHealthAlertGroupingTypeIncident   V2CoreHealthAlertGroupingType = "incident"
	V2CoreHealthAlertGroupingTypeStandalone V2CoreHealthAlertGroupingType = "standalone"
)

// The ingestion method.
type V2CoreHealthAlertMeterEventSummariesDelayedIngestionMethod string

// List of values that V2CoreHealthAlertMeterEventSummariesDelayedIngestionMethod can take
const (
	V2CoreHealthAlertMeterEventSummariesDelayedIngestionMethodImportSets V2CoreHealthAlertMeterEventSummariesDelayedIngestionMethod = "import_sets"
)

// The impacted Metronome billing pipeline.
type V2CoreHealthAlertMetronomeNotificationLatencyPipeline string

// List of values that V2CoreHealthAlertMetronomeNotificationLatencyPipeline can take
const (
	V2CoreHealthAlertMetronomeNotificationLatencyPipelineConfigurationTriggered        V2CoreHealthAlertMetronomeNotificationLatencyPipeline = "configuration_triggered"
	V2CoreHealthAlertMetronomeNotificationLatencyPipelineHighCardinalityUsageTriggered V2CoreHealthAlertMetronomeNotificationLatencyPipeline = "high_cardinality_usage_triggered"
	V2CoreHealthAlertMetronomeNotificationLatencyPipelineStandardUsageTriggered        V2CoreHealthAlertMetronomeNotificationLatencyPipeline = "standard_usage_triggered"
	V2CoreHealthAlertMetronomeNotificationLatencyPipelineTimeTriggered                 V2CoreHealthAlertMetronomeNotificationLatencyPipeline = "time_triggered"
)

// The type of the payment method.
type V2CoreHealthAlertPaymentMethodErrorPaymentMethodType string

// List of values that V2CoreHealthAlertPaymentMethodErrorPaymentMethodType can take
const (
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeACSSDebit            V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "acss_debit"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeAffirm               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "affirm"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeAfterpayClearpay     V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "afterpay_clearpay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeAlipay               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "alipay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeAlma                 V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "alma"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeAmazonPay            V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "amazon_pay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeApplePay             V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "apple_pay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeAUBECSDebit          V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "au_becs_debit"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeBACSDebit            V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "bacs_debit"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeBancontact           V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "bancontact"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeBillie               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "billie"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeBLIK                 V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "blik"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeBoleto               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "boleto"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeCard                 V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "card"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeCardPresent          V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "card_present"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeCartesBancaires      V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "cartes_bancaires"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeCashApp              V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "cashapp"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeCrypto               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "crypto"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeDummyPassthroughCard V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "dummy_passthrough_card"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeEPS                  V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "eps"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeFPX                  V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "fpx"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeGiropay              V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "giropay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeGrabpay              V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "grabpay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeIDEAL                V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "ideal"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeInteracPresent       V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "interac_present"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeKakaoPay             V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "kakao_pay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeKlarna               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "klarna"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeKonbini              V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "konbini"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeKriya                V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "kriya"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeKrCard               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "kr_card"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeLink                 V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "link"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeMbWay                V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "mb_way"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeMobilepay            V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "mobilepay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeMondu                V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "mondu"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeMultibanco           V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "multibanco"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeNaverPay             V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "naver_pay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeNgBank               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "ng_bank"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeNgBankTransfer       V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "ng_bank_transfer"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeNgCard               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "ng_card"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeNgMarket             V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "ng_market"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeNgUssd               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "ng_ussd"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeNgWallet             V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "ng_wallet"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeOXXO                 V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "oxxo"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeP24                  V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "p24"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypePaperCheck           V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "paper_check"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypePayco                V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "payco"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypePayNow               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "paynow"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypePaypal               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "paypal"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypePaypay               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "paypay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypePayto                V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "payto"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypePayByBank            V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "pay_by_bank"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypePix                  V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "pix"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypePromptPay            V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "promptpay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeRechnung             V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "rechnung"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeRevolutPay           V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "revolut_pay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeSamsungPay           V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "samsung_pay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeSatispay             V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "satispay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeScalapay             V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "scalapay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeSEPADebit            V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "sepa_debit"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeSequra               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "sequra"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeSofort               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "sofort"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeSunbit               V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "sunbit"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeSwish                V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "swish"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeTWINT                V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "twint"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeUpi                  V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "upi"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeUSBankAccount        V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "us_bank_account"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeVipps                V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "vipps"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeWeChatPay            V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "wechat_pay"
	V2CoreHealthAlertPaymentMethodErrorPaymentMethodTypeZip                  V2CoreHealthAlertPaymentMethodErrorPaymentMethodType = "zip"
)

// The severity of the alert.
type V2CoreHealthAlertSeverity string

// List of values that V2CoreHealthAlertSeverity can take
const (
	V2CoreHealthAlertSeverityCritical    V2CoreHealthAlertSeverity = "critical"
	V2CoreHealthAlertSeverityNonCritical V2CoreHealthAlertSeverity = "non_critical"
)

// The current status of the alert.
type V2CoreHealthAlertStatus string

// List of values that V2CoreHealthAlertStatus can take
const (
	V2CoreHealthAlertStatusOpen     V2CoreHealthAlertStatus = "open"
	V2CoreHealthAlertStatusResolved V2CoreHealthAlertStatus = "resolved"
)

// The type of the alert. Determines which sub-hash field is populated.
type V2CoreHealthAlertType string

// List of values that V2CoreHealthAlertType can take
const (
	V2CoreHealthAlertTypeAPIError                           V2CoreHealthAlertType = "api_error"
	V2CoreHealthAlertTypeAPILatency                         V2CoreHealthAlertType = "api_latency"
	V2CoreHealthAlertTypeAuthorizationRateDrop              V2CoreHealthAlertType = "authorization_rate_drop"
	V2CoreHealthAlertTypeElementsError                      V2CoreHealthAlertType = "elements_error"
	V2CoreHealthAlertTypeEventGenerationFailure             V2CoreHealthAlertType = "event_generation_failure"
	V2CoreHealthAlertTypeFraudRate                          V2CoreHealthAlertType = "fraud_rate"
	V2CoreHealthAlertTypeInvoiceCountDropped                V2CoreHealthAlertType = "invoice_count_dropped"
	V2CoreHealthAlertTypeIssuingAuthorizationRequestErrors  V2CoreHealthAlertType = "issuing_authorization_request_errors"
	V2CoreHealthAlertTypeIssuingAuthorizationRequestTimeout V2CoreHealthAlertType = "issuing_authorization_request_timeout"
	V2CoreHealthAlertTypeMeterEventSummariesDelayed         V2CoreHealthAlertType = "meter_event_summaries_delayed"
	V2CoreHealthAlertTypeMetronomeNotificationLatency       V2CoreHealthAlertType = "metronome_notification_latency"
	V2CoreHealthAlertTypePaymentMethodError                 V2CoreHealthAlertType = "payment_method_error"
	V2CoreHealthAlertTypeSEPADebitDelayed                   V2CoreHealthAlertType = "sepa_debit_delayed"
	V2CoreHealthAlertTypeTrafficVolumeDrop                  V2CoreHealthAlertType = "traffic_volume_drop"
	V2CoreHealthAlertTypeWebhookLatency                     V2CoreHealthAlertType = "webhook_latency"
)

// The top impacted connected accounts (only for platforms).
type V2CoreHealthAlertAPIErrorTopImpactedAccount struct {
	// The account ID of the impacted connected account.
	Account string `json:"account"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
}

// Populated when type is api_error.
type V2CoreHealthAlertAPIError struct {
	// The canonical path.
	CanonicalPath string `json:"canonical_path"`
	// The error code.
	ErrorCode string `json:"error_code,omitempty"`
	// The HTTP method.
	HTTPMethod V2CoreHealthAlertAPIErrorHTTPMethod `json:"http_method"`
	// The HTTP status.
	HTTPStatus string `json:"http_status"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
	// The top impacted connected accounts (only for platforms).
	TopImpactedAccounts []*V2CoreHealthAlertAPIErrorTopImpactedAccount `json:"top_impacted_accounts,omitempty"`
}

// The top impacted connected accounts (only for platforms).
type V2CoreHealthAlertAPILatencyTopImpactedAccount struct {
	// The account ID of the impacted connected account.
	Account string `json:"account"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
}

// Populated when type is api_latency.
type V2CoreHealthAlertAPILatency struct {
	// The canonical path.
	CanonicalPath string `json:"canonical_path"`
	// The HTTP method.
	HTTPMethod V2CoreHealthAlertAPILatencyHTTPMethod `json:"http_method"`
	// The HTTP status.
	HTTPStatus string `json:"http_status"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
	// The top impacted connected accounts (only for platforms).
	TopImpactedAccounts []*V2CoreHealthAlertAPILatencyTopImpactedAccount `json:"top_impacted_accounts,omitempty"`
}

// Dimensions that describe what subset of payments are impacted.
type V2CoreHealthAlertAuthorizationRateDropDimension struct {
	// Populated when type is issuer.
	Issuer string `json:"issuer,omitempty"`
	// The type of the dimension. Determines which field in dimension_details is populated.
	Type V2CoreHealthAlertAuthorizationRateDropDimensionType `json:"type"`
}

// Populated when type is authorization_rate_drop.
type V2CoreHealthAlertAuthorizationRateDrop struct {
	// The type of the charge.
	ChargeType V2CoreHealthAlertAuthorizationRateDropChargeType `json:"charge_type"`
	// The current authorization rate percentage.
	CurrentPercentage float64 `json:"current_percentage,string"`
	// Dimensions that describe what subset of payments are impacted.
	Dimensions []*V2CoreHealthAlertAuthorizationRateDropDimension `json:"dimensions,omitempty"`
	// The type of the payment method.
	PaymentMethodType V2CoreHealthAlertAuthorizationRateDropPaymentMethodType `json:"payment_method_type"`
	// The previous authorization rate percentage.
	PreviousPercentage float64 `json:"previous_percentage,string"`
}

// Links to relevant documentation for diagnosing and resolving the alert.
type V2CoreHealthAlertDocumentationLink struct {
	// A human-readable label for the link.
	Label string `json:"label"`
	// The URL of the documentation.
	URL string `json:"url"`
}

// Populated when type is elements_error.
type V2CoreHealthAlertElementsError struct {
	// The type of the element.
	ElementType V2CoreHealthAlertElementsErrorElementType `json:"element_type,omitempty"`
	// The number of impacted sessions.
	ImpactedSessions int64 `json:"impacted_sessions"`
}

// The related object details.
type V2CoreHealthAlertEventGenerationFailureRelatedObject struct {
	// The ID of the related object (e.g., "pi_...").
	ID string `json:"id"`
	// The type of the related object (e.g., "payment_intent").
	Type string `json:"type"`
	// The API URL for the related object (e.g., "/v1/payment_intents/pi_...").
	URL string `json:"url"`
}

// Populated when type is event_generation_failure.
type V2CoreHealthAlertEventGenerationFailure struct {
	// The context the event should have been generated for. Only present when the account is a connected account.
	Context string `json:"context,omitempty"`
	// The type of event that Stripe failed to generate.
	EventType string `json:"event_type"`
	// The related object details.
	RelatedObject *V2CoreHealthAlertEventGenerationFailureRelatedObject `json:"related_object"`
}

// Populated when type is fraud_rate.
type V2CoreHealthAlertFraudRate struct {
	// Fraud attack type.
	AttackType V2CoreHealthAlertFraudRateAttackType `json:"attack_type"`
	// The number of impacted requests which are detected.
	ImpactedRequests int64 `json:"impacted_requests"`
	// Estimated aggregated amount for the impacted requests.
	RealizedFraudAmount Amount `json:"realized_fraud_amount"`
}

// Populated when type is invoice_count_dropped.
type V2CoreHealthAlertInvoiceCountDropped struct {
	// The observed number of invoices within the time window.
	ObservedCount float64 `json:"observed_count,string"`
	// The expected threshold number of invoices within the time window.
	ThresholdCount float64 `json:"threshold_count,string"`
	// The size of the observation time window.
	TimeWindow string `json:"time_window"`
}

// Populated when type is issuing_authorization_request_errors.
type V2CoreHealthAlertIssuingAuthorizationRequestErrors struct {
	// Estimated aggregated amount for the approved requests.
	ApprovedAmount Amount `json:"approved_amount,omitempty"`
	// The number of approved requests which are impacted.
	ApprovedImpactedRequests int64 `json:"approved_impacted_requests,omitempty"`
	// Estimated aggregated amount for the declined requests.
	DeclinedAmount Amount `json:"declined_amount,omitempty"`
	// The number of declined requests which are impacted.
	DeclinedImpactedRequests int64 `json:"declined_impacted_requests,omitempty"`
}

// Populated when type is issuing_authorization_request_timeout.
type V2CoreHealthAlertIssuingAuthorizationRequestTimeout struct {
	// Estimated aggregated amount for the approved requests.
	ApprovedAmount Amount `json:"approved_amount,omitempty"`
	// The number of approved requests which are impacted.
	ApprovedImpactedRequests int64 `json:"approved_impacted_requests,omitempty"`
	// Estimated aggregated amount for the declined requests.
	DeclinedAmount Amount `json:"declined_amount,omitempty"`
	// The number of declined requests which are impacted.
	DeclinedImpactedRequests int64 `json:"declined_impacted_requests,omitempty"`
}

// Populated when type is meter_event_summaries_delayed.
type V2CoreHealthAlertMeterEventSummariesDelayed struct {
	// The ingestion method.
	IngestionMethod V2CoreHealthAlertMeterEventSummariesDelayedIngestionMethod `json:"ingestion_method,omitempty"`
}

// Populated when type is metronome_notification_latency.
type V2CoreHealthAlertMetronomeNotificationLatency struct {
	// The impacted Metronome billing pipeline.
	Pipeline V2CoreHealthAlertMetronomeNotificationLatencyPipeline `json:"pipeline"`
}

// The top impacted connected accounts (only for platforms).
type V2CoreHealthAlertPaymentMethodErrorTopImpactedAccount struct {
	// The account ID of the impacted connected account.
	Account string `json:"account"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
}

// Populated when type is payment_method_error.
type V2CoreHealthAlertPaymentMethodError struct {
	// The returned error code.
	ErrorCode string `json:"error_code,omitempty"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
	// The type of the payment method.
	PaymentMethodType V2CoreHealthAlertPaymentMethodErrorPaymentMethodType `json:"payment_method_type"`
	// The top impacted connected accounts (only for platforms).
	TopImpactedAccounts []*V2CoreHealthAlertPaymentMethodErrorTopImpactedAccount `json:"top_impacted_accounts,omitempty"`
}

// Populated when type is sepa_debit_delayed.
type V2CoreHealthAlertSEPADebitDelayed struct {
	// The number of impacted payments.
	ImpactedPayments int64 `json:"impacted_payments"`
	// The percentage of impacted payments.
	ImpactedPaymentsPercentage float64 `json:"impacted_payments_percentage,string"`
}

// Populated when type is traffic_volume_drop.
type V2CoreHealthAlertTrafficVolumeDrop struct {
	// The total volume of payment requests within the latest observation time window.
	ActualTraffic int64 `json:"actual_traffic"`
	// The canonical path.
	CanonicalPath string `json:"canonical_path,omitempty"`
	// The expected volume of payment requests within the latest observation time window.
	ExpectedTraffic int64 `json:"expected_traffic,omitempty"`
	// The size of the observation time window.
	TimeWindow string `json:"time_window"`
}

// Populated when type is webhook_latency.
type V2CoreHealthAlertWebhookLatency struct {
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
}

// A health alert represents a detected problem affecting a merchant's Stripe integration.
type V2CoreHealthAlert struct {
	APIResource
	// Populated when type is api_error.
	APIError *V2CoreHealthAlertAPIError `json:"api_error,omitempty"`
	// Populated when type is api_latency.
	APILatency *V2CoreHealthAlertAPILatency `json:"api_latency,omitempty"`
	// Populated when type is authorization_rate_drop.
	AuthorizationRateDrop *V2CoreHealthAlertAuthorizationRateDrop `json:"authorization_rate_drop,omitempty"`
	// Time at which the health alert was created.
	Created time.Time `json:"created"`
	// Links to relevant documentation for diagnosing and resolving the alert.
	DocumentationLinks []*V2CoreHealthAlertDocumentationLink `json:"documentation_links"`
	// Populated when type is elements_error.
	ElementsError *V2CoreHealthAlertElementsError `json:"elements_error,omitempty"`
	// Populated when type is event_generation_failure.
	EventGenerationFailure *V2CoreHealthAlertEventGenerationFailure `json:"event_generation_failure,omitempty"`
	// Populated when type is fraud_rate.
	FraudRate *V2CoreHealthAlertFraudRate `json:"fraud_rate,omitempty"`
	// The grouping key for the alert.
	GroupingKey string `json:"grouping_key"`
	// Whether the alert is linked to an incident or is a self-contained problem.
	GroupingType V2CoreHealthAlertGroupingType `json:"grouping_type"`
	// Unique identifier for the health alert.
	ID string `json:"id"`
	// Populated when type is invoice_count_dropped.
	InvoiceCountDropped *V2CoreHealthAlertInvoiceCountDropped `json:"invoice_count_dropped,omitempty"`
	// Populated when type is issuing_authorization_request_errors.
	IssuingAuthorizationRequestErrors *V2CoreHealthAlertIssuingAuthorizationRequestErrors `json:"issuing_authorization_request_errors,omitempty"`
	// Populated when type is issuing_authorization_request_timeout.
	IssuingAuthorizationRequestTimeout *V2CoreHealthAlertIssuingAuthorizationRequestTimeout `json:"issuing_authorization_request_timeout,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Populated when type is meter_event_summaries_delayed.
	MeterEventSummariesDelayed *V2CoreHealthAlertMeterEventSummariesDelayed `json:"meter_event_summaries_delayed,omitempty"`
	// Populated when type is metronome_notification_latency.
	MetronomeNotificationLatency *V2CoreHealthAlertMetronomeNotificationLatency `json:"metronome_notification_latency,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Populated when type is payment_method_error.
	PaymentMethodError *V2CoreHealthAlertPaymentMethodError `json:"payment_method_error,omitempty"`
	// The time when the user experience has returned to expected levels. Null if the alert is still open.
	ResolvedAt time.Time `json:"resolved_at,omitempty"`
	// Populated when type is sepa_debit_delayed.
	SEPADebitDelayed *V2CoreHealthAlertSEPADebitDelayed `json:"sepa_debit_delayed,omitempty"`
	// The severity of the alert.
	Severity V2CoreHealthAlertSeverity `json:"severity"`
	// The time when impact on the user experience was first detected.
	StartedAt time.Time `json:"started_at"`
	// The current status of the alert.
	Status V2CoreHealthAlertStatus `json:"status"`
	// A short description of the alert.
	Summary string `json:"summary"`
	// Populated when type is traffic_volume_drop.
	TrafficVolumeDrop *V2CoreHealthAlertTrafficVolumeDrop `json:"traffic_volume_drop,omitempty"`
	// The type of the alert. Determines which sub-hash field is populated.
	Type V2CoreHealthAlertType `json:"type"`
	// Populated when type is webhook_latency.
	WebhookLatency *V2CoreHealthAlertWebhookLatency `json:"webhook_latency,omitempty"`
}
