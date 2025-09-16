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

// v1TaxTransactionService is used to invoke /v1/tax/transactions APIs.
type v1TaxTransactionService struct {
	B   Backend
	Key string
}

// Retrieves a Tax Transaction object.
func (c v1TaxTransactionService) Retrieve(ctx context.Context, id string, params *TaxTransactionRetrieveParams) (*TaxTransaction, error) {
	if params == nil {
		params = &TaxTransactionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax/transactions/%s", id)
	transaction := &TaxTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// Creates a Tax Transaction from a calculation, if that calculation hasn't expired. Calculations expire after 90 days.
func (c v1TaxTransactionService) CreateFromCalculation(ctx context.Context, params *TaxTransactionCreateFromCalculationParams) (*TaxTransaction, error) {
	if params == nil {
		params = &TaxTransactionCreateFromCalculationParams{}
	}
	params.Context = ctx
	transaction := &TaxTransaction{}
	err := c.B.Call(
		http.MethodPost, "/v1/tax/transactions/create_from_calculation", c.Key, params, transaction)
	return transaction, err
}

// Partially or fully reverses a previously created Transaction.
func (c v1TaxTransactionService) CreateReversal(ctx context.Context, params *TaxTransactionCreateReversalParams) (*TaxTransaction, error) {
	if params == nil {
		params = &TaxTransactionCreateReversalParams{}
	}
	params.Context = ctx
	transaction := &TaxTransaction{}
	err := c.B.Call(
		http.MethodPost, "/v1/tax/transactions/create_reversal", c.Key, params, transaction)
	return transaction, err
}

// Retrieves the line items of a committed standalone transaction as a collection.
func (c v1TaxTransactionService) ListLineItems(ctx context.Context, listParams *TaxTransactionListLineItemsParams) Seq2[*TaxTransactionLineItem, error] {
	if listParams == nil {
		listParams = &TaxTransactionListLineItemsParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/tax/transactions/%s/line_items", StringValue(listParams.Transaction))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TaxTransactionLineItem, ListContainer, error) {
		list := &TaxTransactionLineItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
