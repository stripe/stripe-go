//
//
// File generated from our OpenAPI spec
//
//

// Package order provides the /orders APIs
package order

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /orders APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new order.
func New(params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().New(params)
}

// New creates a new order.
func (c Client) New(params *stripe.OrderParams) (*stripe.Order, error) {
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, "/v1/orders", c.Key, params, order)
	return order, err
}

// Get returns the details of an order.
func Get(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().Get(id, params)
}

// Get returns the details of an order.
func (c Client) Get(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, order)
	return order, err
}

// Update updates an order's properties.
func Update(id string, params *stripe.OrderUpdateParams) (*stripe.Order, error) {
	return getC().Update(id, params)
}

// Update updates an order's properties.
func (c Client) Update(id string, params *stripe.OrderUpdateParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Pay is the method for the `POST /v1/orders/{id}/pay` API.
func Pay(id string, params *stripe.OrderPayParams) (*stripe.Order, error) {
	return getC().Pay(id, params)
}

// Pay is the method for the `POST /v1/orders/{id}/pay` API.
func (c Client) Pay(id string, params *stripe.OrderPayParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s/pay", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Return is the method for the `POST /v1/orders/{id}/returns` API.
func Return(id string, params *stripe.OrderReturnParams) (*stripe.OrderReturn, error) {
	return getC().Return(id, params)
}

// Return is the method for the `POST /v1/orders/{id}/returns` API.
func (c Client) Return(id string, params *stripe.OrderReturnParams) (*stripe.OrderReturn, error) {
	path := stripe.FormatURLPath("/v1/orders/%s/returns", id)
	order := &stripe.OrderReturn{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// List returns a list of orders.
func List(params *stripe.OrderListParams) *Iter {
	return getC().List(params)
}

// List returns a list of orders.
func (c Client) List(listParams *stripe.OrderListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.OrderList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/orders", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for orders.
type Iter struct {
	*stripe.Iter
}

// Order returns the order which the iterator is currently pointing to.
func (i *Iter) Order() *stripe.Order {
	return i.Current().(*stripe.Order)
}

// OrderList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) OrderList() *stripe.OrderList {
	return i.List().(*stripe.OrderList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
