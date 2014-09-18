// package coupon provides the /coupons APIs
package coupon

import (
	"errors"
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /coupons APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new coupons.
// For more details see https://stripe.com/docs/api#create_coupon.
func New(params *stripe.CouponParams) (*stripe.Coupon, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.CouponParams) (*stripe.Coupon, error) {
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

	if params.Duration == stripe.Repeating {
		body.Add("duration_in_months", strconv.FormatUint(params.DurationPeriod, 10))
	}

	if params.Redemptions > 0 {
		body.Add("max_redemptions", strconv.FormatUint(params.Redemptions, 10))
	}

	if params.RedeemBy > 0 {
		body.Add("redeem_by", strconv.FormatInt(params.RedeemBy, 10))
	}

	params.AppendTo(body)

	coupon := &stripe.Coupon{}
	err := c.B.Call("POST", "/coupons", c.Key, body, coupon)

	return coupon, err
}

// Get returns the details of a coupon.
// For more details see https://stripe.com/docs/api#retrieve_coupon.
func Get(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	coupon := &stripe.Coupon{}
	err := c.B.Call("GET", "/coupons/"+id, c.Key, body, coupon)

	return coupon, err
}

// Del removes a coupon.
// For more details see https://stripe.com/docs/api#delete_coupon.
func Del(id string) error {
	return getC().Del(id)
}

func (c Client) Del(id string) error {
	return c.B.Call("DELETE", "/coupons/"+id, c.Key, nil, nil)
}

// List returns a list of coupons.
// For more details see https://stripe.com/docs/api#list_coupons.
func List(params *stripe.CouponListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.CouponListParams) *Iter {
	type couponList struct {
		stripe.ListMeta
		Values []*stripe.Coupon `json:"data"`
	}

	var body *url.Values
	var lp *stripe.ListParams

	if params != nil {
		body = &url.Values{}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &couponList{}
		err := c.B.Call("GET", "/coupons", c.Key, &b, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is a iterator for list responses.
type Iter struct {
	Iter *stripe.Iter
}

// Next returns the next value in the list.
func (i *Iter) Next() (*stripe.Coupon, error) {
	c, err := i.Iter.Next()
	if err != nil {
		return nil, err
	}

	return c.(*stripe.Coupon), err
}

// Stop returns true if there are no more iterations to be performed.
func (i *Iter) Stop() bool {
	return i.Iter.Stop()
}

// Meta returns the list metadata.
func (i *Iter) Meta() *stripe.ListMeta {
	return i.Iter.Meta()
}

func getC() Client {
	return Client{stripe.GetBackend(), stripe.Key}
}
