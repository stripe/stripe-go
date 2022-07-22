//
//
// File generated from our OpenAPI spec
//
//

// Package subitem provides the /subscription_items APIs
package subitem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /subscription_items APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new subscription item.
func New(params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	return getC().New(params)
}

// New creates a new subscription item.
func (c Client) New(params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	subscriptionitem := &stripe.SubscriptionItem{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/subscription_items",
		c.Key,
		params,
		subscriptionitem,
	)
	return subscriptionitem, err
}

// Get returns the details of a subscription item.
func Get(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	return getC().Get(id, params)
}

// Get returns the details of a subscription item.
func (c Client) Get(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	path := stripe.FormatURLPath("/v1/subscription_items/%s", id)
	subscriptionitem := &stripe.SubscriptionItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, subscriptionitem)
	return subscriptionitem, err
}

// Update updates a subscription item's properties.
func Update(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	return getC().Update(id, params)
}

// Update updates a subscription item's properties.
func (c Client) Update(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	path := stripe.FormatURLPath("/v1/subscription_items/%s", id)
	subscriptionitem := &stripe.SubscriptionItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscriptionitem)
	return subscriptionitem, err
}

// Del removes a subscription item.
func Del(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	return getC().Del(id, params)
}

// Del removes a subscription item.
func (c Client) Del(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	path := stripe.FormatURLPath("/v1/subscription_items/%s", id)
	subscriptionitem := &stripe.SubscriptionItem{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, subscriptionitem)
	return subscriptionitem, err
}

// List returns a list of subscription items.
func List(params *stripe.SubscriptionItemListParams) *Iter {
	return getC().List(params)
}

// List returns a list of subscription items.
func (c Client) List(listParams *stripe.SubscriptionItemListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SubscriptionItemList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/subscription_items", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for subscription items.
type Iter struct {
	*stripe.Iter
}

// SubscriptionItem returns the subscription item which the iterator is currently pointing to.
func (i *Iter) SubscriptionItem() *stripe.SubscriptionItem {
	return i.Current().(*stripe.SubscriptionItem)
}

// SubscriptionItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SubscriptionItemList() *stripe.SubscriptionItemList {
	return i.List().(*stripe.SubscriptionItemList)
}

// UsageRecordSummaries is the method for the `GET /v1/subscription_items/{subscription_item}/usage_record_summaries` API.
func UsageRecordSummaries(params *stripe.SubscriptionItemUsageRecordSummariesParams) *UsageRecordSummaryIter {
	return getC().UsageRecordSummaries(params)
}

// UsageRecordSummaries is the method for the `GET /v1/subscription_items/{subscription_item}/usage_record_summaries` API.
func (c Client) UsageRecordSummaries(listParams *stripe.SubscriptionItemUsageRecordSummariesParams) *UsageRecordSummaryIter {
	path := stripe.FormatURLPath(
		"/v1/subscription_items/%s/usage_record_summaries",
		stripe.StringValue(listParams.SubscriptionItem),
	)
	return &UsageRecordSummaryIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.UsageRecordSummaryList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// UsageRecordSummaryIter is an iterator for usage record summaries.
type UsageRecordSummaryIter struct {
	*stripe.Iter
}

// UsageRecordSummary returns the usage record summary which the iterator is currently pointing to.
func (i *UsageRecordSummaryIter) UsageRecordSummary() *stripe.UsageRecordSummary {
	return i.Current().(*stripe.UsageRecordSummary)
}

// UsageRecordSummaryList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *UsageRecordSummaryIter) UsageRecordSummaryList() *stripe.UsageRecordSummaryList {
	return i.List().(*stripe.UsageRecordSummaryList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
