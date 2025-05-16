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

// v1TreasuryReceivedDebitService is used to invoke /v1/treasury/received_debits APIs.
type v1TreasuryReceivedDebitService struct {
	B   Backend
	Key string
}

// Retrieves the details of an existing ReceivedDebit by passing the unique ReceivedDebit ID from the ReceivedDebit list
func (c v1TreasuryReceivedDebitService) Retrieve(ctx context.Context, id string, params *TreasuryReceivedDebitRetrieveParams) (*TreasuryReceivedDebit, error) {
	if params == nil {
		params = &TreasuryReceivedDebitRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/received_debits/%s", id)
	receiveddebit := &TreasuryReceivedDebit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receiveddebit)
	return receiveddebit, err
}

// Returns a list of ReceivedDebits.
func (c v1TreasuryReceivedDebitService) List(ctx context.Context, listParams *TreasuryReceivedDebitListParams) Seq2[*TreasuryReceivedDebit, error] {
	if listParams == nil {
		listParams = &TreasuryReceivedDebitListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TreasuryReceivedDebit, ListContainer, error) {
		list := &TreasuryReceivedDebitList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/treasury/received_debits", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
