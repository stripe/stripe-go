package stripe

import "encoding/json"

// CouponDuration is the list of allowed values for the coupon's duration.
// Allowed values are "forever", "once", "repeating".
type CouponDuration string

// CouponParams is the set of parameters that can be used when creating a coupon.
// For more details see https://stripe.com/docs/api#create_coupon.
type CouponParams struct {
	Params
	Duration                                     CouponDuration
	Id                                           string
	Currency                                     Currency
	Amount, Percent, DurationPeriod, Redemptions uint64
	RedeemBy                                     int64
}

// CouponListParams is the set of parameters that can be used when listing coupons.
// For more detail see https://stripe.com/docs/api#list_coupons.
type CouponListParams struct {
	ListParams
}

// Coupon is the resource representing a Stripe coupon.
// For more details see https://stripe.com/docs/api#coupons.
type Coupon struct {
	Id             string            `json:"id"`
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
}

func (c *Coupon) UnmarshalJSON(data []byte) error {
	type coupon Coupon
	var cc coupon
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = Coupon(cc)
	} else {
		// the id is surrounded by escaped \, so ignore those
		c.Id = string(data[1 : len(data)-1])
	}

	return nil
}
