//
//
// File generated from our OpenAPI spec
//
//

// Package activeentitlement provides the /v1/entitlements/active_entitlements APIs
package activeentitlement

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/entitlements/active_entitlements APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve an active entitlement
func Get(id string, params *stripe.EntitlementsActiveEntitlementParams) (*stripe.EntitlementsActiveEntitlement, error) {
	return getC().Get(id, params)
}

// Retrieve an active entitlement
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.EntitlementsActiveEntitlementParams) (*stripe.EntitlementsActiveEntitlement, error) {
	path := stripe.FormatURLPath("/v1/entitlements/active_entitlements/%s", id)
	activeentitlement := &stripe.EntitlementsActiveEntitlement{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, activeentitlement)
	return activeentitlement, err
}

// Retrieve a list of active entitlements for a customer
func List(params *stripe.EntitlementsActiveEntitlementListParams) *Iter {
	return getC().List(params)
}

// Retrieve a list of active entitlements for a customer
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.EntitlementsActiveEntitlementListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.EntitlementsActiveEntitlementList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/entitlements/active_entitlements", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for entitlements active entitlements.
type Iter struct {
	*stripe.Iter
}

// EntitlementsActiveEntitlement returns the entitlements active entitlement which the iterator is currently pointing to.
func (i *Iter) EntitlementsActiveEntitlement() *stripe.EntitlementsActiveEntitlement {
	return i.Current().(*stripe.EntitlementsActiveEntitlement)
}

// EntitlementsActiveEntitlementList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) EntitlementsActiveEntitlementList() *stripe.EntitlementsActiveEntitlementList {
	return i.List().(*stripe.EntitlementsActiveEntitlementList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
