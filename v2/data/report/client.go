//
//
// File generated from our OpenAPI spec
//
//

// Package report provides the report related APIs
package report

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke report related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves metadata about a specific `Report`, including its name, description, and
// the parameters it accepts. It's useful for understanding the capabilities and
// requirements of a particular `Report` before requesting a `ReportRun`.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2DataReportParams) (*stripe.V2DataReport, error) {
	path := stripe.FormatURLPath("/v2/data/reports/%s", id)
	report := &stripe.V2DataReport{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, report)
	return report, err
}

// Returns a list of Stripe-defined reports that the caller can create a `ReportRun` for.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2DataReportListParams) stripe.Seq2[*stripe.V2DataReport, error] {
	if listParams == nil {
		listParams = &stripe.V2DataReportListParams{}
	}
	return stripe.NewV2List("/v2/data/reports", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2DataReport], error) {
		page := &stripe.V2Page[*stripe.V2DataReport]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
