//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Describes whether the quote line is affecting a new schedule or an existing schedule.
type QuotePreviewInvoiceAppliesToType string

// List of values that QuotePreviewInvoiceAppliesToType can take
const (
	QuotePreviewInvoiceAppliesToTypeNewReference         QuotePreviewInvoiceAppliesToType = "new_reference"
	QuotePreviewInvoiceAppliesToTypeSubscriptionSchedule QuotePreviewInvoiceAppliesToType = "subscription_schedule"
)

// Type of the account referenced.
type QuotePreviewInvoiceAutomaticTaxLiabilityType string

// List of values that QuotePreviewInvoiceAutomaticTaxLiabilityType can take
const (
	QuotePreviewInvoiceAutomaticTaxLiabilityTypeAccount QuotePreviewInvoiceAutomaticTaxLiabilityType = "account"
	QuotePreviewInvoiceAutomaticTaxLiabilityTypeSelf    QuotePreviewInvoiceAutomaticTaxLiabilityType = "self"
)

// The status of the most recent automated tax calculation for this invoice.
type QuotePreviewInvoiceAutomaticTaxStatus string

// List of values that QuotePreviewInvoiceAutomaticTaxStatus can take
const (
	QuotePreviewInvoiceAutomaticTaxStatusComplete               QuotePreviewInvoiceAutomaticTaxStatus = "complete"
	QuotePreviewInvoiceAutomaticTaxStatusFailed                 QuotePreviewInvoiceAutomaticTaxStatus = "failed"
	QuotePreviewInvoiceAutomaticTaxStatusRequiresLocationInputs QuotePreviewInvoiceAutomaticTaxStatus = "requires_location_inputs"
)

// Indicates the reason why the invoice was created.
//
// * `manual`: Unrelated to a subscription, for example, created via the invoice editor.
// * `subscription`: No longer in use. Applies to subscriptions from before May 2018 where no distinction was made between updates, cycles, and thresholds.
// * `subscription_create`: A new subscription was created.
// * `subscription_cycle`: A subscription advanced into a new period.
// * `subscription_threshold`: A subscription reached a billing threshold.
// * `subscription_update`: A subscription was updated.
// * `upcoming`: Reserved for simulated invoices, per the upcoming invoice endpoint.
type QuotePreviewInvoiceBillingReason string

// List of values that QuotePreviewInvoiceBillingReason can take
const (
	QuotePreviewInvoiceBillingReasonAutomaticPendingInvoiceItemInvoice QuotePreviewInvoiceBillingReason = "automatic_pending_invoice_item_invoice"
	QuotePreviewInvoiceBillingReasonManual                             QuotePreviewInvoiceBillingReason = "manual"
	QuotePreviewInvoiceBillingReasonQuoteAccept                        QuotePreviewInvoiceBillingReason = "quote_accept"
	QuotePreviewInvoiceBillingReasonSubscription                       QuotePreviewInvoiceBillingReason = "subscription"
	QuotePreviewInvoiceBillingReasonSubscriptionCreate                 QuotePreviewInvoiceBillingReason = "subscription_create"
	QuotePreviewInvoiceBillingReasonSubscriptionCycle                  QuotePreviewInvoiceBillingReason = "subscription_cycle"
	QuotePreviewInvoiceBillingReasonSubscriptionThreshold              QuotePreviewInvoiceBillingReason = "subscription_threshold"
	QuotePreviewInvoiceBillingReasonSubscriptionUpdate                 QuotePreviewInvoiceBillingReason = "subscription_update"
	QuotePreviewInvoiceBillingReasonUpcoming                           QuotePreviewInvoiceBillingReason = "upcoming"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions.
type QuotePreviewInvoiceCollectionMethod string

// List of values that QuotePreviewInvoiceCollectionMethod can take
const (
	QuotePreviewInvoiceCollectionMethodChargeAutomatically QuotePreviewInvoiceCollectionMethod = "charge_automatically"
	QuotePreviewInvoiceCollectionMethodSendInvoice         QuotePreviewInvoiceCollectionMethod = "send_invoice"
)

// The customer's tax exempt status. Until the invoice is finalized, this field will equal `customer.tax_exempt`. Once the invoice is finalized, this field will no longer be updated.
type QuotePreviewInvoiceCustomerTaxExempt string

// List of values that QuotePreviewInvoiceCustomerTaxExempt can take
const (
	QuotePreviewInvoiceCustomerTaxExemptExempt  QuotePreviewInvoiceCustomerTaxExempt = "exempt"
	QuotePreviewInvoiceCustomerTaxExemptNone    QuotePreviewInvoiceCustomerTaxExempt = "none"
	QuotePreviewInvoiceCustomerTaxExemptReverse QuotePreviewInvoiceCustomerTaxExempt = "reverse"
)

// The type of the tax ID, one of `ad_nrt`, `ar_cuit`, `eu_vat`, `bo_tin`, `br_cnpj`, `br_cpf`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `eu_oss_vat`, `pe_ruc`, `ro_tin`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, `vn_tin`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, or `unknown`
type QuotePreviewInvoiceCustomerTaxIDType string

// List of values that QuotePreviewInvoiceCustomerTaxIDType can take
const (
	QuotePreviewInvoiceCustomerTaxIDTypeADNRT    QuotePreviewInvoiceCustomerTaxIDType = "ad_nrt"
	QuotePreviewInvoiceCustomerTaxIDTypeAETRN    QuotePreviewInvoiceCustomerTaxIDType = "ae_trn"
	QuotePreviewInvoiceCustomerTaxIDTypeARCUIT   QuotePreviewInvoiceCustomerTaxIDType = "ar_cuit"
	QuotePreviewInvoiceCustomerTaxIDTypeAUABN    QuotePreviewInvoiceCustomerTaxIDType = "au_abn"
	QuotePreviewInvoiceCustomerTaxIDTypeAUARN    QuotePreviewInvoiceCustomerTaxIDType = "au_arn"
	QuotePreviewInvoiceCustomerTaxIDTypeBGUIC    QuotePreviewInvoiceCustomerTaxIDType = "bg_uic"
	QuotePreviewInvoiceCustomerTaxIDTypeBOTIN    QuotePreviewInvoiceCustomerTaxIDType = "bo_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeBRCNPJ   QuotePreviewInvoiceCustomerTaxIDType = "br_cnpj"
	QuotePreviewInvoiceCustomerTaxIDTypeBRCPF    QuotePreviewInvoiceCustomerTaxIDType = "br_cpf"
	QuotePreviewInvoiceCustomerTaxIDTypeCABN     QuotePreviewInvoiceCustomerTaxIDType = "ca_bn"
	QuotePreviewInvoiceCustomerTaxIDTypeCAGSTHST QuotePreviewInvoiceCustomerTaxIDType = "ca_gst_hst"
	QuotePreviewInvoiceCustomerTaxIDTypeCAPSTBC  QuotePreviewInvoiceCustomerTaxIDType = "ca_pst_bc"
	QuotePreviewInvoiceCustomerTaxIDTypeCAPSTMB  QuotePreviewInvoiceCustomerTaxIDType = "ca_pst_mb"
	QuotePreviewInvoiceCustomerTaxIDTypeCAPSTSK  QuotePreviewInvoiceCustomerTaxIDType = "ca_pst_sk"
	QuotePreviewInvoiceCustomerTaxIDTypeCAQST    QuotePreviewInvoiceCustomerTaxIDType = "ca_qst"
	QuotePreviewInvoiceCustomerTaxIDTypeCHVAT    QuotePreviewInvoiceCustomerTaxIDType = "ch_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeCLTIN    QuotePreviewInvoiceCustomerTaxIDType = "cl_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeCNTIN    QuotePreviewInvoiceCustomerTaxIDType = "cn_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeCONIT    QuotePreviewInvoiceCustomerTaxIDType = "co_nit"
	QuotePreviewInvoiceCustomerTaxIDTypeCRTIN    QuotePreviewInvoiceCustomerTaxIDType = "cr_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeDORCN    QuotePreviewInvoiceCustomerTaxIDType = "do_rcn"
	QuotePreviewInvoiceCustomerTaxIDTypeECRUC    QuotePreviewInvoiceCustomerTaxIDType = "ec_ruc"
	QuotePreviewInvoiceCustomerTaxIDTypeEGTIN    QuotePreviewInvoiceCustomerTaxIDType = "eg_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeESCIF    QuotePreviewInvoiceCustomerTaxIDType = "es_cif"
	QuotePreviewInvoiceCustomerTaxIDTypeEUOSSVAT QuotePreviewInvoiceCustomerTaxIDType = "eu_oss_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeEUVAT    QuotePreviewInvoiceCustomerTaxIDType = "eu_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeGBVAT    QuotePreviewInvoiceCustomerTaxIDType = "gb_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeGEVAT    QuotePreviewInvoiceCustomerTaxIDType = "ge_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeHKBR     QuotePreviewInvoiceCustomerTaxIDType = "hk_br"
	QuotePreviewInvoiceCustomerTaxIDTypeHUTIN    QuotePreviewInvoiceCustomerTaxIDType = "hu_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeIDNPWP   QuotePreviewInvoiceCustomerTaxIDType = "id_npwp"
	QuotePreviewInvoiceCustomerTaxIDTypeILVAT    QuotePreviewInvoiceCustomerTaxIDType = "il_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeINGST    QuotePreviewInvoiceCustomerTaxIDType = "in_gst"
	QuotePreviewInvoiceCustomerTaxIDTypeISVAT    QuotePreviewInvoiceCustomerTaxIDType = "is_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeJPCN     QuotePreviewInvoiceCustomerTaxIDType = "jp_cn"
	QuotePreviewInvoiceCustomerTaxIDTypeJPRN     QuotePreviewInvoiceCustomerTaxIDType = "jp_rn"
	QuotePreviewInvoiceCustomerTaxIDTypeJPTRN    QuotePreviewInvoiceCustomerTaxIDType = "jp_trn"
	QuotePreviewInvoiceCustomerTaxIDTypeKEPIN    QuotePreviewInvoiceCustomerTaxIDType = "ke_pin"
	QuotePreviewInvoiceCustomerTaxIDTypeKRBRN    QuotePreviewInvoiceCustomerTaxIDType = "kr_brn"
	QuotePreviewInvoiceCustomerTaxIDTypeLIUID    QuotePreviewInvoiceCustomerTaxIDType = "li_uid"
	QuotePreviewInvoiceCustomerTaxIDTypeMXRFC    QuotePreviewInvoiceCustomerTaxIDType = "mx_rfc"
	QuotePreviewInvoiceCustomerTaxIDTypeMYFRP    QuotePreviewInvoiceCustomerTaxIDType = "my_frp"
	QuotePreviewInvoiceCustomerTaxIDTypeMYITN    QuotePreviewInvoiceCustomerTaxIDType = "my_itn"
	QuotePreviewInvoiceCustomerTaxIDTypeMYSST    QuotePreviewInvoiceCustomerTaxIDType = "my_sst"
	QuotePreviewInvoiceCustomerTaxIDTypeNOVAT    QuotePreviewInvoiceCustomerTaxIDType = "no_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeNZGST    QuotePreviewInvoiceCustomerTaxIDType = "nz_gst"
	QuotePreviewInvoiceCustomerTaxIDTypePERUC    QuotePreviewInvoiceCustomerTaxIDType = "pe_ruc"
	QuotePreviewInvoiceCustomerTaxIDTypePHTIN    QuotePreviewInvoiceCustomerTaxIDType = "ph_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeROTIN    QuotePreviewInvoiceCustomerTaxIDType = "ro_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeRSPIB    QuotePreviewInvoiceCustomerTaxIDType = "rs_pib"
	QuotePreviewInvoiceCustomerTaxIDTypeRUINN    QuotePreviewInvoiceCustomerTaxIDType = "ru_inn"
	QuotePreviewInvoiceCustomerTaxIDTypeRUKPP    QuotePreviewInvoiceCustomerTaxIDType = "ru_kpp"
	QuotePreviewInvoiceCustomerTaxIDTypeSAVAT    QuotePreviewInvoiceCustomerTaxIDType = "sa_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeSGGST    QuotePreviewInvoiceCustomerTaxIDType = "sg_gst"
	QuotePreviewInvoiceCustomerTaxIDTypeSGUEN    QuotePreviewInvoiceCustomerTaxIDType = "sg_uen"
	QuotePreviewInvoiceCustomerTaxIDTypeSITIN    QuotePreviewInvoiceCustomerTaxIDType = "si_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeSVNIT    QuotePreviewInvoiceCustomerTaxIDType = "sv_nit"
	QuotePreviewInvoiceCustomerTaxIDTypeTHVAT    QuotePreviewInvoiceCustomerTaxIDType = "th_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeTRTIN    QuotePreviewInvoiceCustomerTaxIDType = "tr_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeTWVAT    QuotePreviewInvoiceCustomerTaxIDType = "tw_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeUAVAT    QuotePreviewInvoiceCustomerTaxIDType = "ua_vat"
	QuotePreviewInvoiceCustomerTaxIDTypeUnknown  QuotePreviewInvoiceCustomerTaxIDType = "unknown"
	QuotePreviewInvoiceCustomerTaxIDTypeUSEIN    QuotePreviewInvoiceCustomerTaxIDType = "us_ein"
	QuotePreviewInvoiceCustomerTaxIDTypeUYRUC    QuotePreviewInvoiceCustomerTaxIDType = "uy_ruc"
	QuotePreviewInvoiceCustomerTaxIDTypeVERIF    QuotePreviewInvoiceCustomerTaxIDType = "ve_rif"
	QuotePreviewInvoiceCustomerTaxIDTypeVNTIN    QuotePreviewInvoiceCustomerTaxIDType = "vn_tin"
	QuotePreviewInvoiceCustomerTaxIDTypeZAVAT    QuotePreviewInvoiceCustomerTaxIDType = "za_vat"
)

// Type of the account referenced.
type QuotePreviewInvoiceIssuerType string

// List of values that QuotePreviewInvoiceIssuerType can take
const (
	QuotePreviewInvoiceIssuerTypeAccount QuotePreviewInvoiceIssuerType = "account"
	QuotePreviewInvoiceIssuerTypeSelf    QuotePreviewInvoiceIssuerType = "self"
)

// Transaction type of the mandate.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Bank account verification method.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodInstant       QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// Preferred language of the Bancontact authorization page that the customer is redirected to.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage string

// List of values that QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage can take
const (
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguageDE QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage = "de"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguageEN QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage = "en"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguageFR QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage = "fr"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguageNL QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage = "nl"
)

// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecureAny       QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure = "any"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecureAutomatic QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
)

// The desired country code of the bank account information. Permitted values include: `BE`, `DE`, `ES`, `FR`, `IE`, or `NL`.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry string

// List of values that QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry can take
const (
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryBE QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "BE"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryDE QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "DE"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryES QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "ES"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryFR QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "FR"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryIE QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "IE"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryNL QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "NL"
)

// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType string

// List of values that QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType can take
const (
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceFundingTypeBankTransfer QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType = "bank_transfer"
)

// The list of permissions to request. The `payment_method` permission must be included.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission string

// List of values that QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission can take
const (
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionBalances      QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "balances"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionOwnership     QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "ownership"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionPaymentMethod QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "payment_method"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionTransactions  QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "transactions"
)

// Data features requested to be retrieved upon account creation.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch string

// List of values that QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch can take
const (
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchBalances         QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "balances"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchInferredBalances QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "inferred_balances"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchOwnership        QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "ownership"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchTransactions     QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "transactions"
)

// Bank account verification method.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod string

// List of values that QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod can take
const (
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethodAutomatic     QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "automatic"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethodInstant       QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "instant"
	QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethodMicrodeposits QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "microdeposits"
)

// The list of payment method types (e.g. card) to provide to the invoice's PaymentIntent. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice's default payment method, the subscription's default payment method, the customer's default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).
type QuotePreviewInvoicePaymentSettingsPaymentMethodType string

// List of values that QuotePreviewInvoicePaymentSettingsPaymentMethodType can take
const (
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeACHCreditTransfer  QuotePreviewInvoicePaymentSettingsPaymentMethodType = "ach_credit_transfer"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeACHDebit           QuotePreviewInvoicePaymentSettingsPaymentMethodType = "ach_debit"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeACSSDebit          QuotePreviewInvoicePaymentSettingsPaymentMethodType = "acss_debit"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeAUBECSDebit        QuotePreviewInvoicePaymentSettingsPaymentMethodType = "au_becs_debit"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeBACSDebit          QuotePreviewInvoicePaymentSettingsPaymentMethodType = "bacs_debit"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeBancontact         QuotePreviewInvoicePaymentSettingsPaymentMethodType = "bancontact"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeBoleto             QuotePreviewInvoicePaymentSettingsPaymentMethodType = "boleto"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeCard               QuotePreviewInvoicePaymentSettingsPaymentMethodType = "card"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeCashApp            QuotePreviewInvoicePaymentSettingsPaymentMethodType = "cashapp"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeCustomerBalance    QuotePreviewInvoicePaymentSettingsPaymentMethodType = "customer_balance"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeFPX                QuotePreviewInvoicePaymentSettingsPaymentMethodType = "fpx"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeGiropay            QuotePreviewInvoicePaymentSettingsPaymentMethodType = "giropay"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeGrabpay            QuotePreviewInvoicePaymentSettingsPaymentMethodType = "grabpay"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeIDEAL              QuotePreviewInvoicePaymentSettingsPaymentMethodType = "ideal"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeKonbini            QuotePreviewInvoicePaymentSettingsPaymentMethodType = "konbini"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeLink               QuotePreviewInvoicePaymentSettingsPaymentMethodType = "link"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypePayNow             QuotePreviewInvoicePaymentSettingsPaymentMethodType = "paynow"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypePaypal             QuotePreviewInvoicePaymentSettingsPaymentMethodType = "paypal"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypePromptPay          QuotePreviewInvoicePaymentSettingsPaymentMethodType = "promptpay"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeSEPACreditTransfer QuotePreviewInvoicePaymentSettingsPaymentMethodType = "sepa_credit_transfer"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeSEPADebit          QuotePreviewInvoicePaymentSettingsPaymentMethodType = "sepa_debit"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeSofort             QuotePreviewInvoicePaymentSettingsPaymentMethodType = "sofort"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeUSBankAccount      QuotePreviewInvoicePaymentSettingsPaymentMethodType = "us_bank_account"
	QuotePreviewInvoicePaymentSettingsPaymentMethodTypeWeChatPay          QuotePreviewInvoicePaymentSettingsPaymentMethodType = "wechat_pay"
)

// Page size of invoice pdf. Options include a4, letter, and auto. If set to auto, page size will be switched to a4 or letter based on customer locale.
type QuotePreviewInvoiceRenderingPDFPageSize string

// List of values that QuotePreviewInvoiceRenderingPDFPageSize can take
const (
	QuotePreviewInvoiceRenderingPDFPageSizeA4     QuotePreviewInvoiceRenderingPDFPageSize = "a4"
	QuotePreviewInvoiceRenderingPDFPageSizeAuto   QuotePreviewInvoiceRenderingPDFPageSize = "auto"
	QuotePreviewInvoiceRenderingPDFPageSizeLetter QuotePreviewInvoiceRenderingPDFPageSize = "letter"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type QuotePreviewInvoiceShippingCostTaxTaxabilityReason string

// List of values that QuotePreviewInvoiceShippingCostTaxTaxabilityReason can take
const (
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonCustomerExempt       QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "customer_exempt"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonNotCollecting        QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "not_collecting"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonNotSubjectToTax      QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "not_subject_to_tax"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonNotSupported         QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "not_supported"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonPortionProductExempt QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "portion_product_exempt"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonPortionReducedRated  QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "portion_reduced_rated"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonPortionStandardRated QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "portion_standard_rated"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonProductExempt        QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "product_exempt"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonProductExemptHoliday QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "product_exempt_holiday"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonProportionallyRated  QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "proportionally_rated"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonReducedRated         QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "reduced_rated"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonReverseCharge        QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "reverse_charge"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonStandardRated        QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "standard_rated"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonTaxableBasisReduced  QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "taxable_basis_reduced"
	QuotePreviewInvoiceShippingCostTaxTaxabilityReasonZeroRated            QuotePreviewInvoiceShippingCostTaxTaxabilityReason = "zero_rated"
)

// The status of the invoice, one of `draft`, `open`, `paid`, `uncollectible`, or `void`. [Learn more](https://stripe.com/docs/billing/invoices/workflow#workflow-overview)
type QuotePreviewInvoiceStatus string

// List of values that QuotePreviewInvoiceStatus can take
const (
	QuotePreviewInvoiceStatusDraft         QuotePreviewInvoiceStatus = "draft"
	QuotePreviewInvoiceStatusOpen          QuotePreviewInvoiceStatus = "open"
	QuotePreviewInvoiceStatusPaid          QuotePreviewInvoiceStatus = "paid"
	QuotePreviewInvoiceStatusUncollectible QuotePreviewInvoiceStatus = "uncollectible"
	QuotePreviewInvoiceStatusVoid          QuotePreviewInvoiceStatus = "void"
)

// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
type QuotePreviewInvoiceSubscriptionDetailsPauseCollectionBehavior string

// List of values that QuotePreviewInvoiceSubscriptionDetailsPauseCollectionBehavior can take
const (
	QuotePreviewInvoiceSubscriptionDetailsPauseCollectionBehaviorKeepAsDraft       QuotePreviewInvoiceSubscriptionDetailsPauseCollectionBehavior = "keep_as_draft"
	QuotePreviewInvoiceSubscriptionDetailsPauseCollectionBehaviorMarkUncollectible QuotePreviewInvoiceSubscriptionDetailsPauseCollectionBehavior = "mark_uncollectible"
	QuotePreviewInvoiceSubscriptionDetailsPauseCollectionBehaviorVoid              QuotePreviewInvoiceSubscriptionDetailsPauseCollectionBehavior = "void"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type QuotePreviewInvoiceTotalTaxAmountTaxabilityReason string

// List of values that QuotePreviewInvoiceTotalTaxAmountTaxabilityReason can take
const (
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonCustomerExempt       QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "customer_exempt"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonNotCollecting        QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "not_collecting"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonNotSubjectToTax      QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "not_subject_to_tax"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonNotSupported         QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "not_supported"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonPortionProductExempt QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "portion_product_exempt"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonPortionReducedRated  QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "portion_reduced_rated"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonPortionStandardRated QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "portion_standard_rated"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonProductExempt        QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "product_exempt"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonProductExemptHoliday QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "product_exempt_holiday"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonProportionallyRated  QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "proportionally_rated"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonReducedRated         QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "reduced_rated"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonReverseCharge        QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "reverse_charge"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonStandardRated        QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "standard_rated"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonTaxableBasisReduced  QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "taxable_basis_reduced"
	QuotePreviewInvoiceTotalTaxAmountTaxabilityReasonZeroRated            QuotePreviewInvoiceTotalTaxAmountTaxabilityReason = "zero_rated"
)

// Preview the invoices that would be generated by accepting the quote.
type QuotePreviewInvoiceListParams struct {
	ListParams `form:"*"`
	Quote      *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuotePreviewInvoiceListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type QuotePreviewInvoiceAppliesTo struct {
	// A custom string that identifies a new subscription schedule being created upon quote acceptance. All quote lines with the same `new_reference` field will be applied to the creation of a new subscription schedule.
	NewReference string `json:"new_reference"`
	// The ID of the schedule the line applies to.
	SubscriptionSchedule string `json:"subscription_schedule"`
	// Describes whether the quote line is affecting a new schedule or an existing schedule.
	Type QuotePreviewInvoiceAppliesToType `json:"type"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type QuotePreviewInvoiceAutomaticTaxLiability struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewInvoiceAutomaticTaxLiabilityType `json:"type"`
}
type QuotePreviewInvoiceAutomaticTax struct {
	// Whether Stripe automatically computes tax on this invoice. Note that incompatible invoice items (invoice items with manually specified [tax rates](https://stripe.com/docs/api/tax_rates), negative amounts, or `tax_behavior=unspecified`) cannot be added to automatic tax invoices.
	Enabled bool `json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *QuotePreviewInvoiceAutomaticTaxLiability `json:"liability"`
	// The status of the most recent automated tax calculation for this invoice.
	Status QuotePreviewInvoiceAutomaticTaxStatus `json:"status"`
}

// Custom fields displayed on the invoice.
type QuotePreviewInvoiceCustomField struct {
	// The name of the custom field.
	Name string `json:"name"`
	// The value of the custom field.
	Value string `json:"value"`
}

// The customer's tax IDs. Until the invoice is finalized, this field will contain the same tax IDs as `customer.tax_ids`. Once the invoice is finalized, this field will no longer be updated.
type QuotePreviewInvoiceCustomerTaxID struct {
	// The type of the tax ID, one of `ad_nrt`, `ar_cuit`, `eu_vat`, `bo_tin`, `br_cnpj`, `br_cpf`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `eu_oss_vat`, `pe_ruc`, `ro_tin`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, `vn_tin`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, or `unknown`
	Type QuotePreviewInvoiceCustomerTaxIDType `json:"type"`
	// The value of the tax ID.
	Value string `json:"value"`
}

// Details of the invoice that was cloned. See the [revision documentation](https://stripe.com/docs/invoicing/invoice-revisions) for more details.
type QuotePreviewInvoiceFromInvoice struct {
	// The relation between this invoice and the cloned invoice
	Action string `json:"action"`
	// The invoice that was cloned.
	Invoice *Invoice `json:"invoice"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type QuotePreviewInvoiceIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewInvoiceIssuerType `json:"type"`
}
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions struct {
	// Transaction type of the mandate.
	TransactionType QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}

// If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebit struct {
	MandateOptions *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options"`
	// Bank account verification method.
	VerificationMethod QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the invoice's PaymentIntent.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontact struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage `json:"preferred_language"`
}
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCardInstallments struct {
	// Whether Installments are enabled for this Invoice.
	Enabled bool `json:"enabled"`
}

// If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the invoice's PaymentIntent.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCard struct {
	Installments *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCardInstallments `json:"installments"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer struct {
	// The desired country code of the bank account information. Permitted values include: `BE`, `DE`, `ES`, `FR`, `IE`, or `NL`.
	Country QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry `json:"country"`
}
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransfer struct {
	EUBankTransfer *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer `json:"eu_bank_transfer"`
	// The bank transfer type that can be used for funding. Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.
	Type string `json:"type"`
}

// If paying by `customer_balance`, this sub-hash contains details about the Bank transfer payment method options to pass to the invoice's PaymentIntent.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalance struct {
	BankTransfer *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType `json:"funding_type"`
}

// If paying by `konbini`, this sub-hash contains details about the Konbini payment method options to pass to the invoice's PaymentIntent.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsKonbini struct{}
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	// The list of permissions to request. The `payment_method` permission must be included.
	Permissions []QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
	// Data features requested to be retrieved upon account creation.
	Prefetch []QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch `json:"prefetch"`
}

// If paying by `us_bank_account`, this sub-hash contains details about the ACH direct debit payment method options to pass to the invoice's PaymentIntent.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccount struct {
	FinancialConnections *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	// Bank account verification method.
	VerificationMethod QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment-method-specific configuration to provide to the invoice's PaymentIntent.
type QuotePreviewInvoicePaymentSettingsPaymentMethodOptions struct {
	// If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
	ACSSDebit *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsACSSDebit `json:"acss_debit"`
	// If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the invoice's PaymentIntent.
	Bancontact *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBancontact `json:"bancontact"`
	// If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the invoice's PaymentIntent.
	Card *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCard `json:"card"`
	// If paying by `customer_balance`, this sub-hash contains details about the Bank transfer payment method options to pass to the invoice's PaymentIntent.
	CustomerBalance *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCustomerBalance `json:"customer_balance"`
	// If paying by `konbini`, this sub-hash contains details about the Konbini payment method options to pass to the invoice's PaymentIntent.
	Konbini *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsKonbini `json:"konbini"`
	// If paying by `us_bank_account`, this sub-hash contains details about the ACH direct debit payment method options to pass to the invoice's PaymentIntent.
	USBankAccount *QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsUSBankAccount `json:"us_bank_account"`
}
type QuotePreviewInvoicePaymentSettings struct {
	// ID of the mandate to be used for this invoice. It must correspond to the payment method used to pay the invoice, including the invoice's default_payment_method or default_source, if set.
	DefaultMandate string `json:"default_mandate"`
	// Payment-method-specific configuration to provide to the invoice's PaymentIntent.
	PaymentMethodOptions *QuotePreviewInvoicePaymentSettingsPaymentMethodOptions `json:"payment_method_options"`
	// The list of payment method types (e.g. card) to provide to the invoice's PaymentIntent. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice's default payment method, the subscription's default payment method, the customer's default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).
	PaymentMethodTypes []QuotePreviewInvoicePaymentSettingsPaymentMethodType `json:"payment_method_types"`
}

// Invoice pdf rendering options
type QuotePreviewInvoiceRenderingPDF struct {
	// Page size of invoice pdf. Options include a4, letter, and auto. If set to auto, page size will be switched to a4 or letter based on customer locale.
	PageSize QuotePreviewInvoiceRenderingPDFPageSize `json:"page_size"`
}

// The rendering-related settings that control how the invoice is displayed on customer-facing surfaces such as PDF and Hosted Invoice Page.
type QuotePreviewInvoiceRendering struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.
	AmountTaxDisplay string `json:"amount_tax_display"`
	// Invoice pdf rendering options
	PDF *QuotePreviewInvoiceRenderingPDF `json:"pdf"`
}

// This is a legacy field that will be removed soon. For details about `rendering_options`, refer to `rendering` instead. Options for invoice PDF rendering.
type QuotePreviewInvoiceRenderingOptions struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.
	AmountTaxDisplay string `json:"amount_tax_display"`
}

// The taxes applied to the shipping rate.
type QuotePreviewInvoiceShippingCostTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.
	//
	// Related guide: [Tax rates](https://stripe.com/docs/billing/taxes/tax-rates)
	Rate *TaxRate `json:"rate"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason QuotePreviewInvoiceShippingCostTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
}

// The details of the cost of shipping, including the ShippingRate applied on the invoice.
type QuotePreviewInvoiceShippingCost struct {
	// Total shipping cost before any taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total tax amount applied due to shipping costs. If no tax was applied, defaults to 0.
	AmountTax int64 `json:"amount_tax"`
	// Total shipping cost after taxes are applied.
	AmountTotal int64 `json:"amount_total"`
	// The ID of the ShippingRate for this invoice.
	ShippingRate *ShippingRate `json:"shipping_rate"`
	// The taxes applied to the shipping rate.
	Taxes []*QuotePreviewInvoiceShippingCostTax `json:"taxes"`
}
type QuotePreviewInvoiceStatusTransitions struct {
	// The time that the invoice draft was finalized.
	FinalizedAt int64 `json:"finalized_at"`
	// The time that the invoice was marked uncollectible.
	MarkedUncollectibleAt int64 `json:"marked_uncollectible_at"`
	// The time that the invoice was paid.
	PaidAt int64 `json:"paid_at"`
	// The time that the invoice was voided.
	VoidedAt int64 `json:"voided_at"`
}

// If specified, payment collection for this subscription will be paused.
type QuotePreviewInvoiceSubscriptionDetailsPauseCollection struct {
	// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
	Behavior QuotePreviewInvoiceSubscriptionDetailsPauseCollectionBehavior `json:"behavior"`
	// The time after which the subscription will resume collecting payments.
	ResumesAt int64 `json:"resumes_at"`
}

// Details about the subscription that created this invoice.
type QuotePreviewInvoiceSubscriptionDetails struct {
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that will reflect the metadata of the subscription at the time of invoice creation. *Note: This attribute is populated only for invoices created on or after June 29, 2023.*
	Metadata map[string]string `json:"metadata"`
	// If specified, payment collection for this subscription will be paused.
	PauseCollection *QuotePreviewInvoiceSubscriptionDetailsPauseCollection `json:"pause_collection"`
}

// Indicates which line items triggered a threshold invoice.
type QuotePreviewInvoiceThresholdReasonItemReason struct {
	// The IDs of the line items that triggered the threshold invoice.
	LineItemIDs []string `json:"line_item_ids"`
	// The quantity threshold boundary that applied to the given line item.
	UsageGTE int64 `json:"usage_gte"`
}
type QuotePreviewInvoiceThresholdReason struct {
	// The total invoice amount threshold boundary if it triggered the threshold invoice.
	AmountGTE int64 `json:"amount_gte"`
	// Indicates which line items triggered a threshold invoice.
	ItemReasons []*QuotePreviewInvoiceThresholdReasonItemReason `json:"item_reasons"`
}

// The aggregate amounts calculated per discount across all line items.
type QuotePreviewInvoiceTotalDiscountAmount struct {
	// The amount, in cents (or local equivalent), of the discount.
	Amount int64 `json:"amount"`
	// The discount that was applied to get this discount amount.
	Discount *Discount `json:"discount"`
}

// The aggregate amounts calculated per tax rate for all line items.
type QuotePreviewInvoiceTotalTaxAmount struct {
	// The amount, in cents (or local equivalent), of the tax.
	Amount int64 `json:"amount"`
	// Whether this tax amount is inclusive or exclusive.
	Inclusive bool `json:"inclusive"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason QuotePreviewInvoiceTotalTaxAmountTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
	// The tax rate that was applied to get this tax amount.
	TaxRate *TaxRate `json:"tax_rate"`
}

// The account (if any) the payment will be attributed to for tax reporting, and where funds from the payment will be transferred to for the invoice.
type QuotePreviewInvoiceTransferData struct {
	// The amount in cents (or local equivalent) that will be transferred to the destination account when the invoice is paid. By default, the entire amount is transferred to the destination.
	Amount int64 `json:"amount"`
	// The account where funds from the payment will be transferred to upon payment success.
	Destination *Account `json:"destination"`
}

// Invoices are statements of amounts owed by a customer, and are either
// generated one-off, or generated periodically from a subscription.
//
// They contain [invoice items](https://stripe.com/docs/api#invoiceitems), and proration adjustments
// that may be caused by subscription upgrades/downgrades (if necessary).
//
// If your invoice is configured to be billed through automatic charges,
// Stripe automatically finalizes your invoice and attempts payment. Note
// that finalizing the invoice,
// [when automatic](https://stripe.com/docs/invoicing/integration/automatic-advancement-collection), does
// not happen immediately as the invoice is created. Stripe waits
// until one hour after the last webhook was successfully sent (or the last
// webhook timed out after failing). If you (and the platforms you may have
// connected to) have no webhooks configured, Stripe waits one hour after
// creation to finalize the invoice.
//
// If your invoice is configured to be billed by sending an email, then based on your
// [email settings](https://dashboard.stripe.com/account/billing/automatic),
// Stripe will email the invoice to your customer and await payment. These
// emails can contain a link to a hosted page to pay the invoice.
//
// Stripe applies any customer credit on the account before determining the
// amount due for the invoice (i.e., the amount that will be actually
// charged). If the amount due for the invoice is less than Stripe's [minimum allowed charge
// per currency](https://stripe.com/docs/currencies#minimum-and-maximum-charge-amounts), the
// invoice is automatically marked paid, and we add the amount due to the
// customer's credit balance which is applied to the next invoice.
//
// More details on the customer's credit balance are
// [here](https://stripe.com/docs/billing/customer/balance).
//
// Related guide: [Send invoices to customers](https://stripe.com/docs/billing/invoices/sending)
type QuotePreviewInvoice struct {
	// The country of the business associated with this invoice, most often the business creating the invoice.
	AccountCountry string `json:"account_country"`
	// The public name of the business associated with this invoice, most often the business creating the invoice.
	AccountName string `json:"account_name"`
	// The account tax IDs associated with the invoice. Only editable when the invoice is a draft.
	AccountTaxIDs []*TaxID `json:"account_tax_ids"`
	// Final amount due at this time for this invoice. If the invoice's total is smaller than the minimum charge amount, for example, or if there is account credit that can be applied to the invoice, the `amount_due` may be 0. If there is a positive `starting_balance` for the invoice (the customer owes money), the `amount_due` will also take that into account. The charge that gets generated for the invoice will be for the amount specified in `amount_due`.
	AmountDue int64 `json:"amount_due"`
	// The amount, in cents (or local equivalent), that was paid.
	AmountPaid int64 `json:"amount_paid"`
	// The difference between amount_due and amount_paid, in cents (or local equivalent).
	AmountRemaining int64 `json:"amount_remaining"`
	// This is the sum of all the shipping amounts.
	AmountShipping int64 `json:"amount_shipping"`
	// ID of the Connect Application that created the invoice.
	Application *Application `json:"application"`
	// The fee in cents (or local equivalent) that will be applied to the invoice and transferred to the application owner's Stripe account when the invoice is paid.
	ApplicationFeeAmount int64                         `json:"application_fee_amount"`
	AppliesTo            *QuotePreviewInvoiceAppliesTo `json:"applies_to"`
	// Number of payment attempts made for this invoice, from the perspective of the payment retry schedule. Any payment attempt counts as the first attempt, and subsequently only automatic retries increment the attempt count. In other words, manual payment attempts after the first attempt do not affect the retry schedule.
	AttemptCount int64 `json:"attempt_count"`
	// Whether an attempt has been made to pay the invoice. An invoice is not attempted until 1 hour after the `invoice.created` webhook, for example, so you might not want to display that invoice as unpaid to your users.
	Attempted    bool                             `json:"attempted"`
	AutomaticTax *QuotePreviewInvoiceAutomaticTax `json:"automatic_tax"`
	// Indicates the reason why the invoice was created.
	//
	// * `manual`: Unrelated to a subscription, for example, created via the invoice editor.
	// * `subscription`: No longer in use. Applies to subscriptions from before May 2018 where no distinction was made between updates, cycles, and thresholds.
	// * `subscription_create`: A new subscription was created.
	// * `subscription_cycle`: A subscription advanced into a new period.
	// * `subscription_threshold`: A subscription reached a billing threshold.
	// * `subscription_update`: A subscription was updated.
	// * `upcoming`: Reserved for simulated invoices, per the upcoming invoice endpoint.
	BillingReason QuotePreviewInvoiceBillingReason `json:"billing_reason"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions.
	CollectionMethod QuotePreviewInvoiceCollectionMethod `json:"collection_method"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The customer's address. Until the invoice is finalized, this field will equal `customer.address`. Once the invoice is finalized, this field will no longer be updated.
	CustomerAddress *Address `json:"customer_address"`
	// The customer's email. Until the invoice is finalized, this field will equal `customer.email`. Once the invoice is finalized, this field will no longer be updated.
	CustomerEmail string `json:"customer_email"`
	// The customer's name. Until the invoice is finalized, this field will equal `customer.name`. Once the invoice is finalized, this field will no longer be updated.
	CustomerName string `json:"customer_name"`
	// The customer's phone number. Until the invoice is finalized, this field will equal `customer.phone`. Once the invoice is finalized, this field will no longer be updated.
	CustomerPhone string `json:"customer_phone"`
	// The customer's shipping information. Until the invoice is finalized, this field will equal `customer.shipping`. Once the invoice is finalized, this field will no longer be updated.
	CustomerShipping *ShippingDetails `json:"customer_shipping"`
	// The customer's tax exempt status. Until the invoice is finalized, this field will equal `customer.tax_exempt`. Once the invoice is finalized, this field will no longer be updated.
	CustomerTaxExempt QuotePreviewInvoiceCustomerTaxExempt `json:"customer_tax_exempt"`
	// The customer's tax IDs. Until the invoice is finalized, this field will contain the same tax IDs as `customer.tax_ids`. Once the invoice is finalized, this field will no longer be updated.
	CustomerTaxIDs []*QuotePreviewInvoiceCustomerTaxID `json:"customer_tax_ids"`
	// Custom fields displayed on the invoice.
	CustomFields []*QuotePreviewInvoiceCustomField `json:"custom_fields"`
	// ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription's default payment method, if any, or to the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// ID of the default payment source for the invoice. It must belong to the customer associated with the invoice and be in a chargeable state. If not set, defaults to the subscription's default source, if any, or to the customer's default source.
	DefaultSource *PaymentSource `json:"default_source"`
	// The tax rates applied to this invoice, if any.
	DefaultTaxRates []*TaxRate `json:"default_tax_rates"`
	// An arbitrary string attached to the object. Often useful for displaying to users. Referenced as 'memo' in the Dashboard.
	Description string `json:"description"`
	// Describes the current discount applied to this invoice, if there is one. Not populated if there are multiple discounts.
	Discount *Discount `json:"discount"`
	// The discounts applied to the invoice. Line item discounts are applied before invoice discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*Discount `json:"discounts"`
	// The date on which payment for this invoice is due. This value will be `null` for invoices where `collection_method=charge_automatically`.
	DueDate int64 `json:"due_date"`
	// The date when this invoice is in effect. Same as `finalized_at` unless overwritten. When defined, this value replaces the system-generated 'Date of issue' printed on the invoice PDF and receipt.
	EffectiveAt int64 `json:"effective_at"`
	// Ending customer balance after the invoice is finalized. Invoices are finalized approximately an hour after successful webhook delivery or when payment collection is attempted for the invoice. If the invoice has not been finalized yet, this will be null.
	EndingBalance int64 `json:"ending_balance"`
	// Footer displayed on the invoice.
	Footer string `json:"footer"`
	// Details of the invoice that was cloned. See the [revision documentation](https://stripe.com/docs/invoicing/invoice-revisions) for more details.
	FromInvoice *QuotePreviewInvoiceFromInvoice `json:"from_invoice"`
	// Unique identifier for the object. This property is always present unless the invoice is an upcoming invoice. See [Retrieve an upcoming invoice](https://stripe.com/docs/api/invoices/upcoming) for more details.
	ID string `json:"id"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *QuotePreviewInvoiceIssuer `json:"issuer"`
	// The error encountered during the previous attempt to finalize the invoice. This field is cleared when the invoice is successfully finalized.
	LastFinalizationError *Error `json:"last_finalization_error"`
	// The ID of the most recent non-draft revision of this invoice
	LatestRevision *Invoice `json:"latest_revision"`
	// The individual line items that make up the invoice. `lines` is sorted as follows: (1) pending invoice items (including prorations) in reverse chronological order, (2) subscription items in reverse chronological order, and (3) invoice items added after invoice creation in chronological order.
	Lines *InvoiceLineItemList `json:"lines"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The time at which payment will next be attempted. This value will be `null` for invoices where `collection_method=send_invoice`.
	NextPaymentAttempt int64 `json:"next_payment_attempt"`
	// A unique, identifying string that appears on emails sent to the customer for this invoice. This starts with the customer's unique invoice_prefix if it is specified.
	Number string `json:"number"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account (if any) for which the funds of the invoice payment are intended. If set, the invoice will be presented with the branding and support information of the specified account. See the [Invoices with Connect](https://stripe.com/docs/billing/invoices/connect) documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// Whether payment was successfully collected for this invoice. An invoice can be paid (most commonly) with a charge or with credit from the customer's account balance.
	Paid bool `json:"paid"`
	// Returns true if the invoice was manually marked paid, returns false if the invoice hasn't been paid yet or was paid on Stripe.
	PaidOutOfBand bool `json:"paid_out_of_band"`
	// The PaymentIntent associated with this invoice. The PaymentIntent is generated when the invoice is finalized, and can then be used to pay the invoice. Note that voiding an invoice will cancel the PaymentIntent.
	PaymentIntent   *PaymentIntent                      `json:"payment_intent"`
	PaymentSettings *QuotePreviewInvoicePaymentSettings `json:"payment_settings"`
	// End of the usage period during which invoice items were added to this invoice.
	PeriodEnd int64 `json:"period_end"`
	// Start of the usage period during which invoice items were added to this invoice.
	PeriodStart int64 `json:"period_start"`
	// Total amount of all post-payment credit notes issued for this invoice.
	PostPaymentCreditNotesAmount int64 `json:"post_payment_credit_notes_amount"`
	// Total amount of all pre-payment credit notes issued for this invoice.
	PrePaymentCreditNotesAmount int64 `json:"pre_payment_credit_notes_amount"`
	// The quote this invoice was generated from.
	Quote *Quote `json:"quote"`
	// This is the transaction number that appears on email receipts sent for this invoice.
	ReceiptNumber string `json:"receipt_number"`
	// The rendering-related settings that control how the invoice is displayed on customer-facing surfaces such as PDF and Hosted Invoice Page.
	Rendering *QuotePreviewInvoiceRendering `json:"rendering"`
	// This is a legacy field that will be removed soon. For details about `rendering_options`, refer to `rendering` instead. Options for invoice PDF rendering.
	RenderingOptions *QuotePreviewInvoiceRenderingOptions `json:"rendering_options"`
	// The details of the cost of shipping, including the ShippingRate applied on the invoice.
	ShippingCost *QuotePreviewInvoiceShippingCost `json:"shipping_cost"`
	// Shipping details for the invoice. The Invoice PDF will use the `shipping_details` value if it is set, otherwise the PDF will render the shipping address from the customer.
	ShippingDetails *ShippingDetails `json:"shipping_details"`
	// Starting customer balance before the invoice is finalized. If the invoice has not been finalized yet, this will be the current customer balance. For revision invoices, this also includes any customer balance that was applied to the original invoice.
	StartingBalance int64 `json:"starting_balance"`
	// Extra information about an invoice for the customer's credit card statement.
	StatementDescriptor string `json:"statement_descriptor"`
	// The status of the invoice, one of `draft`, `open`, `paid`, `uncollectible`, or `void`. [Learn more](https://stripe.com/docs/billing/invoices/workflow#workflow-overview)
	Status            QuotePreviewInvoiceStatus             `json:"status"`
	StatusTransitions *QuotePreviewInvoiceStatusTransitions `json:"status_transitions"`
	// The subscription that this invoice was prepared for, if any.
	Subscription *Subscription `json:"subscription"`
	// Details about the subscription that created this invoice.
	SubscriptionDetails *QuotePreviewInvoiceSubscriptionDetails `json:"subscription_details"`
	// Only set for upcoming invoices that preview prorations. The time used to calculate prorations.
	SubscriptionProrationDate int64 `json:"subscription_proration_date"`
	// Total of all subscriptions, invoice items, and prorations on the invoice before any invoice level discount or exclusive tax is applied. Item discounts are already incorporated
	Subtotal int64 `json:"subtotal"`
	// The integer amount in cents (or local equivalent) representing the subtotal of the invoice before any invoice level discount or tax is applied. Item discounts are already incorporated
	SubtotalExcludingTax int64 `json:"subtotal_excluding_tax"`
	// The amount of tax on this invoice. This is the sum of all the tax amounts on this invoice.
	Tax int64 `json:"tax"`
	// ID of the test clock this invoice belongs to.
	TestClock       *TestHelpersTestClock               `json:"test_clock"`
	ThresholdReason *QuotePreviewInvoiceThresholdReason `json:"threshold_reason"`
	// Total after discounts and taxes.
	Total int64 `json:"total"`
	// The aggregate amounts calculated per discount across all line items.
	TotalDiscountAmounts []*QuotePreviewInvoiceTotalDiscountAmount `json:"total_discount_amounts"`
	// The integer amount in cents (or local equivalent) representing the total amount of the invoice including all discounts but excluding all tax.
	TotalExcludingTax int64 `json:"total_excluding_tax"`
	// The aggregate amounts calculated per tax rate for all line items.
	TotalTaxAmounts []*QuotePreviewInvoiceTotalTaxAmount `json:"total_tax_amounts"`
	// The account (if any) the payment will be attributed to for tax reporting, and where funds from the payment will be transferred to for the invoice.
	TransferData *QuotePreviewInvoiceTransferData `json:"transfer_data"`
	// Invoices are automatically paid or sent 1 hour after webhooks are delivered, or until all webhook delivery attempts have [been exhausted](https://stripe.com/docs/billing/webhooks#understand). This field tracks the time when webhooks for this invoice were successfully delivered. If the invoice had no webhooks to deliver, this will be set while the invoice is being created.
	WebhooksDeliveredAt int64 `json:"webhooks_delivered_at"`
}

// QuotePreviewInvoiceList is a list of QuotePreviewInvoices as retrieved from a list endpoint.
type QuotePreviewInvoiceList struct {
	APIResource
	ListMeta
	Data []*QuotePreviewInvoice `json:"data"`
}
