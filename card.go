package stripe

import (
	"encoding/json"

	"net/url"
)

// CardBrand is the list of allowed values for the card's brand.
// Allowed values are "Unknown", "Visa", "American Express", "MasterCard", "Discover"
// "JCB", "Diners Club".
type CardBrand string

// Verification is the list of allowed verification responses.
// Allowed values are "pass", "fail", "unchecked".
type Verification string

// CardFunding is the list of allowed values for the card's funding.
// Allowed values are "credit", "debit", "prepaid", "unknown".
type CardFunding string

const (
	Unknown    CardBrand = "Unknown"
	Visa       CardBrand = "Visa"
	Amex       CardBrand = "American Express"
	MasterCard CardBrand = "MasterCard"
	Discover   CardBrand = "Discover"
	JCB        CardBrand = "JCB"
	DinersClub CardBrand = "Diners Club"

	Pass      Verification = "pass"
	Fail      Verification = "fail"
	Unchecked Verification = "unchecked"

	CreditFunding  CardFunding = "credit"
	DebitFunding   CardFunding = "debit"
	PrepaidFunding CardFunding = "prepaid"
	UnknownFunding CardFunding = "unknown"
)

// CardParams is the set of parameters that can be used when creating or updating a card.
// For more details see https://stripe.com/docs/api#create_card and https://stripe.com/docs/api#update_card.
type CardParams struct {
	Params
	Token                                         string
	Customer, Recipient                           string
	Name, Number, Month, Year, CVC                string
	Address1, Address2, City, State, Zip, Country string
}

// CardListParams is the set of parameters that can be used when listing cards.
// For more details see https://stripe.com/docs/api#list_cards.
type CardListParams struct {
	ListParams
	Customer, Recipient string
}

// Card is the resource representing a Stripe credit/debit card.
// For more details see https://stripe.com/docs/api#cards.
type Card struct {
	Id            string       `json:"id"`
	Month         uint8        `json:"exp_month"`
	Year          uint16       `json:"exp_year"`
	Fingerprint   string       `json:"fingerprint"`
	Funding       CardFunding  `json:"funding"`
	LastFour      string       `json:"last4"`
	Brand         CardBrand    `json:"brand"`
	City          string       `json:"address_city"`
	Country       string       `json:"address_country"`
	Address1      string       `json:"address_line1"`
	Address1Check Verification `json:"address_line1_check"`
	Address2      string       `json:"address_line2"`
	State         string       `json:"address_state"`
	Zip           string       `json:"address_zip"`
	ZipCheck      Verification `json:"address_zip_check"`
	CardCountry   string       `json:"country"`
	Customer      *Customer    `json:"customer"`
	CVCCheck      Verification `json:"cvc_check"`
	Name          string       `json:"name"`
	Recipient     *Recipient   `json:"recipient"`
}

// CardList is a list object for cards.
type CardList struct {
	ListMeta
	Values []*Card `json:"data"`
}

// CardIter is a iterator for list responses.
type CardIter struct {
	Iter *Iter
}

// Next returns the next value in the list.
func (i *CardIter) Next() (*Card, error) {
	c, err := i.Iter.Next()
	if err != nil {
		return nil, err
	}

	return c.(*Card), err
}

// Stop returns true if there are no more iterations to be performed.
func (i *CardIter) Stop() bool {
	return i.Iter.Stop()
}

// Meta returns the list metadata.
func (i *CardIter) Meta() *ListMeta {
	return i.Iter.Meta()
}

// AppendDetails adds the card's details to the query string values.
// When creating a new card, the parameters are passed as a dictionary, but
// on updates they are simply the parameter name.
func (c *CardParams) AppendDetails(values *url.Values, creating bool) {
	if creating {
		if len(c.Token) > 0 {
			values.Add("card", c.Token)
		} else {
			values.Add("card[number]", c.Number)
			values.Add("card[exp_month]", c.Month)
			values.Add("card[exp_year]", c.Year)

			if len(c.CVC) > 0 {
				values.Add("card[cvc]", c.CVC)
			}
		}
	}

	if len(c.Name) > 0 {
		if creating {
			values.Add("card[name]", c.Name)
		} else {
			values.Add("name", c.Name)
		}
	}

	if len(c.Address1) > 0 {
		if creating {
			values.Add("card[address_line1]", c.Address1)
		} else {
			values.Add("address_line1", c.Address1)
		}
	}

	if len(c.Address2) > 0 {
		if creating {
			values.Add("card[address_line2]", c.Address2)
		} else {
			values.Add("address_line2", c.Address2)
		}
	}

	if len(c.City) > 0 {
		if creating {
			values.Add("card[address_city]", c.City)
		} else {
			values.Add("address_city", c.City)
		}
	}

	if len(c.State) > 0 {
		if creating {
			values.Add("card[address_state]", c.State)
		} else {
			values.Add("address_state", c.State)
		}
	}

	if len(c.Zip) > 0 {
		if creating {
			values.Add("card[address_zip]", c.Zip)
		} else {
			values.Add("address_zip", c.Zip)
		}
	}

	if len(c.Country) > 0 {
		if creating {
			values.Add("card[address_country]", c.Country)
		} else {
			values.Add("address_country", c.Country)
		}
	}
}

func (c *Card) UnmarshalJSON(data []byte) error {
	type card Card
	var cc card
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = Card(cc)
	} else {
		// the id is surrounded by escaped \, so ignore those
		c.Id = string(data[1 : len(data)-1])
	}

	return nil
}
