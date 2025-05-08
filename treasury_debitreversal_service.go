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

// v1TreasuryDebitReversalService is used to invoke /v1/treasury/debit_reversals APIs.
type v1TreasuryDebitReversalService struct {
	B   Backend
	Key string
}

// Reverses a ReceivedDebit and creates a DebitReversal object.
func (c v1TreasuryDebitReversalService) Create(ctx context.Context, params *TreasuryDebitReversalCreateParams) (*TreasuryDebitReversal, error) {
	if params == nil {
		params = &TreasuryDebitReversalCreateParams{}
	}
	params.Context = ctx
	debitreversal := &TreasuryDebitReversal{}
	err := c.B.Call(
		http.MethodPost, "/v1/treasury/debit_reversals", c.Key, params, debitreversal)
	return debitreversal, err
}

// Retrieves a DebitReversal object.
func (c v1TreasuryDebitReversalService) Retrieve(ctx context.Context, id string, params *TreasuryDebitReversalRetrieveParams) (*TreasuryDebitReversal, error) {
	if params == nil {
		params = &TreasuryDebitReversalRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/debit_reversals/%s", id)
	debitreversal := &TreasuryDebitReversal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, debitreversal)
	return debitreversal, err
}

// Returns a list of DebitReversals.
func (c v1TreasuryDebitReversalService) List(ctx context.Context, listParams *TreasuryDebitReversalListParams) Seq2[*TreasuryDebitReversal, error] {
	if listParams == nil {
		listParams = &TreasuryDebitReversalListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TreasuryDebitReversal, ListContainer, error) {
		list := &TreasuryDebitReversalList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/treasury/debit_reversals", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
