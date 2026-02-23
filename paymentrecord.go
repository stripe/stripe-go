//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Indicates whether the customer was present in your checkout flow during this payment.
type PaymentRecordCustomerPresence string

// List of values that PaymentRecordCustomerPresence can take
const (
	PaymentRecordCustomerPresenceOffSession PaymentRecordCustomerPresence = "off_session"
	PaymentRecordCustomerPresenceOnSession  PaymentRecordCustomerPresence = "on_session"
)

// Type of entity that holds the account. This can be either `individual` or `company`.
type PaymentRecordPaymentMethodDetailsACHDebitAccountHolderType string

// List of values that PaymentRecordPaymentMethodDetailsACHDebitAccountHolderType can take
const (
	PaymentRecordPaymentMethodDetailsACHDebitAccountHolderTypeCompany    PaymentRecordPaymentMethodDetailsACHDebitAccountHolderType = "company"
	PaymentRecordPaymentMethodDetailsACHDebitAccountHolderTypeIndividual PaymentRecordPaymentMethodDetailsACHDebitAccountHolderType = "individual"
)

// funding type of the underlying payment method.
type PaymentRecordPaymentMethodDetailsAmazonPayFundingType string

// List of values that PaymentRecordPaymentMethodDetailsAmazonPayFundingType can take
const (
	PaymentRecordPaymentMethodDetailsAmazonPayFundingTypeCard PaymentRecordPaymentMethodDetailsAmazonPayFundingType = "card"
)

// Preferred language of the Bancontact authorization page that the customer is redirected to. Can be one of `en`, `de`, `fr`, or `nl`
type PaymentRecordPaymentMethodDetailsBancontactPreferredLanguage string

// List of values that PaymentRecordPaymentMethodDetailsBancontactPreferredLanguage can take
const (
	PaymentRecordPaymentMethodDetailsBancontactPreferredLanguageDE PaymentRecordPaymentMethodDetailsBancontactPreferredLanguage = "de"
	PaymentRecordPaymentMethodDetailsBancontactPreferredLanguageEN PaymentRecordPaymentMethodDetailsBancontactPreferredLanguage = "en"
	PaymentRecordPaymentMethodDetailsBancontactPreferredLanguageFR PaymentRecordPaymentMethodDetailsBancontactPreferredLanguage = "fr"
	PaymentRecordPaymentMethodDetailsBancontactPreferredLanguageNL PaymentRecordPaymentMethodDetailsBancontactPreferredLanguage = "nl"
)

// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
type PaymentRecordPaymentMethodDetailsCardBrand string

// List of values that PaymentRecordPaymentMethodDetailsCardBrand can take
const (
	PaymentRecordPaymentMethodDetailsCardBrandAmex            PaymentRecordPaymentMethodDetailsCardBrand = "amex"
	PaymentRecordPaymentMethodDetailsCardBrandCartesBancaires PaymentRecordPaymentMethodDetailsCardBrand = "cartes_bancaires"
	PaymentRecordPaymentMethodDetailsCardBrandDiners          PaymentRecordPaymentMethodDetailsCardBrand = "diners"
	PaymentRecordPaymentMethodDetailsCardBrandDiscover        PaymentRecordPaymentMethodDetailsCardBrand = "discover"
	PaymentRecordPaymentMethodDetailsCardBrandEFTPOSAU        PaymentRecordPaymentMethodDetailsCardBrand = "eftpos_au"
	PaymentRecordPaymentMethodDetailsCardBrandInterac         PaymentRecordPaymentMethodDetailsCardBrand = "interac"
	PaymentRecordPaymentMethodDetailsCardBrandJCB             PaymentRecordPaymentMethodDetailsCardBrand = "jcb"
	PaymentRecordPaymentMethodDetailsCardBrandLink            PaymentRecordPaymentMethodDetailsCardBrand = "link"
	PaymentRecordPaymentMethodDetailsCardBrandMastercard      PaymentRecordPaymentMethodDetailsCardBrand = "mastercard"
	PaymentRecordPaymentMethodDetailsCardBrandUnionpay        PaymentRecordPaymentMethodDetailsCardBrand = "unionpay"
	PaymentRecordPaymentMethodDetailsCardBrandUnknown         PaymentRecordPaymentMethodDetailsCardBrand = "unknown"
	PaymentRecordPaymentMethodDetailsCardBrandVisa            PaymentRecordPaymentMethodDetailsCardBrand = "visa"
)

// If you provide a value for `address.line1`, the check result is one of `pass`, `fail`, `unavailable`, or `unchecked`.
type PaymentRecordPaymentMethodDetailsCardChecksAddressLine1Check string

// List of values that PaymentRecordPaymentMethodDetailsCardChecksAddressLine1Check can take
const (
	PaymentRecordPaymentMethodDetailsCardChecksAddressLine1CheckFail        PaymentRecordPaymentMethodDetailsCardChecksAddressLine1Check = "fail"
	PaymentRecordPaymentMethodDetailsCardChecksAddressLine1CheckPass        PaymentRecordPaymentMethodDetailsCardChecksAddressLine1Check = "pass"
	PaymentRecordPaymentMethodDetailsCardChecksAddressLine1CheckUnavailable PaymentRecordPaymentMethodDetailsCardChecksAddressLine1Check = "unavailable"
	PaymentRecordPaymentMethodDetailsCardChecksAddressLine1CheckUnchecked   PaymentRecordPaymentMethodDetailsCardChecksAddressLine1Check = "unchecked"
)

// If you provide a address postal code, the check result is one of `pass`, `fail`, `unavailable`, or `unchecked`.
type PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck string

// List of values that PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck can take
const (
	PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckFail        PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "fail"
	PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckPass        PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "pass"
	PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckUnavailable PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "unavailable"
	PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckUnchecked   PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "unchecked"
)

// If you provide a CVC, the check results is one of `pass`, `fail`, `unavailable`, or `unchecked`.
type PaymentRecordPaymentMethodDetailsCardChecksCVCCheck string

// List of values that PaymentRecordPaymentMethodDetailsCardChecksCVCCheck can take
const (
	PaymentRecordPaymentMethodDetailsCardChecksCVCCheckFail        PaymentRecordPaymentMethodDetailsCardChecksCVCCheck = "fail"
	PaymentRecordPaymentMethodDetailsCardChecksCVCCheckPass        PaymentRecordPaymentMethodDetailsCardChecksCVCCheck = "pass"
	PaymentRecordPaymentMethodDetailsCardChecksCVCCheckUnavailable PaymentRecordPaymentMethodDetailsCardChecksCVCCheck = "unavailable"
	PaymentRecordPaymentMethodDetailsCardChecksCVCCheckUnchecked   PaymentRecordPaymentMethodDetailsCardChecksCVCCheck = "unchecked"
)

// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
type PaymentRecordPaymentMethodDetailsCardFunding string

// List of values that PaymentRecordPaymentMethodDetailsCardFunding can take
const (
	PaymentRecordPaymentMethodDetailsCardFundingCredit  PaymentRecordPaymentMethodDetailsCardFunding = "credit"
	PaymentRecordPaymentMethodDetailsCardFundingDebit   PaymentRecordPaymentMethodDetailsCardFunding = "debit"
	PaymentRecordPaymentMethodDetailsCardFundingPrepaid PaymentRecordPaymentMethodDetailsCardFunding = "prepaid"
	PaymentRecordPaymentMethodDetailsCardFundingUnknown PaymentRecordPaymentMethodDetailsCardFunding = "unknown"
)

// For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card. One of `month`.
type PaymentRecordPaymentMethodDetailsCardInstallmentsPlanInterval string

// List of values that PaymentRecordPaymentMethodDetailsCardInstallmentsPlanInterval can take
const (
	PaymentRecordPaymentMethodDetailsCardInstallmentsPlanIntervalMonth PaymentRecordPaymentMethodDetailsCardInstallmentsPlanInterval = "month"
)

// Type of installment plan, one of `fixed_count`, `revolving`, or `bonus`.
type PaymentRecordPaymentMethodDetailsCardInstallmentsPlanType string

// List of values that PaymentRecordPaymentMethodDetailsCardInstallmentsPlanType can take
const (
	PaymentRecordPaymentMethodDetailsCardInstallmentsPlanTypeBonus      PaymentRecordPaymentMethodDetailsCardInstallmentsPlanType = "bonus"
	PaymentRecordPaymentMethodDetailsCardInstallmentsPlanTypeFixedCount PaymentRecordPaymentMethodDetailsCardInstallmentsPlanType = "fixed_count"
	PaymentRecordPaymentMethodDetailsCardInstallmentsPlanTypeRevolving  PaymentRecordPaymentMethodDetailsCardInstallmentsPlanType = "revolving"
)

// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
type PaymentRecordPaymentMethodDetailsCardNetwork string

// List of values that PaymentRecordPaymentMethodDetailsCardNetwork can take
const (
	PaymentRecordPaymentMethodDetailsCardNetworkAmex            PaymentRecordPaymentMethodDetailsCardNetwork = "amex"
	PaymentRecordPaymentMethodDetailsCardNetworkCartesBancaires PaymentRecordPaymentMethodDetailsCardNetwork = "cartes_bancaires"
	PaymentRecordPaymentMethodDetailsCardNetworkDiners          PaymentRecordPaymentMethodDetailsCardNetwork = "diners"
	PaymentRecordPaymentMethodDetailsCardNetworkDiscover        PaymentRecordPaymentMethodDetailsCardNetwork = "discover"
	PaymentRecordPaymentMethodDetailsCardNetworkEFTPOSAU        PaymentRecordPaymentMethodDetailsCardNetwork = "eftpos_au"
	PaymentRecordPaymentMethodDetailsCardNetworkInterac         PaymentRecordPaymentMethodDetailsCardNetwork = "interac"
	PaymentRecordPaymentMethodDetailsCardNetworkJCB             PaymentRecordPaymentMethodDetailsCardNetwork = "jcb"
	PaymentRecordPaymentMethodDetailsCardNetworkLink            PaymentRecordPaymentMethodDetailsCardNetwork = "link"
	PaymentRecordPaymentMethodDetailsCardNetworkMastercard      PaymentRecordPaymentMethodDetailsCardNetwork = "mastercard"
	PaymentRecordPaymentMethodDetailsCardNetworkUnionpay        PaymentRecordPaymentMethodDetailsCardNetwork = "unionpay"
	PaymentRecordPaymentMethodDetailsCardNetworkUnknown         PaymentRecordPaymentMethodDetailsCardNetwork = "unknown"
	PaymentRecordPaymentMethodDetailsCardNetworkVisa            PaymentRecordPaymentMethodDetailsCardNetwork = "visa"
)

// The transaction type that was passed for an off-session, Merchant-Initiated transaction, one of `recurring` or `unscheduled`.
type PaymentRecordPaymentMethodDetailsCardStoredCredentialUsage string

// List of values that PaymentRecordPaymentMethodDetailsCardStoredCredentialUsage can take
const (
	PaymentRecordPaymentMethodDetailsCardStoredCredentialUsageRecurring   PaymentRecordPaymentMethodDetailsCardStoredCredentialUsage = "recurring"
	PaymentRecordPaymentMethodDetailsCardStoredCredentialUsageUnscheduled PaymentRecordPaymentMethodDetailsCardStoredCredentialUsage = "unscheduled"
)

// For authenticated transactions: Indicates how the issuing bank authenticated the customer.
type PaymentRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that PaymentRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take
const (
	PaymentRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    PaymentRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless PaymentRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// Indicates the outcome of 3D Secure authentication.
type PaymentRecordPaymentMethodDetailsCardThreeDSecureResult string

// List of values that PaymentRecordPaymentMethodDetailsCardThreeDSecureResult can take
const (
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged PaymentRecordPaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultAuthenticated       PaymentRecordPaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultExempted            PaymentRecordPaymentMethodDetailsCardThreeDSecureResult = "exempted"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultFailed              PaymentRecordPaymentMethodDetailsCardThreeDSecureResult = "failed"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultNotSupported        PaymentRecordPaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultProcessingError     PaymentRecordPaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// Additional information about why 3D Secure succeeded or failed, based on the `result`.
type PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReason can take
const (
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReasonBypassed            PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReasonCanceled            PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReasonRejected            PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

// The version of 3D Secure that was used.
type PaymentRecordPaymentMethodDetailsCardThreeDSecureVersion string

// List of values that PaymentRecordPaymentMethodDetailsCardThreeDSecureVersion can take
const (
	PaymentRecordPaymentMethodDetailsCardThreeDSecureVersion102 PaymentRecordPaymentMethodDetailsCardThreeDSecureVersion = "1.0.2"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureVersion210 PaymentRecordPaymentMethodDetailsCardThreeDSecureVersion = "2.1.0"
	PaymentRecordPaymentMethodDetailsCardThreeDSecureVersion220 PaymentRecordPaymentMethodDetailsCardThreeDSecureVersion = "2.2.0"
)

// The method used to process this payment method offline. Only deferred is allowed.
type PaymentRecordPaymentMethodDetailsCardPresentOfflineType string

// List of values that PaymentRecordPaymentMethodDetailsCardPresentOfflineType can take
const (
	PaymentRecordPaymentMethodDetailsCardPresentOfflineTypeDeferred PaymentRecordPaymentMethodDetailsCardPresentOfflineType = "deferred"
)

// How card details were read in this transaction.
type PaymentRecordPaymentMethodDetailsCardPresentReadMethod string

// List of values that PaymentRecordPaymentMethodDetailsCardPresentReadMethod can take
const (
	PaymentRecordPaymentMethodDetailsCardPresentReadMethodContactEmv               PaymentRecordPaymentMethodDetailsCardPresentReadMethod = "contact_emv"
	PaymentRecordPaymentMethodDetailsCardPresentReadMethodContactlessEmv           PaymentRecordPaymentMethodDetailsCardPresentReadMethod = "contactless_emv"
	PaymentRecordPaymentMethodDetailsCardPresentReadMethodContactlessMagstripeMode PaymentRecordPaymentMethodDetailsCardPresentReadMethod = "contactless_magstripe_mode"
	PaymentRecordPaymentMethodDetailsCardPresentReadMethodMagneticStripeFallback   PaymentRecordPaymentMethodDetailsCardPresentReadMethod = "magnetic_stripe_fallback"
	PaymentRecordPaymentMethodDetailsCardPresentReadMethodMagneticStripeTrack2     PaymentRecordPaymentMethodDetailsCardPresentReadMethod = "magnetic_stripe_track2"
)

// The type of account being debited or credited
type PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountType string

// List of values that PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountType can take
const (
	PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountTypeChecking PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountType = "checking"
	PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountTypeCredit   PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountType = "credit"
	PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountTypePrepaid  PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountType = "prepaid"
	PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountTypeUnknown  PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountType = "unknown"
)

// The type of mobile wallet, one of `apple_pay`, `google_pay`, `samsung_pay`, or `unknown`.
type PaymentRecordPaymentMethodDetailsCardPresentWalletType string

// List of values that PaymentRecordPaymentMethodDetailsCardPresentWalletType can take
const (
	PaymentRecordPaymentMethodDetailsCardPresentWalletTypeApplePay   PaymentRecordPaymentMethodDetailsCardPresentWalletType = "apple_pay"
	PaymentRecordPaymentMethodDetailsCardPresentWalletTypeGooglePay  PaymentRecordPaymentMethodDetailsCardPresentWalletType = "google_pay"
	PaymentRecordPaymentMethodDetailsCardPresentWalletTypeSamsungPay PaymentRecordPaymentMethodDetailsCardPresentWalletType = "samsung_pay"
	PaymentRecordPaymentMethodDetailsCardPresentWalletTypeUnknown    PaymentRecordPaymentMethodDetailsCardPresentWalletType = "unknown"
)

// The blockchain network that the transaction was sent on.
type PaymentRecordPaymentMethodDetailsCryptoNetwork string

// List of values that PaymentRecordPaymentMethodDetailsCryptoNetwork can take
const (
	PaymentRecordPaymentMethodDetailsCryptoNetworkBase     PaymentRecordPaymentMethodDetailsCryptoNetwork = "base"
	PaymentRecordPaymentMethodDetailsCryptoNetworkEthereum PaymentRecordPaymentMethodDetailsCryptoNetwork = "ethereum"
	PaymentRecordPaymentMethodDetailsCryptoNetworkPolygon  PaymentRecordPaymentMethodDetailsCryptoNetwork = "polygon"
	PaymentRecordPaymentMethodDetailsCryptoNetworkSolana   PaymentRecordPaymentMethodDetailsCryptoNetwork = "solana"
)

// The token currency that the transaction was sent with.
type PaymentRecordPaymentMethodDetailsCryptoTokenCurrency string

// List of values that PaymentRecordPaymentMethodDetailsCryptoTokenCurrency can take
const (
	PaymentRecordPaymentMethodDetailsCryptoTokenCurrencyUsdc PaymentRecordPaymentMethodDetailsCryptoTokenCurrency = "usdc"
	PaymentRecordPaymentMethodDetailsCryptoTokenCurrencyUsdg PaymentRecordPaymentMethodDetailsCryptoTokenCurrency = "usdg"
	PaymentRecordPaymentMethodDetailsCryptoTokenCurrencyUsdp PaymentRecordPaymentMethodDetailsCryptoTokenCurrency = "usdp"
)

// The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `deutsche_bank_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.
type PaymentRecordPaymentMethodDetailsEPSBank string

// List of values that PaymentRecordPaymentMethodDetailsEPSBank can take
const (
	PaymentRecordPaymentMethodDetailsEPSBankArzteUndApothekerBank                PaymentRecordPaymentMethodDetailsEPSBank = "arzte_und_apotheker_bank"
	PaymentRecordPaymentMethodDetailsEPSBankAustrianAnadiBankAg                  PaymentRecordPaymentMethodDetailsEPSBank = "austrian_anadi_bank_ag"
	PaymentRecordPaymentMethodDetailsEPSBankBankAustria                          PaymentRecordPaymentMethodDetailsEPSBank = "bank_austria"
	PaymentRecordPaymentMethodDetailsEPSBankBankhausCarlSpangler                 PaymentRecordPaymentMethodDetailsEPSBank = "bankhaus_carl_spangler"
	PaymentRecordPaymentMethodDetailsEPSBankBankhausSchelhammerUndSchatteraAg    PaymentRecordPaymentMethodDetailsEPSBank = "bankhaus_schelhammer_und_schattera_ag"
	PaymentRecordPaymentMethodDetailsEPSBankBawagPskAg                           PaymentRecordPaymentMethodDetailsEPSBank = "bawag_psk_ag"
	PaymentRecordPaymentMethodDetailsEPSBankBksBankAg                            PaymentRecordPaymentMethodDetailsEPSBank = "bks_bank_ag"
	PaymentRecordPaymentMethodDetailsEPSBankBrullKallmusBankAg                   PaymentRecordPaymentMethodDetailsEPSBank = "brull_kallmus_bank_ag"
	PaymentRecordPaymentMethodDetailsEPSBankBtvVierLanderBank                    PaymentRecordPaymentMethodDetailsEPSBank = "btv_vier_lander_bank"
	PaymentRecordPaymentMethodDetailsEPSBankCapitalBankGraweGruppeAg             PaymentRecordPaymentMethodDetailsEPSBank = "capital_bank_grawe_gruppe_ag"
	PaymentRecordPaymentMethodDetailsEPSBankDeutscheBankAg                       PaymentRecordPaymentMethodDetailsEPSBank = "deutsche_bank_ag"
	PaymentRecordPaymentMethodDetailsEPSBankDolomitenbank                        PaymentRecordPaymentMethodDetailsEPSBank = "dolomitenbank"
	PaymentRecordPaymentMethodDetailsEPSBankEasybankAg                           PaymentRecordPaymentMethodDetailsEPSBank = "easybank_ag"
	PaymentRecordPaymentMethodDetailsEPSBankErsteBankUndSparkassen               PaymentRecordPaymentMethodDetailsEPSBank = "erste_bank_und_sparkassen"
	PaymentRecordPaymentMethodDetailsEPSBankHypoAlpeadriabankInternationalAg     PaymentRecordPaymentMethodDetailsEPSBank = "hypo_alpeadriabank_international_ag"
	PaymentRecordPaymentMethodDetailsEPSBankHypoBankBurgenlandAktiengesellschaft PaymentRecordPaymentMethodDetailsEPSBank = "hypo_bank_burgenland_aktiengesellschaft"
	PaymentRecordPaymentMethodDetailsEPSBankHypoNoeLbFurNiederosterreichUWien    PaymentRecordPaymentMethodDetailsEPSBank = "hypo_noe_lb_fur_niederosterreich_u_wien"
	PaymentRecordPaymentMethodDetailsEPSBankHypoOberosterreichSalzburgSteiermark PaymentRecordPaymentMethodDetailsEPSBank = "hypo_oberosterreich_salzburg_steiermark"
	PaymentRecordPaymentMethodDetailsEPSBankHypoTirolBankAg                      PaymentRecordPaymentMethodDetailsEPSBank = "hypo_tirol_bank_ag"
	PaymentRecordPaymentMethodDetailsEPSBankHypoVorarlbergBankAg                 PaymentRecordPaymentMethodDetailsEPSBank = "hypo_vorarlberg_bank_ag"
	PaymentRecordPaymentMethodDetailsEPSBankMarchfelderBank                      PaymentRecordPaymentMethodDetailsEPSBank = "marchfelder_bank"
	PaymentRecordPaymentMethodDetailsEPSBankOberbankAg                           PaymentRecordPaymentMethodDetailsEPSBank = "oberbank_ag"
	PaymentRecordPaymentMethodDetailsEPSBankRaiffeisenBankengruppeOsterreich     PaymentRecordPaymentMethodDetailsEPSBank = "raiffeisen_bankengruppe_osterreich"
	PaymentRecordPaymentMethodDetailsEPSBankSchoellerbankAg                      PaymentRecordPaymentMethodDetailsEPSBank = "schoellerbank_ag"
	PaymentRecordPaymentMethodDetailsEPSBankSpardaBankWien                       PaymentRecordPaymentMethodDetailsEPSBank = "sparda_bank_wien"
	PaymentRecordPaymentMethodDetailsEPSBankVolksbankGruppe                      PaymentRecordPaymentMethodDetailsEPSBank = "volksbank_gruppe"
	PaymentRecordPaymentMethodDetailsEPSBankVolkskreditbankAg                    PaymentRecordPaymentMethodDetailsEPSBank = "volkskreditbank_ag"
	PaymentRecordPaymentMethodDetailsEPSBankVrBankBraunau                        PaymentRecordPaymentMethodDetailsEPSBank = "vr_bank_braunau"
)

// Account holder type, if provided. Can be one of `individual` or `company`.
type PaymentRecordPaymentMethodDetailsFPXAccountHolderType string

// List of values that PaymentRecordPaymentMethodDetailsFPXAccountHolderType can take
const (
	PaymentRecordPaymentMethodDetailsFPXAccountHolderTypeCompany    PaymentRecordPaymentMethodDetailsFPXAccountHolderType = "company"
	PaymentRecordPaymentMethodDetailsFPXAccountHolderTypeIndividual PaymentRecordPaymentMethodDetailsFPXAccountHolderType = "individual"
)

// The customer's bank. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, `pb_enterprise`, or `bank_of_china`.
type PaymentRecordPaymentMethodDetailsFPXBank string

// List of values that PaymentRecordPaymentMethodDetailsFPXBank can take
const (
	PaymentRecordPaymentMethodDetailsFPXBankAffinBank         PaymentRecordPaymentMethodDetailsFPXBank = "affin_bank"
	PaymentRecordPaymentMethodDetailsFPXBankAgrobank          PaymentRecordPaymentMethodDetailsFPXBank = "agrobank"
	PaymentRecordPaymentMethodDetailsFPXBankAllianceBank      PaymentRecordPaymentMethodDetailsFPXBank = "alliance_bank"
	PaymentRecordPaymentMethodDetailsFPXBankAmbank            PaymentRecordPaymentMethodDetailsFPXBank = "ambank"
	PaymentRecordPaymentMethodDetailsFPXBankBankIslam         PaymentRecordPaymentMethodDetailsFPXBank = "bank_islam"
	PaymentRecordPaymentMethodDetailsFPXBankBankMuamalat      PaymentRecordPaymentMethodDetailsFPXBank = "bank_muamalat"
	PaymentRecordPaymentMethodDetailsFPXBankBankOfChina       PaymentRecordPaymentMethodDetailsFPXBank = "bank_of_china"
	PaymentRecordPaymentMethodDetailsFPXBankBankRakyat        PaymentRecordPaymentMethodDetailsFPXBank = "bank_rakyat"
	PaymentRecordPaymentMethodDetailsFPXBankBsn               PaymentRecordPaymentMethodDetailsFPXBank = "bsn"
	PaymentRecordPaymentMethodDetailsFPXBankCimb              PaymentRecordPaymentMethodDetailsFPXBank = "cimb"
	PaymentRecordPaymentMethodDetailsFPXBankDeutscheBank      PaymentRecordPaymentMethodDetailsFPXBank = "deutsche_bank"
	PaymentRecordPaymentMethodDetailsFPXBankHongLeongBank     PaymentRecordPaymentMethodDetailsFPXBank = "hong_leong_bank"
	PaymentRecordPaymentMethodDetailsFPXBankHsbc              PaymentRecordPaymentMethodDetailsFPXBank = "hsbc"
	PaymentRecordPaymentMethodDetailsFPXBankKfh               PaymentRecordPaymentMethodDetailsFPXBank = "kfh"
	PaymentRecordPaymentMethodDetailsFPXBankMaybank2e         PaymentRecordPaymentMethodDetailsFPXBank = "maybank2e"
	PaymentRecordPaymentMethodDetailsFPXBankMaybank2u         PaymentRecordPaymentMethodDetailsFPXBank = "maybank2u"
	PaymentRecordPaymentMethodDetailsFPXBankOcbc              PaymentRecordPaymentMethodDetailsFPXBank = "ocbc"
	PaymentRecordPaymentMethodDetailsFPXBankPbEnterprise      PaymentRecordPaymentMethodDetailsFPXBank = "pb_enterprise"
	PaymentRecordPaymentMethodDetailsFPXBankPublicBank        PaymentRecordPaymentMethodDetailsFPXBank = "public_bank"
	PaymentRecordPaymentMethodDetailsFPXBankRhb               PaymentRecordPaymentMethodDetailsFPXBank = "rhb"
	PaymentRecordPaymentMethodDetailsFPXBankStandardChartered PaymentRecordPaymentMethodDetailsFPXBank = "standard_chartered"
	PaymentRecordPaymentMethodDetailsFPXBankUob               PaymentRecordPaymentMethodDetailsFPXBank = "uob"
)

// The customer's bank. Can be one of `abn_amro`, `adyen`, `asn_bank`, `bunq`, `buut`, `finom`, `handelsbanken`, `ing`, `knab`, `mollie`, `moneyou`, `n26`, `nn`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, `van_lanschot`, or `yoursafe`.
type PaymentRecordPaymentMethodDetailsIDEALBank string

// List of values that PaymentRecordPaymentMethodDetailsIDEALBank can take
const (
	PaymentRecordPaymentMethodDetailsIDEALBankAbnAmro       PaymentRecordPaymentMethodDetailsIDEALBank = "abn_amro"
	PaymentRecordPaymentMethodDetailsIDEALBankAdyen         PaymentRecordPaymentMethodDetailsIDEALBank = "adyen"
	PaymentRecordPaymentMethodDetailsIDEALBankAsnBank       PaymentRecordPaymentMethodDetailsIDEALBank = "asn_bank"
	PaymentRecordPaymentMethodDetailsIDEALBankBunq          PaymentRecordPaymentMethodDetailsIDEALBank = "bunq"
	PaymentRecordPaymentMethodDetailsIDEALBankBuut          PaymentRecordPaymentMethodDetailsIDEALBank = "buut"
	PaymentRecordPaymentMethodDetailsIDEALBankFinom         PaymentRecordPaymentMethodDetailsIDEALBank = "finom"
	PaymentRecordPaymentMethodDetailsIDEALBankHandelsbanken PaymentRecordPaymentMethodDetailsIDEALBank = "handelsbanken"
	PaymentRecordPaymentMethodDetailsIDEALBankIng           PaymentRecordPaymentMethodDetailsIDEALBank = "ing"
	PaymentRecordPaymentMethodDetailsIDEALBankKnab          PaymentRecordPaymentMethodDetailsIDEALBank = "knab"
	PaymentRecordPaymentMethodDetailsIDEALBankMollie        PaymentRecordPaymentMethodDetailsIDEALBank = "mollie"
	PaymentRecordPaymentMethodDetailsIDEALBankMoneyou       PaymentRecordPaymentMethodDetailsIDEALBank = "moneyou"
	PaymentRecordPaymentMethodDetailsIDEALBankN26           PaymentRecordPaymentMethodDetailsIDEALBank = "n26"
	PaymentRecordPaymentMethodDetailsIDEALBankNn            PaymentRecordPaymentMethodDetailsIDEALBank = "nn"
	PaymentRecordPaymentMethodDetailsIDEALBankRabobank      PaymentRecordPaymentMethodDetailsIDEALBank = "rabobank"
	PaymentRecordPaymentMethodDetailsIDEALBankRegiobank     PaymentRecordPaymentMethodDetailsIDEALBank = "regiobank"
	PaymentRecordPaymentMethodDetailsIDEALBankRevolut       PaymentRecordPaymentMethodDetailsIDEALBank = "revolut"
	PaymentRecordPaymentMethodDetailsIDEALBankSnsBank       PaymentRecordPaymentMethodDetailsIDEALBank = "sns_bank"
	PaymentRecordPaymentMethodDetailsIDEALBankTriodosBank   PaymentRecordPaymentMethodDetailsIDEALBank = "triodos_bank"
	PaymentRecordPaymentMethodDetailsIDEALBankVanLanschot   PaymentRecordPaymentMethodDetailsIDEALBank = "van_lanschot"
	PaymentRecordPaymentMethodDetailsIDEALBankYoursafe      PaymentRecordPaymentMethodDetailsIDEALBank = "yoursafe"
)

// The Bank Identifier Code of the customer's bank.
type PaymentRecordPaymentMethodDetailsIDEALBIC string

// List of values that PaymentRecordPaymentMethodDetailsIDEALBIC can take
const (
	PaymentRecordPaymentMethodDetailsIDEALBICABNANL2A PaymentRecordPaymentMethodDetailsIDEALBIC = "ABNANL2A"
	PaymentRecordPaymentMethodDetailsIDEALBICADYBNL2A PaymentRecordPaymentMethodDetailsIDEALBIC = "ADYBNL2A"
	PaymentRecordPaymentMethodDetailsIDEALBICASNBNL21 PaymentRecordPaymentMethodDetailsIDEALBIC = "ASNBNL21"
	PaymentRecordPaymentMethodDetailsIDEALBICBITSNL2A PaymentRecordPaymentMethodDetailsIDEALBIC = "BITSNL2A"
	PaymentRecordPaymentMethodDetailsIDEALBICBUNQNL2A PaymentRecordPaymentMethodDetailsIDEALBIC = "BUNQNL2A"
	PaymentRecordPaymentMethodDetailsIDEALBICBUUTNL2A PaymentRecordPaymentMethodDetailsIDEALBIC = "BUUTNL2A"
	PaymentRecordPaymentMethodDetailsIDEALBICFNOMNL22 PaymentRecordPaymentMethodDetailsIDEALBIC = "FNOMNL22"
	PaymentRecordPaymentMethodDetailsIDEALBICFVLBNL22 PaymentRecordPaymentMethodDetailsIDEALBIC = "FVLBNL22"
	PaymentRecordPaymentMethodDetailsIDEALBICHANDNL2A PaymentRecordPaymentMethodDetailsIDEALBIC = "HANDNL2A"
	PaymentRecordPaymentMethodDetailsIDEALBICINGBNL2A PaymentRecordPaymentMethodDetailsIDEALBIC = "INGBNL2A"
	PaymentRecordPaymentMethodDetailsIDEALBICKNABNL2H PaymentRecordPaymentMethodDetailsIDEALBIC = "KNABNL2H"
	PaymentRecordPaymentMethodDetailsIDEALBICMLLENL2A PaymentRecordPaymentMethodDetailsIDEALBIC = "MLLENL2A"
	PaymentRecordPaymentMethodDetailsIDEALBICMOYONL21 PaymentRecordPaymentMethodDetailsIDEALBIC = "MOYONL21"
	PaymentRecordPaymentMethodDetailsIDEALBICNNBANL2G PaymentRecordPaymentMethodDetailsIDEALBIC = "NNBANL2G"
	PaymentRecordPaymentMethodDetailsIDEALBICNTSBDEB1 PaymentRecordPaymentMethodDetailsIDEALBIC = "NTSBDEB1"
	PaymentRecordPaymentMethodDetailsIDEALBICRABONL2U PaymentRecordPaymentMethodDetailsIDEALBIC = "RABONL2U"
	PaymentRecordPaymentMethodDetailsIDEALBICRBRBNL21 PaymentRecordPaymentMethodDetailsIDEALBIC = "RBRBNL21"
	PaymentRecordPaymentMethodDetailsIDEALBICREVOIE23 PaymentRecordPaymentMethodDetailsIDEALBIC = "REVOIE23"
	PaymentRecordPaymentMethodDetailsIDEALBICREVOLT21 PaymentRecordPaymentMethodDetailsIDEALBIC = "REVOLT21"
	PaymentRecordPaymentMethodDetailsIDEALBICSNSBNL2A PaymentRecordPaymentMethodDetailsIDEALBIC = "SNSBNL2A"
	PaymentRecordPaymentMethodDetailsIDEALBICTRIONL2U PaymentRecordPaymentMethodDetailsIDEALBIC = "TRIONL2U"
)

// How card details were read in this transaction.
type PaymentRecordPaymentMethodDetailsInteracPresentReadMethod string

// List of values that PaymentRecordPaymentMethodDetailsInteracPresentReadMethod can take
const (
	PaymentRecordPaymentMethodDetailsInteracPresentReadMethodContactEmv               PaymentRecordPaymentMethodDetailsInteracPresentReadMethod = "contact_emv"
	PaymentRecordPaymentMethodDetailsInteracPresentReadMethodContactlessEmv           PaymentRecordPaymentMethodDetailsInteracPresentReadMethod = "contactless_emv"
	PaymentRecordPaymentMethodDetailsInteracPresentReadMethodContactlessMagstripeMode PaymentRecordPaymentMethodDetailsInteracPresentReadMethod = "contactless_magstripe_mode"
	PaymentRecordPaymentMethodDetailsInteracPresentReadMethodMagneticStripeFallback   PaymentRecordPaymentMethodDetailsInteracPresentReadMethod = "magnetic_stripe_fallback"
	PaymentRecordPaymentMethodDetailsInteracPresentReadMethodMagneticStripeTrack2     PaymentRecordPaymentMethodDetailsInteracPresentReadMethod = "magnetic_stripe_track2"
)

// The type of account being debited or credited
type PaymentRecordPaymentMethodDetailsInteracPresentReceiptAccountType string

// List of values that PaymentRecordPaymentMethodDetailsInteracPresentReceiptAccountType can take
const (
	PaymentRecordPaymentMethodDetailsInteracPresentReceiptAccountTypeChecking PaymentRecordPaymentMethodDetailsInteracPresentReceiptAccountType = "checking"
	PaymentRecordPaymentMethodDetailsInteracPresentReceiptAccountTypeSavings  PaymentRecordPaymentMethodDetailsInteracPresentReceiptAccountType = "savings"
	PaymentRecordPaymentMethodDetailsInteracPresentReceiptAccountTypeUnknown  PaymentRecordPaymentMethodDetailsInteracPresentReceiptAccountType = "unknown"
)

// The name of the convenience store chain where the payment was completed.
type PaymentRecordPaymentMethodDetailsKonbiniStoreChain string

// List of values that PaymentRecordPaymentMethodDetailsKonbiniStoreChain can take
const (
	PaymentRecordPaymentMethodDetailsKonbiniStoreChainFamilyMart PaymentRecordPaymentMethodDetailsKonbiniStoreChain = "familymart"
	PaymentRecordPaymentMethodDetailsKonbiniStoreChainLawson     PaymentRecordPaymentMethodDetailsKonbiniStoreChain = "lawson"
	PaymentRecordPaymentMethodDetailsKonbiniStoreChainMinistop   PaymentRecordPaymentMethodDetailsKonbiniStoreChain = "ministop"
	PaymentRecordPaymentMethodDetailsKonbiniStoreChainSeicomart  PaymentRecordPaymentMethodDetailsKonbiniStoreChain = "seicomart"
)

// The local credit or debit card brand.
type PaymentRecordPaymentMethodDetailsKrCardBrand string

// List of values that PaymentRecordPaymentMethodDetailsKrCardBrand can take
const (
	PaymentRecordPaymentMethodDetailsKrCardBrandBc          PaymentRecordPaymentMethodDetailsKrCardBrand = "bc"
	PaymentRecordPaymentMethodDetailsKrCardBrandCiti        PaymentRecordPaymentMethodDetailsKrCardBrand = "citi"
	PaymentRecordPaymentMethodDetailsKrCardBrandHana        PaymentRecordPaymentMethodDetailsKrCardBrand = "hana"
	PaymentRecordPaymentMethodDetailsKrCardBrandHyundai     PaymentRecordPaymentMethodDetailsKrCardBrand = "hyundai"
	PaymentRecordPaymentMethodDetailsKrCardBrandJeju        PaymentRecordPaymentMethodDetailsKrCardBrand = "jeju"
	PaymentRecordPaymentMethodDetailsKrCardBrandJeonbuk     PaymentRecordPaymentMethodDetailsKrCardBrand = "jeonbuk"
	PaymentRecordPaymentMethodDetailsKrCardBrandKakaobank   PaymentRecordPaymentMethodDetailsKrCardBrand = "kakaobank"
	PaymentRecordPaymentMethodDetailsKrCardBrandKbank       PaymentRecordPaymentMethodDetailsKrCardBrand = "kbank"
	PaymentRecordPaymentMethodDetailsKrCardBrandKdbbank     PaymentRecordPaymentMethodDetailsKrCardBrand = "kdbbank"
	PaymentRecordPaymentMethodDetailsKrCardBrandKookmin     PaymentRecordPaymentMethodDetailsKrCardBrand = "kookmin"
	PaymentRecordPaymentMethodDetailsKrCardBrandKwangju     PaymentRecordPaymentMethodDetailsKrCardBrand = "kwangju"
	PaymentRecordPaymentMethodDetailsKrCardBrandLotte       PaymentRecordPaymentMethodDetailsKrCardBrand = "lotte"
	PaymentRecordPaymentMethodDetailsKrCardBrandMg          PaymentRecordPaymentMethodDetailsKrCardBrand = "mg"
	PaymentRecordPaymentMethodDetailsKrCardBrandNh          PaymentRecordPaymentMethodDetailsKrCardBrand = "nh"
	PaymentRecordPaymentMethodDetailsKrCardBrandPost        PaymentRecordPaymentMethodDetailsKrCardBrand = "post"
	PaymentRecordPaymentMethodDetailsKrCardBrandSamsung     PaymentRecordPaymentMethodDetailsKrCardBrand = "samsung"
	PaymentRecordPaymentMethodDetailsKrCardBrandSavingsbank PaymentRecordPaymentMethodDetailsKrCardBrand = "savingsbank"
	PaymentRecordPaymentMethodDetailsKrCardBrandShinhan     PaymentRecordPaymentMethodDetailsKrCardBrand = "shinhan"
	PaymentRecordPaymentMethodDetailsKrCardBrandShinhyup    PaymentRecordPaymentMethodDetailsKrCardBrand = "shinhyup"
	PaymentRecordPaymentMethodDetailsKrCardBrandSuhyup      PaymentRecordPaymentMethodDetailsKrCardBrand = "suhyup"
	PaymentRecordPaymentMethodDetailsKrCardBrandTossbank    PaymentRecordPaymentMethodDetailsKrCardBrand = "tossbank"
	PaymentRecordPaymentMethodDetailsKrCardBrandWoori       PaymentRecordPaymentMethodDetailsKrCardBrand = "woori"
)

// The customer's bank. Can be one of `ing`, `citi_handlowy`, `tmobile_usbugi_bankowe`, `plus_bank`, `etransfer_pocztowy24`, `banki_spbdzielcze`, `bank_nowy_bfg_sa`, `getin_bank`, `velobank`, `blik`, `noble_pay`, `ideabank`, `envelobank`, `santander_przelew24`, `nest_przelew`, `mbank_mtransfer`, `inteligo`, `pbac_z_ipko`, `bnp_paribas`, `credit_agricole`, `toyota_bank`, `bank_pekao_sa`, `volkswagen_bank`, `bank_millennium`, `alior_bank`, or `boz`.
type PaymentRecordPaymentMethodDetailsP24Bank string

// List of values that PaymentRecordPaymentMethodDetailsP24Bank can take
const (
	PaymentRecordPaymentMethodDetailsP24BankAliorBank            PaymentRecordPaymentMethodDetailsP24Bank = "alior_bank"
	PaymentRecordPaymentMethodDetailsP24BankBankMillennium       PaymentRecordPaymentMethodDetailsP24Bank = "bank_millennium"
	PaymentRecordPaymentMethodDetailsP24BankBankNowyBfgSa        PaymentRecordPaymentMethodDetailsP24Bank = "bank_nowy_bfg_sa"
	PaymentRecordPaymentMethodDetailsP24BankBankPekaoSa          PaymentRecordPaymentMethodDetailsP24Bank = "bank_pekao_sa"
	PaymentRecordPaymentMethodDetailsP24BankBankiSpbdzielcze     PaymentRecordPaymentMethodDetailsP24Bank = "banki_spbdzielcze"
	PaymentRecordPaymentMethodDetailsP24BankBLIK                 PaymentRecordPaymentMethodDetailsP24Bank = "blik"
	PaymentRecordPaymentMethodDetailsP24BankBnpParibas           PaymentRecordPaymentMethodDetailsP24Bank = "bnp_paribas"
	PaymentRecordPaymentMethodDetailsP24BankBoz                  PaymentRecordPaymentMethodDetailsP24Bank = "boz"
	PaymentRecordPaymentMethodDetailsP24BankCitiHandlowy         PaymentRecordPaymentMethodDetailsP24Bank = "citi_handlowy"
	PaymentRecordPaymentMethodDetailsP24BankCreditAgricole       PaymentRecordPaymentMethodDetailsP24Bank = "credit_agricole"
	PaymentRecordPaymentMethodDetailsP24BankEnvelobank           PaymentRecordPaymentMethodDetailsP24Bank = "envelobank"
	PaymentRecordPaymentMethodDetailsP24BankEtransferPocztowy24  PaymentRecordPaymentMethodDetailsP24Bank = "etransfer_pocztowy24"
	PaymentRecordPaymentMethodDetailsP24BankGetinBank            PaymentRecordPaymentMethodDetailsP24Bank = "getin_bank"
	PaymentRecordPaymentMethodDetailsP24BankIdeabank             PaymentRecordPaymentMethodDetailsP24Bank = "ideabank"
	PaymentRecordPaymentMethodDetailsP24BankIng                  PaymentRecordPaymentMethodDetailsP24Bank = "ing"
	PaymentRecordPaymentMethodDetailsP24BankInteligo             PaymentRecordPaymentMethodDetailsP24Bank = "inteligo"
	PaymentRecordPaymentMethodDetailsP24BankMbankMtransfer       PaymentRecordPaymentMethodDetailsP24Bank = "mbank_mtransfer"
	PaymentRecordPaymentMethodDetailsP24BankNestPrzelew          PaymentRecordPaymentMethodDetailsP24Bank = "nest_przelew"
	PaymentRecordPaymentMethodDetailsP24BankNoblePay             PaymentRecordPaymentMethodDetailsP24Bank = "noble_pay"
	PaymentRecordPaymentMethodDetailsP24BankPbacZIpko            PaymentRecordPaymentMethodDetailsP24Bank = "pbac_z_ipko"
	PaymentRecordPaymentMethodDetailsP24BankPlusBank             PaymentRecordPaymentMethodDetailsP24Bank = "plus_bank"
	PaymentRecordPaymentMethodDetailsP24BankSantanderPrzelew24   PaymentRecordPaymentMethodDetailsP24Bank = "santander_przelew24"
	PaymentRecordPaymentMethodDetailsP24BankTmobileUsbugiBankowe PaymentRecordPaymentMethodDetailsP24Bank = "tmobile_usbugi_bankowe"
	PaymentRecordPaymentMethodDetailsP24BankToyotaBank           PaymentRecordPaymentMethodDetailsP24Bank = "toyota_bank"
	PaymentRecordPaymentMethodDetailsP24BankVelobank             PaymentRecordPaymentMethodDetailsP24Bank = "velobank"
	PaymentRecordPaymentMethodDetailsP24BankVolkswagenBank       PaymentRecordPaymentMethodDetailsP24Bank = "volkswagen_bank"
)

// An array of conditions that are covered for the transaction, if applicable.
type PaymentRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategory string

// List of values that PaymentRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategory can take
const (
	PaymentRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategoryFraudulent         PaymentRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategory = "fraudulent"
	PaymentRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategoryProductNotReceived PaymentRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategory = "product_not_received"
)

// Indicates whether the transaction is eligible for PayPal's seller protection.
type PaymentRecordPaymentMethodDetailsPaypalSellerProtectionStatus string

// List of values that PaymentRecordPaymentMethodDetailsPaypalSellerProtectionStatus can take
const (
	PaymentRecordPaymentMethodDetailsPaypalSellerProtectionStatusEligible          PaymentRecordPaymentMethodDetailsPaypalSellerProtectionStatus = "eligible"
	PaymentRecordPaymentMethodDetailsPaypalSellerProtectionStatusNotEligible       PaymentRecordPaymentMethodDetailsPaypalSellerProtectionStatus = "not_eligible"
	PaymentRecordPaymentMethodDetailsPaypalSellerProtectionStatusPartiallyEligible PaymentRecordPaymentMethodDetailsPaypalSellerProtectionStatus = "partially_eligible"
)

// funding type of the underlying payment method.
type PaymentRecordPaymentMethodDetailsRevolutPayFundingType string

// List of values that PaymentRecordPaymentMethodDetailsRevolutPayFundingType can take
const (
	PaymentRecordPaymentMethodDetailsRevolutPayFundingTypeCard PaymentRecordPaymentMethodDetailsRevolutPayFundingType = "card"
)

// Preferred language of the SOFORT authorization page that the customer is redirected to.
// Can be one of `de`, `en`, `es`, `fr`, `it`, `nl`, or `pl`
type PaymentRecordPaymentMethodDetailsSofortPreferredLanguage string

// List of values that PaymentRecordPaymentMethodDetailsSofortPreferredLanguage can take
const (
	PaymentRecordPaymentMethodDetailsSofortPreferredLanguageDE PaymentRecordPaymentMethodDetailsSofortPreferredLanguage = "de"
	PaymentRecordPaymentMethodDetailsSofortPreferredLanguageEN PaymentRecordPaymentMethodDetailsSofortPreferredLanguage = "en"
	PaymentRecordPaymentMethodDetailsSofortPreferredLanguageES PaymentRecordPaymentMethodDetailsSofortPreferredLanguage = "es"
	PaymentRecordPaymentMethodDetailsSofortPreferredLanguageFR PaymentRecordPaymentMethodDetailsSofortPreferredLanguage = "fr"
	PaymentRecordPaymentMethodDetailsSofortPreferredLanguageIT PaymentRecordPaymentMethodDetailsSofortPreferredLanguage = "it"
	PaymentRecordPaymentMethodDetailsSofortPreferredLanguageNL PaymentRecordPaymentMethodDetailsSofortPreferredLanguage = "nl"
	PaymentRecordPaymentMethodDetailsSofortPreferredLanguagePL PaymentRecordPaymentMethodDetailsSofortPreferredLanguage = "pl"
)

// The type of entity that holds the account. This can be either 'individual' or 'company'.
type PaymentRecordPaymentMethodDetailsUSBankAccountAccountHolderType string

// List of values that PaymentRecordPaymentMethodDetailsUSBankAccountAccountHolderType can take
const (
	PaymentRecordPaymentMethodDetailsUSBankAccountAccountHolderTypeCompany    PaymentRecordPaymentMethodDetailsUSBankAccountAccountHolderType = "company"
	PaymentRecordPaymentMethodDetailsUSBankAccountAccountHolderTypeIndividual PaymentRecordPaymentMethodDetailsUSBankAccountAccountHolderType = "individual"
)

// The type of the bank account. This can be either 'checking' or 'savings'.
type PaymentRecordPaymentMethodDetailsUSBankAccountAccountType string

// List of values that PaymentRecordPaymentMethodDetailsUSBankAccountAccountType can take
const (
	PaymentRecordPaymentMethodDetailsUSBankAccountAccountTypeChecking PaymentRecordPaymentMethodDetailsUSBankAccountAccountType = "checking"
	PaymentRecordPaymentMethodDetailsUSBankAccountAccountTypeSavings  PaymentRecordPaymentMethodDetailsUSBankAccountAccountType = "savings"
)

// The processor used for this payment attempt.
type PaymentRecordProcessorDetailsType string

// List of values that PaymentRecordProcessorDetailsType can take
const (
	PaymentRecordProcessorDetailsTypeCustom PaymentRecordProcessorDetailsType = "custom"
)

// Indicates who reported the payment.
type PaymentRecordReportedBy string

// List of values that PaymentRecordReportedBy can take
const (
	PaymentRecordReportedBySelf   PaymentRecordReportedBy = "self"
	PaymentRecordReportedByStripe PaymentRecordReportedBy = "stripe"
)

// Retrieves a Payment Record with the given ID
type PaymentRecordParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Information about the payment attempt failure.
type PaymentRecordReportPaymentAttemptFailedParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// When the reported payment failed. Measured in seconds since the Unix epoch.
	FailedAt *int64 `form:"failed_at"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentAttemptFailedParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentAttemptFailedParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Information about the payment attempt guarantee.
type PaymentRecordReportPaymentAttemptGuaranteedParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// When the reported payment was guaranteed. Measured in seconds since the Unix epoch.
	GuaranteedAt *int64 `form:"guaranteed_at"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentAttemptGuaranteedParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentAttemptGuaranteedParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The billing details associated with the method of payment.
type PaymentRecordReportPaymentAttemptPaymentMethodDetailsBillingDetailsParams struct {
	// The billing address associated with the method of payment.
	Address *AddressParams `form:"address"`
	// The billing email associated with the method of payment.
	Email *string `form:"email"`
	// The billing name associated with the method of payment.
	Name *string `form:"name"`
	// The billing phone number associated with the method of payment.
	Phone *string `form:"phone"`
}

// Information about the custom (user-defined) payment method used to make this payment.
type PaymentRecordReportPaymentAttemptPaymentMethodDetailsCustomParams struct {
	// Display name for the custom (user-defined) payment method type used to make this payment.
	DisplayName *string `form:"display_name"`
	// The custom payment method type associated with this payment.
	Type *string `form:"type"`
}

// Information about the Payment Method debited for this payment.
type PaymentRecordReportPaymentAttemptPaymentMethodDetailsParams struct {
	// The billing details associated with the method of payment.
	BillingDetails *PaymentRecordReportPaymentAttemptPaymentMethodDetailsBillingDetailsParams `form:"billing_details"`
	// Information about the custom (user-defined) payment method used to make this payment.
	Custom *PaymentRecordReportPaymentAttemptPaymentMethodDetailsCustomParams `form:"custom"`
	// ID of the Stripe Payment Method used to make this payment.
	PaymentMethod *string `form:"payment_method"`
	// The type of the payment method details. An additional hash is included on the payment_method_details with a name matching this value. It contains additional information specific to the type.
	Type *string `form:"type"`
}

// Shipping information for this payment.
type PaymentRecordReportPaymentAttemptShippingDetailsParams struct {
	// The physical shipping address.
	Address *AddressParams `form:"address"`
	// The shipping recipient's name.
	Name *string `form:"name"`
	// The shipping recipient's phone number.
	Phone *string `form:"phone"`
}

// Report a new payment attempt on the specified Payment Record. A new payment
//
//	attempt can only be specified if all other payment attempts are canceled or failed.
type PaymentRecordReportPaymentAttemptParams struct {
	Params `form:"*"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Information about the payment attempt failure.
	Failed *PaymentRecordReportPaymentAttemptFailedParams `form:"failed"`
	// Information about the payment attempt guarantee.
	Guaranteed *PaymentRecordReportPaymentAttemptGuaranteedParams `form:"guaranteed"`
	// When the reported payment was initiated. Measured in seconds since the Unix epoch.
	InitiatedAt *int64 `form:"initiated_at"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The outcome of the reported payment.
	Outcome *string `form:"outcome"`
	// Information about the Payment Method debited for this payment.
	PaymentMethodDetails *PaymentRecordReportPaymentAttemptPaymentMethodDetailsParams `form:"payment_method_details"`
	// Shipping information for this payment.
	ShippingDetails *PaymentRecordReportPaymentAttemptShippingDetailsParams `form:"shipping_details"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentAttemptParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentAttemptParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was canceled.
type PaymentRecordReportPaymentAttemptCanceledParams struct {
	Params `form:"*"`
	// When the reported payment was canceled. Measured in seconds since the Unix epoch.
	CanceledAt *int64 `form:"canceled_at"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentAttemptCanceledParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentAttemptCanceledParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Customer information for this payment.
type PaymentRecordReportPaymentAttemptInformationalCustomerDetailsParams struct {
	// The customer who made the payment.
	Customer *string `form:"customer"`
	// The customer's phone number.
	Email *string `form:"email"`
	// The customer's name.
	Name *string `form:"name"`
	// The customer's phone number.
	Phone *string `form:"phone"`
}

// Shipping information for this payment.
type PaymentRecordReportPaymentAttemptInformationalShippingDetailsParams struct {
	// The physical shipping address.
	Address *AddressParams `form:"address"`
	// The shipping recipient's name.
	Name *string `form:"name"`
	// The shipping recipient's phone number.
	Phone *string `form:"phone"`
}

// Report informational updates on the specified Payment Record.
type PaymentRecordReportPaymentAttemptInformationalParams struct {
	Params `form:"*"`
	// Customer information for this payment.
	CustomerDetails *PaymentRecordReportPaymentAttemptInformationalCustomerDetailsParams `form:"customer_details"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Shipping information for this payment.
	ShippingDetails *PaymentRecordReportPaymentAttemptInformationalShippingDetailsParams `form:"shipping_details"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentAttemptInformationalParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentAttemptInformationalParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) representing how much of this payment to refund. Can refund only up to the remaining, unrefunded amount of the payment.
type PaymentRecordReportRefundAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value *int64 `form:"value"`
}

// Information about the custom processor used to make this refund.
type PaymentRecordReportRefundProcessorDetailsCustomParams struct {
	// A reference to the external refund. This field must be unique across all refunds.
	RefundReference *string `form:"refund_reference"`
}

// Processor information for this refund.
type PaymentRecordReportRefundProcessorDetailsParams struct {
	// Information about the custom processor used to make this refund.
	Custom *PaymentRecordReportRefundProcessorDetailsCustomParams `form:"custom"`
	// The type of the processor details. An additional hash is included on processor_details with a name matching this value. It contains additional information specific to the processor.
	Type *string `form:"type"`
}

// Information about the payment attempt refund.
type PaymentRecordReportRefundRefundedParams struct {
	// When the reported refund completed. Measured in seconds since the Unix epoch.
	RefundedAt *int64 `form:"refunded_at"`
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was refunded.
type PaymentRecordReportRefundParams struct {
	Params `form:"*"`
	// A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) representing how much of this payment to refund. Can refund only up to the remaining, unrefunded amount of the payment.
	Amount *PaymentRecordReportRefundAmountParams `form:"amount"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// When the reported refund was initiated. Measured in seconds since the Unix epoch.
	InitiatedAt *int64 `form:"initiated_at"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The outcome of the reported refund.
	Outcome *string `form:"outcome"`
	// Processor information for this refund.
	ProcessorDetails *PaymentRecordReportRefundProcessorDetailsParams `form:"processor_details"`
	// Information about the payment attempt refund.
	Refunded *PaymentRecordReportRefundRefundedParams `form:"refunded"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportRefundParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportRefundParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The amount you initially requested for this payment.
type PaymentRecordReportPaymentAmountRequestedParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value *int64 `form:"value"`
}

// Customer information for this payment.
type PaymentRecordReportPaymentCustomerDetailsParams struct {
	// The customer who made the payment.
	Customer *string `form:"customer"`
	// The customer's phone number.
	Email *string `form:"email"`
	// The customer's name.
	Name *string `form:"name"`
	// The customer's phone number.
	Phone *string `form:"phone"`
}

// Information about the payment attempt failure.
type PaymentRecordReportPaymentFailedParams struct {
	// When the reported payment failed. Measured in seconds since the Unix epoch.
	FailedAt *int64 `form:"failed_at"`
}

// Information about the payment attempt guarantee.
type PaymentRecordReportPaymentGuaranteedParams struct {
	// When the reported payment was guaranteed. Measured in seconds since the Unix epoch.
	GuaranteedAt *int64 `form:"guaranteed_at"`
}

// The billing details associated with the method of payment.
type PaymentRecordReportPaymentPaymentMethodDetailsBillingDetailsParams struct {
	// The billing address associated with the method of payment.
	Address *AddressParams `form:"address"`
	// The billing email associated with the method of payment.
	Email *string `form:"email"`
	// The billing name associated with the method of payment.
	Name *string `form:"name"`
	// The billing phone number associated with the method of payment.
	Phone *string `form:"phone"`
}

// Information about the custom (user-defined) payment method used to make this payment.
type PaymentRecordReportPaymentPaymentMethodDetailsCustomParams struct {
	// Display name for the custom (user-defined) payment method type used to make this payment.
	DisplayName *string `form:"display_name"`
	// The custom payment method type associated with this payment.
	Type *string `form:"type"`
}

// Information about the Payment Method debited for this payment.
type PaymentRecordReportPaymentPaymentMethodDetailsParams struct {
	// The billing details associated with the method of payment.
	BillingDetails *PaymentRecordReportPaymentPaymentMethodDetailsBillingDetailsParams `form:"billing_details"`
	// Information about the custom (user-defined) payment method used to make this payment.
	Custom *PaymentRecordReportPaymentPaymentMethodDetailsCustomParams `form:"custom"`
	// ID of the Stripe Payment Method used to make this payment.
	PaymentMethod *string `form:"payment_method"`
	// The type of the payment method details. An additional hash is included on the payment_method_details with a name matching this value. It contains additional information specific to the type.
	Type *string `form:"type"`
}

// Information about the custom processor used to make this payment.
type PaymentRecordReportPaymentProcessorDetailsCustomParams struct {
	// An opaque string for manual reconciliation of this payment, for example a check number or a payment processor ID.
	PaymentReference *string `form:"payment_reference"`
}

// Processor information for this payment.
type PaymentRecordReportPaymentProcessorDetailsParams struct {
	// Information about the custom processor used to make this payment.
	Custom *PaymentRecordReportPaymentProcessorDetailsCustomParams `form:"custom"`
	// The type of the processor details. An additional hash is included on processor_details with a name matching this value. It contains additional information specific to the processor.
	Type *string `form:"type"`
}

// Shipping information for this payment.
type PaymentRecordReportPaymentShippingDetailsParams struct {
	// The physical shipping address.
	Address *AddressParams `form:"address"`
	// The shipping recipient's name.
	Name *string `form:"name"`
	// The shipping recipient's phone number.
	Phone *string `form:"phone"`
}

// Report a new Payment Record. You may report a Payment Record as it is
//
//	initialized and later report updates through the other report_* methods, or report Payment
//	Records in a terminal state directly, through this method.
type PaymentRecordReportPaymentParams struct {
	Params `form:"*"`
	// The amount you initially requested for this payment.
	AmountRequested *PaymentRecordReportPaymentAmountRequestedParams `form:"amount_requested"`
	// Customer information for this payment.
	CustomerDetails *PaymentRecordReportPaymentCustomerDetailsParams `form:"customer_details"`
	// Indicates whether the customer was present in your checkout flow during this payment.
	CustomerPresence *string `form:"customer_presence"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Information about the payment attempt failure.
	Failed *PaymentRecordReportPaymentFailedParams `form:"failed"`
	// Information about the payment attempt guarantee.
	Guaranteed *PaymentRecordReportPaymentGuaranteedParams `form:"guaranteed"`
	// When the reported payment was initiated. Measured in seconds since the Unix epoch.
	InitiatedAt *int64 `form:"initiated_at"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The outcome of the reported payment.
	Outcome *string `form:"outcome"`
	// Information about the Payment Method debited for this payment.
	PaymentMethodDetails *PaymentRecordReportPaymentPaymentMethodDetailsParams `form:"payment_method_details"`
	// Processor information for this payment.
	ProcessorDetails *PaymentRecordReportPaymentProcessorDetailsParams `form:"processor_details"`
	// Shipping information for this payment.
	ShippingDetails *PaymentRecordReportPaymentShippingDetailsParams `form:"shipping_details"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves a Payment Record with the given ID
type PaymentRecordRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmount struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmountAuthorized struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmountCanceled struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmountFailed struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmountGuaranteed struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmountRefunded struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmountRequested struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// Customer information for this payment.
type PaymentRecordCustomerDetails struct {
	// ID of the Stripe Customer associated with this payment.
	Customer string `json:"customer"`
	// The customer's email address.
	Email string `json:"email"`
	// The customer's name.
	Name string `json:"name"`
	// The customer's phone number.
	Phone string `json:"phone"`
}
type PaymentRecordPaymentMethodDetailsACHCreditTransfer struct {
	// Account number to transfer funds to.
	AccountNumber string `json:"account_number"`
	// Name of the bank associated with the routing number.
	BankName string `json:"bank_name"`
	// Routing transit number for the bank account to transfer funds to.
	RoutingNumber string `json:"routing_number"`
	// SWIFT code of the bank associated with the routing number.
	SwiftCode string `json:"swift_code"`
}
type PaymentRecordPaymentMethodDetailsACHDebit struct {
	// Type of entity that holds the account. This can be either `individual` or `company`.
	AccountHolderType PaymentRecordPaymentMethodDetailsACHDebitAccountHolderType `json:"account_holder_type"`
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
type PaymentRecordPaymentMethodDetailsACSSDebit struct {
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
type PaymentRecordPaymentMethodDetailsAffirm struct {
	// ID of the location that this reader is assigned to.
	Location string `json:"location"`
	// ID of the reader this transaction was made on.
	Reader string `json:"reader"`
	// The Affirm transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsAfterpayClearpay struct {
	// The Afterpay order ID associated with this payment intent.
	OrderID string `json:"order_id"`
	// Order identifier shown to the merchant in Afterpay's online portal.
	Reference string `json:"reference"`
}
type PaymentRecordPaymentMethodDetailsAlipay struct {
	// Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.
	BuyerID string `json:"buyer_id"`
	// Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Transaction ID of this particular Alipay transaction.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsAlmaInstallments struct {
	// The number of installments.
	Count int64 `json:"count"`
}
type PaymentRecordPaymentMethodDetailsAlma struct {
	Installments *PaymentRecordPaymentMethodDetailsAlmaInstallments `json:"installments"`
	// The Alma transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsAmazonPayFundingCard struct {
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand string `json:"brand"`
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
type PaymentRecordPaymentMethodDetailsAmazonPayFunding struct {
	Card *PaymentRecordPaymentMethodDetailsAmazonPayFundingCard `json:"card"`
	// funding type of the underlying payment method.
	Type PaymentRecordPaymentMethodDetailsAmazonPayFundingType `json:"type"`
}
type PaymentRecordPaymentMethodDetailsAmazonPay struct {
	Funding *PaymentRecordPaymentMethodDetailsAmazonPayFunding `json:"funding"`
	// The Amazon Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsAUBECSDebit struct {
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
type PaymentRecordPaymentMethodDetailsBACSDebit struct {
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
type PaymentRecordPaymentMethodDetailsBancontact struct {
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
	// Preferred language of the Bancontact authorization page that the customer is redirected to. Can be one of `en`, `de`, `fr`, or `nl`
	PreferredLanguage PaymentRecordPaymentMethodDetailsBancontactPreferredLanguage `json:"preferred_language"`
	// Owner's verified full name. Values are verified or provided by Bancontact directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type PaymentRecordPaymentMethodDetailsBillie struct {
	// The Billie transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}

// The billing details associated with the method of payment.
type PaymentRecordPaymentMethodDetailsBillingDetails struct {
	// A representation of a physical address.
	Address *Address `json:"address"`
	// The billing email associated with the method of payment.
	Email string `json:"email"`
	// The billing name associated with the method of payment.
	Name string `json:"name"`
	// The billing phone number associated with the method of payment.
	Phone string `json:"phone"`
}
type PaymentRecordPaymentMethodDetailsBLIK struct {
	// A unique and immutable identifier assigned by BLIK to every buyer.
	BuyerID string `json:"buyer_id"`
}
type PaymentRecordPaymentMethodDetailsBoleto struct {
	// The tax ID of the customer (CPF for individuals consumers or CNPJ for businesses consumers)
	TaxID string `json:"tax_id"`
}

// Check results by Card networks on Card address and CVC at time of payment.
type PaymentRecordPaymentMethodDetailsCardChecks struct {
	// If you provide a value for `address.line1`, the check result is one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressLine1Check PaymentRecordPaymentMethodDetailsCardChecksAddressLine1Check `json:"address_line1_check"`
	// If you provide a address postal code, the check result is one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressPostalCodeCheck PaymentRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck `json:"address_postal_code_check"`
	// If you provide a CVC, the check results is one of `pass`, `fail`, `unavailable`, or `unchecked`.
	CVCCheck PaymentRecordPaymentMethodDetailsCardChecksCVCCheck `json:"cvc_check"`
}

// Installment plan selected for the payment.
type PaymentRecordPaymentMethodDetailsCardInstallmentsPlan struct {
	// For `fixed_count` installment plans, this is the number of installment payments your customer will make to their credit card.
	Count int64 `json:"count"`
	// For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card. One of `month`.
	Interval PaymentRecordPaymentMethodDetailsCardInstallmentsPlanInterval `json:"interval"`
	// Type of installment plan, one of `fixed_count`, `revolving`, or `bonus`.
	Type PaymentRecordPaymentMethodDetailsCardInstallmentsPlanType `json:"type"`
}

// Installment details for this payment.
type PaymentRecordPaymentMethodDetailsCardInstallments struct {
	// Installment plan selected for the payment.
	Plan *PaymentRecordPaymentMethodDetailsCardInstallmentsPlan `json:"plan"`
}

// If this card has network token credentials, this contains the details of the network token credentials.
type PaymentRecordPaymentMethodDetailsCardNetworkToken struct {
	// Indicates if Stripe used a network token, either user provided or Stripe managed when processing the transaction.
	Used bool `json:"used"`
}

// Populated if this transaction used 3D Secure authentication.
type PaymentRecordPaymentMethodDetailsCardThreeDSecure struct {
	// For authenticated transactions: Indicates how the issuing bank authenticated the customer.
	AuthenticationFlow PaymentRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	// Indicates the outcome of 3D Secure authentication.
	Result PaymentRecordPaymentMethodDetailsCardThreeDSecureResult `json:"result"`
	// Additional information about why 3D Secure succeeded or failed, based on the `result`.
	ResultReason PaymentRecordPaymentMethodDetailsCardThreeDSecureResultReason `json:"result_reason"`
	// The version of 3D Secure that was used.
	Version PaymentRecordPaymentMethodDetailsCardThreeDSecureVersion `json:"version"`
}
type PaymentRecordPaymentMethodDetailsCardWalletApplePay struct {
	// Type of the apple_pay transaction, one of `apple_pay` or `apple_pay_later`.
	Type string `json:"type"`
}
type PaymentRecordPaymentMethodDetailsCardWalletGooglePay struct{}

// If this Card is part of a card wallet, this contains the details of the card wallet.
type PaymentRecordPaymentMethodDetailsCardWallet struct {
	ApplePay *PaymentRecordPaymentMethodDetailsCardWalletApplePay `json:"apple_pay"`
	// (For tokenized numbers only.) The last four digits of the device account number.
	DynamicLast4 string                                                `json:"dynamic_last4"`
	GooglePay    *PaymentRecordPaymentMethodDetailsCardWalletGooglePay `json:"google_pay"`
	// The type of the card wallet, one of `apple_pay` or `google_pay`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
	Type string `json:"type"`
}

// Details of the card used for this payment attempt.
type PaymentRecordPaymentMethodDetailsCard struct {
	// The authorization code of the payment.
	AuthorizationCode string `json:"authorization_code"`
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand PaymentRecordPaymentMethodDetailsCardBrand `json:"brand"`
	// When using manual capture, a future timestamp at which the charge will be automatically refunded if uncaptured.
	CaptureBefore int64 `json:"capture_before"`
	// Check results by Card networks on Card address and CVC at time of payment.
	Checks *PaymentRecordPaymentMethodDetailsCardChecks `json:"checks"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// A high-level description of the type of cards issued in this range.
	Description string `json:"description"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *As of May 1, 2021, card fingerprint in India for Connect changed to allow two fingerprints for the same card---one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding PaymentRecordPaymentMethodDetailsCardFunding `json:"funding"`
	// Issuer identification number of the card.
	IIN string `json:"iin"`
	// Installment details for this payment.
	Installments *PaymentRecordPaymentMethodDetailsCardInstallments `json:"installments"`
	// The name of the card's issuing bank.
	Issuer string `json:"issuer"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// True if this payment was marked as MOTO and out of scope for SCA.
	MOTO bool `json:"moto"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network PaymentRecordPaymentMethodDetailsCardNetwork `json:"network"`
	// Advice code from the card network for the failed payment.
	NetworkAdviceCode string `json:"network_advice_code"`
	// Decline code from the card network for the failed payment.
	NetworkDeclineCode string `json:"network_decline_code"`
	// If this card has network token credentials, this contains the details of the network token credentials.
	NetworkToken *PaymentRecordPaymentMethodDetailsCardNetworkToken `json:"network_token"`
	// This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.
	NetworkTransactionID string `json:"network_transaction_id"`
	// The transaction type that was passed for an off-session, Merchant-Initiated transaction, one of `recurring` or `unscheduled`.
	StoredCredentialUsage PaymentRecordPaymentMethodDetailsCardStoredCredentialUsage `json:"stored_credential_usage"`
	// Populated if this transaction used 3D Secure authentication.
	ThreeDSecure *PaymentRecordPaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
	// If this Card is part of a card wallet, this contains the details of the card wallet.
	Wallet *PaymentRecordPaymentMethodDetailsCardWallet `json:"wallet"`
}

// Details about payments collected offline.
type PaymentRecordPaymentMethodDetailsCardPresentOffline struct {
	// Time at which the payment was collected while offline
	StoredAt int64 `json:"stored_at"`
	// The method used to process this payment method offline. Only deferred is allowed.
	Type PaymentRecordPaymentMethodDetailsCardPresentOfflineType `json:"type"`
}

// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
type PaymentRecordPaymentMethodDetailsCardPresentReceipt struct {
	// The type of account being debited or credited
	AccountType PaymentRecordPaymentMethodDetailsCardPresentReceiptAccountType `json:"account_type"`
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
type PaymentRecordPaymentMethodDetailsCardPresentWallet struct {
	// The type of mobile wallet, one of `apple_pay`, `google_pay`, `samsung_pay`, or `unknown`.
	Type PaymentRecordPaymentMethodDetailsCardPresentWalletType `json:"type"`
}
type PaymentRecordPaymentMethodDetailsCardPresent struct {
	// The authorized amount
	AmountAuthorized int64 `json:"amount_authorized"`
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand string `json:"brand"`
	// The [product code](https://stripe.com/docs/card-product-codes) that identifies the specific program or product associated with a card.
	BrandProduct string `json:"brand_product"`
	// When using manual capture, a future timestamp after which the charge will be automatically refunded if uncaptured.
	CaptureBefore int64 `json:"capture_before"`
	// The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.
	CardholderName string `json:"cardholder_name"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description"`
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
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin"`
	// Whether this [PaymentIntent](https://docs.stripe.com/api/payment_intents) is eligible for incremental authorizations. Request support using [request_incremental_authorization_support](https://docs.stripe.com/api/payment_intents/create#create_payment_intent-payment_method_options-card_present-request_incremental_authorization_support).
	IncrementalAuthorizationSupported bool `json:"incremental_authorization_supported"`
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.
	Location string `json:"location"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network string `json:"network"`
	// This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.
	NetworkTransactionID string `json:"network_transaction_id"`
	// Details about payments collected offline.
	Offline *PaymentRecordPaymentMethodDetailsCardPresentOffline `json:"offline"`
	// Defines whether the authorized amount can be over-captured or not
	OvercaptureSupported bool `json:"overcapture_supported"`
	// The languages that the issuing bank recommends using for localizing any customer-facing text, as read from the card. Referenced from EMV tag 5F2D, data encoded on the card's chip.
	PreferredLocales []string `json:"preferred_locales"`
	// ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.
	Reader string `json:"reader"`
	// How card details were read in this transaction.
	ReadMethod PaymentRecordPaymentMethodDetailsCardPresentReadMethod `json:"read_method"`
	// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
	Receipt *PaymentRecordPaymentMethodDetailsCardPresentReceipt `json:"receipt"`
	Wallet  *PaymentRecordPaymentMethodDetailsCardPresentWallet  `json:"wallet"`
}
type PaymentRecordPaymentMethodDetailsCashApp struct {
	// A unique and immutable identifier assigned by Cash App to every buyer.
	BuyerID string `json:"buyer_id"`
	// A public identifier for buyers using Cash App.
	Cashtag string `json:"cashtag"`
	// A unique and immutable identifier of payments assigned by Cash App.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsCrypto struct {
	// The wallet address of the customer.
	BuyerAddress string `json:"buyer_address"`
	// The blockchain network that the transaction was sent on.
	Network PaymentRecordPaymentMethodDetailsCryptoNetwork `json:"network"`
	// The token currency that the transaction was sent with.
	TokenCurrency PaymentRecordPaymentMethodDetailsCryptoTokenCurrency `json:"token_currency"`
	// The blockchain transaction hash of the crypto payment.
	TransactionHash string `json:"transaction_hash"`
}

// Custom Payment Methods represent Payment Method types not modeled directly in
// the Stripe API. This resource consists of details about the custom payment method
// used for this payment attempt.
type PaymentRecordPaymentMethodDetailsCustom struct {
	// Display name for the custom (user-defined) payment method type used to make this payment.
	DisplayName string `json:"display_name"`
	// The custom payment method type associated with this payment.
	Type string `json:"type"`
}
type PaymentRecordPaymentMethodDetailsCustomerBalance struct{}
type PaymentRecordPaymentMethodDetailsEPS struct {
	// The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `deutsche_bank_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.
	Bank PaymentRecordPaymentMethodDetailsEPSBank `json:"bank"`
	// Owner's verified full name. Values are verified or provided by EPS directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	// EPS rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type PaymentRecordPaymentMethodDetailsFPX struct {
	// Account holder type, if provided. Can be one of `individual` or `company`.
	AccountHolderType PaymentRecordPaymentMethodDetailsFPXAccountHolderType `json:"account_holder_type"`
	// The customer's bank. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, `pb_enterprise`, or `bank_of_china`.
	Bank PaymentRecordPaymentMethodDetailsFPXBank `json:"bank"`
	// Unique transaction id generated by FPX for every request from the merchant
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsGiropay struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	BIC string `json:"bic"`
	// Owner's verified full name. Values are verified or provided by Giropay directly (if supported) at the time of authorization or settlement. They cannot be set or mutated. Giropay rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type PaymentRecordPaymentMethodDetailsGrabpay struct {
	// Unique transaction id generated by GrabPay
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsIDEAL struct {
	// The customer's bank. Can be one of `abn_amro`, `adyen`, `asn_bank`, `bunq`, `buut`, `finom`, `handelsbanken`, `ing`, `knab`, `mollie`, `moneyou`, `n26`, `nn`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, `van_lanschot`, or `yoursafe`.
	Bank PaymentRecordPaymentMethodDetailsIDEALBank `json:"bank"`
	// The Bank Identifier Code of the customer's bank.
	BIC PaymentRecordPaymentMethodDetailsIDEALBIC `json:"bic"`
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
type PaymentRecordPaymentMethodDetailsInteracPresentReceipt struct {
	// The type of account being debited or credited
	AccountType PaymentRecordPaymentMethodDetailsInteracPresentReceiptAccountType `json:"account_type"`
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
type PaymentRecordPaymentMethodDetailsInteracPresent struct {
	// Card brand. Can be `interac`, `mastercard` or `visa`.
	Brand string `json:"brand"`
	// The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.
	CardholderName string `json:"cardholder_name"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description"`
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
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin"`
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.
	Location string `json:"location"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network string `json:"network"`
	// This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.
	NetworkTransactionID string `json:"network_transaction_id"`
	// The languages that the issuing bank recommends using for localizing any customer-facing text, as read from the card. Referenced from EMV tag 5F2D, data encoded on the card's chip.
	PreferredLocales []string `json:"preferred_locales"`
	// ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.
	Reader string `json:"reader"`
	// How card details were read in this transaction.
	ReadMethod PaymentRecordPaymentMethodDetailsInteracPresentReadMethod `json:"read_method"`
	// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
	Receipt *PaymentRecordPaymentMethodDetailsInteracPresentReceipt `json:"receipt"`
}
type PaymentRecordPaymentMethodDetailsKakaoPay struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Kakao Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}

// The payer's address
type PaymentRecordPaymentMethodDetailsKlarnaPayerDetailsAddress struct {
	// The payer address country
	Country string `json:"country"`
}

// The payer details for this transaction.
type PaymentRecordPaymentMethodDetailsKlarnaPayerDetails struct {
	// The payer's address
	Address *PaymentRecordPaymentMethodDetailsKlarnaPayerDetailsAddress `json:"address"`
}
type PaymentRecordPaymentMethodDetailsKlarna struct {
	// The payer details for this transaction.
	PayerDetails *PaymentRecordPaymentMethodDetailsKlarnaPayerDetails `json:"payer_details"`
	// The Klarna payment method used for this transaction.
	// Can be one of `pay_later`, `pay_now`, `pay_with_financing`, or `pay_in_installments`
	PaymentMethodCategory string `json:"payment_method_category"`
	// Preferred language of the Klarna authorization page that the customer is redirected to.
	// Can be one of `de-AT`, `en-AT`, `nl-BE`, `fr-BE`, `en-BE`, `de-DE`, `en-DE`, `da-DK`, `en-DK`, `es-ES`, `en-ES`, `fi-FI`, `sv-FI`, `en-FI`, `en-GB`, `en-IE`, `it-IT`, `en-IT`, `nl-NL`, `en-NL`, `nb-NO`, `en-NO`, `sv-SE`, `en-SE`, `en-US`, `es-US`, `fr-FR`, `en-FR`, `cs-CZ`, `en-CZ`, `ro-RO`, `en-RO`, `el-GR`, `en-GR`, `en-AU`, `en-NZ`, `en-CA`, `fr-CA`, `pl-PL`, `en-PL`, `pt-PT`, `en-PT`, `de-CH`, `fr-CH`, `it-CH`, or `en-CH`
	PreferredLocale string `json:"preferred_locale"`
}

// If the payment succeeded, this contains the details of the convenience store where the payment was completed.
type PaymentRecordPaymentMethodDetailsKonbiniStore struct {
	// The name of the convenience store chain where the payment was completed.
	Chain PaymentRecordPaymentMethodDetailsKonbiniStoreChain `json:"chain"`
}
type PaymentRecordPaymentMethodDetailsKonbini struct {
	// If the payment succeeded, this contains the details of the convenience store where the payment was completed.
	Store *PaymentRecordPaymentMethodDetailsKonbiniStore `json:"store"`
}
type PaymentRecordPaymentMethodDetailsKrCard struct {
	// The local credit or debit card brand.
	Brand PaymentRecordPaymentMethodDetailsKrCardBrand `json:"brand"`
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The last four digits of the card. This may not be present for American Express cards.
	Last4 string `json:"last4"`
	// The Korean Card transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsLink struct {
	// Two-letter ISO code representing the funding source country beneath the Link payment.
	// You could use this attribute to get a sense of international fees.
	Country string `json:"country"`
}
type PaymentRecordPaymentMethodDetailsMbWay struct{}

// Internal card details
type PaymentRecordPaymentMethodDetailsMobilepayCard struct {
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
type PaymentRecordPaymentMethodDetailsMobilepay struct {
	// Internal card details
	Card *PaymentRecordPaymentMethodDetailsMobilepayCard `json:"card"`
}
type PaymentRecordPaymentMethodDetailsMultibanco struct {
	// Entity number associated with this Multibanco payment.
	Entity string `json:"entity"`
	// Reference number associated with this Multibanco payment.
	Reference string `json:"reference"`
}
type PaymentRecordPaymentMethodDetailsNaverPay struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Naver Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsNzBankAccount struct {
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
type PaymentRecordPaymentMethodDetailsOXXO struct {
	// OXXO reference number
	Number string `json:"number"`
}
type PaymentRecordPaymentMethodDetailsP24 struct {
	// The customer's bank. Can be one of `ing`, `citi_handlowy`, `tmobile_usbugi_bankowe`, `plus_bank`, `etransfer_pocztowy24`, `banki_spbdzielcze`, `bank_nowy_bfg_sa`, `getin_bank`, `velobank`, `blik`, `noble_pay`, `ideabank`, `envelobank`, `santander_przelew24`, `nest_przelew`, `mbank_mtransfer`, `inteligo`, `pbac_z_ipko`, `bnp_paribas`, `credit_agricole`, `toyota_bank`, `bank_pekao_sa`, `volkswagen_bank`, `bank_millennium`, `alior_bank`, or `boz`.
	Bank PaymentRecordPaymentMethodDetailsP24Bank `json:"bank"`
	// Unique reference for this Przelewy24 payment.
	Reference string `json:"reference"`
	// Owner's verified full name. Values are verified or provided by Przelewy24 directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	// Przelewy24 rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type PaymentRecordPaymentMethodDetailsPayByBank struct{}
type PaymentRecordPaymentMethodDetailsPayco struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Payco transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsPayNow struct {
	// ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.
	Location string `json:"location"`
	// ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.
	Reader string `json:"reader"`
	// Reference number associated with this PayNow payment
	Reference string `json:"reference"`
}

// The level of protection offered as defined by PayPal Seller Protection for Merchants, for this transaction.
type PaymentRecordPaymentMethodDetailsPaypalSellerProtection struct {
	// An array of conditions that are covered for the transaction, if applicable.
	DisputeCategories []PaymentRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategory `json:"dispute_categories"`
	// Indicates whether the transaction is eligible for PayPal's seller protection.
	Status PaymentRecordPaymentMethodDetailsPaypalSellerProtectionStatus `json:"status"`
}
type PaymentRecordPaymentMethodDetailsPaypal struct {
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
	SellerProtection *PaymentRecordPaymentMethodDetailsPaypalSellerProtection `json:"seller_protection"`
	// A unique ID generated by PayPal for this transaction.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsPayto struct {
	// Bank-State-Branch number of the bank account.
	BSBNumber string `json:"bsb_number"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate string `json:"mandate"`
	// The PayID alias for the bank account.
	PayID string `json:"pay_id"`
}
type PaymentRecordPaymentMethodDetailsPix struct {
	// Unique transaction id generated by BCB
	BankTransactionID string `json:"bank_transaction_id"`
}
type PaymentRecordPaymentMethodDetailsPromptPay struct {
	// Bill reference generated by PromptPay
	Reference string `json:"reference"`
}
type PaymentRecordPaymentMethodDetailsRevolutPayFundingCard struct {
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand string `json:"brand"`
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
type PaymentRecordPaymentMethodDetailsRevolutPayFunding struct {
	Card *PaymentRecordPaymentMethodDetailsRevolutPayFundingCard `json:"card"`
	// funding type of the underlying payment method.
	Type PaymentRecordPaymentMethodDetailsRevolutPayFundingType `json:"type"`
}
type PaymentRecordPaymentMethodDetailsRevolutPay struct {
	Funding *PaymentRecordPaymentMethodDetailsRevolutPayFunding `json:"funding"`
	// The Revolut Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsSamsungPay struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Samsung Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsSatispay struct {
	// The Satispay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsSEPACreditTransfer struct {
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	BIC string `json:"bic"`
	// IBAN of the bank account to transfer funds to.
	IBAN string `json:"iban"`
}
type PaymentRecordPaymentMethodDetailsSEPADebit struct {
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
type PaymentRecordPaymentMethodDetailsSofort struct {
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
	PreferredLanguage PaymentRecordPaymentMethodDetailsSofortPreferredLanguage `json:"preferred_language"`
	// Owner's verified full name. Values are verified or provided by SOFORT directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type PaymentRecordPaymentMethodDetailsStripeAccount struct{}
type PaymentRecordPaymentMethodDetailsSwish struct {
	// Uniquely identifies the payer's Swish account. You can use this attribute to check whether two Swish transactions were paid for by the same payer
	Fingerprint string `json:"fingerprint"`
	// Payer bank reference number for the payment
	PaymentReference string `json:"payment_reference"`
	// The last four digits of the Swish account phone number
	VerifiedPhoneLast4 string `json:"verified_phone_last4"`
}
type PaymentRecordPaymentMethodDetailsTWINT struct{}
type PaymentRecordPaymentMethodDetailsUSBankAccount struct {
	// The type of entity that holds the account. This can be either 'individual' or 'company'.
	AccountHolderType PaymentRecordPaymentMethodDetailsUSBankAccountAccountHolderType `json:"account_holder_type"`
	// The type of the bank account. This can be either 'checking' or 'savings'.
	AccountType PaymentRecordPaymentMethodDetailsUSBankAccountAccountType `json:"account_type"`
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
	// The ACH payment reference for this transaction.
	PaymentReference string `json:"payment_reference"`
	// The routing number for the bank account.
	RoutingNumber string `json:"routing_number"`
}
type PaymentRecordPaymentMethodDetailsWeChat struct{}
type PaymentRecordPaymentMethodDetailsWeChatPay struct {
	// Uniquely identifies this particular WeChat Pay account. You can use this attribute to check whether two WeChat accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.
	Location string `json:"location"`
	// ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.
	Reader string `json:"reader"`
	// Transaction ID of this particular WeChat Pay transaction.
	TransactionID string `json:"transaction_id"`
}
type PaymentRecordPaymentMethodDetailsZip struct{}

// Information about the Payment Method debited for this payment.
type PaymentRecordPaymentMethodDetails struct {
	ACHCreditTransfer *PaymentRecordPaymentMethodDetailsACHCreditTransfer `json:"ach_credit_transfer"`
	ACHDebit          *PaymentRecordPaymentMethodDetailsACHDebit          `json:"ach_debit"`
	ACSSDebit         *PaymentRecordPaymentMethodDetailsACSSDebit         `json:"acss_debit"`
	Affirm            *PaymentRecordPaymentMethodDetailsAffirm            `json:"affirm"`
	AfterpayClearpay  *PaymentRecordPaymentMethodDetailsAfterpayClearpay  `json:"afterpay_clearpay"`
	Alipay            *PaymentRecordPaymentMethodDetailsAlipay            `json:"alipay"`
	Alma              *PaymentRecordPaymentMethodDetailsAlma              `json:"alma"`
	AmazonPay         *PaymentRecordPaymentMethodDetailsAmazonPay         `json:"amazon_pay"`
	AUBECSDebit       *PaymentRecordPaymentMethodDetailsAUBECSDebit       `json:"au_becs_debit"`
	BACSDebit         *PaymentRecordPaymentMethodDetailsBACSDebit         `json:"bacs_debit"`
	Bancontact        *PaymentRecordPaymentMethodDetailsBancontact        `json:"bancontact"`
	Billie            *PaymentRecordPaymentMethodDetailsBillie            `json:"billie"`
	// The billing details associated with the method of payment.
	BillingDetails *PaymentRecordPaymentMethodDetailsBillingDetails `json:"billing_details"`
	BLIK           *PaymentRecordPaymentMethodDetailsBLIK           `json:"blik"`
	Boleto         *PaymentRecordPaymentMethodDetailsBoleto         `json:"boleto"`
	// Details of the card used for this payment attempt.
	Card        *PaymentRecordPaymentMethodDetailsCard        `json:"card"`
	CardPresent *PaymentRecordPaymentMethodDetailsCardPresent `json:"card_present"`
	CashApp     *PaymentRecordPaymentMethodDetailsCashApp     `json:"cashapp"`
	Crypto      *PaymentRecordPaymentMethodDetailsCrypto      `json:"crypto"`
	// Custom Payment Methods represent Payment Method types not modeled directly in
	// the Stripe API. This resource consists of details about the custom payment method
	// used for this payment attempt.
	Custom          *PaymentRecordPaymentMethodDetailsCustom          `json:"custom"`
	CustomerBalance *PaymentRecordPaymentMethodDetailsCustomerBalance `json:"customer_balance"`
	EPS             *PaymentRecordPaymentMethodDetailsEPS             `json:"eps"`
	FPX             *PaymentRecordPaymentMethodDetailsFPX             `json:"fpx"`
	Giropay         *PaymentRecordPaymentMethodDetailsGiropay         `json:"giropay"`
	Grabpay         *PaymentRecordPaymentMethodDetailsGrabpay         `json:"grabpay"`
	IDEAL           *PaymentRecordPaymentMethodDetailsIDEAL           `json:"ideal"`
	InteracPresent  *PaymentRecordPaymentMethodDetailsInteracPresent  `json:"interac_present"`
	KakaoPay        *PaymentRecordPaymentMethodDetailsKakaoPay        `json:"kakao_pay"`
	Klarna          *PaymentRecordPaymentMethodDetailsKlarna          `json:"klarna"`
	Konbini         *PaymentRecordPaymentMethodDetailsKonbini         `json:"konbini"`
	KrCard          *PaymentRecordPaymentMethodDetailsKrCard          `json:"kr_card"`
	Link            *PaymentRecordPaymentMethodDetailsLink            `json:"link"`
	MbWay           *PaymentRecordPaymentMethodDetailsMbWay           `json:"mb_way"`
	Mobilepay       *PaymentRecordPaymentMethodDetailsMobilepay       `json:"mobilepay"`
	Multibanco      *PaymentRecordPaymentMethodDetailsMultibanco      `json:"multibanco"`
	NaverPay        *PaymentRecordPaymentMethodDetailsNaverPay        `json:"naver_pay"`
	NzBankAccount   *PaymentRecordPaymentMethodDetailsNzBankAccount   `json:"nz_bank_account"`
	OXXO            *PaymentRecordPaymentMethodDetailsOXXO            `json:"oxxo"`
	P24             *PaymentRecordPaymentMethodDetailsP24             `json:"p24"`
	PayByBank       *PaymentRecordPaymentMethodDetailsPayByBank       `json:"pay_by_bank"`
	Payco           *PaymentRecordPaymentMethodDetailsPayco           `json:"payco"`
	// ID of the Stripe PaymentMethod used to make this payment.
	PaymentMethod      string                                               `json:"payment_method"`
	PayNow             *PaymentRecordPaymentMethodDetailsPayNow             `json:"paynow"`
	Paypal             *PaymentRecordPaymentMethodDetailsPaypal             `json:"paypal"`
	Payto              *PaymentRecordPaymentMethodDetailsPayto              `json:"payto"`
	Pix                *PaymentRecordPaymentMethodDetailsPix                `json:"pix"`
	PromptPay          *PaymentRecordPaymentMethodDetailsPromptPay          `json:"promptpay"`
	RevolutPay         *PaymentRecordPaymentMethodDetailsRevolutPay         `json:"revolut_pay"`
	SamsungPay         *PaymentRecordPaymentMethodDetailsSamsungPay         `json:"samsung_pay"`
	Satispay           *PaymentRecordPaymentMethodDetailsSatispay           `json:"satispay"`
	SEPACreditTransfer *PaymentRecordPaymentMethodDetailsSEPACreditTransfer `json:"sepa_credit_transfer"`
	SEPADebit          *PaymentRecordPaymentMethodDetailsSEPADebit          `json:"sepa_debit"`
	Sofort             *PaymentRecordPaymentMethodDetailsSofort             `json:"sofort"`
	StripeAccount      *PaymentRecordPaymentMethodDetailsStripeAccount      `json:"stripe_account"`
	Swish              *PaymentRecordPaymentMethodDetailsSwish              `json:"swish"`
	TWINT              *PaymentRecordPaymentMethodDetailsTWINT              `json:"twint"`
	// The type of transaction-specific details of the payment method used in the payment. See [PaymentMethod.type](https://docs.stripe.com/api/payment_methods/object#payment_method_object-type) for the full list of possible types.
	// An additional hash is included on `payment_method_details` with a name matching this value.
	// It contains information specific to the payment method.
	Type          string                                          `json:"type"`
	USBankAccount *PaymentRecordPaymentMethodDetailsUSBankAccount `json:"us_bank_account"`
	WeChat        *PaymentRecordPaymentMethodDetailsWeChat        `json:"wechat"`
	WeChatPay     *PaymentRecordPaymentMethodDetailsWeChatPay     `json:"wechat_pay"`
	Zip           *PaymentRecordPaymentMethodDetailsZip           `json:"zip"`
}

// Custom processors represent payment processors not modeled directly in
// the Stripe API. This resource consists of details about the custom processor
// used for this payment attempt.
type PaymentRecordProcessorDetailsCustom struct {
	// An opaque string for manual reconciliation of this payment, for example a check number or a payment processor ID.
	PaymentReference string `json:"payment_reference"`
}

// Processor information associated with this payment.
type PaymentRecordProcessorDetails struct {
	// Custom processors represent payment processors not modeled directly in
	// the Stripe API. This resource consists of details about the custom processor
	// used for this payment attempt.
	Custom *PaymentRecordProcessorDetailsCustom `json:"custom"`
	// The processor used for this payment attempt.
	Type PaymentRecordProcessorDetailsType `json:"type"`
}

// Shipping information for this payment.
type PaymentRecordShippingDetails struct {
	// A representation of a physical address.
	Address *Address `json:"address"`
	// The shipping recipient's name.
	Name string `json:"name"`
	// The shipping recipient's phone number.
	Phone string `json:"phone"`
}

// A Payment Record is a resource that allows you to represent payments that occur on- or off-Stripe.
// For example, you can create a Payment Record to model a payment made on a different payment processor,
// in order to mark an Invoice as paid and a Subscription as active. Payment Records consist of one or
// more Payment Attempt Records, which represent individual attempts made on a payment network.
type PaymentRecord struct {
	APIResource
	// A representation of an amount of money, consisting of an amount and a currency.
	Amount *PaymentRecordAmount `json:"amount"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountAuthorized *PaymentRecordAmountAuthorized `json:"amount_authorized"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountCanceled *PaymentRecordAmountCanceled `json:"amount_canceled"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountFailed *PaymentRecordAmountFailed `json:"amount_failed"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountGuaranteed *PaymentRecordAmountGuaranteed `json:"amount_guaranteed"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountRefunded *PaymentRecordAmountRefunded `json:"amount_refunded"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountRequested *PaymentRecordAmountRequested `json:"amount_requested"`
	// ID of the Connect application that created the PaymentRecord.
	Application string `json:"application"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Customer information for this payment.
	CustomerDetails *PaymentRecordCustomerDetails `json:"customer_details"`
	// Indicates whether the customer was present in your checkout flow during this payment.
	CustomerPresence PaymentRecordCustomerPresence `json:"customer_presence"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// ID of the latest Payment Attempt Record attached to this Payment Record.
	LatestPaymentAttemptRecord string `json:"latest_payment_attempt_record"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Information about the Payment Method debited for this payment.
	PaymentMethodDetails *PaymentRecordPaymentMethodDetails `json:"payment_method_details"`
	// Processor information associated with this payment.
	ProcessorDetails *PaymentRecordProcessorDetails `json:"processor_details"`
	// Indicates who reported the payment.
	ReportedBy PaymentRecordReportedBy `json:"reported_by"`
	// Shipping information for this payment.
	ShippingDetails *PaymentRecordShippingDetails `json:"shipping_details"`
}

// UnmarshalJSON handles deserialization of a PaymentRecord.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PaymentRecord) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type paymentRecord PaymentRecord
	var v paymentRecord
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PaymentRecord(v)
	return nil
}
