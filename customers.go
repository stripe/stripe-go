package stripe

import (
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

type CustomerListParams struct {
	Created    int64
	Filters    Filters
	Start, End string
	Limit      uint64
}

type Customer struct {
	Id          string            `json:"id"`
	Live        bool              `json:"livemode"`
	Cards       *CardList         `json:"cards"`
	Created     int64             `json:"created"`
	Balance     int64             `json:"account_balance"`
	Currency    Currency          `json:"currency"`
	Delinquent  bool              `json:"delinquent"`
	DefaultCard string            `json:"default_card"`
	Desc        string            `json:"description"`
	Email       string            `json:"email"`
	Meta        map[string]string `json:"metadata"`
	Subs        *SubscriptionList `json:"subscriptions"`
	Discount    *Discount         `json:"discount"`
}

type CustomerList struct {
	Count  uint16      `json:"total_count"`
	More   bool        `json:"has_more"`
	Url    string      `json:"url"`
	Values []*Customer `json:"data"`
}

type CustomerClient struct {
	api   Api
	token string
}

func (c *CustomerClient) Create(params *CustomerParams) (*Customer, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
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
	}

	cust := &Customer{}
	err := c.api.Call("POST", "/customers", c.token, body, cust)

	return cust, err
}

func (c *CustomerClient) Get(id string) (*Customer, error) {
	cust := &Customer{}
	err := c.api.Call("GET", "/customers/"+id, c.token, nil, cust)

	return cust, err
}

func (c *CustomerClient) Update(id string, params *CustomerParams) (*Customer, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

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
	}

	cust := &Customer{}
	err := c.api.Call("POST", "/customers/"+id, c.token, body, cust)

	return cust, err
}

func (c *CustomerClient) Delete(id string) error {
	return c.api.Call("DELETE", "/customers/"+id, c.token, nil, nil)
}

func (c *CustomerClient) List(params *CustomerListParams) (*CustomerList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

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
	}

	list := &CustomerList{}
	err := c.api.Call("GET", "/customers", c.token, body, list)

	return list, err
}
