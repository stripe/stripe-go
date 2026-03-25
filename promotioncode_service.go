//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1PromotionCodeService is used to invoke /v1/promotion_codes APIs.
type v1PromotionCodeService struct {
	B   Backend
	Key string
}

// A promotion code points to an underlying promotion. You can optionally restrict the code to a specific customer, redemption limit, and expiration date.
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

// Serializes a PromotionCode create request into a batch job JSONL line.
func (c v1PromotionCodeService) MarshalBatchCreate(params *PromotionCodeCreateParams) (string, error) {
	itemID, err := newUUID4()
	if err != nil {
		return "", err
	}

	item := struct {
		ID            string            `json:"id"`
		Context       string            `json:"context,omitempty"`
		StripeVersion string            `json:"stripe_version,omitempty"`
		PathParams    map[string]string `json:"path_params,omitempty"`
		Params        interface{}       `json:"params"`
	}{
		ID:            itemID,
		PathParams:    nil,
		StripeVersion: APIVersion,
	}
	if params != nil {
		item.Params = params
		if params.StripeContext != nil {
			item.Context = *params.StripeContext
		}
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Serializes a PromotionCode update request into a batch job JSONL line.
func (c v1PromotionCodeService) MarshalBatchUpdate(id string, params *PromotionCodeUpdateParams) (string, error) {
	itemID, err := newUUID4()
	if err != nil {
		return "", err
	}

	item := struct {
		ID            string            `json:"id"`
		Context       string            `json:"context,omitempty"`
		StripeVersion string            `json:"stripe_version,omitempty"`
		PathParams    map[string]string `json:"path_params,omitempty"`
		Params        interface{}       `json:"params"`
	}{
		ID:            itemID,
		PathParams:    map[string]string{"promotion_code": id},
		StripeVersion: APIVersion,
	}
	if params != nil {
		item.Params = params
		if params.StripeContext != nil {
			item.Context = *params.StripeContext
		}
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Returns a list of your promotion codes.
func (c v1PromotionCodeService) List(ctx context.Context, listParams *PromotionCodeListParams) *V1List[*PromotionCode] {
	if listParams == nil {
		listParams = &PromotionCodeListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*PromotionCode], error) {
		list := &v1Page[*PromotionCode]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/promotion_codes", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
