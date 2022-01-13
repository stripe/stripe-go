//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// One of `forever`, `once`, and `repeating`. Describes how long a customer who applies this coupon will get the discount.
type CouponDuration string

// List of values that CouponDuration can take
const (
	CouponDurationForever   CouponDuration = "forever"
	CouponDurationOnce      CouponDuration = "once"
	CouponDurationRepeating CouponDuration = "repeating"
)

// Returns a list of your coupons.
type CouponListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// A hash containing directions for what this Coupon will apply discounts to.
type CouponAppliesToParams struct {
	Products []*string `form:"products"`
}

// You can create coupons easily via the [coupon management](https://dashboard.stripe.com/coupons) page of the Stripe dashboard. Coupon creation is also accessible via the API if you need to create coupons on the fly.
//
// A coupon has either a percent_off or an amount_off and currency. If you set an amount_off, that amount will be subtracted from any invoice's subtotal. For example, an invoice with a subtotal of 100 will have a final total of 0 if a coupon with an amount_off of 200 is applied to it and an invoice with a subtotal of 300 will have a final total of 100 if a coupon with an amount_off of 200 is applied to it.
type CouponParams struct {
	Params           `form:"*"`
	AmountOff        *int64                 `form:"amount_off"`
	AppliesTo        *CouponAppliesToParams `form:"applies_to"`
	Currency         *string                `form:"currency"`
	Duration         *string                `form:"duration"`
	DurationInMonths *int64                 `form:"duration_in_months"`
	ID               *string                `form:"id"`
	MaxRedemptions   *int64                 `form:"max_redemptions"`
	Name             *string                `form:"name"`
	PercentOff       *float64               `form:"percent_off"`
	RedeemBy         *int64                 `form:"redeem_by"`
}
type CouponAppliesTo struct {
	Products []string `json:"products"`
}

// A coupon contains information about a percent-off or amount-off discount you
// might want to apply to a customer. Coupons may be applied to [invoices](https://stripe.com/docs/api#invoices) or
// [orders](https://stripe.com/docs/api#create_order_legacy-coupon). Coupons do not work with conventional one-off [charges](https://stripe.com/docs/api#create_charge).
type Coupon struct {
	APIResource
	AmountOff        int64             `json:"amount_off"`
	AppliesTo        *CouponAppliesTo  `json:"applies_to"`
	Created          int64             `json:"created"`
	Currency         Currency          `json:"currency"`
	Deleted          bool              `json:"deleted"`
	Duration         CouponDuration    `json:"duration"`
	DurationInMonths int64             `json:"duration_in_months"`
	ID               string            `json:"id"`
	Livemode         bool              `json:"livemode"`
	MaxRedemptions   int64             `json:"max_redemptions"`
	Metadata         map[string]string `json:"metadata"`
	Name             string            `json:"name"`
	Object           string            `json:"object"`
	PercentOff       float64           `json:"percent_off"`
	RedeemBy         int64             `json:"redeem_by"`
	TimesRedeemed    int64             `json:"times_redeemed"`
	Valid            bool              `json:"valid"`
}

// CouponList is a list of Coupons as retrieved from a list endpoint.
type CouponList struct {
	APIResource
	ListMeta
	Data []*Coupon `json:"data"`
}

// UnmarshalJSON handles deserialization of a Coupon.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Coupon) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type coupon Coupon
	var v coupon
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Coupon(v)
	return nil
}
