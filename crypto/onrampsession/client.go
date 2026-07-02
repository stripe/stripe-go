//
//
// File generated from our OpenAPI spec
//
//

// Package onrampsession provides the /v1/crypto/onramp_sessions APIs
package onrampsession

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/crypto/onramp_sessions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a CryptoOnrampSession object.
//
// After the CryptoOnrampSession is created, display the onramp session modal using the client_secret.
//
// Related guide: [Set up an onramp integration](https://docs.stripe.com/docs/crypto/integrate-the-onramp)
func New(params *stripe.CryptoOnrampSessionParams) (*stripe.CryptoOnrampSession, error) {
	return getC().New(params)
}

// Creates a CryptoOnrampSession object.
//
// After the CryptoOnrampSession is created, display the onramp session modal using the client_secret.
//
// Related guide: [Set up an onramp integration](https://docs.stripe.com/docs/crypto/integrate-the-onramp)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.CryptoOnrampSessionParams) (*stripe.CryptoOnrampSession, error) {
	onrampsession := &stripe.CryptoOnrampSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/crypto/onramp_sessions", c.Key, params, onrampsession)
	return onrampsession, err
}

// Retrieves the details of a CryptoOnrampSession that was previously created.
func Get(id string, params *stripe.CryptoOnrampSessionParams) (*stripe.CryptoOnrampSession, error) {
	return getC().Get(id, params)
}

// Retrieves the details of a CryptoOnrampSession that was previously created.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.CryptoOnrampSessionParams) (*stripe.CryptoOnrampSession, error) {
	path := stripe.FormatURLPath("/v1/crypto/onramp_sessions/%s", id)
	onrampsession := &stripe.CryptoOnrampSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, onrampsession)
	return onrampsession, err
}

// Completes a headless CryptoOnrampSession.
//
// This method will attempt to confirm the payment and execute the quote to deliver the crypto to the customer.
func Checkout(id string, params *stripe.CryptoOnrampSessionCheckoutParams) (*stripe.CryptoOnrampSession, error) {
	return getC().Checkout(id, params)
}

// Completes a headless CryptoOnrampSession.
//
// This method will attempt to confirm the payment and execute the quote to deliver the crypto to the customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Checkout(id string, params *stripe.CryptoOnrampSessionCheckoutParams) (*stripe.CryptoOnrampSession, error) {
	path := stripe.FormatURLPath("/v1/crypto/onramp_sessions/%s/checkout", id)
	onrampsession := &stripe.CryptoOnrampSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, onrampsession)
	return onrampsession, err
}

// Refreshes an executable quote for a CryptoOnrampSession.
func Quote(id string, params *stripe.CryptoOnrampSessionQuoteParams) (*stripe.CryptoOnrampSession, error) {
	return getC().Quote(id, params)
}

// Refreshes an executable quote for a CryptoOnrampSession.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Quote(id string, params *stripe.CryptoOnrampSessionQuoteParams) (*stripe.CryptoOnrampSession, error) {
	path := stripe.FormatURLPath("/v1/crypto/onramp_sessions/%s/quote", id)
	onrampsession := &stripe.CryptoOnrampSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, onrampsession)
	return onrampsession, err
}

// Returns a list of onramp sessions that match the filter criteria. The onramp sessions are returned in sorted order, with the most recent onramp sessions appearing first.
func List(params *stripe.CryptoOnrampSessionListParams) *Iter {
	return getC().List(params)
}

// Returns a list of onramp sessions that match the filter criteria. The onramp sessions are returned in sorted order, with the most recent onramp sessions appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CryptoOnrampSessionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CryptoOnrampSessionList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/crypto/onramp_sessions", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for crypto onramp sessions.
type Iter struct {
	*stripe.Iter
}

// CryptoOnrampSession returns the crypto onramp session which the iterator is currently pointing to.
func (i *Iter) CryptoOnrampSession() *stripe.CryptoOnrampSession {
	return i.Current().(*stripe.CryptoOnrampSession)
}

// CryptoOnrampSessionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CryptoOnrampSessionList() *stripe.CryptoOnrampSessionList {
	return i.List().(*stripe.CryptoOnrampSessionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
