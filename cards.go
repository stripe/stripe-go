package stripe

import (
	"encoding/json"
	"errors"
	"fmt"
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
	LastFour    string   `json:"last4"`
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

	var res []byte
	var err error

	if len(params.Customer) > 0 {
		res, err = c.api.Call("POST", fmt.Sprintf("/customers/%v/cards", params.Customer), c.token, body)
	} else if len(params.Recipient) > 0 {
		res, err = c.api.Call("POST", fmt.Sprintf("/recipients/%v/cards", params.Recipient), c.token, body)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	if err != nil {
		return nil, err
	}

	var card Card
	err = json.Unmarshal(res, &card)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (c *CardClient) Get(id string, params *CardParams) (*Card, error) {
	var res []byte
	var err error

	if len(params.Customer) > 0 {
		res, err = c.api.Call("GET", fmt.Sprintf("/customers/%v/cards/%v", params.Customer, id), c.token, nil)
	} else if len(params.Recipient) > 0 {
		res, err = c.api.Call("GET", fmt.Sprintf("/recipients/%v/cards/%v", params.Recipient, id), c.token, nil)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	if err != nil {
		return nil, err
	}

	var card Card
	err = json.Unmarshal(res, &card)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (c *CardClient) Update(id string, params *CardParams) (*Card, error) {
	body := &url.Values{}
	params.appendTo(body, false)

	var res []byte
	var err error

	if len(params.Customer) > 0 {
		res, err = c.api.Call("POST", fmt.Sprintf("/customers/%v/cards/%v", params.Customer, id), c.token, body)
	} else if len(params.Recipient) > 0 {
		res, err = c.api.Call("POST", fmt.Sprintf("/recipients/%v/cards/%v", params.Recipient, id), c.token, body)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	if err != nil {
		return nil, err
	}

	var card Card
	err = json.Unmarshal(res, &card)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (c *CardClient) Delete(id string, params *CardParams) error {
	var err error

	if len(params.Customer) > 0 {
		_, err = c.api.Call("DELETE", fmt.Sprintf("/customers/%v/cards/%v", params.Customer, id), c.token, nil)
	} else if len(params.Recipient) > 0 {
		_, err = c.api.Call("DELETE", fmt.Sprintf("/recipients/%v/cards/%v", params.Recipient, id), c.token, nil)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	return err
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
