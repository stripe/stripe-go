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

// v2MoneyManagementOutboundPaymentQuoteService is used to invoke outboundpaymentquote related APIs.
type v2MoneyManagementOutboundPaymentQuoteService struct {
	B   Backend
	Key string
}

// Creates an OutboundPaymentQuote usable in an OutboundPayment.
func (c v2MoneyManagementOutboundPaymentQuoteService) Create(ctx context.Context, params *V2MoneyManagementOutboundPaymentQuoteCreateParams) (*V2MoneyManagementOutboundPaymentQuote, error) {
	outboundpaymentquote := &V2MoneyManagementOutboundPaymentQuote{}
	if params == nil {
		params = &V2MoneyManagementOutboundPaymentQuoteCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/outbound_payment_quotes", c.Key, params, outboundpaymentquote)
	return outboundpaymentquote, err
}

// Retrieves the details of an existing OutboundPaymentQuote by passing the unique OutboundPaymentQuote ID.
func (c v2MoneyManagementOutboundPaymentQuoteService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementOutboundPaymentQuoteRetrieveParams) (*V2MoneyManagementOutboundPaymentQuote, error) {
	path := FormatURLPath("/v2/money_management/outbound_payment_quotes/%s", id)
	outboundpaymentquote := &V2MoneyManagementOutboundPaymentQuote{}
	if params == nil {
		params = &V2MoneyManagementOutboundPaymentQuoteRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundpaymentquote)
	return outboundpaymentquote, err
}
