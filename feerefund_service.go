//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1FeeRefundService is used to invoke /v1/application_fees/{id}/refunds APIs.
type v1FeeRefundService struct {
	B   Backend
	Key string
}

// Refunds an application fee that has previously been collected but not yet refunded.
// Funds will be refunded to the Stripe account from which the fee was originally collected.
//
// You can optionally refund only part of an application fee.
// You can do so multiple times, until the entire fee has been refunded.
//
// Once entirely refunded, an application fee can't be refunded again.
// This method will raise an error when called on an already-refunded application fee,
// or when trying to refund more money than is left on an application fee.
func (c v1FeeRefundService) Create(ctx context.Context, params *FeeRefundCreateParams) (*FeeRefund, error) {
	if params == nil {
		params = &FeeRefundCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/application_fees/%s/refunds", StringValue(params.ID))
	feerefund := &FeeRefund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, feerefund)
	return feerefund, err
}

// By default, you can see the 10 most recent refunds stored directly on the application fee object, but you can also retrieve details about a specific refund stored on the application fee.
func (c v1FeeRefundService) Retrieve(ctx context.Context, id string, params *FeeRefundRetrieveParams) (*FeeRefund, error) {
	if params == nil {
		params = &FeeRefundRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/application_fees/%s/refunds/%s", StringValue(params.Fee), id)
	feerefund := &FeeRefund{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, feerefund)
	return feerefund, err
}

// Updates the specified application fee refund by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request only accepts metadata as an argument.
func (c v1FeeRefundService) Update(ctx context.Context, id string, params *FeeRefundUpdateParams) (*FeeRefund, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.Fee == nil {
		return nil, fmt.Errorf("params.Fee must be set")
	}
	if params == nil {
		params = &FeeRefundUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/application_fees/%s/refunds/%s", StringValue(params.Fee), id)
	feerefund := &FeeRefund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, feerefund)
	return feerefund, err
}

// You can see a list of the refunds belonging to a specific application fee. Note that the 10 most recent refunds are always available by default on the application fee object. If you need more than those 10, you can use this API method and the limit and starting_after parameters to page through additional refunds.
func (c v1FeeRefundService) List(ctx context.Context, listParams *FeeRefundListParams) Seq2[*FeeRefund, error] {
	if listParams == nil {
		listParams = &FeeRefundListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/application_fees/%s/refunds", StringValue(listParams.ID))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*FeeRefund, ListContainer, error) {
		list := &FeeRefundList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
