// Package bitcointransaction provides the /bitcoin/transactions APIs.
package bitcointransaction

import (
	"fmt"
	stripe "github.com/stripe/stripe-go"
	"net/url"
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
	type receiverList struct {
		stripe.ListMeta
		Values []*stripe.BitcoinTransaction `json:"data"`
	}

	var body *url.Values
	var lp *stripe.ListParams

	if params != nil {
		body = &url.Values{}

		if len(params.Customer) > 0 {
			body.Add("customer", params.Customer)
		}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &receiverList{}
		err := c.B.Call("GET", fmt.Sprintf("/bitcoin/receivers/%v/transactions", params.Receiver), c.Key, &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
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
