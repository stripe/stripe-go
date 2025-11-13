//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The source type of the discount.
type DiscountSourceType string

// List of values that DiscountSourceType can take
const (
	DiscountSourceTypeCoupon DiscountSourceType = "coupon"
)

type DiscountSource struct {
	// The coupon that was redeemed to create this discount.
	Coupon *Coupon `json:"coupon"`
	// The source type of the discount.
	Type DiscountSourceType `json:"type"`
}

// A discount represents the actual application of a [coupon](https://stripe.com/docs/api#coupons) or [promotion code](https://stripe.com/docs/api#promotion_codes).
// It contains information about when the discount began, when it will end, and what it is applied to.
//
// Related guide: [Applying discounts to subscriptions](https://stripe.com/docs/billing/subscriptions/discounts)
type Discount struct {
	// The Checkout session that this coupon is applied to, if it is applied to a particular session in payment mode. Will not be present for subscription mode.
	CheckoutSession string `json:"checkout_session"`
	// The ID of the customer associated with this discount.
	Customer *Customer `json:"customer"`
	// The ID of the account associated with this discount.
	CustomerAccount string `json:"customer_account"`
	Deleted         bool   `json:"deleted"`
	// If the coupon has a duration of `repeating`, the date that this discount will end. If the coupon has a duration of `once` or `forever`, this attribute will be null.
	End int64 `json:"end"`
	// The ID of the discount object. Discounts cannot be fetched by ID. Use `expand[]=discounts` in API calls to expand discount IDs in an array.
	ID string `json:"id"`
	// The invoice that the discount's coupon was applied to, if it was applied directly to a particular invoice.
	Invoice string `json:"invoice"`
	// The invoice item `id` (or invoice line item `id` for invoice line items of type='subscription') that the discount's coupon was applied to, if it was applied directly to a particular invoice item or invoice line item.
	InvoiceItem string `json:"invoice_item"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The promotion code applied to create this discount.
	PromotionCode *PromotionCode `json:"promotion_code"`
	// The subscription schedule that this coupon is applied to, if it is applied to a particular subscription schedule.
	Schedule string          `json:"schedule"`
	Source   *DiscountSource `json:"source"`
	// Date that the coupon was applied.
	Start int64 `json:"start"`
	// The subscription that this coupon is applied to, if it is applied to a particular subscription.
	Subscription string `json:"subscription"`
	// The subscription item that this coupon is applied to, if it is applied to a particular subscription item.
	SubscriptionItem string `json:"subscription_item"`
}

// UnmarshalJSON handles deserialization of a Discount.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (d *Discount) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		d.ID = id
		return nil
	}

	type discount Discount
	var v discount
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*d = Discount(v)
	return nil
}
