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

// v1PaymentLinkService is used to invoke /v1/payment_links APIs.
type v1PaymentLinkService struct {
	B   Backend
	Key string
}

// Creates a payment link.
func (c v1PaymentLinkService) Create(ctx context.Context, params *PaymentLinkCreateParams) (*PaymentLink, error) {
	if params == nil {
		params = &PaymentLinkCreateParams{}
	}
	params.Context = ctx
	paymentlink := &PaymentLink{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_links", c.Key, params, paymentlink)
	return paymentlink, err
}

// Retrieve a payment link.
func (c v1PaymentLinkService) Retrieve(ctx context.Context, id string, params *PaymentLinkRetrieveParams) (*PaymentLink, error) {
	if params == nil {
		params = &PaymentLinkRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_links/%s", id)
	paymentlink := &PaymentLink{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentlink)
	return paymentlink, err
}

// Updates a payment link.
func (c v1PaymentLinkService) Update(ctx context.Context, id string, params *PaymentLinkUpdateParams) (*PaymentLink, error) {
	if params == nil {
		params = &PaymentLinkUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_links/%s", id)
	paymentlink := &PaymentLink{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentlink)
	return paymentlink, err
}

// Returns a list of your payment links.
func (c v1PaymentLinkService) List(ctx context.Context, listParams *PaymentLinkListParams) Seq2[*PaymentLink, error] {
	if listParams == nil {
		listParams = &PaymentLinkListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*PaymentLink, ListContainer, error) {
		list := &PaymentLinkList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_links", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// When retrieving a payment link, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func (c v1PaymentLinkService) ListLineItems(ctx context.Context, listParams *PaymentLinkListLineItemsParams) Seq2[*LineItem, error] {
	if listParams == nil {
		listParams = &PaymentLinkListLineItemsParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/payment_links/%s/line_items", StringValue(listParams.PaymentLink))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*LineItem, ListContainer, error) {
		list := &LineItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
