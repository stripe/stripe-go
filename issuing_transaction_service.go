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

// v1IssuingTransactionService is used to invoke /v1/issuing/transactions APIs.
type v1IssuingTransactionService struct {
	B   Backend
	Key string
}

// Retrieves an Issuing Transaction object.
func (c v1IssuingTransactionService) Retrieve(ctx context.Context, id string, params *IssuingTransactionRetrieveParams) (*IssuingTransaction, error) {
	if params == nil {
		params = &IssuingTransactionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/transactions/%s", id)
	transaction := &IssuingTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// Updates the specified Issuing Transaction object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1IssuingTransactionService) Update(ctx context.Context, id string, params *IssuingTransactionUpdateParams) (*IssuingTransaction, error) {
	if params == nil {
		params = &IssuingTransactionUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/transactions/%s", id)
	transaction := &IssuingTransaction{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}

// Returns a list of Issuing Transaction objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingTransactionService) List(ctx context.Context, listParams *IssuingTransactionListParams) Seq2[*IssuingTransaction, error] {
	if listParams == nil {
		listParams = &IssuingTransactionListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*IssuingTransaction, ListContainer, error) {
		list := &IssuingTransactionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/transactions", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
