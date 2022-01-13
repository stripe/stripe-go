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
	CheckoutSession *CheckoutSession `json:"checkout_session"`
	Coupon          *Coupon          `json:"coupon"`
	Customer        string           `json:"customer"`
	Deleted         bool             `json:"deleted"`
	End             int64            `json:"end"`
	ID              string           `json:"id"`
	Invoice         string           `json:"invoice"`
	InvoiceItem     string           `json:"invoice_item"`
	Object          string           `json:"object"`
	PromotionCode   *PromotionCode   `json:"promotion_code"`
	Start           int64            `json:"start"`
	Subscription    string           `json:"subscription"`
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
