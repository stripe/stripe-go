package stripe

import "encoding/json"

// SubscriptionItemParams is the set of parameters that can be used when creating or updating a subscription item.
// For more details see https://stripe.com/docs/api#create_subscription_item and https://stripe.com/docs/api#update_subscription_item.
type SubscriptionItemParams struct {
	Params
	Subscription           string
	Quantity             int64
	Currency Currency
}

// InvoiceItemListParams is the set of parameters that can be used when listing invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
type SubscriptionItemListParams struct {
	ListParams
	Subscription string
}

// SubscriptionItem is the resource represneting a Stripe subscription item.
// For more details see https://stripe.com/docs/api#subscription_items.
type SubscriptionItem struct {
	ID           string            `json:"id"`
	Plan       string             `json:"plan"`
	Quantity     int64          `json:"quantitiy"`
	Created int64 `json:"created"`
	}

// InvoiceItemList is a list of invoice items as retrieved from a list endpoint.
type SubscriptionItemList struct {
	ListMeta
	Values []*SubscriptionItem `json:"data"`
}
