//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of your subscription items for a given subscription.
type SubscriptionItemListParams struct {
	ListParams   `form:"*"`
	Subscription *string `form:"subscription"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. When updating, pass an empty string to remove previously-defined thresholds.
type SubscriptionItemBillingThresholdsParams struct {
	UsageGTE *int64 `form:"usage_gte"`
}

// The recurring components of a price such as `interval` and `usage_type`.
type SubscriptionItemPriceDataRecurringParams struct {
	Interval      *string `form:"interval"`
	IntervalCount *int64  `form:"interval_count"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
type SubscriptionItemPriceDataParams struct {
	Currency          *string                                   `form:"currency"`
	Product           *string                                   `form:"product"`
	Recurring         *SubscriptionItemPriceDataRecurringParams `form:"recurring"`
	TaxBehavior       *string                                   `form:"tax_behavior"`
	UnitAmount        *int64                                    `form:"unit_amount"`
	UnitAmountDecimal *float64                                  `form:"unit_amount_decimal,high_precision"`
}

// Adds a new item to an existing subscription. No existing items will be changed or replaced.
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

// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
type SubscriptionItemBillingThresholds struct {
	UsageGTE int64 `form:"usage_gte"`
}

// Subscription items allow you to create customer subscriptions with more than
// one plan, making it easy to represent complex billing relationships.
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

// SubscriptionItemList is a list of SubscriptionItems as retrieved from a list endpoint.
type SubscriptionItemList struct {
	APIResource
	ListMeta
	Data []*SubscriptionItem `json:"data"`
}
