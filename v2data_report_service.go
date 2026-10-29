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

// v2DataReportService is used to invoke report related APIs.
type v2DataReportService struct {
	B   Backend
	Key string
}

// Retrieves metadata about a specific `Report`, including its name, description, and
// the parameters it accepts. It's useful for understanding the capabilities and
// requirements of a particular `Report` before requesting a `ReportRun`.
func (c v2DataReportService) Retrieve(ctx context.Context, id string, params *V2DataReportRetrieveParams) (*V2DataReport, error) {
	if params == nil {
		params = &V2DataReportRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/data/reports/%s", id)
	report := &V2DataReport{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, report)
	return report, err
}

// Returns a list of Stripe-defined reports that the caller can create a `ReportRun` for.
func (c v2DataReportService) List(ctx context.Context, listParams *V2DataReportListParams) *V2List[*V2DataReport] {
	if listParams == nil {
		listParams = &V2DataReportListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/data/reports", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2DataReport], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2DataReport]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
