//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v86/form"
)

// v1CryptoOnrampSessionService is used to invoke /v1/crypto/onramp_sessions APIs.
type v1CryptoOnrampSessionService struct {
	B   Backend
	Key string
}

// Creates a CryptoOnrampSession object.
//
// After the CryptoOnrampSession is created, display the onramp session modal using the client_secret.
//
// Related guide: [Set up an onramp integration](https://docs.stripe.com/docs/crypto/integrate-the-onramp)
func (c v1CryptoOnrampSessionService) Create(ctx context.Context, params *CryptoOnrampSessionCreateParams) (*CryptoOnrampSession, error) {
	if params == nil {
		params = &CryptoOnrampSessionCreateParams{}
	}
	params.Context = ctx
	onrampsession := &CryptoOnrampSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/crypto/onramp_sessions", c.Key, params, onrampsession)
	return onrampsession, err
}

// Retrieves the details of a CryptoOnrampSession that was previously created.
func (c v1CryptoOnrampSessionService) Retrieve(ctx context.Context, id string, params *CryptoOnrampSessionRetrieveParams) (*CryptoOnrampSession, error) {
	if params == nil {
		params = &CryptoOnrampSessionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/crypto/onramp_sessions/%s", id)
	onrampsession := &CryptoOnrampSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, onrampsession)
	return onrampsession, err
}

// Completes a headless CryptoOnrampSession.
//
// This method will attempt to confirm the payment and execute the quote to deliver the crypto to the customer.
func (c v1CryptoOnrampSessionService) Checkout(ctx context.Context, id string, params *CryptoOnrampSessionCheckoutParams) (*CryptoOnrampSession, error) {
	if params == nil {
		params = &CryptoOnrampSessionCheckoutParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/crypto/onramp_sessions/%s/checkout", id)
	onrampsession := &CryptoOnrampSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, onrampsession)
	return onrampsession, err
}

// Refreshes an executable quote for a CryptoOnrampSession.
func (c v1CryptoOnrampSessionService) Quote(ctx context.Context, id string, params *CryptoOnrampSessionQuoteParams) (*CryptoOnrampSession, error) {
	if params == nil {
		params = &CryptoOnrampSessionQuoteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/crypto/onramp_sessions/%s/quote", id)
	onrampsession := &CryptoOnrampSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, onrampsession)
	return onrampsession, err
}

// Returns a list of onramp sessions that match the filter criteria. The onramp sessions are returned in sorted order, with the most recent onramp sessions appearing first.
func (c v1CryptoOnrampSessionService) List(ctx context.Context, listParams *CryptoOnrampSessionListParams) *V1List[*CryptoOnrampSession] {
	if listParams == nil {
		listParams = &CryptoOnrampSessionListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*CryptoOnrampSession], error) {
		list := &v1Page[*CryptoOnrampSession]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/crypto/onramp_sessions", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
