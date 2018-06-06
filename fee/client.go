// Package fee provides the /application_fees APIs
package fee

import (
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke application_fees APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an application fee.
// For more details see https://stripe.com/docs/api#retrieve_application_fee.
func Get(id string, params *stripe.ApplicationFeeParams) (*stripe.ApplicationFee, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ApplicationFeeParams) (*stripe.ApplicationFee, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	fee := &stripe.ApplicationFee{}
	err := c.B.Call("GET", fmt.Sprintf("application_fees/%v", id), c.Key, body, commonParams, fee)

	return fee, err
}

// List returns a list of application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
func List(params *stripe.ApplicationFeeListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.ApplicationFeeListParams) *Iter {
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
		list := &stripe.ApplicationFeeList{}
		err := c.B.Call("GET", "/application_fees", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of ApplicationFees.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// ApplicationFee returns the most recent ApplicationFee
// visited by a call to Next.
func (i *Iter) ApplicationFee() *stripe.ApplicationFee {
	return i.Current().(*stripe.ApplicationFee)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
