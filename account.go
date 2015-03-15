package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type LegalEntityType string

type IdentityVerificationStatus string

type Interval string

const (
	Individual LegalEntityType = "individual"
	Corp       LegalEntityType = "corporation"

	IdentityVerificationPending    IdentityVerificationStatus = "pending"
	IdentityVerificationVerified   IdentityVerificationStatus = "verified"
	IdentityVerificationUnverified IdentityVerificationStatus = "unverified"

	Manual Interval = "manual"
	Day    Interval = "daily"
	Week   Interval = "weekly"
	Month  Interval = "monthly"
)

type AccountParams struct {
	Params
	Country, Email, DefaultCurrency, Statement, BusinessName, SupportPhone string
	LegalEntity                                                            *LegalEntity
	Managed                                                                bool
}

type AccountListParams struct {
	ListParams
}

// Account is the resource representing youe Stripe account.
// For more details see https://stripe.com/docs/api/#account.
type Account struct {
	ID             string `json:"id"`
	ChargesEnabled bool   `json:"charges_enabled"`
	Country        string `json:"country"`
	// Currencies is the list of supported currencies.
	Currencies       []string `json:"currencies_supported"`
	DefaultCurrency  string   `json:"default_currency"`
	DetailsSubmitted bool     `json:"details_submitted"`
	TransfersEnabled bool     `json:"transfers_enabled"`
	Name             string   `json:"display_name"`
	Email            string   `json:"email"`
	Statement        string   `json:"statement_descriptor"`
	Timezone         string   `json:"timezone"`
	BusinessName     string   `json:"business_name"`
	SupportPhone     string   `json:"support_phone"`
	ProductDesc      string   `json:"product_description"`
	Managed          bool     `json:"managed"`
	DebitNegativeBal bool     `json:"debit_negative_balances"`
	Keys             *struct {
		Secret  string `json:"secret"`
		Publish string `json:"publishable"`
	} `json:"keys"`
	Verification *struct {
		Fields    []string `json:"fields_needed"`
		Due       *int64   `json:"due_by"`
		Contacted bool     `json:"contacted"`
	} `json:"verification"`
	LegalEntity      *LegalEntity      `json:"legal_entity"`
	TransferSchedule *TransferSchedule `json:"transfer_schedule"`
}

type LegalEntity struct {
	Type             LegalEntityType      `json:"type"`
	BusinessName     string               `json:"business_name"`
	Address          Address              `json:"address"`
	First            string               `json:"first_name"`
	Last             string               `json:"last_name"`
	PersonalAddress  Address              `json:"personal_address"`
	DOB              DOB                  `json:"dob"`
	AdditionalOwners []Owner              `json:"additional_owners"`
	Verification     IdentityVerification `json:"verification"`
	SSN              string               `json:"ssn_last_4"`
	PersonalID       string               `json:"personal_id_number"`
	BusinessTaxID    string               `json:"business_tax_id"`
	BusinessVatID    string               `json:"business_vat_id"`
}

type Address struct {
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"postal_code"`
	Country string `json:"country"`
}

type DOB struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Owner struct {
	First        string               `json:"first_name"`
	Last         string               `json:"last_name"`
	DOB          DOB                  `json:"dob"`
	Address      Address              `json:"address"`
	Verification IdentityVerification `json:"verification"`
}

type IdentityVerification struct {
	Status   IdentityVerificationStatus `json:"status"`
	Document *struct {
		ID      string `json:"id"`
		Created int64  `json:"created"`
		Size    int64  `json:"size"`
	} `json:"document"`
	Details *string `json:"details"`
}

type TransferSchedule struct {
	Delay    int      `json:"delay_days"`
	Interval Interval `json:"interval"`
}

type AccountList struct {
	ListMeta
	Values []*Account `json:"data"`
}

func (a *AccountParams) AppendDetails(values *url.Values) {
	if a.LegalEntity != nil {
		values.Add("legal_entity[type]", string(a.LegalEntity.Type))

		if len(a.LegalEntity.BusinessName) > 0 {
			values.Add("legal_entity[business_name]", a.LegalEntity.BusinessName)
		}

		if len(a.LegalEntity.First) > 0 {
			values.Add("legal_entity[first_name]", a.LegalEntity.First)
		}

		if len(a.LegalEntity.Last) > 0 {
			values.Add("legal_entity[last_name]", a.LegalEntity.Last)
		}

		values.Add("legal_entity[dob][day]", strconv.Itoa(a.LegalEntity.DOB.Day))
		values.Add("legal_entity[dob][month]", strconv.Itoa(a.LegalEntity.DOB.Month))
		values.Add("legal_entity[dob][year]", strconv.Itoa(a.LegalEntity.DOB.Year))

		if len(a.LegalEntity.SSN) > 0 {
			values.Add("legal_entity[ssn_last_4]", a.LegalEntity.SSN)
		}

		if len(a.LegalEntity.PersonalID) > 0 {
			values.Add("legal_entity[personal_id_number]", a.LegalEntity.PersonalID)
		}

		if len(a.LegalEntity.BusinessTaxID) > 0 {
			values.Add("legal_entity[business_tax_id]", a.LegalEntity.BusinessTaxID)
		}

		if len(a.LegalEntity.BusinessVatID) > 0 {
			values.Add("legal_entity[business_vat_id]", a.LegalEntity.BusinessVatID)
		}

		if len(a.LegalEntity.Address.Line1) > 0 {
			values.Add("legal_entity[address][line1]", a.LegalEntity.Address.Line1)
		}

		if len(a.LegalEntity.Address.Line2) > 0 {
			values.Add("legal_entity[address][line2]", a.LegalEntity.Address.Line2)
		}

		if len(a.LegalEntity.Address.City) > 0 {
			values.Add("legal_entity[address][city]", a.LegalEntity.Address.City)
		}

		if len(a.LegalEntity.Address.State) > 0 {
			values.Add("legal_entity[address][state]", a.LegalEntity.Address.State)
		}

		if len(a.LegalEntity.Address.Zip) > 0 {
			values.Add("legal_entity[address][postal_code]", a.LegalEntity.Address.Zip)
		}

		if len(a.LegalEntity.Address.Country) > 0 {
			values.Add("legal_entity[address][country]", a.LegalEntity.Address.Country)
		}

		if len(a.LegalEntity.PersonalAddress.Line1) > 0 {
			values.Add("legal_entity[personal_address][line1]", a.LegalEntity.PersonalAddress.Line1)
		}

		if len(a.LegalEntity.PersonalAddress.Line2) > 0 {
			values.Add("legal_entity[personal_address][line2]", a.LegalEntity.PersonalAddress.Line2)
		}

		if len(a.LegalEntity.PersonalAddress.City) > 0 {
			values.Add("legal_entity[personal_address][city]", a.LegalEntity.PersonalAddress.City)
		}

		if len(a.LegalEntity.PersonalAddress.State) > 0 {
			values.Add("legal_entity[personal_address][state]", a.LegalEntity.PersonalAddress.State)
		}

		if len(a.LegalEntity.PersonalAddress.Zip) > 0 {
			values.Add("legal_entity[personal_address][postal_code]", a.LegalEntity.PersonalAddress.Zip)
		}

		if len(a.LegalEntity.PersonalAddress.Country) > 0 {
			values.Add("legal_entity[personal_address][country]", a.LegalEntity.PersonalAddress.Country)
		}

		for i, owner := range a.LegalEntity.AdditionalOwners {
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
