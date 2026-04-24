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

// v2DataReportingQueryRunService is used to invoke queryrun related APIs.
type v2DataReportingQueryRunService struct {
	B   Backend
	Key string
}

// Creates a query run to execute ad-hoc SQL and returns a `QueryRun` object to track progress and retrieve results.
func (c v2DataReportingQueryRunService) Create(ctx context.Context, params *V2DataReportingQueryRunCreateParams) (*V2DataReportingQueryRun, error) {
	if params == nil {
		params = &V2DataReportingQueryRunCreateParams{}
	}
	params.Context = ctx
	queryrun := &V2DataReportingQueryRun{}
	err := c.B.Call(
		http.MethodPost, "/v2/data/reporting/query_runs", c.Key, params, queryrun)
	return queryrun, err
}

// Fetches the current state and details of a previously created `QueryRun`. If the `QueryRun`
// has succeeded, the endpoint will provide details for how to retrieve the results.
func (c v2DataReportingQueryRunService) Retrieve(ctx context.Context, id string, params *V2DataReportingQueryRunRetrieveParams) (*V2DataReportingQueryRun, error) {
	if params == nil {
		params = &V2DataReportingQueryRunRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/data/reporting/query_runs/%s", id)
	queryrun := &V2DataReportingQueryRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, queryrun)
	return queryrun, err
}
