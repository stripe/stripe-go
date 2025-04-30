//
//
// File generated from our OpenAPI spec
//
//

// Package redactionjob provides the /v1/privacy/redaction_jobs APIs
package redactionjob

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/privacy/redaction_jobs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create redaction job method
func New(params *stripe.PrivacyRedactionJobParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().New(params)
}

// Create redaction job method
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

// Retrieve redaction job method
func Get(id string, params *stripe.PrivacyRedactionJobParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().Get(id, params)
}

// Retrieve redaction job method
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

// Update redaction job method
func Update(id string, params *stripe.PrivacyRedactionJobParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().Update(id, params)
}

// Update redaction job method
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

// Cancel redaction job method
func Cancel(id string, params *stripe.PrivacyRedactionJobCancelParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().Cancel(id, params)
}

// Cancel redaction job method
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

// Run redaction job method
func Run(id string, params *stripe.PrivacyRedactionJobRunParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().Run(id, params)
}

// Run redaction job method
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

// Validate redaction job method
func Validate(id string, params *stripe.PrivacyRedactionJobValidateParams) (*stripe.PrivacyRedactionJob, error) {
	return getC().Validate(id, params)
}

// Validate redaction job method
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

// List redaction jobs method...
func List(params *stripe.PrivacyRedactionJobListParams) *Iter {
	return getC().List(params)
}

// List redaction jobs method...
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
