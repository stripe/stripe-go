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

// v1EntitlementsActiveEntitlementService is used to invoke /v1/entitlements/active_entitlements APIs.
type v1EntitlementsActiveEntitlementService struct {
	B   Backend
	Key string
}

// Retrieve an active entitlement
func (c v1EntitlementsActiveEntitlementService) Retrieve(ctx context.Context, id string, params *EntitlementsActiveEntitlementRetrieveParams) (*EntitlementsActiveEntitlement, error) {
	if params == nil {
		params = &EntitlementsActiveEntitlementRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/entitlements/active_entitlements/%s", id)
	activeentitlement := &EntitlementsActiveEntitlement{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, activeentitlement)
	return activeentitlement, err
}

// Retrieve a list of active entitlements for a customer
func (c v1EntitlementsActiveEntitlementService) List(ctx context.Context, listParams *EntitlementsActiveEntitlementListParams) Seq2[*EntitlementsActiveEntitlement, error] {
	if listParams == nil {
		listParams = &EntitlementsActiveEntitlementListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*EntitlementsActiveEntitlement, ListContainer, error) {
		list := &EntitlementsActiveEntitlementList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/entitlements/active_entitlements", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
