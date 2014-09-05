package coupon

import (
	"errors"
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /coupons APIs.
type Client struct {
	B     Backend
	Token string
}

var c *Client

// Create POSTs new coupons.
// For more details see https://stripe.com/docs/api#create_coupon.
func Create(params *CouponParams) (*Coupon, error) {
	refresh()
	return c.Create(params)
}

func (c *Client) Create(params *CouponParams) (*Coupon, error) {
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

	params.AppendTo(body)

	coupon := &Coupon{}
	err := c.B.Call("POST", "/coupons", c.Token, body, coupon)

	return coupon, err
}

// Get returns the details of a coupon.
// For more details see https://stripe.com/docs/api#retrieve_coupon.
func Get(id string, params *CouponParams) (*Coupon, error) {
	refresh()
	return c.Get(id, params)
}

func (c *Client) Get(id string, params *CouponParams) (*Coupon, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	coupon := &Coupon{}
	err := c.B.Call("GET", "/coupons/"+id, c.Token, body, coupon)

	return coupon, err
}

// Delete removes a coupon.
// For more details see https://stripe.com/docs/api#delete_coupon.
func Delete(id string) error {
	refresh()
	return c.Delete(id)
}

func (c *Client) Delete(id string) error {
	return c.B.Call("DELETE", "/coupons/"+id, c.Token, nil, nil)
}

// List returns a list of coupons.
// For more details see https://stripe.com/docs/api#list_coupons.
func List(params *CouponListParams) (*CouponList, error) {
	refresh()
	return c.List(params)
}

func (c *Client) List(params *CouponListParams) (*CouponList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		params.AppendTo(body)
	}

	list := &CouponList{}
	err := c.B.Call("GET", "/coupons", c.Token, body, list)

	return list, err
}

func refresh() {
	if c == nil {
		c = &Client{B: GetBackend()}
	}

	c.Token = Key
}
