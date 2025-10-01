//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v83/form"
)

// v1PrivacyRedactionJobService is used to invoke /v1/privacy/redaction_jobs APIs.
type v1PrivacyRedactionJobService struct {
	B   Backend
	Key string
}

// Creates a redaction job. When a job is created, it will start to validate.
func (c v1PrivacyRedactionJobService) Create(ctx context.Context, params *PrivacyRedactionJobCreateParams) (*PrivacyRedactionJob, error) {
	if params == nil {
		params = &PrivacyRedactionJobCreateParams{}
	}
	params.Context = ctx
	redactionjob := &PrivacyRedactionJob{}
	err := c.B.Call(
		http.MethodPost, "/v1/privacy/redaction_jobs", c.Key, params, redactionjob)
	return redactionjob, err
}

// Retrieves the details of a previously created redaction job.
func (c v1PrivacyRedactionJobService) Retrieve(ctx context.Context, id string, params *PrivacyRedactionJobRetrieveParams) (*PrivacyRedactionJob, error) {
	if params == nil {
		params = &PrivacyRedactionJobRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/privacy/redaction_jobs/%s", id)
	redactionjob := &PrivacyRedactionJob{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Updates the properties of a redaction job without running or canceling the job.
//
// If the job to update is in a failed status, it will not automatically start to validate. Once you applied all of the changes, use the validate API to start validation again.
func (c v1PrivacyRedactionJobService) Update(ctx context.Context, id string, params *PrivacyRedactionJobUpdateParams) (*PrivacyRedactionJob, error) {
	if params == nil {
		params = &PrivacyRedactionJobUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/privacy/redaction_jobs/%s", id)
	redactionjob := &PrivacyRedactionJob{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// You can cancel a redaction job when it's in one of these statuses: ready, failed.
//
// Canceling the redaction job will abandon its attempt to redact the configured objects. A canceled job cannot be used again.
func (c v1PrivacyRedactionJobService) Cancel(ctx context.Context, id string, params *PrivacyRedactionJobCancelParams) (*PrivacyRedactionJob, error) {
	if params == nil {
		params = &PrivacyRedactionJobCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/privacy/redaction_jobs/%s/cancel", id)
	redactionjob := &PrivacyRedactionJob{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Run a redaction job in a ready status.
//
// When you run a job, the configured objects will be redacted asynchronously. This action is irreversible and cannot be canceled once started.
//
// The status of the job will move to redacting. Once all of the objects are redacted, the status will become succeeded. If the job's validation_behavior is set to fix, the automatic fixes will be applied to objects at this step.
func (c v1PrivacyRedactionJobService) Run(ctx context.Context, id string, params *PrivacyRedactionJobRunParams) (*PrivacyRedactionJob, error) {
	if params == nil {
		params = &PrivacyRedactionJobRunParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/privacy/redaction_jobs/%s/run", id)
	redactionjob := &PrivacyRedactionJob{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Validate a redaction job when it is in a failed status.
//
// When a job is created, it automatically begins to validate on the configured objects' eligibility for redaction. Use this to validate the job again after its validation errors are resolved or the job's validation_behavior is changed.
//
// The status of the job will move to validating. Once all of the objects are validated, the status of the job will become ready. If there are any validation errors preventing the job from running, the status will become failed.
func (c v1PrivacyRedactionJobService) Validate(ctx context.Context, id string, params *PrivacyRedactionJobValidateParams) (*PrivacyRedactionJob, error) {
	if params == nil {
		params = &PrivacyRedactionJobValidateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/privacy/redaction_jobs/%s/validate", id)
	redactionjob := &PrivacyRedactionJob{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Returns a list of redaction jobs.
func (c v1PrivacyRedactionJobService) List(ctx context.Context, listParams *PrivacyRedactionJobListParams) Seq2[*PrivacyRedactionJob, error] {
	if listParams == nil {
		listParams = &PrivacyRedactionJobListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*PrivacyRedactionJob], error) {
		list := &v1Page[*PrivacyRedactionJob]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/privacy/redaction_jobs", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
