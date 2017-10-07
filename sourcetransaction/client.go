// Package sourcetransaction provides the /source/transactions APIs.
package sourcetransaction

import (
	"errors"
	"fmt"

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

func (c Client) List(params *stripe.SourceTransactionListParams) *Iter {
	body := &form.Values{}
	var lp *stripe.ListParams
	var p *stripe.Params

	form.AppendTo(body, params)
	lp = &params.ListParams
	p = params.ToParams()

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.SourceTransactionList{}
		var err error

		if params != nil && len(params.Source) > 0 {
			err = c.B.Call("GET", fmt.Sprintf("/sources/%v/source_transactions", params.Source), c.Key, b, p, list)
		} else {
			err = errors.New("Invalid source transaction params: Source needs to be set")
		}

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
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
