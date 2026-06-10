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

	"github.com/stripe/stripe-go/v86/form"
)

// v1ProductService is used to invoke /v1/products APIs.
type v1ProductService struct {
	B   Backend
	Key string
}

// Creates a new product object.
func (c v1ProductService) Create(ctx context.Context, params *ProductCreateParams) (*Product, error) {
	if params == nil {
		params = &ProductCreateParams{}
	}
	params.Context = ctx
	product := &Product{}
	err := c.B.Call(http.MethodPost, "/v1/products", c.Key, params, product)
	return product, err
}

// Retrieves the details of an existing product. Supply the unique product ID from either a product creation request or the product list, and Stripe will return the corresponding product information.
func (c v1ProductService) Retrieve(ctx context.Context, id string, params *ProductRetrieveParams) (*Product, error) {
	if params == nil {
		params = &ProductRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/products/%s", id)
	product := &Product{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, product)
	return product, err
}

// Updates the specific product by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1ProductService) Update(ctx context.Context, id string, params *ProductUpdateParams) (*Product, error) {
	if params == nil {
		params = &ProductUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/products/%s", id)
	product := &Product{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, product)
	return product, err
}

// Delete a product. Deleting a product is only possible if it has no prices associated with it. Additionally, deleting a product with type=good is only possible if it has no SKUs associated with it.
func (c v1ProductService) Delete(ctx context.Context, id string, params *ProductDeleteParams) (*Product, error) {
	if params == nil {
		params = &ProductDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/products/%s", id)
	product := &Product{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, product)
	return product, err
}

// Serializes a Product create request into a batch job JSONL line.
func (c v1ProductService) MarshalBatchCreate(params *ProductCreateParams) (string, error) {
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

// Serializes a Product delete request into a batch job JSONL line.
func (c v1ProductService) MarshalBatchDelete(id string, params *ProductDeleteParams) (string, error) {
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
		PathParams:    map[string]string{"id": id},
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

// Serializes a Product update request into a batch job JSONL line.
func (c v1ProductService) MarshalBatchUpdate(id string, params *ProductUpdateParams) (string, error) {
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
		PathParams:    map[string]string{"id": id},
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

// Returns a list of your products. The products are returned sorted by creation date, with the most recently created products appearing first.
func (c v1ProductService) List(ctx context.Context, listParams *ProductListParams) *V1List[*Product] {
	if listParams == nil {
		listParams = &ProductListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*Product], error) {
		list := &v1Page[*Product]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/products", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}

// Search for products you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func (c v1ProductService) Search(ctx context.Context, params *ProductSearchParams) *V1SearchList[*Product] {
	if params == nil {
		params = &ProductSearchParams{}
	}
	params.Context = ctx
	return newV1SearchList(ctx, params, func(ctx context.Context, p *Params, b *form.Values) (*v1SearchPage[*Product], error) {
		list := &v1SearchPage[*Product]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/products/search", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
