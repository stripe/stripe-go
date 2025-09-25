//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of parent that generated this invoice item
type InvoiceItemParentType string

// List of values that InvoiceItemParentType can take
const (
	InvoiceItemParentTypeLicenseFeeSubscriptionDetails InvoiceItemParentType = "license_fee_subscription_details"
	InvoiceItemParentTypeRateCardSubscriptionDetails   InvoiceItemParentType = "rate_card_subscription_details"
	InvoiceItemParentTypeSubscriptionDetails           InvoiceItemParentType = "subscription_details"
)

// The type of the pricing details.
type InvoiceItemPricingType string

// List of values that InvoiceItemPricingType can take
const (
	InvoiceItemPricingTypeLicenseFeeDetails   InvoiceItemPricingType = "license_fee_details"
	InvoiceItemPricingTypePriceDetails        InvoiceItemPricingType = "price_details"
	InvoiceItemPricingTypeRateCardRateDetails InvoiceItemPricingType = "rate_card_rate_details"
)

// Deletes an invoice item, removing it from an invoice. Deleting invoice items is only possible when they're not attached to invoices, or if it's attached to a draft invoice.
type InvoiceItemParams struct {
	Params `form:"*"`
	// The integer amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. If you want to apply a credit to the customer's account, pass a negative amount.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the customer who will be billed when this invoice item is billed.
	Customer *string `form:"customer"`
	// The ID of the account who will be billed when this invoice item is billed.
	CustomerAccount *string `form:"customer_account"`
	// An arbitrary string which you can attach to the invoice item. The description is displayed in the invoice for easy tracking.
	Description *string `form:"description"`
	// Controls whether discounts apply to this invoice item. Defaults to false for prorations or negative invoice items, and true for all other invoice items. Cannot be set to true for prorations.
	Discountable *bool `form:"discountable"`
	// The coupons, promotion codes & existing discounts which apply to the invoice item or invoice line item. Item discounts are applied before invoice discounts. Pass an empty string to remove previously-defined discounts.
	Discounts []*InvoiceItemDiscountParams `form:"discounts"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The ID of an existing invoice to add this invoice item to. For subscription invoices, when left blank, the invoice item will be added to the next upcoming scheduled invoice. For standalone invoices, the invoice item won't be automatically added unless you pass `pending_invoice_item_behavior: 'include'` when creating the invoice. This is useful when adding invoice items in response to an invoice.created webhook. You can only add invoice items to draft invoices and there is a maximum of 250 items per invoice.
	Invoice *string `form:"invoice"`
	// The ids of the margins to apply to the invoice item. When set, the `default_margins` on the invoice do not apply to this invoice item.
	Margins []*string `form:"margins"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The period associated with this invoice item. When set to different values, the period will be rendered on the invoice. If you have [Stripe Revenue Recognition](https://stripe.com/docs/revenue-recognition) enabled, the period will be used to recognize and defer revenue. See the [Revenue Recognition documentation](https://stripe.com/docs/revenue-recognition/methodology/subscriptions-and-invoicing) for details.
	Period *InvoiceItemPeriodParams `form:"period"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	// The pricing information for the invoice item.
	Pricing *InvoiceItemPricingParams `form:"pricing"`
	// Non-negative integer. The quantity of units for the invoice item.
	Quantity *int64 `form:"quantity"`
	// The ID of a subscription to add this invoice item to. When left blank, the invoice item is added to the next upcoming scheduled invoice. When set, scheduled invoices for subscriptions other than the specified subscription will ignore the invoice item. Use this when you want to express that an invoice item has been accrued within the context of a particular subscription.
	Subscription *string `form:"subscription"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
	// The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item. Pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
	// The decimal unit amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. This `unit_amount_decimal` will be multiplied by the quantity to get the full amount. Passing in a negative `unit_amount_decimal` will reduce the `amount_due` on the invoice. Accepts at most 12 decimal places.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceItemParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *InvoiceItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Time span for the redeemed discount.
type InvoiceItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type InvoiceItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *InvoiceItemDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The coupons, promotion codes & existing discounts which apply to the invoice item or invoice line item. Item discounts are applied before invoice discounts. Pass an empty string to remove previously-defined discounts.
type InvoiceItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *InvoiceItemDiscountDiscountEndParams `form:"discount_end"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// The period associated with this invoice item. When set to different values, the period will be rendered on the invoice. If you have [Stripe Revenue Recognition](https://stripe.com/docs/revenue-recognition) enabled, the period will be used to recognize and defer revenue. See the [Revenue Recognition documentation](https://stripe.com/docs/revenue-recognition/methodology/subscriptions-and-invoicing) for details.
type InvoiceItemPeriodParams struct {
	// The end of the period, which must be greater than or equal to the start. This value is inclusive.
	End *int64 `form:"end"`
	// The start of the period. This value is inclusive.
	Start *int64 `form:"start"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
type InvoiceItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to.
	Product *string `form:"product"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// The pricing information for the invoice item.
type InvoiceItemPricingParams struct {
	// The ID of the price object.
	Price *string `form:"price"`
}

// Returns a list of your invoice items. Invoice items are returned sorted by creation date, with the most recently created invoice items appearing first.
type InvoiceItemListParams struct {
	ListParams `form:"*"`
	// Only return invoice items that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return invoice items that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// The identifier of the customer whose invoice items to return. If none is provided, all invoice items will be returned.
	Customer *string `form:"customer"`
	// The identifier of the account whose invoice items to return. If none is provided, all invoice items will be returned.
	CustomerAccount *string `form:"customer_account"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return invoice items belonging to this invoice. If none is provided, all invoice items will be returned. If specifying an invoice, no customer identifier is needed.
	Invoice *string `form:"invoice"`
	// Set to `true` to only show pending invoice items, which are not yet attached to any invoices. Set to `false` to only show invoice items already attached to invoices. If unspecified, no filter is applied.
	Pending *bool `form:"pending"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceItemListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Deletes an invoice item, removing it from an invoice. Deleting invoice items is only possible when they're not attached to invoices, or if it's attached to a draft invoice.
type InvoiceItemDeleteParams struct {
	Params `form:"*"`
}

// Retrieves the invoice item with the given ID.
type InvoiceItemRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceItemRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Time span for the redeemed discount.
type InvoiceItemUpdateDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type InvoiceItemUpdateDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *InvoiceItemUpdateDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The coupons, promotion codes & existing discounts which apply to the invoice item or invoice line item. Item discounts are applied before invoice discounts. Pass an empty string to remove previously-defined discounts.
type InvoiceItemUpdateDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *InvoiceItemUpdateDiscountDiscountEndParams `form:"discount_end"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// The period associated with this invoice item. When set to different values, the period will be rendered on the invoice. If you have [Stripe Revenue Recognition](https://stripe.com/docs/revenue-recognition) enabled, the period will be used to recognize and defer revenue. See the [Revenue Recognition documentation](https://stripe.com/docs/revenue-recognition/methodology/subscriptions-and-invoicing) for details.
type InvoiceItemUpdatePeriodParams struct {
	// The end of the period, which must be greater than or equal to the start. This value is inclusive.
	End *int64 `form:"end"`
	// The start of the period. This value is inclusive.
	Start *int64 `form:"start"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
type InvoiceItemUpdatePriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to.
	Product *string `form:"product"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// The pricing information for the invoice item.
type InvoiceItemUpdatePricingParams struct {
	// The ID of the price object.
	Price *string `form:"price"`
}

// Updates the amount or description of an invoice item on an upcoming invoice. Updating an invoice item is only possible before the invoice it's attached to is closed.
type InvoiceItemUpdateParams struct {
	Params `form:"*"`
	// The integer amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. If you want to apply a credit to the customer's account, pass a negative amount.
	Amount *int64 `form:"amount"`
	// An arbitrary string which you can attach to the invoice item. The description is displayed in the invoice for easy tracking.
	Description *string `form:"description"`
	// Controls whether discounts apply to this invoice item. Defaults to false for prorations or negative invoice items, and true for all other invoice items. Cannot be set to true for prorations.
	Discountable *bool `form:"discountable"`
	// The coupons, promotion codes & existing discounts which apply to the invoice item or invoice line item. Item discounts are applied before invoice discounts. Pass an empty string to remove previously-defined discounts.
	Discounts []*InvoiceItemUpdateDiscountParams `form:"discounts"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The ids of the margins to apply to the invoice item. When set, the `default_margins` on the invoice do not apply to this invoice item.
	Margins []*string `form:"margins"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The period associated with this invoice item. When set to different values, the period will be rendered on the invoice. If you have [Stripe Revenue Recognition](https://stripe.com/docs/revenue-recognition) enabled, the period will be used to recognize and defer revenue. See the [Revenue Recognition documentation](https://stripe.com/docs/revenue-recognition/methodology/subscriptions-and-invoicing) for details.
	Period *InvoiceItemUpdatePeriodParams `form:"period"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *InvoiceItemUpdatePriceDataParams `form:"price_data"`
	// The pricing information for the invoice item.
	Pricing *InvoiceItemUpdatePricingParams `form:"pricing"`
	// Non-negative integer. The quantity of units for the invoice item.
	Quantity *int64 `form:"quantity"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
	// The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item. Pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
	// The decimal unit amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. This `unit_amount_decimal` will be multiplied by the quantity to get the full amount. Passing in a negative `unit_amount_decimal` will reduce the `amount_due` on the invoice. Accepts at most 12 decimal places.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceItemUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *InvoiceItemUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Time span for the redeemed discount.
type InvoiceItemCreateDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type InvoiceItemCreateDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *InvoiceItemCreateDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The coupons and promotion codes to redeem into discounts for the invoice item or invoice line item.
type InvoiceItemCreateDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *InvoiceItemCreateDiscountDiscountEndParams `form:"discount_end"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// The period associated with this invoice item. When set to different values, the period will be rendered on the invoice. If you have [Stripe Revenue Recognition](https://stripe.com/docs/revenue-recognition) enabled, the period will be used to recognize and defer revenue. See the [Revenue Recognition documentation](https://stripe.com/docs/revenue-recognition/methodology/subscriptions-and-invoicing) for details.
type InvoiceItemCreatePeriodParams struct {
	// The end of the period, which must be greater than or equal to the start. This value is inclusive.
	End *int64 `form:"end"`
	// The start of the period. This value is inclusive.
	Start *int64 `form:"start"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
type InvoiceItemCreatePriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to.
	Product *string `form:"product"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// The pricing information for the invoice item.
type InvoiceItemCreatePricingParams struct {
	// The ID of the price object.
	Price *string `form:"price"`
}

// Creates an item to be added to a draft invoice (up to 250 items per invoice). If no invoice is specified, the item will be on the next invoice created for the customer specified.
type InvoiceItemCreateParams struct {
	Params `form:"*"`
	// The integer amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. Passing in a negative `amount` will reduce the `amount_due` on the invoice.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the customer who will be billed when this invoice item is billed.
	Customer *string `form:"customer"`
	// The ID of the account who will be billed when this invoice item is billed.
	CustomerAccount *string `form:"customer_account"`
	// An arbitrary string which you can attach to the invoice item. The description is displayed in the invoice for easy tracking.
	Description *string `form:"description"`
	// Controls whether discounts apply to this invoice item. Defaults to false for prorations or negative invoice items, and true for all other invoice items.
	Discountable *bool `form:"discountable"`
	// The coupons and promotion codes to redeem into discounts for the invoice item or invoice line item.
	Discounts []*InvoiceItemCreateDiscountParams `form:"discounts"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The ID of an existing invoice to add this invoice item to. For subscription invoices, when left blank, the invoice item will be added to the next upcoming scheduled invoice. For standalone invoices, the invoice item won't be automatically added unless you pass `pending_invoice_item_behavior: 'include'` when creating the invoice. This is useful when adding invoice items in response to an invoice.created webhook. You can only add invoice items to draft invoices and there is a maximum of 250 items per invoice.
	Invoice *string `form:"invoice"`
	// The ids of the margins to apply to the invoice item. When set, the `default_margins` on the invoice do not apply to this invoice item.
	Margins []*string `form:"margins"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The period associated with this invoice item. When set to different values, the period will be rendered on the invoice. If you have [Stripe Revenue Recognition](https://stripe.com/docs/revenue-recognition) enabled, the period will be used to recognize and defer revenue. See the [Revenue Recognition documentation](https://stripe.com/docs/revenue-recognition/methodology/subscriptions-and-invoicing) for details.
	Period *InvoiceItemCreatePeriodParams `form:"period"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *InvoiceItemCreatePriceDataParams `form:"price_data"`
	// The pricing information for the invoice item.
	Pricing *InvoiceItemCreatePricingParams `form:"pricing"`
	// Non-negative integer. The quantity of units for the invoice item.
	Quantity *int64 `form:"quantity"`
	// The ID of a subscription to add this invoice item to. When left blank, the invoice item is added to the next upcoming scheduled invoice. When set, scheduled invoices for subscriptions other than the specified subscription will ignore the invoice item. Use this when you want to express that an invoice item has been accrued within the context of a particular subscription.
	Subscription *string `form:"subscription"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
	// The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item.
	TaxRates []*string `form:"tax_rates"`
	// The decimal unit amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. This `unit_amount_decimal` will be multiplied by the quantity to get the full amount. Passing in a negative `unit_amount_decimal` will reduce the `amount_due` on the invoice. Accepts at most 12 decimal places.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// AddExpand appends a new field to expand.
func (p *InvoiceItemCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *InvoiceItemCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details about the license fee subscription that generated this invoice item
type InvoiceItemParentLicenseFeeSubscriptionDetails struct {
	// The license fee subscription that generated this invoice item
	LicenseFeeSubscription string `json:"license_fee_subscription"`
	// The license fee version that generated this invoice item
	LicenseFeeVersion string `json:"license_fee_version"`
	// The pricing plan subscription that manages the license fee subscription
	PricingPlanSubscription string `json:"pricing_plan_subscription"`
	// The pricing plan version at the time this invoice item was generated
	PricingPlanVersion string `json:"pricing_plan_version"`
}

// Details about the rate card subscription that generated this invoice item
type InvoiceItemParentRateCardSubscriptionDetails struct {
	// The pricing plan subscription that manages the rate card subscription
	PricingPlanSubscription string `json:"pricing_plan_subscription"`
	// The pricing plan version at the time this invoice item was generated
	PricingPlanVersion string `json:"pricing_plan_version"`
	// The rate card subscription that generated this invoice item
	RateCardSubscription string `json:"rate_card_subscription"`
	// The rate card version that generated this invoice item
	RateCardVersion string `json:"rate_card_version"`
}

// Details about the subscription that generated this invoice item
type InvoiceItemParentSubscriptionDetails struct {
	// The subscription that generated this invoice item
	Subscription string `json:"subscription"`
	// The subscription item that generated this invoice item
	SubscriptionItem string `json:"subscription_item"`
}

// The parent that generated this invoice item.
type InvoiceItemParent struct {
	// Details about the license fee subscription that generated this invoice item
	LicenseFeeSubscriptionDetails *InvoiceItemParentLicenseFeeSubscriptionDetails `json:"license_fee_subscription_details"`
	// Details about the rate card subscription that generated this invoice item
	RateCardSubscriptionDetails *InvoiceItemParentRateCardSubscriptionDetails `json:"rate_card_subscription_details"`
	// Details about the subscription that generated this invoice item
	SubscriptionDetails *InvoiceItemParentSubscriptionDetails `json:"subscription_details"`
	// The type of parent that generated this invoice item
	Type InvoiceItemParentType `json:"type"`
}
type InvoiceItemPricingLicenseFeeDetails struct {
	// The ID of the licensed item this item is associated with
	LicensedItem string `json:"licensed_item"`
	// The ID of the license fee this item is associated with
	LicenseFee string `json:"license_fee"`
	// The version of the license fee this item is associated with
	LicenseFeeVersion string `json:"license_fee_version"`
}
type InvoiceItemPricingPriceDetails struct {
	// The ID of the price this item is associated with.
	Price string `json:"price"`
	// The ID of the product this item is associated with.
	Product string `json:"product"`
}
type InvoiceItemPricingRateCardRateDetails struct {
	// The ID of billable item this item is associated with
	MeteredItem string `json:"metered_item"`
	// The ID of the rate card this item is associated with
	RateCard string `json:"rate_card"`
	// The ID of the rate card rate this item is associated with
	RateCardRate string `json:"rate_card_rate"`
}

// The pricing information of the invoice item.
type InvoiceItemPricing struct {
	LicenseFeeDetails   *InvoiceItemPricingLicenseFeeDetails   `json:"license_fee_details"`
	PriceDetails        *InvoiceItemPricingPriceDetails        `json:"price_details"`
	RateCardRateDetails *InvoiceItemPricingRateCardRateDetails `json:"rate_card_rate_details"`
	// The type of the pricing details.
	Type InvoiceItemPricingType `json:"type"`
	// The unit amount (in the `currency` specified) of the item which contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// Discount amounts applied when the proration was created.
type InvoiceItemProrationDetailsDiscountAmount struct {
	// The amount, in cents (or local equivalent), of the discount.
	Amount int64 `json:"amount"`
	// The discount that was applied to get this discount amount.
	Discount *Discount `json:"discount"`
}
type InvoiceItemProrationDetails struct {
	// Discount amounts applied when the proration was created.
	DiscountAmounts []*InvoiceItemProrationDetailsDiscountAmount `json:"discount_amounts"`
}

// Invoice Items represent the component lines of an [invoice](https://stripe.com/docs/api/invoices). When you create an invoice item with an `invoice` field, it is attached to the specified invoice and included as [an invoice line item](https://stripe.com/docs/api/invoices/line_item) within [invoice.lines](https://stripe.com/docs/api/invoices/object#invoice_object-lines).
//
// Invoice Items can be created before you are ready to actually send the invoice. This can be particularly useful when combined
// with a [subscription](https://stripe.com/docs/api/subscriptions). Sometimes you want to add a charge or credit to a customer, but actually charge
// or credit the customer's card only at the end of a regular billing cycle. This is useful for combining several charges
// (to minimize per-transaction fees), or for having Stripe tabulate your usage-based billing totals.
//
// Related guides: [Integrate with the Invoicing API](https://stripe.com/docs/invoicing/integration), [Subscription Invoices](https://stripe.com/docs/billing/invoices/subscription#adding-upcoming-invoice-items).
type InvoiceItem struct {
	APIResource
	// Amount (in the `currency` specified) of the invoice item. This should always be equal to `unit_amount * quantity`.
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The ID of the customer who will be billed when this invoice item is billed.
	Customer *Customer `json:"customer"`
	// The ID of the account who will be billed when this invoice item is billed.
	CustomerAccount string `json:"customer_account"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Date    int64 `json:"date"`
	Deleted bool  `json:"deleted"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// If true, discounts will apply to this invoice item. Always false for prorations.
	Discountable bool `json:"discountable"`
	// The discounts which apply to the invoice item. Item discounts are applied before invoice discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*Discount `json:"discounts"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The ID of the invoice this invoice item belongs to.
	Invoice *Invoice `json:"invoice"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The margins which apply to the invoice item. When set, the `default_margins` on the invoice do not apply to this invoice item.
	Margins []*Margin `json:"margins"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The amount after discounts, but before credits and taxes. This field is `null` for `discountable=true` items.
	NetAmount int64 `json:"net_amount"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The parent that generated this invoice item.
	Parent *InvoiceItemParent `json:"parent"`
	Period *Period            `json:"period"`
	// The pricing information of the invoice item.
	Pricing *InvoiceItemPricing `json:"pricing"`
	// Whether the invoice item was created automatically as a proration adjustment when the customer switched plans.
	Proration        bool                         `json:"proration"`
	ProrationDetails *InvoiceItemProrationDetails `json:"proration_details"`
	// Quantity of units for the invoice item. If the invoice item is a proration, the quantity of the subscription that the proration was computed for.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item.
	TaxRates []*TaxRate `json:"tax_rates"`
	// ID of the test clock this invoice item belongs to.
	TestClock *TestHelpersTestClock `json:"test_clock"`
}

// InvoiceItemList is a list of InvoiceItems as retrieved from a list endpoint.
type InvoiceItemList struct {
	APIResource
	ListMeta
	Data []*InvoiceItem `json:"data"`
}
