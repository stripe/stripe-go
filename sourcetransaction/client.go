// Package sourcetransaction provides the /source/transactions APIs.
package sourcetransaction

import (
	"errors"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /sources/:source_id/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of source transactions.
// For more details see https://stripe.com/docs/api#retrieve_source.
func List(params *stripe.SourceTransactionListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.SourceTransactionListParams) *Iter {
	var outerErr error
	var path string

	if listParams == nil || listParams.Source == nil {
		outerErr = errors.New("Invalid source transaction params: Source needs to be set")
	} else {
		path = stripe.FormatURLPath("/sources/%s/source_transactions",
			stripe.StringValue(listParams.Source))
	}

	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.SourceTransactionList{}

		if outerErr != nil {
			return nil, list.ListMeta, outerErr
		}

		err := c.B.CallRaw("GET", path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of SourceTransactions.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// SourceTransaction returns the most recent SourceTransaction
// visited by a call to Next.
func (i *Iter) SourceTransaction() *stripe.SourceTransaction {
	return i.Current().(*stripe.SourceTransaction)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
