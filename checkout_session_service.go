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

// v1CheckoutSessionService is used to invoke /v1/checkout/sessions APIs.
type v1CheckoutSessionService struct {
	B   Backend
	Key string
}

// Creates a Checkout Session object.
func (c v1CheckoutSessionService) Create(ctx context.Context, params *CheckoutSessionCreateParams) (*CheckoutSession, error) {
	if params == nil {
		params = &CheckoutSessionCreateParams{}
	}
	params.Context = ctx
	session := &CheckoutSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/checkout/sessions", c.Key, params, session)
	return session, err
}

// Retrieves a Checkout Session object.
func (c v1CheckoutSessionService) Retrieve(ctx context.Context, id string, params *CheckoutSessionRetrieveParams) (*CheckoutSession, error) {
	if params == nil {
		params = &CheckoutSessionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/checkout/sessions/%s", id)
	session := &CheckoutSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, session)
	return session, err
}

// Updates a Checkout Session object.
//
// Related guide: [Dynamically update Checkout](https://docs.stripe.com/payments/checkout/dynamic-updates)
func (c v1CheckoutSessionService) Update(ctx context.Context, id string, params *CheckoutSessionUpdateParams) (*CheckoutSession, error) {
	if params == nil {
		params = &CheckoutSessionUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/checkout/sessions/%s", id)
	session := &CheckoutSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, session)
	return session, err
}

// A Checkout Session can be expired when it is in one of these statuses: open
//
// After it expires, a customer can't complete a Checkout Session and customers loading the Checkout Session see a message saying the Checkout Session is expired.
func (c v1CheckoutSessionService) Expire(ctx context.Context, id string, params *CheckoutSessionExpireParams) (*CheckoutSession, error) {
	if params == nil {
		params = &CheckoutSessionExpireParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/checkout/sessions/%s/expire", id)
	session := &CheckoutSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, session)
	return session, err
}

// Returns a list of Checkout Sessions.
func (c v1CheckoutSessionService) List(ctx context.Context, listParams *CheckoutSessionListParams) Seq2[*CheckoutSession, error] {
	if listParams == nil {
		listParams = &CheckoutSessionListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*CheckoutSession, ListContainer, error) {
		list := &CheckoutSessionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/checkout/sessions", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// When retrieving a Checkout Session, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func (c v1CheckoutSessionService) ListLineItems(ctx context.Context, listParams *CheckoutSessionListLineItemsParams) Seq2[*LineItem, error] {
	if listParams == nil {
		listParams = &CheckoutSessionListLineItemsParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/checkout/sessions/%s/line_items", StringValue(listParams.Session))
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
