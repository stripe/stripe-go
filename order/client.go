package order

import (
	"errors"
	"fmt"

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
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		commonParams = &params.Params
		form.AppendTo(body, params)
	}

	p := &stripe.Order{}
	err := c.B.Call("POST", "/orders", c.Key, body, commonParams, p)

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
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		commonParams = &params.Params
		form.AppendTo(body, params)
	}

	o := &stripe.Order{}
	err := c.B.Call("POST", "/orders/"+id, c.Key, body, commonParams, o)

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
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		if params.Source == nil && len(params.Customer) == 0 {
			err := errors.New("Invalid order pay params: either customer or a source must be set")
			return nil, err
		}

		body = &form.Values{}
		commonParams = &params.Params
		form.AppendTo(body, params)
	}

	o := &stripe.Order{}
	err := c.B.Call("POST", "/orders/"+id+"/pay", c.Key, body, commonParams, o)

	return o, err
}

// Get returns the details of an order
// For more details see https://stripe.com/docs/api#retrieve_order.
func Get(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		commonParams = &params.Params
		form.AppendTo(body, params)
	}

	order := &stripe.Order{}
	err := c.B.Call("GET", "/orders/"+id, c.Key, body, commonParams, order)
	return order, err
}

// List returns a list of orders.
// For more details see https://stripe.com/docs/api#list_orders
func List(params *stripe.OrderListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.OrderListParams) *Iter {
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
		list := &stripe.OrderList{}
		err := c.B.Call("GET", "/orders", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
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
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	ret := &stripe.OrderReturn{}
	err := c.B.Call("POST", fmt.Sprintf("/orders/%s/returns", id), c.Key, body, commonParams, ret)

	return ret, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
