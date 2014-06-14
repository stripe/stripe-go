package stripe

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

type CouponDuration string

const (
	Forever   CouponDuration = "forever"
	Once      CouponDuration = "once"
	Repeating CouponDuration = "repeating"
)

type CouponParams struct {
	Duration                                     CouponDuration
	Id                                           string
	Currency                                     Currency
	Amount, Percent, DurationPeriod, Redemptions uint64
	Meta                                         map[string]string
	RedeemBy                                     int64
}

type CouponListParams struct {
	Filters    Filters
	Start, End string
	Limit      uint64
}

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

type CouponList struct {
	Count  uint16    `json:"total_count"`
	More   bool      `json:"has_more"`
	Url    string    `json:"url"`
	Values []*Coupon `json:"data"`
}

type CouponClient struct {
	api   Api
	token string
}

func (c *CouponClient) Create(params *CouponParams) (*Coupon, error) {
	body := &url.Values{
		"duration": {string(params.Duration)},
	}

	if len(params.Id) > 0 {
		body.Add("id", params.Id)
	}

	if params.Percent > 0 {
		body.Add("percent_off", strconv.FormatUint(params.Percent, 10))
	} else if params.Amount > 0 {
		body.Add("amount_off", strconv.FormatUint(params.Amount, 10))
		body.Add("currency", string(params.Currency))
	} else {
		err := errors.New("Invalid coupon params: either amount and currency or percent need to be set")
		return nil, err
	}

	if params.Duration == Repeating {
		body.Add("duration_in_months", strconv.FormatUint(params.DurationPeriod, 10))
	}

	if params.Redemptions > 0 {
		body.Add("max_redemptions", strconv.FormatUint(params.Redemptions, 10))
	}

	if params.RedeemBy > 0 {
		body.Add("redeem_by", strconv.FormatInt(params.RedeemBy, 10))
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	coupon := &Coupon{}
	err := c.api.Call("POST", "/coupons", c.token, body, coupon)

	return coupon, err
}

func (c *CouponClient) Get(id string) (*Coupon, error) {
	coupon := &Coupon{}
	err := c.api.Call("GET", "/coupons/"+id, c.token, nil, coupon)

	return coupon, err
}

func (c *CouponClient) Delete(id string) error {
	return c.api.Call("DELETE", "/coupons/"+id, c.token, nil, nil)
}

func (c *CouponClient) List(params *CouponListParams) (*CouponList, error) {
	body := &url.Values{}

	if params != nil {
		if len(params.Filters.f) > 0 {
			params.Filters.appendTo(body)
		}

		if len(params.Start) > 0 {
			body.Add("starting_after", params.Start)
		}

		if len(params.End) > 0 {
			body.Add("ending_before", params.End)
		}

		if params.Limit > 0 {
			if params.Limit > 100 {
				params.Limit = 100
			}

			body.Add("limit", strconv.FormatUint(params.Limit, 10))
		}
	}

	list := &CouponList{}
	err := c.api.Call("GET", "/coupons", c.token, body, list)

	return list, err
}
