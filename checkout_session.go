//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The status of the most recent automated tax calculation for this session.
type CheckoutSessionAutomaticTaxStatus string

// List of values that CheckoutSessionAutomaticTaxStatus can take
const (
	CheckoutSessionAutomaticTaxStatusComplete               CheckoutSessionAutomaticTaxStatus = "complete"
	CheckoutSessionAutomaticTaxStatusFailed                 CheckoutSessionAutomaticTaxStatus = "failed"
	CheckoutSessionAutomaticTaxStatusRequiresLocationInputs CheckoutSessionAutomaticTaxStatus = "requires_location_inputs"
)

// Describes whether Checkout should collect the customer's billing address.
type CheckoutSessionBillingAddressCollection string

// List of values that CheckoutSessionBillingAddressCollection can take
const (
	CheckoutSessionBillingAddressCollectionAuto     CheckoutSessionBillingAddressCollection = "auto"
	CheckoutSessionBillingAddressCollectionRequired CheckoutSessionBillingAddressCollection = "required"
)

// If `opt_in`, the customer consents to receiving promotional communications
// from the merchant about this Checkout Session.
type CheckoutSessionConsentPromotions string

// List of values that CheckoutSessionConsentPromotions can take
const (
	CheckoutSessionConsentPromotionsOptIn  CheckoutSessionConsentPromotions = "opt_in"
	CheckoutSessionConsentPromotionsOptOut CheckoutSessionConsentPromotions = "opt_out"
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

// The type of the tax ID, one of `eu_vat`, `br_cnpj`, `br_cpf`, `eu_oss_vat`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, or `unknown`
type CheckoutSessionCustomerDetailsTaxIDsType string

// List of values that CheckoutSessionCustomerDetailsTaxIDsType can take
const (
	CheckoutSessionCustomerDetailsTaxIDsTypeAETRN    CheckoutSessionCustomerDetailsTaxIDsType = "ae_trn"
	CheckoutSessionCustomerDetailsTaxIDsTypeAUABN    CheckoutSessionCustomerDetailsTaxIDsType = "au_abn"
	CheckoutSessionCustomerDetailsTaxIDsTypeAUARN    CheckoutSessionCustomerDetailsTaxIDsType = "au_arn"
	CheckoutSessionCustomerDetailsTaxIDsTypeBGUIC    CheckoutSessionCustomerDetailsTaxIDsType = "bg_uic"
	CheckoutSessionCustomerDetailsTaxIDsTypeBRCNPJ   CheckoutSessionCustomerDetailsTaxIDsType = "br_cnpj"
	CheckoutSessionCustomerDetailsTaxIDsTypeBRCPF    CheckoutSessionCustomerDetailsTaxIDsType = "br_cpf"
	CheckoutSessionCustomerDetailsTaxIDsTypeCABN     CheckoutSessionCustomerDetailsTaxIDsType = "ca_bn"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAGSTHST CheckoutSessionCustomerDetailsTaxIDsType = "ca_gst_hst"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTBC  CheckoutSessionCustomerDetailsTaxIDsType = "ca_pst_bc"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTMB  CheckoutSessionCustomerDetailsTaxIDsType = "ca_pst_mb"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTSK  CheckoutSessionCustomerDetailsTaxIDsType = "ca_pst_sk"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAQST    CheckoutSessionCustomerDetailsTaxIDsType = "ca_qst"
	CheckoutSessionCustomerDetailsTaxIDsTypeCHVAT    CheckoutSessionCustomerDetailsTaxIDsType = "ch_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeCLTIN    CheckoutSessionCustomerDetailsTaxIDsType = "cl_tin"
	CheckoutSessionCustomerDetailsTaxIDsTypeESCIF    CheckoutSessionCustomerDetailsTaxIDsType = "es_cif"
	CheckoutSessionCustomerDetailsTaxIDsTypeEUOSSVAT CheckoutSessionCustomerDetailsTaxIDsType = "eu_oss_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeEUVAT    CheckoutSessionCustomerDetailsTaxIDsType = "eu_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeGBVAT    CheckoutSessionCustomerDetailsTaxIDsType = "gb_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeGEVAT    CheckoutSessionCustomerDetailsTaxIDsType = "ge_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeHKBR     CheckoutSessionCustomerDetailsTaxIDsType = "hk_br"
	CheckoutSessionCustomerDetailsTaxIDsTypeHUTIN    CheckoutSessionCustomerDetailsTaxIDsType = "hu_tin"
	CheckoutSessionCustomerDetailsTaxIDsTypeIDNPWP   CheckoutSessionCustomerDetailsTaxIDsType = "id_npwp"
	CheckoutSessionCustomerDetailsTaxIDsTypeILVAT    CheckoutSessionCustomerDetailsTaxIDsType = "il_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeINGST    CheckoutSessionCustomerDetailsTaxIDsType = "in_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeISVAT    CheckoutSessionCustomerDetailsTaxIDsType = "is_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeJPCN     CheckoutSessionCustomerDetailsTaxIDsType = "jp_cn"
	CheckoutSessionCustomerDetailsTaxIDsTypeJPRN     CheckoutSessionCustomerDetailsTaxIDsType = "jp_rn"
	CheckoutSessionCustomerDetailsTaxIDsTypeKRBRN    CheckoutSessionCustomerDetailsTaxIDsType = "kr_brn"
	CheckoutSessionCustomerDetailsTaxIDsTypeLIUID    CheckoutSessionCustomerDetailsTaxIDsType = "li_uid"
	CheckoutSessionCustomerDetailsTaxIDsTypeMXRFC    CheckoutSessionCustomerDetailsTaxIDsType = "mx_rfc"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYFRP    CheckoutSessionCustomerDetailsTaxIDsType = "my_frp"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYITN    CheckoutSessionCustomerDetailsTaxIDsType = "my_itn"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYSST    CheckoutSessionCustomerDetailsTaxIDsType = "my_sst"
	CheckoutSessionCustomerDetailsTaxIDsTypeNOVAT    CheckoutSessionCustomerDetailsTaxIDsType = "no_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeNZGST    CheckoutSessionCustomerDetailsTaxIDsType = "nz_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeRUINN    CheckoutSessionCustomerDetailsTaxIDsType = "ru_inn"
	CheckoutSessionCustomerDetailsTaxIDsTypeRUKPP    CheckoutSessionCustomerDetailsTaxIDsType = "ru_kpp"
	CheckoutSessionCustomerDetailsTaxIDsTypeSAVAT    CheckoutSessionCustomerDetailsTaxIDsType = "sa_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeSGGST    CheckoutSessionCustomerDetailsTaxIDsType = "sg_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeSGUEN    CheckoutSessionCustomerDetailsTaxIDsType = "sg_uen"
	CheckoutSessionCustomerDetailsTaxIDsTypeSITIN    CheckoutSessionCustomerDetailsTaxIDsType = "si_tin"
	CheckoutSessionCustomerDetailsTaxIDsTypeTHVAT    CheckoutSessionCustomerDetailsTaxIDsType = "th_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeTWVAT    CheckoutSessionCustomerDetailsTaxIDsType = "tw_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeUAVAT    CheckoutSessionCustomerDetailsTaxIDsType = "ua_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeUnknown  CheckoutSessionCustomerDetailsTaxIDsType = "unknown"
	CheckoutSessionCustomerDetailsTaxIDsTypeUSEIN    CheckoutSessionCustomerDetailsTaxIDsType = "us_ein"
	CheckoutSessionCustomerDetailsTaxIDsTypeZAVAT    CheckoutSessionCustomerDetailsTaxIDsType = "za_vat"
)

// The mode of the Checkout Session.
type CheckoutSessionMode string

// List of values that CheckoutSessionMode can take
const (
	CheckoutSessionModePayment      CheckoutSessionMode = "payment"
	CheckoutSessionModeSetup        CheckoutSessionMode = "setup"
	CheckoutSessionModeSubscription CheckoutSessionMode = "subscription"
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
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
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

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsAffirmSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsAffirmSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsAffirmSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsAffirmSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsAfterpayClearpaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsAfterpayClearpaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsAfterpayClearpaySetupFutureUsageNone CheckoutSessionPaymentMethodOptionsAfterpayClearpaySetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsAlipaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsAlipaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsAlipaySetupFutureUsageNone CheckoutSessionPaymentMethodOptionsAlipaySetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsAUBECSDebitSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsAUBECSDebitSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsAUBECSDebitSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsAUBECSDebitSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage = "on_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsBancontactSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsBancontactSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsBancontactSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsBancontactSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage = "on_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsCardSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsCardSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsCardSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage = "on_session"
)

// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
//
// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType string

// List of values that CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType can take
const (
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeIban     CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "iban"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSepa     CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "sepa"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSortCode CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "sort_code"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSpei     CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "spei"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeZengin   CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "zengin"
)

// The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, or `mx_bank_transfer`.
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType string

// List of values that CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType can take
const (
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferTypeEUBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType = "eu_bank_transfer"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferTypeGBBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType = "gb_bank_transfer"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferTypeJPBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType = "jp_bank_transfer"
	CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferTypeMXBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType = "mx_bank_transfer"
)

// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
type CheckoutSessionPaymentMethodOptionsCustomerBalanceFundingType string

// List of values that CheckoutSessionPaymentMethodOptionsCustomerBalanceFundingType can take
const (
	CheckoutSessionPaymentMethodOptionsCustomerBalanceFundingTypeBankTransfer CheckoutSessionPaymentMethodOptionsCustomerBalanceFundingType = "bank_transfer"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsCustomerBalanceSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsCustomerBalanceSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsCustomerBalanceSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsCustomerBalanceSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsEPSSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsEPSSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsEPSSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsEPSSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsFPXSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsFPXSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsFPXSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsFPXSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsGiropaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsGiropaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsGiropaySetupFutureUsageNone CheckoutSessionPaymentMethodOptionsGiropaySetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsGrabpaySetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsGrabpaySetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsGrabpaySetupFutureUsageNone CheckoutSessionPaymentMethodOptionsGrabpaySetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsIdealSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsIdealSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsIdealSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsIdealSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage = "on_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsKonbiniSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsKonbiniSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsKonbiniSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsKonbiniSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsOXXOSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsOXXOSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsOXXOSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsOXXOSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsP24SetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsP24SetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsP24SetupFutureUsageNone CheckoutSessionPaymentMethodOptionsP24SetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsPayNowSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsPayNowSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsPayNowSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsPayNowSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsageNone       CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsage = "none"
	CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsageOffSession CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsage = "off_session"
	CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsageOnSession  CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsage = "on_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type CheckoutSessionPaymentMethodOptionsSofortSetupFutureUsage string

// List of values that CheckoutSessionPaymentMethodOptionsSofortSetupFutureUsage can take
const (
	CheckoutSessionPaymentMethodOptionsSofortSetupFutureUsageNone CheckoutSessionPaymentMethodOptionsSofortSetupFutureUsage = "none"
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

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
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
// specified on Checkout Sessions in `payment` mode, but not Checkout Sessions
// in `subscription` or `setup` mode.
type CheckoutSessionSubmitType string

// List of values that CheckoutSessionSubmitType can take
const (
	CheckoutSessionSubmitTypeAuto   CheckoutSessionSubmitType = "auto"
	CheckoutSessionSubmitTypeBook   CheckoutSessionSubmitType = "book"
	CheckoutSessionSubmitTypeDonate CheckoutSessionSubmitType = "donate"
	CheckoutSessionSubmitTypePay    CheckoutSessionSubmitType = "pay"
)

// Only return the Checkout Sessions for the Customer details specified.
type CheckoutSessionListCustomerDetailsParams struct {
	// Customer's email address.
	Email *string `form:"email"`
}

// Returns a list of Checkout Sessions.
type CheckoutSessionListParams struct {
	ListParams `form:"*"`
	// Only return the Checkout Sessions for the Customer specified.
	Customer *string `form:"customer"`
	// Only return the Checkout Sessions for the Customer details specified.
	CustomerDetails *CheckoutSessionListCustomerDetailsParams `form:"customer_details"`
	// Only return the Checkout Session for the PaymentIntent specified.
	PaymentIntent *string `form:"payment_intent"`
	// Only return the Checkout Session for the subscription specified.
	Subscription *string `form:"subscription"`
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

// Configure actions after a Checkout Session has expired.
type CheckoutSessionAfterExpirationParams struct {
	// Configure a Checkout Session that can be used to recover an expired session.
	Recovery *CheckoutSessionAfterExpirationRecoveryParams `form:"recovery"`
}

// Settings for automatic tax lookup for this session and resulting payments, invoices, and subscriptions.
type CheckoutSessionAutomaticTaxParams struct {
	// Set to true to enable automatic taxes.
	Enabled *bool `form:"enabled"`
}

// Configure fields for the Checkout Session to gather active consent from customers.
type CheckoutSessionConsentCollectionParams struct {
	// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
	// Session will determine whether to display an option to opt into promotional communication
	// from the merchant depending on the customer's locale. Only available to US merchants.
	Promotions *string `form:"promotions"`
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

// When set, provides configuration for this item's quantity to be adjusted by the customer during Checkout.
type CheckoutSessionLineItemAdjustableQuantityParams struct {
	// Set to true if the quantity can be adjusted to any non-negative integer. By default customers will be able to remove the line item by setting the quantity to 0.
	Enabled *bool `form:"enabled"`
	// The maximum quantity the customer can purchase for the Checkout Session. By default this value is 99. You can specify a value up to 999.
	Maximum *int64 `form:"maximum"`
	// The minimum quantity the customer must purchase for the Checkout Session. By default this value is 0.
	Minimum *int64 `form:"minimum"`
}

// Data used to generate a new product object inline. One of `product` or `product_data` is required.
type CheckoutSessionLineItemPriceDataProductDataParams struct {
	// The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description"`
	// A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
}

// The recurring components of a price such as `interval` and `interval_count`.
type CheckoutSessionLineItemPriceDataRecurringParams struct {
	AggregateUsage *string `form:"aggregate_usage"`
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).
	IntervalCount   *int64  `form:"interval_count"`
	TrialPeriodDays *int64  `form:"trial_period_days"`
	UsageType       *string `form:"usage_type"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
type CheckoutSessionLineItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the product that this price will belong to. One of `product` or `product_data` is required.
	Product *string `form:"product"`
	// Data used to generate a new product object inline. One of `product` or `product_data` is required.
	ProductData *CheckoutSessionLineItemPriceDataProductDataParams `form:"product_data"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *CheckoutSessionLineItemPriceDataRecurringParams `form:"recurring"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A non-negative integer in cents (or local equivalent) representing how much to charge. One of `unit_amount` or `unit_amount_decimal` is required.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// A list of items the customer is purchasing. Use this parameter to pass one-time or recurring [Prices](https://stripe.com/docs/api/prices).
//
// For `payment` mode, there is a maximum of 100 line items, however it is recommended to consolidate line items if there are more than a few dozen.
//
// For `subscription` mode, there is a maximum of 20 line items with recurring Prices and 20 line items with one-time Prices. Line items with one-time Prices in will be on the initial invoice only.
type CheckoutSessionLineItemParams struct {
	// When set, provides configuration for this item's quantity to be adjusted by the customer during Checkout.
	AdjustableQuantity *CheckoutSessionLineItemAdjustableQuantityParams `form:"adjustable_quantity"`
	// [Deprecated] The amount to be collected per unit of the line item. If specified, must also pass `currency` and `name`.
	Amount *int64 `form:"amount"`
	// [Deprecated] Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies). Required if `amount` is passed.
	Currency *string `form:"currency"`
	// [Deprecated] The description for the line item, to be displayed on the Checkout page.
	Description *string `form:"description"`
	// The [tax rates](https://stripe.com/docs/api/tax_rates) that will be applied to this line item depending on the customer's billing/shipping address. We currently support the following countries: US, GB, AU, and all countries in the EU.
	DynamicTaxRates []*string `form:"dynamic_tax_rates"`
	// [Deprecated] A list of image URLs representing this line item. Each image can be up to 5 MB in size. If passing `price` or `price_data`, specify images on the associated product instead.
	Images []*string `form:"images"`
	// [Deprecated] The name for the item to be displayed on the Checkout page. Required if `amount` is passed.
	Name *string `form:"name"`
	// The ID of the [Price](https://stripe.com/docs/api/prices) or [Plan](https://stripe.com/docs/api/plans) object. One of `price` or `price_data` is required.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *CheckoutSessionLineItemPriceDataParams `form:"price_data"`
	// The quantity of the line item being purchased. Quantity should not be defined when `recurring.usage_type=metered`.
	Quantity *int64 `form:"quantity"`
	// The [tax rates](https://stripe.com/docs/api/tax_rates) which apply to this line item.
	TaxRates []*string `form:"tax_rates"`
}

// The parameters used to automatically create a Transfer when the payment succeeds.
// For more information, see the PaymentIntents [use case for connected accounts](https://stripe.com/docs/payments/connected-accounts).
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
	Params `form:"*"`
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. The amount of the application fee collected will be capped at the total payment amount. For more information, see the PaymentIntents [use case for connected accounts](https://stripe.com/docs/payments/connected-accounts).
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The Stripe account ID for which these funds are intended. For details,
	// see the PaymentIntents [use case for connected
	// accounts](https://stripe.com/docs/payments/connected-accounts).
	OnBehalfOf *string `form:"on_behalf_of"`
	// Email address that the receipt for the resulting payment will be sent to. If `receipt_email` is specified for a payment in live mode, a receipt will be sent regardless of your [email settings](https://dashboard.stripe.com/account/emails).
	ReceiptEmail *string `form:"receipt_email"`
	// Indicates that you intend to [make future payments](https://stripe.com/docs/payments/payment-intents#future-usage) with the payment
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
	// Extra information about the payment. This will appear on your
	// customer's statement when this payment succeeds in creating a charge.
	StatementDescriptor *string `form:"statement_descriptor"`
	// Provides information about the charge that customers see on their statements. Concatenated with the
	// prefix (shortened descriptor) or statement descriptor that's set on the account to form the complete
	// statement descriptor. Maximum 22 characters for the concatenated descriptor.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix"`
	// The parameters used to automatically create a Transfer when the payment succeeds.
	// For more information, see the PaymentIntents [use case for connected accounts](https://stripe.com/docs/payments/connected-accounts).
	TransferData *CheckoutSessionPaymentIntentDataTransferDataParams `form:"transfer_data"`
	// A string that identifies the resulting payment as part of a group. See the PaymentIntents [use case for connected accounts](https://stripe.com/docs/payments/connected-accounts) for details.
	TransferGroup *string `form:"transfer_group"`
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

// contains details about the ACSS Debit payment method options.
type CheckoutSessionPaymentMethodOptionsACSSDebitParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies). This is only accepted for Checkout Sessions in `setup` mode.
	Currency *string `form:"currency"`
	// Additional fields for Mandate creation
	MandateOptions *CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// contains details about the Affirm payment method options.
type CheckoutSessionPaymentMethodOptionsAffirmParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Afterpay Clearpay payment method options.
type CheckoutSessionPaymentMethodOptionsAfterpayClearpayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Alipay payment method options.
type CheckoutSessionPaymentMethodOptionsAlipayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the AU Becs Debit payment method options.
type CheckoutSessionPaymentMethodOptionsAUBECSDebitParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Bacs Debit payment method options.
type CheckoutSessionPaymentMethodOptionsBACSDebitParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Bancontact payment method options.
type CheckoutSessionPaymentMethodOptionsBancontactParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Boleto payment method options.
type CheckoutSessionPaymentMethodOptionsBoletoParams struct {
	// The number of calendar days before a Boleto voucher expires. For example, if you create a Boleto voucher on Monday and you set expires_after_days to 2, the Boleto invoice will expire on Wednesday at 23:59 America/Sao_Paulo time.
	ExpiresAfterDays *int64 `form:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// The selected installment plan to use for this payment attempt.
// This parameter can only be provided during confirmation.
type CheckoutSessionPaymentMethodOptionsCardInstallmentsPlanParams struct {
	// For `fixed_count` installment plans, this is the number of installment payments your customer will make to their credit card.
	Count *int64 `form:"count"`
	// For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card.
	// One of `month`.
	Interval *string `form:"interval"`
	// Type of installment plan, one of `fixed_count`.
	Type *string `form:"type"`
}

// Installment options for card payments
type CheckoutSessionPaymentMethodOptionsCardInstallmentsParams struct {
	// Setting to true enables installments for this PaymentIntent.
	// This will cause the response to contain a list of available installment plans.
	// Setting to false will prevent any selected plan from applying to a charge.
	Enabled *bool `form:"enabled"`
	// The selected installment plan to use for this payment attempt.
	// This parameter can only be provided during confirmation.
	Plan *CheckoutSessionPaymentMethodOptionsCardInstallmentsPlanParams `form:"plan"`
}

// contains details about the Card payment method options.
type CheckoutSessionPaymentMethodOptionsCardParams struct {
	// Installment options for card payments
	Installments *CheckoutSessionPaymentMethodOptionsCardInstallmentsParams `form:"installments"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kana prefix (shortened Kana descriptor) or Kana statement descriptor that's set on the account to form the complete statement descriptor. Maximum 22 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 22 characters.
	StatementDescriptorSuffixKana *string `form:"statement_descriptor_suffix_kana"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kanji prefix (shortened Kanji descriptor) or Kanji statement descriptor that's set on the account to form the complete statement descriptor. Maximum 17 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 17 characters.
	StatementDescriptorSuffixKanji *string `form:"statement_descriptor_suffix_kanji"`
}
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information. Permitted values include: `DE`, `ES`, `FR`, `IE`, or `NL`.
	Country *string `form:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	EUBankTransfer *CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer"`
	// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
	//
	// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
	RequestedAddressTypes []*string `form:"requested_address_types"`
	// The list of bank transfer types that this PaymentIntent is allowed to use for funding. Permitted values include: `us_bank_account`, `eu_bank_account`, `id_bank_account`, `gb_bank_account`, `jp_bank_account`, `mx_bank_account`, `eu_bank_transfer`, `gb_bank_transfer`, `id_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.
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
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the EPS payment method options.
type CheckoutSessionPaymentMethodOptionsEPSParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the EPS payment method options.
type CheckoutSessionPaymentMethodOptionsFPXParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Giropay payment method options.
type CheckoutSessionPaymentMethodOptionsGiropayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Grabpay payment method options.
type CheckoutSessionPaymentMethodOptionsGrabpayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Ideal payment method options.
type CheckoutSessionPaymentMethodOptionsIdealParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Klarna payment method options.
type CheckoutSessionPaymentMethodOptionsKlarnaParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Konbini payment method options.
type CheckoutSessionPaymentMethodOptionsKonbiniParams struct {
	// The number of calendar days (between 1 and 60) after which Konbini payment instructions will expire. For example, if a PaymentIntent is confirmed with Konbini and `expires_after_days` set to 2 on Monday JST, the instructions will expire on Wednesday 23:59:59 JST. Defaults to 3 days.
	ExpiresAfterDays *int64 `form:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the OXXO payment method options.
type CheckoutSessionPaymentMethodOptionsOXXOParams struct {
	// The number of calendar days before an OXXO voucher expires. For example, if you create an OXXO voucher on Monday and you set expires_after_days to 2, the OXXO invoice will expire on Wednesday at 23:59 America/Mexico_City time.
	ExpiresAfterDays *int64 `form:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the P24 payment method options.
type CheckoutSessionPaymentMethodOptionsP24Params struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Confirm that the payer has accepted the P24 terms and conditions.
	TOSShownAndAccepted *bool `form:"tos_shown_and_accepted"`
}

// contains details about the PayNow payment method options.
type CheckoutSessionPaymentMethodOptionsPayNowParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Confirm that the payer has accepted the P24 terms and conditions.
	TOSShownAndAccepted *bool `form:"tos_shown_and_accepted"`
}

// contains details about the Sepa Debit payment method options.
type CheckoutSessionPaymentMethodOptionsSepaDebitParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// contains details about the Sofort payment method options.
type CheckoutSessionPaymentMethodOptionsSofortParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Additional fields for Financial Connections Session creation
type CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included. Valid permissions include: `balances`, `ownership`, `payment_method`, and `transactions`.
	Permissions []*string `form:"permissions"`
}

// contains details about the Us Bank Account payment method options.
type CheckoutSessionPaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation
	FinancialConnections *CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// contains details about the WeChat Pay payment method options.
type CheckoutSessionPaymentMethodOptionsWechatPayParams struct {
	// The app ID registered with WeChat Pay. Only required when client is ios or android.
	AppID *string `form:"app_id"`
	// The client type that the end customer will pay from
	Client *string `form:"client"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Payment-method-specific configuration.
type CheckoutSessionPaymentMethodOptionsParams struct {
	// contains details about the ACSS Debit payment method options.
	ACSSDebit *CheckoutSessionPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	// contains details about the Affirm payment method options.
	Affirm *CheckoutSessionPaymentMethodOptionsAffirmParams `form:"affirm"`
	// contains details about the Afterpay Clearpay payment method options.
	AfterpayClearpay *CheckoutSessionPaymentMethodOptionsAfterpayClearpayParams `form:"afterpay_clearpay"`
	// contains details about the Alipay payment method options.
	Alipay *CheckoutSessionPaymentMethodOptionsAlipayParams `form:"alipay"`
	// contains details about the AU Becs Debit payment method options.
	AUBECSDebit *CheckoutSessionPaymentMethodOptionsAUBECSDebitParams `form:"au_becs_debit"`
	// contains details about the Bacs Debit payment method options.
	BACSDebit *CheckoutSessionPaymentMethodOptionsBACSDebitParams `form:"bacs_debit"`
	// contains details about the Bancontact payment method options.
	Bancontact *CheckoutSessionPaymentMethodOptionsBancontactParams `form:"bancontact"`
	// contains details about the Boleto payment method options.
	Boleto *CheckoutSessionPaymentMethodOptionsBoletoParams `form:"boleto"`
	// contains details about the Card payment method options.
	Card *CheckoutSessionPaymentMethodOptionsCardParams `form:"card"`
	// contains details about the Customer Balance payment method options.
	CustomerBalance *CheckoutSessionPaymentMethodOptionsCustomerBalanceParams `form:"customer_balance"`
	// contains details about the EPS payment method options.
	EPS *CheckoutSessionPaymentMethodOptionsEPSParams `form:"eps"`
	// contains details about the EPS payment method options.
	FPX *CheckoutSessionPaymentMethodOptionsFPXParams `form:"fpx"`
	// contains details about the Giropay payment method options.
	Giropay *CheckoutSessionPaymentMethodOptionsGiropayParams `form:"giropay"`
	// contains details about the Grabpay payment method options.
	Grabpay *CheckoutSessionPaymentMethodOptionsGrabpayParams `form:"grabpay"`
	// contains details about the Ideal payment method options.
	Ideal *CheckoutSessionPaymentMethodOptionsIdealParams `form:"ideal"`
	// contains details about the Klarna payment method options.
	Klarna *CheckoutSessionPaymentMethodOptionsKlarnaParams `form:"klarna"`
	// contains details about the Konbini payment method options.
	Konbini *CheckoutSessionPaymentMethodOptionsKonbiniParams `form:"konbini"`
	// contains details about the OXXO payment method options.
	OXXO *CheckoutSessionPaymentMethodOptionsOXXOParams `form:"oxxo"`
	// contains details about the P24 payment method options.
	P24 *CheckoutSessionPaymentMethodOptionsP24Params `form:"p24"`
	// contains details about the PayNow payment method options.
	PayNow *CheckoutSessionPaymentMethodOptionsPayNowParams `form:"paynow"`
	// contains details about the Sepa Debit payment method options.
	SepaDebit *CheckoutSessionPaymentMethodOptionsSepaDebitParams `form:"sepa_debit"`
	// contains details about the Sofort payment method options.
	Sofort *CheckoutSessionPaymentMethodOptionsSofortParams `form:"sofort"`
	// contains details about the Us Bank Account payment method options.
	USBankAccount *CheckoutSessionPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account"`
	// contains details about the WeChat Pay payment method options.
	WechatPay *CheckoutSessionPaymentMethodOptionsWechatPayParams `form:"wechat_pay"`
}

// Controls phone number collection settings for the session.
//
// We recommend that you review your privacy policy and check with your legal contacts
// before using this feature. Learn more about [collecting phone numbers with Checkout](https://stripe.com/docs/payments/checkout/phone-numbers).
type CheckoutSessionPhoneNumberCollectionParams struct {
	// Set to `true` to enable phone number collection.
	Enabled *bool `form:"enabled"`
}

// A subset of parameters to be passed to SetupIntent creation for Checkout Sessions in `setup` mode.
type CheckoutSessionSetupIntentDataParams struct {
	Params `form:"*"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// The Stripe account for which the setup is intended.
	OnBehalfOf *string `form:"on_behalf_of"`
}

// When set, provides configuration for Checkout to collect a shipping address from a customer.
type CheckoutSessionShippingAddressCollectionParams struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for
	// shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SD, SY, UM, VI`.
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

// Parameters to be passed to Shipping Rate creation for this shipping option
type CheckoutSessionShippingOptionShippingRateDataParams struct {
	// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DeliveryEstimate *CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateParams `form:"delivery_estimate"`
	// The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DisplayName *string `form:"display_name"`
	// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
	FixedAmount *CheckoutSessionShippingOptionShippingRateDataFixedAmountParams `form:"fixed_amount"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.
	TaxCode *string `form:"tax_code"`
	// The type of calculation to use on the shipping rate. Can only be `fixed_amount` for now.
	Type *string `form:"type"`
}

// The shipping rate options to apply to this Session.
type CheckoutSessionShippingOptionParams struct {
	// The ID of the Shipping Rate to use for this shipping option.
	ShippingRate *string `form:"shipping_rate"`
	// Parameters to be passed to Shipping Rate creation for this shipping option
	ShippingRateData *CheckoutSessionShippingOptionShippingRateDataParams `form:"shipping_rate_data"`
}

// A list of items, each with an attached plan, that the customer is subscribing to. Prefer using `line_items`.
type CheckoutSessionSubscriptionDataItemsParams struct {
	// Plan ID for this item.
	Plan *string `form:"plan"`
	// The quantity of the subscription item being purchased. Quantity should not be defined when `recurring.usage_type=metered`.
	Quantity *int64 `form:"quantity"`
	// The tax rates which apply to this item. When set, the `default_tax_rates`
	// on `subscription_data` do not apply to this item.
	TaxRates []*string `form:"tax_rates"`
}

// If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.
type CheckoutSessionSubscriptionDataTransferDataParams struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the destination account. By default, the entire amount is transferred to the destination.
	AmountPercent *float64 `form:"amount_percent"`
	// ID of an existing, connected Stripe account.
	Destination *string `form:"destination"`
}

// A subset of parameters to be passed to subscription creation for Checkout Sessions in `subscription` mode.
type CheckoutSessionSubscriptionDataParams struct {
	Params `form:"*"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account. To use an application fee percent, the request must be made on behalf of another account, using the `Stripe-Account` header or an OAuth key. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// The ID of the coupon to apply to this subscription. A coupon applied to a subscription will only affect invoices created for that particular subscription.
	Coupon *string `form:"coupon"`
	// The tax rates that will apply to any subscription item that does not have
	// `tax_rates` set. Invoices created will have their `default_tax_rates` populated
	// from the subscription.
	DefaultTaxRates []*string `form:"default_tax_rates"`
	// The subscription's description, meant to be displayable to the customer.
	// Use this field to optionally store an explanation of the subscription
	// for rendering in Stripe hosted surfaces.
	Description *string `form:"description"`
	// A list of items, each with an attached plan, that the customer is subscribing to. Prefer using `line_items`.
	Items []*CheckoutSessionSubscriptionDataItemsParams `form:"items"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.
	TransferData *CheckoutSessionSubscriptionDataTransferDataParams `form:"transfer_data"`
	// Unix timestamp representing the end of the trial period the customer
	// will get before being charged for the first time. Has to be at least
	// 48 hours in the future.
	TrialEnd *int64 `form:"trial_end"`
	// Indicates if a plan's `trial_period_days` should be applied to the subscription. Setting `trial_end` on `subscription_data` is preferred. Defaults to `false`.
	TrialFromPlan *bool `form:"trial_from_plan"`
	// Integer representing the number of trial period days before the
	// customer is charged for the first time. Has to be at least 1.
	TrialPeriodDays *int64 `form:"trial_period_days"`
}

// Controls tax ID collection settings for the session.
type CheckoutSessionTaxIDCollectionParams struct {
	// Set to true to enable Tax ID collection.
	Enabled *bool `form:"enabled"`
}

// Creates a Session object.
type CheckoutSessionParams struct {
	Params `form:"*"`
	// Configure actions after a Checkout Session has expired.
	AfterExpiration *CheckoutSessionAfterExpirationParams `form:"after_expiration"`
	// Enables user redeemable promotion codes.
	AllowPromotionCodes *bool `form:"allow_promotion_codes"`
	// Settings for automatic tax lookup for this session and resulting payments, invoices, and subscriptions.
	AutomaticTax *CheckoutSessionAutomaticTaxParams `form:"automatic_tax"`
	// Specify whether Checkout should collect the customer's billing address.
	BillingAddressCollection *string `form:"billing_address_collection"`
	// The URL the customer will be directed to if they decide to cancel payment and return to your website.
	CancelURL *string `form:"cancel_url"`
	// A unique string to reference the Checkout Session. This can be a
	// customer ID, a cart ID, or similar, and can be used to reconcile the
	// session with your internal systems.
	ClientReferenceID *string `form:"client_reference_id"`
	// Configure fields for the Checkout Session to gather active consent from customers.
	ConsentCollection *CheckoutSessionConsentCollectionParams `form:"consent_collection"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// ID of an existing Customer, if one exists. In `payment` mode, the customer's most recent card
	// payment method will be used to prefill the email, name, card details, and billing address
	// on the Checkout page. In `subscription` mode, the customer's [default payment method](https://stripe.com/docs/api/customers/update#update_customer-invoice_settings-default_payment_method)
	// will be used if it's a card, and otherwise the most recent card will be used. A valid billing address, billing name and billing email are required on the payment method for Checkout to prefill the customer's card details.
	//
	// If the Customer already has a valid [email](https://stripe.com/docs/api/customers/object#customer_object-email) set, the email will be prefilled and not editable in Checkout.
	// If the Customer does not have a valid `email`, Checkout will set the email entered during the session on the Customer.
	//
	// If blank for Checkout Sessions in `payment` or `subscription` mode, Checkout will create a new Customer object based on information provided during the payment flow.
	//
	// You can set [`payment_intent_data.setup_future_usage`](https://stripe.com/docs/api/checkout/sessions/create#create_checkout_session-payment_intent_data-setup_future_usage) to have Checkout automatically attach the payment method to the Customer you pass in for future reuse.
	Customer *string `form:"customer"`
	// Configure whether a Checkout Session creates a [Customer](https://stripe.com/docs/api/customers) during Session confirmation.
	//
	// When a Customer is not created, you can still retrieve email, address, and other customer data entered in Checkout
	// with [customer_details](https://stripe.com/docs/api/checkout/sessions/object#checkout_session_object-customer_details).
	//
	// Sessions that don't create Customers instead create [Guest Customers](https://support.stripe.com/questions/guest-customer-faq)
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
	// The coupon or promotion code to apply to this Session. Currently, only up to one may be specified.
	Discounts []*CheckoutSessionDiscountParams `form:"discounts"`
	// The Epoch time in seconds at which the Checkout Session will expire. It can be anywhere from 30 minutes to 24 hours after Checkout Session creation. By default, this value is 24 hours from creation.
	ExpiresAt *int64 `form:"expires_at"`
	// A list of items the customer is purchasing. Use this parameter to pass one-time or recurring [Prices](https://stripe.com/docs/api/prices).
	//
	// For `payment` mode, there is a maximum of 100 line items, however it is recommended to consolidate line items if there are more than a few dozen.
	//
	// For `subscription` mode, there is a maximum of 20 line items with recurring Prices and 20 line items with one-time Prices. Line items with one-time Prices in will be on the initial invoice only.
	LineItems []*CheckoutSessionLineItemParams `form:"line_items"`
	// The IETF language tag of the locale Checkout is displayed in. If blank or `auto`, the browser's locale is used.
	Locale *string `form:"locale"`
	// The mode of the Checkout Session. Required when using prices or `setup` mode. Pass `subscription` if the Checkout Session includes at least one recurring item.
	Mode *string `form:"mode"`
	// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
	PaymentIntentData *CheckoutSessionPaymentIntentDataParams `form:"payment_intent_data"`
	// Payment-method-specific configuration.
	PaymentMethodOptions *CheckoutSessionPaymentMethodOptionsParams `form:"payment_method_options"`
	// A list of the types of payment methods (e.g., `card`) this Checkout Session can accept.
	//
	// Do not include this attribute if you prefer to manage your payment methods from the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).
	//
	// Read more about the supported payment methods and their requirements in our [payment
	// method details guide](https://stripe.com/docs/payments/checkout/payment-methods).
	//
	// If multiple payment methods are passed, Checkout will dynamically reorder them to
	// prioritize the most relevant payment methods based on the customer's location and
	// other characteristics.
	PaymentMethodTypes []*string `form:"payment_method_types"`
	// Controls phone number collection settings for the session.
	//
	// We recommend that you review your privacy policy and check with your legal contacts
	// before using this feature. Learn more about [collecting phone numbers with Checkout](https://stripe.com/docs/payments/checkout/phone-numbers).
	PhoneNumberCollection *CheckoutSessionPhoneNumberCollectionParams `form:"phone_number_collection"`
	// A subset of parameters to be passed to SetupIntent creation for Checkout Sessions in `setup` mode.
	SetupIntentData *CheckoutSessionSetupIntentDataParams `form:"setup_intent_data"`
	// When set, provides configuration for Checkout to collect a shipping address from a customer.
	ShippingAddressCollection *CheckoutSessionShippingAddressCollectionParams `form:"shipping_address_collection"`
	// The shipping rate options to apply to this Session.
	ShippingOptions []*CheckoutSessionShippingOptionParams `form:"shipping_options"`
	// [Deprecated] The shipping rate to apply to this Session. Only up to one may be specified.
	ShippingRates []*string `form:"shipping_rates"`
	// Describes the type of transaction being performed by Checkout in order to customize
	// relevant text on the page, such as the submit button. `submit_type` can only be
	// specified on Checkout Sessions in `payment` mode, but not Checkout Sessions
	// in `subscription` or `setup` mode.
	SubmitType *string `form:"submit_type"`
	// A subset of parameters to be passed to subscription creation for Checkout Sessions in `subscription` mode.
	SubscriptionData *CheckoutSessionSubscriptionDataParams `form:"subscription_data"`
	// The URL to which Stripe should send customers when payment or setup
	// is complete.
	// If you'd like to use information from the successful Checkout Session on your page,
	// read the guide on [customizing your success page](https://stripe.com/docs/payments/checkout/custom-success-page).
	SuccessURL *string `form:"success_url"`
	// Controls tax ID collection settings for the session.
	TaxIDCollection *CheckoutSessionTaxIDCollectionParams `form:"tax_id_collection"`
}

// A Session can be expired when it is in one of these statuses: open
//
// After it expires, a customer can't complete a Session and customers loading the Session see a message saying the Session is expired.
type CheckoutSessionExpireParams struct {
	Params `form:"*"`
}

// When retrieving a Checkout Session, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type CheckoutSessionListLineItemsParams struct {
	ListParams `form:"*"`
	Session    *string `form:"-"` // Included in URL
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
type CheckoutSessionAutomaticTax struct {
	// Indicates whether automatic tax is enabled for the session
	Enabled bool `json:"enabled"`
	// The status of the most recent automated tax calculation for this session.
	Status CheckoutSessionAutomaticTaxStatus `json:"status"`
}

// Results of `consent_collection` for this session.
type CheckoutSessionConsent struct {
	// If `opt_in`, the customer consents to receiving promotional communications
	// from the merchant about this Checkout Session.
	Promotions CheckoutSessionConsentPromotions `json:"promotions"`
}

// When set, provides configuration for the Checkout Session to gather active consent from customers.
type CheckoutSessionConsentCollection struct {
	// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
	// Session will determine whether to display an option to opt into promotional communication
	// from the merchant depending on the customer's locale. Only available to US merchants.
	Promotions CheckoutSessionConsentCollectionPromotions `json:"promotions"`
}

// The customer's tax IDs after a completed Checkout Session.
type CheckoutSessionCustomerDetailsTaxIDs struct {
	// The type of the tax ID, one of `eu_vat`, `br_cnpj`, `br_cpf`, `eu_oss_vat`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, or `unknown`
	Type CheckoutSessionCustomerDetailsTaxIDsType `json:"type"`
	// The value of the tax ID.
	Value string `json:"value"`
}

// The customer details including the customer's tax exempt status and the customer's tax IDs. Only the customer's email is present on Sessions in `setup` mode.
type CheckoutSessionCustomerDetails struct {
	// The customer's address after a completed Checkout Session. Note: This property is populated only for sessions on or after March 30, 2022.
	Address *Address `json:"address"`
	// The email associated with the Customer, if one exists, on the Checkout Session after a completed Checkout Session or at time of session expiry.
	// Otherwise, if the customer has consented to promotional content, this value is the most recent valid email provided by the customer on the Checkout form.
	Email string `json:"email"`
	// The customer's name after a completed Checkout Session. Note: This property is populated only for sessions on or after March 30, 2022.
	Name string `json:"name"`
	// The customer's phone number after a completed Checkout Session.
	Phone string `json:"phone"`
	// The customer's tax exempt status after a completed Checkout Session.
	TaxExempt CheckoutSessionCustomerDetailsTaxExempt `json:"tax_exempt"`
	// The customer's tax IDs after a completed Checkout Session.
	TaxIDs []*CheckoutSessionCustomerDetailsTaxIDs `json:"tax_ids"`
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
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsACSSDebitSetupFutureUsage `json:"setup_future_usage"`
	// Bank account verification method.
	VerificationMethod CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}
type CheckoutSessionPaymentMethodOptionsAffirm struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsAffirmSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsAfterpayClearpay struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsAfterpayClearpaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsAlipay struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsAlipaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsAUBECSDebit struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsAUBECSDebitSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsBACSDebit struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsBACSDebitSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsBancontact struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsBancontactSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsBoleto struct {
	// The number of calendar days before a Boleto voucher expires. For example, if you create a Boleto voucher on Monday and you set expires_after_days to 2, the Boleto voucher will expire on Wednesday at 23:59 America/Sao_Paulo time.
	ExpiresAfterDays int64 `json:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsBoletoSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsCardInstallments struct {
	// Indicates if installments are enabled
	Enabled bool `json:"enabled"`
}
type CheckoutSessionPaymentMethodOptionsCard struct {
	Installments *CheckoutSessionPaymentMethodOptionsCardInstallments `json:"installments"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsCardSetupFutureUsage `json:"setup_future_usage"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kana prefix (shortened Kana descriptor) or Kana statement descriptor that's set on the account to form the complete statement descriptor. Maximum 22 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 22 characters.
	StatementDescriptorSuffixKana string `json:"statement_descriptor_suffix_kana"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kanji prefix (shortened Kanji descriptor) or Kanji statement descriptor that's set on the account to form the complete statement descriptor. Maximum 17 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 17 characters.
	StatementDescriptorSuffixKanji string `json:"statement_descriptor_suffix_kanji"`
}
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer struct {
	// The desired country code of the bank account information. Permitted values include: `DE`, `ES`, `FR`, `IE`, or `NL`.
	Country string `json:"country"`
}
type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransfer struct {
	EUBankTransfer *CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer `json:"eu_bank_transfer"`
	// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
	//
	// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
	RequestedAddressTypes []CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType `json:"requested_address_types"`
	// The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, or `mx_bank_transfer`.
	Type CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType `json:"type"`
}
type CheckoutSessionPaymentMethodOptionsCustomerBalance struct {
	BankTransfer *CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType CheckoutSessionPaymentMethodOptionsCustomerBalanceFundingType `json:"funding_type"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsCustomerBalanceSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsEPS struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsEPSSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsFPX struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsFPXSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsGiropay struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsGiropaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsGrabpay struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsGrabpaySetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsIdeal struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsIdealSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsKlarna struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsKlarnaSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsKonbini struct {
	// The number of calendar days (between 1 and 60) after which Konbini payment instructions will expire. For example, if a PaymentIntent is confirmed with Konbini and `expires_after_days` set to 2 on Monday JST, the instructions will expire on Wednesday 23:59:59 JST.
	ExpiresAfterDays int64 `json:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsKonbiniSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsOXXO struct {
	// The number of calendar days before an OXXO invoice expires. For example, if you create an OXXO invoice on Monday and you set expires_after_days to 2, the OXXO invoice will expire on Wednesday at 23:59 America/Mexico_City time.
	ExpiresAfterDays int64 `json:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsOXXOSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsP24 struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsP24SetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsPayNow struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsPayNowSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsSepaDebit struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsSofort struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsSofortSetupFutureUsage `json:"setup_future_usage"`
}
type CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	// The list of permissions to request. The `payment_method` permission must be included.
	Permissions []CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL string `json:"return_url"`
}
type CheckoutSessionPaymentMethodOptionsUSBankAccount struct {
	FinancialConnections *CheckoutSessionPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage CheckoutSessionPaymentMethodOptionsUSBankAccountSetupFutureUsage `json:"setup_future_usage"`
	// Bank account verification method.
	VerificationMethod CheckoutSessionPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment-method-specific configuration for the PaymentIntent or SetupIntent of this CheckoutSession.
type CheckoutSessionPaymentMethodOptions struct {
	ACSSDebit        *CheckoutSessionPaymentMethodOptionsACSSDebit        `json:"acss_debit"`
	Affirm           *CheckoutSessionPaymentMethodOptionsAffirm           `json:"affirm"`
	AfterpayClearpay *CheckoutSessionPaymentMethodOptionsAfterpayClearpay `json:"afterpay_clearpay"`
	Alipay           *CheckoutSessionPaymentMethodOptionsAlipay           `json:"alipay"`
	AUBECSDebit      *CheckoutSessionPaymentMethodOptionsAUBECSDebit      `json:"au_becs_debit"`
	BACSDebit        *CheckoutSessionPaymentMethodOptionsBACSDebit        `json:"bacs_debit"`
	Bancontact       *CheckoutSessionPaymentMethodOptionsBancontact       `json:"bancontact"`
	Boleto           *CheckoutSessionPaymentMethodOptionsBoleto           `json:"boleto"`
	Card             *CheckoutSessionPaymentMethodOptionsCard             `json:"card"`
	CustomerBalance  *CheckoutSessionPaymentMethodOptionsCustomerBalance  `json:"customer_balance"`
	EPS              *CheckoutSessionPaymentMethodOptionsEPS              `json:"eps"`
	FPX              *CheckoutSessionPaymentMethodOptionsFPX              `json:"fpx"`
	Giropay          *CheckoutSessionPaymentMethodOptionsGiropay          `json:"giropay"`
	Grabpay          *CheckoutSessionPaymentMethodOptionsGrabpay          `json:"grabpay"`
	Ideal            *CheckoutSessionPaymentMethodOptionsIdeal            `json:"ideal"`
	Klarna           *CheckoutSessionPaymentMethodOptionsKlarna           `json:"klarna"`
	Konbini          *CheckoutSessionPaymentMethodOptionsKonbini          `json:"konbini"`
	OXXO             *CheckoutSessionPaymentMethodOptionsOXXO             `json:"oxxo"`
	P24              *CheckoutSessionPaymentMethodOptionsP24              `json:"p24"`
	PayNow           *CheckoutSessionPaymentMethodOptionsPayNow           `json:"paynow"`
	SepaDebit        *CheckoutSessionPaymentMethodOptionsSepaDebit        `json:"sepa_debit"`
	Sofort           *CheckoutSessionPaymentMethodOptionsSofort           `json:"sofort"`
	USBankAccount    *CheckoutSessionPaymentMethodOptionsUSBankAccount    `json:"us_bank_account"`
}
type CheckoutSessionPhoneNumberCollection struct {
	// Indicates whether phone number collection is enabled for the session
	Enabled bool `json:"enabled"`
}

// When set, provides configuration for Checkout to collect a shipping address from a customer.
type CheckoutSessionShippingAddressCollection struct {
	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for
	// shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SD, SY, UM, VI`.
	AllowedCountries []string `json:"allowed_countries"`
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
}

// The aggregated discounts.
type CheckoutSessionTotalDetailsBreakdownDiscount struct {
	// The amount discounted.
	Amount int64 `json:"amount"`
	// A discount represents the actual application of a [coupon](https://stripe.com/docs/api#coupons) or [promotion code](https://stripe.com/docs/api#promotion_codes).
	// It contains information about when the discount began, when it will end, and what it is applied to.
	//
	// Related guide: [Applying Discounts to Subscriptions](https://stripe.com/docs/billing/subscriptions/discounts).
	Discount *Discount `json:"discount"`
}

// The aggregated tax amounts by rate.
type CheckoutSessionTotalDetailsBreakdownTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.
	//
	// Related guide: [Tax Rates](https://stripe.com/docs/billing/taxes/tax-rates).
	Rate    *TaxRate `json:"rate"`
	TaxRate *TaxRate `json:"tax_rate"` // Do not use: use `Rate`
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

// A Checkout Session represents your customer's session as they pay for
// one-time purchases or subscriptions through [Checkout](https://stripe.com/docs/payments/checkout)
// or [Payment Links](https://stripe.com/docs/payments/payment-links). We recommend creating a
// new Session each time your customer attempts to pay.
//
// Once payment is successful, the Checkout Session will contain a reference
// to the [Customer](https://stripe.com/docs/api/customers), and either the successful
// [PaymentIntent](https://stripe.com/docs/api/payment_intents) or an active
// [Subscription](https://stripe.com/docs/api/subscriptions).
//
// You can create a Checkout Session on your server and pass its ID to the
// client to begin Checkout.
//
// Related guide: [Checkout Server Quickstart](https://stripe.com/docs/payments/checkout/api).
type CheckoutSession struct {
	APIResource
	// When set, provides configuration for actions to take if this Checkout Session expires.
	AfterExpiration *CheckoutSessionAfterExpiration `json:"after_expiration"`
	// Enables user redeemable promotion codes.
	AllowPromotionCodes bool `json:"allow_promotion_codes"`
	// Total of all items before discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total of all items after discounts and taxes are applied.
	AmountTotal  int64                        `json:"amount_total"`
	AutomaticTax *CheckoutSessionAutomaticTax `json:"automatic_tax"`
	// Describes whether Checkout should collect the customer's billing address.
	BillingAddressCollection CheckoutSessionBillingAddressCollection `json:"billing_address_collection"`
	// The URL the customer will be directed to if they decide to cancel payment and return to your website.
	CancelURL string `json:"cancel_url"`
	// A unique string to reference the Checkout Session. This can be a
	// customer ID, a cart ID, or similar, and can be used to reconcile the
	// Session with your internal systems.
	ClientReferenceID string `json:"client_reference_id"`
	// Results of `consent_collection` for this session.
	Consent *CheckoutSessionConsent `json:"consent"`
	// When set, provides configuration for the Checkout Session to gather active consent from customers.
	ConsentCollection *CheckoutSessionConsentCollection `json:"consent_collection"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The ID of the customer for this Session.
	// For Checkout Sessions in `payment` or `subscription` mode, Checkout
	// will create a new customer object based on information provided
	// during the payment flow unless an existing customer was provided when
	// the Session was created.
	Customer *Customer `json:"customer"`
	// Configure whether a Checkout Session creates a Customer when the Checkout Session completes.
	CustomerCreation CheckoutSessionCustomerCreation `json:"customer_creation"`
	// The customer details including the customer's tax exempt status and the customer's tax IDs. Only the customer's email is present on Sessions in `setup` mode.
	CustomerDetails *CheckoutSessionCustomerDetails `json:"customer_details"`
	// If provided, this value will be used when the Customer object is created.
	// If not provided, customers will be asked to enter their email address.
	// Use this parameter to prefill customer data if you already have an email
	// on file. To access information about the customer once the payment flow is
	// complete, use the `customer` attribute.
	CustomerEmail string `json:"customer_email"`
	Deleted       bool   `json:"deleted"`
	// The timestamp at which the Checkout Session will expire.
	ExpiresAt int64 `json:"expires_at"`
	// Unique identifier for the object. Used to pass to `redirectToCheckout`
	// in Stripe.js.
	ID string `json:"id"`
	// The line items purchased by the customer.
	LineItems *LineItemList `json:"line_items"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The IETF language tag of the locale Checkout is displayed in. If blank or `auto`, the browser's locale is used.
	Locale string `json:"locale"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The mode of the Checkout Session.
	Mode CheckoutSessionMode `json:"mode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The ID of the PaymentIntent for Checkout Sessions in `payment` mode.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
	// The ID of the Payment Link that created this Session.
	PaymentLink *PaymentLink `json:"payment_link"`
	// Payment-method-specific configuration for the PaymentIntent or SetupIntent of this CheckoutSession.
	PaymentMethodOptions *CheckoutSessionPaymentMethodOptions `json:"payment_method_options"`
	// A list of the types of payment methods (e.g. card) this Checkout
	// Session is allowed to accept.
	PaymentMethodTypes []string `json:"payment_method_types"`
	// The payment status of the Checkout Session, one of `paid`, `unpaid`, or `no_payment_required`.
	// You can use this value to decide when to fulfill your customer's order.
	PaymentStatus         CheckoutSessionPaymentStatus          `json:"payment_status"`
	PhoneNumberCollection *CheckoutSessionPhoneNumberCollection `json:"phone_number_collection"`
	// The ID of the original expired Checkout Session that triggered the recovery flow.
	RecoveredFrom string `json:"recovered_from"`
	// The ID of the SetupIntent for Checkout Sessions in `setup` mode.
	SetupIntent *SetupIntent `json:"setup_intent"`
	// Shipping information for this Checkout Session.
	Shipping *ShippingDetails `json:"shipping"`
	// When set, provides configuration for Checkout to collect a shipping address from a customer.
	ShippingAddressCollection *CheckoutSessionShippingAddressCollection `json:"shipping_address_collection"`
	// The shipping rate options applied to this Session.
	ShippingOptions []*CheckoutSessionShippingOption `json:"shipping_options"`
	// The ID of the ShippingRate for Checkout Sessions in `payment` mode.
	ShippingRate *ShippingRate `json:"shipping_rate"`
	// The status of the Checkout Session, one of `open`, `complete`, or `expired`.
	Status CheckoutSessionStatus `json:"status"`
	// Describes the type of transaction being performed by Checkout in order to customize
	// relevant text on the page, such as the submit button. `submit_type` can only be
	// specified on Checkout Sessions in `payment` mode, but not Checkout Sessions
	// in `subscription` or `setup` mode.
	SubmitType CheckoutSessionSubmitType `json:"submit_type"`
	// The ID of the subscription for Checkout Sessions in `subscription` mode.
	Subscription *Subscription `json:"subscription"`
	// The URL the customer will be directed to after the payment or
	// subscription creation is successful.
	SuccessURL      string                          `json:"success_url"`
	TaxIDCollection *CheckoutSessionTaxIDCollection `json:"tax_id_collection"`
	// Tax and discount details for the computed total amount.
	TotalDetails *CheckoutSessionTotalDetails `json:"total_details"`
	// The URL to the Checkout Session. Redirect customers to this URL to take them to Checkout. If you're using [Custom Domains](https://stripe.com/docs/payments/checkout/custom-domains), the URL will use your subdomain. Otherwise, it'll use `checkout.stripe.com.`
	URL string `json:"url"`
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

// CheckoutSessionList is a list of Sessions as retrieved from a list endpoint.
type CheckoutSessionList struct {
	APIResource
	ListMeta
	Data []*CheckoutSession `json:"data"`
}
