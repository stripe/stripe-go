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

// v1TreasuryReceivedCreditService is used to invoke /v1/treasury/received_credits APIs.
type v1TreasuryReceivedCreditService struct {
	B   Backend
	Key string
}

// Retrieves the details of an existing ReceivedCredit by passing the unique ReceivedCredit ID from the ReceivedCredit list.
func (c v1TreasuryReceivedCreditService) Retrieve(ctx context.Context, id string, params *TreasuryReceivedCreditRetrieveParams) (*TreasuryReceivedCredit, error) {
	if params == nil {
		params = &TreasuryReceivedCreditRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/received_credits/%s", id)
	receivedcredit := &TreasuryReceivedCredit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receivedcredit)
	return receivedcredit, err
}

// Returns a list of ReceivedCredits.
func (c v1TreasuryReceivedCreditService) List(ctx context.Context, listParams *TreasuryReceivedCreditListParams) Seq2[*TreasuryReceivedCredit, error] {
	if listParams == nil {
		listParams = &TreasuryReceivedCreditListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TreasuryReceivedCredit, ListContainer, error) {
		list := &TreasuryReceivedCreditList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/treasury/received_credits", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
