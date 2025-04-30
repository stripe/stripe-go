//
//
// File generated from our OpenAPI spec
//
//

// Package token provides the /v1/issuing/tokens APIs
package token

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/issuing/tokens APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an Issuing Token object.
func Get(id string, params *stripe.IssuingTokenParams) (*stripe.IssuingToken, error) {
	return getC().Get(id, params)
}

// Retrieves an Issuing Token object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.IssuingTokenParams) (*stripe.IssuingToken, error) {
	path := stripe.FormatURLPath("/v1/issuing/tokens/%s", id)
	token := &stripe.IssuingToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, token)
	return token, err
}

// Attempts to update the specified Issuing Token object to the status specified.
func Update(id string, params *stripe.IssuingTokenParams) (*stripe.IssuingToken, error) {
	return getC().Update(id, params)
}

// Attempts to update the specified Issuing Token object to the status specified.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingTokenParams) (*stripe.IssuingToken, error) {
	path := stripe.FormatURLPath("/v1/issuing/tokens/%s", id)
	token := &stripe.IssuingToken{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, token)
	return token, err
}

// Lists all Issuing Token objects for a given card.
func List(params *stripe.IssuingTokenListParams) *Iter {
	return getC().List(params)
}

// Lists all Issuing Token objects for a given card.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IssuingTokenListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingTokenList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/tokens", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing tokens.
type Iter struct {
	*stripe.Iter
}

// IssuingToken returns the issuing token which the iterator is currently pointing to.
func (i *Iter) IssuingToken() *stripe.IssuingToken {
	return i.Current().(*stripe.IssuingToken)
}

// IssuingTokenList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingTokenList() *stripe.IssuingTokenList {
	return i.List().(*stripe.IssuingTokenList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
