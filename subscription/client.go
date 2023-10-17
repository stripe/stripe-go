//
//
// File generated from our OpenAPI spec
//
//

// Package subscription provides the /subscriptions APIs
package subscription

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
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

// Cancel is the method for the `DELETE /v1/subscriptions/{subscription_exposed_id}` API.
func Cancel(id string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `DELETE /v1/subscriptions/{subscription_exposed_id}` API.
func (c Client) Cancel(id string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, subscription)
	return subscription, err
}

// DeleteDiscount is the method for the `DELETE /v1/subscriptions/{subscription_exposed_id}/discount` API.
func DeleteDiscount(id string, params *stripe.SubscriptionDeleteDiscountParams) (*stripe.Subscription, error) {
	return getC().DeleteDiscount(id, params)
}

// DeleteDiscount is the method for the `DELETE /v1/subscriptions/{subscription_exposed_id}/discount` API.
func (c Client) DeleteDiscount(id string, params *stripe.SubscriptionDeleteDiscountParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s/discount", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, subscription)
	return subscription, err
}

// Resume is the method for the `POST /v1/subscriptions/{subscription}/resume` API.
func Resume(id string, params *stripe.SubscriptionResumeParams) (*stripe.Subscription, error) {
	return getC().Resume(id, params)
}

// Resume is the method for the `POST /v1/subscriptions/{subscription}/resume` API.
func (c Client) Resume(id string, params *stripe.SubscriptionResumeParams) (*stripe.Subscription, error) {
	path := stripe.FormatURLPath("/v1/subscriptions/%s/resume", id)
	subscription := &stripe.Subscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscription)
	return subscription, err
}

// List returns a list of subscriptions.
func List(params *stripe.SubscriptionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of subscriptions.
func (c Client) List(listParams *stripe.SubscriptionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SubscriptionList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/subscriptions", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
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

// Search returns a search result containing subscriptions.
func Search(params *stripe.SubscriptionSearchParams) *SearchIter {
	return getC().Search(params)
}

// Search returns a search result containing subscriptions.
func (c Client) Search(params *stripe.SubscriptionSearchParams) *SearchIter {
	return &SearchIter{
		SearchIter: stripe.GetSearchIter(params, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.SearchContainer, error) {
			list := &stripe.SubscriptionSearchResult{}
			err := c.B.CallRaw(http.MethodGet, "/v1/subscriptions/search", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// SearchIter is an iterator for subscriptions.
type SearchIter struct {
	*stripe.SearchIter
}

// Subscription returns the subscription which the iterator is currently pointing to.
func (i *SearchIter) Subscription() *stripe.Subscription {
	return i.Current().(*stripe.Subscription)
}

// SubscriptionSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) SubscriptionSearchResult() *stripe.SubscriptionSearchResult {
	return i.SearchResult().(*stripe.SubscriptionSearchResult)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
