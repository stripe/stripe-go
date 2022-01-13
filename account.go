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
	Params          `form:"*"`
	AccountToken    *string                       `form:"account_token"`
	BusinessProfile *AccountBusinessProfileParams `form:"business_profile"`
	BusinessType    *string                       `form:"business_type"`
	Capabilities    *AccountCapabilitiesParams    `form:"capabilities"`
	Company         *AccountCompanyParams         `form:"company"`
	Country         *string                       `form:"country"`
	DefaultCurrency *string                       `form:"default_currency"`
	Documents       *AccountDocumentsParams       `form:"documents"`
	Email           *string                       `form:"email"`
	ExternalAccount *AccountExternalAccountParams `form:"external_account"`
	Individual      *PersonParams                 `form:"individual"`
	Settings        *AccountSettingsParams        `form:"settings"`
	TOSAcceptance   *AccountTOSAcceptanceParams   `form:"tos_acceptance"`
	Type            *string                       `form:"type"`
	// This parameter is deprecated. Prefer using Capabilities instead.
	RequestedCapabilities []*string `form:"requested_capabilities"`
}

// Business information about the account.
type AccountBusinessProfileParams struct {
	MCC                *string        `form:"mcc"`
	Name               *string        `form:"name"`
	ProductDescription *string        `form:"product_description"`
	SupportAddress     *AddressParams `form:"support_address"`
	SupportEmail       *string        `form:"support_email"`
	SupportPhone       *string        `form:"support_phone"`
	SupportURL         *string        `form:"support_url"`
	URL                *string        `form:"url"`
}

// The acss_debit_payments capability.
type AccountCapabilitiesACSSDebitPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The afterpay_clearpay_payments capability.
type AccountCapabilitiesAfterpayClearpayPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The au_becs_debit_payments capability.
type AccountCapabilitiesAUBECSDebitPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The bacs_debit_payments capability.
type AccountCapabilitiesBACSDebitPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The bancontact_payments capability.
type AccountCapabilitiesBancontactPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The boleto_payments capability.
type AccountCapabilitiesBoletoPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The card_issuing capability.
type AccountCapabilitiesCardIssuingParams struct {
	Requested *bool `form:"requested"`
}

// The card_payments capability.
type AccountCapabilitiesCardPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The cartes_bancaires_payments capability.
type AccountCapabilitiesCartesBancairesPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The eps_payments capability.
type AccountCapabilitiesEPSPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The fpx_payments capability.
type AccountCapabilitiesFPXPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The giropay_payments capability.
type AccountCapabilitiesGiropayPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The grabpay_payments capability.
type AccountCapabilitiesGrabpayPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The ideal_payments capability.
type AccountCapabilitiesIdealPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The jcb_payments capability.
type AccountCapabilitiesJCBPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The klarna_payments capability.
type AccountCapabilitiesKlarnaPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The legacy_payments capability.
type AccountCapabilitiesLegacyPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The oxxo_payments capability.
type AccountCapabilitiesOXXOPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The p24_payments capability.
type AccountCapabilitiesP24PaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The sepa_debit_payments capability.
type AccountCapabilitiesSEPADebitPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The sofort_payments capability.
type AccountCapabilitiesSofortPaymentsParams struct {
	Requested *bool `form:"requested"`
}

// The tax_reporting_us_1099_k capability.
type AccountCapabilitiesTaxReportingUS1099KParams struct {
	Requested *bool `form:"requested"`
}

// The tax_reporting_us_1099_misc capability.
type AccountCapabilitiesTaxReportingUS1099MISCParams struct {
	Requested *bool `form:"requested"`
}

// The transfers capability.
type AccountCapabilitiesTransfersParams struct {
	Requested *bool `form:"requested"`
}

// Each key of the dictionary represents a capability, and each capability maps to its settings (e.g. whether it has been requested or not). Each capability will be inactive until you have provided its specific requirements and Stripe has verified them. An account may have some of its requested capabilities be active and some be inactive.
type AccountCapabilitiesParams struct {
	ACSSDebitPayments        *AccountCapabilitiesACSSDebitPaymentsParams        `form:"acss_debit_payments"`
	AfterpayClearpayPayments *AccountCapabilitiesAfterpayClearpayPaymentsParams `form:"afterpay_clearpay_payments"`
	AUBECSDebitPayments      *AccountCapabilitiesAUBECSDebitPaymentsParams      `form:"au_becs_debit_payments"`
	BACSDebitPayments        *AccountCapabilitiesBACSDebitPaymentsParams        `form:"bacs_debit_payments"`
	BancontactPayments       *AccountCapabilitiesBancontactPaymentsParams       `form:"bancontact_payments"`
	BoletoPayments           *AccountCapabilitiesBoletoPaymentsParams           `form:"boleto_payments"`
	CardIssuing              *AccountCapabilitiesCardIssuingParams              `form:"card_issuing"`
	CardPayments             *AccountCapabilitiesCardPaymentsParams             `form:"card_payments"`
	CartesBancairesPayments  *AccountCapabilitiesCartesBancairesPaymentsParams  `form:"cartes_bancaires_payments"`
	EPSPayments              *AccountCapabilitiesEPSPaymentsParams              `form:"eps_payments"`
	FPXPayments              *AccountCapabilitiesFPXPaymentsParams              `form:"fpx_payments"`
	GiropayPayments          *AccountCapabilitiesGiropayPaymentsParams          `form:"giropay_payments"`
	GrabpayPayments          *AccountCapabilitiesGrabpayPaymentsParams          `form:"grabpay_payments"`
	IdealPayments            *AccountCapabilitiesIdealPaymentsParams            `form:"ideal_payments"`
	JCBPayments              *AccountCapabilitiesJCBPaymentsParams              `form:"jcb_payments"`
	KlarnaPayments           *AccountCapabilitiesKlarnaPaymentsParams           `form:"klarna_payments"`
	LegacyPayments           *AccountCapabilitiesLegacyPaymentsParams           `form:"legacy_payments"`
	OXXOPayments             *AccountCapabilitiesOXXOPaymentsParams             `form:"oxxo_payments"`
	P24Payments              *AccountCapabilitiesP24PaymentsParams              `form:"p24_payments"`
	SEPADebitPayments        *AccountCapabilitiesSEPADebitPaymentsParams        `form:"sepa_debit_payments"`
	SofortPayments           *AccountCapabilitiesSofortPaymentsParams           `form:"sofort_payments"`
	TaxReportingUS1099K      *AccountCapabilitiesTaxReportingUS1099KParams      `form:"tax_reporting_us_1099_k"`
	TaxReportingUS1099MISC   *AccountCapabilitiesTaxReportingUS1099MISCParams   `form:"tax_reporting_us_1099_misc"`
	Transfers                *AccountCapabilitiesTransfersParams                `form:"transfers"`
}

// The Kana variation of the company's primary address (Japan only).
type AccountAddressParams struct {
	City       *string `form:"city"`
	Country    *string `form:"country"`
	Line1      *string `form:"line1"`
	Line2      *string `form:"line2"`
	PostalCode *string `form:"postal_code"`
	State      *string `form:"state"`
	// Town/cho-me. Note that this is only used for Kana/Kanji representations
	// of an address.
	Town *string `form:"town"`
}

// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
type AccountCompanyOwnershipDeclarationParams struct {
	Date      *int64  `form:"date"`
	IP        *string `form:"ip"`
	UserAgent *string `form:"user_agent"`
}

// A document verifying the business.
type AccountCompanyVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// Information on the verification state of the company.
type AccountCompanyVerificationParams struct {
	Document *AccountCompanyVerificationDocumentParams `form:"document"`
}

// Information about the company or business. This field is available for any `business_type`.
type AccountCompanyParams struct {
	Address              *AccountAddressParams                     `form:"address"`
	AddressKana          *AccountAddressParams                     `form:"address_kana"`
	AddressKanji         *AccountAddressParams                     `form:"address_kanji"`
	DirectorsProvided    *bool                                     `form:"directors_provided"`
	ExecutivesProvided   *bool                                     `form:"executives_provided"`
	Name                 *string                                   `form:"name"`
	NameKana             *string                                   `form:"name_kana"`
	NameKanji            *string                                   `form:"name_kanji"`
	OwnershipDeclaration *AccountCompanyOwnershipDeclarationParams `form:"ownership_declaration"`
	// This parameter can only be used on Token creation.
	OwnershipDeclarationShownAndSigned *bool                             `form:"ownership_declaration_shown_and_signed"`
	OwnersProvided                     *bool                             `form:"owners_provided"`
	Phone                              *string                           `form:"phone"`
	RegistrationNumber                 *string                           `form:"registration_number"`
	Structure                          *string                           `form:"structure"`
	TaxID                              *string                           `form:"tax_id"`
	TaxIDRegistrar                     *string                           `form:"tax_id_registrar"`
	VATID                              *string                           `form:"vat_id"`
	Verification                       *AccountCompanyVerificationParams `form:"verification"`
}

// One or more documents that support the [Bank account ownership verification](https://support.stripe.com/questions/bank-account-ownership-verification) requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a voided check.
type AccountDocumentsBankAccountOwnershipVerificationParams struct {
	Files []*string `form:"files"`
}

// One or more documents that demonstrate proof of a company's license to operate.
type AccountDocumentsCompanyLicenseParams struct {
	Files []*string `form:"files"`
}

// One or more documents showing the company's Memorandum of Association.
type AccountDocumentsCompanyMemorandumOfAssociationParams struct {
	Files []*string `form:"files"`
}

// (Certain countries only) One or more documents showing the ministerial decree legalizing the company's establishment.
type AccountDocumentsCompanyMinisterialDecreeParams struct {
	Files []*string `form:"files"`
}

// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
type AccountDocumentsCompanyRegistrationVerificationParams struct {
	Files []*string `form:"files"`
}

// One or more documents that demonstrate proof of a company's tax ID.
type AccountDocumentsCompanyTaxIDVerificationParams struct {
	Files []*string `form:"files"`
}

// One or more documents showing the company's proof of registration with the national business registry.
type AccountDocumentsProofOfRegistrationParams struct {
	Files []*string `form:"files"`
}

// Documents that may be submitted to satisfy various informational requests.
type AccountDocumentsParams struct {
	BankAccountOwnershipVerification *AccountDocumentsBankAccountOwnershipVerificationParams `form:"bank_account_ownership_verification"`
	CompanyLicense                   *AccountDocumentsCompanyLicenseParams                   `form:"company_license"`
	CompanyMemorandumOfAssocation    *AccountDocumentsCompanyMemorandumOfAssociationParams   `form:"company_memorandum_of_association"`
	CompanyMinisterialDecree         *AccountDocumentsCompanyMinisterialDecreeParams         `form:"company_ministerial_decree"`
	CompanyRegistrationVerification  *AccountDocumentsCompanyRegistrationVerificationParams  `form:"company_registration_verification"`
	CompanyTaxIDVerification         *AccountDocumentsCompanyTaxIDVerificationParams         `form:"company_tax_id_verification"`
	ProofOfRegistration              *AccountDocumentsProofOfRegistrationParams              `form:"proof_of_registration"`
}

// AccountSettingsBACSDebitPaymentsParams represent allowed parameters to configure settings specific to
// BACS Debit charging on the account.
type AccountSettingsBACSDebitPaymentsParams struct {
	DisplayName *string `form:"display_name"`
}

// Settings used to apply the account's branding to email receipts, invoices, Checkout, and other products.
type AccountSettingsBrandingParams struct {
	Icon           *string `form:"icon"`
	Logo           *string `form:"logo"`
	PrimaryColor   *string `form:"primary_color"`
	SecondaryColor *string `form:"secondary_color"`
}

// Details on the account's acceptance of the [Stripe Issuing Terms and Disclosures](https://stripe.com/docs/issuing/connect/tos_acceptance).
type AccountTOSAcceptanceParams struct {
	Date             *int64  `form:"date"`
	IP               *string `form:"ip"`
	ServiceAgreement *string `form:"service_agreement"`
	UserAgent        *string `form:"user_agent"`
}

// Settings specific to the account's use of the Card Issuing product.
type AccountSettingsCardIssuingParams struct {
	TOSAcceptance *AccountTOSAcceptanceParams `form:"tos_acceptance"`
}

// Automatically declines certain charge types regardless of whether the card issuer accepted or declined the charge.
type AccountDeclineSettingsParams struct {
	AVSFailure *bool `form:"avs_failure"`
	CVCFailure *bool `form:"cvc_failure"`
}

// Settings specific to card charging on the account.
type AccountSettingsCardPaymentsParams struct {
	DeclineOn                 *AccountDeclineSettingsParams `form:"decline_on"`
	StatementDescriptorPrefix *string                       `form:"statement_descriptor_prefix"`
}

// AccountSettingsDashboardParams represent allowed parameters to configure settings for the
// account's Dashboard.
type AccountSettingsDashboardParams struct {
	DisplayName *string `form:"display_name"`
	Timezone    *string `form:"timezone"`
}

// Settings that apply across payment methods for charging on the account.
type AccountSettingsPaymentsParams struct {
	StatementDescriptor      *string `form:"statement_descriptor"`
	StatementDescriptorKana  *string `form:"statement_descriptor_kana"`
	StatementDescriptorKanji *string `form:"statement_descriptor_kanji"`
}

// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://stripe.com/docs/connect/bank-transfers#payout-information) documentation.
type PayoutScheduleParams struct {
	DelayDays        *int64  `form:"delay_days"`
	DelayDaysMinimum *bool   `form:"-"` // See custom AppendTo
	Interval         *string `form:"interval"`
	MonthlyAnchor    *int64  `form:"monthly_anchor"`
	WeeklyAnchor     *string `form:"weekly_anchor"`
}

// AppendTo implements custom encoding logic for PayoutScheduleParams.
func (p *PayoutScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.DelayDaysMinimum) {
		body.Add(form.FormatKey(append(keyParts, "delay_days")), "minimum")
	}
}

// Settings specific to the account's payouts.
type AccountSettingsPayoutsParams struct {
	DebitNegativeBalances *bool                 `form:"debit_negative_balances"`
	Schedule              *PayoutScheduleParams `form:"schedule"`
	StatementDescriptor   *string               `form:"statement_descriptor"`
}

// Options for customizing how the account functions within Stripe.
type AccountSettingsParams struct {
	BACSDebitPayments *AccountSettingsBACSDebitPaymentsParams `form:"bacs_debit_payments"`
	Branding          *AccountSettingsBrandingParams          `form:"branding"`
	CardIssuing       *AccountSettingsCardIssuingParams       `form:"card_issuing"`
	CardPayments      *AccountSettingsCardPaymentsParams      `form:"card_payments"`
	Dashboard         *AccountSettingsDashboardParams         `form:"dashboard"`
	Payments          *AccountSettingsPaymentsParams          `form:"payments"`
	Payouts           *AccountSettingsPayoutsParams           `form:"payouts"`
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
	MCC                string   `json:"mcc"`
	Name               string   `json:"name"`
	ProductDescription string   `json:"product_description"`
	SupportAddress     *Address `json:"support_address"`
	SupportEmail       string   `json:"support_email"`
	SupportPhone       string   `json:"support_phone"`
	SupportURL         string   `json:"support_url"`
	URL                string   `json:"url"`
}
type AccountCapabilities struct {
	ACSSDebitPayments        AccountCapabilityStatus `json:"acss_debit_payments"`
	AfterpayClearpayPayments AccountCapabilityStatus `json:"afterpay_clearpay_payments"`
	AUBECSDebitPayments      AccountCapabilityStatus `json:"au_becs_debit_payments"`
	BACSDebitPayments        AccountCapabilityStatus `json:"bacs_debit_payments"`
	BancontactPayments       AccountCapabilityStatus `json:"bancontact_payments"`
	BoletoPayments           AccountCapabilityStatus `json:"boleto_payments"`
	CardIssuing              AccountCapabilityStatus `json:"card_issuing"`
	CardPayments             AccountCapabilityStatus `json:"card_payments"`
	CartesBancairesPayments  AccountCapabilityStatus `json:"cartes_bancaires_payments"`
	EPSPayments              AccountCapabilityStatus `json:"eps_payments"`
	FPXPayments              AccountCapabilityStatus `json:"fpx_payments"`
	GiropayPayments          AccountCapabilityStatus `json:"giropay_payments"`
	GrabpayPayments          AccountCapabilityStatus `json:"grabpay_payments"`
	IdealPayments            AccountCapabilityStatus `json:"ideal_payments"`
	JCBPayments              AccountCapabilityStatus `json:"jcb_payments"`
	KlarnaPayments           AccountCapabilityStatus `json:"klarna_payments"`
	LegacyPayments           AccountCapabilityStatus `json:"legacy_payments"`
	OXXOPayments             AccountCapabilityStatus `json:"oxxo_payments"`
	P24Payments              AccountCapabilityStatus `json:"p24_payments"`
	SEPADebitPayments        AccountCapabilityStatus `json:"sepa_debit_payments"`
	SofortPayments           AccountCapabilityStatus `json:"sofort_payments"`
	TaxReportingUS1099K      AccountCapabilityStatus `json:"tax_reporting_us_1099_k"`
	TaxReportingUS1099MISC   AccountCapabilityStatus `json:"tax_reporting_us_1099_misc"`
	Transfers                AccountCapabilityStatus `json:"transfers"`
}

// The Kana variation of the company's primary address (Japan only).
type AccountAddress struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	// Town/cho-me. Note that this is only used for Kana/Kanji representations
	// of an address.
	Town string `json:"town"`
}

// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
type AccountCompanyOwnershipDeclaration struct {
	Date      int64  `json:"date"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}
type AccountCompanyVerificationDocument struct {
	Back        *File                                         `json:"back"`
	Details     string                                        `json:"details"`
	DetailsCode AccountCompanyVerificationDocumentDetailsCode `json:"details_code"`
	Front       *File                                         `json:"front"`
}

// Information on the verification state of the company.
type AccountCompanyVerification struct {
	Document *AccountCompanyVerificationDocument `json:"document"`
}
type AccountCompany struct {
	Address              *AccountAddress                     `json:"address"`
	AddressKana          *AccountAddress                     `json:"address_kana"`
	AddressKanji         *AccountAddress                     `json:"address_kanji"`
	DirectorsProvided    bool                                `json:"directors_provided"`
	ExecutivesProvided   bool                                `json:"executives_provided"`
	Name                 string                              `json:"name"`
	NameKana             string                              `json:"name_kana"`
	NameKanji            string                              `json:"name_kanji"`
	OwnershipDeclaration *AccountCompanyOwnershipDeclaration `json:"ownership_declaration"`
	OwnersProvided       bool                                `json:"owners_provided"`
	Phone                string                              `json:"phone"`
	RegistrationNumber   string                              `json:"registration_number"`
	Structure            AccountCompanyStructure             `json:"structure"`
	TaxIDProvided        bool                                `json:"tax_id_provided"`
	TaxIDRegistrar       string                              `json:"tax_id_registrar"`
	VATIDProvided        bool                                `json:"vat_id_provided"`
	Verification         *AccountCompanyVerification         `json:"verification"`
}
type AccountController struct {
	IsController bool                  `json:"is_controller"`
	Type         AccountControllerType `json:"type"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type AccountFutureRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type AccountFutureRequirementsError struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Requirement string `json:"requirement"`
}
type AccountFutureRequirements struct {
	Alternatives        []*AccountFutureRequirementsAlternative `json:"alternatives"`
	CurrentDeadline     int64                                   `json:"current_deadline"`
	CurrentlyDue        []string                                `json:"currently_due"`
	DisabledReason      string                                  `json:"disabled_reason"`
	Errors              []*AccountFutureRequirementsError       `json:"errors"`
	EventuallyDue       []string                                `json:"eventually_due"`
	PastDue             []string                                `json:"past_due"`
	PendingVerification []string                                `json:"pending_verification"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type AccountRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type AccountRequirementsError struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Requirement string `json:"requirement"`
}
type AccountRequirements struct {
	Alternatives        []*AccountRequirementsAlternative `json:"alternatives"`
	CurrentDeadline     int64                             `json:"current_deadline"`
	CurrentlyDue        []string                          `json:"currently_due"`
	DisabledReason      AccountRequirementsDisabledReason `json:"disabled_reason"`
	Errors              []*AccountRequirementsError       `json:"errors"`
	EventuallyDue       []string                          `json:"eventually_due"`
	PastDue             []string                          `json:"past_due"`
	PendingVerification []string                          `json:"pending_verification"`
}
type AccountSettingsBACSDebitPayments struct {
	DisplayName string `json:"display_name"`
}
type AccountSettingsBranding struct {
	Icon           *File  `json:"icon"`
	Logo           *File  `json:"logo"`
	PrimaryColor   string `json:"primary_color"`
	SecondaryColor string `json:"secondary_color"`
}
type AccountTOSAcceptance struct {
	Date             int64                                `json:"date"`
	IP               string                               `json:"ip"`
	ServiceAgreement AccountTOSAcceptanceServiceAgreement `json:"service_agreement"`
	UserAgent        string                               `json:"user_agent"`
}
type AccountSettingsCardIssuing struct {
	TOSAcceptance *AccountTOSAcceptance `json:"tos_acceptance"`
}
type AccountDeclineOn struct {
	AVSFailure bool `json:"avs_failure"`
	CVCFailure bool `json:"cvc_failure"`
}
type AccountSettingsCardPayments struct {
	DeclineOn                 *AccountDeclineOn `json:"decline_on"`
	StatementDescriptorPrefix string            `json:"statement_descriptor_prefix"`
}
type AccountSettingsDashboard struct {
	DisplayName string `json:"display_name"`
	Timezone    string `json:"timezone"`
}
type AccountSettingsPayments struct {
	StatementDescriptor      string `json:"statement_descriptor"`
	StatementDescriptorKana  string `json:"statement_descriptor_kana"`
	StatementDescriptorKanji string `json:"statement_descriptor_kanji"`
}
type AccountPayoutSchedule struct {
	DelayDays     int64          `json:"delay_days"`
	Interval      PayoutInterval `json:"interval"`
	MonthlyAnchor int64          `json:"monthly_anchor"`
	WeeklyAnchor  string         `json:"weekly_anchor"`
}
type AccountSettingsPayouts struct {
	DebitNegativeBalances bool                   `json:"debit_negative_balances"`
	Schedule              *AccountPayoutSchedule `json:"schedule"`
	StatementDescriptor   string                 `json:"statement_descriptor"`
}
type AccountSettingsSEPADebitPayments struct {
	CreditorID string `json:"creditor_id"`
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
}

// This is an object representing a Stripe account. You can retrieve it to see
// properties on the account like its current e-mail address or if the account is
// enabled yet to make live charges.
//
// Some properties, marked below, are available only to platforms that want to
// [create and manage Express or Custom accounts](https://stripe.com/docs/connect/accounts).
type Account struct {
	APIResource
	BusinessProfile    *AccountBusinessProfile    `json:"business_profile"`
	BusinessType       AccountBusinessType        `json:"business_type"`
	Capabilities       *AccountCapabilities       `json:"capabilities"`
	ChargesEnabled     bool                       `json:"charges_enabled"`
	Company            *AccountCompany            `json:"company"`
	Controller         *AccountController         `json:"controller"`
	Country            string                     `json:"country"`
	Created            int64                      `json:"created"`
	DefaultCurrency    Currency                   `json:"default_currency"`
	Deleted            bool                       `json:"deleted"`
	DetailsSubmitted   bool                       `json:"details_submitted"`
	Email              string                     `json:"email"`
	ExternalAccounts   *ExternalAccountList       `json:"external_accounts"`
	FutureRequirements *AccountFutureRequirements `json:"future_requirements"`
	ID                 string                     `json:"id"`
	Individual         *Person                    `json:"individual"`
	Metadata           map[string]string          `json:"metadata"`
	Object             string                     `json:"object"`
	PayoutsEnabled     bool                       `json:"payouts_enabled"`
	Requirements       *AccountRequirements       `json:"requirements"`
	Settings           *AccountSettings           `json:"settings"`
	TOSAcceptance      *AccountTOSAcceptance      `json:"tos_acceptance"`
	Type               AccountType                `json:"type"`
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
