//
//
// File generated from our OpenAPI spec
//
//

// Package schema provides the /v1/sigma/schemas APIs
package schema

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/sigma/schemas APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Lists the schemas available to the authorized merchant in Stripe's data products
func List(params *stripe.SigmaSchemaListParams) *Iter {
	return getC().List(params)
}

// Lists the schemas available to the authorized merchant in Stripe's data products
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.SigmaSchemaListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SigmaSchemaList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/sigma/schemas", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for sigma schemas.
type Iter struct {
	*stripe.Iter
}

// SigmaSchema returns the sigma schema which the iterator is currently pointing to.
func (i *Iter) SigmaSchema() *stripe.SigmaSchema {
	return i.Current().(*stripe.SigmaSchema)
}

// SigmaSchemaList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SigmaSchemaList() *stripe.SigmaSchemaList {
	return i.List().(*stripe.SigmaSchemaList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
