// Package bitcointransaction provides the /bitcoin/transactions APIs.
package bitcointransaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to invoke /bitcoin/receivers/:receiver_id/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of bitcoin transactions.
func List(params *stripe.BitcoinTransactionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of bitcoin transactions.
func (c Client) List(listParams *stripe.BitcoinTransactionListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		path := stripe.FormatURLPath("/v1/bitcoin/receivers/%s/transactions",
			stripe.StringValue(listParams.Receiver))
		list := &stripe.BitcoinTransactionList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for bitcoin transactions.
type Iter struct {
	*stripe.Iter
}

// BitcoinTransaction returns the bitcoin transaction which the iterator is currently pointing to.
func (i *Iter) BitcoinTransaction() *stripe.BitcoinTransaction {
	return i.Current().(*stripe.BitcoinTransaction)
}

// BitcoinTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BitcoinTransactionList() *stripe.BitcoinTransactionList {
	return i.List().(*stripe.BitcoinTransactionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
