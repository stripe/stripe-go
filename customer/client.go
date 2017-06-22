// Package customer provides the /customers APIs
package customer

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
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
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		commonParams = &params.Params
		form.AppendTo(body, params)
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
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		commonParams = &params.Params
		form.AppendTo(body, params)
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
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	cust := &stripe.Customer{}
	err := c.B.Call("POST", "/customers/"+id, c.Key, body, commonParams, cust)

	return cust, err
}

// Del removes a customer.
// For more details see https://stripe.com/docs/api#delete_customer.
func Del(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	cust := &stripe.Customer{}
	err := c.B.Call("DELETE", "/customers/"+id, c.Key, body, commonParams, cust)

	return cust, err
}

// List returns a list of customers.
// For more details see https://stripe.com/docs/api#list_customers.
func List(params *stripe.CustomerListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.CustomerListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.CustomerList{}
		err := c.B.Call("GET", "/customers", c.Key, b, p, list)

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
