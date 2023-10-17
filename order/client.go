//
//
// File generated from our OpenAPI spec
//
//

// Package order provides the /orders APIs
package order

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
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
func Update(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().Update(id, params)
}

// Update updates an order's properties.
func (c Client) Update(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Cancel is the method for the `POST /v1/orders/{id}/cancel` API.
func Cancel(id string, params *stripe.OrderCancelParams) (*stripe.Order, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/orders/{id}/cancel` API.
func (c Client) Cancel(id string, params *stripe.OrderCancelParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s/cancel", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Reopen is the method for the `POST /v1/orders/{id}/reopen` API.
func Reopen(id string, params *stripe.OrderReopenParams) (*stripe.Order, error) {
	return getC().Reopen(id, params)
}

// Reopen is the method for the `POST /v1/orders/{id}/reopen` API.
func (c Client) Reopen(id string, params *stripe.OrderReopenParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s/reopen", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Submit is the method for the `POST /v1/orders/{id}/submit` API.
func Submit(id string, params *stripe.OrderSubmitParams) (*stripe.Order, error) {
	return getC().Submit(id, params)
}

// Submit is the method for the `POST /v1/orders/{id}/submit` API.
func (c Client) Submit(id string, params *stripe.OrderSubmitParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s/submit", id)
	order := &stripe.Order{}
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

// ListLineItems is the method for the `GET /v1/orders/{id}/line_items` API.
func ListLineItems(params *stripe.OrderListLineItemsParams) *LineItemIter {
	return getC().ListLineItems(params)
}

// ListLineItems is the method for the `GET /v1/orders/{id}/line_items` API.
func (c Client) ListLineItems(listParams *stripe.OrderListLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/orders/%s/line_items",
		stripe.StringValue(listParams.ID),
	)
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.LineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// LineItemIter is an iterator for line items.
type LineItemIter struct {
	*stripe.Iter
}

// LineItem returns the line item which the iterator is currently pointing to.
func (i *LineItemIter) LineItem() *stripe.LineItem {
	return i.Current().(*stripe.LineItem)
}

// LineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) LineItemList() *stripe.LineItemList {
	return i.List().(*stripe.LineItemList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
