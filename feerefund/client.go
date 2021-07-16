//
//
// File generated from our OpenAPI spec
//
//

// Package feerefund provides the /application_fees/{id}/refunds APIs
package feerefund

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /application_fees/{id}/refunds APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new fee refund.
func New(params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	return getC().New(params)
}

// New creates a new fee refund.
func (c Client) New(params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.ApplicationFee == nil {
		return nil, fmt.Errorf("params.ApplicationFee must be set")
	}
	path := stripe.FormatURLPath(
		"/v1/application_fees/%s/refunds",
		stripe.StringValue(params.ApplicationFee),
	)
	feerefund := &stripe.FeeRefund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, feerefund)
	return feerefund, err
}

// Get returns the details of a fee refund.
func Get(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	return getC().Get(id, params)
}

// Get returns the details of a fee refund.
func (c Client) Get(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.ApplicationFee == nil {
		return nil, fmt.Errorf("params.ApplicationFee must be set")
	}
	path := stripe.FormatURLPath(
		"/v1/application_fees/%s/refunds/%s",
		stripe.StringValue(params.ApplicationFee),
		id,
	)
	feerefund := &stripe.FeeRefund{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, feerefund)
	return feerefund, err
}

// Update updates a fee refund's properties.
func Update(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	return getC().Update(id, params)
}

// Update updates a fee refund's properties.
func (c Client) Update(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.ApplicationFee == nil {
		return nil, fmt.Errorf("params.ApplicationFee must be set")
	}
	path := stripe.FormatURLPath(
		"/v1/application_fees/%s/refunds/%s",
		stripe.StringValue(params.ApplicationFee),
		id,
	)
	feerefund := &stripe.FeeRefund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, feerefund)
	return feerefund, err
}

// List returns a list of fee refunds.
func List(params *stripe.FeeRefundListParams) *Iter {
	return getC().List(params)
}

// List returns a list of fee refunds.
func (c Client) List(listParams *stripe.FeeRefundListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/application_fees/%s/refunds",
		stripe.StringValue(listParams.ApplicationFee),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.FeeRefundList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for fee refunds.
type Iter struct {
	*stripe.Iter
}

// FeeRefund returns the fee refund which the iterator is currently pointing to.
func (i *Iter) FeeRefund() *stripe.FeeRefund {
	return i.Current().(*stripe.FeeRefund)
}

// FeeRefundList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FeeRefundList() *stripe.FeeRefundList {
	return i.List().(*stripe.FeeRefundList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
