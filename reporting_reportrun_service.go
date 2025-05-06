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

// v1ReportingReportRunService is used to invoke /v1/reporting/report_runs APIs.
type v1ReportingReportRunService struct {
	B   Backend
	Key string
}

// Creates a new object and begin running the report. (Certain report types require a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).)
func (c v1ReportingReportRunService) Create(ctx context.Context, params *ReportingReportRunCreateParams) (*ReportingReportRun, error) {
	if params == nil {
		params = &ReportingReportRunCreateParams{}
	}
	params.Context = ctx
	reportrun := &ReportingReportRun{}
	err := c.B.Call(
		http.MethodPost, "/v1/reporting/report_runs", c.Key, params, reportrun)
	return reportrun, err
}

// Retrieves the details of an existing Report Run.
func (c v1ReportingReportRunService) Retrieve(ctx context.Context, id string, params *ReportingReportRunRetrieveParams) (*ReportingReportRun, error) {
	if params == nil {
		params = &ReportingReportRunRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/reporting/report_runs/%s", id)
	reportrun := &ReportingReportRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reportrun)
	return reportrun, err
}

// Returns a list of Report Runs, with the most recent appearing first.
func (c v1ReportingReportRunService) List(ctx context.Context, listParams *ReportingReportRunListParams) Seq2[*ReportingReportRun, error] {
	if listParams == nil {
		listParams = &ReportingReportRunListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*ReportingReportRun, ListContainer, error) {
		list := &ReportingReportRunList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/reporting/report_runs", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
