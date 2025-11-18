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

// v1TreasuryTransactionEntryService is used to invoke /v1/treasury/transaction_entries APIs.
type v1TreasuryTransactionEntryService struct {
	B   Backend
	Key string
}

// Retrieves a TransactionEntry object.
func (c v1TreasuryTransactionEntryService) Retrieve(ctx context.Context, id string, params *TreasuryTransactionEntryRetrieveParams) (*TreasuryTransactionEntry, error) {
	if params == nil {
		params = &TreasuryTransactionEntryRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/transaction_entries/%s", id)
	transactionentry := &TreasuryTransactionEntry{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transactionentry)
	return transactionentry, err
}

// Retrieves a list of TransactionEntry objects.
func (c v1TreasuryTransactionEntryService) List(ctx context.Context, listParams *TreasuryTransactionEntryListParams) Seq2[*TreasuryTransactionEntry, error] {
	if listParams == nil {
		listParams = &TreasuryTransactionEntryListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*TreasuryTransactionEntry], error) {
		list := &v1Page[*TreasuryTransactionEntry]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/treasury/transaction_entries", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
