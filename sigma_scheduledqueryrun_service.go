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

// v1SigmaScheduledQueryRunService is used to invoke /v1/sigma/scheduled_query_runs APIs.
type v1SigmaScheduledQueryRunService struct {
	B   Backend
	Key string
}

// Retrieves the details of an scheduled query run.
func (c v1SigmaScheduledQueryRunService) Retrieve(ctx context.Context, id string, params *SigmaScheduledQueryRunRetrieveParams) (*SigmaScheduledQueryRun, error) {
	if params == nil {
		params = &SigmaScheduledQueryRunRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/sigma/scheduled_query_runs/%s", id)
	scheduledqueryrun := &SigmaScheduledQueryRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, scheduledqueryrun)
	return scheduledqueryrun, err
}

// Returns a list of scheduled query runs.
func (c v1SigmaScheduledQueryRunService) List(ctx context.Context, listParams *SigmaScheduledQueryRunListParams) Seq2[*SigmaScheduledQueryRun, error] {
	if listParams == nil {
		listParams = &SigmaScheduledQueryRunListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*SigmaScheduledQueryRun, ListContainer, error) {
		list := &SigmaScheduledQueryRunList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/sigma/scheduled_query_runs", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
