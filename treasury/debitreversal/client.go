//
//
// File generated from our OpenAPI spec
//
//

// Package debitreversal provides the /treasury/debit_reversals APIs
package debitreversal

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /treasury/debit_reversals APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new treasury debit reversal.
func New(params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	return getC().New(params)
}

// New creates a new treasury debit reversal.
func (c Client) New(params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	debitreversal := &stripe.TreasuryDebitReversal{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/treasury/debit_reversals",
		c.Key,
		params,
		debitreversal,
	)
	return debitreversal, err
}

// Get returns the details of a treasury debit reversal.
func Get(id string, params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	return getC().Get(id, params)
}

// Get returns the details of a treasury debit reversal.
func (c Client) Get(id string, params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	path := stripe.FormatURLPath("/v1/treasury/debit_reversals/%s", id)
	debitreversal := &stripe.TreasuryDebitReversal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, debitreversal)
	return debitreversal, err
}

// List returns a list of treasury debit reversals.
func List(params *stripe.TreasuryDebitReversalListParams) *Iter {
	return getC().List(params)
}

// List returns a list of treasury debit reversals.
func (c Client) List(listParams *stripe.TreasuryDebitReversalListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryDebitReversalList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/debit_reversals", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for treasury debit reversals.
type Iter struct {
	*stripe.Iter
}

// TreasuryDebitReversal returns the treasury debit reversal which the iterator is currently pointing to.
func (i *Iter) TreasuryDebitReversal() *stripe.TreasuryDebitReversal {
	return i.Current().(*stripe.TreasuryDebitReversal)
}

// TreasuryDebitReversalList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryDebitReversalList() *stripe.TreasuryDebitReversalList {
	return i.List().(*stripe.TreasuryDebitReversalList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
