package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// CustomerParams is the set of parameters that can be used when creating or updating a customer.
// For more details see https://stripe.com/docs/api#create_customer and https://stripe.com/docs/api#update_customer.
type CustomerParams struct {
	Params
	Balance       int64
	Token, Coupon string
	Card          *CardParams
	Desc, Email   string
	Plan          string
	Quantity      uint64
	TrialEnd      int64
	DefaultCard   string
}

// CustomerListParams is the set of parameters that can be used when listing customers.
// For more details see https://stripe.com/docs/api#list_customers.
type CustomerListParams struct {
	ListParams
	Created int64
}

// Customer is the resource representing a Stripe customer.
// For more details see https://stripe.com/docs/api#customers.
type Customer struct {
	Id          string            `json:"id"`
	Live        bool              `json:"livemode"`
	Cards       *CardList         `json:"cards"`
	Created     int64             `json:"created"`
	Balance     int64             `json:"account_balance"`
	Currency    Currency          `json:"currency"`
	DefaultCard *Card             `json:"default_card"`
	Delinquent  bool              `json:"delinquent"`
	Desc        string            `json:"description"`
	Discount    *Discount         `json:"discount"`
	Email       string            `json:"email"`
	Meta        map[string]string `json:"metadata"`
	Subs        *SubscriptionList `json:"subscriptions"`
}

// CustomerList is a list object for customers.
type CustomerList struct {
	ListResponse
	Values []*Customer `json:"data"`
}

// CustomerClient is the client used to invoke /customers APIs.
type CustomerClient struct {
	api   Api
	token string
}

// Create POSTs new customers.
// For more details see https://stripe.com/docs/api#create_customer.
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

			if params.TrialEnd > 0 {
				body.Add("trial_end", strconv.FormatInt(params.TrialEnd, 10))
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

// Get returns the details of a customer.
// For more details see https://stripe.com/docs/api#retrieve_customer.
func (c *CustomerClient) Get(id string) (*Customer, error) {
	cust := &Customer{}
	err := c.api.Call("GET", "/customers/"+id, c.token, nil, cust)

	return cust, err
}

// Update updates a customer's properties.
// For more details see	https://stripe.com/docs/api#update_customer.
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
			if len(params.Card.Token) > 0 {
				body.Add("card", params.Card.Token)
			} else {
				params.Card.appendTo(body, true)
			}
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

// Delete removes a customer.
// For more details see https://stripe.com/docs/api#delete_customer.
func (c *CustomerClient) Delete(id string) error {
	return c.api.Call("DELETE", "/customers/"+id, c.token, nil, nil)
}

// List returns a list of customers.
// For more details see https://stripe.com/docs/api#list_customers.
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

func (c *Customer) UnmarshalJSON(data []byte) error {
	type customer Customer
	var cc customer
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = Customer(cc)
	} else {
		c.Id = string(data[1 : len(data)-1])
	}

	return nil
}
