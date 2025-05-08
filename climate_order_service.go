//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1ClimateOrderService is used to invoke /v1/climate/orders APIs.
type v1ClimateOrderService struct {
	B   Backend
	Key string
}

// Creates a Climate order object for a given Climate product. The order will be processed immediately
// after creation and payment will be deducted your Stripe balance.
func (c v1ClimateOrderService) Create(ctx context.Context, params *ClimateOrderCreateParams) (*ClimateOrder, error) {
	if params == nil {
		params = &ClimateOrderCreateParams{}
	}
	params.Context = ctx
	order := &ClimateOrder{}
	err := c.B.Call(http.MethodPost, "/v1/climate/orders", c.Key, params, order)
	return order, err
}

// Retrieves the details of a Climate order object with the given ID.
func (c v1ClimateOrderService) Retrieve(ctx context.Context, id string, params *ClimateOrderRetrieveParams) (*ClimateOrder, error) {
	if params == nil {
		params = &ClimateOrderRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/climate/orders/%s", id)
	order := &ClimateOrder{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, order)
	return order, err
}

// Updates the specified order by setting the values of the parameters passed.
func (c v1ClimateOrderService) Update(ctx context.Context, id string, params *ClimateOrderUpdateParams) (*ClimateOrder, error) {
	if params == nil {
		params = &ClimateOrderUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/climate/orders/%s", id)
	order := &ClimateOrder{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Cancels a Climate order. You can cancel an order within 24 hours of creation. Stripe refunds the
// reservation amount_subtotal, but not the amount_fees for user-triggered cancellations. Frontier
// might cancel reservations if suppliers fail to deliver. If Frontier cancels the reservation, Stripe
// provides 90 days advance notice and refunds the amount_total.
func (c v1ClimateOrderService) Cancel(ctx context.Context, id string, params *ClimateOrderCancelParams) (*ClimateOrder, error) {
	if params == nil {
		params = &ClimateOrderCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/climate/orders/%s/cancel", id)
	order := &ClimateOrder{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, order)
	return order, err
}

// Lists all Climate order objects. The orders are returned sorted by creation date, with the
// most recently created orders appearing first.
func (c v1ClimateOrderService) List(ctx context.Context, listParams *ClimateOrderListParams) Seq2[*ClimateOrder, error] {
	if listParams == nil {
		listParams = &ClimateOrderListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*ClimateOrder, ListContainer, error) {
		list := &ClimateOrderList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/climate/orders", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
