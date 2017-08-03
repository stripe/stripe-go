package stripe

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	// Individual is a constant value representing an individual legal entity
	// type.
	Individual LegalEntityType = "individual"

	// Company is a constant value representing a company legal entity type.
	Company LegalEntityType = "company"

	// IdentityVerificationPending is a constant value indicating that identity
	// verification status is pending.
	IdentityVerificationPending IdentityVerificationStatus = "pending"

	// IdentityVerificationVerified is a constant value indicating that
	// identity verification status is verified.
	IdentityVerificationVerified IdentityVerificationStatus = "verified"

	// IdentityVerificationUnverified is a constant value indicating that
	// identity verification status is unverified.
	IdentityVerificationUnverified IdentityVerificationStatus = "unverified"

	// Manual is a constant value representing a manual payout interval.
	Manual Interval = "manual"

	// Day is a constant value representing a daily payout interval.
	Day Interval = "daily"

	// Week is a constant value representing a weekly payout interval.
	Week Interval = "weekly"

	// Month is a constant value representing a monthly payout interval.
	Month Interval = "monthly"
)

// AccountParams are the parameters allowed during account creation/updates.
type AccountParams struct {
	Params
	Country, Email, DefaultCurrency, Statement, BusinessName, BusinessUrl,
	BusinessPrimaryColor, SupportPhone, SupportEmail, SupportUrl,
	FromRecipient, PayoutStatement string
	ExternalAccount                      *AccountExternalAccountParams
	LegalEntity                          *LegalEntity
	PayoutSchedule                       *PayoutScheduleParams
	DebitNegativeBal, NoDebitNegativeBal bool
	TOSAcceptance                        *TOSAcceptanceParams
	Type                                 AccountType
}

// AccountListParams are the parameters allowed during account listing.
type AccountListParams struct {
	ListParams
}

// AccountExternalAccountParams are the parameters allowed to reference an
// external account when creating an account. It should either have Token set
// or everything else.
type AccountExternalAccountParams struct {
	Params
	Account, Country, Currency, Routing, Token string
}

// PayoutScheduleParams are the parameters allowed for payout schedules.
type PayoutScheduleParams struct {
	Delay, MonthAnchor uint64
	WeekAnchor         string
	Interval           Interval
	MinimumDelay       bool
}

// Account is the resource representing your Stripe account.
// For more details see https://stripe.com/docs/api/#account.
type Account struct {
	ID                   string               `json:"id"`
	ChargesEnabled       bool                 `json:"charges_enabled"`
	Country              string               `json:"country"`
	DefaultCurrency      string               `json:"default_currency"`
	DetailsSubmitted     bool                 `json:"details_submitted"`
	PayoutsEnabled       bool                 `json:"payouts_enabled"`
	Name                 string               `json:"display_name"`
	Email                string               `json:"email"`
	ExternalAccounts     *ExternalAccountList `json:"external_accounts"`
	Statement            string               `json:"statement_descriptor"`
	PayoutStatement      string               `json:"payout_statement_descriptor"`
	Timezone             string               `json:"timezone"`
	BusinessName         string               `json:"business_name"`
	BusinessPrimaryColor string               `json:"business_primary_color"`
	BusinessUrl          string               `json:"business_url"`
	BusinessLogo         string               `json:"business_logo"`
	SupportPhone         string               `json:"support_phone"`
	SupportEmail         string               `json:"support_email"`
	SupportUrl           string               `json:"support_url"`
	ProductDesc          string               `json:"product_description"`
	DebitNegativeBal     bool                 `json:"debit_negative_balances"`
	Type                 AccountType
	Keys                 *struct {
		Secret  string `json:"secret"`
		Publish string `json:"publishable"`
	} `json:"keys"`
	Verification *struct {
		Fields         []string `json:"fields_needed"`
		Due            *int64   `json:"due_by"`
		DisabledReason string   `json:"disabled_reason"`
	} `json:"verification"`
	LegalEntity    *LegalEntity    `json:"legal_entity"`
	PayoutSchedule *PayoutSchedule `json:"payout_schedule"`
	TOSAcceptance  *struct {
		Date      int64  `json:"date"`
		IP        string `json:"ip"`
		UserAgent string `json:"user_agent"`
	} `json:"tos_acceptance"`
	SupportAddress *Address          `json:"support_address"`
	Deleted        bool              `json:"deleted"`
	Meta           map[string]string `json:"metadata"`
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
	Type ExternalAccountType `json:"object"`
	ID   string              `json:"id"`

	// A bank account attached to an account. Populated only if the external
	// account is a bank account.
	BankAccount *BankAccount

	// A card attached to an account. Populated only if the external account is
	// a card.
	Card *Card
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
	Type LegalEntityType `json:"type"`

	AdditionalOwners []Owner `json:"additional_owners"`

	// AdditionalOwnersEmpty can be set to clear a legal entity's additional
	// owners.
	AdditionalOwnersEmpty bool

	BusinessName          string               `json:"business_name"`
	BusinessNameKana      string               `json:"business_name_kana"`
	BusinessNameKanji     string               `json:"business_name_kanji"`
	Address               Address              `json:"address"`
	AddressKana           Address              `json:"address_kana"`
	AddressKanji          Address              `json:"address_kanji"`
	First                 string               `json:"first_name"`
	FirstKana             string               `json:"first_name_kana"`
	FirstKanji            string               `json:"first_name_kanji"`
	Gender                Gender               `json:"gender"`
	Last                  string               `json:"last_name"`
	LastKana              string               `json:"last_name_kana"`
	LastKanji             string               `json:"last_name_kanji"`
	MaidenName            string               `json:"maiden_name"`
	PersonalAddress       Address              `json:"personal_address"`
	PersonalAddressKana   Address              `json:"personal_address_kana"`
	PersonalAddressKanji  Address              `json:"personal_address_kanji"`
	PhoneNumber           string               `json:"phone_number"`
	DOB                   DOB                  `json:"dob"`
	Verification          IdentityVerification `json:"verification"`
	SSN                   string               `json:"ssn_last_4"`
	SSNProvided           bool                 `json:"ssn_last_4_provided"`
	PersonalID            string               `json:"personal_id_number"`
	PersonalIDProvided    bool                 `json:"personal_id_number_provided"`
	BusinessTaxID         string               `json:"business_tax_id"`
	BusinessTaxIDProvided bool                 `json:"business_tax_id_provided"`
	BusinessVatID         string               `json:"business_vat_id"`
}

// Address is the structure for an account address.
type Address struct {
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"postal_code"`
	Country string `json:"country"`

	// Town/cho-me. Note that this is only used for Kana/Kanji representations
	// of an address.
	Town string `json:"town"`
}

func (a *Address) AppendDetails(values *form.Values, prefix string) {
	if len(a.Line1) > 0 {
		values.Add(prefix+"[line1]", a.Line1)
	}

	if len(a.Line2) > 0 {
		values.Add(prefix+"[line2]", a.Line2)
	}

	if len(a.City) > 0 {
		values.Add(prefix+"[city]", a.City)
	}

	if len(a.State) > 0 {
		values.Add(prefix+"[state]", a.State)
	}

	if len(a.Zip) > 0 {
		values.Add(prefix+"[postal_code]", a.Zip)
	}

	if len(a.Country) > 0 {
		values.Add(prefix+"[country]", a.Country)
	}

	if len(a.Town) > 0 {
		values.Add(prefix+"[town]", a.Town)
	}
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

// Owner is the structure for an account owner.
type Owner struct {
	First        string               `json:"first_name"`
	Last         string               `json:"last_name"`
	DOB          DOB                  `json:"dob"`
	Address      Address              `json:"address"`
	Verification IdentityVerification `json:"verification"`
}

// IdentityVerification is the structure for an account's verification.
type IdentityVerification struct {
	DetailsCode IdentityVerificationDetailsCode `json:"details_code"`
	Status      IdentityVerificationStatus      `json:"status"`
	Document    *IdentityDocument               `json:"document"`
	Details     *string                         `json:"details"`
}

// IdentityDocument is the structure for an identity document.
type IdentityDocument struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	Size    int64  `json:"size"`
}

// PayoutSchedule is the structure for an account's payout schedule.
type PayoutSchedule struct {
	Delay       uint64   `json:"delay_days"`
	Interval    Interval `json:"interval"`
	WeekAnchor  string   `json:"weekly_anchor"`
	MonthAnchor uint64   `json:"monthly_anchor"`
}

// TOSAcceptanceParams is the structure for TOS acceptance.
type TOSAcceptanceParams struct {
	Date      int64  `json:"date"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

// AccountRejectParams is the structure for the Reject function.
type AccountRejectParams struct {
	Reason string `json:"reason"`
}

// AppendDetails adds the legal entity to the query string.
func (l *LegalEntity) AppendDetails(values *form.Values) {
	if len(l.Type) > 0 {
		values.Add("legal_entity[type]", string(l.Type))
	}

	if len(l.BusinessName) > 0 {
		values.Add("legal_entity[business_name]", l.BusinessName)
	}

	if len(l.BusinessNameKana) > 0 {
		values.Add("legal_entity[business_name_kana]", l.BusinessNameKana)
	}

	if len(l.BusinessNameKanji) > 0 {
		values.Add("legal_entity[business_name_kanji]", l.BusinessNameKanji)
	}

	if len(l.First) > 0 {
		values.Add("legal_entity[first_name]", l.First)
	}

	if len(l.FirstKana) > 0 {
		values.Add("legal_entity[first_name_kana]", l.FirstKana)
	}

	if len(l.FirstKanji) > 0 {
		values.Add("legal_entity[first_name_kanji]", l.FirstKanji)
	}

	if len(l.Gender) > 0 {
		values.Add("legal_entity[gender]", string(l.Gender))
	}

	if len(l.Last) > 0 {
		values.Add("legal_entity[last_name]", l.Last)
	}

	if len(l.LastKana) > 0 {
		values.Add("legal_entity[last_name_kana]", l.LastKana)
	}

	if len(l.LastKanji) > 0 {
		values.Add("legal_entity[last_name_kanji]", l.LastKanji)
	}

	if len(l.MaidenName) > 0 {
		values.Add("legal_entity[maiden_name]", l.MaidenName)
	}

	if l.DOB.Day > 0 {
		values.Add("legal_entity[dob][day]", strconv.Itoa(l.DOB.Day))
	}

	if l.DOB.Month > 0 {
		values.Add("legal_entity[dob][month]", strconv.Itoa(l.DOB.Month))
	}

	if l.DOB.Year > 0 {
		values.Add("legal_entity[dob][year]", strconv.Itoa(l.DOB.Year))
	}

	if len(l.SSN) > 0 {
		values.Add("legal_entity[ssn_last_4]", l.SSN)
	}

	if len(l.PersonalID) > 0 {
		values.Add("legal_entity[personal_id_number]", l.PersonalID)
	}

	if len(l.PhoneNumber) > 0 {
		values.Add("legal_entity[phone_number]", l.PhoneNumber)
	}

	if len(l.BusinessTaxID) > 0 {
		values.Add("legal_entity[business_tax_id]", l.BusinessTaxID)
	}

	if len(l.BusinessVatID) > 0 {
		values.Add("legal_entity[business_vat_id]", l.BusinessVatID)
	}

	if l.Verification.Document != nil {
		values.Add("legal_entity[verification][document]", l.Verification.Document.ID)
	}

	l.Address.AppendDetails(values, "legal_entity[address]")
	l.AddressKana.AppendDetails(values, "legal_entity[address_kana]")
	l.AddressKanji.AppendDetails(values, "legal_entity[address_kanji]")

	l.PersonalAddress.AppendDetails(values, "legal_entity[personal_address]")
	l.PersonalAddressKana.AppendDetails(values, "legal_entity[personal_address_kana]")
	l.PersonalAddressKanji.AppendDetails(values, "legal_entity[personal_address_kanji]")

	// sending an empty value unsets additional owners
	if l.AdditionalOwnersEmpty {
		values.Add("legal_entity[additional_owners]", "")
	} else {
		for i, owner := range l.AdditionalOwners {
			if len(owner.First) > 0 {
				values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][first_name]", i), owner.First)
			}

			if len(owner.Last) > 0 {
				values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][last_name]", i), owner.Last)
			}

			if owner.DOB.Day > 0 {
				values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][dob][day]", i), strconv.Itoa(owner.DOB.Day))
			}

			if owner.DOB.Month > 0 {
				values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][dob][month]", i), strconv.Itoa(owner.DOB.Month))
			}

			if owner.DOB.Year > 0 {
				values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][dob][year]", i), strconv.Itoa(owner.DOB.Year))
			}

			owner.Address.AppendDetails(values, fmt.Sprintf("legal_entity[additional_owners][%v][address]", i))

			if owner.Verification.Document != nil {
				values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][verification][document]", i), owner.Verification.Document.ID)
			}
		}
	}
}

// AppendDetails adds the payout schedule to the query string.
func (t *PayoutScheduleParams) AppendDetails(values *form.Values) {
	if t.Delay > 0 {
		values.Add("payout_schedule[delay_days]", strconv.FormatUint(t.Delay, 10))
	} else if t.MinimumDelay {
		values.Add("payout_schedule[delay_days]", "minimum")
	}

	values.Add("payout_schedule[interval]", string(t.Interval))
	if t.Interval == Week && len(t.WeekAnchor) > 0 {
		values.Add("payout_schedule[weekly_anchor]", t.WeekAnchor)
	} else if t.Interval == Month && t.MonthAnchor > 0 {
		values.Add("payout_schedule[monthly_anchor]", strconv.FormatUint(t.MonthAnchor, 10))
	}
}

// AppendDetails adds the terms of service acceptance to the query string.
func (t *TOSAcceptanceParams) AppendDetails(values *form.Values) {
	if t.Date > 0 {
		values.Add("tos_acceptance[date]", strconv.FormatInt(t.Date, 10))
	}
	if len(t.IP) > 0 {
		values.Add("tos_acceptance[ip]", t.IP)
	}
	if len(t.UserAgent) > 0 {
		values.Add("tos_acceptance[user_agent]", t.UserAgent)
	}
}

// UnmarshalJSON handles deserialization of an Account.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (a *Account) UnmarshalJSON(data []byte) error {
	type account Account
	var aa account
	err := json.Unmarshal(data, &aa)

	if err == nil {
		*a = Account(aa)
	} else {
		// the id is surrounded by "\" characters, so strip them
		a.ID = string(data[1 : len(data)-1])
	}

	return nil
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
