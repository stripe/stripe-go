// Package customer provides the /customes APIs
package customer

import (
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /customers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new customers.
// For more details see https://stripe.com/docs/api#create_customer.
func New(params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.CustomerParams) (*stripe.Customer, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &url.Values{}
		if params.Balance != 0 {
			body.Add("account_balance", strconv.FormatInt(params.Balance, 10))
		}

		if len(params.Token) > 0 {
			body.Add("card", params.Token)
		} else if params.Card != nil {
			params.Card.AppendDetails(body, true)
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
		commonParams = &params.Params

		params.AppendTo(body)
	}

	cust := &stripe.Customer{}
	err := c.B.Call("POST", "/customers", c.Key, body, commonParams, cust)

	return cust, err
}

// Get returns the details of a customer.
// For more details see https://stripe.com/docs/api#retrieve_customer.
func Get(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &url.Values{}
		commonParams = &params.Params
		params.AppendTo(body)
	}

	cust := &stripe.Customer{}
	err := c.B.Call("GET", "/customers/"+id, c.Key, body, commonParams, cust)

	return cust, err
}

// Update updates a customer's properties.
// For more details see	https://stripe.com/docs/api#update_customer.
func Update(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
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
				params.Card.AppendDetails(body, true)
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
		params.AppendTo(body)
	}

	cust := &stripe.Customer{}
	err := c.B.Call("POST", "/customers/"+id, c.Key, body, commonParams, cust)

	return cust, err
}

// Del removes a customer.
// For more details see https://stripe.com/docs/api#delete_customer.
func Del(id string) error {
	return getC().Del(id)
}

func (c Client) Del(id string) error {
	return c.B.Call("DELETE", "/customers/"+id, c.Key, nil, nil, nil)
}

// List returns a list of customers.
// For more details see https://stripe.com/docs/api#list_customers.
func List(params *stripe.CustomerListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.CustomerListParams) *Iter {
	type customerList struct {
		stripe.ListMeta
		Values []*stripe.Customer `json:"data"`
	}

	var body *url.Values
	var lp *stripe.ListParams

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &customerList{}
		err := c.B.Call("GET", "/customers", c.Key, &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Customers.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Customer returns the most recent Customer
// visited by a call to Next.
func (i *Iter) Customer() *stripe.Customer {
	return i.Current().(*stripe.Customer)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
