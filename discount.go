package stripe

import "encoding/json"

// DiscountParams is the set of parameters that can be used when deleting a discount.
type DiscountParams struct {
	Params `form:"*"`
}

// Discount is the resource representing a Stripe discount.
// For more details see https://stripe.com/docs/api#discounts.
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
func (s *Discount) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type discount Discount
	var v discount
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = Discount(v)
	return nil
}
