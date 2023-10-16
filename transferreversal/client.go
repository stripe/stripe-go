//
//
// File generated from our OpenAPI spec
//
//

// Package transferreversal provides the /transfers/{id}/reversals APIs
package transferreversal

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /transfers/{id}/reversals APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new transfer reversal.
func New(params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	return getC().New(params)
}

// New creates a new transfer reversal.
func (c Client) New(params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	path := stripe.FormatURLPath(
		"/v1/transfers/%s/reversals",
		stripe.StringValue(params.ID),
	)
	transferreversal := &stripe.TransferReversal{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, transferreversal)
	return transferreversal, err
}

// Get returns the details of a transfer reversal.
func Get(id string, params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	return getC().Get(id, params)
}

// Get returns the details of a transfer reversal.
func (c Client) Get(id string, params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	if params == nil {
		return nil, fmt.Errorf(
			"params cannnot be nil, and params.Transfer must be set",
		)
	}
	path := stripe.FormatURLPath(
		"/v1/transfers/%s/reversals/%s",
		stripe.StringValue(params.ID),
		id,
	)
	transferreversal := &stripe.TransferReversal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transferreversal)
	return transferreversal, err
}

// Update updates a transfer reversal's properties.
func Update(id string, params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	return getC().Update(id, params)
}

// Update updates a transfer reversal's properties.
func (c Client) Update(id string, params *stripe.TransferReversalParams) (*stripe.TransferReversal, error) {
	path := stripe.FormatURLPath(
		"/v1/transfers/%s/reversals/%s",
		stripe.StringValue(params.ID),
		id,
	)
	transferreversal := &stripe.TransferReversal{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, transferreversal)
	return transferreversal, err
}

// List returns a list of transfer reversals.
func List(params *stripe.TransferReversalListParams) *Iter {
	return getC().List(params)
}

// List returns a list of transfer reversals.
func (c Client) List(listParams *stripe.TransferReversalListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/transfers/%s/reversals",
		stripe.StringValue(listParams.ID),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TransferReversalList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for transfer reversals.
type Iter struct {
	*stripe.Iter
}

// TransferReversal returns the transfer reversal which the iterator is currently pointing to.
func (i *Iter) TransferReversal() *stripe.TransferReversal {
	return i.Current().(*stripe.TransferReversal)
}

// TransferReversalList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TransferReversalList() *stripe.TransferReversalList {
	return i.List().(*stripe.TransferReversalList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
