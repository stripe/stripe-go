package stripe

// SubItemParams is the set of parameters that can be used when creating or updating a subscription item.
// For more details see https://stripe.com/docs/api#create_subscription_item and https://stripe.com/docs/api#update_subscription_item.
type SubItemParams struct {
	Params
	Sub                     string
	ID                      string
	Quantity                uint64
	Plan                    string
	ProrationDate           int64
	NoProrate, QuantityZero bool
	Deleted                 bool
}

// SubItemListParams is the set of parameters that can be used when listing invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
type SubItemListParams struct {
	ListParams
	Sub string
}

// SubItem is the resource represneting a Stripe subscription item.
// For more details see https://stripe.com/docs/api#subscription_items.
type SubItem struct {
	ID       string `json:"id"`
	Plan     string `json:"plan"`
	Quantity int64  `json:"quantitiy"`
	Created  int64  `json:"created"`
}

// SubItemList is a list of invoice items as retrieved from a list endpoint.
type SubItemList struct {
	ListMeta
	Values []*SubItem `json:"data"`
}
