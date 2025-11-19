//
//
// File generated from our OpenAPI spec
//
//

// Package redactionjob provides the /v1/privacy/redaction_jobs APIs
package redactionjob

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/form"
)

// Client is used to invoke /v1/privacy/redaction_jobs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a redaction job. When a job is created, it will start to validate.
func New(params *stripe.PrivacyRedactionJobParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().New(params)
}

// Creates a redaction job. When a job is created, it will start to validate.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.PrivacyRedactionJobParams) (*stripe.PrivacyRedactionJob, error) {
	redactionjob := &stripe.PrivacyRedactionJob{}
	err := c.B.Call(
		http.MethodPost, "/v1/privacy/redaction_jobs", c.Key, params, redactionjob)
	return redactionjob, err
}

// Retrieves the details of a previously created redaction job.
func Get(id string, params *stripe.PrivacyRedactionJobParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().Get(id, params)
}

// Retrieves the details of a previously created redaction job.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.PrivacyRedactionJobParams) (*stripe.PrivacyRedactionJob, error) {
	path := stripe.FormatURLPath("/v1/privacy/redaction_jobs/%s", id)
	redactionjob := &stripe.PrivacyRedactionJob{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Updates the properties of a redaction job without running or canceling the job.
//
// If the job to update is in a failed status, it will not automatically start to validate. Once you applied all of the changes, use the validate API to start validation again.
func Update(id string, params *stripe.PrivacyRedactionJobParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().Update(id, params)
}

// Updates the properties of a redaction job without running or canceling the job.
//
// If the job to update is in a failed status, it will not automatically start to validate. Once you applied all of the changes, use the validate API to start validation again.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PrivacyRedactionJobParams) (*stripe.PrivacyRedactionJob, error) {
	path := stripe.FormatURLPath("/v1/privacy/redaction_jobs/%s", id)
	redactionjob := &stripe.PrivacyRedactionJob{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// You can cancel a redaction job when it's in one of these statuses: ready, failed.
//
// Canceling the redaction job will abandon its attempt to redact the configured objects. A canceled job cannot be used again.
func Cancel(id string, params *stripe.PrivacyRedactionJobCancelParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().Cancel(id, params)
}

// You can cancel a redaction job when it's in one of these statuses: ready, failed.
//
// Canceling the redaction job will abandon its attempt to redact the configured objects. A canceled job cannot be used again.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.PrivacyRedactionJobCancelParams) (*stripe.PrivacyRedactionJob, error) {
	path := stripe.FormatURLPath("/v1/privacy/redaction_jobs/%s/cancel", id)
	redactionjob := &stripe.PrivacyRedactionJob{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Run a redaction job in a ready status.
//
// When you run a job, the configured objects will be redacted asynchronously. This action is irreversible and cannot be canceled once started.
//
// The status of the job will move to redacting. Once all of the objects are redacted, the status will become succeeded. If the job's validation_behavior is set to fix, the automatic fixes will be applied to objects at this step.
func Run(id string, params *stripe.PrivacyRedactionJobRunParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().Run(id, params)
}

// Run a redaction job in a ready status.
//
// When you run a job, the configured objects will be redacted asynchronously. This action is irreversible and cannot be canceled once started.
//
// The status of the job will move to redacting. Once all of the objects are redacted, the status will become succeeded. If the job's validation_behavior is set to fix, the automatic fixes will be applied to objects at this step.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Run(id string, params *stripe.PrivacyRedactionJobRunParams) (*stripe.PrivacyRedactionJob, error) {
	path := stripe.FormatURLPath("/v1/privacy/redaction_jobs/%s/run", id)
	redactionjob := &stripe.PrivacyRedactionJob{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Validate a redaction job when it is in a failed status.
//
// When a job is created, it automatically begins to validate on the configured objects' eligibility for redaction. Use this to validate the job again after its validation errors are resolved or the job's validation_behavior is changed.
//
// The status of the job will move to validating. Once all of the objects are validated, the status of the job will become ready. If there are any validation errors preventing the job from running, the status will become failed.
func Validate(id string, params *stripe.PrivacyRedactionJobValidateParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().Validate(id, params)
}

// Validate a redaction job when it is in a failed status.
//
// When a job is created, it automatically begins to validate on the configured objects' eligibility for redaction. Use this to validate the job again after its validation errors are resolved or the job's validation_behavior is changed.
//
// The status of the job will move to validating. Once all of the objects are validated, the status of the job will become ready. If there are any validation errors preventing the job from running, the status will become failed.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Validate(id string, params *stripe.PrivacyRedactionJobValidateParams) (*stripe.PrivacyRedactionJob, error) {
	path := stripe.FormatURLPath("/v1/privacy/redaction_jobs/%s/validate", id)
	redactionjob := &stripe.PrivacyRedactionJob{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, redactionjob)
	return redactionjob, err
}

// Returns a list of redaction jobs.
func List(params *stripe.PrivacyRedactionJobListParams) *Iter {
	return getC().List(params)
}

// Returns a list of redaction jobs.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PrivacyRedactionJobListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PrivacyRedactionJobList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/privacy/redaction_jobs", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for privacy redaction jobs.
type Iter struct {
	*stripe.Iter
}

// PrivacyRedactionJob returns the privacy redaction job which the iterator is currently pointing to.
func (i *Iter) PrivacyRedactionJob() *stripe.PrivacyRedactionJob {
	return i.Current().(*stripe.PrivacyRedactionJob)
}

// PrivacyRedactionJobList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PrivacyRedactionJobList() *stripe.PrivacyRedactionJobList {
	return i.List().(*stripe.PrivacyRedactionJobList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
