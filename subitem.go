//
//
// File generated from our OpenAPI spec
//
//

package stripe

// SubscriptionItemPriceDataRecurringParams is a structure representing the parameters to create
// an inline recurring price for a subscription item.
type SubscriptionItemPriceDataRecurringParams struct {
	Interval      *string `form:"interval"`
	IntervalCount *int64  `form:"interval_count"`
}

// SubscriptionItemPriceDataParams is a structure representing the parameters to create an inline
// price for a subscription item.
type SubscriptionItemPriceDataParams struct {
	Currency          *string                                   `form:"currency"`
	Product           *string                                   `form:"product"`
	Recurring         *SubscriptionItemPriceDataRecurringParams `form:"recurring"`
	TaxBehavior       *string                                   `form:"tax_behavior"`
	UnitAmount        *int64                                    `form:"unit_amount"`
	UnitAmountDecimal *float64                                  `form:"unit_amount_decimal,high_precision"`
}

// SubscriptionItemParams is the set of parameters that can be used when creating or updating a subscription item.
// For more details see https://stripe.com/docs/api#create_subscription_item and https://stripe.com/docs/api#update_subscription_item.
type SubscriptionItemParams struct {
	Params            `form:"*"`
	BillingThresholds *SubscriptionItemBillingThresholdsParams `form:"billing_thresholds"`
	ClearUsage        *bool                                    `form:"clear_usage"`
	OffSession        *bool                                    `form:"off_session"` // Only supported on update
	PaymentBehavior   *string                                  `form:"payment_behavior"`
	Plan              *string                                  `form:"plan"`
	Price             *string                                  `form:"price"`
	PriceData         *SubscriptionItemPriceDataParams         `form:"price_data"`
	ProrationBehavior *string                                  `form:"proration_behavior"`
	ProrationDate     *int64                                   `form:"proration_date"`
	Quantity          *int64                                   `form:"quantity"`
	Subscription      *string                                  `form:"subscription"`
	TaxRates          []*string                                `form:"tax_rates"`

	ID *string `form:"-"` // Deprecated
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
	APIResource
	BillingThresholds SubscriptionItemBillingThresholds `json:"billing_thresholds"`
	Created           int64                             `json:"created"`
	Deleted           bool                              `json:"deleted"`
	ID                string                            `json:"id"`
	Metadata          map[string]string                 `json:"metadata"`
	Object            string                            `json:"object"`
	Plan              *Plan                             `json:"plan"`
	Price             *Price                            `json:"price"`
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
	APIResource
	ListMeta
	Data []*SubscriptionItem `json:"data"`
}
