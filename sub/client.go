// Package sub provides the /subscriptions APIs
package sub

import (
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	Trialing stripe.SubStatus = "trialing"
	Active   stripe.SubStatus = "active"
	PastDue  stripe.SubStatus = "past_due"
	Canceled stripe.SubStatus = "canceled"
	Unpaid   stripe.SubStatus = "unpaid"
	All      stripe.SubStatus = "all"
)

// Client is used to invoke /subscriptions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTS a new subscription for a customer.
// For more details see https://stripe.com/docs/api#create_subscription.
func New(params *stripe.SubParams) (*stripe.Sub, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.SubParams) (*stripe.Sub, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	sub := &stripe.Sub{}
	err := c.B.Call("POST", "/subscriptions", c.Key, body, commonParams, sub)

	return sub, err
}

// Get returns the details of a subscription.
// For more details see https://stripe.com/docs/api#retrieve_subscription.
func Get(id string, params *stripe.SubParams) (*stripe.Sub, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.SubParams) (*stripe.Sub, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	sub := &stripe.Sub{}
	err := c.B.Call("GET", fmt.Sprintf("/subscriptions/%v", id), c.Key, body, commonParams, sub)

	return sub, err
}

// Update updates a subscription's properties.
// For more details see https://stripe.com/docs/api#update_subscription.
func Update(id string, params *stripe.SubParams) (*stripe.Sub, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.SubParams) (*stripe.Sub, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	sub := &stripe.Sub{}
	err := c.B.Call("POST", fmt.Sprintf("/subscriptions/%v", id), c.Key, body, commonParams, sub)

	return sub, err
}

// Cancel removes a subscription.
// For more details see https://stripe.com/docs/api#cancel_subscription.
func Cancel(id string, params *stripe.SubParams) (*stripe.Sub, error) {
	return getC().Cancel(id, params)
}

func (c Client) Cancel(id string, params *stripe.SubParams) (*stripe.Sub, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	sub := &stripe.Sub{}
	err := c.B.Call("DELETE", fmt.Sprintf("/subscriptions/%v", id), c.Key, body, commonParams, sub)

	return sub, err
}

// List returns a list of subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
func List(params *stripe.SubListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.SubListParams) *Iter {
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
		list := &stripe.SubList{}
		err := c.B.Call("GET", "/subscriptions", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Subs.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Sub returns the most recent Sub
// visited by a call to Next.
func (i *Iter) Sub() *stripe.Sub {
	return i.Current().(*stripe.Sub)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
