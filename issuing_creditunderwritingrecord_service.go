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

// v1IssuingCreditUnderwritingRecordService is used to invoke /v1/issuing/credit_underwriting_records APIs.
type v1IssuingCreditUnderwritingRecordService struct {
	B   Backend
	Key string
}

// Retrieves a CreditUnderwritingRecord object.
func (c v1IssuingCreditUnderwritingRecordService) Retrieve(ctx context.Context, id string, params *IssuingCreditUnderwritingRecordRetrieveParams) (*IssuingCreditUnderwritingRecord, error) {
	if params == nil {
		params = &IssuingCreditUnderwritingRecordRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/credit_underwriting_records/%s", id)
	creditunderwritingrecord := &IssuingCreditUnderwritingRecord{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Update a CreditUnderwritingRecord object to correct mistakes.
func (c v1IssuingCreditUnderwritingRecordService) Correct(ctx context.Context, id string, params *IssuingCreditUnderwritingRecordCorrectParams) (*IssuingCreditUnderwritingRecord, error) {
	if params == nil {
		params = &IssuingCreditUnderwritingRecordCorrectParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/issuing/credit_underwriting_records/%s/correct", id)
	creditunderwritingrecord := &IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Creates a CreditUnderwritingRecord object with information about a credit application submission.
func (c v1IssuingCreditUnderwritingRecordService) CreateFromApplication(ctx context.Context, params *IssuingCreditUnderwritingRecordCreateFromApplicationParams) (*IssuingCreditUnderwritingRecord, error) {
	if params == nil {
		params = &IssuingCreditUnderwritingRecordCreateFromApplicationParams{}
	}
	params.Context = ctx
	creditunderwritingrecord := &IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/credit_underwriting_records/create_from_application", c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Creates a CreditUnderwritingRecord object from an underwriting decision coming from a proactive review of an existing accountholder.
func (c v1IssuingCreditUnderwritingRecordService) CreateFromProactiveReview(ctx context.Context, params *IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams) (*IssuingCreditUnderwritingRecord, error) {
	if params == nil {
		params = &IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams{}
	}
	params.Context = ctx
	creditunderwritingrecord := &IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/credit_underwriting_records/create_from_proactive_review", c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Update a CreditUnderwritingRecord object from a decision made on a credit application.
func (c v1IssuingCreditUnderwritingRecordService) ReportDecision(ctx context.Context, id string, params *IssuingCreditUnderwritingRecordReportDecisionParams) (*IssuingCreditUnderwritingRecord, error) {
	if params == nil {
		params = &IssuingCreditUnderwritingRecordReportDecisionParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/issuing/credit_underwriting_records/%s/report_decision", id)
	creditunderwritingrecord := &IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Retrieves a list of CreditUnderwritingRecord objects. The objects are sorted in descending order by creation date, with the most-recently-created object appearing first.
func (c v1IssuingCreditUnderwritingRecordService) List(ctx context.Context, listParams *IssuingCreditUnderwritingRecordListParams) Seq2[*IssuingCreditUnderwritingRecord, error] {
	if listParams == nil {
		listParams = &IssuingCreditUnderwritingRecordListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*IssuingCreditUnderwritingRecord, ListContainer, error) {
		list := &IssuingCreditUnderwritingRecordList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/credit_underwriting_records", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
