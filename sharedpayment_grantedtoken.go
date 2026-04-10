//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The reason why the SharedPaymentGrantedToken has been deactivated.
type SharedPaymentGrantedTokenDeactivatedReason string

// List of values that SharedPaymentGrantedTokenDeactivatedReason can take
const (
	SharedPaymentGrantedTokenDeactivatedReasonConsumed SharedPaymentGrantedTokenDeactivatedReason = "consumed"
	SharedPaymentGrantedTokenDeactivatedReasonExpired  SharedPaymentGrantedTokenDeactivatedReason = "expired"
	SharedPaymentGrantedTokenDeactivatedReasonResolved SharedPaymentGrantedTokenDeactivatedReason = "resolved"
	SharedPaymentGrantedTokenDeactivatedReasonRevoked  SharedPaymentGrantedTokenDeactivatedReason = "revoked"
)

// The recurring interval at which the shared payment token's amount usage restrictions reset.
type SharedPaymentGrantedTokenUsageLimitsRecurringInterval string

// List of values that SharedPaymentGrantedTokenUsageLimitsRecurringInterval can take
const (
	SharedPaymentGrantedTokenUsageLimitsRecurringIntervalMonth SharedPaymentGrantedTokenUsageLimitsRecurringInterval = "month"
	SharedPaymentGrantedTokenUsageLimitsRecurringIntervalWeek  SharedPaymentGrantedTokenUsageLimitsRecurringInterval = "week"
	SharedPaymentGrantedTokenUsageLimitsRecurringIntervalYear  SharedPaymentGrantedTokenUsageLimitsRecurringInterval = "year"
)

// The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, `visa_checkout`, or `link`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeAmexExpressCheckout SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "amex_express_checkout"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeApplePay            SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "apple_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeGooglePay           SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "google_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeLink                SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "link"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeMasterpass          SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "masterpass"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeSamsungPay          SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "samsung_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeVisaCheckout        SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "visa_checkout"
)

// The method used to process this payment method offline. Only deferred is allowed.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentOfflineType string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentOfflineType can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentOfflineTypeDeferred SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentOfflineType = "deferred"
)

// How card details were read in this transaction.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethod string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethod can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethodContactEmv               SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethod = "contact_emv"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethodContactlessEmv           SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethod = "contactless_emv"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethodContactlessMagstripeMode SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethod = "contactless_magstripe_mode"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethodMagneticStripeFallback   SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethod = "magnetic_stripe_fallback"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethodMagneticStripeTrack2     SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethod = "magnetic_stripe_track2"
)

// The type of mobile wallet, one of `apple_pay`, `google_pay`, `samsung_pay`, or `unknown`.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletType string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletType can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletTypeApplePay   SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletType = "apple_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletTypeGooglePay  SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletType = "google_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletTypeSamsungPay SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletType = "samsung_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletTypeUnknown    SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletType = "unknown"
)

// The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `deutsche_bank_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.
type SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankArzteUndApothekerBank                SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "arzte_und_apotheker_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankAustrianAnadiBankAg                  SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "austrian_anadi_bank_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankBankAustria                          SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "bank_austria"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankBankhausCarlSpangler                 SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "bankhaus_carl_spangler"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankBankhausSchelhammerUndSchatteraAg    SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "bankhaus_schelhammer_und_schattera_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankBawagPskAg                           SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "bawag_psk_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankBksBankAg                            SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "bks_bank_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankBrullKallmusBankAg                   SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "brull_kallmus_bank_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankBtvVierLanderBank                    SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "btv_vier_lander_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankCapitalBankGraweGruppeAg             SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "capital_bank_grawe_gruppe_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankDeutscheBankAg                       SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "deutsche_bank_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankDolomitenbank                        SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "dolomitenbank"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankEasybankAg                           SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "easybank_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankErsteBankUndSparkassen               SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "erste_bank_und_sparkassen"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankHypoAlpeadriabankInternationalAg     SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "hypo_alpeadriabank_international_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankHypoBankBurgenlandAktiengesellschaft SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "hypo_bank_burgenland_aktiengesellschaft"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankHypoNoeLbFurNiederosterreichUWien    SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "hypo_noe_lb_fur_niederosterreich_u_wien"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankHypoOberosterreichSalzburgSteiermark SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "hypo_oberosterreich_salzburg_steiermark"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankHypoTirolBankAg                      SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "hypo_tirol_bank_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankHypoVorarlbergBankAg                 SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "hypo_vorarlberg_bank_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankMarchfelderBank                      SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "marchfelder_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankOberbankAg                           SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "oberbank_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankRaiffeisenBankengruppeOsterreich     SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "raiffeisen_bankengruppe_osterreich"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankSchoellerbankAg                      SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "schoellerbank_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankSpardaBankWien                       SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "sparda_bank_wien"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankVolksbankGruppe                      SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "volksbank_gruppe"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankVolkskreditbankAg                    SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "volkskreditbank_ag"
	SharedPaymentGrantedTokenPaymentMethodDetailsEPSBankVrBankBraunau                        SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank = "vr_bank_braunau"
)

// Account holder type, if provided. Can be one of `individual` or `company`.
type SharedPaymentGrantedTokenPaymentMethodDetailsFPXAccountHolderType string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsFPXAccountHolderType can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXAccountHolderTypeCompany    SharedPaymentGrantedTokenPaymentMethodDetailsFPXAccountHolderType = "company"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXAccountHolderTypeIndividual SharedPaymentGrantedTokenPaymentMethodDetailsFPXAccountHolderType = "individual"
)

// The customer's bank, if provided. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, `pb_enterprise`, or `bank_of_china`.
type SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankAffinBank         SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "affin_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankAgrobank          SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "agrobank"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankAllianceBank      SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "alliance_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankAmbank            SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "ambank"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankBankIslam         SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "bank_islam"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankBankMuamalat      SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "bank_muamalat"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankBankOfChina       SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "bank_of_china"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankBankRakyat        SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "bank_rakyat"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankBsn               SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "bsn"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankCimb              SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "cimb"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankDeutscheBank      SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "deutsche_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankHongLeongBank     SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "hong_leong_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankHsbc              SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "hsbc"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankKfh               SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "kfh"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankMaybank2e         SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "maybank2e"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankMaybank2u         SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "maybank2u"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankOcbc              SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "ocbc"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankPbEnterprise      SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "pb_enterprise"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankPublicBank        SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "public_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankRhb               SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "rhb"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankStandardChartered SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "standard_chartered"
	SharedPaymentGrantedTokenPaymentMethodDetailsFPXBankUob               SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank = "uob"
)

type SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBank string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBank can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBankBca     SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBank = "bca"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBankBni     SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBank = "bni"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBankBri     SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBank = "bri"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBankCimb    SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBank = "cimb"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBankPermata SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBank = "permata"
)

// The customer's bank, if provided. Can be one of `abn_amro`, `adyen`, `asn_bank`, `bunq`, `buut`, `finom`, `handelsbanken`, `ing`, `knab`, `mollie`, `moneyou`, `n26`, `nn`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, `van_lanschot`, or `yoursafe`.
type SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankAbnAmro       SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "abn_amro"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankAdyen         SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "adyen"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankAsnBank       SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "asn_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankBunq          SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "bunq"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankBuut          SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "buut"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankFinom         SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "finom"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankHandelsbanken SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "handelsbanken"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankIng           SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "ing"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankKnab          SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "knab"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankMollie        SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "mollie"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankMoneyou       SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "moneyou"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankN26           SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "n26"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankNn            SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "nn"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankRabobank      SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "rabobank"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankRegiobank     SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "regiobank"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankRevolut       SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "revolut"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankSnsBank       SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "sns_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankTriodosBank   SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "triodos_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankVanLanschot   SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "van_lanschot"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBankYoursafe      SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank = "yoursafe"
)

// The Bank Identifier Code of the customer's bank, if the bank was provided.
type SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICABNANL2A SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "ABNANL2A"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICADYBNL2A SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "ADYBNL2A"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICASNBNL21 SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "ASNBNL21"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICBITSNL2A SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "BITSNL2A"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICBUNQNL2A SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "BUNQNL2A"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICBUUTNL2A SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "BUUTNL2A"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICFNOMNL22 SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "FNOMNL22"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICFVLBNL22 SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "FVLBNL22"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICHANDNL2A SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "HANDNL2A"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICINGBNL2A SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "INGBNL2A"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICKNABNL2H SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "KNABNL2H"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICMLLENL2A SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "MLLENL2A"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICMOYONL21 SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "MOYONL21"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICNNBANL2G SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "NNBANL2G"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICNTSBDEB1 SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "NTSBDEB1"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICRABONL2U SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "RABONL2U"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICRBRBNL21 SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "RBRBNL21"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICREVOIE23 SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "REVOIE23"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICREVOLT21 SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "REVOLT21"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICSNSBNL2A SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "SNSBNL2A"
	SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBICTRIONL2U SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC = "TRIONL2U"
)

// How card details were read in this transaction.
type SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethod string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethod can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethodContactEmv               SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethod = "contact_emv"
	SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethodContactlessEmv           SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethod = "contactless_emv"
	SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethodContactlessMagstripeMode SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethod = "contactless_magstripe_mode"
	SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethodMagneticStripeFallback   SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethod = "magnetic_stripe_fallback"
	SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethodMagneticStripeTrack2     SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethod = "magnetic_stripe_track2"
)

// The local credit or debit card brand.
type SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandBc          SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "bc"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandCiti        SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "citi"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandHana        SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "hana"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandHyundai     SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "hyundai"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandJeju        SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "jeju"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandJeonbuk     SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "jeonbuk"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandKakaobank   SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "kakaobank"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandKbank       SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "kbank"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandKdbbank     SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "kdbbank"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandKookmin     SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "kookmin"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandKwangju     SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "kwangju"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandLotte       SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "lotte"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandMg          SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "mg"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandNh          SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "nh"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandPost        SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "post"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandSamsung     SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "samsung"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandSavingsbank SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "savingsbank"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandShinhan     SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "shinhan"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandShinhyup    SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "shinhyup"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandSuhyup      SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "suhyup"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandTossbank    SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "tossbank"
	SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrandWoori       SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand = "woori"
)

// Whether to fund this transaction with Naver Pay points or a card.
type SharedPaymentGrantedTokenPaymentMethodDetailsNaverPayFunding string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsNaverPayFunding can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsNaverPayFundingCard   SharedPaymentGrantedTokenPaymentMethodDetailsNaverPayFunding = "card"
	SharedPaymentGrantedTokenPaymentMethodDetailsNaverPayFundingPoints SharedPaymentGrantedTokenPaymentMethodDetailsNaverPayFunding = "points"
)

// The customer's bank, if provided.
type SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankAliorBank            SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "alior_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankBankMillennium       SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "bank_millennium"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankBankNowyBfgSa        SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "bank_nowy_bfg_sa"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankBankPekaoSa          SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "bank_pekao_sa"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankBankiSpbdzielcze     SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "banki_spbdzielcze"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankBLIK                 SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "blik"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankBnpParibas           SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "bnp_paribas"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankBoz                  SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "boz"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankCitiHandlowy         SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "citi_handlowy"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankCreditAgricole       SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "credit_agricole"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankEnvelobank           SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "envelobank"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankEtransferPocztowy24  SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "etransfer_pocztowy24"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankGetinBank            SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "getin_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankIdeabank             SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "ideabank"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankIng                  SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "ing"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankInteligo             SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "inteligo"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankMbankMtransfer       SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "mbank_mtransfer"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankNestPrzelew          SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "nest_przelew"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankNoblePay             SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "noble_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankPbacZIpko            SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "pbac_z_ipko"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankPlusBank             SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "plus_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankSantanderPrzelew24   SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "santander_przelew24"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankTmobileUsbugiBankowe SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "tmobile_usbugi_bankowe"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankToyotaBank           SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "toyota_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankVelobank             SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "velobank"
	SharedPaymentGrantedTokenPaymentMethodDetailsP24BankVolkswagenBank       SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank = "volkswagen_bank"
)

// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
type SharedPaymentGrantedTokenPaymentMethodDetailsType string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsType can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeACSSDebit        SharedPaymentGrantedTokenPaymentMethodDetailsType = "acss_debit"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeAffirm           SharedPaymentGrantedTokenPaymentMethodDetailsType = "affirm"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeAfterpayClearpay SharedPaymentGrantedTokenPaymentMethodDetailsType = "afterpay_clearpay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeAlipay           SharedPaymentGrantedTokenPaymentMethodDetailsType = "alipay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeAlma             SharedPaymentGrantedTokenPaymentMethodDetailsType = "alma"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeAmazonPay        SharedPaymentGrantedTokenPaymentMethodDetailsType = "amazon_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeAUBECSDebit      SharedPaymentGrantedTokenPaymentMethodDetailsType = "au_becs_debit"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeBACSDebit        SharedPaymentGrantedTokenPaymentMethodDetailsType = "bacs_debit"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeBancontact       SharedPaymentGrantedTokenPaymentMethodDetailsType = "bancontact"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeBillie           SharedPaymentGrantedTokenPaymentMethodDetailsType = "billie"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeBLIK             SharedPaymentGrantedTokenPaymentMethodDetailsType = "blik"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeBoleto           SharedPaymentGrantedTokenPaymentMethodDetailsType = "boleto"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeCard             SharedPaymentGrantedTokenPaymentMethodDetailsType = "card"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeCardPresent      SharedPaymentGrantedTokenPaymentMethodDetailsType = "card_present"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeCashApp          SharedPaymentGrantedTokenPaymentMethodDetailsType = "cashapp"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeCrypto           SharedPaymentGrantedTokenPaymentMethodDetailsType = "crypto"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeCustom           SharedPaymentGrantedTokenPaymentMethodDetailsType = "custom"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeCustomerBalance  SharedPaymentGrantedTokenPaymentMethodDetailsType = "customer_balance"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeEPS              SharedPaymentGrantedTokenPaymentMethodDetailsType = "eps"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeFPX              SharedPaymentGrantedTokenPaymentMethodDetailsType = "fpx"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeGiropay          SharedPaymentGrantedTokenPaymentMethodDetailsType = "giropay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeGopay            SharedPaymentGrantedTokenPaymentMethodDetailsType = "gopay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeGrabpay          SharedPaymentGrantedTokenPaymentMethodDetailsType = "grabpay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeIDBankTransfer   SharedPaymentGrantedTokenPaymentMethodDetailsType = "id_bank_transfer"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeIDEAL            SharedPaymentGrantedTokenPaymentMethodDetailsType = "ideal"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeInteracPresent   SharedPaymentGrantedTokenPaymentMethodDetailsType = "interac_present"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeKakaoPay         SharedPaymentGrantedTokenPaymentMethodDetailsType = "kakao_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeKlarna           SharedPaymentGrantedTokenPaymentMethodDetailsType = "klarna"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeKonbini          SharedPaymentGrantedTokenPaymentMethodDetailsType = "konbini"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeKrCard           SharedPaymentGrantedTokenPaymentMethodDetailsType = "kr_card"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeLink             SharedPaymentGrantedTokenPaymentMethodDetailsType = "link"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeMbWay            SharedPaymentGrantedTokenPaymentMethodDetailsType = "mb_way"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeMobilepay        SharedPaymentGrantedTokenPaymentMethodDetailsType = "mobilepay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeMultibanco       SharedPaymentGrantedTokenPaymentMethodDetailsType = "multibanco"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeNaverPay         SharedPaymentGrantedTokenPaymentMethodDetailsType = "naver_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeNzBankAccount    SharedPaymentGrantedTokenPaymentMethodDetailsType = "nz_bank_account"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeOXXO             SharedPaymentGrantedTokenPaymentMethodDetailsType = "oxxo"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeP24              SharedPaymentGrantedTokenPaymentMethodDetailsType = "p24"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypePayByBank        SharedPaymentGrantedTokenPaymentMethodDetailsType = "pay_by_bank"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypePayco            SharedPaymentGrantedTokenPaymentMethodDetailsType = "payco"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypePayNow           SharedPaymentGrantedTokenPaymentMethodDetailsType = "paynow"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypePaypal           SharedPaymentGrantedTokenPaymentMethodDetailsType = "paypal"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypePaypay           SharedPaymentGrantedTokenPaymentMethodDetailsType = "paypay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypePayto            SharedPaymentGrantedTokenPaymentMethodDetailsType = "payto"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypePix              SharedPaymentGrantedTokenPaymentMethodDetailsType = "pix"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypePromptPay        SharedPaymentGrantedTokenPaymentMethodDetailsType = "promptpay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeQris             SharedPaymentGrantedTokenPaymentMethodDetailsType = "qris"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeRechnung         SharedPaymentGrantedTokenPaymentMethodDetailsType = "rechnung"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeRevolutPay       SharedPaymentGrantedTokenPaymentMethodDetailsType = "revolut_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeSamsungPay       SharedPaymentGrantedTokenPaymentMethodDetailsType = "samsung_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeSatispay         SharedPaymentGrantedTokenPaymentMethodDetailsType = "satispay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeSEPADebit        SharedPaymentGrantedTokenPaymentMethodDetailsType = "sepa_debit"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeShopeepay        SharedPaymentGrantedTokenPaymentMethodDetailsType = "shopeepay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeSofort           SharedPaymentGrantedTokenPaymentMethodDetailsType = "sofort"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeStripeBalance    SharedPaymentGrantedTokenPaymentMethodDetailsType = "stripe_balance"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeSunbit           SharedPaymentGrantedTokenPaymentMethodDetailsType = "sunbit"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeSwish            SharedPaymentGrantedTokenPaymentMethodDetailsType = "swish"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeTWINT            SharedPaymentGrantedTokenPaymentMethodDetailsType = "twint"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeUpi              SharedPaymentGrantedTokenPaymentMethodDetailsType = "upi"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeUSBankAccount    SharedPaymentGrantedTokenPaymentMethodDetailsType = "us_bank_account"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeWeChatPay        SharedPaymentGrantedTokenPaymentMethodDetailsType = "wechat_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeZip              SharedPaymentGrantedTokenPaymentMethodDetailsType = "zip"
)

// Account holder type: individual or company.
type SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountHolderType string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountHolderType can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountHolderTypeCompany    SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountHolderType = "company"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountHolderTypeIndividual SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountHolderType = "individual"
)

// Account type: checkings or savings. Defaults to checking if omitted.
type SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountType string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountType can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountTypeChecking SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountType = "checking"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountTypeSavings  SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountType = "savings"
)

// All supported networks.
type SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountNetworksSupported string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountNetworksSupported can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountNetworksSupportedACH            SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountNetworksSupported = "ach"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountNetworksSupportedUSDomesticWire SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountNetworksSupported = "us_domestic_wire"
)

// The ACH network code that resulted in this block.
type SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR02 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R02"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR03 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R03"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR04 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R04"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR05 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R05"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR07 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R07"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR08 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R08"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR10 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R10"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR11 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R11"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR16 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R16"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR20 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R20"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR29 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R29"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCodeR31 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode = "R31"
)

// The reason why this PaymentMethod's fingerprint has been blocked
type SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReason string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReason can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReasonBankAccountClosed                 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReason = "bank_account_closed"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReasonBankAccountFrozen                 SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReason = "bank_account_frozen"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReasonBankAccountInvalidDetails         SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReason = "bank_account_invalid_details"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReasonBankAccountRestricted             SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReason = "bank_account_restricted"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReasonBankAccountUnusable               SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReason = "bank_account_unusable"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReasonDebitNotAuthorized                SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReason = "debit_not_authorized"
	SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReasonTokenizedAccountNumberDeactivated SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReason = "tokenized_account_number_deactivated"
)

// Retrieves an existing SharedPaymentGrantedToken object
type SharedPaymentGrantedTokenParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *SharedPaymentGrantedTokenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an existing SharedPaymentGrantedToken object
type SharedPaymentGrantedTokenRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *SharedPaymentGrantedTokenRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The total amount captured using this SharedPaymentToken.
type SharedPaymentGrantedTokenUsageDetailsAmountCaptured struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Integer value of the amount in the smallest currency unit.
	Value int64 `json:"value"`
}

// Some details about how the SharedPaymentGrantedToken has been used already.
type SharedPaymentGrantedTokenUsageDetails struct {
	// The total amount captured using this SharedPaymentToken.
	AmountCaptured *SharedPaymentGrantedTokenUsageDetailsAmountCaptured `json:"amount_captured"`
}

// Limits on how this SharedPaymentGrantedToken can be used.
type SharedPaymentGrantedTokenUsageLimits struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Time at which this SharedPaymentToken expires and can no longer be used to confirm a PaymentIntent.
	ExpiresAt int64 `json:"expires_at"`
	// Max amount that can be captured using this SharedPaymentToken.
	MaxAmount int64 `json:"max_amount"`
	// The recurring interval at which the shared payment token's amount usage restrictions reset.
	RecurringInterval SharedPaymentGrantedTokenUsageLimitsRecurringInterval `json:"recurring_interval,omitempty"`
}

// Details about the agent that issued this SharedPaymentGrantedToken.
type SharedPaymentGrantedTokenAgentDetails struct {
	// The Stripe Profile ID of the agent that issued this SharedPaymentGrantedToken.
	NetworkBusinessProfile string `json:"network_business_profile"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsACSSDebit struct {
	// Account number of the bank account.
	AccountNumber string `json:"account_number,omitempty"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Institution number of the bank account.
	InstitutionNumber string `json:"institution_number"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// Transit number of the bank account.
	TransitNumber string `json:"transit_number"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsAffirm struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsAfterpayClearpay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsAlipay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsAlma struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsAmazonPay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsAUBECSDebit struct {
	// Six-digit number identifying bank and branch associated with this bank account.
	BSBNumber string `json:"bsb_number"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsBACSDebit struct {
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode string `json:"sort_code"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsBancontact struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsBillie struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type SharedPaymentGrantedTokenPaymentMethodDetailsBillingDetails struct {
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
type SharedPaymentGrantedTokenPaymentMethodDetailsBLIK struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsBoleto struct {
	// Uniquely identifies the customer tax id (CNPJ or CPF)
	TaxID string `json:"tax_id"`
}

// Checks on Card address and CVC if provided.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardChecks struct {
	// If a address line1 was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressLine1Check string `json:"address_line1_check"`
	// If a address postal code was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressPostalCodeCheck string `json:"address_postal_code_check"`
	// If a CVC was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	CVCCheck string `json:"cvc_check"`
}

// Contains information about card networks that can be used to process the payment.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardNetworks struct {
	// All networks available for selection via [payment_method_options.card.network](https://docs.stripe.com/api/payment_intents/confirm#confirm_payment_intent-payment_method_options-card-network).
	Available []string `json:"available"`
	// The preferred network for co-branded cards. Can be `cartes_bancaires`, `mastercard`, `visa` or `invalid_preference` if requested network is not valid for the card.
	Preferred string `json:"preferred"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletAmexExpressCheckout struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletApplePay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletGooglePay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletLink struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletMasterpass struct {
	// Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	BillingAddress *Address `json:"billing_address"`
	// Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Email string `json:"email"`
	// Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Name string `json:"name"`
	// Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	ShippingAddress *Address `json:"shipping_address"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletSamsungPay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletVisaCheckout struct {
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
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWallet struct {
	AmexExpressCheckout *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletAmexExpressCheckout `json:"amex_express_checkout,omitempty"`
	ApplePay            *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletApplePay            `json:"apple_pay,omitempty"`
	// (For tokenized numbers only.) The last four digits of the device account number.
	DynamicLast4 string                                                             `json:"dynamic_last4"`
	GooglePay    *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletGooglePay  `json:"google_pay,omitempty"`
	Link         *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletLink       `json:"link,omitempty"`
	Masterpass   *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletMasterpass `json:"masterpass,omitempty"`
	SamsungPay   *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletSamsungPay `json:"samsung_pay,omitempty"`
	// The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, `visa_checkout`, or `link`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
	Type         SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType          `json:"type"`
	VisaCheckout *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletVisaCheckout `json:"visa_checkout,omitempty"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsCard struct {
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand string `json:"brand"`
	// Checks on Card address and CVC if provided.
	Checks *SharedPaymentGrantedTokenPaymentMethodDetailsCardChecks `json:"checks,omitempty"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description,omitempty"`
	// The brand to use when displaying the card, this accounts for customer's brand choice on dual-branded cards. Can be `american_express`, `cartes_bancaires`, `diners_club`, `discover`, `eftpos_australia`, `interac`, `jcb`, `mastercard`, `union_pay`, `visa`, or `other` and may contain more values in the future.
	DisplayBrand string `json:"display_brand"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *As of May 1, 2021, card fingerprint in India for Connect changed to allow two fingerprints for the same card---one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint,omitempty"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding string `json:"funding"`
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin,omitempty"`
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer,omitempty"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Contains information about card networks that can be used to process the payment.
	Networks *SharedPaymentGrantedTokenPaymentMethodDetailsCardNetworks `json:"networks"`
	// If this Card is part of a card wallet, this contains the details of the card wallet.
	Wallet *SharedPaymentGrantedTokenPaymentMethodDetailsCardWallet `json:"wallet"`
}

// Contains information about card networks that can be used to process the payment.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentNetworks struct {
	// All networks available for selection via [payment_method_options.card.network](https://docs.stripe.com/api/payment_intents/confirm#confirm_payment_intent-payment_method_options-card-network).
	Available []string `json:"available"`
	// The preferred network for the card.
	Preferred string `json:"preferred"`
}

// Details about payment methods collected offline.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentOffline struct {
	// Time at which the payment was collected while offline
	StoredAt int64 `json:"stored_at"`
	// The method used to process this payment method offline. Only deferred is allowed.
	Type SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentOfflineType `json:"type"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWallet struct {
	// The type of mobile wallet, one of `apple_pay`, `google_pay`, `samsung_pay`, or `unknown`.
	Type SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWalletType `json:"type"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardPresent struct {
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand string `json:"brand"`
	// The [product code](https://stripe.com/docs/card-product-codes) that identifies the specific program or product associated with a card.
	BrandProduct string `json:"brand_product"`
	// The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.
	CardholderName string `json:"cardholder_name"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description,omitempty"`
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
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin,omitempty"`
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer,omitempty"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Contains information about card networks that can be used to process the payment.
	Networks *SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentNetworks `json:"networks"`
	// Details about payment methods collected offline.
	Offline *SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentOffline `json:"offline"`
	// The languages that the issuing bank recommends using for localizing any customer-facing text, as read from the card. Referenced from EMV tag 5F2D, data encoded on the card's chip.
	PreferredLocales []string `json:"preferred_locales"`
	// How card details were read in this transaction.
	ReadMethod SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentReadMethod `json:"read_method"`
	Wallet     *SharedPaymentGrantedTokenPaymentMethodDetailsCardPresentWallet    `json:"wallet,omitempty"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsCashApp struct {
	// A unique and immutable identifier assigned by Cash App to every buyer.
	BuyerID string `json:"buyer_id"`
	// A public identifier for buyers using Cash App.
	Cashtag string `json:"cashtag"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsCrypto struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCustomerBalance struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsEPS struct {
	// The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `deutsche_bank_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.
	Bank SharedPaymentGrantedTokenPaymentMethodDetailsEPSBank `json:"bank"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsFPX struct {
	// Account holder type, if provided. Can be one of `individual` or `company`.
	AccountHolderType SharedPaymentGrantedTokenPaymentMethodDetailsFPXAccountHolderType `json:"account_holder_type"`
	// The customer's bank, if provided. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, `pb_enterprise`, or `bank_of_china`.
	Bank SharedPaymentGrantedTokenPaymentMethodDetailsFPXBank `json:"bank"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsGiropay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsGopay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsGrabpay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransfer struct {
	Bank        SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransferBank `json:"bank"`
	BankCode    string                                                          `json:"bank_code"`
	BankName    string                                                          `json:"bank_name"`
	DisplayName string                                                          `json:"display_name"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsIDEAL struct {
	// The customer's bank, if provided. Can be one of `abn_amro`, `adyen`, `asn_bank`, `bunq`, `buut`, `finom`, `handelsbanken`, `ing`, `knab`, `mollie`, `moneyou`, `n26`, `nn`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, `van_lanschot`, or `yoursafe`.
	Bank SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBank `json:"bank"`
	// The Bank Identifier Code of the customer's bank, if the bank was provided.
	BIC SharedPaymentGrantedTokenPaymentMethodDetailsIDEALBIC `json:"bic"`
}

// Contains information about card networks that can be used to process the payment.
type SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentNetworks struct {
	// All networks available for selection via [payment_method_options.card.network](https://docs.stripe.com/api/payment_intents/confirm#confirm_payment_intent-payment_method_options-card-network).
	Available []string `json:"available"`
	// The preferred network for the card.
	Preferred string `json:"preferred"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresent struct {
	// Card brand. Can be `interac`, `mastercard` or `visa`.
	Brand string `json:"brand"`
	// The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.
	CardholderName string `json:"cardholder_name"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description,omitempty"`
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
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin,omitempty"`
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer,omitempty"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Contains information about card networks that can be used to process the payment.
	Networks *SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentNetworks `json:"networks"`
	// The languages that the issuing bank recommends using for localizing any customer-facing text, as read from the card. Referenced from EMV tag 5F2D, data encoded on the card's chip.
	PreferredLocales []string `json:"preferred_locales"`
	// How card details were read in this transaction.
	ReadMethod SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresentReadMethod `json:"read_method"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsKakaoPay struct{}

// The customer's date of birth, if provided.
type SharedPaymentGrantedTokenPaymentMethodDetailsKlarnaDOB struct {
	// The day of birth, between 1 and 31.
	Day int64 `json:"day"`
	// The month of birth, between 1 and 12.
	Month int64 `json:"month"`
	// The four-digit year of birth.
	Year int64 `json:"year"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsKlarna struct {
	// The customer's date of birth, if provided.
	DOB *SharedPaymentGrantedTokenPaymentMethodDetailsKlarnaDOB `json:"dob,omitempty"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsKonbini struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsKrCard struct {
	// The local credit or debit card brand.
	Brand SharedPaymentGrantedTokenPaymentMethodDetailsKrCardBrand `json:"brand"`
	// The last four digits of the card. This may not be present for American Express cards.
	Last4 string `json:"last4"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsLink struct {
	// Account owner's email address.
	Email string `json:"email"`
	// [Deprecated] This is a legacy parameter that no longer has any function.
	// Deprecated:
	PersistentToken string `json:"persistent_token,omitempty"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsMbWay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsMobilepay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsMultibanco struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsNaverPay struct {
	// Uniquely identifies this particular Naver Pay account. You can use this attribute to check whether two Naver Pay accounts are the same.
	BuyerID string `json:"buyer_id"`
	// Whether to fund this transaction with Naver Pay points or a card.
	Funding SharedPaymentGrantedTokenPaymentMethodDetailsNaverPayFunding `json:"funding"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsNzBankAccount struct {
	// The name on the bank account. Only present if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod's billing details.
	AccountHolderName string `json:"account_holder_name"`
	// The numeric code for the bank account's bank.
	BankCode string `json:"bank_code"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// The numeric code for the bank account's bank branch.
	BranchCode string `json:"branch_code"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// The suffix of the bank account number.
	Suffix string `json:"suffix"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsOXXO struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsP24 struct {
	// The customer's bank, if provided.
	Bank SharedPaymentGrantedTokenPaymentMethodDetailsP24Bank `json:"bank"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsPayByBank struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsPayco struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsPayNow struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsPaypal struct {
	// Two-letter ISO code representing the buyer's country. Values are provided by PayPal directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Country string `json:"country"`
	// Uniquely identifies this particular PayPal account. You can use this attribute to check whether two PayPal accounts are the same.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Owner's email. Values are provided by PayPal directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	PayerEmail string `json:"payer_email"`
	// PayPal account PayerID. This identifier uniquely identifies the PayPal customer.
	PayerID string `json:"payer_id"`
	// Owner's verified email. Values are verified or provided by PayPal directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedEmail string `json:"verified_email,omitempty"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsPaypay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsPayto struct {
	// Bank-State-Branch number of the bank account.
	BSBNumber string `json:"bsb_number"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// The PayID alias for the bank account.
	PayID string `json:"pay_id"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsPix struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsPromptPay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsQris struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsRechnungDOB struct {
	// The day of birth, between 1 and 31.
	Day int64 `json:"day"`
	// The month of birth, between 1 and 12.
	Month int64 `json:"month"`
	// The four-digit year of birth.
	Year int64 `json:"year"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsRechnung struct {
	DOB *SharedPaymentGrantedTokenPaymentMethodDetailsRechnungDOB `json:"dob,omitempty"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsRevolutPay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsSamsungPay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsSatispay struct{}

// Information about the object that generated this PaymentMethod.
type SharedPaymentGrantedTokenPaymentMethodDetailsSEPADebitGeneratedFrom struct {
	// The ID of the Charge that generated this PaymentMethod, if any.
	Charge *Charge `json:"charge"`
	// The ID of the SetupAttempt that generated this PaymentMethod, if any.
	SetupAttempt *SetupAttempt `json:"setup_attempt"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsSEPADebit struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Branch code of bank associated with the bank account.
	BranchCode string `json:"branch_code"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Information about the object that generated this PaymentMethod.
	GeneratedFrom *SharedPaymentGrantedTokenPaymentMethodDetailsSEPADebitGeneratedFrom `json:"generated_from"`
	// Last four characters of the IBAN.
	Last4 string `json:"last4"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsShopeepay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsSofort struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsStripeBalance struct {
	// The connected account ID whose Stripe balance to use as the source of payment
	Account string `json:"account,omitempty"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsSunbit struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsSwish struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsTWINT struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsUpi struct {
	// Customer's unique Virtual Payment Address
	Vpa string `json:"vpa"`
}

// Contains information about US bank account networks that can be used.
type SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountNetworks struct {
	// The preferred network.
	Preferred string `json:"preferred"`
	// All supported networks.
	Supported []SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountNetworksSupported `json:"supported"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlocked struct {
	// The ACH network code that resulted in this block.
	NetworkCode SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedNetworkCode `json:"network_code"`
	// The reason why this PaymentMethod's fingerprint has been blocked
	Reason SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlockedReason `json:"reason"`
}

// Contains information about the future reusability of this PaymentMethod.
type SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetails struct {
	Blocked *SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetailsBlocked `json:"blocked,omitempty"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccount struct {
	// Account holder type: individual or company.
	AccountHolderType SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountHolderType `json:"account_holder_type"`
	// Account number of the bank account.
	AccountNumber string `json:"account_number,omitempty"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountAccountType `json:"account_type"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// The ID of the Financial Connections Account used to create the payment method.
	FinancialConnectionsAccount string `json:"financial_connections_account"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// Contains information about US bank account networks that can be used.
	Networks *SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountNetworks `json:"networks"`
	// Routing number of the bank account.
	RoutingNumber string `json:"routing_number"`
	// Contains information about the future reusability of this PaymentMethod.
	StatusDetails *SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccountStatusDetails `json:"status_details"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsWeChatPay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsZip struct{}

// Details of the PaymentMethod that was shared via this token.
type SharedPaymentGrantedTokenPaymentMethodDetails struct {
	ACSSDebit        *SharedPaymentGrantedTokenPaymentMethodDetailsACSSDebit        `json:"acss_debit,omitempty"`
	Affirm           *SharedPaymentGrantedTokenPaymentMethodDetailsAffirm           `json:"affirm,omitempty"`
	AfterpayClearpay *SharedPaymentGrantedTokenPaymentMethodDetailsAfterpayClearpay `json:"afterpay_clearpay,omitempty"`
	Alipay           *SharedPaymentGrantedTokenPaymentMethodDetailsAlipay           `json:"alipay,omitempty"`
	Alma             *SharedPaymentGrantedTokenPaymentMethodDetailsAlma             `json:"alma,omitempty"`
	AmazonPay        *SharedPaymentGrantedTokenPaymentMethodDetailsAmazonPay        `json:"amazon_pay,omitempty"`
	AUBECSDebit      *SharedPaymentGrantedTokenPaymentMethodDetailsAUBECSDebit      `json:"au_becs_debit,omitempty"`
	BACSDebit        *SharedPaymentGrantedTokenPaymentMethodDetailsBACSDebit        `json:"bacs_debit,omitempty"`
	Bancontact       *SharedPaymentGrantedTokenPaymentMethodDetailsBancontact       `json:"bancontact,omitempty"`
	Billie           *SharedPaymentGrantedTokenPaymentMethodDetailsBillie           `json:"billie,omitempty"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails  *SharedPaymentGrantedTokenPaymentMethodDetailsBillingDetails  `json:"billing_details"`
	BLIK            *SharedPaymentGrantedTokenPaymentMethodDetailsBLIK            `json:"blik,omitempty"`
	Boleto          *SharedPaymentGrantedTokenPaymentMethodDetailsBoleto          `json:"boleto,omitempty"`
	Card            *SharedPaymentGrantedTokenPaymentMethodDetailsCard            `json:"card,omitempty"`
	CardPresent     *SharedPaymentGrantedTokenPaymentMethodDetailsCardPresent     `json:"card_present,omitempty"`
	CashApp         *SharedPaymentGrantedTokenPaymentMethodDetailsCashApp         `json:"cashapp,omitempty"`
	Crypto          *SharedPaymentGrantedTokenPaymentMethodDetailsCrypto          `json:"crypto,omitempty"`
	CustomerBalance *SharedPaymentGrantedTokenPaymentMethodDetailsCustomerBalance `json:"customer_balance,omitempty"`
	EPS             *SharedPaymentGrantedTokenPaymentMethodDetailsEPS             `json:"eps,omitempty"`
	FPX             *SharedPaymentGrantedTokenPaymentMethodDetailsFPX             `json:"fpx,omitempty"`
	Giropay         *SharedPaymentGrantedTokenPaymentMethodDetailsGiropay         `json:"giropay,omitempty"`
	Gopay           *SharedPaymentGrantedTokenPaymentMethodDetailsGopay           `json:"gopay,omitempty"`
	Grabpay         *SharedPaymentGrantedTokenPaymentMethodDetailsGrabpay         `json:"grabpay,omitempty"`
	IDBankTransfer  *SharedPaymentGrantedTokenPaymentMethodDetailsIDBankTransfer  `json:"id_bank_transfer,omitempty"`
	IDEAL           *SharedPaymentGrantedTokenPaymentMethodDetailsIDEAL           `json:"ideal,omitempty"`
	InteracPresent  *SharedPaymentGrantedTokenPaymentMethodDetailsInteracPresent  `json:"interac_present,omitempty"`
	KakaoPay        *SharedPaymentGrantedTokenPaymentMethodDetailsKakaoPay        `json:"kakao_pay,omitempty"`
	Klarna          *SharedPaymentGrantedTokenPaymentMethodDetailsKlarna          `json:"klarna,omitempty"`
	Konbini         *SharedPaymentGrantedTokenPaymentMethodDetailsKonbini         `json:"konbini,omitempty"`
	KrCard          *SharedPaymentGrantedTokenPaymentMethodDetailsKrCard          `json:"kr_card,omitempty"`
	Link            *SharedPaymentGrantedTokenPaymentMethodDetailsLink            `json:"link,omitempty"`
	MbWay           *SharedPaymentGrantedTokenPaymentMethodDetailsMbWay           `json:"mb_way,omitempty"`
	Mobilepay       *SharedPaymentGrantedTokenPaymentMethodDetailsMobilepay       `json:"mobilepay,omitempty"`
	Multibanco      *SharedPaymentGrantedTokenPaymentMethodDetailsMultibanco      `json:"multibanco,omitempty"`
	NaverPay        *SharedPaymentGrantedTokenPaymentMethodDetailsNaverPay        `json:"naver_pay,omitempty"`
	NzBankAccount   *SharedPaymentGrantedTokenPaymentMethodDetailsNzBankAccount   `json:"nz_bank_account,omitempty"`
	OXXO            *SharedPaymentGrantedTokenPaymentMethodDetailsOXXO            `json:"oxxo,omitempty"`
	P24             *SharedPaymentGrantedTokenPaymentMethodDetailsP24             `json:"p24,omitempty"`
	PayByBank       *SharedPaymentGrantedTokenPaymentMethodDetailsPayByBank       `json:"pay_by_bank,omitempty"`
	Payco           *SharedPaymentGrantedTokenPaymentMethodDetailsPayco           `json:"payco,omitempty"`
	PayNow          *SharedPaymentGrantedTokenPaymentMethodDetailsPayNow          `json:"paynow,omitempty"`
	Paypal          *SharedPaymentGrantedTokenPaymentMethodDetailsPaypal          `json:"paypal,omitempty"`
	Paypay          *SharedPaymentGrantedTokenPaymentMethodDetailsPaypay          `json:"paypay,omitempty"`
	Payto           *SharedPaymentGrantedTokenPaymentMethodDetailsPayto           `json:"payto,omitempty"`
	Pix             *SharedPaymentGrantedTokenPaymentMethodDetailsPix             `json:"pix,omitempty"`
	PromptPay       *SharedPaymentGrantedTokenPaymentMethodDetailsPromptPay       `json:"promptpay,omitempty"`
	Qris            *SharedPaymentGrantedTokenPaymentMethodDetailsQris            `json:"qris,omitempty"`
	Rechnung        *SharedPaymentGrantedTokenPaymentMethodDetailsRechnung        `json:"rechnung,omitempty"`
	RevolutPay      *SharedPaymentGrantedTokenPaymentMethodDetailsRevolutPay      `json:"revolut_pay,omitempty"`
	SamsungPay      *SharedPaymentGrantedTokenPaymentMethodDetailsSamsungPay      `json:"samsung_pay,omitempty"`
	Satispay        *SharedPaymentGrantedTokenPaymentMethodDetailsSatispay        `json:"satispay,omitempty"`
	SEPADebit       *SharedPaymentGrantedTokenPaymentMethodDetailsSEPADebit       `json:"sepa_debit,omitempty"`
	Shopeepay       *SharedPaymentGrantedTokenPaymentMethodDetailsShopeepay       `json:"shopeepay,omitempty"`
	Sofort          *SharedPaymentGrantedTokenPaymentMethodDetailsSofort          `json:"sofort,omitempty"`
	StripeBalance   *SharedPaymentGrantedTokenPaymentMethodDetailsStripeBalance   `json:"stripe_balance,omitempty"`
	Sunbit          *SharedPaymentGrantedTokenPaymentMethodDetailsSunbit          `json:"sunbit,omitempty"`
	Swish           *SharedPaymentGrantedTokenPaymentMethodDetailsSwish           `json:"swish,omitempty"`
	TWINT           *SharedPaymentGrantedTokenPaymentMethodDetailsTWINT           `json:"twint,omitempty"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type          SharedPaymentGrantedTokenPaymentMethodDetailsType           `json:"type"`
	Upi           *SharedPaymentGrantedTokenPaymentMethodDetailsUpi           `json:"upi,omitempty"`
	USBankAccount *SharedPaymentGrantedTokenPaymentMethodDetailsUSBankAccount `json:"us_bank_account,omitempty"`
	WeChatPay     *SharedPaymentGrantedTokenPaymentMethodDetailsWeChatPay     `json:"wechat_pay,omitempty"`
	Zip           *SharedPaymentGrantedTokenPaymentMethodDetailsZip           `json:"zip,omitempty"`
}

// Bot risk insight (score: Float, recommended_action).
type SharedPaymentGrantedTokenRiskDetailsInsightsBot struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight (float).
	Score float64 `json:"score"`
}

// Card issuer decline risk insight (score: Float, recommended_action).
type SharedPaymentGrantedTokenRiskDetailsInsightsCardIssuerDecline struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight (float).
	Score float64 `json:"score"`
}

// Card testing risk insight (score: Float, recommended_action).
type SharedPaymentGrantedTokenRiskDetailsInsightsCardTesting struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight (float).
	Score float64 `json:"score"`
}

// Fraudulent dispute risk insight (score: Integer, recommended_action).
type SharedPaymentGrantedTokenRiskDetailsInsightsFraudulentDispute struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight (integer).
	Score int64 `json:"score"`
}

// Stolen card risk insight (score: Integer, recommended_action).
type SharedPaymentGrantedTokenRiskDetailsInsightsStolenCard struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight (integer).
	Score int64 `json:"score"`
}

// Risk insights for this token, including scores and recommended actions for each risk type.
type SharedPaymentGrantedTokenRiskDetailsInsights struct {
	// Bot risk insight (score: Float, recommended_action).
	Bot *SharedPaymentGrantedTokenRiskDetailsInsightsBot `json:"bot,omitempty"`
	// Card issuer decline risk insight (score: Float, recommended_action).
	CardIssuerDecline *SharedPaymentGrantedTokenRiskDetailsInsightsCardIssuerDecline `json:"card_issuer_decline,omitempty"`
	// Card testing risk insight (score: Float, recommended_action).
	CardTesting *SharedPaymentGrantedTokenRiskDetailsInsightsCardTesting `json:"card_testing,omitempty"`
	// Fraudulent dispute risk insight (score: Integer, recommended_action).
	FraudulentDispute *SharedPaymentGrantedTokenRiskDetailsInsightsFraudulentDispute `json:"fraudulent_dispute"`
	// Stolen card risk insight (score: Integer, recommended_action).
	StolenCard *SharedPaymentGrantedTokenRiskDetailsInsightsStolenCard `json:"stolen_card,omitempty"`
}

// Risk details of the SharedPaymentGrantedToken.
type SharedPaymentGrantedTokenRiskDetails struct {
	// Risk insights for this token, including scores and recommended actions for each risk type.
	Insights *SharedPaymentGrantedTokenRiskDetailsInsights `json:"insights"`
}

// SharedPaymentGrantedToken is the view-only resource of a SharedPaymentIssuedToken, which is a limited-use reference to a PaymentMethod.
// When another Stripe merchant shares a SharedPaymentIssuedToken with you, you can view attributes of the shared token using the SharedPaymentGrantedToken API, and use it with a PaymentIntent.
type SharedPaymentGrantedToken struct {
	APIResource
	// Details about the agent that issued this SharedPaymentGrantedToken.
	AgentDetails *SharedPaymentGrantedTokenAgentDetails `json:"agent_details,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Time at which this SharedPaymentGrantedToken expires and can no longer be used to confirm a PaymentIntent.
	DeactivatedAt int64 `json:"deactivated_at"`
	// The reason why the SharedPaymentGrantedToken has been deactivated.
	DeactivatedReason SharedPaymentGrantedTokenDeactivatedReason `json:"deactivated_reason"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Details of the PaymentMethod that was shared via this token.
	PaymentMethodDetails *SharedPaymentGrantedTokenPaymentMethodDetails `json:"payment_method_details"`
	// Risk details of the SharedPaymentGrantedToken.
	RiskDetails *SharedPaymentGrantedTokenRiskDetails `json:"risk_details,omitempty"`
	// Metadata about the SharedPaymentGrantedToken.
	SharedMetadata map[string]string `json:"shared_metadata"`
	// Some details about how the SharedPaymentGrantedToken has been used already.
	UsageDetails *SharedPaymentGrantedTokenUsageDetails `json:"usage_details"`
	// Limits on how this SharedPaymentGrantedToken can be used.
	UsageLimits *SharedPaymentGrantedTokenUsageLimits `json:"usage_limits"`
}
