// Package coupon provides the /coupons APIs
package coupon

import (
	"net/url"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	Forever   stripe.CouponDuration = "forever"
	Once      stripe.CouponDuration = "once"
	Repeating stripe.CouponDuration = "repeating"
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
	body := &form.Values{}
	form.AppendTo(body, params)

	coupon := &stripe.Coupon{}
	err := c.B.Call("POST", "/coupons", c.Key, body, &params.Params, coupon)

	return coupon, err
}

// Get returns the details of a coupon.
// For more details see https://stripe.com/docs/api#retrieve_coupon.
func Get(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	coupon := &stripe.Coupon{}

	err := c.B.Call("GET", "/coupons/"+url.QueryEscape(id), c.Key, body, commonParams, coupon)

	return coupon, err
}

// Update updates a coupon's properties.
// For more details see https://stripe.com/docs/api#update_coupon.
func Update(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	coupon := &stripe.Coupon{}
	err := c.B.Call("POST", "/coupons/"+url.QueryEscape(id), c.Key, body, &params.Params, coupon)

	return coupon, err
}

// Del removes a coupon.
// For more details see https://stripe.com/docs/api#delete_coupon.
func Del(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	coupon := &stripe.Coupon{}
	err := c.B.Call("DELETE", "/coupons/"+url.QueryEscape(id), c.Key, body, commonParams, coupon)

	return coupon, err
}

// List returns a list of coupons.
// For more details see https://stripe.com/docs/api#list_coupons.
func List(params *stripe.CouponListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.CouponListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.CouponList{}
		err := c.B.Call("GET", "/coupons", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Coupons.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Coupon returns the most recent Coupon
// visited by a call to Next.
func (i *Iter) Coupon() *stripe.Coupon {
	return i.Current().(*stripe.Coupon)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
