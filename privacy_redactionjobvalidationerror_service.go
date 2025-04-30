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

// v1PrivacyRedactionJobValidationErrorService is used to invoke /v1/privacy/redaction_jobs/{job}/validation_errors APIs.
type v1PrivacyRedactionJobValidationErrorService struct {
	B   Backend
	Key string
}

// List validation errors method
func (c v1PrivacyRedactionJobValidationErrorService) List(ctx context.Context, listParams *PrivacyRedactionJobValidationErrorListParams) Seq2[*PrivacyRedactionJobValidationError, error] {
	path := FormatURLPath(
		"/v1/privacy/redaction_jobs/%s/validation_errors", StringValue(
			listParams.Job))
	if listParams == nil {
		listParams = &PrivacyRedactionJobValidationErrorListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*PrivacyRedactionJobValidationError, ListContainer, error) {
		list := &PrivacyRedactionJobValidationErrorList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
