//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v86/form"
)

// v1ReportingReportTypeService is used to invoke /v1/reporting/report_types APIs.
type v1ReportingReportTypeService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Report Type. (Certain report types require a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).)
func (c v1ReportingReportTypeService) Retrieve(ctx context.Context, id string, params *ReportingReportTypeRetrieveParams) (*ReportingReportType, error) {
	if params == nil {
		params = &ReportingReportTypeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/reporting/report_types/%s", id)
	reporttype := &ReportingReportType{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reporttype)
	return reporttype, err
}

// Returns a full list of Report Types.
func (c v1ReportingReportTypeService) List(ctx context.Context, listParams *ReportingReportTypeListParams) *V1List[*ReportingReportType] {
	if listParams == nil {
		listParams = &ReportingReportTypeListParams{}
	}
	p := listParams.GetParams()
	p.Context = ctx
	queryParams := &form.Values{}
	form.AppendTo(queryParams, listParams)
	return newV1List(ctx, nil, func(ctx context.Context, _ *Params, _ *form.Values) (*v1Page[*ReportingReportType], error) {
		list := &v1Page[*ReportingReportType]{}
		err := c.B.CallRaw(http.MethodGet, "/v1/reporting/report_types", c.Key, []byte(queryParams.Encode()), p, list)
		return list, err
	})
}
