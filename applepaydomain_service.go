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

// v1ApplePayDomainService is used to invoke /v1/apple_pay/domains APIs.
type v1ApplePayDomainService struct {
	B   Backend
	Key string
}

// Create an apple pay domain.
func (c v1ApplePayDomainService) Create(ctx context.Context, params *ApplePayDomainCreateParams) (*ApplePayDomain, error) {
	if params == nil {
		params = &ApplePayDomainCreateParams{}
	}
	params.Context = ctx
	applepaydomain := &ApplePayDomain{}
	err := c.B.Call(
		http.MethodPost, "/v1/apple_pay/domains", c.Key, params, applepaydomain)
	return applepaydomain, err
}

// Retrieve an apple pay domain.
func (c v1ApplePayDomainService) Retrieve(ctx context.Context, id string, params *ApplePayDomainRetrieveParams) (*ApplePayDomain, error) {
	if params == nil {
		params = &ApplePayDomainRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/apple_pay/domains/%s", id)
	applepaydomain := &ApplePayDomain{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, applepaydomain)
	return applepaydomain, err
}

// Delete an apple pay domain.
func (c v1ApplePayDomainService) Delete(ctx context.Context, id string, params *ApplePayDomainDeleteParams) (*ApplePayDomain, error) {
	if params == nil {
		params = &ApplePayDomainDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/apple_pay/domains/%s", id)
	applepaydomain := &ApplePayDomain{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, applepaydomain)
	return applepaydomain, err
}

// List apple pay domains.
func (c v1ApplePayDomainService) List(ctx context.Context, listParams *ApplePayDomainListParams) Seq2[*ApplePayDomain, error] {
	if listParams == nil {
		listParams = &ApplePayDomainListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*ApplePayDomain, ListContainer, error) {
		list := &ApplePayDomainList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/apple_pay/domains", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
