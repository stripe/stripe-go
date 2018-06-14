// Package customer provides the /customers APIs
package customer

import (
	"net/http"

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
	cust := &stripe.Customer{}
	err := c.B.Call(http.MethodPost, "/customers", c.Key, params, cust)
	return cust, err
}

// Get returns the details of a customer.
// For more details see https://stripe.com/docs/api#retrieve_customer.
func Get(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	path := stripe.FormatURLPath("/customers/%s", id)
	cust := &stripe.Customer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cust)
	return cust, err
}

// Update updates a customer's properties.
// For more details see	https://stripe.com/docs/api#update_customer.
func Update(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	path := stripe.FormatURLPath("/customers/%s", id)
	cust := &stripe.Customer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cust)
	return cust, err
}

// Del removes a customer.
// For more details see https://stripe.com/docs/api#delete_customer.
func Del(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	path := stripe.FormatURLPath("/customers/%s", id)
	cust := &stripe.Customer{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, cust)
	return cust, err
}

// List returns a list of customers.
// For more details see https://stripe.com/docs/api#list_customers.
func List(params *stripe.CustomerListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.CustomerListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.CustomerList{}
		err := c.B.CallRaw(http.MethodGet, "/customers", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
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
