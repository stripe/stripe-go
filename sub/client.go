// Package sub provides the /subscriptions APIs
package sub

import (
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	Trialing stripe.SubscriptionStatus = "trialing"
	Active   stripe.SubscriptionStatus = "active"
	PastDue  stripe.SubscriptionStatus = "past_due"
	Canceled stripe.SubscriptionStatus = "canceled"
	Unpaid   stripe.SubscriptionStatus = "unpaid"
	All      stripe.SubscriptionStatus = "all"
)

// Client is used to invoke /subscriptions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTS a new subscription for a customer.
// For more details see https://stripe.com/docs/api#create_subscription.
func New(params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	sub := &stripe.Subscription{}
	err := c.B.Call("POST", "/subscriptions", c.Key, body, commonParams, sub)

	return sub, err
}

// Get returns the details of a subscription.
// For more details see https://stripe.com/docs/api#retrieve_subscription.
func Get(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	sub := &stripe.Subscription{}
	err := c.B.Call("GET", fmt.Sprintf("/subscriptions/%v", id), c.Key, body, commonParams, sub)

	return sub, err
}

// Update updates a subscription's properties.
// For more details see https://stripe.com/docs/api#update_subscription.
func Update(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	sub := &stripe.Subscription{}
	err := c.B.Call("POST", fmt.Sprintf("/subscriptions/%v", id), c.Key, body, commonParams, sub)

	return sub, err
}

// Cancel removes a subscription.
// For more details see https://stripe.com/docs/api#cancel_subscription.
func Cancel(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	return getC().Cancel(id, params)
}

func (c Client) Cancel(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	sub := &stripe.Subscription{}
	err := c.B.Call("DELETE", fmt.Sprintf("/subscriptions/%v", id), c.Key, body, commonParams, sub)

	return sub, err
}

// List returns a list of subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
func List(params *stripe.SubscriptionListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.SubscriptionListParams) *Iter {
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
		list := &stripe.SubscriptionList{}
		err := c.B.Call("GET", "/subscriptions", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
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

// Subscription returns the most recent Subscription
// visited by a call to Next.
func (i *Iter) Subscription() *stripe.Subscription {
	return i.Current().(*stripe.Subscription)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
