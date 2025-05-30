//
//
// File generated from our OpenAPI spec
//
//

// Package applepaydomain provides the /v1/apple_pay/domains APIs
package applepaydomain

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/apple_pay/domains APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an apple pay domain.
func New(params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	return getC().New(params)
}

// Create an apple pay domain.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	applepaydomain := &stripe.ApplePayDomain{}
	err := c.B.Call(
		http.MethodPost, "/v1/apple_pay/domains", c.Key, params, applepaydomain)
	return applepaydomain, err
}

// Retrieve an apple pay domain.
func Get(id string, params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	return getC().Get(id, params)
}

// Retrieve an apple pay domain.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	path := stripe.FormatURLPath("/v1/apple_pay/domains/%s", id)
	applepaydomain := &stripe.ApplePayDomain{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, applepaydomain)
	return applepaydomain, err
}

// Delete an apple pay domain.
func Del(id string, params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	return getC().Del(id, params)
}

// Delete an apple pay domain.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	path := stripe.FormatURLPath("/v1/apple_pay/domains/%s", id)
	applepaydomain := &stripe.ApplePayDomain{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, applepaydomain)
	return applepaydomain, err
}

// List apple pay domains.
func List(params *stripe.ApplePayDomainListParams) *Iter {
	return getC().List(params)
}

// List apple pay domains.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ApplePayDomainListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ApplePayDomainList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/apple_pay/domains", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for apple pay domains.
type Iter struct {
	*stripe.Iter
}

// ApplePayDomain returns the apple pay domain which the iterator is currently pointing to.
func (i *Iter) ApplePayDomain() *stripe.ApplePayDomain {
	return i.Current().(*stripe.ApplePayDomain)
}

// ApplePayDomainList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ApplePayDomainList() *stripe.ApplePayDomainList {
	return i.List().(*stripe.ApplePayDomainList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
