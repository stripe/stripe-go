//
//
// File generated from our OpenAPI spec
//
//

// Package order provides the /climate/orders APIs
package order

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /climate/orders APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new climate order.
func New(params *stripe.ClimateOrderParams) (*stripe.ClimateOrder, error) {
	return getC().New(params)
}

// New creates a new climate order.
func (c Client) New(params *stripe.ClimateOrderParams) (*stripe.ClimateOrder, error) {
	order := &stripe.ClimateOrder{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, "/v1/climate/orders", c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, order)
	return order, err
}

// Get returns the details of a climate order.
func Get(id string, params *stripe.ClimateOrderParams) (*stripe.ClimateOrder, error) {
	return getC().Get(id, params)
}

// Get returns the details of a climate order.
func (c Client) Get(id string, params *stripe.ClimateOrderParams) (*stripe.ClimateOrder, error) {
	path := stripe.FormatURLPath("/v1/climate/orders/%s", id)
	order := &stripe.ClimateOrder{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, order)
	return order, err
}

// Update updates a climate order's properties.
func Update(id string, params *stripe.ClimateOrderParams) (*stripe.ClimateOrder, error) {
	return getC().Update(id, params)
}

// Update updates a climate order's properties.
func (c Client) Update(id string, params *stripe.ClimateOrderParams) (*stripe.ClimateOrder, error) {
	path := stripe.FormatURLPath("/v1/climate/orders/%s", id)
	order := &stripe.ClimateOrder{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, order)
	return order, err
}

// Cancel is the method for the `POST /v1/climate/orders/{order}/cancel` API.
func Cancel(id string, params *stripe.ClimateOrderCancelParams) (*stripe.ClimateOrder, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/climate/orders/{order}/cancel` API.
func (c Client) Cancel(id string, params *stripe.ClimateOrderCancelParams) (*stripe.ClimateOrder, error) {
	path := stripe.FormatURLPath("/v1/climate/orders/%s/cancel", id)
	order := &stripe.ClimateOrder{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, order)
	return order, err
}

// List returns a list of climate orders.
func List(params *stripe.ClimateOrderListParams) *Iter {
	return getC().List(params)
}

// List returns a list of climate orders.
func (c Client) List(listParams *stripe.ClimateOrderListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ClimateOrderList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				"/v1/climate/orders",
				c.Key,
			)
			err := sr.SetRawForm(p, b)
			if err != nil {
				return nil, list, err
			}
			err = c.B.Call(sr, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for climate orders.
type Iter struct {
	*stripe.Iter
}

// ClimateOrder returns the climate order which the iterator is currently pointing to.
func (i *Iter) ClimateOrder() *stripe.ClimateOrder {
	return i.Current().(*stripe.ClimateOrder)
}

// ClimateOrderList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ClimateOrderList() *stripe.ClimateOrderList {
	return i.List().(*stripe.ClimateOrderList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
