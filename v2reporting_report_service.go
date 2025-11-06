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

// v2ReportingReportService is used to invoke report related APIs.
type v2ReportingReportService struct {
	B   Backend
	Key string
}

// Retrieves metadata about a specific `Report` template, including its name, description,
// and the parameters it accepts. It's useful for understanding the capabilities and
// requirements of a particular `Report` before requesting a `ReportRun`.
func (c v2ReportingReportService) Retrieve(ctx context.Context, id string, params *V2ReportingReportRetrieveParams) (*V2ReportingReport, error) {
	if params == nil {
		params = &V2ReportingReportRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/reporting/reports/%s", id)
	report := &V2ReportingReport{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, report)
	return report, err
}
