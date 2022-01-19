//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The specified behavior after the purchase is complete.
type PaymentLinkAfterCompletionType string

// List of values that PaymentLinkAfterCompletionType can take
const (
	PaymentLinkAfterCompletionTypeHostedConfirmation PaymentLinkAfterCompletionType = "hosted_confirmation"
	PaymentLinkAfterCompletionTypeRedirect           PaymentLinkAfterCompletionType = "redirect"
)

// Configuration for collecting the customer's billing address.
type PaymentLinkBillingAddressCollection string

// List of values that PaymentLinkBillingAddressCollection can take
const (
	PaymentLinkBillingAddressCollectionAuto     PaymentLinkBillingAddressCollection = "auto"
	PaymentLinkBillingAddressCollectionRequired PaymentLinkBillingAddressCollection = "required"
)

// The list of payment method types that customers can use. When `null`, your [payment methods settings](https://dashboard.stripe.com/settings/payment_methods) will be used.
type PaymentLinkPaymentMethodType string

// List of values that PaymentLinkPaymentMethodType can take
const (
	PaymentLinkPaymentMethodTypeCard PaymentLinkPaymentMethodType = "card"
)

// An array of two-letter ISO country codes representing which countries Checkout should provide as options for shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SD, SY, UM, VI`.
type PaymentLinkShippingAddressCollectionAllowedCountry string

// List of values that PaymentLinkShippingAddressCollectionAllowedCountry can take
const (
	PaymentLinkShippingAddressCollectionAllowedCountryAC PaymentLinkShippingAddressCollectionAllowedCountry = "AC"
	PaymentLinkShippingAddressCollectionAllowedCountryAD PaymentLinkShippingAddressCollectionAllowedCountry = "AD"
	PaymentLinkShippingAddressCollectionAllowedCountryAE PaymentLinkShippingAddressCollectionAllowedCountry = "AE"
	PaymentLinkShippingAddressCollectionAllowedCountryAF PaymentLinkShippingAddressCollectionAllowedCountry = "AF"
	PaymentLinkShippingAddressCollectionAllowedCountryAG PaymentLinkShippingAddressCollectionAllowedCountry = "AG"
	PaymentLinkShippingAddressCollectionAllowedCountryAI PaymentLinkShippingAddressCollectionAllowedCountry = "AI"
	PaymentLinkShippingAddressCollectionAllowedCountryAL PaymentLinkShippingAddressCollectionAllowedCountry = "AL"
	PaymentLinkShippingAddressCollectionAllowedCountryAM PaymentLinkShippingAddressCollectionAllowedCountry = "AM"
	PaymentLinkShippingAddressCollectionAllowedCountryAO PaymentLinkShippingAddressCollectionAllowedCountry = "AO"
	PaymentLinkShippingAddressCollectionAllowedCountryAQ PaymentLinkShippingAddressCollectionAllowedCountry = "AQ"
	PaymentLinkShippingAddressCollectionAllowedCountryAR PaymentLinkShippingAddressCollectionAllowedCountry = "AR"
	PaymentLinkShippingAddressCollectionAllowedCountryAT PaymentLinkShippingAddressCollectionAllowedCountry = "AT"
	PaymentLinkShippingAddressCollectionAllowedCountryAU PaymentLinkShippingAddressCollectionAllowedCountry = "AU"
	PaymentLinkShippingAddressCollectionAllowedCountryAW PaymentLinkShippingAddressCollectionAllowedCountry = "AW"
	PaymentLinkShippingAddressCollectionAllowedCountryAX PaymentLinkShippingAddressCollectionAllowedCountry = "AX"
	PaymentLinkShippingAddressCollectionAllowedCountryAZ PaymentLinkShippingAddressCollectionAllowedCountry = "AZ"
	PaymentLinkShippingAddressCollectionAllowedCountryBA PaymentLinkShippingAddressCollectionAllowedCountry = "BA"
	PaymentLinkShippingAddressCollectionAllowedCountryBB PaymentLinkShippingAddressCollectionAllowedCountry = "BB"
	PaymentLinkShippingAddressCollectionAllowedCountryBD PaymentLinkShippingAddressCollectionAllowedCountry = "BD"
	PaymentLinkShippingAddressCollectionAllowedCountryBE PaymentLinkShippingAddressCollectionAllowedCountry = "BE"
	PaymentLinkShippingAddressCollectionAllowedCountryBF PaymentLinkShippingAddressCollectionAllowedCountry = "BF"
	PaymentLinkShippingAddressCollectionAllowedCountryBG PaymentLinkShippingAddressCollectionAllowedCountry = "BG"
	PaymentLinkShippingAddressCollectionAllowedCountryBH PaymentLinkShippingAddressCollectionAllowedCountry = "BH"
	PaymentLinkShippingAddressCollectionAllowedCountryBI PaymentLinkShippingAddressCollectionAllowedCountry = "BI"
	PaymentLinkShippingAddressCollectionAllowedCountryBJ PaymentLinkShippingAddressCollectionAllowedCountry = "BJ"
	PaymentLinkShippingAddressCollectionAllowedCountryBL PaymentLinkShippingAddressCollectionAllowedCountry = "BL"
	PaymentLinkShippingAddressCollectionAllowedCountryBM PaymentLinkShippingAddressCollectionAllowedCountry = "BM"
	PaymentLinkShippingAddressCollectionAllowedCountryBN PaymentLinkShippingAddressCollectionAllowedCountry = "BN"
	PaymentLinkShippingAddressCollectionAllowedCountryBO PaymentLinkShippingAddressCollectionAllowedCountry = "BO"
	PaymentLinkShippingAddressCollectionAllowedCountryBQ PaymentLinkShippingAddressCollectionAllowedCountry = "BQ"
	PaymentLinkShippingAddressCollectionAllowedCountryBR PaymentLinkShippingAddressCollectionAllowedCountry = "BR"
	PaymentLinkShippingAddressCollectionAllowedCountryBS PaymentLinkShippingAddressCollectionAllowedCountry = "BS"
	PaymentLinkShippingAddressCollectionAllowedCountryBT PaymentLinkShippingAddressCollectionAllowedCountry = "BT"
	PaymentLinkShippingAddressCollectionAllowedCountryBV PaymentLinkShippingAddressCollectionAllowedCountry = "BV"
	PaymentLinkShippingAddressCollectionAllowedCountryBW PaymentLinkShippingAddressCollectionAllowedCountry = "BW"
	PaymentLinkShippingAddressCollectionAllowedCountryBY PaymentLinkShippingAddressCollectionAllowedCountry = "BY"
	PaymentLinkShippingAddressCollectionAllowedCountryBZ PaymentLinkShippingAddressCollectionAllowedCountry = "BZ"
	PaymentLinkShippingAddressCollectionAllowedCountryCA PaymentLinkShippingAddressCollectionAllowedCountry = "CA"
	PaymentLinkShippingAddressCollectionAllowedCountryCD PaymentLinkShippingAddressCollectionAllowedCountry = "CD"
	PaymentLinkShippingAddressCollectionAllowedCountryCF PaymentLinkShippingAddressCollectionAllowedCountry = "CF"
	PaymentLinkShippingAddressCollectionAllowedCountryCG PaymentLinkShippingAddressCollectionAllowedCountry = "CG"
	PaymentLinkShippingAddressCollectionAllowedCountryCH PaymentLinkShippingAddressCollectionAllowedCountry = "CH"
	PaymentLinkShippingAddressCollectionAllowedCountryCI PaymentLinkShippingAddressCollectionAllowedCountry = "CI"
	PaymentLinkShippingAddressCollectionAllowedCountryCK PaymentLinkShippingAddressCollectionAllowedCountry = "CK"
	PaymentLinkShippingAddressCollectionAllowedCountryCL PaymentLinkShippingAddressCollectionAllowedCountry = "CL"
	PaymentLinkShippingAddressCollectionAllowedCountryCM PaymentLinkShippingAddressCollectionAllowedCountry = "CM"
	PaymentLinkShippingAddressCollectionAllowedCountryCN PaymentLinkShippingAddressCollectionAllowedCountry = "CN"
	PaymentLinkShippingAddressCollectionAllowedCountryCO PaymentLinkShippingAddressCollectionAllowedCountry = "CO"
	PaymentLinkShippingAddressCollectionAllowedCountryCR PaymentLinkShippingAddressCollectionAllowedCountry = "CR"
	PaymentLinkShippingAddressCollectionAllowedCountryCV PaymentLinkShippingAddressCollectionAllowedCountry = "CV"
	PaymentLinkShippingAddressCollectionAllowedCountryCW PaymentLinkShippingAddressCollectionAllowedCountry = "CW"
	PaymentLinkShippingAddressCollectionAllowedCountryCY PaymentLinkShippingAddressCollectionAllowedCountry = "CY"
	PaymentLinkShippingAddressCollectionAllowedCountryCZ PaymentLinkShippingAddressCollectionAllowedCountry = "CZ"
	PaymentLinkShippingAddressCollectionAllowedCountryDE PaymentLinkShippingAddressCollectionAllowedCountry = "DE"
	PaymentLinkShippingAddressCollectionAllowedCountryDJ PaymentLinkShippingAddressCollectionAllowedCountry = "DJ"
	PaymentLinkShippingAddressCollectionAllowedCountryDK PaymentLinkShippingAddressCollectionAllowedCountry = "DK"
	PaymentLinkShippingAddressCollectionAllowedCountryDM PaymentLinkShippingAddressCollectionAllowedCountry = "DM"
	PaymentLinkShippingAddressCollectionAllowedCountryDO PaymentLinkShippingAddressCollectionAllowedCountry = "DO"
	PaymentLinkShippingAddressCollectionAllowedCountryDZ PaymentLinkShippingAddressCollectionAllowedCountry = "DZ"
	PaymentLinkShippingAddressCollectionAllowedCountryEC PaymentLinkShippingAddressCollectionAllowedCountry = "EC"
	PaymentLinkShippingAddressCollectionAllowedCountryEE PaymentLinkShippingAddressCollectionAllowedCountry = "EE"
	PaymentLinkShippingAddressCollectionAllowedCountryEG PaymentLinkShippingAddressCollectionAllowedCountry = "EG"
	PaymentLinkShippingAddressCollectionAllowedCountryEH PaymentLinkShippingAddressCollectionAllowedCountry = "EH"
	PaymentLinkShippingAddressCollectionAllowedCountryER PaymentLinkShippingAddressCollectionAllowedCountry = "ER"
	PaymentLinkShippingAddressCollectionAllowedCountryES PaymentLinkShippingAddressCollectionAllowedCountry = "ES"
	PaymentLinkShippingAddressCollectionAllowedCountryET PaymentLinkShippingAddressCollectionAllowedCountry = "ET"
	PaymentLinkShippingAddressCollectionAllowedCountryFI PaymentLinkShippingAddressCollectionAllowedCountry = "FI"
	PaymentLinkShippingAddressCollectionAllowedCountryFJ PaymentLinkShippingAddressCollectionAllowedCountry = "FJ"
	PaymentLinkShippingAddressCollectionAllowedCountryFK PaymentLinkShippingAddressCollectionAllowedCountry = "FK"
	PaymentLinkShippingAddressCollectionAllowedCountryFO PaymentLinkShippingAddressCollectionAllowedCountry = "FO"
	PaymentLinkShippingAddressCollectionAllowedCountryFR PaymentLinkShippingAddressCollectionAllowedCountry = "FR"
	PaymentLinkShippingAddressCollectionAllowedCountryGA PaymentLinkShippingAddressCollectionAllowedCountry = "GA"
	PaymentLinkShippingAddressCollectionAllowedCountryGB PaymentLinkShippingAddressCollectionAllowedCountry = "GB"
	PaymentLinkShippingAddressCollectionAllowedCountryGD PaymentLinkShippingAddressCollectionAllowedCountry = "GD"
	PaymentLinkShippingAddressCollectionAllowedCountryGE PaymentLinkShippingAddressCollectionAllowedCountry = "GE"
	PaymentLinkShippingAddressCollectionAllowedCountryGF PaymentLinkShippingAddressCollectionAllowedCountry = "GF"
	PaymentLinkShippingAddressCollectionAllowedCountryGG PaymentLinkShippingAddressCollectionAllowedCountry = "GG"
	PaymentLinkShippingAddressCollectionAllowedCountryGH PaymentLinkShippingAddressCollectionAllowedCountry = "GH"
	PaymentLinkShippingAddressCollectionAllowedCountryGI PaymentLinkShippingAddressCollectionAllowedCountry = "GI"
	PaymentLinkShippingAddressCollectionAllowedCountryGL PaymentLinkShippingAddressCollectionAllowedCountry = "GL"
	PaymentLinkShippingAddressCollectionAllowedCountryGM PaymentLinkShippingAddressCollectionAllowedCountry = "GM"
	PaymentLinkShippingAddressCollectionAllowedCountryGN PaymentLinkShippingAddressCollectionAllowedCountry = "GN"
	PaymentLinkShippingAddressCollectionAllowedCountryGP PaymentLinkShippingAddressCollectionAllowedCountry = "GP"
	PaymentLinkShippingAddressCollectionAllowedCountryGQ PaymentLinkShippingAddressCollectionAllowedCountry = "GQ"
	PaymentLinkShippingAddressCollectionAllowedCountryGR PaymentLinkShippingAddressCollectionAllowedCountry = "GR"
	PaymentLinkShippingAddressCollectionAllowedCountryGS PaymentLinkShippingAddressCollectionAllowedCountry = "GS"
	PaymentLinkShippingAddressCollectionAllowedCountryGT PaymentLinkShippingAddressCollectionAllowedCountry = "GT"
	PaymentLinkShippingAddressCollectionAllowedCountryGU PaymentLinkShippingAddressCollectionAllowedCountry = "GU"
	PaymentLinkShippingAddressCollectionAllowedCountryGW PaymentLinkShippingAddressCollectionAllowedCountry = "GW"
	PaymentLinkShippingAddressCollectionAllowedCountryGY PaymentLinkShippingAddressCollectionAllowedCountry = "GY"
	PaymentLinkShippingAddressCollectionAllowedCountryHK PaymentLinkShippingAddressCollectionAllowedCountry = "HK"
	PaymentLinkShippingAddressCollectionAllowedCountryHN PaymentLinkShippingAddressCollectionAllowedCountry = "HN"
	PaymentLinkShippingAddressCollectionAllowedCountryHR PaymentLinkShippingAddressCollectionAllowedCountry = "HR"
	PaymentLinkShippingAddressCollectionAllowedCountryHT PaymentLinkShippingAddressCollectionAllowedCountry = "HT"
	PaymentLinkShippingAddressCollectionAllowedCountryHU PaymentLinkShippingAddressCollectionAllowedCountry = "HU"
	PaymentLinkShippingAddressCollectionAllowedCountryID PaymentLinkShippingAddressCollectionAllowedCountry = "ID"
	PaymentLinkShippingAddressCollectionAllowedCountryIE PaymentLinkShippingAddressCollectionAllowedCountry = "IE"
	PaymentLinkShippingAddressCollectionAllowedCountryIL PaymentLinkShippingAddressCollectionAllowedCountry = "IL"
	PaymentLinkShippingAddressCollectionAllowedCountryIM PaymentLinkShippingAddressCollectionAllowedCountry = "IM"
	PaymentLinkShippingAddressCollectionAllowedCountryIN PaymentLinkShippingAddressCollectionAllowedCountry = "IN"
	PaymentLinkShippingAddressCollectionAllowedCountryIO PaymentLinkShippingAddressCollectionAllowedCountry = "IO"
	PaymentLinkShippingAddressCollectionAllowedCountryIQ PaymentLinkShippingAddressCollectionAllowedCountry = "IQ"
	PaymentLinkShippingAddressCollectionAllowedCountryIS PaymentLinkShippingAddressCollectionAllowedCountry = "IS"
	PaymentLinkShippingAddressCollectionAllowedCountryIT PaymentLinkShippingAddressCollectionAllowedCountry = "IT"
	PaymentLinkShippingAddressCollectionAllowedCountryJE PaymentLinkShippingAddressCollectionAllowedCountry = "JE"
	PaymentLinkShippingAddressCollectionAllowedCountryJM PaymentLinkShippingAddressCollectionAllowedCountry = "JM"
	PaymentLinkShippingAddressCollectionAllowedCountryJO PaymentLinkShippingAddressCollectionAllowedCountry = "JO"
	PaymentLinkShippingAddressCollectionAllowedCountryJP PaymentLinkShippingAddressCollectionAllowedCountry = "JP"
	PaymentLinkShippingAddressCollectionAllowedCountryKE PaymentLinkShippingAddressCollectionAllowedCountry = "KE"
	PaymentLinkShippingAddressCollectionAllowedCountryKG PaymentLinkShippingAddressCollectionAllowedCountry = "KG"
	PaymentLinkShippingAddressCollectionAllowedCountryKH PaymentLinkShippingAddressCollectionAllowedCountry = "KH"
	PaymentLinkShippingAddressCollectionAllowedCountryKI PaymentLinkShippingAddressCollectionAllowedCountry = "KI"
	PaymentLinkShippingAddressCollectionAllowedCountryKM PaymentLinkShippingAddressCollectionAllowedCountry = "KM"
	PaymentLinkShippingAddressCollectionAllowedCountryKN PaymentLinkShippingAddressCollectionAllowedCountry = "KN"
	PaymentLinkShippingAddressCollectionAllowedCountryKR PaymentLinkShippingAddressCollectionAllowedCountry = "KR"
	PaymentLinkShippingAddressCollectionAllowedCountryKW PaymentLinkShippingAddressCollectionAllowedCountry = "KW"
	PaymentLinkShippingAddressCollectionAllowedCountryKY PaymentLinkShippingAddressCollectionAllowedCountry = "KY"
	PaymentLinkShippingAddressCollectionAllowedCountryKZ PaymentLinkShippingAddressCollectionAllowedCountry = "KZ"
	PaymentLinkShippingAddressCollectionAllowedCountryLA PaymentLinkShippingAddressCollectionAllowedCountry = "LA"
	PaymentLinkShippingAddressCollectionAllowedCountryLB PaymentLinkShippingAddressCollectionAllowedCountry = "LB"
	PaymentLinkShippingAddressCollectionAllowedCountryLC PaymentLinkShippingAddressCollectionAllowedCountry = "LC"
	PaymentLinkShippingAddressCollectionAllowedCountryLI PaymentLinkShippingAddressCollectionAllowedCountry = "LI"
	PaymentLinkShippingAddressCollectionAllowedCountryLK PaymentLinkShippingAddressCollectionAllowedCountry = "LK"
	PaymentLinkShippingAddressCollectionAllowedCountryLR PaymentLinkShippingAddressCollectionAllowedCountry = "LR"
	PaymentLinkShippingAddressCollectionAllowedCountryLS PaymentLinkShippingAddressCollectionAllowedCountry = "LS"
	PaymentLinkShippingAddressCollectionAllowedCountryLT PaymentLinkShippingAddressCollectionAllowedCountry = "LT"
	PaymentLinkShippingAddressCollectionAllowedCountryLU PaymentLinkShippingAddressCollectionAllowedCountry = "LU"
	PaymentLinkShippingAddressCollectionAllowedCountryLV PaymentLinkShippingAddressCollectionAllowedCountry = "LV"
	PaymentLinkShippingAddressCollectionAllowedCountryLY PaymentLinkShippingAddressCollectionAllowedCountry = "LY"
	PaymentLinkShippingAddressCollectionAllowedCountryMA PaymentLinkShippingAddressCollectionAllowedCountry = "MA"
	PaymentLinkShippingAddressCollectionAllowedCountryMC PaymentLinkShippingAddressCollectionAllowedCountry = "MC"
	PaymentLinkShippingAddressCollectionAllowedCountryMD PaymentLinkShippingAddressCollectionAllowedCountry = "MD"
	PaymentLinkShippingAddressCollectionAllowedCountryME PaymentLinkShippingAddressCollectionAllowedCountry = "ME"
	PaymentLinkShippingAddressCollectionAllowedCountryMF PaymentLinkShippingAddressCollectionAllowedCountry = "MF"
	PaymentLinkShippingAddressCollectionAllowedCountryMG PaymentLinkShippingAddressCollectionAllowedCountry = "MG"
	PaymentLinkShippingAddressCollectionAllowedCountryMK PaymentLinkShippingAddressCollectionAllowedCountry = "MK"
	PaymentLinkShippingAddressCollectionAllowedCountryML PaymentLinkShippingAddressCollectionAllowedCountry = "ML"
	PaymentLinkShippingAddressCollectionAllowedCountryMM PaymentLinkShippingAddressCollectionAllowedCountry = "MM"
	PaymentLinkShippingAddressCollectionAllowedCountryMN PaymentLinkShippingAddressCollectionAllowedCountry = "MN"
	PaymentLinkShippingAddressCollectionAllowedCountryMO PaymentLinkShippingAddressCollectionAllowedCountry = "MO"
	PaymentLinkShippingAddressCollectionAllowedCountryMQ PaymentLinkShippingAddressCollectionAllowedCountry = "MQ"
	PaymentLinkShippingAddressCollectionAllowedCountryMR PaymentLinkShippingAddressCollectionAllowedCountry = "MR"
	PaymentLinkShippingAddressCollectionAllowedCountryMS PaymentLinkShippingAddressCollectionAllowedCountry = "MS"
	PaymentLinkShippingAddressCollectionAllowedCountryMT PaymentLinkShippingAddressCollectionAllowedCountry = "MT"
	PaymentLinkShippingAddressCollectionAllowedCountryMU PaymentLinkShippingAddressCollectionAllowedCountry = "MU"
	PaymentLinkShippingAddressCollectionAllowedCountryMV PaymentLinkShippingAddressCollectionAllowedCountry = "MV"
	PaymentLinkShippingAddressCollectionAllowedCountryMW PaymentLinkShippingAddressCollectionAllowedCountry = "MW"
	PaymentLinkShippingAddressCollectionAllowedCountryMX PaymentLinkShippingAddressCollectionAllowedCountry = "MX"
	PaymentLinkShippingAddressCollectionAllowedCountryMY PaymentLinkShippingAddressCollectionAllowedCountry = "MY"
	PaymentLinkShippingAddressCollectionAllowedCountryMZ PaymentLinkShippingAddressCollectionAllowedCountry = "MZ"
	PaymentLinkShippingAddressCollectionAllowedCountryNA PaymentLinkShippingAddressCollectionAllowedCountry = "NA"
	PaymentLinkShippingAddressCollectionAllowedCountryNC PaymentLinkShippingAddressCollectionAllowedCountry = "NC"
	PaymentLinkShippingAddressCollectionAllowedCountryNE PaymentLinkShippingAddressCollectionAllowedCountry = "NE"
	PaymentLinkShippingAddressCollectionAllowedCountryNG PaymentLinkShippingAddressCollectionAllowedCountry = "NG"
	PaymentLinkShippingAddressCollectionAllowedCountryNI PaymentLinkShippingAddressCollectionAllowedCountry = "NI"
	PaymentLinkShippingAddressCollectionAllowedCountryNL PaymentLinkShippingAddressCollectionAllowedCountry = "NL"
	PaymentLinkShippingAddressCollectionAllowedCountryNO PaymentLinkShippingAddressCollectionAllowedCountry = "NO"
	PaymentLinkShippingAddressCollectionAllowedCountryNP PaymentLinkShippingAddressCollectionAllowedCountry = "NP"
	PaymentLinkShippingAddressCollectionAllowedCountryNR PaymentLinkShippingAddressCollectionAllowedCountry = "NR"
	PaymentLinkShippingAddressCollectionAllowedCountryNU PaymentLinkShippingAddressCollectionAllowedCountry = "NU"
	PaymentLinkShippingAddressCollectionAllowedCountryNZ PaymentLinkShippingAddressCollectionAllowedCountry = "NZ"
	PaymentLinkShippingAddressCollectionAllowedCountryOM PaymentLinkShippingAddressCollectionAllowedCountry = "OM"
	PaymentLinkShippingAddressCollectionAllowedCountryPA PaymentLinkShippingAddressCollectionAllowedCountry = "PA"
	PaymentLinkShippingAddressCollectionAllowedCountryPE PaymentLinkShippingAddressCollectionAllowedCountry = "PE"
	PaymentLinkShippingAddressCollectionAllowedCountryPF PaymentLinkShippingAddressCollectionAllowedCountry = "PF"
	PaymentLinkShippingAddressCollectionAllowedCountryPG PaymentLinkShippingAddressCollectionAllowedCountry = "PG"
	PaymentLinkShippingAddressCollectionAllowedCountryPH PaymentLinkShippingAddressCollectionAllowedCountry = "PH"
	PaymentLinkShippingAddressCollectionAllowedCountryPK PaymentLinkShippingAddressCollectionAllowedCountry = "PK"
	PaymentLinkShippingAddressCollectionAllowedCountryPL PaymentLinkShippingAddressCollectionAllowedCountry = "PL"
	PaymentLinkShippingAddressCollectionAllowedCountryPM PaymentLinkShippingAddressCollectionAllowedCountry = "PM"
	PaymentLinkShippingAddressCollectionAllowedCountryPN PaymentLinkShippingAddressCollectionAllowedCountry = "PN"
	PaymentLinkShippingAddressCollectionAllowedCountryPR PaymentLinkShippingAddressCollectionAllowedCountry = "PR"
	PaymentLinkShippingAddressCollectionAllowedCountryPS PaymentLinkShippingAddressCollectionAllowedCountry = "PS"
	PaymentLinkShippingAddressCollectionAllowedCountryPT PaymentLinkShippingAddressCollectionAllowedCountry = "PT"
	PaymentLinkShippingAddressCollectionAllowedCountryPY PaymentLinkShippingAddressCollectionAllowedCountry = "PY"
	PaymentLinkShippingAddressCollectionAllowedCountryQA PaymentLinkShippingAddressCollectionAllowedCountry = "QA"
	PaymentLinkShippingAddressCollectionAllowedCountryRE PaymentLinkShippingAddressCollectionAllowedCountry = "RE"
	PaymentLinkShippingAddressCollectionAllowedCountryRO PaymentLinkShippingAddressCollectionAllowedCountry = "RO"
	PaymentLinkShippingAddressCollectionAllowedCountryRS PaymentLinkShippingAddressCollectionAllowedCountry = "RS"
	PaymentLinkShippingAddressCollectionAllowedCountryRU PaymentLinkShippingAddressCollectionAllowedCountry = "RU"
	PaymentLinkShippingAddressCollectionAllowedCountryRW PaymentLinkShippingAddressCollectionAllowedCountry = "RW"
	PaymentLinkShippingAddressCollectionAllowedCountrySA PaymentLinkShippingAddressCollectionAllowedCountry = "SA"
	PaymentLinkShippingAddressCollectionAllowedCountrySB PaymentLinkShippingAddressCollectionAllowedCountry = "SB"
	PaymentLinkShippingAddressCollectionAllowedCountrySC PaymentLinkShippingAddressCollectionAllowedCountry = "SC"
	PaymentLinkShippingAddressCollectionAllowedCountrySE PaymentLinkShippingAddressCollectionAllowedCountry = "SE"
	PaymentLinkShippingAddressCollectionAllowedCountrySG PaymentLinkShippingAddressCollectionAllowedCountry = "SG"
	PaymentLinkShippingAddressCollectionAllowedCountrySH PaymentLinkShippingAddressCollectionAllowedCountry = "SH"
	PaymentLinkShippingAddressCollectionAllowedCountrySI PaymentLinkShippingAddressCollectionAllowedCountry = "SI"
	PaymentLinkShippingAddressCollectionAllowedCountrySJ PaymentLinkShippingAddressCollectionAllowedCountry = "SJ"
	PaymentLinkShippingAddressCollectionAllowedCountrySK PaymentLinkShippingAddressCollectionAllowedCountry = "SK"
	PaymentLinkShippingAddressCollectionAllowedCountrySL PaymentLinkShippingAddressCollectionAllowedCountry = "SL"
	PaymentLinkShippingAddressCollectionAllowedCountrySM PaymentLinkShippingAddressCollectionAllowedCountry = "SM"
	PaymentLinkShippingAddressCollectionAllowedCountrySN PaymentLinkShippingAddressCollectionAllowedCountry = "SN"
	PaymentLinkShippingAddressCollectionAllowedCountrySO PaymentLinkShippingAddressCollectionAllowedCountry = "SO"
	PaymentLinkShippingAddressCollectionAllowedCountrySR PaymentLinkShippingAddressCollectionAllowedCountry = "SR"
	PaymentLinkShippingAddressCollectionAllowedCountrySS PaymentLinkShippingAddressCollectionAllowedCountry = "SS"
	PaymentLinkShippingAddressCollectionAllowedCountryST PaymentLinkShippingAddressCollectionAllowedCountry = "ST"
	PaymentLinkShippingAddressCollectionAllowedCountrySV PaymentLinkShippingAddressCollectionAllowedCountry = "SV"
	PaymentLinkShippingAddressCollectionAllowedCountrySX PaymentLinkShippingAddressCollectionAllowedCountry = "SX"
	PaymentLinkShippingAddressCollectionAllowedCountrySZ PaymentLinkShippingAddressCollectionAllowedCountry = "SZ"
	PaymentLinkShippingAddressCollectionAllowedCountryTA PaymentLinkShippingAddressCollectionAllowedCountry = "TA"
	PaymentLinkShippingAddressCollectionAllowedCountryTC PaymentLinkShippingAddressCollectionAllowedCountry = "TC"
	PaymentLinkShippingAddressCollectionAllowedCountryTD PaymentLinkShippingAddressCollectionAllowedCountry = "TD"
	PaymentLinkShippingAddressCollectionAllowedCountryTF PaymentLinkShippingAddressCollectionAllowedCountry = "TF"
	PaymentLinkShippingAddressCollectionAllowedCountryTG PaymentLinkShippingAddressCollectionAllowedCountry = "TG"
	PaymentLinkShippingAddressCollectionAllowedCountryTH PaymentLinkShippingAddressCollectionAllowedCountry = "TH"
	PaymentLinkShippingAddressCollectionAllowedCountryTJ PaymentLinkShippingAddressCollectionAllowedCountry = "TJ"
	PaymentLinkShippingAddressCollectionAllowedCountryTK PaymentLinkShippingAddressCollectionAllowedCountry = "TK"
	PaymentLinkShippingAddressCollectionAllowedCountryTL PaymentLinkShippingAddressCollectionAllowedCountry = "TL"
	PaymentLinkShippingAddressCollectionAllowedCountryTM PaymentLinkShippingAddressCollectionAllowedCountry = "TM"
	PaymentLinkShippingAddressCollectionAllowedCountryTN PaymentLinkShippingAddressCollectionAllowedCountry = "TN"
	PaymentLinkShippingAddressCollectionAllowedCountryTO PaymentLinkShippingAddressCollectionAllowedCountry = "TO"
	PaymentLinkShippingAddressCollectionAllowedCountryTR PaymentLinkShippingAddressCollectionAllowedCountry = "TR"
	PaymentLinkShippingAddressCollectionAllowedCountryTT PaymentLinkShippingAddressCollectionAllowedCountry = "TT"
	PaymentLinkShippingAddressCollectionAllowedCountryTV PaymentLinkShippingAddressCollectionAllowedCountry = "TV"
	PaymentLinkShippingAddressCollectionAllowedCountryTW PaymentLinkShippingAddressCollectionAllowedCountry = "TW"
	PaymentLinkShippingAddressCollectionAllowedCountryTZ PaymentLinkShippingAddressCollectionAllowedCountry = "TZ"
	PaymentLinkShippingAddressCollectionAllowedCountryUA PaymentLinkShippingAddressCollectionAllowedCountry = "UA"
	PaymentLinkShippingAddressCollectionAllowedCountryUG PaymentLinkShippingAddressCollectionAllowedCountry = "UG"
	PaymentLinkShippingAddressCollectionAllowedCountryUS PaymentLinkShippingAddressCollectionAllowedCountry = "US"
	PaymentLinkShippingAddressCollectionAllowedCountryUY PaymentLinkShippingAddressCollectionAllowedCountry = "UY"
	PaymentLinkShippingAddressCollectionAllowedCountryUZ PaymentLinkShippingAddressCollectionAllowedCountry = "UZ"
	PaymentLinkShippingAddressCollectionAllowedCountryVA PaymentLinkShippingAddressCollectionAllowedCountry = "VA"
	PaymentLinkShippingAddressCollectionAllowedCountryVC PaymentLinkShippingAddressCollectionAllowedCountry = "VC"
	PaymentLinkShippingAddressCollectionAllowedCountryVE PaymentLinkShippingAddressCollectionAllowedCountry = "VE"
	PaymentLinkShippingAddressCollectionAllowedCountryVG PaymentLinkShippingAddressCollectionAllowedCountry = "VG"
	PaymentLinkShippingAddressCollectionAllowedCountryVN PaymentLinkShippingAddressCollectionAllowedCountry = "VN"
	PaymentLinkShippingAddressCollectionAllowedCountryVU PaymentLinkShippingAddressCollectionAllowedCountry = "VU"
	PaymentLinkShippingAddressCollectionAllowedCountryWF PaymentLinkShippingAddressCollectionAllowedCountry = "WF"
	PaymentLinkShippingAddressCollectionAllowedCountryWS PaymentLinkShippingAddressCollectionAllowedCountry = "WS"
	PaymentLinkShippingAddressCollectionAllowedCountryXK PaymentLinkShippingAddressCollectionAllowedCountry = "XK"
	PaymentLinkShippingAddressCollectionAllowedCountryYE PaymentLinkShippingAddressCollectionAllowedCountry = "YE"
	PaymentLinkShippingAddressCollectionAllowedCountryYT PaymentLinkShippingAddressCollectionAllowedCountry = "YT"
	PaymentLinkShippingAddressCollectionAllowedCountryZA PaymentLinkShippingAddressCollectionAllowedCountry = "ZA"
	PaymentLinkShippingAddressCollectionAllowedCountryZM PaymentLinkShippingAddressCollectionAllowedCountry = "ZM"
	PaymentLinkShippingAddressCollectionAllowedCountryZW PaymentLinkShippingAddressCollectionAllowedCountry = "ZW"
	PaymentLinkShippingAddressCollectionAllowedCountryZZ PaymentLinkShippingAddressCollectionAllowedCountry = "ZZ"
)

// Returns a list of your payment links.
type PaymentLinkListParams struct {
	ListParams `form:"*"`
	// Only return payment links that are active or inactive (e.g., pass `false` to list all inactive payment links).
	Active *bool `form:"active"`
}

// Configuration when `type=hosted_confirmation`.
type PaymentLinkAfterCompletionHostedConfirmationParams struct {
	// A custom message to display to the customer after the purchase is complete.
	CustomMessage *string `form:"custom_message"`
}

// Configuration when `type=redirect`.
type PaymentLinkAfterCompletionRedirectParams struct {
	// The URL the customer will be redirected to after the purchase is complete. You can embed `{CHECKOUT_SESSION_ID}` into the URL to have the `id` of the completed [checkout session](https://stripe.com/docs/api/checkout/sessions/object#checkout_session_object-id) included.
	URL *string `form:"url"`
}

// Behavior after the purchase is complete.
type PaymentLinkAfterCompletionParams struct {
	// Configuration when `type=hosted_confirmation`.
	HostedConfirmation *PaymentLinkAfterCompletionHostedConfirmationParams `form:"hosted_confirmation"`
	// Configuration when `type=redirect`.
	Redirect *PaymentLinkAfterCompletionRedirectParams `form:"redirect"`
	// The specified behavior after the purchase is complete. Either `redirect` or `hosted_confirmation`.
	Type *string `form:"type"`
}

// Configuration for automatic tax collection.
type PaymentLinkAutomaticTaxParams struct {
	// If `true`, tax will be calculated automatically using the customer's location.
	Enabled *bool `form:"enabled"`
}

// When set, provides configuration for this item's quantity to be adjusted by the customer during checkout.
type PaymentLinkLineItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative Integer.
	Enabled *bool `form:"enabled"`
	// The maximum quantity the customer can purchase. By default this value is 99. You can specify a value up to 99.
	Maximum *int64 `form:"maximum"`
	// The minimum quantity the customer can purchase. By default this value is 0. You can specify a value up to 98.
	Minimum *int64 `form:"minimum"`
}

// The line items representing what is being sold. Each line item represents an item being sold. Up to 20 line items are supported.
type PaymentLinkLineItemParams struct {
	// When set, provides configuration for this item's quantity to be adjusted by the customer during checkout.
	AdjustableQuantity *PaymentLinkLineItemAdjustableQuantityParams `form:"adjustable_quantity"`
	// The ID of an existing line item on the payment link.
	ID *string `form:"id"`
	// The ID of the [Price](https://stripe.com/docs/api/prices) or [Plan](https://stripe.com/docs/api/plans) object.
	Price *string `form:"price"`
	// The quantity of the line item being purchased. Only `1` is supported.
	Quantity *int64 `form:"quantity"`
}

// Configuration for collecting the customer's shipping address.
type PaymentLinkShippingAddressCollectionParams struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for
	// shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SD, SY, UM, VI`.
	AllowedCountries []*string `form:"allowed_countries"`
}

// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
type PaymentLinkSubscriptionDataParams struct {
	// Integer representing the number of trial period days before the customer is charged for the first time. Has to be at least 1.
	TrialPeriodDays *int64 `form:"trial_period_days"`
}

// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
type PaymentLinkTransferDataParams struct {
	// The amount that will be transferred automatically when a charge succeeds.
	Amount *int64 `form:"amount"`
	// If specified, successful charges will be attributed to the destination
	// account for tax reporting, and the funds from charges will be transferred
	// to the destination account. The ID of the resulting transfer will be
	// returned on the successful charge's `transfer` field.
	Destination *string `form:"destination"`
}

// Creates a payment link.
type PaymentLinkParams struct {
	Params `form:"*"`
	// Whether the payment link's `url` is active. If `false`, customers visiting the url will be redirected.
	Active *bool `form:"active"`
	// Behavior after the purchase is complete.
	AfterCompletion *PaymentLinkAfterCompletionParams `form:"after_completion"`
	// Enables user redeemable promotion codes.
	AllowPromotionCodes *bool `form:"allow_promotion_codes"`
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. Can only be applied when there are no line items with recurring prices.
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account. There must be at least 1 line item with a recurring price to use this field.
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// Configuration for automatic tax collection.
	AutomaticTax *PaymentLinkAutomaticTaxParams `form:"automatic_tax"`
	// Configuration for collecting the customer's billing address.
	BillingAddressCollection *string `form:"billing_address_collection"`
	// The line items representing what is being sold. Each line item represents an item being sold. Up to 20 line items are supported.
	LineItems []*PaymentLinkLineItemParams `form:"line_items"`
	// The account on behalf of which to charge.
	OnBehalfOf *string `form:"on_behalf_of"`
	// The list of payment method types (e.g., card) that customers can use. Only `card` is supported. Pass an empty string to enable automatic payment methods that use your [payment methods settings](https://dashboard.stripe.com/settings/payment_methods).
	PaymentMethodTypes []*string `form:"payment_method_types"`
	// Configuration for collecting the customer's shipping address.
	ShippingAddressCollection *PaymentLinkShippingAddressCollectionParams `form:"shipping_address_collection"`
	// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
	SubscriptionData *PaymentLinkSubscriptionDataParams `form:"subscription_data"`
	// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
	TransferData *PaymentLinkTransferDataParams `form:"transfer_data"`
}

// When retrieving a payment link, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type PaymentLinkListLineItemsParams struct {
	ListParams  `form:"*"`
	PaymentLink *string `form:"-"` // Included in URL
}
type PaymentLinkAfterCompletionHostedConfirmation struct {
	// The custom message that is displayed to the customer after the purchase is complete.
	CustomMessage string `json:"custom_message"`
}
type PaymentLinkAfterCompletionRedirect struct {
	// The `url` the customer will be redirected to after the purchase is complete.
	URL string `json:"url"`
}
type PaymentLinkAfterCompletion struct {
	HostedConfirmation *PaymentLinkAfterCompletionHostedConfirmation `json:"hosted_confirmation"`
	Redirect           *PaymentLinkAfterCompletionRedirect           `json:"redirect"`
	// The specified behavior after the purchase is complete.
	Type PaymentLinkAfterCompletionType `json:"type"`
}
type PaymentLinkAutomaticTax struct {
	// If `true`, tax will be calculated automatically using the customer's location.
	Enabled bool `json:"enabled"`
}

// Configuration for collecting the customer's shipping address.
type PaymentLinkShippingAddressCollection struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SD, SY, UM, VI`.
	AllowedCountries []PaymentLinkShippingAddressCollectionAllowedCountry `json:"allowed_countries"`
}

// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
type PaymentLinkSubscriptionData struct {
	// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
	TrialPeriodDays int64 `json:"trial_period_days"`
}

// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
type PaymentLinkTransferData struct {
	// The amount in %s that will be transferred to the destination account. By default, the entire amount is transferred to the destination.
	Amount int64 `json:"amount"`
	// The connected account receiving the transfer.
	Destination *Account `json:"destination"`
}

// A payment link allows you create payment pages through a url you can share with customers.
type PaymentLink struct {
	APIResource
	// Whether the payment link's `url` is active. If `false`, customers visiting the url will be redirected.
	Active          bool                        `json:"active"`
	AfterCompletion *PaymentLinkAfterCompletion `json:"after_completion"`
	// Whether user redeemable promotion codes are enabled.
	AllowPromotionCodes bool `json:"allow_promotion_codes"`
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account.
	ApplicationFeeAmount int64 `json:"application_fee_amount"`
	// This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account.
	ApplicationFeePercent float64                  `json:"application_fee_percent"`
	AutomaticTax          *PaymentLinkAutomaticTax `json:"automatic_tax"`
	// Configuration for collecting the customer's billing address.
	BillingAddressCollection PaymentLinkBillingAddressCollection `json:"billing_address_collection"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The line items representing what is being sold.
	LineItems *LineItemList `json:"line_items"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account on behalf of which to charge. See the [Connect documentation](https://support.stripe.com/questions/sending-invoices-on-behalf-of-connected-accounts) for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// The list of payment method types that customers can use. When `null`, your [payment methods settings](https://dashboard.stripe.com/settings/payment_methods) will be used.
	PaymentMethodTypes []PaymentLinkPaymentMethodType `json:"payment_method_types"`
	// Configuration for collecting the customer's shipping address.
	ShippingAddressCollection *PaymentLinkShippingAddressCollection `json:"shipping_address_collection"`
	// When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.
	SubscriptionData *PaymentLinkSubscriptionData `json:"subscription_data"`
	// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.
	TransferData *PaymentLinkTransferData `json:"transfer_data"`
	// The public url that can be shared with customers.
	URL string `json:"url"`
}

// PaymentLinkList is a list of PaymentLinks as retrieved from a list endpoint.
type PaymentLinkList struct {
	APIResource
	ListMeta
	Data []*PaymentLink `json:"data"`
}

// UnmarshalJSON handles deserialization of a PaymentLink.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PaymentLink) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type paymentLink PaymentLink
	var v paymentLink
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PaymentLink(v)
	return nil
}
