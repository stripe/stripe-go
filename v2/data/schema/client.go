//
//
// File generated from our OpenAPI spec
//
//

// Package schema provides the schema related APIs
package schema

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke schema related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the schema for a particular table.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2DataSchemaParams) (*stripe.V2DataSchema, error) {
	path := stripe.FormatURLPath("/v2/data/schemas/%s", id)
	schema := &stripe.V2DataSchema{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, schema)
	return schema, err
}

// Returns a list of schemas describing the tables available to query.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2DataSchemaListParams) stripe.Seq2[*stripe.V2DataSchema, error] {
	if listParams == nil {
		listParams = &stripe.V2DataSchemaListParams{}
	}
	return stripe.NewV2List("/v2/data/schemas", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2DataSchema], error) {
		page := &stripe.V2Page[*stripe.V2DataSchema]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
