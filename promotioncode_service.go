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

// v1PromotionCodeService is used to invoke /v1/promotion_codes APIs.
type v1PromotionCodeService struct {
	B   Backend
	Key string
}

// A promotion code points to a coupon. You can optionally restrict the code to a specific customer, redemption limit, and expiration date.
func (c v1PromotionCodeService) Create(ctx context.Context, params *PromotionCodeCreateParams) (*PromotionCode, error) {
	if params == nil {
		params = &PromotionCodeCreateParams{}
	}
	params.Context = ctx
	promotioncode := &PromotionCode{}
	err := c.B.Call(
		http.MethodPost, "/v1/promotion_codes", c.Key, params, promotioncode)
	return promotioncode, err
}

// Retrieves the promotion code with the given ID. In order to retrieve a promotion code by the customer-facing code use [list](https://docs.stripe.com/docs/api/promotion_codes/list) with the desired code.
func (c v1PromotionCodeService) Retrieve(ctx context.Context, id string, params *PromotionCodeRetrieveParams) (*PromotionCode, error) {
	if params == nil {
		params = &PromotionCodeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/promotion_codes/%s", id)
	promotioncode := &PromotionCode{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, promotioncode)
	return promotioncode, err
}

// Updates the specified promotion code by setting the values of the parameters passed. Most fields are, by design, not editable.
func (c v1PromotionCodeService) Update(ctx context.Context, id string, params *PromotionCodeUpdateParams) (*PromotionCode, error) {
	if params == nil {
		params = &PromotionCodeUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/promotion_codes/%s", id)
	promotioncode := &PromotionCode{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, promotioncode)
	return promotioncode, err
}

// Returns a list of your promotion codes.
func (c v1PromotionCodeService) List(ctx context.Context, listParams *PromotionCodeListParams) Seq2[*PromotionCode, error] {
	if listParams == nil {
		listParams = &PromotionCodeListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*PromotionCode, ListContainer, error) {
		list := &PromotionCodeList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/promotion_codes", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
