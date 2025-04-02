//
//
// File generated from our OpenAPI spec
//
//

// Package paymentattemptrecord provides the /v1/payment_attempt_records APIs
package paymentattemptrecord

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/payment_attempt_records APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Payment Attempt Record with the given ID
func Get(id string, params *stripe.PaymentAttemptRecordParams) (*stripe.PaymentAttemptRecord, error) {
	return getC().Get(id, params)
}

// Retrieves a Payment Attempt Record with the given ID
func (c Client) Get(id string, params *stripe.PaymentAttemptRecordParams) (*stripe.PaymentAttemptRecord, error) {
	path := stripe.FormatURLPath("/v1/payment_attempt_records/%s", id)
	paymentattemptrecord := &stripe.PaymentAttemptRecord{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentattemptrecord)
	return paymentattemptrecord, err
}

// List all the Payment Attempt Records attached to the specified Payment Record.
func List(params *stripe.PaymentAttemptRecordListParams) *Iter {
	return getC().List(params)
}

// List all the Payment Attempt Records attached to the specified Payment Record.
func (c Client) List(listParams *stripe.PaymentAttemptRecordListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentAttemptRecordList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_attempt_records", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment attempt records.
type Iter struct {
	*stripe.Iter
}

// PaymentAttemptRecord returns the payment attempt record which the iterator is currently pointing to.
func (i *Iter) PaymentAttemptRecord() *stripe.PaymentAttemptRecord {
	return i.Current().(*stripe.PaymentAttemptRecord)
}

// PaymentAttemptRecordList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentAttemptRecordList() *stripe.PaymentAttemptRecordList {
	return i.List().(*stripe.PaymentAttemptRecordList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
