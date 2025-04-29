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

// v1PrivacyRedactionJobService is used to invoke /v1/privacy/redaction_jobs APIs.
type v1PrivacyRedactionJobService struct {
	B   Backend
	Key string
}

// Create redaction job method
func (c v1PrivacyRedactionJobService) Create(ctx context.Context, params *PrivacyRedactionJobCreateParams) (*PrivacyRedactionJob, error) {
	redactionjob := &PrivacyRedactionJob{}
	if params == nil {
		params = &PrivacyRedactionJobCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v1/privacy/redaction_jobs", c.Key, params, redactionjob)
	return redactionjob, err
}

// Retrieve redaction job method
func (c v1PrivacyRedactionJobService) Retrieve(ctx context.Context, id string, params *PrivacyRedactionJobRetrieveParams) (*PrivacyRedactionJob, error) {
	path := FormatURLPath("/v1/privacy/redaction_jobs/%s", id)
	redactionjob := &PrivacyRedactionJob{}
	if params == nil {
		params = &PrivacyRedactionJobRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Update redaction job method
func (c v1PrivacyRedactionJobService) Update(ctx context.Context, id string, params *PrivacyRedactionJobUpdateParams) (*PrivacyRedactionJob, error) {
	path := FormatURLPath("/v1/privacy/redaction_jobs/%s", id)
	redactionjob := &PrivacyRedactionJob{}
	if params == nil {
		params = &PrivacyRedactionJobUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Cancel redaction job method
func (c v1PrivacyRedactionJobService) Cancel(ctx context.Context, id string, params *PrivacyRedactionJobCancelParams) (*PrivacyRedactionJob, error) {
	path := FormatURLPath("/v1/privacy/redaction_jobs/%s/cancel", id)
	redactionjob := &PrivacyRedactionJob{}
	if params == nil {
		params = &PrivacyRedactionJobCancelParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Run redaction job method
func (c v1PrivacyRedactionJobService) Run(ctx context.Context, id string, params *PrivacyRedactionJobRunParams) (*PrivacyRedactionJob, error) {
	path := FormatURLPath("/v1/privacy/redaction_jobs/%s/run", id)
	redactionjob := &PrivacyRedactionJob{}
	if params == nil {
		params = &PrivacyRedactionJobRunParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Validate redaction job method
func (c v1PrivacyRedactionJobService) Validate(ctx context.Context, id string, params *PrivacyRedactionJobValidateParams) (*PrivacyRedactionJob, error) {
	path := FormatURLPath("/v1/privacy/redaction_jobs/%s/validate", id)
	redactionjob := &PrivacyRedactionJob{}
	if params == nil {
		params = &PrivacyRedactionJobValidateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// List redaction jobs method...
func (c v1PrivacyRedactionJobService) List(ctx context.Context, listParams *PrivacyRedactionJobListParams) Seq2[*PrivacyRedactionJob, error] {
	if listParams == nil {
		listParams = &PrivacyRedactionJobListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*PrivacyRedactionJob, ListContainer, error) {
		list := &PrivacyRedactionJobList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/privacy/redaction_jobs", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
