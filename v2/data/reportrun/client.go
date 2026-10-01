//
//
// File generated from our OpenAPI spec
//
//

// Package reportrun provides the reportrun related APIs
package reportrun

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke reportrun related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Initiates the generation of a `ReportRun` based on the specified `Report` and
// caller-provided parameters. Returns a `ReportRun` object which can be used to track
// the progress and retrieve the results of the report.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2DataReportRunParams) (*stripe.V2DataReportRun, error) {
	reportrun := &stripe.V2DataReportRun{}
	err := c.B.Call(
		http.MethodPost, "/v2/data/report_runs", c.Key, params, reportrun)
	return reportrun, err
}

// Fetches the current state and details of a previously created `ReportRun`. If the
// `ReportRun` has succeeded, the endpoint provides details for how to retrieve the results.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2DataReportRunParams) (*stripe.V2DataReportRun, error) {
	path := stripe.FormatURLPath("/v2/data/report_runs/%s", id)
	reportrun := &stripe.V2DataReportRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reportrun)
	return reportrun, err
}
