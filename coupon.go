package stripe

import "encoding/json"

// CouponDuration is the list of allowed values for the coupon's duration.
type CouponDuration string

// List of values that CouponDuration can take.
const (
	CouponDurationForever   CouponDuration = "forever"
	CouponDurationOnce      CouponDuration = "once"
	CouponDurationRepeating CouponDuration = "repeating"
)

// CouponAppliesToParams controls the products that a coupon applies to.
type CouponAppliesToParams struct {
	Products []*string `form:"products"`
}

// CouponParams is the set of parameters that can be used when creating a coupon.
// For more details see https://stripe.com/docs/api#create_coupon.
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

// CouponListParams is the set of parameters that can be used when listing coupons.
// For more detail see https://stripe.com/docs/api#list_coupons.
type CouponListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// CouponAppliesTo represents the products a coupon applies to.
type CouponAppliesTo struct {
	Products []string `json:"products"`
}

// Coupon is the resource representing a Stripe coupon.
// For more details see https://stripe.com/docs/api#coupons.
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

// CouponList is a list of coupons as retrieved from a list endpoint.
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
