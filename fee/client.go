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
	Key string
}

// Get returns the details of an application fee.
// For more details see https://stripe.com/docs/api#retrieve_application_fee.
func Get(id string, params *FeeParams) (*Fee, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *FeeParams) (*Fee, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	fee := &Fee{}
	err := c.B.Call("POST", fmt.Sprintf("application_fees/%v/refund", id), c.Key, body, fee)

	return fee, err
}

// List returns a list of application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
func List(params *FeeListParams) *FeeIter {
	return getC().List(params)
}

func (c Client) List(params *FeeListParams) *FeeIter {
	type feeList struct {
		ListMeta
		Values []*Fee `json:"data"`
	}

	var body *url.Values
	var lp *ListParams

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

	return &FeeIter{GetIter(lp, body, func(b url.Values) ([]interface{}, ListMeta, error) {
		list := &feeList{}
		err := c.B.Call("GET", "/application_fees", c.Key, &b, list)

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
