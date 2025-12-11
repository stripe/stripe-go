//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Assessments from Stripe. If set, the value is `fraudulent`.
type ChargeFraudStripeReport string

// List of values that ChargeFraudStripeReport can take
const (
	ChargeFraudStripeReportFraudulent ChargeFraudStripeReport = "fraudulent"
)

// Assessments reported by you. If set, possible values of are `safe` and `fraudulent`.
type ChargeFraudUserReport string

// List of values that ChargeFraudUserReport can take
const (
	ChargeFraudUserReportFraudulent ChargeFraudUserReport = "fraudulent"
	ChargeFraudUserReportSafe       ChargeFraudUserReport = "safe"
)

// An enumerated value providing a more detailed explanation on [how to proceed with an error](https://docs.stripe.com/declines#retrying-issuer-declines).
type ChargeOutcomeAdviceCode string

// List of values that ChargeOutcomeAdviceCode can take
const (
	ChargeOutcomeAdviceCodeConfirmCardData ChargeOutcomeAdviceCode = "confirm_card_data"
	ChargeOutcomeAdviceCodeDoNotTryAgain   ChargeOutcomeAdviceCode = "do_not_try_again"
	ChargeOutcomeAdviceCodeTryAgainLater   ChargeOutcomeAdviceCode = "try_again_later"
)

// funding type of the underlying payment method.
type ChargePaymentMethodDetailsAmazonPayFundingType string

// List of values that ChargePaymentMethodDetailsAmazonPayFundingType can take
const (
	ChargePaymentMethodDetailsAmazonPayFundingTypeCard ChargePaymentMethodDetailsAmazonPayFundingType = "card"
)

// If a address line1 was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
type ChargePaymentMethodDetailsCardChecksAddressLine1Check string

// List of values that ChargePaymentMethodDetailsCardChecksAddressLine1Check can take
const (
	ChargePaymentMethodDetailsCardChecksAddressLine1CheckFail        ChargePaymentMethodDetailsCardChecksAddressLine1Check = "fail"
	ChargePaymentMethodDetailsCardChecksAddressLine1CheckPass        ChargePaymentMethodDetailsCardChecksAddressLine1Check = "pass"
	ChargePaymentMethodDetailsCardChecksAddressLine1CheckUnavailable ChargePaymentMethodDetailsCardChecksAddressLine1Check = "unavailable"
	ChargePaymentMethodDetailsCardChecksAddressLine1CheckUnchecked   ChargePaymentMethodDetailsCardChecksAddressLine1Check = "unchecked"
)

// If a address postal code was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
type ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheck string

// List of values that ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheck can take
const (
	ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheckFail        ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheck = "fail"
	ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheckPass        ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheck = "pass"
	ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheckUnavailable ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheck = "unavailable"
	ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheckUnchecked   ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheck = "unchecked"
)

// If a CVC was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
type ChargePaymentMethodDetailsCardChecksCVCCheck string

// List of values that ChargePaymentMethodDetailsCardChecksCVCCheck can take
const (
	ChargePaymentMethodDetailsCardChecksCVCCheckFail        ChargePaymentMethodDetailsCardChecksCVCCheck = "fail"
	ChargePaymentMethodDetailsCardChecksCVCCheckPass        ChargePaymentMethodDetailsCardChecksCVCCheck = "pass"
	ChargePaymentMethodDetailsCardChecksCVCCheckUnavailable ChargePaymentMethodDetailsCardChecksCVCCheck = "unavailable"
	ChargePaymentMethodDetailsCardChecksCVCCheckUnchecked   ChargePaymentMethodDetailsCardChecksCVCCheck = "unchecked"
)

// Indicates whether or not the decremental authorization feature is supported.
type ChargePaymentMethodDetailsCardDecrementalAuthorizationStatus string

// List of values that ChargePaymentMethodDetailsCardDecrementalAuthorizationStatus can take
const (
	ChargePaymentMethodDetailsCardDecrementalAuthorizationStatusAvailable   ChargePaymentMethodDetailsCardDecrementalAuthorizationStatus = "available"
	ChargePaymentMethodDetailsCardDecrementalAuthorizationStatusUnavailable ChargePaymentMethodDetailsCardDecrementalAuthorizationStatus = "unavailable"
)

// Indicates whether or not the capture window is extended beyond the standard authorization.
type ChargePaymentMethodDetailsCardExtendedAuthorizationStatus string

// List of values that ChargePaymentMethodDetailsCardExtendedAuthorizationStatus can take
const (
	ChargePaymentMethodDetailsCardExtendedAuthorizationStatusDisabled ChargePaymentMethodDetailsCardExtendedAuthorizationStatus = "disabled"
	ChargePaymentMethodDetailsCardExtendedAuthorizationStatusEnabled  ChargePaymentMethodDetailsCardExtendedAuthorizationStatus = "enabled"
)

// Indicates whether or not the incremental authorization feature is supported.
type ChargePaymentMethodDetailsCardIncrementalAuthorizationStatus string

// List of values that ChargePaymentMethodDetailsCardIncrementalAuthorizationStatus can take
const (
	ChargePaymentMethodDetailsCardIncrementalAuthorizationStatusAvailable   ChargePaymentMethodDetailsCardIncrementalAuthorizationStatus = "available"
	ChargePaymentMethodDetailsCardIncrementalAuthorizationStatusUnavailable ChargePaymentMethodDetailsCardIncrementalAuthorizationStatus = "unavailable"
)

// Indicates whether or not multiple captures are supported.
type ChargePaymentMethodDetailsCardMulticaptureStatus string

// List of values that ChargePaymentMethodDetailsCardMulticaptureStatus can take
const (
	ChargePaymentMethodDetailsCardMulticaptureStatusAvailable   ChargePaymentMethodDetailsCardMulticaptureStatus = "available"
	ChargePaymentMethodDetailsCardMulticaptureStatusUnavailable ChargePaymentMethodDetailsCardMulticaptureStatus = "unavailable"
)

// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
type ChargePaymentMethodDetailsCardNetwork string

// List of values that ChargePaymentMethodDetailsCardNetwork can take
const (
	ChargePaymentMethodDetailsCardNetworkAmex            ChargePaymentMethodDetailsCardNetwork = "amex"
	ChargePaymentMethodDetailsCardNetworkCartesBancaires ChargePaymentMethodDetailsCardNetwork = "cartes_bancaires"
	ChargePaymentMethodDetailsCardNetworkDiners          ChargePaymentMethodDetailsCardNetwork = "diners"
	ChargePaymentMethodDetailsCardNetworkDiscover        ChargePaymentMethodDetailsCardNetwork = "discover"
	ChargePaymentMethodDetailsCardNetworkInterac         ChargePaymentMethodDetailsCardNetwork = "interac"
	ChargePaymentMethodDetailsCardNetworkJCB             ChargePaymentMethodDetailsCardNetwork = "jcb"
	ChargePaymentMethodDetailsCardNetworkMastercard      ChargePaymentMethodDetailsCardNetwork = "mastercard"
	ChargePaymentMethodDetailsCardNetworkUnionpay        ChargePaymentMethodDetailsCardNetwork = "unionpay"
	ChargePaymentMethodDetailsCardNetworkVisa            ChargePaymentMethodDetailsCardNetwork = "visa"
	ChargePaymentMethodDetailsCardNetworkUnknown         ChargePaymentMethodDetailsCardNetwork = "unknown"
)

// Indicates whether or not the authorized amount can be over-captured.
type ChargePaymentMethodDetailsCardOvercaptureStatus string

// List of values that ChargePaymentMethodDetailsCardOvercaptureStatus can take
const (
	ChargePaymentMethodDetailsCardOvercaptureStatusAvailable   ChargePaymentMethodDetailsCardOvercaptureStatus = "available"
	ChargePaymentMethodDetailsCardOvercaptureStatusUnavailable ChargePaymentMethodDetailsCardOvercaptureStatus = "unavailable"
)

// Indicates whether the transaction requested for partial authorization feature and the authorization outcome.
type ChargePaymentMethodDetailsCardPartialAuthorizationStatus string

// List of values that ChargePaymentMethodDetailsCardPartialAuthorizationStatus can take
const (
	ChargePaymentMethodDetailsCardPartialAuthorizationStatusDeclined            ChargePaymentMethodDetailsCardPartialAuthorizationStatus = "declined"
	ChargePaymentMethodDetailsCardPartialAuthorizationStatusFullyAuthorized     ChargePaymentMethodDetailsCardPartialAuthorizationStatus = "fully_authorized"
	ChargePaymentMethodDetailsCardPartialAuthorizationStatusNotRequested        ChargePaymentMethodDetailsCardPartialAuthorizationStatus = "not_requested"
	ChargePaymentMethodDetailsCardPartialAuthorizationStatusPartiallyAuthorized ChargePaymentMethodDetailsCardPartialAuthorizationStatus = "partially_authorized"
)

// Status of a card based on the card issuer.
type ChargePaymentMethodDetailsCardRegulatedStatus string

// List of values that ChargePaymentMethodDetailsCardRegulatedStatus can take
const (
	ChargePaymentMethodDetailsCardRegulatedStatusRegulated   ChargePaymentMethodDetailsCardRegulatedStatus = "regulated"
	ChargePaymentMethodDetailsCardRegulatedStatusUnregulated ChargePaymentMethodDetailsCardRegulatedStatus = "unregulated"
)

// For authenticated transactions: how the customer was authenticated by
// the issuing bank.
type ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// The Electronic Commerce Indicator (ECI). A protocol-level field
// indicating what degree of authentication was performed.
type ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator01 ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator = "01"
	ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator02 ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator = "02"
	ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator05 ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator = "05"
	ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator06 ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator = "06"
	ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator07 ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator = "07"
)

// The exemption requested via 3DS and accepted by the issuer at authentication time.
type ChargePaymentMethodDetailsCardThreeDSecureExemptionIndicator string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureExemptionIndicator can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureExemptionIndicatorLowRisk ChargePaymentMethodDetailsCardThreeDSecureExemptionIndicator = "low_risk"
	ChargePaymentMethodDetailsCardThreeDSecureExemptionIndicatorNone    ChargePaymentMethodDetailsCardThreeDSecureExemptionIndicator = "none"
)

// Indicates the outcome of 3D Secure authentication.
type ChargePaymentMethodDetailsCardThreeDSecureResult string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResult can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged ChargePaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	ChargePaymentMethodDetailsCardThreeDSecureResultAuthenticated       ChargePaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	ChargePaymentMethodDetailsCardThreeDSecureResultExempted            ChargePaymentMethodDetailsCardThreeDSecureResult = "exempted"
	ChargePaymentMethodDetailsCardThreeDSecureResultFailed              ChargePaymentMethodDetailsCardThreeDSecureResult = "failed"
	ChargePaymentMethodDetailsCardThreeDSecureResultNotSupported        ChargePaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultProcessingError     ChargePaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// Additional information about why 3D Secure succeeded or failed based
// on the `result`.
type ChargePaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResultReason can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           ChargePaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonBypassed            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCanceled            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     ChargePaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported ChargePaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       ChargePaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonRejected            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
type ChargePaymentMethodDetailsCardPresentNetwork string

// List of values that ChargePaymentMethodDetailsCardPresentNetwork can take
const (
	ChargePaymentMethodDetailsCardPresentNetworkAmex            ChargePaymentMethodDetailsCardPresentNetwork = "amex"
	ChargePaymentMethodDetailsCardPresentNetworkCartesBancaires ChargePaymentMethodDetailsCardPresentNetwork = "cartes_bancaires"
	ChargePaymentMethodDetailsCardPresentNetworkDiners          ChargePaymentMethodDetailsCardPresentNetwork = "diners"
	ChargePaymentMethodDetailsCardPresentNetworkDiscover        ChargePaymentMethodDetailsCardPresentNetwork = "discover"
	ChargePaymentMethodDetailsCardPresentNetworkInterac         ChargePaymentMethodDetailsCardPresentNetwork = "interac"
	ChargePaymentMethodDetailsCardPresentNetworkJCB             ChargePaymentMethodDetailsCardPresentNetwork = "jcb"
	ChargePaymentMethodDetailsCardPresentNetworkMastercard      ChargePaymentMethodDetailsCardPresentNetwork = "mastercard"
	ChargePaymentMethodDetailsCardPresentNetworkUnionpay        ChargePaymentMethodDetailsCardPresentNetwork = "unionpay"
	ChargePaymentMethodDetailsCardPresentNetworkVisa            ChargePaymentMethodDetailsCardPresentNetwork = "visa"
	ChargePaymentMethodDetailsCardPresentNetworkUnknown         ChargePaymentMethodDetailsCardPresentNetwork = "unknown"
)

// The method used to process this payment method offline. Only deferred is allowed.
type ChargePaymentMethodDetailsCardPresentOfflineType string

// List of values that ChargePaymentMethodDetailsCardPresentOfflineType can take
const (
	ChargePaymentMethodDetailsCardPresentOfflineTypeDeferred ChargePaymentMethodDetailsCardPresentOfflineType = "deferred"
)

// The type of account being debited or credited
type ChargePaymentMethodDetailsCardPresentReceiptAccountType string

// List of values that ChargePaymentMethodDetailsCardPresentReceiptAccountType can take
const (
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeChecking ChargePaymentMethodDetailsCardPresentReceiptAccountType = "checking"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeCredit   ChargePaymentMethodDetailsCardPresentReceiptAccountType = "credit"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypePrepaid  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "prepaid"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeUnknown  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "unknown"
)

// The type of mobile wallet, one of `apple_pay`, `google_pay`, `samsung_pay`, or `unknown`.
type ChargePaymentMethodDetailsCardPresentWalletType string

// List of values that ChargePaymentMethodDetailsCardPresentWalletType can take
const (
	ChargePaymentMethodDetailsCardPresentWalletTypeApplePay   ChargePaymentMethodDetailsCardPresentWalletType = "apple_pay"
	ChargePaymentMethodDetailsCardPresentWalletTypeGooglePay  ChargePaymentMethodDetailsCardPresentWalletType = "google_pay"
	ChargePaymentMethodDetailsCardPresentWalletTypeSamsungPay ChargePaymentMethodDetailsCardPresentWalletType = "samsung_pay"
	ChargePaymentMethodDetailsCardPresentWalletTypeUnknown    ChargePaymentMethodDetailsCardPresentWalletType = "unknown"
)

// The blockchain network that the transaction was sent on.
type ChargePaymentMethodDetailsCryptoNetwork string

// List of values that ChargePaymentMethodDetailsCryptoNetwork can take
const (
	ChargePaymentMethodDetailsCryptoNetworkBase     ChargePaymentMethodDetailsCryptoNetwork = "base"
	ChargePaymentMethodDetailsCryptoNetworkEthereum ChargePaymentMethodDetailsCryptoNetwork = "ethereum"
	ChargePaymentMethodDetailsCryptoNetworkPolygon  ChargePaymentMethodDetailsCryptoNetwork = "polygon"
	ChargePaymentMethodDetailsCryptoNetworkSolana   ChargePaymentMethodDetailsCryptoNetwork = "solana"
)

// The token currency that the transaction was sent with.
type ChargePaymentMethodDetailsCryptoTokenCurrency string

// List of values that ChargePaymentMethodDetailsCryptoTokenCurrency can take
const (
	ChargePaymentMethodDetailsCryptoTokenCurrencyUsdc ChargePaymentMethodDetailsCryptoTokenCurrency = "usdc"
	ChargePaymentMethodDetailsCryptoTokenCurrencyUsdg ChargePaymentMethodDetailsCryptoTokenCurrency = "usdg"
	ChargePaymentMethodDetailsCryptoTokenCurrencyUsdp ChargePaymentMethodDetailsCryptoTokenCurrency = "usdp"
)

// Bank where the account is located.
type ChargePaymentMethodDetailsIDBankTransferBank string

// List of values that ChargePaymentMethodDetailsIDBankTransferBank can take
const (
	ChargePaymentMethodDetailsIDBankTransferBankBca     ChargePaymentMethodDetailsIDBankTransferBank = "bca"
	ChargePaymentMethodDetailsIDBankTransferBankBni     ChargePaymentMethodDetailsIDBankTransferBank = "bni"
	ChargePaymentMethodDetailsIDBankTransferBankBri     ChargePaymentMethodDetailsIDBankTransferBank = "bri"
	ChargePaymentMethodDetailsIDBankTransferBankCimb    ChargePaymentMethodDetailsIDBankTransferBank = "cimb"
	ChargePaymentMethodDetailsIDBankTransferBankPermata ChargePaymentMethodDetailsIDBankTransferBank = "permata"
)

// The Klarna payment method used for this transaction.
// Can be one of `pay_later`, `pay_now`, `pay_with_financing`, or `pay_in_installments`
type ChargePaymentMethodDetailsKlarnaPaymentMethodCategory string

// List of values that ChargePaymentMethodDetailsKlarnaPaymentMethodCategory can take
const (
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayLater          ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_later"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayNow            ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_now"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayWithFinancing  ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_with_financing"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayInInstallments ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_in_installments"
)

// The name of the convenience store chain where the payment was completed.
type ChargePaymentMethodDetailsKonbiniStoreChain string

// List of values that ChargePaymentMethodDetailsKonbiniStoreChain can take
const (
	ChargePaymentMethodDetailsKonbiniStoreChainFamilyMart ChargePaymentMethodDetailsKonbiniStoreChain = "familymart"
	ChargePaymentMethodDetailsKonbiniStoreChainLawson     ChargePaymentMethodDetailsKonbiniStoreChain = "lawson"
	ChargePaymentMethodDetailsKonbiniStoreChainMinistop   ChargePaymentMethodDetailsKonbiniStoreChain = "ministop"
	ChargePaymentMethodDetailsKonbiniStoreChainSeicomart  ChargePaymentMethodDetailsKonbiniStoreChain = "seicomart"
)

// The local credit or debit card brand.
type ChargePaymentMethodDetailsKrCardBrand string

// List of values that ChargePaymentMethodDetailsKrCardBrand can take
const (
	ChargePaymentMethodDetailsKrCardBrandBc          ChargePaymentMethodDetailsKrCardBrand = "bc"
	ChargePaymentMethodDetailsKrCardBrandCiti        ChargePaymentMethodDetailsKrCardBrand = "citi"
	ChargePaymentMethodDetailsKrCardBrandHana        ChargePaymentMethodDetailsKrCardBrand = "hana"
	ChargePaymentMethodDetailsKrCardBrandHyundai     ChargePaymentMethodDetailsKrCardBrand = "hyundai"
	ChargePaymentMethodDetailsKrCardBrandJeju        ChargePaymentMethodDetailsKrCardBrand = "jeju"
	ChargePaymentMethodDetailsKrCardBrandJeonbuk     ChargePaymentMethodDetailsKrCardBrand = "jeonbuk"
	ChargePaymentMethodDetailsKrCardBrandKakaobank   ChargePaymentMethodDetailsKrCardBrand = "kakaobank"
	ChargePaymentMethodDetailsKrCardBrandKbank       ChargePaymentMethodDetailsKrCardBrand = "kbank"
	ChargePaymentMethodDetailsKrCardBrandKdbbank     ChargePaymentMethodDetailsKrCardBrand = "kdbbank"
	ChargePaymentMethodDetailsKrCardBrandKookmin     ChargePaymentMethodDetailsKrCardBrand = "kookmin"
	ChargePaymentMethodDetailsKrCardBrandKwangju     ChargePaymentMethodDetailsKrCardBrand = "kwangju"
	ChargePaymentMethodDetailsKrCardBrandLotte       ChargePaymentMethodDetailsKrCardBrand = "lotte"
	ChargePaymentMethodDetailsKrCardBrandMg          ChargePaymentMethodDetailsKrCardBrand = "mg"
	ChargePaymentMethodDetailsKrCardBrandNh          ChargePaymentMethodDetailsKrCardBrand = "nh"
	ChargePaymentMethodDetailsKrCardBrandPost        ChargePaymentMethodDetailsKrCardBrand = "post"
	ChargePaymentMethodDetailsKrCardBrandSamsung     ChargePaymentMethodDetailsKrCardBrand = "samsung"
	ChargePaymentMethodDetailsKrCardBrandSavingsbank ChargePaymentMethodDetailsKrCardBrand = "savingsbank"
	ChargePaymentMethodDetailsKrCardBrandShinhan     ChargePaymentMethodDetailsKrCardBrand = "shinhan"
	ChargePaymentMethodDetailsKrCardBrandShinhyup    ChargePaymentMethodDetailsKrCardBrand = "shinhyup"
	ChargePaymentMethodDetailsKrCardBrandSuhyup      ChargePaymentMethodDetailsKrCardBrand = "suhyup"
	ChargePaymentMethodDetailsKrCardBrandTossbank    ChargePaymentMethodDetailsKrCardBrand = "tossbank"
	ChargePaymentMethodDetailsKrCardBrandWoori       ChargePaymentMethodDetailsKrCardBrand = "woori"
)

// An array of conditions that are covered for the transaction, if applicable.
type ChargePaymentMethodDetailsPaypalSellerProtectionDisputeCategory string

// List of values that ChargePaymentMethodDetailsPaypalSellerProtectionDisputeCategory can take
const (
	ChargePaymentMethodDetailsPaypalSellerProtectionDisputeCategoryFraudulent         ChargePaymentMethodDetailsPaypalSellerProtectionDisputeCategory = "fraudulent"
	ChargePaymentMethodDetailsPaypalSellerProtectionDisputeCategoryProductNotReceived ChargePaymentMethodDetailsPaypalSellerProtectionDisputeCategory = "product_not_received"
)

// Indicates whether the transaction is eligible for PayPal's seller protection.
type ChargePaymentMethodDetailsPaypalSellerProtectionStatus string

// List of values that ChargePaymentMethodDetailsPaypalSellerProtectionStatus can take
const (
	ChargePaymentMethodDetailsPaypalSellerProtectionStatusEligible          ChargePaymentMethodDetailsPaypalSellerProtectionStatus = "eligible"
	ChargePaymentMethodDetailsPaypalSellerProtectionStatusNotEligible       ChargePaymentMethodDetailsPaypalSellerProtectionStatus = "not_eligible"
	ChargePaymentMethodDetailsPaypalSellerProtectionStatusPartiallyEligible ChargePaymentMethodDetailsPaypalSellerProtectionStatus = "partially_eligible"
)

// funding type of the underlying payment method.
type ChargePaymentMethodDetailsRevolutPayFundingType string

// List of values that ChargePaymentMethodDetailsRevolutPayFundingType can take
const (
	ChargePaymentMethodDetailsRevolutPayFundingTypeCard ChargePaymentMethodDetailsRevolutPayFundingType = "card"
)

// The [source_type](https://docs.stripe.com/api/balance/balance_object#balance_object-available-source_types) of the balance
type ChargePaymentMethodDetailsStripeBalanceSourceType string

// List of values that ChargePaymentMethodDetailsStripeBalanceSourceType can take
const (
	ChargePaymentMethodDetailsStripeBalanceSourceTypeBankAccount ChargePaymentMethodDetailsStripeBalanceSourceType = "bank_account"
	ChargePaymentMethodDetailsStripeBalanceSourceTypeCard        ChargePaymentMethodDetailsStripeBalanceSourceType = "card"
	ChargePaymentMethodDetailsStripeBalanceSourceTypeFPX         ChargePaymentMethodDetailsStripeBalanceSourceType = "fpx"
)

// The type of transaction-specific details of the payment method used in the payment. See [PaymentMethod.type](https://docs.stripe.com/api/payment_methods/object#payment_method_object-type) for the full list of possible types.
// An additional hash is included on `payment_method_details` with a name matching this value.
// It contains information specific to the payment method.
type ChargePaymentMethodDetailsType string

// List of values that ChargePaymentMethodDetailsType can take
const (
	ChargePaymentMethodDetailsTypeACHCreditTransfer ChargePaymentMethodDetailsType = "ach_credit_transfer"
	ChargePaymentMethodDetailsTypeACHDebit          ChargePaymentMethodDetailsType = "ach_debit"
	ChargePaymentMethodDetailsTypeACSSDebit         ChargePaymentMethodDetailsType = "acss_debit"
	ChargePaymentMethodDetailsTypeAlipay            ChargePaymentMethodDetailsType = "alipay"
	ChargePaymentMethodDetailsTypeAUBECSDebit       ChargePaymentMethodDetailsType = "au_becs_debit"
	ChargePaymentMethodDetailsTypeBACSDebit         ChargePaymentMethodDetailsType = "bacs_debit"
	ChargePaymentMethodDetailsTypeBancontact        ChargePaymentMethodDetailsType = "bancontact"
	ChargePaymentMethodDetailsTypeCard              ChargePaymentMethodDetailsType = "card"
	ChargePaymentMethodDetailsTypeCardPresent       ChargePaymentMethodDetailsType = "card_present"
	ChargePaymentMethodDetailsTypeEPS               ChargePaymentMethodDetailsType = "eps"
	ChargePaymentMethodDetailsTypeFPX               ChargePaymentMethodDetailsType = "fpx"
	ChargePaymentMethodDetailsTypeGiropay           ChargePaymentMethodDetailsType = "giropay"
	ChargePaymentMethodDetailsTypeGrabpay           ChargePaymentMethodDetailsType = "grabpay"
	ChargePaymentMethodDetailsTypeIDEAL             ChargePaymentMethodDetailsType = "ideal"
	ChargePaymentMethodDetailsTypeInteracPresent    ChargePaymentMethodDetailsType = "interac_present"
	ChargePaymentMethodDetailsTypeKlarna            ChargePaymentMethodDetailsType = "klarna"
	ChargePaymentMethodDetailsTypeMultibanco        ChargePaymentMethodDetailsType = "multibanco"
	ChargePaymentMethodDetailsTypeP24               ChargePaymentMethodDetailsType = "p24"
	ChargePaymentMethodDetailsTypeSEPADebit         ChargePaymentMethodDetailsType = "sepa_debit"
	ChargePaymentMethodDetailsTypeSofort            ChargePaymentMethodDetailsType = "sofort"
	ChargePaymentMethodDetailsTypeSwish             ChargePaymentMethodDetailsType = "swish"
	ChargePaymentMethodDetailsTypeStripeAccount     ChargePaymentMethodDetailsType = "stripe_account"
	ChargePaymentMethodDetailsTypeWeChat            ChargePaymentMethodDetailsType = "wechat"
)

// Account holder type: individual or company.
type ChargePaymentMethodDetailsUSBankAccountAccountHolderType string

// List of values that ChargePaymentMethodDetailsUSBankAccountAccountHolderType can take
const (
	ChargePaymentMethodDetailsUSBankAccountAccountHolderTypeCompany    ChargePaymentMethodDetailsUSBankAccountAccountHolderType = "company"
	ChargePaymentMethodDetailsUSBankAccountAccountHolderTypeIndividual ChargePaymentMethodDetailsUSBankAccountAccountHolderType = "individual"
)

// Account type: checkings or savings. Defaults to checking if omitted.
type ChargePaymentMethodDetailsUSBankAccountAccountType string

// List of values that ChargePaymentMethodDetailsUSBankAccountAccountType can take
const (
	ChargePaymentMethodDetailsUSBankAccountAccountTypeChecking ChargePaymentMethodDetailsUSBankAccountAccountType = "checking"
	ChargePaymentMethodDetailsUSBankAccountAccountTypeSavings  ChargePaymentMethodDetailsUSBankAccountAccountType = "savings"
)

// The status of the payment is either `succeeded`, `pending`, or `failed`.
type ChargeStatus string

// List of values that ChargeStatus can take
const (
	ChargeStatusFailed    ChargeStatus = "failed"
	ChargeStatusPending   ChargeStatus = "pending"
	ChargeStatusSucceeded ChargeStatus = "succeeded"
)

// Returns a list of charges you've previously created. The charges are returned in sorted order, with the most recent charges appearing first.
type ChargeListParams struct {
	ListParams `form:"*"`
	// Only return charges that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return charges that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return charges for the customer specified by this customer ID.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return charges that were created by the PaymentIntent specified by this PaymentIntent ID.
	PaymentIntent *string `form:"payment_intent"`
	// Only return charges for this transfer group, limited to 100.
	TransferGroup *string `form:"transfer_group"`
}

// AddExpand appends a new field to expand.
func (p *ChargeListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type ChargeDestinationParams struct {
	// ID of an existing, connected Stripe account.
	Account *string `form:"account"`
	// The amount to transfer to the destination account without creating an `Application Fee` object. Cannot be combined with the `application_fee` parameter. Must be less than or equal to the charge amount.
	Amount *int64 `form:"amount"`
}

// Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.
type ChargeRadarOptionsParams struct {
	// A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.
type ChargeTransferDataParams struct {
	// The amount transferred to the destination account, if specified. By default, the entire charge amount is transferred to the destination account.
	Amount *int64 `form:"amount"`
	// This parameter can only be used on Charge creation.
	// ID of an existing, connected Stripe account.
	Destination *string `form:"destination"`
}
type ChargeLevel3LineItemParams struct {
	DiscountAmount     *int64  `form:"discount_amount"`
	ProductCode        *string `form:"product_code"`
	ProductDescription *string `form:"product_description"`
	Quantity           *int64  `form:"quantity"`
	TaxAmount          *int64  `form:"tax_amount"`
	UnitCost           *int64  `form:"unit_cost"`
}
type ChargeLevel3Params struct {
	CustomerReference  *string                       `form:"customer_reference"`
	LineItems          []*ChargeLevel3LineItemParams `form:"line_items"`
	MerchantReference  *string                       `form:"merchant_reference"`
	ShippingAddressZip *string                       `form:"shipping_address_zip"`
	ShippingAmount     *int64                        `form:"shipping_amount"`
	ShippingFromZip    *string                       `form:"shipping_from_zip"`
}

// This method is no longer recommended—use the [Payment Intents API](https://docs.stripe.com/docs/api/payment_intents)
// to initiate a new payment instead. Confirmation of the PaymentIntent creates the Charge
// object used to request payment.
type ChargeParams struct {
	Params `form:"*"`
	// Amount intended to be collected by this payment. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://docs.stripe.com/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount         *int64 `form:"amount"`
	ApplicationFee *int64 `form:"application_fee"`
	// A fee in cents (or local equivalent) that will be applied to the charge and transferred to the application owner's Stripe account. The request must be made with an OAuth key or the `Stripe-Account` header in order to take an application fee. For more information, see the application fees [documentation](https://docs.stripe.com/connect/direct-charges#collect-fees).
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// Whether to immediately capture the charge. Defaults to `true`. When `false`, the charge issues an authorization (or pre-authorization), and will need to be [captured](https://api.stripe.com#capture_charge) later. Uncaptured charges expire after a set number of days (7 by default). For more information, see the [authorizing charges and settling later](https://docs.stripe.com/charges/placing-a-hold) documentation.
	Capture *bool `form:"capture"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of an existing customer that will be associated with this request. This field may only be updated if there is no existing associated customer with this charge.
	Customer *string `form:"customer"`
	// An arbitrary string which you can attach to a `Charge` object. It is displayed when in the web interface alongside the charge. Note that if you use Stripe to send automatic email receipts to your customers, your receipt emails will include the `description` of the charge(s) that they are describing.
	Description  *string                  `form:"description"`
	Destination  *ChargeDestinationParams `form:"destination"`
	ExchangeRate *float64                 `form:"exchange_rate"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A set of key-value pairs you can attach to a charge giving information about its riskiness. If you believe a charge is fraudulent, include a `user_report` key with a value of `fraudulent`. If you believe a charge is safe, include a `user_report` key with a value of `safe`. Stripe will use the information you send to improve our fraud detection algorithms.
	FraudDetails *ChargeFraudDetailsParams `form:"fraud_details"`
	Level3       *ChargeLevel3Params       `form:"level3"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The Stripe account ID for which these funds are intended. You can specify the business of record as the connected account using the `on_behalf_of` attribute on the charge. For details, see [Creating Separate Charges and Transfers](https://docs.stripe.com/connect/separate-charges-and-transfers#settlement-merchant).
	OnBehalfOf *string `form:"on_behalf_of"`
	// Provides industry-specific information about the charge.
	PaymentDetails *ChargePaymentDetailsParams `form:"payment_details"`
	// Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.
	RadarOptions *ChargeRadarOptionsParams `form:"radar_options"`
	// The email address to which this charge's [receipt](https://docs.stripe.com/dashboard/receipts) will be sent. The receipt will not be sent until the charge is paid, and no receipts will be sent for test mode charges. If this charge is for a [Customer](https://docs.stripe.com/api/customers/object), the email address specified here will override the customer's email address. If `receipt_email` is specified for a charge in live mode, a receipt will be sent regardless of your [email settings](https://dashboard.stripe.com/account/emails).
	ReceiptEmail *string `form:"receipt_email"`
	// Shipping information for the charge. Helps prevent fraud on charges for physical goods.
	Shipping *ShippingDetailsParams     `form:"shipping"`
	Source   *PaymentSourceSourceParams `form:"*"` // PaymentSourceSourceParams has custom encoding so brought to top level with "*"
	// For a non-card charge, text that appears on the customer's statement as the statement descriptor. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	//
	// For a card charge, this value is ignored unless you don't specify a `statement_descriptor_suffix`, in which case this value is used as the suffix.
	StatementDescriptor *string `form:"statement_descriptor"`
	// Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement. If the account has no prefix value, the suffix is concatenated to the account's statement descriptor.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix"`
	// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.
	TransferData *ChargeTransferDataParams `form:"transfer_data"`
	// A string that identifies this transaction as part of a group. `transfer_group` may only be provided if it has not been set. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options) for details.
	TransferGroup *string `form:"transfer_group"`
}

// SetSource adds valid sources to a ChargeParams object,
// returning an error for unsupported sources.
func (p *ChargeParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// AddExpand appends a new field to expand.
func (p *ChargeParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *ChargeParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A set of key-value pairs you can attach to a charge giving information about its riskiness. If you believe a charge is fraudulent, include a `user_report` key with a value of `fraudulent`. If you believe a charge is safe, include a `user_report` key with a value of `safe`. Stripe will use the information you send to improve our fraud detection algorithms.
type ChargeFraudDetailsParams struct {
	// Either `safe` or `fraudulent`.
	UserReport *string `form:"user_report"`
}

// Affiliate details for this purchase.
type ChargePaymentDetailsCarRentalAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargePaymentDetailsCarRentalDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargePaymentDetailsCarRentalDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargePaymentDetailsCarRentalDeliveryRecipientParams `form:"recipient"`
}

// The details of the distance traveled during the rental period.
type ChargePaymentDetailsCarRentalDistanceParams struct {
	// Distance traveled.
	Amount *int64 `form:"amount"`
	// Unit of measurement for the distance traveled. One of `miles` or `kilometers`.
	Unit *string `form:"unit"`
}

// The details of the passengers in the travel reservation
type ChargePaymentDetailsCarRentalDriverParams struct {
	// Driver's identification number.
	DriverIdentificationNumber *string `form:"driver_identification_number"`
	// Driver's tax number.
	DriverTaxNumber *string `form:"driver_tax_number"`
	// Full name of the person or entity on the car reservation.
	Name *string `form:"name"`
}

// Car rental details for this PaymentIntent.
type ChargePaymentDetailsCarRentalParams struct {
	// Affiliate details for this purchase.
	Affiliate *ChargePaymentDetailsCarRentalAffiliateParams `form:"affiliate"`
	// The booking number associated with the car rental.
	BookingNumber *string `form:"booking_number"`
	// Class code of the car.
	CarClassCode *string `form:"car_class_code"`
	// Make of the car.
	CarMake *string `form:"car_make"`
	// Model of the car.
	CarModel *string `form:"car_model"`
	// The name of the rental car company.
	Company *string `form:"company"`
	// The customer service phone number of the car rental company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// Number of days the car is being rented.
	DaysRented *int64 `form:"days_rented"`
	// Delivery details for this purchase.
	Delivery *ChargePaymentDetailsCarRentalDeliveryParams `form:"delivery"`
	// The details of the distance traveled during the rental period.
	Distance *ChargePaymentDetailsCarRentalDistanceParams `form:"distance"`
	// The details of the passengers in the travel reservation
	Drivers []*ChargePaymentDetailsCarRentalDriverParams `form:"drivers"`
	// List of additional charges being billed.
	ExtraCharges []*string `form:"extra_charges"`
	// Indicates if the customer did not keep nor cancel their booking.
	NoShow *bool `form:"no_show"`
	// Car pick-up address.
	PickupAddress *AddressParams `form:"pickup_address"`
	// Car pick-up time. Measured in seconds since the Unix epoch.
	PickupAt *int64 `form:"pickup_at"`
	// Name of the pickup location.
	PickupLocationName *string `form:"pickup_location_name"`
	// Rental rate.
	RateAmount *int64 `form:"rate_amount"`
	// The frequency at which the rate amount is applied. One of `day`, `week` or `month`
	RateInterval *string `form:"rate_interval"`
	// The name of the person or entity renting the car.
	RenterName *string `form:"renter_name"`
	// Car return address.
	ReturnAddress *AddressParams `form:"return_address"`
	// Car return time. Measured in seconds since the Unix epoch.
	ReturnAt *int64 `form:"return_at"`
	// Name of the return location.
	ReturnLocationName *string `form:"return_location_name"`
	// Indicates whether the goods or services are tax-exempt or tax is not collected.
	TaxExempt *bool `form:"tax_exempt"`
	// The vehicle identification number.
	VehicleIdentificationNumber *string `form:"vehicle_identification_number"`
}

// Affiliate (such as travel agency) details for the rental.
type ChargePaymentDetailsCarRentalDatumAffiliateParams struct {
	// Affiliate partner code.
	Code *string `form:"code"`
	// Name of affiliate partner.
	Name *string `form:"name"`
}

// Distance details for the rental.
type ChargePaymentDetailsCarRentalDatumDistanceParams struct {
	// Distance traveled.
	Amount *int64 `form:"amount"`
	// Unit of measurement for the distance traveled. One of `miles` or `kilometers`.
	Unit *string `form:"unit"`
}

// Driver's date of birth.
type ChargePaymentDetailsCarRentalDatumDriverDateOfBirthParams struct {
	// Day of birth (1-31).
	Day *int64 `form:"day"`
	// Month of birth (1-12).
	Month *int64 `form:"month"`
	// Year of birth (must be greater than 1900).
	Year *int64 `form:"year"`
}

// List of drivers for the rental.
type ChargePaymentDetailsCarRentalDatumDriverParams struct {
	// Driver's date of birth.
	DateOfBirth *ChargePaymentDetailsCarRentalDatumDriverDateOfBirthParams `form:"date_of_birth"`
	// Driver's identification number.
	DriverIdentificationNumber *string `form:"driver_identification_number"`
	// Driver's tax number.
	DriverTaxNumber *string `form:"driver_tax_number"`
	// Driver's full name.
	Name *string `form:"name"`
}

// Drop-off location details.
type ChargePaymentDetailsCarRentalDatumDropOffParams struct {
	// Address of the rental location.
	Address *AddressParams `form:"address"`
	// Location name.
	LocationName *string `form:"location_name"`
	// Timestamp for the location.
	Time *int64 `form:"time"`
}

// Insurance details for the rental.
type ChargePaymentDetailsCarRentalDatumInsuranceParams struct {
	// Amount of the insurance coverage in cents.
	Amount *int64 `form:"amount"`
	// Currency of the insurance amount.
	Currency *string `form:"currency"`
	// Name of the insurance company.
	InsuranceCompanyName *string `form:"insurance_company_name"`
	// Type of insurance coverage.
	InsuranceType *string `form:"insurance_type"`
}

// Pickup location details.
type ChargePaymentDetailsCarRentalDatumPickupParams struct {
	// Address of the rental location.
	Address *AddressParams `form:"address"`
	// Location name.
	LocationName *string `form:"location_name"`
	// Timestamp for the location.
	Time *int64 `form:"time"`
}

// Discount details for the rental.
type ChargePaymentDetailsCarRentalDatumTotalDiscountsParams struct {
	// Corporate client discount code.
	CorporateClientCode *string `form:"corporate_client_code"`
	// Coupon code applied to the rental.
	Coupon *string `form:"coupon"`
	// Maximum number of free miles or kilometers included.
	MaximumFreeMilesOrKilometers *int64 `form:"maximum_free_miles_or_kilometers"`
}

// Additional charges for the rental.
type ChargePaymentDetailsCarRentalDatumTotalExtraChargeParams struct {
	// Amount of the extra charge in cents.
	Amount *int64 `form:"amount"`
	// Type of extra charge.
	Type *string `form:"type"`
}

// Array of tax details.
type ChargePaymentDetailsCarRentalDatumTotalTaxTaxParams struct {
	// Tax amount.
	Amount *int64 `form:"amount"`
	// Tax rate applied.
	Rate *int64 `form:"rate"`
	// Type of tax applied.
	Type *string `form:"type"`
}

// Tax breakdown for the rental.
type ChargePaymentDetailsCarRentalDatumTotalTaxParams struct {
	// Array of tax details.
	Taxes []*ChargePaymentDetailsCarRentalDatumTotalTaxTaxParams `form:"taxes"`
	// Indicates if the transaction is tax exempt.
	TaxExemptIndicator *bool `form:"tax_exempt_indicator"`
}

// Total cost breakdown for the rental.
type ChargePaymentDetailsCarRentalDatumTotalParams struct {
	// Total amount in cents.
	Amount *int64 `form:"amount"`
	// Currency of the amount.
	Currency *string `form:"currency"`
	// Discount details for the rental.
	Discounts *ChargePaymentDetailsCarRentalDatumTotalDiscountsParams `form:"discounts"`
	// Additional charges for the rental.
	ExtraCharges []*ChargePaymentDetailsCarRentalDatumTotalExtraChargeParams `form:"extra_charges"`
	// Rate per unit for the rental.
	RatePerUnit *int64 `form:"rate_per_unit"`
	// Unit of measurement for the rate.
	RateUnit *string `form:"rate_unit"`
	// Tax breakdown for the rental.
	Tax *ChargePaymentDetailsCarRentalDatumTotalTaxParams `form:"tax"`
}

// Vehicle details for the rental.
type ChargePaymentDetailsCarRentalDatumVehicleParams struct {
	// Make of the rental vehicle.
	Make *string `form:"make"`
	// Model of the rental vehicle.
	Model *string `form:"model"`
	// Odometer reading at the time of rental.
	Odometer *int64 `form:"odometer"`
	// Type of the rental vehicle.
	Type *string `form:"type"`
	// Class of the rental vehicle.
	VehicleClass *string `form:"vehicle_class"`
	// Vehicle identification number (VIN).
	VehicleIdentificationNumber *string `form:"vehicle_identification_number"`
}

// Car rental data for this PaymentIntent.
type ChargePaymentDetailsCarRentalDatumParams struct {
	// Affiliate (such as travel agency) details for the rental.
	Affiliate *ChargePaymentDetailsCarRentalDatumAffiliateParams `form:"affiliate"`
	// Booking confirmation number for the car rental.
	BookingNumber *string `form:"booking_number"`
	// Name of the car rental company.
	CarrierName *string `form:"carrier_name"`
	// Customer service phone number for the car rental company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// Number of days the car is being rented.
	DaysRented *int64 `form:"days_rented"`
	// Distance details for the rental.
	Distance *ChargePaymentDetailsCarRentalDatumDistanceParams `form:"distance"`
	// List of drivers for the rental.
	Drivers []*ChargePaymentDetailsCarRentalDatumDriverParams `form:"drivers"`
	// Drop-off location details.
	DropOff *ChargePaymentDetailsCarRentalDatumDropOffParams `form:"drop_off"`
	// Insurance details for the rental.
	Insurances []*ChargePaymentDetailsCarRentalDatumInsuranceParams `form:"insurances"`
	// Indicates if the customer was a no-show.
	NoShowIndicator *bool `form:"no_show_indicator"`
	// Pickup location details.
	Pickup *ChargePaymentDetailsCarRentalDatumPickupParams `form:"pickup"`
	// Name of the person renting the vehicle.
	RenterName *string `form:"renter_name"`
	// Total cost breakdown for the rental.
	Total *ChargePaymentDetailsCarRentalDatumTotalParams `form:"total"`
	// Vehicle details for the rental.
	Vehicle *ChargePaymentDetailsCarRentalDatumVehicleParams `form:"vehicle"`
}

// Affiliate details for this purchase.
type ChargePaymentDetailsEventDetailsAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargePaymentDetailsEventDetailsDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargePaymentDetailsEventDetailsDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargePaymentDetailsEventDetailsDeliveryRecipientParams `form:"recipient"`
}

// Event details for this PaymentIntent
type ChargePaymentDetailsEventDetailsParams struct {
	// Indicates if the tickets are digitally checked when entering the venue.
	AccessControlledVenue *bool `form:"access_controlled_venue"`
	// The event location's address.
	Address *AddressParams `form:"address"`
	// Affiliate details for this purchase.
	Affiliate *ChargePaymentDetailsEventDetailsAffiliateParams `form:"affiliate"`
	// The name of the company
	Company *string `form:"company"`
	// Delivery details for this purchase.
	Delivery *ChargePaymentDetailsEventDetailsDeliveryParams `form:"delivery"`
	// Event end time. Measured in seconds since the Unix epoch.
	EndsAt *int64 `form:"ends_at"`
	// Type of the event entertainment (concert, sports event etc)
	Genre *string `form:"genre"`
	// The name of the event.
	Name *string `form:"name"`
	// Event start time. Measured in seconds since the Unix epoch.
	StartsAt *int64 `form:"starts_at"`
}

// Affiliate details for this purchase.
type ChargePaymentDetailsFlightAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargePaymentDetailsFlightDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargePaymentDetailsFlightDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargePaymentDetailsFlightDeliveryRecipientParams `form:"recipient"`
}

// The details of the passengers in the travel reservation.
type ChargePaymentDetailsFlightPassengerParams struct {
	// Full name of the person or entity on the flight reservation.
	Name *string `form:"name"`
}

// The individual flight segments associated with the trip.
type ChargePaymentDetailsFlightSegmentParams struct {
	// The flight segment amount.
	Amount *int64 `form:"amount"`
	// The International Air Transport Association (IATA) airport code for the arrival airport.
	ArrivalAirport *string `form:"arrival_airport"`
	// The arrival time for the flight segment. Measured in seconds since the Unix epoch.
	ArrivesAt *int64 `form:"arrives_at"`
	// The International Air Transport Association (IATA) carrier code of the carrier operating the flight segment.
	Carrier *string `form:"carrier"`
	// The departure time for the flight segment. Measured in seconds since the Unix epoch.
	DepartsAt *int64 `form:"departs_at"`
	// The International Air Transport Association (IATA) airport code for the departure airport.
	DepartureAirport *string `form:"departure_airport"`
	// The flight number associated with the segment
	FlightNumber *string `form:"flight_number"`
	// The fare class for the segment.
	ServiceClass *string `form:"service_class"`
}

// Flight reservation details for this PaymentIntent
type ChargePaymentDetailsFlightParams struct {
	// Affiliate details for this purchase.
	Affiliate *ChargePaymentDetailsFlightAffiliateParams `form:"affiliate"`
	// The agency number (i.e. International Air Transport Association (IATA) agency number) of the travel agency that made the booking.
	AgencyNumber *string `form:"agency_number"`
	// The International Air Transport Association (IATA) carrier code of the carrier that issued the ticket.
	Carrier *string `form:"carrier"`
	// Delivery details for this purchase.
	Delivery *ChargePaymentDetailsFlightDeliveryParams `form:"delivery"`
	// The name of the person or entity on the reservation.
	PassengerName *string `form:"passenger_name"`
	// The details of the passengers in the travel reservation.
	Passengers []*ChargePaymentDetailsFlightPassengerParams `form:"passengers"`
	// The individual flight segments associated with the trip.
	Segments []*ChargePaymentDetailsFlightSegmentParams `form:"segments"`
	// The ticket number associated with the travel reservation.
	TicketNumber *string `form:"ticket_number"`
}

// Affiliate details if applicable.
type ChargePaymentDetailsFlightDatumAffiliateParams struct {
	// Affiliate partner code.
	Code *string `form:"code"`
	// Name of affiliate partner.
	Name *string `form:"name"`
	// Code provided by the company to a travel agent authorizing ticket issuance.
	TravelAuthorizationCode *string `form:"travel_authorization_code"`
}

// List of insurances.
type ChargePaymentDetailsFlightDatumInsuranceParams struct {
	// Insurance cost.
	Amount *int64 `form:"amount"`
	// Insurance currency.
	Currency *string `form:"currency"`
	// Insurance company name.
	InsuranceCompanyName *string `form:"insurance_company_name"`
	// Type of insurance.
	InsuranceType *string `form:"insurance_type"`
}

// List of passengers.
type ChargePaymentDetailsFlightDatumPassengerParams struct {
	// Passenger's full name.
	Name *string `form:"name"`
}

// Arrival details.
type ChargePaymentDetailsFlightDatumSegmentArrivalParams struct {
	// Arrival airport IATA code.
	Airport *string `form:"airport"`
	// Arrival date/time.
	ArrivesAt *int64 `form:"arrives_at"`
	// Arrival city.
	City *string `form:"city"`
	// Arrival country.
	Country *string `form:"country"`
}

// Departure details.
type ChargePaymentDetailsFlightDatumSegmentDepartureParams struct {
	// Departure airport IATA code.
	Airport *string `form:"airport"`
	// Departure city.
	City *string `form:"city"`
	// Departure country.
	Country *string `form:"country"`
	// Departure date/time.
	DepartsAt *int64 `form:"departs_at"`
}

// List of flight segments.
type ChargePaymentDetailsFlightDatumSegmentParams struct {
	// Segment fare amount.
	Amount *int64 `form:"amount"`
	// Arrival details.
	Arrival *ChargePaymentDetailsFlightDatumSegmentArrivalParams `form:"arrival"`
	// Airline carrier code.
	CarrierCode *string `form:"carrier_code"`
	// Carrier name.
	CarrierName *string `form:"carrier_name"`
	// Segment currency.
	Currency *string `form:"currency"`
	// Departure details.
	Departure *ChargePaymentDetailsFlightDatumSegmentDepartureParams `form:"departure"`
	// Exchange ticket number.
	ExchangeTicketNumber *string `form:"exchange_ticket_number"`
	// Fare basis code.
	FareBasisCode *string `form:"fare_basis_code"`
	// Additional fees.
	Fees *int64 `form:"fees"`
	// Flight number.
	FlightNumber *string `form:"flight_number"`
	// Stopover indicator.
	IsStopOverIndicator *bool `form:"is_stop_over_indicator"`
	// Refundable ticket indicator.
	Refundable *bool `form:"refundable"`
	// Class of service.
	ServiceClass *string `form:"service_class"`
	// Tax amount for segment.
	TaxAmount *int64 `form:"tax_amount"`
	// Ticket number.
	TicketNumber *string `form:"ticket_number"`
}

// Discount details.
type ChargePaymentDetailsFlightDatumTotalDiscountsParams struct {
	// Corporate client discount code.
	CorporateClientCode *string `form:"corporate_client_code"`
}

// Additional charges.
type ChargePaymentDetailsFlightDatumTotalExtraChargeParams struct {
	// Amount of additional charges.
	Amount *int64 `form:"amount"`
	// Type of additional charges.
	Type *string `form:"type"`
}

// Array of tax details.
type ChargePaymentDetailsFlightDatumTotalTaxTaxParams struct {
	// Tax amount.
	Amount *int64 `form:"amount"`
	// Tax rate.
	Rate *int64 `form:"rate"`
	// Type of tax.
	Type *string `form:"type"`
}

// Tax breakdown.
type ChargePaymentDetailsFlightDatumTotalTaxParams struct {
	// Array of tax details.
	Taxes []*ChargePaymentDetailsFlightDatumTotalTaxTaxParams `form:"taxes"`
}

// Total cost breakdown.
type ChargePaymentDetailsFlightDatumTotalParams struct {
	// Total flight amount.
	Amount *int64 `form:"amount"`
	// Reason for credit.
	CreditReason *string `form:"credit_reason"`
	// Total currency.
	Currency *string `form:"currency"`
	// Discount details.
	Discounts *ChargePaymentDetailsFlightDatumTotalDiscountsParams `form:"discounts"`
	// Additional charges.
	ExtraCharges []*ChargePaymentDetailsFlightDatumTotalExtraChargeParams `form:"extra_charges"`
	// Tax breakdown.
	Tax *ChargePaymentDetailsFlightDatumTotalTaxParams `form:"tax"`
}

// Flight data for this PaymentIntent.
type ChargePaymentDetailsFlightDatumParams struct {
	// Affiliate details if applicable.
	Affiliate *ChargePaymentDetailsFlightDatumAffiliateParams `form:"affiliate"`
	// Reservation reference.
	BookingNumber *string `form:"booking_number"`
	// Computerized reservation system used to make the reservation and purchase the ticket.
	ComputerizedReservationSystem *string `form:"computerized_reservation_system"`
	// Ticket restrictions.
	EndorsementsAndRestrictions *string `form:"endorsements_and_restrictions"`
	// List of insurances.
	Insurances []*ChargePaymentDetailsFlightDatumInsuranceParams `form:"insurances"`
	// List of passengers.
	Passengers []*ChargePaymentDetailsFlightDatumPassengerParams `form:"passengers"`
	// List of flight segments.
	Segments []*ChargePaymentDetailsFlightDatumSegmentParams `form:"segments"`
	// Electronic ticket indicator.
	TicketElectronicallyIssuedIndicator *bool `form:"ticket_electronically_issued_indicator"`
	// Total cost breakdown.
	Total *ChargePaymentDetailsFlightDatumTotalParams `form:"total"`
	// Type of flight transaction.
	TransactionType *string `form:"transaction_type"`
}

// Affiliate details for this purchase.
type ChargePaymentDetailsLodgingAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargePaymentDetailsLodgingDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargePaymentDetailsLodgingDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargePaymentDetailsLodgingDeliveryRecipientParams `form:"recipient"`
}

// The details of the passengers in the travel reservation
type ChargePaymentDetailsLodgingPassengerParams struct {
	// Full name of the person or entity on the lodging reservation.
	Name *string `form:"name"`
}

// Lodging reservation details for this PaymentIntent
type ChargePaymentDetailsLodgingParams struct {
	// The lodging location's address.
	Address *AddressParams `form:"address"`
	// The number of adults on the booking
	Adults *int64 `form:"adults"`
	// Affiliate details for this purchase.
	Affiliate *ChargePaymentDetailsLodgingAffiliateParams `form:"affiliate"`
	// The booking number associated with the lodging reservation.
	BookingNumber *string `form:"booking_number"`
	// The lodging category
	Category *string `form:"category"`
	// Lodging check-in time. Measured in seconds since the Unix epoch.
	CheckinAt *int64 `form:"checkin_at"`
	// Lodging check-out time. Measured in seconds since the Unix epoch.
	CheckoutAt *int64 `form:"checkout_at"`
	// The customer service phone number of the lodging company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// The daily lodging room rate.
	DailyRoomRateAmount *int64 `form:"daily_room_rate_amount"`
	// Delivery details for this purchase.
	Delivery *ChargePaymentDetailsLodgingDeliveryParams `form:"delivery"`
	// List of additional charges being billed.
	ExtraCharges []*string `form:"extra_charges"`
	// Indicates whether the lodging location is compliant with the Fire Safety Act.
	FireSafetyActCompliance *bool `form:"fire_safety_act_compliance"`
	// The name of the lodging location.
	Name *string `form:"name"`
	// Indicates if the customer did not keep their booking while failing to cancel the reservation.
	NoShow *bool `form:"no_show"`
	// The number of rooms on the booking
	NumberOfRooms *int64 `form:"number_of_rooms"`
	// The details of the passengers in the travel reservation
	Passengers []*ChargePaymentDetailsLodgingPassengerParams `form:"passengers"`
	// The phone number of the lodging location.
	PropertyPhoneNumber *string `form:"property_phone_number"`
	// The room class for this purchase.
	RoomClass *string `form:"room_class"`
	// The number of room nights
	RoomNights *int64 `form:"room_nights"`
	// The total tax amount associating with the room reservation.
	TotalRoomTaxAmount *int64 `form:"total_room_tax_amount"`
	// The total tax amount
	TotalTaxAmount *int64 `form:"total_tax_amount"`
}

// Accommodation details for the lodging.
type ChargePaymentDetailsLodgingDatumAccommodationParams struct {
	// Type of accommodation.
	AccommodationType *string `form:"accommodation_type"`
	// Bed type.
	BedType *string `form:"bed_type"`
	// Daily accommodation rate in cents.
	DailyRateAmount *int64 `form:"daily_rate_amount"`
	// Number of nights.
	Nights *int64 `form:"nights"`
	// Number of rooms, cabanas, apartments, and so on.
	NumberOfRooms *int64 `form:"number_of_rooms"`
	// Rate type.
	RateType *string `form:"rate_type"`
	// Whether smoking is allowed.
	SmokingIndicator *bool `form:"smoking_indicator"`
}

// Affiliate details if applicable.
type ChargePaymentDetailsLodgingDatumAffiliateParams struct {
	// Affiliate partner code.
	Code *string `form:"code"`
	// Affiliate partner name.
	Name *string `form:"name"`
}

// List of guests for the lodging.
type ChargePaymentDetailsLodgingDatumGuestParams struct {
	// Guest's full name.
	Name *string `form:"name"`
}

// Host details for the lodging.
type ChargePaymentDetailsLodgingDatumHostParams struct {
	// Address of the host.
	Address *AddressParams `form:"address"`
	// Host's country of domicile.
	CountryOfDomicile *string `form:"country_of_domicile"`
	// Reference number for the host.
	HostReference *string `form:"host_reference"`
	// Type of host.
	HostType *string `form:"host_type"`
	// Name of the lodging property or host.
	Name *string `form:"name"`
	// Total number of reservations for the host.
	NumberOfReservations *int64 `form:"number_of_reservations"`
	// Property phone number.
	PropertyPhoneNumber *string `form:"property_phone_number"`
	// Host's registration date.
	RegisteredAt *int64 `form:"registered_at"`
}

// List of insurances for the lodging.
type ChargePaymentDetailsLodgingDatumInsuranceParams struct {
	// Price of the insurance coverage in cents.
	Amount *int64 `form:"amount"`
	// Currency of the insurance amount.
	Currency *string `form:"currency"`
	// Name of the insurance company.
	InsuranceCompanyName *string `form:"insurance_company_name"`
	// Type of insurance coverage.
	InsuranceType *string `form:"insurance_type"`
}

// Discount details for the lodging.
type ChargePaymentDetailsLodgingDatumTotalDiscountsParams struct {
	// Corporate client discount code.
	CorporateClientCode *string `form:"corporate_client_code"`
	// Coupon code.
	Coupon *string `form:"coupon"`
}

// Additional charges for the lodging.
type ChargePaymentDetailsLodgingDatumTotalExtraChargeParams struct {
	// Amount of the extra charge in cents.
	Amount *int64 `form:"amount"`
	// Type of extra charge.
	Type *string `form:"type"`
}

// Tax details.
type ChargePaymentDetailsLodgingDatumTotalTaxTaxParams struct {
	// Tax amount in cents.
	Amount *int64 `form:"amount"`
	// Tax rate.
	Rate *int64 `form:"rate"`
	// Type of tax applied.
	Type *string `form:"type"`
}

// Tax breakdown for the lodging reservation.
type ChargePaymentDetailsLodgingDatumTotalTaxParams struct {
	// Tax details.
	Taxes []*ChargePaymentDetailsLodgingDatumTotalTaxTaxParams `form:"taxes"`
	// Indicates whether the transaction is tax exempt.
	TaxExemptIndicator *bool `form:"tax_exempt_indicator"`
}

// Total details for the lodging.
type ChargePaymentDetailsLodgingDatumTotalParams struct {
	// Total price of the lodging reservation in cents.
	Amount *int64 `form:"amount"`
	// Cash advances in cents.
	CashAdvances *int64 `form:"cash_advances"`
	// Currency of the total amount.
	Currency *string `form:"currency"`
	// Discount details for the lodging.
	Discounts *ChargePaymentDetailsLodgingDatumTotalDiscountsParams `form:"discounts"`
	// Additional charges for the lodging.
	ExtraCharges []*ChargePaymentDetailsLodgingDatumTotalExtraChargeParams `form:"extra_charges"`
	// Prepaid amount in cents.
	PrepaidAmount *int64 `form:"prepaid_amount"`
	// Tax breakdown for the lodging reservation.
	Tax *ChargePaymentDetailsLodgingDatumTotalTaxParams `form:"tax"`
}

// Lodging data for this PaymentIntent.
type ChargePaymentDetailsLodgingDatumParams struct {
	// Accommodation details for the lodging.
	Accommodation *ChargePaymentDetailsLodgingDatumAccommodationParams `form:"accommodation"`
	// Affiliate details if applicable.
	Affiliate *ChargePaymentDetailsLodgingDatumAffiliateParams `form:"affiliate"`
	// Booking confirmation number for the lodging.
	BookingNumber *string `form:"booking_number"`
	// Check-in date.
	CheckinAt *int64 `form:"checkin_at"`
	// Check-out date.
	CheckoutAt *int64 `form:"checkout_at"`
	// Customer service phone number for the lodging company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// Whether the lodging is compliant with any hotel fire safety regulations.
	FireSafetyActComplianceIndicator *bool `form:"fire_safety_act_compliance_indicator"`
	// List of guests for the lodging.
	Guests []*ChargePaymentDetailsLodgingDatumGuestParams `form:"guests"`
	// Host details for the lodging.
	Host *ChargePaymentDetailsLodgingDatumHostParams `form:"host"`
	// List of insurances for the lodging.
	Insurances []*ChargePaymentDetailsLodgingDatumInsuranceParams `form:"insurances"`
	// Whether the renter is a no-show.
	NoShowIndicator *bool `form:"no_show_indicator"`
	// Renter ID number for the lodging.
	RenterIDNumber *string `form:"renter_id_number"`
	// Renter name for the lodging.
	RenterName *string `form:"renter_name"`
	// Total details for the lodging.
	Total *ChargePaymentDetailsLodgingDatumTotalParams `form:"total"`
}

// Affiliate details for this purchase.
type ChargePaymentDetailsSubscriptionAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Subscription billing details for this purchase.
type ChargePaymentDetailsSubscriptionBillingIntervalParams struct {
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	Count *int64 `form:"count"`
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
}

// Subscription details for this PaymentIntent
type ChargePaymentDetailsSubscriptionParams struct {
	// Affiliate details for this purchase.
	Affiliate *ChargePaymentDetailsSubscriptionAffiliateParams `form:"affiliate"`
	// Info whether the subscription will be auto renewed upon expiry.
	AutoRenewal *bool `form:"auto_renewal"`
	// Subscription billing details for this purchase.
	BillingInterval *ChargePaymentDetailsSubscriptionBillingIntervalParams `form:"billing_interval"`
	// Subscription end time. Measured in seconds since the Unix epoch.
	EndsAt *int64 `form:"ends_at"`
	// Name of the product on subscription. e.g. Apple Music Subscription
	Name *string `form:"name"`
	// Subscription start time. Measured in seconds since the Unix epoch.
	StartsAt *int64 `form:"starts_at"`
}

// Provides industry-specific information about the charge.
type ChargePaymentDetailsParams struct {
	// Car rental details for this PaymentIntent.
	CarRental *ChargePaymentDetailsCarRentalParams `form:"car_rental"`
	// Car rental data for this PaymentIntent.
	CarRentalData []*ChargePaymentDetailsCarRentalDatumParams `form:"car_rental_data"`
	// A unique value to identify the customer. This field is available only for card payments.
	//
	// This field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.
	CustomerReference *string `form:"customer_reference"`
	// Event details for this PaymentIntent
	EventDetails *ChargePaymentDetailsEventDetailsParams `form:"event_details"`
	// Flight reservation details for this PaymentIntent
	Flight *ChargePaymentDetailsFlightParams `form:"flight"`
	// Flight data for this PaymentIntent.
	FlightData []*ChargePaymentDetailsFlightDatumParams `form:"flight_data"`
	// Lodging reservation details for this PaymentIntent
	Lodging *ChargePaymentDetailsLodgingParams `form:"lodging"`
	// Lodging data for this PaymentIntent.
	LodgingData []*ChargePaymentDetailsLodgingDatumParams `form:"lodging_data"`
	// A unique value assigned by the business to identify the transaction. Required for L2 and L3 rates.
	//
	// Required when the Payment Method Types array contains `card`, including when [automatic_payment_methods.enabled](https://docs.stripe.com/api/payment_intents/create#create_payment_intent-automatic_payment_methods-enabled) is set to `true`.
	//
	// For Cards, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks. For Klarna, this field is truncated to 255 characters and is visible to customers when they view the order in the Klarna app.
	OrderReference *string `form:"order_reference"`
	// Subscription details for this PaymentIntent
	Subscription *ChargePaymentDetailsSubscriptionParams `form:"subscription"`
}

// Search for charges you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
type ChargeSearchParams struct {
	SearchParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A cursor for pagination across multiple pages of results. Don't include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
	Page *string `form:"page"`
}

// AddExpand appends a new field to expand.
func (p *ChargeSearchParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Affiliate details for this purchase.
type ChargeCapturePaymentDetailsCarRentalAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargeCapturePaymentDetailsCarRentalDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargeCapturePaymentDetailsCarRentalDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargeCapturePaymentDetailsCarRentalDeliveryRecipientParams `form:"recipient"`
}

// The details of the distance traveled during the rental period.
type ChargeCapturePaymentDetailsCarRentalDistanceParams struct {
	// Distance traveled.
	Amount *int64 `form:"amount"`
	// Unit of measurement for the distance traveled. One of `miles` or `kilometers`.
	Unit *string `form:"unit"`
}

// The details of the passengers in the travel reservation
type ChargeCapturePaymentDetailsCarRentalDriverParams struct {
	// Driver's identification number.
	DriverIdentificationNumber *string `form:"driver_identification_number"`
	// Driver's tax number.
	DriverTaxNumber *string `form:"driver_tax_number"`
	// Full name of the person or entity on the car reservation.
	Name *string `form:"name"`
}

// Car rental details for this PaymentIntent.
type ChargeCapturePaymentDetailsCarRentalParams struct {
	// Affiliate details for this purchase.
	Affiliate *ChargeCapturePaymentDetailsCarRentalAffiliateParams `form:"affiliate"`
	// The booking number associated with the car rental.
	BookingNumber *string `form:"booking_number"`
	// Class code of the car.
	CarClassCode *string `form:"car_class_code"`
	// Make of the car.
	CarMake *string `form:"car_make"`
	// Model of the car.
	CarModel *string `form:"car_model"`
	// The name of the rental car company.
	Company *string `form:"company"`
	// The customer service phone number of the car rental company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// Number of days the car is being rented.
	DaysRented *int64 `form:"days_rented"`
	// Delivery details for this purchase.
	Delivery *ChargeCapturePaymentDetailsCarRentalDeliveryParams `form:"delivery"`
	// The details of the distance traveled during the rental period.
	Distance *ChargeCapturePaymentDetailsCarRentalDistanceParams `form:"distance"`
	// The details of the passengers in the travel reservation
	Drivers []*ChargeCapturePaymentDetailsCarRentalDriverParams `form:"drivers"`
	// List of additional charges being billed.
	ExtraCharges []*string `form:"extra_charges"`
	// Indicates if the customer did not keep nor cancel their booking.
	NoShow *bool `form:"no_show"`
	// Car pick-up address.
	PickupAddress *AddressParams `form:"pickup_address"`
	// Car pick-up time. Measured in seconds since the Unix epoch.
	PickupAt *int64 `form:"pickup_at"`
	// Name of the pickup location.
	PickupLocationName *string `form:"pickup_location_name"`
	// Rental rate.
	RateAmount *int64 `form:"rate_amount"`
	// The frequency at which the rate amount is applied. One of `day`, `week` or `month`
	RateInterval *string `form:"rate_interval"`
	// The name of the person or entity renting the car.
	RenterName *string `form:"renter_name"`
	// Car return address.
	ReturnAddress *AddressParams `form:"return_address"`
	// Car return time. Measured in seconds since the Unix epoch.
	ReturnAt *int64 `form:"return_at"`
	// Name of the return location.
	ReturnLocationName *string `form:"return_location_name"`
	// Indicates whether the goods or services are tax-exempt or tax is not collected.
	TaxExempt *bool `form:"tax_exempt"`
	// The vehicle identification number.
	VehicleIdentificationNumber *string `form:"vehicle_identification_number"`
}

// Affiliate (such as travel agency) details for the rental.
type ChargeCapturePaymentDetailsCarRentalDatumAffiliateParams struct {
	// Affiliate partner code.
	Code *string `form:"code"`
	// Name of affiliate partner.
	Name *string `form:"name"`
}

// Distance details for the rental.
type ChargeCapturePaymentDetailsCarRentalDatumDistanceParams struct {
	// Distance traveled.
	Amount *int64 `form:"amount"`
	// Unit of measurement for the distance traveled. One of `miles` or `kilometers`.
	Unit *string `form:"unit"`
}

// Driver's date of birth.
type ChargeCapturePaymentDetailsCarRentalDatumDriverDateOfBirthParams struct {
	// Day of birth (1-31).
	Day *int64 `form:"day"`
	// Month of birth (1-12).
	Month *int64 `form:"month"`
	// Year of birth (must be greater than 1900).
	Year *int64 `form:"year"`
}

// List of drivers for the rental.
type ChargeCapturePaymentDetailsCarRentalDatumDriverParams struct {
	// Driver's date of birth.
	DateOfBirth *ChargeCapturePaymentDetailsCarRentalDatumDriverDateOfBirthParams `form:"date_of_birth"`
	// Driver's identification number.
	DriverIdentificationNumber *string `form:"driver_identification_number"`
	// Driver's tax number.
	DriverTaxNumber *string `form:"driver_tax_number"`
	// Driver's full name.
	Name *string `form:"name"`
}

// Drop-off location details.
type ChargeCapturePaymentDetailsCarRentalDatumDropOffParams struct {
	// Address of the rental location.
	Address *AddressParams `form:"address"`
	// Location name.
	LocationName *string `form:"location_name"`
	// Timestamp for the location.
	Time *int64 `form:"time"`
}

// Insurance details for the rental.
type ChargeCapturePaymentDetailsCarRentalDatumInsuranceParams struct {
	// Amount of the insurance coverage in cents.
	Amount *int64 `form:"amount"`
	// Currency of the insurance amount.
	Currency *string `form:"currency"`
	// Name of the insurance company.
	InsuranceCompanyName *string `form:"insurance_company_name"`
	// Type of insurance coverage.
	InsuranceType *string `form:"insurance_type"`
}

// Pickup location details.
type ChargeCapturePaymentDetailsCarRentalDatumPickupParams struct {
	// Address of the rental location.
	Address *AddressParams `form:"address"`
	// Location name.
	LocationName *string `form:"location_name"`
	// Timestamp for the location.
	Time *int64 `form:"time"`
}

// Discount details for the rental.
type ChargeCapturePaymentDetailsCarRentalDatumTotalDiscountsParams struct {
	// Corporate client discount code.
	CorporateClientCode *string `form:"corporate_client_code"`
	// Coupon code applied to the rental.
	Coupon *string `form:"coupon"`
	// Maximum number of free miles or kilometers included.
	MaximumFreeMilesOrKilometers *int64 `form:"maximum_free_miles_or_kilometers"`
}

// Additional charges for the rental.
type ChargeCapturePaymentDetailsCarRentalDatumTotalExtraChargeParams struct {
	// Amount of the extra charge in cents.
	Amount *int64 `form:"amount"`
	// Type of extra charge.
	Type *string `form:"type"`
}

// Array of tax details.
type ChargeCapturePaymentDetailsCarRentalDatumTotalTaxTaxParams struct {
	// Tax amount.
	Amount *int64 `form:"amount"`
	// Tax rate applied.
	Rate *int64 `form:"rate"`
	// Type of tax applied.
	Type *string `form:"type"`
}

// Tax breakdown for the rental.
type ChargeCapturePaymentDetailsCarRentalDatumTotalTaxParams struct {
	// Array of tax details.
	Taxes []*ChargeCapturePaymentDetailsCarRentalDatumTotalTaxTaxParams `form:"taxes"`
	// Indicates if the transaction is tax exempt.
	TaxExemptIndicator *bool `form:"tax_exempt_indicator"`
}

// Total cost breakdown for the rental.
type ChargeCapturePaymentDetailsCarRentalDatumTotalParams struct {
	// Total amount in cents.
	Amount *int64 `form:"amount"`
	// Currency of the amount.
	Currency *string `form:"currency"`
	// Discount details for the rental.
	Discounts *ChargeCapturePaymentDetailsCarRentalDatumTotalDiscountsParams `form:"discounts"`
	// Additional charges for the rental.
	ExtraCharges []*ChargeCapturePaymentDetailsCarRentalDatumTotalExtraChargeParams `form:"extra_charges"`
	// Rate per unit for the rental.
	RatePerUnit *int64 `form:"rate_per_unit"`
	// Unit of measurement for the rate.
	RateUnit *string `form:"rate_unit"`
	// Tax breakdown for the rental.
	Tax *ChargeCapturePaymentDetailsCarRentalDatumTotalTaxParams `form:"tax"`
}

// Vehicle details for the rental.
type ChargeCapturePaymentDetailsCarRentalDatumVehicleParams struct {
	// Make of the rental vehicle.
	Make *string `form:"make"`
	// Model of the rental vehicle.
	Model *string `form:"model"`
	// Odometer reading at the time of rental.
	Odometer *int64 `form:"odometer"`
	// Type of the rental vehicle.
	Type *string `form:"type"`
	// Class of the rental vehicle.
	VehicleClass *string `form:"vehicle_class"`
	// Vehicle identification number (VIN).
	VehicleIdentificationNumber *string `form:"vehicle_identification_number"`
}

// Car rental data for this PaymentIntent.
type ChargeCapturePaymentDetailsCarRentalDatumParams struct {
	// Affiliate (such as travel agency) details for the rental.
	Affiliate *ChargeCapturePaymentDetailsCarRentalDatumAffiliateParams `form:"affiliate"`
	// Booking confirmation number for the car rental.
	BookingNumber *string `form:"booking_number"`
	// Name of the car rental company.
	CarrierName *string `form:"carrier_name"`
	// Customer service phone number for the car rental company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// Number of days the car is being rented.
	DaysRented *int64 `form:"days_rented"`
	// Distance details for the rental.
	Distance *ChargeCapturePaymentDetailsCarRentalDatumDistanceParams `form:"distance"`
	// List of drivers for the rental.
	Drivers []*ChargeCapturePaymentDetailsCarRentalDatumDriverParams `form:"drivers"`
	// Drop-off location details.
	DropOff *ChargeCapturePaymentDetailsCarRentalDatumDropOffParams `form:"drop_off"`
	// Insurance details for the rental.
	Insurances []*ChargeCapturePaymentDetailsCarRentalDatumInsuranceParams `form:"insurances"`
	// Indicates if the customer was a no-show.
	NoShowIndicator *bool `form:"no_show_indicator"`
	// Pickup location details.
	Pickup *ChargeCapturePaymentDetailsCarRentalDatumPickupParams `form:"pickup"`
	// Name of the person renting the vehicle.
	RenterName *string `form:"renter_name"`
	// Total cost breakdown for the rental.
	Total *ChargeCapturePaymentDetailsCarRentalDatumTotalParams `form:"total"`
	// Vehicle details for the rental.
	Vehicle *ChargeCapturePaymentDetailsCarRentalDatumVehicleParams `form:"vehicle"`
}

// Affiliate details for this purchase.
type ChargeCapturePaymentDetailsEventDetailsAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargeCapturePaymentDetailsEventDetailsDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargeCapturePaymentDetailsEventDetailsDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargeCapturePaymentDetailsEventDetailsDeliveryRecipientParams `form:"recipient"`
}

// Event details for this PaymentIntent
type ChargeCapturePaymentDetailsEventDetailsParams struct {
	// Indicates if the tickets are digitally checked when entering the venue.
	AccessControlledVenue *bool `form:"access_controlled_venue"`
	// The event location's address.
	Address *AddressParams `form:"address"`
	// Affiliate details for this purchase.
	Affiliate *ChargeCapturePaymentDetailsEventDetailsAffiliateParams `form:"affiliate"`
	// The name of the company
	Company *string `form:"company"`
	// Delivery details for this purchase.
	Delivery *ChargeCapturePaymentDetailsEventDetailsDeliveryParams `form:"delivery"`
	// Event end time. Measured in seconds since the Unix epoch.
	EndsAt *int64 `form:"ends_at"`
	// Type of the event entertainment (concert, sports event etc)
	Genre *string `form:"genre"`
	// The name of the event.
	Name *string `form:"name"`
	// Event start time. Measured in seconds since the Unix epoch.
	StartsAt *int64 `form:"starts_at"`
}

// Affiliate details for this purchase.
type ChargeCapturePaymentDetailsFlightAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargeCapturePaymentDetailsFlightDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargeCapturePaymentDetailsFlightDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargeCapturePaymentDetailsFlightDeliveryRecipientParams `form:"recipient"`
}

// The details of the passengers in the travel reservation.
type ChargeCapturePaymentDetailsFlightPassengerParams struct {
	// Full name of the person or entity on the flight reservation.
	Name *string `form:"name"`
}

// The individual flight segments associated with the trip.
type ChargeCapturePaymentDetailsFlightSegmentParams struct {
	// The flight segment amount.
	Amount *int64 `form:"amount"`
	// The International Air Transport Association (IATA) airport code for the arrival airport.
	ArrivalAirport *string `form:"arrival_airport"`
	// The arrival time for the flight segment. Measured in seconds since the Unix epoch.
	ArrivesAt *int64 `form:"arrives_at"`
	// The International Air Transport Association (IATA) carrier code of the carrier operating the flight segment.
	Carrier *string `form:"carrier"`
	// The departure time for the flight segment. Measured in seconds since the Unix epoch.
	DepartsAt *int64 `form:"departs_at"`
	// The International Air Transport Association (IATA) airport code for the departure airport.
	DepartureAirport *string `form:"departure_airport"`
	// The flight number associated with the segment
	FlightNumber *string `form:"flight_number"`
	// The fare class for the segment.
	ServiceClass *string `form:"service_class"`
}

// Flight reservation details for this PaymentIntent
type ChargeCapturePaymentDetailsFlightParams struct {
	// Affiliate details for this purchase.
	Affiliate *ChargeCapturePaymentDetailsFlightAffiliateParams `form:"affiliate"`
	// The agency number (i.e. International Air Transport Association (IATA) agency number) of the travel agency that made the booking.
	AgencyNumber *string `form:"agency_number"`
	// The International Air Transport Association (IATA) carrier code of the carrier that issued the ticket.
	Carrier *string `form:"carrier"`
	// Delivery details for this purchase.
	Delivery *ChargeCapturePaymentDetailsFlightDeliveryParams `form:"delivery"`
	// The name of the person or entity on the reservation.
	PassengerName *string `form:"passenger_name"`
	// The details of the passengers in the travel reservation.
	Passengers []*ChargeCapturePaymentDetailsFlightPassengerParams `form:"passengers"`
	// The individual flight segments associated with the trip.
	Segments []*ChargeCapturePaymentDetailsFlightSegmentParams `form:"segments"`
	// The ticket number associated with the travel reservation.
	TicketNumber *string `form:"ticket_number"`
}

// Affiliate details if applicable.
type ChargeCapturePaymentDetailsFlightDatumAffiliateParams struct {
	// Affiliate partner code.
	Code *string `form:"code"`
	// Name of affiliate partner.
	Name *string `form:"name"`
	// Code provided by the company to a travel agent authorizing ticket issuance.
	TravelAuthorizationCode *string `form:"travel_authorization_code"`
}

// List of insurances.
type ChargeCapturePaymentDetailsFlightDatumInsuranceParams struct {
	// Insurance cost.
	Amount *int64 `form:"amount"`
	// Insurance currency.
	Currency *string `form:"currency"`
	// Insurance company name.
	InsuranceCompanyName *string `form:"insurance_company_name"`
	// Type of insurance.
	InsuranceType *string `form:"insurance_type"`
}

// List of passengers.
type ChargeCapturePaymentDetailsFlightDatumPassengerParams struct {
	// Passenger's full name.
	Name *string `form:"name"`
}

// Arrival details.
type ChargeCapturePaymentDetailsFlightDatumSegmentArrivalParams struct {
	// Arrival airport IATA code.
	Airport *string `form:"airport"`
	// Arrival date/time.
	ArrivesAt *int64 `form:"arrives_at"`
	// Arrival city.
	City *string `form:"city"`
	// Arrival country.
	Country *string `form:"country"`
}

// Departure details.
type ChargeCapturePaymentDetailsFlightDatumSegmentDepartureParams struct {
	// Departure airport IATA code.
	Airport *string `form:"airport"`
	// Departure city.
	City *string `form:"city"`
	// Departure country.
	Country *string `form:"country"`
	// Departure date/time.
	DepartsAt *int64 `form:"departs_at"`
}

// List of flight segments.
type ChargeCapturePaymentDetailsFlightDatumSegmentParams struct {
	// Segment fare amount.
	Amount *int64 `form:"amount"`
	// Arrival details.
	Arrival *ChargeCapturePaymentDetailsFlightDatumSegmentArrivalParams `form:"arrival"`
	// Airline carrier code.
	CarrierCode *string `form:"carrier_code"`
	// Carrier name.
	CarrierName *string `form:"carrier_name"`
	// Segment currency.
	Currency *string `form:"currency"`
	// Departure details.
	Departure *ChargeCapturePaymentDetailsFlightDatumSegmentDepartureParams `form:"departure"`
	// Exchange ticket number.
	ExchangeTicketNumber *string `form:"exchange_ticket_number"`
	// Fare basis code.
	FareBasisCode *string `form:"fare_basis_code"`
	// Additional fees.
	Fees *int64 `form:"fees"`
	// Flight number.
	FlightNumber *string `form:"flight_number"`
	// Stopover indicator.
	IsStopOverIndicator *bool `form:"is_stop_over_indicator"`
	// Refundable ticket indicator.
	Refundable *bool `form:"refundable"`
	// Class of service.
	ServiceClass *string `form:"service_class"`
	// Tax amount for segment.
	TaxAmount *int64 `form:"tax_amount"`
	// Ticket number.
	TicketNumber *string `form:"ticket_number"`
}

// Discount details.
type ChargeCapturePaymentDetailsFlightDatumTotalDiscountsParams struct {
	// Corporate client discount code.
	CorporateClientCode *string `form:"corporate_client_code"`
}

// Additional charges.
type ChargeCapturePaymentDetailsFlightDatumTotalExtraChargeParams struct {
	// Amount of additional charges.
	Amount *int64 `form:"amount"`
	// Type of additional charges.
	Type *string `form:"type"`
}

// Array of tax details.
type ChargeCapturePaymentDetailsFlightDatumTotalTaxTaxParams struct {
	// Tax amount.
	Amount *int64 `form:"amount"`
	// Tax rate.
	Rate *int64 `form:"rate"`
	// Type of tax.
	Type *string `form:"type"`
}

// Tax breakdown.
type ChargeCapturePaymentDetailsFlightDatumTotalTaxParams struct {
	// Array of tax details.
	Taxes []*ChargeCapturePaymentDetailsFlightDatumTotalTaxTaxParams `form:"taxes"`
}

// Total cost breakdown.
type ChargeCapturePaymentDetailsFlightDatumTotalParams struct {
	// Total flight amount.
	Amount *int64 `form:"amount"`
	// Reason for credit.
	CreditReason *string `form:"credit_reason"`
	// Total currency.
	Currency *string `form:"currency"`
	// Discount details.
	Discounts *ChargeCapturePaymentDetailsFlightDatumTotalDiscountsParams `form:"discounts"`
	// Additional charges.
	ExtraCharges []*ChargeCapturePaymentDetailsFlightDatumTotalExtraChargeParams `form:"extra_charges"`
	// Tax breakdown.
	Tax *ChargeCapturePaymentDetailsFlightDatumTotalTaxParams `form:"tax"`
}

// Flight data for this PaymentIntent.
type ChargeCapturePaymentDetailsFlightDatumParams struct {
	// Affiliate details if applicable.
	Affiliate *ChargeCapturePaymentDetailsFlightDatumAffiliateParams `form:"affiliate"`
	// Reservation reference.
	BookingNumber *string `form:"booking_number"`
	// Computerized reservation system used to make the reservation and purchase the ticket.
	ComputerizedReservationSystem *string `form:"computerized_reservation_system"`
	// Ticket restrictions.
	EndorsementsAndRestrictions *string `form:"endorsements_and_restrictions"`
	// List of insurances.
	Insurances []*ChargeCapturePaymentDetailsFlightDatumInsuranceParams `form:"insurances"`
	// List of passengers.
	Passengers []*ChargeCapturePaymentDetailsFlightDatumPassengerParams `form:"passengers"`
	// List of flight segments.
	Segments []*ChargeCapturePaymentDetailsFlightDatumSegmentParams `form:"segments"`
	// Electronic ticket indicator.
	TicketElectronicallyIssuedIndicator *bool `form:"ticket_electronically_issued_indicator"`
	// Total cost breakdown.
	Total *ChargeCapturePaymentDetailsFlightDatumTotalParams `form:"total"`
	// Type of flight transaction.
	TransactionType *string `form:"transaction_type"`
}

// Affiliate details for this purchase.
type ChargeCapturePaymentDetailsLodgingAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargeCapturePaymentDetailsLodgingDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargeCapturePaymentDetailsLodgingDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargeCapturePaymentDetailsLodgingDeliveryRecipientParams `form:"recipient"`
}

// The details of the passengers in the travel reservation
type ChargeCapturePaymentDetailsLodgingPassengerParams struct {
	// Full name of the person or entity on the lodging reservation.
	Name *string `form:"name"`
}

// Lodging reservation details for this PaymentIntent
type ChargeCapturePaymentDetailsLodgingParams struct {
	// The lodging location's address.
	Address *AddressParams `form:"address"`
	// The number of adults on the booking
	Adults *int64 `form:"adults"`
	// Affiliate details for this purchase.
	Affiliate *ChargeCapturePaymentDetailsLodgingAffiliateParams `form:"affiliate"`
	// The booking number associated with the lodging reservation.
	BookingNumber *string `form:"booking_number"`
	// The lodging category
	Category *string `form:"category"`
	// Lodging check-in time. Measured in seconds since the Unix epoch.
	CheckinAt *int64 `form:"checkin_at"`
	// Lodging check-out time. Measured in seconds since the Unix epoch.
	CheckoutAt *int64 `form:"checkout_at"`
	// The customer service phone number of the lodging company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// The daily lodging room rate.
	DailyRoomRateAmount *int64 `form:"daily_room_rate_amount"`
	// Delivery details for this purchase.
	Delivery *ChargeCapturePaymentDetailsLodgingDeliveryParams `form:"delivery"`
	// List of additional charges being billed.
	ExtraCharges []*string `form:"extra_charges"`
	// Indicates whether the lodging location is compliant with the Fire Safety Act.
	FireSafetyActCompliance *bool `form:"fire_safety_act_compliance"`
	// The name of the lodging location.
	Name *string `form:"name"`
	// Indicates if the customer did not keep their booking while failing to cancel the reservation.
	NoShow *bool `form:"no_show"`
	// The number of rooms on the booking
	NumberOfRooms *int64 `form:"number_of_rooms"`
	// The details of the passengers in the travel reservation
	Passengers []*ChargeCapturePaymentDetailsLodgingPassengerParams `form:"passengers"`
	// The phone number of the lodging location.
	PropertyPhoneNumber *string `form:"property_phone_number"`
	// The room class for this purchase.
	RoomClass *string `form:"room_class"`
	// The number of room nights
	RoomNights *int64 `form:"room_nights"`
	// The total tax amount associating with the room reservation.
	TotalRoomTaxAmount *int64 `form:"total_room_tax_amount"`
	// The total tax amount
	TotalTaxAmount *int64 `form:"total_tax_amount"`
}

// Accommodation details for the lodging.
type ChargeCapturePaymentDetailsLodgingDatumAccommodationParams struct {
	// Type of accommodation.
	AccommodationType *string `form:"accommodation_type"`
	// Bed type.
	BedType *string `form:"bed_type"`
	// Daily accommodation rate in cents.
	DailyRateAmount *int64 `form:"daily_rate_amount"`
	// Number of nights.
	Nights *int64 `form:"nights"`
	// Number of rooms, cabanas, apartments, and so on.
	NumberOfRooms *int64 `form:"number_of_rooms"`
	// Rate type.
	RateType *string `form:"rate_type"`
	// Whether smoking is allowed.
	SmokingIndicator *bool `form:"smoking_indicator"`
}

// Affiliate details if applicable.
type ChargeCapturePaymentDetailsLodgingDatumAffiliateParams struct {
	// Affiliate partner code.
	Code *string `form:"code"`
	// Affiliate partner name.
	Name *string `form:"name"`
}

// List of guests for the lodging.
type ChargeCapturePaymentDetailsLodgingDatumGuestParams struct {
	// Guest's full name.
	Name *string `form:"name"`
}

// Host details for the lodging.
type ChargeCapturePaymentDetailsLodgingDatumHostParams struct {
	// Address of the host.
	Address *AddressParams `form:"address"`
	// Host's country of domicile.
	CountryOfDomicile *string `form:"country_of_domicile"`
	// Reference number for the host.
	HostReference *string `form:"host_reference"`
	// Type of host.
	HostType *string `form:"host_type"`
	// Name of the lodging property or host.
	Name *string `form:"name"`
	// Total number of reservations for the host.
	NumberOfReservations *int64 `form:"number_of_reservations"`
	// Property phone number.
	PropertyPhoneNumber *string `form:"property_phone_number"`
	// Host's registration date.
	RegisteredAt *int64 `form:"registered_at"`
}

// List of insurances for the lodging.
type ChargeCapturePaymentDetailsLodgingDatumInsuranceParams struct {
	// Price of the insurance coverage in cents.
	Amount *int64 `form:"amount"`
	// Currency of the insurance amount.
	Currency *string `form:"currency"`
	// Name of the insurance company.
	InsuranceCompanyName *string `form:"insurance_company_name"`
	// Type of insurance coverage.
	InsuranceType *string `form:"insurance_type"`
}

// Discount details for the lodging.
type ChargeCapturePaymentDetailsLodgingDatumTotalDiscountsParams struct {
	// Corporate client discount code.
	CorporateClientCode *string `form:"corporate_client_code"`
	// Coupon code.
	Coupon *string `form:"coupon"`
}

// Additional charges for the lodging.
type ChargeCapturePaymentDetailsLodgingDatumTotalExtraChargeParams struct {
	// Amount of the extra charge in cents.
	Amount *int64 `form:"amount"`
	// Type of extra charge.
	Type *string `form:"type"`
}

// Tax details.
type ChargeCapturePaymentDetailsLodgingDatumTotalTaxTaxParams struct {
	// Tax amount in cents.
	Amount *int64 `form:"amount"`
	// Tax rate.
	Rate *int64 `form:"rate"`
	// Type of tax applied.
	Type *string `form:"type"`
}

// Tax breakdown for the lodging reservation.
type ChargeCapturePaymentDetailsLodgingDatumTotalTaxParams struct {
	// Tax details.
	Taxes []*ChargeCapturePaymentDetailsLodgingDatumTotalTaxTaxParams `form:"taxes"`
	// Indicates whether the transaction is tax exempt.
	TaxExemptIndicator *bool `form:"tax_exempt_indicator"`
}

// Total details for the lodging.
type ChargeCapturePaymentDetailsLodgingDatumTotalParams struct {
	// Total price of the lodging reservation in cents.
	Amount *int64 `form:"amount"`
	// Cash advances in cents.
	CashAdvances *int64 `form:"cash_advances"`
	// Currency of the total amount.
	Currency *string `form:"currency"`
	// Discount details for the lodging.
	Discounts *ChargeCapturePaymentDetailsLodgingDatumTotalDiscountsParams `form:"discounts"`
	// Additional charges for the lodging.
	ExtraCharges []*ChargeCapturePaymentDetailsLodgingDatumTotalExtraChargeParams `form:"extra_charges"`
	// Prepaid amount in cents.
	PrepaidAmount *int64 `form:"prepaid_amount"`
	// Tax breakdown for the lodging reservation.
	Tax *ChargeCapturePaymentDetailsLodgingDatumTotalTaxParams `form:"tax"`
}

// Lodging data for this PaymentIntent.
type ChargeCapturePaymentDetailsLodgingDatumParams struct {
	// Accommodation details for the lodging.
	Accommodation *ChargeCapturePaymentDetailsLodgingDatumAccommodationParams `form:"accommodation"`
	// Affiliate details if applicable.
	Affiliate *ChargeCapturePaymentDetailsLodgingDatumAffiliateParams `form:"affiliate"`
	// Booking confirmation number for the lodging.
	BookingNumber *string `form:"booking_number"`
	// Check-in date.
	CheckinAt *int64 `form:"checkin_at"`
	// Check-out date.
	CheckoutAt *int64 `form:"checkout_at"`
	// Customer service phone number for the lodging company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// Whether the lodging is compliant with any hotel fire safety regulations.
	FireSafetyActComplianceIndicator *bool `form:"fire_safety_act_compliance_indicator"`
	// List of guests for the lodging.
	Guests []*ChargeCapturePaymentDetailsLodgingDatumGuestParams `form:"guests"`
	// Host details for the lodging.
	Host *ChargeCapturePaymentDetailsLodgingDatumHostParams `form:"host"`
	// List of insurances for the lodging.
	Insurances []*ChargeCapturePaymentDetailsLodgingDatumInsuranceParams `form:"insurances"`
	// Whether the renter is a no-show.
	NoShowIndicator *bool `form:"no_show_indicator"`
	// Renter ID number for the lodging.
	RenterIDNumber *string `form:"renter_id_number"`
	// Renter name for the lodging.
	RenterName *string `form:"renter_name"`
	// Total details for the lodging.
	Total *ChargeCapturePaymentDetailsLodgingDatumTotalParams `form:"total"`
}

// Affiliate details for this purchase.
type ChargeCapturePaymentDetailsSubscriptionAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Subscription billing details for this purchase.
type ChargeCapturePaymentDetailsSubscriptionBillingIntervalParams struct {
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	Count *int64 `form:"count"`
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
}

// Subscription details for this PaymentIntent
type ChargeCapturePaymentDetailsSubscriptionParams struct {
	// Affiliate details for this purchase.
	Affiliate *ChargeCapturePaymentDetailsSubscriptionAffiliateParams `form:"affiliate"`
	// Info whether the subscription will be auto renewed upon expiry.
	AutoRenewal *bool `form:"auto_renewal"`
	// Subscription billing details for this purchase.
	BillingInterval *ChargeCapturePaymentDetailsSubscriptionBillingIntervalParams `form:"billing_interval"`
	// Subscription end time. Measured in seconds since the Unix epoch.
	EndsAt *int64 `form:"ends_at"`
	// Name of the product on subscription. e.g. Apple Music Subscription
	Name *string `form:"name"`
	// Subscription start time. Measured in seconds since the Unix epoch.
	StartsAt *int64 `form:"starts_at"`
}

// Provides industry-specific information about the charge.
type ChargeCapturePaymentDetailsParams struct {
	// Car rental details for this PaymentIntent.
	CarRental *ChargeCapturePaymentDetailsCarRentalParams `form:"car_rental"`
	// Car rental data for this PaymentIntent.
	CarRentalData []*ChargeCapturePaymentDetailsCarRentalDatumParams `form:"car_rental_data"`
	// A unique value to identify the customer. This field is available only for card payments.
	//
	// This field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.
	CustomerReference *string `form:"customer_reference"`
	// Event details for this PaymentIntent
	EventDetails *ChargeCapturePaymentDetailsEventDetailsParams `form:"event_details"`
	// Flight reservation details for this PaymentIntent
	Flight *ChargeCapturePaymentDetailsFlightParams `form:"flight"`
	// Flight data for this PaymentIntent.
	FlightData []*ChargeCapturePaymentDetailsFlightDatumParams `form:"flight_data"`
	// Lodging reservation details for this PaymentIntent
	Lodging *ChargeCapturePaymentDetailsLodgingParams `form:"lodging"`
	// Lodging data for this PaymentIntent.
	LodgingData []*ChargeCapturePaymentDetailsLodgingDatumParams `form:"lodging_data"`
	// A unique value assigned by the business to identify the transaction. Required for L2 and L3 rates.
	//
	// Required when the Payment Method Types array contains `card`, including when [automatic_payment_methods.enabled](https://docs.stripe.com/api/payment_intents/create#create_payment_intent-automatic_payment_methods-enabled) is set to `true`.
	//
	// For Cards, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks. For Klarna, this field is truncated to 255 characters and is visible to customers when they view the order in the Klarna app.
	OrderReference *string `form:"order_reference"`
	// Subscription details for this PaymentIntent
	Subscription *ChargeCapturePaymentDetailsSubscriptionParams `form:"subscription"`
}

// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.
type ChargeCaptureTransferDataParams struct {
	// The amount transferred to the destination account, if specified. By default, the entire charge amount is transferred to the destination account.
	Amount *int64 `form:"amount"`
}

// Capture the payment of an existing, uncaptured charge that was created with the capture option set to false.
//
// Uncaptured payments expire a set number of days after they are created ([7 by default](https://docs.stripe.com/docs/charges/placing-a-hold)), after which they are marked as refunded and capture attempts will fail.
//
// Don't use this method to capture a PaymentIntent-initiated charge. Use [Capture a PaymentIntent](https://docs.stripe.com/docs/api/payment_intents/capture).
type ChargeCaptureParams struct {
	Params `form:"*"`
	// The amount to capture, which must be less than or equal to the original amount.
	Amount *int64 `form:"amount"`
	// An application fee to add on to this charge.
	ApplicationFee *int64 `form:"application_fee"`
	// An application fee amount to add on to this charge, which must be less than or equal to the original amount.
	ApplicationFeeAmount *int64   `form:"application_fee_amount"`
	ExchangeRate         *float64 `form:"exchange_rate"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Provides industry-specific information about the charge.
	PaymentDetails *ChargeCapturePaymentDetailsParams `form:"payment_details"`
	// The email address to send this charge's receipt to. This will override the previously-specified email address for this charge, if one was set. Receipts will not be sent in test mode.
	ReceiptEmail *string `form:"receipt_email"`
	// For a non-card charge, text that appears on the customer's statement as the statement descriptor. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	//
	// For a card charge, this value is ignored unless you don't specify a `statement_descriptor_suffix`, in which case this value is used as the suffix.
	StatementDescriptor *string `form:"statement_descriptor"`
	// Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement. If the account has no prefix value, the suffix is concatenated to the account's statement descriptor.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix"`
	// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.
	TransferData *ChargeCaptureTransferDataParams `form:"transfer_data"`
	// A string that identifies this transaction as part of a group. `transfer_group` may only be provided if it has not been set. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options) for details.
	TransferGroup *string `form:"transfer_group"`
}

// AddExpand appends a new field to expand.
func (p *ChargeCaptureParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type ChargeCreateDestinationParams struct {
	// ID of an existing, connected Stripe account.
	Account *string `form:"account"`
	// The amount to transfer to the destination account without creating an `Application Fee` object. Cannot be combined with the `application_fee` parameter. Must be less than or equal to the charge amount.
	Amount *int64 `form:"amount"`
}

// Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.
type ChargeCreateRadarOptionsParams struct {
	// A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.
type ChargeCreateTransferDataParams struct {
	// The amount transferred to the destination account, if specified. By default, the entire charge amount is transferred to the destination account.
	Amount *int64 `form:"amount"`
	// ID of an existing, connected Stripe account.
	Destination *string `form:"destination"`
}
type ChargeCreateLevel3LineItemParams struct {
	DiscountAmount     *int64  `form:"discount_amount"`
	ProductCode        *string `form:"product_code"`
	ProductDescription *string `form:"product_description"`
	Quantity           *int64  `form:"quantity"`
	TaxAmount          *int64  `form:"tax_amount"`
	UnitCost           *int64  `form:"unit_cost"`
}
type ChargeCreateLevel3Params struct {
	CustomerReference  *string                             `form:"customer_reference"`
	LineItems          []*ChargeCreateLevel3LineItemParams `form:"line_items"`
	MerchantReference  *string                             `form:"merchant_reference"`
	ShippingAddressZip *string                             `form:"shipping_address_zip"`
	ShippingAmount     *int64                              `form:"shipping_amount"`
	ShippingFromZip    *string                             `form:"shipping_from_zip"`
}

// This method is no longer recommended—use the [Payment Intents API](https://docs.stripe.com/docs/api/payment_intents)
// to initiate a new payment instead. Confirmation of the PaymentIntent creates the Charge
// object used to request payment.
type ChargeCreateParams struct {
	Params `form:"*"`
	// Amount intended to be collected by this payment. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://docs.stripe.com/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount         *int64 `form:"amount"`
	ApplicationFee *int64 `form:"application_fee"`
	// A fee in cents (or local equivalent) that will be applied to the charge and transferred to the application owner's Stripe account. The request must be made with an OAuth key or the `Stripe-Account` header in order to take an application fee. For more information, see the application fees [documentation](https://docs.stripe.com/connect/direct-charges#collect-fees).
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// Whether to immediately capture the charge. Defaults to `true`. When `false`, the charge issues an authorization (or pre-authorization), and will need to be [captured](https://api.stripe.com#capture_charge) later. Uncaptured charges expire after a set number of days (7 by default). For more information, see the [authorizing charges and settling later](https://docs.stripe.com/charges/placing-a-hold) documentation.
	Capture *bool `form:"capture"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of an existing customer that will be charged in this request.
	Customer *string `form:"customer"`
	// An arbitrary string which you can attach to a `Charge` object. It is displayed when in the web interface alongside the charge. Note that if you use Stripe to send automatic email receipts to your customers, your receipt emails will include the `description` of the charge(s) that they are describing.
	Description  *string                        `form:"description"`
	Destination  *ChargeCreateDestinationParams `form:"destination"`
	ExchangeRate *float64                       `form:"exchange_rate"`
	// Specifies which fields in the response should be expanded.
	Expand []*string                 `form:"expand"`
	Level3 *ChargeCreateLevel3Params `form:"level3"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The Stripe account ID for which these funds are intended. You can specify the business of record as the connected account using the `on_behalf_of` attribute on the charge. For details, see [Creating Separate Charges and Transfers](https://docs.stripe.com/connect/separate-charges-and-transfers#settlement-merchant).
	OnBehalfOf *string `form:"on_behalf_of"`
	// Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.
	RadarOptions *ChargeCreateRadarOptionsParams `form:"radar_options"`
	// The email address to which this charge's [receipt](https://docs.stripe.com/dashboard/receipts) will be sent. The receipt will not be sent until the charge is paid, and no receipts will be sent for test mode charges. If this charge is for a [Customer](https://docs.stripe.com/api/customers/object), the email address specified here will override the customer's email address. If `receipt_email` is specified for a charge in live mode, a receipt will be sent regardless of your [email settings](https://dashboard.stripe.com/account/emails).
	ReceiptEmail *string `form:"receipt_email"`
	// Shipping information for the charge. Helps prevent fraud on charges for physical goods.
	Shipping *ShippingDetailsParams     `form:"shipping"`
	Source   *PaymentSourceSourceParams `form:"*"` // PaymentSourceSourceParams has custom encoding so brought to top level with "*"
	// For a non-card charge, text that appears on the customer's statement as the statement descriptor. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	//
	// For a card charge, this value is ignored unless you don't specify a `statement_descriptor_suffix`, in which case this value is used as the suffix.
	StatementDescriptor *string `form:"statement_descriptor"`
	// Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement. If the account has no prefix value, the suffix is concatenated to the account's statement descriptor.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix"`
	// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.
	TransferData *ChargeCreateTransferDataParams `form:"transfer_data"`
	// A string that identifies this transaction as part of a group. For details, see [Grouping transactions](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options).
	TransferGroup *string `form:"transfer_group"`
}

// SetSource adds valid sources to a ChargeCreateParams object,
// returning an error for unsupported sources.
func (p *ChargeCreateParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// AddExpand appends a new field to expand.
func (p *ChargeCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *ChargeCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of a charge that has previously been created. Supply the unique charge ID that was returned from your previous request, and Stripe will return the corresponding charge information. The same information is returned when creating or refunding the charge.
type ChargeRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *ChargeRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A set of key-value pairs you can attach to a charge giving information about its riskiness. If you believe a charge is fraudulent, include a `user_report` key with a value of `fraudulent`. If you believe a charge is safe, include a `user_report` key with a value of `safe`. Stripe will use the information you send to improve our fraud detection algorithms.
type ChargeUpdateFraudDetailsParams struct {
	// Either `safe` or `fraudulent`.
	UserReport *string `form:"user_report"`
}

// Affiliate details for this purchase.
type ChargeUpdatePaymentDetailsCarRentalAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargeUpdatePaymentDetailsCarRentalDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargeUpdatePaymentDetailsCarRentalDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargeUpdatePaymentDetailsCarRentalDeliveryRecipientParams `form:"recipient"`
}

// The details of the distance traveled during the rental period.
type ChargeUpdatePaymentDetailsCarRentalDistanceParams struct {
	// Distance traveled.
	Amount *int64 `form:"amount"`
	// Unit of measurement for the distance traveled. One of `miles` or `kilometers`.
	Unit *string `form:"unit"`
}

// The details of the passengers in the travel reservation
type ChargeUpdatePaymentDetailsCarRentalDriverParams struct {
	// Driver's identification number.
	DriverIdentificationNumber *string `form:"driver_identification_number"`
	// Driver's tax number.
	DriverTaxNumber *string `form:"driver_tax_number"`
	// Full name of the person or entity on the car reservation.
	Name *string `form:"name"`
}

// Car rental details for this PaymentIntent.
type ChargeUpdatePaymentDetailsCarRentalParams struct {
	// Affiliate details for this purchase.
	Affiliate *ChargeUpdatePaymentDetailsCarRentalAffiliateParams `form:"affiliate"`
	// The booking number associated with the car rental.
	BookingNumber *string `form:"booking_number"`
	// Class code of the car.
	CarClassCode *string `form:"car_class_code"`
	// Make of the car.
	CarMake *string `form:"car_make"`
	// Model of the car.
	CarModel *string `form:"car_model"`
	// The name of the rental car company.
	Company *string `form:"company"`
	// The customer service phone number of the car rental company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// Number of days the car is being rented.
	DaysRented *int64 `form:"days_rented"`
	// Delivery details for this purchase.
	Delivery *ChargeUpdatePaymentDetailsCarRentalDeliveryParams `form:"delivery"`
	// The details of the distance traveled during the rental period.
	Distance *ChargeUpdatePaymentDetailsCarRentalDistanceParams `form:"distance"`
	// The details of the passengers in the travel reservation
	Drivers []*ChargeUpdatePaymentDetailsCarRentalDriverParams `form:"drivers"`
	// List of additional charges being billed.
	ExtraCharges []*string `form:"extra_charges"`
	// Indicates if the customer did not keep nor cancel their booking.
	NoShow *bool `form:"no_show"`
	// Car pick-up address.
	PickupAddress *AddressParams `form:"pickup_address"`
	// Car pick-up time. Measured in seconds since the Unix epoch.
	PickupAt *int64 `form:"pickup_at"`
	// Name of the pickup location.
	PickupLocationName *string `form:"pickup_location_name"`
	// Rental rate.
	RateAmount *int64 `form:"rate_amount"`
	// The frequency at which the rate amount is applied. One of `day`, `week` or `month`
	RateInterval *string `form:"rate_interval"`
	// The name of the person or entity renting the car.
	RenterName *string `form:"renter_name"`
	// Car return address.
	ReturnAddress *AddressParams `form:"return_address"`
	// Car return time. Measured in seconds since the Unix epoch.
	ReturnAt *int64 `form:"return_at"`
	// Name of the return location.
	ReturnLocationName *string `form:"return_location_name"`
	// Indicates whether the goods or services are tax-exempt or tax is not collected.
	TaxExempt *bool `form:"tax_exempt"`
	// The vehicle identification number.
	VehicleIdentificationNumber *string `form:"vehicle_identification_number"`
}

// Affiliate (such as travel agency) details for the rental.
type ChargeUpdatePaymentDetailsCarRentalDatumAffiliateParams struct {
	// Affiliate partner code.
	Code *string `form:"code"`
	// Name of affiliate partner.
	Name *string `form:"name"`
}

// Distance details for the rental.
type ChargeUpdatePaymentDetailsCarRentalDatumDistanceParams struct {
	// Distance traveled.
	Amount *int64 `form:"amount"`
	// Unit of measurement for the distance traveled. One of `miles` or `kilometers`.
	Unit *string `form:"unit"`
}

// Driver's date of birth.
type ChargeUpdatePaymentDetailsCarRentalDatumDriverDateOfBirthParams struct {
	// Day of birth (1-31).
	Day *int64 `form:"day"`
	// Month of birth (1-12).
	Month *int64 `form:"month"`
	// Year of birth (must be greater than 1900).
	Year *int64 `form:"year"`
}

// List of drivers for the rental.
type ChargeUpdatePaymentDetailsCarRentalDatumDriverParams struct {
	// Driver's date of birth.
	DateOfBirth *ChargeUpdatePaymentDetailsCarRentalDatumDriverDateOfBirthParams `form:"date_of_birth"`
	// Driver's identification number.
	DriverIdentificationNumber *string `form:"driver_identification_number"`
	// Driver's tax number.
	DriverTaxNumber *string `form:"driver_tax_number"`
	// Driver's full name.
	Name *string `form:"name"`
}

// Drop-off location details.
type ChargeUpdatePaymentDetailsCarRentalDatumDropOffParams struct {
	// Address of the rental location.
	Address *AddressParams `form:"address"`
	// Location name.
	LocationName *string `form:"location_name"`
	// Timestamp for the location.
	Time *int64 `form:"time"`
}

// Insurance details for the rental.
type ChargeUpdatePaymentDetailsCarRentalDatumInsuranceParams struct {
	// Amount of the insurance coverage in cents.
	Amount *int64 `form:"amount"`
	// Currency of the insurance amount.
	Currency *string `form:"currency"`
	// Name of the insurance company.
	InsuranceCompanyName *string `form:"insurance_company_name"`
	// Type of insurance coverage.
	InsuranceType *string `form:"insurance_type"`
}

// Pickup location details.
type ChargeUpdatePaymentDetailsCarRentalDatumPickupParams struct {
	// Address of the rental location.
	Address *AddressParams `form:"address"`
	// Location name.
	LocationName *string `form:"location_name"`
	// Timestamp for the location.
	Time *int64 `form:"time"`
}

// Discount details for the rental.
type ChargeUpdatePaymentDetailsCarRentalDatumTotalDiscountsParams struct {
	// Corporate client discount code.
	CorporateClientCode *string `form:"corporate_client_code"`
	// Coupon code applied to the rental.
	Coupon *string `form:"coupon"`
	// Maximum number of free miles or kilometers included.
	MaximumFreeMilesOrKilometers *int64 `form:"maximum_free_miles_or_kilometers"`
}

// Additional charges for the rental.
type ChargeUpdatePaymentDetailsCarRentalDatumTotalExtraChargeParams struct {
	// Amount of the extra charge in cents.
	Amount *int64 `form:"amount"`
	// Type of extra charge.
	Type *string `form:"type"`
}

// Array of tax details.
type ChargeUpdatePaymentDetailsCarRentalDatumTotalTaxTaxParams struct {
	// Tax amount.
	Amount *int64 `form:"amount"`
	// Tax rate applied.
	Rate *int64 `form:"rate"`
	// Type of tax applied.
	Type *string `form:"type"`
}

// Tax breakdown for the rental.
type ChargeUpdatePaymentDetailsCarRentalDatumTotalTaxParams struct {
	// Array of tax details.
	Taxes []*ChargeUpdatePaymentDetailsCarRentalDatumTotalTaxTaxParams `form:"taxes"`
	// Indicates if the transaction is tax exempt.
	TaxExemptIndicator *bool `form:"tax_exempt_indicator"`
}

// Total cost breakdown for the rental.
type ChargeUpdatePaymentDetailsCarRentalDatumTotalParams struct {
	// Total amount in cents.
	Amount *int64 `form:"amount"`
	// Currency of the amount.
	Currency *string `form:"currency"`
	// Discount details for the rental.
	Discounts *ChargeUpdatePaymentDetailsCarRentalDatumTotalDiscountsParams `form:"discounts"`
	// Additional charges for the rental.
	ExtraCharges []*ChargeUpdatePaymentDetailsCarRentalDatumTotalExtraChargeParams `form:"extra_charges"`
	// Rate per unit for the rental.
	RatePerUnit *int64 `form:"rate_per_unit"`
	// Unit of measurement for the rate.
	RateUnit *string `form:"rate_unit"`
	// Tax breakdown for the rental.
	Tax *ChargeUpdatePaymentDetailsCarRentalDatumTotalTaxParams `form:"tax"`
}

// Vehicle details for the rental.
type ChargeUpdatePaymentDetailsCarRentalDatumVehicleParams struct {
	// Make of the rental vehicle.
	Make *string `form:"make"`
	// Model of the rental vehicle.
	Model *string `form:"model"`
	// Odometer reading at the time of rental.
	Odometer *int64 `form:"odometer"`
	// Type of the rental vehicle.
	Type *string `form:"type"`
	// Class of the rental vehicle.
	VehicleClass *string `form:"vehicle_class"`
	// Vehicle identification number (VIN).
	VehicleIdentificationNumber *string `form:"vehicle_identification_number"`
}

// Car rental data for this PaymentIntent.
type ChargeUpdatePaymentDetailsCarRentalDatumParams struct {
	// Affiliate (such as travel agency) details for the rental.
	Affiliate *ChargeUpdatePaymentDetailsCarRentalDatumAffiliateParams `form:"affiliate"`
	// Booking confirmation number for the car rental.
	BookingNumber *string `form:"booking_number"`
	// Name of the car rental company.
	CarrierName *string `form:"carrier_name"`
	// Customer service phone number for the car rental company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// Number of days the car is being rented.
	DaysRented *int64 `form:"days_rented"`
	// Distance details for the rental.
	Distance *ChargeUpdatePaymentDetailsCarRentalDatumDistanceParams `form:"distance"`
	// List of drivers for the rental.
	Drivers []*ChargeUpdatePaymentDetailsCarRentalDatumDriverParams `form:"drivers"`
	// Drop-off location details.
	DropOff *ChargeUpdatePaymentDetailsCarRentalDatumDropOffParams `form:"drop_off"`
	// Insurance details for the rental.
	Insurances []*ChargeUpdatePaymentDetailsCarRentalDatumInsuranceParams `form:"insurances"`
	// Indicates if the customer was a no-show.
	NoShowIndicator *bool `form:"no_show_indicator"`
	// Pickup location details.
	Pickup *ChargeUpdatePaymentDetailsCarRentalDatumPickupParams `form:"pickup"`
	// Name of the person renting the vehicle.
	RenterName *string `form:"renter_name"`
	// Total cost breakdown for the rental.
	Total *ChargeUpdatePaymentDetailsCarRentalDatumTotalParams `form:"total"`
	// Vehicle details for the rental.
	Vehicle *ChargeUpdatePaymentDetailsCarRentalDatumVehicleParams `form:"vehicle"`
}

// Affiliate details for this purchase.
type ChargeUpdatePaymentDetailsEventDetailsAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargeUpdatePaymentDetailsEventDetailsDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargeUpdatePaymentDetailsEventDetailsDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargeUpdatePaymentDetailsEventDetailsDeliveryRecipientParams `form:"recipient"`
}

// Event details for this PaymentIntent
type ChargeUpdatePaymentDetailsEventDetailsParams struct {
	// Indicates if the tickets are digitally checked when entering the venue.
	AccessControlledVenue *bool `form:"access_controlled_venue"`
	// The event location's address.
	Address *AddressParams `form:"address"`
	// Affiliate details for this purchase.
	Affiliate *ChargeUpdatePaymentDetailsEventDetailsAffiliateParams `form:"affiliate"`
	// The name of the company
	Company *string `form:"company"`
	// Delivery details for this purchase.
	Delivery *ChargeUpdatePaymentDetailsEventDetailsDeliveryParams `form:"delivery"`
	// Event end time. Measured in seconds since the Unix epoch.
	EndsAt *int64 `form:"ends_at"`
	// Type of the event entertainment (concert, sports event etc)
	Genre *string `form:"genre"`
	// The name of the event.
	Name *string `form:"name"`
	// Event start time. Measured in seconds since the Unix epoch.
	StartsAt *int64 `form:"starts_at"`
}

// Affiliate details for this purchase.
type ChargeUpdatePaymentDetailsFlightAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargeUpdatePaymentDetailsFlightDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargeUpdatePaymentDetailsFlightDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargeUpdatePaymentDetailsFlightDeliveryRecipientParams `form:"recipient"`
}

// The details of the passengers in the travel reservation.
type ChargeUpdatePaymentDetailsFlightPassengerParams struct {
	// Full name of the person or entity on the flight reservation.
	Name *string `form:"name"`
}

// The individual flight segments associated with the trip.
type ChargeUpdatePaymentDetailsFlightSegmentParams struct {
	// The flight segment amount.
	Amount *int64 `form:"amount"`
	// The International Air Transport Association (IATA) airport code for the arrival airport.
	ArrivalAirport *string `form:"arrival_airport"`
	// The arrival time for the flight segment. Measured in seconds since the Unix epoch.
	ArrivesAt *int64 `form:"arrives_at"`
	// The International Air Transport Association (IATA) carrier code of the carrier operating the flight segment.
	Carrier *string `form:"carrier"`
	// The departure time for the flight segment. Measured in seconds since the Unix epoch.
	DepartsAt *int64 `form:"departs_at"`
	// The International Air Transport Association (IATA) airport code for the departure airport.
	DepartureAirport *string `form:"departure_airport"`
	// The flight number associated with the segment
	FlightNumber *string `form:"flight_number"`
	// The fare class for the segment.
	ServiceClass *string `form:"service_class"`
}

// Flight reservation details for this PaymentIntent
type ChargeUpdatePaymentDetailsFlightParams struct {
	// Affiliate details for this purchase.
	Affiliate *ChargeUpdatePaymentDetailsFlightAffiliateParams `form:"affiliate"`
	// The agency number (i.e. International Air Transport Association (IATA) agency number) of the travel agency that made the booking.
	AgencyNumber *string `form:"agency_number"`
	// The International Air Transport Association (IATA) carrier code of the carrier that issued the ticket.
	Carrier *string `form:"carrier"`
	// Delivery details for this purchase.
	Delivery *ChargeUpdatePaymentDetailsFlightDeliveryParams `form:"delivery"`
	// The name of the person or entity on the reservation.
	PassengerName *string `form:"passenger_name"`
	// The details of the passengers in the travel reservation.
	Passengers []*ChargeUpdatePaymentDetailsFlightPassengerParams `form:"passengers"`
	// The individual flight segments associated with the trip.
	Segments []*ChargeUpdatePaymentDetailsFlightSegmentParams `form:"segments"`
	// The ticket number associated with the travel reservation.
	TicketNumber *string `form:"ticket_number"`
}

// Affiliate details if applicable.
type ChargeUpdatePaymentDetailsFlightDatumAffiliateParams struct {
	// Affiliate partner code.
	Code *string `form:"code"`
	// Name of affiliate partner.
	Name *string `form:"name"`
	// Code provided by the company to a travel agent authorizing ticket issuance.
	TravelAuthorizationCode *string `form:"travel_authorization_code"`
}

// List of insurances.
type ChargeUpdatePaymentDetailsFlightDatumInsuranceParams struct {
	// Insurance cost.
	Amount *int64 `form:"amount"`
	// Insurance currency.
	Currency *string `form:"currency"`
	// Insurance company name.
	InsuranceCompanyName *string `form:"insurance_company_name"`
	// Type of insurance.
	InsuranceType *string `form:"insurance_type"`
}

// List of passengers.
type ChargeUpdatePaymentDetailsFlightDatumPassengerParams struct {
	// Passenger's full name.
	Name *string `form:"name"`
}

// Arrival details.
type ChargeUpdatePaymentDetailsFlightDatumSegmentArrivalParams struct {
	// Arrival airport IATA code.
	Airport *string `form:"airport"`
	// Arrival date/time.
	ArrivesAt *int64 `form:"arrives_at"`
	// Arrival city.
	City *string `form:"city"`
	// Arrival country.
	Country *string `form:"country"`
}

// Departure details.
type ChargeUpdatePaymentDetailsFlightDatumSegmentDepartureParams struct {
	// Departure airport IATA code.
	Airport *string `form:"airport"`
	// Departure city.
	City *string `form:"city"`
	// Departure country.
	Country *string `form:"country"`
	// Departure date/time.
	DepartsAt *int64 `form:"departs_at"`
}

// List of flight segments.
type ChargeUpdatePaymentDetailsFlightDatumSegmentParams struct {
	// Segment fare amount.
	Amount *int64 `form:"amount"`
	// Arrival details.
	Arrival *ChargeUpdatePaymentDetailsFlightDatumSegmentArrivalParams `form:"arrival"`
	// Airline carrier code.
	CarrierCode *string `form:"carrier_code"`
	// Carrier name.
	CarrierName *string `form:"carrier_name"`
	// Segment currency.
	Currency *string `form:"currency"`
	// Departure details.
	Departure *ChargeUpdatePaymentDetailsFlightDatumSegmentDepartureParams `form:"departure"`
	// Exchange ticket number.
	ExchangeTicketNumber *string `form:"exchange_ticket_number"`
	// Fare basis code.
	FareBasisCode *string `form:"fare_basis_code"`
	// Additional fees.
	Fees *int64 `form:"fees"`
	// Flight number.
	FlightNumber *string `form:"flight_number"`
	// Stopover indicator.
	IsStopOverIndicator *bool `form:"is_stop_over_indicator"`
	// Refundable ticket indicator.
	Refundable *bool `form:"refundable"`
	// Class of service.
	ServiceClass *string `form:"service_class"`
	// Tax amount for segment.
	TaxAmount *int64 `form:"tax_amount"`
	// Ticket number.
	TicketNumber *string `form:"ticket_number"`
}

// Discount details.
type ChargeUpdatePaymentDetailsFlightDatumTotalDiscountsParams struct {
	// Corporate client discount code.
	CorporateClientCode *string `form:"corporate_client_code"`
}

// Additional charges.
type ChargeUpdatePaymentDetailsFlightDatumTotalExtraChargeParams struct {
	// Amount of additional charges.
	Amount *int64 `form:"amount"`
	// Type of additional charges.
	Type *string `form:"type"`
}

// Array of tax details.
type ChargeUpdatePaymentDetailsFlightDatumTotalTaxTaxParams struct {
	// Tax amount.
	Amount *int64 `form:"amount"`
	// Tax rate.
	Rate *int64 `form:"rate"`
	// Type of tax.
	Type *string `form:"type"`
}

// Tax breakdown.
type ChargeUpdatePaymentDetailsFlightDatumTotalTaxParams struct {
	// Array of tax details.
	Taxes []*ChargeUpdatePaymentDetailsFlightDatumTotalTaxTaxParams `form:"taxes"`
}

// Total cost breakdown.
type ChargeUpdatePaymentDetailsFlightDatumTotalParams struct {
	// Total flight amount.
	Amount *int64 `form:"amount"`
	// Reason for credit.
	CreditReason *string `form:"credit_reason"`
	// Total currency.
	Currency *string `form:"currency"`
	// Discount details.
	Discounts *ChargeUpdatePaymentDetailsFlightDatumTotalDiscountsParams `form:"discounts"`
	// Additional charges.
	ExtraCharges []*ChargeUpdatePaymentDetailsFlightDatumTotalExtraChargeParams `form:"extra_charges"`
	// Tax breakdown.
	Tax *ChargeUpdatePaymentDetailsFlightDatumTotalTaxParams `form:"tax"`
}

// Flight data for this PaymentIntent.
type ChargeUpdatePaymentDetailsFlightDatumParams struct {
	// Affiliate details if applicable.
	Affiliate *ChargeUpdatePaymentDetailsFlightDatumAffiliateParams `form:"affiliate"`
	// Reservation reference.
	BookingNumber *string `form:"booking_number"`
	// Computerized reservation system used to make the reservation and purchase the ticket.
	ComputerizedReservationSystem *string `form:"computerized_reservation_system"`
	// Ticket restrictions.
	EndorsementsAndRestrictions *string `form:"endorsements_and_restrictions"`
	// List of insurances.
	Insurances []*ChargeUpdatePaymentDetailsFlightDatumInsuranceParams `form:"insurances"`
	// List of passengers.
	Passengers []*ChargeUpdatePaymentDetailsFlightDatumPassengerParams `form:"passengers"`
	// List of flight segments.
	Segments []*ChargeUpdatePaymentDetailsFlightDatumSegmentParams `form:"segments"`
	// Electronic ticket indicator.
	TicketElectronicallyIssuedIndicator *bool `form:"ticket_electronically_issued_indicator"`
	// Total cost breakdown.
	Total *ChargeUpdatePaymentDetailsFlightDatumTotalParams `form:"total"`
	// Type of flight transaction.
	TransactionType *string `form:"transaction_type"`
}

// Affiliate details for this purchase.
type ChargeUpdatePaymentDetailsLodgingAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Details of the recipient.
type ChargeUpdatePaymentDetailsLodgingDeliveryRecipientParams struct {
	// The email of the recipient the ticket is delivered to.
	Email *string `form:"email"`
	// The name of the recipient the ticket is delivered to.
	Name *string `form:"name"`
	// The phone number of the recipient the ticket is delivered to.
	Phone *string `form:"phone"`
}

// Delivery details for this purchase.
type ChargeUpdatePaymentDetailsLodgingDeliveryParams struct {
	// The delivery method for the payment
	Mode *string `form:"mode"`
	// Details of the recipient.
	Recipient *ChargeUpdatePaymentDetailsLodgingDeliveryRecipientParams `form:"recipient"`
}

// The details of the passengers in the travel reservation
type ChargeUpdatePaymentDetailsLodgingPassengerParams struct {
	// Full name of the person or entity on the lodging reservation.
	Name *string `form:"name"`
}

// Lodging reservation details for this PaymentIntent
type ChargeUpdatePaymentDetailsLodgingParams struct {
	// The lodging location's address.
	Address *AddressParams `form:"address"`
	// The number of adults on the booking
	Adults *int64 `form:"adults"`
	// Affiliate details for this purchase.
	Affiliate *ChargeUpdatePaymentDetailsLodgingAffiliateParams `form:"affiliate"`
	// The booking number associated with the lodging reservation.
	BookingNumber *string `form:"booking_number"`
	// The lodging category
	Category *string `form:"category"`
	// Lodging check-in time. Measured in seconds since the Unix epoch.
	CheckinAt *int64 `form:"checkin_at"`
	// Lodging check-out time. Measured in seconds since the Unix epoch.
	CheckoutAt *int64 `form:"checkout_at"`
	// The customer service phone number of the lodging company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// The daily lodging room rate.
	DailyRoomRateAmount *int64 `form:"daily_room_rate_amount"`
	// Delivery details for this purchase.
	Delivery *ChargeUpdatePaymentDetailsLodgingDeliveryParams `form:"delivery"`
	// List of additional charges being billed.
	ExtraCharges []*string `form:"extra_charges"`
	// Indicates whether the lodging location is compliant with the Fire Safety Act.
	FireSafetyActCompliance *bool `form:"fire_safety_act_compliance"`
	// The name of the lodging location.
	Name *string `form:"name"`
	// Indicates if the customer did not keep their booking while failing to cancel the reservation.
	NoShow *bool `form:"no_show"`
	// The number of rooms on the booking
	NumberOfRooms *int64 `form:"number_of_rooms"`
	// The details of the passengers in the travel reservation
	Passengers []*ChargeUpdatePaymentDetailsLodgingPassengerParams `form:"passengers"`
	// The phone number of the lodging location.
	PropertyPhoneNumber *string `form:"property_phone_number"`
	// The room class for this purchase.
	RoomClass *string `form:"room_class"`
	// The number of room nights
	RoomNights *int64 `form:"room_nights"`
	// The total tax amount associating with the room reservation.
	TotalRoomTaxAmount *int64 `form:"total_room_tax_amount"`
	// The total tax amount
	TotalTaxAmount *int64 `form:"total_tax_amount"`
}

// Accommodation details for the lodging.
type ChargeUpdatePaymentDetailsLodgingDatumAccommodationParams struct {
	// Type of accommodation.
	AccommodationType *string `form:"accommodation_type"`
	// Bed type.
	BedType *string `form:"bed_type"`
	// Daily accommodation rate in cents.
	DailyRateAmount *int64 `form:"daily_rate_amount"`
	// Number of nights.
	Nights *int64 `form:"nights"`
	// Number of rooms, cabanas, apartments, and so on.
	NumberOfRooms *int64 `form:"number_of_rooms"`
	// Rate type.
	RateType *string `form:"rate_type"`
	// Whether smoking is allowed.
	SmokingIndicator *bool `form:"smoking_indicator"`
}

// Affiliate details if applicable.
type ChargeUpdatePaymentDetailsLodgingDatumAffiliateParams struct {
	// Affiliate partner code.
	Code *string `form:"code"`
	// Affiliate partner name.
	Name *string `form:"name"`
}

// List of guests for the lodging.
type ChargeUpdatePaymentDetailsLodgingDatumGuestParams struct {
	// Guest's full name.
	Name *string `form:"name"`
}

// Host details for the lodging.
type ChargeUpdatePaymentDetailsLodgingDatumHostParams struct {
	// Address of the host.
	Address *AddressParams `form:"address"`
	// Host's country of domicile.
	CountryOfDomicile *string `form:"country_of_domicile"`
	// Reference number for the host.
	HostReference *string `form:"host_reference"`
	// Type of host.
	HostType *string `form:"host_type"`
	// Name of the lodging property or host.
	Name *string `form:"name"`
	// Total number of reservations for the host.
	NumberOfReservations *int64 `form:"number_of_reservations"`
	// Property phone number.
	PropertyPhoneNumber *string `form:"property_phone_number"`
	// Host's registration date.
	RegisteredAt *int64 `form:"registered_at"`
}

// List of insurances for the lodging.
type ChargeUpdatePaymentDetailsLodgingDatumInsuranceParams struct {
	// Price of the insurance coverage in cents.
	Amount *int64 `form:"amount"`
	// Currency of the insurance amount.
	Currency *string `form:"currency"`
	// Name of the insurance company.
	InsuranceCompanyName *string `form:"insurance_company_name"`
	// Type of insurance coverage.
	InsuranceType *string `form:"insurance_type"`
}

// Discount details for the lodging.
type ChargeUpdatePaymentDetailsLodgingDatumTotalDiscountsParams struct {
	// Corporate client discount code.
	CorporateClientCode *string `form:"corporate_client_code"`
	// Coupon code.
	Coupon *string `form:"coupon"`
}

// Additional charges for the lodging.
type ChargeUpdatePaymentDetailsLodgingDatumTotalExtraChargeParams struct {
	// Amount of the extra charge in cents.
	Amount *int64 `form:"amount"`
	// Type of extra charge.
	Type *string `form:"type"`
}

// Tax details.
type ChargeUpdatePaymentDetailsLodgingDatumTotalTaxTaxParams struct {
	// Tax amount in cents.
	Amount *int64 `form:"amount"`
	// Tax rate.
	Rate *int64 `form:"rate"`
	// Type of tax applied.
	Type *string `form:"type"`
}

// Tax breakdown for the lodging reservation.
type ChargeUpdatePaymentDetailsLodgingDatumTotalTaxParams struct {
	// Tax details.
	Taxes []*ChargeUpdatePaymentDetailsLodgingDatumTotalTaxTaxParams `form:"taxes"`
	// Indicates whether the transaction is tax exempt.
	TaxExemptIndicator *bool `form:"tax_exempt_indicator"`
}

// Total details for the lodging.
type ChargeUpdatePaymentDetailsLodgingDatumTotalParams struct {
	// Total price of the lodging reservation in cents.
	Amount *int64 `form:"amount"`
	// Cash advances in cents.
	CashAdvances *int64 `form:"cash_advances"`
	// Currency of the total amount.
	Currency *string `form:"currency"`
	// Discount details for the lodging.
	Discounts *ChargeUpdatePaymentDetailsLodgingDatumTotalDiscountsParams `form:"discounts"`
	// Additional charges for the lodging.
	ExtraCharges []*ChargeUpdatePaymentDetailsLodgingDatumTotalExtraChargeParams `form:"extra_charges"`
	// Prepaid amount in cents.
	PrepaidAmount *int64 `form:"prepaid_amount"`
	// Tax breakdown for the lodging reservation.
	Tax *ChargeUpdatePaymentDetailsLodgingDatumTotalTaxParams `form:"tax"`
}

// Lodging data for this PaymentIntent.
type ChargeUpdatePaymentDetailsLodgingDatumParams struct {
	// Accommodation details for the lodging.
	Accommodation *ChargeUpdatePaymentDetailsLodgingDatumAccommodationParams `form:"accommodation"`
	// Affiliate details if applicable.
	Affiliate *ChargeUpdatePaymentDetailsLodgingDatumAffiliateParams `form:"affiliate"`
	// Booking confirmation number for the lodging.
	BookingNumber *string `form:"booking_number"`
	// Check-in date.
	CheckinAt *int64 `form:"checkin_at"`
	// Check-out date.
	CheckoutAt *int64 `form:"checkout_at"`
	// Customer service phone number for the lodging company.
	CustomerServicePhoneNumber *string `form:"customer_service_phone_number"`
	// Whether the lodging is compliant with any hotel fire safety regulations.
	FireSafetyActComplianceIndicator *bool `form:"fire_safety_act_compliance_indicator"`
	// List of guests for the lodging.
	Guests []*ChargeUpdatePaymentDetailsLodgingDatumGuestParams `form:"guests"`
	// Host details for the lodging.
	Host *ChargeUpdatePaymentDetailsLodgingDatumHostParams `form:"host"`
	// List of insurances for the lodging.
	Insurances []*ChargeUpdatePaymentDetailsLodgingDatumInsuranceParams `form:"insurances"`
	// Whether the renter is a no-show.
	NoShowIndicator *bool `form:"no_show_indicator"`
	// Renter ID number for the lodging.
	RenterIDNumber *string `form:"renter_id_number"`
	// Renter name for the lodging.
	RenterName *string `form:"renter_name"`
	// Total details for the lodging.
	Total *ChargeUpdatePaymentDetailsLodgingDatumTotalParams `form:"total"`
}

// Affiliate details for this purchase.
type ChargeUpdatePaymentDetailsSubscriptionAffiliateParams struct {
	// The name of the affiliate that originated the purchase.
	Name *string `form:"name"`
}

// Subscription billing details for this purchase.
type ChargeUpdatePaymentDetailsSubscriptionBillingIntervalParams struct {
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	Count *int64 `form:"count"`
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
}

// Subscription details for this PaymentIntent
type ChargeUpdatePaymentDetailsSubscriptionParams struct {
	// Affiliate details for this purchase.
	Affiliate *ChargeUpdatePaymentDetailsSubscriptionAffiliateParams `form:"affiliate"`
	// Info whether the subscription will be auto renewed upon expiry.
	AutoRenewal *bool `form:"auto_renewal"`
	// Subscription billing details for this purchase.
	BillingInterval *ChargeUpdatePaymentDetailsSubscriptionBillingIntervalParams `form:"billing_interval"`
	// Subscription end time. Measured in seconds since the Unix epoch.
	EndsAt *int64 `form:"ends_at"`
	// Name of the product on subscription. e.g. Apple Music Subscription
	Name *string `form:"name"`
	// Subscription start time. Measured in seconds since the Unix epoch.
	StartsAt *int64 `form:"starts_at"`
}

// Provides industry-specific information about the charge.
type ChargeUpdatePaymentDetailsParams struct {
	// Car rental details for this PaymentIntent.
	CarRental *ChargeUpdatePaymentDetailsCarRentalParams `form:"car_rental"`
	// Car rental data for this PaymentIntent.
	CarRentalData []*ChargeUpdatePaymentDetailsCarRentalDatumParams `form:"car_rental_data"`
	// A unique value to identify the customer. This field is available only for card payments.
	//
	// This field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.
	CustomerReference *string `form:"customer_reference"`
	// Event details for this PaymentIntent
	EventDetails *ChargeUpdatePaymentDetailsEventDetailsParams `form:"event_details"`
	// Flight reservation details for this PaymentIntent
	Flight *ChargeUpdatePaymentDetailsFlightParams `form:"flight"`
	// Flight data for this PaymentIntent.
	FlightData []*ChargeUpdatePaymentDetailsFlightDatumParams `form:"flight_data"`
	// Lodging reservation details for this PaymentIntent
	Lodging *ChargeUpdatePaymentDetailsLodgingParams `form:"lodging"`
	// Lodging data for this PaymentIntent.
	LodgingData []*ChargeUpdatePaymentDetailsLodgingDatumParams `form:"lodging_data"`
	// A unique value assigned by the business to identify the transaction. Required for L2 and L3 rates.
	//
	// Required when the Payment Method Types array contains `card`, including when [automatic_payment_methods.enabled](https://docs.stripe.com/api/payment_intents/create#create_payment_intent-automatic_payment_methods-enabled) is set to `true`.
	//
	// For Cards, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks. For Klarna, this field is truncated to 255 characters and is visible to customers when they view the order in the Klarna app.
	OrderReference *string `form:"order_reference"`
	// Subscription details for this PaymentIntent
	Subscription *ChargeUpdatePaymentDetailsSubscriptionParams `form:"subscription"`
}

// Updates the specified charge by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
type ChargeUpdateParams struct {
	Params `form:"*"`
	// The ID of an existing customer that will be associated with this request. This field may only be updated if there is no existing associated customer with this charge.
	Customer *string `form:"customer"`
	// An arbitrary string which you can attach to a charge object. It is displayed when in the web interface alongside the charge. Note that if you use Stripe to send automatic email receipts to your customers, your receipt emails will include the `description` of the charge(s) that they are describing.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A set of key-value pairs you can attach to a charge giving information about its riskiness. If you believe a charge is fraudulent, include a `user_report` key with a value of `fraudulent`. If you believe a charge is safe, include a `user_report` key with a value of `safe`. Stripe will use the information you send to improve our fraud detection algorithms.
	FraudDetails *ChargeUpdateFraudDetailsParams `form:"fraud_details"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Provides industry-specific information about the charge.
	PaymentDetails *ChargeUpdatePaymentDetailsParams `form:"payment_details"`
	// This is the email address that the receipt for this charge will be sent to. If this field is updated, then a new email receipt will be sent to the updated address.
	ReceiptEmail *string `form:"receipt_email"`
	// Shipping information for the charge. Helps prevent fraud on charges for physical goods.
	Shipping *ShippingDetailsParams `form:"shipping"`
	// A string that identifies this transaction as part of a group. `transfer_group` may only be provided if it has not been set. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options) for details.
	TransferGroup *string `form:"transfer_group"`
}

// AddExpand appends a new field to expand.
func (p *ChargeUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *ChargeUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

type ChargeBillingDetails struct {
	// Billing address.
	Address *Address `json:"address"`
	// Email address.
	Email string `json:"email"`
	// Full name.
	Name string `json:"name"`
	// Billing phone number (including extension).
	Phone string `json:"phone"`
	// Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.
	TaxID string `json:"tax_id"`
}

// Information on fraud assessments for the charge.
type ChargeFraudDetails struct {
	// Assessments from Stripe. If set, the value is `fraudulent`.
	StripeReport ChargeFraudStripeReport `json:"stripe_report"`
	// Assessments reported by you. If set, possible values of are `safe` and `fraudulent`.
	UserReport ChargeFraudUserReport `json:"user_report"`
}
type ChargeLevel3LineItem struct {
	DiscountAmount     int64  `json:"discount_amount"`
	ProductCode        string `json:"product_code"`
	ProductDescription string `json:"product_description"`
	Quantity           int64  `json:"quantity"`
	TaxAmount          int64  `json:"tax_amount"`
	UnitCost           int64  `json:"unit_cost"`
}
type ChargeLevel3 struct {
	CustomerReference  string                  `json:"customer_reference"`
	LineItems          []*ChargeLevel3LineItem `json:"line_items"`
	MerchantReference  string                  `json:"merchant_reference"`
	ShippingAddressZip string                  `json:"shipping_address_zip"`
	ShippingAmount     int64                   `json:"shipping_amount"`
	ShippingFromZip    string                  `json:"shipping_from_zip"`
}

// The ID of the Radar rule that matched the payment, if applicable.
type ChargeOutcomeRule struct {
	// The action taken on the payment.
	Action string `json:"action"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The predicate to evaluate the payment against.
	Predicate string `json:"predicate"`
}

// Details about whether the payment was accepted, and why. See [understanding declines](https://docs.stripe.com/declines) for details.
type ChargeOutcome struct {
	// An enumerated value providing a more detailed explanation on [how to proceed with an error](https://docs.stripe.com/declines#retrying-issuer-declines).
	AdviceCode ChargeOutcomeAdviceCode `json:"advice_code"`
	// For charges declined by the network, a 2 digit code which indicates the advice returned by the network on how to proceed with an error.
	NetworkAdviceCode string `json:"network_advice_code"`
	// For charges declined by the network, an alphanumeric code which indicates the reason the charge failed.
	NetworkDeclineCode string `json:"network_decline_code"`
	// Possible values are `approved_by_network`, `declined_by_network`, `not_sent_to_network`, and `reversed_after_approval`. The value `reversed_after_approval` indicates the payment was [blocked by Stripe](https://docs.stripe.com/declines#blocked-payments) after bank authorization, and may temporarily appear as "pending" on a cardholder's statement.
	NetworkStatus string `json:"network_status"`
	// An enumerated value providing a more detailed explanation of the outcome's `type`. Charges blocked by Radar's default block rule have the value `highest_risk_level`. Charges placed in review by Radar's default review rule have the value `elevated_risk_level`. Charges blocked because the payment is unlikely to be authorized have the value `low_probability_of_authorization`. Charges authorized, blocked, or placed in review by custom rules have the value `rule`. See [understanding declines](https://docs.stripe.com/declines) for more details.
	Reason string `json:"reason"`
	// Stripe Radar's evaluation of the riskiness of the payment. Possible values for evaluated payments are `normal`, `elevated`, `highest`. For non-card payments, and card-based payments predating the public assignment of risk levels, this field will have the value `not_assessed`. In the event of an error in the evaluation, this field will have the value `unknown`. This field is only available with Radar.
	RiskLevel string `json:"risk_level"`
	// Stripe Radar's evaluation of the riskiness of the payment. Possible values for evaluated payments are between 0 and 100. For non-card payments, card-based payments predating the public assignment of risk scores, or in the event of an error during evaluation, this field will not be present. This field is only available with Radar for Fraud Teams.
	RiskScore int64 `json:"risk_score"`
	// The ID of the Radar rule that matched the payment, if applicable.
	Rule *ChargeOutcomeRule `json:"rule"`
	// A human-readable description of the outcome type and reason, designed for you (the recipient of the payment), not your customer.
	SellerMessage string `json:"seller_message"`
	// Possible values are `authorized`, `manual_review`, `issuer_declined`, `blocked`, and `invalid`. See [understanding declines](https://docs.stripe.com/declines) and [Radar reviews](https://docs.stripe.com/radar/reviews) for details.
	Type string `json:"type"`
}

// UnmarshalJSON handles deserialization of a ChargeOutcomeRule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *ChargeOutcomeRule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}
	type chargeOutcomeRule ChargeOutcomeRule
	var v chargeOutcomeRule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*c = ChargeOutcomeRule(v)
	return nil
}

type ChargePaymentMethodDetailsACHCreditTransfer struct {
	// Account number to transfer funds to.
	AccountNumber string `json:"account_number"`
	// Name of the bank associated with the routing number.
	BankName string `json:"bank_name"`
	// Routing transit number for the bank account to transfer funds to.
	RoutingNumber string `json:"routing_number"`
	// SWIFT code of the bank associated with the routing number.
	SwiftCode string `json:"swift_code"`
}
type ChargePaymentMethodDetailsACHDebit struct {
	// Type of entity that holds the account. This can be either `individual` or `company`.
	AccountHolderType BankAccountAccountHolderType `json:"account_holder_type"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// Routing transit number of the bank account.
	RoutingNumber string `json:"routing_number"`
}
type ChargePaymentMethodDetailsACSSDebit struct {
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.
	ExpectedDebitDate string `json:"expected_debit_date"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Institution number of the bank account
	InstitutionNumber string `json:"institution_number"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate string `json:"mandate"`
	// Transit number of the bank account.
	TransitNumber string `json:"transit_number"`
}
type ChargePaymentMethodDetailsAffirm struct {
	// ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.
	Location string `json:"location"`
	// ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.
	Reader string `json:"reader"`
	// The Affirm transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsAfterpayClearpay struct {
	// The Afterpay order ID associated with this payment intent.
	OrderID string `json:"order_id"`
	// Order identifier shown to the merchant in Afterpay's online portal.
	Reference string `json:"reference"`
}
type ChargePaymentMethodDetailsAlipay struct {
	// Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.
	BuyerID string `json:"buyer_id"`
	// Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Transaction ID of this particular Alipay transaction.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsAlmaInstallments struct {
	// The number of installments.
	Count int64 `json:"count"`
}
type ChargePaymentMethodDetailsAlma struct {
	Installments *ChargePaymentMethodDetailsAlmaInstallments `json:"installments"`
	// The Alma transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsAmazonPayFundingCard struct {
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand string `json:"brand"`
	// The [product code](https://stripe.com/docs/card-product-codes) that identifies the specific program or product associated with a card. (For internal use only and not typically available in standard API requests.)
	BrandProduct string `json:"brand_product"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding string `json:"funding"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
}
type ChargePaymentMethodDetailsAmazonPayFunding struct {
	Card *ChargePaymentMethodDetailsAmazonPayFundingCard `json:"card"`
	// funding type of the underlying payment method.
	Type ChargePaymentMethodDetailsAmazonPayFundingType `json:"type"`
}
type ChargePaymentMethodDetailsAmazonPay struct {
	Funding *ChargePaymentMethodDetailsAmazonPayFunding `json:"funding"`
	// The Amazon Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsAUBECSDebit struct {
	// Bank-State-Branch number of the bank account.
	BSBNumber string `json:"bsb_number"`
	// Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.
	ExpectedDebitDate string `json:"expected_debit_date"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate string `json:"mandate"`
}
type ChargePaymentMethodDetailsBACSDebit struct {
	// Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.
	ExpectedDebitDate string `json:"expected_debit_date"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate string `json:"mandate"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode string `json:"sort_code"`
}
type ChargePaymentMethodDetailsBancontact struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	BIC string `json:"bic"`
	// The ID of the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSEPADebit *PaymentMethod `json:"generated_sepa_debit"`
	// The mandate for the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSEPADebitMandate *Mandate `json:"generated_sepa_debit_mandate"`
	// Last four characters of the IBAN.
	IBANLast4 string `json:"iban_last4"`
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	// Can be one of `en`, `de`, `fr`, or `nl`
	PreferredLanguage string `json:"preferred_language"`
	// Owner's verified full name. Values are verified or provided by Bancontact directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsBillie struct {
	// The Billie transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsBLIK struct {
	// A unique and immutable identifier assigned by BLIK to every buyer.
	BuyerID string `json:"buyer_id"`
}
type ChargePaymentMethodDetailsBoleto struct {
	// The tax ID of the customer (CPF for individuals consumers or CNPJ for businesses consumers)
	TaxID string `json:"tax_id"`
}

// Check results by Card networks on Card address and CVC at time of payment.
type ChargePaymentMethodDetailsCardChecks struct {
	// If a address line1 was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressLine1Check ChargePaymentMethodDetailsCardChecksAddressLine1Check `json:"address_line1_check"`
	// If a address postal code was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressPostalCodeCheck ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheck `json:"address_postal_code_check"`
	// If a CVC was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	CVCCheck ChargePaymentMethodDetailsCardChecksCVCCheck `json:"cvc_check"`
}
type ChargePaymentMethodDetailsCardDecrementalAuthorization struct {
	// Indicates whether or not the decremental authorization feature is supported.
	Status ChargePaymentMethodDetailsCardDecrementalAuthorizationStatus `json:"status"`
}
type ChargePaymentMethodDetailsCardExtendedAuthorization struct {
	// Indicates whether or not the capture window is extended beyond the standard authorization.
	Status ChargePaymentMethodDetailsCardExtendedAuthorizationStatus `json:"status"`
}
type ChargePaymentMethodDetailsCardIncrementalAuthorization struct {
	// Indicates whether or not the incremental authorization feature is supported.
	Status ChargePaymentMethodDetailsCardIncrementalAuthorizationStatus `json:"status"`
}

// Installment details for this payment.
//
// For more information, see the [installments integration guide](https://docs.stripe.com/payments/installments).
type ChargePaymentMethodDetailsCardInstallments struct {
	// Installment plan selected for the payment.
	Plan *PaymentIntentPaymentMethodOptionsCardInstallmentsPlan `json:"plan"`
}
type ChargePaymentMethodDetailsCardMulticapture struct {
	// Indicates whether or not multiple captures are supported.
	Status ChargePaymentMethodDetailsCardMulticaptureStatus `json:"status"`
}

// If this card has network token credentials, this contains the details of the network token credentials.
type ChargePaymentMethodDetailsCardNetworkToken struct {
	// Indicates if Stripe used a network token, either user provided or Stripe managed when processing the transaction.
	Used bool `json:"used"`
}
type ChargePaymentMethodDetailsCardOvercapture struct {
	// The maximum amount that can be captured.
	MaximumAmountCapturable int64 `json:"maximum_amount_capturable"`
	// Indicates whether or not the authorized amount can be over-captured.
	Status ChargePaymentMethodDetailsCardOvercaptureStatus `json:"status"`
}
type ChargePaymentMethodDetailsCardPartialAuthorization struct {
	// Indicates whether the transaction requested for partial authorization feature and the authorization outcome.
	Status ChargePaymentMethodDetailsCardPartialAuthorizationStatus `json:"status"`
}

// Populated if this transaction used 3D Secure authentication.
type ChargePaymentMethodDetailsCardThreeDSecure struct {
	// For authenticated transactions: how the customer was authenticated by
	// the issuing bank.
	AuthenticationFlow ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	// The Electronic Commerce Indicator (ECI). A protocol-level field
	// indicating what degree of authentication was performed.
	ElectronicCommerceIndicator ChargePaymentMethodDetailsCardThreeDSecureElectronicCommerceIndicator `json:"electronic_commerce_indicator"`
	// The exemption requested via 3DS and accepted by the issuer at authentication time.
	ExemptionIndicator ChargePaymentMethodDetailsCardThreeDSecureExemptionIndicator `json:"exemption_indicator"`
	// Whether Stripe requested the value of `exemption_indicator` in the transaction. This will depend on
	// the outcome of Stripe's internal risk assessment.
	ExemptionIndicatorApplied bool `json:"exemption_indicator_applied"`
	// Indicates the outcome of 3D Secure authentication.
	Result ChargePaymentMethodDetailsCardThreeDSecureResult `json:"result"`
	// Additional information about why 3D Secure succeeded or failed based
	// on the `result`.
	ResultReason ChargePaymentMethodDetailsCardThreeDSecureResultReason `json:"result_reason"`
	// The 3D Secure 1 XID or 3D Secure 2 Directory Server Transaction ID
	// (dsTransId) for this payment.
	TransactionID string `json:"transaction_id"`
	// The version of 3D Secure that was used.
	Version string `json:"version"`
}
type ChargePaymentMethodDetailsCardWalletAmexExpressCheckout struct{}
type ChargePaymentMethodDetailsCardWalletApplePay struct{}
type ChargePaymentMethodDetailsCardWalletGooglePay struct{}
type ChargePaymentMethodDetailsCardWalletLink struct{}
type ChargePaymentMethodDetailsCardWalletMasterpass struct {
	// Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	BillingAddress *Address `json:"billing_address"`
	// Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Email string `json:"email"`
	// Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Name string `json:"name"`
	// Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	ShippingAddress *Address `json:"shipping_address"`
}
type ChargePaymentMethodDetailsCardWalletSamsungPay struct{}
type ChargePaymentMethodDetailsCardWalletVisaCheckout struct {
	// Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	BillingAddress *Address `json:"billing_address"`
	// Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Email string `json:"email"`
	// Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Name string `json:"name"`
	// Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	ShippingAddress *Address `json:"shipping_address"`
}

// If this Card is part of a card wallet, this contains the details of the card wallet.
type ChargePaymentMethodDetailsCardWallet struct {
	AmexExpressCheckout *ChargePaymentMethodDetailsCardWalletAmexExpressCheckout `json:"amex_express_checkout"`
	ApplePay            *ChargePaymentMethodDetailsCardWalletApplePay            `json:"apple_pay"`
	// (For tokenized numbers only.) The last four digits of the device account number.
	DynamicLast4 string                                          `json:"dynamic_last4"`
	GooglePay    *ChargePaymentMethodDetailsCardWalletGooglePay  `json:"google_pay"`
	Link         *ChargePaymentMethodDetailsCardWalletLink       `json:"link"`
	Masterpass   *ChargePaymentMethodDetailsCardWalletMasterpass `json:"masterpass"`
	SamsungPay   *ChargePaymentMethodDetailsCardWalletSamsungPay `json:"samsung_pay"`
	// The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, `visa_checkout`, or `link`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
	Type         PaymentMethodCardWalletType                       `json:"type"`
	VisaCheckout *ChargePaymentMethodDetailsCardWalletVisaCheckout `json:"visa_checkout"`
}
type ChargePaymentMethodDetailsCard struct {
	// The authorized amount.
	AmountAuthorized int64 `json:"amount_authorized"`
	// The latest amount intended to be authorized by this charge.
	AmountRequested int64 `json:"amount_requested"`
	// Authorization code on the charge.
	AuthorizationCode string `json:"authorization_code"`
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand PaymentMethodCardBrand `json:"brand"`
	// When using manual capture, a future timestamp at which the charge will be automatically refunded if uncaptured.
	CaptureBefore int64 `json:"capture_before"`
	// Check results by Card networks on Card address and CVC at time of payment.
	Checks *ChargePaymentMethodDetailsCardChecks `json:"checks"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country                  string                                                  `json:"country"`
	DecrementalAuthorization *ChargePaymentMethodDetailsCardDecrementalAuthorization `json:"decremental_authorization"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear               int64                                                `json:"exp_year"`
	ExtendedAuthorization *ChargePaymentMethodDetailsCardExtendedAuthorization `json:"extended_authorization"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *As of May 1, 2021, card fingerprint in India for Connect changed to allow two fingerprints for the same card---one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding                  CardFunding                                             `json:"funding"`
	IncrementalAuthorization *ChargePaymentMethodDetailsCardIncrementalAuthorization `json:"incremental_authorization"`
	// Installment details for this payment.
	//
	// For more information, see the [installments integration guide](https://docs.stripe.com/payments/installments).
	Installments *ChargePaymentMethodDetailsCardInstallments `json:"installments"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment or created by it.
	Mandate string `json:"mandate"`
	// True if this payment was marked as MOTO and out of scope for SCA.
	MOTO         bool                                        `json:"moto"`
	Multicapture *ChargePaymentMethodDetailsCardMulticapture `json:"multicapture"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network ChargePaymentMethodDetailsCardNetwork `json:"network"`
	// If this card has network token credentials, this contains the details of the network token credentials.
	NetworkToken *ChargePaymentMethodDetailsCardNetworkToken `json:"network_token"`
	// This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.
	NetworkTransactionID string                                              `json:"network_transaction_id"`
	Overcapture          *ChargePaymentMethodDetailsCardOvercapture          `json:"overcapture"`
	PartialAuthorization *ChargePaymentMethodDetailsCardPartialAuthorization `json:"partial_authorization"`
	// Status of a card based on the card issuer.
	RegulatedStatus ChargePaymentMethodDetailsCardRegulatedStatus `json:"regulated_status"`
	// Populated if this transaction used 3D Secure authentication.
	ThreeDSecure *ChargePaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
	// If this Card is part of a card wallet, this contains the details of the card wallet.
	Wallet *ChargePaymentMethodDetailsCardWallet `json:"wallet"`
	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description"`
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin"`
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer"`
}

// Details about payments collected offline.
type ChargePaymentMethodDetailsCardPresentOffline struct {
	// Time at which the payment was collected while offline
	StoredAt int64 `json:"stored_at"`
	// The method used to process this payment method offline. Only deferred is allowed.
	Type ChargePaymentMethodDetailsCardPresentOfflineType `json:"type"`
}

// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
type ChargePaymentMethodDetailsCardPresentReceipt struct {
	// The type of account being debited or credited
	AccountType ChargePaymentMethodDetailsCardPresentReceiptAccountType `json:"account_type"`
	// The Application Cryptogram, a unique value generated by the card to authenticate the transaction with issuers.
	ApplicationCryptogram string `json:"application_cryptogram"`
	// The Application Identifier (AID) on the card used to determine which networks are eligible to process the transaction. Referenced from EMV tag 9F12, data encoded on the card's chip.
	ApplicationPreferredName string `json:"application_preferred_name"`
	// Identifier for this transaction.
	AuthorizationCode string `json:"authorization_code"`
	// EMV tag 8A. A code returned by the card issuer.
	AuthorizationResponseCode string `json:"authorization_response_code"`
	// Describes the method used by the cardholder to verify ownership of the card. One of the following: `approval`, `failure`, `none`, `offline_pin`, `offline_pin_and_signature`, `online_pin`, or `signature`.
	CardholderVerificationMethod string `json:"cardholder_verification_method"`
	// Similar to the application_preferred_name, identifying the applications (AIDs) available on the card. Referenced from EMV tag 84.
	DedicatedFileName string `json:"dedicated_file_name"`
	// A 5-byte string that records the checks and validations that occur between the card and the terminal. These checks determine how the terminal processes the transaction and what risk tolerance is acceptable. Referenced from EMV Tag 95.
	TerminalVerificationResults string `json:"terminal_verification_results"`
	// An indication of which steps were completed during the card read process. Referenced from EMV Tag 9B.
	TransactionStatusInformation string `json:"transaction_status_information"`
}
type ChargePaymentMethodDetailsCardPresentWallet struct {
	// The type of mobile wallet, one of `apple_pay`, `google_pay`, `samsung_pay`, or `unknown`.
	Type ChargePaymentMethodDetailsCardPresentWalletType `json:"type"`
}
type ChargePaymentMethodDetailsCardPresent struct {
	// The authorized amount
	AmountAuthorized int64 `json:"amount_authorized"`
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand PaymentMethodCardBrand `json:"brand"`
	// The [product code](https://stripe.com/docs/card-product-codes) that identifies the specific program or product associated with a card.
	BrandProduct string `json:"brand_product"`
	// When using manual capture, a future timestamp after which the charge will be automatically refunded if uncaptured.
	CaptureBefore int64 `json:"capture_before"`
	// The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.
	CardholderName string `json:"cardholder_name"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// Authorization response cryptogram.
	EmvAuthData string `json:"emv_auth_data"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *As of May 1, 2021, card fingerprint in India for Connect changed to allow two fingerprints for the same card---one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding CardFunding `json:"funding"`
	// ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod.
	GeneratedCard string `json:"generated_card"`
	// Whether this [PaymentIntent](https://docs.stripe.com/api/payment_intents) is eligible for incremental authorizations. Request support using [request_incremental_authorization_support](https://docs.stripe.com/api/payment_intents/create#create_payment_intent-payment_method_options-card_present-request_incremental_authorization_support).
	IncrementalAuthorizationSupported bool `json:"incremental_authorization_supported"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network ChargePaymentMethodDetailsCardPresentNetwork `json:"network"`
	// This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.
	NetworkTransactionID string `json:"network_transaction_id"`
	// Details about payments collected offline.
	Offline *ChargePaymentMethodDetailsCardPresentOffline `json:"offline"`
	// Defines whether the authorized amount can be over-captured or not
	OvercaptureSupported bool `json:"overcapture_supported"`
	// The languages that the issuing bank recommends using for localizing any customer-facing text, as read from the card. Referenced from EMV tag 5F2D, data encoded on the card's chip.
	PreferredLocales []string `json:"preferred_locales"`
	// How card details were read in this transaction.
	ReadMethod string `json:"read_method"`
	// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
	Receipt *ChargePaymentMethodDetailsCardPresentReceipt `json:"receipt"`
	Wallet  *ChargePaymentMethodDetailsCardPresentWallet  `json:"wallet"`
	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description"`
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin"`
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer"`
}
type ChargePaymentMethodDetailsCashApp struct {
	// A unique and immutable identifier assigned by Cash App to every buyer.
	BuyerID string `json:"buyer_id"`
	// A public identifier for buyers using Cash App.
	Cashtag string `json:"cashtag"`
	// A unique and immutable identifier of payments assigned by Cash App
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsCrypto struct {
	// The wallet address of the customer.
	BuyerAddress string `json:"buyer_address"`
	// The blockchain network that the transaction was sent on.
	Network ChargePaymentMethodDetailsCryptoNetwork `json:"network"`
	// The token currency that the transaction was sent with.
	TokenCurrency ChargePaymentMethodDetailsCryptoTokenCurrency `json:"token_currency"`
	// The blockchain transaction hash of the crypto payment.
	TransactionHash string `json:"transaction_hash"`
}
type ChargePaymentMethodDetailsCustomerBalance struct{}
type ChargePaymentMethodDetailsEPS struct {
	// The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `deutsche_bank_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.
	Bank string `json:"bank"`
	// Owner's verified full name. Values are verified or provided by EPS directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	// EPS rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsFPX struct {
	// Account holder type, if provided. Can be one of `individual` or `company`.
	AccountHolderType PaymentMethodFPXAccountHolderType `json:"account_holder_type"`
	// The customer's bank. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, `pb_enterprise`, or `bank_of_china`.
	Bank string `json:"bank"`
	// Unique transaction id generated by FPX for every request from the merchant
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsGiropay struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	BIC string `json:"bic"`
	// Owner's verified full name. Values are verified or provided by Giropay directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	// Giropay rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsGopay struct{}
type ChargePaymentMethodDetailsGrabpay struct {
	// Unique transaction id generated by GrabPay
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsIDBankTransfer struct {
	// Account number of the bank account to transfer funds to.
	AccountNumber string `json:"account_number"`
	// Bank where the account is located.
	Bank ChargePaymentMethodDetailsIDBankTransferBank `json:"bank"`
	// Local bank code of the bank.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Merchant name and billing details name, for the customer to check for the correct merchant when performing the bank transfer.
	DisplayName string `json:"display_name"`
}
type ChargePaymentMethodDetailsIDEAL struct {
	// The customer's bank. Can be one of `abn_amro`, `asn_bank`, `bunq`, `buut`, `finom`, `handelsbanken`, `ing`, `knab`, `mollie`, `moneyou`, `n26`, `nn`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, `van_lanschot`, or `yoursafe`.
	Bank string `json:"bank"`
	// The Bank Identifier Code of the customer's bank.
	BIC string `json:"bic"`
	// The ID of the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSEPADebit *PaymentMethod `json:"generated_sepa_debit"`
	// The mandate for the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSEPADebitMandate *Mandate `json:"generated_sepa_debit_mandate"`
	// Last four characters of the IBAN.
	IBANLast4 string `json:"iban_last4"`
	// Unique transaction ID generated by iDEAL.
	TransactionID string `json:"transaction_id"`
	// Owner's verified full name. Values are verified or provided by iDEAL directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}

// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
type ChargePaymentMethodDetailsInteracPresentReceipt struct {
	// The type of account being debited or credited
	AccountType string `json:"account_type"`
	// The Application Cryptogram, a unique value generated by the card to authenticate the transaction with issuers.
	ApplicationCryptogram string `json:"application_cryptogram"`
	// The Application Identifier (AID) on the card used to determine which networks are eligible to process the transaction. Referenced from EMV tag 9F12, data encoded on the card's chip.
	ApplicationPreferredName string `json:"application_preferred_name"`
	// Identifier for this transaction.
	AuthorizationCode string `json:"authorization_code"`
	// EMV tag 8A. A code returned by the card issuer.
	AuthorizationResponseCode string `json:"authorization_response_code"`
	// Describes the method used by the cardholder to verify ownership of the card. One of the following: `approval`, `failure`, `none`, `offline_pin`, `offline_pin_and_signature`, `online_pin`, or `signature`.
	CardholderVerificationMethod string `json:"cardholder_verification_method"`
	// Similar to the application_preferred_name, identifying the applications (AIDs) available on the card. Referenced from EMV tag 84.
	DedicatedFileName string `json:"dedicated_file_name"`
	// A 5-byte string that records the checks and validations that occur between the card and the terminal. These checks determine how the terminal processes the transaction and what risk tolerance is acceptable. Referenced from EMV Tag 95.
	TerminalVerificationResults string `json:"terminal_verification_results"`
	// An indication of which steps were completed during the card read process. Referenced from EMV Tag 9B.
	TransactionStatusInformation string `json:"transaction_status_information"`
}
type ChargePaymentMethodDetailsInteracPresent struct {
	// Card brand. Can be `interac`, `mastercard` or `visa`.
	Brand string `json:"brand"`
	// The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.
	CardholderName string `json:"cardholder_name"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// Authorization response cryptogram.
	EmvAuthData string `json:"emv_auth_data"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *As of May 1, 2021, card fingerprint in India for Connect changed to allow two fingerprints for the same card---one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding string `json:"funding"`
	// ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod.
	GeneratedCard string `json:"generated_card"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network string `json:"network"`
	// This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.
	NetworkTransactionID string `json:"network_transaction_id"`
	// The languages that the issuing bank recommends using for localizing any customer-facing text, as read from the card. Referenced from EMV tag 5F2D, data encoded on the card's chip.
	PreferredLocales []string `json:"preferred_locales"`
	// How card details were read in this transaction.
	ReadMethod string `json:"read_method"`
	// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
	Receipt *ChargePaymentMethodDetailsInteracPresentReceipt `json:"receipt"`
	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description"`
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin"`
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer"`
}
type ChargePaymentMethodDetailsKakaoPay struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Kakao Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}

// The payer's address
type ChargePaymentMethodDetailsKlarnaPayerDetailsAddress struct {
	// The payer address country
	Country string `json:"country"`
}

// The payer details for this transaction.
type ChargePaymentMethodDetailsKlarnaPayerDetails struct {
	// The payer's address
	Address *ChargePaymentMethodDetailsKlarnaPayerDetailsAddress `json:"address"`
}
type ChargePaymentMethodDetailsKlarna struct {
	// The payer details for this transaction.
	PayerDetails *ChargePaymentMethodDetailsKlarnaPayerDetails `json:"payer_details"`
	// The Klarna payment method used for this transaction.
	// Can be one of `pay_later`, `pay_now`, `pay_with_financing`, or `pay_in_installments`
	PaymentMethodCategory ChargePaymentMethodDetailsKlarnaPaymentMethodCategory `json:"payment_method_category"`
	// Preferred language of the Klarna authorization page that the customer is redirected to.
	// Can be one of `de-AT`, `en-AT`, `nl-BE`, `fr-BE`, `en-BE`, `de-DE`, `en-DE`, `da-DK`, `en-DK`, `es-ES`, `en-ES`, `fi-FI`, `sv-FI`, `en-FI`, `en-GB`, `en-IE`, `it-IT`, `en-IT`, `nl-NL`, `en-NL`, `nb-NO`, `en-NO`, `sv-SE`, `en-SE`, `en-US`, `es-US`, `fr-FR`, `en-FR`, `cs-CZ`, `en-CZ`, `ro-RO`, `en-RO`, `el-GR`, `en-GR`, `en-AU`, `en-NZ`, `en-CA`, `fr-CA`, `pl-PL`, `en-PL`, `pt-PT`, `en-PT`, `de-CH`, `fr-CH`, `it-CH`, or `en-CH`
	PreferredLocale string `json:"preferred_locale"`
}

// If the payment succeeded, this contains the details of the convenience store where the payment was completed.
type ChargePaymentMethodDetailsKonbiniStore struct {
	// The name of the convenience store chain where the payment was completed.
	Chain ChargePaymentMethodDetailsKonbiniStoreChain `json:"chain"`
}
type ChargePaymentMethodDetailsKonbini struct {
	// If the payment succeeded, this contains the details of the convenience store where the payment was completed.
	Store *ChargePaymentMethodDetailsKonbiniStore `json:"store"`
}
type ChargePaymentMethodDetailsKrCard struct {
	// The local credit or debit card brand.
	Brand ChargePaymentMethodDetailsKrCardBrand `json:"brand"`
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The last four digits of the card. This may not be present for American Express cards.
	Last4 string `json:"last4"`
	// The Korean Card transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsLink struct {
	// Two-letter ISO code representing the funding source country beneath the Link payment.
	// You could use this attribute to get a sense of international fees.
	Country string `json:"country"`
}
type ChargePaymentMethodDetailsMbWay struct{}

// Internal card details
type ChargePaymentMethodDetailsMobilepayCard struct {
	// Brand of the card used in the transaction
	Brand string `json:"brand"`
	// Two-letter ISO code representing the country of the card
	Country string `json:"country"`
	// Two digit number representing the card's expiration month
	ExpMonth int64 `json:"exp_month"`
	// Two digit number representing the card's expiration year
	ExpYear int64 `json:"exp_year"`
	// The last 4 digits of the card
	Last4 string `json:"last4"`
}
type ChargePaymentMethodDetailsMobilepay struct {
	// Internal card details
	Card *ChargePaymentMethodDetailsMobilepayCard `json:"card"`
}
type ChargePaymentMethodDetailsMultibanco struct {
	// Entity number associated with this Multibanco payment.
	Entity string `json:"entity"`
	// Reference number associated with this Multibanco payment.
	Reference string `json:"reference"`
}
type ChargePaymentMethodDetailsNaverPay struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Naver Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsNzBankAccount struct {
	// The name on the bank account. Only present if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod's billing details.
	AccountHolderName string `json:"account_holder_name"`
	// The numeric code for the bank account's bank.
	BankCode string `json:"bank_code"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// The numeric code for the bank account's bank branch.
	BranchCode string `json:"branch_code"`
	// Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.
	ExpectedDebitDate string `json:"expected_debit_date"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// The suffix of the bank account number.
	Suffix string `json:"suffix"`
}
type ChargePaymentMethodDetailsOXXO struct {
	// OXXO reference number
	Number string `json:"number"`
}
type ChargePaymentMethodDetailsP24 struct {
	// The customer's bank. Can be one of `ing`, `citi_handlowy`, `tmobile_usbugi_bankowe`, `plus_bank`, `etransfer_pocztowy24`, `banki_spbdzielcze`, `bank_nowy_bfg_sa`, `getin_bank`, `velobank`, `blik`, `noble_pay`, `ideabank`, `envelobank`, `santander_przelew24`, `nest_przelew`, `mbank_mtransfer`, `inteligo`, `pbac_z_ipko`, `bnp_paribas`, `credit_agricole`, `toyota_bank`, `bank_pekao_sa`, `volkswagen_bank`, `bank_millennium`, `alior_bank`, or `boz`.
	Bank string `json:"bank"`
	// Unique reference for this Przelewy24 payment.
	Reference string `json:"reference"`
	// Owner's verified full name. Values are verified or provided by Przelewy24 directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	// Przelewy24 rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsPayByBank struct{}
type ChargePaymentMethodDetailsPayco struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Payco transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsPayNow struct {
	// ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.
	Location string `json:"location"`
	// ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.
	Reader string `json:"reader"`
	// Reference number associated with this PayNow payment
	Reference string `json:"reference"`
}

// The level of protection offered as defined by PayPal Seller Protection for Merchants, for this transaction.
type ChargePaymentMethodDetailsPaypalSellerProtection struct {
	// An array of conditions that are covered for the transaction, if applicable.
	DisputeCategories []ChargePaymentMethodDetailsPaypalSellerProtectionDisputeCategory `json:"dispute_categories"`
	// Indicates whether the transaction is eligible for PayPal's seller protection.
	Status ChargePaymentMethodDetailsPaypalSellerProtectionStatus `json:"status"`
}
type ChargePaymentMethodDetailsPaypal struct {
	// Two-letter ISO code representing the buyer's country. Values are provided by PayPal directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Country string `json:"country"`
	// Owner's email. Values are provided by PayPal directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	PayerEmail string `json:"payer_email"`
	// PayPal account PayerID. This identifier uniquely identifies the PayPal customer.
	PayerID string `json:"payer_id"`
	// Owner's full name. Values provided by PayPal directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	PayerName string `json:"payer_name"`
	// The level of protection offered as defined by PayPal Seller Protection for Merchants, for this transaction.
	SellerProtection *ChargePaymentMethodDetailsPaypalSellerProtection `json:"seller_protection"`
	// The shipping address for the customer, as supplied by the merchant at the point of payment
	// execution. This shipping address will not be updated if the merchant updates the shipping
	// address on the PaymentIntent after the PaymentIntent was successfully confirmed.
	Shipping *Address `json:"shipping"`
	// A unique ID generated by PayPal for this transaction.
	TransactionID string `json:"transaction_id"`
	// The shipping address for the customer, as supplied by the merchant at the point of payment
	// execution. This shipping address will not be updated if the merchant updates the shipping
	// address on the PaymentIntent after the PaymentIntent was successfully confirmed.
	VerifiedAddress *Address `json:"verified_address"`
	// Owner's verified email. Values are verified or provided by PayPal directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedEmail string `json:"verified_email"`
	// Owner's verified full name. Values are verified or provided by PayPal directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsPaypay struct{}
type ChargePaymentMethodDetailsPayto struct {
	// Bank-State-Branch number of the bank account.
	BSBNumber string `json:"bsb_number"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate string `json:"mandate"`
	// The PayID alias for the bank account.
	PayID string `json:"pay_id"`
}
type ChargePaymentMethodDetailsPix struct {
	// Unique transaction id generated by BCB
	BankTransactionID string `json:"bank_transaction_id"`
	// ID of the multi use Mandate generated by the PaymentIntent
	Mandate string `json:"mandate"`
}
type ChargePaymentMethodDetailsPromptPay struct {
	// Bill reference generated by PromptPay
	Reference string `json:"reference"`
}
type ChargePaymentMethodDetailsQris struct{}
type ChargePaymentMethodDetailsRechnung struct {
	// Payment portal URL.
	PaymentPortalURL string `json:"payment_portal_url"`
}
type ChargePaymentMethodDetailsRevolutPayFundingCard struct {
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand string `json:"brand"`
	// The [product code](https://stripe.com/docs/card-product-codes) that identifies the specific program or product associated with a card. (For internal use only and not typically available in standard API requests.)
	BrandProduct string `json:"brand_product"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding string `json:"funding"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
}
type ChargePaymentMethodDetailsRevolutPayFunding struct {
	Card *ChargePaymentMethodDetailsRevolutPayFundingCard `json:"card"`
	// funding type of the underlying payment method.
	Type ChargePaymentMethodDetailsRevolutPayFundingType `json:"type"`
}
type ChargePaymentMethodDetailsRevolutPay struct {
	Funding *ChargePaymentMethodDetailsRevolutPayFunding `json:"funding"`
	// The Revolut Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsSamsungPay struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Samsung Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsSatispay struct {
	// The Satispay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsSEPACreditTransfer struct {
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	BIC string `json:"bic"`
	// IBAN of the bank account to transfer funds to.
	IBAN string `json:"iban"`
}
type ChargePaymentMethodDetailsSEPADebit struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Branch code of bank associated with the bank account.
	BranchCode string `json:"branch_code"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
	// Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.
	ExpectedDebitDate string `json:"expected_debit_date"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four characters of the IBAN.
	Last4 string `json:"last4"`
	// Find the ID of the mandate used for this payment under the [payment_method_details.sepa_debit.mandate](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-sepa_debit-mandate) property on the Charge. Use this mandate ID to [retrieve the Mandate](https://docs.stripe.com/api/mandates/retrieve).
	Mandate string `json:"mandate"`
}
type ChargePaymentMethodDetailsShopeepay struct{}
type ChargePaymentMethodDetailsSofort struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	BIC string `json:"bic"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
	// The ID of the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSEPADebit *PaymentMethod `json:"generated_sepa_debit"`
	// The mandate for the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSEPADebitMandate *Mandate `json:"generated_sepa_debit_mandate"`
	// Last four characters of the IBAN.
	IBANLast4 string `json:"iban_last4"`
	// Preferred language of the SOFORT authorization page that the customer is redirected to.
	// Can be one of `de`, `en`, `es`, `fr`, `it`, `nl`, or `pl`
	PreferredLanguage string `json:"preferred_language"`
	// Owner's verified full name. Values are verified or provided by SOFORT directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsStripeAccount struct{}
type ChargePaymentMethodDetailsStripeBalance struct {
	// The connected account ID whose Stripe balance to use as the source of payment
	Account string `json:"account"`
	// The [source_type](https://docs.stripe.com/api/balance/balance_object#balance_object-available-source_types) of the balance
	SourceType ChargePaymentMethodDetailsStripeBalanceSourceType `json:"source_type"`
}
type ChargePaymentMethodDetailsSwish struct {
	// Uniquely identifies the payer's Swish account. You can use this attribute to check whether two Swish transactions were paid for by the same payer
	Fingerprint string `json:"fingerprint"`
	// Payer bank reference number for the payment
	PaymentReference string `json:"payment_reference"`
	// The last four digits of the Swish account phone number
	VerifiedPhoneLast4 string `json:"verified_phone_last4"`
}
type ChargePaymentMethodDetailsTWINT struct{}
type ChargePaymentMethodDetailsUSBankAccount struct {
	// Account holder type: individual or company.
	AccountHolderType ChargePaymentMethodDetailsUSBankAccountAccountHolderType `json:"account_holder_type"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType ChargePaymentMethodDetailsUSBankAccountAccountType `json:"account_type"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.
	ExpectedDebitDate string `json:"expected_debit_date"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate *Mandate `json:"mandate"`
	// Reference number to locate ACH payments with customer's bank.
	PaymentReference string `json:"payment_reference"`
	// Routing number of the bank account.
	RoutingNumber string `json:"routing_number"`
}
type ChargePaymentMethodDetailsWeChat struct{}
type ChargePaymentMethodDetailsWeChatPay struct {
	// Uniquely identifies this particular WeChat Pay account. You can use this attribute to check whether two WeChat accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.
	Location string `json:"location"`
	// ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.
	Reader string `json:"reader"`
	// Transaction ID of this particular WeChat Pay transaction.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsZip struct{}

// Details about the payment method at the time of the transaction.
type ChargePaymentMethodDetails struct {
	ACHCreditTransfer  *ChargePaymentMethodDetailsACHCreditTransfer  `json:"ach_credit_transfer"`
	ACHDebit           *ChargePaymentMethodDetailsACHDebit           `json:"ach_debit"`
	ACSSDebit          *ChargePaymentMethodDetailsACSSDebit          `json:"acss_debit"`
	Affirm             *ChargePaymentMethodDetailsAffirm             `json:"affirm"`
	AfterpayClearpay   *ChargePaymentMethodDetailsAfterpayClearpay   `json:"afterpay_clearpay"`
	Alipay             *ChargePaymentMethodDetailsAlipay             `json:"alipay"`
	Alma               *ChargePaymentMethodDetailsAlma               `json:"alma"`
	AmazonPay          *ChargePaymentMethodDetailsAmazonPay          `json:"amazon_pay"`
	AUBECSDebit        *ChargePaymentMethodDetailsAUBECSDebit        `json:"au_becs_debit"`
	BACSDebit          *ChargePaymentMethodDetailsBACSDebit          `json:"bacs_debit"`
	Bancontact         *ChargePaymentMethodDetailsBancontact         `json:"bancontact"`
	Billie             *ChargePaymentMethodDetailsBillie             `json:"billie"`
	BLIK               *ChargePaymentMethodDetailsBLIK               `json:"blik"`
	Boleto             *ChargePaymentMethodDetailsBoleto             `json:"boleto"`
	Card               *ChargePaymentMethodDetailsCard               `json:"card"`
	CardPresent        *ChargePaymentMethodDetailsCardPresent        `json:"card_present"`
	CashApp            *ChargePaymentMethodDetailsCashApp            `json:"cashapp"`
	Crypto             *ChargePaymentMethodDetailsCrypto             `json:"crypto"`
	CustomerBalance    *ChargePaymentMethodDetailsCustomerBalance    `json:"customer_balance"`
	EPS                *ChargePaymentMethodDetailsEPS                `json:"eps"`
	FPX                *ChargePaymentMethodDetailsFPX                `json:"fpx"`
	Giropay            *ChargePaymentMethodDetailsGiropay            `json:"giropay"`
	Gopay              *ChargePaymentMethodDetailsGopay              `json:"gopay"`
	Grabpay            *ChargePaymentMethodDetailsGrabpay            `json:"grabpay"`
	IDBankTransfer     *ChargePaymentMethodDetailsIDBankTransfer     `json:"id_bank_transfer"`
	IDEAL              *ChargePaymentMethodDetailsIDEAL              `json:"ideal"`
	InteracPresent     *ChargePaymentMethodDetailsInteracPresent     `json:"interac_present"`
	KakaoPay           *ChargePaymentMethodDetailsKakaoPay           `json:"kakao_pay"`
	Klarna             *ChargePaymentMethodDetailsKlarna             `json:"klarna"`
	Konbini            *ChargePaymentMethodDetailsKonbini            `json:"konbini"`
	KrCard             *ChargePaymentMethodDetailsKrCard             `json:"kr_card"`
	Link               *ChargePaymentMethodDetailsLink               `json:"link"`
	MbWay              *ChargePaymentMethodDetailsMbWay              `json:"mb_way"`
	Mobilepay          *ChargePaymentMethodDetailsMobilepay          `json:"mobilepay"`
	Multibanco         *ChargePaymentMethodDetailsMultibanco         `json:"multibanco"`
	NaverPay           *ChargePaymentMethodDetailsNaverPay           `json:"naver_pay"`
	NzBankAccount      *ChargePaymentMethodDetailsNzBankAccount      `json:"nz_bank_account"`
	OXXO               *ChargePaymentMethodDetailsOXXO               `json:"oxxo"`
	P24                *ChargePaymentMethodDetailsP24                `json:"p24"`
	PayByBank          *ChargePaymentMethodDetailsPayByBank          `json:"pay_by_bank"`
	Payco              *ChargePaymentMethodDetailsPayco              `json:"payco"`
	PayNow             *ChargePaymentMethodDetailsPayNow             `json:"paynow"`
	Paypal             *ChargePaymentMethodDetailsPaypal             `json:"paypal"`
	Paypay             *ChargePaymentMethodDetailsPaypay             `json:"paypay"`
	Payto              *ChargePaymentMethodDetailsPayto              `json:"payto"`
	Pix                *ChargePaymentMethodDetailsPix                `json:"pix"`
	PromptPay          *ChargePaymentMethodDetailsPromptPay          `json:"promptpay"`
	Qris               *ChargePaymentMethodDetailsQris               `json:"qris"`
	Rechnung           *ChargePaymentMethodDetailsRechnung           `json:"rechnung"`
	RevolutPay         *ChargePaymentMethodDetailsRevolutPay         `json:"revolut_pay"`
	SamsungPay         *ChargePaymentMethodDetailsSamsungPay         `json:"samsung_pay"`
	Satispay           *ChargePaymentMethodDetailsSatispay           `json:"satispay"`
	SEPACreditTransfer *ChargePaymentMethodDetailsSEPACreditTransfer `json:"sepa_credit_transfer"`
	SEPADebit          *ChargePaymentMethodDetailsSEPADebit          `json:"sepa_debit"`
	Shopeepay          *ChargePaymentMethodDetailsShopeepay          `json:"shopeepay"`
	Sofort             *ChargePaymentMethodDetailsSofort             `json:"sofort"`
	StripeAccount      *ChargePaymentMethodDetailsStripeAccount      `json:"stripe_account"`
	StripeBalance      *ChargePaymentMethodDetailsStripeBalance      `json:"stripe_balance"`
	Swish              *ChargePaymentMethodDetailsSwish              `json:"swish"`
	TWINT              *ChargePaymentMethodDetailsTWINT              `json:"twint"`
	// The type of transaction-specific details of the payment method used in the payment. See [PaymentMethod.type](https://docs.stripe.com/api/payment_methods/object#payment_method_object-type) for the full list of possible types.
	// An additional hash is included on `payment_method_details` with a name matching this value.
	// It contains information specific to the payment method.
	Type          ChargePaymentMethodDetailsType           `json:"type"`
	USBankAccount *ChargePaymentMethodDetailsUSBankAccount `json:"us_bank_account"`
	WeChat        *ChargePaymentMethodDetailsWeChat        `json:"wechat"`
	WeChatPay     *ChargePaymentMethodDetailsWeChatPay     `json:"wechat_pay"`
	Zip           *ChargePaymentMethodDetailsZip           `json:"zip"`
}
type ChargePresentmentDetails struct {
	// Amount intended to be collected by this payment, denominated in `presentment_currency`.
	PresentmentAmount int64 `json:"presentment_amount"`
	// Currency presented to the customer during payment.
	PresentmentCurrency Currency `json:"presentment_currency"`
}

// Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.
type ChargeRadarOptions struct {
	// A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session string `json:"session"`
}

// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.
type ChargeTransferData struct {
	// The amount transferred to the destination account, if specified. By default, the entire charge amount is transferred to the destination account.
	Amount int64 `json:"amount"`
	// ID of an existing, connected Stripe account to transfer funds to if `transfer_data` was specified in the charge request.
	Destination *Account `json:"destination"`
}

// The `Charge` object represents a single attempt to move money into your Stripe account.
// PaymentIntent confirmation is the most common way to create Charges, but [Account Debits](https://docs.stripe.com/connect/account-debits) may also create Charges.
// Some legacy payment flows create Charges directly, which is not recommended for new integrations.
type Charge struct {
	APIResource
	// Amount intended to be collected by this payment. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://docs.stripe.com/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount int64 `json:"amount"`
	// Amount in cents (or local equivalent) captured (can be less than the amount attribute on the charge if a partial capture was made).
	AmountCaptured int64 `json:"amount_captured"`
	// Amount in cents (or local equivalent) refunded (can be less than the amount attribute on the charge if a partial refund was issued).
	AmountRefunded int64 `json:"amount_refunded"`
	// ID of the Connect application that created the charge.
	Application *Application `json:"application"`
	// The application fee (if any) for the charge. [See the Connect documentation](https://docs.stripe.com/connect/direct-charges#collect-fees) for details.
	ApplicationFee *ApplicationFee `json:"application_fee"`
	// The amount of the application fee (if any) requested for the charge. [See the Connect documentation](https://docs.stripe.com/connect/direct-charges#collect-fees) for details.
	ApplicationFeeAmount int64 `json:"application_fee_amount"`
	// Authorization code on the charge.
	AuthorizationCode string `json:"authorization_code"`
	// ID of the balance transaction that describes the impact of this charge on your account balance (not including refunds or disputes).
	BalanceTransaction *BalanceTransaction   `json:"balance_transaction"`
	BillingDetails     *ChargeBillingDetails `json:"billing_details"`
	// The full statement descriptor that is passed to card networks, and that is displayed on your customers' credit card and bank statements. Allows you to see what the statement descriptor looks like after the static and dynamic portions are combined. This value only exists for card payments.
	CalculatedStatementDescriptor string `json:"calculated_statement_descriptor"`
	// If the charge was created without capturing, this Boolean represents whether it is still uncaptured or has since been captured.
	Captured bool `json:"captured"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// ID of the customer this charge is for if one exists.
	Customer *Customer `json:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Whether the charge has been disputed.
	Disputed bool `json:"disputed"`
	// ID of the balance transaction that describes the reversal of the balance on your account due to payment failure.
	FailureBalanceTransaction *BalanceTransaction `json:"failure_balance_transaction"`
	// Error code explaining reason for charge failure if available (see [the errors section](https://docs.stripe.com/error-codes) for a list of codes).
	FailureCode string `json:"failure_code"`
	// Message to user further explaining reason for charge failure if available.
	FailureMessage string `json:"failure_message"`
	// Information on fraud assessments for the charge.
	FraudDetails *ChargeFraudDetails `json:"fraud_details"`
	// Unique identifier for the object.
	ID     string        `json:"id"`
	Level3 *ChargeLevel3 `json:"level3"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account (if any) the charge was made on behalf of without triggering an automatic transfer. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// Details about whether the payment was accepted, and why. See [understanding declines](https://docs.stripe.com/declines) for details.
	Outcome *ChargeOutcome `json:"outcome"`
	// `true` if the charge succeeded, or was successfully authorized for later capture.
	Paid bool `json:"paid"`
	// ID of the PaymentIntent associated with this charge, if one exists.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
	// ID of the payment method used in this charge.
	PaymentMethod string `json:"payment_method"`
	// Details about the payment method at the time of the transaction.
	PaymentMethodDetails *ChargePaymentMethodDetails `json:"payment_method_details"`
	PresentmentDetails   *ChargePresentmentDetails   `json:"presentment_details"`
	// Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.
	RadarOptions *ChargeRadarOptions `json:"radar_options"`
	// This is the email address that the receipt for this charge was sent to.
	ReceiptEmail string `json:"receipt_email"`
	// This is the transaction number that appears on email receipts sent for this charge. This attribute will be `null` until a receipt has been sent.
	ReceiptNumber string `json:"receipt_number"`
	// This is the URL to view the receipt for this charge. The receipt is kept up-to-date to the latest state of the charge, including any refunds. If the charge is for an Invoice, the receipt will be stylized as an Invoice receipt.
	ReceiptURL string `json:"receipt_url"`
	// Whether the charge has been fully refunded. If the charge is only partially refunded, this attribute will still be false.
	Refunded bool `json:"refunded"`
	// A list of refunds that have been applied to the charge.
	Refunds *RefundList `json:"refunds"`
	// ID of the review associated with this charge if one exists.
	Review *Review `json:"review"`
	// Shipping information for the charge.
	Shipping *ShippingDetails `json:"shipping"`
	// This is a legacy field that will be removed in the future. It contains the Source, Card, or BankAccount object used for the charge. For details about the payment method used for this charge, refer to `payment_method` or `payment_method_details` instead.
	Source *PaymentSource `json:"source"`
	// The transfer ID which created this charge. Only present if the charge came from another Stripe account. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.
	SourceTransfer *Transfer `json:"source_transfer"`
	// For a non-card charge, text that appears on the customer's statement as the statement descriptor. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	//
	// For a card charge, this value is ignored unless you don't specify a `statement_descriptor_suffix`, in which case this value is used as the suffix.
	StatementDescriptor string `json:"statement_descriptor"`
	// Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement. If the account has no prefix value, the suffix is concatenated to the account's statement descriptor.
	StatementDescriptorSuffix string `json:"statement_descriptor_suffix"`
	// The status of the payment is either `succeeded`, `pending`, or `failed`.
	Status ChargeStatus `json:"status"`
	// ID of the transfer to the `destination` account (only applicable if the charge was created using the `destination` parameter).
	Transfer *Transfer `json:"transfer"`
	// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.
	TransferData *ChargeTransferData `json:"transfer_data"`
	// A string that identifies this transaction as part of a group. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options) for details.
	TransferGroup string `json:"transfer_group"`
}

// ChargeList is a list of Charges as retrieved from a list endpoint.
type ChargeList struct {
	APIResource
	ListMeta
	Data []*Charge `json:"data"`
}

// ChargeSearchResult is a list of Charge search results as retrieved from a search endpoint.
type ChargeSearchResult struct {
	APIResource
	SearchMeta
	Data []*Charge `json:"data"`
}

// UnmarshalJSON handles deserialization of a Charge.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Charge) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type charge Charge
	var v charge
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Charge(v)
	return nil
}
