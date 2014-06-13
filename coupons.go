package stripe

import (
	"encoding/json"
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

	res, err := c.api.Call("POST", "/coupons", c.token, body)

	if err != nil {
		return nil, err
	}

	var coupon Coupon
	err = json.Unmarshal(res, &coupon)
	return &coupon, err
}

func (c *CouponClient) Get(id string) (*Coupon, error) {
	res, err := c.api.Call("GET", "/coupons/"+id, c.token, nil)

	if err != nil {
		return nil, err
	}

	var coupon Coupon
	err = json.Unmarshal(res, &coupon)
	return &coupon, err
}

func (c *CouponClient) Delete(id string) error {
	_, err := c.api.Call("DELETE", "/coupons/"+id, c.token, nil)
	return err
}
