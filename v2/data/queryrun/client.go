//
//
// File generated from our OpenAPI spec
//
//

// Package queryrun provides the queryrun related APIs
package queryrun

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke queryrun related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Submits a SQL query for execution against a dataset and returns a `QueryRun` object
// to track progress and retrieve results.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2DataQueryRunParams) (*stripe.V2DataQueryRun, error) {
	queryrun := &stripe.V2DataQueryRun{}
	err := c.B.Call(
		http.MethodPost, "/v2/data/query_runs", c.Key, params, queryrun)
	return queryrun, err
}

// Retrieves the status and results of a previously created `QueryRun`. Returns a `404` error
// if the `QueryRun` doesn't exist or its retention period has ended.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2DataQueryRunParams) (*stripe.V2DataQueryRun, error) {
	path := stripe.FormatURLPath("/v2/data/query_runs/%s", id)
	queryrun := &stripe.V2DataQueryRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, queryrun)
	return queryrun, err
}

// Returns a list of `QueryRun` objects created through the API, most recent first. Listed
// `QueryRun` objects don't include `result`. Retrieve a `QueryRun` to access its result.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2DataQueryRunListParams) stripe.Seq2[*stripe.V2DataQueryRun, error] {
	if listParams == nil {
		listParams = &stripe.V2DataQueryRunListParams{}
	}
	return stripe.NewV2List("/v2/data/query_runs", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2DataQueryRun], error) {
		page := &stripe.V2Page[*stripe.V2DataQueryRun]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
