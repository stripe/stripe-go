//
//
// File generated from our OpenAPI spec
//
//

// Package queryrun provides the queryrun related APIs
package queryrun

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke queryrun related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a query run to execute ad-hoc SQL and returns a `QueryRun` object to track progress and retrieve results.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2DataReportingQueryRunParams) (*stripe.V2DataReportingQueryRun, error) {
	queryrun := &stripe.V2DataReportingQueryRun{}
	err := c.B.Call(
		http.MethodPost, "/v2/data/reporting/query_runs", c.Key, params, queryrun)
	return queryrun, err
}

// Fetches the current state and details of a previously created `QueryRun`. If the `QueryRun`
// has succeeded, the endpoint will provide details for how to retrieve the results.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2DataReportingQueryRunParams) (*stripe.V2DataReportingQueryRun, error) {
	path := stripe.FormatURLPath("/v2/data/reporting/query_runs/%s", id)
	queryrun := &stripe.V2DataReportingQueryRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, queryrun)
	return queryrun, err
}
