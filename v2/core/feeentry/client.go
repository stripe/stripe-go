//
//
// File generated from our OpenAPI spec
//
//

// Package feeentry provides the feeentry related APIs
package feeentry

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke feeentry related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a FeeEntry by id.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreFeeEntryParams) (*stripe.V2CoreFeeEntry, error) {
	path := stripe.FormatURLPath("/v2/core/fee_entries/%s", id)
	feeentry := &stripe.V2CoreFeeEntry{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, feeentry)
	return feeentry, err
}

// List FeeEntries optionally filtered by incurred_by, fee_batch, or collection_record (at most one filter at a time).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreFeeEntryListParams) stripe.Seq2[*stripe.V2CoreFeeEntry, error] {
	if listParams == nil {
		listParams = &stripe.V2CoreFeeEntryListParams{}
	}
	return stripe.NewV2List("/v2/core/fee_entries", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2CoreFeeEntry], error) {
		page := &stripe.V2Page[*stripe.V2CoreFeeEntry]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
