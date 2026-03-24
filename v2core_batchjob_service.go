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

// v2CoreBatchJobService is used to invoke batchjob related APIs.
type v2CoreBatchJobService struct {
	B   Backend
	Key string
}

// Creates a new batch job.
func (c v2CoreBatchJobService) Create(ctx context.Context, params *V2CoreBatchJobCreateParams) (*V2CoreBatchJob, error) {
	if params == nil {
		params = &V2CoreBatchJobCreateParams{}
	}
	params.Context = ctx
	batchjob := &V2CoreBatchJob{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/batch_jobs", c.Key, params, batchjob)
	return batchjob, err
}

// Retrieves an existing batch job.
func (c v2CoreBatchJobService) Retrieve(ctx context.Context, id string, params *V2CoreBatchJobRetrieveParams) (*V2CoreBatchJob, error) {
	if params == nil {
		params = &V2CoreBatchJobRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/batch_jobs/%s", id)
	batchjob := &V2CoreBatchJob{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, batchjob)
	return batchjob, err
}

// Cancels an existing batch job.
func (c v2CoreBatchJobService) Cancel(ctx context.Context, id string, params *V2CoreBatchJobCancelParams) (*V2CoreBatchJob, error) {
	if params == nil {
		params = &V2CoreBatchJobCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/batch_jobs/%s/cancel", id)
	batchjob := &V2CoreBatchJob{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, batchjob)
	return batchjob, err
}
