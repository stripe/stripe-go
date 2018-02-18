package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/form"
)

// LegalEntityType describes the types for a legal entity.
// Allowed values are "individual", "company".
type LegalEntityType string

// IdentityVerificationDetailsCode is a machine-readable code specifying the
// verification state of a legal entity. Allowed values are
// "failed_keyed_identity", "failed_other", "scan_corrupt",
// "scan_failed_greyscale", "scan_failed_other",
// "scan_id_country_not_supported", "scan_id_type_not_supported",
// "scan_name_mismatch", "scan_not_readable", "scan_not_uploaded".
type IdentityVerificationDetailsCode string

// IdentityVerificationStatus describes the different statuses for identity verification.
// Allowed values are "pending", "verified", "unverified".
type IdentityVerificationStatus string

// Interval describes the payout interval.
// Allowed values are "manual", "daily", "weekly", "monthly".
type Interval string

const (
	// Company is a constant value representing a company legal entity type.
	Company LegalEntityType = "company"

	// Day is a constant value representing a daily payout interval.
	Day Interval = "daily"

	// Individual is a constant value representing an individual legal entity
	// type.
	Individual LegalEntityType = "individual"

	// IdentityVerificationPending is a constant value indicating that identity
	// verification status is pending.
	IdentityVerificationPending IdentityVerificationStatus = "pending"

	// IdentityVerificationUnverified is a constant value indicating that
	// identity verification status is unverified.
	IdentityVerificationUnverified IdentityVerificationStatus = "unverified"

	// IdentityVerificationVerified is a constant value indicating that
	// identity verification status is verified.
	IdentityVerificationVerified IdentityVerificationStatus = "verified"

	// Manual is a constant value representing a manual payout interval.
	Manual Interval = "manual"

	// Monthly is a constant value representing a monthly payout interval.
	Monthly Interval = "monthly"

	// Weekly is a constant value representing a weekly payout interval.
	Weekly Interval = "weekly"
)

// AccountParams are the parameters allowed during account creation/updates.
type AccountParams struct {
	Params                    `form:"*"`
	BusinessName              *string                       `form:"business_name"`
	BusinessPrimaryColor      *string                       `form:"business_primary_color"`
	BusinessURL               *string                       `form:"business_url"`
	Country                   *string                       `form:"country"`
	DebitNegativeBalances     *bool                         `form:"debit_negative_balances"`
	DefaultCurrency           *string                       `form:"default_currency"`
	Email                     *string                       `form:"email"`
	ExternalAccount           *AccountExternalAccountParams `form:"external_account"`
	FromRecipient             *string                       `form:"from_recipient"`
	LegalEntity               *LegalEntityParams            `form:"legal_entity"`
	PayoutSchedule            *PayoutScheduleParams         `form:"payout_schedule"`
	PayoutStatementDescriptor *string                       `form:"payout_statement_descriptor"`
	ProductDescription        *string                       `form:"product_description"`
	StatementDescriptor       *string                       `form:"statement_descriptor"`
	SupportEmail              *string                       `form:"support_email"`
	SupportPhone              *string                       `form:"support_phone"`
	SupportURL                *string                       `form:"support_url"`
	TOSAcceptance             *TOSAcceptanceParams          `form:"tos_acceptance"`
	Type                      *string                       `form:"type"`
}

// LegalEntityParams represents a legal_entity during account creation/updates.
type LegalEntityParams struct {
	AdditionalOwners []AdditionalOwnerParams `form:"additional_owners,indexed"`

	// AdditionalOwnersEmpty can be set to clear a legal entity's additional
	// owners.
	AdditionalOwnersEmpty bool `form:"additional_owners,empty"`

	Address              *AccountAddressParams       `form:"address"`
	AddressKana          *AccountAddressParams       `form:"address_kana"`
	AddressKanji         *AccountAddressParams       `form:"address_kanji"`
	BusinessName         *string                     `form:"business_name"`
	BusinessNameKana     *string                     `form:"business_name_kana"`
	BusinessNameKanji    *string                     `form:"business_name_kanji"`
	BusinessTaxID        *string                     `form:"business_tax_id"`
	BusinessVatID        *string                     `form:"business_vat_id"`
	DOB                  *DOBParams                  `form:"dob"`
	FirstName            *string                     `form:"first_name"`
	FirstNameKana        *string                     `form:"first_name_kana"`
	FirstNameKanji       *string                     `form:"first_name_kanji"`
	Gender               *string                     `form:"gender"`
	LastName             *string                     `form:"last_name"`
	LastNameKana         *string                     `form:"last_name_kana"`
	LastNameKanji        *string                     `form:"last_name_kanji"`
	MaidenName           *string                     `form:"maiden_name"`
	PersonalAddress      *AccountAddressParams       `form:"personal_address"`
	PersonalAddressKana  *AccountAddressParams       `form:"personal_address_kana"`
	PersonalAddressKanji *AccountAddressParams       `form:"personal_address_kanji"`
	PersonalIDNumber     *string                     `form:"personal_id_number"`
	PhoneNumber          *string                     `form:"phone_number"`
	SSNLast4             *string                     `form:"ssn_last_4"`
	Type                 *string                     `form:"type"`
	Verification         *IdentityVerificationParams `form:"verification"`
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

// DOBParams represents a DOB during account creation/updates.
type DOBParams struct {
	Day   *int `form:"day"`
	Month *int `form:"month"`
	Year  *int `form:"year"`
}

// TOSAcceptanceParams represents tos_acceptance during account creation/updates.
type TOSAcceptanceParams struct {
	Date      *int64  `form:"date"`
	IP        *string `form:"ip"`
	UserAgent *string `form:"user_agent"`
}

// AdditionalOwnerParams represents an additional owner during account creation/updates.
type AdditionalOwnerParams struct {
	Address      *AccountAddressParams       `form:"address"`
	DOB          *DOBParams                  `form:"dob"`
	FirstName    *string                     `form:"first_name"`
	LastName     *string                     `form:"last_name"`
	MaidenName   *string                     `form:"maiden_name"`
	Verification *IdentityVerificationParams `form:"verification"`
}

// IdentityVerification represents a verification during account creation/updates.
type IdentityVerificationParams struct {
	Document *IdentityDocument `form:"document"`
}

// AccountListParams are the parameters allowed during account listing.
type AccountListParams struct {
	ListParams `form:"*"`
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
// other specified parameters or the token value
func (p *AccountExternalAccountParams) AppendTo(body *form.Values, keyParts []string) {
	if p.Token != nil {
		body.Add(form.FormatKey(keyParts), StringValue(p.Token))
	} else {
		body.Add(form.FormatKey(append(keyParts, "object")), "bank_account")
	}
}

// PayoutScheduleParams are the parameters allowed for payout schedules.
type PayoutScheduleParams struct {
	DelayDays        *uint64 `form:"delay_days"`
	DelayDaysMinimum *bool   `form:"-"` // See custom AppendTo
	Interval         *string `form:"interval"`
	MonthlyAnchor    *uint64 `form:"monthly_anchor"`
	WeeklyAnchor     *string `form:"weekly_anchor"`
}

func (p *PayoutScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.DelayDaysMinimum) {
		body.Add(form.FormatKey(append(keyParts, "delay_days")), "minimum")
	}
}

// Account is the resource representing your Stripe account.
// For more details see https://stripe.com/docs/api/#account.
type Account struct {
	BusinessLogo          string               `json:"business_logo"`
	BusinessName          string               `json:"business_name"`
	BusinessPrimaryColor  string               `json:"business_primary_color"`
	BusinessURL           string               `json:"business_url"`
	ChargesEnabled        bool                 `json:"charges_enabled"`
	Country               string               `json:"country"`
	DebitNegativeBalances bool                 `json:"debit_negative_balances"`
	DefaultCurrency       string               `json:"default_currency"`
	Deleted               bool                 `json:"deleted"`
	DetailsSubmitted      bool                 `json:"details_submitted"`
	Email                 string               `json:"email"`
	ExternalAccounts      *ExternalAccountList `json:"external_accounts"`
	ID                    string               `json:"id"`

	Keys *struct {
		Publishable string `json:"publishable"`
		Secret      string `json:"secret"`
	} `json:"keys"`

	LegalEntity               *LegalEntity      `json:"legal_entity"`
	Metadata                  map[string]string `json:"metadata"`
	DisplayName               string            `json:"display_name"`
	PayoutSchedule            *PayoutSchedule   `json:"payout_schedule"`
	PayoutStatementDescriptor string            `json:"payout_statement_descriptor"`
	PayoutsEnabled            bool              `json:"payouts_enabled"`
	ProductDescription        string            `json:"product_description"`
	StatementDescriptor       string            `json:"statement_descriptor"`
	SupportAddress            *Address          `json:"support_address"`
	SupportEmail              string            `json:"support_email"`
	SupportPhone              string            `json:"support_phone"`
	SupportURL                string            `json:"support_url"`
	Timezone                  string            `json:"timezone"`

	TOSAcceptance *struct {
		Date      int64  `json:"date"`
		IP        string `json:"ip"`
		UserAgent string `json:"user_agent"`
	} `json:"tos_acceptance"`

	Type AccountType `json:"type"`

	Verification *struct {
		DisabledReason string   `json:"disabled_reason"`
		DueBy          *int64   `json:"due_by"`
		FieldsNeeded   []string `json:"fields_needed"`
	} `json:"verification"`
}

// UnmarshalJSON handles deserialization of an account.
// This custom unmarshaling is needed because the resulting
// property may be an ID or the full struct if it was expanded.
func (a *Account) UnmarshalJSON(data []byte) error {
	type account Account
	var aa account
	err := json.Unmarshal(data, &aa)
	if err == nil {
		*a = Account(aa)
	} else {
		// the ID is surrounded by "\" characters, so strip them
		a.ID = string(data[1 : len(data)-1])
	}
	return nil
}

// ExternalAccountType is the type of an external account.
type ExternalAccountType string

const (
	// ExternalAccountTypeBankAccount is a constant value representing an external
	// account which is a bank account.
	ExternalAccountTypeBankAccount ExternalAccountType = "bank_account"

	// ExternalAccountTypeCard is a constant value representing an external account
	// which is a card.
	ExternalAccountTypeCard ExternalAccountType = "card"
)

// AccountType is the type of an account.
type AccountType string

const (
	// AccountTypeCustom is a constant value representing an account of type custom.
	AccountTypeCustom AccountType = "custom"

	// AccountTypeExpress is a constant value representing an account of type express.
	AccountTypeExpress AccountType = "express"

	// AccountTypeStandard is a constant value representing an account of type standard.
	AccountTypeStandard AccountType = "standard"
)

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
func (ea *ExternalAccount) UnmarshalJSON(b []byte) error {
	type externalAccount ExternalAccount
	var account externalAccount

	err := json.Unmarshal(b, &account)
	if err != nil {
		return err
	}

	*ea = ExternalAccount(account)

	switch ea.Type {
	case ExternalAccountTypeBankAccount:
		err = json.Unmarshal(b, &ea.BankAccount)
	case ExternalAccountTypeCard:
		err = json.Unmarshal(b, &ea.Card)
	}
	return err
}

// LegalEntity is the structure for properties related to an account's legal state.
type LegalEntity struct {
	AdditionalOwners         []AdditionalOwner    `json:"additional_owners"`
	Address                  AccountAddress       `json:"address"`
	AddressKana              AccountAddress       `json:"address_kana"`
	AddressKanji             AccountAddress       `json:"address_kanji"`
	BusinessName             string               `json:"business_name"`
	BusinessNameKana         string               `json:"business_name_kana"`
	BusinessNameKanji        string               `json:"business_name_kanji"`
	BusinessTaxIDProvided    bool                 `json:"business_tax_id_provided"`
	BusinessVatIDProvided    bool                 `json:"business_vat_id_provided"`
	DOB                      DOB                  `json:"dob"`
	FirstName                string               `json:"first_name"`
	FirstNameKana            string               `json:"first_name_kana"`
	FirstNameKanji           string               `json:"first_name_kanji"`
	Gender                   Gender               `json:"gender"`
	LastName                 string               `json:"last_name"`
	LastNameKana             string               `json:"last_name_kana"`
	LastNameKanji            string               `json:"last_name_kanji"`
	MaidenName               string               `json:"maiden_name"`
	PersonalAddress          AccountAddress       `json:"personal_address"`
	PersonalAddressKana      AccountAddress       `json:"personal_address_kana"`
	PersonalAddressKanji     AccountAddress       `json:"personal_address_kanji"`
	PersonalIDNumberProvided bool                 `json:"personal_id_number_provided"`
	PhoneNumber              string               `json:"phone_number"`
	SSNLast4Provided         bool                 `json:"ssn_last_4_provided"`
	Type                     LegalEntityType      `json:"type"`
	Verification             IdentityVerification `json:"verification"`
}

// Address is the structure for an account address.
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

// DOB is a structure for an account owner's date of birth.
type DOB struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

// Gender is the gender of an account owner. International regulations require
// either “male” or “female”.
type Gender string

// AdditionalOwner is the structure for an account owner.
type AdditionalOwner struct {
	Address      AccountAddress       `json:"address"`
	DOB          DOB                  `json:"dob"`
	FirstName    string               `json:"first_name"`
	LastName     string               `json:"last_name"`
	MaidenName   string               `json:"maiden_name"`
	Verification IdentityVerification `json:"verification"`
}

// IdentityVerification is the structure for an account's verification.
type IdentityVerification struct {
	Details     *string                         `json:"details"`
	DetailsCode IdentityVerificationDetailsCode `json:"details_code"`
	Document    *IdentityDocument               `json:"document"`
	Status      IdentityVerificationStatus      `json:"status"`
}

// IdentityDocument is the structure for an identity document.
type IdentityDocument struct {
	Created int64  `json:"created" form:"-"`
	ID      string `json:"id" form:"-"` // See custom AppendTo implementation
	Size    int64  `json:"size" form:"-"`
}

// AppendTo implements custom form encoding for IdentityDocument. In the Go
// library, it's suggested that users initialize an IdentityDocument and fill
// its ID when updating an account, but in the API's account update method,
// there is no subobject; you simply pass an ID into the document field.
//
// The inherent cause of this asymmetry is that instead of creating separate
// structs for parameters (which are normally separate from the structs used
// for responses), someone decided to reuse them instead, and although request
// and response constructs are similar, they're not identical, thus the
// discrepancy.
//
// Long term, we should create separate parameter structs. This isn't hard, but
// is breaking, and will be somewhat painful for users when they upgrade.
func (d *IdentityDocument) AppendTo(body *form.Values, keyParts []string) {
	body.Add(form.FormatKey(keyParts), d.ID)
}

// PayoutSchedule is the structure for an account's payout schedule.
type PayoutSchedule struct {
	DelayDays     uint64   `json:"delay_days"`
	Interval      Interval `json:"interval"`
	MonthlyAnchor uint64   `json:"monthly_anchor"`
	WeeklyAnchor  string   `json:"weekly_anchor"`
}

// AccountRejectParams is the structure for the Reject function.
type AccountRejectParams struct {
	// Reason is the reason that an account was rejected. It should be given a
	// value of one of `fraud`, `terms_of_service`, or `other`.
	Reason *string `form:"reason"`
}

// UnmarshalJSON handles deserialization of an IdentityDocument.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (d *IdentityDocument) UnmarshalJSON(data []byte) error {
	type identityDocument IdentityDocument
	var doc identityDocument
	err := json.Unmarshal(data, &doc)

	if err == nil {
		*d = IdentityDocument(doc)
	} else {
		// the id is surrounded by "\" characters, so strip them
		d.ID = string(data[1 : len(data)-1])
	}

	return nil
}
