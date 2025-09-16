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

// v1BillingCreditGrantService is used to invoke /v1/billing/credit_grants APIs.
type v1BillingCreditGrantService struct {
	B   Backend
	Key string
}

// Creates a credit grant.
func (c v1BillingCreditGrantService) Create(ctx context.Context, params *BillingCreditGrantCreateParams) (*BillingCreditGrant, error) {
	if params == nil {
		params = &BillingCreditGrantCreateParams{}
	}
	params.Context = ctx
	creditgrant := &BillingCreditGrant{}
	err := c.B.Call(
		http.MethodPost, "/v1/billing/credit_grants", c.Key, params, creditgrant)
	return creditgrant, err
}

// Retrieves a credit grant.
func (c v1BillingCreditGrantService) Retrieve(ctx context.Context, id string, params *BillingCreditGrantRetrieveParams) (*BillingCreditGrant, error) {
	if params == nil {
		params = &BillingCreditGrantRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/credit_grants/%s", id)
	creditgrant := &BillingCreditGrant{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, creditgrant)
	return creditgrant, err
}

// Updates a credit grant.
func (c v1BillingCreditGrantService) Update(ctx context.Context, id string, params *BillingCreditGrantUpdateParams) (*BillingCreditGrant, error) {
	if params == nil {
		params = &BillingCreditGrantUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/credit_grants/%s", id)
	creditgrant := &BillingCreditGrant{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, creditgrant)
	return creditgrant, err
}

// Expires a credit grant.
func (c v1BillingCreditGrantService) Expire(ctx context.Context, id string, params *BillingCreditGrantExpireParams) (*BillingCreditGrant, error) {
	if params == nil {
		params = &BillingCreditGrantExpireParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/credit_grants/%s/expire", id)
	creditgrant := &BillingCreditGrant{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, creditgrant)
	return creditgrant, err
}

// Voids a credit grant.
func (c v1BillingCreditGrantService) VoidGrant(ctx context.Context, id string, params *BillingCreditGrantVoidGrantParams) (*BillingCreditGrant, error) {
	if params == nil {
		params = &BillingCreditGrantVoidGrantParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/credit_grants/%s/void", id)
	creditgrant := &BillingCreditGrant{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, creditgrant)
	return creditgrant, err
}

// Retrieve a list of credit grants.
func (c v1BillingCreditGrantService) List(ctx context.Context, listParams *BillingCreditGrantListParams) Seq2[*BillingCreditGrant, error] {
	if listParams == nil {
		listParams = &BillingCreditGrantListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*BillingCreditGrant, ListContainer, error) {
		list := &BillingCreditGrantList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/billing/credit_grants", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
