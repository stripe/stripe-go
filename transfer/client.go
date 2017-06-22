// Package transfer provides the /transfers APIs
package transfer

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	SourceAlipay  stripe.TransferSourceType = "alipay_account"
	SourceBank    stripe.TransferSourceType = "bank_account"
	SourceBitcoin stripe.TransferSourceType = "bitcoin_receiver"
	SourceCard    stripe.TransferSourceType = "card"
)

// Client is used to invoke /transfers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new transfer.
// For more details see https://stripe.com/docs/api#create_transfer.
func New(params *stripe.TransferParams) (*stripe.Transfer, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.TransferParams) (*stripe.Transfer, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	transfer := &stripe.Transfer{}
	err := c.B.Call("POST", "/transfers", c.Key, body, &params.Params, transfer)

	return transfer, err
}

// Get returns the details of a transfer.
// For more details see https://stripe.com/docs/api#retrieve_transfer.
func Get(id string, params *stripe.TransferParams) (*stripe.Transfer, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.TransferParams) (*stripe.Transfer, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	transfer := &stripe.Transfer{}
	err := c.B.Call("GET", "/transfers/"+id, c.Key, body, commonParams, transfer)

	return transfer, err
}

// Update updates a transfer's properties.
// For more details see https://stripe.com/docs/api#update_transfer.
func Update(id string, params *stripe.TransferParams) (*stripe.Transfer, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.TransferParams) (*stripe.Transfer, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params

		body = &form.Values{}

		form.AppendTo(body, params)
	}

	transfer := &stripe.Transfer{}
	err := c.B.Call("POST", "/transfers/"+id, c.Key, body, commonParams, transfer)

	return transfer, err
}

// List returns a list of transfers.
// For more details see https://stripe.com/docs/api#list_transfers.
func List(params *stripe.TransferListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.TransferListParams) *Iter {
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
		list := &stripe.TransferList{}
		err := c.B.Call("GET", "/transfers", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Transfers.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Transfer returns the most recent Transfer
// visited by a call to Next.
func (i *Iter) Transfer() *stripe.Transfer {
	return i.Current().(*stripe.Transfer)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
