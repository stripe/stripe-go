//
//
// File generated from our OpenAPI spec
//
//

// Package customerentitlement provides the /customers/{customer}/entitlements APIs
package customerentitlement

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /customers/{customer}/entitlements APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of customer entitlements.
func List(params *stripe.CustomerEntitlementListParams) *Iter {
	return getC().List(params)
}

// List returns a list of customer entitlements.
func (c Client) List(listParams *stripe.CustomerEntitlementListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/entitlements",
		stripe.StringValue(listParams.Customer),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CustomerEntitlementList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for customer entitlements.
type Iter struct {
	*stripe.Iter
}

// CustomerEntitlement returns the customer entitlement which the iterator is currently pointing to.
func (i *Iter) CustomerEntitlement() *stripe.CustomerEntitlement {
	return i.Current().(*stripe.CustomerEntitlement)
}

// CustomerEntitlementList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CustomerEntitlementList() *stripe.CustomerEntitlementList {
	return i.List().(*stripe.CustomerEntitlementList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
