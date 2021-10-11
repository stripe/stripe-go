//
//
// File generated from our OpenAPI spec
//
//

// Package reporttype provides the /reporting/report_types APIs
package reporttype

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /reporting/report_types APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a report type.
func Get(id string, params *stripe.ReportTypeParams) (*stripe.ReportType, error) {
	return getC().Get(id, params)
}

// Get returns the details of a report type.
func (c Client) Get(id string, params *stripe.ReportTypeParams) (*stripe.ReportType, error) {
	path := stripe.FormatURLPath("/v1/reporting/report_types/%s", id)
	reporttype := &stripe.ReportType{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reporttype)
	return reporttype, err
}

// List returns a list of report types.
func List(params *stripe.ReportTypeListParams) *Iter {
	return getC().List(params)
}

// List returns a list of report types.
func (c Client) List(listParams *stripe.ReportTypeListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ReportTypeList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/reporting/report_types", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for report types.
type Iter struct {
	*stripe.Iter
}

// ReportType returns the report type which the iterator is currently pointing to.
func (i *Iter) ReportType() *stripe.ReportType {
	return i.Current().(*stripe.ReportType)
}

// ReportTypeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ReportTypeList() *stripe.ReportTypeList {
	return i.List().(*stripe.ReportTypeList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
