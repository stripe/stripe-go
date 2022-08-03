//
//
// File generated from our OpenAPI spec
//
//

// Package promotioncode provides the /promotion_codes APIs
package promotioncode

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /promotion_codes APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new promotion code.
func New(params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	return getC().New(params)
}

// New creates a new promotion code.
func (c Client) New(params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	promotioncode := &stripe.PromotionCode{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/promotion_codes",
		c.Key,
		params,
		promotioncode,
	)
	return promotioncode, err
}

// Get returns the details of a promotion code.
func Get(id string, params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	return getC().Get(id, params)
}

// Get returns the details of a promotion code.
func (c Client) Get(id string, params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	path := stripe.FormatURLPath("/v1/promotion_codes/%s", id)
	promotioncode := &stripe.PromotionCode{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, promotioncode)
	return promotioncode, err
}

// Update updates a promotion code's properties.
func Update(id string, params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	return getC().Update(id, params)
}

// Update updates a promotion code's properties.
func (c Client) Update(id string, params *stripe.PromotionCodeParams) (*stripe.PromotionCode, error) {
	path := stripe.FormatURLPath("/v1/promotion_codes/%s", id)
	promotioncode := &stripe.PromotionCode{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, promotioncode)
	return promotioncode, err
}

// List returns a list of promotion codes.
func List(params *stripe.PromotionCodeListParams) *Iter {
	return getC().List(params)
}

// List returns a list of promotion codes.
func (c Client) List(listParams *stripe.PromotionCodeListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PromotionCodeList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/promotion_codes", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for promotion codes.
type Iter struct {
	*stripe.Iter
}

// PromotionCode returns the promotion code which the iterator is currently pointing to.
func (i *Iter) PromotionCode() *stripe.PromotionCode {
	return i.Current().(*stripe.PromotionCode)
}

// PromotionCodeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PromotionCodeList() *stripe.PromotionCodeList {
	return i.List().(*stripe.PromotionCodeList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
