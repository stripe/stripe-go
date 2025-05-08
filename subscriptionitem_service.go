//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1SubscriptionItemService is used to invoke /v1/subscription_items APIs.
type v1SubscriptionItemService struct {
	B   Backend
	Key string
}

// Adds a new item to an existing subscription. No existing items will be changed or replaced.
func (c v1SubscriptionItemService) Create(ctx context.Context, params *SubscriptionItemCreateParams) (*SubscriptionItem, error) {
	if params == nil {
		params = &SubscriptionItemCreateParams{}
	}
	params.Context = ctx
	subscriptionitem := &SubscriptionItem{}
	err := c.B.Call(
		http.MethodPost, "/v1/subscription_items", c.Key, params, subscriptionitem)
	return subscriptionitem, err
}

// Retrieves the subscription item with the given ID.
func (c v1SubscriptionItemService) Retrieve(ctx context.Context, id string, params *SubscriptionItemRetrieveParams) (*SubscriptionItem, error) {
	if params == nil {
		params = &SubscriptionItemRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/subscription_items/%s", id)
	subscriptionitem := &SubscriptionItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, subscriptionitem)
	return subscriptionitem, err
}

// Updates the plan or quantity of an item on a current subscription.
func (c v1SubscriptionItemService) Update(ctx context.Context, id string, params *SubscriptionItemUpdateParams) (*SubscriptionItem, error) {
	if params == nil {
		params = &SubscriptionItemUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/subscription_items/%s", id)
	subscriptionitem := &SubscriptionItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscriptionitem)
	return subscriptionitem, err
}

// Deletes an item from the subscription. Removing a subscription item from a subscription will not cancel the subscription.
func (c v1SubscriptionItemService) Delete(ctx context.Context, id string, params *SubscriptionItemDeleteParams) (*SubscriptionItem, error) {
	if params == nil {
		params = &SubscriptionItemDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/subscription_items/%s", id)
	subscriptionitem := &SubscriptionItem{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, subscriptionitem)
	return subscriptionitem, err
}

// Returns a list of your subscription items for a given subscription.
func (c v1SubscriptionItemService) List(ctx context.Context, listParams *SubscriptionItemListParams) Seq2[*SubscriptionItem, error] {
	if listParams == nil {
		listParams = &SubscriptionItemListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*SubscriptionItem, ListContainer, error) {
		list := &SubscriptionItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/subscription_items", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
