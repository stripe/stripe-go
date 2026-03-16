//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Indicates whether the customer was present in your checkout flow during this payment.
type PaymentAttemptRecordCustomerPresence string

// List of values that PaymentAttemptRecordCustomerPresence can take
const (
	PaymentAttemptRecordCustomerPresenceOffSession PaymentAttemptRecordCustomerPresence = "off_session"
	PaymentAttemptRecordCustomerPresenceOnSession  PaymentAttemptRecordCustomerPresence = "on_session"
)

// Type of entity that holds the account. This can be either `individual` or `company`.
type PaymentAttemptRecordPaymentMethodDetailsACHDebitAccountHolderType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsACHDebitAccountHolderType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsACHDebitAccountHolderTypeCompany    PaymentAttemptRecordPaymentMethodDetailsACHDebitAccountHolderType = "company"
	PaymentAttemptRecordPaymentMethodDetailsACHDebitAccountHolderTypeIndividual PaymentAttemptRecordPaymentMethodDetailsACHDebitAccountHolderType = "individual"
)

// funding type of the underlying payment method.
type PaymentAttemptRecordPaymentMethodDetailsAmazonPayFundingType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsAmazonPayFundingType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsAmazonPayFundingTypeCard PaymentAttemptRecordPaymentMethodDetailsAmazonPayFundingType = "card"
)

// Preferred language of the Bancontact authorization page that the customer is redirected to. Can be one of `en`, `de`, `fr`, or `nl`
type PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguage string

// List of values that PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguage can take
const (
	PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguageDE PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguage = "de"
	PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguageEN PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguage = "en"
	PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguageFR PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguage = "fr"
	PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguageNL PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguage = "nl"
)

// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
type PaymentAttemptRecordPaymentMethodDetailsCardBrand string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardBrand can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardBrandAmex            PaymentAttemptRecordPaymentMethodDetailsCardBrand = "amex"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandCartesBancaires PaymentAttemptRecordPaymentMethodDetailsCardBrand = "cartes_bancaires"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandDiners          PaymentAttemptRecordPaymentMethodDetailsCardBrand = "diners"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandDiscover        PaymentAttemptRecordPaymentMethodDetailsCardBrand = "discover"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandEFTPOSAU        PaymentAttemptRecordPaymentMethodDetailsCardBrand = "eftpos_au"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandInterac         PaymentAttemptRecordPaymentMethodDetailsCardBrand = "interac"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandJCB             PaymentAttemptRecordPaymentMethodDetailsCardBrand = "jcb"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandLink            PaymentAttemptRecordPaymentMethodDetailsCardBrand = "link"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandMastercard      PaymentAttemptRecordPaymentMethodDetailsCardBrand = "mastercard"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandUnionpay        PaymentAttemptRecordPaymentMethodDetailsCardBrand = "unionpay"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandUnknown         PaymentAttemptRecordPaymentMethodDetailsCardBrand = "unknown"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandVisa            PaymentAttemptRecordPaymentMethodDetailsCardBrand = "visa"
)

// If you provide a value for `address.line1`, the check result is one of `pass`, `fail`, `unavailable`, or `unchecked`.
type PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1CheckFail        PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check = "fail"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1CheckPass        PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check = "pass"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1CheckUnavailable PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check = "unavailable"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1CheckUnchecked   PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check = "unchecked"
)

// If you provide a address postal code, the check result is one of `pass`, `fail`, `unavailable`, or `unchecked`.
type PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckFail        PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "fail"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckPass        PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "pass"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckUnavailable PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "unavailable"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckUnchecked   PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "unchecked"
)

// If you provide a CVC, the check results is one of `pass`, `fail`, `unavailable`, or `unchecked`.
type PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheckFail        PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck = "fail"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheckPass        PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck = "pass"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheckUnavailable PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck = "unavailable"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheckUnchecked   PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck = "unchecked"
)

// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
type PaymentAttemptRecordPaymentMethodDetailsCardFunding string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardFunding can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardFundingCredit  PaymentAttemptRecordPaymentMethodDetailsCardFunding = "credit"
	PaymentAttemptRecordPaymentMethodDetailsCardFundingDebit   PaymentAttemptRecordPaymentMethodDetailsCardFunding = "debit"
	PaymentAttemptRecordPaymentMethodDetailsCardFundingPrepaid PaymentAttemptRecordPaymentMethodDetailsCardFunding = "prepaid"
	PaymentAttemptRecordPaymentMethodDetailsCardFundingUnknown PaymentAttemptRecordPaymentMethodDetailsCardFunding = "unknown"
)

// For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card. One of `month`.
type PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanInterval string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanInterval can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanIntervalMonth PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanInterval = "month"
)

// Type of installment plan, one of `fixed_count`, `revolving`, or `bonus`.
type PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanTypeBonus      PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanType = "bonus"
	PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanTypeFixedCount PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanType = "fixed_count"
	PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanTypeRevolving  PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanType = "revolving"
)

// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
type PaymentAttemptRecordPaymentMethodDetailsCardNetwork string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardNetwork can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkAmex            PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "amex"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkCartesBancaires PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "cartes_bancaires"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkDiners          PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "diners"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkDiscover        PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "discover"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkEFTPOSAU        PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "eftpos_au"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkInterac         PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "interac"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkJCB             PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "jcb"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkLink            PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "link"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkMastercard      PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "mastercard"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkUnionpay        PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "unionpay"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkUnknown         PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "unknown"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkVisa            PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "visa"
)

// The transaction type that was passed for an off-session, Merchant-Initiated transaction, one of `recurring` or `unscheduled`.
type PaymentAttemptRecordPaymentMethodDetailsCardStoredCredentialUsage string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardStoredCredentialUsage can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardStoredCredentialUsageRecurring   PaymentAttemptRecordPaymentMethodDetailsCardStoredCredentialUsage = "recurring"
	PaymentAttemptRecordPaymentMethodDetailsCardStoredCredentialUsageUnscheduled PaymentAttemptRecordPaymentMethodDetailsCardStoredCredentialUsage = "unscheduled"
)

// For authenticated transactions: Indicates how the issuing bank authenticated the customer.
type PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// Indicates the outcome of 3D Secure authentication.
type PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultAuthenticated       PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultExempted            PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "exempted"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultFailed              PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "failed"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultNotSupported        PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultProcessingError     PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// Additional information about why 3D Secure succeeded or failed, based on the `result`.
type PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonBypassed            PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonCanceled            PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonRejected            PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

// The version of 3D Secure that was used.
type PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion102 PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion = "1.0.2"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion210 PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion = "2.1.0"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion220 PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion = "2.2.0"
)

// The method used to process this payment method offline. Only deferred is allowed.
type PaymentAttemptRecordPaymentMethodDetailsCardPresentOfflineType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardPresentOfflineType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardPresentOfflineTypeDeferred PaymentAttemptRecordPaymentMethodDetailsCardPresentOfflineType = "deferred"
)

// How card details were read in this transaction.
type PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethod string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethod can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethodContactEmv               PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethod = "contact_emv"
	PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethodContactlessEmv           PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethod = "contactless_emv"
	PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethodContactlessMagstripeMode PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethod = "contactless_magstripe_mode"
	PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethodMagneticStripeFallback   PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethod = "magnetic_stripe_fallback"
	PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethodMagneticStripeTrack2     PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethod = "magnetic_stripe_track2"
)

// The type of account being debited or credited
type PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountTypeChecking PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountType = "checking"
	PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountTypeCredit   PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountType = "credit"
	PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountTypePrepaid  PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountType = "prepaid"
	PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountTypeUnknown  PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountType = "unknown"
)

// The type of mobile wallet, one of `apple_pay`, `google_pay`, `samsung_pay`, or `unknown`.
type PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletTypeApplePay   PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletType = "apple_pay"
	PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletTypeGooglePay  PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletType = "google_pay"
	PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletTypeSamsungPay PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletType = "samsung_pay"
	PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletTypeUnknown    PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletType = "unknown"
)

// The blockchain network that the transaction was sent on.
type PaymentAttemptRecordPaymentMethodDetailsCryptoNetwork string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCryptoNetwork can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCryptoNetworkBase     PaymentAttemptRecordPaymentMethodDetailsCryptoNetwork = "base"
	PaymentAttemptRecordPaymentMethodDetailsCryptoNetworkEthereum PaymentAttemptRecordPaymentMethodDetailsCryptoNetwork = "ethereum"
	PaymentAttemptRecordPaymentMethodDetailsCryptoNetworkPolygon  PaymentAttemptRecordPaymentMethodDetailsCryptoNetwork = "polygon"
	PaymentAttemptRecordPaymentMethodDetailsCryptoNetworkSolana   PaymentAttemptRecordPaymentMethodDetailsCryptoNetwork = "solana"
)

// The token currency that the transaction was sent with.
type PaymentAttemptRecordPaymentMethodDetailsCryptoTokenCurrency string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCryptoTokenCurrency can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCryptoTokenCurrencyUsdc PaymentAttemptRecordPaymentMethodDetailsCryptoTokenCurrency = "usdc"
	PaymentAttemptRecordPaymentMethodDetailsCryptoTokenCurrencyUsdg PaymentAttemptRecordPaymentMethodDetailsCryptoTokenCurrency = "usdg"
	PaymentAttemptRecordPaymentMethodDetailsCryptoTokenCurrencyUsdp PaymentAttemptRecordPaymentMethodDetailsCryptoTokenCurrency = "usdp"
)

// The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `deutsche_bank_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.
type PaymentAttemptRecordPaymentMethodDetailsEPSBank string

// List of values that PaymentAttemptRecordPaymentMethodDetailsEPSBank can take
const (
	PaymentAttemptRecordPaymentMethodDetailsEPSBankArzteUndApothekerBank                PaymentAttemptRecordPaymentMethodDetailsEPSBank = "arzte_und_apotheker_bank"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankAustrianAnadiBankAg                  PaymentAttemptRecordPaymentMethodDetailsEPSBank = "austrian_anadi_bank_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankBankAustria                          PaymentAttemptRecordPaymentMethodDetailsEPSBank = "bank_austria"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankBankhausCarlSpangler                 PaymentAttemptRecordPaymentMethodDetailsEPSBank = "bankhaus_carl_spangler"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankBankhausSchelhammerUndSchatteraAg    PaymentAttemptRecordPaymentMethodDetailsEPSBank = "bankhaus_schelhammer_und_schattera_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankBawagPskAg                           PaymentAttemptRecordPaymentMethodDetailsEPSBank = "bawag_psk_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankBksBankAg                            PaymentAttemptRecordPaymentMethodDetailsEPSBank = "bks_bank_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankBrullKallmusBankAg                   PaymentAttemptRecordPaymentMethodDetailsEPSBank = "brull_kallmus_bank_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankBtvVierLanderBank                    PaymentAttemptRecordPaymentMethodDetailsEPSBank = "btv_vier_lander_bank"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankCapitalBankGraweGruppeAg             PaymentAttemptRecordPaymentMethodDetailsEPSBank = "capital_bank_grawe_gruppe_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankDeutscheBankAg                       PaymentAttemptRecordPaymentMethodDetailsEPSBank = "deutsche_bank_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankDolomitenbank                        PaymentAttemptRecordPaymentMethodDetailsEPSBank = "dolomitenbank"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankEasybankAg                           PaymentAttemptRecordPaymentMethodDetailsEPSBank = "easybank_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankErsteBankUndSparkassen               PaymentAttemptRecordPaymentMethodDetailsEPSBank = "erste_bank_und_sparkassen"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankHypoAlpeadriabankInternationalAg     PaymentAttemptRecordPaymentMethodDetailsEPSBank = "hypo_alpeadriabank_international_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankHypoBankBurgenlandAktiengesellschaft PaymentAttemptRecordPaymentMethodDetailsEPSBank = "hypo_bank_burgenland_aktiengesellschaft"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankHypoNoeLbFurNiederosterreichUWien    PaymentAttemptRecordPaymentMethodDetailsEPSBank = "hypo_noe_lb_fur_niederosterreich_u_wien"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankHypoOberosterreichSalzburgSteiermark PaymentAttemptRecordPaymentMethodDetailsEPSBank = "hypo_oberosterreich_salzburg_steiermark"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankHypoTirolBankAg                      PaymentAttemptRecordPaymentMethodDetailsEPSBank = "hypo_tirol_bank_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankHypoVorarlbergBankAg                 PaymentAttemptRecordPaymentMethodDetailsEPSBank = "hypo_vorarlberg_bank_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankMarchfelderBank                      PaymentAttemptRecordPaymentMethodDetailsEPSBank = "marchfelder_bank"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankOberbankAg                           PaymentAttemptRecordPaymentMethodDetailsEPSBank = "oberbank_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankRaiffeisenBankengruppeOsterreich     PaymentAttemptRecordPaymentMethodDetailsEPSBank = "raiffeisen_bankengruppe_osterreich"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankSchoellerbankAg                      PaymentAttemptRecordPaymentMethodDetailsEPSBank = "schoellerbank_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankSpardaBankWien                       PaymentAttemptRecordPaymentMethodDetailsEPSBank = "sparda_bank_wien"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankVolksbankGruppe                      PaymentAttemptRecordPaymentMethodDetailsEPSBank = "volksbank_gruppe"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankVolkskreditbankAg                    PaymentAttemptRecordPaymentMethodDetailsEPSBank = "volkskreditbank_ag"
	PaymentAttemptRecordPaymentMethodDetailsEPSBankVrBankBraunau                        PaymentAttemptRecordPaymentMethodDetailsEPSBank = "vr_bank_braunau"
)

// Account holder type, if provided. Can be one of `individual` or `company`.
type PaymentAttemptRecordPaymentMethodDetailsFPXAccountHolderType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsFPXAccountHolderType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsFPXAccountHolderTypeCompany    PaymentAttemptRecordPaymentMethodDetailsFPXAccountHolderType = "company"
	PaymentAttemptRecordPaymentMethodDetailsFPXAccountHolderTypeIndividual PaymentAttemptRecordPaymentMethodDetailsFPXAccountHolderType = "individual"
)

// The customer's bank. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, `pb_enterprise`, or `bank_of_china`.
type PaymentAttemptRecordPaymentMethodDetailsFPXBank string

// List of values that PaymentAttemptRecordPaymentMethodDetailsFPXBank can take
const (
	PaymentAttemptRecordPaymentMethodDetailsFPXBankAffinBank         PaymentAttemptRecordPaymentMethodDetailsFPXBank = "affin_bank"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankAgrobank          PaymentAttemptRecordPaymentMethodDetailsFPXBank = "agrobank"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankAllianceBank      PaymentAttemptRecordPaymentMethodDetailsFPXBank = "alliance_bank"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankAmbank            PaymentAttemptRecordPaymentMethodDetailsFPXBank = "ambank"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankBankIslam         PaymentAttemptRecordPaymentMethodDetailsFPXBank = "bank_islam"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankBankMuamalat      PaymentAttemptRecordPaymentMethodDetailsFPXBank = "bank_muamalat"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankBankOfChina       PaymentAttemptRecordPaymentMethodDetailsFPXBank = "bank_of_china"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankBankRakyat        PaymentAttemptRecordPaymentMethodDetailsFPXBank = "bank_rakyat"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankBsn               PaymentAttemptRecordPaymentMethodDetailsFPXBank = "bsn"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankCimb              PaymentAttemptRecordPaymentMethodDetailsFPXBank = "cimb"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankDeutscheBank      PaymentAttemptRecordPaymentMethodDetailsFPXBank = "deutsche_bank"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankHongLeongBank     PaymentAttemptRecordPaymentMethodDetailsFPXBank = "hong_leong_bank"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankHsbc              PaymentAttemptRecordPaymentMethodDetailsFPXBank = "hsbc"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankKfh               PaymentAttemptRecordPaymentMethodDetailsFPXBank = "kfh"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankMaybank2e         PaymentAttemptRecordPaymentMethodDetailsFPXBank = "maybank2e"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankMaybank2u         PaymentAttemptRecordPaymentMethodDetailsFPXBank = "maybank2u"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankOcbc              PaymentAttemptRecordPaymentMethodDetailsFPXBank = "ocbc"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankPbEnterprise      PaymentAttemptRecordPaymentMethodDetailsFPXBank = "pb_enterprise"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankPublicBank        PaymentAttemptRecordPaymentMethodDetailsFPXBank = "public_bank"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankRhb               PaymentAttemptRecordPaymentMethodDetailsFPXBank = "rhb"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankStandardChartered PaymentAttemptRecordPaymentMethodDetailsFPXBank = "standard_chartered"
	PaymentAttemptRecordPaymentMethodDetailsFPXBankUob               PaymentAttemptRecordPaymentMethodDetailsFPXBank = "uob"
)

// The customer's bank. Can be one of `abn_amro`, `adyen`, `asn_bank`, `bunq`, `buut`, `finom`, `handelsbanken`, `ing`, `knab`, `mollie`, `moneyou`, `n26`, `nn`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, `van_lanschot`, or `yoursafe`.
type PaymentAttemptRecordPaymentMethodDetailsIDEALBank string

// List of values that PaymentAttemptRecordPaymentMethodDetailsIDEALBank can take
const (
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankAbnAmro       PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "abn_amro"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankAdyen         PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "adyen"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankAsnBank       PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "asn_bank"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankBunq          PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "bunq"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankBuut          PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "buut"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankFinom         PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "finom"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankHandelsbanken PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "handelsbanken"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankIng           PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "ing"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankKnab          PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "knab"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankMollie        PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "mollie"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankMoneyou       PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "moneyou"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankN26           PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "n26"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankNn            PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "nn"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankRabobank      PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "rabobank"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankRegiobank     PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "regiobank"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankRevolut       PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "revolut"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankSnsBank       PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "sns_bank"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankTriodosBank   PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "triodos_bank"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankVanLanschot   PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "van_lanschot"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBankYoursafe      PaymentAttemptRecordPaymentMethodDetailsIDEALBank = "yoursafe"
)

// The Bank Identifier Code of the customer's bank.
type PaymentAttemptRecordPaymentMethodDetailsIDEALBIC string

// List of values that PaymentAttemptRecordPaymentMethodDetailsIDEALBIC can take
const (
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICABNANL2A PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "ABNANL2A"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICADYBNL2A PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "ADYBNL2A"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICASNBNL21 PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "ASNBNL21"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICBITSNL2A PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "BITSNL2A"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICBUNQNL2A PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "BUNQNL2A"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICBUUTNL2A PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "BUUTNL2A"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICFNOMNL22 PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "FNOMNL22"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICFVLBNL22 PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "FVLBNL22"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICHANDNL2A PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "HANDNL2A"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICINGBNL2A PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "INGBNL2A"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICKNABNL2H PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "KNABNL2H"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICMLLENL2A PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "MLLENL2A"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICMOYONL21 PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "MOYONL21"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICNNBANL2G PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "NNBANL2G"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICNTSBDEB1 PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "NTSBDEB1"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICRABONL2U PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "RABONL2U"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICRBRBNL21 PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "RBRBNL21"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICREVOIE23 PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "REVOIE23"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICREVOLT21 PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "REVOLT21"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICSNSBNL2A PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "SNSBNL2A"
	PaymentAttemptRecordPaymentMethodDetailsIDEALBICTRIONL2U PaymentAttemptRecordPaymentMethodDetailsIDEALBIC = "TRIONL2U"
)

// How card details were read in this transaction.
type PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethod string

// List of values that PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethod can take
const (
	PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethodContactEmv               PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethod = "contact_emv"
	PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethodContactlessEmv           PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethod = "contactless_emv"
	PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethodContactlessMagstripeMode PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethod = "contactless_magstripe_mode"
	PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethodMagneticStripeFallback   PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethod = "magnetic_stripe_fallback"
	PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethodMagneticStripeTrack2     PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethod = "magnetic_stripe_track2"
)

// The type of account being debited or credited
type PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceiptAccountType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceiptAccountType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceiptAccountTypeChecking PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceiptAccountType = "checking"
	PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceiptAccountTypeSavings  PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceiptAccountType = "savings"
	PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceiptAccountTypeUnknown  PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceiptAccountType = "unknown"
)

// The name of the convenience store chain where the payment was completed.
type PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChain string

// List of values that PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChain can take
const (
	PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChainFamilyMart PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChain = "familymart"
	PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChainLawson     PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChain = "lawson"
	PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChainMinistop   PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChain = "ministop"
	PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChainSeicomart  PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChain = "seicomart"
)

// The local credit or debit card brand.
type PaymentAttemptRecordPaymentMethodDetailsKrCardBrand string

// List of values that PaymentAttemptRecordPaymentMethodDetailsKrCardBrand can take
const (
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandBc          PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "bc"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandCiti        PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "citi"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandHana        PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "hana"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandHyundai     PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "hyundai"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandJeju        PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "jeju"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandJeonbuk     PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "jeonbuk"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandKakaobank   PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "kakaobank"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandKbank       PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "kbank"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandKdbbank     PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "kdbbank"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandKookmin     PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "kookmin"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandKwangju     PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "kwangju"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandLotte       PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "lotte"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandMg          PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "mg"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandNh          PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "nh"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandPost        PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "post"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandSamsung     PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "samsung"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandSavingsbank PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "savingsbank"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandShinhan     PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "shinhan"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandShinhyup    PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "shinhyup"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandSuhyup      PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "suhyup"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandTossbank    PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "tossbank"
	PaymentAttemptRecordPaymentMethodDetailsKrCardBrandWoori       PaymentAttemptRecordPaymentMethodDetailsKrCardBrand = "woori"
)

// The customer's bank. Can be one of `ing`, `citi_handlowy`, `tmobile_usbugi_bankowe`, `plus_bank`, `etransfer_pocztowy24`, `banki_spbdzielcze`, `bank_nowy_bfg_sa`, `getin_bank`, `velobank`, `blik`, `noble_pay`, `ideabank`, `envelobank`, `santander_przelew24`, `nest_przelew`, `mbank_mtransfer`, `inteligo`, `pbac_z_ipko`, `bnp_paribas`, `credit_agricole`, `toyota_bank`, `bank_pekao_sa`, `volkswagen_bank`, `bank_millennium`, `alior_bank`, or `boz`.
type PaymentAttemptRecordPaymentMethodDetailsP24Bank string

// List of values that PaymentAttemptRecordPaymentMethodDetailsP24Bank can take
const (
	PaymentAttemptRecordPaymentMethodDetailsP24BankAliorBank            PaymentAttemptRecordPaymentMethodDetailsP24Bank = "alior_bank"
	PaymentAttemptRecordPaymentMethodDetailsP24BankBankMillennium       PaymentAttemptRecordPaymentMethodDetailsP24Bank = "bank_millennium"
	PaymentAttemptRecordPaymentMethodDetailsP24BankBankNowyBfgSa        PaymentAttemptRecordPaymentMethodDetailsP24Bank = "bank_nowy_bfg_sa"
	PaymentAttemptRecordPaymentMethodDetailsP24BankBankPekaoSa          PaymentAttemptRecordPaymentMethodDetailsP24Bank = "bank_pekao_sa"
	PaymentAttemptRecordPaymentMethodDetailsP24BankBankiSpbdzielcze     PaymentAttemptRecordPaymentMethodDetailsP24Bank = "banki_spbdzielcze"
	PaymentAttemptRecordPaymentMethodDetailsP24BankBLIK                 PaymentAttemptRecordPaymentMethodDetailsP24Bank = "blik"
	PaymentAttemptRecordPaymentMethodDetailsP24BankBnpParibas           PaymentAttemptRecordPaymentMethodDetailsP24Bank = "bnp_paribas"
	PaymentAttemptRecordPaymentMethodDetailsP24BankBoz                  PaymentAttemptRecordPaymentMethodDetailsP24Bank = "boz"
	PaymentAttemptRecordPaymentMethodDetailsP24BankCitiHandlowy         PaymentAttemptRecordPaymentMethodDetailsP24Bank = "citi_handlowy"
	PaymentAttemptRecordPaymentMethodDetailsP24BankCreditAgricole       PaymentAttemptRecordPaymentMethodDetailsP24Bank = "credit_agricole"
	PaymentAttemptRecordPaymentMethodDetailsP24BankEnvelobank           PaymentAttemptRecordPaymentMethodDetailsP24Bank = "envelobank"
	PaymentAttemptRecordPaymentMethodDetailsP24BankEtransferPocztowy24  PaymentAttemptRecordPaymentMethodDetailsP24Bank = "etransfer_pocztowy24"
	PaymentAttemptRecordPaymentMethodDetailsP24BankGetinBank            PaymentAttemptRecordPaymentMethodDetailsP24Bank = "getin_bank"
	PaymentAttemptRecordPaymentMethodDetailsP24BankIdeabank             PaymentAttemptRecordPaymentMethodDetailsP24Bank = "ideabank"
	PaymentAttemptRecordPaymentMethodDetailsP24BankIng                  PaymentAttemptRecordPaymentMethodDetailsP24Bank = "ing"
	PaymentAttemptRecordPaymentMethodDetailsP24BankInteligo             PaymentAttemptRecordPaymentMethodDetailsP24Bank = "inteligo"
	PaymentAttemptRecordPaymentMethodDetailsP24BankMbankMtransfer       PaymentAttemptRecordPaymentMethodDetailsP24Bank = "mbank_mtransfer"
	PaymentAttemptRecordPaymentMethodDetailsP24BankNestPrzelew          PaymentAttemptRecordPaymentMethodDetailsP24Bank = "nest_przelew"
	PaymentAttemptRecordPaymentMethodDetailsP24BankNoblePay             PaymentAttemptRecordPaymentMethodDetailsP24Bank = "noble_pay"
	PaymentAttemptRecordPaymentMethodDetailsP24BankPbacZIpko            PaymentAttemptRecordPaymentMethodDetailsP24Bank = "pbac_z_ipko"
	PaymentAttemptRecordPaymentMethodDetailsP24BankPlusBank             PaymentAttemptRecordPaymentMethodDetailsP24Bank = "plus_bank"
	PaymentAttemptRecordPaymentMethodDetailsP24BankSantanderPrzelew24   PaymentAttemptRecordPaymentMethodDetailsP24Bank = "santander_przelew24"
	PaymentAttemptRecordPaymentMethodDetailsP24BankTmobileUsbugiBankowe PaymentAttemptRecordPaymentMethodDetailsP24Bank = "tmobile_usbugi_bankowe"
	PaymentAttemptRecordPaymentMethodDetailsP24BankToyotaBank           PaymentAttemptRecordPaymentMethodDetailsP24Bank = "toyota_bank"
	PaymentAttemptRecordPaymentMethodDetailsP24BankVelobank             PaymentAttemptRecordPaymentMethodDetailsP24Bank = "velobank"
	PaymentAttemptRecordPaymentMethodDetailsP24BankVolkswagenBank       PaymentAttemptRecordPaymentMethodDetailsP24Bank = "volkswagen_bank"
)

// An array of conditions that are covered for the transaction, if applicable.
type PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategory string

// List of values that PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategory can take
const (
	PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategoryFraudulent         PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategory = "fraudulent"
	PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategoryProductNotReceived PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategory = "product_not_received"
)

// Indicates whether the transaction is eligible for PayPal's seller protection.
type PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionStatus string

// List of values that PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionStatus can take
const (
	PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionStatusEligible          PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionStatus = "eligible"
	PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionStatusNotEligible       PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionStatus = "not_eligible"
	PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionStatusPartiallyEligible PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionStatus = "partially_eligible"
)

// funding type of the underlying payment method.
type PaymentAttemptRecordPaymentMethodDetailsRevolutPayFundingType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsRevolutPayFundingType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsRevolutPayFundingTypeCard PaymentAttemptRecordPaymentMethodDetailsRevolutPayFundingType = "card"
)

// Preferred language of the SOFORT authorization page that the customer is redirected to.
// Can be one of `de`, `en`, `es`, `fr`, `it`, `nl`, or `pl`
type PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguage string

// List of values that PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguage can take
const (
	PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguageDE PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguage = "de"
	PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguageEN PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguage = "en"
	PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguageES PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguage = "es"
	PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguageFR PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguage = "fr"
	PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguageIT PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguage = "it"
	PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguageNL PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguage = "nl"
	PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguagePL PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguage = "pl"
)

// The type of entity that holds the account. This can be either 'individual' or 'company'.
type PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountHolderType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountHolderType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountHolderTypeCompany    PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountHolderType = "company"
	PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountHolderTypeIndividual PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountHolderType = "individual"
)

// The type of the bank account. This can be either 'checking' or 'savings'.
type PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountTypeChecking PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountType = "checking"
	PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountTypeSavings  PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountType = "savings"
)

// The processor used for this payment attempt.
type PaymentAttemptRecordProcessorDetailsType string

// List of values that PaymentAttemptRecordProcessorDetailsType can take
const (
	PaymentAttemptRecordProcessorDetailsTypeCustom PaymentAttemptRecordProcessorDetailsType = "custom"
)

// Indicates who reported the payment.
type PaymentAttemptRecordReportedBy string

// List of values that PaymentAttemptRecordReportedBy can take
const (
	PaymentAttemptRecordReportedBySelf   PaymentAttemptRecordReportedBy = "self"
	PaymentAttemptRecordReportedByStripe PaymentAttemptRecordReportedBy = "stripe"
)

// List all the Payment Attempt Records attached to the specified Payment Record.
type PaymentAttemptRecordListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The ID of the Payment Record.
	PaymentRecord *string `form:"payment_record"`
}

// AddExpand appends a new field to expand.
func (p *PaymentAttemptRecordListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a Payment Attempt Record with the given ID
type PaymentAttemptRecordParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PaymentAttemptRecordParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a Payment Attempt Record with the given ID
type PaymentAttemptRecordRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PaymentAttemptRecordRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmount struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmountAuthorized struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmountCanceled struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmountFailed struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmountGuaranteed struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmountRefunded struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmountRequested struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY.
	Value int64 `json:"value"`
}

// Customer information for this payment.
type PaymentAttemptRecordCustomerDetails struct {
	// ID of the Stripe Customer associated with this payment.
	Customer string `json:"customer"`
	// The customer's email address.
	Email string `json:"email"`
	// The customer's name.
	Name string `json:"name"`
	// The customer's phone number.
	Phone string `json:"phone"`
}
type PaymentAttemptRecordPaymentMethodDetailsACHCreditTransfer struct {
	// Account number to transfer funds to.
	AccountNumber string `json:"account_number"`
	// Name of the bank associated with the routing number.
	BankName string `json:"bank_name"`
	// Routing transit number for the bank account to transfer funds to.
	RoutingNumber string `json:"routing_number"`
	// SWIFT code of the bank associated with the routing number.
	SwiftCode string `json:"swift_code"`
}
type PaymentAttemptRecordPaymentMethodDetailsACHDebit struct {
	// Type of entity that holds the account. This can be either `individual` or `company`.
	AccountHolderType PaymentAttemptRecordPaymentMethodDetailsACHDebitAccountHolderType `json:"account_holder_type"`
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
type PaymentAttemptRecordPaymentMethodDetailsACSSDebit struct {
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
type PaymentAttemptRecordPaymentMethodDetailsAffirm struct {
	// ID of the location that this reader is assigned to.
	Location string `json:"location"`
	// ID of the reader this transaction was made on.
	Reader string `json:"reader"`
	// The Affirm transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsAfterpayClearpay struct {
	// The Afterpay order ID associated with this payment intent.
	OrderID string `json:"order_id"`
	// Order identifier shown to the merchant in Afterpay's online portal.
	Reference string `json:"reference"`
}
type PaymentAttemptRecordPaymentMethodDetailsAlipay struct {
	// Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.
	BuyerID string `json:"buyer_id"`
	// Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Transaction ID of this particular Alipay transaction.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsAlmaInstallments struct {
	// The number of installments.
	Count int64 `json:"count"`
}
type PaymentAttemptRecordPaymentMethodDetailsAlma struct {
	Installments *PaymentAttemptRecordPaymentMethodDetailsAlmaInstallments `json:"installments"`
	// The Alma transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsAmazonPayFundingCard struct {
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
type PaymentAttemptRecordPaymentMethodDetailsAmazonPayFunding struct {
	Card *PaymentAttemptRecordPaymentMethodDetailsAmazonPayFundingCard `json:"card"`
	// funding type of the underlying payment method.
	Type PaymentAttemptRecordPaymentMethodDetailsAmazonPayFundingType `json:"type"`
}
type PaymentAttemptRecordPaymentMethodDetailsAmazonPay struct {
	Funding *PaymentAttemptRecordPaymentMethodDetailsAmazonPayFunding `json:"funding"`
	// The Amazon Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsAUBECSDebit struct {
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
type PaymentAttemptRecordPaymentMethodDetailsBACSDebit struct {
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
type PaymentAttemptRecordPaymentMethodDetailsBancontact struct {
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
	PreferredLanguage PaymentAttemptRecordPaymentMethodDetailsBancontactPreferredLanguage `json:"preferred_language"`
	// Owner's verified full name. Values are verified or provided by Bancontact directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type PaymentAttemptRecordPaymentMethodDetailsBillie struct {
	// The Billie transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}

// The billing details associated with the method of payment.
type PaymentAttemptRecordPaymentMethodDetailsBillingDetails struct {
	// A representation of a physical address.
	Address *Address `json:"address"`
	// The billing email associated with the method of payment.
	Email string `json:"email"`
	// The billing name associated with the method of payment.
	Name string `json:"name"`
	// The billing phone number associated with the method of payment.
	Phone string `json:"phone"`
}
type PaymentAttemptRecordPaymentMethodDetailsBLIK struct {
	// A unique and immutable identifier assigned by BLIK to every buyer.
	BuyerID string `json:"buyer_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsBoleto struct {
	// The tax ID of the customer (CPF for individuals consumers or CNPJ for businesses consumers)
	TaxID string `json:"tax_id"`
}

// Check results by Card networks on Card address and CVC at time of payment.
type PaymentAttemptRecordPaymentMethodDetailsCardChecks struct {
	// If you provide a value for `address.line1`, the check result is one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressLine1Check PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check `json:"address_line1_check"`
	// If you provide a address postal code, the check result is one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressPostalCodeCheck PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck `json:"address_postal_code_check"`
	// If you provide a CVC, the check results is one of `pass`, `fail`, `unavailable`, or `unchecked`.
	CVCCheck PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck `json:"cvc_check"`
}

// Installment plan selected for the payment.
type PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlan struct {
	// For `fixed_count` installment plans, this is the number of installment payments your customer will make to their credit card.
	Count int64 `json:"count"`
	// For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card. One of `month`.
	Interval PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanInterval `json:"interval"`
	// Type of installment plan, one of `fixed_count`, `revolving`, or `bonus`.
	Type PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlanType `json:"type"`
}

// Installment details for this payment.
type PaymentAttemptRecordPaymentMethodDetailsCardInstallments struct {
	// Installment plan selected for the payment.
	Plan *PaymentAttemptRecordPaymentMethodDetailsCardInstallmentsPlan `json:"plan"`
}

// If this card has network token credentials, this contains the details of the network token credentials.
type PaymentAttemptRecordPaymentMethodDetailsCardNetworkToken struct {
	// Indicates if Stripe used a network token, either user provided or Stripe managed when processing the transaction.
	Used bool `json:"used"`
}

// Populated if this transaction used 3D Secure authentication.
type PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure struct {
	// For authenticated transactions: Indicates how the issuing bank authenticated the customer.
	AuthenticationFlow PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	// Indicates the outcome of 3D Secure authentication.
	Result PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult `json:"result"`
	// Additional information about why 3D Secure succeeded or failed, based on the `result`.
	ResultReason PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason `json:"result_reason"`
	// The version of 3D Secure that was used.
	Version PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion `json:"version"`
}
type PaymentAttemptRecordPaymentMethodDetailsCardWalletApplePay struct {
	// Type of the apple_pay transaction, one of `apple_pay` or `apple_pay_later`.
	Type string `json:"type"`
}
type PaymentAttemptRecordPaymentMethodDetailsCardWalletGooglePay struct{}

// If this Card is part of a card wallet, this contains the details of the card wallet.
type PaymentAttemptRecordPaymentMethodDetailsCardWallet struct {
	ApplePay *PaymentAttemptRecordPaymentMethodDetailsCardWalletApplePay `json:"apple_pay"`
	// (For tokenized numbers only.) The last four digits of the device account number.
	DynamicLast4 string                                                       `json:"dynamic_last4"`
	GooglePay    *PaymentAttemptRecordPaymentMethodDetailsCardWalletGooglePay `json:"google_pay"`
	// The type of the card wallet, one of `apple_pay` or `google_pay`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
	Type string `json:"type"`
}

// Details of the card used for this payment attempt.
type PaymentAttemptRecordPaymentMethodDetailsCard struct {
	// The authorization code of the payment.
	AuthorizationCode string `json:"authorization_code"`
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand PaymentAttemptRecordPaymentMethodDetailsCardBrand `json:"brand"`
	// When using manual capture, a future timestamp at which the charge will be automatically refunded if uncaptured.
	CaptureBefore int64 `json:"capture_before"`
	// Check results by Card networks on Card address and CVC at time of payment.
	Checks *PaymentAttemptRecordPaymentMethodDetailsCardChecks `json:"checks"`
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
	Funding PaymentAttemptRecordPaymentMethodDetailsCardFunding `json:"funding"`
	// Issuer identification number of the card.
	IIN string `json:"iin"`
	// Installment details for this payment.
	Installments *PaymentAttemptRecordPaymentMethodDetailsCardInstallments `json:"installments"`
	// The name of the card's issuing bank.
	Issuer string `json:"issuer"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// True if this payment was marked as MOTO and out of scope for SCA.
	MOTO bool `json:"moto"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network PaymentAttemptRecordPaymentMethodDetailsCardNetwork `json:"network"`
	// Advice code from the card network for the failed payment.
	NetworkAdviceCode string `json:"network_advice_code"`
	// Decline code from the card network for the failed payment.
	NetworkDeclineCode string `json:"network_decline_code"`
	// If this card has network token credentials, this contains the details of the network token credentials.
	NetworkToken *PaymentAttemptRecordPaymentMethodDetailsCardNetworkToken `json:"network_token"`
	// This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.
	NetworkTransactionID string `json:"network_transaction_id"`
	// The transaction type that was passed for an off-session, Merchant-Initiated transaction, one of `recurring` or `unscheduled`.
	StoredCredentialUsage PaymentAttemptRecordPaymentMethodDetailsCardStoredCredentialUsage `json:"stored_credential_usage"`
	// Populated if this transaction used 3D Secure authentication.
	ThreeDSecure *PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
	// If this Card is part of a card wallet, this contains the details of the card wallet.
	Wallet *PaymentAttemptRecordPaymentMethodDetailsCardWallet `json:"wallet"`
}

// Details about payments collected offline.
type PaymentAttemptRecordPaymentMethodDetailsCardPresentOffline struct {
	// Time at which the payment was collected while offline
	StoredAt int64 `json:"stored_at"`
	// The method used to process this payment method offline. Only deferred is allowed.
	Type PaymentAttemptRecordPaymentMethodDetailsCardPresentOfflineType `json:"type"`
}

// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
type PaymentAttemptRecordPaymentMethodDetailsCardPresentReceipt struct {
	// The type of account being debited or credited
	AccountType PaymentAttemptRecordPaymentMethodDetailsCardPresentReceiptAccountType `json:"account_type"`
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
type PaymentAttemptRecordPaymentMethodDetailsCardPresentWallet struct {
	// The type of mobile wallet, one of `apple_pay`, `google_pay`, `samsung_pay`, or `unknown`.
	Type PaymentAttemptRecordPaymentMethodDetailsCardPresentWalletType `json:"type"`
}
type PaymentAttemptRecordPaymentMethodDetailsCardPresent struct {
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
	Offline *PaymentAttemptRecordPaymentMethodDetailsCardPresentOffline `json:"offline"`
	// Defines whether the authorized amount can be over-captured or not
	OvercaptureSupported bool `json:"overcapture_supported"`
	// The languages that the issuing bank recommends using for localizing any customer-facing text, as read from the card. Referenced from EMV tag 5F2D, data encoded on the card's chip.
	PreferredLocales []string `json:"preferred_locales"`
	// ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.
	Reader string `json:"reader"`
	// How card details were read in this transaction.
	ReadMethod PaymentAttemptRecordPaymentMethodDetailsCardPresentReadMethod `json:"read_method"`
	// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
	Receipt *PaymentAttemptRecordPaymentMethodDetailsCardPresentReceipt `json:"receipt"`
	Wallet  *PaymentAttemptRecordPaymentMethodDetailsCardPresentWallet  `json:"wallet"`
}
type PaymentAttemptRecordPaymentMethodDetailsCashApp struct {
	// A unique and immutable identifier assigned by Cash App to every buyer.
	BuyerID string `json:"buyer_id"`
	// A public identifier for buyers using Cash App.
	Cashtag string `json:"cashtag"`
	// A unique and immutable identifier of payments assigned by Cash App.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsCrypto struct {
	// The wallet address of the customer.
	BuyerAddress string `json:"buyer_address"`
	// The blockchain network that the transaction was sent on.
	Network PaymentAttemptRecordPaymentMethodDetailsCryptoNetwork `json:"network"`
	// The token currency that the transaction was sent with.
	TokenCurrency PaymentAttemptRecordPaymentMethodDetailsCryptoTokenCurrency `json:"token_currency"`
	// The blockchain transaction hash of the crypto payment.
	TransactionHash string `json:"transaction_hash"`
}

// Custom Payment Methods represent Payment Method types not modeled directly in
// the Stripe API. This resource consists of details about the custom payment method
// used for this payment attempt.
type PaymentAttemptRecordPaymentMethodDetailsCustom struct {
	// Display name for the custom (user-defined) payment method type used to make this payment.
	DisplayName string `json:"display_name"`
	// The custom payment method type associated with this payment.
	Type string `json:"type"`
}
type PaymentAttemptRecordPaymentMethodDetailsCustomerBalance struct{}
type PaymentAttemptRecordPaymentMethodDetailsEPS struct {
	// The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `deutsche_bank_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.
	Bank PaymentAttemptRecordPaymentMethodDetailsEPSBank `json:"bank"`
	// Owner's verified full name. Values are verified or provided by EPS directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	// EPS rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type PaymentAttemptRecordPaymentMethodDetailsFPX struct {
	// Account holder type, if provided. Can be one of `individual` or `company`.
	AccountHolderType PaymentAttemptRecordPaymentMethodDetailsFPXAccountHolderType `json:"account_holder_type"`
	// The customer's bank. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, `pb_enterprise`, or `bank_of_china`.
	Bank PaymentAttemptRecordPaymentMethodDetailsFPXBank `json:"bank"`
	// Unique transaction id generated by FPX for every request from the merchant
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsGiropay struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	BIC string `json:"bic"`
	// Owner's verified full name. Values are verified or provided by Giropay directly (if supported) at the time of authorization or settlement. They cannot be set or mutated. Giropay rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type PaymentAttemptRecordPaymentMethodDetailsGrabpay struct {
	// Unique transaction id generated by GrabPay
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsIDEAL struct {
	// The customer's bank. Can be one of `abn_amro`, `adyen`, `asn_bank`, `bunq`, `buut`, `finom`, `handelsbanken`, `ing`, `knab`, `mollie`, `moneyou`, `n26`, `nn`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, `van_lanschot`, or `yoursafe`.
	Bank PaymentAttemptRecordPaymentMethodDetailsIDEALBank `json:"bank"`
	// The Bank Identifier Code of the customer's bank.
	BIC PaymentAttemptRecordPaymentMethodDetailsIDEALBIC `json:"bic"`
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
type PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceipt struct {
	// The type of account being debited or credited
	AccountType PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceiptAccountType `json:"account_type"`
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
type PaymentAttemptRecordPaymentMethodDetailsInteracPresent struct {
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
	ReadMethod PaymentAttemptRecordPaymentMethodDetailsInteracPresentReadMethod `json:"read_method"`
	// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
	Receipt *PaymentAttemptRecordPaymentMethodDetailsInteracPresentReceipt `json:"receipt"`
}
type PaymentAttemptRecordPaymentMethodDetailsKakaoPay struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Kakao Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}

// The payer's address
type PaymentAttemptRecordPaymentMethodDetailsKlarnaPayerDetailsAddress struct {
	// The payer address country
	Country string `json:"country"`
}

// The payer details for this transaction.
type PaymentAttemptRecordPaymentMethodDetailsKlarnaPayerDetails struct {
	// The payer's address
	Address *PaymentAttemptRecordPaymentMethodDetailsKlarnaPayerDetailsAddress `json:"address"`
}
type PaymentAttemptRecordPaymentMethodDetailsKlarna struct {
	// The payer details for this transaction.
	PayerDetails *PaymentAttemptRecordPaymentMethodDetailsKlarnaPayerDetails `json:"payer_details"`
	// The Klarna payment method used for this transaction.
	// Can be one of `pay_later`, `pay_now`, `pay_with_financing`, or `pay_in_installments`
	PaymentMethodCategory string `json:"payment_method_category"`
	// Preferred language of the Klarna authorization page that the customer is redirected to.
	// Can be one of `de-AT`, `en-AT`, `nl-BE`, `fr-BE`, `en-BE`, `de-DE`, `en-DE`, `da-DK`, `en-DK`, `es-ES`, `en-ES`, `fi-FI`, `sv-FI`, `en-FI`, `en-GB`, `en-IE`, `it-IT`, `en-IT`, `nl-NL`, `en-NL`, `nb-NO`, `en-NO`, `sv-SE`, `en-SE`, `en-US`, `es-US`, `fr-FR`, `en-FR`, `cs-CZ`, `en-CZ`, `ro-RO`, `en-RO`, `el-GR`, `en-GR`, `en-AU`, `en-NZ`, `en-CA`, `fr-CA`, `pl-PL`, `en-PL`, `pt-PT`, `en-PT`, `de-CH`, `fr-CH`, `it-CH`, or `en-CH`
	PreferredLocale string `json:"preferred_locale"`
}

// If the payment succeeded, this contains the details of the convenience store where the payment was completed.
type PaymentAttemptRecordPaymentMethodDetailsKonbiniStore struct {
	// The name of the convenience store chain where the payment was completed.
	Chain PaymentAttemptRecordPaymentMethodDetailsKonbiniStoreChain `json:"chain"`
}
type PaymentAttemptRecordPaymentMethodDetailsKonbini struct {
	// If the payment succeeded, this contains the details of the convenience store where the payment was completed.
	Store *PaymentAttemptRecordPaymentMethodDetailsKonbiniStore `json:"store"`
}
type PaymentAttemptRecordPaymentMethodDetailsKrCard struct {
	// The local credit or debit card brand.
	Brand PaymentAttemptRecordPaymentMethodDetailsKrCardBrand `json:"brand"`
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The last four digits of the card. This may not be present for American Express cards.
	Last4 string `json:"last4"`
	// The Korean Card transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsLink struct {
	// Two-letter ISO code representing the funding source country beneath the Link payment.
	// You could use this attribute to get a sense of international fees.
	Country string `json:"country"`
}
type PaymentAttemptRecordPaymentMethodDetailsMbWay struct{}

// Internal card details
type PaymentAttemptRecordPaymentMethodDetailsMobilepayCard struct {
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
type PaymentAttemptRecordPaymentMethodDetailsMobilepay struct {
	// Internal card details
	Card *PaymentAttemptRecordPaymentMethodDetailsMobilepayCard `json:"card"`
}
type PaymentAttemptRecordPaymentMethodDetailsMultibanco struct {
	// Entity number associated with this Multibanco payment.
	Entity string `json:"entity"`
	// Reference number associated with this Multibanco payment.
	Reference string `json:"reference"`
}
type PaymentAttemptRecordPaymentMethodDetailsNaverPay struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Naver Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsNzBankAccount struct {
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
type PaymentAttemptRecordPaymentMethodDetailsOXXO struct {
	// OXXO reference number
	Number string `json:"number"`
}
type PaymentAttemptRecordPaymentMethodDetailsP24 struct {
	// The customer's bank. Can be one of `ing`, `citi_handlowy`, `tmobile_usbugi_bankowe`, `plus_bank`, `etransfer_pocztowy24`, `banki_spbdzielcze`, `bank_nowy_bfg_sa`, `getin_bank`, `velobank`, `blik`, `noble_pay`, `ideabank`, `envelobank`, `santander_przelew24`, `nest_przelew`, `mbank_mtransfer`, `inteligo`, `pbac_z_ipko`, `bnp_paribas`, `credit_agricole`, `toyota_bank`, `bank_pekao_sa`, `volkswagen_bank`, `bank_millennium`, `alior_bank`, or `boz`.
	Bank PaymentAttemptRecordPaymentMethodDetailsP24Bank `json:"bank"`
	// Unique reference for this Przelewy24 payment.
	Reference string `json:"reference"`
	// Owner's verified full name. Values are verified or provided by Przelewy24 directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	// Przelewy24 rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type PaymentAttemptRecordPaymentMethodDetailsPayByBank struct{}
type PaymentAttemptRecordPaymentMethodDetailsPayco struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Payco transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsPayNow struct {
	// ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.
	Location string `json:"location"`
	// ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.
	Reader string `json:"reader"`
	// Reference number associated with this PayNow payment
	Reference string `json:"reference"`
}

// The level of protection offered as defined by PayPal Seller Protection for Merchants, for this transaction.
type PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtection struct {
	// An array of conditions that are covered for the transaction, if applicable.
	DisputeCategories []PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionDisputeCategory `json:"dispute_categories"`
	// Indicates whether the transaction is eligible for PayPal's seller protection.
	Status PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtectionStatus `json:"status"`
}
type PaymentAttemptRecordPaymentMethodDetailsPaypal struct {
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
	SellerProtection *PaymentAttemptRecordPaymentMethodDetailsPaypalSellerProtection `json:"seller_protection"`
	// A unique ID generated by PayPal for this transaction.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsPayto struct {
	// Bank-State-Branch number of the bank account.
	BSBNumber string `json:"bsb_number"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate string `json:"mandate"`
	// The PayID alias for the bank account.
	PayID string `json:"pay_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsPix struct {
	// Unique transaction id generated by BCB
	BankTransactionID string `json:"bank_transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsPromptPay struct {
	// Bill reference generated by PromptPay
	Reference string `json:"reference"`
}
type PaymentAttemptRecordPaymentMethodDetailsRevolutPayFundingCard struct {
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
type PaymentAttemptRecordPaymentMethodDetailsRevolutPayFunding struct {
	Card *PaymentAttemptRecordPaymentMethodDetailsRevolutPayFundingCard `json:"card"`
	// funding type of the underlying payment method.
	Type PaymentAttemptRecordPaymentMethodDetailsRevolutPayFundingType `json:"type"`
}
type PaymentAttemptRecordPaymentMethodDetailsRevolutPay struct {
	Funding *PaymentAttemptRecordPaymentMethodDetailsRevolutPayFunding `json:"funding"`
	// The Revolut Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsSamsungPay struct {
	// A unique identifier for the buyer as determined by the local payment processor.
	BuyerID string `json:"buyer_id"`
	// The Samsung Pay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsSatispay struct {
	// The Satispay transaction ID associated with this payment.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsSEPACreditTransfer struct {
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	BIC string `json:"bic"`
	// IBAN of the bank account to transfer funds to.
	IBAN string `json:"iban"`
}
type PaymentAttemptRecordPaymentMethodDetailsSEPADebit struct {
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
type PaymentAttemptRecordPaymentMethodDetailsSofort struct {
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
	PreferredLanguage PaymentAttemptRecordPaymentMethodDetailsSofortPreferredLanguage `json:"preferred_language"`
	// Owner's verified full name. Values are verified or provided by SOFORT directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type PaymentAttemptRecordPaymentMethodDetailsStripeAccount struct{}
type PaymentAttemptRecordPaymentMethodDetailsSwish struct {
	// Uniquely identifies the payer's Swish account. You can use this attribute to check whether two Swish transactions were paid for by the same payer
	Fingerprint string `json:"fingerprint"`
	// Payer bank reference number for the payment
	PaymentReference string `json:"payment_reference"`
	// The last four digits of the Swish account phone number
	VerifiedPhoneLast4 string `json:"verified_phone_last4"`
}
type PaymentAttemptRecordPaymentMethodDetailsTWINT struct{}
type PaymentAttemptRecordPaymentMethodDetailsUSBankAccount struct {
	// The type of entity that holds the account. This can be either 'individual' or 'company'.
	AccountHolderType PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountHolderType `json:"account_holder_type"`
	// The type of the bank account. This can be either 'checking' or 'savings'.
	AccountType PaymentAttemptRecordPaymentMethodDetailsUSBankAccountAccountType `json:"account_type"`
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
type PaymentAttemptRecordPaymentMethodDetailsWeChat struct{}
type PaymentAttemptRecordPaymentMethodDetailsWeChatPay struct {
	// Uniquely identifies this particular WeChat Pay account. You can use this attribute to check whether two WeChat accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.
	Location string `json:"location"`
	// ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.
	Reader string `json:"reader"`
	// Transaction ID of this particular WeChat Pay transaction.
	TransactionID string `json:"transaction_id"`
}
type PaymentAttemptRecordPaymentMethodDetailsZip struct{}

// Information about the Payment Method debited for this payment.
type PaymentAttemptRecordPaymentMethodDetails struct {
	ACHCreditTransfer *PaymentAttemptRecordPaymentMethodDetailsACHCreditTransfer `json:"ach_credit_transfer"`
	ACHDebit          *PaymentAttemptRecordPaymentMethodDetailsACHDebit          `json:"ach_debit"`
	ACSSDebit         *PaymentAttemptRecordPaymentMethodDetailsACSSDebit         `json:"acss_debit"`
	Affirm            *PaymentAttemptRecordPaymentMethodDetailsAffirm            `json:"affirm"`
	AfterpayClearpay  *PaymentAttemptRecordPaymentMethodDetailsAfterpayClearpay  `json:"afterpay_clearpay"`
	Alipay            *PaymentAttemptRecordPaymentMethodDetailsAlipay            `json:"alipay"`
	Alma              *PaymentAttemptRecordPaymentMethodDetailsAlma              `json:"alma"`
	AmazonPay         *PaymentAttemptRecordPaymentMethodDetailsAmazonPay         `json:"amazon_pay"`
	AUBECSDebit       *PaymentAttemptRecordPaymentMethodDetailsAUBECSDebit       `json:"au_becs_debit"`
	BACSDebit         *PaymentAttemptRecordPaymentMethodDetailsBACSDebit         `json:"bacs_debit"`
	Bancontact        *PaymentAttemptRecordPaymentMethodDetailsBancontact        `json:"bancontact"`
	Billie            *PaymentAttemptRecordPaymentMethodDetailsBillie            `json:"billie"`
	// The billing details associated with the method of payment.
	BillingDetails *PaymentAttemptRecordPaymentMethodDetailsBillingDetails `json:"billing_details"`
	BLIK           *PaymentAttemptRecordPaymentMethodDetailsBLIK           `json:"blik"`
	Boleto         *PaymentAttemptRecordPaymentMethodDetailsBoleto         `json:"boleto"`
	// Details of the card used for this payment attempt.
	Card        *PaymentAttemptRecordPaymentMethodDetailsCard        `json:"card"`
	CardPresent *PaymentAttemptRecordPaymentMethodDetailsCardPresent `json:"card_present"`
	CashApp     *PaymentAttemptRecordPaymentMethodDetailsCashApp     `json:"cashapp"`
	Crypto      *PaymentAttemptRecordPaymentMethodDetailsCrypto      `json:"crypto"`
	// Custom Payment Methods represent Payment Method types not modeled directly in
	// the Stripe API. This resource consists of details about the custom payment method
	// used for this payment attempt.
	Custom          *PaymentAttemptRecordPaymentMethodDetailsCustom          `json:"custom"`
	CustomerBalance *PaymentAttemptRecordPaymentMethodDetailsCustomerBalance `json:"customer_balance"`
	EPS             *PaymentAttemptRecordPaymentMethodDetailsEPS             `json:"eps"`
	FPX             *PaymentAttemptRecordPaymentMethodDetailsFPX             `json:"fpx"`
	Giropay         *PaymentAttemptRecordPaymentMethodDetailsGiropay         `json:"giropay"`
	Grabpay         *PaymentAttemptRecordPaymentMethodDetailsGrabpay         `json:"grabpay"`
	IDEAL           *PaymentAttemptRecordPaymentMethodDetailsIDEAL           `json:"ideal"`
	InteracPresent  *PaymentAttemptRecordPaymentMethodDetailsInteracPresent  `json:"interac_present"`
	KakaoPay        *PaymentAttemptRecordPaymentMethodDetailsKakaoPay        `json:"kakao_pay"`
	Klarna          *PaymentAttemptRecordPaymentMethodDetailsKlarna          `json:"klarna"`
	Konbini         *PaymentAttemptRecordPaymentMethodDetailsKonbini         `json:"konbini"`
	KrCard          *PaymentAttemptRecordPaymentMethodDetailsKrCard          `json:"kr_card"`
	Link            *PaymentAttemptRecordPaymentMethodDetailsLink            `json:"link"`
	MbWay           *PaymentAttemptRecordPaymentMethodDetailsMbWay           `json:"mb_way"`
	Mobilepay       *PaymentAttemptRecordPaymentMethodDetailsMobilepay       `json:"mobilepay"`
	Multibanco      *PaymentAttemptRecordPaymentMethodDetailsMultibanco      `json:"multibanco"`
	NaverPay        *PaymentAttemptRecordPaymentMethodDetailsNaverPay        `json:"naver_pay"`
	NzBankAccount   *PaymentAttemptRecordPaymentMethodDetailsNzBankAccount   `json:"nz_bank_account"`
	OXXO            *PaymentAttemptRecordPaymentMethodDetailsOXXO            `json:"oxxo"`
	P24             *PaymentAttemptRecordPaymentMethodDetailsP24             `json:"p24"`
	PayByBank       *PaymentAttemptRecordPaymentMethodDetailsPayByBank       `json:"pay_by_bank"`
	Payco           *PaymentAttemptRecordPaymentMethodDetailsPayco           `json:"payco"`
	// ID of the Stripe PaymentMethod used to make this payment.
	PaymentMethod      string                                                      `json:"payment_method"`
	PayNow             *PaymentAttemptRecordPaymentMethodDetailsPayNow             `json:"paynow"`
	Paypal             *PaymentAttemptRecordPaymentMethodDetailsPaypal             `json:"paypal"`
	Payto              *PaymentAttemptRecordPaymentMethodDetailsPayto              `json:"payto"`
	Pix                *PaymentAttemptRecordPaymentMethodDetailsPix                `json:"pix"`
	PromptPay          *PaymentAttemptRecordPaymentMethodDetailsPromptPay          `json:"promptpay"`
	RevolutPay         *PaymentAttemptRecordPaymentMethodDetailsRevolutPay         `json:"revolut_pay"`
	SamsungPay         *PaymentAttemptRecordPaymentMethodDetailsSamsungPay         `json:"samsung_pay"`
	Satispay           *PaymentAttemptRecordPaymentMethodDetailsSatispay           `json:"satispay"`
	SEPACreditTransfer *PaymentAttemptRecordPaymentMethodDetailsSEPACreditTransfer `json:"sepa_credit_transfer"`
	SEPADebit          *PaymentAttemptRecordPaymentMethodDetailsSEPADebit          `json:"sepa_debit"`
	Sofort             *PaymentAttemptRecordPaymentMethodDetailsSofort             `json:"sofort"`
	StripeAccount      *PaymentAttemptRecordPaymentMethodDetailsStripeAccount      `json:"stripe_account"`
	Swish              *PaymentAttemptRecordPaymentMethodDetailsSwish              `json:"swish"`
	TWINT              *PaymentAttemptRecordPaymentMethodDetailsTWINT              `json:"twint"`
	// The type of transaction-specific details of the payment method used in the payment. See [PaymentMethod.type](https://docs.stripe.com/api/payment_methods/object#payment_method_object-type) for the full list of possible types.
	// An additional hash is included on `payment_method_details` with a name matching this value.
	// It contains information specific to the payment method.
	Type          string                                                 `json:"type"`
	USBankAccount *PaymentAttemptRecordPaymentMethodDetailsUSBankAccount `json:"us_bank_account"`
	WeChat        *PaymentAttemptRecordPaymentMethodDetailsWeChat        `json:"wechat"`
	WeChatPay     *PaymentAttemptRecordPaymentMethodDetailsWeChatPay     `json:"wechat_pay"`
	Zip           *PaymentAttemptRecordPaymentMethodDetailsZip           `json:"zip"`
}

// Custom processors represent payment processors not modeled directly in
// the Stripe API. This resource consists of details about the custom processor
// used for this payment attempt.
type PaymentAttemptRecordProcessorDetailsCustom struct {
	// An opaque string for manual reconciliation of this payment, for example a check number or a payment processor ID.
	PaymentReference string `json:"payment_reference"`
}

// Processor information associated with this payment.
type PaymentAttemptRecordProcessorDetails struct {
	// Custom processors represent payment processors not modeled directly in
	// the Stripe API. This resource consists of details about the custom processor
	// used for this payment attempt.
	Custom *PaymentAttemptRecordProcessorDetailsCustom `json:"custom"`
	// The processor used for this payment attempt.
	Type PaymentAttemptRecordProcessorDetailsType `json:"type"`
}

// Shipping information for this payment.
type PaymentAttemptRecordShippingDetails struct {
	// A representation of a physical address.
	Address *Address `json:"address"`
	// The shipping recipient's name.
	Name string `json:"name"`
	// The shipping recipient's phone number.
	Phone string `json:"phone"`
}

// A Payment Attempt Record represents an individual attempt at making a payment, on or off Stripe.
// Each payment attempt tries to collect a fixed amount of money from a fixed customer and payment
// method. Payment Attempt Records are attached to Payment Records. Only one attempt per Payment Record
// can have guaranteed funds.
type PaymentAttemptRecord struct {
	APIResource
	// A representation of an amount of money, consisting of an amount and a currency.
	Amount *PaymentAttemptRecordAmount `json:"amount"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountAuthorized *PaymentAttemptRecordAmountAuthorized `json:"amount_authorized"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountCanceled *PaymentAttemptRecordAmountCanceled `json:"amount_canceled"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountFailed *PaymentAttemptRecordAmountFailed `json:"amount_failed"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountGuaranteed *PaymentAttemptRecordAmountGuaranteed `json:"amount_guaranteed"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountRefunded *PaymentAttemptRecordAmountRefunded `json:"amount_refunded"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountRequested *PaymentAttemptRecordAmountRequested `json:"amount_requested"`
	// ID of the Connect application that created the PaymentAttemptRecord.
	Application string `json:"application"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Customer information for this payment.
	CustomerDetails *PaymentAttemptRecordCustomerDetails `json:"customer_details"`
	// Indicates whether the customer was present in your checkout flow during this payment.
	CustomerPresence PaymentAttemptRecordCustomerPresence `json:"customer_presence"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Information about the Payment Method debited for this payment.
	PaymentMethodDetails *PaymentAttemptRecordPaymentMethodDetails `json:"payment_method_details"`
	// ID of the Payment Record this Payment Attempt Record belongs to.
	PaymentRecord string `json:"payment_record"`
	// Processor information associated with this payment.
	ProcessorDetails *PaymentAttemptRecordProcessorDetails `json:"processor_details"`
	// Indicates who reported the payment.
	ReportedBy PaymentAttemptRecordReportedBy `json:"reported_by"`
	// Shipping information for this payment.
	ShippingDetails *PaymentAttemptRecordShippingDetails `json:"shipping_details"`
}

// PaymentAttemptRecordList is a list of PaymentAttemptRecords as retrieved from a list endpoint.
type PaymentAttemptRecordList struct {
	APIResource
	ListMeta
	Data []*PaymentAttemptRecord `json:"data"`
}
