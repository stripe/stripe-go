// package feerefund provides the /application_fees/refunds APIs
package feerefund

import (
	"fmt"
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /application_fees/refunds APIs.
type Client struct {
	B   Backend
	Key string
}

// New refunds the application fee collected.
// For more details see https://stripe.com/docs/api#refund_application_fee.
func New(params *FeeRefundParams) (*FeeRefund, error) {
	return getC().New(params)
}

func (c Client) New(params *FeeRefundParams) (*FeeRefund, error) {
	body := &url.Values{}

	body = &url.Values{}

	if params.Amount > 0 {
		body.Add("amount", strconv.FormatUint(params.Amount, 10))
	}

	params.AppendTo(body)

	refund := &FeeRefund{}
	err := c.B.Call("POST", fmt.Sprintf("application_fees/%v/refunds", params.Fee), c.Key, body, refund)

	return refund, err
}

// Get returns the details of a fee refund.
// For more details see https://stripe.com/docs/api#retrieve_fee_refund.
func Get(id string, params *FeeRefundParams) (*FeeRefund, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *FeeRefundParams) (*FeeRefund, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	refund := &FeeRefund{}
	err := c.B.Call("GET", fmt.Sprintf("/application_fees/%v/refunds/%v", params.Fee, id), c.Key, body, refund)

	return refund, err
}

// Update updates a refund's properties.
// For more details see https://stripe.com/docs/api#update_refund.
func Update(id string, params *FeeRefundParams) (*FeeRefund, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *FeeRefundParams) (*FeeRefund, error) {
	body := &url.Values{}

	params.AppendTo(body)

	refund := &FeeRefund{}
	err := c.B.Call("POST", fmt.Sprintf("/application_fees/%v/refunds/%v", params.Fee, id), c.Key, body, refund)

	return refund, err
}

// List returns a list of fee refunds.
// For more details see https://stripe.com/docs/api#list_fee_refunds.
func List(params *FeeRefundListParams) *FeeRefundIter {
	return getC().List(params)
}

func (c Client) List(params *FeeRefundListParams) *FeeRefundIter {
	body := &url.Values{}
	var lp *ListParams

	params.AppendTo(body)
	lp = &params.ListParams

	return &FeeRefundIter{GetIter(lp, body, func(b url.Values) ([]interface{}, ListMeta, error) {
		list := &FeeRefundList{}
		err := c.B.Call("GET", fmt.Sprintf("/application_fees/%v/refunds", params.Fee), c.Key, &b, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func getC() Client {
	return Client{GetBackend(), Key}
}
