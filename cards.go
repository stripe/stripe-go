package stripe

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

type CardType string
type Verification string

const (
	Unknown    CardType = "Unknown"
	Visa       CardType = "Visa"
	Amex       CardType = "American Express"
	MasterCard CardType = "MasterCard"
	Discover   CardType = "Discover"
	JCB        CardType = "JCB"
	DinersClub CardType = "Diners Club"

	Pass      Verification = "pass"
	Fail      Verification = "fail"
	Unchecked Verification = "unchecked"
)

type CardParams struct {
	Token                                         string
	Customer, Recipient                           string
	Name, Number, Month, Year, CVC                string
	Address1, Address2, City, State, Zip, Country string
}

type CardListParams struct {
	Customer, Recipient string
	Filters             Filters
	Start, End          string
	Limit               uint64
}

type Card struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Type        CardType     `json:"type"`
	Month       uint8        `json:"exp_month"`
	Year        uint16       `json:"exp_year"`
	LastFour    string       `json:"last4"`
	Fingerprint string       `json:"fingerprint"`
	CardCountry string       `json:"country"`
	Customer    string       `json:"customer"`
	Recipient   string       `json:"recipient"`
	Address1    string       `json:"address_line1"`
	Address2    string       `json:"address_line2"`
	Country     string       `json:"address_country"`
	State       string       `json:"address_state"`
	Zip         string       `json:"address_zip"`
	City        string       `json:"address_city"`
	Line1Check  Verification `json:"address_line1_check"`
	ZipCheck    Verification `json:"address_zip_check"`
	CVCCheck    Verification `json:"cvc_check"`
}

type CardList struct {
	Count  uint16  `json:"total_count"`
	More   bool    `json:"has_more"`
	Url    string  `json:"url"`
	Values []*Card `json:"data"`
}

type CardClient struct {
	api   Api
	token string
}

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

func (c *CardClient) Delete(id string, params *CardParams) error {
	if len(params.Customer) > 0 {
		return c.api.Call("DELETE", fmt.Sprintf("/customers/%v/cards/%v", params.Customer, id), c.token, nil, nil)
	} else if len(params.Recipient) > 0 {
		return c.api.Call("DELETE", fmt.Sprintf("/recipients/%v/cards/%v", params.Recipient, id), c.token, nil, nil)
	} else {
		return errors.New("Invalid card params: either customer or recipient need to be set")
	}
}

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

func (c *CardParams) appendTo(values *url.Values, creating bool) {
	if creating {
		values.Add("card[number]", c.Number)
		values.Add("card[exp_month]", c.Month)
		values.Add("card[exp_year]", c.Year)

		if len(c.CVC) > 0 {
			values.Add("card[cvc]", c.CVC)
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
