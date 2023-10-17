//
//
// File generated from our OpenAPI spec
//
//

// Package creditunderwritingrecord provides the /issuing/credit_underwriting_records APIs
package creditunderwritingrecord

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /issuing/credit_underwriting_records APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an issuing credit underwriting record.
func Get(id string, params *stripe.IssuingCreditUnderwritingRecordParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	return getC().Get(id, params)
}

// Get returns the details of an issuing credit underwriting record.
func (c Client) Get(id string, params *stripe.IssuingCreditUnderwritingRecordParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	path := stripe.FormatURLPath("/v1/issuing/credit_underwriting_records/%s", id)
	creditunderwritingrecord := &stripe.IssuingCreditUnderwritingRecord{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Correct is the method for the `POST /v1/issuing/credit_underwriting_records/{credit_underwriting_record}/correct` API.
func Correct(id string, params *stripe.IssuingCreditUnderwritingRecordCorrectParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	return getC().Correct(id, params)
}

// Correct is the method for the `POST /v1/issuing/credit_underwriting_records/{credit_underwriting_record}/correct` API.
func (c Client) Correct(id string, params *stripe.IssuingCreditUnderwritingRecordCorrectParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/issuing/credit_underwriting_records/%s/correct",
		id,
	)
	creditunderwritingrecord := &stripe.IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost,
		path,
		c.Key,
		params,
		creditunderwritingrecord,
	)
	return creditunderwritingrecord, err
}

// CreateFromApplication is the method for the `POST /v1/issuing/credit_underwriting_records/create_from_application` API.
func CreateFromApplication(params *stripe.IssuingCreditUnderwritingRecordCreateFromApplicationParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	return getC().CreateFromApplication(params)
}

// CreateFromApplication is the method for the `POST /v1/issuing/credit_underwriting_records/create_from_application` API.
func (c Client) CreateFromApplication(params *stripe.IssuingCreditUnderwritingRecordCreateFromApplicationParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	creditunderwritingrecord := &stripe.IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/issuing/credit_underwriting_records/create_from_application",
		c.Key,
		params,
		creditunderwritingrecord,
	)
	return creditunderwritingrecord, err
}

// CreateFromProactiveReview is the method for the `POST /v1/issuing/credit_underwriting_records/create_from_proactive_review` API.
func CreateFromProactiveReview(params *stripe.IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	return getC().CreateFromProactiveReview(params)
}

// CreateFromProactiveReview is the method for the `POST /v1/issuing/credit_underwriting_records/create_from_proactive_review` API.
func (c Client) CreateFromProactiveReview(params *stripe.IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	creditunderwritingrecord := &stripe.IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/issuing/credit_underwriting_records/create_from_proactive_review",
		c.Key,
		params,
		creditunderwritingrecord,
	)
	return creditunderwritingrecord, err
}

// ReportDecision is the method for the `POST /v1/issuing/credit_underwriting_records/{credit_underwriting_record}/report_decision` API.
func ReportDecision(id string, params *stripe.IssuingCreditUnderwritingRecordReportDecisionParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	return getC().ReportDecision(id, params)
}

// ReportDecision is the method for the `POST /v1/issuing/credit_underwriting_records/{credit_underwriting_record}/report_decision` API.
func (c Client) ReportDecision(id string, params *stripe.IssuingCreditUnderwritingRecordReportDecisionParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/issuing/credit_underwriting_records/%s/report_decision",
		id,
	)
	creditunderwritingrecord := &stripe.IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost,
		path,
		c.Key,
		params,
		creditunderwritingrecord,
	)
	return creditunderwritingrecord, err
}

// List returns a list of issuing credit underwriting records.
func List(params *stripe.IssuingCreditUnderwritingRecordListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuing credit underwriting records.
func (c Client) List(listParams *stripe.IssuingCreditUnderwritingRecordListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingCreditUnderwritingRecordList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/credit_underwriting_records", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing credit underwriting records.
type Iter struct {
	*stripe.Iter
}

// IssuingCreditUnderwritingRecord returns the issuing credit underwriting record which the iterator is currently pointing to.
func (i *Iter) IssuingCreditUnderwritingRecord() *stripe.IssuingCreditUnderwritingRecord {
	return i.Current().(*stripe.IssuingCreditUnderwritingRecord)
}

// IssuingCreditUnderwritingRecordList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingCreditUnderwritingRecordList() *stripe.IssuingCreditUnderwritingRecordList {
	return i.List().(*stripe.IssuingCreditUnderwritingRecordList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
