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

// v1IssuingFraudLiabilityDebitService is used to invoke /v1/issuing/fraud_liability_debits APIs.
type v1IssuingFraudLiabilityDebitService struct {
	B   Backend
	Key string
}

// Retrieves an Issuing FraudLiabilityDebit object.
func (c v1IssuingFraudLiabilityDebitService) Retrieve(ctx context.Context, id string, params *IssuingFraudLiabilityDebitRetrieveParams) (*IssuingFraudLiabilityDebit, error) {
	if params == nil {
		params = &IssuingFraudLiabilityDebitRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/fraud_liability_debits/%s", id)
	fraudliabilitydebit := &IssuingFraudLiabilityDebit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, fraudliabilitydebit)
	return fraudliabilitydebit, err
}

// Returns a list of Issuing FraudLiabilityDebit objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingFraudLiabilityDebitService) List(ctx context.Context, listParams *IssuingFraudLiabilityDebitListParams) Seq2[*IssuingFraudLiabilityDebit, error] {
	if listParams == nil {
		listParams = &IssuingFraudLiabilityDebitListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*IssuingFraudLiabilityDebit], error) {
		list := &v1Page[*IssuingFraudLiabilityDebit]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/fraud_liability_debits", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
