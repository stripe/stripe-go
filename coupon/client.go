// Package coupon provides the /coupons APIs
package coupon

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /coupons APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new coupon.
func New(params *stripe.CouponParams) (*stripe.Coupon, error) {
	return getC().New(params)
}

// New creates a new coupon.
func (c Client) New(params *stripe.CouponParams) (*stripe.Coupon, error) {
	coupon := &stripe.Coupon{}
	err := c.B.Call(http.MethodPost, "/coupons", c.Key, params, coupon)
	return coupon, err
}

// Get returns the details of a coupon.
func Get(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	return getC().Get(id, params)
}

// Get returns the details of a coupon.
func (c Client) Get(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	path := stripe.FormatURLPath("/coupons/%s", id)
	coupon := &stripe.Coupon{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, coupon)
	return coupon, err
}

// Update updates a coupon's properties.
func Update(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	return getC().Update(id, params)
}

// Update updates a coupon's properties.
func (c Client) Update(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	path := stripe.FormatURLPath("/coupons/%s", id)
	coupon := &stripe.Coupon{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, coupon)
	return coupon, err
}

// Del removes a coupon.
func Del(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	return getC().Del(id, params)
}

// Del removes a coupon.
func (c Client) Del(id string, params *stripe.CouponParams) (*stripe.Coupon, error) {
	path := stripe.FormatURLPath("/coupons/%s", id)
	coupon := &stripe.Coupon{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, coupon)
	return coupon, err
}

// List returns a list of coupons.
func List(params *stripe.CouponListParams) *Iter {
	return getC().List(params)
}

// List returns a list of coupons.
func (c Client) List(listParams *stripe.CouponListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.CouponList{}
		err := c.B.CallRaw(http.MethodGet, "/coupons", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
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
