//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
)

// The business type.
type AccountBusinessType string

// List of values that AccountBusinessType can take
const (
	AccountBusinessTypeCompany          AccountBusinessType = "company"
	AccountBusinessTypeGovernmentEntity AccountBusinessType = "government_entity"
	AccountBusinessTypeIndividual       AccountBusinessType = "individual"
	AccountBusinessTypeNonProfit        AccountBusinessType = "non_profit"
)

// AccountCapability maps to a given capability for an account.
type AccountCapability string

// List of values that AccountCapability can take.
const (
	AccountCapabilityAUBECSDebitPayments     AccountCapability = "au_becs_debit_payments"
	AccountCapabilityBACSDebitPayments       AccountCapability = "bacs_debit_payments"
	AccountCapabilityBancontactPayments      AccountCapability = "bancontact_payments"
	AccountCapabilityCardIssuing             AccountCapability = "card_issuing"
	AccountCapabilityCardPayments            AccountCapability = "card_payments"
	AccountCapabilityCartesBancairesPayments AccountCapability = "cartes_bancaires_payments"
	AccountCapabilityEPSPayments             AccountCapability = "eps_payments"
	AccountCapabilityFPXPayments             AccountCapability = "fpx_payments"
	AccountCapabilityGiropayPayments         AccountCapability = "giropay_payments"
	AccountCapabilityGrabpayPayments         AccountCapability = "grabpay_payments"
	AccountCapabilityIdealPayments           AccountCapability = "ideal_payments"
	AccountCapabilityJCBPayments             AccountCapability = "jcb_payments"
	AccountCapabilityKlarnaPayments          AccountCapability = "klarna_payments"
	AccountCapabilityLegacyPayments          AccountCapability = "legacy_payments"
	AccountCapabilityOXXOPayments            AccountCapability = "oxxo_payments"
	AccountCapabilityP24Payments             AccountCapability = "p24_payments"
	AccountCapabilitySEPADebitPayments       AccountCapability = "sepa_debit_payments"
	AccountCapabilitySofortPayments          AccountCapability = "sofort_payments"
	AccountCapabilityTaxReportingUS1099K     AccountCapability = "tax_reporting_us_1099_k"
	AccountCapabilityTaxReportingUS1099MISC  AccountCapability = "tax_reporting_us_1099_misc"
	AccountCapabilityTransfers               AccountCapability = "transfers"
)

// The status of the Canadian pre-authorized debits payments capability of the account, or whether the account can directly process Canadian pre-authorized debits charges.
type AccountCapabilityStatus string

// List of values that AccountCapabilityStatus can take
const (
	AccountCapabilityStatusActive   AccountCapabilityStatus = "active"
	AccountCapabilityStatusInactive AccountCapabilityStatus = "inactive"
	AccountCapabilityStatusPending  AccountCapabilityStatus = "pending"
)

// The status of the Affirm capability of the account, or whether the account can directly process Affirm charges.
type AccountCapabilitiesAffirmPayments string

// List of values that AccountCapabilitiesAffirmPayments can take
const (
	AccountCapabilitiesAffirmPaymentsActive   AccountCapabilitiesAffirmPayments = "active"
	AccountCapabilitiesAffirmPaymentsInactive AccountCapabilitiesAffirmPayments = "inactive"
	AccountCapabilitiesAffirmPaymentsPending  AccountCapabilitiesAffirmPayments = "pending"
)

// The status of the customer_balance payments capability of the account, or whether the account can directly process customer_balance charges.
type AccountCapabilitiesBankTransferPayments string

// List of values that AccountCapabilitiesBankTransferPayments can take
const (
	AccountCapabilitiesBankTransferPaymentsActive   AccountCapabilitiesBankTransferPayments = "active"
	AccountCapabilitiesBankTransferPaymentsInactive AccountCapabilitiesBankTransferPayments = "inactive"
	AccountCapabilitiesBankTransferPaymentsPending  AccountCapabilitiesBankTransferPayments = "pending"
)

// The status of the blik payments capability of the account, or whether the account can directly process blik charges.
type AccountCapabilitiesBLIKPayments string

// List of values that AccountCapabilitiesBLIKPayments can take
const (
	AccountCapabilitiesBLIKPaymentsActive   AccountCapabilitiesBLIKPayments = "active"
	AccountCapabilitiesBLIKPaymentsInactive AccountCapabilitiesBLIKPayments = "inactive"
	AccountCapabilitiesBLIKPaymentsPending  AccountCapabilitiesBLIKPayments = "pending"
)

// The status of the konbini payments capability of the account, or whether the account can directly process konbini charges.
type AccountCapabilitiesKonbiniPayments string

// List of values that AccountCapabilitiesKonbiniPayments can take
const (
	AccountCapabilitiesKonbiniPaymentsActive   AccountCapabilitiesKonbiniPayments = "active"
	AccountCapabilitiesKonbiniPaymentsInactive AccountCapabilitiesKonbiniPayments = "inactive"
	AccountCapabilitiesKonbiniPaymentsPending  AccountCapabilitiesKonbiniPayments = "pending"
)

// The status of the link_payments capability of the account, or whether the account can directly process Link charges.
type AccountCapabilitiesLinkPayments string

// List of values that AccountCapabilitiesLinkPayments can take
const (
	AccountCapabilitiesLinkPaymentsActive   AccountCapabilitiesLinkPayments = "active"
	AccountCapabilitiesLinkPaymentsInactive AccountCapabilitiesLinkPayments = "inactive"
	AccountCapabilitiesLinkPaymentsPending  AccountCapabilitiesLinkPayments = "pending"
)

// The status of the paynow payments capability of the account, or whether the account can directly process paynow charges.
type AccountCapabilitiesPayNowPayments string

// List of values that AccountCapabilitiesPayNowPayments can take
const (
	AccountCapabilitiesPayNowPaymentsActive   AccountCapabilitiesPayNowPayments = "active"
	AccountCapabilitiesPayNowPaymentsInactive AccountCapabilitiesPayNowPayments = "inactive"
	AccountCapabilitiesPayNowPaymentsPending  AccountCapabilitiesPayNowPayments = "pending"
)

// The status of the promptpay payments capability of the account, or whether the account can directly process promptpay charges.
type AccountCapabilitiesPromptPayPayments string

// List of values that AccountCapabilitiesPromptPayPayments can take
const (
	AccountCapabilitiesPromptPayPaymentsActive   AccountCapabilitiesPromptPayPayments = "active"
	AccountCapabilitiesPromptPayPaymentsInactive AccountCapabilitiesPromptPayPayments = "inactive"
	AccountCapabilitiesPromptPayPaymentsPending  AccountCapabilitiesPromptPayPayments = "pending"
)

// The status of the banking capability, or whether the account can have bank accounts.
type AccountCapabilitiesTreasury string

// List of values that AccountCapabilitiesTreasury can take
const (
	AccountCapabilitiesTreasuryActive   AccountCapabilitiesTreasury = "active"
	AccountCapabilitiesTreasuryInactive AccountCapabilitiesTreasury = "inactive"
	AccountCapabilitiesTreasuryPending  AccountCapabilitiesTreasury = "pending"
)

// The status of the US bank account ACH payments capability of the account, or whether the account can directly process US bank account charges.
type AccountCapabilitiesUSBankAccountAchPayments string

// List of values that AccountCapabilitiesUSBankAccountAchPayments can take
const (
	AccountCapabilitiesUSBankAccountAchPaymentsActive   AccountCapabilitiesUSBankAccountAchPayments = "active"
	AccountCapabilitiesUSBankAccountAchPaymentsInactive AccountCapabilitiesUSBankAccountAchPayments = "inactive"
	AccountCapabilitiesUSBankAccountAchPaymentsPending  AccountCapabilitiesUSBankAccountAchPayments = "pending"
)

// The category identifying the legal structure of the company or legal entity. See [Business structure](https://stripe.com/docs/connect/identity-verification#business-structure) for more details.
type AccountCompanyStructure string

// List of values that AccountCompanyStructure can take
const (
	AccountCompanyStructureFreeZoneEstablishment              AccountCompanyStructure = "free_zone_establishment"
	AccountCompanyStructureFreeZoneLLC                        AccountCompanyStructure = "free_zone_llc"
	AccountCompanyStructureGovernmentInstrumentality          AccountCompanyStructure = "government_instrumentality"
	AccountCompanyStructureGovernmentalUnit                   AccountCompanyStructure = "governmental_unit"
	AccountCompanyStructureIncorporatedNonProfit              AccountCompanyStructure = "incorporated_non_profit"
	AccountCompanyStructureLimitedLiabilityPartnership        AccountCompanyStructure = "limited_liability_partnership"
	AccountCompanyStructureLLC                                AccountCompanyStructure = "llc"
	AccountCompanyStructureMultiMemberLLC                     AccountCompanyStructure = "multi_member_llc"
	AccountCompanyStructurePrivateCompany                     AccountCompanyStructure = "private_company"
	AccountCompanyStructurePrivateCorporation                 AccountCompanyStructure = "private_corporation"
	AccountCompanyStructurePrivatePartnership                 AccountCompanyStructure = "private_partnership"
	AccountCompanyStructurePublicCompany                      AccountCompanyStructure = "public_company"
	AccountCompanyStructurePublicCorporation                  AccountCompanyStructure = "public_corporation"
	AccountCompanyStructurePublicPartnership                  AccountCompanyStructure = "public_partnership"
	AccountCompanyStructureSingleMemberLLC                    AccountCompanyStructure = "single_member_llc"
	AccountCompanyStructureSoleEstablishment                  AccountCompanyStructure = "sole_establishment"
	AccountCompanyStructureSoleProprietorship                 AccountCompanyStructure = "sole_proprietorship"
	AccountCompanyStructureTaxExemptGovernmentInstrumentality AccountCompanyStructure = "tax_exempt_government_instrumentality"
	AccountCompanyStructureUnincorporatedAssociation          AccountCompanyStructure = "unincorporated_association"
	AccountCompanyStructureUnincorporatedNonProfit            AccountCompanyStructure = "unincorporated_non_profit"
)

// AccountRejectReason describes the valid reason to reject an account
type AccountRejectReason string

// List of values that AccountRejectReason can take.
const (
	AccountRejectReasonFraud          AccountRejectReason = "fraud"
	AccountRejectReasonOther          AccountRejectReason = "other"
	AccountRejectReasonTermsOfService AccountRejectReason = "terms_of_service"
)

// One of `document_corrupt`, `document_expired`, `document_failed_copy`, `document_failed_greyscale`, `document_failed_other`, `document_failed_test_mode`, `document_fraudulent`, `document_incomplete`, `document_invalid`, `document_manipulated`, `document_not_readable`, `document_not_uploaded`, `document_type_not_supported`, or `document_too_large`. A machine-readable code specifying the verification state for this document.
type AccountCompanyVerificationDocumentDetailsCode string

// List of values that AccountCompanyVerificationDocumentDetailsCode can take
const (
	AccountCompanyVerificationDocumentDetailsCodeDocumentCorrupt        AccountCompanyVerificationDocumentDetailsCode = "document_corrupt"
	AccountCompanyVerificationDocumentDetailsCodeDocumentFailedCopy     AccountCompanyVerificationDocumentDetailsCode = "document_failed_copy"
	AccountCompanyVerificationDocumentDetailsCodeDocumentFailedOther    AccountCompanyVerificationDocumentDetailsCode = "document_failed_other"
	AccountCompanyVerificationDocumentDetailsCodeDocumentFailedTestMode AccountCompanyVerificationDocumentDetailsCode = "document_failed_test_mode"
	AccountCompanyVerificationDocumentDetailsCodeDocumentFraudulent     AccountCompanyVerificationDocumentDetailsCode = "document_fraudulent"
	AccountCompanyVerificationDocumentDetailsCodeDocumentInvalid        AccountCompanyVerificationDocumentDetailsCode = "document_invalid"
	AccountCompanyVerificationDocumentDetailsCodeDocumentManipulated    AccountCompanyVerificationDocumentDetailsCode = "document_manipulated"
	AccountCompanyVerificationDocumentDetailsCodeDocumentNotReadable    AccountCompanyVerificationDocumentDetailsCode = "document_not_readable"
	AccountCompanyVerificationDocumentDetailsCodeDocumentNotUploaded    AccountCompanyVerificationDocumentDetailsCode = "document_not_uploaded"
	AccountCompanyVerificationDocumentDetailsCodeDocumentTooLarge       AccountCompanyVerificationDocumentDetailsCode = "document_too_large"
)

// The controller type. Can be `application`, if a Connect application controls the account, or `account`, if the account controls itself.
type AccountControllerType string

// List of values that AccountControllerType can take
const (
	AccountControllerTypeAccount     AccountControllerType = "account"
	AccountControllerTypeApplication AccountControllerType = "application"
)

type ExternalAccountType string

// List of values that ExternalAccountType can take
const (
	ExternalAccountTypeBankAccount ExternalAccountType = "bank_account"
	ExternalAccountTypeCard        ExternalAccountType = "card"
)

// If the account is disabled, this string describes why. Can be `requirements.past_due`, `requirements.pending_verification`, `listed`, `platform_paused`, `rejected.fraud`, `rejected.listed`, `rejected.terms_of_service`, `rejected.other`, `under_review`, or `other`.
type AccountRequirementsDisabledReason string

// List of values that AccountRequirementsDisabledReason can take
const (
	AccountRequirementsDisabledReasonFieldsNeeded           AccountRequirementsDisabledReason = "fields_needed"
	AccountRequirementsDisabledReasonListed                 AccountRequirementsDisabledReason = "listed"
	AccountRequirementsDisabledReasonOther                  AccountRequirementsDisabledReason = "other"
	AccountRequirementsDisabledReasonRejectedFraud          AccountRequirementsDisabledReason = "rejected.fraud"
	AccountRequirementsDisabledReasonRejectedListed         AccountRequirementsDisabledReason = "rejected.listed"
	AccountRequirementsDisabledReasonRejectedOther          AccountRequirementsDisabledReason = "rejected.other"
	AccountRequirementsDisabledReasonRejectedTermsOfService AccountRequirementsDisabledReason = "rejected.terms_of_service"
	AccountRequirementsDisabledReasonUnderReview            AccountRequirementsDisabledReason = "under_review"
)

// How frequently funds will be paid out. One of `manual` (payouts only created via API call), `daily`, `weekly`, or `monthly`.
type PayoutInterval string

// List of values that PayoutInterval can take
const (
	PayoutIntervalDaily   PayoutInterval = "daily"
	PayoutIntervalManual  PayoutInterval = "manual"
	PayoutIntervalMonthly PayoutInterval = "monthly"
	PayoutIntervalWeekly  PayoutInterval = "weekly"
)

// The user's service agreement type
type AccountTOSAcceptanceServiceAgreement string

// List of values that AccountTOSAcceptanceServiceAgreement can take
const (
	AccountTOSAcceptanceServiceAgreementFull      AccountTOSAcceptanceServiceAgreement = "full"
	AccountTOSAcceptanceServiceAgreementRecipient AccountTOSAcceptanceServiceAgreement = "recipient"
)

// The Stripe account type. Can be `standard`, `express`, or `custom`.
type AccountType string

// List of values that AccountType can take
const (
	AccountTypeCustom   AccountType = "custom"
	AccountTypeExpress  AccountType = "express"
	AccountTypeStandard AccountType = "standard"
)

// Retrieves the details of an account.
type AccountParams struct {
	Params `form:"*"`
	// An [account token](https://stripe.com/docs/api#create_account_token), used to securely provide details to the account.
	AccountToken *string `form:"account_token"`
	// Business information about the account.
	BusinessProfile *AccountBusinessProfileParams `form:"business_profile"`
	// The business type.
	BusinessType *string `form:"business_type"`
	// Each key of the dictionary represents a capability, and each capability maps to its settings (e.g. whether it has been requested or not). Each capability will be inactive until you have provided its specific requirements and Stripe has verified them. An account may have some of its requested capabilities be active and some be inactive.
	Capabilities *AccountCapabilitiesParams `form:"capabilities"`
	// Information about the company or business. This field is available for any `business_type`.
	Company *AccountCompanyParams `form:"company"`
	// The country in which the account holder resides, or in which the business is legally established. This should be an ISO 3166-1 alpha-2 country code. For example, if you are in the United States and the business for which you're creating an account is legally represented in Canada, you would use `CA` as the country for the account being created. Available countries include [Stripe's global markets](https://stripe.com/global) as well as countries where [cross-border payouts](https://stripe.com/docs/connect/cross-border-payouts) are supported.
	Country *string `form:"country"`
	// Three-letter ISO currency code representing the default currency for the account. This must be a currency that [Stripe supports in the account's country](https://stripe.com/docs/payouts).
	DefaultCurrency *string `form:"default_currency"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *AccountDocumentsParams `form:"documents"`
	// The email address of the account holder. This is only to make the account easier to identify to you. Stripe only emails Custom accounts with your consent.
	Email *string `form:"email"`
	// A card or bank account to attach to the account for receiving [payouts](https://stripe.com/docs/connect/bank-debit-card-payouts) (you won't be able to use it for top-ups). You can provide either a token, like the ones returned by [Stripe.js](https://stripe.com/docs/js), or a dictionary, as documented in the `external_account` parameter for [bank account](https://stripe.com/docs/api#account_create_bank_account) creation.
	//
	// By default, providing an external account sets it as the new default external account for its currency, and deletes the old default if one exists. To add additional external accounts without replacing the existing default for the currency, use the bank account or card creation API.
	ExternalAccount *AccountExternalAccountParams `form:"external_account"`
	// Information about the person represented by the account. This field is null unless `business_type` is set to `individual`.
	Individual *PersonParams `form:"individual"`
	// Options for customizing how the account functions within Stripe.
	Settings *AccountSettingsParams `form:"settings"`
	// Details on the account's acceptance of the [Stripe Services Agreement](https://stripe.com/docs/connect/updating-accounts#tos-acceptance).
	TOSAcceptance *AccountTOSAcceptanceParams `form:"tos_acceptance"`
	// The type of Stripe account to create. May be one of `custom`, `express` or `standard`.
	Type *string `form:"type"`
	// This parameter is deprecated. Prefer using Capabilities instead.
	RequestedCapabilities []*string `form:"requested_capabilities"`
}

// Business information about the account.
type AccountBusinessProfileParams struct {
	// [The merchant category code for the account](https://stripe.com/docs/connect/setting-mcc). MCCs are used to classify businesses based on the goods or services they provide.
	MCC *string `form:"mcc"`
	// The customer-facing business name.
	Name *string `form:"name"`
	// Internal-only description of the product sold by, or service provided by, the business. Used by Stripe for risk and underwriting purposes.
	ProductDescription *string `form:"product_description"`
	// A publicly available mailing address for sending support issues to.
	SupportAddress *AddressParams `form:"support_address"`
	// A publicly available email address for sending support issues to.
	SupportEmail *string `form:"support_email"`
	// A publicly available phone number to call with support issues.
	SupportPhone *string `form:"support_phone"`
	// A publicly available website for handling support issues.
	SupportURL *string `form:"support_url"`
	// The business's publicly available website.
	URL *string `form:"url"`
}

// The acss_debit_payments capability.
type AccountCapabilitiesACSSDebitPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The affirm_payments capability.
type AccountCapabilitiesAffirmPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The afterpay_clearpay_payments capability.
type AccountCapabilitiesAfterpayClearpayPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The au_becs_debit_payments capability.
type AccountCapabilitiesAUBECSDebitPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The bacs_debit_payments capability.
type AccountCapabilitiesBACSDebitPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The bancontact_payments capability.
type AccountCapabilitiesBancontactPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The bank_transfer_payments capability.
type AccountCapabilitiesBankTransferPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The blik_payments capability.
type AccountCapabilitiesBLIKPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The boleto_payments capability.
type AccountCapabilitiesBoletoPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The card_issuing capability.
type AccountCapabilitiesCardIssuingParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The card_payments capability.
type AccountCapabilitiesCardPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The cartes_bancaires_payments capability.
type AccountCapabilitiesCartesBancairesPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The eps_payments capability.
type AccountCapabilitiesEPSPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The fpx_payments capability.
type AccountCapabilitiesFPXPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The giropay_payments capability.
type AccountCapabilitiesGiropayPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The grabpay_payments capability.
type AccountCapabilitiesGrabpayPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The ideal_payments capability.
type AccountCapabilitiesIdealPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The jcb_payments capability.
type AccountCapabilitiesJCBPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The klarna_payments capability.
type AccountCapabilitiesKlarnaPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The konbini_payments capability.
type AccountCapabilitiesKonbiniPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The legacy_payments capability.
type AccountCapabilitiesLegacyPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The link_payments capability.
type AccountCapabilitiesLinkPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The oxxo_payments capability.
type AccountCapabilitiesOXXOPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The p24_payments capability.
type AccountCapabilitiesP24PaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The paynow_payments capability.
type AccountCapabilitiesPayNowPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The promptpay_payments capability.
type AccountCapabilitiesPromptPayPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The sepa_debit_payments capability.
type AccountCapabilitiesSEPADebitPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The sofort_payments capability.
type AccountCapabilitiesSofortPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The tax_reporting_us_1099_k capability.
type AccountCapabilitiesTaxReportingUS1099KParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The tax_reporting_us_1099_misc capability.
type AccountCapabilitiesTaxReportingUS1099MISCParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The transfers capability.
type AccountCapabilitiesTransfersParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The treasury capability.
type AccountCapabilitiesTreasuryParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// The us_bank_account_ach_payments capability.
type AccountCapabilitiesUSBankAccountAchPaymentsParams struct {
	// Passing true requests the capability for the account, if it is not already requested. A requested capability may not immediately become active. Any requirements to activate the capability are returned in the `requirements` arrays.
	Requested *bool `form:"requested"`
}

// Each key of the dictionary represents a capability, and each capability maps to its settings (e.g. whether it has been requested or not). Each capability will be inactive until you have provided its specific requirements and Stripe has verified them. An account may have some of its requested capabilities be active and some be inactive.
type AccountCapabilitiesParams struct {
	// The acss_debit_payments capability.
	ACSSDebitPayments *AccountCapabilitiesACSSDebitPaymentsParams `form:"acss_debit_payments"`
	// The affirm_payments capability.
	AffirmPayments *AccountCapabilitiesAffirmPaymentsParams `form:"affirm_payments"`
	// The afterpay_clearpay_payments capability.
	AfterpayClearpayPayments *AccountCapabilitiesAfterpayClearpayPaymentsParams `form:"afterpay_clearpay_payments"`
	// The au_becs_debit_payments capability.
	AUBECSDebitPayments *AccountCapabilitiesAUBECSDebitPaymentsParams `form:"au_becs_debit_payments"`
	// The bacs_debit_payments capability.
	BACSDebitPayments *AccountCapabilitiesBACSDebitPaymentsParams `form:"bacs_debit_payments"`
	// The bancontact_payments capability.
	BancontactPayments *AccountCapabilitiesBancontactPaymentsParams `form:"bancontact_payments"`
	// The bank_transfer_payments capability.
	BankTransferPayments *AccountCapabilitiesBankTransferPaymentsParams `form:"bank_transfer_payments"`
	// The blik_payments capability.
	BLIKPayments *AccountCapabilitiesBLIKPaymentsParams `form:"blik_payments"`
	// The boleto_payments capability.
	BoletoPayments *AccountCapabilitiesBoletoPaymentsParams `form:"boleto_payments"`
	// The card_issuing capability.
	CardIssuing *AccountCapabilitiesCardIssuingParams `form:"card_issuing"`
	// The card_payments capability.
	CardPayments *AccountCapabilitiesCardPaymentsParams `form:"card_payments"`
	// The cartes_bancaires_payments capability.
	CartesBancairesPayments *AccountCapabilitiesCartesBancairesPaymentsParams `form:"cartes_bancaires_payments"`
	// The eps_payments capability.
	EPSPayments *AccountCapabilitiesEPSPaymentsParams `form:"eps_payments"`
	// The fpx_payments capability.
	FPXPayments *AccountCapabilitiesFPXPaymentsParams `form:"fpx_payments"`
	// The giropay_payments capability.
	GiropayPayments *AccountCapabilitiesGiropayPaymentsParams `form:"giropay_payments"`
	// The grabpay_payments capability.
	GrabpayPayments *AccountCapabilitiesGrabpayPaymentsParams `form:"grabpay_payments"`
	// The ideal_payments capability.
	IdealPayments *AccountCapabilitiesIdealPaymentsParams `form:"ideal_payments"`
	// The jcb_payments capability.
	JCBPayments *AccountCapabilitiesJCBPaymentsParams `form:"jcb_payments"`
	// The klarna_payments capability.
	KlarnaPayments *AccountCapabilitiesKlarnaPaymentsParams `form:"klarna_payments"`
	// The konbini_payments capability.
	KonbiniPayments *AccountCapabilitiesKonbiniPaymentsParams `form:"konbini_payments"`
	// The legacy_payments capability.
	LegacyPayments *AccountCapabilitiesLegacyPaymentsParams `form:"legacy_payments"`
	// The link_payments capability.
	LinkPayments *AccountCapabilitiesLinkPaymentsParams `form:"link_payments"`
	// The oxxo_payments capability.
	OXXOPayments *AccountCapabilitiesOXXOPaymentsParams `form:"oxxo_payments"`
	// The p24_payments capability.
	P24Payments *AccountCapabilitiesP24PaymentsParams `form:"p24_payments"`
	// The paynow_payments capability.
	PayNowPayments *AccountCapabilitiesPayNowPaymentsParams `form:"paynow_payments"`
	// The promptpay_payments capability.
	PromptPayPayments *AccountCapabilitiesPromptPayPaymentsParams `form:"promptpay_payments"`
	// The sepa_debit_payments capability.
	SEPADebitPayments *AccountCapabilitiesSEPADebitPaymentsParams `form:"sepa_debit_payments"`
	// The sofort_payments capability.
	SofortPayments *AccountCapabilitiesSofortPaymentsParams `form:"sofort_payments"`
	// The tax_reporting_us_1099_k capability.
	TaxReportingUS1099K *AccountCapabilitiesTaxReportingUS1099KParams `form:"tax_reporting_us_1099_k"`
	// The tax_reporting_us_1099_misc capability.
	TaxReportingUS1099MISC *AccountCapabilitiesTaxReportingUS1099MISCParams `form:"tax_reporting_us_1099_misc"`
	// The transfers capability.
	Transfers *AccountCapabilitiesTransfersParams `form:"transfers"`
	// The treasury capability.
	Treasury *AccountCapabilitiesTreasuryParams `form:"treasury"`
	// The us_bank_account_ach_payments capability.
	USBankAccountAchPayments *AccountCapabilitiesUSBankAccountAchPaymentsParams `form:"us_bank_account_ach_payments"`
}

// The Kana variation of the company's primary address (Japan only).
type AccountAddressParams struct {
	// City or ward.
	City *string `form:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country"`
	// Block or building number.
	Line1 *string `form:"line1"`
	// Building details.
	Line2 *string `form:"line2"`
	// Postal code.
	PostalCode *string `form:"postal_code"`
	// Prefecture.
	State *string `form:"state"`
	// Town/cho-me. Note that this is only used for Kana/Kanji representations
	// of an address.
	// Town or cho-me.
	Town *string `form:"town"`
}

// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
type AccountCompanyOwnershipDeclarationParams struct {
	// The Unix timestamp marking when the beneficial owner attestation was made.
	Date *int64 `form:"date"`
	// The IP address from which the beneficial owner attestation was made.
	IP *string `form:"ip"`
	// The user agent of the browser from which the beneficial owner attestation was made.
	UserAgent *string `form:"user_agent"`
}

// A document verifying the business.
type AccountCompanyVerificationDocumentParams struct {
	// The back of a document returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `additional_verification`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back"`
	// The front of a document returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `additional_verification`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front"`
}

// Information on the verification state of the company.
type AccountCompanyVerificationParams struct {
	// A document verifying the business.
	Document *AccountCompanyVerificationDocumentParams `form:"document"`
}

// Information about the company or business. This field is available for any `business_type`.
type AccountCompanyParams struct {
	// The company's primary address.
	Address *AccountAddressParams `form:"address"`
	// The Kana variation of the company's primary address (Japan only).
	AddressKana *AccountAddressParams `form:"address_kana"`
	// The Kanji variation of the company's primary address (Japan only).
	AddressKanji *AccountAddressParams `form:"address_kanji"`
	// Whether the company's directors have been provided. Set this Boolean to `true` after creating all the company's directors with [the Persons API](https://stripe.com/docs/api/persons) for accounts with a `relationship.director` requirement. This value is not automatically set to `true` after creating directors, so it needs to be updated to indicate all directors have been provided.
	DirectorsProvided *bool `form:"directors_provided"`
	// Whether the company's executives have been provided. Set this Boolean to `true` after creating all the company's executives with [the Persons API](https://stripe.com/docs/api/persons) for accounts with a `relationship.executive` requirement.
	ExecutivesProvided *bool `form:"executives_provided"`
	// The company's legal name.
	Name *string `form:"name"`
	// The Kana variation of the company's legal name (Japan only).
	NameKana *string `form:"name_kana"`
	// The Kanji variation of the company's legal name (Japan only).
	NameKanji *string `form:"name_kanji"`
	// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
	OwnershipDeclaration *AccountCompanyOwnershipDeclarationParams `form:"ownership_declaration"`
	// This parameter can only be used on Token creation.
	OwnershipDeclarationShownAndSigned *bool `form:"ownership_declaration_shown_and_signed"`
	// Whether the company's owners have been provided. Set this Boolean to `true` after creating all the company's owners with [the Persons API](https://stripe.com/docs/api/persons) for accounts with a `relationship.owner` requirement.
	OwnersProvided *bool `form:"owners_provided"`
	// The company's phone number (used for verification).
	Phone *string `form:"phone"`
	// The identification number given to a company when it is registered or incorporated, if distinct from the identification number used for filing taxes. (Examples are the CIN for companies and LLP IN for partnerships in India, and the Company Registration Number in Hong Kong).
	RegistrationNumber *string `form:"registration_number"`
	// The category identifying the legal structure of the company or legal entity. See [Business structure](https://stripe.com/docs/connect/identity-verification#business-structure) for more details.
	Structure *string `form:"structure"`
	// The business ID number of the company, as appropriate for the company's country. (Examples are an Employer ID Number in the U.S., a Business Number in Canada, or a Company Number in the UK.)
	TaxID *string `form:"tax_id"`
	// The jurisdiction in which the `tax_id` is registered (Germany-based companies only).
	TaxIDRegistrar *string `form:"tax_id_registrar"`
	// The VAT number of the company.
	VATID *string `form:"vat_id"`
	// Information on the verification state of the company.
	Verification *AccountCompanyVerificationParams `form:"verification"`
}

// One or more documents that support the [Bank account ownership verification](https://support.stripe.com/questions/bank-account-ownership-verification) requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a voided check.
type AccountDocumentsBankAccountOwnershipVerificationParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// One or more documents that demonstrate proof of a company's license to operate.
type AccountDocumentsCompanyLicenseParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// One or more documents showing the company's Memorandum of Association.
type AccountDocumentsCompanyMemorandumOfAssociationParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// (Certain countries only) One or more documents showing the ministerial decree legalizing the company's establishment.
type AccountDocumentsCompanyMinisterialDecreeParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
type AccountDocumentsCompanyRegistrationVerificationParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// One or more documents that demonstrate proof of a company's tax ID.
type AccountDocumentsCompanyTaxIDVerificationParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// One or more documents showing the company's proof of registration with the national business registry.
type AccountDocumentsProofOfRegistrationParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// Documents that may be submitted to satisfy various informational requests.
type AccountDocumentsParams struct {
	// One or more documents that support the [Bank account ownership verification](https://support.stripe.com/questions/bank-account-ownership-verification) requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a voided check.
	BankAccountOwnershipVerification *AccountDocumentsBankAccountOwnershipVerificationParams `form:"bank_account_ownership_verification"`
	// One or more documents that demonstrate proof of a company's license to operate.
	CompanyLicense *AccountDocumentsCompanyLicenseParams `form:"company_license"`
	// One or more documents showing the company's Memorandum of Association.
	CompanyMemorandumOfAssocation *AccountDocumentsCompanyMemorandumOfAssociationParams `form:"company_memorandum_of_association"`
	// (Certain countries only) One or more documents showing the ministerial decree legalizing the company's establishment.
	CompanyMinisterialDecree *AccountDocumentsCompanyMinisterialDecreeParams `form:"company_ministerial_decree"`
	// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
	CompanyRegistrationVerification *AccountDocumentsCompanyRegistrationVerificationParams `form:"company_registration_verification"`
	// One or more documents that demonstrate proof of a company's tax ID.
	CompanyTaxIDVerification *AccountDocumentsCompanyTaxIDVerificationParams `form:"company_tax_id_verification"`
	// One or more documents showing the company's proof of registration with the national business registry.
	ProofOfRegistration *AccountDocumentsProofOfRegistrationParams `form:"proof_of_registration"`
}

// AccountSettingsBACSDebitPaymentsParams represent allowed parameters to configure settings specific to
// BACS Debit charging on the account.
type AccountSettingsBACSDebitPaymentsParams struct {
	DisplayName *string `form:"display_name"`
}

// Settings used to apply the account's branding to email receipts, invoices, Checkout, and other products.
type AccountSettingsBrandingParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) An icon for the account. Must be square and at least 128px x 128px.
	Icon *string `form:"icon"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) A logo for the account that will be used in Checkout instead of the icon and without the account's name next to it if provided. Must be at least 128px x 128px.
	Logo *string `form:"logo"`
	// A CSS hex color value representing the primary branding color for this account.
	PrimaryColor *string `form:"primary_color"`
	// A CSS hex color value representing the secondary branding color for this account.
	SecondaryColor *string `form:"secondary_color"`
}

// Details on the account's acceptance of the [Stripe Issuing Terms and Disclosures](https://stripe.com/docs/issuing/connect/tos_acceptance).
type AccountTOSAcceptanceParams struct {
	// The Unix timestamp marking when the account representative accepted their service agreement.
	Date *int64 `form:"date"`
	// The IP address from which the account representative accepted their service agreement.
	IP *string `form:"ip"`
	// The user's service agreement type.
	ServiceAgreement *string `form:"service_agreement"`
	// The user agent of the browser from which the account representative accepted their service agreement.
	UserAgent *string `form:"user_agent"`
}

// Settings specific to the account's use of the Card Issuing product.
type AccountSettingsCardIssuingParams struct {
	// Details on the account's acceptance of the [Stripe Issuing Terms and Disclosures](https://stripe.com/docs/issuing/connect/tos_acceptance).
	TOSAcceptance *AccountTOSAcceptanceParams `form:"tos_acceptance"`
}

// Automatically declines certain charge types regardless of whether the card issuer accepted or declined the charge.
type AccountDeclineSettingsParams struct {
	// Whether Stripe automatically declines charges with an incorrect ZIP or postal code. This setting only applies when a ZIP or postal code is provided and they fail bank verification.
	AVSFailure *bool `form:"avs_failure"`
	// Whether Stripe automatically declines charges with an incorrect CVC. This setting only applies when a CVC is provided and it fails bank verification.
	CVCFailure *bool `form:"cvc_failure"`
}

// Settings specific to card charging on the account.
type AccountSettingsCardPaymentsParams struct {
	// Automatically declines certain charge types regardless of whether the card issuer accepted or declined the charge.
	DeclineOn *AccountDeclineSettingsParams `form:"decline_on"`
	// The default text that appears on credit card statements when a charge is made. This field prefixes any dynamic `statement_descriptor` specified on the charge. `statement_descriptor_prefix` is useful for maximizing descriptor space for the dynamic portion.
	StatementDescriptorPrefix *string `form:"statement_descriptor_prefix"`
	// The Kana variation of the default text that appears on credit card statements when a charge is made (Japan only). This field prefixes any dynamic `statement_descriptor_suffix_kana` specified on the charge. `statement_descriptor_prefix_kana` is useful for maximizing descriptor space for the dynamic portion.
	StatementDescriptorPrefixKana *string `form:"statement_descriptor_prefix_kana"`
	// The Kanji variation of the default text that appears on credit card statements when a charge is made (Japan only). This field prefixes any dynamic `statement_descriptor_suffix_kanji` specified on the charge. `statement_descriptor_prefix_kanji` is useful for maximizing descriptor space for the dynamic portion.
	StatementDescriptorPrefixKanji *string `form:"statement_descriptor_prefix_kanji"`
}

// AccountSettingsDashboardParams represent allowed parameters to configure settings for the
// account's Dashboard.
type AccountSettingsDashboardParams struct {
	DisplayName *string `form:"display_name"`
	Timezone    *string `form:"timezone"`
}

// Settings that apply across payment methods for charging on the account.
type AccountSettingsPaymentsParams struct {
	// The default text that appears on credit card statements when a charge is made. This field prefixes any dynamic `statement_descriptor` specified on the charge.
	StatementDescriptor *string `form:"statement_descriptor"`
	// The Kana variation of the default text that appears on credit card statements when a charge is made (Japan only).
	StatementDescriptorKana *string `form:"statement_descriptor_kana"`
	// The Kanji variation of the default text that appears on credit card statements when a charge is made (Japan only).
	StatementDescriptorKanji *string `form:"statement_descriptor_kanji"`
}

// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://stripe.com/docs/connect/bank-transfers#payout-information) documentation.
type PayoutScheduleParams struct {
	// The number of days charge funds are held before being paid out. May also be set to `minimum`, representing the lowest available value for the account country. Default is `minimum`. The `delay_days` parameter does not apply when the `interval` is `manual`.
	DelayDays        *int64 `form:"delay_days"`
	DelayDaysMinimum *bool  `form:"-"` // See custom AppendTo
	// How frequently available funds are paid out. One of: `daily`, `manual`, `weekly`, or `monthly`. Default is `daily`.
	Interval *string `form:"interval"`
	// The day of the month when available funds are paid out, specified as a number between 1--31. Payouts nominally scheduled between the 29th and 31st of the month are instead sent on the last day of a shorter month. Required and applicable only if `interval` is `monthly`.
	MonthlyAnchor *int64 `form:"monthly_anchor"`
	// The day of the week when available funds are paid out, specified as `monday`, `tuesday`, etc. (required and applicable only if `interval` is `weekly`.)
	WeeklyAnchor *string `form:"weekly_anchor"`
}

// AppendTo implements custom encoding logic for PayoutScheduleParams.
func (p *PayoutScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.DelayDaysMinimum) {
		body.Add(form.FormatKey(append(keyParts, "delay_days")), "minimum")
	}
}

// Settings specific to the account's payouts.
type AccountSettingsPayoutsParams struct {
	// A Boolean indicating whether Stripe should try to reclaim negative balances from an attached bank account. For details, see [Understanding Connect Account Balances](https://stripe.com/docs/connect/account-balances).
	DebitNegativeBalances *bool `form:"debit_negative_balances"`
	// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://stripe.com/docs/connect/bank-transfers#payout-information) documentation.
	Schedule *PayoutScheduleParams `form:"schedule"`
	// The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.
	StatementDescriptor *string `form:"statement_descriptor"`
}

// Details on the account's acceptance of the Stripe Treasury Services Agreement.
type AccountSettingsTreasuryTOSAcceptanceParams struct {
	// The Unix timestamp marking when the account representative accepted the service agreement.
	Date *int64 `form:"date"`
	// The IP address from which the account representative accepted the service agreement.
	IP *string `form:"ip"`
	// The user agent of the browser from which the account representative accepted the service agreement.
	UserAgent *string `form:"user_agent"`
}

// Settings specific to the account's Treasury FinancialAccounts.
type AccountSettingsTreasuryParams struct {
	// Details on the account's acceptance of the Stripe Treasury Services Agreement.
	TOSAcceptance *AccountSettingsTreasuryTOSAcceptanceParams `form:"tos_acceptance"`
}

// Options for customizing how the account functions within Stripe.
type AccountSettingsParams struct {
	BACSDebitPayments *AccountSettingsBACSDebitPaymentsParams `form:"bacs_debit_payments"`
	// Settings used to apply the account's branding to email receipts, invoices, Checkout, and other products.
	Branding *AccountSettingsBrandingParams `form:"branding"`
	// Settings specific to the account's use of the Card Issuing product.
	CardIssuing *AccountSettingsCardIssuingParams `form:"card_issuing"`
	// Settings specific to card charging on the account.
	CardPayments *AccountSettingsCardPaymentsParams `form:"card_payments"`
	Dashboard    *AccountSettingsDashboardParams    `form:"dashboard"`
	// Settings that apply across payment methods for charging on the account.
	Payments *AccountSettingsPaymentsParams `form:"payments"`
	// Settings specific to the account's payouts.
	Payouts *AccountSettingsPayoutsParams `form:"payouts"`
	// Settings specific to the account's Treasury FinancialAccounts.
	Treasury *AccountSettingsTreasuryParams `form:"treasury"`
}

// Returns a list of accounts connected to your platform via [Connect](https://stripe.com/docs/connect). If you're not a platform, the list is empty.
type AccountListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// With [Connect](https://stripe.com/docs/connect), you may flag accounts as suspicious.
//
// Test-mode Custom and Express accounts can be rejected at any time. Accounts created using live-mode keys may only be rejected once all balances are zero.
type AccountRejectParams struct {
	Params `form:"*"`
	// The reason for rejecting the account. Can be `fraud`, `terms_of_service`, or `other`.
	Reason *string `form:"reason"`
}

// AccountExternalAccountParams are the parameters allowed to reference an
// external account when creating an account. It should either have Token set
// or everything else.
type AccountExternalAccountParams struct {
	Params            `form:"*"`
	AccountNumber     *string `form:"account_number"`
	AccountHolderName *string `form:"account_holder_name"`
	AccountHolderType *string `form:"account_holder_type"`
	Country           *string `form:"country"`
	Currency          *string `form:"currency"`
	RoutingNumber     *string `form:"routing_number"`
	Token             *string `form:"token"`
}

// AppendTo implements custom encoding logic for AccountExternalAccountParams
// so that we can send the special required `object` field up along with the
// other specified parameters or the token value.
func (p *AccountExternalAccountParams) AppendTo(body *form.Values, keyParts []string) {
	if p.Token != nil {
		body.Add(form.FormatKey(keyParts), StringValue(p.Token))
	} else {
		body.Add(form.FormatKey(append(keyParts, "object")), "bank_account")
	}
}

// Business information about the account.
type AccountBusinessProfile struct {
	// [The merchant category code for the account](https://stripe.com/docs/connect/setting-mcc). MCCs are used to classify businesses based on the goods or services they provide.
	MCC string `json:"mcc"`
	// The customer-facing business name.
	Name string `json:"name"`
	// Internal-only description of the product sold or service provided by the business. It's used by Stripe for risk and underwriting purposes.
	ProductDescription string `json:"product_description"`
	// A publicly available mailing address for sending support issues to.
	SupportAddress *Address `json:"support_address"`
	// A publicly available email address for sending support issues to.
	SupportEmail string `json:"support_email"`
	// A publicly available phone number to call with support issues.
	SupportPhone string `json:"support_phone"`
	// A publicly available website for handling support issues.
	SupportURL string `json:"support_url"`
	// The business's publicly available website.
	URL string `json:"url"`
}
type AccountCapabilities struct {
	// The status of the Canadian pre-authorized debits payments capability of the account, or whether the account can directly process Canadian pre-authorized debits charges.
	ACSSDebitPayments AccountCapabilityStatus `json:"acss_debit_payments"`
	// The status of the Affirm capability of the account, or whether the account can directly process Affirm charges.
	AffirmPayments AccountCapabilitiesAffirmPayments `json:"affirm_payments"`
	// The status of the Afterpay Clearpay capability of the account, or whether the account can directly process Afterpay Clearpay charges.
	AfterpayClearpayPayments AccountCapabilityStatus `json:"afterpay_clearpay_payments"`
	// The status of the BECS Direct Debit (AU) payments capability of the account, or whether the account can directly process BECS Direct Debit (AU) charges.
	AUBECSDebitPayments AccountCapabilityStatus `json:"au_becs_debit_payments"`
	// The status of the Bacs Direct Debits payments capability of the account, or whether the account can directly process Bacs Direct Debits charges.
	BACSDebitPayments AccountCapabilityStatus `json:"bacs_debit_payments"`
	// The status of the Bancontact payments capability of the account, or whether the account can directly process Bancontact charges.
	BancontactPayments AccountCapabilityStatus `json:"bancontact_payments"`
	// The status of the customer_balance payments capability of the account, or whether the account can directly process customer_balance charges.
	BankTransferPayments AccountCapabilitiesBankTransferPayments `json:"bank_transfer_payments"`
	// The status of the blik payments capability of the account, or whether the account can directly process blik charges.
	BLIKPayments AccountCapabilitiesBLIKPayments `json:"blik_payments"`
	// The status of the boleto payments capability of the account, or whether the account can directly process boleto charges.
	BoletoPayments AccountCapabilityStatus `json:"boleto_payments"`
	// The status of the card issuing capability of the account, or whether you can use Issuing to distribute funds on cards
	CardIssuing AccountCapabilityStatus `json:"card_issuing"`
	// The status of the card payments capability of the account, or whether the account can directly process credit and debit card charges.
	CardPayments AccountCapabilityStatus `json:"card_payments"`
	// The status of the Cartes Bancaires payments capability of the account, or whether the account can directly process Cartes Bancaires card charges in EUR currency.
	CartesBancairesPayments AccountCapabilityStatus `json:"cartes_bancaires_payments"`
	// The status of the EPS payments capability of the account, or whether the account can directly process EPS charges.
	EPSPayments AccountCapabilityStatus `json:"eps_payments"`
	// The status of the FPX payments capability of the account, or whether the account can directly process FPX charges.
	FPXPayments AccountCapabilityStatus `json:"fpx_payments"`
	// The status of the giropay payments capability of the account, or whether the account can directly process giropay charges.
	GiropayPayments AccountCapabilityStatus `json:"giropay_payments"`
	// The status of the GrabPay payments capability of the account, or whether the account can directly process GrabPay charges.
	GrabpayPayments AccountCapabilityStatus `json:"grabpay_payments"`
	// The status of the iDEAL payments capability of the account, or whether the account can directly process iDEAL charges.
	IdealPayments AccountCapabilityStatus `json:"ideal_payments"`
	// The status of the JCB payments capability of the account, or whether the account (Japan only) can directly process JCB credit card charges in JPY currency.
	JCBPayments AccountCapabilityStatus `json:"jcb_payments"`
	// The status of the Klarna payments capability of the account, or whether the account can directly process Klarna charges.
	KlarnaPayments AccountCapabilityStatus `json:"klarna_payments"`
	// The status of the konbini payments capability of the account, or whether the account can directly process konbini charges.
	KonbiniPayments AccountCapabilitiesKonbiniPayments `json:"konbini_payments"`
	// The status of the legacy payments capability of the account.
	LegacyPayments AccountCapabilityStatus `json:"legacy_payments"`
	// The status of the link_payments capability of the account, or whether the account can directly process Link charges.
	LinkPayments AccountCapabilitiesLinkPayments `json:"link_payments"`
	// The status of the OXXO payments capability of the account, or whether the account can directly process OXXO charges.
	OXXOPayments AccountCapabilityStatus `json:"oxxo_payments"`
	// The status of the P24 payments capability of the account, or whether the account can directly process P24 charges.
	P24Payments AccountCapabilityStatus `json:"p24_payments"`
	// The status of the paynow payments capability of the account, or whether the account can directly process paynow charges.
	PayNowPayments AccountCapabilitiesPayNowPayments `json:"paynow_payments"`
	// The status of the promptpay payments capability of the account, or whether the account can directly process promptpay charges.
	PromptPayPayments AccountCapabilitiesPromptPayPayments `json:"promptpay_payments"`
	// The status of the SEPA Direct Debits payments capability of the account, or whether the account can directly process SEPA Direct Debits charges.
	SEPADebitPayments AccountCapabilityStatus `json:"sepa_debit_payments"`
	// The status of the Sofort payments capability of the account, or whether the account can directly process Sofort charges.
	SofortPayments AccountCapabilityStatus `json:"sofort_payments"`
	// The status of the tax reporting 1099-K (US) capability of the account.
	TaxReportingUS1099K AccountCapabilityStatus `json:"tax_reporting_us_1099_k"`
	// The status of the tax reporting 1099-MISC (US) capability of the account.
	TaxReportingUS1099MISC AccountCapabilityStatus `json:"tax_reporting_us_1099_misc"`
	// The status of the transfers capability of the account, or whether your platform can transfer funds to the account.
	Transfers AccountCapabilityStatus `json:"transfers"`
	// The status of the banking capability, or whether the account can have bank accounts.
	Treasury AccountCapabilitiesTreasury `json:"treasury"`
	// The status of the US bank account ACH payments capability of the account, or whether the account can directly process US bank account charges.
	USBankAccountAchPayments AccountCapabilitiesUSBankAccountAchPayments `json:"us_bank_account_ach_payments"`
}

// The Kana variation of the company's primary address (Japan only).
type AccountAddress struct {
	// City/Ward.
	City string `json:"city"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// Block/Building number.
	Line1 string `json:"line1"`
	// Building details.
	Line2 string `json:"line2"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code"`
	// Prefecture.
	State string `json:"state"`
	// Town/cho-me. Note that this is only used for Kana/Kanji representations
	// of an address.
	// Town/cho-me.
	Town string `json:"town"`
}

// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
type AccountCompanyOwnershipDeclaration struct {
	// The Unix timestamp marking when the beneficial owner attestation was made.
	Date int64 `json:"date"`
	// The IP address from which the beneficial owner attestation was made.
	IP string `json:"ip"`
	// The user-agent string from the browser where the beneficial owner attestation was made.
	UserAgent string `json:"user_agent"`
}
type AccountCompanyVerificationDocument struct {
	// The back of a document returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `additional_verification`.
	Back *File `json:"back"`
	// A user-displayable string describing the verification state of this document.
	Details string `json:"details"`
	// One of `document_corrupt`, `document_expired`, `document_failed_copy`, `document_failed_greyscale`, `document_failed_other`, `document_failed_test_mode`, `document_fraudulent`, `document_incomplete`, `document_invalid`, `document_manipulated`, `document_not_readable`, `document_not_uploaded`, `document_type_not_supported`, or `document_too_large`. A machine-readable code specifying the verification state for this document.
	DetailsCode AccountCompanyVerificationDocumentDetailsCode `json:"details_code"`
	// The front of a document returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `additional_verification`.
	Front *File `json:"front"`
}

// Information on the verification state of the company.
type AccountCompanyVerification struct {
	Document *AccountCompanyVerificationDocument `json:"document"`
}
type AccountCompany struct {
	Address *AccountAddress `json:"address"`
	// The Kana variation of the company's primary address (Japan only).
	AddressKana *AccountAddress `json:"address_kana"`
	// The Kanji variation of the company's primary address (Japan only).
	AddressKanji *AccountAddress `json:"address_kanji"`
	// Whether the company's directors have been provided. This Boolean will be `true` if you've manually indicated that all directors are provided via [the `directors_provided` parameter](https://stripe.com/docs/api/accounts/update#update_account-company-directors_provided).
	DirectorsProvided bool `json:"directors_provided"`
	// Whether the company's executives have been provided. This Boolean will be `true` if you've manually indicated that all executives are provided via [the `executives_provided` parameter](https://stripe.com/docs/api/accounts/update#update_account-company-executives_provided), or if Stripe determined that sufficient executives were provided.
	ExecutivesProvided bool `json:"executives_provided"`
	// The company's legal name.
	Name string `json:"name"`
	// The Kana variation of the company's legal name (Japan only).
	NameKana string `json:"name_kana"`
	// The Kanji variation of the company's legal name (Japan only).
	NameKanji string `json:"name_kanji"`
	// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
	OwnershipDeclaration *AccountCompanyOwnershipDeclaration `json:"ownership_declaration"`
	// Whether the company's owners have been provided. This Boolean will be `true` if you've manually indicated that all owners are provided via [the `owners_provided` parameter](https://stripe.com/docs/api/accounts/update#update_account-company-owners_provided), or if Stripe determined that sufficient owners were provided. Stripe determines ownership requirements using both the number of owners provided and their total percent ownership (calculated by adding the `percent_ownership` of each owner together).
	OwnersProvided bool `json:"owners_provided"`
	// The company's phone number (used for verification).
	Phone              string `json:"phone"`
	RegistrationNumber string `json:"registration_number"`
	// The category identifying the legal structure of the company or legal entity. See [Business structure](https://stripe.com/docs/connect/identity-verification#business-structure) for more details.
	Structure AccountCompanyStructure `json:"structure"`
	// Whether the company's business ID number was provided.
	TaxIDProvided bool `json:"tax_id_provided"`
	// The jurisdiction in which the `tax_id` is registered (Germany-based companies only).
	TaxIDRegistrar string `json:"tax_id_registrar"`
	// Whether the company's business VAT number was provided.
	VATIDProvided bool `json:"vat_id_provided"`
	// Information on the verification state of the company.
	Verification *AccountCompanyVerification `json:"verification"`
}
type AccountController struct {
	// `true` if the Connect application retrieving the resource controls the account and can therefore exercise [platform controls](https://stripe.com/docs/connect/platform-controls-for-standard-accounts). Otherwise, this field is null.
	IsController bool `json:"is_controller"`
	// The controller type. Can be `application`, if a Connect application controls the account, or `account`, if the account controls itself.
	Type AccountControllerType `json:"type"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type AccountFutureRequirementsAlternative struct {
	// Fields that can be provided to satisfy all fields in `original_fields_due`.
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	// Fields that are due and can be satisfied by providing all fields in `alternative_fields_due`.
	OriginalFieldsDue []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type AccountFutureRequirementsError struct {
	// The code for the type of error.
	Code string `json:"code"`
	// An informative message that indicates the error type and provides additional details about the error.
	Reason string `json:"reason"`
	// The specific user onboarding requirement field (in the requirements hash) that needs to be resolved.
	Requirement string `json:"requirement"`
}
type AccountFutureRequirements struct {
	// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
	Alternatives []*AccountFutureRequirementsAlternative `json:"alternatives"`
	// Date on which `future_requirements` merges with the main `requirements` hash and `future_requirements` becomes empty. After the transition, `currently_due` requirements may immediately become `past_due`, but the account may also be given a grace period depending on its enablement state prior to transitioning.
	CurrentDeadline int64 `json:"current_deadline"`
	// Fields that need to be collected to keep the account enabled. If not collected by `future_requirements[current_deadline]`, these fields will transition to the main `requirements` hash.
	CurrentlyDue []string `json:"currently_due"`
	// This is typed as a string for consistency with `requirements.disabled_reason`, but it safe to assume `future_requirements.disabled_reason` is empty because fields in `future_requirements` will never disable the account.
	DisabledReason string `json:"disabled_reason"`
	// Fields that are `currently_due` and need to be collected again because validation or verification failed.
	Errors []*AccountFutureRequirementsError `json:"errors"`
	// Fields that need to be collected assuming all volume thresholds are reached. As they become required, they appear in `currently_due` as well.
	EventuallyDue []string `json:"eventually_due"`
	// Fields that weren't collected by `requirements.current_deadline`. These fields need to be collected to enable the capability on the account. New fields will never appear here; `future_requirements.past_due` will always be a subset of `requirements.past_due`.
	PastDue []string `json:"past_due"`
	// Fields that may become required depending on the results of verification or review. Will be an empty array unless an asynchronous verification is pending. If verification fails, these fields move to `eventually_due` or `currently_due`.
	PendingVerification []string `json:"pending_verification"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type AccountRequirementsAlternative struct {
	// Fields that can be provided to satisfy all fields in `original_fields_due`.
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	// Fields that are due and can be satisfied by providing all fields in `alternative_fields_due`.
	OriginalFieldsDue []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type AccountRequirementsError struct {
	// The code for the type of error.
	Code string `json:"code"`
	// An informative message that indicates the error type and provides additional details about the error.
	Reason string `json:"reason"`
	// The specific user onboarding requirement field (in the requirements hash) that needs to be resolved.
	Requirement string `json:"requirement"`
}
type AccountRequirements struct {
	// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
	Alternatives []*AccountRequirementsAlternative `json:"alternatives"`
	// Date by which the fields in `currently_due` must be collected to keep the account enabled. These fields may disable the account sooner if the next threshold is reached before they are collected.
	CurrentDeadline int64 `json:"current_deadline"`
	// Fields that need to be collected to keep the account enabled. If not collected by `current_deadline`, these fields appear in `past_due` as well, and the account is disabled.
	CurrentlyDue []string `json:"currently_due"`
	// If the account is disabled, this string describes why. Can be `requirements.past_due`, `requirements.pending_verification`, `listed`, `platform_paused`, `rejected.fraud`, `rejected.listed`, `rejected.terms_of_service`, `rejected.other`, `under_review`, or `other`.
	DisabledReason AccountRequirementsDisabledReason `json:"disabled_reason"`
	// Fields that are `currently_due` and need to be collected again because validation or verification failed.
	Errors []*AccountRequirementsError `json:"errors"`
	// Fields that need to be collected assuming all volume thresholds are reached. As they become required, they appear in `currently_due` as well, and `current_deadline` becomes set.
	EventuallyDue []string `json:"eventually_due"`
	// Fields that weren't collected by `current_deadline`. These fields need to be collected to enable the account.
	PastDue []string `json:"past_due"`
	// Fields that may become required depending on the results of verification or review. Will be an empty array unless an asynchronous verification is pending. If verification fails, these fields move to `eventually_due`, `currently_due`, or `past_due`.
	PendingVerification []string `json:"pending_verification"`
}
type AccountSettingsBACSDebitPayments struct {
	// The Bacs Direct Debit Display Name for this account. For payments made with Bacs Direct Debit, this will appear on the mandate, and as the statement descriptor.
	DisplayName string `json:"display_name"`
}
type AccountSettingsBranding struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) An icon for the account. Must be square and at least 128px x 128px.
	Icon *File `json:"icon"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) A logo for the account that will be used in Checkout instead of the icon and without the account's name next to it if provided. Must be at least 128px x 128px.
	Logo *File `json:"logo"`
	// A CSS hex color value representing the primary branding color for this account
	PrimaryColor string `json:"primary_color"`
	// A CSS hex color value representing the secondary branding color for this account
	SecondaryColor string `json:"secondary_color"`
}
type AccountTOSAcceptance struct {
	// The Unix timestamp marking when the account representative accepted their service agreement
	Date int64 `json:"date"`
	// The IP address from which the account representative accepted their service agreement
	IP string `json:"ip"`
	// The user's service agreement type
	ServiceAgreement AccountTOSAcceptanceServiceAgreement `json:"service_agreement"`
	// The user agent of the browser from which the account representative accepted their service agreement
	UserAgent string `json:"user_agent"`
}
type AccountSettingsCardIssuing struct {
	TOSAcceptance *AccountTOSAcceptance `json:"tos_acceptance"`
}
type AccountDeclineOn struct {
	// Whether Stripe automatically declines charges with an incorrect ZIP or postal code. This setting only applies when a ZIP or postal code is provided and they fail bank verification.
	AVSFailure bool `json:"avs_failure"`
	// Whether Stripe automatically declines charges with an incorrect CVC. This setting only applies when a CVC is provided and it fails bank verification.
	CVCFailure bool `json:"cvc_failure"`
}
type AccountSettingsCardPayments struct {
	DeclineOn *AccountDeclineOn `json:"decline_on"`
	// The default text that appears on credit card statements when a charge is made. This field prefixes any dynamic `statement_descriptor` specified on the charge. `statement_descriptor_prefix` is useful for maximizing descriptor space for the dynamic portion.
	StatementDescriptorPrefix string `json:"statement_descriptor_prefix"`
	// The Kana variation of the default text that appears on credit card statements when a charge is made (Japan only). This field prefixes any dynamic `statement_descriptor_suffix_kana` specified on the charge. `statement_descriptor_prefix_kana` is useful for maximizing descriptor space for the dynamic portion.
	StatementDescriptorPrefixKana string `json:"statement_descriptor_prefix_kana"`
	// The Kanji variation of the default text that appears on credit card statements when a charge is made (Japan only). This field prefixes any dynamic `statement_descriptor_suffix_kanji` specified on the charge. `statement_descriptor_prefix_kanji` is useful for maximizing descriptor space for the dynamic portion.
	StatementDescriptorPrefixKanji string `json:"statement_descriptor_prefix_kanji"`
}
type AccountSettingsDashboard struct {
	// The display name for this account. This is used on the Stripe Dashboard to differentiate between accounts.
	DisplayName string `json:"display_name"`
	// The timezone used in the Stripe Dashboard for this account. A list of possible time zone values is maintained at the [IANA Time Zone Database](http://www.iana.org/time-zones).
	Timezone string `json:"timezone"`
}
type AccountSettingsPayments struct {
	// The default text that appears on credit card statements when a charge is made. This field prefixes any dynamic `statement_descriptor` specified on the charge.
	StatementDescriptor string `json:"statement_descriptor"`
	// The Kana variation of the default text that appears on credit card statements when a charge is made (Japan only)
	StatementDescriptorKana string `json:"statement_descriptor_kana"`
	// The Kanji variation of the default text that appears on credit card statements when a charge is made (Japan only)
	StatementDescriptorKanji string `json:"statement_descriptor_kanji"`
	// The Kana variation of the default text that appears on credit card statements when a charge is made (Japan only). This field prefixes any dynamic `statement_descriptor_suffix_kana` specified on the charge. `statement_descriptor_prefix_kana` is useful for maximizing descriptor space for the dynamic portion.
	StatementDescriptorPrefixKana string `json:"statement_descriptor_prefix_kana"`
	// The Kanji variation of the default text that appears on credit card statements when a charge is made (Japan only). This field prefixes any dynamic `statement_descriptor_suffix_kanji` specified on the charge. `statement_descriptor_prefix_kanji` is useful for maximizing descriptor space for the dynamic portion.
	StatementDescriptorPrefixKanji string `json:"statement_descriptor_prefix_kanji"`
}
type AccountPayoutSchedule struct {
	// The number of days charges for the account will be held before being paid out.
	DelayDays int64 `json:"delay_days"`
	// How frequently funds will be paid out. One of `manual` (payouts only created via API call), `daily`, `weekly`, or `monthly`.
	Interval PayoutInterval `json:"interval"`
	// The day of the month funds will be paid out. Only shown if `interval` is monthly. Payouts scheduled between the 29th and 31st of the month are sent on the last day of shorter months.
	MonthlyAnchor int64 `json:"monthly_anchor"`
	// The day of the week funds will be paid out, of the style 'monday', 'tuesday', etc. Only shown if `interval` is weekly.
	WeeklyAnchor string `json:"weekly_anchor"`
}
type AccountSettingsPayouts struct {
	// A Boolean indicating if Stripe should try to reclaim negative balances from an attached bank account. See our [Understanding Connect Account Balances](https://stripe.com/docs/connect/account-balances) documentation for details. Default value is `false` for Custom accounts, otherwise `true`.
	DebitNegativeBalances bool                   `json:"debit_negative_balances"`
	Schedule              *AccountPayoutSchedule `json:"schedule"`
	// The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.
	StatementDescriptor string `json:"statement_descriptor"`
}
type AccountSettingsSEPADebitPayments struct {
	// SEPA creditor identifier that identifies the company making the payment.
	CreditorID string `json:"creditor_id"`
}
type AccountSettingsTreasuryTOSAcceptance struct {
	// The Unix timestamp marking when the account representative accepted the service agreement.
	Date int64 `json:"date"`
	// The IP address from which the account representative accepted the service agreement.
	IP string `json:"ip"`
	// The user agent of the browser from which the account representative accepted the service agreement.
	UserAgent string `json:"user_agent"`
}
type AccountSettingsTreasury struct {
	TOSAcceptance *AccountSettingsTreasuryTOSAcceptance `json:"tos_acceptance"`
}

// Options for customizing how the account functions within Stripe.
type AccountSettings struct {
	BACSDebitPayments *AccountSettingsBACSDebitPayments `json:"bacs_debit_payments"`
	Branding          *AccountSettingsBranding          `json:"branding"`
	CardIssuing       *AccountSettingsCardIssuing       `json:"card_issuing"`
	CardPayments      *AccountSettingsCardPayments      `json:"card_payments"`
	Dashboard         *AccountSettingsDashboard         `json:"dashboard"`
	Payments          *AccountSettingsPayments          `json:"payments"`
	Payouts           *AccountSettingsPayouts           `json:"payouts"`
	SEPADebitPayments *AccountSettingsSEPADebitPayments `json:"sepa_debit_payments"`
	Treasury          *AccountSettingsTreasury          `json:"treasury"`
}

// This is an object representing a Stripe account. You can retrieve it to see
// properties on the account like its current e-mail address or if the account is
// enabled yet to make live charges.
//
// Some properties, marked below, are available only to platforms that want to
// [create and manage Express or Custom accounts](https://stripe.com/docs/connect/accounts).
type Account struct {
	APIResource
	// Business information about the account.
	BusinessProfile *AccountBusinessProfile `json:"business_profile"`
	// The business type.
	BusinessType AccountBusinessType  `json:"business_type"`
	Capabilities *AccountCapabilities `json:"capabilities"`
	// Whether the account can create live charges.
	ChargesEnabled bool               `json:"charges_enabled"`
	Company        *AccountCompany    `json:"company"`
	Controller     *AccountController `json:"controller"`
	// The account's country.
	Country string `json:"country"`
	// Time at which the account was connected. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter ISO currency code representing the default currency for the account. This must be a currency that [Stripe supports in the account's country](https://stripe.com/docs/payouts).
	DefaultCurrency Currency `json:"default_currency"`
	Deleted         bool     `json:"deleted"`
	// Whether account details have been submitted. Standard accounts cannot receive payouts before this is true.
	DetailsSubmitted bool `json:"details_submitted"`
	// An email address associated with the account. You can treat this as metadata: it is not used for authentication or messaging account holders.
	Email string `json:"email"`
	// External accounts (bank accounts and debit cards) currently attached to this account
	ExternalAccounts   *ExternalAccountList       `json:"external_accounts"`
	FutureRequirements *AccountFutureRequirements `json:"future_requirements"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// This is an object representing a person associated with a Stripe account.
	//
	// A platform cannot access a Standard or Express account's persons after the account starts onboarding, such as after generating an account link for the account.
	// See the [Standard onboarding](https://stripe.com/docs/connect/standard-accounts) or [Express onboarding documentation](https://stripe.com/docs/connect/express-accounts) for information about platform pre-filling and account onboarding steps.
	//
	// Related guide: [Handling Identity Verification with the API](https://stripe.com/docs/connect/identity-verification-api#person-information).
	Individual *Person `json:"individual"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Whether Stripe can send payouts to this account.
	PayoutsEnabled bool                 `json:"payouts_enabled"`
	Requirements   *AccountRequirements `json:"requirements"`
	// Options for customizing how the account functions within Stripe.
	Settings      *AccountSettings      `json:"settings"`
	TOSAcceptance *AccountTOSAcceptance `json:"tos_acceptance"`
	// The Stripe account type. Can be `standard`, `express`, or `custom`.
	Type AccountType `json:"type"`
}
type ExternalAccount struct {
	ID   string              `json:"id"`
	Type ExternalAccountType `json:"object"`

	// BankAccount is a bank account attached to an account. Populated only if
	// the external account is a bank account.
	BankAccount *BankAccount `json:"-"`
	// Card is a card attached to an account. Populated only if the external
	// account is a card.
	Card *Card `json:"-"`
}

// AccountList is a list of Accounts as retrieved from a list endpoint.
type AccountList struct {
	APIResource
	ListMeta
	Data []*Account `json:"data"`
}

// ExternalAccountList is a list of external accounts that may be either bank
// accounts or cards.
type ExternalAccountList struct {
	APIResource
	ListMeta

	// Values contains any external accounts (bank accounts and/or cards)
	// currently attached to this account.
	Data []*ExternalAccount `json:"data"`
}

// UnmarshalJSON handles deserialization of an Account.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (a *Account) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		a.ID = id
		return nil
	}

	type account Account
	var v account
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*a = Account(v)
	return nil
}

// UnmarshalJSON handles deserialization of an ExternalAccount.
// This custom unmarshaling is needed because the specific type of
// ExternalAccount it refers to is specified in the JSON
func (e *ExternalAccount) UnmarshalJSON(data []byte) error {
	type externalAccount ExternalAccount
	var v externalAccount
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*e = ExternalAccount(v)
	var err error

	switch e.Type {
	case ExternalAccountTypeBankAccount:
		err = json.Unmarshal(data, &e.BankAccount)
	case ExternalAccountTypeCard:
		err = json.Unmarshal(data, &e.Card)
	}
	return err
}
