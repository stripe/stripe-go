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

// v1TreasuryOutboundPaymentService is used to invoke /v1/treasury/outbound_payments APIs.
type v1TreasuryOutboundPaymentService struct {
	B   Backend
	Key string
}

// Creates an OutboundPayment.
func (c v1TreasuryOutboundPaymentService) Create(ctx context.Context, params *TreasuryOutboundPaymentCreateParams) (*TreasuryOutboundPayment, error) {
	if params == nil {
		params = &TreasuryOutboundPaymentCreateParams{}
	}
	params.Context = ctx
	outboundpayment := &TreasuryOutboundPayment{}
	err := c.B.Call(
		http.MethodPost, "/v1/treasury/outbound_payments", c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Retrieves the details of an existing OutboundPayment by passing the unique OutboundPayment ID from either the OutboundPayment creation request or OutboundPayment list.
func (c v1TreasuryOutboundPaymentService) Retrieve(ctx context.Context, id string, params *TreasuryOutboundPaymentRetrieveParams) (*TreasuryOutboundPayment, error) {
	if params == nil {
		params = &TreasuryOutboundPaymentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/outbound_payments/%s", id)
	outboundpayment := &TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Cancel an OutboundPayment.
func (c v1TreasuryOutboundPaymentService) Cancel(ctx context.Context, id string, params *TreasuryOutboundPaymentCancelParams) (*TreasuryOutboundPayment, error) {
	if params == nil {
		params = &TreasuryOutboundPaymentCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/outbound_payments/%s/cancel", id)
	outboundpayment := &TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Returns a list of OutboundPayments sent from the specified FinancialAccount.
func (c v1TreasuryOutboundPaymentService) List(ctx context.Context, listParams *TreasuryOutboundPaymentListParams) Seq2[*TreasuryOutboundPayment, error] {
	if listParams == nil {
		listParams = &TreasuryOutboundPaymentListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TreasuryOutboundPayment, ListContainer, error) {
		list := &TreasuryOutboundPaymentList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/treasury/outbound_payments", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
