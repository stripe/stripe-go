//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The HTTP method.
type V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethod string

// List of values that V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethod can take
const (
	V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethodDELETE V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethod = "DELETE"
	V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethodGET    V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethod = "GET"
	V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethodPOST   V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethod = "POST"
	V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethodPUT    V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethod = "PUT"
)

// The HTTP method.
type V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethod string

// List of values that V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethod can take
const (
	V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethodDELETE V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethod = "DELETE"
	V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethodGET    V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethod = "GET"
	V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethodPOST   V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethod = "POST"
	V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethodPUT    V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethod = "PUT"
)

// The type of the charge.
type V2CoreHealthAlertHistoryEntryAuthorizationRateDropChargeType string

// List of values that V2CoreHealthAlertHistoryEntryAuthorizationRateDropChargeType can take
const (
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropChargeTypeMoneyMoving V2CoreHealthAlertHistoryEntryAuthorizationRateDropChargeType = "money_moving"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropChargeTypeValidation  V2CoreHealthAlertHistoryEntryAuthorizationRateDropChargeType = "validation"
)

// The type of the dimension. Determines which field in dimension_details is populated.
type V2CoreHealthAlertHistoryEntryAuthorizationRateDropDimensionType string

// List of values that V2CoreHealthAlertHistoryEntryAuthorizationRateDropDimensionType can take
const (
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropDimensionTypeIssuer V2CoreHealthAlertHistoryEntryAuthorizationRateDropDimensionType = "issuer"
)

// The type of the payment method.
type V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType string

// List of values that V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType can take
const (
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeACSSDebit            V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "acss_debit"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeAffirm               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "affirm"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeAfterpayClearpay     V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "afterpay_clearpay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeAlipay               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "alipay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeAlma                 V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "alma"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeAmazonPay            V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "amazon_pay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeApplePay             V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "apple_pay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeAUBECSDebit          V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "au_becs_debit"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeBACSDebit            V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "bacs_debit"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeBancontact           V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "bancontact"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeBillie               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "billie"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeBLIK                 V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "blik"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeBoleto               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "boleto"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeCard                 V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "card"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeCardPresent          V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "card_present"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeCartesBancaires      V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "cartes_bancaires"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeCashApp              V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "cashapp"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeCrypto               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "crypto"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeDummyPassthroughCard V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "dummy_passthrough_card"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeEPS                  V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "eps"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeFPX                  V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "fpx"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeGiropay              V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "giropay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeGrabpay              V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "grabpay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeIDEAL                V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "ideal"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeInteracPresent       V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "interac_present"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeKakaoPay             V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "kakao_pay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeKlarna               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "klarna"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeKonbini              V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "konbini"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeKriya                V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "kriya"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeKrCard               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "kr_card"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeLink                 V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "link"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeMbWay                V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "mb_way"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeMobilepay            V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "mobilepay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeMondu                V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "mondu"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeMultibanco           V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "multibanco"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeNaverPay             V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "naver_pay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeNgBank               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "ng_bank"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeNgBankTransfer       V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "ng_bank_transfer"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeNgCard               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "ng_card"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeNgMarket             V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "ng_market"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeNgUssd               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "ng_ussd"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeNgWallet             V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "ng_wallet"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeOXXO                 V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "oxxo"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeP24                  V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "p24"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypePaperCheck           V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "paper_check"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypePayco                V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "payco"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypePayNow               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "paynow"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypePaypal               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "paypal"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypePaypay               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "paypay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypePayto                V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "payto"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypePayByBank            V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "pay_by_bank"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypePix                  V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "pix"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypePromptPay            V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "promptpay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeRechnung             V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "rechnung"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeRevolutPay           V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "revolut_pay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeSamsungPay           V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "samsung_pay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeSatispay             V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "satispay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeScalapay             V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "scalapay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeSEPADebit            V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "sepa_debit"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeSequra               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "sequra"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeSofort               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "sofort"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeSunbit               V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "sunbit"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeSwish                V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "swish"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeTWINT                V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "twint"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeUpi                  V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "upi"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeUSBankAccount        V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "us_bank_account"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeVipps                V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "vipps"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeWeChatPay            V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "wechat_pay"
	V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodTypeZip                  V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType = "zip"
)

// The type of the element.
type V2CoreHealthAlertHistoryEntryElementsErrorElementType string

// List of values that V2CoreHealthAlertHistoryEntryElementsErrorElementType can take
const (
	V2CoreHealthAlertHistoryEntryElementsErrorElementTypeExpressCheckout V2CoreHealthAlertHistoryEntryElementsErrorElementType = "expressCheckout"
	V2CoreHealthAlertHistoryEntryElementsErrorElementTypePayment         V2CoreHealthAlertHistoryEntryElementsErrorElementType = "payment"
)

// Fraud attack type.
type V2CoreHealthAlertHistoryEntryFraudRateAttackType string

// List of values that V2CoreHealthAlertHistoryEntryFraudRateAttackType can take
const (
	V2CoreHealthAlertHistoryEntryFraudRateAttackTypeSpike           V2CoreHealthAlertHistoryEntryFraudRateAttackType = "spike"
	V2CoreHealthAlertHistoryEntryFraudRateAttackTypeSustainedAttack V2CoreHealthAlertHistoryEntryFraudRateAttackType = "sustained_attack"
)

// The ingestion method.
type V2CoreHealthAlertHistoryEntryMeterEventSummariesDelayedIngestionMethod string

// List of values that V2CoreHealthAlertHistoryEntryMeterEventSummariesDelayedIngestionMethod can take
const (
	V2CoreHealthAlertHistoryEntryMeterEventSummariesDelayedIngestionMethodImportSets V2CoreHealthAlertHistoryEntryMeterEventSummariesDelayedIngestionMethod = "import_sets"
)

// The impacted Metronome billing pipeline.
type V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipeline string

// List of values that V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipeline can take
const (
	V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipelineConfigurationTriggered        V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipeline = "configuration_triggered"
	V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipelineHighCardinalityUsageTriggered V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipeline = "high_cardinality_usage_triggered"
	V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipelineStandardUsageTriggered        V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipeline = "standard_usage_triggered"
	V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipelineTimeTriggered                 V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipeline = "time_triggered"
)

// The type of the payment method.
type V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType string

// List of values that V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType can take
const (
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeACSSDebit            V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "acss_debit"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeAffirm               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "affirm"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeAfterpayClearpay     V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "afterpay_clearpay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeAlipay               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "alipay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeAlma                 V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "alma"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeAmazonPay            V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "amazon_pay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeApplePay             V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "apple_pay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeAUBECSDebit          V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "au_becs_debit"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeBACSDebit            V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "bacs_debit"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeBancontact           V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "bancontact"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeBillie               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "billie"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeBLIK                 V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "blik"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeBoleto               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "boleto"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeCard                 V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "card"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeCardPresent          V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "card_present"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeCartesBancaires      V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "cartes_bancaires"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeCashApp              V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "cashapp"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeCrypto               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "crypto"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeDummyPassthroughCard V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "dummy_passthrough_card"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeEPS                  V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "eps"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeFPX                  V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "fpx"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeGiropay              V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "giropay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeGrabpay              V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "grabpay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeIDEAL                V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "ideal"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeInteracPresent       V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "interac_present"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeKakaoPay             V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "kakao_pay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeKlarna               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "klarna"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeKonbini              V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "konbini"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeKriya                V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "kriya"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeKrCard               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "kr_card"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeLink                 V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "link"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeMbWay                V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "mb_way"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeMobilepay            V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "mobilepay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeMondu                V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "mondu"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeMultibanco           V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "multibanco"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeNaverPay             V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "naver_pay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeNgBank               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "ng_bank"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeNgBankTransfer       V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "ng_bank_transfer"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeNgCard               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "ng_card"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeNgMarket             V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "ng_market"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeNgUssd               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "ng_ussd"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeNgWallet             V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "ng_wallet"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeOXXO                 V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "oxxo"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeP24                  V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "p24"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypePaperCheck           V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "paper_check"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypePayco                V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "payco"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypePayNow               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "paynow"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypePaypal               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "paypal"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypePaypay               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "paypay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypePayto                V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "payto"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypePayByBank            V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "pay_by_bank"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypePix                  V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "pix"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypePromptPay            V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "promptpay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeRechnung             V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "rechnung"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeRevolutPay           V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "revolut_pay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeSamsungPay           V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "samsung_pay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeSatispay             V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "satispay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeScalapay             V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "scalapay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeSEPADebit            V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "sepa_debit"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeSequra               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "sequra"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeSofort               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "sofort"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeSunbit               V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "sunbit"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeSwish                V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "swish"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeTWINT                V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "twint"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeUpi                  V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "upi"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeUSBankAccount        V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "us_bank_account"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeVipps                V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "vipps"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeWeChatPay            V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "wechat_pay"
	V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodTypeZip                  V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType = "zip"
)

// The type of transition that occurred.
type V2CoreHealthAlertHistoryEntryTransition string

// List of values that V2CoreHealthAlertHistoryEntryTransition can take
const (
	V2CoreHealthAlertHistoryEntryTransitionImpactUpdated V2CoreHealthAlertHistoryEntryTransition = "impact_updated"
	V2CoreHealthAlertHistoryEntryTransitionOpened        V2CoreHealthAlertHistoryEntryTransition = "opened"
	V2CoreHealthAlertHistoryEntryTransitionResolved      V2CoreHealthAlertHistoryEntryTransition = "resolved"
)

// The type of the alert. Determines which sub-hash field is populated.
type V2CoreHealthAlertHistoryEntryType string

// List of values that V2CoreHealthAlertHistoryEntryType can take
const (
	V2CoreHealthAlertHistoryEntryTypeAPIError                           V2CoreHealthAlertHistoryEntryType = "api_error"
	V2CoreHealthAlertHistoryEntryTypeAPILatency                         V2CoreHealthAlertHistoryEntryType = "api_latency"
	V2CoreHealthAlertHistoryEntryTypeAuthorizationRateDrop              V2CoreHealthAlertHistoryEntryType = "authorization_rate_drop"
	V2CoreHealthAlertHistoryEntryTypeElementsError                      V2CoreHealthAlertHistoryEntryType = "elements_error"
	V2CoreHealthAlertHistoryEntryTypeEventGenerationFailure             V2CoreHealthAlertHistoryEntryType = "event_generation_failure"
	V2CoreHealthAlertHistoryEntryTypeFraudRate                          V2CoreHealthAlertHistoryEntryType = "fraud_rate"
	V2CoreHealthAlertHistoryEntryTypeInvoiceCountDropped                V2CoreHealthAlertHistoryEntryType = "invoice_count_dropped"
	V2CoreHealthAlertHistoryEntryTypeIssuingAuthorizationRequestErrors  V2CoreHealthAlertHistoryEntryType = "issuing_authorization_request_errors"
	V2CoreHealthAlertHistoryEntryTypeIssuingAuthorizationRequestTimeout V2CoreHealthAlertHistoryEntryType = "issuing_authorization_request_timeout"
	V2CoreHealthAlertHistoryEntryTypeMeterEventSummariesDelayed         V2CoreHealthAlertHistoryEntryType = "meter_event_summaries_delayed"
	V2CoreHealthAlertHistoryEntryTypeMetronomeNotificationLatency       V2CoreHealthAlertHistoryEntryType = "metronome_notification_latency"
	V2CoreHealthAlertHistoryEntryTypePaymentMethodError                 V2CoreHealthAlertHistoryEntryType = "payment_method_error"
	V2CoreHealthAlertHistoryEntryTypeSEPADebitDelayed                   V2CoreHealthAlertHistoryEntryType = "sepa_debit_delayed"
	V2CoreHealthAlertHistoryEntryTypeTrafficVolumeDrop                  V2CoreHealthAlertHistoryEntryType = "traffic_volume_drop"
	V2CoreHealthAlertHistoryEntryTypeWebhookLatency                     V2CoreHealthAlertHistoryEntryType = "webhook_latency"
)

// The top impacted connected accounts (only for platforms).
type V2CoreHealthAlertHistoryEntryAPIErrorTopImpactedAccount struct {
	// The account ID of the impacted connected account.
	Account string `json:"account"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
}

// Populated when type is api_error.
type V2CoreHealthAlertHistoryEntryAPIError struct {
	// The canonical path.
	CanonicalPath string `json:"canonical_path"`
	// The error code.
	ErrorCode string `json:"error_code,omitempty"`
	// The HTTP method.
	HTTPMethod V2CoreHealthAlertHistoryEntryAPIErrorHTTPMethod `json:"http_method"`
	// The HTTP status.
	HTTPStatus string `json:"http_status"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
	// The top impacted connected accounts (only for platforms).
	TopImpactedAccounts []*V2CoreHealthAlertHistoryEntryAPIErrorTopImpactedAccount `json:"top_impacted_accounts,omitempty"`
}

// The top impacted connected accounts (only for platforms).
type V2CoreHealthAlertHistoryEntryAPILatencyTopImpactedAccount struct {
	// The account ID of the impacted connected account.
	Account string `json:"account"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
}

// Populated when type is api_latency.
type V2CoreHealthAlertHistoryEntryAPILatency struct {
	// The canonical path.
	CanonicalPath string `json:"canonical_path"`
	// The HTTP method.
	HTTPMethod V2CoreHealthAlertHistoryEntryAPILatencyHTTPMethod `json:"http_method"`
	// The HTTP status.
	HTTPStatus string `json:"http_status"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
	// The top impacted connected accounts (only for platforms).
	TopImpactedAccounts []*V2CoreHealthAlertHistoryEntryAPILatencyTopImpactedAccount `json:"top_impacted_accounts,omitempty"`
}

// Dimensions that describe what subset of payments are impacted.
type V2CoreHealthAlertHistoryEntryAuthorizationRateDropDimension struct {
	// Populated when type is issuer.
	Issuer string `json:"issuer,omitempty"`
	// The type of the dimension. Determines which field in dimension_details is populated.
	Type V2CoreHealthAlertHistoryEntryAuthorizationRateDropDimensionType `json:"type"`
}

// Populated when type is authorization_rate_drop.
type V2CoreHealthAlertHistoryEntryAuthorizationRateDrop struct {
	// The type of the charge.
	ChargeType V2CoreHealthAlertHistoryEntryAuthorizationRateDropChargeType `json:"charge_type"`
	// The current authorization rate percentage.
	CurrentPercentage float64 `json:"current_percentage,string"`
	// Dimensions that describe what subset of payments are impacted.
	Dimensions []*V2CoreHealthAlertHistoryEntryAuthorizationRateDropDimension `json:"dimensions,omitempty"`
	// The type of the payment method.
	PaymentMethodType V2CoreHealthAlertHistoryEntryAuthorizationRateDropPaymentMethodType `json:"payment_method_type"`
	// The previous authorization rate percentage.
	PreviousPercentage float64 `json:"previous_percentage,string"`
}

// Populated when type is elements_error.
type V2CoreHealthAlertHistoryEntryElementsError struct {
	// The type of the element.
	ElementType V2CoreHealthAlertHistoryEntryElementsErrorElementType `json:"element_type,omitempty"`
	// The number of impacted sessions.
	ImpactedSessions int64 `json:"impacted_sessions"`
}

// The related object details.
type V2CoreHealthAlertHistoryEntryEventGenerationFailureRelatedObject struct {
	// The ID of the related object (e.g., "pi_...").
	ID string `json:"id"`
	// The type of the related object (e.g., "payment_intent").
	Type string `json:"type"`
	// The API URL for the related object (e.g., "/v1/payment_intents/pi_...").
	URL string `json:"url"`
}

// Populated when type is event_generation_failure.
type V2CoreHealthAlertHistoryEntryEventGenerationFailure struct {
	// The context the event should have been generated for. Only present when the account is a connected account.
	Context string `json:"context,omitempty"`
	// The type of event that Stripe failed to generate.
	EventType string `json:"event_type"`
	// The related object details.
	RelatedObject *V2CoreHealthAlertHistoryEntryEventGenerationFailureRelatedObject `json:"related_object"`
}

// Populated when type is fraud_rate.
type V2CoreHealthAlertHistoryEntryFraudRate struct {
	// Fraud attack type.
	AttackType V2CoreHealthAlertHistoryEntryFraudRateAttackType `json:"attack_type"`
	// The number of impacted requests which are detected.
	ImpactedRequests int64 `json:"impacted_requests"`
	// Estimated aggregated amount for the impacted requests.
	RealizedFraudAmount Amount `json:"realized_fraud_amount"`
}

// Populated when type is invoice_count_dropped.
type V2CoreHealthAlertHistoryEntryInvoiceCountDropped struct {
	// The observed number of invoices within the time window.
	ObservedCount float64 `json:"observed_count,string"`
	// The expected threshold number of invoices within the time window.
	ThresholdCount float64 `json:"threshold_count,string"`
	// The size of the observation time window.
	TimeWindow string `json:"time_window"`
}

// Populated when type is issuing_authorization_request_errors.
type V2CoreHealthAlertHistoryEntryIssuingAuthorizationRequestErrors struct {
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
type V2CoreHealthAlertHistoryEntryIssuingAuthorizationRequestTimeout struct {
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
type V2CoreHealthAlertHistoryEntryMeterEventSummariesDelayed struct {
	// The ingestion method.
	IngestionMethod V2CoreHealthAlertHistoryEntryMeterEventSummariesDelayedIngestionMethod `json:"ingestion_method,omitempty"`
}

// Populated when type is metronome_notification_latency.
type V2CoreHealthAlertHistoryEntryMetronomeNotificationLatency struct {
	// The impacted Metronome billing pipeline.
	Pipeline V2CoreHealthAlertHistoryEntryMetronomeNotificationLatencyPipeline `json:"pipeline"`
}

// The top impacted connected accounts (only for platforms).
type V2CoreHealthAlertHistoryEntryPaymentMethodErrorTopImpactedAccount struct {
	// The account ID of the impacted connected account.
	Account string `json:"account"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
}

// Populated when type is payment_method_error.
type V2CoreHealthAlertHistoryEntryPaymentMethodError struct {
	// The returned error code.
	ErrorCode string `json:"error_code,omitempty"`
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
	// The percentage of impacted requests.
	ImpactedRequestsPercentage float64 `json:"impacted_requests_percentage,string,omitempty"`
	// The type of the payment method.
	PaymentMethodType V2CoreHealthAlertHistoryEntryPaymentMethodErrorPaymentMethodType `json:"payment_method_type"`
	// The top impacted connected accounts (only for platforms).
	TopImpactedAccounts []*V2CoreHealthAlertHistoryEntryPaymentMethodErrorTopImpactedAccount `json:"top_impacted_accounts,omitempty"`
}

// Populated when type is sepa_debit_delayed.
type V2CoreHealthAlertHistoryEntrySEPADebitDelayed struct {
	// The number of impacted payments.
	ImpactedPayments int64 `json:"impacted_payments"`
	// The percentage of impacted payments.
	ImpactedPaymentsPercentage float64 `json:"impacted_payments_percentage,string"`
}

// Populated when type is traffic_volume_drop.
type V2CoreHealthAlertHistoryEntryTrafficVolumeDrop struct {
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
type V2CoreHealthAlertHistoryEntryWebhookLatency struct {
	// The number of impacted requests.
	ImpactedRequests int64 `json:"impacted_requests"`
}

// An alert history entry representing a state transition of a health alert.
type V2CoreHealthAlertHistoryEntry struct {
	APIResource
	// Populated when type is api_error.
	APIError *V2CoreHealthAlertHistoryEntryAPIError `json:"api_error,omitempty"`
	// Populated when type is api_latency.
	APILatency *V2CoreHealthAlertHistoryEntryAPILatency `json:"api_latency,omitempty"`
	// The time at which this transition occurred.
	At time.Time `json:"at"`
	// Populated when type is authorization_rate_drop.
	AuthorizationRateDrop *V2CoreHealthAlertHistoryEntryAuthorizationRateDrop `json:"authorization_rate_drop,omitempty"`
	// Populated when type is elements_error.
	ElementsError *V2CoreHealthAlertHistoryEntryElementsError `json:"elements_error,omitempty"`
	// Populated when type is event_generation_failure.
	EventGenerationFailure *V2CoreHealthAlertHistoryEntryEventGenerationFailure `json:"event_generation_failure,omitempty"`
	// Populated when type is fraud_rate.
	FraudRate *V2CoreHealthAlertHistoryEntryFraudRate `json:"fraud_rate,omitempty"`
	// Unique identifier for the alert history entry.
	ID string `json:"id"`
	// Populated when type is invoice_count_dropped.
	InvoiceCountDropped *V2CoreHealthAlertHistoryEntryInvoiceCountDropped `json:"invoice_count_dropped,omitempty"`
	// Populated when type is issuing_authorization_request_errors.
	IssuingAuthorizationRequestErrors *V2CoreHealthAlertHistoryEntryIssuingAuthorizationRequestErrors `json:"issuing_authorization_request_errors,omitempty"`
	// Populated when type is issuing_authorization_request_timeout.
	IssuingAuthorizationRequestTimeout *V2CoreHealthAlertHistoryEntryIssuingAuthorizationRequestTimeout `json:"issuing_authorization_request_timeout,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Populated when type is meter_event_summaries_delayed.
	MeterEventSummariesDelayed *V2CoreHealthAlertHistoryEntryMeterEventSummariesDelayed `json:"meter_event_summaries_delayed,omitempty"`
	// Populated when type is metronome_notification_latency.
	MetronomeNotificationLatency *V2CoreHealthAlertHistoryEntryMetronomeNotificationLatency `json:"metronome_notification_latency,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Populated when type is payment_method_error.
	PaymentMethodError *V2CoreHealthAlertHistoryEntryPaymentMethodError `json:"payment_method_error,omitempty"`
	// Populated when type is sepa_debit_delayed.
	SEPADebitDelayed *V2CoreHealthAlertHistoryEntrySEPADebitDelayed `json:"sepa_debit_delayed,omitempty"`
	// Populated when type is traffic_volume_drop.
	TrafficVolumeDrop *V2CoreHealthAlertHistoryEntryTrafficVolumeDrop `json:"traffic_volume_drop,omitempty"`
	// The type of transition that occurred.
	Transition V2CoreHealthAlertHistoryEntryTransition `json:"transition"`
	// The type of the alert. Determines which sub-hash field is populated.
	Type V2CoreHealthAlertHistoryEntryType `json:"type"`
	// Populated when type is webhook_latency.
	WebhookLatency *V2CoreHealthAlertHistoryEntryWebhookLatency `json:"webhook_latency,omitempty"`
}
