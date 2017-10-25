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

	// Month is a constant value representing a monthly payout interval.
	Month Interval = "monthly"

	// Week is a constant value representing a weekly payout interval.
	Week Interval = "weekly"
)

// AccountParams are the parameters allowed during account creation/updates.
type AccountParams struct {
	Params               `form:"*"`
	BusinessName         string                        `form:"business_name"`
	BusinessPrimaryColor string                        `form:"business_primary_color"`
	BusinessUrl          string                        `form:"business_url"`
	Country              string                        `form:"country"`
	DebitNegativeBal     bool                          `form:"debit_negative_balances"`
	DefaultCurrency      string                        `form:"default_currency"`
	Email                string                        `form:"email"`
	ExternalAccount      *AccountExternalAccountParams `form:"external_account"`
	FromRecipient        string                        `form:"from_recipient"`
	LegalEntity          *LegalEntity                  `form:"legal_entity"`
	NoDebitNegativeBal   bool                          `form:"debit_negative_balances,invert"`
	PayoutSchedule       *PayoutScheduleParams         `form:"payout_schedule"`
	PayoutStatement      string                        `form:"payout_statement_descriptor"`
	Statement            string                        `form:"statement_descriptor"`
	SupportEmail         string                        `form:"support_email"`
	SupportPhone         string                        `form:"support_phone"`
	SupportUrl           string                        `form:"support_url"`
	TOSAcceptance        *TOSAcceptanceParams          `form:"tos_acceptance"`
	Type                 AccountType                   `form:"type"`
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
	Account           string `form:"account_number"`
	AccountHolderName string `form:"account_holder_name"`
	AccountHolderType string `form:"account_holder_type"`
	Country           string `form:"country"`
	Currency          string `form:"currency"`
	Routing           string `form:"routing_number"`
	Token             string `form:"token"`
}

// AppendTo implements custom encoding logic for AccountExternalAccountParams
// so that we can send the special required `object` field up along with the
// other specified parameters or the token value
func (p *AccountExternalAccountParams) AppendTo(body *form.Values, keyParts []string) {
	if len(p.Token) > 0 {
		body.Add(form.FormatKey(keyParts), p.Token)
	} else {
		body.Add(form.FormatKey(append(keyParts, "object")), "bank_account")
	}
}

// PayoutScheduleParams are the parameters allowed for payout schedules.
type PayoutScheduleParams struct {
	Delay        uint64   `form:"delay_days"`
	Interval     Interval `form:"interval"`
	MinimumDelay bool     `form:"-"` // See custom AppendTo
	MonthAnchor  uint64   `form:"monthly_anchor"`
	WeekAnchor   string   `form:"weekly_anchor"`
}

func (p *PayoutScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if p.MinimumDelay {
		body.Add(form.FormatKey(append(keyParts, "delay_days")), "minimum")
	}
}

// Account is the resource representing your Stripe account.
// For more details see https://stripe.com/docs/api/#account.
type Account struct {
	BusinessLogo         string               `json:"business_logo"`
	BusinessName         string               `json:"business_name"`
	BusinessPrimaryColor string               `json:"business_primary_color"`
	BusinessUrl          string               `json:"business_url"`
	ChargesEnabled       bool                 `json:"charges_enabled"`
	Country              string               `json:"country"`
	DebitNegativeBal     bool                 `json:"debit_negative_balances"`
	DefaultCurrency      string               `json:"default_currency"`
	Deleted              bool                 `json:"deleted"`
	DetailsSubmitted     bool                 `json:"details_submitted"`
	Email                string               `json:"email"`
	ExternalAccounts     *ExternalAccountList `json:"external_accounts"`
	ID                   string               `json:"id"`

	Keys *struct {
		Publish string `json:"publishable"`
		Secret  string `json:"secret"`
	} `json:"keys"`

	LegalEntity     *LegalEntity      `json:"legal_entity"`
	Meta            map[string]string `json:"metadata"`
	Name            string            `json:"display_name"`
	PayoutSchedule  *PayoutSchedule   `json:"payout_schedule"`
	PayoutStatement string            `json:"payout_statement_descriptor"`
	PayoutsEnabled  bool              `json:"payouts_enabled"`
	ProductDesc     string            `json:"product_description"`
	Statement       string            `json:"statement_descriptor"`
	SupportAddress  *Address          `json:"support_address"`
	SupportEmail    string            `json:"support_email"`
	SupportPhone    string            `json:"support_phone"`
	SupportUrl      string            `json:"support_url"`
	Timezone        string            `json:"timezone"`

	TOSAcceptance *struct {
		Date      int64  `json:"date"`
		IP        string `json:"ip"`
		UserAgent string `json:"user_agent"`
	} `json:"tos_acceptance"`

	Type AccountType

	Verification *struct {
		DisabledReason string   `json:"disabled_reason"`
		Due            *int64   `json:"due_by"`
		Fields         []string `json:"fields_needed"`
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
	Values []*Account `json:"data"`
}

// ExternalAccountList is a list of external accounts that may be either bank
// accounts or cards.
type ExternalAccountList struct {
	ListMeta

	// Values contains any external accounts (bank accounts and/or cards)
	// currently attached to this account.
	Values []*ExternalAccount `json:"data"`
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
	AdditionalOwners []Owner `json:"additional_owners" form:"additional_owners,indexed"`

	// AdditionalOwnersEmpty can be set to clear a legal entity's additional
	// owners.
	AdditionalOwnersEmpty bool `form:"additional_owners,empty"`

	Address               Address              `json:"address" form:"address"`
	AddressKana           Address              `json:"address_kana" form:"address_kana"`
	AddressKanji          Address              `json:"address_kanji" form:"address_kanji"`
	BusinessName          string               `json:"business_name" form:"business_name"`
	BusinessNameKana      string               `json:"business_name_kana" form:"business_name_kana"`
	BusinessNameKanji     string               `json:"business_name_kanji" form:"business_name_kanji"`
	BusinessTaxID         string               `json:"-" form:"business_tax_id"`
	BusinessTaxIDProvided bool                 `json:"business_tax_id_provided" form:"-"`
	BusinessVatID         string               `json:"-" form:"business_vat_id"`
	BusinessVatIDProvided bool                 `json:"business_vat_id_provided" form:"-"`
	DOB                   DOB                  `json:"dob" form:"dob"`
	First                 string               `json:"first_name" form:"first_name"`
	FirstKana             string               `json:"first_name_kana" form:"first_name_kana"`
	FirstKanji            string               `json:"first_name_kanji" form:"first_name_kanji"`
	Gender                Gender               `json:"gender" form:"gender"`
	Last                  string               `json:"last_name" form:"last_name"`
	LastKana              string               `json:"last_name_kana" form:"last_name_kana"`
	LastKanji             string               `json:"last_name_kanji" form:"last_name_kanji"`
	MaidenName            string               `json:"maiden_name" form:"maiden_name"`
	PersonalAddress       Address              `json:"personal_address" form:"personal_address"`
	PersonalAddressKana   Address              `json:"personal_address_kana" form:"personal_address_kana"`
	PersonalAddressKanji  Address              `json:"personal_address_kanji" form:"personal_address_kanji"`
	PersonalID            string               `json:"-" form:"personal_id_number"`
	PersonalIDProvided    bool                 `json:"personal_id_number_provided" form:"-"`
	PhoneNumber           string               `json:"phone_number" form:"phone_number"`
	SSN                   string               `json:"-" form:"ssn_last_4"`
	SSNProvided           bool                 `json:"ssn_last_4_provided" form:"-"`
	Type                  LegalEntityType      `json:"type" form:"type"`
	Verification          IdentityVerification `json:"verification" form:"verification"`
}

// Address is the structure for an account address.
type Address struct {
	City    string `json:"city" form:"city"`
	Country string `json:"country" form:"country"`
	Line1   string `json:"line1" form:"line1"`
	Line2   string `json:"line2" form:"line2"`
	State   string `json:"state" form:"state"`

	// Town/cho-me. Note that this is only used for Kana/Kanji representations
	// of an address.
	Town string `json:"town" form:"town"`

	Zip string `json:"postal_code" form:"postal_code"`
}

// DOB is a structure for an account owner's date of birth.
type DOB struct {
	Day   int `json:"day" form:"day"`
	Month int `json:"month" form:"month"`
	Year  int `json:"year" form:"year"`
}

// Gender is the gender of an account owner. International regulations require
// either “male” or “female”.
type Gender string

// Owner is the structure for an account owner.
type Owner struct {
	Address      Address              `json:"address" form:"address"`
	DOB          DOB                  `json:"dob" form:"dob"`
	First        string               `json:"first_name" form:"first_name"`
	Last         string               `json:"last_name" form:"last_name"`
	Verification IdentityVerification `json:"verification" form:"verification"`
}

// IdentityVerification is the structure for an account's verification.
type IdentityVerification struct {
	Details     *string                         `json:"details" form:"-"`
	DetailsCode IdentityVerificationDetailsCode `json:"details_code" form:"-"`
	Document    *IdentityDocument               `json:"document" form:"document"`
	Status      IdentityVerificationStatus      `json:"status" form:"-"`
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
	Delay       uint64   `json:"delay_days" form:"delay_days"`
	Interval    Interval `json:"interval" form:"interval"`
	MonthAnchor uint64   `json:"monthly_anchor" form:"monthly_anchor"`
	WeekAnchor  string   `json:"weekly_anchor" form:"weekly_anchor"`
}

// TOSAcceptanceParams is the structure for TOS acceptance.
type TOSAcceptanceParams struct {
	Date      int64  `json:"date" form:"date"`
	IP        string `json:"ip" form:"ip"`
	UserAgent string `json:"user_agent" form:"user_agent"`
}

// AccountRejectParams is the structure for the Reject function.
type AccountRejectParams struct {
	// Reason is the reason that an account was rejected. It should be given a
	// value of one of `fraud`, `terms_of_service`, or `other`.
	Reason string `json:"reason" form:"reason"`
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
