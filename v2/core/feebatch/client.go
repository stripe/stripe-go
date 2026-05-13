//
//
// File generated from our OpenAPI spec
//
//

// Package feebatch provides the feebatch related APIs
package feebatch

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke feebatch related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a FeeBatch by id.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreFeeBatchParams) (*stripe.V2CoreFeeBatch, error) {
	path := stripe.FormatURLPath("/v2/core/fee_batches/%s", id)
	feebatch := &stripe.V2CoreFeeBatch{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, feebatch)
	return feebatch, err
}

// List FeeBatches optionally filtered by collection_record.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreFeeBatchListParams) stripe.Seq2[*stripe.V2CoreFeeBatch, error] {
	if listParams == nil {
		listParams = &stripe.V2CoreFeeBatchListParams{}
	}
	return stripe.NewV2List("/v2/core/fee_batches", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2CoreFeeBatch], error) {
		page := &stripe.V2Page[*stripe.V2CoreFeeBatch]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
