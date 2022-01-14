//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Removes the currently applied discount on a customer.
type DiscountParams struct {
	Params `form:"*"`
}

// A discount represents the actual application of a coupon to a particular
// customer. It contains information about when the discount began and when it
// will end.
//
// Related guide: [Applying Discounts to Subscriptions](https://stripe.com/docs/billing/subscriptions/discounts).
type Discount struct {
	APIResource
	// The Checkout session that this coupon is applied to, if it is applied to a particular session in payment mode. Will not be present for subscription mode.
	CheckoutSession *CheckoutSession `json:"checkout_session"`
	// A coupon contains information about a percent-off or amount-off discount you
	// might want to apply to a customer. Coupons may be applied to [invoices](https://stripe.com/docs/api#invoices) or
	// [orders](https://stripe.com/docs/api#create_order_legacy-coupon). Coupons do not work with conventional one-off [charges](https://stripe.com/docs/api#create_charge).
	Coupon *Coupon `json:"coupon"`
	// The ID of the customer associated with this discount.
	Customer string `json:"customer"`
	Deleted  bool   `json:"deleted"`
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
	// Date that the coupon was applied.
	Start int64 `json:"start"`
	// The subscription that this coupon is applied to, if it is applied to a particular subscription.
	Subscription string `json:"subscription"`
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
