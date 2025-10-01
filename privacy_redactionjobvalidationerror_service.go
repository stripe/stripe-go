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

// v1PrivacyRedactionJobValidationErrorService is used to invoke /v1/privacy/redaction_jobs/{job}/validation_errors APIs.
type v1PrivacyRedactionJobValidationErrorService struct {
	B   Backend
	Key string
}

// Returns a list of validation errors for the specified redaction job.
func (c v1PrivacyRedactionJobValidationErrorService) List(ctx context.Context, listParams *PrivacyRedactionJobValidationErrorListParams) Seq2[*PrivacyRedactionJobValidationError, error] {
	if listParams == nil {
		listParams = &PrivacyRedactionJobValidationErrorListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/privacy/redaction_jobs/%s/validation_errors", StringValue(
			listParams.Job))
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*PrivacyRedactionJobValidationError], error) {
		list := &v1Page[*PrivacyRedactionJobValidationError]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
