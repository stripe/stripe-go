//
//
// File generated from our OpenAPI spec
//
//

// Package acceptedfinancing provides the /capital/financing/accepted APIs
package acceptedfinancing

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/form"
)

// Client is used to invoke /capital/financing/accepted APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an accepted financing.
func Get(params *stripe.AcceptedFinancingParams) *Iter {
	return getC().Get(params)
}

// Get returns the details of an accepted financing.
func (c Client) Get(listParams *stripe.AcceptedFinancingParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.AcceptedFinancingList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/capital/financing/accepted", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for accepted financings.
type Iter struct {
	*stripe.Iter
}

// AcceptedFinancing returns the accepted financing which the iterator is currently pointing to.
func (i *Iter) AcceptedFinancing() *stripe.AcceptedFinancing {
	return i.Current().(*stripe.AcceptedFinancing)
}

// AcceptedFinancingList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) AcceptedFinancingList() *stripe.AcceptedFinancingList {
	return i.List().(*stripe.AcceptedFinancingList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
