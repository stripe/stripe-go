//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1CouponService is used to invoke /v1/coupons APIs.
type v1CouponService struct {
	B   Backend
	Key string
}

// You can create coupons easily via the [coupon management](https://dashboard.stripe.com/coupons) page of the Stripe dashboard. Coupon creation is also accessible via the API if you need to create coupons on the fly.
//
// A coupon has either a percent_off or an amount_off and currency. If you set an amount_off, that amount will be subtracted from any invoice's subtotal. For example, an invoice with a subtotal of 100 will have a final total of 0 if a coupon with an amount_off of 200 is applied to it and an invoice with a subtotal of 300 will have a final total of 100 if a coupon with an amount_off of 200 is applied to it.
func (c v1CouponService) Create(ctx context.Context, params *CouponCreateParams) (*Coupon, error) {
	if params == nil {
		params = &CouponCreateParams{}
	}
	params.Context = ctx
	coupon := &Coupon{}
	err := c.B.Call(http.MethodPost, "/v1/coupons", c.Key, params, coupon)
	return coupon, err
}

// Retrieves the coupon with the given ID.
func (c v1CouponService) Retrieve(ctx context.Context, id string, params *CouponRetrieveParams) (*Coupon, error) {
	if params == nil {
		params = &CouponRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/coupons/%s", id)
	coupon := &Coupon{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, coupon)
	return coupon, err
}

// Updates the metadata of a coupon. Other coupon details (currency, duration, amount_off) are, by design, not editable.
func (c v1CouponService) Update(ctx context.Context, id string, params *CouponUpdateParams) (*Coupon, error) {
	if params == nil {
		params = &CouponUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/coupons/%s", id)
	coupon := &Coupon{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, coupon)
	return coupon, err
}

// You can delete coupons via the [coupon management](https://dashboard.stripe.com/coupons) page of the Stripe dashboard. However, deleting a coupon does not affect any customers who have already applied the coupon; it means that new customers can't redeem the coupon. You can also delete coupons via the API.
func (c v1CouponService) Delete(ctx context.Context, id string, params *CouponDeleteParams) (*Coupon, error) {
	if params == nil {
		params = &CouponDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/coupons/%s", id)
	coupon := &Coupon{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, coupon)
	return coupon, err
}

// Returns a list of your coupons.
func (c v1CouponService) List(ctx context.Context, listParams *CouponListParams) Seq2[*Coupon, error] {
	if listParams == nil {
		listParams = &CouponListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Coupon, ListContainer, error) {
		list := &CouponList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/coupons", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
