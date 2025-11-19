//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1PaymentAttemptRecordService is used to invoke /v1/payment_attempt_records APIs.
type v1PaymentAttemptRecordService struct {
	B   Backend
	Key string
}

// Retrieves a Payment Attempt Record with the given ID
func (c v1PaymentAttemptRecordService) Retrieve(ctx context.Context, id string, params *PaymentAttemptRecordRetrieveParams) (*PaymentAttemptRecord, error) {
	if params == nil {
		params = &PaymentAttemptRecordRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_attempt_records/%s", id)
	paymentattemptrecord := &PaymentAttemptRecord{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// List all the Payment Attempt Records attached to the specified Payment Record.
func (c v1PaymentAttemptRecordService) List(ctx context.Context, listParams *PaymentAttemptRecordListParams) Seq2[*PaymentAttemptRecord, error] {
	if listParams == nil {
		listParams = &PaymentAttemptRecordListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*PaymentAttemptRecord], error) {
		list := &v1Page[*PaymentAttemptRecord]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_attempt_records", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
