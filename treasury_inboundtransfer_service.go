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

// v1TreasuryInboundTransferService is used to invoke /v1/treasury/inbound_transfers APIs.
type v1TreasuryInboundTransferService struct {
	B   Backend
	Key string
}

// Creates an InboundTransfer.
func (c v1TreasuryInboundTransferService) Create(ctx context.Context, params *TreasuryInboundTransferCreateParams) (*TreasuryInboundTransfer, error) {
	if params == nil {
		params = &TreasuryInboundTransferCreateParams{}
	}
	params.Context = ctx
	inboundtransfer := &TreasuryInboundTransfer{}
	err := c.B.Call(
		http.MethodPost, "/v1/treasury/inbound_transfers", c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Retrieves the details of an existing InboundTransfer.
func (c v1TreasuryInboundTransferService) Retrieve(ctx context.Context, id string, params *TreasuryInboundTransferRetrieveParams) (*TreasuryInboundTransfer, error) {
	if params == nil {
		params = &TreasuryInboundTransferRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/inbound_transfers/%s", id)
	inboundtransfer := &TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Cancels an InboundTransfer.
func (c v1TreasuryInboundTransferService) Cancel(ctx context.Context, id string, params *TreasuryInboundTransferCancelParams) (*TreasuryInboundTransfer, error) {
	if params == nil {
		params = &TreasuryInboundTransferCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/inbound_transfers/%s/cancel", id)
	inboundtransfer := &TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Returns a list of InboundTransfers sent from the specified FinancialAccount.
func (c v1TreasuryInboundTransferService) List(ctx context.Context, listParams *TreasuryInboundTransferListParams) Seq2[*TreasuryInboundTransfer, error] {
	if listParams == nil {
		listParams = &TreasuryInboundTransferListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TreasuryInboundTransfer, ListContainer, error) {
		list := &TreasuryInboundTransferList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/treasury/inbound_transfers", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
