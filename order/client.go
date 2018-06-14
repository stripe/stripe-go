package order

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /orders APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new order.
// For more details see https://stripe.com/docs/api#create_order.
func New(params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().New(params)
}

// New POSTs a new order.
// For more details see https://stripe.com/docs/api#create_order.
func (c Client) New(params *stripe.OrderParams) (*stripe.Order, error) {
	p := &stripe.Order{}
	err := c.B.Call("POST", "/orders", c.Key, params, p)
	return p, err
}

// Update updates an order's properties.
// For more details see https://stripe.com/docs/api#update_order.
func Update(id string, params *stripe.OrderUpdateParams) (*stripe.Order, error) {
	return getC().Update(id, params)
}

// Update updates an order's properties.
// For more details see https://stripe.com/docs/api#update_order.
func (c Client) Update(id string, params *stripe.OrderUpdateParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/orders/%s", id)
	o := &stripe.Order{}
	err := c.B.Call("POST", path, c.Key, params, o)
	return o, err
}

// Pay pays an order
// For more details see https://stripe.com/docs/api#pay_order.
func Pay(id string, params *stripe.OrderPayParams) (*stripe.Order, error) {
	return getC().Pay(id, params)
}

// Pay pays an order
// For more details see https://stripe.com/docs/api#pay_order.
func (c Client) Pay(id string, params *stripe.OrderPayParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/orders/%s/pay", id)
	o := &stripe.Order{}
	err := c.B.Call("POST", path, c.Key, params, o)
	return o, err
}

// Get returns the details of an order
// For more details see https://stripe.com/docs/api#retrieve_order.
func Get(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/orders/%s", id)
	order := &stripe.Order{}
	err := c.B.Call("GET", path, c.Key, params, order)
	return order, err
}

// List returns a list of orders.
// For more details see https://stripe.com/docs/api#list_orders
func List(params *stripe.OrderListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.OrderListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.OrderList{}
		err := c.B.CallRaw("GET", "/orders", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Orders.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Order returns the most recent Order
// visited by a call to Next.
func (i *Iter) Order() *stripe.Order {
	return i.Current().(*stripe.Order)
}

// Return returns all or part of an order.
// For more details see https://stripe.com/docs/api#return_order.
func Return(id string, params *stripe.OrderReturnParams) (*stripe.OrderReturn, error) {
	return getC().Return(id, params)
}

// Return returns all or part of an order.
// For more details see https://stripe.com/docs/api#return_order.
func (c Client) Return(id string, params *stripe.OrderReturnParams) (*stripe.OrderReturn, error) {
	path := stripe.FormatURLPath("/orders/%s/returns", id)
	ret := &stripe.OrderReturn{}
	err := c.B.Call("POST", path, c.Key, params, ret)
	return ret, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
