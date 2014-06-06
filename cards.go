package stripe

import (
	"net/url"
)

type CardType string

const (
	Unknown    CardType = "Unknown"
	Visa       CardType = "Visa"
	Amex       CardType = "American Express"
	MasterCard CardType = "MasterCard"
	Discover   CardType = "Discover"
	JCB        CardType = "JCB"
	DinersClub CardType = "Diners Club"
)

type CardParams struct {
	Token                                         string
	Customer, Recipient                           string
	Name, Number, Month, Year, CVC                string
	Address1, Address2, City, State, Zip, Country string
}

type Card struct {
	Id          string   `json:"id"`
	Name        string   `json:"name,omitempty"`
	Type        CardType `json:"type"`
	Month       uint8    `json:"exp_month"`
	Year        uint16   `json:"exp_year"`
	Last4       string   `json:"last4"`
	Fingerprint string   `json:"fingerprint"`
	CardCountry string   `json:"country,omitempty"`
	Customer    string   `json:"customer,omitempty"`
	Recipient   string   `json:"recipient,omitempty"`
	Address1    string   `json:"address_line1,omitempty"`
	Address2    string   `json:"address_line2,omitempty"`
	Country     string   `json:"address_country,omitempty"`
	State       string   `json:"address_state,omitempty"`
	Zip         string   `json:"address_zip,omitempty"`
	City        string   `json:"address_city"`
	Line1Check  string   `json:"address_line1_check,omitempty"`
	ZipCheck    string   `json:"address_zip_check,omitempty"`
	CVCCheck    string   `json:"cvc_check,omitempty"`
}

func (c *CardParams) appendTo(values *url.Values) {
	values.Add("card[number]", c.Number)
	values.Add("card[exp_month]", c.Month)
	values.Add("card[exp_year]", c.Year)

	if len(c.Name) > 0 {
		values.Add("card[name]", c.Name)
	}

	if len(c.CVC) > 0 {
		values.Add("card[cvc]", c.CVC)
	}

	if len(c.Address1) > 0 {
		values.Add("card[address1]", c.Address1)
	}

	if len(c.Address2) > 0 {
		values.Add("card[address2]", c.Address2)
	}

	if len(c.City) > 0 {
		values.Add("card[address_city]", c.City)
	}

	if len(c.State) > 0 {
		values.Add("card[address_state", c.State)
	}

	if len(c.Zip) > 0 {
		values.Add("card[address_zip]", c.Zip)
	}

	if len(c.Country) > 0 {
		values.Add("card[address_country]", c.Country)
	}
}
