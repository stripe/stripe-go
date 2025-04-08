//
//
// File generated from our OpenAPI spec
//
//

// Package redactionjobvalidationerror provides the /v1/privacy/redaction_jobs/{job}/validation_errors APIs
package redactionjobvalidationerror

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/privacy/redaction_jobs/{job}/validation_errors APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve validation error method
func Get(id string, params *stripe.PrivacyRedactionJobValidationErrorParams) (*stripe.PrivacyRedactionJobValidationError, error) {
	return getC().Get(id, params)
}

// Retrieve validation error method
func (c Client) Get(id string, params *stripe.PrivacyRedactionJobValidationErrorParams) (*stripe.PrivacyRedactionJobValidationError, error) {
	path := stripe.FormatURLPath(
		"/v1/privacy/redaction_jobs/%s/validation_errors/%s", stripe.StringValue(
			params.Job), id)
	redactionjobvalidationerror := &stripe.PrivacyRedactionJobValidationError{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, redactionjobvalidationerror)
	return redactionjobvalidationerror, err
}

// List validation errors method
func List(params *stripe.PrivacyRedactionJobValidationErrorListParams) *Iter {
	return getC().List(params)
}

// List validation errors method
func (c Client) List(listParams *stripe.PrivacyRedactionJobValidationErrorListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/privacy/redaction_jobs/%s/validation_errors", stripe.StringValue(
			listParams.Job))
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PrivacyRedactionJobValidationErrorList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for privacy redaction job validation errors.
type Iter struct {
	*stripe.Iter
}

// PrivacyRedactionJobValidationError returns the privacy redaction job validation error which the iterator is currently pointing to.
func (i *Iter) PrivacyRedactionJobValidationError() *stripe.PrivacyRedactionJobValidationError {
	return i.Current().(*stripe.PrivacyRedactionJobValidationError)
}

// PrivacyRedactionJobValidationErrorList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PrivacyRedactionJobValidationErrorList() *stripe.PrivacyRedactionJobValidationErrorList {
	return i.List().(*stripe.PrivacyRedactionJobValidationErrorList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
