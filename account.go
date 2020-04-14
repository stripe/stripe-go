package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/form"
)

// AccountType is the type of an account.
type AccountType string

// List of values that AccountType can take.
const (
	AccountTypeCustom   AccountType = "custom"
	AccountTypeExpress  AccountType = "express"
	AccountTypeStandard AccountType = "standard"
)

// AccountCapability maps to a given capability for an account.
type AccountCapability string

// List of values that AccountCapability can take.
const (
	AccountCapabilityAUBECSDebitPayments    AccountCapability = "au_becs_debit_payments"
	AccountCapabilityCardIssuing            AccountCapability = "card_issuing"
	AccountCapabilityCardPayments           AccountCapability = "card_payments"
	AccountCapabilityLegacyPayments         AccountCapability = "legacy_payments"
	AccountCapabilityTaxReportingUS1099K    AccountCapability = "tax_reporting_us_1099_k"
	AccountCapabilityTaxReportingUS1099MISC AccountCapability = "tax_reporting_us_1099_misc"
	AccountCapabilityTransfers              AccountCapability = "transfers"
)

// AccountCapabilityStatus is the status a given capability can have
type AccountCapabilityStatus string

// List of values that AccountCapabilityStatus can take.
const (
	AccountCapabilityStatusActive   AccountCapabilityStatus = "active"
	AccountCapabilityStatusInactive AccountCapabilityStatus = "inactive"
	AccountCapabilityStatusPending  AccountCapabilityStatus = "pending"
)

// ExternalAccountType is the type of an external account.
type ExternalAccountType string

// List of values that ExternalAccountType can take.
const (
	ExternalAccountTypeBankAccount ExternalAccountType = "bank_account"
	ExternalAccountTypeCard        ExternalAccountType = "card"
)

// AccountBusinessType describes the business type associated with an account.
type AccountBusinessType string

// List of values that AccountBusinessType can take.
const (
	AccountBusinessTypeCompany          AccountBusinessType = "company"
	AccountBusinessTypeGovernmentEntity AccountBusinessType = "government_entity"
	AccountBusinessTypeIndividual       AccountBusinessType = "individual"
	AccountBusinessTypeNonProfit        AccountBusinessType = "non_profit"
)

// AccountCompanyStructure describes the structure associated with a company.
type AccountCompanyStructure string

// List of values that AccountCompanyStructure can take.
const (
	AccountCompanyStructureGovernmentInstrumentality          AccountCompanyStructure = "government_instrumentality"
	AccountCompanyStructureGovernmentalUnit                   AccountCompanyStructure = "governmental_unit"
	AccountCompanyStructureIncorporatedNonProfit              AccountCompanyStructure = "incorporated_non_profit"
	AccountCompanyStructureMultiMemberLLC                     AccountCompanyStructure = "multi_member_llc"
	AccountCompanyStructurePrivateCorporation                 AccountCompanyStructure = "private_corporation"
	AccountCompanyStructurePrivatePartnership                 AccountCompanyStructure = "private_partnership"
	AccountCompanyStructurePublicCorporation                  AccountCompanyStructure = "public_corporation"
	AccountCompanyStructurePublicPartnership                  AccountCompanyStructure = "public_partnership"
	AccountCompanyStructureTaxExemptGovernmentInstrumentality AccountCompanyStructure = "tax_exempt_government_instrumentality"
	AccountCompanyStructureUnincorporatedAssociation          AccountCompanyStructure = "unincorporated_association"
	AccountCompanyStructureUnincorporatedNonProfit            AccountCompanyStructure = "unincorporated_non_profit"
)

// AccountRequirementsDisabledReason describes why an account is disabled.
type AccountRequirementsDisabledReason string

// List of values that AccountRequirementsDisabledReason can take.
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

// PayoutInterval describes the payout interval.
type PayoutInterval string

// List of values that PayoutInterval can take.
const (
	PayoutIntervalDaily   PayoutInterval = "daily"
	PayoutIntervalManual  PayoutInterval = "manual"
	PayoutIntervalMonthly PayoutInterval = "monthly"
	PayoutIntervalWeekly  PayoutInterval = "weekly"
)

// AccountRejectReason describes the valid reason to reject an account
type AccountRejectReason string

// List of values that AccountRejectReason can take.
const (
	AccountRejectReasonFraud          AccountRejectReason = "fraud"
	AccountRejectReasonOther          AccountRejectReason = "other"
	AccountRejectReasonTermsOfService AccountRejectReason = "terms_of_service"
)

// AccountCompanyVerificationDocumentDetailsCode is a machine-readable code specifying the
// verification state of a document associated with a company.
type AccountCompanyVerificationDocumentDetailsCode string

// List of values that AccountCompanyVerificationDocumentDetailsCode can take.
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

// AccountBusinessProfileParams are the parameters allowed for an account's business information
type AccountBusinessProfileParams struct {
	MCC                *string `form:"mcc"`
	Name               *string `form:"name"`
	ProductDescription *string `form:"product_description"`
	SupportEmail       *string `form:"support_email"`
	SupportPhone       *string `form:"support_phone"`
	SupportURL         *string `form:"support_url"`
	URL                *string `form:"url"`
}

// AccountCompanyVerificationDocumentParams are the parameters allowed to pass for a document
// verifying a company.
type AccountCompanyVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// AccountCompanyVerificationParams are the parameters allowed to verify a company.
type AccountCompanyVerificationParams struct {
	Document *AccountCompanyVerificationDocumentParams `form:"document"`
}

// AccountCompanyParams are the parameters describing the company associated with the account.
type AccountCompanyParams struct {
	Address            *AccountAddressParams             `form:"address"`
	AddressKana        *AccountAddressParams             `form:"address_kana"`
	AddressKanji       *AccountAddressParams             `form:"address_kanji"`
	DirectorsProvided  *bool                             `form:"directors_provided"`
	ExecutivesProvided *bool                             `form:"executives_provided"`
	Name               *string                           `form:"name"`
	NameKana           *string                           `form:"name_kana"`
	NameKanji          *string                           `form:"name_kanji"`
	OwnersProvided     *bool                             `form:"owners_provided"`
	Structure          *string                           `form:"structure"`
	Phone              *string                           `form:"phone"`
	TaxID              *string                           `form:"tax_id"`
	TaxIDRegistrar     *string                           `form:"tax_id_registrar"`
	VATID              *string                           `form:"vat_id"`
	Verification       *AccountCompanyVerificationParams `form:"verification"`
}

// AccountDeclineSettingsParams represents the parameters allowed for configuring
// card declines on connected accounts.
type AccountDeclineSettingsParams struct {
	AVSFailure *bool `form:"avs_failure"`
	CVCFailure *bool `form:"cvc_failure"`
}

// AccountSettingsBrandingParams represent allowed parameters to configure settings specific to the
// account’s branding.
type AccountSettingsBrandingParams struct {
	Icon           *string `form:"icon"`
	Logo           *string `form:"logo"`
	PrimaryColor   *string `form:"primary_color"`
	SecondaryColor *string `form:"secondary_color"`
}

// AccountSettingsCardPaymentsParams represent allowed parameters to configure settings specific to
// card charging on the account.
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

// AccountSettingsPaymentsParams represent allowed parameters to configure settings  across payment
// methods for charging on the account.
type AccountSettingsPaymentsParams struct {
	StatementDescriptor      *string `form:"statement_descriptor"`
	StatementDescriptorKana  *string `form:"statement_descriptor_kana"`
	StatementDescriptorKanji *string `form:"statement_descriptor_kanji"`
}

// AccountSettingsPayoutsParams represent allowed parameters to configure settings specific to the
// account’s payouts.
type AccountSettingsPayoutsParams struct {
	DebitNegativeBalances *bool                 `form:"debit_negative_balances"`
	Schedule              *PayoutScheduleParams `form:"schedule"`
	StatementDescriptor   *string               `form:"statement_descriptor"`
}

// AccountSettingsParams are the parameters allowed for the account's settings.
type AccountSettingsParams struct {
	Branding     *AccountSettingsBrandingParams     `form:"branding"`
	CardPayments *AccountSettingsCardPaymentsParams `form:"card_payments"`
	Dashboard    *AccountSettingsDashboardParams    `form:"dashboard"`
	Payments     *AccountSettingsPaymentsParams     `form:"payments"`
	Payouts      *AccountSettingsPayoutsParams      `form:"payouts"`
}

// PayoutScheduleParams are the parameters allowed for payout schedules.
type PayoutScheduleParams struct {
	DelayDays        *int64  `form:"delay_days"`
	DelayDaysMinimum *bool   `form:"-"` // See custom AppendTo
	Interval         *string `form:"interval"`
	MonthlyAnchor    *int64  `form:"monthly_anchor"`
	WeeklyAnchor     *string `form:"weekly_anchor"`
}

// AppendTo implements custom encoding logic for PayoutScheduleParams
// so that we can send a special value for `delay_days` field if needed.
func (p *PayoutScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.DelayDaysMinimum) {
		body.Add(form.FormatKey(append(keyParts, "delay_days")), "minimum")
	}
}

// AccountParams are the parameters allowed during account creation/updates.
type AccountParams struct {
	Params                `form:"*"`
	AccountToken          *string                       `form:"account_token"`
	BusinessProfile       *AccountBusinessProfileParams `form:"business_profile"`
	BusinessType          *string                       `form:"business_type"`
	Company               *AccountCompanyParams         `form:"company"`
	Country               *string                       `form:"country"`
	DefaultCurrency       *string                       `form:"default_currency"`
	Email                 *string                       `form:"email"`
	ExternalAccount       *AccountExternalAccountParams `form:"external_account"`
	Individual            *PersonParams                 `form:"individual"`
	RequestedCapabilities []*string                     `form:"requested_capabilities"`
	Settings              *AccountSettingsParams        `form:"settings"`
	TOSAcceptance         *AccountTOSAcceptanceParams   `form:"tos_acceptance"`
	Type                  *string                       `form:"type"`
}

// AccountAddressParams represents an address during account creation/updates.
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

// AccountTOSAcceptanceParams represents tos_acceptance during account creation/updates.
type AccountTOSAcceptanceParams struct {
	Date      *int64  `form:"date"`
	IP        *string `form:"ip"`
	UserAgent *string `form:"user_agent"`
}

// AccountListParams are the parameters allowed during account listing.
type AccountListParams struct {
	ListParams `form:"*"`
}

// AccountRejectParams is the structure for the Reject function.
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

// AccountBusinessProfile represents optional information related to the business.
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

// AccountCapabilities is the resource representing the capabilities enabled on that account.
type AccountCapabilities struct {
	AUBECSDebitPayments    AccountCapabilityStatus `json:"au_becs_debit_payments"`
	CardPayments           AccountCapabilityStatus `json:"card_payments"`
	LegacyPayments         AccountCapabilityStatus `json:"legacy_payments"`
	TaxReportingUS1099K    AccountCapabilityStatus `json:"tax_reporting_us_1099_k"`
	TaxReportingUS1099MISC AccountCapabilityStatus `json:"tax_reporting_us_1099_misc"`
	Transfers              AccountCapabilityStatus `json:"transfers"`
}

// AccountCompanyVerificationDocument represents details about a company's verification state.
type AccountCompanyVerificationDocument struct {
	Back        *File                                         `json:"back"`
	Details     string                                        `json:"details"`
	DetailsCode AccountCompanyVerificationDocumentDetailsCode `json:"details_code"`
	Front       *File                                         `json:"front"`
}

// AccountCompanyVerification represents details about a company's verification state.
type AccountCompanyVerification struct {
	Document *AccountCompanyVerificationDocument `json:"document"`
}

// AccountCompany represents details about the company or business associated with the account.
type AccountCompany struct {
	Address            *AccountAddress             `json:"address"`
	AddressKana        *AccountAddress             `json:"address_kana"`
	AddressKanji       *AccountAddress             `json:"address_kanji"`
	DirectorsProvided  bool                        `json:"directors_provided"`
	ExecutivesProvided bool                        `json:"executives_provided"`
	Name               string                      `json:"name"`
	NameKana           string                      `json:"name_kana"`
	NameKanji          string                      `json:"name_kanji"`
	OwnersProvided     bool                        `json:"owners_provided"`
	Phone              string                      `json:"phone"`
	Structure          AccountCompanyStructure     `json:"structure"`
	TaxIDProvided      bool                        `json:"tax_id_provided"`
	TaxIDRegistrar     string                      `json:"tax_id_registrar"`
	VATIDProvided      bool                        `json:"vat_id_provided"`
	Verification       *AccountCompanyVerification `json:"verification"`
}

// AccountDeclineOn represents card charges decline behavior for that account.
type AccountDeclineOn struct {
	AVSFailure bool `json:"avs_failure"`
	CVCFailure bool `json:"cvc_failure"`
}

// AccountPayoutSchedule is the structure for an account's payout schedule.
type AccountPayoutSchedule struct {
	DelayDays     int64          `json:"delay_days"`
	Interval      PayoutInterval `json:"interval"`
	MonthlyAnchor int64          `json:"monthly_anchor"`
	WeeklyAnchor  string         `json:"weekly_anchor"`
}

// AccountRequirementsError represents details about an error with a requirement.
type AccountRequirementsError struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Requirement string `json:"requirement"`
}

// AccountRequirements represents information that needs to be collected for an account.
type AccountRequirements struct {
	CurrentDeadline     int64                             `json:"current_deadline"`
	CurrentlyDue        []string                          `json:"currently_due"`
	DisabledReason      AccountRequirementsDisabledReason `json:"disabled_reason"`
	Errors              []*AccountRequirementsError       `json:"errors"`
	EventuallyDue       []string                          `json:"eventually_due"`
	PastDue             []string                          `json:"past_due"`
	PendingVerification []string                          `json:"pending_verification"`
}

// AccountSettingsBranding represents settings specific to the account's branding.
type AccountSettingsBranding struct {
	Icon           *File  `json:"icon"`
	Logo           *File  `json:"logo"`
	PrimaryColor   string `json:"primary_color"`
	SecondaryColor string `json:"secondary_color"`
}

// AccountSettingsCardPayments represents settings specific to card charging on the account.
type AccountSettingsCardPayments struct {
	DeclineOn                 *AccountDeclineOn `json:"decline_on"`
	StatementDescriptorPrefix string            `json:"statement_descriptor_prefix"`
}

// AccountSettingsDashboard represents settings specific to the account's Dashboard.
type AccountSettingsDashboard struct {
	DisplayName string `json:"display_name"`
	Timezone    string `json:"timezone"`
}

// AccountSettingsPayments represents settings that apply across payment methods for charging on
// the account.
type AccountSettingsPayments struct {
	StatementDescriptor      string `json:"statement_descriptor"`
	StatementDescriptorKana  string `json:"statement_descriptor_kana"`
	StatementDescriptorKanji string `json:"statement_descriptor_kanji"`
}

// AccountSettingsPayouts represents settings specific to the account’s payouts.
type AccountSettingsPayouts struct {
	DebitNegativeBalances bool                   `json:"debit_negative_balances"`
	Schedule              *AccountPayoutSchedule `json:"schedule"`
	StatementDescriptor   string                 `json:"statement_descriptor"`
}

// AccountSettings represents options for customizing how the account functions within Stripe.
type AccountSettings struct {
	Branding     *AccountSettingsBranding     `json:"branding"`
	CardPayments *AccountSettingsCardPayments `json:"card_payments"`
	Dashboard    *AccountSettingsDashboard    `json:"dashboard"`
	Payments     *AccountSettingsPayments     `json:"payments"`
	Payouts      *AccountSettingsPayouts      `json:"payouts"`
}

// AccountTOSAcceptance represents status of acceptance of our terms of services for the account.
type AccountTOSAcceptance struct {
	Date      int64  `json:"date"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

// Account is the resource representing your Stripe account.
// For more details see https://stripe.com/docs/api/#account.
type Account struct {
	BusinessProfile  *AccountBusinessProfile `json:"business_profile"`
	BusinessType     AccountBusinessType     `json:"business_type"`
	Capabilities     *AccountCapabilities    `json:"capabilities"`
	ChargesEnabled   bool                    `json:"charges_enabled"`
	Company          *AccountCompany         `json:"company"`
	Country          string                  `json:"country"`
	Created          int64                   `json:"created"`
	DefaultCurrency  Currency                `json:"default_currency"`
	Deleted          bool                    `json:"deleted"`
	DetailsSubmitted bool                    `json:"details_submitted"`
	Email            string                  `json:"email"`
	ExternalAccounts *ExternalAccountList    `json:"external_accounts"`
	ID               string                  `json:"id"`
	Individual       *Person                 `json:"individual"`
	Metadata         map[string]string       `json:"metadata"`
	Object           string                  `json:"object"`
	PayoutsEnabled   bool                    `json:"payouts_enabled"`
	Requirements     *AccountRequirements    `json:"requirements"`
	Settings         *AccountSettings        `json:"settings"`
	TOSAcceptance    *AccountTOSAcceptance   `json:"tos_acceptance"`
	Type             AccountType             `json:"type"`
}

// UnmarshalJSON handles deserialization of an account.
// This custom unmarshaling is needed because the resulting
// property may be an ID or the full struct if it was expanded.
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

// AccountList is a list of accounts as returned from a list endpoint.
type AccountList struct {
	ListMeta
	Data []*Account `json:"data"`
}

// ExternalAccountList is a list of external accounts that may be either bank
// accounts or cards.
type ExternalAccountList struct {
	ListMeta

	// Values contains any external accounts (bank accounts and/or cards)
	// currently attached to this account.
	Data []*ExternalAccount `json:"data"`
}

// ExternalAccount is an external account (a bank account or card) that's
// attached to an account. It contains fields that will be conditionally
// populated depending on its type.
type ExternalAccount struct {
	// BankAccount is a bank account attached to an account. Populated only if
	// the external account is a bank account.
	BankAccount *BankAccount

	// Card is a card attached to an account. Populated only if the external
	// account is a card.
	Card *Card

	ID   string              `json:"id"`
	Type ExternalAccountType `json:"object"`
}

// UnmarshalJSON implements Unmarshaler.UnmarshalJSON.
func (ea *ExternalAccount) UnmarshalJSON(data []byte) error {
	type externalAccount ExternalAccount
	var v externalAccount
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var err error
	*ea = ExternalAccount(v)

	switch ea.Type {
	case ExternalAccountTypeBankAccount:
		err = json.Unmarshal(data, &ea.BankAccount)
	case ExternalAccountTypeCard:
		err = json.Unmarshal(data, &ea.Card)
	}

	return err
}

// AccountAddress is the structure for an account address.
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
