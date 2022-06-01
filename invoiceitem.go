//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Returns a list of your invoice items. Invoice items are returned sorted by creation date, with the most recently created invoice items appearing first.
type InvoiceItemListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	// The identifier of the customer whose invoice items to return. If none is provided, all invoice items will be returned.
	Customer *string `form:"customer"`
	// Only return invoice items belonging to this invoice. If none is provided, all invoice items will be returned. If specifying an invoice, no customer identifier is needed.
	Invoice *string `form:"invoice"`
	// Set to `true` to only show pending invoice items, which are not yet attached to any invoices. Set to `false` to only show invoice items already attached to invoices. If unspecified, no filter is applied.
	Pending *bool `form:"pending"`
}

// The coupons to redeem into discounts for the invoice item or invoice line item.
type InvoiceItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
}

// The period associated with this invoice item. When set to different values, the period will be rendered on the invoice.
type InvoiceItemPeriodParams struct {
	// The end of the period, which must be greater than or equal to the start.
	End *int64 `form:"end"`
	// The start of the period.
	Start *int64 `form:"start"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
type InvoiceItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the product that this price will belong to.
	Product *string `form:"product"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Creates an item to be added to a draft invoice (up to 250 items per invoice). If no invoice is specified, the item will be on the next invoice created for the customer specified.
type InvoiceItemParams struct {
	Params `form:"*"`
	// The integer amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. If you want to apply a credit to the customer's account, pass a negative amount.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the customer who will be billed when this invoice item is billed.
	Customer *string `form:"customer"`
	// An arbitrary string which you can attach to the invoice item. The description is displayed in the invoice for easy tracking.
	Description *string `form:"description"`
	// Controls whether discounts apply to this invoice item. Defaults to false for prorations or negative invoice items, and true for all other invoice items. Cannot be set to true for prorations.
	Discountable *bool `form:"discountable"`
	// The coupons & existing discounts which apply to the invoice item or invoice line item. Item discounts are applied before invoice discounts. Pass an empty string to remove previously-defined discounts.
	Discounts []*InvoiceItemDiscountParams `form:"discounts"`
	// The ID of an existing invoice to add this invoice item to. When left blank, the invoice item will be added to the next upcoming scheduled invoice. This is useful when adding invoice items in response to an invoice.created webhook. You can only add invoice items to draft invoices and there is a maximum of 250 items per invoice.
	Invoice *string `form:"invoice"`
	// The period associated with this invoice item. When set to different values, the period will be rendered on the invoice.
	Period *InvoiceItemPeriodParams `form:"period"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	// Non-negative integer. The quantity of units for the invoice item.
	Quantity *int64 `form:"quantity"`
	// The ID of a subscription to add this invoice item to. When left blank, the invoice item will be be added to the next upcoming scheduled invoice. When set, scheduled invoices for subscriptions other than the specified subscription will ignore the invoice item. Use this when you want to express that an invoice item has been accrued within the context of a particular subscription.
	Subscription *string `form:"subscription"`
	// The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item. Pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
	// The integer unit amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. This unit_amount will be multiplied by the quantity to get the full amount. If you want to apply a credit to the customer's account, pass a negative unit_amount.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Sometimes you want to add a charge or credit to a customer, but actually
// charge or credit the customer's card only at the end of a regular billing
// cycle. This is useful for combining several charges (to minimize
// per-transaction fees), or for having Stripe tabulate your usage-based billing
// totals.
//
// Related guide: [Subscription Invoices](https://stripe.com/docs/billing/invoices/subscription#adding-upcoming-invoice-items).
type InvoiceItem struct {
	APIResource
	// Amount (in the `currency` specified) of the invoice item. This should always be equal to `unit_amount * quantity`.
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The ID of the customer who will be billed when this invoice item is billed.
	Customer *Customer `json:"customer"`
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
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string  `json:"object"`
	Period *Period `json:"period"`
	// If the invoice item is a proration, the plan of the subscription that the proration was computed for.
	Plan *Plan `json:"plan"`
	// The price of the invoice item.
	Price *Price `json:"price"`
	// Whether the invoice item was created automatically as a proration adjustment when the customer switched plans.
	Proration bool `json:"proration"`
	// Quantity of units for the invoice item. If the invoice item is a proration, the quantity of the subscription that the proration was computed for.
	Quantity int64 `json:"quantity"`
	// The subscription that this invoice item has been created for, if any.
	Subscription *Subscription `json:"subscription"`
	// The subscription item that this invoice item has been created for, if any.
	SubscriptionItem string `json:"subscription_item"`
	// The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item.
	TaxRates []*TaxRate `json:"tax_rates"`
	// ID of the test clock this invoice item belongs to.
	TestClock *TestHelpersTestClock `json:"test_clock"`
	// Unit amount (in the `currency` specified) of the invoice item.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// UnmarshalJSON handles deserialization of an InvoiceItem.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *InvoiceItem) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type invoiceItem InvoiceItem
	var v invoiceItem
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = InvoiceItem(v)
	return nil
}

// InvoiceItemList is a list of InvoiceItems as retrieved from a list endpoint.
type InvoiceItemList struct {
	APIResource
	ListMeta
	Data []*InvoiceItem `json:"data"`
}
