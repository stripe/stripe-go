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

// v2DataReportRunService is used to invoke reportrun related APIs.
type v2DataReportRunService struct {
	B   Backend
	Key string
}

// Initiates the generation of a `ReportRun` based on the specified `Report` and
// caller-provided parameters. Returns a `ReportRun` object which can be used to track
// the progress and retrieve the results of the report.
func (c v2DataReportRunService) Create(ctx context.Context, params *V2DataReportRunCreateParams) (*V2DataReportRun, error) {
	if params == nil {
		params = &V2DataReportRunCreateParams{}
	}
	params.Context = ctx
	reportrun := &V2DataReportRun{}
	err := c.B.Call(
		http.MethodPost, "/v2/data/report_runs", c.Key, params, reportrun)
	return reportrun, err
}

// Fetches the current state and details of a previously created `ReportRun`. If the
// `ReportRun` has succeeded, the endpoint provides details for how to retrieve the results.
func (c v2DataReportRunService) Retrieve(ctx context.Context, id string, params *V2DataReportRunRetrieveParams) (*V2DataReportRun, error) {
	if params == nil {
		params = &V2DataReportRunRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/data/report_runs/%s", id)
	reportrun := &V2DataReportRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reportrun)
	return reportrun, err
}
