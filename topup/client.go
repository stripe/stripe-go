package topup

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /topups APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new topups.
// For more details see https://stripe.com/docs/api#create_topup.
func New(params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.TopupParams) (*stripe.Topup, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	topup := &stripe.Topup{}
	err := c.B.Call("POST", "/topups", c.Key, body, &params.Params, topup)

	return topup, err
}

// Get returns the details of a topup.
// For more details see https://stripe.com/docs/api#retrieve_topup.
func Get(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	topup := &stripe.Topup{}
	err := c.B.Call("GET", stripe.FormatURLPath("/topups/%s", id), c.Key, body, commonParams, topup)

	return topup, err
}

// Update updates a topup's properties.
// For more details see https://stripe.com/docs/api#update_topup.
func Update(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	topup := &stripe.Topup{}
	err := c.B.Call("POST", stripe.FormatURLPath("/topups/%s", id), c.Key, body, commonParams, topup)

	return topup, err
}

// List returns a list of topups.
// For more details see https://stripe.com/docs/api#list_topups.
func List(params *stripe.TopupListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.TopupListParams) *Iter {
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
		list := &stripe.TopupList{}
		err := c.B.Call("GET", "/topups", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Topups.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
