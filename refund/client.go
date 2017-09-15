// Package refund provides the /refunds APIs
package refund

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	RefundFraudulent          stripe.RefundReason = "fraudulent"
	RefundDuplicate           stripe.RefundReason = "duplicate"
	RefundRequestedByCustomer stripe.RefundReason = "requested_by_customer"
)

// Client is used to invoke /refunds APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New refunds a charge previously created.
// For more details see https://stripe.com/docs/api#refund_charge.
func New(params *stripe.RefundParams) (*stripe.Refund, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.RefundParams) (*stripe.Refund, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	refund := &stripe.Refund{}
	err := c.B.Call("POST", "/refunds", c.Key, body, &params.Params, refund)

	return refund, err
}

// Get returns the details of a refund.
// For more details see https://stripe.com/docs/api#retrieve_refund.
func Get(id string, params *stripe.RefundParams) (*stripe.Refund, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.RefundParams) (*stripe.Refund, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	refund := &stripe.Refund{}
	err := c.B.Call("GET", "/refunds/"+id, c.Key, body, commonParams, refund)

	return refund, err
}

// Update updates a refund's properties.
// For more details see https://stripe.com/docs/api#update_refund.
func Update(id string, params *stripe.RefundParams) (*stripe.Refund, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.RefundParams) (*stripe.Refund, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	refund := &stripe.Refund{}
	err := c.B.Call("POST", "/refunds/"+id, c.Key, body, &params.Params, refund)

	return refund, err
}

// List returns a list of refunds.
// For more details see https://stripe.com/docs/api#list_refunds.
func List(params *stripe.RefundListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.RefundListParams) *Iter {
	body := &form.Values{}
	var lp *stripe.ListParams
	var p *stripe.Params

	form.AppendTo(body, params)
	lp = &params.ListParams
	p = params.ToParams()

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.RefundList{}
		err := c.B.Call("GET", "/refunds", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Refunds.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Refund returns the most recent Refund
// visited by a call to Next.
func (i *Iter) Refund() *stripe.Refund {
	return i.Current().(*stripe.Refund)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
