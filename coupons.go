package stripe

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// CouponDuration is the list of allowed values for the coupon's duration.
// Allowed values are "forever", "once", "repeating".
type CouponDuration string

const (
	Forever   CouponDuration = "forever"
	Once      CouponDuration = "once"
	Repeating CouponDuration = "repeating"
)

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

// CouponList is a list object for coupons.
type CouponList struct {
	ListResponse
	Values []*Coupon `json:"data"`
}

// CouponClient is the client used to invoke /coupons APIs.
type CouponClient struct {
	api   Api
	token string
}

// Create POSTs new coupons.
// For more details see https://stripe.com/docs/api#create_coupon.
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

// Get returns the details of a coupon.
// For more details see https://stripe.com/docs/api#retrieve_coupon.
func (c *CouponClient) Get(id string) (*Coupon, error) {
	coupon := &Coupon{}
	err := c.api.Call("GET", "/coupons/"+id, c.token, nil, coupon)

	return coupon, err
}

// Delete removes a coupon.
// For more details see https://stripe.com/docs/api#delete_coupon.
func (c *CouponClient) Delete(id string) error {
	return c.api.Call("DELETE", "/coupons/"+id, c.token, nil, nil)
}

// List returns a list of coupons.
// For more details see https://stripe.com/docs/api#list_coupons.
func (c *CouponClient) List(params *CouponListParams) (*CouponList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

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
