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

// v1ProductFeatureService is used to invoke /v1/products/{product}/features APIs.
type v1ProductFeatureService struct {
	B   Backend
	Key string
}

// Creates a product_feature, which represents a feature attachment to a product
func (c v1ProductFeatureService) Create(ctx context.Context, params *ProductFeatureCreateParams) (*ProductFeature, error) {
	if params == nil {
		params = &ProductFeatureCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/products/%s/features", StringValue(params.Product))
	productfeature := &ProductFeature{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, productfeature)
	return productfeature, err
}

// Retrieves a product_feature, which represents a feature attachment to a product
func (c v1ProductFeatureService) Retrieve(ctx context.Context, id string, params *ProductFeatureRetrieveParams) (*ProductFeature, error) {
	if params == nil {
		params = &ProductFeatureRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/products/%s/features/%s", StringValue(params.Product), id)
	productfeature := &ProductFeature{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, productfeature)
	return productfeature, err
}

// Deletes the feature attachment to a product
func (c v1ProductFeatureService) Delete(ctx context.Context, id string, params *ProductFeatureDeleteParams) (*ProductFeature, error) {
	if params == nil {
		params = &ProductFeatureDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/products/%s/features/%s", StringValue(params.Product), id)
	productfeature := &ProductFeature{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, productfeature)
	return productfeature, err
}

// Retrieve a list of features for a product
func (c v1ProductFeatureService) List(ctx context.Context, listParams *ProductFeatureListParams) Seq2[*ProductFeature, error] {
	if listParams == nil {
		listParams = &ProductFeatureListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/products/%s/features", StringValue(listParams.Product))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*ProductFeature, ListContainer, error) {
		list := &ProductFeatureList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
