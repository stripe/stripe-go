//
//
// File generated from our OpenAPI spec
//
//

// Package sub provides the /subscriptions APIs
package sub

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /subscriptions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new subscription.
func New(params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	return getC().New(params)
}

// New creates a new subscription.
func (c Client) New(params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	subscription := &stripe.Subscription{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/subscriptions",
		c.Key,
		params,
		subscription,
	)
	return subscription, err
}

// Get returns the details of a subscription.
func Get(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	return getC().Get(id, params)
}

// Get returns the details of a subscription.
func (c Client) Get(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, subscription)
	return subscription, err
}

// Update updates a subscription's properties.
func Update(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	return getC().Update(id, params)
}

// Update updates a subscription's properties.
func (c Client) Update(id string, params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscription)
	return subscription, err
}

// Cancel removes a subscription.
func Cancel(id string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
	return getC().Cancel(id, params)
}

// Cancel removes a subscription.
func (c Client) Cancel(id string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, subscription)
	return subscription, err
}

// List returns a list of subscriptions.
func List(params *stripe.SubscriptionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of subscriptions.
func (c Client) List(listParams *stripe.SubscriptionListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.SubscriptionList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/subscriptions", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for subscriptions.
type Iter struct {
	*stripe.Iter
}

// Subscription returns the subscription which the iterator is currently pointing to.
func (i *Iter) Subscription() *stripe.Subscription {
	return i.Current().(*stripe.Subscription)
}

// SubscriptionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SubscriptionList() *stripe.SubscriptionList {
	return i.List().(*stripe.SubscriptionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
