// package fee provides the /application_fees APIs
package fee

import (
	"fmt"
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke application_fees APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an application fee.
// For more details see https://stripe.com/docs/api#retrieve_application_fee.
func Get(id string, params *stripe.FeeParams) (*stripe.Fee, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.FeeParams) (*stripe.Fee, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	fee := &stripe.Fee{}
	err := c.B.Call("POST", fmt.Sprintf("application_fees/%v/refund", id), c.Key, body, fee)

	return fee, err
}

// List returns a list of application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
func List(params *stripe.FeeListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.FeeListParams) *Iter {
	type feeList struct {
		stripe.ListMeta
		Values []*stripe.Fee `json:"data"`
	}

	var body *url.Values
	var lp *stripe.ListParams

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Charge) > 0 {
			body.Add("charge", params.Charge)
		}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &feeList{}
		err := c.B.Call("GET", "/application_fees", c.Key, &b, list)

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
func (i *Iter) Next() (*stripe.Fee, error) {
	f, err := i.Iter.Next()
	if err != nil {
		return nil, err
	}

	return f.(*stripe.Fee), err
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
