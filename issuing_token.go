//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The token service provider / card network associated with the token.
type IssuingTokenNetwork string

// List of values that IssuingTokenNetwork can take
const (
	IssuingTokenNetworkMastercard IssuingTokenNetwork = "mastercard"
	IssuingTokenNetworkVisa       IssuingTokenNetwork = "visa"
)

// The ISO 639-1 language code of the device associated with the tokenization request.
type IssuingTokenNetworkDataDeviceLanguage string

// List of values that IssuingTokenNetworkDataDeviceLanguage can take
const (
	IssuingTokenNetworkDataDeviceLanguageAa IssuingTokenNetworkDataDeviceLanguage = "aa"
	IssuingTokenNetworkDataDeviceLanguageAb IssuingTokenNetworkDataDeviceLanguage = "ab"
	IssuingTokenNetworkDataDeviceLanguageAe IssuingTokenNetworkDataDeviceLanguage = "ae"
	IssuingTokenNetworkDataDeviceLanguageAf IssuingTokenNetworkDataDeviceLanguage = "af"
	IssuingTokenNetworkDataDeviceLanguageAk IssuingTokenNetworkDataDeviceLanguage = "ak"
	IssuingTokenNetworkDataDeviceLanguageAm IssuingTokenNetworkDataDeviceLanguage = "am"
	IssuingTokenNetworkDataDeviceLanguageAn IssuingTokenNetworkDataDeviceLanguage = "an"
	IssuingTokenNetworkDataDeviceLanguageAR IssuingTokenNetworkDataDeviceLanguage = "ar"
	IssuingTokenNetworkDataDeviceLanguageAs IssuingTokenNetworkDataDeviceLanguage = "as"
	IssuingTokenNetworkDataDeviceLanguageAv IssuingTokenNetworkDataDeviceLanguage = "av"
	IssuingTokenNetworkDataDeviceLanguageAy IssuingTokenNetworkDataDeviceLanguage = "ay"
	IssuingTokenNetworkDataDeviceLanguageAz IssuingTokenNetworkDataDeviceLanguage = "az"
	IssuingTokenNetworkDataDeviceLanguageBa IssuingTokenNetworkDataDeviceLanguage = "ba"
	IssuingTokenNetworkDataDeviceLanguageBe IssuingTokenNetworkDataDeviceLanguage = "be"
	IssuingTokenNetworkDataDeviceLanguageBG IssuingTokenNetworkDataDeviceLanguage = "bg"
	IssuingTokenNetworkDataDeviceLanguageBi IssuingTokenNetworkDataDeviceLanguage = "bi"
	IssuingTokenNetworkDataDeviceLanguageBm IssuingTokenNetworkDataDeviceLanguage = "bm"
	IssuingTokenNetworkDataDeviceLanguageBn IssuingTokenNetworkDataDeviceLanguage = "bn"
	IssuingTokenNetworkDataDeviceLanguageBo IssuingTokenNetworkDataDeviceLanguage = "bo"
	IssuingTokenNetworkDataDeviceLanguageBr IssuingTokenNetworkDataDeviceLanguage = "br"
	IssuingTokenNetworkDataDeviceLanguageBs IssuingTokenNetworkDataDeviceLanguage = "bs"
	IssuingTokenNetworkDataDeviceLanguageCa IssuingTokenNetworkDataDeviceLanguage = "ca"
	IssuingTokenNetworkDataDeviceLanguageCe IssuingTokenNetworkDataDeviceLanguage = "ce"
	IssuingTokenNetworkDataDeviceLanguageCh IssuingTokenNetworkDataDeviceLanguage = "ch"
	IssuingTokenNetworkDataDeviceLanguageCo IssuingTokenNetworkDataDeviceLanguage = "co"
	IssuingTokenNetworkDataDeviceLanguageCr IssuingTokenNetworkDataDeviceLanguage = "cr"
	IssuingTokenNetworkDataDeviceLanguageCS IssuingTokenNetworkDataDeviceLanguage = "cs"
	IssuingTokenNetworkDataDeviceLanguageCu IssuingTokenNetworkDataDeviceLanguage = "cu"
	IssuingTokenNetworkDataDeviceLanguageCv IssuingTokenNetworkDataDeviceLanguage = "cv"
	IssuingTokenNetworkDataDeviceLanguageCy IssuingTokenNetworkDataDeviceLanguage = "cy"
	IssuingTokenNetworkDataDeviceLanguageDA IssuingTokenNetworkDataDeviceLanguage = "da"
	IssuingTokenNetworkDataDeviceLanguageDE IssuingTokenNetworkDataDeviceLanguage = "de"
	IssuingTokenNetworkDataDeviceLanguageDv IssuingTokenNetworkDataDeviceLanguage = "dv"
	IssuingTokenNetworkDataDeviceLanguageDz IssuingTokenNetworkDataDeviceLanguage = "dz"
	IssuingTokenNetworkDataDeviceLanguageEe IssuingTokenNetworkDataDeviceLanguage = "ee"
	IssuingTokenNetworkDataDeviceLanguageEL IssuingTokenNetworkDataDeviceLanguage = "el"
	IssuingTokenNetworkDataDeviceLanguageEN IssuingTokenNetworkDataDeviceLanguage = "en"
	IssuingTokenNetworkDataDeviceLanguageEo IssuingTokenNetworkDataDeviceLanguage = "eo"
	IssuingTokenNetworkDataDeviceLanguageES IssuingTokenNetworkDataDeviceLanguage = "es"
	IssuingTokenNetworkDataDeviceLanguageET IssuingTokenNetworkDataDeviceLanguage = "et"
	IssuingTokenNetworkDataDeviceLanguageEU IssuingTokenNetworkDataDeviceLanguage = "eu"
	IssuingTokenNetworkDataDeviceLanguageFa IssuingTokenNetworkDataDeviceLanguage = "fa"
	IssuingTokenNetworkDataDeviceLanguageFf IssuingTokenNetworkDataDeviceLanguage = "ff"
	IssuingTokenNetworkDataDeviceLanguageFI IssuingTokenNetworkDataDeviceLanguage = "fi"
	IssuingTokenNetworkDataDeviceLanguageFj IssuingTokenNetworkDataDeviceLanguage = "fj"
	IssuingTokenNetworkDataDeviceLanguageFo IssuingTokenNetworkDataDeviceLanguage = "fo"
	IssuingTokenNetworkDataDeviceLanguageFR IssuingTokenNetworkDataDeviceLanguage = "fr"
	IssuingTokenNetworkDataDeviceLanguageFy IssuingTokenNetworkDataDeviceLanguage = "fy"
	IssuingTokenNetworkDataDeviceLanguageGa IssuingTokenNetworkDataDeviceLanguage = "ga"
	IssuingTokenNetworkDataDeviceLanguageGd IssuingTokenNetworkDataDeviceLanguage = "gd"
	IssuingTokenNetworkDataDeviceLanguageGl IssuingTokenNetworkDataDeviceLanguage = "gl"
	IssuingTokenNetworkDataDeviceLanguageGn IssuingTokenNetworkDataDeviceLanguage = "gn"
	IssuingTokenNetworkDataDeviceLanguageGu IssuingTokenNetworkDataDeviceLanguage = "gu"
	IssuingTokenNetworkDataDeviceLanguageGv IssuingTokenNetworkDataDeviceLanguage = "gv"
	IssuingTokenNetworkDataDeviceLanguageHa IssuingTokenNetworkDataDeviceLanguage = "ha"
	IssuingTokenNetworkDataDeviceLanguageHE IssuingTokenNetworkDataDeviceLanguage = "he"
	IssuingTokenNetworkDataDeviceLanguageHi IssuingTokenNetworkDataDeviceLanguage = "hi"
	IssuingTokenNetworkDataDeviceLanguageHo IssuingTokenNetworkDataDeviceLanguage = "ho"
	IssuingTokenNetworkDataDeviceLanguageHR IssuingTokenNetworkDataDeviceLanguage = "hr"
	IssuingTokenNetworkDataDeviceLanguageHt IssuingTokenNetworkDataDeviceLanguage = "ht"
	IssuingTokenNetworkDataDeviceLanguageHU IssuingTokenNetworkDataDeviceLanguage = "hu"
	IssuingTokenNetworkDataDeviceLanguageHy IssuingTokenNetworkDataDeviceLanguage = "hy"
	IssuingTokenNetworkDataDeviceLanguageHz IssuingTokenNetworkDataDeviceLanguage = "hz"
	IssuingTokenNetworkDataDeviceLanguageIa IssuingTokenNetworkDataDeviceLanguage = "ia"
	IssuingTokenNetworkDataDeviceLanguageID IssuingTokenNetworkDataDeviceLanguage = "id"
	IssuingTokenNetworkDataDeviceLanguageIe IssuingTokenNetworkDataDeviceLanguage = "ie"
	IssuingTokenNetworkDataDeviceLanguageIg IssuingTokenNetworkDataDeviceLanguage = "ig"
	IssuingTokenNetworkDataDeviceLanguageIi IssuingTokenNetworkDataDeviceLanguage = "ii"
	IssuingTokenNetworkDataDeviceLanguageIk IssuingTokenNetworkDataDeviceLanguage = "ik"
	IssuingTokenNetworkDataDeviceLanguageIo IssuingTokenNetworkDataDeviceLanguage = "io"
	IssuingTokenNetworkDataDeviceLanguageIs IssuingTokenNetworkDataDeviceLanguage = "is"
	IssuingTokenNetworkDataDeviceLanguageIT IssuingTokenNetworkDataDeviceLanguage = "it"
	IssuingTokenNetworkDataDeviceLanguageIu IssuingTokenNetworkDataDeviceLanguage = "iu"
	IssuingTokenNetworkDataDeviceLanguageJA IssuingTokenNetworkDataDeviceLanguage = "ja"
	IssuingTokenNetworkDataDeviceLanguageJv IssuingTokenNetworkDataDeviceLanguage = "jv"
	IssuingTokenNetworkDataDeviceLanguageKa IssuingTokenNetworkDataDeviceLanguage = "ka"
	IssuingTokenNetworkDataDeviceLanguageKg IssuingTokenNetworkDataDeviceLanguage = "kg"
	IssuingTokenNetworkDataDeviceLanguageKi IssuingTokenNetworkDataDeviceLanguage = "ki"
	IssuingTokenNetworkDataDeviceLanguageKj IssuingTokenNetworkDataDeviceLanguage = "kj"
	IssuingTokenNetworkDataDeviceLanguageKk IssuingTokenNetworkDataDeviceLanguage = "kk"
	IssuingTokenNetworkDataDeviceLanguageKl IssuingTokenNetworkDataDeviceLanguage = "kl"
	IssuingTokenNetworkDataDeviceLanguageKm IssuingTokenNetworkDataDeviceLanguage = "km"
	IssuingTokenNetworkDataDeviceLanguageKn IssuingTokenNetworkDataDeviceLanguage = "kn"
	IssuingTokenNetworkDataDeviceLanguageKO IssuingTokenNetworkDataDeviceLanguage = "ko"
	IssuingTokenNetworkDataDeviceLanguageKr IssuingTokenNetworkDataDeviceLanguage = "kr"
	IssuingTokenNetworkDataDeviceLanguageKs IssuingTokenNetworkDataDeviceLanguage = "ks"
	IssuingTokenNetworkDataDeviceLanguageKu IssuingTokenNetworkDataDeviceLanguage = "ku"
	IssuingTokenNetworkDataDeviceLanguageKv IssuingTokenNetworkDataDeviceLanguage = "kv"
	IssuingTokenNetworkDataDeviceLanguageKw IssuingTokenNetworkDataDeviceLanguage = "kw"
	IssuingTokenNetworkDataDeviceLanguageKy IssuingTokenNetworkDataDeviceLanguage = "ky"
	IssuingTokenNetworkDataDeviceLanguageLa IssuingTokenNetworkDataDeviceLanguage = "la"
	IssuingTokenNetworkDataDeviceLanguageLb IssuingTokenNetworkDataDeviceLanguage = "lb"
	IssuingTokenNetworkDataDeviceLanguageLg IssuingTokenNetworkDataDeviceLanguage = "lg"
	IssuingTokenNetworkDataDeviceLanguageLi IssuingTokenNetworkDataDeviceLanguage = "li"
	IssuingTokenNetworkDataDeviceLanguageLn IssuingTokenNetworkDataDeviceLanguage = "ln"
	IssuingTokenNetworkDataDeviceLanguageLo IssuingTokenNetworkDataDeviceLanguage = "lo"
	IssuingTokenNetworkDataDeviceLanguageLT IssuingTokenNetworkDataDeviceLanguage = "lt"
	IssuingTokenNetworkDataDeviceLanguageLu IssuingTokenNetworkDataDeviceLanguage = "lu"
	IssuingTokenNetworkDataDeviceLanguageLV IssuingTokenNetworkDataDeviceLanguage = "lv"
	IssuingTokenNetworkDataDeviceLanguageMg IssuingTokenNetworkDataDeviceLanguage = "mg"
	IssuingTokenNetworkDataDeviceLanguageMh IssuingTokenNetworkDataDeviceLanguage = "mh"
	IssuingTokenNetworkDataDeviceLanguageMi IssuingTokenNetworkDataDeviceLanguage = "mi"
	IssuingTokenNetworkDataDeviceLanguageMk IssuingTokenNetworkDataDeviceLanguage = "mk"
	IssuingTokenNetworkDataDeviceLanguageMl IssuingTokenNetworkDataDeviceLanguage = "ml"
	IssuingTokenNetworkDataDeviceLanguageMn IssuingTokenNetworkDataDeviceLanguage = "mn"
	IssuingTokenNetworkDataDeviceLanguageMr IssuingTokenNetworkDataDeviceLanguage = "mr"
	IssuingTokenNetworkDataDeviceLanguageMS IssuingTokenNetworkDataDeviceLanguage = "ms"
	IssuingTokenNetworkDataDeviceLanguageMT IssuingTokenNetworkDataDeviceLanguage = "mt"
	IssuingTokenNetworkDataDeviceLanguageMy IssuingTokenNetworkDataDeviceLanguage = "my"
	IssuingTokenNetworkDataDeviceLanguageNa IssuingTokenNetworkDataDeviceLanguage = "na"
	IssuingTokenNetworkDataDeviceLanguageNB IssuingTokenNetworkDataDeviceLanguage = "nb"
	IssuingTokenNetworkDataDeviceLanguageNd IssuingTokenNetworkDataDeviceLanguage = "nd"
	IssuingTokenNetworkDataDeviceLanguageNe IssuingTokenNetworkDataDeviceLanguage = "ne"
	IssuingTokenNetworkDataDeviceLanguageNg IssuingTokenNetworkDataDeviceLanguage = "ng"
	IssuingTokenNetworkDataDeviceLanguageNL IssuingTokenNetworkDataDeviceLanguage = "nl"
	IssuingTokenNetworkDataDeviceLanguageNn IssuingTokenNetworkDataDeviceLanguage = "nn"
	IssuingTokenNetworkDataDeviceLanguageNo IssuingTokenNetworkDataDeviceLanguage = "no"
	IssuingTokenNetworkDataDeviceLanguageNr IssuingTokenNetworkDataDeviceLanguage = "nr"
	IssuingTokenNetworkDataDeviceLanguageNv IssuingTokenNetworkDataDeviceLanguage = "nv"
	IssuingTokenNetworkDataDeviceLanguageNy IssuingTokenNetworkDataDeviceLanguage = "ny"
	IssuingTokenNetworkDataDeviceLanguageOc IssuingTokenNetworkDataDeviceLanguage = "oc"
	IssuingTokenNetworkDataDeviceLanguageOj IssuingTokenNetworkDataDeviceLanguage = "oj"
	IssuingTokenNetworkDataDeviceLanguageOm IssuingTokenNetworkDataDeviceLanguage = "om"
	IssuingTokenNetworkDataDeviceLanguageOr IssuingTokenNetworkDataDeviceLanguage = "or"
	IssuingTokenNetworkDataDeviceLanguageOs IssuingTokenNetworkDataDeviceLanguage = "os"
	IssuingTokenNetworkDataDeviceLanguagePa IssuingTokenNetworkDataDeviceLanguage = "pa"
	IssuingTokenNetworkDataDeviceLanguagePi IssuingTokenNetworkDataDeviceLanguage = "pi"
	IssuingTokenNetworkDataDeviceLanguagePL IssuingTokenNetworkDataDeviceLanguage = "pl"
	IssuingTokenNetworkDataDeviceLanguagePs IssuingTokenNetworkDataDeviceLanguage = "ps"
	IssuingTokenNetworkDataDeviceLanguagePT IssuingTokenNetworkDataDeviceLanguage = "pt"
	IssuingTokenNetworkDataDeviceLanguageQu IssuingTokenNetworkDataDeviceLanguage = "qu"
	IssuingTokenNetworkDataDeviceLanguageRm IssuingTokenNetworkDataDeviceLanguage = "rm"
	IssuingTokenNetworkDataDeviceLanguageRn IssuingTokenNetworkDataDeviceLanguage = "rn"
	IssuingTokenNetworkDataDeviceLanguageRO IssuingTokenNetworkDataDeviceLanguage = "ro"
	IssuingTokenNetworkDataDeviceLanguageRU IssuingTokenNetworkDataDeviceLanguage = "ru"
	IssuingTokenNetworkDataDeviceLanguageRw IssuingTokenNetworkDataDeviceLanguage = "rw"
	IssuingTokenNetworkDataDeviceLanguageSa IssuingTokenNetworkDataDeviceLanguage = "sa"
	IssuingTokenNetworkDataDeviceLanguageSc IssuingTokenNetworkDataDeviceLanguage = "sc"
	IssuingTokenNetworkDataDeviceLanguageSd IssuingTokenNetworkDataDeviceLanguage = "sd"
	IssuingTokenNetworkDataDeviceLanguageSe IssuingTokenNetworkDataDeviceLanguage = "se"
	IssuingTokenNetworkDataDeviceLanguageSg IssuingTokenNetworkDataDeviceLanguage = "sg"
	IssuingTokenNetworkDataDeviceLanguageSi IssuingTokenNetworkDataDeviceLanguage = "si"
	IssuingTokenNetworkDataDeviceLanguageSK IssuingTokenNetworkDataDeviceLanguage = "sk"
	IssuingTokenNetworkDataDeviceLanguageSL IssuingTokenNetworkDataDeviceLanguage = "sl"
	IssuingTokenNetworkDataDeviceLanguageSm IssuingTokenNetworkDataDeviceLanguage = "sm"
	IssuingTokenNetworkDataDeviceLanguageSn IssuingTokenNetworkDataDeviceLanguage = "sn"
	IssuingTokenNetworkDataDeviceLanguageSo IssuingTokenNetworkDataDeviceLanguage = "so"
	IssuingTokenNetworkDataDeviceLanguageSq IssuingTokenNetworkDataDeviceLanguage = "sq"
	IssuingTokenNetworkDataDeviceLanguageSr IssuingTokenNetworkDataDeviceLanguage = "sr"
	IssuingTokenNetworkDataDeviceLanguageSs IssuingTokenNetworkDataDeviceLanguage = "ss"
	IssuingTokenNetworkDataDeviceLanguageSt IssuingTokenNetworkDataDeviceLanguage = "st"
	IssuingTokenNetworkDataDeviceLanguageSu IssuingTokenNetworkDataDeviceLanguage = "su"
	IssuingTokenNetworkDataDeviceLanguageSV IssuingTokenNetworkDataDeviceLanguage = "sv"
	IssuingTokenNetworkDataDeviceLanguageSw IssuingTokenNetworkDataDeviceLanguage = "sw"
	IssuingTokenNetworkDataDeviceLanguageTa IssuingTokenNetworkDataDeviceLanguage = "ta"
	IssuingTokenNetworkDataDeviceLanguageTe IssuingTokenNetworkDataDeviceLanguage = "te"
	IssuingTokenNetworkDataDeviceLanguageTg IssuingTokenNetworkDataDeviceLanguage = "tg"
	IssuingTokenNetworkDataDeviceLanguageTH IssuingTokenNetworkDataDeviceLanguage = "th"
	IssuingTokenNetworkDataDeviceLanguageTi IssuingTokenNetworkDataDeviceLanguage = "ti"
	IssuingTokenNetworkDataDeviceLanguageTk IssuingTokenNetworkDataDeviceLanguage = "tk"
	IssuingTokenNetworkDataDeviceLanguageTl IssuingTokenNetworkDataDeviceLanguage = "tl"
	IssuingTokenNetworkDataDeviceLanguageTn IssuingTokenNetworkDataDeviceLanguage = "tn"
	IssuingTokenNetworkDataDeviceLanguageTo IssuingTokenNetworkDataDeviceLanguage = "to"
	IssuingTokenNetworkDataDeviceLanguageTR IssuingTokenNetworkDataDeviceLanguage = "tr"
	IssuingTokenNetworkDataDeviceLanguageTS IssuingTokenNetworkDataDeviceLanguage = "ts"
	IssuingTokenNetworkDataDeviceLanguageTt IssuingTokenNetworkDataDeviceLanguage = "tt"
	IssuingTokenNetworkDataDeviceLanguageTw IssuingTokenNetworkDataDeviceLanguage = "tw"
	IssuingTokenNetworkDataDeviceLanguageTy IssuingTokenNetworkDataDeviceLanguage = "ty"
	IssuingTokenNetworkDataDeviceLanguageUg IssuingTokenNetworkDataDeviceLanguage = "ug"
	IssuingTokenNetworkDataDeviceLanguageUk IssuingTokenNetworkDataDeviceLanguage = "uk"
	IssuingTokenNetworkDataDeviceLanguageUr IssuingTokenNetworkDataDeviceLanguage = "ur"
	IssuingTokenNetworkDataDeviceLanguageUz IssuingTokenNetworkDataDeviceLanguage = "uz"
	IssuingTokenNetworkDataDeviceLanguageVe IssuingTokenNetworkDataDeviceLanguage = "ve"
	IssuingTokenNetworkDataDeviceLanguageVI IssuingTokenNetworkDataDeviceLanguage = "vi"
	IssuingTokenNetworkDataDeviceLanguageVo IssuingTokenNetworkDataDeviceLanguage = "vo"
	IssuingTokenNetworkDataDeviceLanguageWa IssuingTokenNetworkDataDeviceLanguage = "wa"
	IssuingTokenNetworkDataDeviceLanguageWo IssuingTokenNetworkDataDeviceLanguage = "wo"
	IssuingTokenNetworkDataDeviceLanguageXh IssuingTokenNetworkDataDeviceLanguage = "xh"
	IssuingTokenNetworkDataDeviceLanguageYi IssuingTokenNetworkDataDeviceLanguage = "yi"
	IssuingTokenNetworkDataDeviceLanguageYo IssuingTokenNetworkDataDeviceLanguage = "yo"
	IssuingTokenNetworkDataDeviceLanguageZa IssuingTokenNetworkDataDeviceLanguage = "za"
	IssuingTokenNetworkDataDeviceLanguageZH IssuingTokenNetworkDataDeviceLanguage = "zh"
	IssuingTokenNetworkDataDeviceLanguageZu IssuingTokenNetworkDataDeviceLanguage = "zu"
)

// The type of device used for tokenization.
type IssuingTokenNetworkDataDeviceType string

// List of values that IssuingTokenNetworkDataDeviceType can take
const (
	IssuingTokenNetworkDataDeviceTypeOther IssuingTokenNetworkDataDeviceType = "other"
	IssuingTokenNetworkDataDeviceTypePhone IssuingTokenNetworkDataDeviceType = "phone"
	IssuingTokenNetworkDataDeviceTypeWatch IssuingTokenNetworkDataDeviceType = "watch"
)

// The network that the token is associated with. An additional hash is included with a name matching this value, containing tokenization data specific to the card network.
type IssuingTokenNetworkDataType string

// List of values that IssuingTokenNetworkDataType can take
const (
	IssuingTokenNetworkDataTypeMastercard IssuingTokenNetworkDataType = "mastercard"
	IssuingTokenNetworkDataTypeVisa       IssuingTokenNetworkDataType = "visa"
)

// The network's recommendation to Stripe for this token activation request.
type IssuingTokenNetworkDataVisaTokenDecisionRecommendation string

// List of values that IssuingTokenNetworkDataVisaTokenDecisionRecommendation can take
const (
	IssuingTokenNetworkDataVisaTokenDecisionRecommendationApprove         IssuingTokenNetworkDataVisaTokenDecisionRecommendation = "approve"
	IssuingTokenNetworkDataVisaTokenDecisionRecommendationDecline         IssuingTokenNetworkDataVisaTokenDecisionRecommendation = "decline"
	IssuingTokenNetworkDataVisaTokenDecisionRecommendationRecommendIDAndV IssuingTokenNetworkDataVisaTokenDecisionRecommendation = "recommend_id_and_v"
)

// The method used for tokenizing a card.
type IssuingTokenNetworkDataWalletProviderCardNumberSource string

// List of values that IssuingTokenNetworkDataWalletProviderCardNumberSource can take
const (
	IssuingTokenNetworkDataWalletProviderCardNumberSourceApp    IssuingTokenNetworkDataWalletProviderCardNumberSource = "app"
	IssuingTokenNetworkDataWalletProviderCardNumberSourceManual IssuingTokenNetworkDataWalletProviderCardNumberSource = "manual"
	IssuingTokenNetworkDataWalletProviderCardNumberSourceOnFile IssuingTokenNetworkDataWalletProviderCardNumberSource = "on_file"
	IssuingTokenNetworkDataWalletProviderCardNumberSourceOther  IssuingTokenNetworkDataWalletProviderCardNumberSource = "other"
)

// The reasons for suggested tokenization given by the card network.
type IssuingTokenNetworkDataWalletProviderReasonCode string

// List of values that IssuingTokenNetworkDataWalletProviderReasonCode can take
const (
	IssuingTokenNetworkDataWalletProviderReasonCodeAccountCardTooNew                       IssuingTokenNetworkDataWalletProviderReasonCode = "account_card_too_new"
	IssuingTokenNetworkDataWalletProviderReasonCodeAccountRecentlyChanged                  IssuingTokenNetworkDataWalletProviderReasonCode = "account_recently_changed"
	IssuingTokenNetworkDataWalletProviderReasonCodeAccountTooNew                           IssuingTokenNetworkDataWalletProviderReasonCode = "account_too_new"
	IssuingTokenNetworkDataWalletProviderReasonCodeAccountTooNewSinceLaunch                IssuingTokenNetworkDataWalletProviderReasonCode = "account_too_new_since_launch"
	IssuingTokenNetworkDataWalletProviderReasonCodeAdditionalDevice                        IssuingTokenNetworkDataWalletProviderReasonCode = "additional_device"
	IssuingTokenNetworkDataWalletProviderReasonCodeDataExpired                             IssuingTokenNetworkDataWalletProviderReasonCode = "data_expired"
	IssuingTokenNetworkDataWalletProviderReasonCodeDeferIDVDecision                        IssuingTokenNetworkDataWalletProviderReasonCode = "defer_id_v_decision"
	IssuingTokenNetworkDataWalletProviderReasonCodeDeviceRecentlyLost                      IssuingTokenNetworkDataWalletProviderReasonCode = "device_recently_lost"
	IssuingTokenNetworkDataWalletProviderReasonCodeGoodActivityHistory                     IssuingTokenNetworkDataWalletProviderReasonCode = "good_activity_history"
	IssuingTokenNetworkDataWalletProviderReasonCodeHasSuspendedTokens                      IssuingTokenNetworkDataWalletProviderReasonCode = "has_suspended_tokens"
	IssuingTokenNetworkDataWalletProviderReasonCodeHighRisk                                IssuingTokenNetworkDataWalletProviderReasonCode = "high_risk"
	IssuingTokenNetworkDataWalletProviderReasonCodeInactiveAccount                         IssuingTokenNetworkDataWalletProviderReasonCode = "inactive_account"
	IssuingTokenNetworkDataWalletProviderReasonCodeLongAccountTenure                       IssuingTokenNetworkDataWalletProviderReasonCode = "long_account_tenure"
	IssuingTokenNetworkDataWalletProviderReasonCodeLowAccountScore                         IssuingTokenNetworkDataWalletProviderReasonCode = "low_account_score"
	IssuingTokenNetworkDataWalletProviderReasonCodeLowDeviceScore                          IssuingTokenNetworkDataWalletProviderReasonCode = "low_device_score"
	IssuingTokenNetworkDataWalletProviderReasonCodeLowPhoneNumberScore                     IssuingTokenNetworkDataWalletProviderReasonCode = "low_phone_number_score"
	IssuingTokenNetworkDataWalletProviderReasonCodeNetworkServiceError                     IssuingTokenNetworkDataWalletProviderReasonCode = "network_service_error"
	IssuingTokenNetworkDataWalletProviderReasonCodeOutsideHomeTerritory                    IssuingTokenNetworkDataWalletProviderReasonCode = "outside_home_territory"
	IssuingTokenNetworkDataWalletProviderReasonCodeProvisioningCardholderMismatch          IssuingTokenNetworkDataWalletProviderReasonCode = "provisioning_cardholder_mismatch"
	IssuingTokenNetworkDataWalletProviderReasonCodeProvisioningDeviceAndCardholderMismatch IssuingTokenNetworkDataWalletProviderReasonCode = "provisioning_device_and_cardholder_mismatch"
	IssuingTokenNetworkDataWalletProviderReasonCodeProvisioningDeviceMismatch              IssuingTokenNetworkDataWalletProviderReasonCode = "provisioning_device_mismatch"
	IssuingTokenNetworkDataWalletProviderReasonCodeSameDeviceNoPriorAuthentication         IssuingTokenNetworkDataWalletProviderReasonCode = "same_device_no_prior_authentication"
	IssuingTokenNetworkDataWalletProviderReasonCodeSameDeviceSuccessfulPriorAuthentication IssuingTokenNetworkDataWalletProviderReasonCode = "same_device_successful_prior_authentication"
	IssuingTokenNetworkDataWalletProviderReasonCodeSoftwareUpdate                          IssuingTokenNetworkDataWalletProviderReasonCode = "software_update"
	IssuingTokenNetworkDataWalletProviderReasonCodeSuspiciousActivity                      IssuingTokenNetworkDataWalletProviderReasonCode = "suspicious_activity"
	IssuingTokenNetworkDataWalletProviderReasonCodeTooManyDifferentCardholders             IssuingTokenNetworkDataWalletProviderReasonCode = "too_many_different_cardholders"
	IssuingTokenNetworkDataWalletProviderReasonCodeTooManyRecentAttempts                   IssuingTokenNetworkDataWalletProviderReasonCode = "too_many_recent_attempts"
	IssuingTokenNetworkDataWalletProviderReasonCodeTooManyRecentTokens                     IssuingTokenNetworkDataWalletProviderReasonCode = "too_many_recent_tokens"
)

// The recommendation on responding to the tokenization request.
type IssuingTokenNetworkDataWalletProviderSuggestedDecision string

// List of values that IssuingTokenNetworkDataWalletProviderSuggestedDecision can take
const (
	IssuingTokenNetworkDataWalletProviderSuggestedDecisionApprove     IssuingTokenNetworkDataWalletProviderSuggestedDecision = "approve"
	IssuingTokenNetworkDataWalletProviderSuggestedDecisionDecline     IssuingTokenNetworkDataWalletProviderSuggestedDecision = "decline"
	IssuingTokenNetworkDataWalletProviderSuggestedDecisionRequireAuth IssuingTokenNetworkDataWalletProviderSuggestedDecision = "require_auth"
)

// The decision made during token provisioning.
type IssuingTokenProvisioningDecision string

// List of values that IssuingTokenProvisioningDecision can take
const (
	IssuingTokenProvisioningDecisionApprove              IssuingTokenProvisioningDecision = "approve"
	IssuingTokenProvisioningDecisionApprovePendingIDAndV IssuingTokenProvisioningDecision = "approve_pending_id_and_v"
	IssuingTokenProvisioningDecisionDecline              IssuingTokenProvisioningDecision = "decline"
)

// The usage state of the token.
type IssuingTokenStatus string

// List of values that IssuingTokenStatus can take
const (
	IssuingTokenStatusActive    IssuingTokenStatus = "active"
	IssuingTokenStatusDeleted   IssuingTokenStatus = "deleted"
	IssuingTokenStatusRequested IssuingTokenStatus = "requested"
	IssuingTokenStatusSuspended IssuingTokenStatus = "suspended"
)

// The type of the token, indicating how it is used.
type IssuingTokenTokenType string

// List of values that IssuingTokenTokenType can take
const (
	IssuingTokenTokenTypeCardOnFile               IssuingTokenTokenType = "card_on_file"
	IssuingTokenTokenTypeCloudBased               IssuingTokenTokenType = "cloud_based"
	IssuingTokenTokenTypeCommercePlatform         IssuingTokenTokenType = "commerce_platform"
	IssuingTokenTokenTypeCommercialVirtualAccount IssuingTokenTokenType = "commercial_virtual_account"
	IssuingTokenTokenTypeSecureElement            IssuingTokenTokenType = "secure_element"
	IssuingTokenTokenTypeStaticCredential         IssuingTokenTokenType = "static_credential"
)

// The digital wallet for this token, if one was used.
type IssuingTokenWalletProvider string

// List of values that IssuingTokenWalletProvider can take
const (
	IssuingTokenWalletProviderApplePay   IssuingTokenWalletProvider = "apple_pay"
	IssuingTokenWalletProviderGooglePay  IssuingTokenWalletProvider = "google_pay"
	IssuingTokenWalletProviderSamsungPay IssuingTokenWalletProvider = "samsung_pay"
)

// Lists all Issuing Token objects for a given card.
type IssuingTokenListParams struct {
	ListParams `form:"*"`
	// The Issuing card identifier to list tokens for.
	Card *string `form:"card" json:"card"`
	// Only return Issuing tokens that were created during the given date interval.
	Created *int64 `form:"created" json:"created,omitempty"`
	// Only return Issuing tokens that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created" json:"-"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Select Issuing tokens with the given status.
	Status *string `form:"status" json:"status,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *IssuingTokenListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an Issuing Token object.
type IssuingTokenParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Specifies which status the token should be updated to.
	Status *string `form:"status" json:"status,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *IssuingTokenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an Issuing Token object.
type IssuingTokenRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *IssuingTokenRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Attempts to update the specified Issuing Token object to the status specified.
type IssuingTokenUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Specifies which status the token should be updated to.
	Status *string `form:"status" json:"status"`
}

// AddExpand appends a new field to expand.
func (p *IssuingTokenUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type IssuingTokenNetworkDataDevice struct {
	// An obfuscated ID derived from the device ID.
	DeviceFingerprint string `json:"device_fingerprint,omitempty"`
	// An identifier for the device used during wallet provisioning.
	DeviceID string `json:"device_id,omitempty"`
	// The IP address of the device at provisioning time.
	IPAddress string `json:"ip_address,omitempty"`
	// The ISO 639-1 language code of the device associated with the tokenization request.
	Language IssuingTokenNetworkDataDeviceLanguage `json:"language,omitempty"`
	// The geographic latitude/longitude coordinates of the device at provisioning time. The format is [+-]decimal/[+-]decimal.
	Location string `json:"location,omitempty"`
	// The name of the device used for tokenization.
	Name string `json:"name,omitempty"`
	// The phone number of the device used for tokenization.
	PhoneNumber string `json:"phone_number,omitempty"`
	// The type of device used for tokenization.
	Type IssuingTokenNetworkDataDeviceType `json:"type,omitempty"`
}
type IssuingTokenNetworkDataMastercard struct {
	// A unique reference ID from MasterCard to represent the card account number.
	CardReferenceID string `json:"card_reference_id,omitempty"`
	// The network-unique identifier for the token.
	TokenReferenceID string `json:"token_reference_id"`
	// The ID of the entity requesting tokenization, specific to MasterCard.
	TokenRequestorID string `json:"token_requestor_id"`
	// The name of the entity requesting tokenization, if known. This is directly provided from MasterCard.
	TokenRequestorName string `json:"token_requestor_name,omitempty"`
}
type IssuingTokenNetworkDataVisa struct {
	// A unique reference ID from Visa to represent the card account number.
	CardReferenceID string `json:"card_reference_id"`
	// The network's recommendation to Stripe for this token activation request.
	TokenDecisionRecommendation IssuingTokenNetworkDataVisaTokenDecisionRecommendation `json:"token_decision_recommendation,omitempty"`
	// The network-unique identifier for the token.
	TokenReferenceID string `json:"token_reference_id"`
	// The ID of the entity requesting tokenization, specific to Visa.
	TokenRequestorID string `json:"token_requestor_id"`
	// Degree of risk associated with the token between `01` and `99`, with higher number indicating higher risk. A `00` value indicates the token was not scored by Visa.
	TokenRiskScore string `json:"token_risk_score,omitempty"`
}
type IssuingTokenNetworkDataWalletProviderCardholderAddress struct {
	// The street address of the cardholder tokenizing the card.
	Line1 string `json:"line1"`
	// The postal code of the cardholder tokenizing the card.
	PostalCode string `json:"postal_code"`
}
type IssuingTokenNetworkDataWalletProvider struct {
	// The wallet provider-given account ID of the digital wallet the token belongs to.
	AccountID string `json:"account_id,omitempty"`
	// An evaluation on the trustworthiness of the wallet account between 1 and 5. A higher score indicates more trustworthy.
	AccountTrustScore int64                                                   `json:"account_trust_score,omitempty"`
	CardholderAddress *IssuingTokenNetworkDataWalletProviderCardholderAddress `json:"cardholder_address,omitempty"`
	// The name of the cardholder tokenizing the card.
	CardholderName string `json:"cardholder_name,omitempty"`
	// The method used for tokenizing a card.
	CardNumberSource IssuingTokenNetworkDataWalletProviderCardNumberSource `json:"card_number_source,omitempty"`
	// An evaluation on the trustworthiness of the device. A higher score indicates more trustworthy.
	DeviceTrustScore int64 `json:"device_trust_score,omitempty"`
	// The hashed email address of the cardholder's account with the wallet provider.
	HashedAccountEmailAddress string `json:"hashed_account_email_address,omitempty"`
	// The reasons for suggested tokenization given by the card network.
	ReasonCodes []IssuingTokenNetworkDataWalletProviderReasonCode `json:"reason_codes,omitempty"`
	// The recommendation on responding to the tokenization request.
	SuggestedDecision IssuingTokenNetworkDataWalletProviderSuggestedDecision `json:"suggested_decision,omitempty"`
	// The version of the standard for mapping reason codes followed by the wallet provider.
	SuggestedDecisionVersion string `json:"suggested_decision_version,omitempty"`
}
type IssuingTokenNetworkData struct {
	Device     *IssuingTokenNetworkDataDevice     `json:"device,omitempty"`
	Mastercard *IssuingTokenNetworkDataMastercard `json:"mastercard,omitempty"`
	// The network that the token is associated with. An additional hash is included with a name matching this value, containing tokenization data specific to the card network.
	Type           IssuingTokenNetworkDataType            `json:"type"`
	Visa           *IssuingTokenNetworkDataVisa           `json:"visa,omitempty"`
	WalletProvider *IssuingTokenNetworkDataWalletProvider `json:"wallet_provider,omitempty"`
}

// An issuing token object is created when an issued card is added to a digital wallet. As a [card issuer](https://docs.stripe.com/issuing), you can [view and manage these tokens](https://docs.stripe.com/issuing/controls/token-management) through Stripe.
type IssuingToken struct {
	APIResource
	// Card associated with this token.
	Card *IssuingCard `json:"card"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The hashed ID derived from the device ID from the card network associated with the token.
	DeviceFingerprint string `json:"device_fingerprint"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The last four digits of the token.
	Last4 string `json:"last4,omitempty"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// The token service provider / card network associated with the token.
	Network     IssuingTokenNetwork      `json:"network"`
	NetworkData *IssuingTokenNetworkData `json:"network_data,omitempty"`
	// Time at which the token was last updated by the card network. Measured in seconds since the Unix epoch.
	NetworkUpdatedAt int64 `json:"network_updated_at"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The decision made during token provisioning.
	ProvisioningDecision IssuingTokenProvisioningDecision `json:"provisioning_decision,omitempty"`
	// The usage state of the token.
	Status IssuingTokenStatus `json:"status"`
	// The type of the token, indicating how it is used.
	TokenType IssuingTokenTokenType `json:"token_type,omitempty"`
	// The digital wallet for this token, if one was used.
	WalletProvider IssuingTokenWalletProvider `json:"wallet_provider,omitempty"`
}

// IssuingTokenList is a list of Tokens as retrieved from a list endpoint.
type IssuingTokenList struct {
	APIResource
	ListMeta
	Data []*IssuingToken `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingToken.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingToken) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingToken IssuingToken
	var v issuingToken
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingToken(v)
	return nil
}
