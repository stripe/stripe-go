//
//
// File generated from our OpenAPI spec
//
//

// Package taxid provides the /v1/tax_ids APIs
package taxid

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/tax_ids APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new tax_id object for a customer.
func New(params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	return getC().New(params)
}

// Creates a new tax_id object for a customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	path := "/v1/tax_ids"
	if params.Customer != nil {
		path = stripe.FormatURLPath(
			"/v1/customers/%s/tax_ids",
			stripe.StringValue(params.Customer),
		)
	}
	taxid := &stripe.TaxID{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, taxid)
	return taxid, err
}

// Retrieves the tax_id object with the given identifier.
func Get(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	return getC().Get(id, params)
}

// Retrieves the tax_id object with the given identifier.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	path := stripe.FormatURLPath(
		"/v1/tax_ids/%s",
		id,
	)
	if params.Customer != nil {
		path = stripe.FormatURLPath(
			"/v1/customers/%s/tax_ids/%s",
			stripe.StringValue(params.Customer),
			id,
		)
	}
	taxid := &stripe.TaxID{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, taxid)
	return taxid, err
}

// Deletes an existing tax_id object.
func Del(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	return getC().Del(id, params)
}

// Deletes an existing tax_id object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	path := stripe.FormatURLPath(
		"/v1/tax_ids/%s",
		id,
	)
	if params.Customer != nil {
		path = stripe.FormatURLPath(
			"/v1/customers/%s/tax_ids/%s",
			stripe.StringValue(params.Customer),
			id,
		)
	}
	taxid := &stripe.TaxID{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, taxid)
	return taxid, err
}

// Returns a list of tax IDs for a customer.
func List(params *stripe.TaxIDListParams) *Iter {
	return getC().List(params)
}

// Returns a list of tax IDs for a customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TaxIDListParams) *Iter {
	path := "/v1/tax_ids"
	if listParams != nil && listParams.Customer != nil {
		path = stripe.FormatURLPath(
			"/v1/customers/%s/tax_ids",
			stripe.StringValue(listParams.Customer),
		)
	}
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxIDList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for tax ids.
type Iter struct {
	*stripe.Iter
}

// TaxID returns the tax id which the iterator is currently pointing to.
func (i *Iter) TaxID() *stripe.TaxID {
	return i.Current().(*stripe.TaxID)
}

// TaxIDList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxIDList() *stripe.TaxIDList {
	return i.List().(*stripe.TaxIDList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
