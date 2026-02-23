//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Type of the account referenced.
type CheckoutSessionAutomaticTaxLiabilityType string

// List of values that CheckoutSessionAutomaticTaxLiabilityType can take
const (
	CheckoutSessionAutomaticTaxLiabilityTypeAccount CheckoutSessionAutomaticTaxLiabilityType = "account"
	CheckoutSessionAutomaticTaxLiabilityTypeSelf    CheckoutSessionAutomaticTaxLiabilityType = "self"
)

// The status of the most recent automated tax calculation for this session.
type CheckoutSessionAutomaticTaxStatus string

// List of values that CheckoutSessionAutomaticTaxStatus can take
const (
	CheckoutSessionAutomaticTaxStatusComplete               CheckoutSessionAutomaticTaxStatus = "complete"
	CheckoutSessionAutomaticTaxStatusFailed                 CheckoutSessionAutomaticTaxStatus = "failed"
	CheckoutSessionAutomaticTaxStatusRequiresLocationInputs CheckoutSessionAutomaticTaxStatus = "requires_location_inputs"
)

// Describes whether Checkout should collect the customer's billing address. Defaults to `auto`.
type CheckoutSessionBillingAddressCollection string

// List of values that CheckoutSessionBillingAddressCollection can take
const (
	CheckoutSessionBillingAddressCollectionAuto     CheckoutSessionBillingAddressCollection = "auto"
	CheckoutSessionBillingAddressCollectionRequired CheckoutSessionBillingAddressCollection = "required"
)

// The border style for the Checkout Session. Must be one of `rounded`, `rectangular`, or `pill`.
type CheckoutSessionBrandingSettingsBorderStyle string

// List of values that CheckoutSessionBrandingSettingsBorderStyle can take
const (
	CheckoutSessionBrandingSettingsBorderStylePill        CheckoutSessionBrandingSettingsBorderStyle = "pill"
	CheckoutSessionBrandingSettingsBorderStyleRectangular CheckoutSessionBrandingSettingsBorderStyle = "rectangular"
	CheckoutSessionBrandingSettingsBorderStyleRounded     CheckoutSessionBrandingSettingsBorderStyle = "rounded"
)

// The type of image for the icon. Must be one of `file` or `url`.
type CheckoutSessionBrandingSettingsIconType string

// List of values that CheckoutSessionBrandingSettingsIconType can take
const (
	CheckoutSessionBrandingSettingsIconTypeFile CheckoutSessionBrandingSettingsIconType = "file"
	CheckoutSessionBrandingSettingsIconTypeURL  CheckoutSessionBrandingSettingsIconType = "url"
)

// The type of image for the logo. Must be one of `file` or `url`.
type CheckoutSessionBrandingSettingsLogoType string

// List of values that CheckoutSessionBrandingSettingsLogoType can take
const (
	CheckoutSessionBrandingSettingsLogoTypeFile CheckoutSessionBrandingSettingsLogoType = "file"
	CheckoutSessionBrandingSettingsLogoTypeURL  CheckoutSessionBrandingSettingsLogoType = "url"
)

// If `opt_in`, the customer consents to receiving promotional communications
// from the merchant about this Checkout Session.
type CheckoutSessionConsentPromotions string

// List of values that CheckoutSessionConsentPromotions can take
const (
	CheckoutSessionConsentPromotionsOptIn  CheckoutSessionConsentPromotions = "opt_in"
	CheckoutSessionConsentPromotionsOptOut CheckoutSessionConsentPromotions = "opt_out"
)

// If `accepted`, the customer in this Checkout Session has agreed to the merchant's terms of service.
type CheckoutSessionConsentTermsOfService string

// List of values that CheckoutSessionConsentTermsOfService can take
const (
	CheckoutSessionConsentTermsOfServiceAccepted CheckoutSessionConsentTermsOfService = "accepted"
)

// Determines the position and visibility of the payment method reuse agreement in the UI. When set to `auto`, Stripe's defaults will be used.
//
// When set to `hidden`, the payment method reuse agreement text will always be hidden in the UI.
type CheckoutSessionConsentCollectionPaymentMethodReuseAgreementPosition string

// List of values that CheckoutSessionConsentCollectionPaymentMethodReuseAgreementPosition can take
const (
	CheckoutSessionConsentCollectionPaymentMethodReuseAgreementPositionAuto   CheckoutSessionConsentCollectionPaymentMethodReuseAgreementPosition = "auto"
	CheckoutSessionConsentCollectionPaymentMethodReuseAgreementPositionHidden CheckoutSessionConsentCollectionPaymentMethodReuseAgreementPosition = "hidden"
)

// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
// Session will determine whether to display an option to opt into promotional communication
// from the merchant depending on the customer's locale. Only available to US merchants.
type CheckoutSessionConsentCollectionPromotions string

// List of values that CheckoutSessionConsentCollectionPromotions can take
const (
	CheckoutSessionConsentCollectionPromotionsAuto CheckoutSessionConsentCollectionPromotions = "auto"
	CheckoutSessionConsentCollectionPromotionsNone CheckoutSessionConsentCollectionPromotions = "none"
)

// If set to `required`, it requires customers to accept the terms of service before being able to pay.
type CheckoutSessionConsentCollectionTermsOfService string

// List of values that CheckoutSessionConsentCollectionTermsOfService can take
const (
	CheckoutSessionConsentCollectionTermsOfServiceNone     CheckoutSessionConsentCollectionTermsOfService = "none"
	CheckoutSessionConsentCollectionTermsOfServiceRequired CheckoutSessionConsentCollectionTermsOfService = "required"
)

// The type of the label.
type CheckoutSessionCustomFieldLabelType string

// List of values that CheckoutSessionCustomFieldLabelType can take
const (
	CheckoutSessionCustomFieldLabelTypeCustom CheckoutSessionCustomFieldLabelType = "custom"
)

// The type of the field.
type CheckoutSessionCustomFieldType string

// List of values that CheckoutSessionCustomFieldType can take
const (
	CheckoutSessionCustomFieldTypeDropdown CheckoutSessionCustomFieldType = "dropdown"
	CheckoutSessionCustomFieldTypeNumeric  CheckoutSessionCustomFieldType = "numeric"
	CheckoutSessionCustomFieldTypeText     CheckoutSessionCustomFieldType = "text"
)

// Configure whether a Checkout Session creates a Customer when the Checkout Session completes.
type CheckoutSessionCustomerCreation string

// List of values that CheckoutSessionCustomerCreation can take
const (
	CheckoutSessionCustomerCreationAlways     CheckoutSessionCustomerCreation = "always"
	CheckoutSessionCustomerCreationIfRequired CheckoutSessionCustomerCreation = "if_required"
)

// The customer's tax exempt status after a completed Checkout Session.
type CheckoutSessionCustomerDetailsTaxExempt string

// List of values that CheckoutSessionCustomerDetailsTaxExempt can take
const (
	CheckoutSessionCustomerDetailsTaxExemptExempt  CheckoutSessionCustomerDetailsTaxExempt = "exempt"
	CheckoutSessionCustomerDetailsTaxExemptNone    CheckoutSessionCustomerDetailsTaxExempt = "none"
	CheckoutSessionCustomerDetailsTaxExemptReverse CheckoutSessionCustomerDetailsTaxExempt = "reverse"
)

// The type of the tax ID, one of `ad_nrt`, `ar_cuit`, `eu_vat`, `bo_tin`, `br_cnpj`, `br_cpf`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `eu_oss_vat`, `hr_oib`, `pe_ruc`, `ro_tin`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, `vn_tin`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `no_voec`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `pl_nip`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `li_vat`, `lk_vat`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, `al_tin`, `bh_vat`, `kz_bin`, `ng_tin`, `om_vat`, `de_stn`, `ch_uid`, `tz_vat`, `uz_vat`, `uz_tin`, `md_vat`, `ma_vat`, `by_tin`, `ao_tin`, `bs_tin`, `bb_tin`, `cd_nif`, `mr_nif`, `me_pib`, `zw_tin`, `ba_tin`, `gn_nif`, `mk_vat`, `sr_fin`, `sn_ninea`, `am_tin`, `np_pan`, `tj_tin`, `ug_tin`, `zm_tin`, `kh_tin`, `aw_tin`, `az_tin`, `bd_bin`, `bj_ifu`, `et_tin`, `kg_tin`, `la_tin`, `cm_niu`, `cv_nif`, `bf_ifu`, or `unknown`
type CheckoutSessionCustomerDetailsTaxIDType string

// List of values that CheckoutSessionCustomerDetailsTaxIDType can take
const (
	CheckoutSessionCustomerDetailsTaxIDTypeADNRT    CheckoutSessionCustomerDetailsTaxIDType = "ad_nrt"
	CheckoutSessionCustomerDetailsTaxIDTypeAETRN    CheckoutSessionCustomerDetailsTaxIDType = "ae_trn"
	CheckoutSessionCustomerDetailsTaxIDTypeAlTin    CheckoutSessionCustomerDetailsTaxIDType = "al_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeAmTin    CheckoutSessionCustomerDetailsTaxIDType = "am_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeAoTin    CheckoutSessionCustomerDetailsTaxIDType = "ao_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeARCUIT   CheckoutSessionCustomerDetailsTaxIDType = "ar_cuit"
	CheckoutSessionCustomerDetailsTaxIDTypeAUABN    CheckoutSessionCustomerDetailsTaxIDType = "au_abn"
	CheckoutSessionCustomerDetailsTaxIDTypeAUARN    CheckoutSessionCustomerDetailsTaxIDType = "au_arn"
	CheckoutSessionCustomerDetailsTaxIDTypeAwTin    CheckoutSessionCustomerDetailsTaxIDType = "aw_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeAzTin    CheckoutSessionCustomerDetailsTaxIDType = "az_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeBaTin    CheckoutSessionCustomerDetailsTaxIDType = "ba_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeBbTin    CheckoutSessionCustomerDetailsTaxIDType = "bb_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeBdBin    CheckoutSessionCustomerDetailsTaxIDType = "bd_bin"
	CheckoutSessionCustomerDetailsTaxIDTypeBfIfu    CheckoutSessionCustomerDetailsTaxIDType = "bf_ifu"
	CheckoutSessionCustomerDetailsTaxIDTypeBGUIC    CheckoutSessionCustomerDetailsTaxIDType = "bg_uic"
	CheckoutSessionCustomerDetailsTaxIDTypeBhVAT    CheckoutSessionCustomerDetailsTaxIDType = "bh_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeBjIfu    CheckoutSessionCustomerDetailsTaxIDType = "bj_ifu"
	CheckoutSessionCustomerDetailsTaxIDTypeBOTIN    CheckoutSessionCustomerDetailsTaxIDType = "bo_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeBRCNPJ   CheckoutSessionCustomerDetailsTaxIDType = "br_cnpj"
	CheckoutSessionCustomerDetailsTaxIDTypeBRCPF    CheckoutSessionCustomerDetailsTaxIDType = "br_cpf"
	CheckoutSessionCustomerDetailsTaxIDTypeBsTin    CheckoutSessionCustomerDetailsTaxIDType = "bs_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeByTin    CheckoutSessionCustomerDetailsTaxIDType = "by_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeCABN     CheckoutSessionCustomerDetailsTaxIDType = "ca_bn"
	CheckoutSessionCustomerDetailsTaxIDTypeCAGSTHST CheckoutSessionCustomerDetailsTaxIDType = "ca_gst_hst"
	CheckoutSessionCustomerDetailsTaxIDTypeCAPSTBC  CheckoutSessionCustomerDetailsTaxIDType = "ca_pst_bc"
	CheckoutSessionCustomerDetailsTaxIDTypeCAPSTMB  CheckoutSessionCustomerDetailsTaxIDType = "ca_pst_mb"
	CheckoutSessionCustomerDetailsTaxIDTypeCAPSTSK  CheckoutSessionCustomerDetailsTaxIDType = "ca_pst_sk"
	CheckoutSessionCustomerDetailsTaxIDTypeCAQST    CheckoutSessionCustomerDetailsTaxIDType = "ca_qst"
	CheckoutSessionCustomerDetailsTaxIDTypeCdNif    CheckoutSessionCustomerDetailsTaxIDType = "cd_nif"
	CheckoutSessionCustomerDetailsTaxIDTypeCHUID    CheckoutSessionCustomerDetailsTaxIDType = "ch_uid"
	CheckoutSessionCustomerDetailsTaxIDTypeCHVAT    CheckoutSessionCustomerDetailsTaxIDType = "ch_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeCLTIN    CheckoutSessionCustomerDetailsTaxIDType = "cl_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeCmNiu    CheckoutSessionCustomerDetailsTaxIDType = "cm_niu"
	CheckoutSessionCustomerDetailsTaxIDTypeCNTIN    CheckoutSessionCustomerDetailsTaxIDType = "cn_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeCONIT    CheckoutSessionCustomerDetailsTaxIDType = "co_nit"
	CheckoutSessionCustomerDetailsTaxIDTypeCRTIN    CheckoutSessionCustomerDetailsTaxIDType = "cr_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeCvNif    CheckoutSessionCustomerDetailsTaxIDType = "cv_nif"
	CheckoutSessionCustomerDetailsTaxIDTypeDEStn    CheckoutSessionCustomerDetailsTaxIDType = "de_stn"
	CheckoutSessionCustomerDetailsTaxIDTypeDORCN    CheckoutSessionCustomerDetailsTaxIDType = "do_rcn"
	CheckoutSessionCustomerDetailsTaxIDTypeECRUC    CheckoutSessionCustomerDetailsTaxIDType = "ec_ruc"
	CheckoutSessionCustomerDetailsTaxIDTypeEGTIN    CheckoutSessionCustomerDetailsTaxIDType = "eg_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeESCIF    CheckoutSessionCustomerDetailsTaxIDType = "es_cif"
	CheckoutSessionCustomerDetailsTaxIDTypeETTin    CheckoutSessionCustomerDetailsTaxIDType = "et_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeEUOSSVAT CheckoutSessionCustomerDetailsTaxIDType = "eu_oss_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeEUVAT    CheckoutSessionCustomerDetailsTaxIDType = "eu_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeGBVAT    CheckoutSessionCustomerDetailsTaxIDType = "gb_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeGEVAT    CheckoutSessionCustomerDetailsTaxIDType = "ge_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeGnNif    CheckoutSessionCustomerDetailsTaxIDType = "gn_nif"
	CheckoutSessionCustomerDetailsTaxIDTypeHKBR     CheckoutSessionCustomerDetailsTaxIDType = "hk_br"
	CheckoutSessionCustomerDetailsTaxIDTypeHROIB    CheckoutSessionCustomerDetailsTaxIDType = "hr_oib"
	CheckoutSessionCustomerDetailsTaxIDTypeHUTIN    CheckoutSessionCustomerDetailsTaxIDType = "hu_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeIDNPWP   CheckoutSessionCustomerDetailsTaxIDType = "id_npwp"
	CheckoutSessionCustomerDetailsTaxIDTypeILVAT    CheckoutSessionCustomerDetailsTaxIDType = "il_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeINGST    CheckoutSessionCustomerDetailsTaxIDType = "in_gst"
	CheckoutSessionCustomerDetailsTaxIDTypeISVAT    CheckoutSessionCustomerDetailsTaxIDType = "is_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeJPCN     CheckoutSessionCustomerDetailsTaxIDType = "jp_cn"
	CheckoutSessionCustomerDetailsTaxIDTypeJPRN     CheckoutSessionCustomerDetailsTaxIDType = "jp_rn"
	CheckoutSessionCustomerDetailsTaxIDTypeJPTRN    CheckoutSessionCustomerDetailsTaxIDType = "jp_trn"
	CheckoutSessionCustomerDetailsTaxIDTypeKEPIN    CheckoutSessionCustomerDetailsTaxIDType = "ke_pin"
	CheckoutSessionCustomerDetailsTaxIDTypeKgTin    CheckoutSessionCustomerDetailsTaxIDType = "kg_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeKhTin    CheckoutSessionCustomerDetailsTaxIDType = "kh_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeKRBRN    CheckoutSessionCustomerDetailsTaxIDType = "kr_brn"
	CheckoutSessionCustomerDetailsTaxIDTypeKzBin    CheckoutSessionCustomerDetailsTaxIDType = "kz_bin"
	CheckoutSessionCustomerDetailsTaxIDTypeLaTin    CheckoutSessionCustomerDetailsTaxIDType = "la_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeLIUID    CheckoutSessionCustomerDetailsTaxIDType = "li_uid"
	CheckoutSessionCustomerDetailsTaxIDTypeLiVAT    CheckoutSessionCustomerDetailsTaxIDType = "li_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeLkVAT    CheckoutSessionCustomerDetailsTaxIDType = "lk_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeMaVAT    CheckoutSessionCustomerDetailsTaxIDType = "ma_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeMdVAT    CheckoutSessionCustomerDetailsTaxIDType = "md_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeMePib    CheckoutSessionCustomerDetailsTaxIDType = "me_pib"
	CheckoutSessionCustomerDetailsTaxIDTypeMkVAT    CheckoutSessionCustomerDetailsTaxIDType = "mk_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeMrNif    CheckoutSessionCustomerDetailsTaxIDType = "mr_nif"
	CheckoutSessionCustomerDetailsTaxIDTypeMXRFC    CheckoutSessionCustomerDetailsTaxIDType = "mx_rfc"
	CheckoutSessionCustomerDetailsTaxIDTypeMYFRP    CheckoutSessionCustomerDetailsTaxIDType = "my_frp"
	CheckoutSessionCustomerDetailsTaxIDTypeMYITN    CheckoutSessionCustomerDetailsTaxIDType = "my_itn"
	CheckoutSessionCustomerDetailsTaxIDTypeMYSST    CheckoutSessionCustomerDetailsTaxIDType = "my_sst"
	CheckoutSessionCustomerDetailsTaxIDTypeNgTin    CheckoutSessionCustomerDetailsTaxIDType = "ng_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeNOVAT    CheckoutSessionCustomerDetailsTaxIDType = "no_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeNOVOEC   CheckoutSessionCustomerDetailsTaxIDType = "no_voec"
	CheckoutSessionCustomerDetailsTaxIDTypeNpPan    CheckoutSessionCustomerDetailsTaxIDType = "np_pan"
	CheckoutSessionCustomerDetailsTaxIDTypeNZGST    CheckoutSessionCustomerDetailsTaxIDType = "nz_gst"
	CheckoutSessionCustomerDetailsTaxIDTypeOmVAT    CheckoutSessionCustomerDetailsTaxIDType = "om_vat"
	CheckoutSessionCustomerDetailsTaxIDTypePERUC    CheckoutSessionCustomerDetailsTaxIDType = "pe_ruc"
	CheckoutSessionCustomerDetailsTaxIDTypePHTIN    CheckoutSessionCustomerDetailsTaxIDType = "ph_tin"
	CheckoutSessionCustomerDetailsTaxIDTypePLNip    CheckoutSessionCustomerDetailsTaxIDType = "pl_nip"
	CheckoutSessionCustomerDetailsTaxIDTypeROTIN    CheckoutSessionCustomerDetailsTaxIDType = "ro_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeRSPIB    CheckoutSessionCustomerDetailsTaxIDType = "rs_pib"
	CheckoutSessionCustomerDetailsTaxIDTypeRUINN    CheckoutSessionCustomerDetailsTaxIDType = "ru_inn"
	CheckoutSessionCustomerDetailsTaxIDTypeRUKPP    CheckoutSessionCustomerDetailsTaxIDType = "ru_kpp"
	CheckoutSessionCustomerDetailsTaxIDTypeSAVAT    CheckoutSessionCustomerDetailsTaxIDType = "sa_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeSGGST    CheckoutSessionCustomerDetailsTaxIDType = "sg_gst"
	CheckoutSessionCustomerDetailsTaxIDTypeSGUEN    CheckoutSessionCustomerDetailsTaxIDType = "sg_uen"
	CheckoutSessionCustomerDetailsTaxIDTypeSITIN    CheckoutSessionCustomerDetailsTaxIDType = "si_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeSnNinea  CheckoutSessionCustomerDetailsTaxIDType = "sn_ninea"
	CheckoutSessionCustomerDetailsTaxIDTypeSrFin    CheckoutSessionCustomerDetailsTaxIDType = "sr_fin"
	CheckoutSessionCustomerDetailsTaxIDTypeSVNIT    CheckoutSessionCustomerDetailsTaxIDType = "sv_nit"
	CheckoutSessionCustomerDetailsTaxIDTypeTHVAT    CheckoutSessionCustomerDetailsTaxIDType = "th_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeTjTin    CheckoutSessionCustomerDetailsTaxIDType = "tj_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeTRTIN    CheckoutSessionCustomerDetailsTaxIDType = "tr_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeTWVAT    CheckoutSessionCustomerDetailsTaxIDType = "tw_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeTzVAT    CheckoutSessionCustomerDetailsTaxIDType = "tz_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeUAVAT    CheckoutSessionCustomerDetailsTaxIDType = "ua_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeUgTin    CheckoutSessionCustomerDetailsTaxIDType = "ug_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeUnknown  CheckoutSessionCustomerDetailsTaxIDType = "unknown"
	CheckoutSessionCustomerDetailsTaxIDTypeUSEIN    CheckoutSessionCustomerDetailsTaxIDType = "us_ein"
	CheckoutSessionCustomerDetailsTaxIDTypeUYRUC    CheckoutSessionCustomerDetailsTaxIDType = "uy_ruc"
	CheckoutSessionCustomerDetailsTaxIDTypeUzTin    CheckoutSessionCustomerDetailsTaxIDType = "uz_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeUzVAT    CheckoutSessionCustomerDetailsTaxIDType = "uz_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeVERIF    CheckoutSessionCustomerDetailsTaxIDType = "ve_rif"
	CheckoutSessionCustomerDetailsTaxIDTypeVNTIN    CheckoutSessionCustomerDetailsTaxIDType = "vn_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeZAVAT    CheckoutSessionCustomerDetailsTaxIDType = "za_vat"
	CheckoutSessionCustomerDetailsTaxIDTypeZmTin    CheckoutSessionCustomerDetailsTaxIDType = "zm_tin"
	CheckoutSessionCustomerDetailsTaxIDTypeZwTin    CheckoutSessionCustomerDetailsTaxIDType = "zw_tin"
)

// Type of the account referenced.
type CheckoutSessionInvoiceCreationInvoiceDataIssuerType string

// List of values that CheckoutSessionInvoiceCreationInvoiceDataIssuerType can take
const (
	CheckoutSessionInvoiceCreationInvoiceDataIssuerTypeAccount CheckoutSessionInvoiceCreationInvoiceDataIssuerType = "account"
	CheckoutSessionInvoiceCreationInvoiceDataIssuerTypeSelf    CheckoutSessionInvoiceCreationInvoiceDataIssuerType = "self"
)

// The mode of the Checkout Session.
type CheckoutSessionMode string

// List of values that CheckoutSessionMode can take
const (
	CheckoutSessionModePayment      CheckoutSessionMode = "payment"
	CheckoutSessionModeSetup        CheckoutSessionMode = "setup"
	CheckoutSessionModeSubscription CheckoutSessionMode = "subscription"
)

// Where the user is coming from. This informs the optimizations that are applied to the session.
type CheckoutSessionOriginContext string

// List of values that CheckoutSessionOriginContext can take
const (
	CheckoutSessionOriginContextMobileApp CheckoutSessionOriginContext = "mobile_app"
	CheckoutSessionOriginContextWeb       CheckoutSessionOriginContext = "web"
)

// Configure whether a Checkout Session should collect a payment method. Defaults to `always`.
type CheckoutSessionPaymentMethodCollection string

// List of values that CheckoutSessionPaymentMethodCollection can take
const (
	CheckoutSessionPaymentMethodCollectionAlways     CheckoutSessionPaymentMethodCollection = "always"
	CheckoutSessionPaymentMethodCollectionIfRequired CheckoutSessionPaymentMethodCollection = "if_required"
)

// List of Stripe products where this mandate can be selected automatically. Returned when the Session is in `setup` mode.
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultForInvoice      CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor = "invoice"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultForSubscription CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor = "subscription"
)

// Payment schedule for the mandate.
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleCombined CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "combined"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleInterval CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "interval"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleSporadic CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "sporadic"
)

// Transaction type of the mandate.
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsACSSDebitSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsACSSDebitSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsACSSDebitSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsACSSDebitSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsACSSDebitSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsACSSDebitSetupFutureUsage = "on_session"
)

// Bank account verification method.
type CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethodInstant       CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsAffirmCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsAffirmCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsAffirmCaptureMethodManual CheckoutSessionPaymentMethodOptionsAffirmCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsAffirmSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsAffirmSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsAffirmSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsAffirmSetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsAfterpayClearpayCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsAfterpayClearpayCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsAfterpayClearpayCaptureMethodManual CheckoutSessionPaymentMethodOptionsAfterpayClearpayCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsAfterpayClearpaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsAfterpayClearpaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsAfterpayClearpaySetupFutureUsageNone CheckoutSessionPaymentMethodOptionsAfterpayClearpaySetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsAlipaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsAlipaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsAlipaySetupFutureUsageNone CheckoutSessionPaymentMethodOptionsAlipaySetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsAlmaCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsAlmaCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsAlmaCaptureMethodManual CheckoutSessionPaymentMethodOptionsAlmaCaptureMethod = "manual"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsAmazonPayCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsAmazonPayCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsAmazonPayCaptureMethodManual CheckoutSessionPaymentMethodOptionsAmazonPayCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsAmazonPaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsAmazonPaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsAmazonPaySetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsAmazonPaySetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsAmazonPaySetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsAmazonPaySetupFutureUsage = "off_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsAUBECSDebitSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsAUBECSDebitSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsAUBECSDebitSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsAUBECSDebitSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage = "on_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsBancontactSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsBancontactSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsBancontactSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsBancontactSetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsBillieCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsBillieCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsBillieCaptureMethodManual CheckoutSessionPaymentMethodOptionsBillieCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage = "on_session"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsCardCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsCardCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsCardCaptureMethodManual CheckoutSessionPaymentMethodOptionsCardCaptureMethod = "manual"
)

// Request ability to [capture beyond the standard authorization validity window](https://docs.stripe.com/payments/extended-authorization) for this CheckoutSession.
type CheckoutSessionPaymentMethodOptionsCardRequestExtendedAuthorization string

// List of values that CheckoutSessionPaymentMethodOptionsCardRequestExtendedAuthorization can take
const (
	CheckoutSessionPaymentMethodOptionsCardRequestExtendedAuthorizationIfAvailable CheckoutSessionPaymentMethodOptionsCardRequestExtendedAuthorization = "if_available"
	CheckoutSessionPaymentMethodOptionsCardRequestExtendedAuthorizationNever       CheckoutSessionPaymentMethodOptionsCardRequestExtendedAuthorization = "never"
)

// Request ability to [increment the authorization](https://docs.stripe.com/payments/incremental-authorization) for this CheckoutSession.
type CheckoutSessionPaymentMethodOptionsCardRequestIncrementalAuthorization string

// List of values that CheckoutSessionPaymentMethodOptionsCardRequestIncrementalAuthorization can take
const (
	CheckoutSessionPaymentMethodOptionsCardRequestIncrementalAuthorizationIfAvailable CheckoutSessionPaymentMethodOptionsCardRequestIncrementalAuthorization = "if_available"
	CheckoutSessionPaymentMethodOptionsCardRequestIncrementalAuthorizationNever       CheckoutSessionPaymentMethodOptionsCardRequestIncrementalAuthorization = "never"
)

// Request ability to make [multiple captures](https://docs.stripe.com/payments/multicapture) for this CheckoutSession.
type CheckoutSessionPaymentMethodOptionsCardRequestMulticapture string

// List of values that CheckoutSessionPaymentMethodOptionsCardRequestMulticapture can take
const (
	CheckoutSessionPaymentMethodOptionsCardRequestMulticaptureIfAvailable CheckoutSessionPaymentMethodOptionsCardRequestMulticapture = "if_available"
	CheckoutSessionPaymentMethodOptionsCardRequestMulticaptureNever       CheckoutSessionPaymentMethodOptionsCardRequestMulticapture = "never"
)

// Request ability to [overcapture](https://docs.stripe.com/payments/overcapture) for this CheckoutSession.
type CheckoutSessionPaymentMethodOptionsCardRequestOvercapture string

// List of values that CheckoutSessionPaymentMethodOptionsCardRequestOvercapture can take
const (
	CheckoutSessionPaymentMethodOptionsCardRequestOvercaptureIfAvailable CheckoutSessionPaymentMethodOptionsCardRequestOvercapture = "if_available"
	CheckoutSessionPaymentMethodOptionsCardRequestOvercaptureNever       CheckoutSessionPaymentMethodOptionsCardRequestOvercapture = "never"
)

// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type CheckoutSessionPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that CheckoutSessionPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	CheckoutSessionPaymentMethodOptionsCardRequestThreeDSecureAny       CheckoutSessionPaymentMethodOptionsCardRequestThreeDSecure = "any"
	CheckoutSessionPaymentMethodOptionsCardRequestThreeDSecureAutomatic CheckoutSessionPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	CheckoutSessionPaymentMethodOptionsCardRequestThreeDSecureChallenge CheckoutSessionPaymentMethodOptionsCardRequestThreeDSecure = "challenge"
)

// Specify the card brands to block in the Checkout Session. If a customer enters or selects a card belonging to a blocked brand, they can't complete the Session.
type CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlocked string

// List of values that CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlocked can take
const (
	CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlockedAmericanExpress       CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlocked = "american_express"
	CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlockedDiscoverGlobalNetwork CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlocked = "discover_global_network"
	CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlockedMastercard            CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlocked = "mastercard"
	CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlockedVisa                  CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlocked = "visa"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsCardSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsCardSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsCardSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage = "on_session"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsCashAppCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsCashAppCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsCashAppCaptureMethodManual CheckoutSessionPaymentMethodOptionsCashAppCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsCashAppSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsCashAppSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsCashAppSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsCashAppSetupFutureUsage = "none"
)

// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
//
// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType string

// List of values that CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType can take
const (
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeABA      CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "aba"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeIBAN     CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "iban"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSEPA     CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "sepa"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSortCode CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "sort_code"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSpei     CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "spei"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSwift    CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "swift"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeZengin   CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "zengin"
)

// The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType string

// List of values that CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType can take
const (
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferTypeEUBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType = "eu_bank_transfer"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferTypeGBBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType = "gb_bank_transfer"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferTypeJPBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType = "jp_bank_transfer"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferTypeMXBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType = "mx_bank_transfer"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferTypeUSBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType = "us_bank_transfer"
)

// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
type CheckoutSessionPaymentMethodOptionsCustomerBalanceFundingType string

// List of values that CheckoutSessionPaymentMethodOptionsCustomerBalanceFundingType can take
const (
	CheckoutSessionPaymentMethodOptionsCustomerBalanceFundingTypeBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceFundingType = "bank_transfer"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsCustomerBalanceSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsCustomerBalanceSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsCustomerBalanceSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsCustomerBalanceSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsEPSSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsEPSSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsEPSSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsEPSSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsFPXSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsFPXSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsFPXSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsFPXSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsGiropaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsGiropaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsGiropaySetupFutureUsageNone CheckoutSessionPaymentMethodOptionsGiropaySetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsGrabpaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsGrabpaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsGrabpaySetupFutureUsageNone CheckoutSessionPaymentMethodOptionsGrabpaySetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsIDEALSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsIDEALSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsIDEALSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsIDEALSetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsKakaoPayCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsKakaoPayCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsKakaoPayCaptureMethodManual CheckoutSessionPaymentMethodOptionsKakaoPayCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsKakaoPaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsKakaoPaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsKakaoPaySetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsKakaoPaySetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsKakaoPaySetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsKakaoPaySetupFutureUsage = "off_session"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsKlarnaCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsKlarnaCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsKlarnaCaptureMethodManual CheckoutSessionPaymentMethodOptionsKlarnaCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage = "on_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsKonbiniSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsKonbiniSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsKonbiniSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsKonbiniSetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsKrCardCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsKrCardCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsKrCardCaptureMethodManual CheckoutSessionPaymentMethodOptionsKrCardCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsKrCardSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsKrCardSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsKrCardSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsKrCardSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsKrCardSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsKrCardSetupFutureUsage = "off_session"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsLinkCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsLinkCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsLinkCaptureMethodManual CheckoutSessionPaymentMethodOptionsLinkCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsLinkSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsLinkSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsLinkSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsLinkSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsLinkSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsLinkSetupFutureUsage = "off_session"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsMobilepayCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsMobilepayCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsMobilepayCaptureMethodManual CheckoutSessionPaymentMethodOptionsMobilepayCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsMobilepaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsMobilepaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsMobilepaySetupFutureUsageNone CheckoutSessionPaymentMethodOptionsMobilepaySetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsMultibancoSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsMultibancoSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsMultibancoSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsMultibancoSetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsNaverPayCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsNaverPayCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsNaverPayCaptureMethodManual CheckoutSessionPaymentMethodOptionsNaverPayCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsNaverPaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsNaverPaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsNaverPaySetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsNaverPaySetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsNaverPaySetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsNaverPaySetupFutureUsage = "off_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsOXXOSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsOXXOSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsOXXOSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsOXXOSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsP24SetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsP24SetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsP24SetupFutureUsageNone CheckoutSessionPaymentMethodOptionsP24SetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsPaycoCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsPaycoCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsPaycoCaptureMethodManual CheckoutSessionPaymentMethodOptionsPaycoCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsPayNowSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsPayNowSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsPayNowSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsPayNowSetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsPaypalCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsPaypalCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsPaypalCaptureMethodManual CheckoutSessionPaymentMethodOptionsPaypalCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsPaypalSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsPaypalSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsPaypalSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsPaypalSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsPaypalSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsPaypalSetupFutureUsage = "off_session"
)

// The type of amount that will be collected. The amount charged must be exact or up to the value of `amount` param for `fixed` or `maximum` type respectively. Defaults to `maximum`.
type CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsAmountType string

// List of values that CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsAmountType can take
const (
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsAmountTypeFixed   CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsAmountType = "fixed"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsAmountTypeMaximum CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsAmountType = "maximum"
)

// The periodicity at which payments will be collected. Defaults to `adhoc`.
type CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule string

// List of values that CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule can take
const (
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentScheduleAdhoc       CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule = "adhoc"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentScheduleAnnual      CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule = "annual"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentScheduleDaily       CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule = "daily"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentScheduleFortnightly CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule = "fortnightly"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentScheduleMonthly     CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule = "monthly"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentScheduleQuarterly   CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule = "quarterly"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentScheduleSemiAnnual  CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule = "semi_annual"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentScheduleWeekly      CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule = "weekly"
)

// The purpose for which payments are made. Has a default value based on your merchant category code.
type CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose string

// List of values that CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose can take
const (
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposeDependantSupport CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "dependant_support"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposeGovernment       CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "government"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposeLoan             CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "loan"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposeMortgage         CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "mortgage"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposeOther            CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "other"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposePension          CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "pension"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposePersonal         CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "personal"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposeRetail           CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "retail"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposeSalary           CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "salary"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposeTax              CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "tax"
	CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurposeUtility          CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose = "utility"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsPaytoSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsPaytoSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsPaytoSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsPaytoSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsPaytoSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsPaytoSetupFutureUsage = "off_session"
)

// Determines if the amount includes the IOF tax.
type CheckoutSessionPaymentMethodOptionsPixAmountIncludesIof string

// List of values that CheckoutSessionPaymentMethodOptionsPixAmountIncludesIof can take
const (
	CheckoutSessionPaymentMethodOptionsPixAmountIncludesIofAlways CheckoutSessionPaymentMethodOptionsPixAmountIncludesIof = "always"
	CheckoutSessionPaymentMethodOptionsPixAmountIncludesIofNever  CheckoutSessionPaymentMethodOptionsPixAmountIncludesIof = "never"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsPixSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsPixSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsPixSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsPixSetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsRevolutPayCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsRevolutPayCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsRevolutPayCaptureMethodManual CheckoutSessionPaymentMethodOptionsRevolutPayCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsRevolutPaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsRevolutPaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsRevolutPaySetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsRevolutPaySetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsRevolutPaySetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsRevolutPaySetupFutureUsage = "off_session"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsSamsungPayCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsSamsungPayCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsSamsungPayCaptureMethodManual CheckoutSessionPaymentMethodOptionsSamsungPayCaptureMethod = "manual"
)

// Controls when the funds will be captured from the customer's account.
type CheckoutSessionPaymentMethodOptionsSatispayCaptureMethod string

// List of values that CheckoutSessionPaymentMethodOptionsSatispayCaptureMethod can take
const (
	CheckoutSessionPaymentMethodOptionsSatispayCaptureMethodManual CheckoutSessionPaymentMethodOptionsSatispayCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsage = "on_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsSofortSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsSofortSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsSofortSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsSofortSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsTWINTSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsTWINTSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsTWINTSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsTWINTSetupFutureUsage = "none"
)

// The account subcategories to use to filter for possible accounts to link. Valid subcategories are `checking` and `savings`.
type CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory string

// List of values that CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory can take
const (
	CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategoryChecking CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "checking"
	CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategorySavings  CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "savings"
)

// The list of permissions to request. The `payment_method` permission must be included.
type CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission string

// List of values that CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission can take
const (
	CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionBalances      CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "balances"
	CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionOwnership     CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "ownership"
	CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionPaymentMethod CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "payment_method"
	CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionTransactions  CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "transactions"
)

// Data features requested to be retrieved upon account creation.
type CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch string

// List of values that CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch can take
const (
	CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchBalances     CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "balances"
	CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchOwnership    CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "ownership"
	CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchTransactions CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "transactions"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
//
// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
//
// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsUSBankAccountSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsUSBankAccountSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsUSBankAccountSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsUSBankAccountSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsUSBankAccountSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsUSBankAccountSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsUSBankAccountSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsUSBankAccountSetupFutureUsage = "on_session"
)

// Bank account verification method.
type CheckoutSessionPaymentMethodOptionsUSBankAccountVerificationMethod string

// List of values that CheckoutSessionPaymentMethodOptionsUSBankAccountVerificationMethod can take
const (
	CheckoutSessionPaymentMethodOptionsUSBankAccountVerificationMethodAutomatic CheckoutSessionPaymentMethodOptionsUSBankAccountVerificationMethod = "automatic"
	CheckoutSessionPaymentMethodOptionsUSBankAccountVerificationMethodInstant   CheckoutSessionPaymentMethodOptionsUSBankAccountVerificationMethod = "instant"
)

// The payment status of the Checkout Session, one of `paid`, `unpaid`, or `no_payment_required`.
// You can use this value to decide when to fulfill your customer's order.
type CheckoutSessionPaymentStatus string

// List of values that CheckoutSessionPaymentStatus can take
const (
	CheckoutSessionPaymentStatusNoPaymentRequired CheckoutSessionPaymentStatus = "no_payment_required"
	CheckoutSessionPaymentStatusPaid              CheckoutSessionPaymentStatus = "paid"
	CheckoutSessionPaymentStatusUnpaid            CheckoutSessionPaymentStatus = "unpaid"
)

// Determines which entity is allowed to update the shipping details.
//
// Default is `client_only`. Stripe Checkout client will automatically update the shipping details. If set to `server_only`, only your server is allowed to update the shipping details.
//
// When set to `server_only`, you must add the onShippingDetailsChange event handler when initializing the Stripe Checkout client and manually update the shipping details from your server using the Stripe API.
type CheckoutSessionPermissionsUpdateShippingDetails string

// List of values that CheckoutSessionPermissionsUpdateShippingDetails can take
const (
	CheckoutSessionPermissionsUpdateShippingDetailsClientOnly CheckoutSessionPermissionsUpdateShippingDetails = "client_only"
	CheckoutSessionPermissionsUpdateShippingDetailsServerOnly CheckoutSessionPermissionsUpdateShippingDetails = "server_only"
)

// This parameter applies to `ui_mode: embedded`. Learn more about the [redirect behavior](https://docs.stripe.com/payments/checkout/custom-success-page?payment-ui=embedded-form) of embedded sessions. Defaults to `always`.
type CheckoutSessionRedirectOnCompletion string

// List of values that CheckoutSessionRedirectOnCompletion can take
const (
	CheckoutSessionRedirectOnCompletionAlways     CheckoutSessionRedirectOnCompletion = "always"
	CheckoutSessionRedirectOnCompletionIfRequired CheckoutSessionRedirectOnCompletion = "if_required"
	CheckoutSessionRedirectOnCompletionNever      CheckoutSessionRedirectOnCompletion = "never"
)

// Uses the `allow_redisplay` value of each saved payment method to filter the set presented to a returning customer. By default, only saved payment methods with 'allow_redisplay: ‘always' are shown in Checkout.
type CheckoutSessionSavedPaymentMethodOptionsAllowRedisplayFilter string

// List of values that CheckoutSessionSavedPaymentMethodOptionsAllowRedisplayFilter can take
const (
	CheckoutSessionSavedPaymentMethodOptionsAllowRedisplayFilterAlways      CheckoutSessionSavedPaymentMethodOptionsAllowRedisplayFilter = "always"
	CheckoutSessionSavedPaymentMethodOptionsAllowRedisplayFilterLimited     CheckoutSessionSavedPaymentMethodOptionsAllowRedisplayFilter = "limited"
	CheckoutSessionSavedPaymentMethodOptionsAllowRedisplayFilterUnspecified CheckoutSessionSavedPaymentMethodOptionsAllowRedisplayFilter = "unspecified"
)

// Enable customers to choose if they wish to remove their saved payment methods. Disabled by default.
type CheckoutSessionSavedPaymentMethodOptionsPaymentMethodRemove string

// List of values that CheckoutSessionSavedPaymentMethodOptionsPaymentMethodRemove can take
const (
	CheckoutSessionSavedPaymentMethodOptionsPaymentMethodRemoveDisabled CheckoutSessionSavedPaymentMethodOptionsPaymentMethodRemove = "disabled"
	CheckoutSessionSavedPaymentMethodOptionsPaymentMethodRemoveEnabled  CheckoutSessionSavedPaymentMethodOptionsPaymentMethodRemove = "enabled"
)

// Enable customers to choose if they wish to save their payment method for future use. Disabled by default.
type CheckoutSessionSavedPaymentMethodOptionsPaymentMethodSave string

// List of values that CheckoutSessionSavedPaymentMethodOptionsPaymentMethodSave can take
const (
	CheckoutSessionSavedPaymentMethodOptionsPaymentMethodSaveDisabled CheckoutSessionSavedPaymentMethodOptionsPaymentMethodSave = "disabled"
	CheckoutSessionSavedPaymentMethodOptionsPaymentMethodSaveEnabled  CheckoutSessionSavedPaymentMethodOptionsPaymentMethodSave = "enabled"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type CheckoutSessionShippingCostTaxTaxabilityReason string

// List of values that CheckoutSessionShippingCostTaxTaxabilityReason can take
const (
	CheckoutSessionShippingCostTaxTaxabilityReasonCustomerExempt       CheckoutSessionShippingCostTaxTaxabilityReason = "customer_exempt"
	CheckoutSessionShippingCostTaxTaxabilityReasonNotCollecting        CheckoutSessionShippingCostTaxTaxabilityReason = "not_collecting"
	CheckoutSessionShippingCostTaxTaxabilityReasonNotSubjectToTax      CheckoutSessionShippingCostTaxTaxabilityReason = "not_subject_to_tax"
	CheckoutSessionShippingCostTaxTaxabilityReasonNotSupported         CheckoutSessionShippingCostTaxTaxabilityReason = "not_supported"
	CheckoutSessionShippingCostTaxTaxabilityReasonPortionProductExempt CheckoutSessionShippingCostTaxTaxabilityReason = "portion_product_exempt"
	CheckoutSessionShippingCostTaxTaxabilityReasonPortionReducedRated  CheckoutSessionShippingCostTaxTaxabilityReason = "portion_reduced_rated"
	CheckoutSessionShippingCostTaxTaxabilityReasonPortionStandardRated CheckoutSessionShippingCostTaxTaxabilityReason = "portion_standard_rated"
	CheckoutSessionShippingCostTaxTaxabilityReasonProductExempt        CheckoutSessionShippingCostTaxTaxabilityReason = "product_exempt"
	CheckoutSessionShippingCostTaxTaxabilityReasonProductExemptHoliday CheckoutSessionShippingCostTaxTaxabilityReason = "product_exempt_holiday"
	CheckoutSessionShippingCostTaxTaxabilityReasonProportionallyRated  CheckoutSessionShippingCostTaxTaxabilityReason = "proportionally_rated"
	CheckoutSessionShippingCostTaxTaxabilityReasonReducedRated         CheckoutSessionShippingCostTaxTaxabilityReason = "reduced_rated"
	CheckoutSessionShippingCostTaxTaxabilityReasonReverseCharge        CheckoutSessionShippingCostTaxTaxabilityReason = "reverse_charge"
	CheckoutSessionShippingCostTaxTaxabilityReasonStandardRated        CheckoutSessionShippingCostTaxTaxabilityReason = "standard_rated"
	CheckoutSessionShippingCostTaxTaxabilityReasonTaxableBasisReduced  CheckoutSessionShippingCostTaxTaxabilityReason = "taxable_basis_reduced"
	CheckoutSessionShippingCostTaxTaxabilityReasonZeroRated            CheckoutSessionShippingCostTaxTaxabilityReason = "zero_rated"
)

// The status of the Checkout Session, one of `open`, `complete`, or `expired`.
type CheckoutSessionStatus string

// List of values that CheckoutSessionStatus can take
const (
	CheckoutSessionStatusComplete CheckoutSessionStatus = "complete"
	CheckoutSessionStatusExpired  CheckoutSessionStatus = "expired"
	CheckoutSessionStatusOpen     CheckoutSessionStatus = "open"
)

// Describes the type of transaction being performed by Checkout in order to customize
// relevant text on the page, such as the submit button. `submit_type` can only be
// specified on Checkout Sessions in `payment` mode. If blank or `auto`, `pay` is used.
type CheckoutSessionSubmitType string

// List of values that CheckoutSessionSubmitType can take
const (
	CheckoutSessionSubmitTypeAuto      CheckoutSessionSubmitType = "auto"
	CheckoutSessionSubmitTypeBook      CheckoutSessionSubmitType = "book"
	CheckoutSessionSubmitTypeDonate    CheckoutSessionSubmitType = "donate"
	CheckoutSessionSubmitTypePay       CheckoutSessionSubmitType = "pay"
	CheckoutSessionSubmitTypeSubscribe CheckoutSessionSubmitType = "subscribe"
)

// Indicates whether a tax ID is required on the payment page
type CheckoutSessionTaxIDCollectionRequired string

// List of values that CheckoutSessionTaxIDCollectionRequired can take
const (
	CheckoutSessionTaxIDCollectionRequiredIfSupported CheckoutSessionTaxIDCollectionRequired = "if_supported"
	CheckoutSessionTaxIDCollectionRequiredNever       CheckoutSessionTaxIDCollectionRequired = "never"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason string

// List of values that CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason can take
const (
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonCustomerExempt       CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "customer_exempt"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonNotCollecting        CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "not_collecting"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonNotSubjectToTax      CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "not_subject_to_tax"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonNotSupported         CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "not_supported"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonPortionProductExempt CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "portion_product_exempt"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonPortionReducedRated  CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "portion_reduced_rated"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonPortionStandardRated CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "portion_standard_rated"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonProductExempt        CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonProductExemptHoliday CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt_holiday"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonProportionallyRated  CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "proportionally_rated"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonReducedRated         CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "reduced_rated"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonReverseCharge        CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "reverse_charge"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonStandardRated        CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "standard_rated"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonTaxableBasisReduced  CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "taxable_basis_reduced"
	CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonZeroRated            CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason = "zero_rated"
)

// The UI mode of the Session. Defaults to `hosted`.
type CheckoutSessionUIMode string

// List of values that CheckoutSessionUIMode can take
const (
	CheckoutSessionUIModeCustom   CheckoutSessionUIMode = "custom"
	CheckoutSessionUIModeEmbedded CheckoutSessionUIMode = "embedded"
	CheckoutSessionUIModeHosted   CheckoutSessionUIMode = "hosted"
)

// Describes whether Checkout should display Link. Defaults to `auto`.
type CheckoutSessionWalletOptionsLinkDisplay string

// List of values that CheckoutSessionWalletOptionsLinkDisplay can take
const (
	CheckoutSessionWalletOptionsLinkDisplayAuto  CheckoutSessionWalletOptionsLinkDisplay = "auto"
	CheckoutSessionWalletOptionsLinkDisplayNever CheckoutSessionWalletOptionsLinkDisplay = "never"
)

// Only return the Checkout Sessions for the Customer details specified.
type CheckoutSessionListCustomerDetailsParams struct {
	// Customer's email address.
	Email *string `form:"email"`
}

// Returns a list of Checkout Sessions.
type CheckoutSessionListParams struct {
	ListParams `form:"*"`
	// Only return Checkout Sessions that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return Checkout Sessions that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return the Checkout Sessions for the Customer specified.
	Customer *string `form:"customer"`
	// Only return the Checkout Sessions for the Account specified.
	CustomerAccount *string `form:"customer_account"`
	// Only return the Checkout Sessions for the Customer details specified.
	CustomerDetails *CheckoutSessionListCustomerDetailsParams `form:"customer_details"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return the Checkout Session for the PaymentIntent specified.
	PaymentIntent *string `form:"payment_intent"`
	// Only return the Checkout Sessions for the Payment Link specified.
	PaymentLink *string `form:"payment_link"`
	// Only return the Checkout Sessions matching the given status.
	Status *string `form:"status"`
	// Only return the Checkout Session for the subscription specified.
	Subscription *string `form:"subscription"`
}

// AddExpand appends a new field to expand.
func (p *CheckoutSessionListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Settings for price localization with [Adaptive Pricing](https://docs.stripe.com/payments/checkout/adaptive-pricing).
type CheckoutSessionAdaptivePricingParams struct {
	// If set to `true`, Adaptive Pricing is available on [eligible sessions](https://docs.stripe.com/payments/currencies/localize-prices/adaptive-pricing?payment-ui=stripe-hosted#restrictions). Defaults to your [dashboard setting](https://dashboard.stripe.com/settings/adaptive-pricing).
	Enabled *bool `form:"enabled"`
}

// Configure a Checkout Session that can be used to recover an expired session.
type CheckoutSessionAfterExpirationRecoveryParams struct {
	// Enables user redeemable promotion codes on the recovered Checkout Sessions. Defaults to `false`
	AllowPromotionCodes *bool `form:"allow_promotion_codes"`
	// If `true`, a recovery URL will be generated to recover this Checkout Session if it
	// expires before a successful transaction is completed. It will be attached to the
	// Checkout Session object upon expiration.
	Enabled *bool `form:"enabled"`
}

// Configure actions after a Checkout Session has expired. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionAfterExpirationParams struct {
	// Configure a Checkout Session that can be used to recover an expired session.
	Recovery *CheckoutSessionAfterExpirationRecoveryParams `form:"recovery"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type CheckoutSessionAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// Settings for automatic tax lookup for this session and resulting payments, invoices, and subscriptions.
type CheckoutSessionAutomaticTaxParams struct {
	// Set to `true` to [calculate tax automatically](https://docs.stripe.com/tax) using the customer's location.
	//
	// Enabling this parameter causes Checkout to collect any billing address information necessary for tax calculation.
	Enabled *bool `form:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *CheckoutSessionAutomaticTaxLiabilityParams `form:"liability"`
}

// The icon for the Checkout Session. For best results, use a square image.
type CheckoutSessionBrandingSettingsIconParams struct {
	// The ID of a [File upload](https://stripe.com/docs/api/files) representing the icon. Purpose must be `business_icon`. Required if `type` is `file` and disallowed otherwise.
	File *string `form:"file"`
	// The type of image for the icon. Must be one of `file` or `url`.
	Type *string `form:"type"`
	// The URL of the image. Required if `type` is `url` and disallowed otherwise.
	URL *string `form:"url"`
}

// The logo for the Checkout Session.
type CheckoutSessionBrandingSettingsLogoParams struct {
	// The ID of a [File upload](https://stripe.com/docs/api/files) representing the logo. Purpose must be `business_logo`. Required if `type` is `file` and disallowed otherwise.
	File *string `form:"file"`
	// The type of image for the logo. Must be one of `file` or `url`.
	Type *string `form:"type"`
	// The URL of the image. Required if `type` is `url` and disallowed otherwise.
	URL *string `form:"url"`
}

// The branding settings for the Checkout Session. This parameter is not allowed if ui_mode is `custom`.
type CheckoutSessionBrandingSettingsParams struct {
	// A hex color value starting with `#` representing the background color for the Checkout Session.
	BackgroundColor *string `form:"background_color"`
	// The border style for the Checkout Session.
	BorderStyle *string `form:"border_style"`
	// A hex color value starting with `#` representing the button color for the Checkout Session.
	ButtonColor *string `form:"button_color"`
	// A string to override the business name shown on the Checkout Session. This only shows at the top of the Checkout page, and your business name still appears in terms, receipts, and other places.
	DisplayName *string `form:"display_name"`
	// The font family for the Checkout Session corresponding to one of the [supported font families](https://docs.stripe.com/payments/checkout/customization/appearance?payment-ui=stripe-hosted#font-compatibility).
	FontFamily *string `form:"font_family"`
	// The icon for the Checkout Session. For best results, use a square image.
	Icon *CheckoutSessionBrandingSettingsIconParams `form:"icon"`
	// The logo for the Checkout Session.
	Logo *CheckoutSessionBrandingSettingsLogoParams `form:"logo"`
}

// Determines the display of payment method reuse agreement text in the UI. If set to `hidden`, it will hide legal text related to the reuse of a payment method.
type CheckoutSessionConsentCollectionPaymentMethodReuseAgreementParams struct {
	// Determines the position and visibility of the payment method reuse agreement in the UI. When set to `auto`, Stripe's
	// defaults will be used. When set to `hidden`, the payment method reuse agreement text will always be hidden in the UI.
	Position *string `form:"position"`
}

// Configure fields for the Checkout Session to gather active consent from customers.
type CheckoutSessionConsentCollectionParams struct {
	// Determines the display of payment method reuse agreement text in the UI. If set to `hidden`, it will hide legal text related to the reuse of a payment method.
	PaymentMethodReuseAgreement *CheckoutSessionConsentCollectionPaymentMethodReuseAgreementParams `form:"payment_method_reuse_agreement"`
	// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
	// Session will determine whether to display an option to opt into promotional communication
	// from the merchant depending on the customer's locale. Only available to US merchants.
	Promotions *string `form:"promotions"`
	// If set to `required`, it requires customers to check a terms of service checkbox before being able to pay.
	// There must be a valid terms of service URL set in your [Dashboard settings](https://dashboard.stripe.com/settings/public).
	TermsOfService *string `form:"terms_of_service"`
}

// The options available for the customer to select. Up to 200 options allowed.
type CheckoutSessionCustomFieldDropdownOptionParams struct {
	// The label for the option, displayed to the customer. Up to 100 characters.
	Label *string `form:"label"`
	// The value for this option, not displayed to the customer, used by your integration to reconcile the option selected by the customer. Must be unique to this option, alphanumeric, and up to 100 characters.
	Value *string `form:"value"`
}

// Configuration for `type=dropdown` fields.
type CheckoutSessionCustomFieldDropdownParams struct {
	// The value that pre-fills the field on the payment page.Must match a `value` in the `options` array.
	DefaultValue *string `form:"default_value"`
	// The options available for the customer to select. Up to 200 options allowed.
	Options []*CheckoutSessionCustomFieldDropdownOptionParams `form:"options"`
}

// The label for the field, displayed to the customer.
type CheckoutSessionCustomFieldLabelParams struct {
	// Custom text for the label, displayed to the customer. Up to 50 characters.
	Custom *string `form:"custom"`
	// The type of the label.
	Type *string `form:"type"`
}

// Configuration for `type=numeric` fields.
type CheckoutSessionCustomFieldNumericParams struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue *string `form:"default_value"`
	// The maximum character length constraint for the customer's input.
	MaximumLength *int64 `form:"maximum_length"`
	// The minimum character length requirement for the customer's input.
	MinimumLength *int64 `form:"minimum_length"`
}

// Configuration for `type=text` fields.
type CheckoutSessionCustomFieldTextParams struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue *string `form:"default_value"`
	// The maximum character length constraint for the customer's input.
	MaximumLength *int64 `form:"maximum_length"`
	// The minimum character length requirement for the customer's input.
	MinimumLength *int64 `form:"minimum_length"`
}

// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionCustomFieldParams struct {
	// Configuration for `type=dropdown` fields.
	Dropdown *CheckoutSessionCustomFieldDropdownParams `form:"dropdown"`
	// String of your choice that your integration can use to reconcile this field. Must be unique to this field, alphanumeric, and up to 200 characters.
	Key *string `form:"key"`
	// The label for the field, displayed to the customer.
	Label *CheckoutSessionCustomFieldLabelParams `form:"label"`
	// Configuration for `type=numeric` fields.
	Numeric *CheckoutSessionCustomFieldNumericParams `form:"numeric"`
	// Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.
	Optional *bool `form:"optional"`
	// Configuration for `type=text` fields.
	Text *CheckoutSessionCustomFieldTextParams `form:"text"`
	// The type of the field.
	Type *string `form:"type"`
}

// Custom text that should be displayed after the payment confirmation button.
type CheckoutSessionCustomTextAfterSubmitParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message"`
}

// Custom text that should be displayed alongside shipping address collection.
type CheckoutSessionCustomTextShippingAddressParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message"`
}

// Custom text that should be displayed alongside the payment confirmation button.
type CheckoutSessionCustomTextSubmitParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message"`
}

// Custom text that should be displayed in place of the default terms of service agreement text.
type CheckoutSessionCustomTextTermsOfServiceAcceptanceParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message"`
}

// Display additional text for your customers using custom text. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionCustomTextParams struct {
	// Custom text that should be displayed after the payment confirmation button.
	AfterSubmit *CheckoutSessionCustomTextAfterSubmitParams `form:"after_submit"`
	// Custom text that should be displayed alongside shipping address collection.
	ShippingAddress *CheckoutSessionCustomTextShippingAddressParams `form:"shipping_address"`
	// Custom text that should be displayed alongside the payment confirmation button.
	Submit *CheckoutSessionCustomTextSubmitParams `form:"submit"`
	// Custom text that should be displayed in place of the default terms of service agreement text.
	TermsOfServiceAcceptance *CheckoutSessionCustomTextTermsOfServiceAcceptanceParams `form:"terms_of_service_acceptance"`
}

// Controls what fields on Customer can be updated by the Checkout Session. Can only be provided when `customer` is provided.
type CheckoutSessionCustomerUpdateParams struct {
	// Describes whether Checkout saves the billing address onto `customer.address`.
	// To always collect a full billing address, use `billing_address_collection`. Defaults to `never`.
	Address *string `form:"address"`
	// Describes whether Checkout saves the name onto `customer.name`. Defaults to `never`.
	Name *string `form:"name"`
	// Describes whether Checkout saves shipping information onto `customer.shipping`.
	// To collect shipping information, use `shipping_address_collection`. Defaults to `never`.
	Shipping *string `form:"shipping"`
}

// The coupon or promotion code to apply to this Session. Currently, only up to one may be specified.
type CheckoutSessionDiscountParams struct {
	// The ID of the coupon to apply to this Session.
	Coupon *string `form:"coupon"`
	// The ID of a promotion code to apply to this Session.
	PromotionCode *string `form:"promotion_code"`
}

// Default custom fields to be displayed on invoices for this customer.
type CheckoutSessionInvoiceCreationInvoiceDataCustomFieldParams struct {
	// The name of the custom field. This may be up to 40 characters.
	Name *string `form:"name"`
	// The value of the custom field. This may be up to 140 characters.
	Value *string `form:"value"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type CheckoutSessionInvoiceCreationInvoiceDataIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// Default options for invoice PDF rendering for this customer.
type CheckoutSessionInvoiceCreationInvoiceDataRenderingOptionsParams struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs. One of `exclude_tax` or `include_inclusive_tax`. `include_inclusive_tax` will include inclusive tax (and exclude exclusive tax) in invoice PDF amounts. `exclude_tax` will exclude all tax (inclusive and exclusive alike) from invoice PDF amounts.
	AmountTaxDisplay *string `form:"amount_tax_display"`
	// ID of the invoice rendering template to use for this invoice.
	Template *string `form:"template"`
}

// Parameters passed when creating invoices for payment-mode Checkout Sessions.
type CheckoutSessionInvoiceCreationInvoiceDataParams struct {
	// The account tax IDs associated with the invoice.
	AccountTaxIDs []*string `form:"account_tax_ids"`
	// Default custom fields to be displayed on invoices for this customer.
	CustomFields []*CheckoutSessionInvoiceCreationInvoiceDataCustomFieldParams `form:"custom_fields"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Default footer to be displayed on invoices for this customer.
	Footer *string `form:"footer"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *CheckoutSessionInvoiceCreationInvoiceDataIssuerParams `form:"issuer"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Default options for invoice PDF rendering for this customer.
	RenderingOptions *CheckoutSessionInvoiceCreationInvoiceDataRenderingOptionsParams `form:"rendering_options"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionInvoiceCreationInvoiceDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Generate a post-purchase Invoice for one-time payments.
type CheckoutSessionInvoiceCreationParams struct {
	// Set to `true` to enable invoice creation.
	Enabled *bool `form:"enabled"`
	// Parameters passed when creating invoices for payment-mode Checkout Sessions.
	InvoiceData *CheckoutSessionInvoiceCreationInvoiceDataParams `form:"invoice_data"`
}

// When set, provides configuration for this item's quantity to be adjusted by the customer during Checkout.
type CheckoutSessionLineItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any positive integer. Setting to false will remove any previously specified constraints on quantity.
	Enabled *bool `form:"enabled"`
	// The maximum quantity the customer can purchase for the Checkout Session. By default this value is 99. You can specify a value up to 999999.
	Maximum *int64 `form:"maximum"`
	// The minimum quantity the customer must purchase for the Checkout Session. By default this value is 0.
	Minimum *int64 `form:"minimum"`
}

// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.
type CheckoutSessionLineItemPriceDataProductDataParams struct {
	// The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description"`
	// A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
	// A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.
	UnitLabel *string `form:"unit_label"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionLineItemPriceDataProductDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The recurring components of a price such as `interval` and `interval_count`.
type CheckoutSessionLineItemPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
type CheckoutSessionLineItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to. One of `product` or `product_data` is required.
	Product *string `form:"product"`
	// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.
	ProductData *CheckoutSessionLineItemPriceDataProductDataParams `form:"product_data"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *CheckoutSessionLineItemPriceDataRecurringParams `form:"recurring"`
	// Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A non-negative integer in cents (or local equivalent) representing how much to charge. One of `unit_amount` or `unit_amount_decimal` is required.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// A list of items the customer is purchasing. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices). The parameter is required for `payment` and `subscription` mode.
//
// For `payment` mode, there is a maximum of 100 line items, however it is recommended to consolidate line items if there are more than a few dozen.
//
// For `subscription` mode, there is a maximum of 20 line items with recurring Prices and 20 line items with one-time Prices. Line items with one-time Prices will be on the initial invoice only.
type CheckoutSessionLineItemParams struct {
	// When set, provides configuration for this item's quantity to be adjusted by the customer during Checkout.
	AdjustableQuantity *CheckoutSessionLineItemAdjustableQuantityParams `form:"adjustable_quantity"`
	// The [tax rates](https://docs.stripe.com/api/tax_rates) that will be applied to this line item depending on the customer's billing/shipping address. We currently support the following countries: US, GB, AU, and all countries in the EU. You can't set this parameter if `ui_mode` is `custom`.
	DynamicTaxRates []*string `form:"dynamic_tax_rates"`
	// ID of an existing line item.
	ID *string `form:"id"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object. One of `price` or `price_data` is required.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required when creating a new line item.
	PriceData *CheckoutSessionLineItemPriceDataParams `form:"price_data"`
	// The quantity of the line item being purchased. Quantity should not be defined when `recurring.usage_type=metered`.
	Quantity *int64 `form:"quantity"`
	// The [tax rates](https://docs.stripe.com/api/tax_rates) which apply to this line item.
	TaxRates []*string `form:"tax_rates"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionLineItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Controls settings applied for collecting the customer's business name on the session.
type CheckoutSessionNameCollectionBusinessParams struct {
	// Enable business name collection on the Checkout Session. Defaults to `false`.
	Enabled *bool `form:"enabled"`
	// Whether the customer is required to provide a business name before completing the Checkout Session. Defaults to `false`.
	Optional *bool `form:"optional"`
}

// Controls settings applied for collecting the customer's individual name on the session.
type CheckoutSessionNameCollectionIndividualParams struct {
	// Enable individual name collection on the Checkout Session. Defaults to `false`.
	Enabled *bool `form:"enabled"`
	// Whether the customer is required to provide their name before completing the Checkout Session. Defaults to `false`.
	Optional *bool `form:"optional"`
}

// Controls name collection settings for the session.
//
// You can configure Checkout to collect your customers' business names, individual names, or both. Each name field can be either required or optional.
//
// If a [Customer](https://docs.stripe.com/api/customers) is created or provided, the names can be saved to the Customer object as well.
//
// You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionNameCollectionParams struct {
	// Controls settings applied for collecting the customer's business name on the session.
	Business *CheckoutSessionNameCollectionBusinessParams `form:"business"`
	// Controls settings applied for collecting the customer's individual name on the session.
	Individual *CheckoutSessionNameCollectionIndividualParams `form:"individual"`
}

// When set, provides configuration for the customer to adjust the quantity of the line item created when a customer chooses to add this optional item to their order.
type CheckoutSessionOptionalItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative integer.
	Enabled *bool `form:"enabled"`
	// The maximum quantity of this item the customer can purchase. By default this value is 99. You can specify a value up to 999999.
	Maximum *int64 `form:"maximum"`
	// The minimum quantity of this item the customer must purchase, if they choose to purchase it. Because this item is optional, the customer will always be able to remove it from their order, even if the `minimum` configured here is greater than 0. By default this value is 0.
	Minimum *int64 `form:"minimum"`
}

// A list of optional items the customer can add to their order at checkout. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices).
//
// There is a maximum of 10 optional items allowed on a Checkout Session, and the existing limits on the number of line items allowed on a Checkout Session apply to the combined number of line items and optional items.
//
// For `payment` mode, there is a maximum of 100 combined line items and optional items, however it is recommended to consolidate items if there are more than a few dozen.
//
// For `subscription` mode, there is a maximum of 20 line items and optional items with recurring Prices and 20 line items and optional items with one-time Prices.
//
// You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionOptionalItemParams struct {
	// When set, provides configuration for the customer to adjust the quantity of the line item created when a customer chooses to add this optional item to their order.
	AdjustableQuantity *CheckoutSessionOptionalItemAdjustableQuantityParams `form:"adjustable_quantity"`
	// The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object.
	Price *string `form:"price"`
	// The initial quantity of the line item created when a customer chooses to add this optional item to their order.
	Quantity *int64 `form:"quantity"`
}

// The parameters used to automatically create a Transfer when the payment succeeds.
// For more information, see the PaymentIntents [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts).
type CheckoutSessionPaymentIntentDataTransferDataParams struct {
	// The amount that will be transferred automatically when a charge succeeds.
	Amount *int64 `form:"amount"`
	// If specified, successful charges will be attributed to the destination
	// account for tax reporting, and the funds from charges will be transferred
	// to the destination account. The ID of the resulting transfer will be
	// returned on the successful charge's `transfer` field.
	Destination *string `form:"destination"`
}

// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
type CheckoutSessionPaymentIntentDataParams struct {
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. The amount of the application fee collected will be capped at the total amount captured. For more information, see the PaymentIntents [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts).
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The Stripe account ID for which these funds are intended. For details,
	// see the PaymentIntents [use case for connected
	// accounts](https://docs.stripe.com/docs/payments/connected-accounts).
	OnBehalfOf *string `form:"on_behalf_of"`
	// Email address that the receipt for the resulting payment will be sent to. If `receipt_email` is specified for a payment in live mode, a receipt will be sent regardless of your [email settings](https://dashboard.stripe.com/account/emails).
	ReceiptEmail *string `form:"receipt_email"`
	// Indicates that you intend to [make future payments](https://docs.stripe.com/payments/payment-intents#future-usage) with the payment
	// method collected by this Checkout Session.
	//
	// When setting this to `on_session`, Checkout will show a notice to the
	// customer that their payment details will be saved.
	//
	// When setting this to `off_session`, Checkout will show a notice to the
	// customer that their payment details will be saved and used for future
	// payments.
	//
	// If a Customer has been provided or Checkout creates a new Customer,
	// Checkout will attach the payment method to the Customer.
	//
	// If Checkout does not create a Customer, the payment method is not attached
	// to a Customer. To reuse the payment method, you can retrieve it from the
	// Checkout Session's PaymentIntent.
	//
	// When processing card payments, Checkout also uses `setup_future_usage`
	// to dynamically optimize your payment flow and comply with regional
	// legislation and network rules, such as SCA.
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Shipping information for this payment.
	Shipping *ShippingDetailsParams `form:"shipping"`
	// Text that appears on the customer's statement as the statement descriptor for a non-card charge. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	//
	// Setting this value for a card charge returns an error. For card charges, set the [statement_descriptor_suffix](https://docs.stripe.com/get-started/account/statement-descriptors#dynamic) instead.
	StatementDescriptor *string `form:"statement_descriptor"`
	// Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix"`
	// The parameters used to automatically create a Transfer when the payment succeeds.
	// For more information, see the PaymentIntents [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts).
	TransferData *CheckoutSessionPaymentIntentDataTransferDataParams `form:"transfer_data"`
	// A string that identifies the resulting payment as part of a group. See the PaymentIntents [use case for connected accounts](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.
	TransferGroup *string `form:"transfer_group"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionPaymentIntentDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// This parameter allows you to set some attributes on the payment method created during a Checkout session.
type CheckoutSessionPaymentMethodDataParams struct {
	// Allow redisplay will be set on the payment method on confirmation and indicates whether this payment method can be shown again to the customer in a checkout flow. Only set this field if you wish to override the allow_redisplay value determined by Checkout.
	AllowRedisplay *string `form:"allow_redisplay"`
}

// Additional fields for Mandate creation
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// A URL for custom mandate text to render during confirmation step.
	// The URL will be rendered with additional GET parameters `payment_intent` and `payment_intent_client_secret` when confirming a Payment Intent,
	// or `setup_intent` and `setup_intent_client_secret` when confirming a Setup Intent.
	CustomMandateURL *string `form:"custom_mandate_url"`
	// List of Stripe products where this mandate can be selected automatically. Only usable in `setup` mode.
	DefaultFor []*string `form:"default_for"`
	// Description of the mandate interval. Only required if 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription *string `form:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule *string `form:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type"`
}

// contains details about the ACSS Debit payment method options. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionPaymentMethodOptionsACSSDebitParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies). This is only accepted for Checkout Sessions in `setup` mode.
	Currency *string `form:"currency"`
	// Additional fields for Mandate creation
	MandateOptions *CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate *string `form:"target_date"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// contains details about the Affirm payment method options.
type CheckoutSessionPaymentMethodOptionsAffirmParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Afterpay Clearpay payment method options.
type CheckoutSessionPaymentMethodOptionsAfterpayClearpayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Alipay payment method options.
type CheckoutSessionPaymentMethodOptionsAlipayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Alma payment method options.
type CheckoutSessionPaymentMethodOptionsAlmaParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
}

// contains details about the AmazonPay payment method options.
type CheckoutSessionPaymentMethodOptionsAmazonPayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the AU Becs Debit payment method options.
type CheckoutSessionPaymentMethodOptionsAUBECSDebitParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate *string `form:"target_date"`
}

// Additional fields for Mandate creation
type CheckoutSessionPaymentMethodOptionsBACSDebitMandateOptionsParams struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'DDIC' or 'STRIPE'.
	ReferencePrefix *string `form:"reference_prefix"`
}

// contains details about the Bacs Debit payment method options.
type CheckoutSessionPaymentMethodOptionsBACSDebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *CheckoutSessionPaymentMethodOptionsBACSDebitMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate *string `form:"target_date"`
}

// contains details about the Bancontact payment method options.
type CheckoutSessionPaymentMethodOptionsBancontactParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Billie payment method options.
type CheckoutSessionPaymentMethodOptionsBillieParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
}

// contains details about the Boleto payment method options.
type CheckoutSessionPaymentMethodOptionsBoletoParams struct {
	// The number of calendar days before a Boleto voucher expires. For example, if you create a Boleto voucher on Monday and you set expires_after_days to 2, the Boleto invoice will expire on Wednesday at 23:59 America/Sao_Paulo time.
	ExpiresAfterDays *int64 `form:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Installment options for card payments
type CheckoutSessionPaymentMethodOptionsCardInstallmentsParams struct {
	// Setting to true enables installments for this Checkout Session.
	// Setting to false will prevent any installment plan from applying to a payment.
	Enabled *bool `form:"enabled"`
}

// Restrictions to apply to the card payment method. For example, you can block specific card brands. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionPaymentMethodOptionsCardRestrictionsParams struct {
	// Specify the card brands to block in the Checkout Session. If a customer enters or selects a card belonging to a blocked brand, they can't complete the Session.
	BrandsBlocked []*string `form:"brands_blocked"`
}

// contains details about the Card payment method options.
type CheckoutSessionPaymentMethodOptionsCardParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Installment options for card payments
	Installments *CheckoutSessionPaymentMethodOptionsCardInstallmentsParams `form:"installments"`
	// Request ability to [capture beyond the standard authorization validity window](https://docs.stripe.com/payments/extended-authorization) for this CheckoutSession.
	RequestExtendedAuthorization *string `form:"request_extended_authorization"`
	// Request ability to [increment the authorization](https://docs.stripe.com/payments/incremental-authorization) for this CheckoutSession.
	RequestIncrementalAuthorization *string `form:"request_incremental_authorization"`
	// Request ability to make [multiple captures](https://docs.stripe.com/payments/multicapture) for this CheckoutSession.
	RequestMulticapture *string `form:"request_multicapture"`
	// Request ability to [overcapture](https://docs.stripe.com/payments/overcapture) for this CheckoutSession.
	RequestOvercapture *string `form:"request_overcapture"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure"`
	// Restrictions to apply to the card payment method. For example, you can block specific card brands. You can't set this parameter if `ui_mode` is `custom`.
	Restrictions *CheckoutSessionPaymentMethodOptionsCardRestrictionsParams `form:"restrictions"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kana prefix (shortened Kana descriptor) or Kana statement descriptor that's set on the account to form the complete statement descriptor. Maximum 22 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 22 characters.
	StatementDescriptorSuffixKana *string `form:"statement_descriptor_suffix_kana"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kanji prefix (shortened Kanji descriptor) or Kanji statement descriptor that's set on the account to form the complete statement descriptor. Maximum 17 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 17 characters.
	StatementDescriptorSuffixKanji *string `form:"statement_descriptor_suffix_kanji"`
}

// contains details about the Cashapp Pay payment method options.
type CheckoutSessionPaymentMethodOptionsCashAppParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Configuration for eu_bank_transfer funding type.
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information. Permitted values include: `DE`, `FR`, `IE`, or `NL`.
	Country *string `form:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for eu_bank_transfer funding type.
	EUBankTransfer *CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer"`
	// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
	//
	// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
	RequestedAddressTypes []*string `form:"requested_address_types"`
	// The list of bank transfer types that this PaymentIntent is allowed to use for funding.
	Type *string `form:"type"`
}

// contains details about the Customer Balance payment method options.
type CheckoutSessionPaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType *string `form:"funding_type"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the DemoPay payment method options.
type CheckoutSessionPaymentMethodOptionsDemoPayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the EPS payment method options.
type CheckoutSessionPaymentMethodOptionsEPSParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the FPX payment method options.
type CheckoutSessionPaymentMethodOptionsFPXParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Giropay payment method options.
type CheckoutSessionPaymentMethodOptionsGiropayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Grabpay payment method options.
type CheckoutSessionPaymentMethodOptionsGrabpayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Ideal payment method options.
type CheckoutSessionPaymentMethodOptionsIDEALParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Kakao Pay payment method options.
type CheckoutSessionPaymentMethodOptionsKakaoPayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Describes the upcoming charge for this subscription.
type CheckoutSessionPaymentMethodOptionsKlarnaSubscriptionNextBillingParams struct {
	// The amount of the next charge for the subscription.
	Amount *int64 `form:"amount"`
	// The date of the next charge for the subscription in YYYY-MM-DD format.
	Date *string `form:"date"`
}

// Subscription details if the Checkout Session sets up a future subscription.
type CheckoutSessionPaymentMethodOptionsKlarnaSubscriptionParams struct {
	// Unit of time between subscription charges.
	Interval *string `form:"interval"`
	// The number of intervals (specified in the `interval` attribute) between subscription charges. For example, `interval=month` and `interval_count=3` charges every 3 months.
	IntervalCount *int64 `form:"interval_count"`
	// Name for subscription.
	Name *string `form:"name"`
	// Describes the upcoming charge for this subscription.
	NextBilling *CheckoutSessionPaymentMethodOptionsKlarnaSubscriptionNextBillingParams `form:"next_billing"`
	// A non-customer-facing reference to correlate subscription charges in the Klarna app. Use a value that persists across subscription charges.
	Reference *string `form:"reference"`
}

// contains details about the Klarna payment method options.
type CheckoutSessionPaymentMethodOptionsKlarnaParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Subscription details if the Checkout Session sets up a future subscription.
	Subscriptions []*CheckoutSessionPaymentMethodOptionsKlarnaSubscriptionParams `form:"subscriptions"`
}

// contains details about the Konbini payment method options.
type CheckoutSessionPaymentMethodOptionsKonbiniParams struct {
	// The number of calendar days (between 1 and 60) after which Konbini payment instructions will expire. For example, if a PaymentIntent is confirmed with Konbini and `expires_after_days` set to 2 on Monday JST, the instructions will expire on Wednesday 23:59:59 JST. Defaults to 3 days.
	ExpiresAfterDays *int64 `form:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Korean card payment method options.
type CheckoutSessionPaymentMethodOptionsKrCardParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Link payment method options.
type CheckoutSessionPaymentMethodOptionsLinkParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Mobilepay payment method options.
type CheckoutSessionPaymentMethodOptionsMobilepayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Multibanco payment method options.
type CheckoutSessionPaymentMethodOptionsMultibancoParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Naver Pay payment method options.
type CheckoutSessionPaymentMethodOptionsNaverPayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the OXXO payment method options.
type CheckoutSessionPaymentMethodOptionsOXXOParams struct {
	// The number of calendar days before an OXXO voucher expires. For example, if you create an OXXO voucher on Monday and you set expires_after_days to 2, the OXXO invoice will expire on Wednesday at 23:59 America/Mexico_City time.
	ExpiresAfterDays *int64 `form:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the P24 payment method options.
type CheckoutSessionPaymentMethodOptionsP24Params struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Confirm that the payer has accepted the P24 terms and conditions.
	TOSShownAndAccepted *bool `form:"tos_shown_and_accepted"`
}

// contains details about the Pay By Bank payment method options.
type CheckoutSessionPaymentMethodOptionsPayByBankParams struct{}

// contains details about the PAYCO payment method options.
type CheckoutSessionPaymentMethodOptionsPaycoParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
}

// contains details about the PayNow payment method options.
type CheckoutSessionPaymentMethodOptionsPayNowParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the PayPal payment method options.
type CheckoutSessionPaymentMethodOptionsPaypalParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// [Preferred locale](https://docs.stripe.com/payments/paypal/supported-locales) of the PayPal checkout page that the customer is redirected to.
	PreferredLocale *string `form:"preferred_locale"`
	// A reference of the PayPal transaction visible to customer which is mapped to PayPal's invoice ID. This must be a globally unique ID if you have configured in your PayPal settings to block multiple payments per invoice ID.
	Reference *string `form:"reference"`
	// The risk correlation ID for an on-session payment using a saved PayPal payment method.
	RiskCorrelationID *string `form:"risk_correlation_id"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	//
	// If you've already set `setup_future_usage` and you're performing a request using a publishable key, you can only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Additional fields for Mandate creation
type CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsParams struct {
	// Amount that will be collected. It is required when `amount_type` is `fixed`.
	Amount *int64 `form:"amount"`
	// The type of amount that will be collected. The amount charged must be exact or up to the value of `amount` param for `fixed` or `maximum` type respectively. Defaults to `maximum`.
	AmountType *string `form:"amount_type"`
	// Date, in YYYY-MM-DD format, after which payments will not be collected. Defaults to no end date.
	EndDate *string `form:"end_date"`
	// The periodicity at which payments will be collected. Defaults to `adhoc`.
	PaymentSchedule *string `form:"payment_schedule"`
	// The number of payments that will be made during a payment period. Defaults to 1 except for when `payment_schedule` is `adhoc`. In that case, it defaults to no limit.
	PaymentsPerPeriod *int64 `form:"payments_per_period"`
	// The purpose for which payments are made. Has a default value based on your merchant category code.
	Purpose *string `form:"purpose"`
	// Date, in YYYY-MM-DD format, from which payments will be collected. Defaults to confirmation time.
	StartDate *string `form:"start_date"`
}

// contains details about the PayTo payment method options.
type CheckoutSessionPaymentMethodOptionsPaytoParams struct {
	// Additional fields for Mandate creation
	MandateOptions *CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Pix payment method options.
type CheckoutSessionPaymentMethodOptionsPixParams struct {
	// Determines if the amount includes the IOF tax. Defaults to `never`.
	AmountIncludesIof *string `form:"amount_includes_iof"`
	// The number of seconds (between 10 and 1209600) after which Pix payment will expire. Defaults to 86400 seconds.
	ExpiresAfterSeconds *int64 `form:"expires_after_seconds"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the RevolutPay payment method options.
type CheckoutSessionPaymentMethodOptionsRevolutPayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Samsung Pay payment method options.
type CheckoutSessionPaymentMethodOptionsSamsungPayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
}

// contains details about the Satispay payment method options.
type CheckoutSessionPaymentMethodOptionsSatispayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
}

// Additional fields for Mandate creation
type CheckoutSessionPaymentMethodOptionsSEPADebitMandateOptionsParams struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'STRIPE'.
	ReferencePrefix *string `form:"reference_prefix"`
}

// contains details about the Sepa Debit payment method options.
type CheckoutSessionPaymentMethodOptionsSEPADebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *CheckoutSessionPaymentMethodOptionsSEPADebitMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate *string `form:"target_date"`
}

// contains details about the Sofort payment method options.
type CheckoutSessionPaymentMethodOptionsSofortParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Swish payment method options.
type CheckoutSessionPaymentMethodOptionsSwishParams struct {
	// The order reference that will be displayed to customers in the Swish application. Defaults to the `id` of the Payment Intent.
	Reference *string `form:"reference"`
}

// contains details about the TWINT payment method options.
type CheckoutSessionPaymentMethodOptionsTWINTParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Additional fields for Financial Connections Session creation
type CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included. Valid permissions include: `balances`, `ownership`, `payment_method`, and `transactions`.
	Permissions []*string `form:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch"`
}

// contains details about the Us Bank Account payment method options.
type CheckoutSessionPaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation
	FinancialConnections *CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate *string `form:"target_date"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// contains details about the WeChat Pay payment method options.
type CheckoutSessionPaymentMethodOptionsWeChatPayParams struct {
	// The app ID registered with WeChat Pay. Only required when client is ios or android.
	AppID *string `form:"app_id"`
	// The client type that the end customer will pay from
	Client *string `form:"client"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Payment-method-specific configuration.
type CheckoutSessionPaymentMethodOptionsParams struct {
	// contains details about the ACSS Debit payment method options. You can't set this parameter if `ui_mode` is `custom`.
	ACSSDebit *CheckoutSessionPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	// contains details about the Affirm payment method options.
	Affirm *CheckoutSessionPaymentMethodOptionsAffirmParams `form:"affirm"`
	// contains details about the Afterpay Clearpay payment method options.
	AfterpayClearpay *CheckoutSessionPaymentMethodOptionsAfterpayClearpayParams `form:"afterpay_clearpay"`
	// contains details about the Alipay payment method options.
	Alipay *CheckoutSessionPaymentMethodOptionsAlipayParams `form:"alipay"`
	// contains details about the Alma payment method options.
	Alma *CheckoutSessionPaymentMethodOptionsAlmaParams `form:"alma"`
	// contains details about the AmazonPay payment method options.
	AmazonPay *CheckoutSessionPaymentMethodOptionsAmazonPayParams `form:"amazon_pay"`
	// contains details about the AU Becs Debit payment method options.
	AUBECSDebit *CheckoutSessionPaymentMethodOptionsAUBECSDebitParams `form:"au_becs_debit"`
	// contains details about the Bacs Debit payment method options.
	BACSDebit *CheckoutSessionPaymentMethodOptionsBACSDebitParams `form:"bacs_debit"`
	// contains details about the Bancontact payment method options.
	Bancontact *CheckoutSessionPaymentMethodOptionsBancontactParams `form:"bancontact"`
	// contains details about the Billie payment method options.
	Billie *CheckoutSessionPaymentMethodOptionsBillieParams `form:"billie"`
	// contains details about the Boleto payment method options.
	Boleto *CheckoutSessionPaymentMethodOptionsBoletoParams `form:"boleto"`
	// contains details about the Card payment method options.
	Card *CheckoutSessionPaymentMethodOptionsCardParams `form:"card"`
	// contains details about the Cashapp Pay payment method options.
	CashApp *CheckoutSessionPaymentMethodOptionsCashAppParams `form:"cashapp"`
	// contains details about the Customer Balance payment method options.
	CustomerBalance *CheckoutSessionPaymentMethodOptionsCustomerBalanceParams `form:"customer_balance"`
	// contains details about the DemoPay payment method options.
	DemoPay *CheckoutSessionPaymentMethodOptionsDemoPayParams `form:"demo_pay"`
	// contains details about the EPS payment method options.
	EPS *CheckoutSessionPaymentMethodOptionsEPSParams `form:"eps"`
	// contains details about the FPX payment method options.
	FPX *CheckoutSessionPaymentMethodOptionsFPXParams `form:"fpx"`
	// contains details about the Giropay payment method options.
	Giropay *CheckoutSessionPaymentMethodOptionsGiropayParams `form:"giropay"`
	// contains details about the Grabpay payment method options.
	Grabpay *CheckoutSessionPaymentMethodOptionsGrabpayParams `form:"grabpay"`
	// contains details about the Ideal payment method options.
	IDEAL *CheckoutSessionPaymentMethodOptionsIDEALParams `form:"ideal"`
	// contains details about the Kakao Pay payment method options.
	KakaoPay *CheckoutSessionPaymentMethodOptionsKakaoPayParams `form:"kakao_pay"`
	// contains details about the Klarna payment method options.
	Klarna *CheckoutSessionPaymentMethodOptionsKlarnaParams `form:"klarna"`
	// contains details about the Konbini payment method options.
	Konbini *CheckoutSessionPaymentMethodOptionsKonbiniParams `form:"konbini"`
	// contains details about the Korean card payment method options.
	KrCard *CheckoutSessionPaymentMethodOptionsKrCardParams `form:"kr_card"`
	// contains details about the Link payment method options.
	Link *CheckoutSessionPaymentMethodOptionsLinkParams `form:"link"`
	// contains details about the Mobilepay payment method options.
	Mobilepay *CheckoutSessionPaymentMethodOptionsMobilepayParams `form:"mobilepay"`
	// contains details about the Multibanco payment method options.
	Multibanco *CheckoutSessionPaymentMethodOptionsMultibancoParams `form:"multibanco"`
	// contains details about the Naver Pay payment method options.
	NaverPay *CheckoutSessionPaymentMethodOptionsNaverPayParams `form:"naver_pay"`
	// contains details about the OXXO payment method options.
	OXXO *CheckoutSessionPaymentMethodOptionsOXXOParams `form:"oxxo"`
	// contains details about the P24 payment method options.
	P24 *CheckoutSessionPaymentMethodOptionsP24Params `form:"p24"`
	// contains details about the Pay By Bank payment method options.
	PayByBank *CheckoutSessionPaymentMethodOptionsPayByBankParams `form:"pay_by_bank"`
	// contains details about the PAYCO payment method options.
	Payco *CheckoutSessionPaymentMethodOptionsPaycoParams `form:"payco"`
	// contains details about the PayNow payment method options.
	PayNow *CheckoutSessionPaymentMethodOptionsPayNowParams `form:"paynow"`
	// contains details about the PayPal payment method options.
	Paypal *CheckoutSessionPaymentMethodOptionsPaypalParams `form:"paypal"`
	// contains details about the PayTo payment method options.
	Payto *CheckoutSessionPaymentMethodOptionsPaytoParams `form:"payto"`
	// contains details about the Pix payment method options.
	Pix *CheckoutSessionPaymentMethodOptionsPixParams `form:"pix"`
	// contains details about the RevolutPay payment method options.
	RevolutPay *CheckoutSessionPaymentMethodOptionsRevolutPayParams `form:"revolut_pay"`
	// contains details about the Samsung Pay payment method options.
	SamsungPay *CheckoutSessionPaymentMethodOptionsSamsungPayParams `form:"samsung_pay"`
	// contains details about the Satispay payment method options.
	Satispay *CheckoutSessionPaymentMethodOptionsSatispayParams `form:"satispay"`
	// contains details about the Sepa Debit payment method options.
	SEPADebit *CheckoutSessionPaymentMethodOptionsSEPADebitParams `form:"sepa_debit"`
	// contains details about the Sofort payment method options.
	Sofort *CheckoutSessionPaymentMethodOptionsSofortParams `form:"sofort"`
	// contains details about the Swish payment method options.
	Swish *CheckoutSessionPaymentMethodOptionsSwishParams `form:"swish"`
	// contains details about the TWINT payment method options.
	TWINT *CheckoutSessionPaymentMethodOptionsTWINTParams `form:"twint"`
	// contains details about the Us Bank Account payment method options.
	USBankAccount *CheckoutSessionPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account"`
	// contains details about the WeChat Pay payment method options.
	WeChatPay *CheckoutSessionPaymentMethodOptionsWeChatPayParams `form:"wechat_pay"`
}

// This property is used to set up permissions for various actions (e.g., update) on the CheckoutSession object. Can only be set when creating `embedded` or `custom` sessions.
//
// For specific permissions, please refer to their dedicated subsections, such as `permissions.update_shipping_details`.
type CheckoutSessionPermissionsParams struct {
	// Determines which entity is allowed to update the shipping details.
	//
	// Default is `client_only`. Stripe Checkout client will automatically update the shipping details. If set to `server_only`, only your server is allowed to update the shipping details.
	//
	// When set to `server_only`, you must add the onShippingDetailsChange event handler when initializing the Stripe Checkout client and manually update the shipping details from your server using the Stripe API.
	UpdateShippingDetails *string `form:"update_shipping_details"`
}

// Controls phone number collection settings for the session.
//
// We recommend that you review your privacy policy and check with your legal contacts
// before using this feature. Learn more about [collecting phone numbers with Checkout](https://docs.stripe.com/payments/checkout/phone-numbers).
type CheckoutSessionPhoneNumberCollectionParams struct {
	// Set to `true` to enable phone number collection.
	//
	// Can only be set in `payment` and `subscription` mode.
	Enabled *bool `form:"enabled"`
}

// Controls saved payment method settings for the session. Only available in `payment` and `subscription` mode.
type CheckoutSessionSavedPaymentMethodOptionsParams struct {
	// Uses the `allow_redisplay` value of each saved payment method to filter the set presented to a returning customer. By default, only saved payment methods with 'allow_redisplay: ‘always' are shown in Checkout.
	AllowRedisplayFilters []*string `form:"allow_redisplay_filters"`
	// Enable customers to choose if they wish to remove their saved payment methods. Disabled by default.
	PaymentMethodRemove *string `form:"payment_method_remove"`
	// Enable customers to choose if they wish to save their payment method for future use. Disabled by default.
	PaymentMethodSave *string `form:"payment_method_save"`
}

// A subset of parameters to be passed to SetupIntent creation for Checkout Sessions in `setup` mode.
type CheckoutSessionSetupIntentDataParams struct {
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The Stripe account for which the setup is intended.
	OnBehalfOf *string `form:"on_behalf_of"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionSetupIntentDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// When set, provides configuration for Checkout to collect a shipping address from a customer.
type CheckoutSessionShippingAddressCollectionParams struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for
	// shipping locations.
	AllowedCountries []*string `form:"allowed_countries"`
}

// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
type CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMaximumParams struct {
	// A unit of time.
	Unit *string `form:"unit"`
	// Must be greater than 0.
	Value *int64 `form:"value"`
}

// The lower bound of the estimated range. If empty, represents no lower bound.
type CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMinimumParams struct {
	// A unit of time.
	Unit *string `form:"unit"`
	// Must be greater than 0.
	Value *int64 `form:"value"`
}

// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
type CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateParams struct {
	// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
	Maximum *CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMaximumParams `form:"maximum"`
	// The lower bound of the estimated range. If empty, represents no lower bound.
	Minimum *CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMinimumParams `form:"minimum"`
}

// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type CheckoutSessionShippingOptionShippingRateDataFixedAmountCurrencyOptionsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
type CheckoutSessionShippingOptionShippingRateDataFixedAmountParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions map[string]*CheckoutSessionShippingOptionShippingRateDataFixedAmountCurrencyOptionsParams `form:"currency_options"`
}

// Parameters to be passed to Shipping Rate creation for this shipping option.
type CheckoutSessionShippingOptionShippingRateDataParams struct {
	// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DeliveryEstimate *CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateParams `form:"delivery_estimate"`
	// The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DisplayName *string `form:"display_name"`
	// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
	FixedAmount *CheckoutSessionShippingOptionShippingRateDataFixedAmountParams `form:"fixed_amount"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.
	TaxCode *string `form:"tax_code"`
	// The type of calculation to use on the shipping rate.
	Type *string `form:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionShippingOptionShippingRateDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The shipping rate options to apply to this Session. Up to a maximum of 5.
type CheckoutSessionShippingOptionParams struct {
	// The ID of the Shipping Rate to use for this shipping option.
	ShippingRate *string `form:"shipping_rate"`
	// Parameters to be passed to Shipping Rate creation for this shipping option.
	ShippingRateData *CheckoutSessionShippingOptionShippingRateDataParams `form:"shipping_rate_data"`
}

// Configure behavior for flexible billing mode.
type CheckoutSessionSubscriptionDataBillingModeFlexibleParams struct {
	// Controls how invoices and invoice items display proration amounts and discount amounts.
	ProrationDiscounts *string `form:"proration_discounts"`
}

// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
type CheckoutSessionSubscriptionDataBillingModeParams struct {
	// Configure behavior for flexible billing mode.
	Flexible *CheckoutSessionSubscriptionDataBillingModeFlexibleParams `form:"flexible"`
	// Controls the calculation and orchestration of prorations and invoices for subscriptions. If no value is passed, the default is `flexible`.
	Type *string `form:"type"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type CheckoutSessionSubscriptionDataInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type CheckoutSessionSubscriptionDataInvoiceSettingsParams struct {
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *CheckoutSessionSubscriptionDataInvoiceSettingsIssuerParams `form:"issuer"`
}

// If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.
type CheckoutSessionSubscriptionDataTransferDataParams struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount is transferred to the destination.
	AmountPercent *float64 `form:"amount_percent"`
	// ID of an existing, connected Stripe account.
	Destination *string `form:"destination"`
}

// Defines how the subscription should behave when the user's free trial ends.
type CheckoutSessionSubscriptionDataTrialSettingsEndBehaviorParams struct {
	// Indicates how the subscription should change when the trial ends if the user did not provide a payment method.
	MissingPaymentMethod *string `form:"missing_payment_method"`
}

// Settings related to subscription trials.
type CheckoutSessionSubscriptionDataTrialSettingsParams struct {
	// Defines how the subscription should behave when the user's free trial ends.
	EndBehavior *CheckoutSessionSubscriptionDataTrialSettingsEndBehaviorParams `form:"end_behavior"`
}

// A subset of parameters to be passed to subscription creation for Checkout Sessions in `subscription` mode.
type CheckoutSessionSubscriptionDataParams struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. To use an application fee percent, the request must be made on behalf of another account, using the `Stripe-Account` header or an OAuth key. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// A future timestamp to anchor the subscription's billing cycle for new subscriptions. You can't set this parameter if `ui_mode` is `custom`.
	BillingCycleAnchor *int64 `form:"billing_cycle_anchor"`
	// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
	BillingMode *CheckoutSessionSubscriptionDataBillingModeParams `form:"billing_mode"`
	// The tax rates that will apply to any subscription item that does not have
	// `tax_rates` set. Invoices created will have their `default_tax_rates` populated
	// from the subscription.
	DefaultTaxRates []*string `form:"default_tax_rates"`
	// The subscription's description, meant to be displayable to the customer.
	// Use this field to optionally store an explanation of the subscription
	// for rendering in the [customer portal](https://docs.stripe.com/customer-management).
	Description *string `form:"description"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *CheckoutSessionSubscriptionDataInvoiceSettingsParams `form:"invoice_settings"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The account on behalf of which to charge, for each of the subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of"`
	// Determines how to handle prorations resulting from the `billing_cycle_anchor`. If no value is passed, the default is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior"`
	// If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.
	TransferData *CheckoutSessionSubscriptionDataTransferDataParams `form:"transfer_data"`
	// Unix timestamp representing the end of the trial period the customer will get before being charged for the first time. Has to be at least 48 hours in the future.
	TrialEnd *int64 `form:"trial_end"`
	// Integer representing the number of trial period days before the customer is charged for the first time. Has to be at least 1.
	TrialPeriodDays *int64 `form:"trial_period_days"`
	// Settings related to subscription trials.
	TrialSettings *CheckoutSessionSubscriptionDataTrialSettingsParams `form:"trial_settings"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionSubscriptionDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Controls tax ID collection during checkout.
type CheckoutSessionTaxIDCollectionParams struct {
	// Enable tax ID collection during checkout. Defaults to `false`.
	Enabled *bool `form:"enabled"`
	// Describes whether a tax ID is required during checkout. Defaults to `never`. You can't set this parameter if `ui_mode` is `custom`.
	Required *string `form:"required"`
}

// contains details about the Link wallet options.
type CheckoutSessionWalletOptionsLinkParams struct {
	// Specifies whether Checkout should display Link as a payment option. By default, Checkout will display all the supported wallets that the Checkout Session was created with. This is the `auto` behavior, and it is the default choice.
	Display *string `form:"display"`
}

// Wallet-specific configuration.
type CheckoutSessionWalletOptionsParams struct {
	// contains details about the Link wallet options.
	Link *CheckoutSessionWalletOptionsLinkParams `form:"link"`
}

// Creates a Checkout Session object.
type CheckoutSessionParams struct {
	Params `form:"*"`
	// Settings for price localization with [Adaptive Pricing](https://docs.stripe.com/payments/checkout/adaptive-pricing).
	AdaptivePricing *CheckoutSessionAdaptivePricingParams `form:"adaptive_pricing"`
	// Configure actions after a Checkout Session has expired. You can't set this parameter if `ui_mode` is `custom`.
	AfterExpiration *CheckoutSessionAfterExpirationParams `form:"after_expiration"`
	// Enables user redeemable promotion codes.
	AllowPromotionCodes *bool `form:"allow_promotion_codes"`
	// Settings for automatic tax lookup for this session and resulting payments, invoices, and subscriptions.
	AutomaticTax *CheckoutSessionAutomaticTaxParams `form:"automatic_tax"`
	// Specify whether Checkout should collect the customer's billing address. Defaults to `auto`.
	BillingAddressCollection *string `form:"billing_address_collection"`
	// The branding settings for the Checkout Session. This parameter is not allowed if ui_mode is `custom`.
	BrandingSettings *CheckoutSessionBrandingSettingsParams `form:"branding_settings"`
	// If set, Checkout displays a back button and customers will be directed to this URL if they decide to cancel payment and return to your website. This parameter is not allowed if ui_mode is `embedded` or `custom`.
	CancelURL *string `form:"cancel_url"`
	// A unique string to reference the Checkout Session. This can be a
	// customer ID, a cart ID, or similar, and can be used to reconcile the
	// session with your internal systems.
	ClientReferenceID *string `form:"client_reference_id"`
	// Information about the customer collected within the Checkout Session. Can only be set when updating `embedded` or `custom` sessions.
	CollectedInformation *CheckoutSessionCollectedInformationParams `form:"collected_information"`
	// Configure fields for the Checkout Session to gather active consent from customers.
	ConsentCollection *CheckoutSessionConsentCollectionParams `form:"consent_collection"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies). Required in `setup` mode when `payment_method_types` is not set.
	Currency *string `form:"currency"`
	// ID of an existing Customer, if one exists. In `payment` mode, the customer's most recently saved card
	// payment method will be used to prefill the email, name, card details, and billing address
	// on the Checkout page. In `subscription` mode, the customer's [default payment method](https://docs.stripe.com/api/customers/update#update_customer-invoice_settings-default_payment_method)
	// will be used if it's a card, otherwise the most recently saved card will be used. A valid billing address, billing name and billing email are required on the payment method for Checkout to prefill the customer's card details.
	//
	// If the Customer already has a valid [email](https://docs.stripe.com/api/customers/object#customer_object-email) set, the email will be prefilled and not editable in Checkout.
	// If the Customer does not have a valid `email`, Checkout will set the email entered during the session on the Customer.
	//
	// If blank for Checkout Sessions in `subscription` mode or with `customer_creation` set as `always` in `payment` mode, Checkout will create a new Customer object based on information provided during the payment flow.
	//
	// You can set [`payment_intent_data.setup_future_usage`](https://docs.stripe.com/api/checkout/sessions/create#create_checkout_session-payment_intent_data-setup_future_usage) to have Checkout automatically attach the payment method to the Customer you pass in for future reuse.
	Customer *string `form:"customer"`
	// ID of an existing Account, if one exists. Has the same behavior as `customer`.
	CustomerAccount *string `form:"customer_account"`
	// Configure whether a Checkout Session creates a [Customer](https://docs.stripe.com/api/customers) during Session confirmation.
	//
	// When a Customer is not created, you can still retrieve email, address, and other customer data entered in Checkout
	// with [customer_details](https://docs.stripe.com/api/checkout/sessions/object#checkout_session_object-customer_details).
	//
	// Sessions that don't create Customers instead are grouped by [guest customers](https://docs.stripe.com/payments/checkout/guest-customers)
	// in the Dashboard. Promotion codes limited to first time customers will return invalid for these Sessions.
	//
	// Can only be set in `payment` and `setup` mode.
	CustomerCreation *string `form:"customer_creation"`
	// If provided, this value will be used when the Customer object is created.
	// If not provided, customers will be asked to enter their email address.
	// Use this parameter to prefill customer data if you already have an email
	// on file. To access information about the customer once a session is
	// complete, use the `customer` field.
	CustomerEmail *string `form:"customer_email"`
	// Controls what fields on Customer can be updated by the Checkout Session. Can only be provided when `customer` is provided.
	CustomerUpdate *CheckoutSessionCustomerUpdateParams `form:"customer_update"`
	// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
	CustomFields []*CheckoutSessionCustomFieldParams `form:"custom_fields"`
	// Display additional text for your customers using custom text. You can't set this parameter if `ui_mode` is `custom`.
	CustomText *CheckoutSessionCustomTextParams `form:"custom_text"`
	// The coupon or promotion code to apply to this Session. Currently, only up to one may be specified.
	Discounts []*CheckoutSessionDiscountParams `form:"discounts"`
	// A list of the types of payment methods (e.g., `card`) that should be excluded from this Checkout Session. This should only be used when payment methods for this Checkout Session are managed through the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).
	ExcludedPaymentMethodTypes []*string `form:"excluded_payment_method_types"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The Epoch time in seconds at which the Checkout Session will expire. It can be anywhere from 30 minutes to 24 hours after Checkout Session creation. By default, this value is 24 hours from creation.
	ExpiresAt *int64 `form:"expires_at"`
	// Generate a post-purchase Invoice for one-time payments.
	InvoiceCreation *CheckoutSessionInvoiceCreationParams `form:"invoice_creation"`
	// A list of items the customer is purchasing. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices). The parameter is required for `payment` and `subscription` mode.
	//
	// For `payment` mode, there is a maximum of 100 line items, however it is recommended to consolidate line items if there are more than a few dozen.
	//
	// For `subscription` mode, there is a maximum of 20 line items with recurring Prices and 20 line items with one-time Prices. Line items with one-time Prices will be on the initial invoice only.
	LineItems []*CheckoutSessionLineItemParams `form:"line_items"`
	// The IETF language tag of the locale Checkout is displayed in. If blank or `auto`, the browser's locale is used.
	Locale *string `form:"locale"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The mode of the Checkout Session. Pass `subscription` if the Checkout Session includes at least one recurring item.
	Mode *string `form:"mode"`
	// Controls name collection settings for the session.
	//
	// You can configure Checkout to collect your customers' business names, individual names, or both. Each name field can be either required or optional.
	//
	// If a [Customer](https://docs.stripe.com/api/customers) is created or provided, the names can be saved to the Customer object as well.
	//
	// You can't set this parameter if `ui_mode` is `custom`.
	NameCollection *CheckoutSessionNameCollectionParams `form:"name_collection"`
	// A list of optional items the customer can add to their order at checkout. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices).
	//
	// There is a maximum of 10 optional items allowed on a Checkout Session, and the existing limits on the number of line items allowed on a Checkout Session apply to the combined number of line items and optional items.
	//
	// For `payment` mode, there is a maximum of 100 combined line items and optional items, however it is recommended to consolidate items if there are more than a few dozen.
	//
	// For `subscription` mode, there is a maximum of 20 line items and optional items with recurring Prices and 20 line items and optional items with one-time Prices.
	//
	// You can't set this parameter if `ui_mode` is `custom`.
	OptionalItems []*CheckoutSessionOptionalItemParams `form:"optional_items"`
	// Where the user is coming from. This informs the optimizations that are applied to the session. You can't set this parameter if `ui_mode` is `custom`.
	OriginContext *string `form:"origin_context"`
	// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
	PaymentIntentData *CheckoutSessionPaymentIntentDataParams `form:"payment_intent_data"`
	// Specify whether Checkout should collect a payment method. When set to `if_required`, Checkout will not collect a payment method when the total due for the session is 0.
	// This may occur if the Checkout Session includes a free trial or a discount.
	//
	// Can only be set in `subscription` mode. Defaults to `always`.
	//
	// If you'd like information on how to collect a payment method outside of Checkout, read the guide on configuring [subscriptions with a free trial](https://docs.stripe.com/payments/checkout/free-trials).
	PaymentMethodCollection *string `form:"payment_method_collection"`
	// The ID of the payment method configuration to use with this Checkout session.
	PaymentMethodConfiguration *string `form:"payment_method_configuration"`
	// This parameter allows you to set some attributes on the payment method created during a Checkout session.
	PaymentMethodData *CheckoutSessionPaymentMethodDataParams `form:"payment_method_data"`
	// Payment-method-specific configuration.
	PaymentMethodOptions *CheckoutSessionPaymentMethodOptionsParams `form:"payment_method_options"`
	// A list of the types of payment methods (e.g., `card`) this Checkout Session can accept.
	//
	// You can omit this attribute to manage your payment methods from the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).
	// See [Dynamic Payment Methods](https://docs.stripe.com/payments/payment-methods/integration-options#using-dynamic-payment-methods) for more details.
	//
	// Read more about the supported payment methods and their requirements in our [payment
	// method details guide](https://docs.stripe.com/docs/payments/checkout/payment-methods).
	//
	// If multiple payment methods are passed, Checkout will dynamically reorder them to
	// prioritize the most relevant payment methods based on the customer's location and
	// other characteristics.
	PaymentMethodTypes []*string `form:"payment_method_types"`
	// This property is used to set up permissions for various actions (e.g., update) on the CheckoutSession object. Can only be set when creating `embedded` or `custom` sessions.
	//
	// For specific permissions, please refer to their dedicated subsections, such as `permissions.update_shipping_details`.
	Permissions *CheckoutSessionPermissionsParams `form:"permissions"`
	// Controls phone number collection settings for the session.
	//
	// We recommend that you review your privacy policy and check with your legal contacts
	// before using this feature. Learn more about [collecting phone numbers with Checkout](https://docs.stripe.com/payments/checkout/phone-numbers).
	PhoneNumberCollection *CheckoutSessionPhoneNumberCollectionParams `form:"phone_number_collection"`
	// This parameter applies to `ui_mode: embedded`. Learn more about the [redirect behavior](https://docs.stripe.com/payments/checkout/custom-success-page?payment-ui=embedded-form) of embedded sessions. Defaults to `always`.
	RedirectOnCompletion *string `form:"redirect_on_completion"`
	// The URL to redirect your customer back to after they authenticate or cancel their payment on the
	// payment method's app or site. This parameter is required if `ui_mode` is `embedded` or `custom`
	// and redirect-based payment methods are enabled on the session.
	ReturnURL *string `form:"return_url"`
	// Controls saved payment method settings for the session. Only available in `payment` and `subscription` mode.
	SavedPaymentMethodOptions *CheckoutSessionSavedPaymentMethodOptionsParams `form:"saved_payment_method_options"`
	// A subset of parameters to be passed to SetupIntent creation for Checkout Sessions in `setup` mode.
	SetupIntentData *CheckoutSessionSetupIntentDataParams `form:"setup_intent_data"`
	// When set, provides configuration for Checkout to collect a shipping address from a customer.
	ShippingAddressCollection *CheckoutSessionShippingAddressCollectionParams `form:"shipping_address_collection"`
	// The shipping rate options to apply to this Session. Up to a maximum of 5.
	ShippingOptions []*CheckoutSessionShippingOptionParams `form:"shipping_options"`
	// Describes the type of transaction being performed by Checkout in order
	// to customize relevant text on the page, such as the submit button.
	//  `submit_type` can only be specified on Checkout Sessions in
	// `payment` or `subscription` mode. If blank or `auto`, `pay` is used.
	// You can't set this parameter if `ui_mode` is `custom`.
	SubmitType *string `form:"submit_type"`
	// A subset of parameters to be passed to subscription creation for Checkout Sessions in `subscription` mode.
	SubscriptionData *CheckoutSessionSubscriptionDataParams `form:"subscription_data"`
	// The URL to which Stripe should send customers when payment or setup
	// is complete.
	// This parameter is not allowed if ui_mode is `embedded` or `custom`. If you'd like to use
	// information from the successful Checkout Session on your page, read the
	// guide on [customizing your success page](https://docs.stripe.com/payments/checkout/custom-success-page).
	SuccessURL *string `form:"success_url"`
	// Controls tax ID collection during checkout.
	TaxIDCollection *CheckoutSessionTaxIDCollectionParams `form:"tax_id_collection"`
	// The UI mode of the Session. Defaults to `hosted`.
	UIMode *string `form:"ui_mode"`
	// Wallet-specific configuration.
	WalletOptions *CheckoutSessionWalletOptionsParams `form:"wallet_options"`
}

// AddExpand appends a new field to expand.
func (p *CheckoutSessionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The shipping details to apply to this Session.
type CheckoutSessionCollectedInformationShippingDetailsParams struct {
	// The address of the customer
	Address *AddressParams `form:"address"`
	// The name of customer
	Name *string `form:"name"`
}

// Information about the customer collected within the Checkout Session. Can only be set when updating `embedded` or `custom` sessions.
type CheckoutSessionCollectedInformationParams struct {
	// The shipping details to apply to this Session.
	ShippingDetails *CheckoutSessionCollectedInformationShippingDetailsParams `form:"shipping_details"`
}

// When retrieving a Checkout Session, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type CheckoutSessionListLineItemsParams struct {
	ListParams `form:"*"`
	Session    *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CheckoutSessionListLineItemsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A Checkout Session can be expired when it is in one of these statuses: open
//
// After it expires, a customer can't complete a Checkout Session and customers loading the Checkout Session see a message saying the Checkout Session is expired.
type CheckoutSessionExpireParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CheckoutSessionExpireParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Settings for price localization with [Adaptive Pricing](https://docs.stripe.com/payments/checkout/adaptive-pricing).
type CheckoutSessionCreateAdaptivePricingParams struct {
	// If set to `true`, Adaptive Pricing is available on [eligible sessions](https://docs.stripe.com/payments/currencies/localize-prices/adaptive-pricing?payment-ui=stripe-hosted#restrictions). Defaults to your [dashboard setting](https://dashboard.stripe.com/settings/adaptive-pricing).
	Enabled *bool `form:"enabled"`
}

// Configure a Checkout Session that can be used to recover an expired session.
type CheckoutSessionCreateAfterExpirationRecoveryParams struct {
	// Enables user redeemable promotion codes on the recovered Checkout Sessions. Defaults to `false`
	AllowPromotionCodes *bool `form:"allow_promotion_codes"`
	// If `true`, a recovery URL will be generated to recover this Checkout Session if it
	// expires before a successful transaction is completed. It will be attached to the
	// Checkout Session object upon expiration.
	Enabled *bool `form:"enabled"`
}

// Configure actions after a Checkout Session has expired. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionCreateAfterExpirationParams struct {
	// Configure a Checkout Session that can be used to recover an expired session.
	Recovery *CheckoutSessionCreateAfterExpirationRecoveryParams `form:"recovery"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type CheckoutSessionCreateAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// Settings for automatic tax lookup for this session and resulting payments, invoices, and subscriptions.
type CheckoutSessionCreateAutomaticTaxParams struct {
	// Set to `true` to [calculate tax automatically](https://docs.stripe.com/tax) using the customer's location.
	//
	// Enabling this parameter causes Checkout to collect any billing address information necessary for tax calculation.
	Enabled *bool `form:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *CheckoutSessionCreateAutomaticTaxLiabilityParams `form:"liability"`
}

// The icon for the Checkout Session. For best results, use a square image.
type CheckoutSessionCreateBrandingSettingsIconParams struct {
	// The ID of a [File upload](https://stripe.com/docs/api/files) representing the icon. Purpose must be `business_icon`. Required if `type` is `file` and disallowed otherwise.
	File *string `form:"file"`
	// The type of image for the icon. Must be one of `file` or `url`.
	Type *string `form:"type"`
	// The URL of the image. Required if `type` is `url` and disallowed otherwise.
	URL *string `form:"url"`
}

// The logo for the Checkout Session.
type CheckoutSessionCreateBrandingSettingsLogoParams struct {
	// The ID of a [File upload](https://stripe.com/docs/api/files) representing the logo. Purpose must be `business_logo`. Required if `type` is `file` and disallowed otherwise.
	File *string `form:"file"`
	// The type of image for the logo. Must be one of `file` or `url`.
	Type *string `form:"type"`
	// The URL of the image. Required if `type` is `url` and disallowed otherwise.
	URL *string `form:"url"`
}

// The branding settings for the Checkout Session. This parameter is not allowed if ui_mode is `custom`.
type CheckoutSessionCreateBrandingSettingsParams struct {
	// A hex color value starting with `#` representing the background color for the Checkout Session.
	BackgroundColor *string `form:"background_color"`
	// The border style for the Checkout Session.
	BorderStyle *string `form:"border_style"`
	// A hex color value starting with `#` representing the button color for the Checkout Session.
	ButtonColor *string `form:"button_color"`
	// A string to override the business name shown on the Checkout Session. This only shows at the top of the Checkout page, and your business name still appears in terms, receipts, and other places.
	DisplayName *string `form:"display_name"`
	// The font family for the Checkout Session corresponding to one of the [supported font families](https://docs.stripe.com/payments/checkout/customization/appearance?payment-ui=stripe-hosted#font-compatibility).
	FontFamily *string `form:"font_family"`
	// The icon for the Checkout Session. For best results, use a square image.
	Icon *CheckoutSessionCreateBrandingSettingsIconParams `form:"icon"`
	// The logo for the Checkout Session.
	Logo *CheckoutSessionCreateBrandingSettingsLogoParams `form:"logo"`
}

// Determines the display of payment method reuse agreement text in the UI. If set to `hidden`, it will hide legal text related to the reuse of a payment method.
type CheckoutSessionCreateConsentCollectionPaymentMethodReuseAgreementParams struct {
	// Determines the position and visibility of the payment method reuse agreement in the UI. When set to `auto`, Stripe's
	// defaults will be used. When set to `hidden`, the payment method reuse agreement text will always be hidden in the UI.
	Position *string `form:"position"`
}

// Configure fields for the Checkout Session to gather active consent from customers.
type CheckoutSessionCreateConsentCollectionParams struct {
	// Determines the display of payment method reuse agreement text in the UI. If set to `hidden`, it will hide legal text related to the reuse of a payment method.
	PaymentMethodReuseAgreement *CheckoutSessionCreateConsentCollectionPaymentMethodReuseAgreementParams `form:"payment_method_reuse_agreement"`
	// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
	// Session will determine whether to display an option to opt into promotional communication
	// from the merchant depending on the customer's locale. Only available to US merchants.
	Promotions *string `form:"promotions"`
	// If set to `required`, it requires customers to check a terms of service checkbox before being able to pay.
	// There must be a valid terms of service URL set in your [Dashboard settings](https://dashboard.stripe.com/settings/public).
	TermsOfService *string `form:"terms_of_service"`
}

// The options available for the customer to select. Up to 200 options allowed.
type CheckoutSessionCreateCustomFieldDropdownOptionParams struct {
	// The label for the option, displayed to the customer. Up to 100 characters.
	Label *string `form:"label"`
	// The value for this option, not displayed to the customer, used by your integration to reconcile the option selected by the customer. Must be unique to this option, alphanumeric, and up to 100 characters.
	Value *string `form:"value"`
}

// Configuration for `type=dropdown` fields.
type CheckoutSessionCreateCustomFieldDropdownParams struct {
	// The value that pre-fills the field on the payment page.Must match a `value` in the `options` array.
	DefaultValue *string `form:"default_value"`
	// The options available for the customer to select. Up to 200 options allowed.
	Options []*CheckoutSessionCreateCustomFieldDropdownOptionParams `form:"options"`
}

// The label for the field, displayed to the customer.
type CheckoutSessionCreateCustomFieldLabelParams struct {
	// Custom text for the label, displayed to the customer. Up to 50 characters.
	Custom *string `form:"custom"`
	// The type of the label.
	Type *string `form:"type"`
}

// Configuration for `type=numeric` fields.
type CheckoutSessionCreateCustomFieldNumericParams struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue *string `form:"default_value"`
	// The maximum character length constraint for the customer's input.
	MaximumLength *int64 `form:"maximum_length"`
	// The minimum character length requirement for the customer's input.
	MinimumLength *int64 `form:"minimum_length"`
}

// Configuration for `type=text` fields.
type CheckoutSessionCreateCustomFieldTextParams struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue *string `form:"default_value"`
	// The maximum character length constraint for the customer's input.
	MaximumLength *int64 `form:"maximum_length"`
	// The minimum character length requirement for the customer's input.
	MinimumLength *int64 `form:"minimum_length"`
}

// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionCreateCustomFieldParams struct {
	// Configuration for `type=dropdown` fields.
	Dropdown *CheckoutSessionCreateCustomFieldDropdownParams `form:"dropdown"`
	// String of your choice that your integration can use to reconcile this field. Must be unique to this field, alphanumeric, and up to 200 characters.
	Key *string `form:"key"`
	// The label for the field, displayed to the customer.
	Label *CheckoutSessionCreateCustomFieldLabelParams `form:"label"`
	// Configuration for `type=numeric` fields.
	Numeric *CheckoutSessionCreateCustomFieldNumericParams `form:"numeric"`
	// Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.
	Optional *bool `form:"optional"`
	// Configuration for `type=text` fields.
	Text *CheckoutSessionCreateCustomFieldTextParams `form:"text"`
	// The type of the field.
	Type *string `form:"type"`
}

// Custom text that should be displayed after the payment confirmation button.
type CheckoutSessionCreateCustomTextAfterSubmitParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message"`
}

// Custom text that should be displayed alongside shipping address collection.
type CheckoutSessionCreateCustomTextShippingAddressParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message"`
}

// Custom text that should be displayed alongside the payment confirmation button.
type CheckoutSessionCreateCustomTextSubmitParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message"`
}

// Custom text that should be displayed in place of the default terms of service agreement text.
type CheckoutSessionCreateCustomTextTermsOfServiceAcceptanceParams struct {
	// Text can be up to 1200 characters in length.
	Message *string `form:"message"`
}

// Display additional text for your customers using custom text. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionCreateCustomTextParams struct {
	// Custom text that should be displayed after the payment confirmation button.
	AfterSubmit *CheckoutSessionCreateCustomTextAfterSubmitParams `form:"after_submit"`
	// Custom text that should be displayed alongside shipping address collection.
	ShippingAddress *CheckoutSessionCreateCustomTextShippingAddressParams `form:"shipping_address"`
	// Custom text that should be displayed alongside the payment confirmation button.
	Submit *CheckoutSessionCreateCustomTextSubmitParams `form:"submit"`
	// Custom text that should be displayed in place of the default terms of service agreement text.
	TermsOfServiceAcceptance *CheckoutSessionCreateCustomTextTermsOfServiceAcceptanceParams `form:"terms_of_service_acceptance"`
}

// Controls what fields on Customer can be updated by the Checkout Session. Can only be provided when `customer` is provided.
type CheckoutSessionCreateCustomerUpdateParams struct {
	// Describes whether Checkout saves the billing address onto `customer.address`.
	// To always collect a full billing address, use `billing_address_collection`. Defaults to `never`.
	Address *string `form:"address"`
	// Describes whether Checkout saves the name onto `customer.name`. Defaults to `never`.
	Name *string `form:"name"`
	// Describes whether Checkout saves shipping information onto `customer.shipping`.
	// To collect shipping information, use `shipping_address_collection`. Defaults to `never`.
	Shipping *string `form:"shipping"`
}

// The coupon or promotion code to apply to this Session. Currently, only up to one may be specified.
type CheckoutSessionCreateDiscountParams struct {
	// The ID of the coupon to apply to this Session.
	Coupon *string `form:"coupon"`
	// The ID of a promotion code to apply to this Session.
	PromotionCode *string `form:"promotion_code"`
}

// Default custom fields to be displayed on invoices for this customer.
type CheckoutSessionCreateInvoiceCreationInvoiceDataCustomFieldParams struct {
	// The name of the custom field. This may be up to 40 characters.
	Name *string `form:"name"`
	// The value of the custom field. This may be up to 140 characters.
	Value *string `form:"value"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type CheckoutSessionCreateInvoiceCreationInvoiceDataIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// Default options for invoice PDF rendering for this customer.
type CheckoutSessionCreateInvoiceCreationInvoiceDataRenderingOptionsParams struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs. One of `exclude_tax` or `include_inclusive_tax`. `include_inclusive_tax` will include inclusive tax (and exclude exclusive tax) in invoice PDF amounts. `exclude_tax` will exclude all tax (inclusive and exclusive alike) from invoice PDF amounts.
	AmountTaxDisplay *string `form:"amount_tax_display"`
	// ID of the invoice rendering template to use for this invoice.
	Template *string `form:"template"`
}

// Parameters passed when creating invoices for payment-mode Checkout Sessions.
type CheckoutSessionCreateInvoiceCreationInvoiceDataParams struct {
	// The account tax IDs associated with the invoice.
	AccountTaxIDs []*string `form:"account_tax_ids"`
	// Default custom fields to be displayed on invoices for this customer.
	CustomFields []*CheckoutSessionCreateInvoiceCreationInvoiceDataCustomFieldParams `form:"custom_fields"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Default footer to be displayed on invoices for this customer.
	Footer *string `form:"footer"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *CheckoutSessionCreateInvoiceCreationInvoiceDataIssuerParams `form:"issuer"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Default options for invoice PDF rendering for this customer.
	RenderingOptions *CheckoutSessionCreateInvoiceCreationInvoiceDataRenderingOptionsParams `form:"rendering_options"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionCreateInvoiceCreationInvoiceDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Generate a post-purchase Invoice for one-time payments.
type CheckoutSessionCreateInvoiceCreationParams struct {
	// Set to `true` to enable invoice creation.
	Enabled *bool `form:"enabled"`
	// Parameters passed when creating invoices for payment-mode Checkout Sessions.
	InvoiceData *CheckoutSessionCreateInvoiceCreationInvoiceDataParams `form:"invoice_data"`
}

// When set, provides configuration for this item's quantity to be adjusted by the customer during Checkout.
type CheckoutSessionCreateLineItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative integer.
	Enabled *bool `form:"enabled"`
	// The maximum quantity the customer can purchase for the Checkout Session. By default this value is 99. You can specify a value up to 999999.
	Maximum *int64 `form:"maximum"`
	// The minimum quantity the customer must purchase for the Checkout Session. By default this value is 0.
	Minimum *int64 `form:"minimum"`
}

// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.
type CheckoutSessionCreateLineItemPriceDataProductDataParams struct {
	// The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description"`
	// A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
	// A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.
	UnitLabel *string `form:"unit_label"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionCreateLineItemPriceDataProductDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The recurring components of a price such as `interval` and `interval_count`.
type CheckoutSessionCreateLineItemPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
type CheckoutSessionCreateLineItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to. One of `product` or `product_data` is required.
	Product *string `form:"product"`
	// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.
	ProductData *CheckoutSessionCreateLineItemPriceDataProductDataParams `form:"product_data"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *CheckoutSessionCreateLineItemPriceDataRecurringParams `form:"recurring"`
	// Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A non-negative integer in cents (or local equivalent) representing how much to charge. One of `unit_amount` or `unit_amount_decimal` is required.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// A list of items the customer is purchasing. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices). The parameter is required for `payment` and `subscription` mode.
//
// For `payment` mode, there is a maximum of 100 line items, however it is recommended to consolidate line items if there are more than a few dozen.
//
// For `subscription` mode, there is a maximum of 20 line items with recurring Prices and 20 line items with one-time Prices. Line items with one-time Prices will be on the initial invoice only.
type CheckoutSessionCreateLineItemParams struct {
	// When set, provides configuration for this item's quantity to be adjusted by the customer during Checkout.
	AdjustableQuantity *CheckoutSessionCreateLineItemAdjustableQuantityParams `form:"adjustable_quantity"`
	// The [tax rates](https://docs.stripe.com/api/tax_rates) that will be applied to this line item depending on the customer's billing/shipping address. We currently support the following countries: US, GB, AU, and all countries in the EU. You can't set this parameter if `ui_mode` is `custom`.
	DynamicTaxRates []*string `form:"dynamic_tax_rates"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object. One of `price` or `price_data` is required.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *CheckoutSessionCreateLineItemPriceDataParams `form:"price_data"`
	// The quantity of the line item being purchased. Quantity should not be defined when `recurring.usage_type=metered`.
	Quantity *int64 `form:"quantity"`
	// The [tax rates](https://docs.stripe.com/api/tax_rates) which apply to this line item.
	TaxRates []*string `form:"tax_rates"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionCreateLineItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Controls settings applied for collecting the customer's business name on the session.
type CheckoutSessionCreateNameCollectionBusinessParams struct {
	// Enable business name collection on the Checkout Session. Defaults to `false`.
	Enabled *bool `form:"enabled"`
	// Whether the customer is required to provide a business name before completing the Checkout Session. Defaults to `false`.
	Optional *bool `form:"optional"`
}

// Controls settings applied for collecting the customer's individual name on the session.
type CheckoutSessionCreateNameCollectionIndividualParams struct {
	// Enable individual name collection on the Checkout Session. Defaults to `false`.
	Enabled *bool `form:"enabled"`
	// Whether the customer is required to provide their name before completing the Checkout Session. Defaults to `false`.
	Optional *bool `form:"optional"`
}

// Controls name collection settings for the session.
//
// You can configure Checkout to collect your customers' business names, individual names, or both. Each name field can be either required or optional.
//
// If a [Customer](https://docs.stripe.com/api/customers) is created or provided, the names can be saved to the Customer object as well.
//
// You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionCreateNameCollectionParams struct {
	// Controls settings applied for collecting the customer's business name on the session.
	Business *CheckoutSessionCreateNameCollectionBusinessParams `form:"business"`
	// Controls settings applied for collecting the customer's individual name on the session.
	Individual *CheckoutSessionCreateNameCollectionIndividualParams `form:"individual"`
}

// When set, provides configuration for the customer to adjust the quantity of the line item created when a customer chooses to add this optional item to their order.
type CheckoutSessionCreateOptionalItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative integer.
	Enabled *bool `form:"enabled"`
	// The maximum quantity of this item the customer can purchase. By default this value is 99. You can specify a value up to 999999.
	Maximum *int64 `form:"maximum"`
	// The minimum quantity of this item the customer must purchase, if they choose to purchase it. Because this item is optional, the customer will always be able to remove it from their order, even if the `minimum` configured here is greater than 0. By default this value is 0.
	Minimum *int64 `form:"minimum"`
}

// A list of optional items the customer can add to their order at checkout. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices).
//
// There is a maximum of 10 optional items allowed on a Checkout Session, and the existing limits on the number of line items allowed on a Checkout Session apply to the combined number of line items and optional items.
//
// For `payment` mode, there is a maximum of 100 combined line items and optional items, however it is recommended to consolidate items if there are more than a few dozen.
//
// For `subscription` mode, there is a maximum of 20 line items and optional items with recurring Prices and 20 line items and optional items with one-time Prices.
//
// You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionCreateOptionalItemParams struct {
	// When set, provides configuration for the customer to adjust the quantity of the line item created when a customer chooses to add this optional item to their order.
	AdjustableQuantity *CheckoutSessionCreateOptionalItemAdjustableQuantityParams `form:"adjustable_quantity"`
	// The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object.
	Price *string `form:"price"`
	// The initial quantity of the line item created when a customer chooses to add this optional item to their order.
	Quantity *int64 `form:"quantity"`
}

// The parameters used to automatically create a Transfer when the payment succeeds.
// For more information, see the PaymentIntents [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts).
type CheckoutSessionCreatePaymentIntentDataTransferDataParams struct {
	// The amount that will be transferred automatically when a charge succeeds.
	Amount *int64 `form:"amount"`
	// If specified, successful charges will be attributed to the destination
	// account for tax reporting, and the funds from charges will be transferred
	// to the destination account. The ID of the resulting transfer will be
	// returned on the successful charge's `transfer` field.
	Destination *string `form:"destination"`
}

// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
type CheckoutSessionCreatePaymentIntentDataParams struct {
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. The amount of the application fee collected will be capped at the total amount captured. For more information, see the PaymentIntents [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts).
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The Stripe account ID for which these funds are intended. For details,
	// see the PaymentIntents [use case for connected
	// accounts](https://docs.stripe.com/docs/payments/connected-accounts).
	OnBehalfOf *string `form:"on_behalf_of"`
	// Email address that the receipt for the resulting payment will be sent to. If `receipt_email` is specified for a payment in live mode, a receipt will be sent regardless of your [email settings](https://dashboard.stripe.com/account/emails).
	ReceiptEmail *string `form:"receipt_email"`
	// Indicates that you intend to [make future payments](https://docs.stripe.com/payments/payment-intents#future-usage) with the payment
	// method collected by this Checkout Session.
	//
	// When setting this to `on_session`, Checkout will show a notice to the
	// customer that their payment details will be saved.
	//
	// When setting this to `off_session`, Checkout will show a notice to the
	// customer that their payment details will be saved and used for future
	// payments.
	//
	// If a Customer has been provided or Checkout creates a new Customer,
	// Checkout will attach the payment method to the Customer.
	//
	// If Checkout does not create a Customer, the payment method is not attached
	// to a Customer. To reuse the payment method, you can retrieve it from the
	// Checkout Session's PaymentIntent.
	//
	// When processing card payments, Checkout also uses `setup_future_usage`
	// to dynamically optimize your payment flow and comply with regional
	// legislation and network rules, such as SCA.
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Shipping information for this payment.
	Shipping *ShippingDetailsParams `form:"shipping"`
	// Text that appears on the customer's statement as the statement descriptor for a non-card charge. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).
	//
	// Setting this value for a card charge returns an error. For card charges, set the [statement_descriptor_suffix](https://docs.stripe.com/get-started/account/statement-descriptors#dynamic) instead.
	StatementDescriptor *string `form:"statement_descriptor"`
	// Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix"`
	// The parameters used to automatically create a Transfer when the payment succeeds.
	// For more information, see the PaymentIntents [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts).
	TransferData *CheckoutSessionCreatePaymentIntentDataTransferDataParams `form:"transfer_data"`
	// A string that identifies the resulting payment as part of a group. See the PaymentIntents [use case for connected accounts](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.
	TransferGroup *string `form:"transfer_group"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionCreatePaymentIntentDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// This parameter allows you to set some attributes on the payment method created during a Checkout session.
type CheckoutSessionCreatePaymentMethodDataParams struct {
	// Allow redisplay will be set on the payment method on confirmation and indicates whether this payment method can be shown again to the customer in a checkout flow. Only set this field if you wish to override the allow_redisplay value determined by Checkout.
	AllowRedisplay *string `form:"allow_redisplay"`
}

// Additional fields for Mandate creation
type CheckoutSessionCreatePaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// A URL for custom mandate text to render during confirmation step.
	// The URL will be rendered with additional GET parameters `payment_intent` and `payment_intent_client_secret` when confirming a Payment Intent,
	// or `setup_intent` and `setup_intent_client_secret` when confirming a Setup Intent.
	CustomMandateURL *string `form:"custom_mandate_url"`
	// List of Stripe products where this mandate can be selected automatically. Only usable in `setup` mode.
	DefaultFor []*string `form:"default_for"`
	// Description of the mandate interval. Only required if 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription *string `form:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule *string `form:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type"`
}

// contains details about the ACSS Debit payment method options. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionCreatePaymentMethodOptionsACSSDebitParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies). This is only accepted for Checkout Sessions in `setup` mode.
	Currency *string `form:"currency"`
	// Additional fields for Mandate creation
	MandateOptions *CheckoutSessionCreatePaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate *string `form:"target_date"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// contains details about the Affirm payment method options.
type CheckoutSessionCreatePaymentMethodOptionsAffirmParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Afterpay Clearpay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsAfterpayClearpayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Alipay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsAlipayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Alma payment method options.
type CheckoutSessionCreatePaymentMethodOptionsAlmaParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
}

// contains details about the AmazonPay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsAmazonPayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the AU Becs Debit payment method options.
type CheckoutSessionCreatePaymentMethodOptionsAUBECSDebitParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate *string `form:"target_date"`
}

// Additional fields for Mandate creation
type CheckoutSessionCreatePaymentMethodOptionsBACSDebitMandateOptionsParams struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'DDIC' or 'STRIPE'.
	ReferencePrefix *string `form:"reference_prefix"`
}

// contains details about the Bacs Debit payment method options.
type CheckoutSessionCreatePaymentMethodOptionsBACSDebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *CheckoutSessionCreatePaymentMethodOptionsBACSDebitMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate *string `form:"target_date"`
}

// contains details about the Bancontact payment method options.
type CheckoutSessionCreatePaymentMethodOptionsBancontactParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Billie payment method options.
type CheckoutSessionCreatePaymentMethodOptionsBillieParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
}

// contains details about the Boleto payment method options.
type CheckoutSessionCreatePaymentMethodOptionsBoletoParams struct {
	// The number of calendar days before a Boleto voucher expires. For example, if you create a Boleto voucher on Monday and you set expires_after_days to 2, the Boleto invoice will expire on Wednesday at 23:59 America/Sao_Paulo time.
	ExpiresAfterDays *int64 `form:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Installment options for card payments
type CheckoutSessionCreatePaymentMethodOptionsCardInstallmentsParams struct {
	// Setting to true enables installments for this Checkout Session.
	// Setting to false will prevent any installment plan from applying to a payment.
	Enabled *bool `form:"enabled"`
}

// Restrictions to apply to the card payment method. For example, you can block specific card brands. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionCreatePaymentMethodOptionsCardRestrictionsParams struct {
	// Specify the card brands to block in the Checkout Session. If a customer enters or selects a card belonging to a blocked brand, they can't complete the Session.
	BrandsBlocked []*string `form:"brands_blocked"`
}

// contains details about the Card payment method options.
type CheckoutSessionCreatePaymentMethodOptionsCardParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Installment options for card payments
	Installments *CheckoutSessionCreatePaymentMethodOptionsCardInstallmentsParams `form:"installments"`
	// Request ability to [capture beyond the standard authorization validity window](https://docs.stripe.com/payments/extended-authorization) for this CheckoutSession.
	RequestExtendedAuthorization *string `form:"request_extended_authorization"`
	// Request ability to [increment the authorization](https://docs.stripe.com/payments/incremental-authorization) for this CheckoutSession.
	RequestIncrementalAuthorization *string `form:"request_incremental_authorization"`
	// Request ability to make [multiple captures](https://docs.stripe.com/payments/multicapture) for this CheckoutSession.
	RequestMulticapture *string `form:"request_multicapture"`
	// Request ability to [overcapture](https://docs.stripe.com/payments/overcapture) for this CheckoutSession.
	RequestOvercapture *string `form:"request_overcapture"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure"`
	// Restrictions to apply to the card payment method. For example, you can block specific card brands. You can't set this parameter if `ui_mode` is `custom`.
	Restrictions *CheckoutSessionCreatePaymentMethodOptionsCardRestrictionsParams `form:"restrictions"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kana prefix (shortened Kana descriptor) or Kana statement descriptor that's set on the account to form the complete statement descriptor. Maximum 22 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 22 characters.
	StatementDescriptorSuffixKana *string `form:"statement_descriptor_suffix_kana"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kanji prefix (shortened Kanji descriptor) or Kanji statement descriptor that's set on the account to form the complete statement descriptor. Maximum 17 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 17 characters.
	StatementDescriptorSuffixKanji *string `form:"statement_descriptor_suffix_kanji"`
}

// contains details about the Cashapp Pay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsCashAppParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Configuration for eu_bank_transfer funding type.
type CheckoutSessionCreatePaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information. Permitted values include: `DE`, `FR`, `IE`, or `NL`.
	Country *string `form:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type CheckoutSessionCreatePaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for eu_bank_transfer funding type.
	EUBankTransfer *CheckoutSessionCreatePaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer"`
	// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
	//
	// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
	RequestedAddressTypes []*string `form:"requested_address_types"`
	// The list of bank transfer types that this PaymentIntent is allowed to use for funding.
	Type *string `form:"type"`
}

// contains details about the Customer Balance payment method options.
type CheckoutSessionCreatePaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *CheckoutSessionCreatePaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType *string `form:"funding_type"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the DemoPay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsDemoPayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the EPS payment method options.
type CheckoutSessionCreatePaymentMethodOptionsEPSParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the FPX payment method options.
type CheckoutSessionCreatePaymentMethodOptionsFPXParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Giropay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsGiropayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Grabpay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsGrabpayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Ideal payment method options.
type CheckoutSessionCreatePaymentMethodOptionsIDEALParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Kakao Pay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsKakaoPayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Describes the upcoming charge for this subscription.
type CheckoutSessionCreatePaymentMethodOptionsKlarnaSubscriptionNextBillingParams struct {
	// The amount of the next charge for the subscription.
	Amount *int64 `form:"amount"`
	// The date of the next charge for the subscription in YYYY-MM-DD format.
	Date *string `form:"date"`
}

// Subscription details if the Checkout Session sets up a future subscription.
type CheckoutSessionCreatePaymentMethodOptionsKlarnaSubscriptionParams struct {
	// Unit of time between subscription charges.
	Interval *string `form:"interval"`
	// The number of intervals (specified in the `interval` attribute) between subscription charges. For example, `interval=month` and `interval_count=3` charges every 3 months.
	IntervalCount *int64 `form:"interval_count"`
	// Name for subscription.
	Name *string `form:"name"`
	// Describes the upcoming charge for this subscription.
	NextBilling *CheckoutSessionCreatePaymentMethodOptionsKlarnaSubscriptionNextBillingParams `form:"next_billing"`
	// A non-customer-facing reference to correlate subscription charges in the Klarna app. Use a value that persists across subscription charges.
	Reference *string `form:"reference"`
}

// contains details about the Klarna payment method options.
type CheckoutSessionCreatePaymentMethodOptionsKlarnaParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Subscription details if the Checkout Session sets up a future subscription.
	Subscriptions []*CheckoutSessionCreatePaymentMethodOptionsKlarnaSubscriptionParams `form:"subscriptions"`
}

// contains details about the Konbini payment method options.
type CheckoutSessionCreatePaymentMethodOptionsKonbiniParams struct {
	// The number of calendar days (between 1 and 60) after which Konbini payment instructions will expire. For example, if a PaymentIntent is confirmed with Konbini and `expires_after_days` set to 2 on Monday JST, the instructions will expire on Wednesday 23:59:59 JST. Defaults to 3 days.
	ExpiresAfterDays *int64 `form:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Korean card payment method options.
type CheckoutSessionCreatePaymentMethodOptionsKrCardParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Link payment method options.
type CheckoutSessionCreatePaymentMethodOptionsLinkParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Mobilepay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsMobilepayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Multibanco payment method options.
type CheckoutSessionCreatePaymentMethodOptionsMultibancoParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Naver Pay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsNaverPayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the OXXO payment method options.
type CheckoutSessionCreatePaymentMethodOptionsOXXOParams struct {
	// The number of calendar days before an OXXO voucher expires. For example, if you create an OXXO voucher on Monday and you set expires_after_days to 2, the OXXO invoice will expire on Wednesday at 23:59 America/Mexico_City time.
	ExpiresAfterDays *int64 `form:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the P24 payment method options.
type CheckoutSessionCreatePaymentMethodOptionsP24Params struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Confirm that the payer has accepted the P24 terms and conditions.
	TOSShownAndAccepted *bool `form:"tos_shown_and_accepted"`
}

// contains details about the Pay By Bank payment method options.
type CheckoutSessionCreatePaymentMethodOptionsPayByBankParams struct{}

// contains details about the PAYCO payment method options.
type CheckoutSessionCreatePaymentMethodOptionsPaycoParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
}

// contains details about the PayNow payment method options.
type CheckoutSessionCreatePaymentMethodOptionsPayNowParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the PayPal payment method options.
type CheckoutSessionCreatePaymentMethodOptionsPaypalParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// [Preferred locale](https://docs.stripe.com/payments/paypal/supported-locales) of the PayPal checkout page that the customer is redirected to.
	PreferredLocale *string `form:"preferred_locale"`
	// A reference of the PayPal transaction visible to customer which is mapped to PayPal's invoice ID. This must be a globally unique ID if you have configured in your PayPal settings to block multiple payments per invoice ID.
	Reference *string `form:"reference"`
	// The risk correlation ID for an on-session payment using a saved PayPal payment method.
	RiskCorrelationID *string `form:"risk_correlation_id"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	//
	// If you've already set `setup_future_usage` and you're performing a request using a publishable key, you can only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Additional fields for Mandate creation
type CheckoutSessionCreatePaymentMethodOptionsPaytoMandateOptionsParams struct {
	// Amount that will be collected. It is required when `amount_type` is `fixed`.
	Amount *int64 `form:"amount"`
	// The type of amount that will be collected. The amount charged must be exact or up to the value of `amount` param for `fixed` or `maximum` type respectively. Defaults to `maximum`.
	AmountType *string `form:"amount_type"`
	// Date, in YYYY-MM-DD format, after which payments will not be collected. Defaults to no end date.
	EndDate *string `form:"end_date"`
	// The periodicity at which payments will be collected. Defaults to `adhoc`.
	PaymentSchedule *string `form:"payment_schedule"`
	// The number of payments that will be made during a payment period. Defaults to 1 except for when `payment_schedule` is `adhoc`. In that case, it defaults to no limit.
	PaymentsPerPeriod *int64 `form:"payments_per_period"`
	// The purpose for which payments are made. Has a default value based on your merchant category code.
	Purpose *string `form:"purpose"`
	// Date, in YYYY-MM-DD format, from which payments will be collected. Defaults to confirmation time.
	StartDate *string `form:"start_date"`
}

// contains details about the PayTo payment method options.
type CheckoutSessionCreatePaymentMethodOptionsPaytoParams struct {
	// Additional fields for Mandate creation
	MandateOptions *CheckoutSessionCreatePaymentMethodOptionsPaytoMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Pix payment method options.
type CheckoutSessionCreatePaymentMethodOptionsPixParams struct {
	// Determines if the amount includes the IOF tax. Defaults to `never`.
	AmountIncludesIof *string `form:"amount_includes_iof"`
	// The number of seconds (between 10 and 1209600) after which Pix payment will expire. Defaults to 86400 seconds.
	ExpiresAfterSeconds *int64 `form:"expires_after_seconds"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the RevolutPay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsRevolutPayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Samsung Pay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsSamsungPayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
}

// contains details about the Satispay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsSatispayParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
}

// Additional fields for Mandate creation
type CheckoutSessionCreatePaymentMethodOptionsSEPADebitMandateOptionsParams struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'STRIPE'.
	ReferencePrefix *string `form:"reference_prefix"`
}

// contains details about the Sepa Debit payment method options.
type CheckoutSessionCreatePaymentMethodOptionsSEPADebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *CheckoutSessionCreatePaymentMethodOptionsSEPADebitMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate *string `form:"target_date"`
}

// contains details about the Sofort payment method options.
type CheckoutSessionCreatePaymentMethodOptionsSofortParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Swish payment method options.
type CheckoutSessionCreatePaymentMethodOptionsSwishParams struct {
	// The order reference that will be displayed to customers in the Swish application. Defaults to the `id` of the Payment Intent.
	Reference *string `form:"reference"`
}

// contains details about the TWINT payment method options.
type CheckoutSessionCreatePaymentMethodOptionsTWINTParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Additional fields for Financial Connections Session creation
type CheckoutSessionCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included. Valid permissions include: `balances`, `ownership`, `payment_method`, and `transactions`.
	Permissions []*string `form:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch"`
}

// contains details about the Us Bank Account payment method options.
type CheckoutSessionCreatePaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation
	FinancialConnections *CheckoutSessionCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate *string `form:"target_date"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// contains details about the WeChat Pay payment method options.
type CheckoutSessionCreatePaymentMethodOptionsWeChatPayParams struct {
	// The app ID registered with WeChat Pay. Only required when client is ios or android.
	AppID *string `form:"app_id"`
	// The client type that the end customer will pay from
	Client *string `form:"client"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Payment-method-specific configuration.
type CheckoutSessionCreatePaymentMethodOptionsParams struct {
	// contains details about the ACSS Debit payment method options. You can't set this parameter if `ui_mode` is `custom`.
	ACSSDebit *CheckoutSessionCreatePaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	// contains details about the Affirm payment method options.
	Affirm *CheckoutSessionCreatePaymentMethodOptionsAffirmParams `form:"affirm"`
	// contains details about the Afterpay Clearpay payment method options.
	AfterpayClearpay *CheckoutSessionCreatePaymentMethodOptionsAfterpayClearpayParams `form:"afterpay_clearpay"`
	// contains details about the Alipay payment method options.
	Alipay *CheckoutSessionCreatePaymentMethodOptionsAlipayParams `form:"alipay"`
	// contains details about the Alma payment method options.
	Alma *CheckoutSessionCreatePaymentMethodOptionsAlmaParams `form:"alma"`
	// contains details about the AmazonPay payment method options.
	AmazonPay *CheckoutSessionCreatePaymentMethodOptionsAmazonPayParams `form:"amazon_pay"`
	// contains details about the AU Becs Debit payment method options.
	AUBECSDebit *CheckoutSessionCreatePaymentMethodOptionsAUBECSDebitParams `form:"au_becs_debit"`
	// contains details about the Bacs Debit payment method options.
	BACSDebit *CheckoutSessionCreatePaymentMethodOptionsBACSDebitParams `form:"bacs_debit"`
	// contains details about the Bancontact payment method options.
	Bancontact *CheckoutSessionCreatePaymentMethodOptionsBancontactParams `form:"bancontact"`
	// contains details about the Billie payment method options.
	Billie *CheckoutSessionCreatePaymentMethodOptionsBillieParams `form:"billie"`
	// contains details about the Boleto payment method options.
	Boleto *CheckoutSessionCreatePaymentMethodOptionsBoletoParams `form:"boleto"`
	// contains details about the Card payment method options.
	Card *CheckoutSessionCreatePaymentMethodOptionsCardParams `form:"card"`
	// contains details about the Cashapp Pay payment method options.
	CashApp *CheckoutSessionCreatePaymentMethodOptionsCashAppParams `form:"cashapp"`
	// contains details about the Customer Balance payment method options.
	CustomerBalance *CheckoutSessionCreatePaymentMethodOptionsCustomerBalanceParams `form:"customer_balance"`
	// contains details about the DemoPay payment method options.
	DemoPay *CheckoutSessionCreatePaymentMethodOptionsDemoPayParams `form:"demo_pay"`
	// contains details about the EPS payment method options.
	EPS *CheckoutSessionCreatePaymentMethodOptionsEPSParams `form:"eps"`
	// contains details about the FPX payment method options.
	FPX *CheckoutSessionCreatePaymentMethodOptionsFPXParams `form:"fpx"`
	// contains details about the Giropay payment method options.
	Giropay *CheckoutSessionCreatePaymentMethodOptionsGiropayParams `form:"giropay"`
	// contains details about the Grabpay payment method options.
	Grabpay *CheckoutSessionCreatePaymentMethodOptionsGrabpayParams `form:"grabpay"`
	// contains details about the Ideal payment method options.
	IDEAL *CheckoutSessionCreatePaymentMethodOptionsIDEALParams `form:"ideal"`
	// contains details about the Kakao Pay payment method options.
	KakaoPay *CheckoutSessionCreatePaymentMethodOptionsKakaoPayParams `form:"kakao_pay"`
	// contains details about the Klarna payment method options.
	Klarna *CheckoutSessionCreatePaymentMethodOptionsKlarnaParams `form:"klarna"`
	// contains details about the Konbini payment method options.
	Konbini *CheckoutSessionCreatePaymentMethodOptionsKonbiniParams `form:"konbini"`
	// contains details about the Korean card payment method options.
	KrCard *CheckoutSessionCreatePaymentMethodOptionsKrCardParams `form:"kr_card"`
	// contains details about the Link payment method options.
	Link *CheckoutSessionCreatePaymentMethodOptionsLinkParams `form:"link"`
	// contains details about the Mobilepay payment method options.
	Mobilepay *CheckoutSessionCreatePaymentMethodOptionsMobilepayParams `form:"mobilepay"`
	// contains details about the Multibanco payment method options.
	Multibanco *CheckoutSessionCreatePaymentMethodOptionsMultibancoParams `form:"multibanco"`
	// contains details about the Naver Pay payment method options.
	NaverPay *CheckoutSessionCreatePaymentMethodOptionsNaverPayParams `form:"naver_pay"`
	// contains details about the OXXO payment method options.
	OXXO *CheckoutSessionCreatePaymentMethodOptionsOXXOParams `form:"oxxo"`
	// contains details about the P24 payment method options.
	P24 *CheckoutSessionCreatePaymentMethodOptionsP24Params `form:"p24"`
	// contains details about the Pay By Bank payment method options.
	PayByBank *CheckoutSessionCreatePaymentMethodOptionsPayByBankParams `form:"pay_by_bank"`
	// contains details about the PAYCO payment method options.
	Payco *CheckoutSessionCreatePaymentMethodOptionsPaycoParams `form:"payco"`
	// contains details about the PayNow payment method options.
	PayNow *CheckoutSessionCreatePaymentMethodOptionsPayNowParams `form:"paynow"`
	// contains details about the PayPal payment method options.
	Paypal *CheckoutSessionCreatePaymentMethodOptionsPaypalParams `form:"paypal"`
	// contains details about the PayTo payment method options.
	Payto *CheckoutSessionCreatePaymentMethodOptionsPaytoParams `form:"payto"`
	// contains details about the Pix payment method options.
	Pix *CheckoutSessionCreatePaymentMethodOptionsPixParams `form:"pix"`
	// contains details about the RevolutPay payment method options.
	RevolutPay *CheckoutSessionCreatePaymentMethodOptionsRevolutPayParams `form:"revolut_pay"`
	// contains details about the Samsung Pay payment method options.
	SamsungPay *CheckoutSessionCreatePaymentMethodOptionsSamsungPayParams `form:"samsung_pay"`
	// contains details about the Satispay payment method options.
	Satispay *CheckoutSessionCreatePaymentMethodOptionsSatispayParams `form:"satispay"`
	// contains details about the Sepa Debit payment method options.
	SEPADebit *CheckoutSessionCreatePaymentMethodOptionsSEPADebitParams `form:"sepa_debit"`
	// contains details about the Sofort payment method options.
	Sofort *CheckoutSessionCreatePaymentMethodOptionsSofortParams `form:"sofort"`
	// contains details about the Swish payment method options.
	Swish *CheckoutSessionCreatePaymentMethodOptionsSwishParams `form:"swish"`
	// contains details about the TWINT payment method options.
	TWINT *CheckoutSessionCreatePaymentMethodOptionsTWINTParams `form:"twint"`
	// contains details about the Us Bank Account payment method options.
	USBankAccount *CheckoutSessionCreatePaymentMethodOptionsUSBankAccountParams `form:"us_bank_account"`
	// contains details about the WeChat Pay payment method options.
	WeChatPay *CheckoutSessionCreatePaymentMethodOptionsWeChatPayParams `form:"wechat_pay"`
}

// This property is used to set up permissions for various actions (e.g., update) on the CheckoutSession object. Can only be set when creating `embedded` or `custom` sessions.
//
// For specific permissions, please refer to their dedicated subsections, such as `permissions.update_shipping_details`.
type CheckoutSessionCreatePermissionsParams struct {
	// Determines which entity is allowed to update the shipping details.
	//
	// Default is `client_only`. Stripe Checkout client will automatically update the shipping details. If set to `server_only`, only your server is allowed to update the shipping details.
	//
	// When set to `server_only`, you must add the onShippingDetailsChange event handler when initializing the Stripe Checkout client and manually update the shipping details from your server using the Stripe API.
	UpdateShippingDetails *string `form:"update_shipping_details"`
}

// Controls phone number collection settings for the session.
//
// We recommend that you review your privacy policy and check with your legal contacts
// before using this feature. Learn more about [collecting phone numbers with Checkout](https://docs.stripe.com/payments/checkout/phone-numbers).
type CheckoutSessionCreatePhoneNumberCollectionParams struct {
	// Set to `true` to enable phone number collection.
	//
	// Can only be set in `payment` and `subscription` mode.
	Enabled *bool `form:"enabled"`
}

// Controls saved payment method settings for the session. Only available in `payment` and `subscription` mode.
type CheckoutSessionCreateSavedPaymentMethodOptionsParams struct {
	// Uses the `allow_redisplay` value of each saved payment method to filter the set presented to a returning customer. By default, only saved payment methods with 'allow_redisplay: ‘always' are shown in Checkout.
	AllowRedisplayFilters []*string `form:"allow_redisplay_filters"`
	// Enable customers to choose if they wish to remove their saved payment methods. Disabled by default.
	PaymentMethodRemove *string `form:"payment_method_remove"`
	// Enable customers to choose if they wish to save their payment method for future use. Disabled by default.
	PaymentMethodSave *string `form:"payment_method_save"`
}

// A subset of parameters to be passed to SetupIntent creation for Checkout Sessions in `setup` mode.
type CheckoutSessionCreateSetupIntentDataParams struct {
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The Stripe account for which the setup is intended.
	OnBehalfOf *string `form:"on_behalf_of"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionCreateSetupIntentDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// When set, provides configuration for Checkout to collect a shipping address from a customer.
type CheckoutSessionCreateShippingAddressCollectionParams struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for
	// shipping locations.
	AllowedCountries []*string `form:"allowed_countries"`
}

// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
type CheckoutSessionCreateShippingOptionShippingRateDataDeliveryEstimateMaximumParams struct {
	// A unit of time.
	Unit *string `form:"unit"`
	// Must be greater than 0.
	Value *int64 `form:"value"`
}

// The lower bound of the estimated range. If empty, represents no lower bound.
type CheckoutSessionCreateShippingOptionShippingRateDataDeliveryEstimateMinimumParams struct {
	// A unit of time.
	Unit *string `form:"unit"`
	// Must be greater than 0.
	Value *int64 `form:"value"`
}

// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
type CheckoutSessionCreateShippingOptionShippingRateDataDeliveryEstimateParams struct {
	// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
	Maximum *CheckoutSessionCreateShippingOptionShippingRateDataDeliveryEstimateMaximumParams `form:"maximum"`
	// The lower bound of the estimated range. If empty, represents no lower bound.
	Minimum *CheckoutSessionCreateShippingOptionShippingRateDataDeliveryEstimateMinimumParams `form:"minimum"`
}

// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type CheckoutSessionCreateShippingOptionShippingRateDataFixedAmountCurrencyOptionsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
type CheckoutSessionCreateShippingOptionShippingRateDataFixedAmountParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions map[string]*CheckoutSessionCreateShippingOptionShippingRateDataFixedAmountCurrencyOptionsParams `form:"currency_options"`
}

// Parameters to be passed to Shipping Rate creation for this shipping option.
type CheckoutSessionCreateShippingOptionShippingRateDataParams struct {
	// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DeliveryEstimate *CheckoutSessionCreateShippingOptionShippingRateDataDeliveryEstimateParams `form:"delivery_estimate"`
	// The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DisplayName *string `form:"display_name"`
	// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
	FixedAmount *CheckoutSessionCreateShippingOptionShippingRateDataFixedAmountParams `form:"fixed_amount"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.
	TaxCode *string `form:"tax_code"`
	// The type of calculation to use on the shipping rate.
	Type *string `form:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionCreateShippingOptionShippingRateDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The shipping rate options to apply to this Session. Up to a maximum of 5.
type CheckoutSessionCreateShippingOptionParams struct {
	// The ID of the Shipping Rate to use for this shipping option.
	ShippingRate *string `form:"shipping_rate"`
	// Parameters to be passed to Shipping Rate creation for this shipping option.
	ShippingRateData *CheckoutSessionCreateShippingOptionShippingRateDataParams `form:"shipping_rate_data"`
}

// Configure behavior for flexible billing mode.
type CheckoutSessionCreateSubscriptionDataBillingModeFlexibleParams struct {
	// Controls how invoices and invoice items display proration amounts and discount amounts.
	ProrationDiscounts *string `form:"proration_discounts"`
}

// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
type CheckoutSessionCreateSubscriptionDataBillingModeParams struct {
	// Configure behavior for flexible billing mode.
	Flexible *CheckoutSessionCreateSubscriptionDataBillingModeFlexibleParams `form:"flexible"`
	// Controls the calculation and orchestration of prorations and invoices for subscriptions. If no value is passed, the default is `flexible`.
	Type *string `form:"type"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type CheckoutSessionCreateSubscriptionDataInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type CheckoutSessionCreateSubscriptionDataInvoiceSettingsParams struct {
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *CheckoutSessionCreateSubscriptionDataInvoiceSettingsIssuerParams `form:"issuer"`
}

// If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.
type CheckoutSessionCreateSubscriptionDataTransferDataParams struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount is transferred to the destination.
	AmountPercent *float64 `form:"amount_percent"`
	// ID of an existing, connected Stripe account.
	Destination *string `form:"destination"`
}

// Defines how the subscription should behave when the user's free trial ends.
type CheckoutSessionCreateSubscriptionDataTrialSettingsEndBehaviorParams struct {
	// Indicates how the subscription should change when the trial ends if the user did not provide a payment method.
	MissingPaymentMethod *string `form:"missing_payment_method"`
}

// Settings related to subscription trials.
type CheckoutSessionCreateSubscriptionDataTrialSettingsParams struct {
	// Defines how the subscription should behave when the user's free trial ends.
	EndBehavior *CheckoutSessionCreateSubscriptionDataTrialSettingsEndBehaviorParams `form:"end_behavior"`
}

// A subset of parameters to be passed to subscription creation for Checkout Sessions in `subscription` mode.
type CheckoutSessionCreateSubscriptionDataParams struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. To use an application fee percent, the request must be made on behalf of another account, using the `Stripe-Account` header or an OAuth key. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// A future timestamp to anchor the subscription's billing cycle for new subscriptions. You can't set this parameter if `ui_mode` is `custom`.
	BillingCycleAnchor *int64 `form:"billing_cycle_anchor"`
	// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
	BillingMode *CheckoutSessionCreateSubscriptionDataBillingModeParams `form:"billing_mode"`
	// The tax rates that will apply to any subscription item that does not have
	// `tax_rates` set. Invoices created will have their `default_tax_rates` populated
	// from the subscription.
	DefaultTaxRates []*string `form:"default_tax_rates"`
	// The subscription's description, meant to be displayable to the customer.
	// Use this field to optionally store an explanation of the subscription
	// for rendering in the [customer portal](https://docs.stripe.com/customer-management).
	Description *string `form:"description"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *CheckoutSessionCreateSubscriptionDataInvoiceSettingsParams `form:"invoice_settings"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The account on behalf of which to charge, for each of the subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of"`
	// Determines how to handle prorations resulting from the `billing_cycle_anchor`. If no value is passed, the default is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior"`
	// If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.
	TransferData *CheckoutSessionCreateSubscriptionDataTransferDataParams `form:"transfer_data"`
	// Unix timestamp representing the end of the trial period the customer will get before being charged for the first time. Has to be at least 48 hours in the future.
	TrialEnd *int64 `form:"trial_end"`
	// Integer representing the number of trial period days before the customer is charged for the first time. Has to be at least 1.
	TrialPeriodDays *int64 `form:"trial_period_days"`
	// Settings related to subscription trials.
	TrialSettings *CheckoutSessionCreateSubscriptionDataTrialSettingsParams `form:"trial_settings"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionCreateSubscriptionDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Controls tax ID collection during checkout.
type CheckoutSessionCreateTaxIDCollectionParams struct {
	// Enable tax ID collection during checkout. Defaults to `false`.
	Enabled *bool `form:"enabled"`
	// Describes whether a tax ID is required during checkout. Defaults to `never`. You can't set this parameter if `ui_mode` is `custom`.
	Required *string `form:"required"`
}

// contains details about the Link wallet options.
type CheckoutSessionCreateWalletOptionsLinkParams struct {
	// Specifies whether Checkout should display Link as a payment option. By default, Checkout will display all the supported wallets that the Checkout Session was created with. This is the `auto` behavior, and it is the default choice.
	Display *string `form:"display"`
}

// Wallet-specific configuration.
type CheckoutSessionCreateWalletOptionsParams struct {
	// contains details about the Link wallet options.
	Link *CheckoutSessionCreateWalletOptionsLinkParams `form:"link"`
}

// Creates a Checkout Session object.
type CheckoutSessionCreateParams struct {
	Params `form:"*"`
	// Settings for price localization with [Adaptive Pricing](https://docs.stripe.com/payments/checkout/adaptive-pricing).
	AdaptivePricing *CheckoutSessionCreateAdaptivePricingParams `form:"adaptive_pricing"`
	// Configure actions after a Checkout Session has expired. You can't set this parameter if `ui_mode` is `custom`.
	AfterExpiration *CheckoutSessionCreateAfterExpirationParams `form:"after_expiration"`
	// Enables user redeemable promotion codes.
	AllowPromotionCodes *bool `form:"allow_promotion_codes"`
	// Settings for automatic tax lookup for this session and resulting payments, invoices, and subscriptions.
	AutomaticTax *CheckoutSessionCreateAutomaticTaxParams `form:"automatic_tax"`
	// Specify whether Checkout should collect the customer's billing address. Defaults to `auto`.
	BillingAddressCollection *string `form:"billing_address_collection"`
	// The branding settings for the Checkout Session. This parameter is not allowed if ui_mode is `custom`.
	BrandingSettings *CheckoutSessionCreateBrandingSettingsParams `form:"branding_settings"`
	// If set, Checkout displays a back button and customers will be directed to this URL if they decide to cancel payment and return to your website. This parameter is not allowed if ui_mode is `embedded` or `custom`.
	CancelURL *string `form:"cancel_url"`
	// A unique string to reference the Checkout Session. This can be a
	// customer ID, a cart ID, or similar, and can be used to reconcile the
	// session with your internal systems.
	ClientReferenceID *string `form:"client_reference_id"`
	// Configure fields for the Checkout Session to gather active consent from customers.
	ConsentCollection *CheckoutSessionCreateConsentCollectionParams `form:"consent_collection"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies). Required in `setup` mode when `payment_method_types` is not set.
	Currency *string `form:"currency"`
	// ID of an existing Customer, if one exists. In `payment` mode, the customer's most recently saved card
	// payment method will be used to prefill the email, name, card details, and billing address
	// on the Checkout page. In `subscription` mode, the customer's [default payment method](https://docs.stripe.com/api/customers/update#update_customer-invoice_settings-default_payment_method)
	// will be used if it's a card, otherwise the most recently saved card will be used. A valid billing address, billing name and billing email are required on the payment method for Checkout to prefill the customer's card details.
	//
	// If the Customer already has a valid [email](https://docs.stripe.com/api/customers/object#customer_object-email) set, the email will be prefilled and not editable in Checkout.
	// If the Customer does not have a valid `email`, Checkout will set the email entered during the session on the Customer.
	//
	// If blank for Checkout Sessions in `subscription` mode or with `customer_creation` set as `always` in `payment` mode, Checkout will create a new Customer object based on information provided during the payment flow.
	//
	// You can set [`payment_intent_data.setup_future_usage`](https://docs.stripe.com/api/checkout/sessions/create#create_checkout_session-payment_intent_data-setup_future_usage) to have Checkout automatically attach the payment method to the Customer you pass in for future reuse.
	Customer *string `form:"customer"`
	// ID of an existing Account, if one exists. Has the same behavior as `customer`.
	CustomerAccount *string `form:"customer_account"`
	// Configure whether a Checkout Session creates a [Customer](https://docs.stripe.com/api/customers) during Session confirmation.
	//
	// When a Customer is not created, you can still retrieve email, address, and other customer data entered in Checkout
	// with [customer_details](https://docs.stripe.com/api/checkout/sessions/object#checkout_session_object-customer_details).
	//
	// Sessions that don't create Customers instead are grouped by [guest customers](https://docs.stripe.com/payments/checkout/guest-customers)
	// in the Dashboard. Promotion codes limited to first time customers will return invalid for these Sessions.
	//
	// Can only be set in `payment` and `setup` mode.
	CustomerCreation *string `form:"customer_creation"`
	// If provided, this value will be used when the Customer object is created.
	// If not provided, customers will be asked to enter their email address.
	// Use this parameter to prefill customer data if you already have an email
	// on file. To access information about the customer once a session is
	// complete, use the `customer` field.
	CustomerEmail *string `form:"customer_email"`
	// Controls what fields on Customer can be updated by the Checkout Session. Can only be provided when `customer` is provided.
	CustomerUpdate *CheckoutSessionCreateCustomerUpdateParams `form:"customer_update"`
	// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
	CustomFields []*CheckoutSessionCreateCustomFieldParams `form:"custom_fields"`
	// Display additional text for your customers using custom text. You can't set this parameter if `ui_mode` is `custom`.
	CustomText *CheckoutSessionCreateCustomTextParams `form:"custom_text"`
	// The coupon or promotion code to apply to this Session. Currently, only up to one may be specified.
	Discounts []*CheckoutSessionCreateDiscountParams `form:"discounts"`
	// A list of the types of payment methods (e.g., `card`) that should be excluded from this Checkout Session. This should only be used when payment methods for this Checkout Session are managed through the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).
	ExcludedPaymentMethodTypes []*string `form:"excluded_payment_method_types"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The Epoch time in seconds at which the Checkout Session will expire. It can be anywhere from 30 minutes to 24 hours after Checkout Session creation. By default, this value is 24 hours from creation.
	ExpiresAt *int64 `form:"expires_at"`
	// Generate a post-purchase Invoice for one-time payments.
	InvoiceCreation *CheckoutSessionCreateInvoiceCreationParams `form:"invoice_creation"`
	// A list of items the customer is purchasing. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices). The parameter is required for `payment` and `subscription` mode.
	//
	// For `payment` mode, there is a maximum of 100 line items, however it is recommended to consolidate line items if there are more than a few dozen.
	//
	// For `subscription` mode, there is a maximum of 20 line items with recurring Prices and 20 line items with one-time Prices. Line items with one-time Prices will be on the initial invoice only.
	LineItems []*CheckoutSessionCreateLineItemParams `form:"line_items"`
	// The IETF language tag of the locale Checkout is displayed in. If blank or `auto`, the browser's locale is used.
	Locale *string `form:"locale"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The mode of the Checkout Session. Pass `subscription` if the Checkout Session includes at least one recurring item.
	Mode *string `form:"mode"`
	// Controls name collection settings for the session.
	//
	// You can configure Checkout to collect your customers' business names, individual names, or both. Each name field can be either required or optional.
	//
	// If a [Customer](https://docs.stripe.com/api/customers) is created or provided, the names can be saved to the Customer object as well.
	//
	// You can't set this parameter if `ui_mode` is `custom`.
	NameCollection *CheckoutSessionCreateNameCollectionParams `form:"name_collection"`
	// A list of optional items the customer can add to their order at checkout. Use this parameter to pass one-time or recurring [Prices](https://docs.stripe.com/api/prices).
	//
	// There is a maximum of 10 optional items allowed on a Checkout Session, and the existing limits on the number of line items allowed on a Checkout Session apply to the combined number of line items and optional items.
	//
	// For `payment` mode, there is a maximum of 100 combined line items and optional items, however it is recommended to consolidate items if there are more than a few dozen.
	//
	// For `subscription` mode, there is a maximum of 20 line items and optional items with recurring Prices and 20 line items and optional items with one-time Prices.
	//
	// You can't set this parameter if `ui_mode` is `custom`.
	OptionalItems []*CheckoutSessionCreateOptionalItemParams `form:"optional_items"`
	// Where the user is coming from. This informs the optimizations that are applied to the session. You can't set this parameter if `ui_mode` is `custom`.
	OriginContext *string `form:"origin_context"`
	// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
	PaymentIntentData *CheckoutSessionCreatePaymentIntentDataParams `form:"payment_intent_data"`
	// Specify whether Checkout should collect a payment method. When set to `if_required`, Checkout will not collect a payment method when the total due for the session is 0.
	// This may occur if the Checkout Session includes a free trial or a discount.
	//
	// Can only be set in `subscription` mode. Defaults to `always`.
	//
	// If you'd like information on how to collect a payment method outside of Checkout, read the guide on configuring [subscriptions with a free trial](https://docs.stripe.com/payments/checkout/free-trials).
	PaymentMethodCollection *string `form:"payment_method_collection"`
	// The ID of the payment method configuration to use with this Checkout session.
	PaymentMethodConfiguration *string `form:"payment_method_configuration"`
	// This parameter allows you to set some attributes on the payment method created during a Checkout session.
	PaymentMethodData *CheckoutSessionCreatePaymentMethodDataParams `form:"payment_method_data"`
	// Payment-method-specific configuration.
	PaymentMethodOptions *CheckoutSessionCreatePaymentMethodOptionsParams `form:"payment_method_options"`
	// A list of the types of payment methods (e.g., `card`) this Checkout Session can accept.
	//
	// You can omit this attribute to manage your payment methods from the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).
	// See [Dynamic Payment Methods](https://docs.stripe.com/payments/payment-methods/integration-options#using-dynamic-payment-methods) for more details.
	//
	// Read more about the supported payment methods and their requirements in our [payment
	// method details guide](https://docs.stripe.com/docs/payments/checkout/payment-methods).
	//
	// If multiple payment methods are passed, Checkout will dynamically reorder them to
	// prioritize the most relevant payment methods based on the customer's location and
	// other characteristics.
	PaymentMethodTypes []*string `form:"payment_method_types"`
	// This property is used to set up permissions for various actions (e.g., update) on the CheckoutSession object. Can only be set when creating `embedded` or `custom` sessions.
	//
	// For specific permissions, please refer to their dedicated subsections, such as `permissions.update_shipping_details`.
	Permissions *CheckoutSessionCreatePermissionsParams `form:"permissions"`
	// Controls phone number collection settings for the session.
	//
	// We recommend that you review your privacy policy and check with your legal contacts
	// before using this feature. Learn more about [collecting phone numbers with Checkout](https://docs.stripe.com/payments/checkout/phone-numbers).
	PhoneNumberCollection *CheckoutSessionCreatePhoneNumberCollectionParams `form:"phone_number_collection"`
	// This parameter applies to `ui_mode: embedded`. Learn more about the [redirect behavior](https://docs.stripe.com/payments/checkout/custom-success-page?payment-ui=embedded-form) of embedded sessions. Defaults to `always`.
	RedirectOnCompletion *string `form:"redirect_on_completion"`
	// The URL to redirect your customer back to after they authenticate or cancel their payment on the
	// payment method's app or site. This parameter is required if `ui_mode` is `embedded` or `custom`
	// and redirect-based payment methods are enabled on the session.
	ReturnURL *string `form:"return_url"`
	// Controls saved payment method settings for the session. Only available in `payment` and `subscription` mode.
	SavedPaymentMethodOptions *CheckoutSessionCreateSavedPaymentMethodOptionsParams `form:"saved_payment_method_options"`
	// A subset of parameters to be passed to SetupIntent creation for Checkout Sessions in `setup` mode.
	SetupIntentData *CheckoutSessionCreateSetupIntentDataParams `form:"setup_intent_data"`
	// When set, provides configuration for Checkout to collect a shipping address from a customer.
	ShippingAddressCollection *CheckoutSessionCreateShippingAddressCollectionParams `form:"shipping_address_collection"`
	// The shipping rate options to apply to this Session. Up to a maximum of 5.
	ShippingOptions []*CheckoutSessionCreateShippingOptionParams `form:"shipping_options"`
	// Describes the type of transaction being performed by Checkout in order
	// to customize relevant text on the page, such as the submit button.
	//  `submit_type` can only be specified on Checkout Sessions in
	// `payment` or `subscription` mode. If blank or `auto`, `pay` is used.
	// You can't set this parameter if `ui_mode` is `custom`.
	SubmitType *string `form:"submit_type"`
	// A subset of parameters to be passed to subscription creation for Checkout Sessions in `subscription` mode.
	SubscriptionData *CheckoutSessionCreateSubscriptionDataParams `form:"subscription_data"`
	// The URL to which Stripe should send customers when payment or setup
	// is complete.
	// This parameter is not allowed if ui_mode is `embedded` or `custom`. If you'd like to use
	// information from the successful Checkout Session on your page, read the
	// guide on [customizing your success page](https://docs.stripe.com/payments/checkout/custom-success-page).
	SuccessURL *string `form:"success_url"`
	// Controls tax ID collection during checkout.
	TaxIDCollection *CheckoutSessionCreateTaxIDCollectionParams `form:"tax_id_collection"`
	// The UI mode of the Session. Defaults to `hosted`.
	UIMode *string `form:"ui_mode"`
	// Wallet-specific configuration.
	WalletOptions *CheckoutSessionCreateWalletOptionsParams `form:"wallet_options"`
}

// AddExpand appends a new field to expand.
func (p *CheckoutSessionCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves a Checkout Session object.
type CheckoutSessionRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CheckoutSessionRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The shipping details to apply to this Session.
type CheckoutSessionUpdateCollectedInformationShippingDetailsParams struct {
	// The address of the customer
	Address *AddressParams `form:"address"`
	// The name of customer
	Name *string `form:"name"`
}

// Information about the customer collected within the Checkout Session. Can only be set when updating `embedded` or `custom` sessions.
type CheckoutSessionUpdateCollectedInformationParams struct {
	// The shipping details to apply to this Session.
	ShippingDetails *CheckoutSessionUpdateCollectedInformationShippingDetailsParams `form:"shipping_details"`
}

// When set, provides configuration for this item's quantity to be adjusted by the customer during Checkout.
type CheckoutSessionUpdateLineItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any positive integer. Setting to false will remove any previously specified constraints on quantity.
	Enabled *bool `form:"enabled"`
	// The maximum quantity the customer can purchase for the Checkout Session. By default this value is 99. You can specify a value up to 999999.
	Maximum *int64 `form:"maximum"`
	// The minimum quantity the customer must purchase for the Checkout Session. By default this value is 0.
	Minimum *int64 `form:"minimum"`
}

// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.
type CheckoutSessionUpdateLineItemPriceDataProductDataParams struct {
	// The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description"`
	// A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
	// A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.
	UnitLabel *string `form:"unit_label"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionUpdateLineItemPriceDataProductDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The recurring components of a price such as `interval` and `interval_count`.
type CheckoutSessionUpdateLineItemPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required when creating a new line item.
type CheckoutSessionUpdateLineItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to. One of `product` or `product_data` is required.
	Product *string `form:"product"`
	// Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.
	ProductData *CheckoutSessionUpdateLineItemPriceDataProductDataParams `form:"product_data"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *CheckoutSessionUpdateLineItemPriceDataRecurringParams `form:"recurring"`
	// Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A non-negative integer in cents (or local equivalent) representing how much to charge. One of `unit_amount` or `unit_amount_decimal` is required.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// A list of items the customer is purchasing.
//
// When updating line items, you must retransmit the entire array of line items.
//
// To retain an existing line item, specify its `id`.
//
// To update an existing line item, specify its `id` along with the new values of the fields to update.
//
// To add a new line item, specify one of `price` or `price_data` and `quantity`.
//
// To remove an existing line item, omit the line item's ID from the retransmitted array.
//
// To reorder a line item, specify it at the desired position in the retransmitted array.
type CheckoutSessionUpdateLineItemParams struct {
	// When set, provides configuration for this item's quantity to be adjusted by the customer during Checkout.
	AdjustableQuantity *CheckoutSessionUpdateLineItemAdjustableQuantityParams `form:"adjustable_quantity"`
	// ID of an existing line item.
	ID *string `form:"id"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The ID of the [Price](https://docs.stripe.com/api/prices). One of `price` or `price_data` is required when creating a new line item.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required when creating a new line item.
	PriceData *CheckoutSessionUpdateLineItemPriceDataParams `form:"price_data"`
	// The quantity of the line item being purchased. Quantity should not be defined when `recurring.usage_type=metered`.
	Quantity *int64 `form:"quantity"`
	// The [tax rates](https://docs.stripe.com/api/tax_rates) which apply to this line item.
	TaxRates []*string `form:"tax_rates"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionUpdateLineItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
type CheckoutSessionUpdateShippingOptionShippingRateDataDeliveryEstimateMaximumParams struct {
	// A unit of time.
	Unit *string `form:"unit"`
	// Must be greater than 0.
	Value *int64 `form:"value"`
}

// The lower bound of the estimated range. If empty, represents no lower bound.
type CheckoutSessionUpdateShippingOptionShippingRateDataDeliveryEstimateMinimumParams struct {
	// A unit of time.
	Unit *string `form:"unit"`
	// Must be greater than 0.
	Value *int64 `form:"value"`
}

// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
type CheckoutSessionUpdateShippingOptionShippingRateDataDeliveryEstimateParams struct {
	// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
	Maximum *CheckoutSessionUpdateShippingOptionShippingRateDataDeliveryEstimateMaximumParams `form:"maximum"`
	// The lower bound of the estimated range. If empty, represents no lower bound.
	Minimum *CheckoutSessionUpdateShippingOptionShippingRateDataDeliveryEstimateMinimumParams `form:"minimum"`
}

// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type CheckoutSessionUpdateShippingOptionShippingRateDataFixedAmountCurrencyOptionsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
type CheckoutSessionUpdateShippingOptionShippingRateDataFixedAmountParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions map[string]*CheckoutSessionUpdateShippingOptionShippingRateDataFixedAmountCurrencyOptionsParams `form:"currency_options"`
}

// Parameters to be passed to Shipping Rate creation for this shipping option.
type CheckoutSessionUpdateShippingOptionShippingRateDataParams struct {
	// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DeliveryEstimate *CheckoutSessionUpdateShippingOptionShippingRateDataDeliveryEstimateParams `form:"delivery_estimate"`
	// The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DisplayName *string `form:"display_name"`
	// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
	FixedAmount *CheckoutSessionUpdateShippingOptionShippingRateDataFixedAmountParams `form:"fixed_amount"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://docs.stripe.com/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.
	TaxCode *string `form:"tax_code"`
	// The type of calculation to use on the shipping rate.
	Type *string `form:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionUpdateShippingOptionShippingRateDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The shipping rate options to apply to this Session. Up to a maximum of 5.
type CheckoutSessionUpdateShippingOptionParams struct {
	// The ID of the Shipping Rate to use for this shipping option.
	ShippingRate *string `form:"shipping_rate"`
	// Parameters to be passed to Shipping Rate creation for this shipping option.
	ShippingRateData *CheckoutSessionUpdateShippingOptionShippingRateDataParams `form:"shipping_rate_data"`
}

// Updates a Checkout Session object.
//
// Related guide: [Dynamically update a Checkout Session](https://docs.stripe.com/payments/advanced/dynamic-updates)
type CheckoutSessionUpdateParams struct {
	Params `form:"*"`
	// Information about the customer collected within the Checkout Session. Can only be set when updating `embedded` or `custom` sessions.
	CollectedInformation *CheckoutSessionUpdateCollectedInformationParams `form:"collected_information"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A list of items the customer is purchasing.
	//
	// When updating line items, you must retransmit the entire array of line items.
	//
	// To retain an existing line item, specify its `id`.
	//
	// To update an existing line item, specify its `id` along with the new values of the fields to update.
	//
	// To add a new line item, specify one of `price` or `price_data` and `quantity`.
	//
	// To remove an existing line item, omit the line item's ID from the retransmitted array.
	//
	// To reorder a line item, specify it at the desired position in the retransmitted array.
	LineItems []*CheckoutSessionUpdateLineItemParams `form:"line_items"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The shipping rate options to apply to this Session. Up to a maximum of 5.
	ShippingOptions []*CheckoutSessionUpdateShippingOptionParams `form:"shipping_options"`
}

// AddExpand appends a new field to expand.
func (p *CheckoutSessionUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CheckoutSessionUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Settings for price localization with [Adaptive Pricing](https://docs.stripe.com/payments/checkout/adaptive-pricing).
type CheckoutSessionAdaptivePricing struct {
	// If enabled, Adaptive Pricing is available on [eligible sessions](https://docs.stripe.com/payments/currencies/localize-prices/adaptive-pricing?payment-ui=stripe-hosted#restrictions).
	Enabled bool `json:"enabled"`
}

// When set, configuration used to recover the Checkout Session on expiry.
type CheckoutSessionAfterExpirationRecovery struct {
	// Enables user redeemable promotion codes on the recovered Checkout Sessions. Defaults to `false`
	AllowPromotionCodes bool `json:"allow_promotion_codes"`
	// If `true`, a recovery url will be generated to recover this Checkout Session if it
	// expires before a transaction is completed. It will be attached to the
	// Checkout Session object upon expiration.
	Enabled bool `json:"enabled"`
	// The timestamp at which the recovery URL will expire.
	ExpiresAt int64 `json:"expires_at"`
	// URL that creates a new Checkout Session when clicked that is a copy of this expired Checkout Session
	URL string `json:"url"`
}

// When set, provides configuration for actions to take if this Checkout Session expires.
type CheckoutSessionAfterExpiration struct {
	// When set, configuration used to recover the Checkout Session on expiry.
	Recovery *CheckoutSessionAfterExpirationRecovery `json:"recovery"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type CheckoutSessionAutomaticTaxLiability struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type CheckoutSessionAutomaticTaxLiabilityType `json:"type"`
}
type CheckoutSessionAutomaticTax struct {
	// Indicates whether automatic tax is enabled for the session
	Enabled bool `json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *CheckoutSessionAutomaticTaxLiability `json:"liability"`
	// The tax provider powering automatic tax.
	Provider string `json:"provider"`
	// The status of the most recent automated tax calculation for this session.
	Status CheckoutSessionAutomaticTaxStatus `json:"status"`
}

// The icon for the Checkout Session. You cannot set both `logo` and `icon`.
type CheckoutSessionBrandingSettingsIcon struct {
	// The ID of a [File upload](https://stripe.com/docs/api/files) representing the icon. Purpose must be `business_icon`. Required if `type` is `file` and disallowed otherwise.
	File string `json:"file"`
	// The type of image for the icon. Must be one of `file` or `url`.
	Type CheckoutSessionBrandingSettingsIconType `json:"type"`
	// The URL of the image. Present when `type` is `url`.
	URL string `json:"url"`
}

// The logo for the Checkout Session. You cannot set both `logo` and `icon`.
type CheckoutSessionBrandingSettingsLogo struct {
	// The ID of a [File upload](https://stripe.com/docs/api/files) representing the logo. Purpose must be `business_logo`. Required if `type` is `file` and disallowed otherwise.
	File string `json:"file"`
	// The type of image for the logo. Must be one of `file` or `url`.
	Type CheckoutSessionBrandingSettingsLogoType `json:"type"`
	// The URL of the image. Present when `type` is `url`.
	URL string `json:"url"`
}
type CheckoutSessionBrandingSettings struct {
	// A hex color value starting with `#` representing the background color for the Checkout Session.
	BackgroundColor string `json:"background_color"`
	// The border style for the Checkout Session. Must be one of `rounded`, `rectangular`, or `pill`.
	BorderStyle CheckoutSessionBrandingSettingsBorderStyle `json:"border_style"`
	// A hex color value starting with `#` representing the button color for the Checkout Session.
	ButtonColor string `json:"button_color"`
	// The display name shown on the Checkout Session.
	DisplayName string `json:"display_name"`
	// The font family for the Checkout Session. Must be one of the [supported font families](https://docs.stripe.com/payments/checkout/customization/appearance?payment-ui=stripe-hosted#font-compatibility).
	FontFamily string `json:"font_family"`
	// The icon for the Checkout Session. You cannot set both `logo` and `icon`.
	Icon *CheckoutSessionBrandingSettingsIcon `json:"icon"`
	// The logo for the Checkout Session. You cannot set both `logo` and `icon`.
	Logo *CheckoutSessionBrandingSettingsLogo `json:"logo"`
}

// Shipping information for this Checkout Session.
type CheckoutSessionCollectedInformationShippingDetails struct {
	Address *Address `json:"address"`
	// Customer name.
	Name string `json:"name"`
}

// Information about the customer collected within the Checkout Session.
type CheckoutSessionCollectedInformation struct {
	// Customer's business name for this Checkout Session
	BusinessName string `json:"business_name"`
	// Customer's individual name for this Checkout Session
	IndividualName string `json:"individual_name"`
	// Shipping information for this Checkout Session.
	ShippingDetails *CheckoutSessionCollectedInformationShippingDetails `json:"shipping_details"`
}

// Results of `consent_collection` for this session.
type CheckoutSessionConsent struct {
	// If `opt_in`, the customer consents to receiving promotional communications
	// from the merchant about this Checkout Session.
	Promotions CheckoutSessionConsentPromotions `json:"promotions"`
	// If `accepted`, the customer in this Checkout Session has agreed to the merchant's terms of service.
	TermsOfService CheckoutSessionConsentTermsOfService `json:"terms_of_service"`
}

// If set to `hidden`, it will hide legal text related to the reuse of a payment method.
type CheckoutSessionConsentCollectionPaymentMethodReuseAgreement struct {
	// Determines the position and visibility of the payment method reuse agreement in the UI. When set to `auto`, Stripe's defaults will be used.
	//
	// When set to `hidden`, the payment method reuse agreement text will always be hidden in the UI.
	Position CheckoutSessionConsentCollectionPaymentMethodReuseAgreementPosition `json:"position"`
}

// When set, provides configuration for the Checkout Session to gather active consent from customers.
type CheckoutSessionConsentCollection struct {
	// If set to `hidden`, it will hide legal text related to the reuse of a payment method.
	PaymentMethodReuseAgreement *CheckoutSessionConsentCollectionPaymentMethodReuseAgreement `json:"payment_method_reuse_agreement"`
	// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
	// Session will determine whether to display an option to opt into promotional communication
	// from the merchant depending on the customer's locale. Only available to US merchants.
	Promotions CheckoutSessionConsentCollectionPromotions `json:"promotions"`
	// If set to `required`, it requires customers to accept the terms of service before being able to pay.
	TermsOfService CheckoutSessionConsentCollectionTermsOfService `json:"terms_of_service"`
}

// Currency conversion details for [Adaptive Pricing](https://docs.stripe.com/payments/checkout/adaptive-pricing) sessions created before 2025-03-31.
type CheckoutSessionCurrencyConversion struct {
	// Total of all items in source currency before discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total of all items in source currency after discounts and taxes are applied.
	AmountTotal int64 `json:"amount_total"`
	// Exchange rate used to convert source currency amounts to customer currency amounts
	FxRate float64 `json:"fx_rate,string"`
	// Creation currency of the CheckoutSession before localization
	SourceCurrency Currency `json:"source_currency"`
}

// The options available for the customer to select. Up to 200 options allowed.
type CheckoutSessionCustomFieldDropdownOption struct {
	// The label for the option, displayed to the customer. Up to 100 characters.
	Label string `json:"label"`
	// The value for this option, not displayed to the customer, used by your integration to reconcile the option selected by the customer. Must be unique to this option, alphanumeric, and up to 100 characters.
	Value string `json:"value"`
}
type CheckoutSessionCustomFieldDropdown struct {
	// The value that pre-fills on the payment page.
	DefaultValue string `json:"default_value"`
	// The options available for the customer to select. Up to 200 options allowed.
	Options []*CheckoutSessionCustomFieldDropdownOption `json:"options"`
	// The option selected by the customer. This will be the `value` for the option.
	Value string `json:"value"`
}
type CheckoutSessionCustomFieldLabel struct {
	// Custom text for the label, displayed to the customer. Up to 50 characters.
	Custom string `json:"custom"`
	// The type of the label.
	Type CheckoutSessionCustomFieldLabelType `json:"type"`
}
type CheckoutSessionCustomFieldNumeric struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue string `json:"default_value"`
	// The maximum character length constraint for the customer's input.
	MaximumLength int64 `json:"maximum_length"`
	// The minimum character length requirement for the customer's input.
	MinimumLength int64 `json:"minimum_length"`
	// The value entered by the customer, containing only digits.
	Value string `json:"value"`
}
type CheckoutSessionCustomFieldText struct {
	// The value that pre-fills the field on the payment page.
	DefaultValue string `json:"default_value"`
	// The maximum character length constraint for the customer's input.
	MaximumLength int64 `json:"maximum_length"`
	// The minimum character length requirement for the customer's input.
	MinimumLength int64 `json:"minimum_length"`
	// The value entered by the customer.
	Value string `json:"value"`
}

// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
type CheckoutSessionCustomField struct {
	Dropdown *CheckoutSessionCustomFieldDropdown `json:"dropdown"`
	// String of your choice that your integration can use to reconcile this field. Must be unique to this field, alphanumeric, and up to 200 characters.
	Key     string                             `json:"key"`
	Label   *CheckoutSessionCustomFieldLabel   `json:"label"`
	Numeric *CheckoutSessionCustomFieldNumeric `json:"numeric"`
	// Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.
	Optional bool                            `json:"optional"`
	Text     *CheckoutSessionCustomFieldText `json:"text"`
	// The type of the field.
	Type CheckoutSessionCustomFieldType `json:"type"`
}

// Custom text that should be displayed after the payment confirmation button.
type CheckoutSessionCustomTextAfterSubmit struct {
	// Text can be up to 1200 characters in length.
	Message string `json:"message"`
}

// Custom text that should be displayed alongside shipping address collection.
type CheckoutSessionCustomTextShippingAddress struct {
	// Text can be up to 1200 characters in length.
	Message string `json:"message"`
}

// Custom text that should be displayed alongside the payment confirmation button.
type CheckoutSessionCustomTextSubmit struct {
	// Text can be up to 1200 characters in length.
	Message string `json:"message"`
}

// Custom text that should be displayed in place of the default terms of service agreement text.
type CheckoutSessionCustomTextTermsOfServiceAcceptance struct {
	// Text can be up to 1200 characters in length.
	Message string `json:"message"`
}
type CheckoutSessionCustomText struct {
	// Custom text that should be displayed after the payment confirmation button.
	AfterSubmit *CheckoutSessionCustomTextAfterSubmit `json:"after_submit"`
	// Custom text that should be displayed alongside shipping address collection.
	ShippingAddress *CheckoutSessionCustomTextShippingAddress `json:"shipping_address"`
	// Custom text that should be displayed alongside the payment confirmation button.
	Submit *CheckoutSessionCustomTextSubmit `json:"submit"`
	// Custom text that should be displayed in place of the default terms of service agreement text.
	TermsOfServiceAcceptance *CheckoutSessionCustomTextTermsOfServiceAcceptance `json:"terms_of_service_acceptance"`
}

// The customer's tax IDs after a completed Checkout Session.
type CheckoutSessionCustomerDetailsTaxID struct {
	// The type of the tax ID, one of `ad_nrt`, `ar_cuit`, `eu_vat`, `bo_tin`, `br_cnpj`, `br_cpf`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `eu_oss_vat`, `hr_oib`, `pe_ruc`, `ro_tin`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, `vn_tin`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `no_voec`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `pl_nip`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `li_vat`, `lk_vat`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, `al_tin`, `bh_vat`, `kz_bin`, `ng_tin`, `om_vat`, `de_stn`, `ch_uid`, `tz_vat`, `uz_vat`, `uz_tin`, `md_vat`, `ma_vat`, `by_tin`, `ao_tin`, `bs_tin`, `bb_tin`, `cd_nif`, `mr_nif`, `me_pib`, `zw_tin`, `ba_tin`, `gn_nif`, `mk_vat`, `sr_fin`, `sn_ninea`, `am_tin`, `np_pan`, `tj_tin`, `ug_tin`, `zm_tin`, `kh_tin`, `aw_tin`, `az_tin`, `bd_bin`, `bj_ifu`, `et_tin`, `kg_tin`, `la_tin`, `cm_niu`, `cv_nif`, `bf_ifu`, or `unknown`
	Type CheckoutSessionCustomerDetailsTaxIDType `json:"type"`
	// The value of the tax ID.
	Value string `json:"value"`
}

// The customer details including the customer's tax exempt status and the customer's tax IDs. Customer's address details are not present on Sessions in `setup` mode.
type CheckoutSessionCustomerDetails struct {
	// The customer's address after a completed Checkout Session. Note: This property is populated only for sessions on or after March 30, 2022.
	Address *Address `json:"address"`
	// The customer's business name after a completed Checkout Session.
	BusinessName string `json:"business_name"`
	// The email associated with the Customer, if one exists, on the Checkout Session after a completed Checkout Session or at time of session expiry.
	// Otherwise, if the customer has consented to promotional content, this value is the most recent valid email provided by the customer on the Checkout form.
	Email string `json:"email"`
	// The customer's individual name after a completed Checkout Session.
	IndividualName string `json:"individual_name"`
	// The customer's name after a completed Checkout Session. Note: This property is populated only for sessions on or after March 30, 2022.
	Name string `json:"name"`
	// The customer's phone number after a completed Checkout Session.
	Phone string `json:"phone"`
	// The customer's tax exempt status after a completed Checkout Session.
	TaxExempt CheckoutSessionCustomerDetailsTaxExempt `json:"tax_exempt"`
	// The customer's tax IDs after a completed Checkout Session.
	TaxIDs []*CheckoutSessionCustomerDetailsTaxID `json:"tax_ids"`
}

// List of coupons and promotion codes attached to the Checkout Session.
type CheckoutSessionDiscount struct {
	// Coupon attached to the Checkout Session.
	Coupon *Coupon `json:"coupon"`
	// Promotion code attached to the Checkout Session.
	PromotionCode *PromotionCode `json:"promotion_code"`
}

// Custom fields displayed on the invoice.
type CheckoutSessionInvoiceCreationInvoiceDataCustomField struct {
	// The name of the custom field.
	Name string `json:"name"`
	// The value of the custom field.
	Value string `json:"value"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type CheckoutSessionInvoiceCreationInvoiceDataIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type CheckoutSessionInvoiceCreationInvoiceDataIssuerType `json:"type"`
}

// Options for invoice PDF rendering.
type CheckoutSessionInvoiceCreationInvoiceDataRenderingOptions struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.
	AmountTaxDisplay string `json:"amount_tax_display"`
	// ID of the invoice rendering template to be used for the generated invoice.
	Template string `json:"template"`
}
type CheckoutSessionInvoiceCreationInvoiceData struct {
	// The account tax IDs associated with the invoice.
	AccountTaxIDs []*TaxID `json:"account_tax_ids"`
	// Custom fields displayed on the invoice.
	CustomFields []*CheckoutSessionInvoiceCreationInvoiceDataCustomField `json:"custom_fields"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Footer displayed on the invoice.
	Footer string `json:"footer"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *CheckoutSessionInvoiceCreationInvoiceDataIssuer `json:"issuer"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Options for invoice PDF rendering.
	RenderingOptions *CheckoutSessionInvoiceCreationInvoiceDataRenderingOptions `json:"rendering_options"`
}

// Details on the state of invoice creation for the Checkout Session.
type CheckoutSessionInvoiceCreation struct {
	// Indicates whether invoice creation is enabled for the Checkout Session.
	Enabled     bool                                       `json:"enabled"`
	InvoiceData *CheckoutSessionInvoiceCreationInvoiceData `json:"invoice_data"`
}
type CheckoutSessionNameCollectionBusiness struct {
	// Indicates whether business name collection is enabled for the session
	Enabled bool `json:"enabled"`
	// Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.
	Optional bool `json:"optional"`
}
type CheckoutSessionNameCollectionIndividual struct {
	// Indicates whether individual name collection is enabled for the session
	Enabled bool `json:"enabled"`
	// Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.
	Optional bool `json:"optional"`
}
type CheckoutSessionNameCollection struct {
	Business   *CheckoutSessionNameCollectionBusiness   `json:"business"`
	Individual *CheckoutSessionNameCollectionIndividual `json:"individual"`
}
type CheckoutSessionOptionalItemAdjustableQuantity struct {
	// Set to true if the quantity can be adjusted to any non-negative integer.
	Enabled bool `json:"enabled"`
	// The maximum quantity of this item the customer can purchase. By default this value is 99. You can specify a value up to 999999.
	Maximum int64 `json:"maximum"`
	// The minimum quantity of this item the customer must purchase, if they choose to purchase it. Because this item is optional, the customer will always be able to remove it from their order, even if the `minimum` configured here is greater than 0. By default this value is 0.
	Minimum int64 `json:"minimum"`
}

// The optional items presented to the customer at checkout.
type CheckoutSessionOptionalItem struct {
	AdjustableQuantity *CheckoutSessionOptionalItemAdjustableQuantity `json:"adjustable_quantity"`
	Price              string                                         `json:"price"`
	Quantity           int64                                          `json:"quantity"`
}

// Information about the payment method configuration used for this Checkout session if using dynamic payment methods.
type CheckoutSessionPaymentMethodConfigurationDetails struct {
	// ID of the payment method configuration used.
	ID string `json:"id"`
	// ID of the parent payment method configuration used.
	Parent string `json:"parent"`
}
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptions struct {
	// A URL for custom mandate text
	CustomMandateURL string `json:"custom_mandate_url"`
	// List of Stripe products where this mandate can be selected automatically. Returned when the Session is in `setup` mode.
	DefaultFor []CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor `json:"default_for"`
	// Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription string `json:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule `json:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}
type CheckoutSessionPaymentMethodOptionsACSSDebit struct {
	Currency       string                                                      `json:"currency"`
	MandateOptions *CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsACSSDebitSetupFutureUsage `json:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate string `json:"target_date"`
	// Bank account verification method.
	VerificationMethod CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}
type CheckoutSessionPaymentMethodOptionsAffirm struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsAffirmCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsAffirmSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsAfterpayClearpay struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsAfterpayClearpayCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsAfterpayClearpaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsAlipay struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsAlipaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsAlma struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsAlmaCaptureMethod `json:"capture_method"`
}
type CheckoutSessionPaymentMethodOptionsAmazonPay struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsAmazonPayCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsAmazonPaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsAUBECSDebit struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsAUBECSDebitSetupFutureUsage `json:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate string `json:"target_date"`
}
type CheckoutSessionPaymentMethodOptionsBACSDebitMandateOptions struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'DDIC' or 'STRIPE'.
	ReferencePrefix string `json:"reference_prefix"`
}
type CheckoutSessionPaymentMethodOptionsBACSDebit struct {
	MandateOptions *CheckoutSessionPaymentMethodOptionsBACSDebitMandateOptions `json:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage `json:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate string `json:"target_date"`
}
type CheckoutSessionPaymentMethodOptionsBancontact struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsBancontactSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsBillie struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsBillieCaptureMethod `json:"capture_method"`
}
type CheckoutSessionPaymentMethodOptionsBoleto struct {
	// The number of calendar days before a Boleto voucher expires. For example, if you create a Boleto voucher on Monday and you set expires_after_days to 2, the Boleto voucher will expire on Wednesday at 23:59 America/Sao_Paulo time.
	ExpiresAfterDays int64 `json:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsCardInstallments struct {
	// Indicates if installments are enabled
	Enabled bool `json:"enabled"`
}
type CheckoutSessionPaymentMethodOptionsCardRestrictions struct {
	// Specify the card brands to block in the Checkout Session. If a customer enters or selects a card belonging to a blocked brand, they can't complete the Session.
	BrandsBlocked []CheckoutSessionPaymentMethodOptionsCardRestrictionsBrandsBlocked `json:"brands_blocked"`
}
type CheckoutSessionPaymentMethodOptionsCard struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsCardCaptureMethod `json:"capture_method"`
	Installments  *CheckoutSessionPaymentMethodOptionsCardInstallments `json:"installments"`
	// Request ability to [capture beyond the standard authorization validity window](https://docs.stripe.com/payments/extended-authorization) for this CheckoutSession.
	RequestExtendedAuthorization CheckoutSessionPaymentMethodOptionsCardRequestExtendedAuthorization `json:"request_extended_authorization"`
	// Request ability to [increment the authorization](https://docs.stripe.com/payments/incremental-authorization) for this CheckoutSession.
	RequestIncrementalAuthorization CheckoutSessionPaymentMethodOptionsCardRequestIncrementalAuthorization `json:"request_incremental_authorization"`
	// Request ability to make [multiple captures](https://docs.stripe.com/payments/multicapture) for this CheckoutSession.
	RequestMulticapture CheckoutSessionPaymentMethodOptionsCardRequestMulticapture `json:"request_multicapture"`
	// Request ability to [overcapture](https://docs.stripe.com/payments/overcapture) for this CheckoutSession.
	RequestOvercapture CheckoutSessionPaymentMethodOptionsCardRequestOvercapture `json:"request_overcapture"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure CheckoutSessionPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
	Restrictions        *CheckoutSessionPaymentMethodOptionsCardRestrictions       `json:"restrictions"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage `json:"setup_future_usage"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kana prefix (shortened Kana descriptor) or Kana statement descriptor that's set on the account to form the complete statement descriptor. Maximum 22 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 22 characters.
	StatementDescriptorSuffixKana string `json:"statement_descriptor_suffix_kana"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kanji prefix (shortened Kanji descriptor) or Kanji statement descriptor that's set on the account to form the complete statement descriptor. Maximum 17 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 17 characters.
	StatementDescriptorSuffixKanji string `json:"statement_descriptor_suffix_kanji"`
}
type CheckoutSessionPaymentMethodOptionsCashApp struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsCashAppCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsCashAppSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer struct {
	// The desired country code of the bank account information. Permitted values include: `DE`, `FR`, `IE`, or `NL`.
	Country string `json:"country"`
}
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransfer struct {
	EUBankTransfer *CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer `json:"eu_bank_transfer"`
	// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
	//
	// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
	RequestedAddressTypes []CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType `json:"requested_address_types"`
	// The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.
	Type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType `json:"type"`
}
type CheckoutSessionPaymentMethodOptionsCustomerBalance struct {
	BankTransfer *CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType CheckoutSessionPaymentMethodOptionsCustomerBalanceFundingType `json:"funding_type"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsCustomerBalanceSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsEPS struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsEPSSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsFPX struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsFPXSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsGiropay struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsGiropaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsGrabpay struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsGrabpaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsIDEAL struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsIDEALSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsKakaoPay struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsKakaoPayCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsKakaoPaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsKlarna struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsKlarnaCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsKonbini struct {
	// The number of calendar days (between 1 and 60) after which Konbini payment instructions will expire. For example, if a PaymentIntent is confirmed with Konbini and `expires_after_days` set to 2 on Monday JST, the instructions will expire on Wednesday 23:59:59 JST.
	ExpiresAfterDays int64 `json:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsKonbiniSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsKrCard struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsKrCardCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsKrCardSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsLink struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsLinkCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsLinkSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsMobilepay struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsMobilepayCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsMobilepaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsMultibanco struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsMultibancoSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsNaverPay struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsNaverPayCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsNaverPaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsOXXO struct {
	// The number of calendar days before an OXXO invoice expires. For example, if you create an OXXO invoice on Monday and you set expires_after_days to 2, the OXXO invoice will expire on Wednesday at 23:59 America/Mexico_City time.
	ExpiresAfterDays int64 `json:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsOXXOSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsP24 struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsP24SetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsPayco struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsPaycoCaptureMethod `json:"capture_method"`
}
type CheckoutSessionPaymentMethodOptionsPayNow struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsPayNowSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsPaypal struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsPaypalCaptureMethod `json:"capture_method"`
	// Preferred locale of the PayPal checkout page that the customer is redirected to.
	PreferredLocale string `json:"preferred_locale"`
	// A reference of the PayPal transaction visible to customer which is mapped to PayPal's invoice ID. This must be a globally unique ID if you have configured in your PayPal settings to block multiple payments per invoice ID.
	Reference string `json:"reference"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsPaypalSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsPaytoMandateOptions struct {
	// Amount that will be collected. It is required when `amount_type` is `fixed`.
	Amount int64 `json:"amount"`
	// The type of amount that will be collected. The amount charged must be exact or up to the value of `amount` param for `fixed` or `maximum` type respectively. Defaults to `maximum`.
	AmountType CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsAmountType `json:"amount_type"`
	// Date, in YYYY-MM-DD format, after which payments will not be collected. Defaults to no end date.
	EndDate string `json:"end_date"`
	// The periodicity at which payments will be collected. Defaults to `adhoc`.
	PaymentSchedule CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPaymentSchedule `json:"payment_schedule"`
	// The number of payments that will be made during a payment period. Defaults to 1 except for when `payment_schedule` is `adhoc`. In that case, it defaults to no limit.
	PaymentsPerPeriod int64 `json:"payments_per_period"`
	// The purpose for which payments are made. Has a default value based on your merchant category code.
	Purpose CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsPurpose `json:"purpose"`
	// Date, in YYYY-MM-DD format, from which payments will be collected. Defaults to confirmation time.
	StartDate string `json:"start_date"`
}
type CheckoutSessionPaymentMethodOptionsPayto struct {
	MandateOptions *CheckoutSessionPaymentMethodOptionsPaytoMandateOptions `json:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsPaytoSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsPix struct {
	// Determines if the amount includes the IOF tax.
	AmountIncludesIof CheckoutSessionPaymentMethodOptionsPixAmountIncludesIof `json:"amount_includes_iof"`
	// The number of seconds after which Pix payment will expire.
	ExpiresAfterSeconds int64 `json:"expires_after_seconds"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsPixSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsRevolutPay struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsRevolutPayCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsRevolutPaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsSamsungPay struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsSamsungPayCaptureMethod `json:"capture_method"`
}
type CheckoutSessionPaymentMethodOptionsSatispay struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod CheckoutSessionPaymentMethodOptionsSatispayCaptureMethod `json:"capture_method"`
}
type CheckoutSessionPaymentMethodOptionsSEPADebitMandateOptions struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'STRIPE'.
	ReferencePrefix string `json:"reference_prefix"`
}
type CheckoutSessionPaymentMethodOptionsSEPADebit struct {
	MandateOptions *CheckoutSessionPaymentMethodOptionsSEPADebitMandateOptions `json:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsage `json:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate string `json:"target_date"`
}
type CheckoutSessionPaymentMethodOptionsSofort struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsSofortSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsSwish struct {
	// The order reference that will be displayed to customers in the Swish application. Defaults to the `id` of the Payment Intent.
	Reference string `json:"reference"`
}
type CheckoutSessionPaymentMethodOptionsTWINT struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsTWINTSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters struct {
	// The account subcategories to use to filter for possible accounts to link. Valid subcategories are `checking` and `savings`.
	AccountSubcategories []CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory `json:"account_subcategories"`
}
type CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	Filters *CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters `json:"filters"`
	// The list of permissions to request. The `payment_method` permission must be included.
	Permissions []CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
	// Data features requested to be retrieved upon account creation.
	Prefetch []CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch `json:"prefetch"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL string `json:"return_url"`
}
type CheckoutSessionPaymentMethodOptionsUSBankAccount struct {
	FinancialConnections *CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// If you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](https://docs.stripe.com/api/payment_methods/attach) the payment method to a Customer after the transaction completes.
	//
	// If the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.
	//
	// When processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](https://docs.stripe.com/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsUSBankAccountSetupFutureUsage `json:"setup_future_usage"`
	// Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.
	TargetDate string `json:"target_date"`
	// Bank account verification method.
	VerificationMethod CheckoutSessionPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment-method-specific configuration for the PaymentIntent or SetupIntent of this CheckoutSession.
type CheckoutSessionPaymentMethodOptions struct {
	ACSSDebit        *CheckoutSessionPaymentMethodOptionsACSSDebit        `json:"acss_debit"`
	Affirm           *CheckoutSessionPaymentMethodOptionsAffirm           `json:"affirm"`
	AfterpayClearpay *CheckoutSessionPaymentMethodOptionsAfterpayClearpay `json:"afterpay_clearpay"`
	Alipay           *CheckoutSessionPaymentMethodOptionsAlipay           `json:"alipay"`
	Alma             *CheckoutSessionPaymentMethodOptionsAlma             `json:"alma"`
	AmazonPay        *CheckoutSessionPaymentMethodOptionsAmazonPay        `json:"amazon_pay"`
	AUBECSDebit      *CheckoutSessionPaymentMethodOptionsAUBECSDebit      `json:"au_becs_debit"`
	BACSDebit        *CheckoutSessionPaymentMethodOptionsBACSDebit        `json:"bacs_debit"`
	Bancontact       *CheckoutSessionPaymentMethodOptionsBancontact       `json:"bancontact"`
	Billie           *CheckoutSessionPaymentMethodOptionsBillie           `json:"billie"`
	Boleto           *CheckoutSessionPaymentMethodOptionsBoleto           `json:"boleto"`
	Card             *CheckoutSessionPaymentMethodOptionsCard             `json:"card"`
	CashApp          *CheckoutSessionPaymentMethodOptionsCashApp          `json:"cashapp"`
	CustomerBalance  *CheckoutSessionPaymentMethodOptionsCustomerBalance  `json:"customer_balance"`
	EPS              *CheckoutSessionPaymentMethodOptionsEPS              `json:"eps"`
	FPX              *CheckoutSessionPaymentMethodOptionsFPX              `json:"fpx"`
	Giropay          *CheckoutSessionPaymentMethodOptionsGiropay          `json:"giropay"`
	Grabpay          *CheckoutSessionPaymentMethodOptionsGrabpay          `json:"grabpay"`
	IDEAL            *CheckoutSessionPaymentMethodOptionsIDEAL            `json:"ideal"`
	KakaoPay         *CheckoutSessionPaymentMethodOptionsKakaoPay         `json:"kakao_pay"`
	Klarna           *CheckoutSessionPaymentMethodOptionsKlarna           `json:"klarna"`
	Konbini          *CheckoutSessionPaymentMethodOptionsKonbini          `json:"konbini"`
	KrCard           *CheckoutSessionPaymentMethodOptionsKrCard           `json:"kr_card"`
	Link             *CheckoutSessionPaymentMethodOptionsLink             `json:"link"`
	Mobilepay        *CheckoutSessionPaymentMethodOptionsMobilepay        `json:"mobilepay"`
	Multibanco       *CheckoutSessionPaymentMethodOptionsMultibanco       `json:"multibanco"`
	NaverPay         *CheckoutSessionPaymentMethodOptionsNaverPay         `json:"naver_pay"`
	OXXO             *CheckoutSessionPaymentMethodOptionsOXXO             `json:"oxxo"`
	P24              *CheckoutSessionPaymentMethodOptionsP24              `json:"p24"`
	Payco            *CheckoutSessionPaymentMethodOptionsPayco            `json:"payco"`
	PayNow           *CheckoutSessionPaymentMethodOptionsPayNow           `json:"paynow"`
	Paypal           *CheckoutSessionPaymentMethodOptionsPaypal           `json:"paypal"`
	Payto            *CheckoutSessionPaymentMethodOptionsPayto            `json:"payto"`
	Pix              *CheckoutSessionPaymentMethodOptionsPix              `json:"pix"`
	RevolutPay       *CheckoutSessionPaymentMethodOptionsRevolutPay       `json:"revolut_pay"`
	SamsungPay       *CheckoutSessionPaymentMethodOptionsSamsungPay       `json:"samsung_pay"`
	Satispay         *CheckoutSessionPaymentMethodOptionsSatispay         `json:"satispay"`
	SEPADebit        *CheckoutSessionPaymentMethodOptionsSEPADebit        `json:"sepa_debit"`
	Sofort           *CheckoutSessionPaymentMethodOptionsSofort           `json:"sofort"`
	Swish            *CheckoutSessionPaymentMethodOptionsSwish            `json:"swish"`
	TWINT            *CheckoutSessionPaymentMethodOptionsTWINT            `json:"twint"`
	USBankAccount    *CheckoutSessionPaymentMethodOptionsUSBankAccount    `json:"us_bank_account"`
}

// This property is used to set up permissions for various actions (e.g., update) on the CheckoutSession object.
//
// For specific permissions, please refer to their dedicated subsections, such as `permissions.update_shipping_details`.
type CheckoutSessionPermissions struct {
	// Determines which entity is allowed to update the shipping details.
	//
	// Default is `client_only`. Stripe Checkout client will automatically update the shipping details. If set to `server_only`, only your server is allowed to update the shipping details.
	//
	// When set to `server_only`, you must add the onShippingDetailsChange event handler when initializing the Stripe Checkout client and manually update the shipping details from your server using the Stripe API.
	UpdateShippingDetails CheckoutSessionPermissionsUpdateShippingDetails `json:"update_shipping_details"`
}
type CheckoutSessionPhoneNumberCollection struct {
	// Indicates whether phone number collection is enabled for the session
	Enabled bool `json:"enabled"`
}
type CheckoutSessionPresentmentDetails struct {
	// Amount intended to be collected by this payment, denominated in `presentment_currency`.
	PresentmentAmount int64 `json:"presentment_amount"`
	// Currency presented to the customer during payment.
	PresentmentCurrency Currency `json:"presentment_currency"`
}

// Controls saved payment method settings for the session. Only available in `payment` and `subscription` mode.
type CheckoutSessionSavedPaymentMethodOptions struct {
	// Uses the `allow_redisplay` value of each saved payment method to filter the set presented to a returning customer. By default, only saved payment methods with 'allow_redisplay: ‘always' are shown in Checkout.
	AllowRedisplayFilters []CheckoutSessionSavedPaymentMethodOptionsAllowRedisplayFilter `json:"allow_redisplay_filters"`
	// Enable customers to choose if they wish to remove their saved payment methods. Disabled by default.
	PaymentMethodRemove CheckoutSessionSavedPaymentMethodOptionsPaymentMethodRemove `json:"payment_method_remove"`
	// Enable customers to choose if they wish to save their payment method for future use. Disabled by default.
	PaymentMethodSave CheckoutSessionSavedPaymentMethodOptionsPaymentMethodSave `json:"payment_method_save"`
}

// When set, provides configuration for Checkout to collect a shipping address from a customer.
type CheckoutSessionShippingAddressCollection struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for
	// shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SY, UM, VI`.
	AllowedCountries []string `json:"allowed_countries"`
}

// The taxes applied to the shipping rate.
type CheckoutSessionShippingCostTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://docs.stripe.com/invoicing/taxes/tax-rates), [subscriptions](https://docs.stripe.com/billing/taxes/tax-rates) and [Checkout Sessions](https://docs.stripe.com/payments/checkout/use-manual-tax-rates) to collect tax.
	//
	// Related guide: [Tax rates](https://docs.stripe.com/billing/taxes/tax-rates)
	Rate *TaxRate `json:"rate"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason CheckoutSessionShippingCostTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
}

// The details of the customer cost of shipping, including the customer chosen ShippingRate.
type CheckoutSessionShippingCost struct {
	// Total shipping cost before any discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total tax amount applied due to shipping costs. If no tax was applied, defaults to 0.
	AmountTax int64 `json:"amount_tax"`
	// Total shipping cost after discounts and taxes are applied.
	AmountTotal int64 `json:"amount_total"`
	// The ID of the ShippingRate for this order.
	ShippingRate *ShippingRate `json:"shipping_rate"`
	// The taxes applied to the shipping rate.
	Taxes []*CheckoutSessionShippingCostTax `json:"taxes"`
}

// The shipping rate options applied to this Session.
type CheckoutSessionShippingOption struct {
	// A non-negative integer in cents representing how much to charge.
	ShippingAmount int64 `json:"shipping_amount"`
	// The shipping rate.
	ShippingRate *ShippingRate `json:"shipping_rate"`
}
type CheckoutSessionTaxIDCollection struct {
	// Indicates whether tax ID collection is enabled for the session
	Enabled bool `json:"enabled"`
	// Indicates whether a tax ID is required on the payment page
	Required CheckoutSessionTaxIDCollectionRequired `json:"required"`
}

// The aggregated discounts.
type CheckoutSessionTotalDetailsBreakdownDiscount struct {
	// The amount discounted.
	Amount int64 `json:"amount"`
	// A discount represents the actual application of a [coupon](https://api.stripe.com#coupons) or [promotion code](https://api.stripe.com#promotion_codes).
	// It contains information about when the discount began, when it will end, and what it is applied to.
	//
	// Related guide: [Applying discounts to subscriptions](https://docs.stripe.com/billing/subscriptions/discounts)
	Discount *Discount `json:"discount"`
}

// The aggregated tax amounts by rate.
type CheckoutSessionTotalDetailsBreakdownTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://docs.stripe.com/invoicing/taxes/tax-rates), [subscriptions](https://docs.stripe.com/billing/taxes/tax-rates) and [Checkout Sessions](https://docs.stripe.com/payments/checkout/use-manual-tax-rates) to collect tax.
	//
	// Related guide: [Tax rates](https://docs.stripe.com/billing/taxes/tax-rates)
	Rate *TaxRate `json:"rate"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
}
type CheckoutSessionTotalDetailsBreakdown struct {
	// The aggregated discounts.
	Discounts []*CheckoutSessionTotalDetailsBreakdownDiscount `json:"discounts"`
	// The aggregated tax amounts by rate.
	Taxes []*CheckoutSessionTotalDetailsBreakdownTax `json:"taxes"`
}

// Tax and discount details for the computed total amount.
type CheckoutSessionTotalDetails struct {
	// This is the sum of all the discounts.
	AmountDiscount int64 `json:"amount_discount"`
	// This is the sum of all the shipping amounts.
	AmountShipping int64 `json:"amount_shipping"`
	// This is the sum of all the tax amounts.
	AmountTax int64                                 `json:"amount_tax"`
	Breakdown *CheckoutSessionTotalDetailsBreakdown `json:"breakdown"`
}
type CheckoutSessionWalletOptionsLink struct {
	// Describes whether Checkout should display Link. Defaults to `auto`.
	Display CheckoutSessionWalletOptionsLinkDisplay `json:"display"`
}

// Wallet-specific configuration for this Checkout Session.
type CheckoutSessionWalletOptions struct {
	Link *CheckoutSessionWalletOptionsLink `json:"link"`
}

// A Checkout Session represents your customer's session as they pay for
// one-time purchases or subscriptions through [Checkout](https://docs.stripe.com/payments/checkout)
// or [Payment Links](https://docs.stripe.com/payments/payment-links). We recommend creating a
// new Session each time your customer attempts to pay.
//
// Once payment is successful, the Checkout Session will contain a reference
// to the [Customer](https://docs.stripe.com/api/customers), and either the successful
// [PaymentIntent](https://docs.stripe.com/api/payment_intents) or an active
// [Subscription](https://docs.stripe.com/api/subscriptions).
//
// You can create a Checkout Session on your server and redirect to its URL
// to begin Checkout.
//
// Related guide: [Checkout quickstart](https://docs.stripe.com/checkout/quickstart)
type CheckoutSession struct {
	APIResource
	// Settings for price localization with [Adaptive Pricing](https://docs.stripe.com/payments/checkout/adaptive-pricing).
	AdaptivePricing *CheckoutSessionAdaptivePricing `json:"adaptive_pricing"`
	// When set, provides configuration for actions to take if this Checkout Session expires.
	AfterExpiration *CheckoutSessionAfterExpiration `json:"after_expiration"`
	// Enables user redeemable promotion codes.
	AllowPromotionCodes bool `json:"allow_promotion_codes"`
	// Total of all items before discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total of all items after discounts and taxes are applied.
	AmountTotal  int64                        `json:"amount_total"`
	AutomaticTax *CheckoutSessionAutomaticTax `json:"automatic_tax"`
	// Describes whether Checkout should collect the customer's billing address. Defaults to `auto`.
	BillingAddressCollection CheckoutSessionBillingAddressCollection `json:"billing_address_collection"`
	BrandingSettings         *CheckoutSessionBrandingSettings        `json:"branding_settings"`
	// If set, Checkout displays a back button and customers will be directed to this URL if they decide to cancel payment and return to your website.
	CancelURL string `json:"cancel_url"`
	// A unique string to reference the Checkout Session. This can be a
	// customer ID, a cart ID, or similar, and can be used to reconcile the
	// Session with your internal systems.
	ClientReferenceID string `json:"client_reference_id"`
	// The client secret of your Checkout Session. Applies to Checkout Sessions with `ui_mode: embedded` or `ui_mode: custom`. For `ui_mode: embedded`, the client secret is to be used when initializing Stripe.js embedded checkout.
	//  For `ui_mode: custom`, use the client secret with [initCheckout](https://docs.stripe.com/js/custom_checkout/init) on your front end.
	ClientSecret string `json:"client_secret"`
	// Information about the customer collected within the Checkout Session.
	CollectedInformation *CheckoutSessionCollectedInformation `json:"collected_information"`
	// Results of `consent_collection` for this session.
	Consent *CheckoutSessionConsent `json:"consent"`
	// When set, provides configuration for the Checkout Session to gather active consent from customers.
	ConsentCollection *CheckoutSessionConsentCollection `json:"consent_collection"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Currency conversion details for [Adaptive Pricing](https://docs.stripe.com/payments/checkout/adaptive-pricing) sessions created before 2025-03-31.
	CurrencyConversion *CheckoutSessionCurrencyConversion `json:"currency_conversion"`
	// The ID of the customer for this Session.
	// For Checkout Sessions in `subscription` mode or Checkout Sessions with `customer_creation` set as `always` in `payment` mode, Checkout
	// will create a new customer object based on information provided
	// during the payment flow unless an existing customer was provided when
	// the Session was created.
	Customer *Customer `json:"customer"`
	// The ID of the account for this Session.
	CustomerAccount string `json:"customer_account"`
	// Configure whether a Checkout Session creates a Customer when the Checkout Session completes.
	CustomerCreation CheckoutSessionCustomerCreation `json:"customer_creation"`
	// The customer details including the customer's tax exempt status and the customer's tax IDs. Customer's address details are not present on Sessions in `setup` mode.
	CustomerDetails *CheckoutSessionCustomerDetails `json:"customer_details"`
	// If provided, this value will be used when the Customer object is created.
	// If not provided, customers will be asked to enter their email address.
	// Use this parameter to prefill customer data if you already have an email
	// on file. To access information about the customer once the payment flow is
	// complete, use the `customer` attribute.
	CustomerEmail string `json:"customer_email"`
	// Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.
	CustomFields []*CheckoutSessionCustomField `json:"custom_fields"`
	CustomText   *CheckoutSessionCustomText    `json:"custom_text"`
	// List of coupons and promotion codes attached to the Checkout Session.
	Discounts []*CheckoutSessionDiscount `json:"discounts"`
	// A list of the types of payment methods (e.g., `card`) that should be excluded from this Checkout Session. This should only be used when payment methods for this Checkout Session are managed through the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).
	ExcludedPaymentMethodTypes []string `json:"excluded_payment_method_types"`
	// The timestamp at which the Checkout Session will expire.
	ExpiresAt int64 `json:"expires_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// ID of the invoice created by the Checkout Session, if it exists.
	Invoice *Invoice `json:"invoice"`
	// Details on the state of invoice creation for the Checkout Session.
	InvoiceCreation *CheckoutSessionInvoiceCreation `json:"invoice_creation"`
	// The line items purchased by the customer.
	LineItems *LineItemList `json:"line_items"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The IETF language tag of the locale Checkout is displayed in. If blank or `auto`, the browser's locale is used.
	Locale string `json:"locale"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The mode of the Checkout Session.
	Mode           CheckoutSessionMode            `json:"mode"`
	NameCollection *CheckoutSessionNameCollection `json:"name_collection"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The optional items presented to the customer at checkout.
	OptionalItems []*CheckoutSessionOptionalItem `json:"optional_items"`
	// Where the user is coming from. This informs the optimizations that are applied to the session.
	OriginContext CheckoutSessionOriginContext `json:"origin_context"`
	// The ID of the PaymentIntent for Checkout Sessions in `payment` mode. You can't confirm or cancel the PaymentIntent for a Checkout Session. To cancel, [expire the Checkout Session](https://docs.stripe.com/api/checkout/sessions/expire) instead.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
	// The ID of the Payment Link that created this Session.
	PaymentLink *PaymentLink `json:"payment_link"`
	// Configure whether a Checkout Session should collect a payment method. Defaults to `always`.
	PaymentMethodCollection CheckoutSessionPaymentMethodCollection `json:"payment_method_collection"`
	// Information about the payment method configuration used for this Checkout session if using dynamic payment methods.
	PaymentMethodConfigurationDetails *CheckoutSessionPaymentMethodConfigurationDetails `json:"payment_method_configuration_details"`
	// Payment-method-specific configuration for the PaymentIntent or SetupIntent of this CheckoutSession.
	PaymentMethodOptions *CheckoutSessionPaymentMethodOptions `json:"payment_method_options"`
	// A list of the types of payment methods (e.g. card) this Checkout
	// Session is allowed to accept.
	PaymentMethodTypes []string `json:"payment_method_types"`
	// The payment status of the Checkout Session, one of `paid`, `unpaid`, or `no_payment_required`.
	// You can use this value to decide when to fulfill your customer's order.
	PaymentStatus CheckoutSessionPaymentStatus `json:"payment_status"`
	// This property is used to set up permissions for various actions (e.g., update) on the CheckoutSession object.
	//
	// For specific permissions, please refer to their dedicated subsections, such as `permissions.update_shipping_details`.
	Permissions           *CheckoutSessionPermissions           `json:"permissions"`
	PhoneNumberCollection *CheckoutSessionPhoneNumberCollection `json:"phone_number_collection"`
	PresentmentDetails    *CheckoutSessionPresentmentDetails    `json:"presentment_details"`
	// The ID of the original expired Checkout Session that triggered the recovery flow.
	RecoveredFrom string `json:"recovered_from"`
	// This parameter applies to `ui_mode: embedded`. Learn more about the [redirect behavior](https://docs.stripe.com/payments/checkout/custom-success-page?payment-ui=embedded-form) of embedded sessions. Defaults to `always`.
	RedirectOnCompletion CheckoutSessionRedirectOnCompletion `json:"redirect_on_completion"`
	// Applies to Checkout Sessions with `ui_mode: embedded` or `ui_mode: custom`. The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method's app or site.
	ReturnURL string `json:"return_url"`
	// Controls saved payment method settings for the session. Only available in `payment` and `subscription` mode.
	SavedPaymentMethodOptions *CheckoutSessionSavedPaymentMethodOptions `json:"saved_payment_method_options"`
	// The ID of the SetupIntent for Checkout Sessions in `setup` mode. You can't confirm or cancel the SetupIntent for a Checkout Session. To cancel, [expire the Checkout Session](https://docs.stripe.com/api/checkout/sessions/expire) instead.
	SetupIntent *SetupIntent `json:"setup_intent"`
	// When set, provides configuration for Checkout to collect a shipping address from a customer.
	ShippingAddressCollection *CheckoutSessionShippingAddressCollection `json:"shipping_address_collection"`
	// The details of the customer cost of shipping, including the customer chosen ShippingRate.
	ShippingCost *CheckoutSessionShippingCost `json:"shipping_cost"`
	// The shipping rate options applied to this Session.
	ShippingOptions []*CheckoutSessionShippingOption `json:"shipping_options"`
	// The status of the Checkout Session, one of `open`, `complete`, or `expired`.
	Status CheckoutSessionStatus `json:"status"`
	// Describes the type of transaction being performed by Checkout in order to customize
	// relevant text on the page, such as the submit button. `submit_type` can only be
	// specified on Checkout Sessions in `payment` mode. If blank or `auto`, `pay` is used.
	SubmitType CheckoutSessionSubmitType `json:"submit_type"`
	// The ID of the [Subscription](https://docs.stripe.com/api/subscriptions) for Checkout Sessions in `subscription` mode.
	Subscription *Subscription `json:"subscription"`
	// The URL the customer will be directed to after the payment or
	// subscription creation is successful.
	SuccessURL      string                          `json:"success_url"`
	TaxIDCollection *CheckoutSessionTaxIDCollection `json:"tax_id_collection"`
	// Tax and discount details for the computed total amount.
	TotalDetails *CheckoutSessionTotalDetails `json:"total_details"`
	// The UI mode of the Session. Defaults to `hosted`.
	UIMode CheckoutSessionUIMode `json:"ui_mode"`
	// The URL to the Checkout Session. Applies to Checkout Sessions with `ui_mode: hosted`. Redirect customers to this URL to take them to Checkout. If you're using [Custom Domains](https://docs.stripe.com/payments/checkout/custom-domains), the URL will use your subdomain. Otherwise, it'll use `checkout.stripe.com.`
	// This value is only present when the session is active.
	URL string `json:"url"`
	// Wallet-specific configuration for this Checkout Session.
	WalletOptions *CheckoutSessionWalletOptions `json:"wallet_options"`
}

// CheckoutSessionList is a list of Sessions as retrieved from a list endpoint.
type CheckoutSessionList struct {
	APIResource
	ListMeta
	Data []*CheckoutSession `json:"data"`
}

// UnmarshalJSON handles deserialization of a CheckoutSession.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *CheckoutSession) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type checkoutSession CheckoutSession
	var v checkoutSession
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = CheckoutSession(v)
	return nil
}
