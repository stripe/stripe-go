package stripe

import "encoding/json"

// CouponDuration is the list of allowed values for the coupon's duration.
// Allowed values are "forever", "once", "repeating".
type CouponDuration string

// CouponParams is the set of parameters that can be used when creating a coupon.
// For more details see https://stripe.com/docs/api#create_coupon.
type CouponParams struct {
	Params         `form:"*"`
	Duration       CouponDuration `form:"duration"`
	ID             string         `form:"id"`
	Currency       Currency       `form:"currency"`
	Amount         uint64         `form:"amount_off"`
	Percent        uint64         `form:"percent_off"`
	DurationPeriod uint64         `form:"duration_in_months"`
	Redemptions    uint64         `form:"max_redemptions"`
	RedeemBy       int64          `form:"redeem_by"`
}

// CouponListParams is the set of parameters that can be used when listing coupons.
// For more detail see https://stripe.com/docs/api#list_coupons.
type CouponListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// Coupon is the resource representing a Stripe coupon.
// For more details see https://stripe.com/docs/api#coupons.
type Coupon struct {
	ID             string            `json:"id"`
	Live           bool              `json:"livemode"`
	Created        int64             `json:"created"`
	Duration       CouponDuration    `json:"duration"`
	Amount         uint64            `json:"amount_off"`
	Currency       Currency          `json:"currency"`
	DurationPeriod uint64            `json:"duration_in_months"`
	Redemptions    uint64            `json:"max_redemptions"`
	Meta           map[string]string `json:"metadata"`
	Percent        uint64            `json:"percent_off"`
	RedeemBy       int64             `json:"redeem_by"`
	Redeemed       uint64            `json:"times_redeemed"`
	Valid          bool              `json:"valid"`
	Deleted        bool              `json:"deleted"`
}

// CouponList is a list of coupons as retrieved from a list endpoint.
type CouponList struct {
	ListMeta
	Values []*Coupon `json:"data"`
}

// UnmarshalJSON handles deserialization of a Coupon.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Coupon) UnmarshalJSON(data []byte) error {
	type coupon Coupon
	var cc coupon
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = Coupon(cc)
	} else {
		// the id is surrounded by "\" characters, so strip them
		c.ID = string(data[1 : len(data)-1])
	}

	return nil
}
