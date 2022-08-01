//
//
// File generated from our OpenAPI spec
//
//

// Package reportrun provides the /reporting/report_runs APIs
package reportrun

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /reporting/report_runs APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new reporting report run.
func New(params *stripe.ReportingReportRunParams) (*stripe.ReportingReportRun, error) {
	return getC().New(params)
}

// New creates a new reporting report run.
func (c Client) New(params *stripe.ReportingReportRunParams) (*stripe.ReportingReportRun, error) {
	reportrun := &stripe.ReportingReportRun{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/reporting/report_runs",
		c.Key,
		params,
		reportrun,
	)
	return reportrun, err
}

// Get returns the details of a reporting report run.
func Get(id string, params *stripe.ReportingReportRunParams) (*stripe.ReportingReportRun, error) {
	return getC().Get(id, params)
}

// Get returns the details of a reporting report run.
func (c Client) Get(id string, params *stripe.ReportingReportRunParams) (*stripe.ReportingReportRun, error) {
	path := stripe.FormatURLPath("/v1/reporting/report_runs/%s", id)
	reportrun := &stripe.ReportingReportRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reportrun)
	return reportrun, err
}

// List returns a list of reporting report runs.
func List(params *stripe.ReportingReportRunListParams) *Iter {
	return getC().List(params)
}

// List returns a list of reporting report runs.
func (c Client) List(listParams *stripe.ReportingReportRunListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ReportingReportRunList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/reporting/report_runs", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for reporting report runs.
type Iter struct {
	*stripe.Iter
}

// ReportingReportRun returns the reporting report run which the iterator is currently pointing to.
func (i *Iter) ReportingReportRun() *stripe.ReportingReportRun {
	return i.Current().(*stripe.ReportingReportRun)
}

// ReportingReportRunList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ReportingReportRunList() *stripe.ReportingReportRunList {
	return i.List().(*stripe.ReportingReportRunList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
