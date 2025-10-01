//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v83/form"
)

// v1OrderService is used to invoke /v1/orders APIs.
type v1OrderService struct {
	B   Backend
	Key string
}

// Creates a new open order object.
func (c v1OrderService) Create(ctx context.Context, params *OrderCreateParams) (*Order, error) {
	if params == nil {
		params = &OrderCreateParams{}
	}
	params.Context = ctx
	order := &Order{}
	err := c.B.Call(http.MethodPost, "/v1/orders", c.Key, params, order)
	return order, err
}

// Retrieves the details of an existing order. Supply the unique order ID from either an order creation request or the order list, and Stripe will return the corresponding order information.
func (c v1OrderService) Retrieve(ctx context.Context, id string, params *OrderRetrieveParams) (*Order, error) {
	if params == nil {
		params = &OrderRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/orders/%s", id)
	order := &Order{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, order)
	return order, err
}

// Updates the specific order by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1OrderService) Update(ctx context.Context, id string, params *OrderUpdateParams) (*Order, error) {
	if params == nil {
		params = &OrderUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/orders/%s", id)
	order := &Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Cancels the order as well as the payment intent if one is attached.
func (c v1OrderService) Cancel(ctx context.Context, id string, params *OrderCancelParams) (*Order, error) {
	if params == nil {
		params = &OrderCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/orders/%s/cancel", id)
	order := &Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Reopens a submitted order.
func (c v1OrderService) Reopen(ctx context.Context, id string, params *OrderReopenParams) (*Order, error) {
	if params == nil {
		params = &OrderReopenParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/orders/%s/reopen", id)
	order := &Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Submitting an Order transitions the status to processing and creates a PaymentIntent object so the order can be paid. If the Order has an amount_total of 0, no PaymentIntent object will be created. Once the order is submitted, its contents cannot be changed, unless the [reopen](https://docs.stripe.com/api#reopen_order) method is called.
func (c v1OrderService) Submit(ctx context.Context, id string, params *OrderSubmitParams) (*Order, error) {
	if params == nil {
		params = &OrderSubmitParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/orders/%s/submit", id)
	order := &Order{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Returns a list of your orders. The orders are returned sorted by creation date, with the most recently created orders appearing first.
func (c v1OrderService) List(ctx context.Context, listParams *OrderListParams) Seq2[*Order, error] {
	if listParams == nil {
		listParams = &OrderListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*Order], error) {
		list := &v1Page[*Order]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/orders", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}

// When retrieving an order, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func (c v1OrderService) ListLineItems(ctx context.Context, listParams *OrderListLineItemsParams) Seq2[*LineItem, error] {
	if listParams == nil {
		listParams = &OrderListLineItemsParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath("/v1/orders/%s/line_items", StringValue(listParams.ID))
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*LineItem], error) {
		list := &v1Page[*LineItem]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
