// Package bitcoinreceiver provides the /bitcoin/receivers APIs.
//
// Note that this entire package is deprecated. Please use the new sources API
// instead.
package bitcoinreceiver

import (
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /bitcoin/receivers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new bitcoin receivers.
// For more details see https://stripe.com/docs/api/#create_bitcoin_receiver
func New(params *stripe.BitcoinReceiverParams) (*stripe.BitcoinReceiver, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.BitcoinReceiverParams) (*stripe.BitcoinReceiver, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	receiver := &stripe.BitcoinReceiver{}
	err := c.B.Call("POST", "/bitcoin/receivers", c.Key, body, &params.Params, receiver)

	return receiver, err
}

// Get returns the details of a bitcoin receiver.
// For more details see https://stripe.com/docs/api/#retrieve_bitcoin_receiver
func Get(id string, params *stripe.BitcoinReceiverParams) (*stripe.BitcoinReceiver, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.BitcoinReceiverParams) (*stripe.BitcoinReceiver, error) {
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
	}

	bitcoinReceiver := &stripe.BitcoinReceiver{}
	err := c.B.Call("GET", "/bitcoin/receivers/"+id, c.Key, nil, commonParams, bitcoinReceiver)

	return bitcoinReceiver, err
}

// Update updates a bitcoin receiver's properties.
// For more details see https://stripe.com/docs/api#update_bitcoin_receiver.
func Update(id string, params *stripe.BitcoinReceiverUpdateParams) (*stripe.BitcoinReceiver, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.BitcoinReceiverUpdateParams) (*stripe.BitcoinReceiver, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	receiver := &stripe.BitcoinReceiver{}
	err := c.B.Call("POST", fmt.Sprintf("/bitcoin/receivers/%v", id), c.Key, body, &params.Params, receiver)

	return receiver, err
}

// List returns a list of bitcoin receivers.
// For more details see https://stripe.com/docs/api/#list_bitcoin_receivers
func List(params *stripe.BitcoinReceiverListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.BitcoinReceiverListParams) *Iter {
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
		list := &stripe.BitcoinReceiverList{}
		err := c.B.Call("GET", "/bitcoin/receivers", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of BitcoinReceivers.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// BitcoinReceiver returns the most recent BitcoinReceiver
// visited by a call to Next.
func (i *Iter) BitcoinReceiver() *stripe.BitcoinReceiver {
	return i.Current().(*stripe.BitcoinReceiver)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
