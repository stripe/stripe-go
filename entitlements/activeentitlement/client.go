//
//
// File generated from our OpenAPI spec
//
//

// Package activeentitlement provides the /entitlements/active_entitlements APIs
package activeentitlement

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /entitlements/active_entitlements APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of entitlements active entitlements.
func List(params *stripe.EntitlementsActiveEntitlementListParams) *Iter {
	return getC().List(params)
}

// List returns a list of entitlements active entitlements.
func (c Client) List(listParams *stripe.EntitlementsActiveEntitlementListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.EntitlementsActiveEntitlementList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/entitlements/active_entitlements", c.Key, b, p, list)

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
