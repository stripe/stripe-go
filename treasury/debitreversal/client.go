//
//
// File generated from our OpenAPI spec
//
//

// Package debitreversal provides the /v1/treasury/debit_reversals APIs
package debitreversal

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/treasury/debit_reversals APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Reverses a ReceivedDebit and creates a DebitReversal object.
func New(params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	return getC().New(params)
}

// Reverses a ReceivedDebit and creates a DebitReversal object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	debitreversal := &stripe.TreasuryDebitReversal{}
	err := c.B.Call(
		http.MethodPost, "/v1/treasury/debit_reversals", c.Key, params, debitreversal)
	return debitreversal, err
}

// Retrieves a DebitReversal object.
func Get(id string, params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	return getC().Get(id, params)
}

// Retrieves a DebitReversal object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TreasuryDebitReversalParams) (*stripe.TreasuryDebitReversal, error) {
	path := stripe.FormatURLPath("/v1/treasury/debit_reversals/%s", id)
	debitreversal := &stripe.TreasuryDebitReversal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, debitreversal)
	return debitreversal, err
}

// Returns a list of DebitReversals.
func List(params *stripe.TreasuryDebitReversalListParams) *Iter {
	return getC().List(params)
}

// Returns a list of DebitReversals.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TreasuryDebitReversalListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryDebitReversalList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/debit_reversals", c.Key, []byte(b.Encode()), p, list)

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
