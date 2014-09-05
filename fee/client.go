// package fee provides the /application_fees APIs
package fee

import (
	"fmt"
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke application_fees APIs.
type Client struct {
	B   Backend
	Tok string
}

// Get returns the details of an application fee.
// For more details see https://stripe.com/docs/api#retrieve_application_fee.
func Get(id string, params *AppFeeParams) (*AppFee, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *AppFeeParams) (*AppFee, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	fee := &AppFee{}
	err := c.B.Call("POST", fmt.Sprintf("application_fees/%v/refund", id), c.Tok, body, fee)

	return fee, err
}

// RefundFee refunds the application fee collected.
// For more details see https://stripe.com/docs/api#refund_application_fee.
func RefundFee(id string, params *AppFeeParams) (*FeeRefund, error) {
	return getC().Refund(id, params)
}

func (c Client) Refund(id string, params *AppFeeParams) (*FeeRefund, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Amount > 0 {
			body.Add("amount", strconv.FormatUint(params.Amount, 10))
		}

		params.AppendTo(body)
	}

	refund := &FeeRefund{}
	err := c.B.Call("POST", fmt.Sprintf("application_fees/%v/refunds", id), c.Tok, body, refund)

	return refund, err
}

// List returns a list of application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
func List(params *AppFeeListParams) (*AppFeeList, error) {
	return getC().List(params)
}

func (c Client) List(params *AppFeeListParams) (*AppFeeList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Charge) > 0 {
			body.Add("charge", params.Charge)
		}

		params.AppendTo(body)
	}

	list := &AppFeeList{}
	err := c.B.Call("GET", "/application_fees", c.Tok, body, list)

	return list, err
}

func getC() Client {
	return Client{GetBackend(), Key}
}
