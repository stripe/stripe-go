package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// LegalEntityType describes the types for a legal entity.
// Current values are "individual", "company".
type LegalEntityType string

// IdentityVerificationStatus describes the different statuses for identity verification.
// Current values are "pending", "verified", "unverified".
type IdentityVerificationStatus string

// Interval describes the payout interval.
// Current values are "manual", "daily", "weekly", "monthly".
type Interval string

const (
	Individual LegalEntityType = "individual"
	Company    LegalEntityType = "company"

	IdentityVerificationPending    IdentityVerificationStatus = "pending"
	IdentityVerificationVerified   IdentityVerificationStatus = "verified"
	IdentityVerificationUnverified IdentityVerificationStatus = "unverified"

	Manual Interval = "manual"
	Day    Interval = "daily"
	Week   Interval = "weekly"
	Month  Interval = "monthly"
)

// AccountParams are the parameters allowed during account creation/updates.
type AccountParams struct {
	Params
	Country, Email, DefaultCurrency, Statement, BusinessName, BusinessUrl,
	BusinessPrimaryColor, SupportPhone, SupportEmail, SupportUrl string
	ExternalAccount           *AccountExternalAccountParams
	LegalEntity               *LegalEntity
	TransferSchedule          *TransferScheduleParams
	Managed, DebitNegativeBal bool
	TOSAcceptance             *TOSAcceptanceParams
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

// TransferScheduleParams are the parameters allowed for transfer schedules.
type TransferScheduleParams struct {
	Delay, MonthAnchor uint64
	WeekAnchor         string
	Interval           Interval
	MinimumDelay       bool
}

// Account is the resource representing your Stripe account.
// For more details see https://stripe.com/docs/api/#account.
type Account struct {
	ID             string `json:"id"`
	ChargesEnabled bool   `json:"charges_enabled"`
	Country        string `json:"country"`
	// Currencies is the list of supported currencies.
	Currencies           []string `json:"currencies_supported"`
	DefaultCurrency      string   `json:"default_currency"`
	DetailsSubmitted     bool     `json:"details_submitted"`
	TransfersEnabled     bool     `json:"transfers_enabled"`
	Name                 string   `json:"display_name"`
	Email                string   `json:"email"`
	Statement            string   `json:"statement_descriptor"`
	Timezone             string   `json:"timezone"`
	BusinessName         string   `json:"business_name"`
	BusinessPrimaryColor string   `json:"business_primary_color"`
	BusinessUrl          string   `json:"business_url"`
	SupportPhone         string   `json:"support_phone"`
	SupportEmail         string   `json:"support_email"`
	SupportUrl           string   `json:"support_url"`
	ProductDesc          string   `json:"product_description"`
	Managed              bool     `json:"managed"`
	DebitNegativeBal     bool     `json:"debit_negative_balances"`
	Keys                 *struct {
		Secret  string `json:"secret"`
		Publish string `json:"publishable"`
	} `json:"keys"`
	Verification *struct {
		Fields         []string `json:"fields_needed"`
		Due            *int64   `json:"due_by"`
		DisabledReason string   `json:"disabled_reason"`
	} `json:"verification"`
	LegalEntity      *LegalEntity      `json:"legal_entity"`
	TransferSchedule *TransferSchedule `json:"transfer_schedule"`
	BankAccounts     *BankAccountList  `json:"bank_accounts"`
	TOSAcceptance    *struct {
		Date      int64  `json:"date"`
		IP        string `json:"ip"`
		UserAgent string `json:"user_agent"`
	} `json:"tos_acceptance"`
	SupportAddress *Address `json:"support_address"`
	Deleted        bool     `json:"deleted"`
}

// LegalEntity is the structure for properties related to an account's legal state.
type LegalEntity struct {
	Type                  LegalEntityType      `json:"type"`
	BusinessName          string               `json:"business_name"`
	Address               Address              `json:"address"`
	First                 string               `json:"first_name"`
	Last                  string               `json:"last_name"`
	PersonalAddress       Address              `json:"personal_address"`
	DOB                   DOB                  `json:"dob"`
	AdditionalOwners      []Owner              `json:"additional_owners"`
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
}

// DOB is a structure for an account owner's date of birth.
type DOB struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

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
	Status   IdentityVerificationStatus `json:"status"`
	Document *IdentityDocument          `json:"document"`
	Details  *string                    `json:"details"`
}

// IdentityDocument is the structure for an identity document.
type IdentityDocument struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	Size    int64  `json:"size"`
}

// TransferSchedule is the structure for an account's transfer schedule.
type TransferSchedule struct {
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

// AppendDetails adds the legal entity to the query string.
func (l *LegalEntity) AppendDetails(values *url.Values) {
	values.Add("legal_entity[type]", string(l.Type))

	if len(l.BusinessName) > 0 {
		values.Add("legal_entity[business_name]", l.BusinessName)
	}

	if len(l.First) > 0 {
		values.Add("legal_entity[first_name]", l.First)
	}

	if len(l.Last) > 0 {
		values.Add("legal_entity[last_name]", l.Last)
	}

	values.Add("legal_entity[dob][day]", strconv.Itoa(l.DOB.Day))
	values.Add("legal_entity[dob][month]", strconv.Itoa(l.DOB.Month))
	values.Add("legal_entity[dob][year]", strconv.Itoa(l.DOB.Year))

	if len(l.SSN) > 0 {
		values.Add("legal_entity[ssn_last_4]", l.SSN)
	}

	if len(l.PersonalID) > 0 {
		values.Add("legal_entity[personal_id_number]", l.PersonalID)
	}

	if len(l.BusinessTaxID) > 0 {
		values.Add("legal_entity[business_tax_id]", l.BusinessTaxID)
	}

	if len(l.BusinessVatID) > 0 {
		values.Add("legal_entity[business_vat_id]", l.BusinessVatID)
	}

	if len(l.Address.Line1) > 0 {
		values.Add("legal_entity[address][line1]", l.Address.Line1)
	}

	if len(l.Address.Line2) > 0 {
		values.Add("legal_entity[address][line2]", l.Address.Line2)
	}

	if len(l.Address.City) > 0 {
		values.Add("legal_entity[address][city]", l.Address.City)
	}

	if len(l.Address.State) > 0 {
		values.Add("legal_entity[address][state]", l.Address.State)
	}

	if len(l.Address.Zip) > 0 {
		values.Add("legal_entity[address][postal_code]", l.Address.Zip)
	}

	if len(l.Address.Country) > 0 {
		values.Add("legal_entity[address][country]", l.Address.Country)
	}

	if len(l.PersonalAddress.Line1) > 0 {
		values.Add("legal_entity[personal_address][line1]", l.PersonalAddress.Line1)
	}

	if len(l.PersonalAddress.Line2) > 0 {
		values.Add("legal_entity[personal_address][line2]", l.PersonalAddress.Line2)
	}

	if len(l.PersonalAddress.City) > 0 {
		values.Add("legal_entity[personal_address][city]", l.PersonalAddress.City)
	}

	if len(l.PersonalAddress.State) > 0 {
		values.Add("legal_entity[personal_address][state]", l.PersonalAddress.State)
	}

	if len(l.PersonalAddress.Zip) > 0 {
		values.Add("legal_entity[personal_address][postal_code]", l.PersonalAddress.Zip)
	}

	if len(l.PersonalAddress.Country) > 0 {
		values.Add("legal_entity[personal_address][country]", l.PersonalAddress.Country)
	}

	for i, owner := range l.AdditionalOwners {
		if len(owner.First) > 0 {
			values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][first_name]", i), owner.First)
		}

		if len(owner.Last) > 0 {
			values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][last_name]", i), owner.Last)
		}

		values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][dob][day]", i), strconv.Itoa(owner.DOB.Day))
		values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][dob][month]", i), strconv.Itoa(owner.DOB.Month))
		values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][dob][year]", i), strconv.Itoa(owner.DOB.Year))

		if len(owner.Address.Line1) > 0 {
			values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][address][line1]", i), owner.Address.Line1)
		}

		if len(owner.Address.Line2) > 0 {
			values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][address][line2]", i), owner.Address.Line2)
		}

		if len(owner.Address.City) > 0 {
			values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][address][city]", i), owner.Address.City)
		}

		if len(owner.Address.State) > 0 {
			values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][address][state]", i), owner.Address.State)
		}

		if len(owner.Address.Zip) > 0 {
			values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][address][postal_code]", i), owner.Address.Zip)
		}

		if len(owner.Address.Country) > 0 {
			values.Add(fmt.Sprintf("legal_entity[additional_owners][%v][address][country]", i), owner.Address.Country)
		}
	}
}

// AppendDetails adds the transfer schedule to the query string.
func (t *TransferScheduleParams) AppendDetails(values *url.Values) {
	if t.Delay > 0 {
		values.Add("transfer_schedule[delay_days]", strconv.FormatUint(t.Delay, 10))
	} else if t.MinimumDelay {
		values.Add("transfer_schedule[delay_days]", "minimum")
	}

	values.Add("transfer_schedule[interval]", string(t.Interval))
	if t.Interval == Week && len(t.WeekAnchor) > 0 {
		values.Add("transfer_schedule[weekly_anchor]", t.WeekAnchor)
	} else if t.Interval == Month && t.MonthAnchor > 0 {
		values.Add("transfer_schedule[monthly_anchor]", strconv.FormatUint(t.MonthAnchor, 10))
	}
}

// AppendDetails adds the terms of service acceptance to the query string.
func (t *TOSAcceptanceParams) AppendDetails(values *url.Values) {
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
