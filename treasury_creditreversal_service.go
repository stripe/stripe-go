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

// v1TreasuryCreditReversalService is used to invoke /v1/treasury/credit_reversals APIs.
type v1TreasuryCreditReversalService struct {
	B   Backend
	Key string
}

// Reverses a ReceivedCredit and creates a CreditReversal object.
func (c v1TreasuryCreditReversalService) Create(ctx context.Context, params *TreasuryCreditReversalCreateParams) (*TreasuryCreditReversal, error) {
	if params == nil {
		params = &TreasuryCreditReversalCreateParams{}
	}
	params.Context = ctx
	creditreversal := &TreasuryCreditReversal{}
	err := c.B.Call(
		http.MethodPost, "/v1/treasury/credit_reversals", c.Key, params, creditreversal)
	return creditreversal, err
}

// Retrieves the details of an existing CreditReversal by passing the unique CreditReversal ID from either the CreditReversal creation request or CreditReversal list
func (c v1TreasuryCreditReversalService) Retrieve(ctx context.Context, id string, params *TreasuryCreditReversalRetrieveParams) (*TreasuryCreditReversal, error) {
	if params == nil {
		params = &TreasuryCreditReversalRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/credit_reversals/%s", id)
	creditreversal := &TreasuryCreditReversal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, creditreversal)
	return creditreversal, err
}

// Returns a list of CreditReversals.
func (c v1TreasuryCreditReversalService) List(ctx context.Context, listParams *TreasuryCreditReversalListParams) Seq2[*TreasuryCreditReversal, error] {
	if listParams == nil {
		listParams = &TreasuryCreditReversalListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TreasuryCreditReversal, ListContainer, error) {
		list := &TreasuryCreditReversalList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/treasury/credit_reversals", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
