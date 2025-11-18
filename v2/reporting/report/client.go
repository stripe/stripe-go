//
//
// File generated from our OpenAPI spec
//
//

// Package report provides the report related APIs
package report

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke report related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves metadata about a specific `Report` template, including its name, description,
// and the parameters it accepts. It's useful for understanding the capabilities and
// requirements of a particular `Report` before requesting a `ReportRun`.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2ReportingReportParams) (*stripe.V2ReportingReport, error) {
	path := stripe.FormatURLPath("/v2/reporting/reports/%s", id)
	report := &stripe.V2ReportingReport{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, report)
	return report, err
}
