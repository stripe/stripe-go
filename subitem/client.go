// Package subitem provides the /subscription_items APIs
package subitem

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /subscriptions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTS a new subscription for a customer.
// For more details see https://stripe.com/docs/api#create_subscription_item.
func New(params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	item := &stripe.SubscriptionItem{}
	err := c.B.Call2("POST", "/subscription_items", c.Key, params, item)
	return item, err
}

// Get returns the details of a subscription.
// For more details see https://stripe.com/docs/api#retrieve_subscription.
func Get(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	path := stripe.FormatURLPath("/subscription_items/%s", id)
	item := &stripe.SubscriptionItem{}
	err := c.B.Call2("GET", path, c.Key, params, item)
	return item, err
}

// Update updates a subscription's properties.
// For more details see https://stripe.com/docs/api#update_subscription.
func Update(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	path := stripe.FormatURLPath("/subscription_items/%s", id)
	subi := &stripe.SubscriptionItem{}
	err := c.B.Call2("POST", path, c.Key, params, subi)
	return subi, err
}

// Del removes a subscription item.
// For more details see https://stripe.com/docs/api#cancel_subscription.
func Del(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.SubscriptionItemParams) (*stripe.SubscriptionItem, error) {
	path := stripe.FormatURLPath("/subscription_items/%s", id)
	item := &stripe.SubscriptionItem{}
	err := c.B.Call2("DELETE", path, c.Key, params, item)

	return item, err
}

// List returns a list of subscription items.
// For more details see https://stripe.com/docs/api#list_subscription_items.
func List(params *stripe.SubscriptionItemListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.SubscriptionItemListParams) *Iter {
	return &Iter{stripe.GetIter2(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.SubscriptionItemList{}
		err := c.B.CallRaw("GET", "/subscription_items", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Subscriptions.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// SubscriptionItem returns the most recent SubscriptionItem
// visited by a call to Next.
func (i *Iter) SubscriptionItem() *stripe.SubscriptionItem {
	return i.Current().(*stripe.SubscriptionItem)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
