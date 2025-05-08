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

// v1IdentityVerificationReportService is used to invoke /v1/identity/verification_reports APIs.
type v1IdentityVerificationReportService struct {
	B   Backend
	Key string
}

// Retrieves an existing VerificationReport
func (c v1IdentityVerificationReportService) Retrieve(ctx context.Context, id string, params *IdentityVerificationReportRetrieveParams) (*IdentityVerificationReport, error) {
	if params == nil {
		params = &IdentityVerificationReportRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/identity/verification_reports/%s", id)
	verificationreport := &IdentityVerificationReport{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, verificationreport)
	return verificationreport, err
}

// List all verification reports.
func (c v1IdentityVerificationReportService) List(ctx context.Context, listParams *IdentityVerificationReportListParams) Seq2[*IdentityVerificationReport, error] {
	if listParams == nil {
		listParams = &IdentityVerificationReportListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*IdentityVerificationReport, ListContainer, error) {
		list := &IdentityVerificationReportList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/identity/verification_reports", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
