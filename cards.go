package stripe

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
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
	ListResponse
	Values []*Card `json:"data"`
}

// CardClient is the client used to invoke /cards APIs.
type CardClient struct {
	api   Api
	token string
}

// Create POSTs new cards either for a customer or recipient.
// For more details see https://stripe.com/docs/api#create_card.
func (c *CardClient) Create(params *CardParams) (*Card, error) {
	body := &url.Values{}
	params.appendTo(body, true)

	card := &Card{}
	var err error

	if len(params.Customer) > 0 {
		err = c.api.Call("POST", fmt.Sprintf("/customers/%v/cards", params.Customer), c.token, body, card)
	} else if len(params.Recipient) > 0 {
		err = c.api.Call("POST", fmt.Sprintf("/recipients/%v/cards", params.Recipient), c.token, body, card)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	return card, err
}

// Get returns the details of a card.
// For more details see https://stripe.com/docs/api#retrieve_card.
func (c *CardClient) Get(id string, params *CardParams) (*Card, error) {
	card := &Card{}
	var err error

	if len(params.Customer) > 0 {
		err = c.api.Call("GET", fmt.Sprintf("/customers/%v/cards/%v", params.Customer, id), c.token, nil, card)
	} else if len(params.Recipient) > 0 {
		err = c.api.Call("GET", fmt.Sprintf("/recipients/%v/cards/%v", params.Recipient, id), c.token, nil, card)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	return card, err
}

// Update updates a card's properties.
// For more details see	https://stripe.com/docs/api#update_card.
func (c *CardClient) Update(id string, params *CardParams) (*Card, error) {
	body := &url.Values{}
	params.appendTo(body, false)

	card := &Card{}
	var err error

	if len(params.Customer) > 0 {
		err = c.api.Call("POST", fmt.Sprintf("/customers/%v/cards/%v", params.Customer, id), c.token, body, card)
	} else if len(params.Recipient) > 0 {
		err = c.api.Call("POST", fmt.Sprintf("/recipients/%v/cards/%v", params.Recipient, id), c.token, body, card)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	return card, err
}

// Delete remotes a card.
// For more details see https://stripe.com/docs/api#delete_card.
func (c *CardClient) Delete(id string, params *CardParams) error {
	if len(params.Customer) > 0 {
		return c.api.Call("DELETE", fmt.Sprintf("/customers/%v/cards/%v", params.Customer, id), c.token, nil, nil)
	} else if len(params.Recipient) > 0 {
		return c.api.Call("DELETE", fmt.Sprintf("/recipients/%v/cards/%v", params.Recipient, id), c.token, nil, nil)
	} else {
		return errors.New("Invalid card params: either customer or recipient need to be set")
	}
}

// List returns a list of cards.
// For more details see https://stripe.com/docs/api#list_cards.
func (c *CardClient) List(params *CardListParams) (*CardList, error) {
	body := &url.Values{}

	if len(params.Filters.f) > 0 {
		params.Filters.appendTo(body)
	}

	if len(params.Start) > 0 {
		body.Add("starting_after", params.Start)
	}

	if len(params.End) > 0 {
		body.Add("ending_before", params.End)
	}

	if params.Limit > 0 {
		if params.Limit > 100 {
			params.Limit = 100
		}

		body.Add("limit", strconv.FormatUint(params.Limit, 10))
	}

	list := &CardList{}
	var err error

	if len(params.Customer) > 0 {
		err = c.api.Call("GET", fmt.Sprintf("/customers/%v/cards", params.Customer), c.token, body, list)
	} else if len(params.Recipient) > 0 {
		err = c.api.Call("GET", fmt.Sprintf("/recipients/%v/cards", params.Recipient), c.token, body, list)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	return list, err
}

// appendTo adds the card's details to the querey string values.
// When creating a new card, the parameters are passed as a dictionary, but
// on updates they are simply the parameter name.
func (c *CardParams) appendTo(values *url.Values, creating bool) {
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
