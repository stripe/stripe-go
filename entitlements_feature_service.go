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

// v1EntitlementsFeatureService is used to invoke /v1/entitlements/features APIs.
type v1EntitlementsFeatureService struct {
	B   Backend
	Key string
}

// Creates a feature
func (c v1EntitlementsFeatureService) Create(ctx context.Context, params *EntitlementsFeatureCreateParams) (*EntitlementsFeature, error) {
	if params == nil {
		params = &EntitlementsFeatureCreateParams{}
	}
	params.Context = ctx
	feature := &EntitlementsFeature{}
	err := c.B.Call(
		http.MethodPost, "/v1/entitlements/features", c.Key, params, feature)
	return feature, err
}

// Retrieves a feature
func (c v1EntitlementsFeatureService) Retrieve(ctx context.Context, id string, params *EntitlementsFeatureRetrieveParams) (*EntitlementsFeature, error) {
	if params == nil {
		params = &EntitlementsFeatureRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/entitlements/features/%s", id)
	feature := &EntitlementsFeature{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, feature)
	return feature, err
}

// Update a feature's metadata or permanently deactivate it.
func (c v1EntitlementsFeatureService) Update(ctx context.Context, id string, params *EntitlementsFeatureUpdateParams) (*EntitlementsFeature, error) {
	if params == nil {
		params = &EntitlementsFeatureUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/entitlements/features/%s", id)
	feature := &EntitlementsFeature{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, feature)
	return feature, err
}

// Retrieve a list of features
func (c v1EntitlementsFeatureService) List(ctx context.Context, listParams *EntitlementsFeatureListParams) Seq2[*EntitlementsFeature, error] {
	if listParams == nil {
		listParams = &EntitlementsFeatureListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*EntitlementsFeature, ListContainer, error) {
		list := &EntitlementsFeatureList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/entitlements/features", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
