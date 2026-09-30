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

// v2DataQueryRunService is used to invoke queryrun related APIs.
type v2DataQueryRunService struct {
	B   Backend
	Key string
}

// Submits a SQL query for execution against a dataset and returns a `QueryRun` object
// to track progress and retrieve results.
func (c v2DataQueryRunService) Create(ctx context.Context, params *V2DataQueryRunCreateParams) (*V2DataQueryRun, error) {
	if params == nil {
		params = &V2DataQueryRunCreateParams{}
	}
	params.Context = ctx
	queryrun := &V2DataQueryRun{}
	err := c.B.Call(
		http.MethodPost, "/v2/data/query_runs", c.Key, params, queryrun)
	return queryrun, err
}

// Retrieves the status and results of a previously created `QueryRun`.
func (c v2DataQueryRunService) Retrieve(ctx context.Context, id string, params *V2DataQueryRunRetrieveParams) (*V2DataQueryRun, error) {
	if params == nil {
		params = &V2DataQueryRunRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/data/query_runs/%s", id)
	queryrun := &V2DataQueryRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, queryrun)
	return queryrun, err
}
