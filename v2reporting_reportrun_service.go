//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2ReportingReportRunService is used to invoke reportrun related APIs.
type v2ReportingReportRunService struct {
	B   Backend
	Key string
}

// Initiates the generation of a `ReportRun` based on the specified report template
// and user-provided parameters. It's the starting point for obtaining report data,
// and returns a `ReportRun` object which can be used to track the progress and retrieve
// the results of the report.
func (c v2ReportingReportRunService) Create(ctx context.Context, params *V2ReportingReportRunCreateParams) (*V2ReportingReportRun, error) {
	if params == nil {
		params = &V2ReportingReportRunCreateParams{}
	}
	params.Context = ctx
	reportrun := &V2ReportingReportRun{}
	err := c.B.Call(
		http.MethodPost, "/v2/reporting/report_runs", c.Key, params, reportrun)
	return reportrun, err
}

// Fetches the current state and details of a previously created `ReportRun`. If the `ReportRun`
// has succeeded, the endpoint will provide details for how to retrieve the results.
func (c v2ReportingReportRunService) Retrieve(ctx context.Context, id string, params *V2ReportingReportRunRetrieveParams) (*V2ReportingReportRun, error) {
	if params == nil {
		params = &V2ReportingReportRunRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/reporting/report_runs/%s", id)
	reportrun := &V2ReportingReportRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reportrun)
	return reportrun, err
}
