// Package payout provides the /payouts APIs
package payout

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /payouts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new payout.
// For more details see https://stripe.com/docs/api#create_payout.
func New(params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.PayoutParams) (*stripe.Payout, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	payout := &stripe.Payout{}
	err := c.B.Call("POST", "/payouts", c.Key, body, &params.Params, payout)

	return payout, err
}

// Get returns the details of a payout.
// For more details see https://stripe.com/docs/api#retrieve_payout.
func Get(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	payout := &stripe.Payout{}
	err := c.B.Call("GET", stripe.FormatURLPath("/payouts/%s", id), c.Key, body, commonParams, payout)

	return payout, err
}

// Update updates a payout's properties.
// For more details see https://stripe.com/docs/api#update_payout.
func Update(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	payout := &stripe.Payout{}
	err := c.B.Call("POST", stripe.FormatURLPath("/payouts/%s", id), c.Key, body, commonParams, payout)

	return payout, err
}

// Cancel cancels a pending payout.
// For more details see https://stripe.com/docs/api#cancel_payout.
func Cancel(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().Cancel(id, params)
}

func (c Client) Cancel(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	payout := &stripe.Payout{}
	err := c.B.Call("POST", stripe.FormatURLPath("/payouts/%s/cancel", id), c.Key, body, commonParams, payout)

	return payout, err
}

// List returns a list of payouts.
// For more details see https://stripe.com/docs/api#list_payouts.
func List(params *stripe.PayoutListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.PayoutListParams) *Iter {
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
		list := &stripe.PayoutList{}
		err := c.B.Call("GET", "/payouts", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Payouts.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Payout returns the most recent Payout
// visited by a call to Next.
func (i *Iter) Payout() *stripe.Payout {
	return i.Current().(*stripe.Payout)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
