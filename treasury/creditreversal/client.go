//
//
// File generated from our OpenAPI spec
//
//

// Package creditreversal provides the /treasury/credit_reversals APIs
package creditreversal

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /treasury/credit_reversals APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new treasury credit reversal.
func New(params *stripe.TreasuryCreditReversalParams) (*stripe.TreasuryCreditReversal, error) {
	return getC().New(params)
}

// New creates a new treasury credit reversal.
func (c Client) New(params *stripe.TreasuryCreditReversalParams) (*stripe.TreasuryCreditReversal, error) {
	creditreversal := &stripe.TreasuryCreditReversal{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/treasury/credit_reversals",
		c.Key,
		params,
		creditreversal,
	)
	return creditreversal, err
}

// Get returns the details of a treasury credit reversal.
func Get(id string, params *stripe.TreasuryCreditReversalParams) (*stripe.TreasuryCreditReversal, error) {
	return getC().Get(id, params)
}

// Get returns the details of a treasury credit reversal.
func (c Client) Get(id string, params *stripe.TreasuryCreditReversalParams) (*stripe.TreasuryCreditReversal, error) {
	path := stripe.FormatURLPath("/v1/treasury/credit_reversals/%s", id)
	creditreversal := &stripe.TreasuryCreditReversal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, creditreversal)
	return creditreversal, err
}

// List returns a list of treasury credit reversals.
func List(params *stripe.TreasuryCreditReversalListParams) *Iter {
	return getC().List(params)
}

// List returns a list of treasury credit reversals.
func (c Client) List(listParams *stripe.TreasuryCreditReversalListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryCreditReversalList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/credit_reversals", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for treasury credit reversals.
type Iter struct {
	*stripe.Iter
}

// TreasuryCreditReversal returns the treasury credit reversal which the iterator is currently pointing to.
func (i *Iter) TreasuryCreditReversal() *stripe.TreasuryCreditReversal {
	return i.Current().(*stripe.TreasuryCreditReversal)
}

// TreasuryCreditReversalList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryCreditReversalList() *stripe.TreasuryCreditReversalList {
	return i.List().(*stripe.TreasuryCreditReversalList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
