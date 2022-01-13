//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Retrieves the promotion code with the given ID. In order to retrieve a promotion code by the customer-facing code use [list](https://stripe.com/docs/api/promotion_codes/list) with the desired code.
type PromotionCodeParams struct {
	Params         `form:"*"`
	Active         *bool                            `form:"active"`
	Code           *string                          `form:"code"`
	Coupon         *string                          `form:"coupon"`
	Customer       *string                          `form:"customer"`
	ExpiresAt      *int64                           `form:"expires_at"`
	MaxRedemptions *int64                           `form:"max_redemptions"`
	Restrictions   *PromotionCodeRestrictionsParams `form:"restrictions"`
}

// Settings that restrict the redemption of the promotion code.
type PromotionCodeRestrictionsParams struct {
	FirstTimeTransaction  *bool   `form:"first_time_transaction"`
	MinimumAmount         *int64  `form:"minimum_amount"`
	MinimumAmountCurrency *string `form:"minimum_amount_currency"`
}

// Returns a list of your promotion codes.
type PromotionCodeListParams struct {
	ListParams   `form:"*"`
	Active       *bool             `form:"active"`
	Code         *string           `form:"code"`
	Coupon       *string           `form:"coupon"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Customer     *string           `form:"customer"`
}
type PromotionCodeRestrictions struct {
	FirstTimeTransaction  bool     `json:"first_time_transaction"`
	MinimumAmount         int64    `json:"minimum_amount"`
	MinimumAmountCurrency Currency `json:"minimum_amount_currency"`
}

// A Promotion Code represents a customer-redeemable code for a coupon. It can be used to
// create multiple codes for a single coupon.
type PromotionCode struct {
	APIResource
	Active         bool                       `json:"active"`
	Code           string                     `json:"code"`
	Coupon         *Coupon                    `json:"coupon"`
	Created        int64                      `json:"created"`
	Customer       *Customer                  `json:"customer"`
	ExpiresAt      int64                      `json:"expires_at"`
	ID             string                     `json:"id"`
	Livemode       bool                       `json:"livemode"`
	MaxRedemptions int64                      `json:"max_redemptions"`
	Metadata       map[string]string          `json:"metadata"`
	Object         string                     `json:"object"`
	Restrictions   *PromotionCodeRestrictions `json:"restrictions"`
	TimesRedeemed  int64                      `json:"times_redeemed"`
}

// PromotionCodeList is a list of PromotionCodes as retrieved from a list endpoint.
type PromotionCodeList struct {
	APIResource
	ListMeta
	Data []*PromotionCode `json:"data"`
}

// UnmarshalJSON handles deserialization of a PromotionCode.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PromotionCode) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type promotionCode PromotionCode
	var v promotionCode
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PromotionCode(v)
	return nil
}
