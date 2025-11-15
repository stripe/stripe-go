//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v83/form"
)

// v1TreasuryOutboundTransferService is used to invoke /v1/treasury/outbound_transfers APIs.
type v1TreasuryOutboundTransferService struct {
	B   Backend
	Key string
}

// Creates an OutboundTransfer.
func (c v1TreasuryOutboundTransferService) Create(ctx context.Context, params *TreasuryOutboundTransferCreateParams) (*TreasuryOutboundTransfer, error) {
	if params == nil {
		params = &TreasuryOutboundTransferCreateParams{}
	}
	params.Context = ctx
	outboundtransfer := &TreasuryOutboundTransfer{}
	err := c.B.Call(
		http.MethodPost, "/v1/treasury/outbound_transfers", c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Retrieves the details of an existing OutboundTransfer by passing the unique OutboundTransfer ID from either the OutboundTransfer creation request or OutboundTransfer list.
func (c v1TreasuryOutboundTransferService) Retrieve(ctx context.Context, id string, params *TreasuryOutboundTransferRetrieveParams) (*TreasuryOutboundTransfer, error) {
	if params == nil {
		params = &TreasuryOutboundTransferRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/outbound_transfers/%s", id)
	outboundtransfer := &TreasuryOutboundTransfer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// An OutboundTransfer can be canceled if the funds have not yet been paid out.
func (c v1TreasuryOutboundTransferService) Cancel(ctx context.Context, id string, params *TreasuryOutboundTransferCancelParams) (*TreasuryOutboundTransfer, error) {
	if params == nil {
		params = &TreasuryOutboundTransferCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/outbound_transfers/%s/cancel", id)
	outboundtransfer := &TreasuryOutboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Returns a list of OutboundTransfers sent from the specified FinancialAccount.
func (c v1TreasuryOutboundTransferService) List(ctx context.Context, listParams *TreasuryOutboundTransferListParams) *V1List[*TreasuryOutboundTransfer] {
	if listParams == nil {
		listParams = &TreasuryOutboundTransferListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*TreasuryOutboundTransfer], error) {
		list := &v1Page[*TreasuryOutboundTransfer]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/treasury/outbound_transfers", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
