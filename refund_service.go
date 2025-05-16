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

// v1RefundService is used to invoke /v1/refunds APIs.
type v1RefundService struct {
	B   Backend
	Key string
}

// When you create a new refund, you must specify a Charge or a PaymentIntent object on which to create it.
//
// Creating a new refund will refund a charge that has previously been created but not yet refunded.
// Funds will be refunded to the credit or debit card that was originally charged.
//
// You can optionally refund only part of a charge.
// You can do so multiple times, until the entire charge has been refunded.
//
// Once entirely refunded, a charge can't be refunded again.
// This method will raise an error when called on an already-refunded charge,
// or when trying to refund more money than is left on a charge.
func (c v1RefundService) Create(ctx context.Context, params *RefundCreateParams) (*Refund, error) {
	if params == nil {
		params = &RefundCreateParams{}
	}
	params.Context = ctx
	refund := &Refund{}
	err := c.B.Call(http.MethodPost, "/v1/refunds", c.Key, params, refund)
	return refund, err
}

// Retrieves the details of an existing refund.
func (c v1RefundService) Retrieve(ctx context.Context, id string, params *RefundRetrieveParams) (*Refund, error) {
	if params == nil {
		params = &RefundRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/refunds/%s", id)
	refund := &Refund{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, refund)
	return refund, err
}

// Updates the refund that you specify by setting the values of the passed parameters. Any parameters that you don't provide remain unchanged.
//
// This request only accepts metadata as an argument.
func (c v1RefundService) Update(ctx context.Context, id string, params *RefundUpdateParams) (*Refund, error) {
	if params == nil {
		params = &RefundUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/refunds/%s", id)
	refund := &Refund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, refund)
	return refund, err
}

// Cancels a refund with a status of requires_action.
//
// You can't cancel refunds in other states. Only refunds for payment methods that require customer action can enter the requires_action state.
func (c v1RefundService) Cancel(ctx context.Context, id string, params *RefundCancelParams) (*Refund, error) {
	if params == nil {
		params = &RefundCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/refunds/%s/cancel", id)
	refund := &Refund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, refund)
	return refund, err
}

// Returns a list of all refunds you created. We return the refunds in sorted order, with the most recent refunds appearing first. The 10 most recent refunds are always available by default on the Charge object.
func (c v1RefundService) List(ctx context.Context, listParams *RefundListParams) Seq2[*Refund, error] {
	if listParams == nil {
		listParams = &RefundListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Refund, ListContainer, error) {
		list := &RefundList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/refunds", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
