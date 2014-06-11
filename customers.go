package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type CustomerParams struct {
	Balance       int64
	Token, Coupon string
	Card          *CardParams
	Desc, Email   string
	Meta          map[string]string
	Plan          string
	Quantity      uint64
	End           int64
	DefaultCard   string
}

type Customer struct {
	Id            string            `json:"id"`
	Live          bool              `json:"livemode"`
	Cards         *CardList         `json:"cards"`
	Created       int64             `json:"created"`
	Balance       int64             `json:"account_balance"`
	Currency      Currency          `json:"currency"`
	Delinquent    bool              `json:"delinquent"`
	DefaultCard   string            `json:"default_card,omitempty"`
	Desc          string            `json:"description,omitempty"`
	Email         string            `json:"email,omitempty"`
	Meta          map[string]string `json:"metadata,omitempty"`
	Subscriptions *SubscriptionList `json:"subscriptions,omitempty"`
	//Discount
}

type CustomerClient struct {
	api   Api
	token string
}

func (c *CustomerClient) Create(params *CustomerParams) (*Customer, error) {
	body := &url.Values{}

	if params.Balance != 0 {
		body.Add("account_balance", strconv.FormatInt(params.Balance, 10))
	}

	if len(params.Token) > 0 {
		body.Add("card", params.Token)
	} else if params.Card != nil {
		params.Card.appendTo(body, true)
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	if len(params.Coupon) > 0 {
		body.Add("coupon", params.Coupon)
	}

	if len(params.Email) > 0 {
		body.Add("email", params.Email)
	}

	if len(params.Plan) > 0 {
		body.Add("plan", params.Plan)

		if params.Quantity > 0 {
			body.Add("quantity", strconv.FormatUint(params.Quantity, 10))
		}

		if params.End > 0 {
			body.Add("trial_end", strconv.FormatInt(params.End, 10))
		}

	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	res, err := c.api.Call("POST", "/customers", c.token, body)
	if err != nil {
		return nil, err
	}

	var cust Customer
	err = json.Unmarshal(res, &cust)
	return &cust, err
}

func (c *CustomerClient) Get(id string) (*Customer, error) {
	res, err := c.api.Call("GET", "/customers/"+id, c.token, nil)
	if err != nil {
		return nil, err
	}

	var cust Customer
	err = json.Unmarshal(res, &cust)
	return &cust, err
}

func (c *CustomerClient) Update(id string, params *CustomerParams) (*Customer, error) {
	body := &url.Values{}

	if params.Balance != 0 {
		body.Add("account_balance", strconv.FormatInt(params.Balance, 10))
	}

	if len(params.Token) > 0 {
		body.Add("card", params.Token)
	} else if params.Card != nil {
		params.Card.appendTo(body, true)
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	if len(params.Coupon) > 0 {
		body.Add("coupon", params.Coupon)
	}

	if len(params.Email) > 0 {
		body.Add("email", params.Email)
	}

	if len(params.DefaultCard) > 0 {
		body.Add("default_card", params.DefaultCard)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	res, err := c.api.Call("POST", "/customers/"+id, c.token, body)
	if err != nil {
		return nil, err
	}

	var cust Customer
	err = json.Unmarshal(res, &cust)
	return &cust, err
}

func (c *CustomerClient) Delete(id string) error {
	_, err := c.api.Call("DELETE", "/customers/"+id, c.token, nil)
	return err
}
