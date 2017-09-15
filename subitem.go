package stripe

// SubItemParams is the set of parameters that can be used when creating or updating a subscription item.
// For more details see https://stripe.com/docs/api#create_subscription_item and https://stripe.com/docs/api#update_subscription_item.
type SubItemParams struct {
	Params        `form:"*"`
	ID            string `form:"-"` // Handled in URL
	NoProrate     bool   `form:"prorate,invert"`
	Plan          string `form:"plan"`
	ProrationDate int64  `form:"proration_date"`
	Quantity      uint64 `form:"quantity"`
	QuantityZero  bool   `form:"quantity,zero"`
	Sub           string `form:"subscription"`
}

// SubItemListParams is the set of parameters that can be used when listing invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
type SubItemListParams struct {
	ListParams `form:"*"`
	Sub        string `form:"subscription"`
}

// SubItem is the resource representing a Stripe subscription item.
// For more details see https://stripe.com/docs/api#subscription_items.
type SubItem struct {
	Created  int64             `json:"created"`
	Deleted  bool              `json:"deleted"`
	ID       string            `json:"id"`
	Meta     map[string]string `json:"metadata"`
	Plan     *Plan             `json:"plan"`
	Quantity uint64            `json:"quantity"`
}

// SubItemList is a list of invoice items as retrieved from a list endpoint.
type SubItemList struct {
	ListMeta
	Values []*SubItem `json:"data"`
}
