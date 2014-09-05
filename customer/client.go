package customer

import (
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /customers APIs.
type Client struct {
	B   Backend
	Tok string
}

// Create POSTs new customers.
// For more details see https://stripe.com/docs/api#create_customer.
func Create(params *CustomerParams) (*Customer, error) {
	return getC().Create(params)
}

func (c Client) Create(params *CustomerParams) (*Customer, error) {
	var body *url.Values

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

		params.AppendTo(body)
	}

	cust := &Customer{}
	err := c.B.Call("POST", "/customers", c.Tok, body, cust)

	return cust, err
}

// Get returns the details of a customer.
// For more details see https://stripe.com/docs/api#retrieve_customer.
func Get(id string, params *CustomerParams) (*Customer, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *CustomerParams) (*Customer, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	cust := &Customer{}
	err := c.B.Call("GET", "/customers/"+id, c.Tok, body, cust)

	return cust, err
}

// Update updates a customer's properties.
// For more details see	https://stripe.com/docs/api#update_customer.
func Update(id string, params *CustomerParams) (*Customer, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *CustomerParams) (*Customer, error) {
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

	cust := &Customer{}
	err := c.B.Call("POST", "/customers/"+id, c.Tok, body, cust)

	return cust, err
}

// Delete removes a customer.
// For more details see https://stripe.com/docs/api#delete_customer.
func Delete(id string) error {
	return getC().Delete(id)
}

func (c Client) Delete(id string) error {
	return c.B.Call("DELETE", "/customers/"+id, c.Tok, nil, nil)
}

// List returns a list of customers.
// For more details see https://stripe.com/docs/api#list_customers.
func List(params *CustomerListParams) (*CustomerList, error) {
	return getC().List(params)
}

func (c Client) List(params *CustomerListParams) (*CustomerList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		params.AppendTo(body)
	}

	list := &CustomerList{}
	err := c.B.Call("GET", "/customers", c.Tok, body, list)

	return list, err
}

func getC() Client {
	return Client{GetBackend(), Key}
}
