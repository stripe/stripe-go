//
//
// File generated from our OpenAPI spec
//
//

// Package order provides the /v1/orders APIs
package order

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/form"
)

// Client is used to invoke /v1/orders APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new open order object.
func New(params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().New(params)
}

// Creates a new open order object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.OrderParams) (*stripe.Order, error) {
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, "/v1/orders", c.Key, params, order)
	return order, err
}

// Retrieves the details of an existing order. Supply the unique order ID from either an order creation request or the order list, and Stripe will return the corresponding order information.
func Get(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().Get(id, params)
}

// Retrieves the details of an existing order. Supply the unique order ID from either an order creation request or the order list, and Stripe will return the corresponding order information.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, order)
	return order, err
}

// Updates the specific order by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	return getC().Update(id, params)
}

// Updates the specific order by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.OrderParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Cancels the order as well as the payment intent if one is attached.
func Cancel(id string, params *stripe.OrderCancelParams) (*stripe.Order, error) {
	return getC().Cancel(id, params)
}

// Cancels the order as well as the payment intent if one is attached.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.OrderCancelParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s/cancel", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Reopens a submitted order.
func Reopen(id string, params *stripe.OrderReopenParams) (*stripe.Order, error) {
	return getC().Reopen(id, params)
}

// Reopens a submitted order.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Reopen(id string, params *stripe.OrderReopenParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s/reopen", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Submitting an Order transitions the status to processing and creates a PaymentIntent object so the order can be paid. If the Order has an amount_total of 0, no PaymentIntent object will be created. Once the order is submitted, its contents cannot be changed, unless the [reopen](https://docs.stripe.com/api#reopen_order) method is called.
func Submit(id string, params *stripe.OrderSubmitParams) (*stripe.Order, error) {
	return getC().Submit(id, params)
}

// Submitting an Order transitions the status to processing and creates a PaymentIntent object so the order can be paid. If the Order has an amount_total of 0, no PaymentIntent object will be created. Once the order is submitted, its contents cannot be changed, unless the [reopen](https://docs.stripe.com/api#reopen_order) method is called.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Submit(id string, params *stripe.OrderSubmitParams) (*stripe.Order, error) {
	path := stripe.FormatURLPath("/v1/orders/%s/submit", id)
	order := &stripe.Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Returns a list of your orders. The orders are returned sorted by creation date, with the most recently created orders appearing first.
func List(params *stripe.OrderListParams) *Iter {
	return getC().List(params)
}

// Returns a list of your orders. The orders are returned sorted by creation date, with the most recently created orders appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.OrderListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.OrderList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/orders", c.Key, []byte(b.Encode()), p, list)

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

// When retrieving an order, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func ListLineItems(params *stripe.OrderListLineItemsParams) *LineItemIter {
	return getC().ListLineItems(params)
}

// When retrieving an order, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLineItems(listParams *stripe.OrderListLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/orders/%s/line_items", stripe.StringValue(listParams.ID))
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.LineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

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
