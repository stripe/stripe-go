// Package feerefund provides the /application_fees/refunds APIs
package feerefund

import (
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /application_fees/refunds APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New refunds the application fee collected.
// For more details see https://stripe.com/docs/api#refund_application_fee.
func New(params *stripe.ApplicationFeeRefundParams) (*stripe.ApplicationFeeRefund, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.ApplicationFeeRefundParams) (*stripe.ApplicationFeeRefund, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.ApplicationFee == "" {
		return nil, fmt.Errorf("params.ApplicationFee must be set")
	}

	body := &form.Values{}
	form.AppendTo(body, params)

	refund := &stripe.ApplicationFeeRefund{}
	err := c.B.Call("POST", fmt.Sprintf("application_fees/%v/refunds", params.ApplicationFee), c.Key, body, &params.Params, refund)

	return refund, err
}

// Get returns the details of a fee refund.
// For more details see https://stripe.com/docs/api#retrieve_fee_refund.
func Get(id string, params *stripe.ApplicationFeeRefundParams) (*stripe.ApplicationFeeRefund, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ApplicationFeeRefundParams) (*stripe.ApplicationFeeRefund, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.ApplicationFee == "" {
		return nil, fmt.Errorf("params.ApplicationFee must be set")
	}

	body := &form.Values{}
	form.AppendTo(body, params)

	refund := &stripe.ApplicationFeeRefund{}
	err := c.B.Call("GET", fmt.Sprintf("/application_fees/%v/refunds/%v", params.ApplicationFee, id), c.Key, body, &params.Params, refund)

	return refund, err
}

// Update updates a refund's properties.
// For more details see https://stripe.com/docs/api#update_refund.
func Update(id string, params *stripe.ApplicationFeeRefundParams) (*stripe.ApplicationFeeRefund, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.ApplicationFeeRefundParams) (*stripe.ApplicationFeeRefund, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.ApplicationFee == "" {
		return nil, fmt.Errorf("params.ApplicationFee must be set")
	}

	body := &form.Values{}
	form.AppendTo(body, params)

	refund := &stripe.ApplicationFeeRefund{}
	err := c.B.Call("POST", fmt.Sprintf("/application_fees/%v/refunds/%v", params.ApplicationFee, id), c.Key, body, &params.Params, refund)

	return refund, err
}

// List returns a list of fee refunds.
// For more details see https://stripe.com/docs/api#list_fee_refunds.
func List(params *stripe.ApplicationFeeRefundListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.ApplicationFeeRefundListParams) *Iter {
	body := &form.Values{}
	var lp *stripe.ListParams
	var p *stripe.Params

	form.AppendTo(body, params)
	lp = &params.ListParams
	p = params.ToParams()

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.ApplicationFeeRefundList{}
		err := c.B.Call("GET", fmt.Sprintf("/application_fees/%v/refunds", params.ApplicationFee), c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of ApplicationFeeRefunds.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// ApplicationFeeRefund returns the most recent ApplicationFeeRefund
// visited by a call to Next.
func (i *Iter) ApplicationFeeRefund() *stripe.ApplicationFeeRefund {
	return i.Current().(*stripe.ApplicationFeeRefund)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
