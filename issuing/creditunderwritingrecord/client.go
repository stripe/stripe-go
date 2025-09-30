//
//
// File generated from our OpenAPI spec
//
//

// Package creditunderwritingrecord provides the /v1/issuing/credit_underwriting_records APIs
package creditunderwritingrecord

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
	"github.com/stripe/stripe-go/v83/form"
)

// Client is used to invoke /v1/issuing/credit_underwriting_records APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a CreditUnderwritingRecord object.
func Get(id string, params *stripe.IssuingCreditUnderwritingRecordParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	return getC().Get(id, params)
}

// Retrieves a CreditUnderwritingRecord object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.IssuingCreditUnderwritingRecordParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	path := stripe.FormatURLPath("/v1/issuing/credit_underwriting_records/%s", id)
	creditunderwritingrecord := &stripe.IssuingCreditUnderwritingRecord{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Update a CreditUnderwritingRecord object to correct mistakes.
func Correct(id string, params *stripe.IssuingCreditUnderwritingRecordCorrectParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	return getC().Correct(id, params)
}

// Update a CreditUnderwritingRecord object to correct mistakes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Correct(id string, params *stripe.IssuingCreditUnderwritingRecordCorrectParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/issuing/credit_underwriting_records/%s/correct", id)
	creditunderwritingrecord := &stripe.IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Creates a CreditUnderwritingRecord object with information about a credit application submission.
func CreateFromApplication(params *stripe.IssuingCreditUnderwritingRecordCreateFromApplicationParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	return getC().CreateFromApplication(params)
}

// Creates a CreditUnderwritingRecord object with information about a credit application submission.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CreateFromApplication(params *stripe.IssuingCreditUnderwritingRecordCreateFromApplicationParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	creditunderwritingrecord := &stripe.IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/credit_underwriting_records/create_from_application", c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Creates a CreditUnderwritingRecord object from an underwriting decision coming from a proactive review of an existing accountholder.
func CreateFromProactiveReview(params *stripe.IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	return getC().CreateFromProactiveReview(params)
}

// Creates a CreditUnderwritingRecord object from an underwriting decision coming from a proactive review of an existing accountholder.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CreateFromProactiveReview(params *stripe.IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	creditunderwritingrecord := &stripe.IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/credit_underwriting_records/create_from_proactive_review", c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Update a CreditUnderwritingRecord object from a decision made on a credit application.
func ReportDecision(id string, params *stripe.IssuingCreditUnderwritingRecordReportDecisionParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	return getC().ReportDecision(id, params)
}

// Update a CreditUnderwritingRecord object from a decision made on a credit application.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReportDecision(id string, params *stripe.IssuingCreditUnderwritingRecordReportDecisionParams) (*stripe.IssuingCreditUnderwritingRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/issuing/credit_underwriting_records/%s/report_decision", id)
	creditunderwritingrecord := &stripe.IssuingCreditUnderwritingRecord{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, creditunderwritingrecord)
	return creditunderwritingrecord, err
}

// Retrieves a list of CreditUnderwritingRecord objects. The objects are sorted in descending order by creation date, with the most-recently-created object appearing first.
func List(params *stripe.IssuingCreditUnderwritingRecordListParams) *Iter {
	return getC().List(params)
}

// Retrieves a list of CreditUnderwritingRecord objects. The objects are sorted in descending order by creation date, with the most-recently-created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IssuingCreditUnderwritingRecordListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingCreditUnderwritingRecordList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/credit_underwriting_records", c.Key, []byte(b.Encode()), p, list)

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
