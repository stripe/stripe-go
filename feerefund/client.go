// package feerefund provides the /application_fees/refunds APIs
package feerefund

import (
	"fmt"
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
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
	body := &url.Values{}

	if params.Amount > 0 {
		body.Add("amount", strconv.FormatUint(params.Amount, 10))
	}

	params.AppendTo(body)

	refund := &stripe.FeeRefund{}
	err := c.B.Call("POST", fmt.Sprintf("application_fees/%v/refunds", params.Fee), c.Key, body, refund)

	return refund, err
}

// Get returns the details of a fee refund.
// For more details see https://stripe.com/docs/api#retrieve_fee_refund.
func Get(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	refund := &stripe.FeeRefund{}
	err := c.B.Call("GET", fmt.Sprintf("/application_fees/%v/refunds/%v", params.Fee, id), c.Key, body, refund)

	return refund, err
}

// Update updates a refund's properties.
// For more details see https://stripe.com/docs/api#update_refund.
func Update(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.FeeRefundParams) (*stripe.FeeRefund, error) {
	body := &url.Values{}

	params.AppendTo(body)

	refund := &stripe.FeeRefund{}
	err := c.B.Call("POST", fmt.Sprintf("/application_fees/%v/refunds/%v", params.Fee, id), c.Key, body, refund)

	return refund, err
}

// List returns a list of fee refunds.
// For more details see https://stripe.com/docs/api#list_fee_refunds.
func List(params *stripe.FeeRefundListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.FeeRefundListParams) *Iter {
	body := &url.Values{}
	var lp *stripe.ListParams

	params.AppendTo(body)
	lp = &params.ListParams

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.FeeRefundList{}
		err := c.B.Call("GET", fmt.Sprintf("/application_fees/%v/refunds", params.Fee), c.Key, &b, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is a iterator for list responses.
type Iter struct {
	Iter *stripe.Iter
}

// Next returns the next value in the list.
func (i *Iter) Next() (*stripe.FeeRefund, error) {
	f, err := i.Iter.Next()
	if err != nil {
		return nil, err
	}

	return f.(*stripe.FeeRefund), err
}

// Stop returns true if there are no more iterations to be performed.
func (i *Iter) Stop() bool {
	return i.Iter.Stop()
}

// Meta returns the list metadata.
func (i *Iter) Meta() *stripe.ListMeta {
	return i.Iter.Meta()
}

func getC() Client {
	return Client{stripe.GetBackend(), stripe.Key}
}
