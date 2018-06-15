// Package feerefund provides the /application_fees/refunds APIs
package feerefund

import (
	"fmt"
	"net/http"

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
func New(params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.ApplicationFee == nil {
		return nil, fmt.Errorf("params.ApplicationFee must be set")
	}

	path := stripe.FormatURLPath("/application_fees/%s/refunds",
		stripe.StringValue(params.ApplicationFee))
	refund := &stripe.FeeRefund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, refund)
	return refund, err
}

// Get returns the details of a fee refund.
// For more details see https://stripe.com/docs/api#retrieve_fee_refund.
func Get(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.ApplicationFee == nil {
		return nil, fmt.Errorf("params.ApplicationFee must be set")
	}

	path := stripe.FormatURLPath("/application_fees/%s/refunds/%s",
		stripe.StringValue(params.ApplicationFee), id)
	refund := &stripe.FeeRefund{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, refund)
	return refund, err
}

// Update updates a refund's properties.
// For more details see https://stripe.com/docs/api#update_refund.
func Update(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.ApplicationFee == nil {
		return nil, fmt.Errorf("params.ApplicationFee must be set")
	}

	path := stripe.FormatURLPath("/application_fees/%s/refunds/%s",
		stripe.StringValue(params.ApplicationFee), id)
	refund := &stripe.FeeRefund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, refund)

	return refund, err
}

// List returns a list of fee refunds.
// For more details see https://stripe.com/docs/api#list_fee_refunds.
func List(params *stripe.FeeRefundListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.FeeRefundListParams) *Iter {
	path := stripe.FormatURLPath("/application_fees/%s/refunds",
		stripe.StringValue(listParams.ApplicationFee))

	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.FeeRefundList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of FeeRefunds.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// FeeRefund returns the most recent FeeRefund
// visited by a call to Next.
func (i *Iter) FeeRefund() *stripe.FeeRefund {
	return i.Current().(*stripe.FeeRefund)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
