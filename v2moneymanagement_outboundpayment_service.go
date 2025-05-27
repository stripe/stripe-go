//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2MoneyManagementOutboundPaymentService is used to invoke outboundpayment related APIs.
type v2MoneyManagementOutboundPaymentService struct {
	B   Backend
	Key string
}

// Creates an OutboundPayment.
func (c v2MoneyManagementOutboundPaymentService) Create(ctx context.Context, params *V2MoneyManagementOutboundPaymentCreateParams) (*V2MoneyManagementOutboundPayment, error) {
	if params == nil {
		params = &V2MoneyManagementOutboundPaymentCreateParams{}
	}
	params.Context = ctx
	outboundpayment := &V2MoneyManagementOutboundPayment{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/outbound_payments", c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Retrieves the details of an existing OutboundPayment by passing the unique OutboundPayment ID from either the OutboundPayment create or list response.
func (c v2MoneyManagementOutboundPaymentService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementOutboundPaymentRetrieveParams) (*V2MoneyManagementOutboundPayment, error) {
	if params == nil {
		params = &V2MoneyManagementOutboundPaymentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/outbound_payments/%s", id)
	outboundpayment := &V2MoneyManagementOutboundPayment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Cancels an OutboundPayment. Only processing OutboundPayments can be canceled.
func (c v2MoneyManagementOutboundPaymentService) Cancel(ctx context.Context, id string, params *V2MoneyManagementOutboundPaymentCancelParams) (*V2MoneyManagementOutboundPayment, error) {
	if params == nil {
		params = &V2MoneyManagementOutboundPaymentCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/outbound_payments/%s/cancel", id)
	outboundpayment := &V2MoneyManagementOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Returns a list of OutboundPayments that match the provided filters.
func (c v2MoneyManagementOutboundPaymentService) List(ctx context.Context, listParams *V2MoneyManagementOutboundPaymentListParams) Seq2[*V2MoneyManagementOutboundPayment, error] {
	if listParams == nil {
		listParams = &V2MoneyManagementOutboundPaymentListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/money_management/outbound_payments", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementOutboundPayment], error) {
		page := &V2Page[*V2MoneyManagementOutboundPayment]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
