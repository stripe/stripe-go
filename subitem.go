package stripe

// SubscriptionItemParams is the set of parameters that can be used when creating or updating a subscription item.
// For more details see https://stripe.com/docs/api#create_subscription_item and https://stripe.com/docs/api#update_subscription_item.
type SubscriptionItemParams struct {
	Params        `form:"*"`
	ID            string `form:"-"` // Handled in URL
	NoProrate     bool   `form:"prorate,invert"`
	Plan          string `form:"plan"`
	ProrationDate int64  `form:"proration_date"`
	Quantity      uint64 `form:"quantity"`
	QuantityZero  bool   `form:"quantity,zero"`
	Subscription  string `form:"subscription"`
}

// SubscriptionItemListParams is the set of parameters that can be used when listing invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
type SubscriptionItemListParams struct {
	ListParams   `form:"*"`
	Subscription string `form:"subscription"`
}

// SubscriptionItem is the resource representing a Stripe subscription item.
// For more details see https://stripe.com/docs/api#subscription_items.
type SubscriptionItem struct {
	Created  int64             `json:"created"`
	Deleted  bool              `json:"deleted"`
	ID       string            `json:"id"`
	Metadata map[string]string `json:"metadata"`
	Plan     *Plan             `json:"plan"`
	Quantity uint64            `json:"quantity"`
}

// SubscriptionItemList is a list of invoice items as retrieved from a list endpoint.
type SubscriptionItemList struct {
	ListMeta
	Data []*SubscriptionItem `json:"data"`
}
