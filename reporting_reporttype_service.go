//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
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
func (c v1ReportingReportTypeService) List(ctx context.Context, listParams *ReportingReportTypeListParams) Seq2[*ReportingReportType, error] {
	if listParams == nil {
		listParams = &ReportingReportTypeListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*ReportingReportType, ListContainer, error) {
		list := &ReportingReportTypeList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/reporting/report_types", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
