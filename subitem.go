package stripe

// SubscriptionItemParams is the set of parameters that can be used when creating or updating a subscription item.
// For more details see https://stripe.com/docs/api#create_subscription_item and https://stripe.com/docs/api#update_subscription_item.
type SubscriptionItemParams struct {
	Params            `form:"*"`
	ID                *string                                  `form:"-"` // Handled in URL
	BillingThresholds *SubscriptionItemBillingThresholdsParams `form:"billing_thresholds"`
	ClearUsage        *bool                                    `form:"clear_usage"`
	PaymentBehavior   *string                                  `form:"payment_behavior"`
	Plan              *string                                  `form:"plan"`
	Prorate           *bool                                    `form:"prorate"`
	ProrationDate     *int64                                   `form:"proration_date"`
	ProrationBehavior *string                                  `form:"proration_behavior"`
	Quantity          *int64                                   `form:"quantity"`
	Subscription      *string                                  `form:"subscription"`
	TaxRates          []*string                                `form:"tax_rates"`

	// The following parameters are only supported on updates
	OffSession *bool `form:"off_session"`
}

// SubscriptionItemBillingThresholdsParams is a structure representing the parameters allowed to
// control billing thresholds for a subscription item.
type SubscriptionItemBillingThresholdsParams struct {
	UsageGTE *int64 `form:"usage_gte"`
}

// SubscriptionItemListParams is the set of parameters that can be used when listing invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
type SubscriptionItemListParams struct {
	ListParams   `form:"*"`
	Subscription *string `form:"subscription"`
}

// SubscriptionItem is the resource representing a Stripe subscription item.
// For more details see https://stripe.com/docs/api#subscription_items.
type SubscriptionItem struct {
	BillingThresholds SubscriptionItemBillingThresholds `json:"billing_thresholds"`
	Created           int64                             `json:"created"`
	Deleted           bool                              `json:"deleted"`
	ID                string                            `json:"id"`
	Metadata          map[string]string                 `json:"metadata"`
	Plan              *Plan                             `json:"plan"`
	Quantity          int64                             `json:"quantity"`
	Subscription      string                            `json:"subscription"`
	TaxRates          []*TaxRate                        `json:"tax_rates"`
}

// SubscriptionItemBillingThresholds is a structure representing the billing thresholds for a
// subscription item.
type SubscriptionItemBillingThresholds struct {
	UsageGTE int64 `form:"usage_gte"`
}

// SubscriptionItemList is a list of invoice items as retrieved from a list endpoint.
type SubscriptionItemList struct {
	ListMeta
	Data []*SubscriptionItem `json:"data"`
}
