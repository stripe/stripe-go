// Package bitcointransaction provides the /bitcoin/transactions APIs.
package bitcointransaction

import (
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /bitcoin/receivers/:receiver_id/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of bitcoin transactions.
// For more details see https://stripe.com/docs/api#retrieve_bitcoin_receiver.
func List(params *stripe.BitcoinTransactionListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.BitcoinTransactionListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.BitcoinTransactionList{}
		err := c.B.Call("GET", fmt.Sprintf("/bitcoin/receivers/%v/transactions", params.Receiver), c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of BitcoinTransactions.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// BitcoinTransaction returns the most recent BitcoinTransaction
// visited by a call to Next.
func (i *Iter) BitcoinTransaction() *stripe.BitcoinTransaction {
	return i.Current().(*stripe.BitcoinTransaction)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
