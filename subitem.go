//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of your subscription items for a given subscription.
type SubscriptionItemListParams struct {
	ListParams `form:"*"`
	// The ID of the subscription whose items will be retrieved.
	Subscription *string `form:"subscription"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. When updating, pass an empty string to remove previously-defined thresholds.
type SubscriptionItemBillingThresholdsParams struct {
	// Usage threshold that triggers the subscription to advance to a new billing period
	UsageGTE *int64 `form:"usage_gte"`
}

// The recurring components of a price such as `interval` and `interval_count`.
type SubscriptionItemPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
type SubscriptionItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the product that this price will belong to.
	Product *string `form:"product"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *SubscriptionItemPriceDataRecurringParams `form:"recurring"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Adds a new item to an existing subscription. No existing items will be changed or replaced.
type SubscriptionItemParams struct {
	Params `form:"*"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. When updating, pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionItemBillingThresholdsParams `form:"billing_thresholds"`
	// Delete all usage for the given subscription item. Allowed only when the current plan's `usage_type` is `metered`.
	ClearUsage *bool `form:"clear_usage"`
	OffSession *bool `form:"off_session"` // Only supported on update
	// Use `allow_incomplete` to transition the subscription to `status=past_due` if a payment is required but cannot be paid. This allows you to manage scenarios where additional user actions are needed to pay a subscription's invoice. For example, SCA regulation may require 3DS authentication to complete payment. See the [SCA Migration Guide](https://stripe.com/docs/billing/migration/strong-customer-authentication) for Billing to learn more. This is the default behavior.
	//
	// Use `default_incomplete` to transition the subscription to `status=past_due` when payment is required and await explicit confirmation of the invoice's payment intent. This allows simpler management of scenarios where additional user actions are needed to pay a subscription's invoice. Such as failed payments, [SCA regulation](https://stripe.com/docs/billing/migration/strong-customer-authentication), or collecting a mandate for a bank debit payment method.
	//
	// Use `pending_if_incomplete` to update the subscription using [pending updates](https://stripe.com/docs/billing/subscriptions/pending-updates). When you use `pending_if_incomplete` you can only pass the parameters [supported by pending updates](https://stripe.com/docs/billing/pending-updates-reference#supported-attributes).
	//
	// Use `error_if_incomplete` if you want Stripe to return an HTTP 402 status code if a subscription's invoice cannot be paid. For example, if a payment method requires 3DS authentication due to SCA regulation and further user action is needed, this parameter does not update the subscription and returns an error instead. This was the default behavior for API versions prior to 2019-03-14. See the [changelog](https://stripe.com/docs/upgrades#2019-03-14) to learn more.
	PaymentBehavior *string `form:"payment_behavior"`
	// The identifier of the new plan for this subscription item.
	Plan *string `form:"plan"`
	// The ID of the price object. When changing a subscription item's price, `quantity` is set to 1 unless a `quantity` parameter is provided.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *SubscriptionItemPriceDataParams `form:"price_data"`
	// Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes.
	ProrationBehavior *string `form:"proration_behavior"`
	// If set, the proration will be calculated as though the subscription was updated at the given time. This can be used to apply the same proration that was previewed with the [upcoming invoice](https://stripe.com/docs/api#retrieve_customer_invoice) endpoint.
	ProrationDate *int64 `form:"proration_date"`
	// The quantity you'd like to apply to the subscription item you're creating.
	Quantity *int64 `form:"quantity"`
	// The identifier of the subscription to modify.
	Subscription *string `form:"subscription"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`

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
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	Deleted bool  `json:"deleted"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// You can now model subscriptions more flexibly using the [Prices API](https://stripe.com/docs/api#prices). It replaces the Plans API and is backwards compatible to simplify your migration.
	//
	// Plans define the base price, currency, and billing cycle for recurring purchases of products.
	// [Products](https://stripe.com/docs/api#products) help you track inventory or provisioning, and plans help you track pricing. Different physical goods or levels of service should be represented by products, and pricing options should be represented by plans. This approach lets you change prices without having to change your provisioning scheme.
	//
	// For example, you might have a single "gold" product that has plans for $10/month, $100/year, €9/month, and €90/year.
	//
	// Related guides: [Set up a subscription](https://stripe.com/docs/billing/subscriptions/set-up-subscription) and more about [products and prices](https://stripe.com/docs/products-prices/overview).
	Plan *Plan `json:"plan"`
	// Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.
	// [Products](https://stripe.com/docs/api#products) help you track inventory or provisioning, and prices help you track payment terms. Different physical goods or levels of service should be represented by products, and pricing options should be represented by prices. This approach lets you change prices without having to change your provisioning scheme.
	//
	// For example, you might have a single "gold" product that has prices for $10/month, $100/year, and €9 once.
	//
	// Related guides: [Set up a subscription](https://stripe.com/docs/billing/subscriptions/set-up-subscription), [create an invoice](https://stripe.com/docs/billing/invoices/create), and more about [products and prices](https://stripe.com/docs/products-prices/overview).
	Price *Price `json:"price"`
	// The [quantity](https://stripe.com/docs/subscriptions/quantities) of the plan to which the customer should be subscribed.
	Quantity int64 `json:"quantity"`
	// The `subscription` this `subscription_item` belongs to.
	Subscription string `json:"subscription"`
	// The tax rates which apply to this `subscription_item`. When set, the `default_tax_rates` on the subscription do not apply to this `subscription_item`.
	TaxRates []*TaxRate `json:"tax_rates"`
}

// SubscriptionItemList is a list of SubscriptionItems as retrieved from a list endpoint.
type SubscriptionItemList struct {
	APIResource
	ListMeta
	Data []*SubscriptionItem `json:"data"`
}
