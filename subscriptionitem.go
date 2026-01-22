//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Determines the type of trial for this item.
type SubscriptionItemTrialType string

// List of values that SubscriptionItemTrialType can take
const (
	SubscriptionItemTrialTypeFree SubscriptionItemTrialType = "free"
	SubscriptionItemTrialTypePaid SubscriptionItemTrialType = "paid"
)

// Deletes an item from the subscription. Removing a subscription item from a subscription will not cancel the subscription.
type SubscriptionItemParams struct {
	Params `form:"*"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionItemBillingThresholdsParams `form:"billing_thresholds"`
	// Delete all usage for the given subscription item. Allowed only when the current plan's `usage_type` is `metered`.
	ClearUsage *bool `form:"clear_usage"`
	// The trial offer to apply to this subscription item.
	CurrentTrial *SubscriptionItemCurrentTrialParams `form:"current_trial"`
	// The coupons to redeem into discounts for the subscription item.
	Discounts []*SubscriptionItemDiscountParams `form:"discounts"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Only supported on update
	// Indicates if a customer is on or off-session while an invoice payment is attempted. Defaults to `false` (on-session).
	OffSession *bool `form:"off_session"`
	// Use `allow_incomplete` to transition the subscription to `status=past_due` if a payment is required but cannot be paid. This allows you to manage scenarios where additional user actions are needed to pay a subscription's invoice. For example, SCA regulation may require 3DS authentication to complete payment. See the [SCA Migration Guide](https://docs.stripe.com/billing/migration/strong-customer-authentication) for Billing to learn more. This is the default behavior.
	//
	// Use `default_incomplete` to transition the subscription to `status=past_due` when payment is required and await explicit confirmation of the invoice's payment intent. This allows simpler management of scenarios where additional user actions are needed to pay a subscription's invoice. Such as failed payments, [SCA regulation](https://docs.stripe.com/billing/migration/strong-customer-authentication), or collecting a mandate for a bank debit payment method.
	//
	// Use `pending_if_incomplete` to update the subscription using [pending updates](https://docs.stripe.com/billing/subscriptions/pending-updates). When you use `pending_if_incomplete` you can only pass the parameters [supported by pending updates](https://docs.stripe.com/billing/pending-updates-reference#supported-attributes).
	//
	// Use `error_if_incomplete` if you want Stripe to return an HTTP 402 status code if a subscription's invoice cannot be paid. For example, if a payment method requires 3DS authentication due to SCA regulation and further user action is needed, this parameter does not update the subscription and returns an error instead. This was the default behavior for API versions prior to 2019-03-14. See the [changelog](https://docs.stripe.com/changelog/2019-03-14) to learn more.
	PaymentBehavior *string `form:"payment_behavior"`
	// The identifier of the new plan for this subscription item.
	Plan *string `form:"plan"`
	// The ID of the price object. One of `price` or `price_data` is required. When changing a subscription item's price, `quantity` is set to 1 unless a `quantity` parameter is provided.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *SubscriptionItemPriceDataParams `form:"price_data"`
	// Determines how to handle [prorations](https://docs.stripe.com/billing/subscriptions/prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes. The default value is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior"`
	// If set, the proration will be calculated as though the subscription was updated at the given time. This can be used to apply the same proration that was previewed with the [upcoming invoice](https://api.stripe.com#retrieve_customer_invoice) endpoint.
	ProrationDate *int64 `form:"proration_date"`
	// The quantity you'd like to apply to the subscription item you're creating.
	Quantity *int64 `form:"quantity"`
	// The identifier of the subscription to modify.
	Subscription *string `form:"subscription"`
	// A list of [Tax Rate](https://docs.stripe.com/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://docs.stripe.com/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionItemTrialParams `form:"trial"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionItemParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionItemBillingThresholdsParams struct {
	// Number of units that meets the billing threshold to advance the subscription to a new billing period (e.g., it takes 10 $5 units to meet a $50 [monetary threshold](https://docs.stripe.com/api/subscriptions/update#update_subscription-billing_thresholds-amount_gte))
	UsageGTE *int64 `form:"usage_gte"`
}

// The trial offer to apply to this subscription item.
type SubscriptionItemCurrentTrialParams struct {
	// Unix timestamp representing the end of the trial offer period. Required when the trial offer has `duration.type=timestamp`. Cannot be specified when `duration.type=relative`.
	TrialEnd *int64 `form:"trial_end"`
	// The ID of the trial offer to apply to the subscription item.
	TrialOffer *string `form:"trial_offer"`
}

// Time span for the redeemed discount.
type SubscriptionItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionItemDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// Anchor the service period to a custom date. Type must be `custom` to specify.
type SubscriptionItemDiscountSettingsServicePeriodAnchorConfigCustomParams struct {
	// The day of the month the anchor should be. Ranges from 1 to 31.
	DayOfMonth *int64 `form:"day_of_month"`
	// The hour of the day the anchor should be. Ranges from 0 to 23.
	Hour *int64 `form:"hour"`
	// The minute of the hour the anchor should be. Ranges from 0 to 59.
	Minute *int64 `form:"minute"`
	// The month to start full cycle periods. Ranges from 1 to 12.
	Month *int64 `form:"month"`
	// The second of the minute the anchor should be. Ranges from 0 to 59.
	Second *int64 `form:"second"`
}

// Configures service period cycle anchoring.
type SubscriptionItemDiscountSettingsServicePeriodAnchorConfigParams struct {
	// Anchor the service period to a custom date. Type must be `custom` to specify.
	Custom *SubscriptionItemDiscountSettingsServicePeriodAnchorConfigCustomParams `form:"custom"`
	// The type of service period anchor config. Defaults to `subscription_service_cycle_anchor` if omitted.
	Type *string `form:"type"`
}

// Settings for discount application including service period anchoring.
type SubscriptionItemDiscountSettingsParams struct {
	// Configures service period cycle anchoring.
	ServicePeriodAnchorConfig *SubscriptionItemDiscountSettingsServicePeriodAnchorConfigParams `form:"service_period_anchor_config"`
	// The start date of the discount's service period when applying a coupon or promotion code with a service period duration. Defaults to `now` if omitted.
	StartDate *string `form:"start_date"`
}

// The coupons to redeem into discounts for the subscription item.
type SubscriptionItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionItemDiscountDiscountEndParams `form:"discount_end"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
	// Settings for discount application including service period anchoring.
	Settings *SubscriptionItemDiscountSettingsParams `form:"settings"`
}

// The recurring components of a price such as `interval` and `interval_count`.
type SubscriptionItemPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
type SubscriptionItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to.
	Product *string `form:"product"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *SubscriptionItemPriceDataRecurringParams `form:"recurring"`
	// Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Returns a list of your subscription items for a given subscription.
type SubscriptionItemListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The ID of the subscription whose items will be retrieved.
	Subscription *string `form:"subscription"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionItemListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Options that configure the trial on the subscription item.
type SubscriptionItemTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial. Currently only supports at most 1 price ID.
	ConvertsTo []*string `form:"converts_to"`
	// Determines the type of trial for this item.
	Type *string `form:"type"`
}

// Deletes an item from the subscription. Removing a subscription item from a subscription will not cancel the subscription.
type SubscriptionItemDeleteParams struct {
	Params `form:"*"`
	// Delete all usage for the given subscription item. Allowed only when the current plan's `usage_type` is `metered`.
	ClearUsage *bool `form:"clear_usage"`
	// Determines how to handle [prorations](https://docs.stripe.com/billing/subscriptions/prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes. The default value is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior"`
	// If set, the proration will be calculated as though the subscription was updated at the given time. This can be used to apply the same proration that was previewed with the [upcoming invoice](https://api.stripe.com#retrieve_customer_invoice) endpoint.
	ProrationDate *int64 `form:"proration_date"`
}

// Retrieves the subscription item with the given ID.
type SubscriptionItemRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionItemRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionItemUpdateBillingThresholdsParams struct {
	// Number of units that meets the billing threshold to advance the subscription to a new billing period (e.g., it takes 10 $5 units to meet a $50 [monetary threshold](https://docs.stripe.com/api/subscriptions/update#update_subscription-billing_thresholds-amount_gte))
	UsageGTE *int64 `form:"usage_gte"`
}

// The trial offer to apply to this subscription item.
type SubscriptionItemUpdateCurrentTrialParams struct {
	// Unix timestamp representing the end of the trial offer period. Required when the trial offer has `duration.type=timestamp`. Cannot be specified when `duration.type=relative`.
	TrialEnd *int64 `form:"trial_end"`
	// The ID of the trial offer to apply to the subscription item.
	TrialOffer *string `form:"trial_offer"`
}

// Time span for the redeemed discount.
type SubscriptionItemUpdateDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionItemUpdateDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionItemUpdateDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// Anchor the service period to a custom date. Type must be `custom` to specify.
type SubscriptionItemUpdateDiscountSettingsServicePeriodAnchorConfigCustomParams struct {
	// The day of the month the anchor should be. Ranges from 1 to 31.
	DayOfMonth *int64 `form:"day_of_month"`
	// The hour of the day the anchor should be. Ranges from 0 to 23.
	Hour *int64 `form:"hour"`
	// The minute of the hour the anchor should be. Ranges from 0 to 59.
	Minute *int64 `form:"minute"`
	// The month to start full cycle periods. Ranges from 1 to 12.
	Month *int64 `form:"month"`
	// The second of the minute the anchor should be. Ranges from 0 to 59.
	Second *int64 `form:"second"`
}

// Configures service period cycle anchoring.
type SubscriptionItemUpdateDiscountSettingsServicePeriodAnchorConfigParams struct {
	// Anchor the service period to a custom date. Type must be `custom` to specify.
	Custom *SubscriptionItemUpdateDiscountSettingsServicePeriodAnchorConfigCustomParams `form:"custom"`
	// The type of service period anchor config. Defaults to `subscription_service_cycle_anchor` if omitted.
	Type *string `form:"type"`
}

// Settings for discount application including service period anchoring.
type SubscriptionItemUpdateDiscountSettingsParams struct {
	// Configures service period cycle anchoring.
	ServicePeriodAnchorConfig *SubscriptionItemUpdateDiscountSettingsServicePeriodAnchorConfigParams `form:"service_period_anchor_config"`
	// The start date of the discount's service period when applying a coupon or promotion code with a service period duration. Defaults to `now` if omitted.
	StartDate *string `form:"start_date"`
}

// The coupons to redeem into discounts for the subscription item.
type SubscriptionItemUpdateDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionItemUpdateDiscountDiscountEndParams `form:"discount_end"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
	// Settings for discount application including service period anchoring.
	Settings *SubscriptionItemUpdateDiscountSettingsParams `form:"settings"`
}

// The recurring components of a price such as `interval` and `interval_count`.
type SubscriptionItemUpdatePriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
type SubscriptionItemUpdatePriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to.
	Product *string `form:"product"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *SubscriptionItemUpdatePriceDataRecurringParams `form:"recurring"`
	// Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Updates the plan or quantity of an item on a current subscription.
type SubscriptionItemUpdateParams struct {
	Params `form:"*"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionItemUpdateBillingThresholdsParams `form:"billing_thresholds"`
	// The trial offer to apply to this subscription item.
	CurrentTrial *SubscriptionItemUpdateCurrentTrialParams `form:"current_trial"`
	// The coupons to redeem into discounts for the subscription item.
	Discounts []*SubscriptionItemUpdateDiscountParams `form:"discounts"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Indicates if a customer is on or off-session while an invoice payment is attempted. Defaults to `false` (on-session).
	OffSession *bool `form:"off_session"`
	// Use `allow_incomplete` to transition the subscription to `status=past_due` if a payment is required but cannot be paid. This allows you to manage scenarios where additional user actions are needed to pay a subscription's invoice. For example, SCA regulation may require 3DS authentication to complete payment. See the [SCA Migration Guide](https://docs.stripe.com/billing/migration/strong-customer-authentication) for Billing to learn more. This is the default behavior.
	//
	// Use `default_incomplete` to transition the subscription to `status=past_due` when payment is required and await explicit confirmation of the invoice's payment intent. This allows simpler management of scenarios where additional user actions are needed to pay a subscription's invoice. Such as failed payments, [SCA regulation](https://docs.stripe.com/billing/migration/strong-customer-authentication), or collecting a mandate for a bank debit payment method.
	//
	// Use `pending_if_incomplete` to update the subscription using [pending updates](https://docs.stripe.com/billing/subscriptions/pending-updates). When you use `pending_if_incomplete` you can only pass the parameters [supported by pending updates](https://docs.stripe.com/billing/pending-updates-reference#supported-attributes).
	//
	// Use `error_if_incomplete` if you want Stripe to return an HTTP 402 status code if a subscription's invoice cannot be paid. For example, if a payment method requires 3DS authentication due to SCA regulation and further user action is needed, this parameter does not update the subscription and returns an error instead. This was the default behavior for API versions prior to 2019-03-14. See the [changelog](https://docs.stripe.com/changelog/2019-03-14) to learn more.
	PaymentBehavior *string `form:"payment_behavior"`
	// The identifier of the new plan for this subscription item.
	Plan *string `form:"plan"`
	// The ID of the price object. One of `price` or `price_data` is required. When changing a subscription item's price, `quantity` is set to 1 unless a `quantity` parameter is provided.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *SubscriptionItemUpdatePriceDataParams `form:"price_data"`
	// Determines how to handle [prorations](https://docs.stripe.com/billing/subscriptions/prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes. The default value is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior"`
	// If set, the proration will be calculated as though the subscription was updated at the given time. This can be used to apply the same proration that was previewed with the [upcoming invoice](https://api.stripe.com#retrieve_customer_invoice) endpoint.
	ProrationDate *int64 `form:"proration_date"`
	// The quantity you'd like to apply to the subscription item you're creating.
	Quantity *int64 `form:"quantity"`
	// A list of [Tax Rate](https://docs.stripe.com/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://docs.stripe.com/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionItemUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionItemUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionItemCreateBillingThresholdsParams struct {
	// Number of units that meets the billing threshold to advance the subscription to a new billing period (e.g., it takes 10 $5 units to meet a $50 [monetary threshold](https://docs.stripe.com/api/subscriptions/update#update_subscription-billing_thresholds-amount_gte))
	UsageGTE *int64 `form:"usage_gte"`
}

// The trial offer to apply to this subscription item.
type SubscriptionItemCreateCurrentTrialParams struct {
	// Unix timestamp representing the end of the trial offer period. Required when the trial offer has `duration.type=timestamp`. Cannot be specified when `duration.type=relative`.
	TrialEnd *int64 `form:"trial_end"`
	// The ID of the trial offer to apply to the subscription item.
	TrialOffer *string `form:"trial_offer"`
}

// Time span for the redeemed discount.
type SubscriptionItemCreateDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionItemCreateDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionItemCreateDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// Anchor the service period to a custom date. Type must be `custom` to specify.
type SubscriptionItemCreateDiscountSettingsServicePeriodAnchorConfigCustomParams struct {
	// The day of the month the anchor should be. Ranges from 1 to 31.
	DayOfMonth *int64 `form:"day_of_month"`
	// The hour of the day the anchor should be. Ranges from 0 to 23.
	Hour *int64 `form:"hour"`
	// The minute of the hour the anchor should be. Ranges from 0 to 59.
	Minute *int64 `form:"minute"`
	// The month to start full cycle periods. Ranges from 1 to 12.
	Month *int64 `form:"month"`
	// The second of the minute the anchor should be. Ranges from 0 to 59.
	Second *int64 `form:"second"`
}

// Configures service period cycle anchoring.
type SubscriptionItemCreateDiscountSettingsServicePeriodAnchorConfigParams struct {
	// Anchor the service period to a custom date. Type must be `custom` to specify.
	Custom *SubscriptionItemCreateDiscountSettingsServicePeriodAnchorConfigCustomParams `form:"custom"`
	// The type of service period anchor config. Defaults to `subscription_service_cycle_anchor` if omitted.
	Type *string `form:"type"`
}

// Settings for discount application including service period anchoring.
type SubscriptionItemCreateDiscountSettingsParams struct {
	// Configures service period cycle anchoring.
	ServicePeriodAnchorConfig *SubscriptionItemCreateDiscountSettingsServicePeriodAnchorConfigParams `form:"service_period_anchor_config"`
	// The start date of the discount's service period when applying a coupon or promotion code with a service period duration. Defaults to `now` if omitted.
	StartDate *string `form:"start_date"`
}

// The coupons to redeem into discounts for the subscription item.
type SubscriptionItemCreateDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionItemCreateDiscountDiscountEndParams `form:"discount_end"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
	// Settings for discount application including service period anchoring.
	Settings *SubscriptionItemCreateDiscountSettingsParams `form:"settings"`
}

// The recurring components of a price such as `interval` and `interval_count`.
type SubscriptionItemCreatePriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline.
type SubscriptionItemCreatePriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to.
	Product *string `form:"product"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *SubscriptionItemCreatePriceDataRecurringParams `form:"recurring"`
	// Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Options that configure the trial on the subscription item.
type SubscriptionItemCreateTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial. Currently only supports at most 1 price ID.
	ConvertsTo []*string `form:"converts_to"`
	// Determines the type of trial for this item.
	Type *string `form:"type"`
}

// Adds a new item to an existing subscription. No existing items will be changed or replaced.
type SubscriptionItemCreateParams struct {
	Params `form:"*"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionItemCreateBillingThresholdsParams `form:"billing_thresholds"`
	// The trial offer to apply to this subscription item.
	CurrentTrial *SubscriptionItemCreateCurrentTrialParams `form:"current_trial"`
	// The coupons to redeem into discounts for the subscription item.
	Discounts []*SubscriptionItemCreateDiscountParams `form:"discounts"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Use `allow_incomplete` to transition the subscription to `status=past_due` if a payment is required but cannot be paid. This allows you to manage scenarios where additional user actions are needed to pay a subscription's invoice. For example, SCA regulation may require 3DS authentication to complete payment. See the [SCA Migration Guide](https://docs.stripe.com/billing/migration/strong-customer-authentication) for Billing to learn more. This is the default behavior.
	//
	// Use `default_incomplete` to transition the subscription to `status=past_due` when payment is required and await explicit confirmation of the invoice's payment intent. This allows simpler management of scenarios where additional user actions are needed to pay a subscription's invoice. Such as failed payments, [SCA regulation](https://docs.stripe.com/billing/migration/strong-customer-authentication), or collecting a mandate for a bank debit payment method.
	//
	// Use `pending_if_incomplete` to update the subscription using [pending updates](https://docs.stripe.com/billing/subscriptions/pending-updates). When you use `pending_if_incomplete` you can only pass the parameters [supported by pending updates](https://docs.stripe.com/billing/pending-updates-reference#supported-attributes).
	//
	// Use `error_if_incomplete` if you want Stripe to return an HTTP 402 status code if a subscription's invoice cannot be paid. For example, if a payment method requires 3DS authentication due to SCA regulation and further user action is needed, this parameter does not update the subscription and returns an error instead. This was the default behavior for API versions prior to 2019-03-14. See the [changelog](https://docs.stripe.com/changelog/2019-03-14) to learn more.
	PaymentBehavior *string `form:"payment_behavior"`
	// The identifier of the plan to add to the subscription.
	Plan *string `form:"plan"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline.
	PriceData *SubscriptionItemCreatePriceDataParams `form:"price_data"`
	// Determines how to handle [prorations](https://docs.stripe.com/billing/subscriptions/prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes. The default value is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior"`
	// If set, the proration will be calculated as though the subscription was updated at the given time. This can be used to apply the same proration that was previewed with the [upcoming invoice](https://api.stripe.com#retrieve_customer_invoice) endpoint.
	ProrationDate *int64 `form:"proration_date"`
	// The quantity you'd like to apply to the subscription item you're creating.
	Quantity *int64 `form:"quantity"`
	// The identifier of the subscription to modify.
	Subscription *string `form:"subscription"`
	// A list of [Tax Rate](https://docs.stripe.com/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://docs.stripe.com/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionItemCreateTrialParams `form:"trial"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionItemCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionItemCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
type SubscriptionItemBillingThresholds struct {
	// Usage threshold that triggers the subscription to create an invoice
	UsageGTE int64 `json:"usage_gte"`
}

// The current trial that is applied to this subscription item.
type SubscriptionItemCurrentTrial struct {
	EndDate    int64  `json:"end_date"`
	StartDate  int64  `json:"start_date"`
	TrialOffer string `json:"trial_offer"`
}

// Options that configure the trial on the subscription item.
type SubscriptionItemTrial struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []string `json:"converts_to"`
	// Determines the type of trial for this item.
	Type SubscriptionItemTrialType `json:"type"`
}

// Subscription items allow you to create customer subscriptions with more than
// one plan, making it easy to represent complex billing relationships.
type SubscriptionItem struct {
	APIResource
	// The time period the subscription item has been billed for.
	BilledUntil int64 `json:"billed_until"`
	// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
	BillingThresholds *SubscriptionItemBillingThresholds `json:"billing_thresholds"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The end time of this subscription item's current billing period.
	CurrentPeriodEnd int64 `json:"current_period_end"`
	// The start time of this subscription item's current billing period.
	CurrentPeriodStart int64 `json:"current_period_start"`
	// The current trial that is applied to this subscription item.
	CurrentTrial *SubscriptionItemCurrentTrial `json:"current_trial"`
	Deleted      bool                          `json:"deleted"`
	// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*Discount `json:"discounts"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// You can now model subscriptions more flexibly using the [Prices API](https://api.stripe.com#prices). It replaces the Plans API and is backwards compatible to simplify your migration.
	//
	// Plans define the base price, currency, and billing cycle for recurring purchases of products.
	// [Products](https://api.stripe.com#products) help you track inventory or provisioning, and plans help you track pricing. Different physical goods or levels of service should be represented by products, and pricing options should be represented by plans. This approach lets you change prices without having to change your provisioning scheme.
	//
	// For example, you might have a single "gold" product that has plans for $10/month, $100/year, €9/month, and €90/year.
	//
	// Related guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription) and more about [products and prices](https://docs.stripe.com/products-prices/overview).
	Plan *Plan `json:"plan"`
	// Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.
	// [Products](https://api.stripe.com#products) help you track inventory or provisioning, and prices help you track payment terms. Different physical goods or levels of service should be represented by products, and pricing options should be represented by prices. This approach lets you change prices without having to change your provisioning scheme.
	//
	// For example, you might have a single "gold" product that has prices for $10/month, $100/year, and €9 once.
	//
	// Related guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription), [create an invoice](https://docs.stripe.com/billing/invoices/create), and more about [products and prices](https://docs.stripe.com/products-prices/overview).
	Price *Price `json:"price"`
	// The [quantity](https://docs.stripe.com/subscriptions/quantities) of the plan to which the customer should be subscribed.
	Quantity int64 `json:"quantity"`
	// The `subscription` this `subscription_item` belongs to.
	Subscription string `json:"subscription"`
	// The tax rates which apply to this `subscription_item`. When set, the `default_tax_rates` on the subscription do not apply to this `subscription_item`.
	TaxRates []*TaxRate `json:"tax_rates"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionItemTrial `json:"trial"`
}

// SubscriptionItemList is a list of SubscriptionItems as retrieved from a list endpoint.
type SubscriptionItemList struct {
	APIResource
	ListMeta
	Data []*SubscriptionItem `json:"data"`
}
